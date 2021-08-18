package polygon

import (
	"github.com/mraron/njudge/utils/problems"
	"github.com/spf13/afero"
	"testing"
)

func TestDummyParsing(t *testing.T) {
	memFs := afero.NewMemMapFs()

	cs := problems.NewConfigStore()
	parser, identifier := ParserAndIdentifier(UseFS(memFs))

	cs.Register("polygon", parser, identifier)

	memFs.Mkdir("testproblem", 0700)
	if p, err := cs.Parse("testproblem/"); p != nil || err != problems.ErrorNoMatch {
		t.Error("no problem.xml but found it")
	}

	f, _ := memFs.Create("testproblem/problem.xml")
	f.Close()
	if p, _ := cs.Parse("testproblem/"); p != nil {
		t.Error("empty problem.xml")
	}

}

func TestParsing(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./tests/")
	cs := problems.NewConfigStore()
	parser, identifier := ParserAndIdentifier(CompileBinaries(false), UseFS(fs))

	cs.Register("polygon", parser, identifier)

	polygonTestcases := []func(*testing.T, Problem) {
		func(t *testing.T, p Problem) {
			assert := func(name string, got, want interface{}) {
				if got != want {
					t.Errorf("wrong %s, want %v got %v", name, want, got)
				}
			}

			assert("title", p.Titles()[0].String(), "Fancy Fence")
			assert("title lang", p.Titles()[0].Locale, "english")
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
					}else {
						t.Errorf("can't find group %s, %v", k, err)
					}
				}

				skeleton, err := p.StatusSkeleton("tests")
				if err == nil {
					assert("sk_feedbacktype", skeleton.FeedbackType, problems.FEEDBACK_IOI)
					assert("sk_maxscore", skeleton.Feedback[0].MaxScore(), float64(100))
					assert("sk_testsetname", skeleton.Feedback[0].Name, "tests")
				}else {
					t.Error(err)
				}
			}else {
				t.Error("\"tests\" testset not found")
			}
		},
	}

	if p, err := cs.Parse("problemxml/"); err != nil {
		t.Fatalf("got error: %v", err)
	}else {
		for _, tc := range polygonTestcases {
			tc(t, p.(Problem))
		}
	}
}

func TestJSONStatement(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./tests/")

	_, err := ParseJSONStatement(fs, "hablaty/")
	if err == nil {
		t.Error("found problem-properties.json")
	}

	_, err = ParseJSONStatement(fs, "json_statement/")
	if err != nil {
		t.Error(err)
	}
}