package node

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

var nodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Node",
	Run: func(cmd *cobra.Command, args []string) {

		execution := &httpclient.HttpExecution{
			Method:             http.MethodDelete,
			Format:             httpclient.None,
			Url:                nodeUrlPath + nodeId,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			fmt.Println(_error)
			os.Exit(1)
		}

	},
}

func init() {
	nodeCmd.AddCommand(nodeDeleteCmd)
}
