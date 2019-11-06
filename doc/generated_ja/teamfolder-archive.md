# teamfolder archive 

チームフォルダのアーカイブ

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
.\tbx.exe teamfolder archive -name チームフォルダ名
```

macOS, Linux:

```bash
$HOME/Desktop/tbx teamfolder archive -name チームフォルダ名
```

## Options

| オプション | 説明             | デフォルト |
|------------|------------------|------------|
| `-name`    | チームフォルダ名 |            |
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

