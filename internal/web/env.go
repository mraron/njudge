package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mraron/njudge/internal/njudge/cached"
	templates2 "github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	slogecho "github.com/samber/slog-echo"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/mraron/njudge/internal/web/helpers/templates"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/quasoft/memstore"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Server) SetupDataAccess() {
	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
	_ = s.ProblemStore.UpdateProblems()

	if s.Mode == "demo" {
		s.PartialsStore = templates2.Empty{}

		s.Categories = memory.NewCategories()
		s.Tags = memory.NewTags()
		s.Problems = memory.NewProblems()
		s.Submissions = memory.NewSubmissions()
		s.Users = memory.NewUsers()

		s.ProblemQuery = memory.NewProblemQuery(s.Problems)
		s.ProblemInfoQuery = memory.NewProblemInfoQuery(s.Submissions)
		s.ProblemListQuery = memory.NewProblemListQuery(s.ProblemStore, s.Problems, s.Tags, s.Categories)
		s.SubmissionListQuery = memory.NewSubmissionListQuery(s.Submissions, s.Problems)

		s.RegisterService = njudge.NewRegisterService(s.Users)
		s.SubmitService = memory.NewSubmitService(s.Submissions, s.Users, s.ProblemQuery, s.ProblemStore)
		s.TagsService = memory.NewTagsService(s.Tags, s.Problems, s.ProblemInfoQuery)

		nt1 := njudge.NewCategory("NT1", nil)
		nt1, _ = s.Categories.Insert(context.Background(), *nt1)
		nt1_2021 := njudge.NewCategory("2021", nt1)
		nt1_2021, _ = s.Categories.Insert(context.Background(), *nt1_2021)

		t := njudge.NewTag("constructive")
		t, _ = s.Tags.Insert(context.Background(), *t)

		t2 := njudge.NewTag("dp")
		t2, _ = s.Tags.Insert(context.Background(), *t2)

		p := njudge.NewProblem("main", "NT21_Atvagas")
		p.AddTag(*t, 1)
		p.SetCategory(*nt1_2021)

		u, _ := njudge.NewUser("mraron", "email@email.com", "admin")
		u.SetPassword("abc")
		u.Activate()
		u, _ = s.Users.Insert(context.Background(), *u)

		s.Problems.Insert(context.Background(), njudge.NewProblem("main", "is1"))
		prob, _ := s.Problems.Insert(context.Background(), p)

		sub, _ := njudge.NewSubmission(*u, *prob, cpp.Std14)
		sub.SetSource([]byte("#include<bits/stdc++.h>"))
		sub.Verdict = njudge.VerdictAC
		sdata, _ := prob.WithStoredData(s.ProblemStore)
		ss, _ := sdata.StatusSkeleton("")
		ss.Compiled = true
		sub.Status = *ss
		sub, _ = s.Submissions.Insert(context.Background(), *sub)
	} else {
		s.PartialsStore = templates2.NewCached(s.DB, 1*time.Minute)

		s.Categories = db.NewCategories(s.DB)
		s.Tags = db.NewTags(s.DB)
		s.SolvedStatusQuery = cached.NewSolvedStatusQuery(db.NewSolvedStatusQuery(s.DB), 30*time.Second)
		s.Problems = db.NewProblems(
			s.DB,
			s.SolvedStatusQuery,
		)
		s.Submissions = db.NewSubmissions(s.DB)
		s.Users = db.NewUsers(s.DB)

		s.ProblemQuery = s.Problems.(*db.Problems)
		s.ProblemInfoQuery = s.Problems.(*db.Problems)
		s.ProblemListQuery = memory.NewProblemListQuery(s.ProblemStore, s.Problems, s.Tags, s.Categories)
		s.SubmissionListQuery = db.NewSubmissionListQuery(s.DB)

		s.RegisterService = njudge.NewRegisterService(s.Users)
		s.SubmitService = memory.NewSubmitService(s.Submissions, s.Users, s.ProblemQuery, s.ProblemStore)
		s.TagsService = memory.NewTagsService(s.Tags, s.Problems, s.ProblemInfoQuery)
	}
}

func (s *Server) SetupEnvironment() {
	if s.Mode == "development" || s.Mode == "demo" {
		boil.DebugMode = true
	}

	loc, err := time.LoadLocation("Europe/Budapest")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	boil.SetLocation(loc)

	if s.GoogleAuth.Enabled {
		goth.UseProviders(
			google.New(s.GoogleAuth.ClientKey, s.GoogleAuth.Secret, s.GoogleAuth.Callback, "email", "profile"),
		)
	}

	if s.Mode != "demo" {
		s.ConnectToDB()
	}

	s.SetupDataAccess()

	if s.SMTP.Enabled {
		port, err := strconv.Atoi(s.SMTP.MailServerPort)
		if err != nil {
			panic(err)
		}

		s.MailService = email.SMTPService{
			From:     s.SMTP.MailAccount,
			Host:     s.SMTP.MailServerHost,
			Port:     port,
			User:     s.SMTP.MailAccount,
			Password: s.SMTP.MailAccountPassword,
		}
	} else if s.Sendgrid.Enabled {
		s.MailService = email.SendgridService{
			SenderName:    s.Sendgrid.SenderName,
			SenderAddress: s.Sendgrid.SenderAddress,
			APIKey:        s.Sendgrid.ApiKey,
		}
	} else {
		if s.Mode == "development" || s.Mode == "demo" {
			s.MailService = email.LogService{Logger: log.Default()}
		} else {
			s.MailService = email.ErrorService{}
		}
	}
}

func (s *Server) ConnectToDB() {
	var err error

	sslmode := "require"
	if !s.SSLMode {
		sslmode = "disable"
	}

	if s.Database.Port == 0 {
		s.Database.Port = 5432
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", s.User, s.Password, s.Host, s.Name, s.Database.Port, sslmode)
	s.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	boil.SetDB(s.DB)

	for {
		log.Print("Trying to ping database...")
		if err := s.DB.Ping(); err == nil {
			log.Print("OK, connected to database")
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func (s *Server) setupEcho() {
	if s.Mode == "development" || s.Mode == "demo" {
		s.e.Debug = true
	} else {
		s.e.HTTPErrorHandler = func(err error, c echo.Context) {
			code := http.StatusInternalServerError
			var he *echo.HTTPError
			if errors.As(err, &he) {
				code = he.Code
			}

			templates2.Render(c, code, templates2.Error("Hiba történt."))
			c.Logger().Error(err)
		}
	}

	s.e.Renderer = templates.New(s.Server, s.ProblemStore, s.Users, s.Problems, s.Tags, s.PartialsStore)

	var (
		store sessions.Store
		err   error
	)

	if s.Mode == "development" || s.Mode == "production" {
		store, err = pgstore.NewPGStoreFromPool(s.DB, []byte(s.CookieSecret))
		if err != nil {
			panic(err)
		}
	} else {
		store = memstore.NewMemStore(
			[]byte("authkey123"),
			[]byte("enckey12341234567890123456789012"),
		)
	}

	s.e.Use(slogecho.New(slog.Default()))
	s.e.Use(middleware.Recover())
	s.e.Use(session.Middleware(store))
	s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	s.prepareRoutes(s.e)

}
