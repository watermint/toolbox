---
layout: command
title: Command `team namespace file size`
lang: en
---

# team namespace file size

List all files and folders of the team namespaces 

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
| Dropbox for teams: View your team membership                                                             |
| Dropbox for teams: View structure of your team's and members' folders                                    |
| Dropbox for teams: View and edit content of your team's files and folders                                |
| Dropbox for teams: View basic information about your team including names, user count, and team settings |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account.
Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2024 Takayuki Okazaki
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
.\tbx.exe team namespace file size 
```

macOS, Linux:
```
$HOME/Desktop/tbx team namespace file size 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                   | Description                                                                       | Default |
|--------------------------|-----------------------------------------------------------------------------------|---------|
| `-depth`                 | Report entry for all files and directories depth directories deep                 | 3       |
| `-folder-name`           | List only for the folder matched to the name. Filter by exact match to the name.  |         |
| `-folder-name-prefix`    | List only for the folder matched to the name. Filter by name match to the prefix. |         |
| `-folder-name-suffix`    | List only for the folder matched to the name. Filter by name match to the suffix. |         |
| `-include-app-folder`    | If true, include app folders                                                      | false   |
| `-include-member-folder` | if true, include team member folders                                              | false   |
| `-include-shared-folder` | If true, include shared folders                                                   | true    |
| `-include-team-folder`   | If true, include team folders                                                     | true    |
| `-peer`                  | Account alias                                                                     | default |

## Common options:

| Option             | Description                                                                               | Default              |
|--------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                 |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                          | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited           | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint)                            | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage)                      | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                 | Number of processors |
| `-debug`           | Enable debug mode                                                                         | false                |
| `-experiment`      | Enable experimental feature(s).                                                           |                      |
| `-extra`           | Extra parameter file path                                                                 |                      |
| `-lang`            | Display language                                                                          | auto                 |
| `-output`          | Output format (none/text/markdown/json)                                                   | text                 |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy. |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)          | false                |
| `-retain-job-data` | Job data retain policy                                                                    | default              |
| `-secure`          | Do not store tokens into a file                                                           | false                |
| `-skip-logging`    | Skip logging in the local storage                                                         | false                |
| `-verbose`         | Show current operations for more detail.                                                  | false                |
| `-workspace`       | Workspace path                                                                            |                      |

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: namespace_size

Namespace size
The command will generate a report in three different formats. `namespace_size.csv`, `namespace_size.json`, and `namespace_size.xlsx`.

| Column               | Description                                                                                |
|----------------------|--------------------------------------------------------------------------------------------|
| namespace_name       | The name of this namespace                                                                 |
| namespace_id         | The ID of this namespace.                                                                  |
| namespace_type       | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| owner_team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |
| path                 | Path to the folder                                                                         |
| count_file           | Number of files under the folder                                                           |
| count_folder         | Number of folders under the folder                                                         |
| count_descendant     | Number of files and folders under the folder                                               |
| size                 | Size of the folder                                                                         |
| depth                | Folder depth                                                                               |
| mod_time_earliest    | The earliest modification time of a file in this folder or child folders.                  |
| mod_time_latest      | The latest modification time of a file in this folder or child folders                     |
| api_complexity       | Folder complexity index for API operations                                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


