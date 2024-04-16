package cached

import (
	"context"
	"github.com/erni27/imcache"
	"github.com/mraron/njudge/internal/njudge"
	"time"
)

type problemAndUser struct {
	ProblemID, UserID int
}

type SolvedStatusQuery struct {
	solvedStatusQuery njudge.SolvedStatusQuery
	cache             *imcache.Cache[problemAndUser, njudge.SolvedStatus]
}

func (ssq *SolvedStatusQuery) GetSolvedStatus(ctx context.Context, problemID, userID int) (njudge.SolvedStatus, error) {
	key := problemAndUser{problemID, userID}
	if val, ok := ssq.cache.Get(key); ok {
		return val, nil
	}

	val, err := ssq.solvedStatusQuery.GetSolvedStatus(ctx, problemID, userID)
	if err != nil {
		return 0, nil
	}

	ssq.cache.Set(key, val, imcache.WithDefaultExpiration())
	return val, nil
}

func NewSolvedStatusQuery(solvedStatusQuery njudge.SolvedStatusQuery, ttl time.Duration) *SolvedStatusQuery {
	return &SolvedStatusQuery{
		cache: imcache.New[problemAndUser, njudge.SolvedStatus](
			imcache.WithDefaultExpirationOption[problemAndUser, njudge.SolvedStatus](ttl),
		),
		solvedStatusQuery: solvedStatusQuery,
	}
}
