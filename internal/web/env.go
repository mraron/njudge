package web

import (
	"context"
	"time"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Server) SetupEnvironment(ctx context.Context) error {
	if s.Mode == ModeDebug {
		s.Logger.Debug("debug mode enabled")
		boil.DebugMode = true
	}

	s.Logger.Info("timezone from config: " + s.TimeZone)
	loc, err := time.LoadLocation(s.TimeZone)
	if err != nil {
		return err
	}
	time.Local = loc
	boil.SetLocation(loc)

	if s.GoogleAuth.Enabled {
		goth.UseProviders(
			google.New(s.GoogleAuth.ClientKey, s.GoogleAuth.Secret, s.Url+"/user/auth/callback?provider=google", "email", "profile"),
		)
	} else {
		s.Logger.Info("no google auth enabled")
	}

	return nil
}
