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
* Individual account: https://help.dropbox.com/installs-integrations/third-party/third-party-apps

This command use following access type(s) during the operation:

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

| オプション      | 説明                     | デフォルト |
|-----------------|--------------------------|------------|
| `-dropbox-path` | Destination Dropbox path |            |
| `-local-path`   | Local file path          |            |
| `-peer`         | Account alias            | default    |

Common options:

| オプション      | 説明                                                                                             | デフォルト     |
|-----------------|--------------------------------------------------------------------------------------------------|----------------|
| `-bandwidth-kb` | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒)0の場合、制限を行わない | 0              |
| `-concurrency`  | 指定した並列度で並列処理を行います                                                               | プロセッサー数 |
| `-debug`        | デバッグモードを有効にする                                                                       | false          |
| `-low-memory`   | Low memory footprint mode                                                                        | false          |
| `-proxy`        | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                        |                |
| `-quiet`        | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                              | false          |
| `-secure`       | トークンをファイルに保存しません                                                                 | false          |
| `-workspace`    | ワークスペースへのパス                                                                           |                |

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

## Report: skipped 

Report files are generated in three formats, `skipped.csv`, `skipped.xlsx` and `skipped.json`.
But if you run with `-low-memory` option, the command will generate only `skipped.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`...   

| 列                             | 説明                                                                                                   |
|--------------------------------|--------------------------------------------------------------------------------------------------------|
| status                         | 処理の状態                                                                                             |
| reason                         | 失敗またはスキップの理由                                                                               |
| input.file                     | Local file path                                                                                        |
| input.size                     | Local file size                                                                                        |
| result.id                      | A unique identifier for the file.                                                                      |
| result.tag                     | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.name                    | The last component of the path (including extension).                                                  |
| result.path_lower              | The lowercased full path in the user's Dropbox. This always starts with a slash.                       |
| result.path_display            | The cased path to be used for display purposes only.                                                   |
| result.client_modified         | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified         | The last time the file was modified on Dropbox.                                                        |
| result.revision                | A unique identifier for the current revision of a file.                                                |
| result.size                    | The file size in bytes.                                                                                |
| result.content_hash            | A hash of the file content.                                                                            |
| result.shared_folder_id        | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.   |
| result.parent_shared_folder_id | ID of shared folder that holds this file.                                                              |

## Report: summary 

Report files are generated in three formats, `summary.csv`, `summary.xlsx` and `summary.json`.
But if you run with `-low-memory` option, the command will generate only `summary.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`...   

| 列               | 説明                                                |
|------------------|-----------------------------------------------------|
| upload_start     | Time of start uploading                             |
| upload_end       | Time of finish uploading                            |
| num_bytes        | Total upload size (Bytes)                           |
| num_files_error  | The number of files failed or got an error.         |
| num_files_upload | The number of files uploaded or to upload.          |
| num_files_skip   | The number of files skipped or to skip.             |
| num_api_call     | The number of estimated upload API call for upload. |

## Report: uploaded 

Report files are generated in three formats, `uploaded.csv`, `uploaded.xlsx` and `uploaded.json`.
But if you run with `-low-memory` option, the command will generate only `uploaded.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`...   

| 列                             | 説明                                                                                                   |
|--------------------------------|--------------------------------------------------------------------------------------------------------|
| status                         | 処理の状態                                                                                             |
| reason                         | 失敗またはスキップの理由                                                                               |
| input.file                     | Local file path                                                                                        |
| input.size                     | Local file size                                                                                        |
| result.id                      | A unique identifier for the file.                                                                      |
| result.tag                     | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.name                    | The last component of the path (including extension).                                                  |
| result.path_lower              | The lowercased full path in the user's Dropbox. This always starts with a slash.                       |
| result.path_display            | The cased path to be used for display purposes only.                                                   |
| result.client_modified         | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified         | The last time the file was modified on Dropbox.                                                        |
| result.revision                | A unique identifier for the current revision of a file.                                                |
| result.size                    | The file size in bytes.                                                                                |
| result.content_hash            | A hash of the file content.                                                                            |
| result.shared_folder_id        | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.   |
| result.parent_shared_folder_id | ID of shared folder that holds this file.                                                              |

