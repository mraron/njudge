package feladat_txt_test

import (
	"path/filepath"
	"testing"

	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"github.com/spf13/afero"
)

func TestParsing(t *testing.T) {
	tests := []struct {
		path       string
		name       string
		configFile string
		check      func(t *testing.T, p problems.Problem, err error)
	}{
		{
			"baratnok/",
			"simple",
			`Barátnők (50 pont);
NO;
29;64000;1.0;1;
0;0;2;2;2;2;2;1;2;2;2;2;2;2;2;1;1;1;2;2;2;2;2;2;2;2;2;2;2;`,
			func(t *testing.T, p problems.Problem, err error) {
				if err != nil {
					t.Error(err)
					return
				}

				if p.Name() != "baratnok" {
					t.Error(p.Name())
				}

				if p.Titles()[0].String() != "Barátnők (50 pont)" {
					t.Error(p.Titles())
				}

				if len(p.Attachments()) != 1 {
					t.Error(p.Attachments())
				}

				if len(p.Statements()) != 1 {
					t.Error(p.Statements())
				}

				if p.MemoryLimit() != 64000*1024 {
					t.Error("wrong memory limit:", p.MemoryLimit())
				}

				if p.TimeLimit() != 1000 {
					t.Error("wrong time limit")
				}

				if i, o := p.InputOutputFiles(); i != "" || o != "" {
					t.Error("output files")
				}

				if p.GetTaskType().Name() != "batch" {
					t.Error(p.GetTaskType())
				}

				sk, err := p.StatusSkeleton("")
				if err != nil {
					t.Error(err)
				}

				if sk.FeedbackType != problems.FeedbackIOI {
					t.Error("Wrong feedback type")
				}
				if sk.Feedback[0].MaxScore() != 50.0 {
					t.Error("Wrong max score")
				}
			},
		},
		{
			"mobilnet/",
			"subtest",
			`MobilNet (50 pont);
NO;
22;32000;1.0;2;
0;0;1;1;1;1;1;1;1;1;1;1;2;2;2;2;2;2;2;2;2;2;
0;0;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;`,
			func(t *testing.T, p problems.Problem, err error) {
				if err != nil {
					t.Error(err)
					return
				}

				sk, err := p.StatusSkeleton("")
				if err != nil {
					t.Error(err)
				}

				if tc := sk.Feedback[0].IndexTestcase(22); tc.MaxScore != 3 {
					t.Error(tc)
				}
			},
		},
		{
			"abc/",
			"module",
			`abc;
ayaya;
1;1;0.1;1;
0;
0;`,
			func(t *testing.T, p problems.Problem, err error) {
				if err == nil {
					t.Error("can parse module")
					return
				}
			},
		},
	}

	for ind := range tests {
		t.Run(tests[ind].name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			afero.WriteFile(fs, filepath.Join(tests[ind].path, "feladat.txt"), []byte(tests[ind].configFile), 0777)

			afero.WriteFile(fs, filepath.Join(tests[ind].path, "feladat.pdf"), []byte(""), 0777)
			afero.WriteFile(fs, filepath.Join(tests[ind].path, "minta.zip"), []byte(""), 0777)
			afero.WriteFile(fs, filepath.Join(tests[ind].path, "ellen.cpp"), []byte("main(){}"), 0777)

			p, err := feladat_txt.Parse(fs, tests[ind].path)
			tests[ind].check(t, p, err)
		})
	}

	fs := afero.NewMemMapFs()
	p, err := feladat_txt.Parse(fs, "./")
	if err == nil || p != nil {
		t.Error("Can parse non-existing feladat.txt")
	}
}
