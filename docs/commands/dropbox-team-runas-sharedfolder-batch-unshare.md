---
layout: command
title: Command `dropbox team runas sharedfolder batch unshare`
lang: en
---

# dropbox team runas sharedfolder batch unshare

Remove sharing from multiple folders on behalf of team members, managing folder access in bulk 

Admin tool to revoke folder sharing in bulk for security or compliance. Removes sharing while preserving folder contents for the owner. Critical for incident response or preventing data leaks. All unshare actions create audit records.

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
| Dropbox for teams: View content of your Dropbox files and folders                                        |
| Dropbox for teams: View your team membership                                                             |
| Dropbox for teams: View your Dropbox sharing settings and collaborators                                  |
| Dropbox for teams: View and manage your Dropbox sharing settings and collaborators                       |
| Dropbox for teams: View structure of your team's and members' folders                                    |
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
.\tbx.exe dropbox team runas sharedfolder batch unshare -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team runas sharedfolder batch unshare -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-base-path**
: Base path of the shared folder to unshare.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-leave-copy**
: Leave a copy after unsharing.. Default: false

**-peer**
: Account alias. Default: default

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

# File formats

## Format: File

List of member folders for batch operations.

| Column       | Description                  | Example                      |
|--------------|------------------------------|------------------------------|
| member_email | Email address of the member. | member@example.com           |
| path         | Path to the member's folder. | /Team Folder/Shared/file.txt |

The first line is a header line. The program will accept a file without the header.
```
member_email,path
member@example.com,/Team Folder/Shared/file.txt
```

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                         | Description                                                                                                             |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| status                         | Status of the operation                                                                                                 |
| reason                         | Reason of failure or skipped operation                                                                                  |
| input.member_email             | Email address of the member.                                                                                            |
| input.path                     | Path to the member's folder.                                                                                            |
| result.shared_folder_id        | The ID of the shared folder.                                                                                            |
| result.parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| result.name                    | The name of this shared folder.                                                                                         |
| result.access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| result.path_lower              | The lower-cased full path of this shared folder.                                                                        |
| result.is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| result.is_team_folder          | Whether this folder is a team folder.                                                                                   |
| result.policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| result.policy_shared_link      | Who links can be shared with.                                                                                           |
| result.policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| result.policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| result.policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| result.owner_team_id           | Team ID of the folder owner team                                                                                        |
| result.owner_team_name         | Team name of the team that owns the folder                                                                              |
| result.access_inheritance      | Access inheritance type                                                                                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...


