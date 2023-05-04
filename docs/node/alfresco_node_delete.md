---
date: 2023-05-04T12:40:28+02:00
title: "alfresco node delete"
slug: alfresco_node_delete
---
## alfresco node delete

Delete a Node existing in the repository

### Synopsis

Removes an existing node from the repository.
Both metadata and content resources are removed.

```
alfresco node delete [flags]
```

### Options

```
  -h, --help                  help for delete
  -i, --nodeId string         Node Id in Alfresco Repository to be deleted. You can also use one of these well-known aliases: -my-, -shared-, -root-
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

