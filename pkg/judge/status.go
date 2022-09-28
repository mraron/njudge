package judge

import (
	"encoding/json"
	"time"
)

type Status struct {
	Id          string        `json:"id" mapstructure:"id"`
	Host        string        `json:"host" mapstructure:"host"`
	Port        string        `json:"port" mapstructure:"port"`
	Url         string        `json:"url"`
	Load        float64       `json:"load"`
	Uptime      time.Duration `json:"uptime"`
	ProblemList []string      `json:"problem_list"`
}

func ParseStatus(s string) (res Status, err error) {
	err = json.Unmarshal([]byte(s), &res)
	return
}

func (s Status) SupportsProblem(want string) bool {
	for _, key := range s.ProblemList {
		if key == want {
			return true
		}
	}

	return false
}

func (s Status) String() string {
	res, _ := json.Marshal(s)
	return string(res)
}
