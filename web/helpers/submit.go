package helpers

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/web/extmodels"
	"github.com/mraron/njudge/web/helpers/config"
	"time"
)

func Submit(cfg config.Server, DB *sqlx.DB, problemStore problems.Store, uid int, problemset, problem, language string, source []byte) (int, error) {
	var (
		tx  *sql.Tx
		id  int
		err error
	)
	mustPanic := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	transaction := func() {
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				err = p.(error)
			}
		}()

		tx, err = DB.Begin()
		mustPanic(err)

		id = 0
		res, err := tx.Query("INSERT INTO submissions (status,\"user_id\",verdict,ontest,submitted,judged,problem,language,private,problemset,source,started) VALUES ($1,$2,$3,NULL,$4,NULL,$5,$6,false,$7, $8,false) RETURNING id", problems.Status{}, uid, extmodels.VERDICT_UP, time.Now(), problemStore.MustGet(problem).Name(), language, problemset, source)

		mustPanic(err)

		res.Next()

		err = res.Scan(&id)
		mustPanic(err)

		err = res.Close()
		mustPanic(err)

		err = tx.Commit()
		mustPanic(err)
	}

	if transaction(); err != nil {
		return -1, err
	}

	return id, nil
}
