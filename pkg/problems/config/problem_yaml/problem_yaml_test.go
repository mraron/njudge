package problem_yaml_test

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/config/problem_yaml"
	"github.com/spf13/afero"
)

func TestParsing(t *testing.T) {
	problemYAML := `{
		"shortname": "covering",
		"titles": [
			{"title": "T-lefedés", "language": "hungarian"}
		],
		"statements": [
			{"path": "covering-HUN.pdf", "language": "hungarian", "type": "application/pdf"}
		],
		"tests": {
			"feedback_type": "ioi",
			"time_limit": 1.0,
			"memory_limit": 100,
			"subtasks": [
				{
					"test_count": 5,
					"input_pattern": "input.test_01_%02d",
					"output_pattern": "output.test_01_%02d",
					"scoring": "group",
					"max_score": 5
				},
				{
					"test_count": 5,
					"input_pattern": "input.test_02_%02d",
					"output_pattern": "output.test_02_%02d",
					"scoring": "group",
					"max_score": 10
				},
				{
					"test_count": 5,
					"input_list": ["1.in", "2.in"],
					"output_list": ["1.out", "2.out"],
					"max_score": 85
				}
			]
		}
	}`

	fs := afero.NewMemMapFs()

	_ = afero.WriteFile(fs, "covering-HUN.pdf", []byte(""), 0777)

	parser, identifier := problem_yaml.ParserAndIdentifier()

	if identifier(fs, "./") {
		t.Fatal("can identify???")
	}
	_ = afero.WriteFile(fs, "problem.yaml", []byte(problemYAML), 0777)
	if !identifier(fs, "./") {
		t.Fatal("can't identify")
	}

	p, err := parser(fs, "./")
	if err != nil {
		t.Fatal(err)
	}

	if p.MemoryLimit() != 100*1024*1024 {
		t.Error(p.MemoryLimit())
	}

	if p.TimeLimit() != 1000 {
		t.Error(p.TimeLimit())
	}

	sk, err := p.StatusSkeleton("")
	if err != nil {
		t.Error(err)
	}

	if sk.FeedbackType != problems.FeedbackIOI {
		t.Error(sk.FeedbackType)
	}

	if x := sk.Feedback[0].IndexTestcase(5).InputPath; x != "tests/input.test_01_05" {
		t.Error(x)
	}

	if x := sk.Feedback[0].IndexTestcase(9).InputPath; x != "tests/input.test_02_04" {
		t.Error(x)
	}

	if x := sk.Feedback[0].Groups[0].Scoring; x != problems.ScoringGroup {
		t.Error(x)
	}

	if x := sk.Feedback[0].IndexTestcase(11); x.InputPath != "tests/1.in" || x.MaxScore != 85.0/2.0 {
		t.Error(x)
	}

	problemYAML = `{
		"shortname": "covering",
		"titles": [
			{"title": "T-lefedés", "language": "hungarian"}
		],
		"statements": [
			{"path": "covering-HUN.pdf", "language": "hungarian", "type": "application/pdf"}
		],
		"attachments": [
			{"name": "minta.zip", "path": "minta.zip"}
		],
		"tests": {
			"feedback_type": "ioi",
			"time_limit": 1.0,
			"memory_limit": 100,
			"input_pattern": "%d.in",
			"output_pattern": "%d.out",
			"subtasks": [
				{
					"test_count": 5,
					"scoring": "group",
					"max_score": 5
				},
				{
					"test_count": 5,
					"scoring": "group",
					"max_score": 10
				},
				{
					"test_count": 5,
					"max_score": 85
				}
			]
		}
	}`

	_ = afero.WriteFile(fs, "problem.yaml", []byte(problemYAML), 0777)
	_ = afero.WriteFile(fs, "minta.zip", []byte(""), 0777)
	p, err = parser(fs, "./")
	if err != nil {
		t.Fatal(err)
	}

	sk, err = p.StatusSkeleton("")
	if err != nil {
		t.Error(err)
	}

	if tc := sk.Feedback[0].IndexTestcase(15); tc.InputPath != "tests/15.in" {
		t.Error(tc)
	}

	problemYAML = `{
		"shortname": "covering",
		"titles": [
			{"title": "T-lefedés", "language": "hungarian"}
		],
		"statements": [
			{"path": "covering-HUN.pdf", "language": "hungarian", "type": "application/pdf"}
		],
		"attachments": [
			{"name": "minta.zip", "path": "minta.zip"}
		],
		"tests": {
			"feedback_type": "ioi",
			"time_limit": 1.0,
			"memory_limit": 100,
			"input_pattern": "%d.in",
			"output_pattern": "%d.out",
			"test_count": 100
		}
	}`

	_ = afero.WriteFile(fs, "problem.yaml", []byte(problemYAML), 0777)
	_ = afero.WriteFile(fs, "minta.zip", []byte(""), 0777)
	p, err = parser(fs, "./")
	if err != nil {
		t.Fatal(err)
	}

	sk, err = p.StatusSkeleton("")
	if err != nil {
		t.Error(err)
	}

	if tc := sk.Feedback[0].IndexTestcase(56); tc.InputPath != "tests/56.in" {
		t.Error(tc)
	}

	assert.Equal(t, sk.Feedback[0].Testcases()[0].TimeLimit, 1000*time.Millisecond)
	assert.Equal(t, sk.Feedback[0].Testcases()[0].MemoryLimit, 100*memory.MiB)
	assert.Equal(t, sk.Feedback[0].Testcases()[0].Group, "base")
	assert.Equal(t, sk.Feedback[0].Testcases()[0].InputPath, "tests/1.in")
	assert.Equal(t, sk.Feedback[0].Testcases()[0].AnswerPath, "tests/1.out")
	assert.Equal(t, sk.Feedback[0].Testcases()[0].Index, 1)
	assert.Equal(t, sk.Feedback[0].Testcases()[0].MaxScore, 0.0)
	assert.Equal(t, sk.Feedback[0].Testcases()[0].Testset, "")
}
