# member update email 

Member email operation 

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

This command use following access type(s) during the operation:

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe member update email -file /path/to/data/file.csv
```

macOS, Linux:

```bash
$HOME/Desktop/tbx member update email -file /path/to/data/file.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option               | Description                                                                                                                                | Default |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-file`              | Data file                                                                                                                                  |         |
| `-peer`              | Account alias                                                                                                                              | default |
| `-update-unverified` | Update an account which didn't verified email. If an account email unverified, email address change may affect lose invitation to folders. | false   |

Common options:

| Option          | Description                                                                      | Default              |
|-----------------|----------------------------------------------------------------------------------|----------------------|
| `-auto-open`    | Auto open URL or artifact folder                                                 | false                |
| `-bandwidth-kb` | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited  | 0                    |
| `-concurrency`  | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`        | Enable debug mode                                                                | false                |
| `-low-memory`   | Low memory footprint mode                                                        | false                |
| `-output`       | Output format (none/text/markdown/json)                                          | text                 |
| `-proxy`        | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`        | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`       | Do not store tokens into a file                                                  | false                |
| `-workspace`    | Workspace path                                                                   |                      |

# File formats

## Format: File

Data file for updating team member email addresses. 

| Column     | Description           | Value example          |
|------------|-----------------------|------------------------|
| from_email | Current Email address | john@example.com       |
| to_email   | New Email address     | john.smith@example.net |

The first line is a header line. The program will accept file without the header.

```csv
from_email,to_email
john@example.com,john.smith@example.net
```

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.

```

watermint toolbox xx.x.xxx
==========================

© 2016-2020 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code

```

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: operation_log 
This report shows the transaction result.
Report files are generated in three formats like below;
* `operation_log.csv`
* `operation_log.xlsx`
* `operation_log.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`...   

| Column                | Description                                                                                    |
|-----------------------|------------------------------------------------------------------------------------------------|
| status                | Status of the operation                                                                        |
| reason                | Reason of failure or skipped operation                                                         |
| input.from_email      | Current Email address                                                                          |
| input.to_email        | New Email address                                                                              |
| result.email          | Email address of user.                                                                         |
| result.email_verified | Is true if the user's email is verified to be owned by the user.                               |
| result.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| result.given_name     | Also known as a first name                                                                     |
| result.surname        | Also known as a last name or family name.                                                      |
| result.display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| result.joined_on      | The date and time the user joined as a member of a specific team.                              |
| result.role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| result.tag            | Operation tag                                                                                  |

