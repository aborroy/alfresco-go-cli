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

const ListPeopleCmdId string = "[PEOPLE LIST]"

var skipCount int
var maxItems int
var peopleListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get Person list from repository",
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

		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                peopleUrlPath,
			Parameters:         httpclient.GetUrlParams(params),
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(ListPeopleCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(ListPeopleCmdId, "Person list has been retrieved")
	},
}

func init() {
	peopleCmd.AddCommand(peopleListCmd)
	peopleListCmd.Flags().IntVar(&skipCount, "skipCount", 0, "Skip a number of initial nodes from the list")
	peopleListCmd.Flags().IntVar(&maxItems, "maxItems", -1, "Maximum number of nodes in the response list (max. 1000)")
	peopleListCmd.Flags().SortFlags = false
}
