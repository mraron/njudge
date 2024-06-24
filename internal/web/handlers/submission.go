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

func GetSubmission(s njudge.Submissions, probs njudge.Problems, psets njudge.Problemsets, solvedStatus njudge.SolvedStatusQuery) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get(templates.UserContextKey).(*njudge.User)
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
			p, err := sub.GetProblem(c.Request().Context(), probs)
			if err != nil {
				return err
			}
			pset, err := psets.GetByName(c.Request().Context(), p.Problemset)
			if err != nil {
				return err
			}
			if u != nil && (u.Role == "admin" || u.ID == sub.UserID) {
				vm.DisplaySource = true
			} else {
				switch pset.CodeVisibility {
				case njudge.CodeVisibilityPrivate:
				case njudge.CodeVisibilitySolved:
					if u != nil {
						ss, err := solvedStatus.GetSolvedStatus(c.Request().Context(), p.ID, u.ID)
						if err != nil {
							return err
						}
						if ss == njudge.Solved {
							vm.DisplaySource = true
						}
					}
				case njudge.CodeVisibilityPublic:
					vm.DisplaySource = true
				}
			}
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
		u := c.Get(templates.UserContextKey).(*njudge.User)
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
