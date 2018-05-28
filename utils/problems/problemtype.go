package problems

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ERROR_NAME_USED = errors.New("Problem name already in use")
	ERROR_AMBIGUOUS_STRUCTURE = errors.New("Parser can't decide problem type, because of many candidates")
	ERROR_NO_MATCH = errors.New("Parser can't decide problem type, no match")
)

type Identifier func(string) bool

type Parser func(string) (Problem, error)

type problemType struct {
	name string
	parser Parser
	identifiers []Identifier
}

var problemTypes []problemType

func CombineIdentifiers(identifiers ...Identifier) Identifier {
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

func DefaultIdentifier(name string) Identifier {
	return func(s string) bool {
		_, err := os.Stat(filepath.Join(s, name + ".problem"))
		return !os.IsNotExist(err)
	}

}

func RegisterType(name string, parser Parser, identifiers ...Identifier) (error) {
	for _, val := range problemTypes {
		if val.name == name {
			return ERROR_NAME_USED
		}
	}

	if len(identifiers) == 0 {
		identifiers = append(identifiers, DefaultIdentifier(name))
	}

	problemTypes = append(problemTypes, problemType{name, parser, identifiers})
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
			matches ++
			if first_match == -1 {
				first_match = ind
			}
		}
	}

	if matches > 1 {
		return nil, ERROR_AMBIGUOUS_STRUCTURE
	}else if matches == 0 {
		return nil, ERROR_NO_MATCH
	}

	return problemTypes[first_match].parser(dir)
}


func init() {
	problemTypes = make([]problemType, 0)
}
