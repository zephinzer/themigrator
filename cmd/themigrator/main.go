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

// versioning stuff
var (
	Commit    string
	Version   string
	Timestamp string
)

// logging stuff
var (
	loggerConfig   logger.Config
	loggerInstance logger.Logger
	logEntries     chan log.Entry
	logOptions     log.Options
)

func init() {
	// setup the logs channel that will store all logs for processing
	logEntries = make(chan log.Entry, 256)

	// add the sub-commands of the `themigrator` command
	themigrator.AddCommand(verify.Get(logEntries))
	themigrator.AddCommand(initialise.Get(logEntries))
	themigrator.AddCommand(apply.Get(logEntries))
	themigrator.AddCommand(new.Get(logEntries))

	themigrator.PersistentFlags().IntVarP(&logOptions.Level, "log-level", "l", 0, "specifies the log level from 0-5 (from {silence, error, warning, info, debug, trace})")
	themigrator.PersistentFlags().StringVarP(&logOptions.Format, "log-format", "f", "text", "specifies the log format (from {json, text})")
}

func main() {
	themigrator.Execute()
}

var themigrator = cobra.Command{
	Use:   "themigrator",
	Short: "the migrator",
	Long: strings.Trim(`themigrator
	When the only way to go is up!
	`, " \n\t"),
	PersistentPreRun: func(command *cobra.Command, args []string) {
		var logLevel logger.Level
		if logOptions.Level > 0 && logOptions.Level <= 5 {
			logLevel = log.Level[logOptions.Level]
		}
		var logFormat logger.Format = logger.FormatText
		if logOptions.Format == "json" {
			logFormat = log.Format[logOptions.Format]
		}
		// setup the logger instance that will process all logs
		loggerConfig = logger.Config{
			Level:  logLevel,
			Format: logFormat,
		}
		if logOptions.Level != 0 {
			loggerInstance = logger.New(loggerConfig)
		}
		go log.Handle(loggerInstance, logEntries)
	},
	Run: func(command *cobra.Command, args []string) {
		command.Help()
	},
}
