# team device unlink 

Unlink device sessions

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
.\tbx.exe team device unlink -file /path/to/data/file.csv
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team device unlink -file /path/to/data/file.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

## Options

| Option              | Description            | Default   |
|---------------------|------------------------|-----------|
| `-delete-on-unlink` | Delete files on unlink | false     |
| `-file`             | Data file              |           |
| `-peer`             | Account alias          | {default} |

Common options:

| Option          | Description                                                                      | Default              |
|-----------------|----------------------------------------------------------------------------------|----------------------|
| `-bandwidth-kb` | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited  | 0                    |
| `-concurrency`  | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`        | Enable debug mode                                                                | false                |
| `-low-memory`   | Low memory footprint mode                                                        | false                |
| `-proxy`        | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`        | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`       | Do not store tokens into a file                                                  | false                |
| `-workspace`    | Workspace path                                                                   |                      |

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

## Report: unlink 

Report files are generated in `unlink.csv`, `unlink.xlsx` and `unlink.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `unlink_0000.xlsx`, `unlink_0001.xlsx`, `unlink_0002.xlsx`...   

| Column                              | Description                                                                          |
|-------------------------------------|--------------------------------------------------------------------------------------|
| status                              | Status of the operation                                                              |
| reason                              | Reason of failure or skipped operation                                               |
| input.team_member_id                | ID of user as a member of a team.                                                    |
| input.email                         | Email address of user.                                                               |
| input.status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| input.given_name                    | Also known as a first name                                                           |
| input.surname                       | Also known as a last name or family name.                                            |
| input.familiar_name                 | Locale-dependent name                                                                |
| input.display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| input.abbreviated_name              | An abbreviated form of the person's name.                                            |
| input.external_id                   | External ID that a team can attach to the user.                                      |
| input.account_id                    | A user's account identifier.                                                         |
| input.device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  |
| input.id                            | The session id.                                                                      |
| input.user_agent                    | Information on the hosting device.                                                   |
| input.os                            | Information on the hosting operating system                                          |
| input.browser                       | Information on the browser used for this web session.                                |
| input.ip_address                    | The IP address of the last activity from this session.                               |
| input.country                       | The country from which the last activity from this session was made.                 |
| input.created                       | The time this session was created.                                                   |
| input.updated                       | The time of the last activity from this session.                                     |
| input.expires                       | The time this session expires                                                        |
| input.host_name                     | Name of the hosting desktop.                                                         |
| input.client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             |
| input.client_version                | The Dropbox client version.                                                          |
| input.platform                      | Information on the hosting platform.                                                 |
| input.is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             |
| input.device_name                   | The device name.                                                                     |
| input.os_version                    | The hosting OS version.                                                              |
| input.last_carrier                  | Last carrier used by the device.                                                     |

