package templates

import (
	"database/sql"
	"time"

	"github.com/mraron/njudge/pkg/web/models"

	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PartialCache struct {
	DB *sql.DB

	validFor time.Duration
	cache    map[string]string
	accessed map[string]time.Time
}

func NewPartialCache(db *sql.DB, validFor time.Duration) *PartialCache {
	return &PartialCache{
		DB:       db,
		validFor: validFor,
		cache:    make(map[string]string),
		accessed: make(map[string]time.Time),
	}
}

func (pc *PartialCache) Get(name string) (string, error) {
	if time.Since(pc.accessed[name]) > pc.validFor {
		p, err := models.Partials(Where("name = ?", name)).One(pc.DB)
		if err != nil {
			return "", nil
		}

		pc.accessed[name] = time.Now()
		pc.cache[name] = p.HTML
		return p.HTML, nil
	}

	return pc.cache[name], nil
}
