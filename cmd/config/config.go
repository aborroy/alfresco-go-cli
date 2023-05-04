package config

import (
	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage ACS connection details",
	Long: `ACS Client configuration is stored on a local ".alfresco" file.
Credentials (username and password) are stored on a Native Store depending on the OS.
The access to the Native Store may require typing OS credentials.`,
}

func init() {
	cmd.RootCmd.AddCommand(ConfigCmd)
}
