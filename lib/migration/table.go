package migration

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/zephinzer/themigrator/lib/errors"
)

const (
	TableName = "themigrations"
)

func CreateTable(connection *sql.DB) error {
	stmt, err := connection.Prepare(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
			id INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
			content TEXT NOT NULL,
			content_hash VARCHAR(64) NOT NULL,
			created_on TIMESTAMP NOT NULL DEFAULT NOW(),
			applied_on DATETIME
		) ENGINE=InnoDB default charset utf8mb4;
	`, TableName))
	if err != nil {
		return errors.New(errors.ErrorDatabaseStatementPrep, err.Error())
	}
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		if r := recover(); r != nil {
			cancel()
		}
	}()
	_, err = stmt.ExecContext(timeoutContext)
	if err != nil {
		return errors.New(errors.ErrorDatabaseOpCreateTable, err.Error())
	}
	return nil
}

func IsTableCreated(connection *sql.DB) error {
	stmt, err := connection.Prepare(fmt.Sprintf("SELECT * FROM %s", TableName))
	if err != nil {
		return errors.New(errors.ErrorDatabaseStatementPrep, err.Error())
	}
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		if r := recover(); r != nil {
			cancel()
		}
	}()
	_, err = stmt.ExecContext(timeoutContext)
	if err != nil {
		return errors.New(errors.ErrorDatabaseOpQuery, err.Error())
	}
	return nil
}
