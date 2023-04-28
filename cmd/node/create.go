package node

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const CreateNodeCmdId string = "[NODE CREATE]"

func CreateNode(
	command *cobra.Command,
	nodeId string,
	nodeName string,
	nodeType string,
	aspects []string,
	properties []string,
	fileName string) bytes.Buffer {

	var responseBody bytes.Buffer

	var nodeCreate NodeUpdate
	if nodeName != "" {
		nodeCreate.Name = nodeName
	}
	if nodeType != "" {
		nodeCreate.NodeType = nodeType
	}
	if aspects != nil {
		nodeCreate.AspectNames = aspects
	}
	if properties != nil {
		m := make(map[string](string))
		for _, property := range properties {
			pair := strings.Split(property, "=")
			m[pair[0]] = pair[1]
		}
		nodeCreate.Properties = m
	}
	jsonNodeCreate, _ := json.Marshal(nodeCreate)

	execution := &httpclient.HttpExecution{
		Method:             http.MethodPost,
		Data:               string(jsonNodeCreate),
		Format:             httpclient.Json,
		Url:                nodeUrlPath + nodeId + "/children",
		ResponseBodyOutput: &responseBody,
	}

	_error := httpclient.Execute(execution)
	if _error != nil {
		cmd.ExitWithError(CreateNodeCmdId, _error)
	}

	if fileName != "" {
		var node Node
		json.Unmarshal(responseBody.Bytes(), &node)

		var responseBodyContent bytes.Buffer
		uploadExecution := &httpclient.HttpExecution{
			Method:             http.MethodPut,
			Data:               fileName,
			Format:             httpclient.Content,
			Url:                nodeUrlPath + node.Entry.ID + "/content",
			ResponseBodyOutput: &responseBodyContent,
		}
		_error = httpclient.ExecuteUploadContent(uploadExecution)
		if _error != nil {
			cmd.ExitWithError(CreateNodeCmdId, _error)
		}
		return responseBodyContent
	} else {
		return responseBody
	}
}

var fileName string
var nodeCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Node",
	Run: func(command *cobra.Command, args []string) {
		response := CreateNode(command, nodeId, nodeName, nodeType, aspects, properties, fileName)
		var format, _ = command.Flags().GetString("output")
		outputNode(response.Bytes(), format)
		log.Println(CreateNodeCmdId, "Node "+nodeName+" has been created under "+nodeId)
	},
}

func init() {
	nodeCmd.AddCommand(nodeCreateCmd)
	nodeCreateCmd.Flags().StringVarP(&nodeName, "name", "n", "", "Node Name")
	nodeCreateCmd.Flags().StringVarP(&nodeType, "type", "t", "", "Node Type")
	nodeCreateCmd.Flags().StringArrayVarP(&aspects, "aspects", "a", nil, "Complete aspect list to be set")
	nodeCreateCmd.Flags().StringArrayVarP(&properties, "properties", "p", nil, "Property=Value list containing properties to be updated")
	nodeCreateCmd.Flags().StringVarP(&fileName, "file", "f", "", "Filename to be uploaded (complete or local path)")
	nodeCreateCmd.Flags().SortFlags = false
}
