package web

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/utils/problems/config/polygon"
	_ "github.com/mraron/njudge/utils/problems/config/task_yaml"
	"github.com/mraron/njudge/web/extmodels"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/helpers/templates"

	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	_ "github.com/mraron/njudge/utils/language/golang"
	_ "github.com/mraron/njudge/utils/language/julia"
	_ "github.com/mraron/njudge/utils/language/nim"
	_ "github.com/mraron/njudge/utils/language/octave"
	_ "github.com/mraron/njudge/utils/language/pascal"
	_ "github.com/mraron/njudge/utils/language/python3"
	_ "github.com/mraron/njudge/utils/language/zip"

	"github.com/mraron/njudge/web/models"
	_ "github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	_ "mime"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	config.Server

	ProblemStore problems.Store

	judges        []*models.Judge
	db            *sqlx.DB
}

func (s *Server) UpdateProblem(pr string) error {
	return s.ProblemStore.UpdateProblem(pr)
}

func (s *Server) GetProblem(pr string) problems.Problem {
	p, _ := s.ProblemStore.Get(pr)
	return p
}

func (s *Server) runUpdateProblems() {
	for {
		if err := s.ProblemStore.Update(); err != nil {
			log.Print(err)
		}

		time.Sleep(20 * time.Second)
	}
}

func (s *Server) ConnectToDB() {
	var err error
	s.db, err = sqlx.Open("postgres", "postgres://"+s.DBAccount+":"+s.DBPassword+"@"+s.DBHost+"/"+s.DBName)

	if err != nil {
		panic(err)
	}

	boil.SetDB(s.db)
}

func (s *Server) GetDB() *sqlx.DB {
	return s.db
}

func (s *Server) loadJudgesFromDB() {
	var err error
	s.judges, err = models.Judges().All(s.db)

	if err != nil {
		panic(err)
	}
}

