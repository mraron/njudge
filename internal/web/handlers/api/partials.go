package api

import (
	"context"
	"database/sql"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PartialDataProvider struct {
	DB *sql.DB
}

func (PartialDataProvider) EndpointURL() string {
	return "api/v1/partials"
}

func (PartialDataProvider) Identifier() string {
	return "name"
}

func (dp PartialDataProvider) List(data *pagination.Data) ([]*models.Partial, error) {
	qms := make([]QueryMod, 0)
	if data.SortField != "" {
		qms = append(qms, OrderBy("name "+data.SortDir))
	}

	if data.PerPage != 0 {
		qms = append(qms, Limit(data.PerPage))
		qms = append(qms, Offset(data.PerPage*(data.Page-1)))
	}

	return models.Partials(qms...).All(context.TODO(), dp.DB)
}

func (dp PartialDataProvider) Count() (int64, error) {
	return models.Partials().Count(context.TODO(), dp.DB)
}

func (dp PartialDataProvider) Get(name string) (*models.Partial, error) {
	return models.Partials(Where("name = ?", name)).One(context.TODO(), dp.DB)
}

func (dp PartialDataProvider) Insert(elem *models.Partial) error {
	return elem.Insert(context.TODO(), dp.DB, boil.Infer())
}

func (dp PartialDataProvider) Delete(name string) error {
	elem, err := models.Partials(Where("name=?", name)).One(context.TODO(), dp.DB)
	if err != nil {
		return err
	}

	_, err = elem.Delete(context.TODO(), dp.DB)
	return err
}

func (dp PartialDataProvider) Update(name string, elem *models.Partial) error {
	elem.Name = name
	_, err := elem.Update(context.TODO(), dp.DB, boil.Infer())
	return err
}
