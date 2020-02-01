package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/initialise"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/new"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/plan"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/verify"
	"gitlab.com/zephinzer/themigrator/lib/log"
)

var logs chan log.Entry

func init() {
	logs = make(chan log.Entry, 256)
	themigrator.AddCommand(verify.Get(logs))
	themigrator.AddCommand(initialise.Get(logs))
	themigrator.AddCommand(plan.Get(logs))
	themigrator.AddCommand(new.Get(logs))
}

func main() {
	go common.HandleLogs(logs)
	themigrator.Execute()
}

var themigrator = cobra.Command{
	Use: "themigrator",
	Run: func(command *cobra.Command, args []string) {
		command.Help()
	},
}
