---
date: 2023-05-04T12:40:28+02:00
title: "alfresco node download-folder"
slug: alfresco_node_download-folder
---
## alfresco node download-folder

Download an Alfresco Repository folder to a local folder

### Synopsis

Folders and files nodes from the repository are retrieved recursively.
Only content is downloaded, while metadata is not available on the local download

```
alfresco node download-folder [flags]
```

### Options

```
  -d, --directory string      Local folder path to download Alfresco content
  -h, --help                  help for download-folder
  -i, --nodeId string         Node Id in Alfresco Repository to download to local folder. You can also use one of these well-known aliases: -my-, -shared-, -root-
  -r, --relativePath string   A path in Alfresco Repository relative to the nodeId.
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco node](/alfresco_node.md)	 - Manage nodes in ACS Repository

