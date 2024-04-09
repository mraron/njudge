package api

import (
	"context"
	"database/sql"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/helpers/pagination"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type JudgeDataProvider struct {
	DB *sql.DB
}

func (JudgeDataProvider) EndpointURL() string {
	return "api/v1/judges"
}

func (JudgeDataProvider) Identifier() string {
	return "id"
}

func (dp JudgeDataProvider) List(data *pagination.Data) ([]*models.Judge, error) {
	qms := make([]QueryMod, 0)
	if data.SortField != "" {
		qms = append(qms, OrderBy(data.SortField+" "+data.SortDir))
	}

	if data.PerPage != 0 {
		qms = append(qms, Limit(data.PerPage))
		qms = append(qms, Offset(data.PerPage*(data.Page-1)))
	}

	return models.Judges(qms...).All(context.TODO(), dp.DB)
}

func (dp JudgeDataProvider) Count() (int64, error) {
	return models.Judges().Count(context.TODO(), dp.DB)
}

func (dp JudgeDataProvider) Get(id string) (*models.Judge, error) {
	return models.Judges(Where("id = ?", id)).One(context.TODO(), dp.DB)
}

func (dp JudgeDataProvider) Insert(elem *models.Judge) error {
	return elem.Insert(context.TODO(), dp.DB, boil.Infer())
}

func (dp JudgeDataProvider) Delete(id string) error {
	elem, err := models.Judges(Where("id=?", id)).One(context.TODO(), dp.DB)
	if err != nil {
		return err
	}

	_, err = elem.Delete(context.TODO(), dp.DB)
	return err
}

func (dp JudgeDataProvider) Update(id string, elem *models.Judge) error {
	_, err := elem.Update(context.TODO(), dp.DB, boil.Infer())
	return err
}
