package web

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/markbates/goth/gothic"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/mail"
	"github.com/mraron/njudge/web/helpers/pagination"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/url"
	"strconv"
)

func (s *Server) getUserLogin(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	return c.Render(http.StatusOK, "login.gohtml", nil)
}

func (s *Server) postUserLogin(c echo.Context) error {
	var (
		u   *models.User = &models.User{}
		err error
	)

	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	u, err = models.Users(Where("name=?", c.FormValue("name"))).One(s.DB)
	if err != nil {
		log.Error("Possible just wrong credentials, but", err)
		return c.Render(http.StatusOK, "login.gohtml", []string{"Hibás felhasználónév és jelszó páros."})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.FormValue("password"))); err != nil {
		return c.Render(http.StatusOK, "login.gohtml", []string{"Hibás felhasználónév és jelszó páros."})
	}

	if u.ActivationKey.Valid {
		return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: az account nincs aktiválva."})
	}

	storage, _ := session.Get("user", c)
	storage.Values["id"] = u.ID

	if err = storage.Save(c.Request(), c.Response()); err != nil {
		return helpers.InternalError(c, err, "Belső hiba #2")
	}

	c.Set("user", u)

	return c.Render(http.StatusOK, "message.gohtml", "Sikeres belépés.")
}

func (s *Server) getUserRegister(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	return c.Render(http.StatusOK, "register.gohtml", nil)
}

func (s *Server) postUserRegister(c echo.Context) error {
	var (
		errStrings = make([]string, 0)
		key             = helpers.GenerateActivationKey(255)
		err    error
		tx     *sql.Tx
	)

	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	used := func(col, value, msg string) {
		u := ""
		if s.DB.Get(&u, "SELECT name FROM users WHERE "+col+"=$1", value); u != "" {
			errStrings = append(errStrings, msg)
		}
	}

	required := func(value, msg string) {
		if c.FormValue(value) == "" {
			errStrings = append(errStrings, msg)
		}
	}

	used("name", c.FormValue("name"), "A név foglalt")
	used("email", c.FormValue("email"), "Az email cím foglalt")

	required("name", "A név mező szükséges")
	required("password", "A jelszó mező szükséges")
	required("password2", "A jelszó ellenörző mező szükséges")
	required("email", "Az email mező szükséges")

	if c.FormValue("password") != c.FormValue("password2") {
		errStrings = append(errStrings, "A két jelszó nem egyezik meg")
	}

	if len(errStrings) > 0 {
		return c.Render(http.StatusOK, "register.gohtml", errStrings)
	}

	mustPanic := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	transaction := func() {
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()

				var ok bool
				if err, ok = p.(error); !ok {
					err = errors.New("hiba")
				}
			}
		}()

		tx, err := s.DB.Begin()
		mustPanic(err)

		hashed, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
		mustPanic(err)

		_, err = tx.Exec("INSERT INTO users (name,password,email,activation_key,role) VALUES ($1,$2,$3,$4,$5)", c.FormValue("name"), hashed, c.FormValue("email"), key, "user")
		mustPanic(err)

		m := mail.Mail{}
		m.Recipients = []string{c.FormValue("email")}
		m.Message = fmt.Sprintf(`Kedves %s!<br> Köszönjük regisztrációd. Aktiváló link: <a href="http://`+s.Hostname+`/user/activate/%s/%s">http://`+s.Hostname+`/user/activate/%s/%s</a>`, c.FormValue("name"), c.FormValue("name"), key, c.FormValue("name"), key)
		m.Subject = "Regisztráció aktiválása"
		mustPanic(m.Send(s.Server))

		mustPanic(tx.Commit())
	}

	if transaction(); err != nil {
		return helpers.InternalError(c, err, "Belső hiba.")
	}

	return c.Redirect(http.StatusFound, "/user/activate")
}

func (s *Server) getUserLogout(c echo.Context) error {
	if u := c.Get("user").(*models.User); u == nil {
		return c.Render(http.StatusOK, "error.gohtml", "Ahhoz hogy kijelentkezz előbb be kell hogy jelentkezz...")
	}

	storage, _ := session.Get("user", c)
	storage.Options.MaxAge = -1
	storage.Values["id"] = -1

	if err := storage.Save(c.Request(), c.Response()); err != nil {
		return helpers.InternalError(c, err, "Belső hiba")
	}

	return c.Redirect(http.StatusFound, "/")
}

func (s *Server) getUserActivate(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	return c.Render(http.StatusOK, "activate.gohtml", nil)
}

