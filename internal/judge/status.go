package judge

import (
	"encoding/json"
	"time"

	"golang.org/x/exp/slices"
)

type Status struct {
	Id           string        `json:"id" mapstructure:"id"`
	Host         string        `json:"host" mapstructure:"host"`
	Port         string        `json:"port" mapstructure:"port"`
	Url          string        `json:"url"`
	Load         float64       `json:"load"`
	Uptime       time.Duration `json:"uptime"`
	ProblemList  []string      `json:"problem_list"`
	LanguageList []string      `json:"language_list"`
}

func ParseStatus(s string) (res Status, err error) {
	err = json.Unmarshal([]byte(s), &res)
	return
}

func (s Status) SupportsProblem(want string) bool {
	return slices.Contains(s.ProblemList, want)
}

func (s Status) SupportsLanguage(want string) bool {
	return slices.Contains(s.LanguageList, want)
}

func (s Status) String() string {
	res, _ := json.Marshal(s)
	return string(res)
}
