package memory

import (
	"context"
	"sort"
	"strings"
	"unicode"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type ProblemListQuery struct {
	store problems.Store
	ps    njudge.Problems
	tags  njudge.Tags
	cs    njudge.Categories
}

func NewProblemListQuery(store problems.Store, ps njudge.Problems, tags njudge.Tags, cs njudge.Categories) *ProblemListQuery {
	return &ProblemListQuery{
		store: store,
		ps:    ps,
		tags:  tags,
		cs:    cs,
	}
}

func (p *ProblemListQuery) filterProblemset(ctx context.Context, req njudge.ProblemListRequest, pr njudge.Problem) (bool, error) {
	if req.Problemset != "" {
		if pr.Problemset != req.Problemset {
			return false, nil
		}
	}

	return true, nil
}

func (p *ProblemListQuery) filterTitle(ctx context.Context, req njudge.ProblemListRequest, pr njudge.Problem) (bool, error) {
	data, err := pr.WithStoredData(p.store)
	if err != nil {
		return false, err
	}

	curr := strings.ToLower(i18n.TranslateContent("hungarian", data.Titles()).String())
	want := strings.ToLower(req.TitleFilter)

	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	a, _, err := transform.String(t, curr)
	if err != nil {
		return false, err
	}
	b, _, err := transform.String(t, want)
	if err != nil {
		return false, err
	}

	return strings.Contains(a, b), nil
}

func (p *ProblemListQuery) filterTags(ctx context.Context, req njudge.ProblemListRequest, pr njudge.Problem) (bool, error) {
	for _, tagName := range req.TagFilter {
		t, err := p.tags.GetByName(ctx, tagName)
		if err != nil {
			return false, nil
		}

		if !pr.HasTag(*t) {
			return false, nil
		}
	}

	return true, nil
}

func (p *ProblemListQuery) filterCategory(ctx context.Context, req njudge.ProblemListRequest, pr njudge.Problem) (bool, error) {
	switch req.CategoryFilter.Type {
	case njudge.CategoryFilterID:
		if pr.Category == nil {
			return false, nil
		}

		categories, err := p.cs.GetAll(ctx)
		if err != nil {
			return false, err
		}

		par := make(map[int]int)
		for ind := range categories {
			if categories[ind].ParentID.Valid {
				par[categories[ind].ID] = categories[ind].ParentID.Int
			}
		}

		curr := pr.Category.ID
		found := false
		for {
			if curr == req.CategoryFilter.Value.(int) {
				found = true
				break
			}

			if _, ok := par[curr]; !ok {
				break
			}
			curr = par[curr]
		}

		return found, nil
	case njudge.CategoryFilterEmpty:
		return pr.Category == nil, nil
	default:
		return true, nil
	}
}

func (p *ProblemListQuery) GetProblemList(ctx context.Context, req njudge.ProblemListRequest) (*njudge.ProblemList, error) {
	allProblems, err := p.ps.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	filters := []func(context.Context, njudge.ProblemListRequest, njudge.Problem) (bool, error){
		p.filterProblemset,
		p.filterTags,
		p.filterTitle,
		p.filterCategory,
	}

	problems := make([]njudge.Problem, 0)
	for ind := range allProblems {
		ok := true
		for _, filter := range filters {
			currOk, err := filter(ctx, req, allProblems[ind])
			if err != nil {
				return nil, err
			}

			ok = ok && currOk
		}

		if ok {
			problems = append(problems, allProblems[ind])
		}
	}

	inv := req.SortDir == njudge.SortDESC
	sort.Slice(problems, func(i, j int) bool {
		switch req.SortField {
		case njudge.ProblemSortFieldSolverCount:
			return inv != (problems[i].SolverCount < problems[j].SolverCount)
		default:
			return inv != (problems[i].ID < problems[j].ID)
		}
	})

	var pdata njudge.PaginationData
	problems, pdata = Paginate(problems, req.Page, req.PerPage)

	return &njudge.ProblemList{
		PaginationData: pdata,
		Problems:       problems,
	}, nil
}
