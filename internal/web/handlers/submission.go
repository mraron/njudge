package handlers

import (
	"github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/internal/web/templates/i18n"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/roles"
)

type GetSubmissionRequest struct {
	ID int `param:"id"`
}

func GetSubmission(s njudge.Submissions) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := &GetSubmissionRequest{}
		if err := c.Bind(data); err != nil {
			return err
		}

		sub, err := s.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}

		vm := templates.SubmissionViewModel{
			Submission: *sub,
		}
		if u != nil && roles.Can(roles.Role(u.Role), roles.ActionCreate, "submissions/rejudge") {
			vm.CanRejudge = true
		}
		if sub.Language != "zip" {
			vm.DisplaySource = true
		}

		c.Set(templates.TitleContextKey, tr.Translate("Submission #%d", data.ID))
		return templates.Render(c, http.StatusOK, templates.Submission(vm))
	}
}

func RejudgeSubmission(s njudge.Submissions) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "submissions/rejudge") {
			return c.NoContent(http.StatusUnauthorized)
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
