package language_test

import (
	"testing"

	"github.com/mraron/njudge/pkg/language"
)

func TestStoreAllExcept(t *testing.T) {
	store := language.NewListStore()
	store.Register("a", nil)
	store.Register("b", nil)
	store.Register("c", nil)

	if x := len(language.ListExcept(store, []string{})); x != 3 {
		t.Error("x !=", 3)
	}
	if x := len(language.ListExcept(store, []string{"a"})); x != 2 {
		t.Error("x !=", 2)
	}
	if x := len(language.ListExcept(store, []string{"c"})); x != 2 {
		t.Error("x !=", 2)
	}
	if x := len(language.ListExcept(store, []string{"b", "c"})); x != 1 {
		t.Error("x !=", 1)
	}
	if x := len(language.ListExcept(store, []string{"a", "b", "c"})); x != 0 {
		t.Error("x !=", 0)
	}
}
