package language_test

import (
	"flag"
	"github.com/mraron/njudge/pkg/language"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	_ "github.com/mraron/njudge/pkg/language/langs/csharp"
	_ "github.com/mraron/njudge/pkg/language/langs/cython3"
	_ "github.com/mraron/njudge/pkg/language/langs/golang"
	_ "github.com/mraron/njudge/pkg/language/langs/java"
	_ "github.com/mraron/njudge/pkg/language/langs/julia"
	_ "github.com/mraron/njudge/pkg/language/langs/nim"
	_ "github.com/mraron/njudge/pkg/language/langs/pascal"
	_ "github.com/mraron/njudge/pkg/language/langs/pypy3"
	_ "github.com/mraron/njudge/pkg/language/langs/python3"
	_ "github.com/mraron/njudge/pkg/language/langs/zip"

	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
)

var useIsolate = flag.Bool("isolate", false, "run isolate integration tests")
var testAllLanguages = flag.Bool("all_languages", false, "run tests for all languages")

func TestIsolateWithCpp17(t *testing.T) {
	if !*useIsolate {
		t.Skip("-isolate is not set")
	}
	if *testAllLanguages {
		t.Skip("running all languages instead (which includes cpp17)")
	}

	s, err := sandbox.NewIsolate(555)
	if err != nil {
		t.Error(err)
	}
	if err := cpp.Std17.Test(t, s); err != nil {
		t.Error(err)
	}
}

func TestIsolateWithAllLanguages(t *testing.T) {
	if !*useIsolate {
		t.Skip("-isolate is not set")
	}
	if !*testAllLanguages {
		t.Skip("-all_languages is not set")
	}

	s, err := sandbox.NewIsolate(556)
	if err != nil {
		t.Error(err)
	}
	for _, lang := range language.DefaultStore.List() {
		t.Logf("Running %s", lang.ID())
		l := lang.(language.Wrapper).Language
		if _, ok := l.(language.Testable); ok {
			if err := l.(language.Testable).Test(t, s); err != nil {
				t.Error(err)
			}
		}
	}
}
