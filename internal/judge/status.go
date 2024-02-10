package judge

import (
	"encoding/json"
	"slices"
	"time"
)

type ServerStatus struct {
	Host         string        `json:"host"`
	Port         string        `json:"port"`
	Url          string        `json:"url"`
	Load         float64       `json:"load"`
	Uptime       time.Duration `json:"uptime"`
	ProblemList  []string      `json:"problem_list"`
	LanguageList []string      `json:"language_list"`
}

func ParseServerStatus(s string) (res ServerStatus, err error) {
	err = json.Unmarshal([]byte(s), &res)
	return
}

func (s ServerStatus) SupportsProblem(want string) bool {
	return slices.Contains(s.ProblemList, want)
}

func (s ServerStatus) SupportsLanguage(want string) bool {
	return slices.Contains(s.LanguageList, want)
}

func (s ServerStatus) String() string {
	res, _ := json.Marshal(s)
	return string(res)
}
