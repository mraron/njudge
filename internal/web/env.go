package web

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Server) SetupDataAccess() {
	if s.Mode == "development" {
		s.Problems = memory.NewProblems()
		s.Submissions = memory.NewSubmissions()
		s.Users = memory.NewUsers()
		s.ProblemQuery = memory.NewProblemQuery(s.Problems)
		s.ProblemInfoQuery = memory.NewProblemInfoQuery(s.Submissions)

		s.Problems.Insert(context.Background(), njudge.NewProblem("main", "is1"))
	} else {
		panic("not supported yet :P")
	}
}

func (s *Server) SetupEnvironment() {
	if s.Mode == "development" {
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

	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
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
		if s.Mode == "development" {
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
