package group

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const UpdateGroupCmdId string = "[GROUP UPDATE]"

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Group properties in ACS Repository",
	Long: `Updates an existing group in the repository.
Add only properties that require modification.`,
	Run: func(command *cobra.Command, args []string) {

		var groupUpdate GroupUpdate
		groupUpdate.DisplayName = displayName
		jsonGroupUpdate, _ := json.Marshal(groupUpdate)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodPut,
			Data:               string(jsonGroupUpdate),
			Format:             httpclient.Json,
			Url:                groupsUrlPath + groupId,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(CreateGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(CreateGroupCmdId, "Group "+groupId+" has been updated")
	},
}

func init() {
	groupCmd.AddCommand(groupUpdateCmd)
	groupUpdateCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository.")
	groupUpdateCmd.Flags().StringVarP(&displayName, "displayName", "d", "", "Display name of the group in Alfresco Repository.")
	groupUpdateCmd.Flags().SortFlags = false
	groupUpdateCmd.MarkFlagRequired("groupId")
	groupUpdateCmd.MarkFlagRequired("displayName")
}
