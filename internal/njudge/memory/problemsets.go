package memory

import (
	"cmp"
	"github.com/mraron/njudge/internal/njudge"
	"golang.org/x/net/context"
	"slices"
	"sync"
)

type Problemsets struct {
	mutex sync.Mutex
	data  []njudge.Problemset
}

func NewProblemsets() *Problemsets {
	return &Problemsets{}
}

func (p *Problemsets) GetByName(ctx context.Context, problemsetName string) (*njudge.Problemset, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, probset := range p.data {
		if probset.Name == problemsetName {
			return &probset, nil
		}
	}
	return nil, njudge.ErrorProblemsetNotFound
}

func (p *Problemsets) GetAll(ctx context.Context) ([]njudge.Problemset, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	res := make([]njudge.Problemset, len(p.data))
	copy(res, p.data)
	return res, nil
}

func (p *Problemsets) Insert(ctx context.Context, problemset njudge.Problemset) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.data = append(p.data, problemset)
	return nil
}

type ProblemsetRanklistService struct {
	users njudge.Users
}

func NewProblemsetRanklistService(users njudge.Users) *ProblemsetRanklistService {
	return &ProblemsetRanklistService{users: users}
}

func (p ProblemsetRanklistService) GetRanklist(ctx context.Context, req njudge.ProblemsetRanklistRequest) (*njudge.ProblemsetRanklist, error) {
	// TODO req.Problemset
	users, err := p.users.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(users, func(a, b njudge.User) int {
		if a.Points == b.Points {
			return cmp.Compare(a.ID, b.ID)
		}
		return -cmp.Compare(a.Points, b.Points)
	})
	res := njudge.ProblemsetRanklist{}
	ind := 1
	for _, user := range users {
		if user.Role == "admin" && req.FilterAdmin {
			continue
		}
		res.Rows = append(res.Rows, njudge.ProblemsetRanklistRow{
			Place: ind,
			ID:    user.ID,
			Name:  user.Name,
			Score: float64(user.Points),
		})
		ind++
	}
	if req.PerPage > 0 {
		res.PaginationData.Count = len(res.Rows)
		res.PaginationData.PerPage = req.PerPage
		res.PaginationData.Page = req.Page
		res.PaginationData.Pages = (len(res.Rows) + req.PerPage - 1) / req.PerPage

		st, en := (req.Page-1)*req.PerPage, req.Page*req.PerPage
		if st < 0 {
			st = 0
		}
		if en > len(res.Rows) {
			en = len(res.Rows)
		}
		res.Rows = res.Rows[st:en]
	}
	return &res, nil
}
