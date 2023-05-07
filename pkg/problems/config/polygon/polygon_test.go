package polygon

import (
	"testing"

	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
)

func TestDummyParsing(t *testing.T) {
	handleError := func(err error) {
		if err != nil {
			t.Error("error occurred", err)
		}
	}

	memFs := afero.NewMemMapFs()

	cs := problems.NewConfigStore()
	parser, identifier := ParserAndIdentifier()

	handleError(cs.Register("polygon", parser, identifier))
	handleError(memFs.Mkdir("testproblem", 0700))

	if p, err := cs.Parse(memFs, "testproblem/"); p != nil || err != problems.ErrorNoMatch {
		t.Error("no problem.xml but found it")
	}

	f, _ := memFs.Create("testproblem/problem.xml")
	handleError(f.Close())

	if p, _ := cs.Parse(memFs, "testproblem/"); p != nil {
		t.Error("empty problem.xml")
	}
}

func testProblemXML(t *testing.T, p Problem) {
	assert := func(name string, got, want interface{}) {
		if got != want {
			t.Errorf("wrong %s, want %v got %v", name, want, got)
		}
	}

	assert("title", p.Titles()[0].String(), "Fancy Fence")
	assert("title lang", p.Titles()[0].Locale(), "english")
	assert("name", p.Name(), "fancyfence")
	assert("path", p.Path, "problemxml/")
	assert("tasktype", p.TaskType, "batch")
	assert("feedbacktype", p.FeedbackType, "ioi")

	if len(p.Judging.Testsets) != 1 {
		t.Error("wrong number of testsets")
	}

	assert("inputfile", p.Judging.InputFile, "")
	assert("outputfile", p.Judging.OutputFile, "")

	_, err := p.Judging.GetTestset("asd")
	if err == nil {
		t.Error("found testset \"asd\"")
	}

	ts, err := p.Judging.GetTestset("tests")
	if ts != nil {
		assert("testset_name", ts.Name, "tests")
		assert("testset_timelimit", ts.TimeLimit, 1000)

		// default polygon pattern
		assert("testset_inputpattern", ts.InputPathPattern, "tests/%02d")
		assert("testset_answerpattern", ts.AnswerPathPattern, "tests/%02d.a")

		assert("testset_count", ts.TestCount, len(ts.Tests))

		groupCnt := make(map[string]int)
		for _, test := range ts.Tests {
			groupCnt[test.Group]++
		}

		assert("number of groups", len(groupCnt), len(ts.Groups))

		for k := range groupCnt {
			g, err := ts.Group(k)
			if err == nil {
				assert(k, k, g.Name)
			} else {
				t.Errorf("can't find group %s, %v", k, err)
			}
		}

		skeleton, err := p.StatusSkeleton("tests")
		if err == nil {
			assert("sk_feedbacktype", skeleton.FeedbackType, problems.FeedbackIOI)
			assert("sk_maxscore", skeleton.Feedback[0].MaxScore(), float64(100))
			assert("sk_testsetname", skeleton.Feedback[0].Name, "tests")
		} else {
			t.Error(err)
		}
	} else {
		t.Error("\"tests\" testset not found", err)
	}
}

func TestFSParsing(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata/")
	cs := problems.NewConfigStore()
	parser, identifier := ParserAndIdentifier(CompileBinaries(false))

	if err := cs.Register("polygon", parser, identifier); err != nil {
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

	stmt, _ := ParseJSONStatement(fs, "hablaty/")
	if stmt != nil {
		t.Error("found problem-properties.json")
	}

	_, err := ParseJSONStatement(fs, "json_statement/")
	if err != nil {
		t.Error(err)
	}
}

func TestCompileBinaries(t *testing.T) {
	fs := afero.NewCopyOnWriteFs(afero.NewBasePathFs(afero.NewOsFs(), "./testdata/"), afero.NewMemMapFs())

	if err := compileIfNotCompiled(fs, "", "check.cpp", "check"); err != nil {
		t.Error(err)
	}

	if err := compileIfNotCompiled(fs, "", "check_syntaxerr.cpp", "check_syntaxerr"); err == nil {
		t.Error("wanted error, but compiled fine")
	}

	// check already exists because of the first compilation
	if err := compileIfNotCompiled(fs, "", "check_syntaxerr.cpp", "check"); err != nil {
		t.Error(err)
	}

	f, err := fs.Create("check")
	if err != nil {
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}

	if err := compileIfNotCompiled(fs, "", "check_syntaxerr.cpp", "check"); err == nil {
		t.Error("should have compile error")
	}
}
