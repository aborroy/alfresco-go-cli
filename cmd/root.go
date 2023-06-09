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
	Short: "A Command Line Interface for Alfresco Content Services",
	Long: `Alfresco CLI provides access to Alfresco REST API services via command line.
A running ACS server is required to use this program (commonly available in http://localhost:8080/alfresco).`,
	Version: "0.0.4",
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

	RootCmd.PersistentFlags().StringVar(&UsernameParam, "username", "", "Alfresco Username (overrides default stored config value)")
	RootCmd.PersistentFlags().StringVar(&PasswordParam, "password", "", "Alfresco Password for the Username (overrides default stored config value)")
	RootCmd.MarkFlagsRequiredTogether("username", "password")
	RootCmd.PersistentFlags().StringP("output", "o", "default", "Output format. E.g.: 'default', 'json' or 'id'.")
}
