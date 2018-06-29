package web

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/models"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (s *Server) postProblemsetSubmit(c echo.Context) error {
	var (
		u   *models.User
		err error
		tx  *sql.Tx
		p   problems.Problem
		ok  bool
	)

	if u = c.Get("user").(*models.User); u == nil {
		return c.Render(http.StatusForbidden, "error.html", "Előbb lépj be.")
	}

	problemName := c.FormValue("problem")
	if p, ok = s.problems[problemName]; !ok {
		return c.Render(http.StatusOK, "error.html", "Hibás feladatazonosító.")
	}

	languageName := c.FormValue("language")

	found := false
	for _, lang := range p.Languages() {
		if lang.Id() == languageName {
			found = true
			break
		}
	}

	if !found {
		return c.Render(http.StatusOK, "error.html", "Hibás nyelvazonosító.")
	}

	fileHeader, err := c.FormFile("source")
	if err != nil {
		return s.internalError(c, err, "Belső hiba #0")
	}

	f, err := fileHeader.Open()
	if err != nil {
		return s.internalError(c, err, "Belső hiba #1")
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return s.internalError(c, err, "Belső hiba #2")
	}

	mustPanic := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	ok = true

	transaction := func() {
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				ok = false
				err = p.(error)
			}
		}()

		tx, err = s.db.Begin()
		mustPanic(err)

		last := 0
		res, err := tx.Query("INSERT INTO submissions (status,\"user\",verdict,ontest,submitted,judged,problem,language,private,problemset,source,started) VALUES ($1,$2,$3,NULL,$4,NULL,$5,$6,false,$7, $8,false) RETURNING id", problems.Status{}, u, models.VERDICT_UP, time.Now(), s.problems[c.FormValue("problem")].Name(), c.FormValue("language"), c.Get("problemset"), contents)
		mustPanic(err)

		res.Next()

		err = res.Scan(&last)
		mustPanic(err)

		err = res.Close()
		mustPanic(err)

		fs, err := os.Create("submissions/" + strconv.Itoa(int(last)))
		mustPanic(err)

		_, err = fs.Write(contents)
		mustPanic(err)

		err = tx.Commit()
		mustPanic(err)
	}

	if transaction(); !ok {
		return s.internalError(c, err, "Belső hiba #4")
	}

	return c.Redirect(http.StatusFound, "/problemset/status")
}

func (s *Server) getSubmission(c echo.Context) error {
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.internalError(c, err, "ajaj")
	}

	sub, err := models.SubmissionFromId(s.db, int64(val))
	if err != nil {
		return s.internalError(c, err, "ajaj")
	}

	return c.Render(http.StatusOK, "submission.html", sub)
}
