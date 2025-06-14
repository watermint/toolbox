---
layout: command
title: コマンド `dropbox team teamfolder member add`
lang: ja
---

# dropbox team teamfolder member add

チームフォルダへのユーザー/グループの一括追加 (非可逆な操作です)

このコマンドは、(1)チームフォルダが存在しない場合、新しいチームフォルダまたは新しいサブフォルダを作成します. このコマンドは、(2)フォルダのアクセス継承設定の変更、(3)フォルダが存在しない場合のグループ作成は行いません. このコマンドは冪等性を持つように設計されています. 操作上のエラーが発生した場合は、安全に再試行することができます. コマンドは、冪等性を保持するためのエラーを報告しません. たとえば、このコマンドは、「メンバーはすでにそのフォルダにアクセスしている」というようなエラーを報告しません.

例:

* Sales（チームフォルダ、グループ「Sales」の編集者アクセス)
	* Sydney (個人アカウント sydney@example.com の閲覧アクセス)
	* Tokyo (グループ "Tokyo Deal Desk" の編集者アクセス)
		* Monthly (個人アカウント success@example.com の閲覧アクセス )
* Marketing (チームフォルダ、グループ "Marketing "の編集者アクセス)
	* Sydney (グループ "Sydney Sales" の編集者アクセス)
	* Tokyo (グループ "Tokyo Sales "のビューアアクセス)

1. 次のようなCSVファイルを準備します

```
Sales,,editor,Sales
Sales,Sydney,editor,sydney@example.com
Sales,Tokyo,editor,Tokyo Deal Desk
Sales,Tokyo/Monthly,viewer,success@example.com
Marketing,,editor,Marketing
Marketing,Sydney,editor,Sydney Sales
Marketing,Tokyo,viewer,Tokyo Sales
```

2. その後、以下のようにコマンドを実行します.

```
tbx teamfolder member add -file /PATH/TO/DATA.csv
```

注: このコマンドは、チームフォルダが存在しない場合には、チームフォルダを作成します. しかし、このコマンドは、グループが見つからない場合には、グループを作成しません. このコマンドを実行する前に、グループが存在している必要があります.

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

| 説明                                                                                               |
|----------------------------------------------------------------------------------------------------|
| Dropbox for teams：Dropboxのファイルやフォルダの内容を表示                                         |
| Dropbox for teams：Dropboxのファイルやフォルダの内容を編集する                                     |
| Dropbox for teams：チームのグループ メンバーシップを表示                                           |
| Dropbox for teams：メンバー アカウントの削除や復元など、チーム グループ メンバーシップの表示と管理 |
| Dropbox for teams：Dropboxの共有設定と共同作業者を表示                                             |
| Dropbox for teams：Dropboxの共有設定と共同作業者の表示と管理                                       |
| Dropbox for teams：チームやメンバーのフォルダ構造を表示                                            |
| Dropbox for teams：チーム内のファイルやフォルダのコンテンツを閲覧・編集できます。                  |
| Dropbox for teams：名前、ユーザー数、チーム設定など、チームの基本情報を表示します。                |

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
.\tbx.exe dropbox team teamfolder member add -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team teamfolder member add -file /PATH/TO/DATA_FILE.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-admin-group-name**
: 管理者操作のための仮グループ名. Default: watermint-toolbox-admin

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

アクセスを追加するためのチームフォルダとメンバーリスト. 各行には、1つのメンバーと1つのフォルダを対応させます. フォルダに2人以上のメンバーを追加したい場合は、それぞれのメンバー用の行を作成してください. 同様に、2つ以上のフォルダにメンバーを追加したい場合は、それぞれのフォルダの行を作成してください.

| 列                         | 説明                                                                                                            | 例       |
|----------------------------|-----------------------------------------------------------------------------------------------------------------|----------|
| team_folder_name           | チームフォルダ名                                                                                                | 営業     |
| path                       | チームフォルダのルートからの相対パス. チームフォルダのルートにメンバーを追加したい場合は空のままにしておきます. | レポート |
| access_type                | アクセス権限 (viewer/editor)                                                                                    | editor   |
| group_name_or_member_email | グループ名またはメンバーのメールアドレス                                                                        | 営業     |

最初の行はヘッダ行です. プログラムは、ヘッダのないファイルを受け入れます.
```
team_folder_name,path,access_type,group_name_or_member_email
営業,レポート,editor,営業
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

| 列                               | 説明                                                                                                            |
|----------------------------------|-----------------------------------------------------------------------------------------------------------------|
| status                           | 処理の状態                                                                                                      |
| reason                           | 失敗またはスキップの理由                                                                                        |
| input.team_folder_name           | チームフォルダ名                                                                                                |
| input.path                       | チームフォルダのルートからの相対パス. チームフォルダのルートにメンバーを追加したい場合は空のままにしておきます. |
| input.access_type                | アクセス権限 (viewer/editor)                                                                                    |
| input.group_name_or_member_email | グループ名またはメンバーのメールアドレス                                                                        |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...


