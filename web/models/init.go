package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDatabase(db_ *gorm.DB) {
	db = db_
}
