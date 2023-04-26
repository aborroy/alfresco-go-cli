package node

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/aborroy/alfresco-cli/cmd"
)

const nodeUrlPath string = "/api/-default-/public/alfresco/versions/1/nodes/"

const TypeFolder string = "cm:folder"
const TypeContent string = "cm:content"

type Node struct {
	Entry struct {
		AspectNames   []string
		CreatedAt     string
		IsFolder      bool
		IsFile        bool
		CreatedByUser struct {
			ID          string
			DisplayName string
		}
		ModifiedAt     string
		ModifiedByUser struct {
			ID          string
			DisplayName string
		}
		Name       string
		ID         string
		NodeType   string
		Properties map[string](string)
	}
}

type NodeList struct {
	List struct {
		Pagination struct {
			Count        int
			HasMoreItems bool
			TotalItems   int
			SkipCount    int
			MaxItems     int
		}
		Entries []struct {
			Node
		}
	}
}

type NodeUpdate struct {
	Name        string              `json:"name,omitempty"`
	NodeType    string              `json:"nodeType,omitempty"`
	AspectNames []string            `json:"aspectNames,omitempty"`
	Properties  map[string](string) `json:"properties,omitempty"`
}

func outputNode(data []byte, format string) {

	var node Node

	switch format {
	case string(cmd.Id):
		json.Unmarshal(data, &node)
		fmt.Println(node.Entry.ID)
	case string(cmd.Default):
		json.Unmarshal(data, &node)
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tMODIFIED AT\tUSER\t")
		fmt.Fprintln(w, node.Entry.ID+"\t"+node.Entry.Name+"\t"+node.Entry.ModifiedAt+"\t"+node.Entry.ModifiedByUser.ID)
		w.Flush()
	case string(cmd.Json):
		fmt.Println(string(data[:]))
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}

}

func outputNodeList(data []byte, format string) {

	var nodeList NodeList

	switch format {
	case string(cmd.Id):
		json.Unmarshal(data, &nodeList)
		for _, node := range nodeList.List.Entries {
			fmt.Println(node.Entry.ID)
		}
	case string(cmd.Default):
		json.Unmarshal(data, &nodeList)
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tMODIFIED AT\tUSER\t")
		for _, node := range nodeList.List.Entries {
			fmt.Fprintln(w, node.Entry.ID+"\t"+node.Entry.Name+"\t"+node.Entry.ModifiedAt+"\t"+node.Entry.ModifiedByUser.ID)
		}
		w.Flush()
	case string(cmd.Json):
		fmt.Println(string(data[:]))
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}

}
