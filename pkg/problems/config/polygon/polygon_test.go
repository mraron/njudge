package polygon

import (
	"errors"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
)

func TestDummyParsing(t *testing.T) {
	memFs := afero.NewMemMapFs()

	cs := problems.NewConfigList()
	parser := Parser{
		CompileBinaries: true,
		EnableHTMLGen:   true,
	}

	assert.NoError(t, cs.Register("polygon", parser.Parse, parser.Identifier))
	assert.NoError(t, memFs.Mkdir("testproblem", 0700))

	if p, err := cs.Parse(memFs, "testproblem/"); p != nil || !errors.Is(err, problems.ErrorNoMatch) {
		t.Error("no problem.xml but found it")
	}

	f, _ := memFs.Create("testproblem/problem.xml")
	assert.NoError(t, f.Close())

	if p, _ := cs.Parse(memFs, "testproblem/"); p != nil {
		t.Error("empty problem.xml")
	}
}

func testProblemXML(t *testing.T, p Problem) {
	assert.Equal(t, "Fancy Fence", p.Titles()[0].String())
	assert.Equal(t, "english", p.Titles()[0].Locale())
	assert.Equal(t, "fancyfence", p.Name())
	assert.Equal(t, "problemxml/", p.Path)
	assert.Equal(t, batch.Name, p.TaskType)
	assert.Equal(t, problems.FeedbackIOI, problems.FeedbackTypeFromShortString(p.FeedbackType))

	if len(p.Judging.Testsets) != 1 {
		t.Error("wrong number of testsets")
	}

	assert.Empty(t, p.Judging.InputFile)
	assert.Empty(t, p.Judging.OutputFile)

	_, err := p.Judging.GetTestset("asd")
	if err == nil {
		t.Error("found testset \"asd\"")
	}

	ts, err := p.Judging.GetTestset("tests")
	if ts != nil {
		assert.Equal(t, "tests", ts.Name)
		assert.Equal(t, 1000, ts.TimeLimit)

		// default polygon pattern
		assert.Equal(t, "tests/%02d", ts.InputPathPattern)
		assert.Equal(t, "tests/%02d.a", ts.AnswerPathPattern)

		assert.Equal(t, len(ts.Tests), ts.TestCount)

		groupCnt := make(map[string]int)
		for _, test := range ts.Tests {
			groupCnt[test.Group]++
		}

		assert.Equal(t, len(ts.Groups), len(groupCnt))

		for k := range groupCnt {
			g, err := ts.Group(k)
			if err == nil {
				assert.Equal(t, g.Name, k)
			} else {
				t.Errorf("can't find group %s, %v", k, err)
			}
		}

		skeleton, err := p.StatusSkeleton("tests")
		if err == nil {
			assert.Equal(t, problems.FeedbackIOI, skeleton.FeedbackType)
			assert.InDelta(t, float64(100), skeleton.Feedback[0].MaxScore(), 0.001)
			assert.Equal(t, "tests", skeleton.Feedback[0].Name)
		} else {
			t.Error(err)
		}
	} else {
		t.Error("\"tests\" testset not found", err)
	}
}

func TestFSParsing(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata/")
	cs := problems.NewConfigList()
	parser := Parser{
		CompileBinaries: false,
		EnableHTMLGen:   true,
	}
	if err := cs.Register("polygon", parser.Parse, parser.Identifier); err != nil {
		t.Error(err)
	}

	if p, err := cs.Parse(fs, "problemxml/"); err != nil {
		t.Fatalf("got error: %v", err)
	} else {
		t.Run("problemXML", func(t *testing.T) {
			testProblemXML(t, p.(Problem))
		})
	}
}

func TestJSONStatement(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata/")

	stmt, _ := ParseJSONStatement(fs, "hablaty/", "hungarian")
	if stmt != nil {
		t.Error("found problem-properties.json")
	}

	_, err := ParseJSONStatement(fs, "json_statement/", "hungarian")
	if err != nil {
		t.Error(err)
	}
}
