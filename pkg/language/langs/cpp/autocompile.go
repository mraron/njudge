package cpp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mraron/njudge/pkg/language"
	"github.com/spf13/afero"
	"go.uber.org/multierr"
)

func AutoCompile(fs afero.Fs, s language.Sandbox, workDir, src, dst string) error {
	if src == "" {
		return nil
	}

	if st, err := fs.Stat(dst); os.IsNotExist(err) || st.Size() == 0 || st.Mode()&0100 != 0100 {
		if binary, err := fs.Create(dst); err == nil {
			if file, err := fs.Open(src); err == nil {
				var buf bytes.Buffer
				if err := s.Init(log.New(ioutil.Discard, "", 0)); err != nil {
					return multierr.Combine(err, binary.Close(), file.Close())
				}
				defer s.Cleanup()

				conts, err := afero.ReadFile(fs, src)
				if err != nil {
					return multierr.Combine(err, file.Close(), binary.Close())
				}

				var headers []language.File
				for _, header := range ExtractHeaderNames(fs, workDir, conts) {
					headerConts, err := afero.ReadFile(fs, filepath.Join(workDir, header))
					if err != nil {
						return multierr.Combine(err, file.Close(), binary.Close())
					}

					headers = append(headers, language.File{
						Name:   header,
						Source: bytes.NewReader(headerConts),
					})
				}

				if err := Std17.Compile(s, language.File{
					Name:   filepath.Base(src),
					Source: file,
				}, binary, &buf, headers); err != nil {
					return multierr.Combine(err, binary.Close(), file.Close(), fmt.Errorf("compile error: %v", buf.String()))
				}

				if err := fs.Chmod(dst, 0755); err != nil {
					return multierr.Combine(err, binary.Close(), file.Close())
				}

				return multierr.Combine(binary.Close(), file.Close())
			} else {
				return multierr.Combine(err, binary.Close())
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
