package node

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

var skipCount string
var maxItems string
var nodeChildrenCmd = &cobra.Command{
	Use:   "list",
	Short: "Get children nodes",
	Run: func(cmd *cobra.Command, args []string) {

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
			fmt.Println(_error)
			os.Exit(1)
		}

		var format, _ = cmd.Flags().GetString("output")
		outputNodeList(responseBody.Bytes(), format)

	},
}

func init() {
	nodeCmd.AddCommand(nodeChildrenCmd)
	nodeChildrenCmd.Flags().StringVar(&skipCount, "skipCount", "0", "Skip a number of initial nodes from the list")
	nodeChildrenCmd.Flags().StringVar(&maxItems, "maxItems", "100", "Number of nodes in the response list (max. 1000)")
}
