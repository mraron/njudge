package glue

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge/db"
	"io"
	"log/slog"
	"strconv"
	"time"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"

	"github.com/volatiletech/null/v8"
)

type Glue struct {
	Judge judge.Judger

	Logger *slog.Logger

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
		glue.Submissions = db.NewSubmissions(conn)
		glue.Problems = db.NewProblems(conn, db.NewSolvedStatusQuery(conn))
		glue.SubmissionsQuery = glue.Submissions.(njudge.SubmissionsQuery)
		return nil
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(glue *Glue) error {
		glue.Logger = logger.With("service", "glue")
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
	g.Logger.Info("submission started", "id", sub.ID)

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
			fmt.Sprintf("callback %d received for submission", result.Index),
			"id", sub.ID,
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
	if !status.Compiled {
		verdict = problems.VerdictName(njudge.VerdictCE)
	} else {
		verdict = status.Feedback[0].Verdict()
		score = float32(status.Feedback[0].Score())
	}

	g.Logger.Info("finished judging submission", "id", sub.ID)

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
	for {
		g.Logger.Info("looking for submissions")
		subs, err := g.SubmissionsQuery.GetUnstarted(ctx, 10)
		if err != nil {
			g.Logger.Error("looking for submissions", "error", err)
			continue
		}

		for _, s := range subs {
			s := s
			go func() {
				// create some kind of token system
				// and also a collection of judges
				err := g.ProcessSubmission(ctx, s)
				if err != nil {
					g.Logger.Error("processing submission", "id", s.ID, "error", err)
					return
				}
			}()
		}

		time.Sleep(5 * time.Second)
	}
}
