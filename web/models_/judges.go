package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/judge"
)

type Judge struct {
	Id     int64
	State  *judge.Server
	Host   string
	Port   string
	Ping   int
	Online bool
}

func getJudges(db *sqlx.DB, query string, args ...interface{}) (ans []*Judge, err error) {
	err = db.Select(&ans, query, args...)
	return
}

func JudgeFromId(db *sqlx.DB, id int) (ans *Judge, err error) {
	ans = new(Judge)
	err = db.Get(ans, "SELECT * FROM judges WHERE id=$1", id)
	return
}

func GetJudges(db *sqlx.DB) ([]*Judge, error) {
	return getJudges(db, "SELECT * FROM judges")
}

func (j Judge) Delete(db *sqlx.DB) (err error) {
	_, err = db.Exec("DELETE FROM judges WHERE id=$1", j.Id)
	return
}

func (j Judge) Update(db *sqlx.DB) (err error) {
	_, err = db.Exec("UPDATE judges SET state=$1, host=$2, port=$3, ping=$4, online=$5 WHERE id=$6", j.State, j.Host, j.Port, j.Ping, j.Online, j.Id)
	return
}

func (j *Judge) Insert(db *sqlx.DB) error {
	res := db.QueryRow("INSERT INTO judges (state, host, port, ping, online) VALUES ($1,$2,$3,$4,$5) RETURNING id", j.State, j.Host, j.Port, j.Ping, j.Online)

	var id int64
	err := res.Scan(&id)

	if err != nil {
		return err
	}

	j.Id = id

	return nil
}

func JudgesAPIGet(db *sqlx.DB, _page int, _perPage int, _sortDir string, _sortField string) ([]*Judge, error) {
	lst, err := getJudges(db, fmt.Sprintf("SELECT * FROM judges ORDER BY %s %s LIMIT %d OFFSET %d ", _sortField, _sortDir, _perPage, _perPage*(_page-1))) //@TODO: SQL Injection!!!
	if err != nil {
		return nil, err
	}

	return lst, nil
}
