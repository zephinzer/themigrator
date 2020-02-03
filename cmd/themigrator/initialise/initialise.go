package initialise

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
	"gitlab.com/zephinzer/themigrator/lib/migration"
)

const (
	CommandCode = "initialise"
)

var (
	done        chan int
	eventStream connection.EventStream
)

// Get retrieves the `init` command
func Get(logs chan log.Entry) *cobra.Command {
	var connectionOptions connection.Options
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialises the migration table",
		Run: func(command *cobra.Command, args []string) {
			done = make(chan int)
			eventStream = connection.NewEventStream()
			go handleErrors(logs, eventStream, done)
			go handleConnectionLogs(logs, eventStream)
			go connection.Establish(connectionOptions, eventStream, 500*time.Millisecond)
			go waitForConnectionThenRun(logs)
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

func waitForConnectionThenRun(logs chan log.Entry) {
	conn := <-eventStream.Connection

	logs <- log.NewEntry(logger.LevelDebug, CommandCode, "checking if migration table exists...")
	err := migration.IsTableCreated(conn)
	if err == nil {
		logs <- log.NewEntry(logger.LevelInfo, CommandCode, "migration table already exists")
		done <- common.ExitCodeOK
		return
	}

	logs <- log.NewEntry(logger.LevelDebug, CommandCode, "creating migration table...")
	err = migration.CreateTable(conn)
	if err != nil {
		logs <- log.NewEntry(logger.LevelError, CommandCode, "creation of migration table failed", err)
		done <- common.ExitCodeCreateMigrationsTableFailed
		return
	}

	logs <- log.NewEntry(logger.LevelInfo, CommandCode, "migration table successfully created")
	done <- common.ExitCodeOK
}

func handleErrors(logs chan log.Entry, eventStream connection.EventStream, done chan int) {
	for {
		select {
		case err := <-eventStream.Error:
			logs <- log.NewEntry(logger.LevelError, CommandCode, "connection failed", err)
			done <- common.ExitCodeDatabaseConnectionError
		}
	}
}

func handleConnectionLogs(logs chan log.Entry, eventStream connection.EventStream) {
	for {
		select {
		case log := <-eventStream.Logs:
			logs <- log
		}
	}
}
