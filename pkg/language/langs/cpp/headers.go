package cpp

import (
	"path/filepath"
	"regexp"

	"github.com/spf13/afero"
)

var includeRegexp = regexp.MustCompile(`#include\s*[<"]([^">]+)[>"]`)

func ExtractHeaderNames(fs afero.Fs, dir string, src []byte) []string {
	names := includeRegexp.FindAllStringSubmatch(string(src), -1)
	res := make([]string, 0, len(names))
	for ind := range names {
		if _, err := fs.Stat(filepath.Join(dir, names[ind][1])); err == nil {
			res = append(res, names[ind][1])
		}
	}

	return res
}
