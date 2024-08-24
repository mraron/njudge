package web

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (s *Server) StartBackgroundJobs(ctx context.Context) {
	go s.runUpdateProblems(ctx)
	go s.pruneUnactivatedUsers(ctx)
}

func (s *Server) runUpdateProblems(ctx context.Context) {
	for {
		if err := s.ProblemStore.UpdateProblems(); err != nil {
			s.Logger.ErrorContext(ctx, "error updating problems", "error", err)
		}

		time.Sleep(20 * time.Second)
	}
}

func (s *Server) pruneUnactivatedUsers(ctx context.Context) {
	for {
		lst, err := s.Users.GetAll(ctx)
		if err != nil {
			s.Logger.ErrorContext(ctx, "error pruning unactivated users", err)
		}
		cnt := 0
		for ind := range lst {
			if !lst[ind].ActivationInfo.Activated && time.Since(lst[ind].Created) > 48 * time.Hour {
				err = errors.Join(err, s.Users.Delete(ctx, lst[ind].ID))
				cnt ++ 
			}
		}
		if err != nil {
			s.Logger.ErrorContext(ctx, "error occured while pruning", err)
		}else {
			s.Logger.InfoContext(ctx, fmt.Sprintf("pruned %d inactive, unactivated accounts", cnt))
		}
		time.Sleep(1 * time.Hour)
	}
}