package verify

import (
	"github.com/spf13/cobra"
	connectionCmd "gitlab.com/zephinzer/themigrator/cmd/themigrator/verify/connection"
	migrationsCmd "gitlab.com/zephinzer/themigrator/cmd/themigrator/verify/migrations"
	"gitlab.com/zephinzer/themigrator/lib/log"
)

func Get(logs chan log.Entry) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Verifies things",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(connectionCmd.Get(logs))
	cmd.AddCommand(migrationsCmd.Get(logs))
	return cmd
}
