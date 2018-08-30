package models

import (
	"github.com/mraron/njudge/judge"
	"github.com/jinzhu/gorm"
)

type Judge struct {
	gorm.Model
	State  *judge.Server
	Host   string
	Port   string
	Ping   int
	Online bool
}

/*
func getJudges(db *sqlx.DB, query string, args ...interface{}) (ans []*Judge, err error) {
	err = db.Select(&ans, query, args...)
	return
}
*/
func JudgeFromId(db *gorm.DB, id int) (ans *Judge, err error) {
	ans = new(Judge)
	err = db.Model(&Judge{}).Where("id = ?", id).First(ans).Error
	return
}

func GetJudges(db *gorm.DB) ([]*Judge, error) {
	lst := make([]*Judge, 0)
	err := db.Find(&lst).Error
	return lst, err
}
/*
func (j Judge) Delete(db *sqlx.DB) (err error) {
	_, err = db.Exec("DELETE FROM judges WHERE id=$1", j.ID)
	return
}

func (j Judge) Update(db *sqlx.DB) (err error) {
	_, err = db.Exec("UPDATE judges SET state=$1, host=$2, port=$3, ping=$4, online=$5 WHERE id=$6", j.State, j.Host, j.Port, j.Ping, j.Online, j.ID)
	return
}

func (j *Judge) Insert(db *sqlx.DB) error {
	res := db.QueryRow("INSERT INTO judges (state, host, port, ping, online) VALUES ($1,$2,$3,$4,$5) RETURNING id", j.State, j.Host, j.Port, j.Ping, j.Online)

	var id int64
	err := res.Scan(&id)

	if err != nil {
		return err
	}

	j.ID = id

	return nil
}
*/

// @TODO add filter and count (to every other model too :))
func JudgesAPIGet(db *gorm.DB, _page int, _perPage int, _sortDir string, _sortField string) ([]*Judge, error) {
	lst := make([]*Judge, 0)
	err := db.Model(&Judge{}).Order(_sortField+" "+_sortDir).Limit(_perPage).Offset(_perPage*(_page-1)).Find(&lst).Error
	if err != nil {
		return nil, err
	}

	return lst, nil
}
