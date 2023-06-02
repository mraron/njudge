package services

import (
	"context"
	"database/sql"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/pagination"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"net/url"
	"strings"
	"unicode"
)

type ProblemListRequest struct {
	Problemset     string
	Pagination     pagination.Data
	TitleFilter    string
	TagFilter      []string
	CategoryFilter problem.CategoryFilter
	GETData        url.Values
}

func (r ProblemListRequest) IsFiltered() bool {
	return r.TitleFilter != "" || len(r.TagFilter) > 0 || r.CategoryFilter.Type != problem.CategoryFilterNone
}

type ProblemListService interface {
	GetProblemList(ctx context.Context, request ProblemListRequest) (*problem.PaginationList, error)
}

type SQLProblemListService struct {
	db                *sql.DB
	ps                problems.Store
	problemRepository problem.Repository
}

func NewSQLProblemListService(db *sql.DB, ps problems.Store, repo problem.Repository) *SQLProblemListService {
	return &SQLProblemListService{db, ps, repo}
}

func (s SQLProblemListService) GetProblemList(ctx context.Context, request ProblemListRequest) (*problem.PaginationList, error) {
	order := qm.OrderBy(request.Pagination.SortField + " " + request.Pagination.SortDir)

	var queries []qm.QueryMod
	if request.Problemset != "" {
		queries = append(queries, models.ProblemRelWhere.Problemset.EQ(request.Problemset))
	}
	if request.TitleFilter != "" {
		problemList, err := s.problemRepository.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		var nameList []interface{}
		for ind := range problemList {
			curr := strings.ToLower(i18n.TranslateContent("hungarian", problemList[ind].Titles()).String())
			want := strings.ToLower(request.TitleFilter)

			t := transform.Chain(
				norm.NFD,
				runes.Remove(runes.In(unicode.Mn)),
				norm.NFC,
			)

			a, _, err := transform.String(t, curr)
			if err != nil {
				return nil, err
			}
			b, _, err := transform.String(t, want)
			if err != nil {
				return nil, err
			}

			if strings.Contains(a, b) {
				nameList = append(nameList, problemList[ind].ProblemRel.Problem)
			}
		}

		queries = append(queries, qm.WhereIn("problem in ?", nameList...))
	}
	if request.CategoryFilter.Type != problem.CategoryFilterNone {
		switch request.CategoryFilter.Type {
		case problem.CategoryFilterID:
			categories, err := models.ProblemCategories().All(ctx, s.db)
			if err != nil {
				return nil, err
			}

			par := make(map[int]int)
			for ind := range categories {
				if categories[ind].ParentID.Valid {
					par[categories[ind].ID] = categories[ind].ParentID.Int
				}
			}

			wantID := request.CategoryFilter.Value.(int)
			var pars []interface{}
			for ind := range categories {
				currID, found := categories[ind].ID, false
				for {
					if currID == wantID {
						found = true
						break
					}

					if _, ok := par[currID]; ok {
						currID = par[currID]
					} else {
						break
					}
				}
				if found {
					pars = append(pars, categories[ind].ID)
				}
			}
			queries = append(queries, qm.WhereIn("category_id IN ?", pars...))
		case problem.CategoryFilterEmpty:
			queries = append(queries, models.ProblemRelWhere.CategoryID.IsNull())
		}
	}
	if request.TagFilter != nil {
		var tagList []interface{}
		for ind := range request.TagFilter {
			tagList = append(tagList, request.TagFilter[ind])
		}

		problemRelList, err := models.ProblemRels(qm.InnerJoin("problem_tags pt on pt.problem_id = problem_rels.id"),
			qm.InnerJoin("tags t on pt.tag_id = t.id"), qm.WhereIn("t.name in ?", tagList...),
			qm.GroupBy("problem_rels.id"), qm.Having("COUNT(DISTINCT t.name) = ?", len(tagList)),
		).All(ctx, s.db)
		if err != nil {
			return nil, err
		}

		var problemIDList []interface{}
		for ind := range problemRelList {
			problemIDList = append(problemIDList, problemRelList[ind].ID)
		}

		queries = append(queries, qm.WhereIn("id IN ?", problemIDList...))
	}

	problemRelList, err := models.ProblemRels(append(append(
		[]qm.QueryMod{
			qm.Limit(request.Pagination.PerPage),
			qm.Offset((request.Pagination.Page - 1) * request.Pagination.PerPage),
		}, queries...), order)...,
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	var problemList []problem.Problem
	for ind := range problemRelList {
		p, err := s.ps.Get(problemRelList[ind].Problem)
		if err != nil {
			return nil, err
		}

		problemList = append(problemList, problem.Problem{
			Problem:    p,
			ProblemRel: *problemRelList[ind],
		})
	}

	cnt, err := models.ProblemRels(queries...).Count(ctx, s.db)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Links(request.Pagination.Page, request.Pagination.PerPage, cnt, request.GETData)
	if err != nil {
		return nil, err
	}

	return &problem.PaginationList{
		Pages:    pages,
		Problems: problemList,
	}, nil
}
