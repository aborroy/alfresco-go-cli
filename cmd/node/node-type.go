package node

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"

	"github.com/aborroy/alfresco-cli/cmd"
)

const nodeUrlPath string = "/api/-default-/public/alfresco/versions/1/nodes/"

const TypeFolder string = "cm:folder"
const TypeContent string = "cm:content"

type Node struct {
	Entry struct {
		AspectNames   []string `json:"aspectNames"`
		CreatedAt     string   `json:"createdAt"`
		IsFolder      bool     `json:"isFolder"`
		IsFile        bool     `json:"isFile"`
		CreatedByUser struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"createdByUser"`
		ModifiedAt     string `json:"modifiedAt"`
		ModifiedByUser struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"modifiedByUser"`
		Name       string              `json:"name"`
		ID         string              `json:"id"`
		NodeType   string              `json:"nodeType"`
		Properties map[string](string) `json:"properties"`
	} `json:"entry"`
}

type NodeList struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []struct {
			Node
		} `json:"entries"`
	} `json:"list"`
}

type NodeUpdate struct {
	Name         string              `json:"name,omitempty"`
	NodeType     string              `json:"nodeType,omitempty"`
	RelativePath string              `json:"relativePath,omitempty"`
	AspectNames  []string            `json:"aspectNames,omitempty"`
	Properties   map[string](string) `json:"properties,omitempty"`
}

func output(data []byte, format string) {
	if format == string(cmd.Json) {
		fmt.Println(string(data[:]))
	} else {
		var nodeList NodeList
		err := json.Unmarshal(data, &nodeList)
		if err != nil || reflect.DeepEqual(nodeList, NodeList{}) {
			outputNode(data, format)
		} else {
			outputNodeList(data, format, nodeList)
		}
	}
}

func outputNode(data []byte, format string) {

	var node Node

	switch format {
	case string(cmd.Id):
		json.Unmarshal(data, &node)
		fmt.Println(node.Entry.ID)
	case string(cmd.Default):
		json.Unmarshal(data, &node)
		if !reflect.DeepEqual(node, Node{}) {
			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tMODIFIED AT\tUSER\t")
			fmt.Fprintln(w, node.Entry.ID+"\t"+node.Entry.Name+"\t"+node.Entry.ModifiedAt+"\t"+node.Entry.ModifiedByUser.ID)
			w.Flush()
		}
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}

}

func outputNodeList(data []byte, format string, nodeList NodeList) {

	switch format {
	case string(cmd.Id):
		for _, node := range nodeList.List.Entries {
			fmt.Println(node.Entry.ID)
		}
	case string(cmd.Default):
		w := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tMODIFIED AT\tUSER\t")
		for _, node := range nodeList.List.Entries {
			fmt.Fprintln(w, node.Entry.ID+"\t"+node.Entry.Name+"\t"+node.Entry.ModifiedAt+"\t"+node.Entry.ModifiedByUser.ID)
		}
		w.Flush()
		fmt.Printf("# Count=%d, HasMoreItems=%t, TotalItems=%d, SkipCount=%d, MaxItems=%d",
			nodeList.List.Pagination.Count,
			nodeList.List.Pagination.HasMoreItems,
			nodeList.List.Pagination.TotalItems,
			nodeList.List.Pagination.SkipCount,
			nodeList.List.Pagination.MaxItems)
	default:
		fmt.Println("Format '" + format + "' is not an option, allowed values are 'id', 'json' or 'default'")
	}

}
