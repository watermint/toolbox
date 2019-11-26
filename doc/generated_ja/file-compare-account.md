# file compare account 

二つのアカウントのファイルを比較します

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
.\tbx.exe file compare account -left left -left-path /比較対象パス -right right -right-path /比較対象パス
```

macOS, Linux:

```bash
$HOME/Desktop/tbx file compare account -left left -left-path /比較対象パス -right right -right-path /比較対象パス
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

同じアカウント内の別パス同士を比較したい場合には `-left` と `-right` に同じ別名を指定してください

## Options

| オプション    | 説明                              | デフォルト |
|---------------|-----------------------------------|------------|
| `-left`       | アカウントの別名 (左)             | {default}  |
| `-left-path`  | アカウントのルートからのパス (左) |            |
| `-right`      | アカウントの別名 (右)             | {default}  |
| `-right-path` | アカウントのルートからのパス (右) |            |

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

## Report: diff 

Report files are generated in `diff.csv`, `diff.xlsx` and `diff.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `diff_0000.xlsx`, `diff_0001.xlsx`, `diff_0002.xlsx`...   

| 列         | 説明                                                                                                                                                                                           |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | 差分のタイプ`file_content_diff`: コンテンツハッシュの差分, `{left|right}_file_missing`: 左または右のファイルが見つからない, `{left|right}_folder_missing`: 左または右のフォルダが見つからない. |
| left_path  | 左のパス                                                                                                                                                                                       |
| left_kind  | フォルダまたはファイル                                                                                                                                                                         |
| left_size  | 左ファイルのサイズ                                                                                                                                                                             |
| left_hash  | 左ファイルのコンテンツハッシュ                                                                                                                                                                 |
| right_path | 右のパス                                                                                                                                                                                       |
| right_kind | フォルダまたはファイル                                                                                                                                                                         |
| right_size | 右ファイルのサイズ                                                                                                                                                                             |
| right_hash | 右ファイルのコンテンツハッシュ                                                                                                                                                                 |

