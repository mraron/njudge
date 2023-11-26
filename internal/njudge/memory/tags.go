package memory

import (
	"context"
	"sync"

	"github.com/mraron/njudge/internal/njudge"
)

type Tags struct {
	sync.Mutex
	nextId int
	data   []njudge.Tag
}

func NewTags() *Tags {
	return &Tags{
		nextId: 1,
		data:   make([]njudge.Tag, 0),
	}
}

func (m *Tags) Get(ctx context.Context, ID int) (*njudge.Tag, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorTagNotFound
}

func (m *Tags) GetAll(ctx context.Context) ([]njudge.Tag, error) {
	m.Lock()
	defer m.Unlock()
	res := make([]njudge.Tag, len(m.data))
	copy(res, m.data)

	return res, nil
}

func (m *Tags) GetByName(ctx context.Context, name string) (*njudge.Tag, error) {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].Name == name {
			res := m.data[ind]
			return &res, nil
		}
	}

	return nil, njudge.ErrorTagNotFound
}

func (m *Tags) Insert(ctx context.Context, u njudge.Tag) (*njudge.Tag, error) {
	m.Lock()
	defer m.Unlock()
	u.ID = m.nextId
	m.nextId++
	m.data = append(m.data, u)

	res := m.data[len(m.data)-1]
	return &res, nil
}

func (m *Tags) Delete(ctx context.Context, ID int) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == ID {
			m.data[ind] = m.data[len(m.data)-1]
			m.data = m.data[:len(m.data)-1]
			return nil
		}
	}

	return njudge.ErrorTagNotFound
}

func (m *Tags) Update(ctx context.Context, user njudge.Tag) error {
	m.Lock()
	defer m.Unlock()
	for ind := range m.data {
		if m.data[ind].ID == user.ID {
			m.data[ind] = user
			return nil
		}
	}
	return njudge.ErrorTagNotFound
}

type TagsService struct {
	tags             njudge.Tags
	problems         njudge.Problems
	problemInfoQuery njudge.ProblemInfoQuery
}

func NewTagsService(tags njudge.Tags, problems njudge.Problems, problemInfoQuery njudge.ProblemInfoQuery) *TagsService {
	return &TagsService{
		tags:             tags,
		problems:         problems,
		problemInfoQuery: problemInfoQuery,
	}
}

func (ts *TagsService) Add(ctx context.Context, tagID int, problemID int, userID int) error {
	p, err := ts.problems.Get(ctx, problemID)
	if err != nil {
		return err
	}
	t, err := ts.tags.Get(ctx, tagID)
	if err != nil {
		return err
	}

	if p.HasTag(*t) {
		return nil
	}

	pinfo, err := ts.problemInfoQuery.GetProblemData(ctx, p.ID, userID)
	if err != nil {
		return err
	}

	if pinfo.UserInfo.SolvedStatus != njudge.Solved {
		return njudge.ErrorUnableToModifyProblemTags
	}

	p.AddTag(*t, userID)
	return ts.problems.Update(ctx, *p, njudge.Fields(njudge.ProblemFields.Tags))
}

func (ts *TagsService) Delete(ctx context.Context, tagID int, problemID int, userID int) error {
	p, err := ts.problems.Get(ctx, problemID)
	if err != nil {
		return err
	}
	t, err := ts.tags.Get(ctx, tagID)
	if err != nil {
		return err
	}

	pinfo, err := ts.problemInfoQuery.GetProblemData(ctx, p.ID, userID)
	if err != nil {
		return err
	}

	if pinfo.UserInfo.SolvedStatus != njudge.Solved {
		return njudge.ErrorUnableToModifyProblemTags
	}

	if err := p.DeleteTag(*t); err != nil {
		return err
	}

	return ts.problems.Update(ctx, *p, njudge.Fields(njudge.ProblemFields.Tags))
}
