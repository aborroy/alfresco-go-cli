package config

import (
	"log"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ConfigDeleteCmdId string = "[CONFIG DELETE]"

var configDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "ACS connection details removal",
	Long: `ACS Client configuration is removed from local ".alfresco" file.
Credentials (username and password) are removed from Native Store.`,
	Run: func(command *cobra.Command, args []string) {
		storedServer := viper.GetString(nativestore.UrlLabel)
		_err := nativestore.Delete(storedServer)
		if _err != nil {
			cmd.ExitWithError(ConfigDeleteCmdId, _err)
		}
		viper.Set(nativestore.DefaultLabel, "")
		log.Println(ConfigDeleteCmdId, "Configuration deleted for", storedServer)
	},
}

func init() {
	ConfigCmd.AddCommand(configDeleteCmd)
}
