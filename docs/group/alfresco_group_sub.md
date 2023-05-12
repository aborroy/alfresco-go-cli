---
title: "alfresco group sub"
slug: alfresco_group_sub
---
## alfresco group sub

Removes an authority (person or group) from a Group in repository

### Synopsis

The authority is removed as children of the Group.

```
alfresco group sub [flags]
```

### Options

```
  -i, --groupId string    ID of the group in Alfresco Repository.
  -m, --memberId string   ID of the authority (group or person) in Alfresco Repository to be added.
  -h, --help              help for sub
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco group](alfresco_group.md)	 - Manage groups in ACS Repository

