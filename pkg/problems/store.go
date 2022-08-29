package problems

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/afero"
)

var ErrorProblemNotFound = errors.New("problem not found")

type problemNotFoundError struct {
	name string
}

func (perr problemNotFoundError) Error() string {
	return "problem not found: " + perr.name
}

func (perr problemNotFoundError) Is(target error) bool {
	return target == ErrorProblemNotFound
}

var ErrorProblemParse = errors.New("can't parse problems")

type ProblemParseError struct {
	Errors   []error
	Problems []string
}

func (perr ProblemParseError) Error() string {
	return fmt.Sprintf("can't parse problems: %v", perr.Errors)
}

func (perr ProblemParseError) Is(target error) bool {
	return target == ErrorProblemParse
}

//Store is an interface which is used to access a bunch of problems for example from the filesystem
type Store interface {
	List() ([]string, error)
	Has(string) (bool, error)
	Get(string) (Problem, error)
	MustGet(string) Problem
	Update() error
	UpdateProblem(path string, id string) error
}

type FsStore struct {
	cs           ConfigStore
	fs           afero.Fs
	dir          string
	ignorePrefix bool

	ByPath map[string]string

	problems     map[string]Problem
	problemsList []string

	sync.RWMutex
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

func FsStoreIgnorePreifx() FsStoreOptions {
	return func(fs *FsStore) {
		fs.ignorePrefix = true
	}
}

func NewFsStore(dir string, options ...FsStoreOptions) *FsStore {
	fsStore := &FsStore{
		cs:           globalConfigStore,
		fs:           afero.NewOsFs(),
		dir:          dir,
		problems:     make(map[string]Problem),
		problemsList: make([]string, 0),
		ByPath:       make(map[string]string),
	}
	for _, opt := range options {
		opt(fsStore)
	}

	return fsStore
}

func (s *FsStore) List() ([]string, error) {
	s.RLock()
	defer s.RUnlock()

	lst := make([]string, len(s.problemsList))
	copy(lst, s.problemsList)

	return lst, nil
}

func (s *FsStore) Has(p string) (bool, error) {
	if _, err := s.Get(p); err != nil {
		return false, nil
	}

	return true, nil
}

func (s *FsStore) Get(p string) (Problem, error) {
	s.RLock()
	defer s.RUnlock()

	if res, ok := s.problems[p]; ok {
		return res, nil
	}

	for _, prob := range s.problems {
		if prob.Name() == p {
			return prob, nil
		}
	}

	return nil, problemNotFoundError{p}
}

func (s *FsStore) MustGet(p string) Problem {
	res, err := s.Get(p)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *FsStore) Update() error {
	errs := ProblemParseError{Errors: make([]error, 0), Problems: make([]string, 0)}
	lst := make([]string, 0)

	prefix, atPath := "", ""
	getId := func(path string) string {
		rel, err := filepath.Rel(atPath, path)
		if err != nil || prefix == "" {
			return filepath.Base(path)
		}

		if len(rel) == 1 || rel[0] != '.' {
			return prefix + "_" + filepath.Base(path)
		}

		return filepath.Base(path)
	}

	if err := afero.Walk(s.fs, s.dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return err
		}

		if strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if s.dir == path {
			return nil
		}

		//skip directories recursively with the .njudge_ignore file
		if _, err := os.Stat(filepath.Join(path, ".njudge_ignore")); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		} else {
			return nil
		}

		if !s.ignorePrefix {
			prefixFile := filepath.Join(path, ".njudge_prefix")
			if _, err := os.Stat(prefixFile); err != nil {
				if !errors.Is(err, os.ErrNotExist) {
					return err
				}
			} else {
				prefixBytes, err := ioutil.ReadFile(prefixFile)
				if err != nil {
					return err
				}

				prefix = strings.TrimSpace(string(prefixBytes))
				atPath = path
			}
		}

		if err := s.UpdateProblem(path, getId(path)); err != nil {
			if !errors.Is(err, ErrorNoMatch) {
				errs.Errors = append(errs.Errors, fmt.Errorf("%s: %w", info.Name(), err))
				errs.Problems = append(errs.Problems, info.Name())

				if err == ErrorNoMatch {
					return nil
				} else {
					return filepath.SkipDir
				}
			} else {
				return nil
			}
		} else {
			lst = append(lst, info.Name())
			return filepath.SkipDir
		}
	}); err != nil {
		return err
	}

	s.Lock()
	s.problemsList = make([]string, len(lst))
	copy(s.problemsList, lst)
	s.Unlock()

	if len(errs.Errors) == 0 {
		return nil
	}
	return errs
}

type problemWrapper struct {
	Problem

	nameOverride string
}

func (pw problemWrapper) Name() string {
	return pw.nameOverride
}

func (s *FsStore) UpdateProblem(path string, id string) error {
	s.Lock()
	defer s.Unlock()

	prob, err := s.cs.Parse(path)
	if err != nil {
		return err
	}

	s.ByPath[path] = id
	s.problems[id] = problemWrapper{prob, id}
	return nil
}
