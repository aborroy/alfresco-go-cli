package people

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/httpclient"
	"github.com/spf13/cobra"
)

const CreatePeopleCmdId string = "[PEOPLE CREATE]"

var password string
var firstName string
var lastName string
var email string
var properties []string
var peopleCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Person in ACS Repository",
	Long: `Creates a new person in the repository.
The person can be created setting only required properties.`,
	Run: func(command *cobra.Command, args []string) {

		var personCreate PersonUpdate
		personCreate.ID = personId
		personCreate.Password = password
		personCreate.FirstName = firstName
		personCreate.LastName = lastName
		personCreate.Email = email
		personCreate.Enabled = true
		PopulatePersonUpdate(properties, &personCreate)
		jsonPersonCreate, _ := json.Marshal(personCreate)

		execution := &httpclient.HttpExecution{
			Method:             http.MethodPost,
			Data:               string(jsonPersonCreate),
			Format:             httpclient.Json,
			Url:                strings.TrimSuffix(peopleUrlPath, "/"),
			ResponseBodyOutput: &responseBody,
		}

		_error := httpclient.Execute(execution)
		if _error != nil {
			cmd.ExitWithError(CreatePeopleCmdId, _error)
		}
	},
	PostRun: func(command *cobra.Command, args []string) {
		log.Println(CreatePeopleCmdId, "Person "+personId+" has been created")
	},
}

func init() {
	peopleCmd.AddCommand(peopleCreateCmd)
	peopleCreateCmd.Flags().StringVarP(&personId, "personId", "i", "", "Username of the user in Alfresco Repository.")
	peopleCreateCmd.Flags().StringVarP(&password, "password", "s", "", "Password of the user in Alfresco Repository.")
	peopleCreateCmd.Flags().StringVarP(&firstName, "firstName", "f", "", "First Name of the user in Alfresco Repository.")
	peopleCreateCmd.Flags().StringVarP(&lastName, "lastName", "l", "", "Last Name of the user in Alfresco Repository.")
	peopleCreateCmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user in Alfresco Repository.")
	peopleCreateCmd.Flags().StringArrayVarP(&properties, "properties", "p", nil,
		"Property=Value list containing properties to be created. Property strings accepted: "+
			"description, skypeID, googleID, instantMessageID, jobTitle, location, mobile, telephone, "+
			"company.organization, company.address1, company.address2, company.address3, company.postcode, "+
			"company.telephone, company.fax, company.email")
	peopleCreateCmd.Flags().SortFlags = false
	peopleCreateCmd.MarkFlagRequired("personId")
	peopleCreateCmd.MarkFlagRequired("password")
	peopleCreateCmd.MarkFlagRequired("firstName")
	peopleCreateCmd.MarkFlagRequired("lastName")
	peopleCreateCmd.MarkFlagRequired("email")

}
