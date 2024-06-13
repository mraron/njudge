package judge

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	_ "github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/config/polygon"
	_ "github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	_ "github.com/mraron/njudge/pkg/problems/config/task_yaml"
	slogchi "github.com/samber/slog-chi"
	"log/slog"
	"net"
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

func (s Server) PostJudgeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := Submission{}
		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		inited := false
		initResponse := func(statusCode int) {
			if inited {
				return
			}
			inited = true
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
		}
		enc := json.NewEncoder(w)

		st, err := s.Judger.Judge(r.Context(), sub, func(result Result) error {
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
				_ = enc.Encode(res)
				return
			}
			if st == nil {
				initResponse(http.StatusInternalServerError)
				_ = enc.Encode(res)
				return
			}
		}
		initResponse(http.StatusOK)
		_ = enc.Encode(res)
		return
	}
}

func (s Server) Run() error {
	go func() {
		for {
			if err := s.ProblemStore.UpdateProblems(); err != nil {
				s.Logger.Error("failed to update problemStore", "error", err.Error())
			}
			time.Sleep(30 * time.Second)
		}
	}()

	r := chi.NewRouter()

	r.Use(slogchi.New(s.Logger))
	r.Use(middleware.Recoverer)

	r.Post("/judge", s.PostJudgeHandler())

	return http.ListenAndServe(net.JoinHostPort("0.0.0.0", s.Config.Port), r)
}
