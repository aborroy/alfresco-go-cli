---
title: "alfresco group add"
slug: alfresco_group_add
---
## alfresco group add

Add an authority (person or group) to a Group in repository

### Synopsis

The authority is added as children of the Group.

```
alfresco group add [flags]
```

### Options

```
  -i, --groupId string         ID of the group in Alfresco Repository.
  -a, --authorityId string     ID of the authority (group or person) in Alfresco Repository to be added.
  -t, --authorityType string   Type of the authority: GROUP or PERSON.
  -h, --help                   help for add
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco group](alfresco_group.md)	 - Manage groups in ACS Repository

