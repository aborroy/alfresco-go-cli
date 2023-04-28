package node

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const GetNodeCmdId string = "[NODE GET]"

var nodeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Node information",
	Run: func(command *cobra.Command, args []string) {

		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                nodeUrlPath + nodeId,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(GetNodeCmdId, _error)
		}

		var format, _ = command.Flags().GetString("output")
		outputNode(responseBody.Bytes(), format)

		log.Println(GetNodeCmdId, "Details for node "+nodeId+" have been retrieved")

	},
}

func init() {
	nodeCmd.AddCommand(nodeGetCmd)
}
