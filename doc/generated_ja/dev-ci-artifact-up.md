# dev ci artifact up

CI成果物をアップロードします 

# Usage

This document uses the Desktop folder for command example.
## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT
```

macOS, Linux:
```
$HOME/Desktop/tbx dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option          | Description                            | Default |
|-----------------|----------------------------------------|---------|
| `-dropbox-path` | アップロード先Dropboxパス              |         |
| `-local-path`   | アップロードするローカルファイルのパス |         |
| `-peer-name`    | アカウントの別名                       | deploy  |

## Common options:

| Option            | Description                                                                                        | Default        |
|-------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`      | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`  | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage` | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`    | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`          | デバッグモードを有効にする                                                                         | false          |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: summary

このレポートはアップロード結果の概要を出力します.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column           | Description                                      |
|------------------|--------------------------------------------------|
| upload_start     | アップロード開始日時                             |
| upload_end       | アップロード終了日時                             |
| num_bytes        | 合計アップロードサイズ (バイト)                  |
| num_files_error  | 失敗またはエラーが発生したファイル数.            |
| num_files_upload | アップロード済みまたはアップロード対象ファイル数 |
| num_files_skip   | スキップ対象またはスキップ予定のファイル数       |
| num_api_call     | この処理によって消費される見積アップロードAPI数  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...
## Report: uploaded

このレポートは処理結果を出力します.
The command will generate a report in three different formats. `uploaded.csv`, `uploaded.json`, and `uploaded.xlsx`.

| Column                 | Description                                                  |
|------------------------|--------------------------------------------------------------|
| status                 | 処理の状態                                                   |
| reason                 | 失敗またはスキップの理由                                     |
| input.file             | ローカルファイルのパス                                       |
| input.size             | ローカルファイルのサイズ                                     |
| result.name            | 名称                                                         |
| result.path_display    | パス (表示目的で大文字小文字を区別する).                     |
| result.client_modified | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ |
| result.server_modified | Dropbox上で最後に更新された日時                              |
| result.size            | ファイルサイズ(バイト単位)                                   |
| result.content_hash    | ファイルコンテンツのハッシュ                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...
## Report: skipped

このレポートは処理結果を出力します.
The command will generate a report in three different formats. `skipped.csv`, `skipped.json`, and `skipped.xlsx`.

| Column                 | Description                                                  |
|------------------------|--------------------------------------------------------------|
| status                 | 処理の状態                                                   |
| reason                 | 失敗またはスキップの理由                                     |
| input.file             | ローカルファイルのパス                                       |
| input.size             | ローカルファイルのサイズ                                     |
| result.name            | 名称                                                         |
| result.path_display    | パス (表示目的で大文字小文字を区別する).                     |
| result.client_modified | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ |
| result.server_modified | Dropbox上で最後に更新された日時                              |
| result.size            | ファイルサイズ(バイト単位)                                   |
| result.content_hash    | ファイルコンテンツのハッシュ                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

