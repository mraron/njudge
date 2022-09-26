package glue

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/pkg/judge"
	"github.com/mraron/njudge/pkg/web/extmodels"
	"github.com/mraron/njudge/pkg/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Config struct {
	Port string
}

type Server struct {
	Config

	DB            *sql.DB
	JudgesUpdater JudgesUpdater
	FindJudger    FindJudger

	judges      []*models.Judge
	judgesMutex sync.RWMutex
}

func (s *Server) Run() {
	go s.runSyncJudges()
	go s.runJudger()
	s.runServer()
}

func (s *Server) runSyncJudges() {
	var err error
	for {
		s.judgesMutex.Lock()
		if s.judges, err = s.JudgesUpdater.JudgesUpdate(); err != nil {
			log.Print(err)
		}
		s.judgesMutex.Unlock()

		time.Sleep(10 * time.Second)
	}
}

func (s *Server) runServer() {
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
			if !st.Status.Compiled {
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

	panic(g.Start(":" + s.Port))
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
			s.judgesMutex.RLock()
			j, err := s.FindJudger.FindJudge(s.judges, sub)
			if err != nil {
				log.Print(err)
				continue
			}
			s.judgesMutex.RUnlock()

			var st judge.ServerStatus
			st, err = judge.ParseServerStatus(j.State)
			if err != nil {
				log.Print(err)
				continue
			}

			client := judge.NewClient(st.Url, "")
			if err := client.SubmitCallback(context.TODO(), judge.Submission{Id: strconv.Itoa(sub.ID), Problem: sub.Problem, Language: sub.Language, Source: sub.Source}, "http://glue:"+s.Port+"/callback/"+strconv.Itoa(sub.ID)); err != nil {
				log.Print("Trying to submit to server", j.Host, j.Port, "Error", err)
				continue
			}

			if _, err := s.DB.Exec("UPDATE submissions SET started=true WHERE id=$1", sub.ID); err != nil {
				log.Print(err)
				continue
			}
		}
	}
}
