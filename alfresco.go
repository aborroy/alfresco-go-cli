package main

import (
	"github.com/aborroy/alfresco-cli/cmd"
	_ "github.com/aborroy/alfresco-cli/cmd/config"
	_ "github.com/aborroy/alfresco-cli/cmd/node"
)

func main() {
	cmd.Execute()
}
