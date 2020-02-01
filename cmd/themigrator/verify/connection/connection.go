package connection

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
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
			logs <- log.Entry{
				Code: "VERIFY_CONNECTION",
				Message: fmt.Sprintf(
					"connecting as '%s' to '%s:%s/%s' with parameters %v",
					connectionOptions.User,
					connectionOptions.Host,
					connectionOptions.Port,
					connectionOptions.Database,
					connectionOptions.Params,
				),
			}
			go connection.Establish(
				connectionOptions,
				eventStream,
				500*time.Millisecond,
			)
			go func() {
				<-eventStream.Connection
				logs <- log.Entry{
					Code:    common.ErrorOK,
					Message: "connection credentials verified",
				}
				done <- common.ExitCodeOK
			}()
			exitCode := <-done
			<-time.After(time.Second)
			os.Exit(exitCode)
		},
	}
	connection.AddCobraFlags(cmd, &connectionOptions)
	return cmd
}

func handleLogs(eventStream connection.EventStream, logs chan log.Entry) {
	for logEntry := range eventStream.Logs {
		logs <- logEntry
	}
}

func handleErrors(eventStream connection.EventStream, done chan int) {
	for range eventStream.Error {
		done <- common.ExitCodeDatabaseConnectionError
	}
}
