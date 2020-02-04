package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/apply"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/initialise"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/new"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/verify"
	"gitlab.com/zephinzer/themigrator/lib/log"
)

var loggerInstance logger.Logger
var logEntries chan log.Entry

func init() {
	// setup the logger instance that will process all logs
	loggerInstance = logger.New(logger.Config{
		Level: logger.LevelTrace,
	})

	// setup the logs channel that will store all logs for processing
	logEntries = make(chan log.Entry, 256)

	// add the sub-commands of the `themigrator` command
	themigrator.AddCommand(verify.Get(logEntries))
	themigrator.AddCommand(initialise.Get(logEntries))
	themigrator.AddCommand(apply.Get(logEntries))
	themigrator.AddCommand(new.Get(logEntries))
}

func main() {
	go log.Handle(loggerInstance, logEntries)
	themigrator.Execute()
}

var themigrator = cobra.Command{
	Use:   "themigrator",
	Short: "the migrator",
	Long: strings.Trim(`themigrator
	When the only way to go is up!
	`, " \n\t"),
	Run: func(command *cobra.Command, args []string) {
		command.Help()
	},
}
