package migration

import (
	"time"
)

type Migration struct {
	ID          string
	Content     string
	ContentHash string
	CreatedOn   time.Time
	AppliedOn   time.Time
}
