package memory

import (
	"github.com/mraron/njudge/internal/njudge"
	"golang.org/x/net/context"
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
