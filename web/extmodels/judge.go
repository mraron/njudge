package extmodels

import (
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/web/models"
)

type Judge struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Ping        int           `json:"ping"`
	Online      bool          `json:"online"`

	judge.ServerStatus
}

func NewJudgeFromModelsJudge(j *models.Judge) (res Judge) {
	res.Id = int64(j.ID)
	res.Host = j.Host
	res.Port = j.Port
	res.Ping = j.Ping
	res.Online = j.Online

	if server, err := judge.ParseServerStatus(j.State); err == nil {
		res.Name = server.Id
		res.Load = server.Load
		res.ProblemList = server.ProblemList
		res.Uptime = server.Uptime
	}

	return
}