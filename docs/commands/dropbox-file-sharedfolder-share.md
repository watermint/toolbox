---
layout: command
title: Command `dropbox file sharedfolder share`
lang: en
---

# dropbox file sharedfolder share

Share a folder 

Creates a shared folder from an existing folder with configurable sharing policies and permissions.

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
* Dropbox (Individual account): https://help.dropbox.com/installs-integrations/third-party/third-party-apps

## Auth scopes

| Description                                                                                          |
|------------------------------------------------------------------------------------------------------|
| Dropbox: View basic information about your Dropbox account such as your username, email, and country |
| Dropbox: View and manage your Dropbox sharing settings and collaborators                             |

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
.\tbx.exe dropbox file sharedfolder share -path /DROPBOX/PATH/TO/SHARE
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox file sharedfolder share -path /DROPBOX/PATH/TO/SHARE
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-acl-update-policy**
: Who can change a shared folder's access control list (ACL).. Options: owner (aclupdatepolicy: owner), editor (aclupdatepolicy: editor). Default: owner

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-policy**
: Who can be a member of this shared folder.. Options: team (memberpolicy: team), anyone (memberpolicy: anyone). Default: anyone

**-path**
: Path to be shared

**-peer**
: Account alias. Default: default

**-shared-link-policy**
: Who can view shared links in this folder.. Options: anyone (sharedlinkpolicy: anyone), members (sharedlinkpolicy: members). Default: anyone

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

## Report: shared

This report shows a list of shared folders.
The command will generate a report in three different formats. `shared.csv`, `shared.json`, and `shared.xlsx`.

| Column                  | Description                                                                                                             |
|-------------------------|-------------------------------------------------------------------------------------------------------------------------|
| shared_folder_id        | The ID of the shared folder.                                                                                            |
| parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| name                    | The name of this shared folder.                                                                                         |
| access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| path_lower              | The lower-cased full path of this shared folder.                                                                        |
| is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| is_team_folder          | Whether this folder is a team folder.                                                                                   |
| policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| policy_shared_link      | Who links can be shared with.                                                                                           |
| policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| owner_team_id           | Team ID of the folder owner team                                                                                        |
| owner_team_name         | Team name of the team that owns the folder                                                                              |
| access_inheritance      | Access inheritance type                                                                                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_0000.xlsx`, `shared_0001.xlsx`, `shared_0002.xlsx`, ...


