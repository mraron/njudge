package web

import (
	"context"
	_ "mime"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/web/services"

	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/internal/web/helpers/templates/partials"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	_ "github.com/mraron/njudge/pkg/problems/config/task_yaml"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/batch"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/communication"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/output_only"
	_ "github.com/mraron/njudge/pkg/problems/tasktype/stub"
)

type Server struct {
	config.Server
	DB *sqlx.DB

	ProblemStore  problems.Store
	MailService   email.Service
	PartialsStore partials.Store

	Categories          njudge.Categories
	Tags                njudge.Tags
	Problems            njudge.Problems
	Users               njudge.Users
	Submissions         njudge.Submissions
	ProblemInfoQuery    njudge.ProblemInfoQuery
	ProblemQuery        njudge.ProblemQuery
	ProblemListQuery    njudge.ProblemListQuery
	SubmissionListQuery njudge.SubmissionListQuery

	RegisterService njudge.RegisterService
	SubmitService   njudge.SubmitService
	TagsService     njudge.TagsService

	e *echo.Echo
}

func (s *Server) Run() {
	s.e = echo.New()

	s.SetupEnvironment()
	s.StartBackgroundJobs()

	s.setupEcho()

	panic(s.e.Start(":" + s.Port))
}

func (s *Server) Submit(uid int, problemset, problem, language string, source []byte) (int, error) {
	subService := services.NewSQLSubmitService(s.DB.DB, s.ProblemStore)
	sub, err := subService.Submit(context.Background(), services.SubmitRequest{
		UserID:     uid,
		Problemset: problemset,
		Problem:    problem,
		Language:   language,
		Source:     source,
	})

	if err != nil {
		return -1, err
	}
	return sub.ID, nil
}
