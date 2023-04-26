package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

var nodeName string
var nodeType string
var aspects []string
var properties []string
var nodeUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Node information",
	Run: func(cmd *cobra.Command, args []string) {

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
			fmt.Println(_error)
			os.Exit(1)
		}

		var format, _ = cmd.Flags().GetString("output")
		outputNode(responseBody.Bytes(), format)

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
