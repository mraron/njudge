package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/roles"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
)

func GetSubmission(r submission.Repository) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}

		sub, err := r.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}

		c.Set("title", tr.Translate("Submission #%d", data.ID))
		return c.Render(http.StatusOK, "submission.gohtml", sub)
	}
}

func RejudgeSubmission(rs services.RejudgeService) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "submissions/rejudge") {
			return helpers.UnauthorizedError(c)
		}

		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}

		if err := rs.Rejudge(c.Request().Context(), data.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getSubmission", data.ID))
	}
}
