package roles

type Action int

const (
	ActionCreate Action = iota
	ActionDelete
	ActionEdit
	ActionView
)

type Role string

type Entity string

type Rule struct {
	Action Action
	Entity Entity
}

var Rules map[Role][]Rule

func Can(role Role, e Entity, a Action) bool {
	for _, rule := range Rules[role] {
		if rule.Action == a && rule.Entity == e {
			return true
		}
	}

	return false
}

func AddRule(e Entity, a Action, r Role) {
	if _, ok := Rules[r]; !ok {
		Rules[r] = make([]Rule, 0)
	}

	Rules[r] = append(Rules[r], Rule{a, e})
}

func init() {
	Rules = make(map[Role][]Rule)

	AddRule("admin_panel", ActionView, "admin")

	AddRule("api/v1/problem_rels", ActionView, "admin")
	AddRule("api/v1/problem_rels", ActionCreate, "admin")
	AddRule("api/v1/problem_rels", ActionDelete, "admin")
	AddRule("api/v1/problem_rels", ActionEdit, "admin")
}