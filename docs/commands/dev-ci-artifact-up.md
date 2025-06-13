---
layout: command
title: Command `dev ci artifact up`
lang: en
---

# dev ci artifact up

Upload CI artifact 

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT -timeout NUM
```

macOS, Linux:
```
$HOME/Desktop/tbx dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT -timeout NUM
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option          | Description            | Default |
|-----------------|------------------------|---------|
| `-dropbox-path` | Dropbox path to upload |         |
| `-local-path`   | Local path to upload   |         |
| `-peer-name`    | Account alias          | deploy  |

## Common options:

| Option             | Description                                                                                                                                           | Default              |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                                                                             |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                                                                                      | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited                                                                       | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint) (Options: low, normal)                                                                 | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage) (Options: low, normal, unlimited)                                                | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                                                                             | Number of processors |
| `-debug`           | Enable debug mode                                                                                                                                     | false                |
| `-experiment`      | Enable experimental feature(s).                                                                                                                       |                      |
| `-extra`           | Extra parameter file path                                                                                                                             |                      |
| `-lang`            | Display language (Options: auto, en, ja)                                                                                                              | auto                 |
| `-output`          | Output format (none/text/markdown/json) (Options: text, markdown, json, none)                                                                         | text                 |
| `-output-filter`   | Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON. |                      |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.                                                          |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy (Options: default, on_error, none)                                                                                             | default              |
| `-secure`          | Do not store tokens into a file                                                                                                                       | false                |
| `-skip-logging`    | Skip logging in the local storage                                                                                                                     | false                |
| `-verbose`         | Show current operations for more detail.                                                                                                              | false                |
| `-workspace`       | Workspace path                                                                                                                                        |                      |

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## Report: summary

This report shows a summary of the upload results.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column                | Description                                   |
|-----------------------|-----------------------------------------------|
| start                 | Time of start                                 |
| end                   | Time of finish                                |
| num_bytes             | Total upload size (Bytes)                     |
| num_files_error       | The number of files failed or got an error.   |
| num_files_transferred | The number of files uploaded/downloaded.      |
| num_files_skip        | The number of files skipped or to skip.       |
| num_folder_created    | Number of created folders.                    |
| num_delete            | Number of deleted entries.                    |
| num_api_call          | The number of estimated API calls for upload. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...


