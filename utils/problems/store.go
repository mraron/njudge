package problems

import (
	"errors"
	"fmt"
	"github.com/spf13/afero"
	"path/filepath"
	"sync"
)

var ErrorProblemNotFound = errors.New("problem not found")

//Store is an interface which is used to access a bunch of problems for example from the filesystem
type Store interface {
	List() ([]string, error)
	Has(string) (bool, error)
	Get(string) (Problem, error)
	Update() error
	UpdateProblem(string) error
}

type FsStore struct {
	cs ConfigStore
	fs afero.Fs
	dir string

	problems      map[string]Problem
	problemsList []string

	sync.Mutex
}

type FsStoreOptions func(*FsStore)

func FsStoreUseConfigStore(cs ConfigStore) FsStoreOptions {
	return func(fs *FsStore) {
		fs.cs = cs
	}
}

func FsStoreUseFs(afs afero.Fs) FsStoreOptions {
	return func(fs *FsStore) {
		fs.fs = afs
	}
}

func NewFsStore(dir string, options... FsStoreOptions) *FsStore {
	fs := &FsStore{cs: globalConfigStore, fs: afero.NewOsFs(), dir: dir, problems: make(map[string]Problem), problemsList: make([]string, 0)}
	for _, opt := range options {
		opt(fs)
	}

	return fs
}

func (s *FsStore) List() ([]string, error) {
	return s.problemsList, nil
}

func (s *FsStore) Has(p string) (bool, error) {
	for _, key := range s.problemsList {
		if key == p {
			return true, nil
		}
	}

	return false, nil
}

func (s *FsStore) Get(p string) (Problem, error) {
	s.Lock()
	defer s.Unlock()

	if res, ok := s.problems[p]; ok {
		return res, nil
	}
	return nil, ErrorProblemNotFound
}

func (s *FsStore) Update() error {
	files, err := afero.ReadDir(s.fs, s.dir)
	if err != nil {
		return err
	}

	s.problemsList = s.problemsList[:0]

	errs := make([]error, 0)
	for _, file := range files {
		if file.IsDir() {
			name := filepath.Base(file.Name())
			if err := s.UpdateProblem(name); err != nil {
				errs = append(errs, err)
			}

			s.problemsList = append(s.problemsList, name)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	err = errs[0]
	for i := 1; i < len(errs); i++ {
		err = fmt.Errorf("%v; %v", err, errs[i])
	}

	return err
}

func (s *FsStore) UpdateProblem(p string) error {
	s.Lock()
	defer s.Unlock()

	path := filepath.Join(s.dir, p)
	prob, err := s.cs.Parse(path)
	if err != nil {
		return err
	}

	s.problems[p] = prob
	return nil
}





