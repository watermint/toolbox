---
layout: command
title: コマンド
lang: ja
---

# dev ci artifact up

CI成果物をアップロードします 

# インストール

[最新リリース](https://github.com/watermint/toolbox/releases/latest)からコンパイル済みのバイナリをダウンロードしてください. Windowsをお使いの方は、`tbx-xx.x.xxx-win.zip`のようなzipファイルをダウンロードしてください. その後、アーカイブを解凍し、デスクトップ フォルダに `tbx.exe` を配置します.
watermint toolboxは、システムで許可されていれば、システム内のどのパスからでも実行できます. しかし、説明書のサンプルでは、デスクトップ フォルダを使用しています. デスクトップ フォルダ以外にバイナリを配置した場合は、パスを読み替えてください.

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT
```

macOS, Linux:
```
$HOME/Desktop/tbx dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション      | 説明                                   | デフォルト |
|-----------------|----------------------------------------|------------|
| `-dropbox-path` | アップロード先Dropboxパス              |            |
| `-local-path`   | アップロードするローカルファイルのパス |            |
| `-peer-name`    | アカウントの別名                       | deploy     |
| `-timeout`      | 処理タイムアウト(秒)                   | 60         |

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
| `-extra`          | 追加パラメータファイルのパス                                                                       |                |
| `-lang`           | 表示言語                                                                                           | auto           |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください    |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-verbose`        | 現在の操作を詳細に表示します.                                                                      | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: deleted

パス
このコマンドはレポートを3種類の書式で出力します. `deleted.csv`, `deleted.json`, ならびに `deleted.xlsx`.

| 列                           | 説明                   |
|------------------------------|------------------------|
| entry_path                   | パス                   |
| entry_shard.file_system_type | ファイルシステムの種別 |
| entry_shard.shard_id         | シャードID             |
| entry_shard.attributes       | シャードの属性         |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

## レポート: skipped

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `skipped.csv`, `skipped.json`, ならびに `skipped.xlsx`.

| 列                                 | 説明                     |
|------------------------------------|--------------------------|
| status                             | 処理の状態               |
| reason                             | 失敗またはスキップの理由 |
| input.entry_path                   | パス                     |
| input.entry_shard.file_system_type | ファイルシステムの種別   |
| input.entry_shard.shard_id         | シャードID               |
| input.entry_shard.attributes       | シャードの属性           |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## レポート: summary

このレポートはアップロード結果の概要を出力します.
このコマンドはレポートを3種類の書式で出力します. `summary.csv`, `summary.json`, ならびに `summary.xlsx`.

| 列                    | 説明                                            |
|-----------------------|-------------------------------------------------|
| start                 | 開始時間                                        |
| end                   | 完了時間                                        |
| num_bytes             | 合計アップロードサイズ (バイト)                 |
| num_files_error       | 失敗またはエラーが発生したファイル数.           |
| num_files_transferred | アップロード/ダウンロードされたファイル数.      |
| num_files_skip        | スキップ対象またはスキップ予定のファイル数      |
| num_folder_created    | 作成されたフォルダ数.                           |
| num_delete            | 削除されたエントリ数.                           |
| num_api_call          | この処理によって消費される見積アップロードAPI数 |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

## レポート: uploaded

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `uploaded.csv`, `uploaded.json`, ならびに `uploaded.xlsx`.

| 列                     | 説明                                                         |
|------------------------|--------------------------------------------------------------|
| status                 | 処理の状態                                                   |
| reason                 | 失敗またはスキップの理由                                     |
| input.path             | パス                                                         |
| result.name            | 名称                                                         |
| result.path_display    | パス (表示目的で大文字小文字を区別する).                     |
| result.client_modified | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ |
| result.server_modified | Dropbox上で最後に更新された日時                              |
| result.size            | ファイルサイズ(バイト単位)                                   |
| result.content_hash    | ファイルコンテンツのハッシュ                                 |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.


