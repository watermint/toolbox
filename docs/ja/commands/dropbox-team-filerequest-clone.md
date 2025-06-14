---
layout: command
title: コマンド `dropbox team filerequest clone`
lang: ja
---

# dropbox team filerequest clone

ファイルリクエストを入力データに従い複製します (試験的実装かつ非可逆な操作です)

既存のテンプレートに基づいて設定を変更した新しいファイルリクエストを作成します。月次レポートや定期的な提出などの標準化された収集プロセスを効率化します。受信者ごとのカスタマイズを可能にしながら、フォルダ構造を保持します。

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
* Dropbox for teams: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## 認可スコープ

| 説明 |
|------|

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します.
リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

1. 次のURLを開き認証ダイアログを開いてください:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. 'Allow'をクリックします (先にログインしておく必要があります):
3. 認証コードをコピーします:
認証コードを入力してください
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
.\tbx.exe dropbox team filerequest clone -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team filerequest clone -file /PATH/TO/DATA_FILE.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-base-path**
: ファイルパス標準を選択します。これは、特にDropbox for Teams向けのオプションです。Dropboxの個人版を使用している場合は、選択したものが何であってもほとんど問題ありません。Dropbox for Teamsでは、更新されたチームスペースで`home`を選択すると、ユーザー名付きの個人フォルダが選択されます。これは、個人フォルダ内のファイルを参照したりアップロードしたりする場合に便利です。なぜなら、パスにユーザー名付きのフォルダ名を指定する必要がないからです。一方、`root`を指定すると、アクセス権のあるすべてのフォルダにアクセスできます。ただし、個人フォルダにアクセスする場合には、個人フォルダの名前を含むパスを指定する必要があります。. Options: root (すべてのフォルダーへのフルアクセス権限), home (個人用ホームフォルダーへの限定アクセス). Default: root

**-file**
: データファイルへのパス

**-peer**
: アカウントの別名. Default: default

## 共通のオプション:

**-auth-database**
: 認証データベースへのカスタムパス (デフォルト: $HOME/.toolbox/secrets/secrets.db)

**-auto-open**
: 成果物フォルダまたはURLを自動で開く. Default: false

**-bandwidth-kb**
: コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない. Default: 0

**-budget-memory**
: メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます). Options: low, normal. Default: normal

**-budget-storage**
: ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します). Options: low, normal, unlimited. Default: normal

**-concurrency**
: 指定した並列度で並列処理を行います. Default: プロセッサー数

**-debug**
: デバッグモードを有効にする. Default: false

**-experiment**
: 実験的機能を有効化する

**-extra**
: 追加パラメータファイルのパス

**-lang**
: 表示言語. Options: auto, en, ja. Default: auto

**-output**
: 出力書式 (none/text/markdown/json). Options: text, markdown, json, none. Default: text

**-output-filter**
: 出力フィルタ・クエリ（jq構文）。レポートの出力はjq構文を使ってフィルタリングされる。このオプションは、レポートがJSONとして出力される場合にのみ適用される。

**-proxy**
: HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください

**-quiet**
: エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します. Default: false

**-retain-job-data**
: ジョブデータ保持ポリシー. Options: default, on_error, none. Default: default

**-secure**
: トークンをファイルに保存しません. Default: false

**-skip-logging**
: ローカルストレージへのログ保存をスキップ. Default: false

**-verbose**
: 現在の操作を詳細に表示します.. Default: false

**-workspace**
: ワークスペースへのパス

# ファイル書式

## 書式: File

このレポートはチームメンバーのもつファイルリクエストを一覧します.

