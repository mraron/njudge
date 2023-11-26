package glue

import (
	"context"
	"database/sql"
	"log"

	"github.com/mraron/njudge/internal/judge"
	"github.com/mraron/njudge/internal/njudge/db/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/multierr"
)

type JudgesUpdater interface {
	UpdateJudges(ctx context.Context) ([]*models.Judge, error)
}

type JudgesUpdaterFromDB struct {
	DB *sql.DB
}

func (o *JudgesUpdaterFromDB) UpdateJudges(ctx context.Context) ([]*models.Judge, error) {
	judges, err := models.Judges().All(ctx, o.DB)
	if err != nil {
		return nil, err
	}

	var judgesError error
	for ind := range judges {
		c := judge.NewClient("http://" + judges[ind].Host + ":" + judges[ind].Port)
		st, err := c.Status(ctx)
		if err != nil {
			judges[ind].Online = false
			judges[ind].Ping = -1

			_, err2 := judges[ind].Update(ctx, o.DB, boil.Infer())
			judgesError = multierr.Combine(judgesError, err, err2)
		}

		judges[ind].Online = true
		judges[ind].State = st.String()
		judges[ind].Ping = 1

		_, err = judges[ind].Update(ctx, o.DB, boil.Infer())
		if err != nil {
			log.Print("trying to access judge on", judges[ind].Host, judges[ind].Port, " unsuccessful update in database", err)
			continue
		}
	}

	return judges, judgesError
}
