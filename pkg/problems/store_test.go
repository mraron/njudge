package problems_test

import (
	"path/filepath"
	"testing"

	"github.com/mraron/njudge/pkg/problems"
	"github.com/spf13/afero"
)

type dummyProblemConfig struct {
	configName string
}

func newDummyProblemConfig(configName string) *dummyProblemConfig {
	return &dummyProblemConfig{configName: configName}
}

func (d dummyProblemConfig) identifier() problems.ConfigIdentifier {
	return func(f afero.Fs, s string) bool {
		_, err := f.Stat(filepath.Join(s, d.configName))
		return err == nil
	}
}

func (d dummyProblemConfig) parser() problems.ConfigParser {
	return func(f afero.Fs, s string) (problems.Problem, error) {
		return nil, nil
	}
}

func TestFSStore(t *testing.T) {
	config := newDummyProblemConfig("feladat.xhtml")
	configStore := problems.NewConfigList()
	configStore.Register("dummy", config.parser(), config.identifier())

	f := afero.NewMemMapFs()
	f.MkdirAll("problems/aplusb", 0755)
	afero.WriteFile(f, "problems/aplusb/feladat.xhtml", []byte("lalal"), 0644)
	f.MkdirAll("problems/aplusb2", 0755)
	afero.WriteFile(f, "problems/aplusb2/feladat.xhtml", []byte("lalal"), 0644)
	f.MkdirAll("problems/aplusb2/aplusb3", 0755)
	afero.WriteFile(f, "problems/aplusb2/aplusb3/feladat.xhtml", []byte("lalal"), 0644)

	f.MkdirAll("problems/prefixed", 0755)
	afero.WriteFile(f, "problems/prefixed/.njudge_prefix", []byte("XX"), 0644)
	f.MkdirAll("problems/prefixed/first", 0755)
	afero.WriteFile(f, "problems/prefixed/first/feladat.xhtml", []byte("lalal"), 0644)
	f.MkdirAll("problems/prefixed/second", 0755)
	afero.WriteFile(f, "problems/prefixed/second/feladat.xhtml", []byte("lalal"), 0644)
	f.MkdirAll("problems/prefixed/third", 0755)
	f.MkdirAll("problems/prefixed/.hidden", 0755)
	afero.WriteFile(f, "problems/prefixed/.hidden/feladat.xhtml", []byte("lalal"), 0644)
	f.MkdirAll("problems/prefixed/ignored", 0755)
	afero.WriteFile(f, "problems/prefixed/ignored/feladat.xhtml", []byte("lalal"), 0644)
	afero.WriteFile(f, "problems/prefixed/ignored/.njudge_ignore", []byte("lalal"), 0644)

	f.MkdirAll("problems/recursive", 0755)
	afero.WriteFile(f, "problems/recursive/.njudge_ignore", []byte("lalal"), 0644)
	f.MkdirAll("problems/recursive/nono", 0755)
	afero.WriteFile(f, "problems/recursive/nono/feladat.xhtml", []byte("lalal"), 0644)

	store := problems.NewFsStore("problems/", problems.FsStoreUseFs(f), problems.FsStoreUseConfigStore(configStore))
	if err := store.Update(); err != nil {
		t.Error(err)
	}
	if has, err := store.Has("aplusb"); !has || err != nil {
		t.Error("aplusb", has, err)
	}
	if has, err := store.Has("aplusb2"); !has || err != nil {
		t.Error("aplusb2", has, err)
	}
	if has, err := store.Has("aplusb3"); has {
		t.Error("aplusb3", has, err)
	}

	if has, err := store.Has("XX_first"); !has || err != nil {
		t.Error("XX_first", has, err)
	}
	if has, err := store.Has("XX_second"); !has || err != nil {
		t.Error("XX_second", has, err)
	}
	if has, err := store.Has("XX_third"); has {
		t.Error("XX_third", has, err)
	}
	if has, err := store.Has("XX_.hidden"); has {
		t.Error("XX_.hidden", has, err)
	}
	if has, err := store.Has("XX_hidden"); has {
		t.Error("XX_hidden", has, err)
	}
	if has, err := store.Has("XX_ignored"); has {
		t.Error("XX_ignored", has, err)
	}
	if has, err := store.Has("nono"); has {
		t.Error("nono", has, err)
	}
}
