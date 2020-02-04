package migration

// GetUnappliedFrom returns a slice of Migration objects that represent
// migrations which are in the :source but not in the :applied slice
func GetUnappliedFrom(source, applied []Migration) Migrations {
	var migrationsToApply Migrations
	for i := 0; i < len(source); i++ {
		localMigrationAlreadyApplied := false
		currentLocalMigration := source[i]
		for _, currentRemoteMigration := range applied {
			if currentLocalMigration.ContentHash == currentRemoteMigration.ContentHash {
				localMigrationAlreadyApplied = true
			}
		}
		if !localMigrationAlreadyApplied {
			migrationsToApply = append(migrationsToApply, currentLocalMigration)
		}
	}
	return migrationsToApply
}
