package migrations

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
	"gitlab.com/zephinzer/themigrator/lib/migration"
)

const (
	CommandCode = "verify-migrations"
)

func Get(logs chan log.Entry) *cobra.Command {
	var connectionOptions connection.Options
	cmd := &cobra.Command{
		Use:   "migrations",
		Short: "verifies that migrations locally and remotely are in order",
		Long: strings.Trim(`themigrator verify migrations
		verifies that migrations locally and remotely are in order
		`, " \n\t"),
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

				// verify that everything in the remote can be found locally with its integrity in-tact
				for i := 0; i < len(remoteMigrations); i++ {
					remoteMigration := remoteMigrations[i]
					localMigration := localMigrations[i]
					uuidMatches := remoteMigration.HasSameUUIDAs(localMigration)
					contentHashMatches := remoteMigration.HasSameContentHashAs(localMigration)
					if uuidMatches && contentHashMatches {
						logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("processed %s/%s", remoteMigration.UUID, remoteMigration.ContentHash))
					} else if !uuidMatches && !contentHashMatches {
						migrationFoundElsewhere := false
						for _, anotherLocalMigration := range localMigrations {
							extraCheckUUIDMatches := remoteMigration.HasSameUUIDAs(anotherLocalMigration)
							extraCheckContentHashMatches := remoteMigration.HasSameContentHashAs(anotherLocalMigration)
							if extraCheckUUIDMatches && extraCheckContentHashMatches {
								migrationFoundElsewhere = true
							}
						}
						if migrationFoundElsewhere {
							logs <- log.NewEntry(logger.LevelWarn, CommandCode, fmt.Sprintf("remote migration with ID '%s' and hash '%s' has a local copy but it seems like the order may be wonky", remoteMigration.UUID, remoteMigration.ContentHash))
						} else {
							logs <- log.NewEntry(logger.LevelError, CommandCode, fmt.Sprintf("remote migration with ID '%s' and hash '%s' does not have local copy", remoteMigration.UUID, remoteMigration.ContentHash))
						}
					} else {
						logs <- log.NewEntry(logger.LevelWarn, CommandCode, fmt.Sprintf("uuid/content mismatch between %s/%s and %s/%s", remoteMigration.UUID, remoteMigration.ContentHash, localMigration.UUID, localMigration.ContentHash))
					}
				}

				done <- common.ExitCodeOK
			}()
			exitCode := <-done
			<-time.After(time.Second)
			os.Exit(exitCode)
		},
	}

	connection.AddCobraFlags(connection.AddCobraFlagsOptions{
		Command:           cmd,
		ConnectionOptions: &connectionOptions,
		RequiredFlags:     []string{connection.FlagDatabase.Long},
	})
	return cmd
}

func handleErrors(eventStream connection.EventStream, done chan int) {
	for {
		select {
		case <-eventStream.Error:
			done <- common.ExitCodeDatabaseConnectionError
		}
	}
}
