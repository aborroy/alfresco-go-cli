package node

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const DeleteNodeCmdId string = "[NODE DELETE]"

var nodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Node",
	Run: func(command *cobra.Command, args []string) {
		execution := &httpclient.HttpExecution{
			Method:             http.MethodDelete,
			Format:             httpclient.None,
			Url:                nodeUrlPath + nodeId,
			ResponseBodyOutput: &responseBody,
		}
		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(DeleteNodeCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(DeleteNodeCmdId, "Node "+nodeId+" has been deleted")
	},
}

func init() {
	nodeCmd.AddCommand(nodeDeleteCmd)
	nodeDeleteCmd.Flags().StringVarP(&nodeId, "nodeId", "i", "", "Node Id in Alfresco Repository to be deleted.")
	nodeDeleteCmd.MarkFlagRequired("nodeId")
}
