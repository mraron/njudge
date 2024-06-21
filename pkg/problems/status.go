package problems

import (
	"bytes"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"math"
	"time"
)

// VerdictName represents the verdict of a testcase i.e. the outcome which happened in result of running the testcase (or the lack of running it in the case of VerdictDR)
type VerdictName string

const (
	VerdictUnknown VerdictName = ""
	VerdictAC      VerdictName = "AC"
	VerdictWA      VerdictName = "WA"
	VerdictRE      VerdictName = "RE"
	VerdictTL      VerdictName = "TL"
	VerdictML      VerdictName = "ML"
	VerdictXX      VerdictName = "XX"
	VerdictDR      VerdictName = "DR"
	VerdictPC      VerdictName = "PC"
	VerdictPE      VerdictName = "PE"
	VerdictSK      VerdictName = "SK"
)

func (v *VerdictName) UnmarshalJSON(i []byte) error {
	switch string(i) {
	case "0", fmt.Sprintf("%q", VerdictAC):
		*v = VerdictAC
	case "1", fmt.Sprintf("%q", VerdictWA):
		*v = VerdictWA
	case "2", fmt.Sprintf("%q", VerdictRE):
		*v = VerdictRE
	case "3", fmt.Sprintf("%q", VerdictTL):
		*v = VerdictTL
	case "4", fmt.Sprintf("%q", VerdictML):
		*v = VerdictML
	case "5", fmt.Sprintf("%q", VerdictXX):
		*v = VerdictXX
	case "6", fmt.Sprintf("%q", VerdictDR):
		*v = VerdictDR
	case "7", fmt.Sprintf("%q", VerdictPC):
		*v = VerdictPC
	case "8", fmt.Sprintf("%q", VerdictPE):
		*v = VerdictPE
	case fmt.Sprintf("%q", VerdictSK):
		*v = VerdictSK
	case "null":
	default:
		return fmt.Errorf("unknown VerdictName: %q", i)
	}
	return nil
}

// FeedbackType is mainly for displaying to the end user.
// In FeedbackCF we actually output the contestant's output and the jury's output,
// too whereas in for example FeedbackACM we only use standard ACM feedback (just the verdict),
// in FeedbackIOI we display all testcases along information about groups.
type FeedbackType string

const (
	FeedbackUnknown FeedbackType = ""
	FeedbackCF      FeedbackType = "FeedbackCF"
	FeedbackIOI     FeedbackType = "FeedbackIOI"
	FeedbackACM     FeedbackType = "FeedbackACM"
	FeedbackLazyIOI FeedbackType = "FeedbackLazyIOI"
)

// FeedbackTypeFromShortString parses a string into FeedbackType, the default is FEEDBACK_CF, "ioi" is for FeedbackIOI and "acm" is for FeedbackACM
func FeedbackTypeFromShortString(str string) FeedbackType {
	if str == "ioi" {
		return FeedbackIOI
	} else if str == "acm" {
		return FeedbackACM
	} else if str == "lazyioi" {
		return FeedbackLazyIOI
	}

	return FeedbackCF
}

func (f *FeedbackType) UnmarshalJSON(i []byte) error {
	switch string(i) {
	case "0", fmt.Sprintf("%q", FeedbackCF):
		*f = FeedbackCF
	case "1", fmt.Sprintf("%q", FeedbackIOI):
		*f = FeedbackIOI
	case "2", fmt.Sprintf("%q", FeedbackACM):
		*f = FeedbackACM
	case "3", fmt.Sprintf("%q", FeedbackLazyIOI):
		*f = FeedbackLazyIOI
	case "null", "\"\"":
		fallthrough
	default:
		*f = FeedbackUnknown
	}
	return nil
}

// ScoringType represents the scoring of a group of tests,
//
//   - ScoringGroup means that if there's a non-accepted (or partially accepted) testcase in the group then the whole group scores 0 points,
//   - ScoringSum means that the score of the group is the sum of the scores.
//   - ScoringMin means that the score of the group is the minimum of the scores.
type ScoringType string

const (
	ScoringUnknown ScoringType = ""
	ScoringGroup   ScoringType = "ScoringGroup"
	ScoringSum     ScoringType = "ScoringSum"
	ScoringMin     ScoringType = "ScoringMin"
)

func ScoringFromString(str string) ScoringType {
	if str == "group" {
		return ScoringGroup
	} else if str == "min" {
		return ScoringMin
	}

	return ScoringSum
}

