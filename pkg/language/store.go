package language

import (
	"errors"
	"sort"

	"slices"
)

var ErrorLanguageNotFound = errors.New("language not found")

type NotFoundError struct {
	ID string
}

func (n NotFoundError) Error() string {
	return "language not found: " + n.ID
}

func (n NotFoundError) Is(err error) bool {
	return err == ErrorLanguageNotFound
}

// Store is an interface which is used to capture the notion of storing languages.
// Via it's Store.Register method it's possible to override the underlying ID of the language for the outside world.
type Store interface {
	Register(id string, l Language)
	List() []Language
	Get(id string) (Language, error)
}

// ListExcept returns a slice of languages except some.
func ListExcept(s Store, except []string) []Language {
	var res []Language
	for _, elem := range s.List() {
		if !slices.Contains(except, elem.ID()) {
			res = append(res, elem)
		}
	}

	return res
}

// Wrapper overrides a Language's ID.
type Wrapper struct {
	IDWrapper string
	Language
}

func (w Wrapper) ID() string {
	return w.IDWrapper
}

// ListStore is a basic implementation (and probably only realistic, so maybe an interface is not really necessary) of a Store.
type ListStore struct {
	LanguageList []Language
}

func NewListStore() *ListStore {
	return &ListStore{make([]Language, 0)}
}

func (m *ListStore) Register(id string, l Language) {
	m.LanguageList = append(m.LanguageList, Wrapper{id, l})
}

func (m *ListStore) List() []Language {
	ans := make([]Language, len(m.LanguageList))

	ind := 0
	for _, val := range m.LanguageList {
		ans[ind] = val
		ind++
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].ID() < ans[j].ID()
	})

	return ans
}

func (m *ListStore) Get(id string) (Language, error) {
	for ind := range m.LanguageList {
		if m.LanguageList[ind].ID() == id {
			return m.LanguageList[ind], nil
		}
	}

	return nil, NotFoundError{id}
}

// DefaultStore is a store which all Language objects should register themselves in.
var DefaultStore Store

func init() {
	DefaultStore = NewListStore()
}
