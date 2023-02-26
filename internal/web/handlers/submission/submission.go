package submission

import (
	"fmt"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/roles"
	"github.com/mraron/njudge/internal/web/models"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
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

		c.Set("title", fmt.Sprintf("Beküldés #%d", val))
		return c.Render(http.StatusOK, "submission.gohtml", sub)
	}
}

func Rejudge(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "submissions/rejudge") {
			return helpers.UnauthorizedError(c)
		}

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
