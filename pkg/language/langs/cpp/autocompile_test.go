package cpp_test

import (
	"testing"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/spf13/afero"
)

func TestAutoCompile(t *testing.T) {
	fs := afero.NewMemMapFs()
	s := sandbox.NewDummy()

	afero.WriteFile(fs, "main.cpp", []byte(`#include "teszt.h"
int main() {x*=2;}`), 0644)
	afero.WriteFile(fs, "main_syntaxerror.cpp", []byte(`#include "teszt.h"
int main() {x*****asdasd=2;}`), 0644)

	fs.MkdirAll("headers", 0755)
	afero.WriteFile(fs, "headers/teszt.h", []byte(`#ifndef _TESZT_H
#define _TESZT_H
int x = 11;
#endif //_TESZT_H`), 0644)

	if err := cpp.AutoCompile(fs, s, "./headers", "main.cpp", "main"); err != nil {
		t.Error(err)
	}

	if finfo, err := fs.Stat("main"); err != nil || finfo.Mode()&0100 != 0100 || finfo.Size() == 0 {
		t.Error("no execbit or wrong size")
	}

	afero.WriteFile(fs, "main", []byte(""), 0644)
	if err := cpp.AutoCompile(fs, s, "./headers", "main.cpp", "main"); err != nil {
		t.Error(err)
	}

	if finfo, err := fs.Stat("main"); err != nil || finfo.Mode()&0100 != 0100 || finfo.Size() == 0 {
		t.Error("no execbit or wrong size")
	}

	afero.WriteFile(fs, "main_not_executable", []byte("has"), 0766)
	if err := cpp.AutoCompile(fs, s, "./headers", "main_syntaxerror.cpp", "main_not_executable"); err != nil {
		t.Error(err)
	}
	fs.Chmod("main_not_executable", 0666)
	if err := cpp.AutoCompile(fs, s, "./headers", "main_syntaxerror.cpp", "main_not_executable"); err == nil {
		t.Error("No error?")
	}

	afero.WriteFile(fs, "main", []byte(""), 0777)
	if err := cpp.AutoCompile(fs, s, "./headers", "main_syntaxerror.cpp", "main"); err == nil {
		t.Error("Compiled fine???")
	}
}
