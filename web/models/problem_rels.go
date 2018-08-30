package models

import (
	"github.com/jinzhu/gorm"
)

type ProblemRel struct {
	gorm.Model
	Problemset string
	Problem    string
}

/*
func getProblemRels(db *sqlx.DB, query string, args ...interface{}) (ans []ProblemRel, err error) {
	err = db.Select(&ans, query, args...)
	return
}
*/
func ProblemRelFromId(db *gorm.DB, id int) (ans ProblemRel, err error) {
	err = db.Model(&ProblemRel{}).Where("id = ?", id).First(&ans).Error
	return
}
/*
func (pr ProblemRel) Delete(db *sqlx.DB) (err error) {
	_, err = db.Exec("DELETE FROM problem_rels WHERE id=$1", pr.ID)
	return
}

func (pr ProblemRel) Update(db *sqlx.DB) (err error) {
	_, err = db.Exec("UPDATE problem_rels SET problemset=$1, problem=$2 WHERE id=$3", pr.Problemset, pr.Problem, pr.ID)
	return
}

func (pr *ProblemRel) Insert(db *sqlx.DB) error {
	res := db.QueryRow("INSERT INTO problem_rels (problemset, problem) VALUES ($1,$2) RETURNING id", pr.Problemset, pr.Problem)

	var id int64
	err := res.Scan(&id)

	if err != nil {
		return err
	}

	pr.ID = id

	return nil
} */

func ProblemsFromProblemset(db *gorm.DB, problemset string) ([]string, error) {
	lst := make([]ProblemRel, 0)
	err := db.Model(&ProblemRel{}).Where("problemset = ?", problemset).Find(&lst).Error
	if err != nil {
		return nil, err
	}

	ret := make([]string, len(lst))
	for ind, val := range lst {
		ret[ind] = val.Problem
	}

	return ret, nil
}

func ProblemRelAPIGet(db *gorm.DB, _page int, _perPage int, _sortDir string, _sortField string) ([]ProblemRel, error) {
	lst := make([]ProblemRel, 0)
	err := db.Model(&ProblemRel{}).Order(_sortField+" "+_sortDir).Limit(_perPage).Offset(_perPage*(_page-1)).Find(&lst).Error
	if err != nil {
		return nil, err
	}

	return lst, nil
}
