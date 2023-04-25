package config

import (
	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Connection details",
}

func init() {
	cmd.RootCmd.AddCommand(ConfigCmd)
}
