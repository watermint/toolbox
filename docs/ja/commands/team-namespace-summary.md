---
layout: command
title: コマンド
lang: ja
---

# team namespace summary

チーム・ネームスペースの状態概要を報告する. 

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
* Dropbox Business: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## 認可スコープ

| 説明                                                                           |
|--------------------------------------------------------------------------------|
| Dropbox Business: チームメンバーの確認                                         |
| Dropbox Business: Dropboxの共有設定と共同作業者の表示                          |
| Dropbox Business: チームやメンバーのフォルダの構造を閲覧                       |
| Dropbox Business: チーム内のファイルやフォルダーのコンテンツを閲覧・編集       |
| Dropbox Business: 名前、ユーザー数、チーム設定など、チームの基本的な情報を確認 |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します.
リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2022 Takayuki Okazaki
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
.\tbx.exe team namespace summary 
```

macOS, Linux:
```
$HOME/Desktop/tbx team namespace summary 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション             | 説明                                     | デフォルト |
|------------------------|------------------------------------------|------------|
| `-peer`                | アカウントの別名                         | default    |
| `-skip-member-summary` | メンバー名前空間のスキャンをスキップする | false      |

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

## レポート: folder_without_parent

親フォルダーがないフォルダー.
このコマンドはレポートを3種類の書式で出力します. `folder_without_parent.csv`, `folder_without_parent.json`, ならびに `folder_without_parent.xlsx`.

| 列                      | 説明                                                                                                 |
|-------------------------|------------------------------------------------------------------------------------------------------|
| shared_folder_id        | 共有フォルダのID                                                                                     |
| parent_shared_folder_id | 親共有フォルダのID. このフィールドはフォルダが他の共有フォルダに含まれる場合のみ設定されます.        |
| name                    | 共有フォルダの名称                                                                                   |
| access_type             | ユーザーの共有ファイル・フォルダへのアクセスレベル (owner, editor, viewer, または viewer_no_comment) |
| path_lower              | 共有フォルダのフルパス(小文字に変換済み).                                                            |
| is_inside_team_folder   | フォルダがチームフォルダに内包されているかどうか                                                     |
| is_team_folder          | このフォルダがチームフォルダであるかどうか                                                           |
| policy_manage_access    | このフォルダへメンバーを追加したり削除できるユーザー                                                 |
| policy_shared_link      | このフォルダの共有リンクを誰が利用できるか                                                           |
| policy_member           | だれがこの共有フォルダのメンバーに参加できるか (team, または anyone)                                 |
| policy_viewer_info      | だれが閲覧社情報を有効化・無効化できるか                                                             |
| owner_team_id           | フォルダ所有チームのチームID                                                                         |
| owner_team_name         | このフォルダを所有するチームの名前                                                                   |
| access_inheritance      | アクセス継承タイプ                                                                                   |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `folder_without_parent_0000.xlsx`, `folder_without_parent_0001.xlsx`, `folder_without_parent_0002.xlsx`, ...

## レポート: member

メンバーネームスペースの概要
このコマンドはレポートを3種類の書式で出力します. `member.csv`, `member.json`, ならびに `member.xlsx`.

| 列                  | 説明                                                          |
|---------------------|---------------------------------------------------------------|
| email               | メンバーのメールアドレス                                      |
| total_namespaces    | 合計ネームスペース数（メンバールート・ネームスペースを除く）. |
| mounted_namespaces  | マウントされているフォルダーの数                              |
| owner_namespaces    | このメンバーが所有する共有フォルダーの数.                     |
| team_folders        | チームフォルダー数                                            |
| inside_team_folders | チームフォルダー内のフォルダー数                              |
| external_folders    | チーム外のユーザーから共有されたフォルダーの数                |
| app_folders         | アプリフォルダの数.                                           |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...

## レポート: team

チームネームスペースの概要.
このコマンドはレポートを3種類の書式で出力します. `team.csv`, `team.json`, ならびに `team.xlsx`.

| 列              | 説明                 |
|-----------------|----------------------|
| namespace_type  | ネームスペースの種類 |
| namespace_count | ネームスペースの数   |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `team_0000.xlsx`, `team_0001.xlsx`, `team_0002.xlsx`, ...

## レポート: team_folder

チームフォルダーの概要.
このコマンドはレポートを3種類の書式で出力します. `team_folder.csv`, `team_folder.json`, ならびに `team_folder.xlsx`.

| 列                    | 説明                                      |
|-----------------------|-------------------------------------------|
| name                  | チームフォルダ名                          |
| num_namespaces_inside | このチームフォルダ内のネームスペースの数. |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.


