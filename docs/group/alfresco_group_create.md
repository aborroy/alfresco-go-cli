---
title: "alfresco group create"
slug: alfresco_group_create
---
## alfresco group create

Create new Group in ACS Repository

### Synopsis

Creates a new group in the repository.
The group can be created setting only required properties.
When specifying parentIds, the group is created associated to those parentIds, not as a children of them.

```
alfresco group create [flags]
```

### Options

```
  -i, --groupId string          ID of the group in Alfresco Repository.
  -d, --displayName string      Display name of the group in Alfresco Repository.
  -p, --parentIds stringArray   List containing the IDs of parent groups.
  -h, --help                    help for create
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco group](alfresco_group.md)	 - Manage groups in ACS Repository

