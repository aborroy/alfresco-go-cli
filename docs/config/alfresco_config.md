---
title: "alfresco config"
slug: alfresco_config
---
## alfresco config

Manage ACS connection details

### Synopsis

ACS Client configuration is stored on a local ".alfresco" file.
Credentials (username and password) are stored on a Native Store depending on the OS.
The access to the Native Store may require typing OS credentials.

### Options

```
  -h, --help   help for config
```

### Options inherited from parent commands

```
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
```

### SEE ALSO

* [alfresco](../alfresco.md)	 - A Command Line Interface for Alfresco Content Services
* [alfresco config delete](alfresco_config_delete.md)	 - ACS connection details removal
* [alfresco config get](alfresco_config_get.md)	 - Get ACS connection details
* [alfresco config set](alfresco_config_set.md)	 - ACS connection details storage

