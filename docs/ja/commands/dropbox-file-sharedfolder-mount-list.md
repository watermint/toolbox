---
layout: command
title: コマンド `dropbox file sharedfolder mount list`
lang: ja
---

# dropbox file sharedfolder mount list

現在のユーザーがマウントしているすべての共有フォルダーを一覧表示 

Dropboxにマウントされているすべての共有フォルダをリスト表示します。

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
* Dropbox (個人アカウント): https://help.dropbox.com/installs-integrations/third-party/third-party-apps

## 認可スコープ

| 説明                                                                                   |
|----------------------------------------------------------------------------------------|
| Dropbox: ユーザー名、メールアドレス、国名など、Dropboxアカウントの基本情報を表示します |
| Dropbox: Dropboxの共有設定と共同作業者の表示                                           |

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
.\tbx.exe dropbox file sharedfolder mount list 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox file sharedfolder mount list 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-base-path**
: ファイルパス標準を選択します。これは、特にDropbox for Teams向けのオプションです。Dropboxの個人版を使用している場合は、選択したものが何であってもほとんど問題ありません。Dropbox for Teamsでは、更新されたチームスペースで`home`を選択すると、ユーザー名付きの個人フォルダが選択されます。これは、個人フォルダ内のファイルを参照したりアップロードしたりする場合に便利です。なぜなら、パスにユーザー名付きのフォルダ名を指定する必要がないからです。一方、`root`を指定すると、アクセス権のあるすべてのフォルダにアクセスできます。ただし、個人フォルダにアクセスする場合には、個人フォルダの名前を含むパスを指定する必要があります。. Options: root (権限を持つすべてのフォルダへのフルアクセス), home (個人用ホームフォルダへの限定アクセス). Default: root

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

## レポート: mounts

このレポートは共有フォルダの一覧を出力します.
このコマンドはレポートを3種類の書式で出力します. `mounts.csv`, `mounts.json`, ならびに `mounts.xlsx`.

| 列                    | 説明                                                                                                 |
|-----------------------|------------------------------------------------------------------------------------------------------|
| shared_folder_id      | 共有フォルダのID                                                                                     |
| name                  | 共有フォルダの名称                                                                                   |
| access_type           | ユーザーの共有ファイル・フォルダへのアクセスレベル (owner, editor, viewer, または viewer_no_comment) |
| path_lower            | 共有フォルダのフルパス(小文字に変換済み).                                                            |
| is_inside_team_folder | フォルダがチームフォルダに内包されているかどうか                                                     |
| is_team_folder        | このフォルダがチームフォルダであるかどうか                                                           |
| policy_manage_access  | このフォルダへメンバーを追加したり削除できるユーザー                                                 |
| policy_shared_link    | このフォルダの共有リンクを誰が利用できるか                                                           |
| policy_member_folder  | フォルダ自体に設定されている、この共有フォルダのメンバーになれる人.                                  |
| policy_member         | だれがこの共有フォルダのメンバーに参加できるか (team, または anyone)                                 |
| policy_viewer_info    | だれが閲覧社情報を有効化・無効化できるか                                                             |
| owner_team_name       | このフォルダを所有するチームの名前                                                                   |
| access_inheritance    | アクセス継承タイプ                                                                                   |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `mounts_0000.xlsx`, `mounts_0001.xlsx`, `mounts_0002.xlsx`, ...


