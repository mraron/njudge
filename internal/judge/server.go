package judge

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	_ "github.com/mraron/njudge/pkg/problems/config/task_yaml"
	slogecho "github.com/samber/slog-echo"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/mraron/njudge/pkg/language/langs/csharp"
	_ "github.com/mraron/njudge/pkg/language/langs/golang"
	_ "github.com/mraron/njudge/pkg/language/langs/java"
	_ "github.com/mraron/njudge/pkg/language/langs/julia"
	_ "github.com/mraron/njudge/pkg/language/langs/pascal"
	_ "github.com/mraron/njudge/pkg/language/langs/pypy3"
	_ "github.com/mraron/njudge/pkg/language/langs/python3"
	_ "github.com/mraron/njudge/pkg/language/langs/zip"
)

type Submission struct {
	ID       string `json:"id"`
	Problem  string `json:"problem"`
	Language string `json:"language"`
	Source   []byte `json:"source"`
	//TODO add status skeleton here?
}

type Result struct {
	Index  int              `json:"index"`
	Test   string           `json:"test"`
	Status *problems.Status `json:"status"`
	Error  string           `json:"error"`
}

type Server struct {
	Logger       *slog.Logger
	Judger       Judger
	ProblemStore problems.Store

	Config struct {
		Port string
	}
}

type ServerOption func(*Server)

func WithPortServerOption(port int) ServerOption {
	return func(server *Server) {
		server.Config.Port = fmt.Sprintf("%d", port)
	}
}

func NewServer(logger *slog.Logger, judger Judger, problemStore problems.Store, opts ...ServerOption) Server {
	res := Server{Logger: logger, Judger: judger, ProblemStore: problemStore}
	res.Config.Port = "8080"
	for _, opt := range opts {
		opt(&res)
	}
	return res
}

func (s Server) PostJudgeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		sub := Submission{}
		if err := c.Bind(&sub); err != nil {
			return err
		}

		inited := false
		initResponse := func(statusCode int) {
			if inited {
				return
			}
			inited = true
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c.Response().WriteHeader(statusCode)
		}
		enc := json.NewEncoder(c.Response().Writer)

		st, err := s.Judger.Judge(c.Request().Context(), sub, func(result Result) error {
			initResponse(http.StatusOK)
			return enc.Encode(result)
		})
		res := Result{
			Status: st,
		}
		if err != nil {
			res.Error = err.Error()
			if errors.Is(err, problems.ErrorProblemNotFound) || errors.Is(err, language.ErrorLanguageNotFound) {
				initResponse(http.StatusBadRequest)
				return enc.Encode(res)
			}
			if st == nil {
				initResponse(http.StatusInternalServerError)
				return enc.Encode(res)
			}
		}
		initResponse(http.StatusOK)
		return enc.Encode(res)
	}
}

func (s Server) Run() error {
	go func() {
		for {
			if err := s.ProblemStore.UpdateProblems(); err != nil {
				s.Logger.Error("failed to update problemStore", err)
			}
			time.Sleep(30 * time.Second)
		}
	}()

	e := echo.New()
	e.Use(slogecho.New(s.Logger))
	e.Use(middleware.Recover())

	e.POST("/judge", s.PostJudgeHandler())

	return e.Start(":" + s.Config.Port)
}

/*
func main() {
	s1, _ := sandbox.NewIsolate(104)
	s2, _ := sandbox.NewIsolate(105)
	provider := sandbox.NewProvider().Put(s1).Put(s2)

	problemStore := problems.NewFsStore("/home/aron/Projects/njudge/njudge_problems_git")
	_ = problemStore.UpdateProblems()
	languageStore := language.DefaultStore

	judge := Judge{
		SandboxProvider: provider,
		ProblemStore:    problemStore,
		LanguageStore:   languageStore,
		RateLimit:       5 * time.Second,
	}

	server := NewServer(slog.Default(), &judge, problemStore, WithPortServerOption(8081))
	go func() {
		fmt.Println("start sleep")
		time.Sleep(5 * time.Second)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(10 * time.Second)
			cancel()
		}()
		client := NewClient("http://localhost:8081")
		res, err := client.Judge(ctx, Submission{
			Problem:  "KK24_csoki2",
			Language: "python3",
			Source:   []byte(`while True: pass`),
		}, func(result Result) error {
			for _, tc := range result.Status.Feedback[0].Testcases() {
				fmt.Print(tc.VerdictName)
			}
			fmt.Println("")
			return nil
		})
		fmt.Println("ends: ", res, err)
	}()
	panic(server.Run())
	/*
	   	fmt.Println(judge.Judge(context.Background(), Submission{
	   		ID:       "",
	   		Problem:  "KK24_csoki22",
	   		Language: "python3",
	   		Source: []byte(`// @check-accepted: examples N=0 no-limits

	   #include <fstream>
	   #include <iostream>
	   #include <vector>

	   using namespace std;

	   int main() {
	       int M, N, K;

	       cin >> M >> N >> K;

	       if ( (N+M)%K == 0 ) {
	           cout << "IGEN" << endl;
	       } else {
	           cout << "NEM" << endl;
	         }

	       return 0;
	   }
	   `),
	   	}, nil))
}
*/
