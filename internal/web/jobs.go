package web

import (
	"context"
	"time"
)

func (s *Server) StartBackgroundJobs(ctx context.Context) {
	go s.runUpdateProblems(ctx)
}

func (s *Server) runUpdateProblems(ctx context.Context) {
	for {
		if err := s.ProblemStore.UpdateProblems(); err != nil {
			s.Logger.ErrorContext(ctx, "error updating problems", err)
		}

		time.Sleep(20 * time.Second)
	}
}
