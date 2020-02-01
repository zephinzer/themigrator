package connection

import "github.com/spf13/cobra"

func AddCobraFlags(command *cobra.Command, connectionOptions *Options) {
	command.PersistentFlags().StringVarP(
		&connectionOptions.User,
		"user",
		"u",
		"user",
		"username of the database user",
	)
	command.PersistentFlags().StringVarP(
		&connectionOptions.Password,
		"password",
		"p",
		"password",
		"password of the database user",
	)
	command.PersistentFlags().StringVarP(
		&connectionOptions.Host,
		"host",
		"H",
		"localhost",
		"hostname of the database service",
	)
	command.PersistentFlags().StringVarP(
		&connectionOptions.Port,
		"port",
		"P",
		"3306",
		"port that the database service is listening on",
	)
	command.PersistentFlags().StringVarP(
		&connectionOptions.Database,
		"database",
		"d",
		"",
		"database table to use by default",
	)
}
