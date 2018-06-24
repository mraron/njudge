package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProblemRel struct {
	Id         int64
	Problemset string
	Problem    string
}

func getProblemRels(db *sqlx.DB, query string, args ...interface{}) (ans []ProblemRel, err error) {
	err = db.Select(&ans, query, args...)
	return
}

func ProblemRelFromId(db *sqlx.DB, id int) (ans ProblemRel, err error) {
	err = db.Get(&ans, "SELECT * FROM problem_rels WHERE id=$1", id)
	return
}

func (pr ProblemRel) Delete(db *sqlx.DB) (err error) {
	_, err = db.Exec("DELETE FROM problem_rels WHERE id=$1", pr.Id)
	return
}

func (pr ProblemRel) Update(db *sqlx.DB) (err error) {
	_, err = db.Exec("UPDATE problem_rels SET problemset=$1, problem=$2 WHERE id=$3", pr.Problemset, pr.Problem, pr.Id)
	return
}

func (pr *ProblemRel) Insert(db *sqlx.DB) error {
	res := db.QueryRow("INSERT INTO problem_rels (problemset, problem) VALUES ($1,$2) RETURNING id", pr.Problemset, pr.Problem)

	var id int64
	err := res.Scan(&id)

	if err != nil {
		fmt.Println(err, "safsdf2")
		return err
	}

	pr.Id = id

	return nil
}

func ProblemsFromProblemset(db *sqlx.DB, problemset string) ([]string, error) {
	lst, err := getProblemRels(db, "SELECT * FROM problem_rels WHERE problemset=$1", problemset)
	if err != nil {
		return nil, err
	}

	ret := make([]string, len(lst))
	for ind, val := range lst {
		ret[ind] = val.Problem
	}

	return ret, nil
}

func ProblemRelAPIGet(db *sqlx.DB, _page int, _perPage int, _sortDir string, _sortField string) ([]ProblemRel, error) {
	fmt.Println(fmt.Sprintf("SELECT * FROM problem_rels ORDER BY %s %s LIMIT %d OFFSET %d ", _sortField, _sortDir, _perPage, _perPage*(_page-1)))
	lst, err := getProblemRels(db, fmt.Sprintf("SELECT * FROM problem_rels ORDER BY %s %s LIMIT %d OFFSET %d ", _sortField, _sortDir, _perPage, _perPage*(_page-1))) //@TODO: SQL Injection!!!
	if err != nil {
		return nil, err
	}

	return lst, nil
}
