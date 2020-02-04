package migration

import (
	"strings"
	"time"
)

type Migration struct {
	ID          string
	Content     string
	ContentHash string
	CreatedOn   time.Time
	AppliedOn   time.Time
}

// Migrations is an instance of an array of Migration instances
// that can be sorted
type Migrations []Migration

func (m Migrations) Len() int {
	return len(m)
}

func (m Migrations) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Migrations) Less(i, j int) bool {
	return strings.Compare(m[i].ID, m[j].ID) < 0
}
