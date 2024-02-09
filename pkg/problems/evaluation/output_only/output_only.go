package output_only

import (
	"archive/zip"
	"bytes"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"go.uber.org/multierr"
)

type OutputOnly struct {
	batch.Batch
}

func New() OutputOnly {
	oo := OutputOnly{Batch: batch.New()}
	oo.Batch.InitF = func(rc *batch.RunContext) error {
		readerAt := bytes.NewReader(rc.Binary)

		zip, err := zip.NewReader(readerAt, int64(len(rc.Binary)))
		if err != nil {
			return err
		}

		rc.Store["zip"] = zip
		return nil
	}

	oo.Batch.RunF = func(rc *batch.RunContext, g *problems.Group, t *problems.Testcase) (sandbox.Status, error) {
		var (
			f     io.ReadCloser
			err   error
			conts []byte
		)

		for _, file := range rc.Store["zip"].(*zip.Reader).File {
			if file.Name == filepath.Base(t.OutputPath) {
				f, err = file.Open()
				if err != nil {
					break
				}
				defer f.Close()

				conts, err = io.ReadAll(f)
				if err != nil {
					err = multierr.Combine(err, f.Close())
					break
				}

				rc.Stdout = &bytes.Buffer{}
				rc.Stdout.Write(conts)

				return sandbox.Status{Verdict: sandbox.VerdictOK}, nil
			}
		}

		return sandbox.Status{Verdict: sandbox.VerdictRE}, err
	}

	oo.Batch.CheckFailF = func(rc *batch.RunContext, s sandbox.Status, g *problems.Group, t *problems.Testcase) error {
		t.VerdictName = problems.VerdictPE
		return nil
	}

	return oo
}

func (o OutputOnly) Name() string {
	return "outputonly"
}

func (o OutputOnly) Compile(jinfo problems.Judgeable, s sandbox.Sandbox, l language.Language, src io.Reader, errw io.Writer) (io.Reader, error) {
	zipContents, err := ioutil.ReadAll(src)
	if err != nil {
		errw.Write([]byte(err.Error()))
		return nil, err
	}

	readerAt := bytes.NewReader(zipContents)

	_, err = zip.NewReader(readerAt, int64(len(zipContents)))
	if err != nil {
		errw.Write([]byte(err.Error()))
		return nil, err
	}

	return bytes.NewReader(zipContents), nil
}
