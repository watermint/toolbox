# team device list 

チーム内全てのデバイス/セッションを一覧します

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
.\tbx.exe team device list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team device list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

## Options

| オプション | 説明             | デフォルト |
|------------|------------------|------------|
| `-peer`    | アカウントの別名 | {default}  |

Common options:

| オプション     | 説明                                                                   | デフォルト     |
|----------------|------------------------------------------------------------------------|----------------|
| `-bandwidth`   | {"key":"infra.control.app_opt.common_opts.flag.bandwidth","params":{}} | 0              |
| `-concurrency` | 指定した並列度で並列処理を行います                                     | プロセッサー数 |
| `-debug`       | デバッグモードを有効にする                                             | false          |
| `-proxy`       | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                              |                |
| `-quiet`       | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します    | false          |
| `-secure`      | トークンをファイルに保存しません                                       | false          |
| `-workspace`   | ワークスペースへのパス                                                 |                |

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

## Report: device 

Report files are generated in `device.csv`, `device.xlsx` and `device.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`...   

| 列                            | 説明                                                                                 |
|-------------------------------|--------------------------------------------------------------------------------------|
| team_member_id                | ID of user as a member of a team.                                                    |
| email                         | Email address of user.                                                               |
| status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name                    | Also known as a first name                                                           |
| surname                       | Also known as a last name or family name.                                            |
| familiar_name                 | Locale-dependent name                                                                |
| display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| abbreviated_name              | An abbreviated form of the person's name.                                            |
| external_id                   | External ID that a team can attach to the user.                                      |
| account_id                    | A user's account identifier.                                                         |
| device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  |
| id                            | The session id.                                                                      |
| user_agent                    | Information on the hosting device.                                                   |
| os                            | Information on the hosting operating system                                          |
| browser                       | Information on the browser used for this web session.                                |
| ip_address                    | The IP address of the last activity from this session.                               |
| country                       | The country from which the last activity from this session was made.                 |
| created                       | The time this session was created.                                                   |
| updated                       | The time of the last activity from this session.                                     |
| expires                       | The time this session expires                                                        |
| host_name                     | Name of the hosting desktop.                                                         |
| client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             |
| client_version                | The Dropbox client version.                                                          |
| platform                      | Information on the hosting platform.                                                 |
| is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             |
| device_name                   | The device name.                                                                     |
| os_version                    | The hosting OS version.                                                              |
| last_carrier                  | Last carrier used by the device.                                                     |

