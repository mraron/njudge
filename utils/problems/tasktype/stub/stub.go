package stub

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/utils/problems/tasktype/batch"
	"io"
	"io/ioutil"
)

type Stub struct {
}

func (s Stub) Name() string {
	return "stub"
}

func (s Stub) Compile(jinfo problems.JudgingInformation, sandbox language.Sandbox, lang language.Language, src io.Reader, dest io.Writer) (io.Reader, error) {
	lst, found := jinfo.Languages(), false

	for _, l := range lst {
		if l.Name() == lang.Name() {
			found = true
		}
	}

	if !found {
		return nil, errors.New(fmt.Sprintf("running problem %s on %s tasktype, language %s is not supported", jinfo.Name(), s.Name(), lang.Name()))
	}

	files := jinfo.Files()
	needed := make([]problems.File, 0)
	for _, f := range files {
		if f.Role == "stub_"+lang.Id() {
			needed = append(needed, f)
		}
	}

	language_files := make([]language.File, 0, len(needed))
	for _, n := range needed {
		conts, err := ioutil.ReadFile(n.Path)
		if err != nil {
			return nil, err
		}

		language_files = append(language_files, language.File{n.Name, bytes.NewBuffer(conts)})
	}

	buf := &bytes.Buffer{}

	err := lang.Compile(sandbox, language.File{"main", src}, buf, dest, language_files)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (Stub) Run(jinfo problems.JudgingInformation, sp *language.SandboxProvider, lang language.Language, bin io.Reader, testNotifier chan string, statusNotifier chan problems.Status) (problems.Status, error) {
	return batch.Batch{}.Run(jinfo, sp, lang, bin, testNotifier, statusNotifier)
}

func init() {
	problems.RegisterTaskType(Stub{})
}
