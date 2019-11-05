# team filerequest list 

List all file requests in the team

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it.
If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/ja-jp/teams-admins/admin/app-integrations

This command use following access type(s) during the operation:
* Dropbox Business File access

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team filerequest list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team filerequest list 
```

## Options

| Option  | Description   | Default   |
|---------|---------------|-----------|
| `-peer` | Account alias | {default} |

Common options:

| Option         | Description                                                                      | Default              |
|----------------|----------------------------------------------------------------------------------|----------------------|
| `-concurrency` | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`       | Enable debug mode                                                                | false                |
| `-proxy`       | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`       | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`      | Do not store tokens into a file                                                  | false                |
| `-workspace`   | Workspace path                                                                   |                      |

## Authentication

For the first run, `toolbox` will ask you an authentication with your Dropbox account. 
Please copy the link and paste it into your browser. Then proceed to authorization.
After authorization, Dropbox will show you an authorization code.
Please copy that code and paste it to the `toolbox`.

```
watermint toolbox xx.x.xxx
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Testing network connection...
Done

1. Visit the URL for the auth dialog:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

# Result

Report file path will be displayed last line of the command line output.
If you missed command line output, please see path below.
[job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: file_request 

Report files are generated in `file_request.csv`, `file_request.xlsx` and `file_request.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`...   

| Column                      | Description                                                                   |
|-----------------------------|-------------------------------------------------------------------------------|
| account_id                  | Account ID of this file request owner.                                        |
| team_member_id              | ID of file request owner user as a member of a team                           |
| email                       | Email address of this file request owner.                                     |
| status                      | The user status of this file request owner (active/invited/suspended/removed) |
| surname                     | Surname of this file request owner.                                           |
| given_name                  | Given name of this file request owner.                                        |
| file_request_id             | The ID of the file request.                                                   |
| url                         | The URL of the file request.                                                  |
| title                       | The title of the file request.                                                |
| created                     | When this file request was created.                                           |
| is_open                     | Whether or not the file request is open.                                      |
| file_count                  | The number of files this file request has received.                           |
| destination                 | The path of the folder in the Dropbox where uploaded files will be sent       |
| deadline                    | The deadline for this file request.                                           |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           |

