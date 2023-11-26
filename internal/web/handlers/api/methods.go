package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/helpers/roles"

	"github.com/labstack/echo/v4"
)

type Provider[T any] interface {
	EndpointURL() string

	Identifier() string

	List(*pagination.Data) ([]*T, error)
	Count() (int64, error)
	Get(string) (*T, error)
}

type WritableProvider[T any] interface {
	Provider[T]

	Insert(*T) error
	Delete(string) error
	Update(string, *T) error
}

func GetList[T any](dp Provider[T]) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)

		if !roles.Can(roles.Role(u.Role), roles.ActionView, roles.Entity(dp.EndpointURL())) {
			return helpers.UnauthorizedError(c)
		}

		data, err := pagination.Parse(c)
		if err != nil {
			return err
		}

		lst, err := dp.List(data)
		if err != nil {
			return err
		}

		cnt, err := dp.Count()
		if err != nil {
			return err
		}

		c.Response().Header().Set("X-Total-Count", strconv.Itoa(int(cnt)))

		if lst == nil {
			lst = make([]*T, 0)
		}
		return c.JSON(http.StatusOK, lst)
	}
}

func Get[T any](dp Provider[T]) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionView, roles.Entity(dp.EndpointURL())) {
			return helpers.UnauthorizedError(c)
		}

		id, err := url.QueryUnescape(c.Param(dp.Identifier()))
		if err != nil {
			return err
		}

		elem, err := dp.Get(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, elem)
	}
}

func Post[T any](dp WritableProvider[T]) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionCreate, roles.Entity(dp.EndpointURL())) {
			return helpers.UnauthorizedError(c)
		}

		elem := new(T)
		if err := c.Bind(elem); err != nil {
			return err
		}

		if err := dp.Insert(elem); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, elem)
	}
}

func Put[T any](dp WritableProvider[T]) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionEdit, roles.Entity(dp.EndpointURL())) {
			return helpers.UnauthorizedError(c)
		}

		id, err := url.QueryUnescape(c.Param(dp.Identifier()))
		if err != nil {
			return err
		}

		elem := new(T)
		if err := c.Bind(elem); err != nil {
			return err
		}

		err = dp.Update(id, elem)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{"updated"})
	}
}

func Delete[T any](dp WritableProvider[T]) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionDelete, roles.Entity(dp.EndpointURL())) {
			return helpers.UnauthorizedError(c)
		}

		id, err := url.QueryUnescape(c.Param(dp.Identifier()))
		if err != nil {
			return err
		}

		if err = dp.Delete(id); err != nil {
			return err
		}

		return c.String(http.StatusOK, "ok")
	}
}
