package node

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const NodeChildrenCmdId string = "[NODE LIST]"

func ListNode(
	nodeId string,
	relativePath string,
	skipCount int,
	maxItems int,
	responseBody *bytes.Buffer) {

	var params = make(map[string]string)
	if relativePath != "" {
		params["relativePath"] = relativePath
	}
	params["skipCount"] = strconv.Itoa(skipCount)
	params["maxItems"] = strconv.Itoa(maxItems)

	execution := &httpclient.HttpExecution{
		Method:             http.MethodGet,
		Format:             httpclient.None,
		Url:                nodeUrlPath + nodeId + "/children",
		Parameters:         GetUrlParams(params),
		ResponseBodyOutput: responseBody,
	}

	_error := httpclient.Execute(execution)
	if _error != nil {
		cmd.ExitWithError(NodeChildrenCmdId, _error)
	}

}

var skipCount int
var maxItems int
var nodeChildrenCmd = &cobra.Command{
	Use:   "list",
	Short: "Get children nodes",
	Run: func(command *cobra.Command, args []string) {
		if maxItems == -1 {
			maxItems = viper.GetInt(nativestore.MaxItemsLabel)
		}
		ListNode(nodeId, relativePath, skipCount, maxItems, &responseBody)
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(NodeChildrenCmdId, "Details for children nodes of "+nodeId+" have been retrieved")
	},
}

func init() {
	nodeCmd.AddCommand(nodeChildrenCmd)
	nodeChildrenCmd.Flags().StringVarP(&nodeId, "nodeId", "i", "", "Node Id in Alfresco Repository to get children nodes")
	nodeChildrenCmd.Flags().StringVarP(&relativePath, "relativePath", "r", "", "A path relative to the nodeId.")
	nodeChildrenCmd.Flags().IntVar(&skipCount, "skipCount", 0, "Skip a number of initial nodes from the list")
	nodeChildrenCmd.Flags().IntVar(&maxItems, "maxItems", -1, "Maximum number of nodes in the response list (max. 1000)")
	nodeChildrenCmd.Flags().SortFlags = false
	nodeChildrenCmd.MarkFlagRequired("nodeId")
}
