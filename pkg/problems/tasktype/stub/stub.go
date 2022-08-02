package stub

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/tasktype/batch"
)

type Stub struct {
	batch.Batch
}

func New() Stub {
	return Stub{batch.New()}
}

func (s Stub) Name() string {
	return "stub"
}

func (s Stub) Compile(jinfo problems.Judgeable, sandbox language.Sandbox, lang language.Language, src io.Reader, dest io.Writer) (io.Reader, error) {
	lst, found := jinfo.Languages(), false

	for _, l := range lst {
		if l.Name() == lang.Name() {
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("%s tasktype: language %s is not supported", s.Name(), lang.Name())
	}

	files := jinfo.Files()
	needed := make([]problems.File, 0)
	for _, f := range files {
		if f.Role == "stub_"+lang.Id() || (strings.HasPrefix(lang.Id(), "cpp") && f.Role == "stub_cpp") {
			needed = append(needed, f)
		}
	}

	language_files := make([]language.File, 0, len(needed))
	for _, n := range needed {
		conts, err := ioutil.ReadFile(n.Path)
		if err != nil {
			return nil, err
		}

		language_files = append(language_files, language.File{Name: n.Name, Source: bytes.NewBuffer(conts)})
	}

	buf := &bytes.Buffer{}

	err := lang.Compile(sandbox, language.File{Name: "main", Source: src}, buf, dest, language_files)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func init() {
	problems.RegisterTaskType(New())
}
