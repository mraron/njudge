package web

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/web/extmodels"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) StartBackgroundJobs() {
	go s.runUpdateProblems()
	go s.runSyncJudges()
	go s.runGlue()
	go s.runJudger()
	go s.runStatisticsUpdate()
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
		s.judgesMutex.Lock()
		defer s.judgesMutex.Unlock()

		var err error
		s.judges, err = models.Judges().All(s.DB)

		if err != nil {
			panic(err)
		}
	}

	updateJudges := func() {
		s.judgesMutex.RLock()
		defer s.judgesMutex.RUnlock()

		for _, j := range s.judges {
			jwt, err := helpers.GetJWT(s.Keys)
			if err != nil {
				log.Print(err)
				continue
			}

			c := judge.NewClient("http://"+j.Host+":"+j.Port, jwt)
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
	}

	for {
		loadJudgesFromDB()
		updateJudges()

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

			sub, err := models.Submissions(Where("id=?", id)).One(s.DB)
			if err != nil {
				return err
			}

			if _, err = s.DB.Exec("UPDATE problem_rels SET solver_count = (SELECT COUNT(distinct user_id) from submissions where problemset = problem_rels.problemset and problem = problem_rels.problem and verdict = 0) WHERE problemset = $1 and problem = $2", sub.Problemset, sub.Problem); err != nil {
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
	findJudger := func(sub *models.Submission) {
		s.judgesMutex.RLock()
		defer s.judgesMutex.RUnlock()

		for _, j := range s.judges {
			st, err := judge.ParseServerStatus(j.State)
			if err != nil {
				log.Print("malformed judge: ", j.State, err)
				continue
			}

			if j.Online && st.SupportsProblem(sub.Problem) {
				token, err := helpers.GetJWT(s.Keys)
				if err != nil {
					log.Print("can't get jwt token", err)
					continue
				}

				client := judge.NewClient(st.Url, token)

				//@TODO web is a placeholder until migrating to streaming response
				if err = client.SubmitCallback(context.TODO(), judge.Submission{Id: strconv.Itoa(sub.ID), Problem: sub.Problem, Language: sub.Language, Source: sub.Source}, "http://web:"+s.GluePort+"/callback/"+strconv.Itoa(sub.ID)); err != nil {
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
			findJudger(sub)
		}
	}
}

func (s *Server) runStatisticsUpdate() {
	for {
		probs, err := models.ProblemRels().All(s.DB)
		if err != nil {
			log.Print(err)
			continue
		}

		userPoints := make(map[int]float64)
		for _, p := range probs {
			points := math.Sqrt(1.0 / float64(p.SolverCount))
			solvedBy, err := models.Submissions(Distinct("user_id"), Where("verdict = 0"), Where("problemset = ?", p.Problemset), Where("problem = ?", p.Problem)).All(s.DB)
			if err != nil {
				log.Print(err)
				continue
			}

			for _, uid := range solvedBy {
				userPoints[uid.UserID] += points
			}
		}

		for uid, pts := range userPoints {
			var user models.User
			user.ID = uid
			user.Points = null.Float32{Float32: float32(pts), Valid: true}
			if _, err := user.Update(s.DB, boil.Whitelist("points")); err != nil {
				log.Print(err)
				continue
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
