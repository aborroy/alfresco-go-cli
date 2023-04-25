package node

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

var nodeChildrenCmd = &cobra.Command{
	Use:   "list",
	Short: "Get children nodes",
	Run: func(cmd *cobra.Command, args []string) {

		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                nodeUrlPath + nodeId + "/children",
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
}
