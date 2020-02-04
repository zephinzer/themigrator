package apply

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
	"gitlab.com/zephinzer/themigrator/lib/migration"
)

const (
	CommandCode = "apply"
)

func Get(logs chan log.Entry) *cobra.Command {
	var connectionOptions connection.Options
	var confirmApply bool
	var migrateUpTo int
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Applies pending migrations",
		Run: func(command *cobra.Command, args []string) {
			done := make(chan int)
			eventStream := connection.NewEventStream()
			go handleErrors(eventStream, done)
			go connection.Establish(connectionOptions, eventStream, 500*time.Millisecond)
			go func() {
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("validating provided input arguments (%v)", args))
				// check whether provided arguments are a valid absolute file path
				pathToMigrations, err := filepath.Abs(path.Join(args...))
				if err != nil {
					logs <- log.NewEntry(logger.LevelError, CommandCode, "failed to get absolute path of migrations for reasons that appear below this", err)
					done <- common.ExitCodeInsufficientPermissions
					return
				}

				// check whether provided path to migrations is a valid directory
				migrationsDirectoryFileInfo, err := os.Lstat(pathToMigrations)
				if err != nil {
					logs <- log.NewEntry(logger.LevelError, CommandCode, fmt.Sprintf("failed to retrieve information about the migrations directory at '%s', check your permissions", pathToMigrations), err)
					done <- common.ExitCodeInsufficientPermissions
					return
				} else if !migrationsDirectoryFileInfo.IsDir() {
					logs <- log.NewEntry(logger.LevelError, CommandCode, fmt.Sprintf("provided path '%s' should be a directory", pathToMigrations), err)
					done <- common.ExitCodeInvalidUserInput
					return
				}

				// start retrieving local migrations
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("retrieving local migrations from '%s'...", pathToMigrations))
				localMigrations, err := migration.LoadFilesystem(pathToMigrations)
				if err != nil {
					logs <- log.NewEntry(logger.LevelError, CommandCode, fmt.Sprintf("failed to retrieve local migrations from the path '%s'", pathToMigrations), err)
					done <- common.ExitCodeGeneric
					return
				}
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("found %v local migrations", len(localMigrations)))
				for i := 0; i < len(localMigrations); i++ {
					localMigration := localMigrations[i]
					logs <- log.NewEntry(logger.LevelDebug, CommandCode, fmt.Sprintf("%s/%s: '%s'", localMigration.UUID, localMigration.ContentHash, localMigration.Content))
				}

				establishedConnection := <-eventStream.Connection
				// start retrieving remote migrations
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, "retrieving remote migrations from the database...")
				remoteMigrations, err := migration.LoadRemote(establishedConnection)
				if err != nil {
					logs <- log.NewEntry(logger.LevelError, CommandCode, fmt.Sprintf("failed to retrieve remote migrations from the path '%s'", pathToMigrations), err)
					done <- common.ExitCodeGeneric
					return
				}
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("found %v remote migrations", len(remoteMigrations)))
				for i := 0; i < len(remoteMigrations); i++ {
					remoteMigration := remoteMigrations[i]
					logs <- log.NewEntry(logger.LevelDebug, CommandCode, fmt.Sprintf("%s/%s: '%s'", remoteMigration.UUID, remoteMigration.ContentHash, remoteMigration.Content))
				}

				// deduce the migrations that need applying
				applicableMigrations := migration.GetUnappliedFrom(localMigrations, remoteMigrations)
				if len(applicableMigrations) == 0 {
					logs <- log.NewEntry(logger.LevelInfo, CommandCode, "all migrations that could've been applied have been")
					done <- common.ExitCodeOK
					return
				}
				sort.Sort(applicableMigrations)
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("%v migrations have yet to be applied", len(applicableMigrations)))
				for i := 0; i < len(applicableMigrations); i++ {
					applicableMigration := applicableMigrations[i]
					logs <- log.NewEntry(logger.LevelDebug, CommandCode, fmt.Sprintf("%s/%s: '%s'", applicableMigration.UUID, applicableMigration.ContentHash, applicableMigration.Content))
				}

				if !confirmApply {
					logs <- log.NewEntry(logger.LevelWarn, CommandCode, "refusing to apply because --confirm was not specified")
					done <- common.ExitCodeSavedYourAss
					return
				}

				logs <- log.NewEntry(logger.LevelInfo, CommandCode, "crossing fingers and applying the migrations now...")
				if migrateUpTo < 0 {
					migrateUpTo = len(applicableMigrations)
				}
				for i := 0; i < migrateUpTo; i++ {
					applicableMigration := applicableMigrations[i]
					if err = applicableMigration.Apply(establishedConnection); err != nil {
						logs <- log.NewEntry(logger.LevelError, CommandCode, fmt.Sprintf("could not apply migration '%s'", applicableMigration.UUID), err)
					} else {
						logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("successfully applied migration '%s'", applicableMigration.UUID), err)
					}
				}
				done <- common.ExitCodeOK
			}()
			exitCode := <-done
			<-time.After(time.Second)
			os.Exit(exitCode)
		},
	}
	cmd.Flags().IntVarP(&migrateUpTo, "steps", "S", -1, "when 0 or positive, this will be how many migrations to run")
	cmd.Flags().BoolVarP(&confirmApply, "confirm", "y", false, "when specified, applies the migrations as planned")
	connection.AddCobraFlags(connection.AddCobraFlagsOptions{
		Command:           cmd,
		ConnectionOptions: &connectionOptions,
		RequiredFlags:     []string{connection.FlagDatabase.Long},
	})
	return cmd
}

func getRemoteMigrations(dbConnection *sql.DB, logs chan log.Entry) ([]migration.Migration, error) {
	remoteMigrations, err := migration.LoadRemote(dbConnection)
	if err != nil {
		logs <- log.NewEntry(logger.LevelError, CommandCode, "error loading remote migrations", err)
		return nil, err
	}
	logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("found %v remote migrations as follows", len(remoteMigrations)))
	for i := 0; i < len(remoteMigrations); i++ {
		logs <- log.NewEntry(logger.LevelDebug, CommandCode, fmt.Sprintf("%s: '%s'", remoteMigrations[i].ContentHash, remoteMigrations[i].Content))
	}
	return remoteMigrations, nil
}

func handleErrors(eventStream connection.EventStream, done chan int) {
	for {
		select {
		case <-eventStream.Error:
			done <- common.ExitCodeDatabaseConnectionError
		}
	}
}
