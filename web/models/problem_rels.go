package models

import (
	"github.com/jmoiron/sqlx"
)

type ProblemRel struct {
	Problemset string
	Problem    string
}

func getProblemRels(db *sqlx.DB, query string, args ...interface{}) (ans []ProblemRel, err error) {
	err = db.Select(&ans, query, args...)
	return
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
