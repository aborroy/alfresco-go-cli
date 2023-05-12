package people

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const UpdatePeopleCmdId string = "[PEOPLE UPDATE]"

var peopleUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Person properties in ACS Repository",
	Long: `Updates an existing person in the repository.
Add only properties that require modification.`,
	Run: func(command *cobra.Command, args []string) {

		var personUpdate PersonUpdate
		personUpdate.Password = password
		personUpdate.FirstName = firstName
		personUpdate.LastName = lastName
		personUpdate.Email = email
		PopulatePersonUpdate(properties, &personUpdate)
		jsonPersonCreate, _ := json.Marshal(personUpdate)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodPut,
			Data:               string(jsonPersonCreate),
			Format:             httpclient.Json,
			Url:                peopleUrlPath + personId,
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(CreatePeopleCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(CreatePeopleCmdId, "Person "+personId+" has been updated")
	},
}

func init() {
	peopleCmd.AddCommand(peopleUpdateCmd)
	peopleUpdateCmd.Flags().StringVarP(&personId, "personId", "i", "", "Username of the user in Alfresco Repository.")
	peopleUpdateCmd.Flags().StringVarP(&password, "password", "s", "", "Password of the user in Alfresco Repository.")
	peopleUpdateCmd.Flags().StringVarP(&firstName, "firstName", "f", "", "First Name of the user in Alfresco Repository.")
	peopleUpdateCmd.Flags().StringVarP(&lastName, "lastName", "l", "", "Last Name of the user in Alfresco Repository.")
	peopleUpdateCmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user in Alfresco Repository.")
	peopleUpdateCmd.Flags().StringArrayVarP(&properties, "properties", "p", nil,
		"Property=Value list containing properties to be updated. Property strings accepted: "+
			"description, skypeID, googleID, instantMessageID, jobTitle, location, mobile, telephone, "+
			"company.organization, company.address1, company.address2, company.address3, company.postcode, "+
			"company.telephone, company.fax, company.email")
	peopleUpdateCmd.Flags().SortFlags = false
}
