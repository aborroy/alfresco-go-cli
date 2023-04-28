package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Format string

const (
	Json    Format = "json"
	Id      Format = "id"
	Default Format = "default"
)

var cfgFile string
var RootCmd = &cobra.Command{
	Use:   "alfresco",
	Short: "A Command Line Interface for Alfresco Content Services.",
}

var UsernameParam string
var PasswordParam string

func Execute() {
	_err := RootCmd.Execute()
	if _err != nil {
		os.Exit(1)
	}
}

func ExitWithError(CmdId string, err error) {
	fmt.Println("ERROR", CmdId, err)
	log.Fatal("ERROR " + CmdId + " " + err.Error())
	os.Exit(1)
}

func init() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".alfresco")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		os.WriteFile(".alfresco", nil, 0644)
	}

	RootCmd.PersistentFlags().StringVarP(&UsernameParam, "username", "u", "", "Alfresco Username (overrides default stored config value)")
	RootCmd.PersistentFlags().StringVarP(&PasswordParam, "password", "p", "", "Alfresco Password for the Username (overrides default stored config value)")
	RootCmd.PersistentFlags().StringP("output", "o", "default", "Output format. E.g.: 'default', 'json' or 'id'.")
}
