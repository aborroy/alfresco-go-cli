---
title: "alfresco people update"
slug: alfresco_people_update
---
## alfresco people update

Update Person properties in ACS Repository

### Synopsis

Updates an existing person in the repository.
Add only properties that require modification.

```
alfresco people update [flags]
```

### Options

```
  -i, --personId string          Username of the user in Alfresco Repository.
  -s, --password string          Password of the user in Alfresco Repository.
  -f, --firstName string         First Name of the user in Alfresco Repository.
  -l, --lastName string          Last Name of the user in Alfresco Repository.
  -e, --email string             Email of the user in Alfresco Repository.
  -p, --properties stringArray   Property=Value list containing properties to be updated. Property strings accepted: description, skypeID, googleID, instantMessageID, jobTitle, location, mobile, telephone, company.organization, company.address1, company.address2, company.address3, company.postcode, company.telephone, company.fax, company.email
  -h, --help                     help for update
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco people](alfresco_people.md)	 - Manage people in ACS Repository

