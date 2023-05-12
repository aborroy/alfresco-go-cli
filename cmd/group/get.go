package group

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const GetGroupCmdId string = "[GROUP GET]"

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Group information from repository",
	Long: `Properties of the Group are retrieved. 
Properties are provided as output of the command.`,
	Run: func(command *cobra.Command, args []string) {
		execution := &httpclient.HttpExecution{
			Method:             http.MethodGet,
			Format:             httpclient.None,
			Url:                groupsUrlPath + groupId,
			ResponseBodyOutput: &responseBody,
		}
		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(GetGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(GetGroupCmdId, "Details for group "+groupId+" have been retrieved")
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)
	groupGetCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository.")
	groupGetCmd.MarkFlagRequired("groupId")
}
