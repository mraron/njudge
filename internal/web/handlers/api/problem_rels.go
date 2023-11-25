package api

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/helpers/pagination"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ProblemRelDataProvider struct {
	DB *sql.DB
}

func (ProblemRelDataProvider) EndpointURL() string {
	return "api/v1/problem_rels"
}

func (ProblemRelDataProvider) Identifier() string {
	return "id"
}

func (dp ProblemRelDataProvider) List(data *pagination.Data) ([]*models.ProblemRel, error) {
	qms := make([]QueryMod, 0)
	if data.SortField != "" {
		qms = append(qms, OrderBy(data.SortField+" "+data.SortDir))
	}

	if data.PerPage != 0 {
		qms = append(qms, Limit(data.PerPage))
		qms = append(qms, Offset(data.PerPage*(data.Page-1)))
	}

	return models.ProblemRels(qms...).All(context.TODO(), dp.DB)
}

func (dp ProblemRelDataProvider) Count() (int64, error) {
	return models.ProblemRels().Count(context.TODO(), dp.DB)
}

func (dp ProblemRelDataProvider) Get(id string) (*models.ProblemRel, error) {
	return models.ProblemRels(Where("id = ?", id)).One(context.TODO(), dp.DB)
}

func (dp ProblemRelDataProvider) Insert(elem *models.ProblemRel) error {
	return elem.Insert(context.TODO(), dp.DB, boil.Infer())
}

func (dp ProblemRelDataProvider) Delete(id string) error {
	elem, err := models.ProblemRels(Where("id=?", id)).One(context.TODO(), dp.DB)
	if err != nil {
		return err
	}

	_, err = elem.Delete(context.TODO(), dp.DB)
	return err
}

func (dp ProblemRelDataProvider) Update(id string, elem *models.ProblemRel) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	elem.ID = idInt
	_, err = elem.Update(context.TODO(), dp.DB, boil.Infer())
	return err
}
