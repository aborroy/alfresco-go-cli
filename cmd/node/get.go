package node

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const GetNodeCmdId string = "[NODE GET]"

func GetNodeProperties(nodeId string) bytes.Buffer {

	var responseBody bytes.Buffer

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

	return responseBody
}

func GetNodeContent(
	nodeId string,
	downloadFolderName string,
	fileName string) {

	execution := &httpclient.HttpExecution{
		Method:             http.MethodGet,
		Format:             httpclient.None,
		Data:               downloadFolderName + "/" + fileName,
		Url:                nodeUrlPath + nodeId + "/content",
		ResponseBodyOutput: &responseBody,
	}

	_error := httpclient.ExecuteDownloadContent(execution)
	if _error != nil {
		cmd.ExitWithError(GetNodeCmdId, _error)
	}

}

var downloadFolderName string
var nodeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Node information (properties and content)",
	Run: func(command *cobra.Command, args []string) {

		responseBody := GetNodeProperties(nodeId)
		var format, _ = command.Flags().GetString("output")
		outputNode(responseBody.Bytes(), format)
		log.Println(GetNodeCmdId, "Details for node "+nodeId+" have been retrieved")

		var node Node
		json.Unmarshal(responseBody.Bytes(), &node)

		if node.Entry.IsFile && downloadFolderName != "" {
			GetNodeContent(nodeId, downloadFolderName, node.Entry.Name)
			log.Println(GetNodeCmdId, "Node "+node.Entry.Name+" ("+nodeId+") has been downloaded to folder "+downloadFolderName)
		}

	},
}

func init() {
	nodeCmd.AddCommand(nodeGetCmd)
	nodeGetCmd.Flags().StringVarP(&downloadFolderName, "directory", "d", "", "Folder to download the content (complete or local path). When empty, only properties are retrieved.")
}