func (s *ScoringType) UnmarshalJSON(i []byte) error {
	switch string(i) {
	case "0", fmt.Sprintf("%q", ScoringGroup):
		*s = ScoringGroup
	case "1", fmt.Sprintf("%q", ScoringSum):
		*s = ScoringSum
	case "2", fmt.Sprintf("%q", ScoringMin):
		*s = ScoringMin
	case "null":
	default:
		return fmt.Errorf("unknown ScoringType: %q", i)
	}
	return nil
}

// Base64String is used to store certain specific characters (null bytes and such) in DBMS like PostgreSQL which might
// sometimes be produced by faulty programs.
type Base64String string

func (b *Base64String) UnmarshalJSON(i []byte) error {
	if len(i) < 2 {
		return errors.New("Base64String too short")
	}
	if i[0] != '"' || i[len(i)-1] != '"' {
		return errors.New("Base64String not quoted")
	}
	res, err := base64.StdEncoding.DecodeString(string(i[1 : len(i)-1]))
	if err != nil {
		return err
	}
	*b = Base64String(res)
	return nil
}

func (b *Base64String) MarshalJSON() ([]byte, error) {
	res := "\"" + base64.StdEncoding.EncodeToString([]byte(*b)) + "\""
	return []byte(res), nil
}

func (b *Base64String) String() string {
	return string(*b)
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
	Output         Base64String
	ExpectedOutput Base64String
	CheckerOutput  Base64String
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

func (ts *Testset) Verdict() VerdictName {
	if ts.IsAC() {
		return VerdictAC
	}
	return ts.IndexTestcase(ts.FirstNonAC()).VerdictName
}

func (ts *Testset) IndexTestcase(ind int) *Testcase {
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

func (ts *Testset) Testcases() (testcases []*Testcase) {
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

func (ts *Testset) Score() (res float64) {
	for _, g := range ts.Groups {
		res += g.Score()
	}

	return
}

func (ts *Testset) MaxScore() (res float64) {
	for _, g := range ts.Groups {
		res += g.MaxScore()
	}

	return
}

func (ts *Testset) FirstNonAC() int {
	until := 0
	for _, g := range ts.Groups {
		if g.FirstNonAC() != -1 {
			return g.FirstNonAC() + until
		}
		until += len(g.Testcases)
	}

	return -1
}

func (ts *Testset) MaxMemoryUsage() memory.Amount {
	mx := memory.Amount(0)
	for _, g := range ts.Groups {
		if mx < g.MaxMemoryUsage() {
			mx = g.MaxMemoryUsage()
		}
	}

	return mx
}

func (ts *Testset) IsAC() bool {
	return ts.FirstNonAC() == -1
}

func (ts *Testset) MaxTimeSpent() time.Duration {
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

func (g *Group) Score() float64 {
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

func (g *Group) MaxScore() float64 {
	sum := 0.0
	for _, val := range g.Testcases {
		sum += val.MaxScore
		if g.Scoring == ScoringMin {
			sum = val.MaxScore
		}
	}

	return sum
}

func (g *Group) FirstNonAC() int {
	for ind, val := range g.Testcases {
		if val.VerdictName != VerdictAC && val.VerdictName != VerdictDR {
			return ind + 1
		}
	}

	return -1
}

func (g *Group) IsAC() bool {
	return g.FirstNonAC() == -1
}

func (g *Group) MaxMemoryUsage() memory.Amount {
	mx := memory.Amount(0)
	for _, val := range g.Testcases {
		if mx < val.MemoryUsed {
			mx = val.MemoryUsed
		}
	}

	return mx
}

func (g *Group) MaxTimeSpent() time.Duration {
	mx := time.Duration(0)
	for _, val := range g.Testcases {
		if mx < val.TimeSpent {
			mx = val.TimeSpent
		}
	}

	return mx
}

type CompilationStatus int

const (
	BeforeCompilation CompilationStatus = 0
	DuringCompilation CompilationStatus = 1
	AfterCompilation  CompilationStatus = 2
)

// A Status represents the status of a submission after judging
// It contains the information about compilation and the feedback
// The main Testset is always the first one in Feedback.
type Status struct {
	CompilationStatus CompilationStatus
	Compiled          bool
	CompilerOutput    Base64String
	FeedbackType      FeedbackType
	Feedback          []Testset
}

func (v *Status) Value() (driver.Value, error) {
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
