---
title: "alfresco people create"
slug: alfresco_people_create
---
## alfresco people create

Create new Person in ACS Repository

### Synopsis

Creates a new person in the repository.
The person can be created setting only required properties.

```
alfresco people create [flags]
```

### Options

```
  -i, --personId string          Username of the user in Alfresco Repository.
  -s, --password string          Password of the user in Alfresco Repository.
  -f, --firstName string         First Name of the user in Alfresco Repository.
  -l, --lastName string          Last Name of the user in Alfresco Repository.
  -e, --email string             Email of the user in Alfresco Repository.
  -p, --properties stringArray   Property=Value list containing properties to be created. Property strings accepted: description, skypeID, googleID, instantMessageID, jobTitle, location, mobile, telephone, company.organization, company.address1, company.address2, company.address3, company.postcode, company.telephone, company.fax, company.email
  -h, --help                     help for create
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco people](alfresco_people.md)	 - Manage people in ACS Repository

