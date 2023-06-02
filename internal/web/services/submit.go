package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

var ErrorUnsupportedLanguage = errors.New("unsupported language")

type SubmitRequest struct {
	UserID     int
	Problemset string
	Problem    string
	Language   string
	Source     []byte
}

type SubmitService interface {
	Submit(ctx context.Context, subRequest SubmitRequest) (*submission.Submission, error)
}

type SQLSubmitService struct {
	db           *sql.DB
	problemStore problems.Store
}

func NewSQLSubmitService(db *sql.DB, problemStore problems.Store) *SQLSubmitService {
	return &SQLSubmitService{
		db:           db,
		problemStore: problemStore,
	}
}

func (s *SQLSubmitService) Submit(ctx context.Context, req SubmitRequest) (*submission.Submission, error) {
	p, err := s.problemStore.Get(req.Problem)
	if err != nil {
		return nil, err
	}

	found := false
	for _, lang := range p.Languages() {
		if lang.Id() == req.Language {
			found = true
			break
		}
	}

	if !found {
		return nil, ErrorUnsupportedLanguage
	}

	sub := models.Submission{
		UserID:     req.UserID,
		Problemset: req.Problemset,
		Problem:    req.Problem,
		Language:   req.Language,
		Source:     req.Source,
		Status:     "{}",
		Verdict:    int(problem.VerdictUP),
		Submitted:  time.Now(),
		Private:    false,
		Started:    false,
	}
	if err := sub.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &submission.Submission{Submission: sub}, nil
}
