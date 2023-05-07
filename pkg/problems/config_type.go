package problems

import (
	"errors"
	"fmt"

	"github.com/spf13/afero"
)

var (
	ErrorNameUsed = errors.New("config type name already in use")
	ErrorNoMatch  = errors.New("parser can't decide config type, no match")
)

// ConfigIdentifier is a function for some config type which takes a path and returns true if it thinks that its respective parser can parse it
type ConfigIdentifier func(afero.Fs, string) bool

// ConfigParser is a function for some config type which parses the problem from some path provided to it
type ConfigParser func(afero.Fs, string) (Problem, error)

// ConfigStore is an interface with which you can register/deregister certain config types and parse a problem using these config types
type ConfigStore interface {
	Register(string, ConfigParser, ConfigIdentifier) error
	Deregister(string) error

	Parse(afero.Fs, string) (Problem, error)
}

type problemConfigType struct {
	name        string
	parser      ConfigParser
	identifiers ConfigIdentifier
}

type configStore struct {
	configTypes []problemConfigType
}

// NewConfigStore returns the default implementation of ConfigStore
func NewConfigStore() ConfigStore {
	return &configStore{make([]problemConfigType, 0)}
}

func (cs *configStore) Register(name string, parser ConfigParser, identifier ConfigIdentifier) error {
	for _, val := range cs.configTypes {
		if val.name == name {
			return ErrorNameUsed
		}
	}

	if parser == nil {
		return fmt.Errorf("parser can't be nil")
	}

	if identifier == nil {
		return fmt.Errorf("identifier can't be nil")
	}

	cs.configTypes = append(cs.configTypes, problemConfigType{name, parser, identifier})
	return nil
}

func (cs *configStore) Deregister(name string) error {
	index := -1
	for ind := range cs.configTypes {
		if cs.configTypes[ind].name == name {
			index = ind
		}
	}

	if index == -1 {
		return fmt.Errorf("config type name not found")
	}

	cs.configTypes = append(cs.configTypes[:index], cs.configTypes[index+1:]...)
	return nil
}

func (cs *configStore) Parse(fs afero.Fs, path string) (Problem, error) {
	match := -1
	for ind := range cs.configTypes {
		if cs.configTypes[ind].identifiers(fs, path) {
			match = ind
			break
		}
	}

	if match == -1 {
		return nil, ErrorNoMatch
	}

	return cs.configTypes[match].parser(fs, path)
}

var globalConfigStore ConfigStore

// RegisterConfigType registers a config type to the global ConfigStore
func RegisterConfigType(name string, parser ConfigParser, identifier ConfigIdentifier) error {
	return globalConfigStore.Register(name, parser, identifier)
}

// DeregisterConfigType unregisters a config type to the global ConfigStore
func DeregisterConfigType(name string) error {
	return globalConfigStore.Deregister(name)
}

// Parse tries to parse a problem with the help of the global ConfigStore
func Parse(path string) (Problem, error) {
	return globalConfigStore.Parse(afero.NewOsFs(), path)
}

func init() {
	globalConfigStore = NewConfigStore()
}
