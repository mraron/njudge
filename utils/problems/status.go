package problems

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// VerdictName represents the verdict of a testcase i.e the outcome which happened in result of running the testcase (or the lack of running it in the case of VERDICT_DR)
type VerdictName int

const (
	VERDICT_AC VerdictName = iota
	VERDICT_WA
	VERDICT_RE
	VERDICT_TL
	VERDICT_ML
	VERDICT_XX
	VERDICT_DR
	VERDICT_PC
	VERDICT_PE
)

func (v VerdictName) String() string {
	switch v {
	case VERDICT_AC:
		return "Elfogadva"
	case VERDICT_WA:
		return "Rossz válasz"
	case VERDICT_RE:
		return "Futási hiba"
	case VERDICT_TL:
		return "Időlimit túllépés"
	case VERDICT_ML:
		return "Memória limit túllépés"
	case VERDICT_XX:
		return "Belső hiba"
	case VERDICT_DR:
		return "Nem fut"
	case VERDICT_PC:
		return "Részben elfogadva"
	case VERDICT_PE:
		return "Prezentációs hiba"
	}

	return "..."
}

// FeedbackType is mainly for displaying to the end user.
// In FEEDBACK_CF we actually output the contestant's output and the jury's output,
// too whereas in for example FEEDBACK_ACM we only use standard ACM feedback (just the verdict),
// in FEEDBACK_IOI we display all testcases along information about groups.
type FeedbackType int

const (
	FEEDBACK_CF FeedbackType = iota
	FEEDBACK_IOI
	FEEDBACK_ACM
)

// FeedbackFromString parses a string into FeedbackType, the default is FEEDBACK_CF, "ioi" is for FEEDBACK_IOI and "acm" is for FEEDBACK_ACM
func FeedbackFromString(str string) FeedbackType {
	if str == "ioi" {
		return FEEDBACK_IOI
	} else if str == "acm" {
		return FEEDBACK_ACM
	}

	return FEEDBACK_CF
}

// ScoringType represents the scoring of a group of tests,
// * SCORING_GROUP means that if there's a non accepted testcase in the group then the whole group scores 0 points,
// * SCORING_SUM means that the score of the group is the sum of scores of individual scores.
type ScoringType int

const (
	SCORING_GROUP ScoringType = iota
	SCORING_SUM
)

func ScoringFromString(str string) ScoringType {
	if str == "group" {
		return SCORING_GROUP
	}

	return SCORING_SUM
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
	MemoryUsed     int
	TimeLimit      time.Duration
	MemoryLimit    int
}

// Testset represents some set of tests, for example pretests and system tests should be testsets.
// They are also used to seamlessly rejudge for contestants.
type Testset struct {
	Name      string
	Groups    []Group
	Testcases []Testcase
}

func (ts *Testset) SetTimeLimit(tl time.Duration) {
	for ind := range ts.Testcases {
		ts.Testcases[ind].TimeLimit = tl
	}

	for ind := range ts.Groups {
		ts.Groups[ind].SetTimeLimit(tl)
	}
}

func (ts *Testset) SetMemoryLimit(ml int) {
	for ind := range ts.Testcases {
		ts.Testcases[ind].MemoryLimit = ml
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
	for _, g := range ts.Groups {
		if g.FirstNonAC() != -1 {
			return g.FirstNonAC()
		}
	}

	return -1
}

func (ts Testset) MaxMemoryUsage() int {
	mx := 0
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
	Dependencies []string //@TODO: actually support this while calculating score
}

func (g *Group) SetTimeLimit(tl time.Duration) {
	for ind := range g.Testcases {
		g.Testcases[ind].TimeLimit = tl
	}
}

func (g *Group) SetMemoryLimit(ml int) {
	for ind := range g.Testcases {
		g.Testcases[ind].MemoryLimit = ml
	}
}

func (g Group) Score() float64 {
	switch g.Scoring {
	case SCORING_GROUP:
		for _, val := range g.Testcases {
			if val.VerdictName != VERDICT_AC {
				return 0.0
			}
		}

		return g.MaxScore()
	case SCORING_SUM:
		sum := 0.0
		for _, val := range g.Testcases {
			sum += val.Score
		}

		return sum
	}

	return -1.0
}

func (g Group) MaxScore() float64 {
	sum := 0.0
	for _, val := range g.Testcases {
		sum += val.MaxScore
	}

	return sum
}

func (g Group) FirstNonAC() int {
	for ind, val := range g.Testcases {
		if val.VerdictName != VERDICT_AC {
			return ind + 1
		}
	}

	return -1
}

func (g Group) IsAC() bool {
	return g.FirstNonAC() == -1
}

func (g Group) MaxMemoryUsage() int {
	mx := 0
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
// (Multiple testsets are currently outside of focus, but possible to implement)
type Status struct {
	Compiled       bool
	CompilerOutput string
	FeedbackType   FeedbackType
	Feedback       []Testset
}

func (v Status) Score() (ans float64) {
	for _, f := range v.Feedback {
		ans += f.Score()
	}

	return
}

func (v Status) MaxScore() (ans float64) {
	for _, f := range v.Feedback {
		ans += f.MaxScore()
	}

	return
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

		switch sv.(type) {
		case []uint8:
			orig := sv.([]uint8)

			arr = make([]byte, len(orig))
			for i := 0; i < len(orig); i++ {
				arr[i] = byte(orig[i])
			}
		case string:
			arr = []byte(sv.(string))
		}

		if err := json.Unmarshal(arr, v); err != nil {
			return err
		}

		return nil
	}

	return errors.New("can't scan from non string type")
}

func (v Status) Verdict() VerdictName {
	if v.FirstNonAC() == -1 {
		return VERDICT_AC
	}

	return v.IndexTestcase(v.FirstNonAC()).VerdictName
}

func (v Status) IsAC() bool {
	for _, val := range v.Feedback {
		if !val.IsAC() {
			return false
		}
	}

	return true
}

func (v Status) MaxMemoryUsage() int {
	res := 0
	for _, val := range v.Feedback {
		if res < val.MaxMemoryUsage() {
			res = val.MaxMemoryUsage()
		}
	}

	return res
}

func (v Status) MaxTimeSpent() time.Duration {
	res := time.Duration(0)
	for _, val := range v.Feedback {
		if res < val.MaxTimeSpent() {
			res = val.MaxTimeSpent()
		}
	}

	return res
}

func (v Status) FirstNonAC() int {
	ind := 1
	for _, val := range v.Feedback {
		for _, val2 := range val.Testcases {
			if val2.VerdictName != VERDICT_AC {
				return ind
			}

			ind++
		}
	}

	return -1
}

func (v Status) IndexTestcase(ind int) Testcase {
	curr := 1
	for _, val := range v.Feedback {
		for _, val2 := range val.Testcases {
			if curr == ind {
				return val2
			}

			curr++
		}
	}

	return Testcase{}
}
