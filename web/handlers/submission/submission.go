package submission

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
)

func Get(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		val, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		sub, err := models.Submissions(Where("id=?", val)).One(DB)
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "submission.gohtml", sub)
	}
}

func Rejudge(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		sub, err := models.Submissions(Where("id = ?", id)).One(DB)
		if err != nil {
			return err
		}

		sub.ID = id
		sub.Judged = null.Time{Valid: false}
		sub.Started = false
		if _, err := sub.Update(DB, boil.Infer()); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/submission/"+strconv.Itoa(id))
	}
}