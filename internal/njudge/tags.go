package njudge

import (
	"time"
)

type Tag struct {
	ID   int
	Name string
}

type ProblemTag struct {
	ID        int
	ProblemID int
	Tag       Tag
	UserID    int
	Added     time.Time
}
