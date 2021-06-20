package polygon

import (
	"github.com/mraron/njudge/utils/problems"
	"github.com/spf13/afero"
	"log"
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

func TestTestDataParsing(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./tests/")

	bs, err := afero.ReadFile(fs, "problemxml/problem.xml")
	log.Print(string(bs), err)

	cs := problems.NewConfigStore()
	parser, identifier := ParserAndIdentifier(CompileBinaries(false), UseFS(fs))

	cs.Register("polygon", parser, identifier)

	if p, err := cs.Parse("problemxml/"); err != nil {
		t.Errorf("got error: %v", err)
	}else {
		if p.Titles()[0].String() != "Fancy Fence" {
			t.Errorf("Wanted %q got %q", "Fancy Fence", p.Titles()[0].String())
		}
		if p.Name() != "fancyfence" {
			t.Errorf("Wanted %q got %q", "fancyfence", p.Name())
		}

		log.Print(p.(Problem))
	}
}