func (s *Server) runSyncJudges() {
	for {
		s.loadJudgesFromDB()
		for _, j := range s.judges {
			jwt, err := s.getJWT()
			if err != nil {
				log.Print(err)
				continue
			}

			c := judge.NewClient("http://" + j.Host + ":" + j.Port, jwt)
			st, err := c.Status(context.TODO())
			if err != nil {
				log.Print("trying to access judge on ", j.Host, j.Port, " getting error ", err)
				j.Online = false
				j.Ping = -1
				_, err = j.Update(s.db, boil.Infer())
				if err != nil {
					log.Print("also error occurred while updating", err)
				}

				continue
			}

			j.Online = true
			j.State = st.String()
			j.Ping = 1

			_, err = j.Update(s.db, boil.Infer())
			if err != nil {
				log.Print("trying to access judge on", j.Host, j.Port, " unsuccesful update in database", err)
				continue
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) runGlue() {
	g := echo.New()
	g.Use(middleware.Logger())

	g.POST("/callback/:id", func(c echo.Context) error {
		id_ := c.Param("id")

		id, err := strconv.Atoi(id_)
		if err != nil {
			return helpers.InternalError(c, err, "err")
		}

		st := judge.Status{}
		if err = c.Bind(&st); err != nil {
			return helpers.InternalError(c, err, "err")
		}

		if st.Done {
			verdict := st.Status.Verdict()
			if st.Status.Compiled == false {
				verdict = extmodels.VERDICT_CE
			}

			if _, err := s.db.Exec("UPDATE submissions SET verdict=$1, status=$2, ontest=NULL, judged=$3, score=$5 WHERE id=$4", verdict, st.Status, time.Now(), id, st.Status.Score()); err != nil {
				return helpers.InternalError(c, err, "err")
			}
		} else {
			if _, err := s.db.Exec("UPDATE submissions SET ontest=$1, status=$2, verdict=$3 WHERE id=$4", st.Test, st.Status, extmodels.VERDICT_RU, id); err != nil {
				log.Print("can't realtime update status", err)
			}
		}

		return c.String(http.StatusOK, "ok")
	})

	panic(g.Start(":" + s.GluePort))
}

func (s *Server) runJudger() {
	for {
		time.Sleep(1 * time.Second)

		ss, err := models.Submissions(Where("started=?", false), OrderBy("id ASC"), Limit(1)).All(s.db)
		if err != nil {
			log.Print("judger query error", err)
			continue
		}

		if len(ss) == 0 {
			continue
		}

		for _, sub := range ss {
			for _, j := range s.judges {
				st, err := judge.ParseServerStatus(j.State)
				if err != nil {
					log.Print("malformed judge: ", j.State, err)
					continue
				}

				if st.SupportsProblem(sub.Problem) {
					token, err := s.getJWT()
					if err != nil {
						log.Print("can't get jwt token", err)
					}

					client := judge.NewClient(st.Url, token)
					if err = client.SubmitCallback(context.TODO(), judge.Submission{Id:strconv.Itoa(sub.ID), Problem: sub.Problem, Language: sub.Language, Source: sub.Source}, "http://" + s.Hostname + ":" + s.GluePort + "/callback/" + strconv.Itoa(sub.ID)); err != nil {
						log.Print("Trying to submit to server", j.Host, j.Port, "Error", err)
						continue
					}

					if _, err = s.db.Exec("UPDATE submissions SET started=true WHERE id=$1", sub.ID); err != nil {
						log.Print("FATAL: ", err)
					}
					break
				}
			}
		}
	}
}

func (s *Server) StartBackgroundProcesses() {
	go s.runUpdateProblems()
	go s.runSyncJudges()
	go s.runGlue()
	go s.runJudger()
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

	s.ConnectToDB()
	s.parseKeys()

	s.ProblemStore = problems.NewFsStore(s.ProblemsDir)
}

func (s *Server) Run() {
	s.SetupEnvironment()
	s.StartBackgroundProcesses()

	e := echo.New()
	store := sessions.NewCookieStore([]byte(s.CookieSecret))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := s.currentUser(c)
			if err != nil {
				return helpers.InternalError(c, err, "belső hiba")
			}
			c.Set("user", user)

			return next(c)
		}
	})

	e.Renderer = templates.New(s.TemplatesDir, s.ProblemStore)

	e.GET("/", s.getHome)

	e.Static("/static", "public")

	e.GET("/submission/:id", s.getSubmission)
	e.GET("/submission/rejudge/:id", s.getSubmissionRejudge)
	e.GET("/task_archive", s.getTaskArchive)

	ps := e.Group("/problemset", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("problemset", c.Param("name"))
			return next(c)
		}
	})

	ps.GET("/:name/", s.getProblemsetList)
	ps.GET("/:name/:problem/", s.getProblemsetProblem)
	ps.GET("/:name/:problem/problem", s.getProblemsetProblem)
	ps.GET("/:name/:problem/status", s.getProblemsetProblemStatus)
	ps.GET("/:name/:problem/ranklist", s.getProblemsetProblemRanklist)
	ps.GET("/:name/:problem/pdf/:language/", s.getProblemsetProblemPDFLanguage)
	ps.GET("/:name/:problem/attachment/:attachment/", s.getProblemsetProblemAttachment)
	ps.GET("/:name/:problem/:file", s.getProblemsetProblemFile)
	ps.POST("/:name/submit", s.postProblemsetSubmit)
	ps.GET("/status", s.getProblemsetStatus)

	u := e.Group("/user")

	u.GET("/auth/callback", s.getUserAuthCallback)
	u.GET("/auth", s.getUserAuth)

	u.GET("/login", s.getUserLogin)
	u.POST("/login", s.postUserLogin)
	u.GET("/logout", s.getUserLogout)
	u.GET("/register", s.getUserRegister)
	u.POST("/register", s.postUserRegister)
	u.GET("/activate", s.getUserActivate)
	u.GET("/activate/:name/:key", s.getActivateUser)

	u.GET("/profile/:name/", s.getUserProfile)

	v1 := e.Group("/api/v1")

	v1.GET("/problem_rels", s.getAPIProblemRels)
	v1.POST("/problem_rels", s.postAPIProblemRel)
	v1.GET("/problem_rels/:id", s.getAPIProblemRel)
	v1.PUT("/problem_rels/:id", s.putAPIProblemRel)
	v1.DELETE("/problem_rels/:id", s.deleteAPIProblemRel)

	v1.GET("/judges", s.getAPIJudges)
	v1.POST("/judges", s.postAPIJudge)
	v1.GET("/judges/:id", s.getAPIJudge)
	v1.PUT("/judges/:id", s.putAPIJudge)
	v1.DELETE("/judges/:id", s.deleteAPIJudge)

	v1.GET("/users", s.getAPIUsers)
	v1.POST("/users", s.postAPIUser)
	v1.GET("/users/:id", s.getAPIUser)
	v1.PUT("/users/:id", s.putAPIUser)
	v1.DELETE("/users/:id", s.deleteAPIUser)

	v1.GET("/submissions", s.getAPISubmissions)
	v1.POST("/submissions", s.postAPISubmission)
	v1.GET("/submissions/:id", s.getAPISubmission)
	v1.PUT("/submissions/:id", s.putAPISubmission)
	v1.DELETE("/submissions/:id", s.deleteAPISubmission)

	e.GET("/admin", s.getAdmin)

	panic(e.Start(":" + s.Port))
}

func (s *Server) getHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.gohtml", nil)
}

func (s *Server) getAdmin(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "admin_panel") {
		return c.Render(http.StatusUnauthorized, "error.gohtml", "Engedély megtagadva.")
	}

	return c.Render(http.StatusOK, "admin.gohtml", struct {
		Host string
	}{s.Hostname + ":" + s.Port})
}
