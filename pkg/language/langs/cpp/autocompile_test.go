package cpp_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/spf13/afero"
)

func TestAutoCompile(t *testing.T) {
	fs := afero.NewMemMapFs()
	s, err := sandbox.NewDummy()
	assert.Nil(t, err)

	_ = afero.WriteFile(fs, "main.cpp", []byte(`#include "teszt.h"
int main() {x*=2;}`), 0644)
	_ = afero.WriteFile(fs, "main_syntaxerror.cpp", []byte(`#include "teszt.h"
int main() {x*****asdasd=2;}`), 0644)

	_ = fs.MkdirAll("headers", 0755)
	_ = afero.WriteFile(fs, "headers/teszt.h", []byte(`#ifndef _TESZT_H
#define _TESZT_H
int x = 11;
#endif //_TESZT_H`), 0644)

	if err := cpp.AutoCompile(context.TODO(), fs, s, "./headers", "main.cpp", "main"); err != nil {
		t.Error(err)
	}

	var fileInfo os.FileInfo
	fileInfo, err = fs.Stat("main")
	assert.Nil(t, err)
	assert.Equal(t, os.FileMode(0100), fileInfo.Mode()&0100, "no execbit")
	assert.NotEqual(t, 0, fileInfo.Size(), "wrong size")

	_ = afero.WriteFile(fs, "main", []byte(""), 0644)
	if err := cpp.AutoCompile(context.TODO(), fs, s, "./headers", "main.cpp", "main"); err != nil {
		t.Error(err)
	}

	fileInfo, err = fs.Stat("main")
	assert.Nil(t, err)
	assert.Equal(t, os.FileMode(0100), fileInfo.Mode()&0100, "no execbit")
	assert.NotEqual(t, 0, fileInfo.Size(), "wrong size")

	_ = afero.WriteFile(fs, "main_not_executable", []byte("has"), 0766)
	if err := cpp.AutoCompile(context.TODO(), fs, s, "./headers", "main_syntaxerror.cpp", "main_not_executable"); err != nil {
		t.Error(err)
	}

	_ = fs.Chmod("main_not_executable", 0666)
	if err := cpp.AutoCompile(context.TODO(), fs, s, "./headers", "main_syntaxerror.cpp", "main_not_executable"); err == nil {
		t.Error("No error?")
	}

	_ = afero.WriteFile(fs, "main", []byte(""), 0777)
	if err := cpp.AutoCompile(context.TODO(), fs, s, "./headers", "main_syntaxerror.cpp", "main"); err == nil {
		t.Error("Compiled fine???")
	}
}
