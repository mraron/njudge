package web

import (
	"context"
	"database/sql"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/cached"
	"github.com/mraron/njudge/internal/njudge/db"
	"github.com/mraron/njudge/internal/njudge/email"
	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/mraron/njudge/internal/web/templates"
	"github.com/mraron/njudge/pkg/language/langs/cpp"
	"github.com/mraron/njudge/pkg/problems"
	"time"
)

// DataAccess provides access to the business logic in handlers
type DataAccess struct {
	ProblemStore  problems.Store
	MailService   email.Service
	PartialsStore templates.Store

	Categories          njudge.Categories
	Tags                njudge.Tags
	Problems            njudge.Problems
	Users               njudge.Users
	Submissions         njudge.Submissions
	SolvedStatusQuery   njudge.SolvedStatusQuery
	ProblemInfoQuery    njudge.ProblemInfoQuery
	ProblemQuery        njudge.ProblemQuery
	ProblemListQuery    njudge.ProblemListQuery
	SubmissionListQuery njudge.SubmissionListQuery

	SubmitService      *njudge.SubmitService
	TagsService        njudge.TagsService
	TaskArchiveService njudge.TaskArchiveService
}

// NewDemoDataAccess creates in-memory (internal/njudge/memory) "demo" data
func NewDemoDataAccess(ctx context.Context, ps problems.Store, ms email.Service) (*DataAccess, error) {
	s, err := &DataAccess{}, error(nil)
	s.ProblemStore = ps
	s.MailService = ms
	if err = s.ProblemStore.UpdateProblems(); err != nil {
		return nil, err
	}
	s.PartialsStore = templates.Empty{}

	s.Categories = memory.NewCategories()
	s.Tags = memory.NewTags()
	s.Problems = memory.NewProblems()
	s.Submissions = memory.NewSubmissions()
	s.Users = memory.NewUsers()

	s.ProblemQuery = memory.NewProblemQuery(s.Problems)
	s.ProblemInfoQuery = memory.NewProblemInfoQuery(s.Submissions)
	s.ProblemListQuery = memory.NewProblemListQuery(s.ProblemStore, s.Problems, s.Tags, s.Categories)
	s.SubmissionListQuery = memory.NewSubmissionListQuery(s.Submissions, s.Problems)

	s.SubmitService = njudge.NewSubmitService(s.Users, s.ProblemQuery, s.ProblemStore)
	s.TagsService = memory.NewTagsService(s.Tags, s.Problems, s.ProblemInfoQuery)

	// Create dummy category NT1/2021
	nt1 := njudge.NewCategory("NT1", nil)
	nt1, err = s.Categories.Insert(ctx, *nt1)
	if err != nil {
		return nil, err
	}
	nt1_2021 := njudge.NewCategory("2021", nt1)
	nt1_2021, err = s.Categories.Insert(ctx, *nt1_2021)
	if err != nil {
		return nil, err
	}

	// Create two tags, constructive and dp
	t := njudge.NewTag("constructive")
	t, err = s.Tags.Insert(ctx, *t)
	if err != nil {
		return nil, err
	}

	t2 := njudge.NewTag("dp")
	_, err = s.Tags.Insert(ctx, *t2)
	if err != nil {
		return nil, err
	}

	// Create NT21_Atvagas in problemset main, setting its category
	p := njudge.NewProblem("main", "NT21_Atvagas")
	err = p.AddTag(*t, 1)
	if err != nil {
		return nil, err
	}
	p.SetCategory(*nt1_2021)

	// Create a user and activating
	u, err := njudge.NewUser("mraron", "email@email.com", "admin")
	if err != nil {
		return nil, err
	}
	err = u.SetPassword("abc")
	if err != nil {
		return nil, err
	}
	u.Activate()
	u, err = s.Users.Insert(ctx, *u)
	if err != nil {
		return nil, err
	}

	// Creating another problem is1
	_, err = s.Problems.Insert(ctx, njudge.NewProblem("main", "is1"))
	if err != nil {
		return nil, err
	}
	prob, err := s.Problems.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	// Creating a (AC) submission for it
	sub, err := njudge.NewSubmission(*u, *prob, cpp.Std14)
	if err != nil {
		return nil, err
	}
	sub.SetSource([]byte("#include<bits/stdc++.h>"))
	sub.Verdict = njudge.VerdictAC
	storedData, err := prob.WithStoredData(s.ProblemStore)
	if err != nil {
		return nil, err
	}
	ss, err := storedData.StatusSkeleton("")
	if err != nil {
		return nil, err
	}
	ss.Compiled = true
	sub.Status = *ss
	_, err = s.Submissions.Insert(ctx, *sub)
	if err != nil {
		return nil, err
	}

	s.TaskArchiveService = njudge.TaskArchiveService{
		Categories:        s.Categories,
		Problems:          s.Problems,
		SolvedStatusQuery: s.SolvedStatusQuery,
		ProblemQuery:      s.ProblemQuery,
		ProblemStore:      s.ProblemStore,
	}

	return s, nil
}

// NewDBDataAccess creates a DataAccess backed by "database"-kind business logic (internal/njudge/db)
func NewDBDataAccess(ctx context.Context, ps problems.Store, DB *sql.DB, ms email.Service) (*DataAccess, error) {
	s, err := &DataAccess{}, error(nil)
	s.ProblemStore = ps
	s.MailService = ms
	if err = s.ProblemStore.UpdateProblems(); err != nil {
		return nil, err
	}

	s.PartialsStore = templates.NewCached(DB, 1*time.Minute)

	s.Categories = db.NewCategories(DB)
	s.Tags = db.NewTags(DB)
	s.SolvedStatusQuery = cached.NewSolvedStatusQuery(db.NewSolvedStatusQuery(DB), 30*time.Second)
	s.Problems = db.NewProblems(
		DB,
		s.SolvedStatusQuery,
	)
	s.Submissions = db.NewSubmissions(DB)
	s.Users = db.NewUsers(DB)

	s.ProblemQuery = s.Problems.(*db.Problems)
	s.ProblemInfoQuery = s.Problems.(*db.Problems)
	s.ProblemListQuery = memory.NewProblemListQuery(s.ProblemStore, s.Problems, s.Tags, s.Categories)
	s.SubmissionListQuery = db.NewSubmissionListQuery(DB)

	s.SubmitService = njudge.NewSubmitService(s.Users, s.ProblemQuery, s.ProblemStore)
	s.TagsService = memory.NewTagsService(s.Tags, s.Problems, s.ProblemInfoQuery)

	s.TaskArchiveService = njudge.TaskArchiveService{
		Categories:        s.Categories,
		Problems:          s.Problems,
		SolvedStatusQuery: s.SolvedStatusQuery,
		ProblemQuery:      s.ProblemQuery,
		ProblemStore:      s.ProblemStore,
	}

	return s, nil
}
