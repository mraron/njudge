package checker

import (
	"golang.org/x/net/context"
	"io"
	"os"
	"strings"

	"github.com/karrick/gobls"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/multierr"
)

// Whitediff is the [default checker] built into CMS
//
// [default checker]: https://cms.readthedocs.io/en/v1.4/Task%20types.html#tasktypes-white-diff
type Whitediff struct{}

func (Whitediff) Name() string {
	return "whitediff"
}

func (Whitediff) Check(ctx context.Context, testcase *problems.Testcase) error {
	tc := testcase

	ans, err := os.Open(tc.AnswerPath)
	if err != nil {
		return multierr.Combine(err, ans.Close())
	}
	defer ans.Close()

	out, err := os.Open(tc.OutputPath)
	if err != nil {
		return multierr.Combine(err, out.Close())
	}
	defer out.Close()

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
		return 1, multierr.Combine(x.Err(), y.Err())
	}

	return 0, multierr.Combine(x.Err(), y.Err())
}
