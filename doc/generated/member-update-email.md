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
You can delete those files after use if you want to remove it.
If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/ja-jp/teams-admins/admin/app-integrations

This command use following access type(s) during the operation:
* Dropbox Business management

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe member update email 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx member update email 
```

## Options

| Option                    | Description                                          | Default   |
|---------------------------|------------------------------------------------------|-----------|
| `-dont-update-unverified` | Do not update an account which didn't verified email | true      |
| `-file`                   | Data file                                            |           |
| `-peer`                   | Account alias                                        | {default} |

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

## Network configuration: Proxy

The executable automatically detects your proxy configuration from the environment.
However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port.
Currently, the executable doesn't support proxies which require authentication.

# Result

Report file path will be displayed last line of the command line output.
If you missed command line output, please see path below.
[job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: update 

Report files are generated in `update.csv`, `update.xlsx` and `update.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `update_0000.xlsx`, `update_0001.xlsx`, `update_0002.xlsx`...   

| Column                  | Description                                                                                                          |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                              |
| reason                  | Reason of failure or skipped operation                                                                               |
| input.from_email        | Current Email address                                                                                                |
| input.to_email          | New Email address                                                                                                    |
| result.team_member_id   | ID of user as a member of a team.                                                                                    |
| result.email            | Email address of user.                                                                                               |
| result.email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| result.status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| result.given_name       | Also known as a first name                                                                                           |
| result.surname          | Also known as a last name or family name.                                                                            |
| result.familiar_name    | Locale-dependent name                                                                                                |
| result.display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| result.abbreviated_name | An abbreviated form of the person's name.                                                                            |
| result.member_folder_id | The namespace id of the user's root folder.                                                                          |
| result.external_id      | External ID that a team can attach to the user.                                                                      |
| result.account_id       | A user's account identifier.                                                                                         |
| result.persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| result.joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| result.role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| result.tag              | Operation tag                                                                                                        |

