package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/jmoiron/sqlx"
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

func UserFromId(db *sqlx.DB, id int) (ans User, err error) {
	err = db.Get(&ans, "SELECT * FROM users WHERE id=$1", id)
	return
}

func (u User) Delete(db *sqlx.DB) (err error) {
	_, err = db.Exec("DELETE FROM users WHERE id=$1", u.Id)
	return
}

func (u User) Update(db *sqlx.DB) (err error) {
	_, err = db.Exec("UPDATE users SET name=$1, password=$2, email=$3, activation_key=$4, role=$5 WHERE id=$6", u.Name, u.HashedPassword, u.Email, u.ActivationKey, u.Role, u.Id)
	return
}

func (u *User) Insert(db *sqlx.DB) error {
	res := db.QueryRow("INSERT INTO users (name, password, email, activation_key, role) VALUES ($1,$2,$3,$4,$5) RETURNING id", u.Name, u.HashedPassword, u.Email, u.ActivationKey, u.Role)

	var id int64
	err := res.Scan(&id)

	if err != nil {
		return err
	}

	u.Id = id

	return nil
}
