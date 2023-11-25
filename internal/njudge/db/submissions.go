package db

import (
	"context"
	"database/sql"

	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
)

type Submissions struct {
	db *sql.DB
}

func NewSubmissions(db *sql.DB) *Submissions {
	return &Submissions{
		db: db,
	}
}

func (ss *Submissions) toNjudge(ctx context.Context, s *models.Submission) (*njudge.Submission, error) {
	panic("not impl")
	/*res := njudge.Submission{
		ID: s.ID,

		UserID:   s.UserID,
		Language: s.Language,
		Source:   s.Source,
		Private:  s.Private,

		Started:   s.Started,
		Verdict:   njudge.Verdict(s.Verdict),
		Ontest:    s.Ontest,
		Submitted: s.Submitted,
		Judged:    s.Judged,
		Score:     s.Score.Float32,
	}
	*/

}

func (ss *Submissions) Get(ctx context.Context, ID int) (*njudge.Submission, error) {
	panic("not implemented") // TODO: Implement
}

func (ss *Submissions) GetAll(ctx context.Context) ([]njudge.Submission, error) {
	panic("not implemented") // TODO: Implement
}

func (ss *Submissions) Insert(ctx context.Context, s njudge.Submission) (*njudge.Submission, error) {
	panic("not implemented") // TODO: Implement
}

func (ss *Submissions) Delete(ctx context.Context, ID int) error {
	panic("not implemented") // TODO: Implement
}

func (ss *Submissions) Update(ctx context.Context, s njudge.Submission, fields []string) error {
	panic("not implemented") // TODO: Implement
}
