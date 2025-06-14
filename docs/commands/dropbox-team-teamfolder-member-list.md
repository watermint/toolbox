---
layout: command
title: Command `dropbox team teamfolder member list`
lang: en
---

# dropbox team teamfolder member list

Display all members with access to each team folder, showing permission levels and access types 

Shows complete membership for all team folders including permission levels and whether access is direct or through groups. Critical for access audits, security reviews, and understanding who can access sensitive content. Identifies over-privileged access.

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
.\tbx.exe dropbox team teamfolder member list 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team teamfolder member list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-member-type-external**
: Filter folder members. Keep only members that are external (not in the same team). Note: Invited members are marked as external member.

**-member-type-internal**
: Filter folder members. Keep only members that are internal (in the same team). Note: Invited members are marked as external member.

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

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

## Report: membership

This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.
The command will generate a report in three different formats. `membership.csv`, `membership.json`, and `membership.xlsx`.

| Column          | Description                                                                                                                          |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| path            | Path                                                                                                                                 |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)                             |
| owner_team_name | Team name of the team that owns the folder                                                                                           |
| access_type     | User's access level for this folder                                                                                                  |
| member_type     | Type of this member (user, group, or invitee)                                                                                        |
| member_name     | Name of this member                                                                                                                  |
| member_email    | Email address of this member                                                                                                         |
| same_team       | Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `membership_0000.xlsx`, `membership_0001.xlsx`, `membership_0002.xlsx`, ...

## Report: no_member

This report shows folders without members.
The command will generate a report in three different formats. `no_member.csv`, `no_member.json`, and `no_member.xlsx`.

| Column          | Description                                                                                              |
|-----------------|----------------------------------------------------------------------------------------------------------|
| owner_team_name | Team name of the team that owns the folder                                                               |
| path            | Path                                                                                                     |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `no_member_0000.xlsx`, `no_member_0001.xlsx`, `no_member_0002.xlsx`, ...


