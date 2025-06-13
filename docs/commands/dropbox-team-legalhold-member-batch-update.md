---
layout: command
title: Command `dropbox team legalhold member batch update`
lang: en
---

# dropbox team legalhold member batch update

Add or remove multiple team members from legal hold policies in batch for efficient compliance management 

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

| Description |
|-------------|

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
.\tbx.exe dropbox team legalhold member batch update -member /PATH/TO/MEMBER_LIST.csv -policy-id POLICY_ID
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team legalhold member batch update -member /PATH/TO/MEMBER_LIST.csv -policy-id POLICY_ID
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-member**
: Path to member list file

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID

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

## Format: Member

Member email address

| Column | Description   | Example          |
|--------|---------------|------------------|
| email  | Email address | emma@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
emma@example.com
```

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: policy

This report shows a list of members.
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column           | Description                                                                                                          |
|------------------|----------------------------------------------------------------------------------------------------------------------|
| team_member_id   | ID of user as a member of a team.                                                                                    |
| email            | Email address of user.                                                                                               |
| email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| given_name       | Also known as a first name                                                                                           |
| surname          | Also known as a last name or family name.                                                                            |
| familiar_name    | Locale-dependent name                                                                                                |
| display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| abbreviated_name | An abbreviated form of the person's name.                                                                            |
| member_folder_id | The namespace id of the user's root folder.                                                                          |
| external_id      | External ID that a team can attach to the user.                                                                      |
| account_id       | A user's account identifier.                                                                                         |
| persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| invited_on       | The date and time the user was invited to the team                                                                   |
| role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| tag              | Operation tag                                                                                                        |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...


