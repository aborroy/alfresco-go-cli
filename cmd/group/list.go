package group

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ListGroupCmdId string = "[GROUP LIST]"

var skipCount int
var maxItems int
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get Group list from repository",
	Long: `Properties List is provided as output of the command.
If list elements count is greater than "maxItems" flag, output includes "HasMoreItems" field set to true.
Incrementing the "skipCount" flag on a loop will allow to retrieve all the children nodes.`,
	Run: func(command *cobra.Command, args []string) {
		if maxItems == -1 {
			maxItems = viper.GetInt(nativestore.MaxItemsLabel)
		}
		var params = make(map[string]string)
		params["skipCount"] = strconv.Itoa(skipCount)
		params["maxItems"] = strconv.Itoa(maxItems)

		var url = groupsUrlPath + groupId + "/members"
		if groupId == "" {
			url = groupsUrlPath
		}

		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                url,
			Parameters:         httpclient.GetUrlParams(params),
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(ListGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(ListGroupCmdId, "Group list has been retrieved")
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
	groupListCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository. When this parameter is omitted, group list is recovered from root node.")
	groupListCmd.Flags().IntVar(&skipCount, "skipCount", 0, "Skip a number of initial nodes from the list")
	groupListCmd.Flags().IntVar(&maxItems, "maxItems", -1, "Maximum number of nodes in the response list (max. 1000)")
	groupListCmd.Flags().SortFlags = false
}
