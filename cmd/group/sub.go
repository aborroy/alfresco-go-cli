package group

import (
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const SubGroupCmdId string = "[GROUP SUB]"

var memberId string
var groupSubCmd = &cobra.Command{
	Use:   "sub",
	Short: "Removes an authority (person or group) from a Group in repository",
	Long:  `The authority is removed as children of the Group.`,
	Run: func(command *cobra.Command, args []string) {

		execution := &httpclient.HttpExecution{
			Method:             http.MethodDelete,
			Format:             httpclient.Json,
			Url:                groupsUrlPath + groupId + "/members/" + memberId,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(SubGroupCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(SubGroupCmdId, "Auhtority "+authorityId+" ("+authorityType+") has been removed from group "+groupId)
	},
}

func init() {
	groupCmd.AddCommand(groupSubCmd)
	groupSubCmd.Flags().StringVarP(&groupId, "groupId", "i", "", "ID of the group in Alfresco Repository.")
	groupSubCmd.Flags().StringVarP(&memberId, "memberId", "m", "", "ID of the authority (group or person) in Alfresco Repository to be added.")
	groupSubCmd.Flags().SortFlags = false
	groupSubCmd.MarkFlagRequired("groupId")
	groupSubCmd.MarkFlagRequired("memberId")
}
