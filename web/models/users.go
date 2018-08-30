package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/mraron/njudge/web/roles"
	"strconv"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"column:name"`
	HashedPassword []byte `gorm:"column:password"`
	Email          string `gorm:"column:email"`
	ActivationKey  sql.NullString `gorm:"column:activation_key"`
	Role           roles.Role `gorm:"column:role"`
}

func (u User) Value() (driver.Value, error) {
	return driver.Value(u.ID), nil
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

	if err = db.Model(&User{}).Where("id = ?", id).First(u).Error; err != nil {
		return err
	}

	return nil
}

func UserFromId(db *gorm.DB, id int) (ans User, err error) {
	err = db.Model(&User{}).Where("id = ?", id).First(&ans).Error
	return
}

/*
func (u User) Delete(db *sqlx.DB) (err error) {
	_, err = db.Exec("DELETE FROM users WHERE id=$1", u.ID)
	return
}

func (u User) Update(db *sqlx.DB) (err error) {
	_, err = db.Exec("UPDATE users SET name=$1, password=$2, email=$3, activation_key=$4, role=$5 WHERE id=$6", u.Name, u.HashedPassword, u.Email, u.ActivationKey, u.Role, u.ID)
	return
}

func (u *User) Insert(db *sqlx.DB) error {
	res := db.QueryRow("INSERT INTO users (name, password, email, activation_key, role) VALUES ($1,$2,$3,$4,$5) RETURNING id", u.Name, u.HashedPassword, u.Email, u.ActivationKey, u.Role)

	var id int64
	err := res.Scan(&id)

	if err != nil {
		return err
	}

	u.ID = id

	return nil
}*/
