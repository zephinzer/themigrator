package initialise

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
	"gitlab.com/zephinzer/themigrator/lib/migration"
)

const (
	CommandCode = "INIT"
)

func Get(logs chan log.Entry) *cobra.Command {
	var connectionOptions connection.Options
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialises the migration table",
		Run: func(command *cobra.Command, args []string) {
			done := make(chan int)
			eventStream := connection.NewEventStream()
			go handleErrors(eventStream, done)
			go connection.Establish(
				connectionOptions,
				eventStream,
				500*time.Millisecond,
			)
			go func() {
				err := migration.CreateTable(<-eventStream.Connection)
				if err != nil {
					logs <- log.Entry{
						Code:    CommandCode,
						Level:   log.LevelError,
						Message: err.Error(),
					}
					done <- common.ExitCodeCreateMigrationsTableFailed
				} else {
					logs <- log.Entry{
						Code:    CommandCode,
						Level:   log.LevelInfo,
						Message: "migration table created",
					}
					done <- common.ExitCodeOK
				}
			}()
			exitCode := <-done
			<-time.After(time.Second)
			os.Exit(exitCode)
		},
	}
	connection.AddCobraFlags(cmd, &connectionOptions)
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
