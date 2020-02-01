package migration

import (
	"database/sql"

	"gitlab.com/zephinzer/themigrator/lib/errors"
)

func CreateTable(connection *sql.DB) error {
	stmt, err := connection.Prepare(`
		CREATE TABLE IF NOT EXISTS themigrations(
			id INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
			content TEXT NOT NULL,
			content_hash VARCHAR(64) NOT NULL,
			created_on TIMESTAMP NOT NULL DEFAULT NOW(),
			applied_on DATETIME
		) ENGINE=InnoDB default charset utf8mb4;
	`)
	if err != nil {
		return errors.New(errors.ErrorDatabaseStatementPrep, err.Error(), nil)
	}
	_, err = stmt.Exec()
	if err != nil {
		return errors.New(errors.ErrorDatabaseOpCreateTable, err.Error(), nil)
	}
	return nil
}
