package config

import (
	"fmt"
	"log"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ConfigGetCmdId string = "[CONFIG GET]"

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get ACS connection details",
	Long: `ACS Client configuration and credentials are retrieved.
The access to the Native Store may require typing OS credentials.`,
	Run: func(command *cobra.Command, args []string) {
		storedServer := viper.GetString(nativestore.UrlLabel)
		username, password, _err := nativestore.Get(storedServer)
		if _err == nil {
			fmt.Println("server:", storedServer)
			if viper.GetString(nativestore.ProtocolLabel) == "https" {
				fmt.Println("insecure:", viper.GetString(nativestore.InsecureLabel))
			}
			fmt.Println("username:", username)
			fmt.Println("password:", password)
			fmt.Println("maxItems:", viper.GetInt(nativestore.MaxItemsLabel))
		} else {
			cmd.ExitWithError(ConfigGetCmdId, _err)
		}
		log.Println(ConfigGetCmdId, "Configuration get for", storedServer)
	},
}

func init() {
	ConfigCmd.AddCommand(configGetCmd)
}
