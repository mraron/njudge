package profile

import (
	"database/sql"
	"fmt"
	"github.com/mraron/njudge/internal/web/handlers/problemset"
	"github.com/mraron/njudge/internal/web/helpers/ui"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserData struct {
	Username        string                 `json:"username"`
	PictureSrc      string                 `json:"pictureSrc"`
	Rating          int                    `json:"rating"`
	Score           float64                `json:"score"`
	NumSolved       int                    `json:"numSolved"`
	LastSubmissions []problemset.StatusRow `json:"lastSubmissions"`
}

func UserDataFromUser(u *models.User) (*UserData, error) {
	return &UserData{
		Username:        u.Name,
		PictureSrc:      "/assets/profile.webp",
		LastSubmissions: []problemset.StatusRow{},
	}, nil
}

type ProfileData struct {
	UserData UserData  `json:"userData"`
	Solved   []ui.Link `json:"solved"`
	Unsolved []ui.Link `json:"unsolved"`
}

func ProfileDataFromUser(ctx context.Context, DB *sql.DB, u *models.User) (*ProfileData, error) {
	var (
		solved, attempted models.SubmissionSlice
		err               error
	)

	solved, err = models.Submissions(Select("max(submissions.id) as id, problemset, problem"),
		Where("user_id = ?", u.ID), Where("verdict = ?", 0),
		GroupBy("submissions.problemset, submissions.problem"),
	).All(ctx, DB)
	if err != nil {
		return nil, err
	}

	attempted, err = models.Submissions(Select("max(submissions.id) as id, problemset, problem"),
		Where("user_id = ?", u.ID), Where("verdict <> ?", 0),
		Where("not exists(select id from submissions as other where other.user_id = ? and "+
			"verdict = 0 and other.problem = submissions.problem and other.problemset=submissions.problemset)", u.ID),
		GroupBy("submissions.problemset, submissions.problem"),
	).All(ctx, DB)
	if err != nil {
		return nil, err
	}

	toLinks := func(s models.SubmissionSlice) []ui.Link {
		res := make([]ui.Link, len(s))
		for i := range res {
			res[i] = ui.Link{
				Text: s[i].Problem,
				Href: fmt.Sprintf("/submission/%d/", s[i].ID),
			}
		}

		return res
	}

	ud, err := UserDataFromUser(u)
	if err != nil {
		return nil, err
	}

	return &ProfileData{
		UserData: *ud,
		Solved:   toLinks(solved),
		Unsolved: toLinks(attempted),
	}, nil
}

func GetProfile(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("profile").(*models.User)
		pd, err := ProfileDataFromUser(c.Request().Context(), DB.DB, u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, pd)
	}
}

func GetSubmissions(DB *sqlx.DB, problemStore problems.Store, statusPageService services.StatusPageService) echo.HandlerFunc {
	type request struct {
		Page int `query:"page"`
	}
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		u := c.Get("profile").(*models.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if data.Page <= 0 {
			data.Page = 1
		}

		statusReq := services.StatusPageRequest{
			Pagination: pagination.Data{
				Page:      data.Page,
				PerPage:   20,
				SortDir:   "DESC",
				SortField: "id",
			},
			UserID: u.ID,
		}

		statusPage, err := statusPageService.GetStatusPage(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		statusRows := make([]*problemset.StatusRow, len(statusPage.Submissions))
		for i := range statusPage.Submissions {
			sub := &statusPage.Submissions[i]
			if statusRows[i], err = problemset.StatusRowFromSubmission(c.Request().Context(), DB.DB,
				problemStore, tr, sub); err != nil {
				return err
			}
		}

		return c.JSON(http.StatusOK, struct {
			PaginationData pagination.Data         `json:"paginationData"`
			Submissions    []*problemset.StatusRow `json:"submissions"`
		}{statusPage.PaginationData, statusRows})
	}
}

func GetSettings(DB *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		return c.JSON(http.StatusOK, struct {
			ShowUnsolved bool `json:"showUnsolved"`
		}{u.ShowUnsolvedTags})
	}
}

func PostSettingsChangePassword(DB *sqlx.DB) echo.HandlerFunc {
	type request struct {
		PasswordOld  string `json:"oldPw"`
		PasswordNew1 string `json:"newPw"`
		PasswordNew2 string `json:"newPwConfirm"`
	}
	type response struct {
		Message string `json:"message"`
	}
	return func(c echo.Context) error {
		tr := c.Get("translator").(i18n.Translator)
		u := c.Get("user").(*models.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.PasswordOld)); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, response{
				Message: tr.Translate("Wrong old password."),
			})
		}

		if len(data.PasswordNew1) == 0 {
			return c.JSON(http.StatusUnprocessableEntity, response{
				Message: tr.Translate("It's required to give a new password."),
			})
		}

		if data.PasswordNew1 != data.PasswordNew2 {
			return c.JSON(http.StatusUnprocessableEntity, response{
				Message: tr.Translate("The two given passwords doesn't match."),
			})
		}

		res, err := bcrypt.GenerateFromPassword([]byte(data.PasswordNew1), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		u.Password = string(res)
		if _, err = u.Update(c.Request().Context(), DB, boil.Whitelist(models.UserColumns.Password)); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "")
	}
}

func PostSettingsMisc(DB *sqlx.DB) echo.HandlerFunc {
	type request struct {
		ShowTagsForUnsolved bool `json:"showUnsolved"`
	}
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		u.ShowUnsolvedTags = data.ShowTagsForUnsolved
		if _, err := u.Update(c.Request().Context(), DB, boil.Whitelist(models.UserColumns.ShowUnsolvedTags)); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "../")
	}
}
