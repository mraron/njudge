package stub

import (
	"bytes"
	"fmt"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
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

func (s Stub) Compile(jinfo problems.Judgeable, sbox sandbox.Sandbox, lang language.Language, src io.Reader, dest io.Writer) (io.Reader, error) {
	lst, found := jinfo.Languages(), false

	for _, l := range lst {
		if l.ID() == lang.ID() {
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("%s evaluation: language %s is not supported", s.Name(), lang.ID())
	}

	files := jinfo.Files()
	needed := make([]problems.File, 0)
	for _, f := range files {
		if f.Role == "stub_"+lang.ID() || (strings.HasPrefix(lang.ID(), "cpp") && f.Role == "stub_cpp") {
			needed = append(needed, f)
		}
	}

	language_files := make([]sandbox.File, 0, len(needed))
	for _, n := range needed {
		conts, err := ioutil.ReadFile(n.Path)
		if err != nil {
			return nil, err
		}

		language_files = append(language_files, sandbox.File{Name: n.Name, Source: bytes.NewBuffer(conts)})
	}

	buf := &bytes.Buffer{}

	_, err := lang.Compile(sbox, sandbox.File{Name: "main", Source: src}, dest, nil)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
