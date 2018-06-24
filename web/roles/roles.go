package roles

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Action int

const (
	ActionCreate Action = iota
	ActionDelete
	ActionEdit
	ActionView
)

type Role string

func (r Role) Value() (driver.Value, error) {
	return driver.Value(r), nil
}

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		return errors.New("can't scan role from nil")
	}

	*r = Role(fmt.Sprint(value))

	return nil
}

func (r Role) Equal(r2 Role) bool {
	if len(r) != len(r2) {
		return false
	}

	for ind := range r {
		if r[ind] != r2[ind] {
			return false
		}
	}

	return true
}

type Entity string

type Rule struct {
	Role   Role
	Action Action
	Entity Entity
}

var Rules []Rule

func Can(role Role, a Action, e Entity) bool {
	for _, rule := range Rules {
		if rule.Role.Equal(role) && rule.Action == a && rule.Entity == e {
			return true
		}
	}

	return false
}

func AddRule(e Entity, a Action, r Role) {
	Rules = append(Rules, Rule{r, a, e})
}

func init() {
	Rules = make([]Rule, 0)

	AddRule("admin_panel", ActionView, Role("admin"))

	AddRule("api/v1/problem_rels", ActionView, Role("admin"))
	AddRule("api/v1/problem_rels", ActionCreate, Role("admin"))
	AddRule("api/v1/problem_rels", ActionDelete, Role("admin"))
	AddRule("api/v1/problem_rels", ActionEdit, Role("admin"))
}
