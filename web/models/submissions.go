package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/lib/pq"
	"strconv"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/mraron/njudge/utils/problems"
)

type Verdict int

const (
	VERDICT_AC Verdict = iota
	VERDICT_WA
	VERDICT_RE
	VERDICT_TL
	VERDICT_ML
	VERDICT_XX
	VERDICT_CE
	VERDICT_RU
	VERDICT_UP
)

func (v *Verdict) Scan(value interface{}) error {
	var (
		val int
		err error
	)

	if value == nil {
		return errors.New("can't scan verdicct from nil")
	}

	switch value.(type) {
	case int:
		val = value.(int)
	case int64:
		val = int(value.(int64))
	case []uint8:
		val, err = strconv.Atoi(string(value.([]uint8)))
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported type")
	}

	*v = Verdict(val)

	return nil
}

func (v Verdict) Value() (driver.Value, error) {
	return driver.Value(int64(v)), nil
}

type Submission struct {
	gorm.Model
	Status     problems.Status `gorm:"column:status"`
	Verdict    Verdict `gorm:"column:verdict"`
	OnTest     sql.NullString `gorm:"column:ontest"`
	User       *User `sql:"type:integer" gorm:"column:userid"`
	Submitted  time.Time `gorm:"column:submitted"`
	Judged     pq.NullTime `gorm:"column:judged"`
	Problemset string `gorm:"column:problemset"`
	Problem    string `gorm:"column:problem"`
	Language   string `gorm:"column:language"`
	Private    bool `gorm:"column:private"`
	Source     string `gorm:"column:source"`
	Started    bool `gorm:"column:started"`
}


func SubmissionFromId(db *gorm.DB, id int64) (*Submission, error) {
	s := new(Submission)
	err := db.Model(&Submission{}).Where("id = ?", id).First(s).Error

	return s, err
}
