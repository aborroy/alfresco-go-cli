package people

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/aborroy/alfresco-cli/cmd"
)

const peopleUrlPath string = "/api/-default-/public/alfresco/versions/1/people/"

type Company struct {
	Organization string `json:"organization,omitempty"`
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty"`
	Address3     string `json:"address3,omitempty"`
	Postcode     string `json:"postcode,omitempty"`
	Telephone    string `json:"telephone,omitempty"`
	Fax          string `json:"fax,omitempty"`
	Email        string `json:"email,omitempty"`
}

type PersonUpdate struct {
	ID                        string              `json:"id,omitempty"`
	FirstName                 string              `json:"firstName,omitempty"`
	LastName                  string              `json:"lastName,omitempty"`
	Description               string              `json:"description,omitempty"`
	Email                     string              `json:"email,omitempty"`
	SkypeID                   string              `json:"skypeId,omitempty"`
	GoogleID                  string              `json:"googleId,omitempty"`
	InstantMessageID          string              `json:"instantMessageId,omitempty"`
	JobTitle                  string              `json:"jobTitle,omitempty"`
	Location                  string              `json:"location,omitempty"`
	Company                   *Company            `json:"company,omitempty"`
	Mobile                    string              `json:"mobile,omitempty"`
	Telephone                 string              `json:"telephone,omitempty"`
	UserStatus                string              `json:"userStatus,omitempty"`
	Enabled                   bool                `json:"enabled,omitempty"`
	EmailNotificationsEnabled bool                `json:"emailNotificationsEnabled,omitempty"`
	Password                  string              `json:"password,omitempty"`
	AspectNames               []string            `json:"aspectNames,omitempty"`
	Properties                map[string](string) `json:"properties,omitempty"`
}

func PopulatePersonUpdate(properties []string, personUpdate *PersonUpdate) {
	for _, property := range properties {
		pair := strings.Split(property, "=")
		switch pair[0] {
		case "description":
			personUpdate.Description = pair[1]
		case "skypeID":
			personUpdate.SkypeID = pair[1]
		case "googleID":
			personUpdate.GoogleID = pair[1]
		case "instantMessageID":
			personUpdate.InstantMessageID = pair[1]
		case "jobTitle":
			personUpdate.JobTitle = pair[1]
		case "location":
			personUpdate.Location = pair[1]
		case "mobile":
			personUpdate.Mobile = pair[1]
		case "telephone":
			personUpdate.Telephone = pair[1]
		case "company.organization":
			personUpdate.Company.Organization = pair[1]
		case "company.address1":
			personUpdate.Company.Address1 = pair[1]
		case "company.address2":
			personUpdate.Company.Address2 = pair[1]
		case "company.address3":
			personUpdate.Company.Address3 = pair[1]
		case "company.postcode":
			personUpdate.Company.Postcode = pair[1]
		case "company.telephone":
			personUpdate.Company.Telephone = pair[1]
		case "company.fax":
			personUpdate.Company.Fax = pair[1]
		case "company.email":
			personUpdate.Company.Email = pair[1]
		default:
			log.Fatal("Property", pair[0], "is not supported")
		}
	}

}

type Person struct {
	Entry struct {
		LastName     string `json:"lastName"`
		UserStatus   string `json:"userStatus"`
		Capabilities struct {
			IsGuest   bool `json:"isGuest"`
			IsAdmin   bool `json:"isAdmin"`
			IsMutable bool `json:"isMutable"`
		} `json:"capabilities"`
		DisplayName               string `json:"displayName"`
		JobTitle                  string `json:"jobTitle"`
		StatusUpdatedAt           string `json:"statusUpdatedAt"`
		Mobile                    string `json:"mobile"`
		EmailNotificationsEnabled bool   `json:"emailNotificationsEnabled"`
		Description               string `json:"description"`
		Telephone                 string `json:"telephone"`
		Enabled                   bool   `json:"enabled"`
		FirstName                 string `json:"firstName"`
		SkypeID                   string `json:"skypeId"`
		AvatarID                  string `json:"avatarId"`
		Location                  string `json:"location"`
		Company                   struct {
			Organization string `json:"organization"`
			Address1     string `json:"address1"`
			Address2     string `json:"address2"`
			Address3     string `json:"address3"`
			Postcode     string `json:"postcode"`
		} `json:"company"`
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"entry"`
}

type PersonList struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []struct {
			Person
		} `json:"entries"`
	} `json:"list"`
}

func output(data []byte, format string) {
	if format == string(cmd.Json) {
		fmt.Println(string(data[:]))
	} else {
		var personList PersonList
		err := json.Unmarshal(data, &personList)
		if err != nil || reflect.DeepEqual(personList, PersonList{}) {
			outputPerson(data, format)
		} else {
			outputPersonList(data, format, personList)
		}
	}
}

func outputPerson(data []byte, format string) {
	switch format {
	case string(cmd.Id):
		var person Person
		json.Unmarshal(data, &person)
		fmt.Println(person.Entry.ID)
	case string(cmd.Default):
		var person Person
		json.Unmarshal(data, &person)
		if !reflect.DeepEqual(person, Person{}) {
			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tIS ADMIN\tEMAIL\t")
			fmt.Fprintln(w,
				person.Entry.ID+"\t"+
					person.Entry.DisplayName+"\t"+
					strconv.FormatBool(person.Entry.Capabilities.IsAdmin)+"\t"+
					person.Entry.Email)
			w.Flush()
		}
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}
}

func outputPersonList(data []byte, format string, personList PersonList) {
	switch format {
	case string(cmd.Id):
		for _, person := range personList.List.Entries {
			fmt.Println(person.Entry.ID)
		}
	case string(cmd.Default):
		w := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tIS ADMIN\tEMAIL\t")
		for _, person := range personList.List.Entries {
			fmt.Fprintln(w,
				person.Entry.ID+"\t"+
					person.Entry.DisplayName+"\t"+
					strconv.FormatBool(person.Entry.Capabilities.IsAdmin)+"\t"+
					person.Entry.Email)
		}
		w.Flush()
		fmt.Printf("# Count=%d, HasMoreItems=%t, TotalItems=%d, SkipCount=%d, MaxItems=%d",
			personList.List.Pagination.Count,
			personList.List.Pagination.HasMoreItems,
			personList.List.Pagination.TotalItems,
			personList.List.Pagination.SkipCount,
			personList.List.Pagination.MaxItems)
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}

}
