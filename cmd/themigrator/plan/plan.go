package plan

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
	"gitlab.com/zephinzer/themigrator/lib/migration"
)

const (
	CommandCode = "PLAN"
)

func Get(logs chan log.Entry) *cobra.Command {
	var connectionOptions connection.Options
	cmd := &cobra.Command{
		Use:   "plan [PATH TO MIGRATIONS]",
		Short: "Dumps the migraiton plan if the migration is run",
		Run: func(command *cobra.Command, args []string) {
			done := make(chan int)
			eventStream := connection.NewEventStream()
			go handleErrors(eventStream, done)
			go connection.Establish(connectionOptions, eventStream, 500*time.Millisecond)
			go func() {
				pathToMigrations, err := filepath.Abs(path.Join(args...))
				if err != nil {
					logs <- log.NewEntry(logger.LevelError, CommandCode, "failed to get absolute path of migrations", err)
					done <- common.ExitCodeInsufficientPermissions
					return
				}
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, fmt.Sprintf("using '%s' as the migrations directory", pathToMigrations))
				_, err = getLocalMigrations(pathToMigrations, logs)
				if err != nil {
					done <- common.ExitCodeGeneric
					return
				}
				_, err = getRemoteMigrations(<-eventStream.Connection, logs)
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
		logs <- log.NewEntry(logger.LevelDebug, "MIGRATION", fmt.Sprintf("%s: '%s'", remoteMigrations[i].ContentHash, remoteMigrations[i].Content))
	}
	return remoteMigrations, nil
}

func getLocalMigrations(fromPath string, logs chan log.Entry) ([]migration.Migration, error) {
	localMigrations, err := migration.LoadFilesystem(fromPath)
	if err != nil {
		logs <- log.NewEntry(logger.LevelError, CommandCode, "error loading local migrations", err)
		return nil, err
	}
	logs <- log.NewEntry(logger.LevelInfo, CommandCode,
		fmt.Sprintf("found %v local migrations as follows", len(localMigrations)),
	)
	for i := 0; i < len(localMigrations); i++ {
		logs <- log.NewEntry(logger.LevelDebug, "MIGRATION", fmt.Sprintf("%s: '%s'", localMigrations[i].ContentHash, localMigrations[i].Content))
	}
	return localMigrations, nil
}

func handleErrors(eventStream connection.EventStream, done chan int) {
	for {
		select {
		case <-eventStream.Error:
			done <- common.ExitCodeDatabaseConnectionError
		}
	}
}
