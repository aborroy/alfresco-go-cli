// This file is intended for use with "go run"; it isn't really part of the package.

package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/aborroy/alfresco-cli/cmd"
	_ "github.com/aborroy/alfresco-cli/cmd/config"
	_ "github.com/aborroy/alfresco-cli/cmd/node"
	"github.com/spf13/cobra/doc"
)

const fmTemplate = `---
date: %s
title: "%s"
slug: %s
---
`

var currentFileName string

var filePrepender = func(filename string) string {
	now := time.Now().Format(time.RFC3339)
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	currentFileName = filename
	return fmt.Sprintf(fmTemplate, now, strings.Replace(base, "_", " ", -1), base)
}

var linkHandler = func(name string) string {

	var prefix = ""
	level := len(strings.Split(currentFileName, "_"))
	if level == 3 {
		return prefix + strings.ToLower(name)
	}

	base := strings.TrimSuffix(name, path.Ext(name))
	parts := strings.Split(base, "_")
	switch len(parts) {
	case 1:
		prefix = "../"
	case 2:
		prefix = prefix + parts[1] + "/"
	case 3:
		prefix = ""
	}
	return prefix + strings.ToLower(name)
}

func main() {
	cmd.RootCmd.DisableAutoGenTag = true
	err := doc.GenMarkdownTreeCustom(cmd.RootCmd, "../", filePrepender, linkHandler)
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir("../",
		func(pathName string, info fs.DirEntry, err error) error {
			if path.Ext(pathName) == ".md" {
				base := strings.TrimSuffix(pathName, path.Ext(pathName))
				parts := strings.Split(base, "_")
				if len(parts) > 1 {
					if _, err := os.Stat(path.Dir(pathName) + "/" + parts[1]); os.IsNotExist(err) {
						os.Mkdir(path.Dir(pathName)+"/"+parts[1], os.ModePerm)
					}
					err := os.Rename(pathName, path.Dir(pathName)+"/"+parts[1]+"/"+path.Base(pathName))
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
