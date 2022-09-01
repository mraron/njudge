package api

import (
	"database/sql"
	"strconv"

	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/pagination"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserDataProvider struct {
	DB *sql.DB
}

func (UserDataProvider) EndpointURL() string {
	return "api/v1/users"
}

func (UserDataProvider) Identifier() string {
	return "id"
}

func (dp UserDataProvider) List(data *pagination.Data) ([]*models.User, error) {
	qms := make([]QueryMod, 0)
	if data.SortField != "" {
		qms = append(qms, OrderBy(data.SortField+" "+data.SortDir))
	}

	if data.PerPage != 0 {
		qms = append(qms, Limit(data.PerPage))
		qms = append(qms, Offset(data.PerPage*(data.Page-1)))
	}

	res, err := models.Users(qms...).All(dp.DB)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(res); i++ {
		helpers.CensorUserPassword(res[i])
	}

	return res, nil
}

func (dp UserDataProvider) Count() (int64, error) {
	return models.Users().Count(dp.DB)
}

func (dp UserDataProvider) Get(id string) (*models.User, error) {
	elem, err := models.Users(Where("id = ?", id)).One(dp.DB)
	if err != nil {
		return nil, err
	}

	helpers.CensorUserPassword(elem)
	return elem, nil
}

func (dp UserDataProvider) Insert(elem *models.User) error {
	return elem.Insert(dp.DB, boil.Infer())
}

func (dp UserDataProvider) Delete(id string) error {
	elem, err := models.Users(Where("id=?", id)).One(dp.DB)
	if err != nil {
		return err
	}

	_, err = elem.Delete(dp.DB)
	return err
}

func (dp UserDataProvider) Update(id string, elem *models.User) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	elem.ID = idInt
	_, err = elem.Update(dp.DB, boil.Infer())
	return err
}
