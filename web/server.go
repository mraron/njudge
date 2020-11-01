package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/judge"

	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/utils/problems/config/polygon"

	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	_ "github.com/mraron/njudge/utils/language/golang"
	_ "github.com/mraron/njudge/utils/language/julia"
	_ "github.com/mraron/njudge/utils/language/octave"
	_ "github.com/mraron/njudge/utils/language/python3"
	_ "github.com/mraron/njudge/utils/language/zip"

	"github.com/mraron/njudge/web/models"
	_ "github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"html/template"
	"io/ioutil"
	_ "mime"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	Hostname     string
	Port         string
	ProblemsDir  string
	SubmissionsDir string
	TemplatesDir string

	MailAccount         string
	MailServerHost      string
	MailServerPort      string
	MailAccountPassword string

	DBAccount  string
	DBPassword string
	DBHost     string
	DBName     string

	GluePort string

	judges        []*models.Judge
	problems      map[string]problems.Problem
	problemsMutex sync.Mutex
	db            *sqlx.DB
}

/*
func New(port string, problemsDir string, templatesDir string, mailServerAccount, mailServerHost, mailServerPort, mailAccountPassword string, glueport string) *Server {
	return &Server{port, problemsDir, templatesDir, mailServerAccount, mailServerHost, mailServerPort, mailAccountPassword, glueport, make([]*models.Judge, 0), make(map[string]problems.Problem), nil}
}*/

func (s *Server) getProblem(name string) problems.Problem {
	s.problemsMutex.Lock()
	p := s.problems[name]
	s.problemsMutex.Unlock()

	return p
}

func (s *Server) getProblemExists(name string) (problems.Problem, bool) {
	s.problemsMutex.Lock()
	p, ok := s.problems[name]
	s.problemsMutex.Unlock()

	return p, ok
}

func (s *Server) runUpdateProblems() {
	for {
		s.problemsMutex.Lock()
		if s.problems == nil {
			s.problems = make(map[string]problems.Problem)
		}

		files, err := ioutil.ReadDir(s.ProblemsDir)
		if err != nil {
			panic(err)
		}

		pList := make([]string, 0)

		for _, f := range files {
			if f.IsDir() {
				path := filepath.Join(s.ProblemsDir, f.Name())
				p, err := problems.Parse(path)
				if err == nil {
					s.problems[p.Name()] = p
					pList = append(pList, p.Name())
				} else {
					log.Print(err)
				}
			}
		}

		s.problemsMutex.Unlock()

		time.Sleep(20 * time.Second)
	}
}

func (s *Server) connectToDB() {
	var err error
	s.db, err = sqlx.Open("postgres", "postgres://"+s.DBAccount+":"+s.DBPassword+"@"+s.DBHost+"/"+s.DBName)

	if err != nil {
		panic(err)
	}

	boil.SetDB(s.db)
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
			st, err := judge.NewFromUrl("http://" + j.Host + ":" + j.Port)

			if err != nil {
				log.Print("trying to access judge on ", j.Host, j.Port, " getting error ", err)
				j.Online = false
				j.Ping = -1
				_, err = j.Update(s.db, boil.Infer())
				if err != nil {
					log.Print("also error occured while updating", err)
				}

				continue
			}

			j.Online = true
			j.State, _ = st.ToString()
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

func (s *Server) internalError(c echo.Context, err error, msg string) error {
	c.Logger().Print("internal error:", err)
	return c.Render(http.StatusInternalServerError, "error.html", msg)
}

func (s *Server) unauthorizedError(c echo.Context) error {
	return c.String(http.StatusUnauthorized, "unauthorized")
}

func (s *Server) runGlue() {
	g := echo.New()
	g.Use(middleware.Logger())

	g.POST("/callback/:id", func(c echo.Context) error {
		id_ := c.Param("id")

		id, err := strconv.Atoi(id_)
		if err != nil {
			return s.internalError(c, err, "err")
		}

		st := judge.Status{}
		if err = c.Bind(&st); err != nil {
			return s.internalError(c, err, "err")
		}

		if st.Done {
			verdict := Verdict(st.Status.Verdict())
			if st.Status.Compiled == false {
				verdict = VERDICT_CE
			}

			if _, err := s.db.Exec("UPDATE submissions SET verdict=$1, status=$2, ontest=NULL, judged=$3 WHERE id=$4", verdict, st.Status, time.Now(), id); err != nil {
				return s.internalError(c, err, "err")
			}
		} else {
			if _, err := s.db.Exec("UPDATE submissions SET ontest=$1, status=$2, verdict=$3 WHERE id=$4", st.Test, st.Status, VERDICT_RU, id); err != nil {
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
				server := &judge.Server{}
				server.FromString(j.State)

				if server.SupportsProblem(sub.Problem) {
					err := server.Submit(judge.Submission{sub.ID, sub.Problem, sub.Language, sub.Source, "http://" + s.Hostname + ":" + s.GluePort + "/callback/" + strconv.Itoa(int(sub.ID))})
					if err != nil {
						log.Print("Trying to submit to server", j.Host, j.Port, "Error", err)
						continue
					}
					if _, err := s.db.Exec("UPDATE submissions SET started=true WHERE id=$1", sub.ID); err != nil {
						log.Print("FATAL: ", err)
					}
					break
				}
			}
		}
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("titkosdolog"))))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := s.currentUser(c)
			if err != nil {
				return s.internalError(c, err, "belső hiba")
			}

			c.Set("user", user)

			return next(c)
		}
	})

	t := &Template{
		templates: template.Must(template.New("templater").Funcs(s.templatefuncs()).ParseGlob(filepath.Join(s.TemplatesDir, "*.html"))),
	}

	e.Renderer = t

	e.GET("/", s.getHome)

	e.Static("/static", "public")
	e.GET("/submission/:id", s.getSubmission)

	ps := e.Group("/problemset", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("problemset", c.Param("name"))
			return next(c)
		}
	})

	ps.GET("/:name/", s.getProblemsetMain)
	ps.GET("/:name/:problem/", s.getProblemsetProblem)
	ps.GET("/:name/:problem/pdf/:language/", s.getProblemsetProblemPDFLanguage)
	ps.GET("/:name/:problem/attachment/:attachment/", s.getProblemsetProblemAttachment)
	ps.GET("/:name/:problem/:file", s.getProblemsetProblemFile)
	ps.POST("/:name/submit", s.postProblemsetSubmit)
	ps.GET("/status", s.getProblemsetStatus)

	u := e.Group("/user")

	u.GET("/login", s.getUserLogin)
	u.POST("/login", s.postUserLogin)
	u.GET("/logout", s.getUserLogout)
	u.GET("/register", s.getUserRegister)
	u.POST("/register", s.postUserRegister)
	u.GET("/activate", s.getUserActivate)
	u.GET("/activate/:name/:key", s.getActivateUser)

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

	s.connectToDB()

	go s.runUpdateProblems()
	go s.runSyncJudges()
	go s.runGlue()
	go s.runJudger()

	for idx, judge := range s.judges {
		fmt.Println(idx, judge)
	}

	panic(e.Start(":" + s.Port))
}

func (s *Server) getHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func (s *Server) getAdmin(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "admin_panel") {
		return c.Render(http.StatusUnauthorized, "error.html", "Engedély megtagadva.")
	}

	return c.Render(http.StatusOK, "admin.html", struct {
		Host string
	}{s.Hostname + ":" + s.Port})
}
