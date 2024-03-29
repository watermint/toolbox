---
layout: command
title: コマンド `google sheets sheet import`
lang: ja
---

# google sheets sheet import

スプレッドシートにデータをインポート 

# セキュリティ

`watermint toolbox`は認証情報をファイルシステム上に保存します. それは次のパスです:

| OS      | パス                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

これらの認証情報ファイルはDropboxサポートを含め誰にも共有しないでください.
不必要になった場合にはこれらのファイルを削除しても問題ありません. 認証情報の削除を確実にしたい場合には、アプリケーションアクセス設定または管理コンソールからアプリケーションへの許可を取り消してください.

方法は次のヘルプセンター記事をご参照ください:
* Google: https://support.google.com/accounts/answer/3466521

## 認可スコープ

| 説明                                                                    |
|-------------------------------------------------------------------------|
| Google Sheets: Google Driveでのスプレッドシートの閲覧、編集、作成、削除 |

# 認可

最初の実行では、`tbx`はあなたのGoogleアカウントへの認可を要求します.
Enterキーを押すと、ブラウザが起動します。その後、サービスが認証を行い、tbxがその結果を受け取ります。認証成功のメッセージが表示されたら、ブラウザのウィンドウを閉じてもかまいません。
```

watermint toolbox xx.x.xxx
==========================

© 2016-2024 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

認可URLを開きます:
https://accounts.google.com/o/oauth2/auth?client_id=xxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&state=xxxxxxxx

```

# インストール

[最新リリース](https://github.com/watermint/toolbox/releases/latest)からコンパイル済みのバイナリをダウンロードしてください. Windowsをお使いの方は、`tbx-xx.x.xxx-win.zip`のようなzipファイルをダウンロードしてください. その後、アーカイブを解凍し、デスクトップ フォルダに `tbx.exe` を配置します.
watermint toolboxは、システムで許可されていれば、システム内のどのパスからでも実行できます. しかし、説明書のサンプルでは、デスクトップ フォルダを使用しています. デスクトップ フォルダ以外にバイナリを配置した場合は、パスを読み替えてください.

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe google sheets sheet import -data /LOCAL/PATH/TO/INPUT.csv -id VALUE -range VALUE
```

macOS, Linux:
```
$HOME/Desktop/tbx google sheets sheet import -data /LOCAL/PATH/TO/INPUT.csv -id VALUE -range VALUE
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション   | 説明                                                                                                                                                                                                                                                                                                                                            | デフォルト |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------|
| `-data`      | データファイルのパス                                                                                                                                                                                                                                                                                                                            |            |
| `-id`        | スプレッドシートID                                                                                                                                                                                                                                                                                                                              |            |
| `-input-raw` | Raw入力                                                                                                                                                                                                                                                                                                                                         | false      |
| `-peer`      | アカウントの別名                                                                                                                                                                                                                                                                                                                                | default    |
| `-range`     | 値がカバーする範囲をA1表記で表します. これは、Sheet1!A1:B2のような文字列で、スプレッドシート内のセルのグループを参照するもので、一般的には数式の中で使用されます. `Sheet1!A1:B2` は、Sheet1の上2行の最初の2つのセルを指します. `A1:B2` は、最初に見えるシートの上2行の最初の2つのセルを指します. `Sheet1` は Sheet1 のすべてのセルを参照します. |            |

## 共通のオプション:

| オプション         | 説明                                                                                               | デフォルト     |
|--------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auth-database`   | 認証データベースへのカスタムパス (デフォルト: $HOME/.toolbox/secrets/secrets.db)                   |                |
| `-auto-open`       | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`    | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`   | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage`  | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`     | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`           | デバッグモードを有効にする                                                                         | false          |
| `-experiment`      | 実験的機能を有効化する                                                                             |                |
| `-extra`           | 追加パラメータファイルのパス                                                                       |                |
| `-lang`            | 表示言語                                                                                           | auto           |
| `-output`          | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`           | HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください    |                |
| `-quiet`           | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-retain-job-data` | ジョブデータ保持ポリシー                                                                           | default        |
| `-secure`          | トークンをファイルに保存しません                                                                   | false          |
| `-skip-logging`    | ローカルストレージへのログ保存をスキップ                                                           | false          |
| `-verbose`         | 現在の操作を詳細に表示します.                                                                      | false          |
| `-workspace`       | ワークスペースへのパス                                                                             |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: updated

更新された値
このコマンドはレポートを3種類の書式で出力します. `updated.csv`, `updated.json`, ならびに `updated.xlsx`.

| 列              | 説明               |
|-----------------|--------------------|
| spreadsheet_id  | スプレッドシートID |
| updated_range   | 更新された範囲     |
| updated_rows    | 更新された行数     |
| updated_columns | 更新されたカラム数 |
| updated_cells   | 更新されたセル数   |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `updated_0000.xlsx`, `updated_0001.xlsx`, `updated_0002.xlsx`, ...

# コマンドのグリッドデータ入力

## グリッドデータの入力: Data

入力データファイル

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.


