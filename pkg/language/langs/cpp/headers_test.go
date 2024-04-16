package cpp_test

import (
	"testing"

	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/spf13/afero"
)

func TestExtractHeaderNames(t *testing.T) {
	fs := afero.NewMemMapFs()

	src := []byte(`#include "abc.h"
#include<kaki>
#include<iostream>
#include <halo.h>`)
	_ = afero.WriteFile(fs, "main.cpp", src, 0644)

	_ = fs.MkdirAll("headers", 0755)
	_ = afero.WriteFile(fs, "headers/abc.h", []byte(""), 0644)
	_ = afero.WriteFile(fs, "headers/halo.h", []byte(""), 0644)

	_ = fs.MkdirAll("viccelek1", 0755)
	_ = afero.WriteFile(fs, "viccelek1/kaki", []byte(""), 0644)

	_ = fs.MkdirAll("viccelek2", 0755)
	_ = afero.WriteFile(fs, "viccelek2/iostream", []byte(""), 0644)

	if res := cpp.ExtractHeaderNames(fs, "./", src); len(res) != 0 {
		t.Errorf("0 != %d: %v", len(res), res)
	}

	if res := cpp.ExtractHeaderNames(fs, "./headers", src); len(res) != 2 {
		t.Errorf("2 != %d: %v", len(res), res)
	}

	if res := cpp.ExtractHeaderNames(fs, "./viccelek1", src); len(res) != 1 {
		t.Errorf("0 != %d: %v", len(res), res)
	}

	if res := cpp.ExtractHeaderNames(fs, "./viccelek2", src); len(res) != 1 {
		t.Errorf("0 != %d: %v", len(res), res)
	}
}
