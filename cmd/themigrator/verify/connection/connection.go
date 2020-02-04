package connection

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
)

const (
	CommandCode = "verify-connection"
)

func Get(logs chan log.Entry) *cobra.Command {
	var connectionOptions connection.Options
	cmd := &cobra.Command{
		Use:   "connection",
		Short: "Verifies the connection configuration",
		Run: func(command *cobra.Command, args []string) {
			done := make(chan int)
			eventStream := connection.NewEventStream()
			go handleErrors(eventStream, done)
			go handleLogs(eventStream, logs)
			logs <- log.NewEntry(logger.LevelDebug, CommandCode, fmt.Sprintf(
				"connecting as '%s' to '%s:%s/%s' with parameters %v",
				connectionOptions.User,
				connectionOptions.Host,
				connectionOptions.Port,
				connectionOptions.Database,
				connectionOptions.Params,
			))
			go connection.Establish(
				connectionOptions,
				eventStream,
				500*time.Millisecond,
			)
			go func() {
				<-eventStream.Connection
				logs <- log.NewEntry(logger.LevelInfo, CommandCode, "connection credentials verified")
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

func handleLogs(eventStream connection.EventStream, logs chan log.Entry) {
	for logEntry := range eventStream.Logs {
		logs <- log.NewEntry(logEntry.GetLevel(), CommandCode, logEntry.Error())
	}
}

func handleErrors(eventStream connection.EventStream, done chan int) {
	for range eventStream.Error {
		done <- common.ExitCodeDatabaseConnectionError
	}
}
