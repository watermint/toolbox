---
layout: command
title: Command `dropbox team runas file sync batch up`
lang: en
---

# dropbox team runas file sync batch up

Batch sync up that run as members (Irreversible operation)

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS      | Path                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox for teams: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## Auth scopes

| Description                                                                                              |
|----------------------------------------------------------------------------------------------------------|
| Dropbox for teams: View and edit basic information about your Dropbox account such as your profile photo |
| Dropbox for teams: View content of your Dropbox files and folders                                        |
| Dropbox for teams: Edit content of your Dropbox files and folders                                        |
| Dropbox for teams: View your team membership                                                             |
| Dropbox for teams: View structure of your team's and members' folders                                    |
| Dropbox for teams: View basic information about your team including names, user count, and team settings |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account.
Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe dropbox team runas file sync batch up -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team runas file sync batch up -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                 | Description                                                | Default |
|------------------------|------------------------------------------------------------|---------|
| `-batch-size`          | Batch commit size                                          | 250     |
| `-delete`              | Delete Dropbox file if a file removed locally              | false   |
| `-exit-on-failure`     | Exit the program on failure                                | false   |
| `-file`                | Path to data file                                          |         |
| `-name-disable-ignore` | Filter by name. Filter system file or ignore files.        |         |
| `-name-name`           | Filter by name. Filter by exact match to the name.         |         |
| `-name-name-prefix`    | Filter by name. Filter by name match to the prefix.        |         |
| `-name-name-suffix`    | Filter by name. Filter by name match to the suffix.        |         |
| `-overwrite`           | Overwrite existing file on the target path if that exists. | false   |
| `-peer`                | Account alias                                              | default |

## Common options:

| Option             | Description                                                                                                                                           | Default              |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                                                                             |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                                                                                      | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited                                                                       | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint)                                                                                        | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage)                                                                                  | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                                                                             | Number of processors |
| `-debug`           | Enable debug mode                                                                                                                                     | false                |
| `-experiment`      | Enable experimental feature(s).                                                                                                                       |                      |
| `-extra`           | Extra parameter file path                                                                                                                             |                      |
| `-lang`            | Display language                                                                                                                                      | auto                 |
| `-output`          | Output format (none/text/markdown/json)                                                                                                               | text                 |
| `-output-filter`   | Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON. |                      |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy.                                                             |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy                                                                                                                                | default              |
| `-secure`          | Do not store tokens into a file                                                                                                                       | false                |
| `-skip-logging`    | Skip logging in the local storage                                                                                                                     | false                |
| `-verbose`         | Show current operations for more detail.                                                                                                              | false                |
| `-workspace`       | Workspace path                                                                                                                                        |                      |

# File formats

## Format: File

Mapping of local to destination paths

| Column       | Description                     | Example           |
|--------------|---------------------------------|-------------------|
| member_email | The email address of the member | emma@example.com  |
| local_path   | Local file path                 | /file_server/emma |
| dropbox_path | Destination Dropbox path        | /data             |

The first line is a header line. The program will accept a file without the header.
```
member_email,local_path,dropbox_path
emma@example.com,/file_server/emma,/data
```

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: deleted

Path
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                       | Description      |
|------------------------------|------------------|
| entry_path                   | Path             |
| entry_shard.file_system_type | File system type |
| entry_shard.shard_id         | Shard ID         |
| entry_shard.attributes       | Shard attributes |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column             | Description                            |
|--------------------|----------------------------------------|
| status             | Status of the operation                |
| reason             | Reason of failure or skipped operation |
| input.member_email | The email address of the member        |
| input.local_path   | Local file path                        |
| input.dropbox_path | Destination Dropbox path               |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

## Report: skipped

This report shows the transaction result.
The command will generate a report in three different formats. `skipped.csv`, `skipped.json`, and `skipped.xlsx`.

| Column                             | Description                            |
|------------------------------------|----------------------------------------|
| status                             | Status of the operation                |
| reason                             | Reason of failure or skipped operation |
| input.entry_path                   | Path                                   |
| input.entry_shard.file_system_type | File system type                       |
| input.entry_shard.shard_id         | Shard ID                               |
| input.entry_shard.attributes       | Shard attributes                       |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## Report: summary

This report shows a summary of the upload results.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column                | Description                                         |
|-----------------------|-----------------------------------------------------|
| start                 | Time of start                                       |
| end                   | Time of finish                                      |
| num_bytes             | Total upload size (Bytes)                           |
| num_files_error       | The number of files failed or got an error.         |
| num_files_transferred | The number of files uploaded/downloaded.            |
| num_files_skip        | The number of files skipped or to skip.             |
| num_folder_created    | Number of created folders.                          |
| num_delete            | Number of deleted entry.                            |
| num_api_call          | The number of estimated upload API call for upload. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

## Report: uploaded

This report shows the transaction result.
The command will generate a report in three different formats. `uploaded.csv`, `uploaded.json`, and `uploaded.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.path                         | Path                                                                                                                 |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.content_hash                | A hash of the file content.                                                                                          |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


