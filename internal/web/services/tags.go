package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

type TagsService interface {
	Add(ctx context.Context, tagID int, problemID int, userID int) error
	Delete(ctx context.Context, tagID int, userID int) error
}

var (
	ErrorUnableToModifyTags = errors.New("user can't modify tags")
	ErrorNoSuchTag          = errors.New("no such tag")
)

type SQLTagsService struct {
	db *sql.DB
}

func NewSQLTagsService(db *sql.DB) *SQLTagsService {
	return &SQLTagsService{db}
}

func (tg SQLTagsService) Add(ctx context.Context, tagID int, problemID int, userID int) error {
	pr, err := models.ProblemRels(models.ProblemRelWhere.ID.EQ(problemID)).One(ctx, tg.db)
	if err != nil {
		return err
	}

	st, err := helpers.HasUserSolved(tg.db, userID, pr.Problemset, pr.Problem)
	if err != nil {
		return err
	}

	if st != problem.Solved {
		return ErrorUnableToModifyTags
	}

	tag := models.ProblemTag{
		TagID:     tagID,
		ProblemID: problemID,
		Added:     time.Now(),
		UserID:    userID,
	}

	return tag.Insert(ctx, tg.db, boil.Whitelist(models.ProblemTagColumns.TagID,
		models.ProblemTagColumns.ProblemID, models.ProblemTagColumns.Added, models.ProblemTagColumns.UserID))
}

func (tg SQLTagsService) Delete(ctx context.Context, tagID int, userID int) error {
	tag, err := models.ProblemTags(models.ProblemTagWhere.ID.EQ(tagID)).One(ctx, tg.db)
	if err != nil {
		return err
	}

	pr, err := tag.Problem().One(ctx, tg.db)
	if err != nil {
		return err
	}

	st, err := helpers.HasUserSolved(tg.db, userID, pr.Problemset, pr.Problem)
	if err != nil {
		return err
	}

	if st != problem.Solved {
		return ErrorUnableToModifyTags
	}

	if cnt, err := models.ProblemTags(models.ProblemTagWhere.ProblemID.EQ(pr.ID), models.ProblemTagWhere.TagID.EQ(tagID)).DeleteAll(ctx, tg.db); err != nil {
		return err
	} else if cnt == 0 {
		return ErrorNoSuchTag
	}

	return nil
}
