package web

import (
	"crypto/rsa"
	"fmt"
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
	"github.com/mraron/njudge/web/roles"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	Mode string
	Hostname       string
	Port           string
	ProblemsDir    string
	SubmissionsDir string
	TemplatesDir   string

	CookieSecret string

	GoogleAuth struct {
		Enabled   bool
		ClientKey string
		Secret    string
		Callback  string
	}

	Sendgrid struct {
		Enabled bool
		ApiKey string `json:"api_key"`
		SenderName string `json:"sender_name"`
		SenderAddress string `json:"sender_address"`
	}

	SMTP struct {
		Enabled bool
		MailAccount         string `json:"mail_account"`
		MailServerHost      string `json:"mail_server"`
		MailServerPort      string `json:"mail_port"`
		MailAccountPassword string `json:"mail_password"`
	} `json:"smtp"`

	DBAccount  string
	DBPassword string
	DBHost     string
	DBName     string

	GluePort string

	Keys struct {
		PrivateKeyLocation string `json:"private_key"`
		PublicKeyLocation string `json:"public_key"`
		privateKey *rsa.PrivateKey
		publicKey *rsa.PublicKey
	}

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

func (s *Server) AddProblem(dir string) (problems.Problem, error) {
	if s.problems == nil {
		s.problems = make(map[string]problems.Problem)
	}

	path := filepath.Join(s.ProblemsDir, dir)
	p, err := problems.Parse(path)
	if err == nil {
		s.problems[p.Name()] = p
	}

	return p, err
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
		for _, f := range files {
			if f.IsDir() {
				if _, err := s.AddProblem(f.Name()); err != nil {
					log.Print(err)
				}
			}
		}

		s.problemsMutex.Unlock()

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

			st, err := judge.NewFromUrl("http://" + j.Host + ":" + j.Port, jwt)

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
	return c.Render(http.StatusInternalServerError, "error.gohtml", msg)
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

			if _, err := s.db.Exec("UPDATE submissions SET verdict=$1, status=$2, ontest=NULL, judged=$3, score=$5 WHERE id=$4", verdict, st.Status, time.Now(), id, st.Status.Score()); err != nil {
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
				server, err := judge.NewFromString(j.State)
				if err != nil {
					log.Print("malformed judge: ", j.State, err)
					continue
				}

				if server.SupportsProblem(sub.Problem) {
					token, err := s.getJWT()
					if err != nil {
						log.Print("can't get jwt token", err)
					}

					if err = server.Submit(judge.Submission{sub.ID, sub.Problem, sub.Language, sub.Source, "http://" + s.Hostname + ":" + s.GluePort + "/callback/" + strconv.Itoa(int(sub.ID))}, token); err != nil {
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

func (s *Server) Run() {
	if s.Mode == "development" {
		boil.DebugMode = true
	}

	//@TODO add a member to Server
	loc, err := time.LoadLocation("Europe/Budapest")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	boil.SetLocation(loc)

	e := echo.New()
	store := sessions.NewCookieStore([]byte(s.CookieSecret))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
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

	if s.GoogleAuth.Enabled {
		goth.UseProviders(
			google.New(s.GoogleAuth.ClientKey, s.GoogleAuth.Secret, s.GoogleAuth.Callback, "email", "profile"),
		)
	}

	t := &Template{
		templates: template.Must(template.New("templater").Funcs(s.templatefuncs()).ParseGlob(filepath.Join(s.TemplatesDir, "*.gohtml"))),
	}

	e.Renderer = t

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

	s.ConnectToDB()
	s.parseKeys()

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
