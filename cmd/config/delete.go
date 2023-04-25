package config

import (
	"fmt"
	"os"

	"github.com/aborroy/alfresco-cli/nativestore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Connection details removal",
	Run: func(cmd *cobra.Command, args []string) {
		storedServer := viper.GetString(nativestore.DefaultLabel)
		_err := nativestore.Delete(storedServer)
		if _err != nil {
			fmt.Println(_err)
			os.Exit(1)
		}
		viper.Set(nativestore.DefaultLabel, "")
	},
}

func init() {
	ConfigCmd.AddCommand(configDeleteCmd)
}
