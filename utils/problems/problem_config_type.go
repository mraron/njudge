package problems

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ERROR_NAME_USED           = errors.New("Problem name already in use")
	ERROR_AMBIGUOUS_STRUCTURE = errors.New("Parser can't decide problem type, because of many candidates")
	ERROR_NO_MATCH            = errors.New("Parser can't decide problem type, no match")
)

type ConfigIdentifier func(string) bool

type ConfigParser func(string) (Problem, error)

type problemConfigType struct {
	name        string
	parser      ConfigParser
	identifiers []ConfigIdentifier
}

var problemTypes []problemConfigType

func CombineConfigIdentifiers(identifiers ...ConfigIdentifier) ConfigIdentifier {
	return func(s string) bool {
		ok := true
		for _, val := range identifiers {
			if !val(s) {
				ok = false
			}
		}

		return ok
	}
}

func DefaultConfigIdentifier(name string) ConfigIdentifier {
	return func(s string) bool {
		_, err := os.Stat(filepath.Join(s, name+".problem"))
		return !os.IsNotExist(err)
	}

}

func RegisterConfigType(name string, parser ConfigParser, identifiers ...ConfigIdentifier) error {
	for _, val := range problemTypes {
		if val.name == name {
			return ERROR_NAME_USED
		}
	}

	if len(identifiers) == 0 {
		identifiers = append(identifiers, DefaultConfigIdentifier(name))
	}

	problemTypes = append(problemTypes, problemConfigType{name, parser, identifiers})
	return nil
}

func Parse(dir string) (Problem, error) {
	matches, first_match := 0, -1
	for ind, val := range problemTypes {
		ok := true
		for _, identifier := range val.identifiers {
			if !identifier(dir) {
				ok = false
			}
		}

		if ok {
			matches++
			if first_match == -1 {
				first_match = ind
			}
		}
	}

	if matches > 1 {
		return nil, fmt.Errorf("%w: %s", ERROR_AMBIGUOUS_STRUCTURE, dir)
	} else if matches == 0 {
		return nil, fmt.Errorf("%w: %s", ERROR_NO_MATCH, dir)
	}

	return problemTypes[first_match].parser(dir)
}

func init() {
	problemTypes = make([]problemConfigType, 0)
}
