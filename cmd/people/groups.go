package people

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

const GroupPeopleCmdId string = "[PEOPLE GROUP]"

var skipCountGroup int
var maxItemsGroup int
var peopleGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Get Group list for a person in repository",
	Long: `Properties List is provided as output of the command.
If list elements count is greater than "maxItems" flag, output includes "HasMoreItems" field set to true.
Incrementing the "skipCount" flag on a loop will allow to retrieve all the children nodes.`,
	Run: func(command *cobra.Command, args []string) {
		if maxItems == -1 {
			maxItemsGroup = viper.GetInt(nativestore.MaxItemsLabel)
		}
		var params = make(map[string]string)
		params["skipCount"] = strconv.Itoa(skipCountGroup)
		params["maxItems"] = strconv.Itoa(maxItemsGroup)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                peopleUrlPath + personId + "/groups",
			Parameters:         httpclient.GetUrlParams(params),
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(GroupPeopleCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(GroupPeopleCmdId, "Group list for person "+personId+"has been retrieved")
	},
}

func init() {
	peopleCmd.AddCommand(peopleGroupCmd)
	peopleGroupCmd.Flags().StringVarP(&personId, "personId", "i", "", "Username of the user in Alfresco Repository. You can use the -me- string in place of <personId> to specify the currently authenticated user.")
	peopleGroupCmd.Flags().IntVar(&skipCountGroup, "skipCount", 0, "Skip a number of initial nodes from the list")
	peopleGroupCmd.Flags().IntVar(&maxItemsGroup, "maxItems", -1, "Maximum number of nodes in the response list (max. 1000)")
	peopleGroupCmd.Flags().SortFlags = false
	peopleGroupCmd.MarkFlagRequired("personId")
}
