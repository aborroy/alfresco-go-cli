---
date: 2023-05-04T12:46:43+02:00
title: "alfresco node list"
slug: alfresco_node_list
---
## alfresco node list

Get children nodes from a Node in the repository

### Synopsis

Metadata List for direct children nodes of a Node in the repository.
Metadata List is provided as output of the command.
If list elements count is greater than "maxItems" flag, output includes "HasMoreItems" field set to true.
Incrementing the "skipCount" flag on a loop will allow to retrieve all the children nodes.

```
alfresco node list [flags]
```

### Options

```
  -i, --nodeId string         Node Id in Alfresco Repository to get children nodes. You can also use one of these well-known aliases: -my-, -shared-, -root-
  -r, --relativePath string   A path in Alfresco Repository relative to the nodeId.
      --skipCount int         Skip a number of initial nodes from the list
      --maxItems int          Maximum number of nodes in the response list (max. 1000) (default -1)
  -h, --help                  help for list
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco node](alfresco_node.md)	 - Manage nodes in ACS Repository

