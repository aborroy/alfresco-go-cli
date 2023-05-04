---
date: 2023-05-04T12:46:43+02:00
title: "alfresco node create"
slug: alfresco_node_create
---
## alfresco node create

Create new Node in ACS Repository

### Synopsis

Creates a new node as children of a parent node in the repository.
The node can be created setting only metadata (name, type, aspects and properties) or
a local file can be also specified to be associated as the content of the new node.

```
alfresco node create [flags]
```

### Options

```
  -i, --nodeId string            Parent Node Id in Alfresco Repository (commonly a folder node). The node is created under this Parent Node. You can also use one of these well-known aliases: -my-, -shared-, -root-
  -r, --relativePath string      A path in Alfresco Repository relative to the nodeId for the Parent Node.
  -n, --name string              New Node Name
  -t, --type string              New Node Type
  -a, --aspects stringArray      Complete aspect list to be set for the New Node
  -p, --properties stringArray   Property=Value list containing properties to be created for the New Node
  -f, --file string              Filename to be uploaded (complete or local path)
  -h, --help                     help for create
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco node](alfresco_node.md)	 - Manage nodes in ACS Repository

