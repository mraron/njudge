package web

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/web/extmodels"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) StartBackgroundJobs() {
	go s.runUpdateProblems()
	go s.runSyncJudges()
	go s.runGlue()
	go s.runJudger()
}

func (s *Server) runUpdateProblems() {
	for {
		if err := s.ProblemStore.Update(); err != nil {
			log.Print(err)
		}

		time.Sleep(20 * time.Second)
	}
}

func (s *Server) runSyncJudges() {
	loadJudgesFromDB := func() {
		var err error
		s.judges, err = models.Judges().All(s.DB)

		if err != nil {
			panic(err)
		}
	}

	for {
		loadJudgesFromDB()

		for _, j := range s.judges {
			jwt, err := helpers.GetJWT(s.Keys)
			if err != nil {
				log.Print(err)
				continue
			}

			c := judge.NewClient("http://" + j.Host + ":" + j.Port, jwt)
			st, err := c.Status(context.TODO())
			if err != nil {
				log.Print("trying to access judge on ", j.Host, j.Port, " getting error ", err)
				j.Online = false
				j.Ping = -1
				_, err = j.Update(s.DB, boil.Infer())
				if err != nil {
					log.Print("also error occurred while updating", err)
				}

				continue
			}

			j.Online = true
			j.State = st.String()
			j.Ping = 1

			_, err = j.Update(s.DB, boil.Infer())
			if err != nil {
				log.Print("trying to access judge on", j.Host, j.Port, " unsuccesful update in database", err)
				continue
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) runGlue() {
	g := echo.New()
	g.Use(middleware.Logger())

	g.POST("/callback/:id", func(c echo.Context) error {
		id_ := c.Param("id")

		id, err := strconv.Atoi(id_)
		if err != nil {
			return err
		}

		st := judge.Status{}
		if err = c.Bind(&st); err != nil {
			return err
		}

		if st.Done {
			verdict := st.Status.Verdict()
			if st.Status.Compiled == false {
				verdict = extmodels.VERDICT_CE
			}

			if _, err := s.DB.Exec("UPDATE submissions SET verdict=$1, status=$2, ontest=NULL, judged=$3, score=$5 WHERE id=$4", verdict, st.Status, time.Now(), id, st.Status.Score()); err != nil {
				return err
			}
		} else {
			if _, err := s.DB.Exec("UPDATE submissions SET ontest=$1, status=$2, verdict=$3 WHERE id=$4", st.Test, st.Status, extmodels.VERDICT_RU, id); err != nil {
				log.Print("can't realtime update status", err)
			}
		}

		return c.String(http.StatusOK, "ok")
	})

	panic(g.Start(":" + s.GluePort))
}

func (s *Server) runJudger() {
	for {
		time.Sleep(1 * time.Second)

		ss, err := models.Submissions(Where("started=?", false), OrderBy("id ASC"), Limit(1)).All(s.DB)
		if err != nil {
			log.Print("judger query error", err)
			continue
		}

		if len(ss) == 0 {
			continue
		}

		for _, sub := range ss {
			for _, j := range s.judges {
				st, err := judge.ParseServerStatus(j.State)
				if err != nil {
					log.Print("malformed judge: ", j.State, err)
					continue
				}

				if st.SupportsProblem(sub.Problem) {
					token, err := helpers.GetJWT(s.Keys)
					if err != nil {
						log.Print("can't get jwt token", err)
					}

					client := judge.NewClient(st.Url, token)
					if err = client.SubmitCallback(context.TODO(), judge.Submission{Id:strconv.Itoa(sub.ID), Problem: sub.Problem, Language: sub.Language, Source: sub.Source}, "http://" + s.Hostname + ":" + s.GluePort + "/callback/" + strconv.Itoa(sub.ID)); err != nil {
						log.Print("Trying to submit to server", j.Host, j.Port, "Error", err)
						continue
					}

					if _, err = s.DB.Exec("UPDATE submissions SET started=true WHERE id=$1", sub.ID); err != nil {
						log.Print("FATAL: ", err)
					}
					break
				}
			}
		}
	}
}
