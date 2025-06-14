---
layout: command
title: Command `dropbox team device unlink`
lang: en
---

# dropbox team device unlink

Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access (Irreversible operation)

Immediately terminates device sessions, forcing re-authentication. Critical security tool for lost devices, departing employees, or suspicious activity. Device must reconnect and re-sync after unlinking. Consider member communication before bulk unlinking.

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
| Dropbox for teams: View and manage your team's sessions, devices, and apps                               |
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
.\tbx.exe dropbox team device unlink -file /path/to/data/file.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team device unlink -file /path/to/data/file.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-delete-on-unlink**
: Delete files on unlink. Default: false

**-file**
: Data file

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

This report shows a list of current existing sessions in the team with team member information.

| Column                        | Description                                                                          | Example                                                                                                         |
|-------------------------------|--------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| team_member_id                | ID of user as a member of a team.                                                    | dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                                                       |
| email                         | Email address of user.                                                               | john.smith@example.com                                                                                          |
| status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) | active                                                                                                          |
| given_name                    | Also known as a first name                                                           | John                                                                                                            |
| surname                       | Also known as a last name or family name.                                            | Smith                                                                                                           |
| familiar_name                 | Locale-dependent name                                                                | John Smith                                                                                                      |
| display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  | John Smith                                                                                                      |
| abbreviated_name              | An abbreviated form of the person's name.                                            | JS                                                                                                              |
| external_id                   | External ID that a team can attach to the user (optional)                            | (empty string if not set)                                                                                       |
| account_id                    | A user's account identifier.                                                         | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                                                        |
| device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  | desktop_client                                                                                                  |
| id                            | The session id.                                                                      | dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                                                      |
| user_agent                    | Information on the hosting device.                                                   | Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 |
| os                            | Information on the hosting operating system                                          | Windows                                                                                                         |
| browser                       | Information on the browser used for this web session.                                | Chrome                                                                                                          |
| ip_address                    | The IP address of the last activity from this session.                               | xx.xxx.x.xxx                                                                                                    |
| country                       | The country from which the last activity from this session was made.                 | United States                                                                                                   |
| created                       | The time this session was created.                                                   | 2019-09-20T23:47:33Z                                                                                            |
| updated                       | The time of the last activity from this session.                                     | 2019-10-25T04:42:16Z                                                                                            |
| expires                       | The time this session expires (optional)                                             | 2024-03-22T10:30:56Z                                                                                            |
| host_name                     | Name of the hosting desktop.                                                         | nihonbashi                                                                                                      |
| client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             | windows                                                                                                         |
| client_version                | The Dropbox client version.                                                          | 83.4.152                                                                                                        |
| platform                      | Information on the hosting platform.                                                 | Windows 10 1903                                                                                                 |
| is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             | TRUE                                                                                                            |
| device_name                   | The device name.                                                                     | My Awesome PC                                                                                                   |
| os_version                    | The hosting OS version.                                                              | (empty string if not set)                                                                                       |
| last_carrier                  | Last carrier used by the device (optional).                                          | AT&T                                                                                                            |

The first line is a header line. The program will accept a file without the header.
```
team_member_id,email,status,given_name,surname,familiar_name,display_name,abbreviated_name,external_id,account_id,device_tag,id,user_agent,os,browser,ip_address,country,created,updated,expires,host_name,client_type,client_version,platform,is_delete_on_unlink_supported,device_name,os_version,last_carrier
dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,john.smith@example.com,active,John,Smith,John Smith,John Smith,JS,(empty string if not set),dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,desktop_client,dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",Windows,Chrome,xx.xxx.x.xxx,United States,2019-09-20T23:47:33Z,2019-10-25T04:42:16Z,2024-03-22T10:30:56Z,nihonbashi,windows,83.4.152,Windows 10 1903,TRUE,My Awesome PC,(empty string if not set),AT&T
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

| Column                              | Description                                                                          |
|-------------------------------------|--------------------------------------------------------------------------------------|
| status                              | Status of the operation                                                              |
| reason                              | Reason of failure or skipped operation                                               |
| input.team_member_id                | ID of user as a member of a team.                                                    |
| input.email                         | Email address of user.                                                               |
| input.status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| input.given_name                    | Also known as a first name                                                           |
| input.surname                       | Also known as a last name or family name.                                            |
| input.display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| input.device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  |
| input.id                            | The session id.                                                                      |
| input.user_agent                    | Information on the hosting device.                                                   |
| input.os                            | Information on the hosting operating system                                          |
| input.browser                       | Information on the browser used for this web session.                                |
| input.ip_address                    | The IP address of the last activity from this session.                               |
| input.country                       | The country from which the last activity from this session was made.                 |
| input.created                       | The time this session was created.                                                   |
| input.updated                       | The time of the last activity from this session.                                     |
| input.expires                       | The time this session expires (optional)                                             |
| input.host_name                     | Name of the hosting desktop.                                                         |
| input.client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             |
| input.client_version                | The Dropbox client version.                                                          |
| input.platform                      | Information on the hosting platform.                                                 |
| input.is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             |
| input.device_name                   | The device name.                                                                     |
| input.os_version                    | The hosting OS version.                                                              |
| input.last_carrier                  | Last carrier used by the device (optional).                                          |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...


