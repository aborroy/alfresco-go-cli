package group

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const DeleteGroupCmdId string = "[GROUP DELETE]"

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Group existing in the repository",
	Long:  `Removes an existing group from the repository.`,
	Run: func(command *cobra.Command, args []string) {
		execution := &httpclient.HttpExecution{
			Method:             http.MethodDelete,
			Format:             httpclient.None,
			Url:                groupsUrlPath + groupId,
			ResponseBodyOutput: &responseBody,
		}
		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(DeleteGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(DeleteGroupCmdId, "Group "+groupId+" has been deleted")
	},
}

func init() {
	groupCmd.AddCommand(groupDeleteCmd)
	groupDeleteCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository.")
	groupDeleteCmd.MarkFlagRequired("groupId")
}
