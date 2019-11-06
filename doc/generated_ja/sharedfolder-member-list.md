# sharedfolder member list 

共有フォルダのメンバーを一覧します

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
* Individual account: https://help.dropbox.com/ja-jp/installs-integrations/third-party/third-party-apps

This command use following access type(s) during the operation:
* Dropbox Full access

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe sharedfolder member list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx sharedfolder member list 
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

## Report: sharedfolder_member 

Report files are generated in `sharedfolder_member.csv`, `sharedfolder_member.xlsx` and `sharedfolder_member.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `sharedfolder_member_0000.xlsx`, `sharedfolder_member_0001.xlsx`, `sharedfolder_member_0002.xlsx`...   

| 列                      | 説明                                                                                                                    |
|-------------------------|-------------------------------------------------------------------------------------------------------------------------|
| shared_folder_id        | The ID of the shared folder.                                                                                            |
| parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| name                    | The name of the this shared folder.                                                                                     |
| path_lower              | The lower-cased full path of this shared folder.                                                                        |
| is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| is_team_folder          | Whether this folder is a team folder.                                                                                   |
| access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| is_inherited            | True if the member has access from a parent folder                                                                      |
| account_id              | A user's account identifier.                                                                                            |
| email                   | Email address of user.                                                                                                  |
| display_name            | A name that can be used directly to represent the name of a user's Dropbox account.                                     |
| group_name              | Name of a group                                                                                                         |
| group_id                | A group's identifier                                                                                                    |
| invitee_email           | Email address of invitee for this folder                                                                                |

