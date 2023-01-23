package web

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

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

	s.ConnectToDB()
	s.Keys.MustParse()

	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
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
		time.Sleep(5*time.Second)
	}
}
