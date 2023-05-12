package group

import (
	"bytes"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

var responseBody bytes.Buffer
var groupId string
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage groups in ACS Repository",
	Long:  `ACS Repository handles a set of Groups that main contain another groups or persons.`,
	PersistentPostRun: func(command *cobra.Command, args []string) {
		var format, _ = command.Flags().GetString("output")
		output(responseBody.Bytes(), format)
	},
}

func init() {
	cmd.RootCmd.AddCommand(groupCmd)
}
