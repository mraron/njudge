package web

import (
	"context"
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	_ "github.com/mraron/njudge/pkg/problems/config/task_yaml"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/batch"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/communication"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/stub"
	"log/slog"
	_ "mime"
	"net"
	"net/http"
)

type Server struct {
	Logger *slog.Logger
	Config
	DataAccess
	DB *sql.DB
}

func NewServer(logger *slog.Logger, cfg Config, dataAccess DataAccess, db *sql.DB) (*Server, error) {
	res := Server{
		Logger:     logger,
		Config:     cfg,
		DataAccess: dataAccess,
		DB:         db,
	}
	if cfg.Mode.UsesDB() && db == nil {
		return nil, errors.New("db connection is required")
	}
	return &res, nil
}

func (s *Server) Run(ctx context.Context) error {
	e := echo.New()
	if err := s.SetupEnvironment(ctx); err != nil {
		return err
	}

	s.StartBackgroundJobs(ctx)
	s.SetupEcho(ctx, e)

	s.Logger.Info("listening on 0.0.0.0:" + s.Port)
	server := http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", s.Port),
		Handler: e,
	}
	// todo graceful shutdown
	return server.ListenAndServe()
}
