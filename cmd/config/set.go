package config

import (
	"log"
	"strings"

	"github.com/aborroy/alfresco-cli/cmd"
	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ConfigSetCmdId string = "[CONFIG SET]"

var server string
var username string
var password string
var insecure bool
var maxItems int
var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "ACS connection details storage",
	Long: `ACS Client configuration is stored on a local ".alfresco" file.
Credentials (username and password) are stored on a Native Store depending on the OS.
The access to the Native Store may require typing OS credentials.
When using TLS, "insecure" flag can be set to "true" to allow connections to ACS servers using self-signed certificates.`,
	Run: func(command *cobra.Command, args []string) {
		_err := nativestore.Set(server, username, password)
		if _err != nil {
			cmd.ExitWithError(ConfigSetCmdId, _err)
		} else {
			viper.Set(nativestore.UrlLabel, server)
			protocol := "http"
			if strings.HasPrefix(server, "https") {
				protocol = "https"
			}
			viper.Set(nativestore.ProtocolLabel, protocol)
			viper.Set(nativestore.InsecureLabel, insecure)
			viper.Set(nativestore.MaxItemsLabel, maxItems)
			viper.WriteConfig()
		}
		log.Println(ConfigSetCmdId, "Configuration set for", server)
	},
}

func init() {
	ConfigCmd.AddCommand(configSetCmd)
	configSetCmd.Flags().StringVarP(&server, "server", "s", "", "Alfresco Server URL (e.g. http://localhost:8080/alfresco)")
	configSetCmd.Flags().StringVarP(&username, "username", "u", "", "Alfresco Username")
	configSetCmd.Flags().StringVarP(&password, "password", "p", "", "Alfresco Password for the Username")
	configSetCmd.Flags().BoolVar(&insecure, "insecure", false, "Accept insecure TLS connections (to use with self-signed certificates)")
	configSetCmd.Flags().IntVar(&maxItems, "maxItems", 100, "Maximum number of nodes in response lists (max. 1000) ")
	configSetCmd.Flags().SortFlags = false
	configSetCmd.MarkFlagRequired("server")
	configSetCmd.MarkFlagRequired("username")
	configSetCmd.MarkFlagRequired("password")
}
