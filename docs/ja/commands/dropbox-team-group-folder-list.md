---
layout: command
title: コマンド `dropbox team group folder list`
lang: ja
---

# dropbox team group folder list

各グループのフォルダーを一覧表示 

グループの権限をフォルダにマッピングし、コンテンツアクセスパターンを明らかにします。アクセスレビューと権限の継承の理解に不可欠です。過剰な権限を持つグループを特定し、セキュリティのためにフォルダ構造を最適化するのに役立ちます。

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

| 説明                                                                                |
|-------------------------------------------------------------------------------------|
| Dropbox for teams：Dropboxのファイルやフォルダに関する情報を表示                    |
| Dropbox for teams：チームのグループ メンバーシップを表示                            |
| Dropbox for teams：チームメンバーの確認                                             |
| Dropbox for teams：Dropboxの共有設定と共同作業者を表示                              |
| Dropbox for teams：チームやメンバーのフォルダ構造を表示                             |
| Dropbox for teams：名前、ユーザー数、チーム設定など、チームの基本情報を表示します。 |

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
.\tbx.exe dropbox team group folder list 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team group folder list 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-base-path**
: ファイルパス標準を選択します。これは、特にDropbox for Teams向けのオプションです。Dropboxの個人版を使用している場合は、選択したものが何であってもほとんど問題ありません。Dropbox for Teamsでは、更新されたチームスペースで`home`を選択すると、ユーザー名付きの個人フォルダが選択されます。これは、個人フォルダ内のファイルを参照したりアップロードしたりする場合に便利です。なぜなら、パスにユーザー名付きのフォルダ名を指定する必要がないからです。一方、`root`を指定すると、アクセス権のあるすべてのフォルダにアクセスできます。ただし、個人フォルダにアクセスする場合には、個人フォルダの名前を含むパスを指定する必要があります。. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: フォルダ名によるフィルター. 名前による完全一致でフィルター.

**-folder-name-prefix**
: フォルダ名によるフィルター. 名前の前方一致によるフィルター.

**-folder-name-suffix**
: フォルダ名によるフィルター. 名前の後方一致によるフィルター.

**-group-name**
: グループ名でフィルタリングします. 名前による完全一致でフィルター.

**-group-name-prefix**
: グループ名でフィルタリングします. 名前の前方一致によるフィルター.

**-group-name-suffix**
: グループ名でフィルタリングします. 名前の後方一致によるフィルター.

**-include-external-groups**
: レポートに外部のグループを含める.. Default: false

**-peer**
: アカウントの別名. Default: default

**-scan-timeout**
: スキャンのタイムアウト設定. スキャンタイムアウトした場合、チームフォルダのサブフォルダのパスは `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME` のようなダミーパスに置き換えられます.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

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

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: group_to_folder

グループからフォルダへのマッピング.
このコマンドはレポートを3種類の書式で出力します. `group_to_folder.csv`, `group_to_folder.json`, ならびに `group_to_folder.xlsx`.

| 列                 | 説明                                                                                                             |
|--------------------|------------------------------------------------------------------------------------------------------------------|
| group_name         | グループ名称                                                                                                     |
| group_type         | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed)                          |
| group_is_same_team | グループが同じチームにいる場合は 'true' となります. それ以外の場合はfalseです.                                   |
| access_type        | このフォルダに対するグループのアクセスレベル.                                                                    |
| namespace_name     | 名前空間の名称                                                                                                   |
| path               | パス                                                                                                             |
| folder_type        | フォルダの種別. (`team_folder`: チームフォルダまたはチームフォルダ以下のフォルダ, `shared_folder`: 共有フォルダ) |
| owner_team_name    | このフォルダを所有するチームの名前                                                                               |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `group_to_folder_0000.xlsx`, `group_to_folder_0001.xlsx`, `group_to_folder_0002.xlsx`, ...

## レポート: group_with_no_folders

このレポートはチーム内のグループを一覧します.
このコマンドはレポートを3種類の書式で出力します. `group_with_no_folders.csv`, `group_with_no_folders.json`, ならびに `group_with_no_folders.xlsx`.

| 列                    | 説明                                                                                    |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_name            | グループ名称                                                                            |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| member_count          | グループ内のメンバー数                                                                  |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `group_with_no_folders_0000.xlsx`, `group_with_no_folders_0001.xlsx`, `group_with_no_folders_0002.xlsx`, ...


