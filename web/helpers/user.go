package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HasUserSolved(DB *sqlx.DB, u *models.User, problemSet, problem string) (int, error) {
	solvedStatus := -1
	if u != nil {
		cnt, err := models.Submissions(Where("problemset = ?", problemSet), Where("problem = ?", problem), Where("verdict = 0"), Where("user_id = ?", u.ID)).Count(DB)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return -1, fmt.Errorf("can't get solvedstatus for %s %s%s: %w", u.Name, problemSet, problem, err)
		}else {
			if cnt > 0 {
				solvedStatus = 0
			} else {
				cnt, err := models.Submissions(Where("problemset = ?", problemSet), Where("problem = ?", problem), Where("user_id = ?", u.ID)).Count(DB)
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					return -1, fmt.Errorf("can't get solvedstatus for %s %s %s: %w", u.Name, problemSet, problem, err)
				} else {
					if cnt>0 {
						solvedStatus = 1
					}
				}
			}
		}
	}

	return solvedStatus, nil
}
