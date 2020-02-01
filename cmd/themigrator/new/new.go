package new

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zephinzer/themigrator/cmd/themigrator/common"
	"gitlab.com/zephinzer/themigrator/lib/connection"
	"gitlab.com/zephinzer/themigrator/lib/log"
	"gitlab.com/zephinzer/themigrator/lib/utils"
)

func Get(logs chan log.Entry) *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Creates a new migration",
		Run: func(command *cobra.Command, args []string) {
			done := make(chan int)
			go func() {
				pathToMigrations, err := filepath.Abs(path.Join(args...))
				if err != nil {
					done <- common.ExitCodeInsufficientPermissions
					return
				}
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("selected directory to place new migration in: '%s'\n", pathToMigrations)
				fmt.Printf("describe your migration: ")
				userInputName, err := reader.ReadString('\n')
				if err != nil {
					done <- common.ExitCodeInvalidUserInput
					return
				}
				migrationName := time.Now().Format("20060102150405")
				providedName := utils.FormatMigrationName(userInputName)
				migrationName = fmt.Sprintf("%s_%s.sql", migrationName, providedName)
				pathToMigration := migrationName
				if len(args) > 0 {
					args = append(args, migrationName)
					pathToMigration = path.Join(args...)
				}
				err = ioutil.WriteFile(pathToMigration, []byte{}, os.ModePerm)
				if err != nil {
					done <- common.ExitCodeErrorCreatingFile
					return
				}
				fmt.Printf("created migration file at: %s\n", pathToMigration)
				done <- common.ExitCodeOK
			}()
			exitCode := <-done
			<-time.After(time.Second)
			os.Exit(exitCode)
		},
	}
}

func handleErrors(eventStream connection.EventStream, done chan int) {
	for {
		select {
		case <-eventStream.Error:
			done <- common.ExitCodeDatabaseConnectionError
		}
	}
}
