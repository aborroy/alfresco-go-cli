package people

import (
	"bytes"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

var responseBody bytes.Buffer
var personId string
var peopleCmd = &cobra.Command{
	Use:   "people",
	Short: "Manage people in ACS Repository",
	Long:  `ACS Repository handles a set of Persons that may be associated to groups, permissions and roles.`,
	PersistentPostRun: func(command *cobra.Command, args []string) {
		var format, _ = command.Flags().GetString("output")
		output(responseBody.Bytes(), format)
	},
}

func init() {
	cmd.RootCmd.AddCommand(peopleCmd)
}
