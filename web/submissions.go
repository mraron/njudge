package web

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (s *Server) Submit(uid int, problemset, problem, language string, source []byte) (int, error) {
	var (
		tx  *sql.Tx
		id  int
		err error
	)
	mustPanic := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	ok := true

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

		id = 0
		res, err := tx.Query("INSERT INTO submissions (status,\"user_id\",verdict,ontest,submitted,judged,problem,language,private,problemset,source,started) VALUES ($1,$2,$3,NULL,$4,NULL,$5,$6,false,$7, $8,false) RETURNING id", problems.Status{}, uid, VERDICT_UP, time.Now(), s.GetProblem(problem).Name(), language, problemset, source)

		mustPanic(err)

		res.Next()

		err = res.Scan(&id)
		mustPanic(err)

		err = res.Close()
		mustPanic(err)

		fs, err := os.Create(filepath.Join(s.SubmissionsDir, strconv.Itoa(int(id))))
		mustPanic(err)

		_, err = fs.Write([]byte(source))
		mustPanic(err)

		err = tx.Commit()
		mustPanic(err)
	}

	if transaction(); err != nil {
		return -1, err
	}

	return id, nil
}

func (s *Server) postProblemsetSubmit(c echo.Context) error {
	var (
		u   *models.User
		err error
		id  int
		p   problems.Problem
	)

	if u = c.Get("user").(*models.User); u == nil {
		return c.Render(http.StatusForbidden, "error.gohtml", "Előbb lépj be.")
	}

	problemName := c.FormValue("problem")
	if has, _ := s.ProblemStore.Has(problemName); !has {
		return c.Render(http.StatusOK, "error.gohtml", "Hibás feladatazonosító.")
	}else {
		p, _ = s.ProblemStore.Get(problemName)
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
		return c.Render(http.StatusOK, "error.gohtml", "Hibás nyelvazonosító.")
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

	if id, err = s.Submit(u.ID, c.Get("problemset").(string), s.GetProblem(c.FormValue("problem")).Name(), languageName, contents); err != nil {
		return s.internalError(c, err, "Belső hiba #4")
	}

	return c.Redirect(http.StatusFound, "/problemset/status#submission"+strconv.Itoa(id))
}

func (s *Server) getSubmission(c echo.Context) error {
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.internalError(c, err, "ajaj")
	}

	sub, err := models.Submissions(Where("id=?", val)).One(s.db)
	//sub, err := models.SubmissionFromId(s.db, int64(val))
	if err != nil {
		return s.internalError(c, err, "ajaj")
	}

	return c.Render(http.StatusOK, "submission.gohtml", sub)
}

func (s *Server) getAPISubmissions(c echo.Context) error {
	u := c.Get("user").(*models.User)

	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/submissions") {
		return s.unauthorizedError(c)
	}

	data, err := parsePaginationData(c)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	lst, err := models.Submissions(OrderBy(data._sortField+" "+data._sortDir), Limit(data._perPage), Offset(data._perPage*(data._page-1))).All(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}
	//models.Submissions().Count(s.db)

	//source code is quiet big to serve for lists
	for i := 0; i < len(lst); i++ {
		lst[i].Source = []byte("-")
	}

	return c.JSON(http.StatusOK, lst)
}

func (s *Server) postAPISubmission(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "api/v1/submissions") {
		return s.unauthorizedError(c)
	}

	pr := new(models.Submission)
	if err := c.Bind(pr); err != nil {
		return s.internalError(c, err, "error")
	}

	return pr.Insert(s.db, boil.Infer())
}

func (s *Server) getAPISubmission(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "api/v1/submissions") {
		return s.unauthorizedError(c)
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr, err := models.Submissions(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, pr)
}

func (s *Server) deleteAPISubmission(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionDelete, "api/v1/submissions") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr, err := models.Submissions(Where("id=?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	_, err = pr.Delete(s.db)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.String(http.StatusOK, "ok")
}

func (s *Server) putAPISubmission(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionEdit, "api/v1/submissions") {
		return s.unauthorizedError(c)
	}

	id_ := c.Param("id")

	id, err := strconv.Atoi(id_)
	if err != nil {
		return s.internalError(c, err, "error")
	}

	pr := new(models.Submission)
	if err = c.Bind(pr); err != nil {
		return s.internalError(c, err, "error")
	}

	pr.ID = id
	_, err = pr.Update(s.db, boil.Infer())

	if err != nil {
		return s.internalError(c, err, "error")
	}

	return c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{"updated"})
}

func (s *Server) getSubmissionRejudge(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.internalError(c, err, "can't parse id")
	}

	sub, err := models.Submissions(Where("id = ?", id)).One(s.db)
	if err != nil {
		return s.internalError(c, err, "can't find submission")
	}

	sub.ID = id
	sub.Judged = null.Time{Valid: false}
	sub.Started = false
	sub.Update(s.db, boil.Infer())

	return c.Redirect(http.StatusFound, "/submission/"+strconv.Itoa(id))
}