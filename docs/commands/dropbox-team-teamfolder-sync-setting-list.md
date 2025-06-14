---
layout: command
title: Command `dropbox team teamfolder sync setting list`
lang: en
---

# dropbox team teamfolder sync setting list

Display sync configuration for all team folders, showing default sync behavior for members 

Shows current sync settings for all team folders indicating whether they automatically sync to new members' devices. Helps understand bandwidth impact, storage requirements, and ensures appropriate content distribution policies.

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
| Dropbox for teams: View content of your team's files and folders                                         |
| Dropbox for teams: View structure of your team's and members' folders                                    |
| Dropbox for teams: View and edit content of your team's files and folders                                |
| Dropbox for teams: View basic information about your team including names, user count, and team settings |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account.
Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the application.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:\n\nhttps://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx\n\n2. Click 'Allow' (you might have to login first):\n3. Copy the authorization code:
Enter the authorization code
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
.\tbx.exe dropbox team teamfolder sync setting list 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team teamfolder sync setting list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-peer**
: Account alias. Default: default

**-scan-all**
: Perform a scan for all depths (can take considerable time depending on folder structure). Default: false

**-show-all**
: Show all scanned folders. Default: false

## Common options:

**-auth-database**
: Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)

**-auto-open**
: Auto open URL or artifact folder. Default: false

**-bandwidth-kb**
: Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited. Default: 0

**-budget-memory**
: Memory budget (limits some feature to reduce memory footprint). Options: low, normal. Default: normal

**-budget-storage**
: Storage budget (limits logs or some feature to reduce storage usage). Options: low, normal, unlimited. Default: normal

**-concurrency**
: Maximum concurrency for running operation. Default: Number of processors

**-debug**
: Enable debug mode. Default: false

**-experiment**
: Enable experimental feature(s).

**-extra**
: Extra parameter file path

**-lang**
: Display language. Options: auto, en, ja. Default: auto

**-output**
: Output format (none/text/markdown/json). Options: text, markdown, json, none. Default: text

**-output-filter**
: Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON.

**-proxy**
: HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.

**-quiet**
: Suppress non-error messages, and make output readable by a machine (JSON format). Default: false

**-retain-job-data**
: Job data retain policy. Options: default, on_error, none. Default: default

**-secure**
: Do not store tokens into a file. Default: false

**-skip-logging**
: Skip logging in the local storage. Default: false

**-verbose**
: Show current operations for more detail.. Default: false

**-workspace**
: Workspace path

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: folders

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `folders.csv`, `folders.json`, and `folders.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| id                          | A unique identifier for the file.                                                                                    |
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_lower                  | The lowercased full path in the user's Dropbox. This always starts with a slash.                                     |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |
| shared_folder_id            | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.                 |
| parent_shared_folder_id     | ID of shared folder that holds this file.                                                                            |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `folders_0000.xlsx`, `folders_0001.xlsx`, `folders_0002.xlsx`, ...

## Report: settings

Folder settings
The command will generate a report in three different formats. `settings.csv`, `settings.json`, and `settings.xlsx`.

| Column       | Description                                               |
|--------------|-----------------------------------------------------------|
| team_folder  | Team folder name                                          |
| path         | Path (Relative to the team folder. Blank for first level) |
| sync_setting | Sync setting                                              |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `settings_0000.xlsx`, `settings_0001.xlsx`, `settings_0002.xlsx`, ...


