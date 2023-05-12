package group

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const CreateGroupCmdId string = "[GROUP CREATE]"

var displayName string
var parentIds []string
var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Group in ACS Repository",
	Long: `Creates a new group in the repository.
The group can be created setting only required properties.
When specifying parentIds, the group is created associated to those parentIds, not as a children of them.`,
	Run: func(command *cobra.Command, args []string) {

		var groupCreate GroupUpdate
		groupCreate.ID = groupId
		groupCreate.DisplayName = displayName
		if parentIds != nil {
			groupCreate.ParentIds = parentIds
		}
		jsonGroupCreate, _ := json.Marshal(groupCreate)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodPost,
			Data:               string(jsonGroupCreate),
			Format:             httpclient.Json,
			Url:                strings.TrimSuffix(groupsUrlPath, "/"),
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(CreateGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(CreateGroupCmdId, "Group "+groupId+" has been created")
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)
	groupCreateCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository.")
	groupCreateCmd.Flags().StringVarP(&displayName, "displayName", "d", "", "Display name of the group in Alfresco Repository.")
	groupCreateCmd.Flags().StringArrayVarP(&parentIds, "parentIds", "p", nil, "List containing the IDs of parent groups.")
	groupCreateCmd.Flags().SortFlags = false
	groupCreateCmd.MarkFlagRequired("groupId")

}
