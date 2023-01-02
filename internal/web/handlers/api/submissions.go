package api

import (
	"database/sql"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SubmissionDataProvider struct {
	DB *sql.DB
}

func (SubmissionDataProvider) EndpointURL() string {
	return "api/v1/submissions"
}

func (SubmissionDataProvider) Identifier() string {
	return "id"
}

func (dp SubmissionDataProvider) List(data *pagination.Data) ([]*models.Submission, error) {
	qms := make([]QueryMod, 0)
	if data.SortField != "" {
		qms = append(qms, OrderBy(data.SortField+" "+data.SortDir))
	}

	if data.PerPage != 0 {
		qms = append(qms, Limit(data.PerPage))
		qms = append(qms, Offset(data.PerPage*(data.Page-1)))
	}

	return models.Submissions(qms...).All(dp.DB)
}

func (dp SubmissionDataProvider) Count() (int64, error) {
	return models.Submissions().Count(dp.DB)
}

func (dp SubmissionDataProvider) Get(id string) (*models.Submission, error) {
	return models.Submissions(Where("id = ?", id)).One(dp.DB)
}

func (dp SubmissionDataProvider) Insert(elem *models.Submission) error {
	return elem.Insert(dp.DB, boil.Infer())
}

func (dp SubmissionDataProvider) Delete(id string) error {
	elem, err := models.Submissions(Where("id=?", id)).One(dp.DB)
	if err != nil {
		return err
	}

	_, err = elem.Delete(dp.DB)
	return err
}

func (dp SubmissionDataProvider) Update(id string, elem *models.Submission) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	elem.ID = idInt
	_, err = elem.Update(dp.DB, boil.Infer())
	return err
}
