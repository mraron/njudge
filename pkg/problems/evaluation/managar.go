package evaluation

import (
	"context"
	"errors"
	"sync"
)

type Step func(context.Context) error

type Manager struct {
	steps         map[string]Step
	stepWaitGroup map[string]*sync.WaitGroup
	nextSteps     map[string][]string
}

func NewManager() *Manager {
	return &Manager{
		steps:         make(map[string]Step),
		stepWaitGroup: make(map[string]*sync.WaitGroup),
		nextSteps:     make(map[string][]string),
	}
}

func (m *Manager) AddStep(name string, fun Step, deps []string) error {
	m.steps[name] = fun
	m.stepWaitGroup[name] = &sync.WaitGroup{}
	for i := range deps {
		m.stepWaitGroup[name].Add(1)
		m.nextSteps[deps[i]] = append(m.nextSteps[deps[i]], name)
	}
	return nil
}

func (m *Manager) Run(ctx context.Context) error {
	var errs []error
	ctx2, cancelFunc := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}
	for k, v := range m.steps {
		k, v := k, v
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.stepWaitGroup[k].Wait()
			if ctx2.Err() == nil {
				if err := v(ctx2); err != nil {
					cancelFunc()
					errs = append(errs, err)
				}
			}

			for next := range m.nextSteps[k] {
				m.stepWaitGroup[m.nextSteps[k][next]].Done()
			}
		}()
	}
	wg.Wait()
	cancelFunc()

	if ctx2.Err() != nil {
		errs = append(errs, ctx2.Err())
	}
	return errors.Join(errs...)
}
