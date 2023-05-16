package language

import "sort"

type Store interface {
	Register(id string, l Language)
	List() []Language
	Get(id string) Language
}

type LanguageWrapper struct {
	id string
	Language
}

func (lw LanguageWrapper) Id() string {
	return lw.id
}

type ListStore struct {
	langList []Language
}

func NewListStore() *ListStore {
	return &ListStore{make([]Language, 0)}
}

func (m *ListStore) Register(id string, l Language) {
	m.langList = append(m.langList, LanguageWrapper{id, l})
}

func (m *ListStore) List() []Language {
	ans := make([]Language, len(m.langList))

	ind := 0
	for _, val := range m.langList {
		ans[ind] = val
		ind++
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Id() < ans[j].Id()
	})

	return ans
}

func (m *ListStore) Get(id string) Language {
	for ind := range m.langList {
		if m.langList[ind].Id() == id {
			return m.langList[ind]
		}
	}

	return nil
}

var DefaultStore Store

func init() {
	DefaultStore = NewListStore()
}
