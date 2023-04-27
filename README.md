# Alfresco CLI
A Command Line Interface for Alfresco Repository using REST API.

## Usage

Download the binary compiled for your architecture (Linux, Windows or Mac OS) from **Releases**.

>> You may rename the binary to `alfresco`, all the following samples are using this command name by default.

Using `-h` flag provides detail on the use of the different commands available:

```
$ ./alfresco -h
A Command Line Interface for Alfresco Content Services.

Usage:
  alfresco [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Connection details
  help        Help about any command
  node        Manage nodes

Flags:
  -h, --help            help for alfresco
  -o, --output string   Output format. E.g.: 'default', 'json' or 'id'. (default "default")

Use "alfresco [command] --help" for more information about a command.
```
![sample](https://user-images.githubusercontent.com/48685308/234789201-59f39749-da46-4630-9562-089f826e8ea9.gif)

## Configuration

Before using interactive commands with Alfresco Repository, defining connection and credentails is required.

Following command will store a connection for ACS deployed in `localhost` with default credentials (admin/admin)

```
./alfresco config set -s http://localhost:8080/alfresco -u admin -p admin
```

>> Note that this command will create a `.alfresco` configuration file on the same folder

## Commands

**Node**

The `node` command provides access to Node handling in Alfresco Repository.

```
./alfresco node -h
Manage nodes

Usage:
  alfresco node [command]

Available Commands:
  create        Create new Node
  delete        Delete Node
  get           Get Node information
  list          Get children nodes
  update        Update Node information
  upload-folder Upload local folder to Alfresco Repository

Flags:
  -i, --nodeId string   Node Id in Alfresco Repository
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

## Sample script

Sample bash script for testing purposes is provided in [sample/test.sh](sample/test.sh) file.

## TODO

* TLS protocol support
* Specific credentials per command
* Progress log
* Download folder
* Site commands
* Person commands
* Group commands
* Search commands
