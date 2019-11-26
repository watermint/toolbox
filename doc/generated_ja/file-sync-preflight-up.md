# file sync preflight up 

上り方向同期のための事前チェックを実施します

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
.\tbx.exe file sync preflight up 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx file sync preflight up 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

## Options

| オプション      | 説明                   | デフォルト |
|-----------------|------------------------|------------|
| `-dropbox-path` | 転送先のDropboxパス    |            |
| `-local-path`   | ローカルファイルのパス |            |
| `-peer`         | アカウントの別名       | {default}  |

Common options:

| オプション     | 説明                                                                                             | デフォルト     |
|----------------|--------------------------------------------------------------------------------------------------|----------------|
| `-bandwidth`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒)0の場合、制限を行わない | 0              |
| `-concurrency` | 指定した並列度で並列処理を行います                                                               | プロセッサー数 |
| `-debug`       | デバッグモードを有効にする                                                                       | false          |
| `-proxy`       | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                        |                |
| `-quiet`       | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                              | false          |
| `-secure`      | トークンをファイルに保存しません                                                                 | false          |
| `-workspace`   | ワークスペースへのパス                                                                           |                |

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

## Report: skip 

Report files are generated in `skip.csv`, `skip.xlsx` and `skip.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `skip_0000.xlsx`, `skip_0001.xlsx`, `skip_0002.xlsx`...   

| 列                             | 説明                                                                                       |
|--------------------------------|--------------------------------------------------------------------------------------------|
| status                         | 処理の状態                                                                                 |
| reason                         | 失敗またはスキップの理由                                                                   |
| input.file                     | ローカルファイルのパス                                                                     |
| input.size                     | ローカルファイルのサイズ                                                                   |
| result.id                      | ファイルへの一意なID                                                                       |
| result.tag                     | エントリーの種別`file`, `folder`, または `deleted`                                         |
| result.name                    | 名称                                                                                       |
| result.path_lower              | パス (すべて小文字に変換). これは常にスラッシュで始まります.                               |
| result.path_display            | パス (表示目的で大文字小文字を区別する).                                                   |
| result.client_modified         | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ                               |
| result.server_modified         | Dropbox上で最後に更新された日時                                                            |
| result.revision                | ファイルの現在バージョンの一意な識別子                                                     |
| result.size                    | ファイルサイズ(バイト単位)                                                                 |
| result.content_hash            | ファイルコンテンツのハッシュ                                                               |
| result.shared_folder_id        | これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダのID。 |
| result.parent_shared_folder_id | このファイルを含む共有フォルダのID.                                                        |

## Report: summary 

Report files are generated in `summary.csv`, `summary.xlsx` and `summary.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`...   

| 列               | 説明                                             |
|------------------|--------------------------------------------------|
| num_bytes        | 合計アップロードサイズ (バイト)                  |
| num_files_error  | 失敗またはエラーが発生したファイル数.            |
| num_files_upload | アップロード済みまたはアップロード対象ファイル数 |
| num_files_skip   | スキップ対象またはスキップ予定のファイル数       |
| num_api_call     | この処理によって消費される見積アップロードAPI数  |

## Report: upload 

Report files are generated in `upload.csv`, `upload.xlsx` and `upload.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `upload_0000.xlsx`, `upload_0001.xlsx`, `upload_0002.xlsx`...   

| 列                             | 説明                                                                                       |
|--------------------------------|--------------------------------------------------------------------------------------------|
| status                         | 処理の状態                                                                                 |
| reason                         | 失敗またはスキップの理由                                                                   |
| input.file                     | ローカルファイルのパス                                                                     |
| input.size                     | ローカルファイルのサイズ                                                                   |
| result.id                      | ファイルへの一意なID                                                                       |
| result.tag                     | エントリーの種別`file`, `folder`, または `deleted`                                         |
| result.name                    | 名称                                                                                       |
| result.path_lower              | パス (すべて小文字に変換). これは常にスラッシュで始まります.                               |
| result.path_display            | パス (表示目的で大文字小文字を区別する).                                                   |
| result.client_modified         | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ                               |
| result.server_modified         | Dropbox上で最後に更新された日時                                                            |
| result.revision                | ファイルの現在バージョンの一意な識別子                                                     |
| result.size                    | ファイルサイズ(バイト単位)                                                                 |
| result.content_hash            | ファイルコンテンツのハッシュ                                                               |
| result.shared_folder_id        | これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダのID。 |
| result.parent_shared_folder_id | このファイルを含む共有フォルダのID.                                                        |

