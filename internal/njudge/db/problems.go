package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/multierr"
)

type Problems struct {
	db *sql.DB
}

func NewProblems(db *sql.DB) *Problems {
	return &Problems{
		db: db,
	}
}

func (ps *Problems) catToNjudge(p *models.ProblemCategory) *njudge.Category {
	if p == nil {
		return nil
	}

	return &njudge.Category{
		ID:       p.ID,
		Name:     p.Name,
		ParentID: p.ParentID,
	}
}

func (ps *Problems) problemTagsToNjudge(pts models.ProblemTagSlice) []njudge.ProblemTag {
	res := make([]njudge.ProblemTag, len(pts))
	for ind := range pts {
		res[ind].Tag = njudge.Tag{
			ID:   pts[ind].R.Tag.ID,
			Name: pts[ind].R.Tag.Name,
		}
		res[ind].ProblemID = pts[ind].ProblemID
		res[ind].UserID = pts[ind].UserID
		res[ind].Added = pts[ind].Added
	}
	return res
}

func (ps *Problems) toNjudge(p *models.ProblemRel) (*njudge.Problem, error) {
	return &njudge.Problem{
		ID:          p.ID,
		Problemset:  p.Problemset,
		Problem:     p.Problem,
		Category:    ps.catToNjudge(p.R.Category),
		SolverCount: p.SolverCount,
		Tags:        ps.problemTagsToNjudge(p.R.ProblemProblemTags),
	}, nil
}

func (ps *Problems) toModel(p njudge.Problem) (*models.ProblemRel, error) {
	res := &models.ProblemRel{
		ID:          p.ID,
		Problemset:  p.Problemset,
		Problem:     p.Problem,
		SolverCount: p.SolverCount,
	}

	if p.Category != nil {
		res.CategoryID = null.IntFrom(p.Category.ID)
	}

	return res, nil
}

func (ps *Problems) get(ctx context.Context, mods ...qm.QueryMod) (*njudge.Problem, error) {
	problem, err := models.ProblemRels(
		append(
			mods,
			qm.Load(models.ProblemRelRels.Category),
			qm.Load("ProblemProblemTags.Tag"),
		)...).One(ctx, ps.db)
	if err != nil {
		return nil, err
	}

	return ps.toNjudge(problem)
}

func (ps *Problems) Get(ctx context.Context, ID int) (*njudge.Problem, error) {
	return ps.get(ctx, models.ProblemRelWhere.ID.EQ(ID))
}

func (ps *Problems) getAll(ctx context.Context, mods ...qm.QueryMod) ([]njudge.Problem, error) {
	problems, err := models.ProblemRels(
		append(
			mods,
			qm.Load(models.ProblemRelRels.Category),
			qm.Load("ProblemProblemTags.Tag"),
		)...).All(ctx, ps.db)
	if err != nil {
		return nil, err
	}

	res := make([]njudge.Problem, len(problems))
	for ind := range problems {
		var curr *njudge.Problem
		curr, err = ps.toNjudge(problems[ind])
		if err != nil {
			return nil, err
		}

		res[ind] = *curr
	}

	return res, nil
}

func (ps *Problems) GetAll(ctx context.Context) ([]njudge.Problem, error) {
	return ps.getAll(ctx)
}

func (ps *Problems) Insert(ctx context.Context, p njudge.Problem) (*njudge.Problem, error) {
	dbobj, err := ps.toModel(p)
	if err != nil {
		return nil, err
	}

	if err = dbobj.Insert(ctx, ps.db, boil.Infer()); err != nil {
		return nil, err
	}

	p.ID = dbobj.ID
	return &p, nil
}

func (ps *Problems) Delete(ctx context.Context, ID int) error {
	_, err := models.ProblemRels(models.ProblemRelWhere.ID.EQ(ID)).DeleteAll(ctx, ps.db)
	return err
}

