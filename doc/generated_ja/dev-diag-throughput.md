# dev diag throughput

キャプチャログからスループットを評価 

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.
## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe dev diag throughput 
```

macOS, Linux:
```
$HOME/Desktop/tbx dev diag throughput 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション              | 説明                                                    | デフォルト              |
|-------------------------|---------------------------------------------------------|-------------------------|
| `-bucket`               | バケットサイズ (ミリ秒)                                 | 1000                    |
| `-endpoint-name`        | Filter by endpoint. Filter by exact match to the name.  |                         |
| `-endpoint-name-prefix` | Filter by endpoint. Filter by name match to the prefix. |                         |
| `-endpoint-name-suffix` | Filter by endpoint. Filter by name match to the suffix. |                         |
| `-job-id`               | ジョブIDの指定                                          |                         |
| `-path`                 | ワークスペースへのパス.                                 |                         |
| `-time-format`          | 日時フォーマット (Goの日付フォーマット)                 | 2006-01-02 15:04:05.999 |

## 共通のオプション:

| オプション        | 説明                                                                                               | デフォルト     |
|-------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`      | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`  | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage` | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`    | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`          | デバッグモードを有効にする                                                                         | false          |
| `-experiment`     | 実験的機能を有効化する                                                                             |                |
| `-lang`           | 表示言語                                                                                           | auto           |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: report

スループット
このコマンドはレポートを3種類の書式で出力します. `report.csv`, `report.json`, ならびに `report.xlsx`.

| 列                  | 説明                                             |
|---------------------|--------------------------------------------------|
| time                | タイムスタンプ                                   |
| concurrency         | 並列数                                           |
| success_concurrency | 成功したリクエストの並列数                       |
| success_sent        | 成功したリクエストのバケットあたりの送信バイト数 |
| success_received    | 成功したリクエストのバケットあたりの受信バイト数 |
| failure_concurrency | 失敗したリクエストの並列数                       |
| failure_sent        | 失敗したリクエストのバケットあたりの送信バイト数 |
| failure_received    | 失敗したリクエストのバケットあたりの受信バイト数 |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `report_0000.xlsx`, `report_0001.xlsx`, `report_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

