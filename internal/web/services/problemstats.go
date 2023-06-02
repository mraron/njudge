package services

import (
	"context"
	"github.com/mraron/njudge/internal/web/domain/problem"
)

type ProblemStatsService interface {
	GetStatsData(ctx context.Context, p problem.Problem, userID int) (*problem.StatsData, error)
}
