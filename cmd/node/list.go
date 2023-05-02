package node

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
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
	skipCount int,
	maxItems int) bytes.Buffer {

	var responseBody bytes.Buffer

	params := url.Values{}
	params.Add("skipCount", strconv.Itoa(skipCount))
	params.Add("maxItems", strconv.Itoa(maxItems))

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

	return responseBody
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
		response := ListNode(nodeId, skipCount, maxItems)
		var format, _ = command.Flags().GetString("output")
		outputNodeList(response.Bytes(), format)
		log.Println(NodeChildrenCmdId, "Details for children nodes of "+nodeId+" have been retrieved")
	},
}

func init() {
	nodeCmd.AddCommand(nodeChildrenCmd)
	nodeChildrenCmd.Flags().IntVar(&skipCount, "skipCount", 0, "Skip a number of initial nodes from the list")
	nodeChildrenCmd.Flags().IntVar(&maxItems, "maxItems", -1, "Maximum number of nodes in the response list (max. 1000)")
}
