---
layout: command
title: コマンド `dropbox team namespace file size`
lang: ja
---

# dropbox team namespace file size

チーム内全ての名前空間でのファイル・フォルダを一覧 

設定可能な深度スキャンでチームネームスペース全体のストレージ消費を分析します。ネームスペースタイプ（チーム、共有、メンバー、アプリフォルダ）別のサイズ分布を表示します。ストレージの最適化、大きなフォルダの特定、アーカイブ戦略の計画に重要です。

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
| Dropbox for teams：チームメンバーの確認                                             |
| Dropbox for teams：チームやメンバーのフォルダ構造を表示                             |
| Dropbox for teams：チーム内のファイルやフォルダのコンテンツを閲覧・編集できます。   |
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
.\tbx.exe dropbox team namespace file size 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team namespace file size 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-depth**
: フォルダ階層数の指定. Default: 3

**-folder-name**
: 名前に一致するフォルダのみをリストアップします. 名前による完全一致でフィルター.

**-folder-name-prefix**
: 名前に一致するフォルダのみをリストアップします. 名前の前方一致によるフィルター.

**-folder-name-suffix**
: 名前に一致するフォルダのみをリストアップします. 名前の後方一致によるフィルター.

**-include-app-folder**
: Trueの場合、アプリフォルダを含めます. Default: false

**-include-member-folder**
: Trueの場合、チームメンバーフォルダを含めます. Default: false

**-include-shared-folder**
: Trueの場合、共有フォルダを含めます. Default: true

**-include-team-folder**
: Trueの場合、チームフォルダを含めます. Default: true

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

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: namespace_size

名前空間のサイズ.
このコマンドはレポートを3種類の書式で出力します. `namespace_size.csv`, `namespace_size.json`, ならびに `namespace_size.xlsx`.

| 列                   | 説明                                                                                   |
|----------------------|----------------------------------------------------------------------------------------|
| namespace_name       | 名前空間の名称                                                                         |
| namespace_id         | 名前空間ID                                                                             |
| namespace_type       | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| owner_team_member_id | メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID           |
| path                 | フォルダへのパス                                                                       |
| count_file           | このフォルダに含まれるファイル数                                                       |
| count_folder         | このフォルダに含まれるフォルダ数                                                       |
| count_descendant     | このフォルダに含まれるファイル・フォルダ数                                             |
| size                 | フォルダのサイズ                                                                       |
| depth                | フォルダの深さ                                                                         |
| mod_time_earliest    | このフォルダまたは子フォルダ内のファイルの最も古い更新日時                             |
| mod_time_latest      | このフォルダまたは子フォルダ内のファイルの最も新しい更新日時                           |
| api_complexity       | APIを用いて操作する場合のフォルダ複雑度の指標                                          |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`, ...


