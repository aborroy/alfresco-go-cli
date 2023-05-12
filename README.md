# Alfresco CLI
A Command Line Interface for Alfresco Repository using REST API.

## Usage

Download the binary compiled for your architecture (Linux, Windows or Mac OS) from [**Releases**](https://github.com/aborroy/alfresco-go-cli/releases).

>> You may rename the binary to `alfresco`, all the following samples are using this command name by default.

Using `-h` flag provides detail on the use of the different commands available:

```
$ ./alfresco -h
Alfresco CLI provides access to Alfresco REST API services via command line.
A running ACS server is required to use this program (commonly available in http://localhost:8080/alfresco).

Usage:
  alfresco [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Manage ACS connection details
  group       Manage groups in ACS Repository
  help        Help about any command
  node        Manage nodes in ACS Repository
  people      Manage people in ACS Repository

Flags:
  -h, --help              help for alfresco
  -o, --output string     Output format. E.g.: 'default', 'json' or 'id'. (default "default")
      --password string   Alfresco Password for the Username (overrides default stored config value)
      --username string   Alfresco Username (overrides default stored config value)
  -v, --version           version for alfresco

Use "alfresco [command] --help" for more information about a command.
```

**Sample working session**

![sample](https://user-images.githubusercontent.com/48685308/234789201-59f39749-da46-4630-9562-089f826e8ea9.gif)

## Configuration

Before using interactive commands with Alfresco Repository, defining connection and credentails is required.

Following command will store a connection for ACS deployed in `localhost` with default credentials (admin/admin)

```
./alfresco config set -s http://localhost:8080/alfresco -u admin -p admin
```

When using TLS protocol, an additional boolean flag `insecure` (`false` by default) can be enabled to use Alfresco CLI against Alfresco Servers configured with self-signed certificates

```
./alfresco config set -s https://localhost/alfresco -u admin -p admin --insecure true
```

Default number of results when retrieving a list of nodes from Alfresco Repository can be customized using integer flag `maxItems`. Default value is `100` and maximum value is `1000`.

```
./alfresco config set -s http://localhost:8080/alfresco -u admin -p password --maxItems 1000
```

>> Note that this command will create a `.alfresco` configuration file on the same folder

These credentials will be used by default for every command, however using specific credentials for a single command can be achived by setting `username` and `password` parameters. 

For instance, executing the `node list` command with user `test` and password `test` can be done using the following command:

```
./alfresco node list -i -root- --username test --password test
```

## Logs

Local file `alfresco.log` is capturing command logs. The log file includes activity entries and error details.

```
$ tail -f alfresco.log
2023/04/28 10:12:48 [NODE UPLOAD FOLDER] Uploading local folder ../tmp to Alfresco Repository folder 72d25c97-4ab9-47ff-8b70-b67c5f4939e8
2023/04/28 10:12:56 [NODE UPLOAD FOLDER] Folder ../tmp has been uploaded
2023/04/28 10:12:56 ERROR [NODE CREATE] Post "http://localhost:8080/alfresco/api/-default-/public/alfresco/versions/1/nodes/4960a8a5-a258-4fed-84b4-73ae8cf9b2de/children": EOF
```

## Generic flags

```
Flags:
  -h, --help            help for alfresco
  -o, --output string   Output format. E.g.: 'default', 'json' or 'id'. (default "default")
```

In addition to `help`, output formatting can be selected by using `-o` or `--output` in every command.

By default formatting is *grepable* text (`default`).

```
$ ./alfresco node list -i -root-
ID                                   NAME             MODIFIED AT                  USER
9bcd0a02-bbed-44f6-baa5-2a27974f9bdf Data Dictionary  2023-04-11T09:21:51.378+0000 System
3cb0e1bc-f349-4651-be79-c5aa4e63465c Guest Home       2023-04-11T09:21:42.711+0000 System
5aaa845d-aaa3-4550-a2f6-2cf78d12c907 Imap Attachments 2023-04-11T09:21:42.761+0000 System
a8e32ff6-6140-4a19-84a5-c157820fc376 IMAP Home        2023-04-11T09:21:42.768+0000 System
08870377-0da9-4a2d-964c-e77e4f3b5e21 Shared           2023-04-25T13:48:13.831+0000 admin
44847591-6db1-44f8-a09d-e91385de3583 Sites            2023-04-11T09:21:50.198+0000 System
308b599d-6653-4190-aee3-18e6a0f7e9fd User Homes       2023-04-13T14:21:40.945+0000 admin
# Count=7, HasMoreItems=false, TotalItems=7, SkipCount=0, MaxItems=100
```

In order to get raw JSON Response, `json` parameter can be used.

```
$ ./alfresco node list -i -root- -o json
{
  "pagination" : {
    "count" : 7,
    "hasMoreItems" : false,
    "totalItems" : 7,
    "skipCount" : 0,
    "maxItems" : 100
  },
  "entries" : [ {
    "entry" : {
      "id" : "9bcd0a02-bbed-44f6-baa5-2a27974f9bdf",
      "name" : "Data Dictionary",
      "nodeType" : "cm:folder",
      "isFolder" : true,
      "isFile" : false,
      "isLocked" : false,
...
}
```

And finally, a list of IDs can be obtained using `id` option.

```
$ ./alfresco node list -i -root- -o id
9bcd0a02-bbed-44f6-baa5-2a27974f9bdf
3cb0e1bc-f349-4651-be79-c5aa4e63465c
5aaa845d-aaa3-4550-a2f6-2cf78d12c907
a8e32ff6-6140-4a19-84a5-c157820fc376
08870377-0da9-4a2d-964c-e77e4f3b5e21
44847591-6db1-44f8-a09d-e91385de3583
308b599d-6653-4190-aee3-18e6a0f7e9fd
```

## Testing and sample scripts

Sample bash scripts for testing purposes are provided in [test](test) folder.

## Documentation

Detailed documentation for available commands in [docs/alfresco.md](docs/alfresco.md)

## TODO

* Site commands
* Search commands
* Provide pre-built programs for Windows, Linux, Mac AMD64 & Mac ARM64
* Control concurrency rate