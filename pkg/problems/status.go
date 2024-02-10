package problems

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"math"
	"time"
)

// VerdictName represents the verdict of a testcase i.e the outcome which happened in result of running the testcase (or the lack of running it in the case of VerdictDR)
type VerdictName int

const (
	VerdictAC VerdictName = iota
	VerdictWA
	VerdictRE
	VerdictTL
	VerdictML
	VerdictXX
	VerdictDR
	VerdictPC
	VerdictPE
)

func (v VerdictName) String() string {
	switch v {
	case VerdictAC:
		return "Elfogadva"
	case VerdictWA:
		return "Rossz válasz"
	case VerdictRE:
		return "Futási hiba"
	case VerdictTL:
		return "Időlimit túllépés"
	case VerdictML:
		return "Memória limit túllépés"
	case VerdictXX:
		return "Belső hiba"
	case VerdictDR:
		return "Nem futott"
	case VerdictPC:
		return "Részben elfogadva"
	case VerdictPE:
		return "Prezentációs hiba"
	}

	return "..."
}

// FeedbackType is mainly for displaying to the end user.
// In FeedbackCF we actually output the contestant's output and the jury's output,
// too whereas in for example FeedbackACM we only use standard ACM feedback (just the verdict),
// in FeedbackIOI we display all testcases along information about groups.
type FeedbackType int

const (
	FeedbackCF FeedbackType = iota
	FeedbackIOI
	FeedbackACM
	FeedbackLazyIOI
)

// FeedbackFromString parses a string into FeedbackType, the default is FEEDBACK_CF, "ioi" is for FeedbackIOI and "acm" is for FeedbackACM
func FeedbackFromString(str string) FeedbackType {
	if str == "ioi" {
		return FeedbackIOI
	} else if str == "acm" {
		return FeedbackACM
	} else if str == "lazyioi" {
		return FeedbackLazyIOI
	}

	return FeedbackCF
}

func (f FeedbackType) String() string {
	switch f {
	case FeedbackACM:
		return "FeedbackACM"
	case FeedbackCF:
		return "FeedbackCF"
	case FeedbackIOI:
		return "FeedbackIOI"
	case FeedbackLazyIOI:
		return "FeedbackLazyIOI"
	}

	return fmt.Sprintf("Feedback(%d)", f)
}

// ScoringType represents the scoring of a group of tests,
//
//   - ScoringGroup means that if there's a non-accepted (or partially accepted) testcase in the group then the whole group scores 0 points,
//   - ScoringSum means that the score of the group is the sum of scores of individual scores.
type ScoringType int

const (
	ScoringGroup ScoringType = iota
	ScoringSum
	ScoringMin
)

func ScoringFromString(str string) ScoringType {
	if str == "group" {
		return ScoringGroup
	} else if str == "min" {
		return ScoringMin
	}

	return ScoringSum
}

