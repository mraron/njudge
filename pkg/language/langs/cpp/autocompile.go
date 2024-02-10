package cpp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

func AutoCompile(ctx context.Context, fs afero.Fs, s sandbox.Sandbox, workDir, src, dst string) error {
	if src == "" {
		return nil
	}

	if st, err := fs.Stat(dst); os.IsNotExist(err) || st.Size() == 0 || st.Mode()&0100 != 0100 {
		if binary, err := fs.Create(dst); err == nil {
			if file, err := fs.Open(src); err == nil {
				var buf bytes.Buffer
				if err := s.Init(ctx); err != nil {
					return errors.Join(err, binary.Close(), file.Close())
				}
				defer func(s sandbox.Sandbox, ctx context.Context) {
					_ = s.Cleanup(ctx)
				}(s, ctx)

				contents, err := afero.ReadFile(fs, src)
				if err != nil {
					return errors.Join(err, file.Close(), binary.Close())
				}

				var headers []sandbox.File
				for _, header := range ExtractHeaderNames(fs, workDir, contents) {
					headerContents, err := afero.ReadFile(fs, filepath.Join(workDir, header))
					if err != nil {
						return errors.Join(err, file.Close(), binary.Close())
					}

					headers = append(headers, sandbox.File{
						Name:   header,
						Source: bytes.NewReader(headerContents),
					})
				}

				if _, err := Std17.Compile(s, sandbox.File{
					Name:   filepath.Base(src),
					Source: file,
				}, &buf, nil); err != nil {
					return errors.Join(err, binary.Close(), file.Close(), fmt.Errorf("compile error: %v", buf.String()))
				}

				if err := fs.Chmod(dst, 0755); err != nil {
					return errors.Join(err, binary.Close(), file.Close())
				}

				return errors.Join(binary.Close(), file.Close())
			} else {
				return errors.Join(err, binary.Close())
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
