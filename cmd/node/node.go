package node

import (
	"bytes"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

var nodeId string
var responseBody bytes.Buffer
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage nodes",
	PersistentPostRun: func(command *cobra.Command, args []string) {
		var format, _ = command.Flags().GetString("output")
		output(responseBody.Bytes(), format)
	},
}

func init() {
	cmd.RootCmd.AddCommand(nodeCmd)
}
