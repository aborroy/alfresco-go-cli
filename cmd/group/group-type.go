package group

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"

	"github.com/aborroy/alfresco-cli/cmd"
)

const groupsUrlPath string = "/api/-default-/public/alfresco/versions/1/groups/"

type GroupUpdate struct {
	ID          string   `json:"id,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	ParentIds   []string `json:"parentIds,omitempty"`
}

type GroupAdd struct {
	ID         string `json:"id"`
	MemberType string `json:"memberType,omitempty"`
}

type Group struct {
	Entry struct {
		IsRoot      bool   `json:"isRoot"`
		DisplayName string `json:"displayName"`
		ID          string `json:"id"`
		MemberType  string `json:"memberType,omitempty"`
	} `json:"entry"`
}

type GroupList struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []struct {
			Group
		} `json:"entries"`
	} `json:"list"`
}

func output(data []byte, format string) {
	if format == string(cmd.Json) {
		fmt.Println(string(data[:]))
	} else {
		var groupList GroupList
		err := json.Unmarshal(data, &groupList)
		if err != nil || reflect.DeepEqual(groupList, GroupList{}) {
			outputGroup(data, format)
		} else {
			outputGroupList(data, format, groupList)
		}
	}
}

func outputGroup(data []byte, format string) {
	switch format {
	case string(cmd.Id):
		var group Group
		json.Unmarshal(data, &group)
		fmt.Println(group.Entry.ID)
	case string(cmd.Default):
		var group Group
		json.Unmarshal(data, &group)
		if !reflect.DeepEqual(group, Group{}) {
			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			if group.Entry.MemberType == "" {
				group.Entry.MemberType = "GROUP"
			}
			fmt.Fprintln(w, "ID\tNAME\tIS ROOT\tTYPE\t")
			fmt.Fprintln(w,
				group.Entry.ID+"\t"+
					group.Entry.DisplayName+"\t"+
					strconv.FormatBool(group.Entry.IsRoot)+"\t"+
					group.Entry.MemberType+"\t")
			w.Flush()
		}
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}
}

func outputGroupList(data []byte, format string, groupList GroupList) {
	switch format {
	case string(cmd.Id):
		for _, group := range groupList.List.Entries {
			fmt.Println(group.Entry.ID)
		}
	case string(cmd.Default):
		w := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tIS ADMIN\tTYPE\t")
		for _, group := range groupList.List.Entries {
			if group.Entry.MemberType == "" {
				group.Entry.MemberType = "GROUP"
			}
			fmt.Fprintln(w,
				group.Entry.ID+"\t"+
					group.Entry.DisplayName+"\t"+
					strconv.FormatBool(group.Entry.IsRoot)+"\t"+
					group.Entry.MemberType+"\t")
		}
		w.Flush()
		fmt.Printf("# Count=%d, HasMoreItems=%t, TotalItems=%d, SkipCount=%d, MaxItems=%d",
			groupList.List.Pagination.Count,
			groupList.List.Pagination.HasMoreItems,
			groupList.List.Pagination.TotalItems,
			groupList.List.Pagination.SkipCount,
			groupList.List.Pagination.MaxItems)
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}

}
