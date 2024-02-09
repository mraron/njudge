package cpp_test

import (
	context2 "golang.org/x/net/context"
	"testing"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/spf13/afero"
)

func TestAutoCompile(t *testing.T) {
	fs := afero.NewMemMapFs()
	s, _ := sandbox.NewDummy()

	_ = afero.WriteFile(fs, "main.cpp", []byte(`#include "teszt.h"
int main() {x*=2;}`), 0644)
	_ = afero.WriteFile(fs, "main_syntaxerror.cpp", []byte(`#include "teszt.h"
int main() {x*****asdasd=2;}`), 0644)

	_ = fs.MkdirAll("headers", 0755)
	_ = afero.WriteFile(fs, "headers/teszt.h", []byte(`#ifndef _TESZT_H
#define _TESZT_H
int x = 11;
#endif //_TESZT_H`), 0644)

	if err := cpp.AutoCompile(context2.TODO(), fs, s, "./headers", "main.cpp", "main"); err != nil {
		t.Error(err)
	}

	if fileInfo, err := fs.Stat("main"); err != nil || fileInfo.Mode()&0100 != 0100 || fileInfo.Size() == 0 {
		t.Error("no execbit or wrong size")
	}

	_ = afero.WriteFile(fs, "main", []byte(""), 0644)
	if err := cpp.AutoCompile(context2.TODO(), fs, s, "./headers", "main.cpp", "main"); err != nil {
		t.Error(err)
	}

	if fileInfo, err := fs.Stat("main"); err != nil || fileInfo.Mode()&0100 != 0100 || fileInfo.Size() == 0 {
		t.Error("no execbit or wrong size")
	}

	_ = afero.WriteFile(fs, "main_not_executable", []byte("has"), 0766)
	if err := cpp.AutoCompile(context2.TODO(), fs, s, "./headers", "main_syntaxerror.cpp", "main_not_executable"); err != nil {
		t.Error(err)
	}

	_ = fs.Chmod("main_not_executable", 0666)
	if err := cpp.AutoCompile(context2.TODO(), fs, s, "./headers", "main_syntaxerror.cpp", "main_not_executable"); err == nil {
		t.Error("No error?")
	}

	_ = afero.WriteFile(fs, "main", []byte(""), 0777)
	if err := cpp.AutoCompile(context2.TODO(), fs, s, "./headers", "main_syntaxerror.cpp", "main"); err == nil {
		t.Error("Compiled fine???")
	}
}
