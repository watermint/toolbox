# teamfolder replication 

チームフォルダを他のチームに複製します

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe teamfolder replication 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx teamfolder replication 
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
| `-name`    | Team folder name |            |

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

## Report: verification 

Report files are generated in three formats, `verification.csv`, `verification.xlsx` and `verification.json`.
But if you run with `-low-memory` option, the command will generate only `verification.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `verification_0000.xlsx`, `verification_0001.xlsx`, `verification_0002.xlsx`...   

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

