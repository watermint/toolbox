# team sharedlink update expiry 

チーム内の公開されている共有リンクについて有効期限を更新します

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
.\tbx.exe team sharedlink update expiry -days 28
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team sharedlink update expiry -days 28
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

## Options

| オプション    | 説明                                                                          | デフォルト |
|---------------|-------------------------------------------------------------------------------|------------|
| `-at`         | 新しい有効期限の日時                                                          |            |
| `-days`       | 新しい有効期限までの日時                                                      | 0          |
| `-peer`       | アカウントの別名                                                              | {default}  |
| `-visibility` | {"key":"recipe.team.sharedlink.update.expiry_vo.flag.visibility","params":{}} | public     |

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

## Report: skipped_sharedlink 

Report files are generated in `skipped_sharedlink.csv`, `skipped_sharedlink.xlsx` and `skipped_sharedlink.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `skipped_sharedlink_0000.xlsx`, `skipped_sharedlink_0001.xlsx`, `skipped_sharedlink_0002.xlsx`...   

| 列         | 説明                                                                                                                                                                                                                    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| tag        | Entry type (file, or folder)                                                                                                                                                                                            |
| url        | URL of the shared link.                                                                                                                                                                                                 |
| name       | The linked file name (including extension).                                                                                                                                                                             |
| expires    | Expiration time, if set.                                                                                                                                                                                                |
| path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                         |
| visibility | The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| email      | Email address of user.                                                                                                                                                                                                  |
| status     | The user's status as a member of a specific team. (active/invited/suspended/removed)                                                                                                                                    |
| surname    | Surname of the link owner                                                                                                                                                                                               |
| given_name | Given name of the link owner                                                                                                                                                                                            |

## Report: updated_sharedlink 

Report files are generated in `updated_sharedlink.csv`, `updated_sharedlink.xlsx` and `updated_sharedlink.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `updated_sharedlink_0000.xlsx`, `updated_sharedlink_0001.xlsx`, `updated_sharedlink_0002.xlsx`...   

| 列                   | 説明                                                                                                                                                                                                                    |
|----------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status               | Status of the operation                                                                                                                                                                                                 |
| reason               | Reason of failure or skipped operation                                                                                                                                                                                  |
| input.shared_link_id | A unique identifier for the linked file or folder                                                                                                                                                                       |
| input.tag            | Entry type (file, or folder)                                                                                                                                                                                            |
| input.url            | URL of the shared link.                                                                                                                                                                                                 |
| input.name           | The linked file name (including extension).                                                                                                                                                                             |
| input.expires        | Expiration time, if set.                                                                                                                                                                                                |
| input.path_lower     | The lowercased full path in the user's Dropbox.                                                                                                                                                                         |
| input.visibility     | The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| input.account_id     | A user's account identifier.                                                                                                                                                                                            |
| input.team_member_id | ID of user as a member of a team.                                                                                                                                                                                       |
| input.email          | Email address of user.                                                                                                                                                                                                  |
| input.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)                                                                                                                                    |
| input.surname        | Surname of the link owner                                                                                                                                                                                               |
| input.given_name     | Given name of the link owner                                                                                                                                                                                            |
| result.id            | A unique identifier for the linked file or folder                                                                                                                                                                       |
| result.tag           | Entry type (file, or folder)                                                                                                                                                                                            |
| result.url           | URL of the shared link.                                                                                                                                                                                                 |
| result.name          | The linked file name (including extension).                                                                                                                                                                             |
| result.expires       | Expiration time, if set.                                                                                                                                                                                                |
| result.path_lower    | The lowercased full path in the user's Dropbox.                                                                                                                                                                         |
| result.visibility    | The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
