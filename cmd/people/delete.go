package people

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const DeletePeopleCmdId string = "[NODE DELETE]"
const peopleDeleteUrlPath string = "/s/api/people/"

var peopleDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Person existing in the repository",
	Long:  `Removes an existing person from the repository.`,
	Run: func(command *cobra.Command, args []string) {
		execution := &httpclient.HttpExecution{
			Method:             http.MethodDelete,
			Format:             httpclient.None,
			Url:                peopleDeleteUrlPath + personId,
			ResponseBodyOutput: &responseBody,
		}
		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(DeletePeopleCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(DeletePeopleCmdId, "Person "+personId+" has been deleted")
	},
}

func init() {
	peopleCmd.AddCommand(peopleDeleteCmd)
	peopleDeleteCmd.Flags().StringVarP(&personId, "personId", "i", "", "Username of the user in Alfresco Repository.")
	peopleDeleteCmd.MarkFlagRequired("personId")
}
