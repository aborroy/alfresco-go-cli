package node

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const NodeUpdateCmdId string = "[NODE UPDATE]"

var nodeName string
var nodeType string
var aspects []string
var properties []string
var nodeUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Node information",
	Run: func(command *cobra.Command, args []string) {

		var nodeUpdate NodeUpdate
		if nodeName != "" {
			nodeUpdate.Name = nodeName
		}
		if nodeType != "" {
			nodeUpdate.NodeType = nodeType
		}
		if aspects != nil {
			nodeUpdate.AspectNames = aspects
		}
		if properties != nil {
			m := make(map[string](string))
			for _, property := range properties {
				pair := strings.Split(property, "=")
				m[pair[0]] = pair[1]
			}
			nodeUpdate.Properties = m
		}
		jsonNodeUpdate, _ := json.Marshal(nodeUpdate)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodPut,
			Data:               string(jsonNodeUpdate),
			Format:             httpclient.Json,
			Url:                nodeUrlPath + nodeId,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(NodeUpdateCmdId, _error)
		}

		var format, _ = command.Flags().GetString("output")
		outputNode(responseBody.Bytes(), format)

		log.Println(NodeUpdateCmdId, "Node "+nodeName+" has been updated")

	},
}

func init() {
	nodeCmd.AddCommand(nodeUpdateCmd)
	nodeUpdateCmd.Flags().StringVarP(&nodeName, "name", "n", "", "Change Node Name")
	nodeUpdateCmd.Flags().StringVarP(&nodeType, "type", "t", "", "Change Node Type")
	nodeUpdateCmd.Flags().StringArrayVarP(&aspects, "aspects", "a", nil, "Complete aspect list to be set")
	nodeUpdateCmd.Flags().StringArrayVarP(&properties, "properties", "p", nil, "Property=Value list containing properties to be updated")
	nodeUpdateCmd.Flags().SortFlags = false
}
