package glue

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge/db"
	"io"
	"log/slog"
	"math"
	"strconv"
	"time"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"

	"github.com/volatiletech/null/v8"
)

type Glue struct {
	Judge judge.Judger

	Logger *slog.Logger

	Users            njudge.Users
	Submissions      njudge.Submissions
	Problems         njudge.Problems
	SubmissionsQuery njudge.SubmissionsQuery
}

type Option func(*Glue) error

func WithDatabaseOption(conn *sql.DB) Option {
	return func(glue *Glue) error {
		if err := conn.Ping(); err != nil {
			return err
		}
		glue.Users = db.NewUsers(conn)
		glue.Submissions = db.NewSubmissions(conn)
		glue.Problems = db.NewProblems(conn, db.NewSolvedStatusQuery(conn))
		glue.SubmissionsQuery = glue.Submissions.(njudge.SubmissionsQuery)
		return nil
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(glue *Glue) error {
		glue.Logger = logger
		return nil
	}
}

func New(judge judge.Judger, opts ...Option) (*Glue, error) {
	glue := &Glue{
		Judge:  judge,
		Logger: slog.New(slog.NewJSONHandler(io.Discard, nil)),
	}
	for _, opt := range opts {
		if err := opt(glue); err != nil {
			return nil, err
		}
	}
	return glue, nil
}

func (g *Glue) ProcessSubmission(ctx context.Context, sub njudge.Submission) error {
	sub.Started = true
	if err := g.Submissions.Update(
		ctx,
		sub,
		njudge.Fields(njudge.SubmissionFields.Started),
	); err != nil {
		return err
	}
	g.Logger.Info("🟢\tstarted submission", "submission_id", sub.ID)

	prob, err := g.Problems.Get(ctx, sub.ProblemID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	status, err := g.Judge.Judge(ctx, judge.Submission{
		ID:       strconv.Itoa(sub.ID),
		Problem:  prob.Problem,
		Language: sub.Language,
		Source:   sub.Source,
	}, func(result judge.Result) error {
		if result.Status == nil {
			return fmt.Errorf("received nil status, error: %v", result.Error)
		}

		sub := njudge.Submission{
			ID:      sub.ID,
			Verdict: njudge.VerdictRU,
			Status:  *result.Status,
			Ontest:  null.NewString(result.Test, true),
		}
		g.Logger.Info(
			fmt.Sprintf("↪️\tcallback %d received", result.Index),
			"submission_id", sub.ID,
		)

		return g.Submissions.Update(ctx, sub, njudge.Fields(
			njudge.SubmissionFields.Verdict,
			njudge.SubmissionFields.Status,
			njudge.SubmissionFields.Ontest,
		))
	})
	defer cancel()
	if err != nil {
		return err
	}
	var (
		verdict problems.VerdictName
		score   float32 = 0.0
	)
	if status.CompilationStatus != problems.AfterCompilation {
		return errors.New("invalid compilation status after judging submission")
	}
	if !status.Compiled {
		verdict = problems.VerdictName(njudge.VerdictCE)
	} else {
		verdict = status.Feedback[0].Verdict()
		score = float32(status.Feedback[0].Score())
	}

	g.Logger.Info("🏁\tfinished judging", "submission_id", sub.ID)

	sub.Verdict = njudge.Verdict(verdict)
	sub.Status = *status
	sub.Ontest = null.NewString("", false)
	sub.Judged = null.NewTime(time.Now(), true)
	sub.Score = score

	return g.Submissions.Update(ctx, sub, njudge.Fields(
		njudge.SubmissionFields.Verdict,
		njudge.SubmissionFields.Status,
		njudge.SubmissionFields.Ontest,
		njudge.SubmissionFields.Judged,
		njudge.SubmissionFields.Score,
	))
}

func (g *Glue) Start(ctx context.Context) {
	go func() {
		for {
			var (
				problemList []njudge.Problem
				subs        []njudge.Submission
				err         error
			)

			g.Logger.Info("🖩\tcalculating user points")
			problemList, err = g.Problems.GetAll(ctx)
			if err != nil {
				g.Logger.ErrorContext(ctx, err.Error())
				continue
			}
			g.Logger.Info(fmt.Sprintf("🔎\tfound %d problems", len(problemList)))
			userPoints := make(map[int]float64)
			for _, problem := range problemList {
				subs, err = g.SubmissionsQuery.GetACSubmissionsOf(ctx, problem.ID)
				if err != nil {
					break
				}
				userSolved := make(map[int]struct{})
				users := make([]int, 0)
				for _, sub := range subs {
					userSolved[sub.UserID] = struct{}{}
					users = append(users, sub.UserID)
				}
				solvedBy := len(users)
				if solvedBy > 0 {
					points := math.Sqrt(1.0 / float64(solvedBy))

					for _, uid := range users {
						userPoints[uid] += points
					}
				}

				problem.SolverCount = solvedBy
				if err = g.Problems.Update(ctx, problem, njudge.Fields(njudge.ProblemFields.SolverCount)); err != nil {
					break
				}
			}
			if err != nil {
				g.Logger.ErrorContext(ctx, err.Error())
				continue
			}
			g.Logger.Info(fmt.Sprintf("👨\tassigning %d users points", len(userPoints)))
			for uid, points := range userPoints {
				user := njudge.User{ID: uid}
				user.Points = float32(points)
				if err = g.Users.Update(ctx, &user, njudge.Fields(njudge.UserFields.Points)); err != nil {
					break
				}
			}
			if err != nil {
				g.Logger.ErrorContext(ctx, err.Error())
				continue
			}
			g.Logger.Info("✔️\tsuccessfully assigned solver count and user points")
			time.Sleep(5 * time.Minute)
		}
	}()

	for {
		g.Logger.Info("🔎\tlooking for submissions")
		subs, err := g.SubmissionsQuery.GetUnstarted(ctx, 20)
		if err != nil {
			g.Logger.Error("‼️\tlooking for submissions", "error", err)
			continue
		}

		for _, s := range subs {
			s := s
			go func() {
				// TODO create some kind of token system and also a collection of judges
				err := g.ProcessSubmission(ctx, s)
				if err != nil {
					g.Logger.Error("‼️\tprocessing submission", "submission_id", s.ID, "error", err)

					s.Verdict = njudge.VerdictXX
					s.Judged = null.NewTime(time.Now(), true)
					_ = g.Submissions.Update(ctx, s, njudge.Fields(njudge.SubmissionFields.Verdict, njudge.SubmissionFields.Judged, njudge.SubmissionFields.Status))
					return
				}
			}()
		}

		time.Sleep(5 * time.Second)
	}
}
