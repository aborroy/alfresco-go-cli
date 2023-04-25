package node

import (
	"bytes"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

var nodeId string
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage nodes",
}

var responseBody bytes.Buffer

func init() {
	cmd.RootCmd.AddCommand(nodeCmd)
	nodeCmd.PersistentFlags().StringVarP(&nodeId, "nodeId", "i", "", "Node Id in Alfresco Repository")
	nodeCmd.Flags().SortFlags = false
}
