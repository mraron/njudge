package helpers

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/pkg/web/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SolvedStatus int

const (
	Unattempted SolvedStatus = iota
	Attempted
	PartiallySolved
	Solved
	Unknown
)

func HasUserSolved(DB *sqlx.DB, u *models.User, problemSet, problem string) (SolvedStatus, error) {
	solvedStatus := Unattempted
	if u != nil {
		cnt, err := models.Submissions(Where("problemset = ?", problemSet), Where("problem = ?", problem), Where("verdict = 0"), Where("user_id = ?", u.ID)).Count(DB)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return Unknown, fmt.Errorf("can't get solvedstatus for %s %s%s: %w", u.Name, problemSet, problem, err)
		} else {
			if cnt > 0 {
				solvedStatus = Solved
			} else {
				cnt, err := models.Submissions(Where("problemset = ?", problemSet), Where("problem = ?", problem), Where("user_id = ?", u.ID)).Count(DB)
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					return Unknown, fmt.Errorf("can't get solvedstatus for %s %s %s: %w", u.Name, problemSet, problem, err)
				} else {
					if cnt > 0 {
						solvedStatus = Attempted
					}
				}
			}
		}
	}

	return solvedStatus, nil
}
