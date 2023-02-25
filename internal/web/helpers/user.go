package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/jmoiron/sqlx"
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

func GetUserLastLanguage(c echo.Context, DB *sqlx.DB) string {
	if res := c.Get("last_language"); res != nil {
		return c.Get("last_language").(string)
	}

	res := ""
	if u := c.Get("user").(*models.User); u != nil {
		sub, err := models.Submissions(Select("language"), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(1)).One(DB)
		if err == nil {
			c.Set("last_language", sub.Language)
			res = sub.Language
		}
	}

	return res
}