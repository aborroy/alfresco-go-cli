package node

import (
	"log"
	"net/http"
	"net/url"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const NodeChildrenCmdId string = "[NODE LIST]"

var skipCount string
var maxItems string
var nodeChildrenCmd = &cobra.Command{
	Use:   "list",
	Short: "Get children nodes",
	Run: func(command *cobra.Command, args []string) {

		params := url.Values{}
		params.Add("skipCount", skipCount)
		params.Add("maxItems", maxItems)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                nodeUrlPath + nodeId + "/children",
			Parameters:         params,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(NodeChildrenCmdId, _error)
		}

		var format, _ = command.Flags().GetString("output")
		outputNodeList(responseBody.Bytes(), format)

		log.Println(NodeChildrenCmdId, "Details for children nodes of "+nodeId+" have been retrieved")

	},
}

func init() {
	nodeCmd.AddCommand(nodeChildrenCmd)
	nodeChildrenCmd.Flags().StringVar(&skipCount, "skipCount", "0", "Skip a number of initial nodes from the list")
	nodeChildrenCmd.Flags().StringVar(&maxItems, "maxItems", "100", "Number of nodes in the response list (max. 1000)")
}
