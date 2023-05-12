package group

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const AddGroupCmdId string = "[GROUP ADD]"

var authorityId string
var authorityType string
var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an authority (person or group) to a Group in repository",
	Long:  `The authority is added as children of the Group.`,
	Run: func(command *cobra.Command, args []string) {

		var groupAdd GroupAdd
		groupAdd.ID = authorityId
		groupAdd.MemberType = authorityType
		jsonGroupAdd, _ := json.Marshal(groupAdd)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodPost,
			Data:               string(jsonGroupAdd),
			Format:             httpclient.Json,
			Url:                groupsUrlPath + groupId + "/members",
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(AddGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(AddGroupCmdId, "Authority "+authorityId+" ("+authorityType+") has been added to group "+groupId)
	},
}

func init() {
	groupCmd.AddCommand(groupAddCmd)
	groupAddCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository.")
	groupAddCmd.Flags().StringVarP(&authorityId, "authorityId", "a", "", "ID of the authority (group or person) in Alfresco Repository to be added.")
	groupAddCmd.Flags().StringVarP(&authorityType, "authorityType", "t", "", "Type of the authority: GROUP or PERSON.")
	groupAddCmd.Flags().SortFlags = false
	groupAddCmd.MarkFlagRequired("groupId")
	groupAddCmd.MarkFlagRequired("authorityId")
	groupAddCmd.MarkFlagRequired("authorityType")
}
