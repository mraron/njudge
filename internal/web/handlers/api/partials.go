package api

import (
	"database/sql"
	"fmt"
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

	return models.Partials(qms...).All(dp.DB)
}

func (dp PartialDataProvider) Count() (int64, error) {
	return models.Partials().Count(dp.DB)
}

func (dp PartialDataProvider) Get(name string) (*models.Partial, error) {
	return models.Partials(Where("name = ?", name)).One(dp.DB)
}

func (dp PartialDataProvider) Insert(elem *models.Partial) error {
	fmt.Println(elem)
	return elem.Insert(dp.DB, boil.Infer())
}

func (dp PartialDataProvider) Delete(name string) error {
	elem, err := models.Partials(Where("name=?", name)).One(dp.DB)
	if err != nil {
		return err
	}

	_, err = elem.Delete(dp.DB)
	return err
}

func (dp PartialDataProvider) Update(name string, elem *models.Partial) error {
	elem.Name = name
	_, err := elem.Update(dp.DB, boil.Infer())
	return err
}
