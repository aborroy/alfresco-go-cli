---
date: 2023-05-04T12:40:28+02:00
title: "alfresco node update"
slug: alfresco_node_update
---
## alfresco node update

Update Node information in the repository

### Synopsis

Updates an existing node in the repository.
The node can be updating setting only modified metadata (name, type and properties) but
to modify "aspects" the full list of "aspects" to be set to the node is required.
A local file can be also specified to be replace the content of the new node.

```
alfresco node update [flags]
```

### Options

```
  -i, --nodeId string            Node Id in Alfresco Repository to be updated. You can also use one of these well-known aliases: -my-, -shared-, -root-
  -r, --relativePath string      A path in Alfresco Repository relative to the nodeId.
  -n, --name string              Updated Node Name
  -t, --type string              Updated Node Type
  -a, --aspects stringArray      Complete aspect list to be set
  -p, --properties stringArray   Property=Value list containing properties to be updated
  -f, --file string              Filename to be uploaded (complete or local path)
  -h, --help                     help for update
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco node](/alfresco_node.md)	 - Manage nodes in ACS Repository

