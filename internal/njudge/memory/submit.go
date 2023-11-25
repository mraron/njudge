package memory

import (
	"context"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
)

type SubmitService struct {
	submissions  njudge.Submissions
	users        njudge.Users
	problemQuery njudge.ProblemQuery
	problemStore problems.Store
}

func NewSubmitService(submissions njudge.Submissions, users njudge.Users, problemQuery njudge.ProblemQuery, problemStore problems.Store) *SubmitService {
	return &SubmitService{
		submissions:  submissions,
		users:        users,
		problemQuery: problemQuery,
		problemStore: problemStore,
	}
}

func (s *SubmitService) Submit(ctx context.Context, req njudge.SubmitRequest) (*njudge.Submission, error) {
	pr, err := s.problemQuery.GetProblem(ctx, req.Problemset, req.Problem)
	if err != nil {
		return nil, err
	}

	sdata, err := pr.WithStoredData(s.problemStore)
	if err != nil {
		return nil, err
	}

	var (
		found     = false
		foundLang language.Language
	)
	for _, lang := range sdata.Languages() {
		if lang.Id() == req.Language {
			found = true
			foundLang = lang
			break
		}
	}

	if !found {
		return nil, njudge.ErrorUnsupportedLanguage
	}

	u, err := s.users.Get(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	sub, err := njudge.NewSubmission(*u, *pr, foundLang)
	if err != nil {
		return nil, err
	}

	sub.SetSource(req.Source)
	sub.Verdict = njudge.VerdictUP

	return s.submissions.Insert(ctx, *sub)
}
