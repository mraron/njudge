package sandbox_test

import (
	"flag"
	_ "github.com/mraron/njudge/internal/web"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
)

var isolateInstalled = flag.Bool("isolate", false, "run isolate tests")
var allLanguages = flag.Bool("all_languages", false, "run all languaegs")

func TestIsolateWithCpp17(t *testing.T) {
	if !*isolateInstalled {
		t.Skip("-isolate is not set	")
	}

	s := sandbox.NewIsolate(555)
	if err := cpp.Std17.Test(s); err != nil {
		t.Error(err)
	}
}

func TestIsolateWithAllLanguages(t *testing.T) {
	if !*isolateInstalled {
		t.Skip("-isolate is not set")
	}
	if !*allLanguages {
		t.Skip("-all_languages is not set")
	}

	s := sandbox.NewIsolate(556)
	for _, lang := range language.DefaultStore.List() {
		t.Logf("Running %s", lang.Id())
		if err := lang.Test(s); err != nil {
			t.Error(err)
		}
	}
}
