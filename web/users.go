package web

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/mraron/njudge/web/models"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
	"github.com/jinzhu/gorm"
)

func GenerateActivationKey(length int) string {
	var (
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678901234567890123456789"
		ans      = make([]byte, length)
	)

	src := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		ans[i] = alphabet[(int(src.Int63()))%len(alphabet)]
	}

	return string(ans)
}

func (s *Server) currentUser(c echo.Context) (*models.User, error) {
	var (
		u   *models.User = &models.User{}
		err error
	)

	storage, _ := session.Get("user", c)

	if _, ok := storage.Values["id"]; !ok {
		return nil, nil
	}

	err = s.db.Where("id = ?", storage.Values["id"]).First(u).Error
	return u, err
}

func (s *Server) getUserLogin(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.html", "Már be vagy lépve...")
	}

	return c.Render(http.StatusOK, "login.html", nil)
}

func (s *Server) postUserLogin(c echo.Context) error {
	var (
		u   *models.User = &models.User{}
		err error
	)

	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.html", "Már be vagy lépve...")
	}

	if err = s.db.Where("name = ?", c.FormValue("name")).First(&u).Error; err != nil {
		return s.internalError(c, err, "Belső hiba #1")
	}

	if err = bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(c.FormValue("password"))); err != nil {
		return c.Render(http.StatusOK, "login.html", []string{"Hibás felhasználónév és jelszó páros."})
	}

	if u.ActivationKey.Valid {
		return c.Render(http.StatusOK, "login.html", []string{"Hiba: az account nincs aktiválva."})
	}

	storage, _ := session.Get("user", c)
	storage.Values["id"] = u.ID

	if err = storage.Save(c.Request(), c.Response()); err != nil {
		return s.internalError(c, err, "Belső hiba #2")
	}

	c.Set("user", u)

	return c.Render(http.StatusOK, "message.html", "Sikeres belépés.")
}

func (s *Server) getUserRegister(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.html", "Már be vagy lépve...")
	}

	return c.Render(http.StatusOK, "register.html", nil)
}

func (s *Server) postUserRegister(c echo.Context) error {
	var (
		errors []string = make([]string, 0)
		key             = GenerateActivationKey(255)
		err    error
		tx     *sql.Tx
	)

	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.html", "Már be vagy lépve...")
	}

	used := func(col, value, msg string) {
		u := &models.User{}
		if err := s.db.Where("col = ?", value).First(u); err != nil && u.Name!="" {
			errors = append(errors, msg)
		}
	}

	required := func(value, msg string) {
		if c.FormValue(value) == "" {
			errors = append(errors, msg)
		}
	}

	used("name", c.FormValue("name"), "A név foglalt")
	used("email", c.FormValue("email"), "Az email cím foglalt")

	required("name", "A név mező szükséges")
	required("password", "A jelszó mező szükséges")
	required("password2", "A jelszó ellenörző mező szükséges")
	required("email", "Az email mező szükséges")

	if c.FormValue("password") != c.FormValue("password2") {
		errors = append(errors, "A két jelszó nem egyezik meg")
	}

	if len(errors) > 0 {
		return c.Render(http.StatusOK, "register.html", errors)
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
				err = p.(error)
			}
		}()

		tx := s.db.Begin()

		hashed, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
		mustPanic(err)

		user := &models.User{}
		user.Name = c.FormValue("name")
		user.HashedPassword = hashed
		user.Email = c.FormValue("email")
		user.ActivationKey.String = key
		user.ActivationKey.Valid = true
		user.Role = "user"

		err = s.db.Save(user).Error
		mustPanic(err)

		//_, err = tx.Exec("INSERT INTO users (name,password,email,activation_key,role) VALUES ($1,$2,$3,$4,$5)", c.FormValue("name"), hashed, c.FormValue("email"), key, "user")


		m := Mail{}
		m.Recipients = []string{c.FormValue("email")}
		m.Message = fmt.Sprintf(`Kedves %s!<br> Köszönjük a registrációt. Aktiváló link: <a href="http://localhost:8080/user/activate/%s/%s">http://localhost:8080/user/activate/%s/%s</a>`, c.FormValue("name"), c.FormValue("name"), key, c.FormValue("name"), key)
		m.Subject = "Regisztráció aktiválása"
		mustPanic(s.SendMail(m))

		mustPanic(tx.Commit().Error)
	}

	if transaction(); err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	return c.Redirect(http.StatusFound, "/user/activate")
}

func (s *Server) getUserLogout(c echo.Context) error {
	if u := c.Get("user").(*models.User); u == nil {
		return c.Render(http.StatusOK, "error.html", "Ahhoz hogy kijelentkezz előbb be kell hogy jelentkezz...")
	}

	storage, _ := session.Get("user", c)
	storage.Options.MaxAge = -1
	storage.Values["id"] = -1

	if err := storage.Save(c.Request(), c.Response()); err != nil {
		return s.internalError(c, err, "Belső hiba")
	}

	return c.Redirect(http.StatusFound, "/")
}

func (s *Server) getUserActivate(c echo.Context) error {
	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.html", "Már be vagy lépve...")
	}

	return c.Render(http.StatusOK, "activate.html", nil)
}

func (s *Server) getActivateUser(c echo.Context) error {
	var (
		user *models.User
		err  error
		tx   *gorm.DB
	)

	if u := c.Get("user").(*models.User); u != nil {
		return c.Render(http.StatusOK, "error.html", "Már be vagy lépve...")
	}

	if err = s.db.Where("name = ?", c.Param("name")).First(&user).Error; err != nil {
		return s.internalError(c, err, "Belső hiba #1")
	}

	if !user.ActivationKey.Valid {
		return c.Render(http.StatusOK, "error.html", "Ez a regisztráció már aktív!")
	}

	if user.ActivationKey.String != c.Param("key") {
		return c.Render(http.StatusOK, "error.html", "Hibás aktiválási kulcs. Biztos jó linkre kattintottál?")
	}

	if tx = s.db.Begin(); err != nil {
		return s.internalError(c, err, "Belső hiba #2")
	}

	user.ActivationKey.String = ""
	user.ActivationKey.Valid = false
	if err = tx.Save(user).Error; err != nil {
		return s.internalError(c, err, "Belső hiba #3")
	}

	if err = tx.Commit().Error; err != nil {
		return s.internalError(c, err, "Belső hiba #4")
	}

	return c.Render(http.StatusOK, "message.html", "Sikeres aktiválás, mostmár beléphetsz.")
}
