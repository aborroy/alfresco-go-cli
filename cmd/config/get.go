package config

import (
	"fmt"
	"os"

	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get connection details",
	Run: func(cmd *cobra.Command, args []string) {
		storedServer := viper.GetString(nativestore.DefaultLabel)
		username, password, _err := nativestore.Get(storedServer)
		if _err == nil {
			fmt.Println("server:", storedServer)
			fmt.Println("username:", username)
			fmt.Println("password:", password)
		} else {
			fmt.Println(_err)
			os.Exit(1)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(configGetCmd)
}
