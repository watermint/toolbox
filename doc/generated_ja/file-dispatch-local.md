# file dispatch local

ローカルファイルを整理します (非可逆な操作です)

# Usage

This document uses the Desktop folder for command example.
## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe file dispatch local -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx file dispatch local -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option     | Description            | Default |
|------------|------------------------|---------|
| `-file`    | データファイルへのパス |         |
| `-preview` | プレビューモード       | false   |

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

# File formats

## Format: File

整理ルールのデータファイル.

| Column              | Description                           | Example                                       |
|---------------------|---------------------------------------|-----------------------------------------------|
| suffix              | ファイル名の拡張子                    | .pdf                                          |
| source_path         | 元のパス                              | <no value>/ダウンロード                       |
| source_file_pattern | 転送元ファイル名のパターン (正規表現) | toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2})      |
| dest_path_pattern   | 転送先パスのパターン                  | <no value>/ドキュメント/<no value>-<no value> |
| dest_file_pattern   | 転送先ファイル名のパターン            | TBX_<no value>-<no value>-<no value>          |

The first line is a header line. The program will accept file without the header.
```
suffix,source_path,source_file_pattern,dest_path_pattern,dest_file_pattern
.pdf,<no value>/ダウンロード,toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2}),<no value>/ドキュメント/<no value>-<no value>,TBX_<no value>-<no value>-<no value>
```

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