// Testcase represents a testcase in the status of a submission.
type Testcase struct {
	Index          int
	InputPath      string
	OutputPath     string
	AnswerPath     string
	Testset        string
	Group          string
	VerdictName    VerdictName
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

// Testset represents some set of tests, for example pretests and system tests should be testsets.
// They can also be used to seamlessly rejudge for contestants.
type Testset struct {
	Name   string
	Groups []Group
}

func (ts Testset) Verdict() VerdictName {
	if ts.IsAC() {
		return VerdictAC
	}
	return VerdictName(ts.IndexTestcase(ts.FirstNonAC()).VerdictName)
}

func (ts Testset) IndexTestcase(ind int) *Testcase {
	curr := 1
	tcs := ts.Testcases()
	for idx := range tcs {
		if curr == ind {
			return tcs[idx]
		}

		curr++
	}

	return &Testcase{VerdictName: VerdictDR}
}

func (ts Testset) Testcases() (testcases []*Testcase) {
	testcaseCount := 0
	for i := range ts.Groups {
		testcaseCount += len(ts.Groups[i].Testcases)
	}

	testcases = make([]*Testcase, 0, testcaseCount)
	for i := range ts.Groups {
		for j := range ts.Groups[i].Testcases {
			testcases = append(testcases, &ts.Groups[i].Testcases[j])
		}
	}

	return
}

func (ts *Testset) SetTimeLimit(tl time.Duration) {
	for _, tc := range ts.Testcases() {
		tc.TimeLimit = tl
	}

	for ind := range ts.Groups {
		ts.Groups[ind].SetTimeLimit(tl)
	}
}

func (ts *Testset) SetMemoryLimit(ml memory.Amount) {
	for _, tc := range ts.Testcases() {
		tc.MemoryLimit = ml
	}

	for ind := range ts.Groups {
		ts.Groups[ind].SetMemoryLimit(ml)
	}
}

func (ts Testset) Score() (res float64) {
	for _, g := range ts.Groups {
		res += g.Score()
	}

	return
}

func (ts Testset) MaxScore() (res float64) {
	for _, g := range ts.Groups {
		res += g.MaxScore()
	}

	return
}

func (ts Testset) FirstNonAC() int {
	until := 0
	for _, g := range ts.Groups {
		if g.FirstNonAC() != -1 {
			return g.FirstNonAC() + until
		}
		until += len(g.Testcases)
	}

	return -1
}

func (ts Testset) MaxMemoryUsage() memory.Amount {
	mx := memory.Amount(0)
	for _, g := range ts.Groups {
		if mx < g.MaxMemoryUsage() {
			mx = g.MaxMemoryUsage()
		}
	}

	return mx
}

func (ts Testset) IsAC() bool {
	return ts.FirstNonAC() == -1
}

func (ts Testset) MaxTimeSpent() time.Duration {
	mx := time.Duration(0)
	for _, g := range ts.Groups {
		if mx < g.MaxTimeSpent() {
			mx = g.MaxTimeSpent()
		}
	}

	return mx
}

// A Group is named group of tests for which there is a scoring policy defined.
type Group struct {
	Name         string
	Scoring      ScoringType
	Testcases    []Testcase
	Dependencies []string
}

func (g *Group) SetTimeLimit(tl time.Duration) {
	for ind := range g.Testcases {
		g.Testcases[ind].TimeLimit = tl
	}
}

func (g *Group) SetMemoryLimit(ml memory.Amount) {
	for ind := range g.Testcases {
		g.Testcases[ind].MemoryLimit = ml
	}
}

func (g Group) Score() float64 {
	sum := 0.0
	for _, val := range g.Testcases {
		sum += val.Score
	}

	switch g.Scoring {
	case ScoringGroup:
		for _, val := range g.Testcases {
			if val.VerdictName != VerdictAC && val.VerdictName != VerdictPC {
				return 0.0
			}
		}

		return sum
	case ScoringSum:
		return sum
	case ScoringMin:
		res := math.MaxFloat64
		for _, val := range g.Testcases {
			if val.Score < res {
				res = val.Score
			}
		}

		return res
	}

	return -1.0
}

func (g Group) MaxScore() float64 {
	sum := 0.0
	for _, val := range g.Testcases {
		sum += val.MaxScore
		if g.Scoring == ScoringMin {
			sum = val.MaxScore
		}
	}

	return sum
}

func (g Group) FirstNonAC() int {
	for ind, val := range g.Testcases {
		if val.VerdictName != VerdictAC && val.VerdictName != VerdictDR {
			return ind + 1
		}
	}

	return -1
}

func (g Group) IsAC() bool {
	return g.FirstNonAC() == -1
}

func (g Group) MaxMemoryUsage() memory.Amount {
	mx := memory.Amount(0)
	for _, val := range g.Testcases {
		if mx < val.MemoryUsed {
			mx = val.MemoryUsed
		}
	}

	return mx
}

func (g Group) MaxTimeSpent() time.Duration {
	mx := time.Duration(0)
	for _, val := range g.Testcases {
		if mx < val.TimeSpent {
			mx = val.TimeSpent
		}
	}

	return mx
}

// A Status represents the status of a submission after judging
// It contains the information about compilation and the feedback
// The main Testset is always the first one in Feedback.
type Status struct {
	Compiled       bool
	CompilerOutput string
	FeedbackType   FeedbackType
	Feedback       []Testset
}

func (v Status) Value() (driver.Value, error) {
	buf := &bytes.Buffer{}

	enc := json.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.String(), nil
}

func (v *Status) Scan(value interface{}) error {
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
