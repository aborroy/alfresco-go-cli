package node

import (
	"bytes"
	"encoding/json"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

func GetNodeId(rootNodeId string, relativePath string) string {
	var responseBodyNodeId bytes.Buffer
	GetNodeProperties(rootNodeId, relativePath, &responseBodyNodeId)
	var node Node
	json.Unmarshal(responseBodyNodeId.Bytes(), &node)
	return node.Entry.ID
}

var nodeId string
var relativePath string
var responseBody bytes.Buffer
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage nodes in ACS Repository",
	Long: `ACS Repository handles a set of Nodes of different types (folders, files...)
This command provides the ability to create, update, retrieve and delete ACS Nodes.`,
	PersistentPostRun: func(command *cobra.Command, args []string) {
		var format, _ = command.Flags().GetString("output")
		output(responseBody.Bytes(), format)
	},
}

func init() {
	cmd.RootCmd.AddCommand(nodeCmd)
}
