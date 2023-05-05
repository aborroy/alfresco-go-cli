---
title: "alfresco config set"
slug: alfresco_config_set
---
## alfresco config set

ACS connection details storage

### Synopsis

ACS Client configuration is stored on a local ".alfresco" file.
Credentials (username and password) are stored on a Native Store depending on the OS.
The access to the Native Store may require typing OS credentials.
When using TLS, "insecure" flag can be set to "true" to allow connections to ACS servers using self-signed certificates.

```
alfresco config set [flags]
```

### Options

```
  -s, --server string     Alfresco Server URL (e.g. http://localhost:8080/alfresco)
  -u, --username string   Alfresco Username
  -p, --password string   Alfresco Password for the Username
      --insecure          Accept insecure TLS connections (to use with self-signed certificates)
      --maxItems int      Maximum number of nodes in response lists (max. 1000)  (default 100)
  -h, --help              help for set
```

### Options inherited from parent commands

```
  -o, --output string   Output format. E.g.: 'default', 'json' or 'id'. (default "default")
```

### SEE ALSO

* [alfresco config](alfresco_config.md)	 - Manage ACS connection details

