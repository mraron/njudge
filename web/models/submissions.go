package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/mraron/njudge/utils/problems"
	"strconv"
	"time"
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
	Id         int64
	Status     problems.Status
	Verdict    Verdict
	OnTest     sql.NullString `db:"ontest"`
	User       *User
	Submitted  time.Time
	Judged     pq.NullTime
	Problemset string
	Problem    string
	Language   string
	Private    bool
	Source     string
	Started    bool
}

func SubmissionFromId(db *sqlx.DB, id int64) (*Submission, error) {
	s := new(Submission)
	err := db.Get(s, "SELECT * FROM submissions WHERE id=$1", id)

	return s, err
}