| 列                          | 説明                                                                      | 例                                                 |
|-----------------------------|---------------------------------------------------------------------------|----------------------------------------------------|
| account_id                  | ファイルリクエスト所有者のアカウントID                                    | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx           |
| team_member_id              | ファイルリクエスト所有者のチームメンバーとしてのID                        | dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx          |
| email                       | ファイルリクエスト所有者のメールアドレス                                  | john@example.com                                   |
| status                      | ファイルリクエスト所有者ユーザーの状態 (active/invited/suspended/removed) | active                                             |
| surname                     | ファイルリクエスト所有者の名字                                            | Smith                                              |
| given_name                  | ファイルリクエスト所有者の名                                              | John                                               |
| file_request_id             | ファイルリクエストID                                                      | xxxxxxxxxxxxxxxxxx                                 |
| url                         | ファイルリクエストのURL                                                   | https://www.dropbox.com/request/xxxxxxxxxxxxxxxxxx |
| title                       | ファイルリクエストのタイトル                                              | 写真コンテスト                                     |
| created                     | このファイルリクエストが作成された日時                                    | 2019-09-20T23:47:33Z                               |
| is_open                     | このファイルリクエストがオープンしているかどうか                          | true                                               |
| file_count                  | このファイルリクエストが受け取ったファイル数                              | 3                                                  |
| destination                 | アップロードされたファイルが置かれるDropboxフォルダのパス                 | /Photo contest entries                             |
| deadline                    | ファイルリクエストの締め切り.                                             | 2019-10-20T23:47:33Z                               |
| deadline_allow_late_uploads | 設定されている場合、締め切りを超えてもアップロードが許可されている        | seven_days                                         |

最初の行はヘッダ行です. プログラムは、ヘッダのないファイルを受け入れます.
```
account_id,team_member_id,email,status,surname,given_name,file_request_id,url,title,created,is_open,file_count,destination,deadline,deadline_allow_late_uploads
dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,john@example.com,active,Smith,John,xxxxxxxxxxxxxxxxxx,https://www.dropbox.com/request/xxxxxxxxxxxxxxxxxx,写真コンテスト,2019-09-20T23:47:33Z,true,3,/Photo contest entries,2019-10-20T23:47:33Z,seven_days
```

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: operation_log

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `operation_log.csv`, `operation_log.json`, ならびに `operation_log.xlsx`.

| 列                                 | 説明                                                                      |
|------------------------------------|---------------------------------------------------------------------------|
| status                             | 処理の状態                                                                |
| reason                             | 失敗またはスキップの理由                                                  |
| input.email                        | ファイルリクエスト所有者のメールアドレス                                  |
| input.status                       | ファイルリクエスト所有者ユーザーの状態 (active/invited/suspended/removed) |
| input.surname                      | ファイルリクエスト所有者の名字                                            |
| input.given_name                   | ファイルリクエスト所有者の名                                              |
| input.url                          | ファイルリクエストのURL                                                   |
| input.title                        | ファイルリクエストのタイトル                                              |
| input.created                      | このファイルリクエストが作成された日時                                    |
| input.is_open                      | このファイルリクエストがオープンしているかどうか                          |
| input.file_count                   | このファイルリクエストが受け取ったファイル数                              |
| input.destination                  | アップロードされたファイルが置かれるDropboxフォルダのパス                 |
| input.deadline                     | ファイルリクエストの締め切り.                                             |
| input.deadline_allow_late_uploads  | 設定されている場合、締め切りを超えてもアップロードが許可されている        |
| result.email                       | ファイルリクエスト所有者のメールアドレス                                  |
| result.status                      | ファイルリクエスト所有者ユーザーの状態 (active/invited/suspended/removed) |
| result.surname                     | ファイルリクエスト所有者の名字                                            |
| result.given_name                  | ファイルリクエスト所有者の名                                              |
| result.url                         | ファイルリクエストのURL                                                   |
| result.title                       | ファイルリクエストのタイトル                                              |
| result.created                     | このファイルリクエストが作成された日時                                    |
| result.is_open                     | このファイルリクエストがオープンしているかどうか                          |
| result.file_count                  | このファイルリクエストが受け取ったファイル数                              |
| result.destination                 | アップロードされたファイルが置かれるDropboxフォルダのパス                 |
| result.deadline                    | ファイルリクエストの締め切り.                                             |
| result.deadline_allow_late_uploads | 設定されている場合、締め切りを超えてもアップロードが許可されている        |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...


