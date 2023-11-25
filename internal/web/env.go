package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/mraron/njudge/internal/web/helpers/templates"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/quasoft/memstore"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Server) SetupDataAccess() {
	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
	s.ProblemStore.Update()

	if s.Mode == "demo" {
		s.PartialsStore = partials.Empty{}

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

		sub, _ := njudge.NewSubmission(*u, *prob, language.DefaultStore.Get("cpp14"))
		sub.SetSource([]byte("#include<bits/stdc++.h>"))
		sub.Verdict = njudge.VerdictAC
		sdata, _ := prob.WithStoredData(s.ProblemStore)
		ss, _ := sdata.StatusSkeleton("")
		sub.Status = *ss
		sub, _ = s.Submissions.Insert(context.Background(), *sub)
	} else {
		panic("not supported yet :P")
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

	//TODO
	//s.ConnectToDB()
	s.SetupDataAccess()
	s.Keys.MustParse()

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
	if !s.DBSSLMode {
		sslmode = "disable"
	}

	if s.DBPort == 0 {
		s.DBPort = 5432
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", s.DBAccount, s.DBPassword, s.DBHost, s.DBName, s.DBPort, sslmode)
	s.DB, err = sqlx.Open("postgres", connStr)
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
			if he, ok := err.(*echo.HTTPError); ok {
				code = he.Code
			}

			if err := c.Render(code, "error.gohtml", "Hiba történt"); err != nil {
				c.Logger().Error(err)
			}

			c.Logger().Error(err)
		}
	}

	s.e.Renderer = templates.New(s.Server, s.ProblemStore, s.Users, s.Problems, s.Tags, s.PartialsStore)

	var (
		store sessions.Store
		err   error
	)

	if s.Mode == "development" || s.Mode == "production" {
		store, err = pgstore.NewPGStoreFromPool(s.DB.DB, []byte(s.CookieSecret))
		if err != nil {
			panic(err)
		}
	} else {
		store = memstore.NewMemStore(
			[]byte("authkey123"),
			[]byte("enckey12341234567890123456789012"),
		)
	}

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(session.Middleware(store))

	s.prepareRoutes(s.e)

}
