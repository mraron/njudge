package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (ss *Submissions) toModel(ctx context.Context, s njudge.Submission) (*models.Submission, error) {
	res := &models.Submission{
		ID: s.ID,

		UserID:    s.UserID,
		ProblemID: s.ProblemID,
		Language:  s.Language,
		Source:    s.Source,
		Private:   s.Private,

		Started:   s.Started,
		Verdict:   int(s.Verdict),
		Ontest:    s.Ontest,
		Submitted: s.Submitted,
		Judged:    s.Judged,
		Score:     null.Float32From(s.Score),
	}

	var (
		statusBytes []byte
		err         error
	)

	if statusBytes, err = json.Marshal(s.Status); err != nil {
		return nil, err
	}
	res.Status = string(statusBytes)

	return res, nil
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
	dbobj, err := ss.toModel(ctx, s)
	if err != nil {
		return nil, err
	}

	if err := dbobj.Insert(ctx, ss.db, boil.Infer()); err != nil {
		return nil, err
	}

	s.ID = dbobj.ID
	return &s, nil
}

func (ss *Submissions) Delete(ctx context.Context, ID int) error {
	_, err := models.Submissions(models.SubmissionWhere.ID.EQ(ID)).DeleteAll(ctx, ss.db)
	return err
}

func (ss *Submissions) Update(ctx context.Context, s njudge.Submission, fields []string) error {
	whitelist := []string{}
	for _, field := range fields {
		switch field {
		case njudge.SubmissionFields.UserID:
			whitelist = append(whitelist, models.SubmissionColumns.UserID)
		case njudge.SubmissionFields.ProblemID:
			whitelist = append(whitelist, models.SubmissionColumns.ProblemID)
		case njudge.SubmissionFields.Language:
			whitelist = append(whitelist, models.SubmissionColumns.Language)
		case njudge.SubmissionFields.Source:
			whitelist = append(whitelist, models.SubmissionColumns.Source)
		case njudge.SubmissionFields.Private:
			whitelist = append(whitelist, models.SubmissionColumns.Private)

		case njudge.SubmissionFields.Started:
			whitelist = append(whitelist, models.SubmissionColumns.Started)
		case njudge.SubmissionFields.Verdict:
			whitelist = append(whitelist, models.SubmissionColumns.Verdict)
		case njudge.SubmissionFields.Ontest:
			whitelist = append(whitelist, models.SubmissionColumns.Ontest)
		case njudge.SubmissionFields.Submitted:
			whitelist = append(whitelist, models.SubmissionColumns.Submitted)
		case njudge.SubmissionFields.Judged:
			whitelist = append(whitelist, models.SubmissionColumns.Judged)
		case njudge.SubmissionFields.Score:
			whitelist = append(whitelist, models.SubmissionColumns.Score)
		case njudge.SubmissionFields.Status:
			whitelist = append(whitelist, models.SubmissionColumns.Status)
		}
	}

	dbobj, err := ss.toModel(ctx, s)
	if err != nil {
		return err
	}

	_, err = dbobj.Update(ctx, ss.db, boil.Whitelist(whitelist...))
	return err
}
