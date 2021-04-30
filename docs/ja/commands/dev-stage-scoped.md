---
layout: command
title: コマンド
lang: ja
---

# dev stage scoped

Dropboxのスコープ付きOAuthアプリテスト 

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
* Dropbox Business: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## 認可スコープ

| 説明                                                                           |
|--------------------------------------------------------------------------------|
| Dropbox: Dropboxのファイルやフォルダのコンテンツを表示                         |
| Dropbox Business: チームメンバーの確認                                         |
| Dropbox Business: 名前、ユーザー数、チーム設定など、チームの基本的な情報を確認 |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します. リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2021 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

1. 次のURLを開き認証ダイアログを開いてください:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. 'Allow'をクリックします (先にログインしておく必要があります):
3. 認証コードをコピーします:
認証コードを入力してください
```
最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します. リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2021 Takayuki Okazaki
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
.\tbx.exe dev stage scoped 
```

macOS, Linux:
```
$HOME/Desktop/tbx dev stage scoped 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション    | 説明                         | デフォルト |
|---------------|------------------------------|------------|
| `-individual` | 個人向けのアカウントの別名   | default    |
| `-team`       | チーム向けのアカウントの別名 | default    |

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

## レポート: file_list

このレポートはファイルとフォルダのメタデータを出力します.
このコマンドはレポートを3種類の書式で出力します. `file_list.csv`, `file_list.json`, ならびに `file_list.xlsx`.

| 列                      | 説明                                                                                       |
|-------------------------|--------------------------------------------------------------------------------------------|
| id                      | ファイルへの一意なID                                                                       |
| tag                     | エントリーの種別`file`, `folder`, または `deleted`                                         |
| name                    | 名称                                                                                       |
| path_lower              | パス (すべて小文字に変換). これは常にスラッシュで始まります.                               |
| path_display            | パス (表示目的で大文字小文字を区別する).                                                   |
| client_modified         | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ                               |
| server_modified         | Dropbox上で最後に更新された日時                                                            |
| revision                | ファイルの現在バージョンの一意な識別子                                                     |
| size                    | ファイルサイズ(バイト単位)                                                                 |
| content_hash            | ファイルコンテンツのハッシュ                                                               |
| shared_folder_id        | これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダのID。 |
| parent_shared_folder_id | このファイルを含む共有フォルダのID.                                                        |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `file_list_0000.xlsx`, `file_list_0001.xlsx`, `file_list_0002.xlsx`, ...

## レポート: member_list

このレポートはメンバー一覧を出力します.
このコマンドはレポートを3種類の書式で出力します. `member_list.csv`, `member_list.json`, ならびに `member_list.xlsx`.

| 列               | 説明                                                                                            |
|------------------|-------------------------------------------------------------------------------------------------|
| team_member_id   | チームにおけるメンバーのID                                                                      |
| email            | ユーザーのメールアドレス                                                                        |
| email_verified   | trueの場合、ユーザーのメールアドレスはユーザーによって所有されていることが確認されています.     |
| status           | チームにおけるメンバーのステータス(active/invited/suspended/removed)                            |
| given_name       | 名                                                                                              |
| surname          | 名字                                                                                            |
| familiar_name    | ロケール依存の名前                                                                              |
| display_name     | ユーザーのDropboxアカウントの表示名称                                                           |
| abbreviated_name | ユーザーの省略名称                                                                              |
| member_folder_id | ユーザールートフォルダの名前空間ID.                                                             |
| external_id      | このユーザーに関連づけられた外部ID                                                              |
| account_id       | ユーザーのアカウントID                                                                          |
| persistent_id    | ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.                             |
| joined_on        | メンバーがチームに参加した日時.                                                                 |
| role             | ユーザーのチームでの役割 (team_admin, user_management_admin, support_admin, または member_only) |
| tag              | 処理のタグ                                                                                      |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `member_list_0000.xlsx`, `member_list_0001.xlsx`, `member_list_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

