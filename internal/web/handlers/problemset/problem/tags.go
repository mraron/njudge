package problem

import (
	"errors"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/models"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func PostTag(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.FormValue("tagID"))
		if err != nil {
			return err
		}

		name, problem := c.Param("name"), c.Param("problem")
		rel, err := models.ProblemRels(Where("problemset=?", name), Where("problem=?", problem)).One(DB)
		if err != nil {
			return err
		}

		if rel == nil {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("no such problem in problemset"))
		}

		u := c.Get("user").(*models.User)
		if u == nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		st, err := helpers.HasUserSolved(DB, u, name, problem)
		if err != nil {
			return err
		}

		if st != helpers.Solved {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		tag := models.ProblemTag{
			TagID:     id,
			ProblemID: rel.ID,
			Added:     time.Now(),
			UserID:    u.ID,
		}

		if err := tag.Insert(DB, boil.Infer()); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, filepath.Dir(c.Request().URL.Path)+"/")
	}
}

func DeleteTag(DB *sqlx.DB, problemStore problems.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		name, problem := c.Param("name"), c.Param("problem")

		u := c.Get("user").(*models.User)
		if u == nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		st, err := helpers.HasUserSolved(DB, u, name, problem)
		if err != nil {
			return err
		}

		if st != helpers.Solved {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		rel, err := models.ProblemRels(Where("problemset=?", name), Where("problem=?", problem)).One(DB)
		if err != nil {
			return err
		}

		if rel == nil {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("no such problem in problemset"))
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		if cnt, err := models.ProblemTags(Where("problem_id=?", rel.ID), Where("tag_id = ?", id)).DeleteAll(DB); err != nil {
			return err
		} else if cnt == 0 {
			return c.JSON(http.StatusNotFound, nil)
		}

		return c.Redirect(http.StatusFound, filepath.Dir(filepath.Dir(c.Request().URL.Path))+"/")
	}
}
