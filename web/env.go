package web

import (
	"github.com/jmoiron/sqlx"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/utils/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
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
	//@TODO create a config entry for sslmode
	s.DB, err = sqlx.Open("postgres", "postgres://"+s.DBAccount+":"+s.DBPassword+"@"+s.DBHost+"/"+s.DBName+"?sslmode=disable")

	if err != nil {
		panic(err)
	}

	boil.SetDB(s.DB)
}
