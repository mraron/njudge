package extmodels

import (
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/web/models"
)

type Judge struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Ping   int    `json:"ping"`
	Online bool   `json:"online"`

	judge.ServerStatus
}

func NewJudgeFromModelsJudge(j *models.Judge) *Judge {
	res := &Judge{}

	res.Id = int64(j.ID)
	res.Host = j.Host
	res.Port = j.Port
	res.Ping = j.Ping
	res.Online = j.Online

	if server, err := judge.ParseServerStatus(j.State); err == nil {
		res.Name = "placeholder"
		res.Load = server.Load
		res.ProblemList = server.ProblemList
		res.Uptime = server.Uptime
	}

	return res
}
