package glue

import (
	"context"
	"fmt"
	"github.com/mraron/njudge/internal/judge2"
	"github.com/mraron/njudge/internal/njudge/db"
	"github.com/mraron/njudge/internal/web/helpers/config"
	"strconv"
	"time"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/problems"

	"github.com/volatiletech/null/v8"
)

type Glue struct {
	Judge judge2.Judger

	Submissions      njudge.Submissions
	Problems         njudge.Problems
	SubmissionsQuery njudge.SubmissionsQuery
}

type Option func(*Glue) error

func WithDatabaseOption(cfg config.Database) Option {
	return func(glue *Glue) error {
		conn, err := cfg.Connect()
		if err != nil {
			return err
		}
		if err = conn.Ping(); err != nil {
			return err
		}
		glue.Submissions = db.NewSubmissions(conn)
		glue.Problems = db.NewProblems(conn, db.NewSolvedStatusQuery(conn))
		glue.SubmissionsQuery = glue.Submissions.(njudge.SubmissionsQuery)
		return nil
	}
}

func New(judge judge2.Judger, opts ...Option) (*Glue, error) {
	glue := &Glue{
		Judge: judge,
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

	prob, err := g.Problems.Get(ctx, sub.ProblemID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	status, err := g.Judge.Judge(ctx, judge2.Submission{
		ID:       strconv.Itoa(sub.ID),
		Problem:  prob.Problem,
		Language: sub.Language,
		Source:   sub.Source,
	}, func(result judge2.Result) error {
		if result.Status == nil {
			return fmt.Errorf("received nil status, error: %v", result.Error)
		}

		sub := njudge.Submission{
			ID:      sub.ID,
			Verdict: njudge.VerdictRU,
			Status:  *result.Status,
			Ontest:  null.NewString(result.Test, true),
		}

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
		subs, err := g.SubmissionsQuery.GetUnstarted(ctx, 10)
		if err != nil {
			fmt.Println(err)
			// log it
			continue
		}

		for _, s := range subs {
			s := s
			go func() {
				// create some kind of token system
				// and also a collection of judges
				err := g.ProcessSubmission(ctx, s)
				if err != nil {
					fmt.Println(err)
					return
				}
			}()
		}

		time.Sleep(5 * time.Second)
	}
}