func (ps *Problems) Update(ctx context.Context, p njudge.Problem, fields []string) error {
	whitelist := []string{}
	updateTags := false
	for _, field := range fields {
		switch field {
		case njudge.ProblemFields.Problemset:
			whitelist = append(whitelist, models.ProblemRelColumns.Problemset)
		case njudge.ProblemFields.Problem:
			whitelist = append(whitelist, models.ProblemRelColumns.Problem)
		case njudge.ProblemFields.Category:
			whitelist = append(whitelist, models.ProblemRelColumns.CategoryID)
		case njudge.ProblemFields.SolverCount:
			whitelist = append(whitelist, models.ProblemRelColumns.SolverCount)
		case njudge.ProblemFields.Tags:
			updateTags = true
		}
	}

	tx, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if !updateTags || len(fields) > 1 {
		updateObj, err := ps.toModel(p)
		if err != nil {
			return multierr.Combine(err, tx.Rollback())
		}

		_, err = updateObj.Update(ctx, tx, boil.Whitelist(whitelist...))
		if err != nil {
			return multierr.Combine(err, tx.Rollback())
		}
	}
	if updateTags {
		updateObj, err := models.ProblemRels(
			models.ProblemRelWhere.ID.EQ(p.ID),
			qm.Load("ProblemProblemTags.Tag"),
		).One(ctx, tx)
		if err != nil {
			return multierr.Combine(err, tx.Rollback())
		}
		for _, oldTag := range updateObj.R.ProblemProblemTags {
			if _, err = oldTag.Delete(ctx, tx); err != nil {
				return multierr.Combine(err, tx.Rollback())
			}
		}
		for _, newTag := range p.Tags {
			ptag := &models.ProblemTag{
				ProblemID: p.ID,
				UserID:    newTag.UserID,
				Added:     newTag.Added,
				TagID:     newTag.Tag.ID,
			}

			if err = ptag.Insert(ctx, tx, boil.Infer()); err != nil {
				return multierr.Combine(err, tx.Rollback())
			}
		}
	}

	return tx.Commit()
}

func (ps *Problems) GetProblem(ctx context.Context, problemset string, problem string) (*njudge.Problem, error) {
	return ps.get(ctx, models.ProblemRelWhere.Problemset.EQ(problemset), models.ProblemRelWhere.Problem.EQ(problem))
}

func (ps *Problems) GetProblemsWithCategory(ctx context.Context, f njudge.CategoryFilter) ([]njudge.Problem, error) {
	switch f.Type {
	case njudge.CategoryFilterID:
		return ps.getAll(ctx, models.ProblemRelWhere.CategoryID.EQ(null.IntFrom(f.Value.(int))))
	case njudge.CategoryFilterEmpty:
		return ps.getAll(ctx, models.ProblemRelWhere.CategoryID.IsNull())
	default: //case njudge.CategoryFilterNone:
		return ps.getAll(ctx)
	}
}

func (ps *Problems) hasUserSolved(ctx context.Context, userID, problemID int) (njudge.SolvedStatus, error) {
	solvedStatus := njudge.Unattempted

	cnt, err := models.Submissions(
		models.SubmissionWhere.ProblemID.EQ(problemID),
		models.SubmissionWhere.Verdict.EQ(int(problems.VerdictAC)),
		models.SubmissionWhere.UserID.EQ(userID),
	).Count(ctx, ps.db)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return njudge.Unknown, err
	} else {
		if cnt > 0 {
			solvedStatus = njudge.Solved
		} else {
			cnt, err := models.Submissions(
				models.SubmissionWhere.ProblemID.EQ(problemID),
				models.SubmissionWhere.UserID.EQ(userID),
			).Count(ctx, ps.db)

			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return njudge.Unknown, err
			} else if cnt > 0 {
				solvedStatus = njudge.Attempted
			}
		}
	}

	return solvedStatus, nil
}

func (ps *Problems) getUserLastLanguage(ctx context.Context, userID int) (string, error) {
	if userID > 0 {
		sub, err := models.Submissions(
			qm.Select(models.SubmissionColumns.Language),
			qm.OrderBy("id DESC"),
			qm.Limit(1),
		).One(ctx, ps.db)

		if err != nil && errors.Is(err, sql.ErrNoRows) {
			return "", nil
		} else if err != nil {
			return "", err
		}

		return sub.Language, nil
	}

	return "", nil
}

func (ps *Problems) GetProblemData(ctx context.Context, problemID int, userID int) (*njudge.ProblemInfo, error) {
	var (
		res njudge.ProblemInfo
		err error
	)

	if userID > 0 {
		res.UserInfo = &njudge.ProblemUserInfo{}

		if res.UserInfo.SolvedStatus, err = ps.hasUserSolved(ctx, userID, problemID); err != nil {
			return nil, err
		}

		if res.UserInfo.LastLanguage, err = ps.getUserLastLanguage(ctx, userID); err != nil {
			return nil, err
		}

		if res.UserInfo.Submissions, err = NewSubmissions(ps.db).getAll(ctx,
			models.SubmissionWhere.ProblemID.EQ(problemID),
			models.SubmissionWhere.UserID.EQ(userID),
			qm.OrderBy("id DESC"),
			qm.Limit(5),
		); err != nil {
			return nil, err
		}
	}

	return &res, nil
}
