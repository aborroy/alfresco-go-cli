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

func GetNodeProperties(
	nodeId string,
	relativePath string,
	responseBody *bytes.Buffer) {

	var params = make(map[string]string)
	if relativePath != "" {
		params["relativePath"] = relativePath
	}

	execution := &httpclient.HttpExecution{
		Method:             http.MethodGet,
		Format:             httpclient.None,
		Url:                nodeUrlPath + nodeId,
		Parameters:         GetUrlParams(params),
		ResponseBodyOutput: responseBody,
	}

	_error := httpclient.Execute(execution)
	if _error != nil {
		cmd.ExitWithError(GetNodeCmdId, _error)
	}

}

func GetNodeContent(
	nodeId string,
	downloadFolderName string,
	fileName string) {

	var responseBodyContent bytes.Buffer

	execution := &httpclient.HttpExecution{
		Method:             http.MethodGet,
		Format:             httpclient.None,
		Data:               downloadFolderName + "/" + fileName,
		Url:                nodeUrlPath + nodeId + "/content",
		ResponseBodyOutput: &responseBodyContent,
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

		GetNodeProperties(nodeId, relativePath, &responseBody)

		var node Node
		json.Unmarshal(responseBody.Bytes(), &node)

		if node.Entry.IsFile && downloadFolderName != "" {
			GetNodeContent(nodeId, downloadFolderName, node.Entry.Name)
			log.Println(GetNodeCmdId, "Node "+node.Entry.Name+" ("+nodeId+") has been downloaded to folder "+downloadFolderName)
		}

	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(GetNodeCmdId, "Details for node "+nodeId+" have been retrieved")
	},
}

func init() {
	nodeCmd.AddCommand(nodeGetCmd)
	nodeGetCmd.Flags().StringVarP(&nodeId, "nodeId", "i", "", "Node Id in Alfresco Repository to be retrieved.")
	nodeGetCmd.Flags().StringVarP(&relativePath, "relativePath", "r", "", "A path relative to the nodeId.")
	nodeGetCmd.Flags().StringVarP(&downloadFolderName, "directory", "d", "", "Folder to download the content (complete or local path). When empty, only properties are retrieved.")
	nodeGetCmd.MarkFlagRequired("nodeId")
}
