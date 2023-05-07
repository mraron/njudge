package language

import "sort"

type Store interface {
	Register(name string, l Language)
	List() []Language
	Get(name string) Language
}

type MapStore struct {
	langList map[string]Language
}

func NewMapStore() *MapStore {
	return &MapStore{make(map[string]Language)}
}

func (m *MapStore) Register(name string, l Language) {
	m.langList[name] = l
}

func (m *MapStore) List() []Language {
	ans := make([]Language, len(m.langList))

	ind := 0
	for _, val := range m.langList {
		ans[ind] = val
		ind++
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Name() < ans[j].Name()
	})

	return ans
}

func (m *MapStore) Get(name string) Language {
	if val, ok := m.langList[name]; ok {
		return val
	}

	return nil
}

var DefaultStore Store

func init() {
	DefaultStore = NewMapStore()
}
