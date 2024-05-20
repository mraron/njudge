package profile

import (
	"github.com/mraron/njudge/internal/web/templates"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
)

func GetProfile(sublist njudge.SubmissionListQuery) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("profile").(*njudge.User)

		solved, err := sublist.GetSolvedSubmissionList(c.Request().Context(), u.ID)
		if err != nil {
			return err
		}

		attempted, err := sublist.GetAttemptedSubmissionList(c.Request().Context(), u.ID)
		if err != nil {
			return err
		}

		c.Set("title", tr.Translate("%s's profile", u.Name))
		return c.Render(http.StatusOK, "user/profile/main", struct {
			User                       *njudge.User
			SolvedProblems             []njudge.Submission
			AttemptedNotSolvedProblems []njudge.Submission
		}{u, solved.Submissions, attempted.Submissions})
	}
}

func GetSubmissions(sublist njudge.SubmissionListQuery) echo.HandlerFunc {
	type request struct {
		Page int `query:"page"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("profile").(*njudge.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if data.Page <= 0 {
			data.Page = 1
		}

		statusReq := njudge.SubmissionListRequest{
			Page:      data.Page,
			PerPage:   20,
			SortDir:   njudge.SortDESC,
			SortField: njudge.SubmissionSortFieldID,
			UserID:    u.ID,
		}

		submissionList, err := sublist.GetPagedSubmissionList(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		qu := (*c.Request().URL).Query()
		links, err := templates.Links(submissionList.PaginationData.Page, submissionList.PaginationData.PerPage, int64(submissionList.PaginationData.Count), qu)
		if err != nil {
			return err
		}

		result := templates.SubmissionsViewModel{
			Submissions: submissionList.Submissions,
			Pages:       links,
		}

		c.Set("title", tr.Translate("%s's submissions", u.Name))
		return c.Render(http.StatusOK, "user/profile/submissions", struct {
			User       *njudge.User
			StatusPage templates.SubmissionsViewModel
		}{u, result})
	}
}

func GetSettings() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)

		templates.DeleteFlash(c, "ChangePassword")
		return c.Render(http.StatusOK, "user/profile/settings", struct {
			User *njudge.User
		}{u})
	}
}

func PostSettingsChangePassword(us njudge.Users) echo.HandlerFunc {
	type request struct {
		PasswordOld  string `form:"passwordOld"`
		PasswordNew1 string `form:"passwordNew1"`
		PasswordNew2 string `form:"passwordNew2"`
	}
	return func(c echo.Context) error {
		tr := c.Get("translator").(i18n.Translator)
		u := c.Get("user").(*njudge.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.PasswordOld)); err != nil {
			templates.SetFlash(c, "ChangePassword", tr.Translate("Wrong old password."))
			return c.Redirect(http.StatusFound, "../")
		}

		if len(data.PasswordNew1) == 0 {
			templates.SetFlash(c, "ChangePassword", tr.Translate("It's required to give a new password."))
			return c.Redirect(http.StatusFound, "../")
		}

		if data.PasswordNew1 != data.PasswordNew2 {
			templates.SetFlash(c, "ChangePassword", tr.Translate("The two given passwords doesn't match."))
			return c.Redirect(http.StatusFound, "../")
		}

		u.SetPassword(data.PasswordNew1)
		if err := us.Update(c.Request().Context(), u, njudge.Fields(njudge.UserFields.Password)); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "../")
	}
}

func PostSettingsMisc(us njudge.Users) echo.HandlerFunc {
	type request struct {
		ShowTagsForUnsolved string `form:"showTagsForUnsolved"`

		ShowTagsForUnsolvedBool bool
	}
	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}
		if data.ShowTagsForUnsolved == "true" {
			data.ShowTagsForUnsolvedBool = true
		}

		u.Settings.ShowUnsolvedTags = data.ShowTagsForUnsolvedBool
		if err := us.Update(c.Request().Context(), u, njudge.Fields(njudge.UserFields.Settings)); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "../")
	}
}
