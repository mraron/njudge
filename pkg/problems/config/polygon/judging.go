package polygon

import (
	"fmt"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/problems/evaluation/output_only"
	"github.com/mraron/njudge/pkg/problems/executable/checker"
	"path/filepath"
	"time"

	"github.com/mraron/njudge/pkg/problems"
)

type Test struct {
	Method string  `xml:"method,attr"`
	Cmd    string  `xml:"cmd,attr"`
	Sample bool    `xml:"sample,attr"`
	Points float64 `xml:"points,attr"`
	Group  string  `xml:"group,attr"`

	Input  string
	Answer string
	Index  int
}

func (tc Test) Testcase() problems.Testcase {
	return problems.Testcase{
		InputPath:   tc.Input,
		AnswerPath:  tc.Answer,
		VerdictName: problems.VerdictDR,
		MaxScore:    tc.Points,
	}
}

type Dependency struct {
	Group string `xml:"group,attr"`
}

type Group struct {
	Name         string       `xml:"name,attr"`
	PointsPolicy string       `xml:"points-policy,attr"`
	Points       float64      `xml:"points,attr"`
	Dependencies []Dependency `xml:"dependencies>dependency"`
}

type Testset struct {
	Name              string  `xml:"name,attr"`
	TimeLimit         int     `xml:"time-limit"`
	MemoryLimit       int     `xml:"memory-limit"`
	TestCount         int     `xml:"test-count"`
	InputPathPattern  string  `xml:"input-path-pattern"`
	AnswerPathPattern string  `xml:"answer-path-pattern"`
	Tests             []Test  `xml:"tests>test"`
	Groups            []Group `xml:"groups>group"`
}

func (ts Testset) Testset(path string) problems.Testset {
	testset := problems.Testset{Name: ts.Name}

	testcases := make(map[string][]Test)
	for ind, tc := range ts.Tests {
		tc.Input, tc.Answer = fmt.Sprintf(filepath.Join(path, ts.InputPathPattern), ind+1), fmt.Sprintf(filepath.Join(path, ts.AnswerPathPattern), ind+1)
		tc.Index = ind + 1

		if len(testcases[tc.Group]) == 0 {
			testcases[tc.Group] = make([]Test, 0)
		}

		testcases[tc.Group] = append(testcases[tc.Group], tc)
	}

	if len(ts.Groups) == 0 {
		ts.Groups = append(ts.Groups, Group{Name: "", PointsPolicy: "sum", Points: -1.0})
	}

	idx := 1
	for _, grp := range ts.Groups {
		testset.Groups = append(testset.Groups, problems.Group{})
		group := &testset.Groups[len(testset.Groups)-1]

		group.Name = grp.Name
		if grp.PointsPolicy == "complete-group" {
			group.Scoring = problems.ScoringGroup
		} else if grp.PointsPolicy == "min-group" {
			group.Scoring = problems.ScoringMin
		} else {
			group.Scoring = problems.ScoringSum
		}

		group.Dependencies = make([]string, 0, len(grp.Dependencies))
		for _, dep := range grp.Dependencies {
			group.Dependencies = append(group.Dependencies, dep.Group)
		}

		for _, tc := range testcases[grp.Name] {
			testcase := tc.Testcase()

			testcase.Index = idx
			testcase.Testset = ts.Name
			testcase.Group = group.Name
			testcase.TimeLimit = time.Duration(ts.TimeLimit) * time.Millisecond
			testcase.MemoryLimit = memory.Amount(ts.MemoryLimit)

			if group.Scoring == problems.ScoringMin {
				testcase.MaxScore = grp.Points
			}

			group.Testcases = append(group.Testcases, testcase)

			idx++
		}
	}

	return testset
}

func (ts Testset) Group(name string) (*Group, error) {
	for ind, g := range ts.Groups {
		if g.Name == name {
			return &ts.Groups[ind], nil
		}
	}

	return nil, fmt.Errorf("can't find group %s", name)
}

type Judging struct {
	CpuName    string `xml:"cpu-name,attr"`
	CpuSpeed   int    `xml:"cpu-speed,attr"`
	InputFile  string `xml:"input-file,attr"`
	OutputFile string `xml:"output-file,attr"`

	Testsets []Testset `xml:"testset"`
}

func (j Judging) GetTestset(name string) (*Testset, error) {
	for _, ts := range j.Testsets {
		if name == ts.Name {
			return &ts, nil
		}
	}

	return nil, fmt.Errorf("no such testset %q", name)
}

func (p Problem) StatusSkeleton(name string) (*problems.Status, error) {
	if name == "" {
		name = "tests"
	}

	testset, err := p.Judging.GetTestset(name)
	if err != nil {
		return nil, err
	}

	feedback := []problems.Testset{testset.Testset(p.Path)}
	if p.TaskType == output_only.Name {
		for _, tc := range feedback[0].Testcases() {
			tc.OutputPath = filepath.Base(tc.AnswerPath)
		}
	}

	return &problems.Status{
		Compiled:       false,
		CompilerOutput: "status skeleton",
		FeedbackType:   problems.FeedbackTypeFromShortString(p.FeedbackType),
		Feedback:       feedback,
	}, nil
}

func (p Problem) Checker() problems.Checker {
	return checker.NewTestlib(filepath.Join(p.Path, "check"))
}
