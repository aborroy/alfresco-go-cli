package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var server string
var username string
var password string
var insecure bool
var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Connection details storage",
	Run: func(cmd *cobra.Command, args []string) {
		_err := nativestore.Set(server, username, password)
		if _err != nil {
			fmt.Println(_err)
			os.Exit(1)
		} else {
			viper.Set(nativestore.UrlLabel, server)
			protocol := "http"
			if strings.HasPrefix(server, "https") {
				protocol = "https"
			}
			viper.Set(nativestore.ProtocolLabel, protocol)
			viper.Set(nativestore.InsecureLabel, insecure)
			viper.WriteConfig()
		}
	},
}

func init() {
	ConfigCmd.AddCommand(configSetCmd)
	configSetCmd.Flags().StringVarP(&server, "server", "s", "", "Alfresco Server URL (e.g. http://localhost:8080/alfresco)")
	configSetCmd.Flags().StringVarP(&username, "username", "u", "", "Alfresco Username")
	configSetCmd.Flags().StringVarP(&password, "password", "p", "", "Alfresco Password for the Username")
	configSetCmd.Flags().BoolVar(&insecure, "insecure", false, "Accept insecure TLS connections (to use with self-signed certificates)")
	configSetCmd.Flags().SortFlags = false
}
