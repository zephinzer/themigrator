package migration

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"

	"gitlab.com/zephinzer/themigrator/lib/errors"
	"gitlab.com/zephinzer/themigrator/lib/utils"
)

func LoadFilesystem(filepath string) ([]Migration, error) {
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, errors.New(
			errors.ErrorFilesystemListDirectory,
			fmt.Sprintf("unable to read directory listing at '%s'", filepath),
		)
	}
	format, err := regexp.Compile(`[\d]{14}_[a-zA-Z0-9\-_\.]+.sql`)
	if err != nil {
		return nil, errors.New(
			errors.ErrorRegexp,
			"regex issue! why you no check?",
		)
	}
	var migrations []Migration
	for _, file := range files {
		if format.MatchString(file.Name()) {
			fullFilePath := path.Join(filepath, file.Name())
			fileContent, err := ioutil.ReadFile(fullFilePath)
			if err != nil {
				return nil, errors.New(
					errors.FilesystemReadFile,
					fmt.Sprintf("unable to read contents of file at '%s'", fullFilePath),
				)
			}
			content := utils.CompressWhitespace(string(fileContent))
			contentHash := utils.Hash(content)
			migrations = append(migrations, Migration{
				Content:     content,
				ContentHash: contentHash,
			})
		}
	}
	return migrations, nil
}

func LoadRemote(connection *sql.DB) ([]Migration, error) {
	stmt, err := connection.Prepare(`
		SELECT 
			id,
			content,
			content_hash,
			created_on,
			applied_on
			FROM themigrations
	`)
	if err != nil {
		return nil, errors.New(
			errors.ErrorDatabaseStatementPrep,
			err.Error(),
		)
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New(
			errors.ErrorDatabaseOpQuery,
			err.Error(),
		)
	}
	var results []Migration
	for rows.Next() {
		var migration Migration
		if err = rows.Scan(
			&migration.ID,
			&migration.Content,
			&migration.ContentHash,
			&migration.CreatedOn,
			&migration.AppliedOn,
		); err != nil {
			return nil, errors.New(
				errors.ErrorDatabaseResultRetrieval,
				err.Error(),
			)
		}
		results = append(results, migration)
	}
	return results, nil
}