func (s *Server) getActivateUser(c echo.Context) error {
	var (
		user *models.User
		err  error
		tx   *sql.Tx
	)

	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	if user, err = models.Users(Where("name=?", c.Param("name"))).One(s.DB); err != nil {
		return helpers.InternalError(c, err, "Belső hiba #1")
	}

	if !user.ActivationKey.Valid {
		return c.Render(http.StatusOK, "error.gohtml", "Ez a regisztráció már aktív!")
	}

	if user.ActivationKey.String != c.Param("key") {
		return c.Render(http.StatusOK, "error.gohtml", "Hibás aktiválási kulcs. Biztos jó linkre kattintottál?")
	}

	if tx, err = s.DB.Begin(); err != nil {
		return helpers.InternalError(c, err, "Belső hiba #2")
	}

	if _, err = tx.Exec("UPDATE users SET activation_key=NULL WHERE name=$1", c.Param("name")); err != nil {
		return helpers.InternalError(c, err, "Belső hiba #3")
	}

	if err = tx.Commit(); err != nil {
		return helpers.InternalError(c, err, "Belső hiba #4")
	}

	return c.Render(http.StatusOK, "message.gohtml", "Sikeres aktiválás, mostmár beléphetsz.")
}

func (s *Server) getUserAuthCallback(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: érvénytelen token."})
	}

	lst, err := models.Users(Where("email = ?", user.Email)).All(s.DB)
	if len(lst) == 0 {
		return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: a felhasználó nincs regisztrálva."})
	}

	if lst[0].ActivationKey.Valid {
		return c.Render(http.StatusOK, "login.gohtml", []string{"Hiba: az account nincs aktiválva."})
	}

	storage, _ := session.Get("user", c)
	storage.Values["id"] = lst[0].ID

	if err = storage.Save(c.Request(), c.Response()); err != nil {
		return helpers.InternalError(c, err, "Belső hiba #2")
	}

	c.Set("user", lst[0])

	return c.Render(http.StatusOK, "message.gohtml", "Sikeres belépés.")
}

func (s *Server) getUserAuth(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.gohtml", "Már be vagy lépve...")
	}

	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

func (s *Server) getAPIUsers(c echo.Context) error {
	u := c.Get("user").(*models.User)

	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	data, err := pagination.Parse(c)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	lst, err := models.Users(OrderBy(data.SortField+" "+data.SortDir), Limit(data.PerPage), Offset(data.PerPage*(data.Page-1))).All(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	for i := 0; i < len(lst); i++ {
		helpers.CensorUserPassword(lst[i])
	}

	return c.JSON(http.StatusOK, lst)
}

func (s *Server) postAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	pr := new(models.User)
	if err := c.Bind(pr); err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return pr.Insert(s.DB, boil.Infer())
}

func (s *Server) getAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr, err := models.Users(Where("id=?", id)).One(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	helpers.CensorUserPassword(pr)

	return c.JSON(http.StatusOK, pr)
}

func (s *Server) deleteAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionDelete, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr, err := models.Users(Where("id=?", id)).One(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	_, err = pr.Delete(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return c.String(http.StatusOK, "ok")
}

func (s *Server) putAPIUser(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionEdit, "api/v1/users") {
		return helpers.UnauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr := new(models.User)
	if err = c.Bind(pr); err != nil {
		return helpers.InternalError(c, err, "error")
	}

	pr.ID = id
	_, err = pr.Update(s.DB, boil.Infer())

	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"updated"})
}

func (s *Server) getUserProfile(c echo.Context) error {
	name, err := url.QueryUnescape(c.Param("name"))
	if err != nil {
		return helpers.InternalError(c, err, "hiba")
	}

	user, err := models.Users(Where("name = ?", name)).One(s.DB)
	if err != nil {
		return helpers.InternalError(c, err, "error")
	}

	return c.Render(http.StatusOK, "profile.gohtml", struct {
		User *models.User
	}{user})
}

func (s *Server) UserSolvedStatus(problemSet, problem string, u *models.User) (int, error) {
	solvedStatus := -1
	if u != nil {
		cnt, err := models.Submissions(Where("problemset = ?", problemSet), Where("problem = ?", problem), Where("verdict = 0"), Where("user_id = ?", u.ID)).Count(s.DB)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return -1, fmt.Errorf("can't get solvedstatus for %s %s%s: %w", u.Name, problemSet, problem, err)
		}else {
			if cnt > 0 {
				solvedStatus = 0
			} else {
				cnt, err := models.Submissions(Where("problemset = ?", problemSet), Where("problem = ?", problem), Where("user_id = ?", u.ID)).Count(s.DB)
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					return -1, fmt.Errorf("can't get solvedstatus for %s %s%s: %w", u.Name, problemSet, problem, err)
				} else {
					if cnt>0 {
						solvedStatus = 1
					}
				}
			}
		}
	}

	return solvedStatus, nil
}