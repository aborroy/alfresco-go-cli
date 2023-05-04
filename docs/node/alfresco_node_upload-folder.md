---
date: 2023-05-04T12:40:28+02:00
title: "alfresco node upload-folder"
slug: alfresco_node_upload-folder
---
## alfresco node upload-folder

Upload local folder to Alfresco Repository

### Synopsis

Folders and files from local folder are retrieved recursively.
Only content (folders and files) is uploaded, but only basic metadata (name and type) is created in Repository.

```
alfresco node upload-folder [flags]
```

### Options

```
  -i, --nodeId string         Parent Node Id in Alfresco Repository to upload local folder. You can also use one of these well-known aliases: -my-, -shared-, -root-
  -r, --relativePath string   A path in Alfresco Repository relative to the nodeId.
  -d, --directory string      Local folder to be uploaded (complete path)
  -h, --help                  help for upload-folder
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco node](/alfresco_node.md)	 - Manage nodes in ACS Repository

