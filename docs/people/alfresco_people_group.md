---
title: "alfresco people group"
slug: alfresco_people_group
---
## alfresco people group

Get Group list for a person in repository

### Synopsis

Properties List is provided as output of the command.
If list elements count is greater than "maxItems" flag, output includes "HasMoreItems" field set to true.
Incrementing the "skipCount" flag on a loop will allow to retrieve all the children nodes.

```
alfresco people group [flags]
```

### Options

```
  -i, --personId string   Username of the user in Alfresco Repository. You can use the -me- string in place of <personId> to specify the currently authenticated user.
      --skipCount int     Skip a number of initial nodes from the list
      --maxItems int      Maximum number of nodes in the response list (max. 1000) (default -1)
  -h, --help              help for group
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco people](alfresco_people.md)	 - Manage people in ACS Repository

