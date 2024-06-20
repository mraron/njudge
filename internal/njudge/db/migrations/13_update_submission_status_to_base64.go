package migrations

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems"
	"time"
)

type oldTestcase struct {
	Index          int
	InputPath      string
	OutputPath     string
	AnswerPath     string
	Testset        string
	Group          string
	VerdictName    problems.VerdictName
	Score          float64
	MaxScore       float64
	Output         string
	ExpectedOutput string
	CheckerOutput  string
	TimeSpent      time.Duration
	MemoryUsed     memory.Amount
	TimeLimit      time.Duration
	MemoryLimit    memory.Amount
}

func (t *oldTestcase) ToNewTestcase() problems.Testcase {
	return problems.Testcase{
		Index:          t.Index,
		InputPath:      t.InputPath,
		OutputPath:     t.OutputPath,
		AnswerPath:     t.AnswerPath,
		Testset:        t.Testset,
		Group:          t.Group,
		VerdictName:    t.VerdictName,
		Score:          t.Score,
		MaxScore:       t.MaxScore,
		Output:         problems.Base64String(t.Output),
		ExpectedOutput: problems.Base64String(t.ExpectedOutput),
		CheckerOutput:  problems.Base64String(t.CheckerOutput),
		TimeSpent:      t.TimeSpent,
		MemoryUsed:     t.MemoryUsed,
		TimeLimit:      t.TimeLimit,
		MemoryLimit:    t.MemoryLimit,
	}
}

type oldTestcases []oldTestcase

func (t oldTestcases) ToNewTestcases() []problems.Testcase {
	res := make([]problems.Testcase, 0, len(t))
	for _, testcase := range t {
		res = append(res, testcase.ToNewTestcase())
	}
	return res
}

type oldGroup struct {
	Name         string
	Scoring      problems.ScoringType
	Testcases    oldTestcases
	Dependencies []string
}

func (t *oldGroup) ToNewGroup() problems.Group {
	return problems.Group{
		Name:         t.Name,
		Scoring:      t.Scoring,
		Testcases:    t.Testcases.ToNewTestcases(),
		Dependencies: t.Dependencies,
	}
}

type oldGroups []oldGroup

func (t oldGroups) ToNewGroups() []problems.Group {
	res := make([]problems.Group, 0, len(t))
	for _, testcase := range t {
		res = append(res, testcase.ToNewGroup())
	}
	return res
}

type oldTestset struct {
	Name   string
	Groups oldGroups
}

func (t oldTestset) ToNewTestset() problems.Testset {
	return problems.Testset{
		Name:   t.Name,
		Groups: t.Groups.ToNewGroups(),
	}
}

type oldTestsets []oldTestset

func (t oldTestsets) ToNewTestsets() []problems.Testset {
	res := make([]problems.Testset, 0, len(t))
	for _, testset := range t {
		res = append(res, testset.ToNewTestset())
	}
	return res
}

type oldStatus struct {
	CompilationStatus problems.CompilationStatus
	Compiled          bool
	CompilerOutput    string
	FeedbackType      problems.FeedbackType
	Feedback          oldTestsets
}

func (v *oldStatus) ToNewStatus() problems.Status {
	return problems.Status{
		CompilationStatus: problems.AfterCompilation, // fix older submissions
		Compiled:          v.Compiled,
		CompilerOutput:    problems.Base64String(v.CompilerOutput),
		FeedbackType:      v.FeedbackType,
		Feedback:          v.Feedback.ToNewTestsets(),
	}
}

func (v *oldStatus) Scan(value any) error {
	if value == nil {
		return errors.New("can't scan status from nil")
	}

	if sv, err := driver.String.ConvertValue(value); err == nil {
		var arr []byte

		switch sv := sv.(type) {
		case []uint8:
			orig := sv

			arr = make([]byte, len(orig))
			for i := 0; i < len(orig); i++ {
				arr[i] = byte(orig[i])
			}
		case string:
			arr = []byte(sv)
		}

		if err := json.Unmarshal(arr, v); err != nil {
			return err
		}

		return nil
	}

	return errors.New("can't scan from non string type")
}

func up13(i any) error {
	db := i.(*sql.DB)
	rows, err := db.Query("SELECT id, status from submissions")
	if err != nil {
		return err
	}
	for rows.Next() {
		s := struct {
			id int64
			st *oldStatus
		}{}
		if err := rows.Scan(&s.id, &s.st); err != nil {
			return err
		}
		newStatus := s.st.ToNewStatus()
		_, err := db.Exec("UPDATE submissions SET status = $2 WHERE id = $1", s.id, &newStatus)
		if err != nil {
			return fmt.Errorf("error migration submission %d: %w", s.id, err)
		}
	}
	return nil
}

func down13(i any) error {
	return nil
}
