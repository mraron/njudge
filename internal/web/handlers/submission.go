package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/roles"
)

func GetSubmission(s njudge.Submissions) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}

		sub, err := s.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}

		c.Set("title", tr.Translate("Submission #%d", data.ID))
		return c.Render(http.StatusOK, "submission.gohtml", sub)
	}
}

func RejudgeSubmission(s njudge.Submissions) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "submissions/rejudge") {
			return helpers.UnauthorizedError(c)
		}

		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}

		sub, err := s.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}

		sub.MarkForRejudge()
		if err := s.Update(c.Request().Context(), *sub, njudge.SubmissionRejudgeFields); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getSubmission", data.ID))
	}
}
