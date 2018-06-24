package models

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func SetDatabase(db_ *sqlx.DB) {
	db = db_
}
