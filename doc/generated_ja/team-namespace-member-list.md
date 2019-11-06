# team namespace member list 

チームフォルダ以下のファイル・フォルダを一覧

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
.\tbx.exe team namespace member list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team namespace member list 
```

## Options

| オプション     | 説明                     | デフォルト |
|----------------|--------------------------|------------|
| `-all-columns` | 全てのカラムを表示します | false      |
| `-peer`        | アカウントの別名         | {default}  |

Common options:

| オプション     | 説明                                                                | デフォルト     |
|----------------|---------------------------------------------------------------------|----------------|
| `-concurrency` | 指定した並列度で並列処理を行います                                  | プロセッサー数 |
| `-debug`       | デバッグモードを有効にする                                          | false          |
| `-proxy`       | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                           |                |
| `-quiet`       | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します | false          |
| `-secure`      | トークンをファイルに保存しません                                    | false          |
| `-workspace`   | ワークスペースへのパス                                              |                |

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

## Report: namespace_member 

Report files are generated in `namespace_member.csv`, `namespace_member.xlsx` and `namespace_member.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `namespace_member_0000.xlsx`, `namespace_member_0001.xlsx`, `namespace_member_0002.xlsx`...   

| 列                 | 説明                                                                                                      |
|--------------------|-----------------------------------------------------------------------------------------------------------|
| namespace_name     | The name of this namespace                                                                                |
| namespace_id       | The ID of this namespace.                                                                                 |
| namespace_type     | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)                |
| entry_access_type  | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| entry_is_inherited | True if the member has access from a parent folder                                                        |
| email              | Email address of user.                                                                                    |
| display_name       | Type of the session (web_session, desktop_client, or mobile_client)                                       |
| group_name         | Name of the group                                                                                         |
| invitee_email      | Email address of invitee for this folder                                                                  |

