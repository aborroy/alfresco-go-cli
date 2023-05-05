---
title: "alfresco node get"
slug: alfresco_node_get
---
## alfresco node get

Get Node information (properties and content) from repository

### Synopsis

Metadata and content of a Node is downloaded. 
Metadata is provided as output of the command.
Content is retrieved optionally (only when "d" flag is populated).

```
alfresco node get [flags]
```

### Options

```
  -d, --directory string      Folder to download the content (complete or local path). When empty, only properties are retrieved.
  -h, --help                  help for get
  -i, --nodeId string         Node Id in Alfresco Repository to be retrieved. You can also use one of these well-known aliases: -my-, -shared-, -root-
  -r, --relativePath string   A path in Alfresco Repository relative to the nodeId.
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco node](alfresco_node.md)	 - Manage nodes in ACS Repository

