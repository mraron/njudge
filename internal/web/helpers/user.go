package helpers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"

	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HasUserSolved(DB *sql.DB, userID int, problemSet, problemName string) (problem.SolvedStatus, error) {
	solvedStatus := problem.Unattempted

	cnt, err := models.Submissions(models.SubmissionWhere.Problemset.EQ(problemSet), models.SubmissionWhere.Problem.EQ(problemName),
		models.SubmissionWhere.Verdict.EQ(int(problems.VerdictAC)), models.SubmissionWhere.UserID.EQ(userID)).Count(context.TODO(), DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return problem.Unknown, fmt.Errorf("can't get solvedstatus for %d %s%s: %w", userID, problemSet, problemName, err)
	} else {
		if cnt > 0 {
			solvedStatus = problem.Solved
		} else {
			cnt, err := models.Submissions(models.SubmissionWhere.Problemset.EQ(problemSet), models.SubmissionWhere.Problem.EQ(problemName),
				models.SubmissionWhere.UserID.EQ(userID)).Count(context.TODO(), DB)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return problem.Unknown, fmt.Errorf("can't get solvedstatus for %d %s %s: %w", userID, problemSet, problemName, err)
			} else {
				if cnt > 0 {
					solvedStatus = problem.Attempted
				}
			}
		}
	}

	return solvedStatus, nil
}

func GetUserLastLanguage(ctx context.Context, DB *sql.DB, userID int) (string, error) {
	if userID > 0 {
		sub, err := models.Submissions(Select("language"), Where("user_id = ?", userID), OrderBy("id DESC"), Limit(1)).One(ctx, DB)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			return "", nil
		} else if err != nil {
			return "", err
		}

		return sub.Language, nil
	}
	return "", nil
}
