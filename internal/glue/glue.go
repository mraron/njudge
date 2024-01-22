package glue

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"

	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge/db"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/helpers/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volatiletech/null/v8"
)

type Config struct {
	Port string

	config.Database `mapstructure:",squash"`
}

type Server struct {
	Config

	DB            *sql.DB
	JudgesUpdater JudgesUpdater
	JudgeFinder   JudgeFinder

	Submissions njudge.Submissions
	Problems    njudge.Problems

	SubmissionsQuery njudge.SubmissionsQuery

	judges      []*models.Judge
	judgesMutex sync.RWMutex
}

func (s *Server) ConnectToDB() {
	var err error

	sslmode := "require"
	if !s.DBSSLMode {
		sslmode = "disable"
	}

	if s.DBPort == 0 {
		s.DBPort = 5432
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", s.DBAccount, s.DBPassword, s.DBHost, s.DBName, s.DBPort, sslmode)
	s.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	s.Submissions = db.NewSubmissions(s.DB)
	s.Problems = db.NewProblems(s.DB, db.NewSolvedStatusQuery(s.DB))
	s.SubmissionsQuery = s.Submissions.(*db.Submissions)
}

func (s *Server) Run() {
	s.ConnectToDB()
	s.JudgesUpdater = &JudgesUpdaterFromDB{s.DB}
	s.JudgeFinder = &FindJudgerNaive{}

	go s.runSyncJudges()
	go s.runJudger()
	s.runServer()
}

func (s *Server) runSyncJudges() {
	var err error
	for {
		s.judgesMutex.Lock()
		if s.judges, err = s.JudgesUpdater.UpdateJudges(context.Background()); err != nil {
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

		st := judge.SubmissionStatus{}
		if err = c.Bind(&st); err != nil {
			return err
		}

		if st.Done {
			var (
				verdict problems.VerdictName
				score   float32 = 0.0
			)

			if !st.Status.Compiled {
				verdict = problems.VerdictName(njudge.VerdictCE)
			} else {
				verdict = st.Status.Feedback[0].Verdict()
				score = float32(st.Status.Feedback[0].Score())
			}

			sub := njudge.Submission{
				ID:      id,
				Verdict: njudge.Verdict(verdict),
				Status:  st.Status,
				Ontest: null.String{
					Valid:  false,
					String: "",
				},
				Judged: null.NewTime(time.Now(), true),
				Score:  score,
			}

			if err := s.Submissions.Update(c.Request().Context(), sub, njudge.Fields(
				njudge.SubmissionFields.Verdict,
				njudge.SubmissionFields.Status,
				njudge.SubmissionFields.Ontest,
				njudge.SubmissionFields.Judged,
				njudge.SubmissionFields.Score,
			)); err != nil {
				return err
			}
		} else {
			sub := njudge.Submission{
				ID:      id,
				Verdict: njudge.VerdictRU,
				Status:  st.Status,
				Ontest:  null.NewString(st.Test, true),
			}

			if err := s.Submissions.Update(c.Request().Context(), sub, njudge.Fields(
				njudge.SubmissionFields.Verdict,
				njudge.SubmissionFields.Status,
				njudge.SubmissionFields.Ontest,
			)); err != nil {
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

		ss, err := s.SubmissionsQuery.GetUnstarted(context.Background(), 5)
		if err != nil {
			log.Print("judger query error", err)
			continue
		}

		if len(ss) == 0 {
			continue
		}

		for _, sub := range ss {
			p, err := s.Problems.Get(context.Background(), sub.ProblemID)
			if err != nil {
				log.Print(err)
				continue
			}

			s.judgesMutex.RLock()
			j, err := s.JudgeFinder.FindJudge(s.judges, p.Problem)
			if err != nil {
				log.Print(err)
				continue
			}
			s.judgesMutex.RUnlock()
			if j == nil {
				continue
			}

			var st judge.ServerStatus
			st, err = judge.ParseServerStatus(j.State)
			if err != nil {
				log.Print(err)
				continue
			}

			client := judge.NewClient(st.Url)
			if err := client.SubmitCallback(context.Background(),
				judge.Submission{
					Id:       strconv.Itoa(sub.ID),
					Problem:  p.Problem,
					Language: sub.Language,
					Source:   sub.Source,
				}, fmt.Sprintf("http://glue:%s/callback/%d", s.Port, sub.ID)); err != nil {
				log.Print("Trying to submit to server", j.Host, j.Port, "Error", err)
				continue
			}

			sub.Started = true
			if err := s.Submissions.Update(
				context.Background(),
				sub,
				njudge.Fields(njudge.SubmissionFields.Started),
			); err != nil {
				log.Print(err)
				continue
			}
		}
	}
}
