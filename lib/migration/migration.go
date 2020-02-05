package migration

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Migration struct {
	ID          int
	UUID        string
	Content     string
	ContentHash string
	Warning     error
	CreatedOn   time.Time
	AppliedOn   time.Time
}

func (m Migration) HasSameUUIDAs(n Migration) bool {
	return strings.Compare(m.UUID, n.UUID) == 0
}

func (m Migration) HasSameContentHashAs(n Migration) bool {
	return strings.Compare(m.ContentHash, n.ContentHash) == 0
}

func (m Migration) IsEqual(n Migration) bool {
	return m.HasSameUUIDAs(n) && m.HasSameContentHashAs(n)
}

func (m Migration) Apply(connections *sql.DB) error {
	migrationTx, err := connections.Begin()
	if err != nil {
		return err
	}
	stmt, err := migrationTx.Prepare(m.Content)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(); err != nil {
		return err
	}
	if err = migrationTx.Commit(); err != nil {
		return err
	}

	recordingMigrationTx, err := connections.Begin()
	if err != nil {
		return err
	}
	victoryStmt, err := recordingMigrationTx.Prepare(fmt.Sprintf(`
		INSERT INTO %s (uuid, content, content_hash) VALUES (?, ?, ?)
	`, TableName))
	if err != nil {
		return err
	}
	if _, err = victoryStmt.Exec(m.UUID, m.Content, m.ContentHash); err != nil {
		return err
	}
	if err = recordingMigrationTx.Commit(); err != nil {
		return err
	}
	return nil
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
	return strings.Compare(m[i].UUID, m[j].UUID) < 0
}
