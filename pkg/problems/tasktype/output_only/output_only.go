package output_only

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/tasktype/batch"
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

	oo.Batch.RunF = func(rc *batch.RunContext, g *problems.Group, t *problems.Testcase) (language.Status, error) {
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

				conts, err = ioutil.ReadAll(f)
				if err != nil {
					err = multierr.Combine(err, f.Close())
					break
				}

				rc.Stdout = &bytes.Buffer{}
				rc.Stdout.Write(conts)

				return language.Status{Verdict: language.VerdictOK}, nil
			}
		}

		return language.Status{Verdict: language.VerdictRE}, err
	}

	oo.Batch.CheckFailF = func(rc *batch.RunContext, s language.Status, g *problems.Group, t *problems.Testcase) error {
		t.VerdictName = problems.VerdictDR
		return nil
	}

	return oo
}

func (o OutputOnly) Name() string {
	return "outputonly"
}

func (o OutputOnly) Compile(jinfo problems.Judgeable, s language.Sandbox, l language.Language, src io.Reader, errw io.Writer) (io.Reader, error) {
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

func init() {
	problems.RegisterTaskType(New())
}
