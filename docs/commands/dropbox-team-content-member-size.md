---
layout: command
title: Command `dropbox team content member size`
lang: en
---

# dropbox team content member size

Count number of members of team folders and shared folders 

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
| Dropbox for teams: View information about your Dropbox files and folders                                 |
| Dropbox for teams: View your team group membership                                                       |
| Dropbox for teams: View your team membership                                                             |
| Dropbox for teams: View your Dropbox sharing settings and collaborators                                  |
| Dropbox for teams: View structure of your team's and members' folders                                    |
| Dropbox for teams: View and edit content of your team's files and folders                                |
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
.\tbx.exe dropbox team content member size 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team content member size 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Default |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-base-path`           | Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder. | root    |
| `-folder-name`         | Filter by folder name. Filter by exact match to the name.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |         |
| `-folder-name-prefix`  | Filter by folder name. Filter by name match to the prefix.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |         |
| `-folder-name-suffix`  | Filter by folder name. Filter by name match to the suffix.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |         |
| `-include-sub-folders` | Include sub-folders to the report.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | false   |
| `-peer`                | Account alias                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | default |
| `-scan-timeout`        | Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | short   |

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

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: member_count

Folder member count
The command will generate a report in three different formats. `member_count.csv`, `member_count.json`, and `member_count.xlsx`.

| Column                | Description                                                                                                                        |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------|
| path                  | Path                                                                                                                               |
| folder_type           | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)                           |
| owner_team_name       | Team name of the team that owns the folder                                                                                         |
| has_no_inherit        | True if the folder or any sub-folder does not inherit the access permission from the parent folder                                 |
| is_no_inherit         | True if the folder does not inherit the access from the parent folder                                                              |
| capacity              | Capacity number to add members. Empty if it's not able to determine by your permission (e.g. a folder contains an external group). |
| count_total           | Total number of members                                                                                                            |
| count_external_groups | Number of external teams' group                                                                                                    |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `member_count_0000.xlsx`, `member_count_0001.xlsx`, `member_count_0002.xlsx`, ...


