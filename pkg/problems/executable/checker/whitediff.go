package checker

import (
	"context"
	"errors"
	"github.com/spf13/afero"
	"io"
	"strings"

	"github.com/karrick/gobls"
	"github.com/mraron/njudge/pkg/problems"
)

// Whitediff is the [default checker] built into CMS
//
// [default checker]: https://cms.readthedocs.io/en/v1.4/Task%20types.html#tasktypes-white-diff
type Whitediff struct {
	answerFs afero.Fs
	outputFs afero.Fs
}

type WhitediffOption func(*Whitediff)

func WhiteDiffWithFs(answerFs, outputFs afero.Fs) WhitediffOption {
	return func(whitediff *Whitediff) {
		whitediff.answerFs = answerFs
		whitediff.outputFs = outputFs
	}
}

func NewWhitediff(opts ...WhitediffOption) Whitediff {
	res := Whitediff{
		answerFs: afero.NewOsFs(),
		outputFs: afero.NewOsFs(),
	}
	for _, opt := range opts {
		opt(&res)
	}

	return res
}

func (Whitediff) Name() string {
	return "whitediff"
}

func (w Whitediff) Check(ctx context.Context, testcase *problems.Testcase) error {
	tc := testcase

	ans, err := w.answerFs.Open(tc.AnswerPath)
	if err != nil {
		return errors.Join(err, ans.Close())
	}
	defer func(ans afero.File) {
		_ = ans.Close()
	}(ans)

	out, err := w.outputFs.Open(tc.OutputPath)
	if err != nil {
		return errors.Join(err, out.Close())
	}
	defer func(out afero.File) {
		_ = out.Close()
	}(out)

	outcome, err := DoWhitediff(ans, out)
	tc.Score = outcome * tc.MaxScore

	if outcome == 1.0 {
		tc.VerdictName = problems.VerdictAC
	} else {
		tc.VerdictName = problems.VerdictWA
	}

	return err
}

var whitespaces = []byte{' ', '\t', '\n', '\x0b', '\x0c', '\r'}

func canonicalize(str string) string {
	for _, w := range whitespaces {
		str = strings.ReplaceAll(str, string(w), " ")
	}

	arr := []string{}
	for _, tok := range strings.Split(str, " ") {
		if len(tok) > 0 {
			arr = append(arr, tok)
		}
	}

	return strings.Join(arr, " ")
}

// DoWhitediff performs a [whitediff] comparision on the two input readers.
// It returns:
//
//   - "1.0" if the streams match,
//   - "0.0" if not,
//
// and any errors encountered while trying to exhaust the readers.
//
// [whitediff]: https://cms.readthedocs.io/en/v1.4/Task%20types.html#tasktypes-white-diff
func DoWhitediff(a io.Reader, b io.Reader) (float64, error) {
	x, y := gobls.NewScanner(a), gobls.NewScanner(b)
	eq := true
	for {
		hasA, hasB := x.Scan(), y.Scan()
		if hasA || hasB {
			a, b := string(x.Bytes()), string(y.Bytes())
			if len(a) == 0 || len(b) == 0 {
				a = strings.Trim(a, string(whitespaces))
				b = strings.Trim(b, string(whitespaces))
				if len(a) > 0 || len(b) > 0 {
					eq = false
					break
				}
			} else {
				a, b = canonicalize(a), canonicalize(b)
				if a != b {
					eq = false
					break
				}
			}
		} else {
			break
		}
	}

	if eq {
		return 1, errors.Join(x.Err(), y.Err())
	}

	return 0, errors.Join(x.Err(), y.Err())
}
