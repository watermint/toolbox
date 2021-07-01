---
layout: command
title: Command
lang: en
---

# team filerequest clone

Clone file requests by given data (Experimental, and Irreversible operation)

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
* Dropbox Business: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## Auth scopes

| Description                  |
|------------------------------|
| Dropbox Business File access |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2021 Takayuki Okazaki
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
.\tbx.exe team filerequest clone -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx team filerequest clone -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option  | Description       | Default |
|---------|-------------------|---------|
| `-file` | Path to data file |         |
| `-peer` | Account alias     | default |

## Common options:

| Option            | Description                                                                               | Default              |
|-------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auto-open`      | Auto open URL or artifact folder                                                          | false                |
| `-bandwidth-kb`   | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited           | 0                    |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                            | normal               |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)                      | normal               |
| `-concurrency`    | Maximum concurrency for running operation                                                 | Number of processors |
| `-debug`          | Enable debug mode                                                                         | false                |
| `-experiment`     | Enable experimental feature(s).                                                           |                      |
| `-extra`          | Extra parameter file path                                                                 |                      |
| `-lang`           | Display language                                                                          | auto                 |
| `-output`         | Output format (none/text/markdown/json)                                                   | text                 |
| `-proxy`          | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy. |                      |
| `-quiet`          | Suppress non-error messages, and make output readable by a machine (JSON format)          | false                |
| `-secure`         | Do not store tokens into a file                                                           | false                |
| `-verbose`        | Show current operations for more detail.                                                  | false                |
| `-workspace`      | Workspace path                                                                            |                      |

# File formats

## Format: File

This report shows a list of file requests with the file request owner team member.

| Column                      | Description                                                                   | Example                                            |
|-----------------------------|-------------------------------------------------------------------------------|----------------------------------------------------|
| account_id                  | Account ID of this file request owner.                                        | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx           |
| team_member_id              | ID of file request owner user as a member of a team                           | dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx          |
| email                       | Email address of this file request owner.                                     | john@example.com                                   |
| status                      | The user status of this file request owner (active/invited/suspended/removed) | active                                             |
| surname                     | Surname of this file request owner.                                           | Smith                                              |
| given_name                  | Given name of this file request owner.                                        | John                                               |
| file_request_id             | The ID of the file request.                                                   | xxxxxxxxxxxxxxxxxx                                 |
| url                         | The URL of the file request.                                                  | https://www.dropbox.com/request/xxxxxxxxxxxxxxxxxx |
| title                       | The title of the file request.                                                | Photo contest submission                           |
| created                     | When this file request was created.                                           | 2019-09-20T23:47:33Z                               |
| is_open                     | Whether or not the file request is open.                                      | true                                               |
| file_count                  | The number of files this file request has received.                           | 3                                                  |
| destination                 | The path of the folder in the Dropbox where uploaded files will be sent       | /Photo contest entries                             |
| deadline                    | The deadline for this file request.                                           | 2019-10-20T23:47:33Z                               |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           | seven_days                                         |

The first line is a header line. The program will accept a file without the header.
```
account_id,team_member_id,email,status,surname,given_name,file_request_id,url,title,created,is_open,file_count,destination,deadline,deadline_allow_late_uploads
dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,john@example.com,active,Smith,John,xxxxxxxxxxxxxxxxxx,https://www.dropbox.com/request/xxxxxxxxxxxxxxxxxx,Photo contest submission,2019-09-20T23:47:33Z,true,3,/Photo contest entries,2019-10-20T23:47:33Z,seven_days
```

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                             | Description                                                                   |
|------------------------------------|-------------------------------------------------------------------------------|
| status                             | Status of the operation                                                       |
| reason                             | Reason of failure or skipped operation                                        |
| input.email                        | Email address of this file request owner.                                     |
| input.status                       | The user status of this file request owner (active/invited/suspended/removed) |
| input.surname                      | Surname of this file request owner.                                           |
| input.given_name                   | Given name of this file request owner.                                        |
| input.url                          | The URL of the file request.                                                  |
| input.title                        | The title of the file request.                                                |
| input.created                      | When this file request was created.                                           |
| input.is_open                      | Whether or not the file request is open.                                      |
| input.file_count                   | The number of files this file request has received.                           |
| input.destination                  | The path of the folder in the Dropbox where uploaded files will be sent       |
| input.deadline                     | The deadline for this file request.                                           |
| input.deadline_allow_late_uploads  | If set, allow uploads after the deadline has passed                           |
| result.email                       | Email address of this file request owner.                                     |
| result.status                      | The user status of this file request owner (active/invited/suspended/removed) |
| result.surname                     | Surname of this file request owner.                                           |
| result.given_name                  | Given name of this file request owner.                                        |
| result.url                         | The URL of the file request.                                                  |
| result.title                       | The title of the file request.                                                |
| result.created                     | When this file request was created.                                           |
| result.is_open                     | Whether or not the file request is open.                                      |
| result.file_count                  | The number of files this file request has received.                           |
| result.destination                 | The path of the folder in the Dropbox where uploaded files will be sent       |
| result.deadline                    | The deadline for this file request.                                           |
| result.deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


