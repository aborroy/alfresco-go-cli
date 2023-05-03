package node

import (
	"bytes"
	"net/url"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

func GetUrlParams(params map[string]string) url.Values {
	var parameters = url.Values{}
	for key, value := range params {
		parameters.Add(key, value)
	}
	return parameters
}

var nodeId string
var relativePath string
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
