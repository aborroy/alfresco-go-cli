package nativestore

import (
	"fmt"
	"os"

	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/spf13/viper"
)

const DefaultLabel string = "alfresco"

func Set(url, username, secret string) error {
	creds := credentials.Credentials{
		ServerURL: url,
		Username:  username,
		Secret:    secret,
	}
	credentials.SetCredsLabel(DefaultLabel)
	return store.Add(&creds)
}

func Get(url string) (string, string, error) {
	credentials.SetCredsLabel(DefaultLabel)
	return store.Get(url)
}

func Delete(url string) error {
	credentials.SetCredsLabel(DefaultLabel)
	return store.Delete(url)
}

func GetDetails() (string, string, string) {
	var storedServer = viper.GetString(DefaultLabel)
	if storedServer == "" {
		fmt.Println("Use 'alfresco config set' to provide connection details")
		os.Exit(1)
	}
	var username, password, _err = Get(storedServer)
	if _err != nil {
		fmt.Println(_err)
		os.Exit(1)
	}
	return storedServer, username, password
}
