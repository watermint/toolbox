---
layout: command
title: コマンド `dropbox team runas sharedfolder batch share`
lang: ja
---

# dropbox team runas sharedfolder batch share

メンバーのフォルダを一括で共有 

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
| Dropbox for teams：Dropboxのファイルやフォルダの内容を表示                          |
| Dropbox for teams：チームメンバーの確認                                             |
| Dropbox for teams：Dropboxの共有設定と共同作業者を表示                              |
| Dropbox for teams：Dropboxの共有設定と共同作業者の表示と管理                        |
| Dropbox for teams：チームやメンバーのフォルダ構造を表示                             |
| Dropbox for teams：名前、ユーザー数、チーム設定など、チームの基本情報を表示します。 |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します.
リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2024 Takayuki Okazaki
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
.\tbx.exe dropbox team runas sharedfolder batch share -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team runas sharedfolder batch share -file /PATH/TO/DATA_FILE.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション            | 説明                                                                          | デフォルト |
|-----------------------|-------------------------------------------------------------------------------|------------|
| `-acl-update-policy`  | この共有フォルダのメンバーを追加・削除できる人.                               | owner      |
| `-file`               | データファイルへのパス                                                        |            |
| `-member-policy`      | この共有フォルダーのメンバーになれる人.                                       | anyone     |
| `-peer`               | アカウントの別名                                                              | default    |
| `-shared-link-policy` | この共有フォルダー内のコンテンツに作成された共有リンクに適用するポリシーです. | anyone     |

## 共通のオプション:

| オプション         | 説明                                                                                                                                                       | デフォルト     |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------|
| `-auth-database`   | 認証データベースへのカスタムパス (デフォルト: $HOME/.toolbox/secrets/secrets.db)                                                                           |                |
| `-auto-open`       | 成果物フォルダまたはURLを自動で開く                                                                                                                        | false          |
| `-bandwidth-kb`    | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない                                                         | 0              |
| `-budget-memory`   | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                                                                                | normal         |
| `-budget-storage`  | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                                                                                    | normal         |
| `-concurrency`     | 指定した並列度で並列処理を行います                                                                                                                         | プロセッサー数 |
| `-debug`           | デバッグモードを有効にする                                                                                                                                 | false          |
| `-experiment`      | 実験的機能を有効化する                                                                                                                                     |                |
| `-extra`           | 追加パラメータファイルのパス                                                                                                                               |                |
| `-lang`            | 表示言語                                                                                                                                                   | auto           |
| `-output`          | 出力書式 (none/text/markdown/json)                                                                                                                         | text           |
| `-output-filter`   | 出力フィルタ・クエリ（jq構文）。レポートの出力はjq構文を使ってフィルタリングされる。このオプションは、レポートがJSONとして出力される場合にのみ適用される。 |                |
| `-proxy`           | HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください                                                            |                |
| `-quiet`           | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                                                                        | false          |
| `-retain-job-data` | ジョブデータ保持ポリシー                                                                                                                                   | default        |
| `-secure`          | トークンをファイルに保存しません                                                                                                                           | false          |
| `-skip-logging`    | ローカルストレージへのログ保存をスキップ                                                                                                                   | false          |
| `-verbose`         | 現在の操作を詳細に表示します.                                                                                                                              | false          |
| `-workspace`       | ワークスペースへのパス                                                                                                                                     |                |

# ファイル書式

## 書式: File

メンバーフォルダデータ

| 列           | 説明                     | 例                   |
|--------------|--------------------------|----------------------|
| member_email | メンバーのメールアドレス | john@example.com     |
| path         | Path for share           | /projects/my_project |

最初の行はヘッダ行です. プログラムは、ヘッダのないファイルを受け入れます.
```
member_email,path
john@example.com,/projects/my_project
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

| 列                             | 説明                                                                                                 |
|--------------------------------|------------------------------------------------------------------------------------------------------|
| status                         | 処理の状態                                                                                           |
| reason                         | 失敗またはスキップの理由                                                                             |
| input.member_email             | メンバーのメールアドレス                                                                             |
| input.path                     | Path for share                                                                                       |
| result.shared_folder_id        | 共有フォルダのID                                                                                     |
| result.parent_shared_folder_id | 親共有フォルダのID. このフィールドはフォルダが他の共有フォルダに含まれる場合のみ設定されます.        |
| result.name                    | 共有フォルダの名称                                                                                   |
| result.access_type             | ユーザーの共有ファイル・フォルダへのアクセスレベル (owner, editor, viewer, または viewer_no_comment) |
| result.path_lower              | 共有フォルダのフルパス(小文字に変換済み).                                                            |
| result.is_inside_team_folder   | フォルダがチームフォルダに内包されているかどうか                                                     |
| result.is_team_folder          | このフォルダがチームフォルダであるかどうか                                                           |
| result.policy_manage_access    | このフォルダへメンバーを追加したり削除できるユーザー                                                 |
| result.policy_shared_link      | このフォルダの共有リンクを誰が利用できるか                                                           |
| result.policy_member_folder    | フォルダ自体に設定されている、この共有フォルダのメンバーになれる人.                                  |
| result.policy_member           | だれがこの共有フォルダのメンバーに参加できるか (team, または anyone)                                 |
| result.policy_viewer_info      | だれが閲覧社情報を有効化・無効化できるか                                                             |
| result.owner_team_id           | フォルダ所有チームのチームID                                                                         |
| result.owner_team_name         | このフォルダを所有するチームの名前                                                                   |
| result.access_inheritance      | アクセス継承タイプ                                                                                   |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.


