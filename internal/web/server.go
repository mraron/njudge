package web

import (
	"database/sql"
	"github.com/mraron/njudge/internal/web/templates"
	_ "mime"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/email"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	_ "github.com/mraron/njudge/pkg/problems/config/task_yaml"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/batch"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/communication"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/stub"
)

type Server struct {
	config.Server
	DB *sql.DB

	ProblemStore  problems.Store
	MailService   email.Service
	PartialsStore templates.Store

	Categories          njudge.Categories
	Tags                njudge.Tags
	Problems            njudge.Problems
	Users               njudge.Users
	Submissions         njudge.Submissions
	SolvedStatusQuery   njudge.SolvedStatusQuery
	ProblemInfoQuery    njudge.ProblemInfoQuery
	ProblemQuery        njudge.ProblemQuery
	ProblemListQuery    njudge.ProblemListQuery
	SubmissionListQuery njudge.SubmissionListQuery

	RegisterService    njudge.RegisterService
	SubmitService      njudge.SubmitService
	TagsService        njudge.TagsService
	TaskArchiveService njudge.TaskArchiveService

	e *echo.Echo
}

func (s *Server) Run() {
	s.e = echo.New()

	s.SetupEnvironment()
	s.StartBackgroundJobs()

	s.setupEcho()

	panic(s.e.Start(":" + s.Port))
}
