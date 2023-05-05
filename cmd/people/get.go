package people

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const GetPeopleCmdId string = "[PEOPLE GET]"

var peopleGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Person information from repository",
	Long: `Properties of the Person are retrieved. 
Properties are provided as output of the command.`,
	Run: func(command *cobra.Command, args []string) {
		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                peopleUrlPath + personId,
			ResponseBodyOutput: &responseBody,
		}
		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(GetPeopleCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(GetPeopleCmdId, "Details for person "+personId+" have been retrieved")
	},
}

func init() {
	peopleCmd.AddCommand(peopleGetCmd)
	peopleGetCmd.Flags().StringVarP(&personId, "personId", "i", "", "Username of the user in Alfresco Repository. You can use the -me- string in place of <personId> to specify the currently authenticated user.")
	peopleGetCmd.MarkFlagRequired("personId")
}
