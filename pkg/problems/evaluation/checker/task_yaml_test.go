package checker

import (
	"errors"
	"github.com/mraron/njudge/pkg/language/runner"
	"github.com/mraron/njudge/pkg/problems"
	"io"
	"testing"
)

func TestTaskYAML_Check(t1 *testing.T) {
	type fields struct {
		path       string
		executable runner.Executable
	}

	tests := []struct {
		name              string
		fields            fields
		tc                *problems.Testcase
		wantVerdictName   problems.VerdictName
		wantScore         float64
		wantErr           bool
		wantCheckerOutput string
	}{
		{
			name: "wrong answer",
			fields: fields{
				path: "",
				executable: runner.NewFunction(func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
					_, err := stdout.Write([]byte("0.0"))
					return 0, err
				}),
			},
			tc: &problems.Testcase{
				Score:       0.0,
				MaxScore:    15.0,
				VerdictName: problems.VerdictDR,
			},
			wantVerdictName: problems.VerdictWA,
			wantScore:       0.0,
		},
		{
			name: "partially correct",
			fields: fields{
				path: "",
				executable: runner.NewFunction(func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
					_, err := stdout.Write([]byte("0.5"))
					return 0, err
				}),
			},
			tc: &problems.Testcase{
				Score:       0.0,
				MaxScore:    15.0,
				VerdictName: problems.VerdictDR,
			},
			wantVerdictName: problems.VerdictPC,
			wantScore:       7.5,
		},
		{
			name: "accepted",
			fields: fields{
				path: "",
				executable: runner.NewFunction(func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
					_, err := stdout.Write([]byte("1.0"))
					return 0, err
				}),
			},
			tc: &problems.Testcase{
				Score:       0.0,
				MaxScore:    15.0,
				VerdictName: problems.VerdictDR,
			},
			wantVerdictName: problems.VerdictAC,
			wantScore:       15,
		},
		{
			name: "executable crash",
			fields: fields{
				path: "",
				executable: runner.NewFunction(func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
					return 0, errors.New("")
				}),
			},
			tc: &problems.Testcase{
				Score:       0.0,
				MaxScore:    15.0,
				VerdictName: problems.VerdictDR,
			},
			wantVerdictName: problems.VerdictDR,
			wantErr:         true,
		},
		{
			name: "wrong format",
			fields: fields{
				path: "",
				executable: runner.NewFunction(func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
					_, err := stdout.Write([]byte("xxx"))
					return 0, err
				}),
			},
			tc: &problems.Testcase{
				Score:       0.0,
				MaxScore:    100.0,
				VerdictName: problems.VerdictDR,
			},
			wantVerdictName: problems.VerdictDR,
			wantErr:         true,
		},
		{
			name: "checker output",
			fields: fields{
				path: "",
				executable: runner.NewFunction(func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
					_, err := stdout.Write([]byte("1.0"))
					_, err2 := stderr.Write([]byte("checker output"))
					return 0, errors.Join(err, err2)
				}),
			},
			tc: &problems.Testcase{
				Score:       0.0,
				MaxScore:    100.0,
				VerdictName: problems.VerdictDR,
			},
			wantVerdictName:   problems.VerdictAC,
			wantScore:         100,
			wantCheckerOutput: "checker output",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaskYAML{
				path:       tt.fields.path,
				executable: tt.fields.executable,
			}
			if err := t.Check(nil, tt.tc); (err != nil) != tt.wantErr {
				t1.Errorf("task_yaml.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.tc.VerdictName != tt.wantVerdictName {
				t1.Errorf("task_yaml.Check() verdictName %v != %v", tt.tc.VerdictName, tt.wantVerdictName)
			}
			if tt.tc.Score != tt.wantScore {
				t1.Errorf("task_yaml.Check() score %v != %v", tt.tc.VerdictName, tt.wantScore)
			}
			if tt.tc.CheckerOutput != tt.wantCheckerOutput {
				t1.Errorf("task_yaml.Check() checkerOutput %v != %v", tt.tc.CheckerOutput, tt.wantCheckerOutput)
			}
		})
	}
}
