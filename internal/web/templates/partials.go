package templates

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/mraron/njudge/internal/njudge/db/models"

	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	CustomHeadPartial   = "custom_head"
	CustomFooterPartial = "custom_footer"
	CustomMenuPartial   = "custom_menu"
)

type Store interface {
	Get(name string) (string, error)
}

type Cached struct {
	DB *sql.DB

	validFor time.Duration
	cache    map[string]string
	accessed map[string]time.Time
}

func NewCached(db *sql.DB, validFor time.Duration) *Cached {
	return &Cached{
		DB:       db,
		validFor: validFor,
		cache:    make(map[string]string),
		accessed: make(map[string]time.Time),
	}
}

func (pc *Cached) Get(name string) (string, error) {
	if time.Since(pc.accessed[name]) > pc.validFor {
		p, err := models.Partials(Where("name = ?", name)).One(context.TODO(), pc.DB)
		if err != nil {
			return "", err
		}

		pc.accessed[name] = time.Now()
		pc.cache[name] = p.HTML
		return p.HTML, nil
	}

	return pc.cache[name], nil
}

type Empty struct{}

func (e Empty) Get(name string) (string, error) {
	return "", fmt.Errorf("no such partial: %s", name)
}
