package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Submissions struct {
	db *sql.DB
}

func NewSubmissions(db *sql.DB) *Submissions {
	return &Submissions{
		db: db,
	}
}

func (ss *Submissions) toNjudge(ctx context.Context, s *models.Submission) (*njudge.Submission, error) {
	res := &njudge.Submission{
		ID: s.ID,

		UserID:    s.UserID,
		ProblemID: s.ProblemID,
		Language:  s.Language,
		Source:    s.Source,
		Private:   s.Private,

		Started:   s.Started,
		Verdict:   njudge.Verdict(s.Verdict),
		Ontest:    s.Ontest,
		Submitted: s.Submitted,
		Judged:    s.Judged,
		Score:     s.Score.Float32,
	}

	if err := json.Unmarshal([]byte(s.Status), &res.Status); err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *Submissions) Get(ctx context.Context, ID int) (*njudge.Submission, error) {
	dbobj, err := models.Submissions(models.SubmissionWhere.ID.EQ(ID)).One(ctx, ss.db)
	if err != nil {
		return nil, MaskNotFoundError(err, njudge.ErrorSubmissionNotFound)
	}

	res, err := ss.toNjudge(ctx, dbobj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *Submissions) getAll(ctx context.Context, mods ...qm.QueryMod) ([]njudge.Submission, error) {
	dbobjs, err := models.Submissions(mods...).All(ctx, ss.db)
	if err != nil {
		return nil, err
	}

	res := make([]njudge.Submission, len(dbobjs))
	for ind := range dbobjs {
		var curr *njudge.Submission
		curr, err = ss.toNjudge(ctx, dbobjs[ind])
		if err != nil {
			return nil, err
		}

		res[ind] = *curr
	}

	return res, nil
}

func (ss *Submissions) GetAll(ctx context.Context) ([]njudge.Submission, error) {
	return ss.getAll(ctx)
}

func (ss *Submissions) Insert(ctx context.Context, s njudge.Submission) (*njudge.Submission, error) {
	panic("not implemented") // TODO: Implement
}

func (ss *Submissions) Delete(ctx context.Context, ID int) error {
	panic("not implemented") // TODO: Implement
}

func (ss *Submissions) Update(ctx context.Context, s njudge.Submission, fields []string) error {
	panic("not implemented") // TODO: Implement
}
