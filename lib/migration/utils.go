package migration

import "gitlab.com/zephinzer/themigrator/lib/errors"

// GetUnappliedFrom returns a slice of Migration objects that represent
// migrations which are in the :source but not in the :applied slice
func GetUnappliedFrom(source, applied []Migration) Migrations {
	var migrationsToApply Migrations
	for i := 0; i < len(source); i++ {
		localMigrationAlreadyApplied := false
		currentLocalMigration := source[i]
		for _, currentRemoteMigration := range applied {
			if currentLocalMigration.UUID == currentRemoteMigration.UUID {
				if currentLocalMigration.ContentHash != currentRemoteMigration.ContentHash {
					currentLocalMigration.Warning = errors.New(ErrorIDMatchHashMismatch, "UUID matches but content hashes do not")
				} else {
					localMigrationAlreadyApplied = true
				}
			} else if currentLocalMigration.ContentHash == currentRemoteMigration.ContentHash &&
				currentLocalMigration.UUID != currentRemoteMigration.UUID {
				localMigrationAlreadyApplied = true
				currentLocalMigration.Warning = errors.New(ErrorIDMatchHashMismatch, "content hash matches but UUID does not")
			}
		}
		if !localMigrationAlreadyApplied {
			migrationsToApply = append(migrationsToApply, currentLocalMigration)
		}
	}
	return migrationsToApply
}
