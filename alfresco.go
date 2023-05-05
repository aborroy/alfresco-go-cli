package main

import (
	"log"
	"os"

	"github.com/aborroy/alfresco-cli/cmd"
	_ "github.com/aborroy/alfresco-cli/cmd/config"
	_ "github.com/aborroy/alfresco-cli/cmd/node"
	_ "github.com/aborroy/alfresco-cli/cmd/people"
)

func main() {
	cmd.Execute()
}

func init() {
	file, err := os.OpenFile("alfresco.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Default().SetOutput(file)
}
