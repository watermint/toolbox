# services google sheets sheet list

スプレッドシートのシート一覧

# セキュリティ

`watermint toolbox`は認証情報をファイルシステム上に保存します. それは次のパスです:

| OS      | パス                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

これらの認証情報ファイルはDropboxサポートを含め誰にも共有しないでください. 不必要になった場合にはこれらのファイルを削除しても問題ありません.
認証情報の削除を確実にしたい場合には、アプリケーションアクセス設定または管理コンソールからアプリケーションへの許可を取り消してください.

方法は次のヘルプセンター記事をご参照ください:
* Google: https://support.google.com/accounts/answer/3466521

## 認可スコープ

| ラベル        | 説明          |
|---------------|---------------|
| google_sheets | Google Sheets |

# 認可

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it
into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code.
Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2020 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

1. 次のURLを開き認証ダイアログを開いてください:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. 'Allow'をクリックします (先にログインしておく必要があります):
3. 認証コードをコピーします:
認証コードを入力してください
```

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe services google sheets sheet list 
```

macOS, Linux:
```
$HOME/Desktop/tbx services google sheets sheet list 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"
のセキュリティーとプライバシーから一般タブを選択します. 次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション | 説明               | デフォルト                                                               |
|------------|--------------------|--------------------------------------------------------------------------|
| `-id`      | スプレッドシートID |                                                                          |
| `-peer`    | アカウントの別名   | &{default [https://www.googleapis.com/auth/spreadsheets.readonly] <nil>} |

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

## レポート: sheets

シートデータ このコマンドはレポートを3種類の書式で出力します. `sheets.csv`, `sheets.json`, ならびに `sheets.xlsx`.

| 列           | 説明                                      |
|--------------|-------------------------------------------|
| sheet_id     | シートID                                  |
| title        | シートの名前                              |
| index        | スプレッドシート内のシートのインデックス. |
| sheet_type   | シートの種類                              |
| row_count    | 行数                                      |
| column_count | カラム数                                  |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `sheets_0000.xlsx`, `sheets_0001.xlsx`, `sheets_0002.xlsx`,
...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`
のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

