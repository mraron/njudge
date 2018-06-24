package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/mraron/njudge/web/roles"
	"strconv"
)

type User struct {
	Id             int64
	Name           string
	HashedPassword []byte `db:"password"`
	Email          string
	ActivationKey  sql.NullString `db:"activation_key"`
	Role           roles.Role
}

func (u User) Value() (driver.Value, error) {
	return driver.Value(u.Id), nil
}

func (u *User) Scan(value interface{}) error {
	var (
		err error
		id  int
	)

	if value == nil {
		return errors.New("can't scan user from nil")
	}

	switch value.(type) {
	case int64:
		id = int(value.(int64))
	case int:
		id = value.(int)
	case []uint8:
		if id, err = strconv.Atoi(string(value.([]uint8))); err != nil {
			return err
		}
	default:
		return errors.New("can't scan user from this type")
	}

	row := db.QueryRow("SELECT * FROM users WHERE id=$1", id)
	if err = row.Scan(&u.Id, &u.Name, &u.HashedPassword, &u.Email, &u.ActivationKey, &u.Role); err != nil {
		return err
	}

	return nil
}
