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
	nodeId string,
	nodeName string,
	nodeType string,
	relativePath string,
	aspects []string,
	properties []string,
	fileName string,
	responseBody *bytes.Buffer) {

	var nodeCreate NodeUpdate
	nodeCreate.Name = nodeName
	nodeCreate.NodeType = nodeType
	nodeCreate.RelativePath = relativePath
	nodeCreate.AspectNames = aspects
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
		ResponseBodyOutput: responseBody,
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
	}
}

var nodeNameCreate string
var nodeTypeCreate string
var aspectsCreate []string
var propertiesCreate []string
var fileNameCreate string
var nodeCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Node",
	Run: func(command *cobra.Command, args []string) {
		CreateNode(nodeId,
			nodeNameCreate,
			nodeTypeCreate,
			relativePath,
			aspectsCreate,
			propertiesCreate,
			fileNameCreate,
			&responseBody)
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(CreateNodeCmdId, "Node "+nodeName+" has been created under "+nodeId)
	},
}

func init() {
	nodeCmd.AddCommand(nodeCreateCmd)
	nodeCreateCmd.Flags().StringVarP(&nodeId, "nodeId", "i", "", "Parent Node Id in Alfresco Repository. The node is created under this Parent Node.")
	nodeCreateCmd.Flags().StringVarP(&relativePath, "relativePath", "r", "", "A path relative to the nodeId.")
	nodeCreateCmd.Flags().StringVarP(&nodeNameCreate, "name", "n", "", "Node Name")
	nodeCreateCmd.Flags().StringVarP(&nodeTypeCreate, "type", "t", "", "Node Type")
	nodeCreateCmd.Flags().StringArrayVarP(&aspectsCreate, "aspects", "a", nil, "Complete aspect list to be set")
	nodeCreateCmd.Flags().StringArrayVarP(&propertiesCreate, "properties", "p", nil, "Property=Value list containing properties to be updated")
	nodeCreateCmd.Flags().StringVarP(&fileNameCreate, "file", "f", "", "Filename to be uploaded (complete or local path)")
	nodeCreateCmd.Flags().SortFlags = false
	nodeCreateCmd.MarkFlagRequired("nodeId")
	nodeCreateCmd.MarkFlagRequired("name")
	nodeCreateCmd.MarkFlagRequired("type")
}
