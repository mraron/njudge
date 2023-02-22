package problem

import (
	"errors"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrorUnableToModifyTags = errors.New("user can't modify tags")
	ErrorNoSuchTag          = errors.New("no such tag")
)

type TagManager struct {
	Problem *Problem
}

func (tg TagManager) CreateTag(DB *sqlx.DB, tagID int, user *models.User) error {
	st, err := helpers.HasUserSolved(DB, user, tg.Problem.ProblemRel.Problemset, tg.Problem.ProblemRel.Problem)
	if err != nil {
		return err
	}

	if st != helpers.Solved {
		return ErrorUnableToModifyTags
	}

	tag := models.ProblemTag{
		TagID:     tagID,
		ProblemID: tg.Problem.ProblemRel.ID,
		Added:     time.Now(),
		UserID:    user.ID,
	}

	return tag.Insert(DB, boil.Infer())
}

func (tg TagManager) DeleteTag(DB *sqlx.DB, tagID int, user *models.User) error {
	st, err := helpers.HasUserSolved(DB, user, tg.Problem.ProblemRel.Problemset, tg.Problem.ProblemRel.Problem)
	if err != nil {
		return err
	}

	if st != helpers.Solved {
		return ErrorUnableToModifyTags
	}

	if cnt, err := models.ProblemTags(Where("problem_id=?", tg.Problem.ProblemRel.ID), Where("tag_id = ?", tagID)).DeleteAll(DB); err != nil {
		return err
	} else if cnt == 0 {
		return ErrorNoSuchTag
	}

	return nil
}
