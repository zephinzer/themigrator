package connection

import (
	"github.com/spf13/cobra"
)

type Flag struct {
	Long    string
	Short   string
	Default string
}

var (
	FlagUser     = Flag{"user", "u", "user"}
	FlagPassword = Flag{"password", "p", "password"}
	FlagHost     = Flag{"host", "H", "localhost"}
	FlagPort     = Flag{"port", "P", "3306"}
	FlagDatabase = Flag{"database", "d", ""}
)

type AddCobraFlagsOptions struct {
	Command           *cobra.Command
	ConnectionOptions *Options
	RequiredFlags     []string
}

func AddCobraFlags(options AddCobraFlagsOptions) {
	options.Command.Flags().StringVarP(
		&options.ConnectionOptions.User,
		FlagUser.Long,
		FlagUser.Short,
		FlagUser.Default,
		"username of the database user",
	)

	options.Command.Flags().StringVarP(
		&options.ConnectionOptions.Password,
		FlagPassword.Long,
		FlagPassword.Short,
		FlagPassword.Default,
		"password of the database user",
	)
	options.Command.Flags().StringVarP(
		&options.ConnectionOptions.Host,
		FlagHost.Long,
		FlagHost.Short,
		FlagHost.Default,
		"hostname of the database service",
	)
	options.Command.Flags().StringVarP(
		&options.ConnectionOptions.Port,
		FlagPort.Long,
		FlagPort.Short,
		FlagPort.Default,
		"port that the database service is listening on",
	)
	options.Command.Flags().StringVarP(
		&options.ConnectionOptions.Database,
		FlagDatabase.Long,
		FlagDatabase.Short,
		FlagDatabase.Default,
		"database table to use by default",
	)
	for _, flag := range options.RequiredFlags {
		err := options.Command.MarkFlagRequired(flag)
		if err != nil {
			panic(err)
		}
	}
}
