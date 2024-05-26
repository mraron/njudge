package glue

import (
	"context"
	"database/sql"
	"errors"
	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/mraron/njudge/pkg/problems"
	"log/slog"
	"math/rand"
	"sync"
)

type Judges struct {
	Conn   *sql.DB
	Logger *slog.Logger

	JudgeClients map[int]*judge.Client
	JudgeModels  models.JudgeSlice

	mutex sync.RWMutex
}

func NewJudges(conn *sql.DB, logger *slog.Logger) *Judges {
	return &Judges{
		Conn:   conn,
		Logger: logger.With("service", "glue_judges"),

		JudgeClients: make(map[int]*judge.Client),
	}
}

func (j *Judges) Update(ctx context.Context) {
	j.mutex.Lock()
	j.Logger.Info("🔄\tupdating judges")
	defer j.mutex.Unlock()

	var err error
	if j.JudgeModels, err = models.Judges().All(ctx, j.Conn); err != nil {
		j.Logger.Error("‼️\tFailed to get judges", "error", err)
		return
	}
	for _, curr := range j.JudgeModels {
		j.JudgeClients[curr.ID] = judge.NewClient(curr.URL)
	}
}

func (j *Judges) Judge(ctx context.Context, sub judge.Submission, callback judge.ResultCallback) (*problems.Status, error) {
	j.mutex.RLock()
	if len(j.JudgeClients) == 0 {
		j.mutex.RUnlock()
		return nil, errors.New("no judge found")
	}

	rnd := rand.Intn(len(j.JudgeModels))
	client := j.JudgeClients[j.JudgeModels[rnd].ID]

	j.mutex.RUnlock()
	return client.Judge(ctx, sub, callback)
}
