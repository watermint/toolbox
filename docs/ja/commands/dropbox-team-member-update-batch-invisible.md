---
layout: command
title: コマンド `dropbox team member update batch invisible`
lang: ja
---

# dropbox team member update batch invisible

メンバーへのディレクトリ制限を有効にします (非可逆な操作です)

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
| Dropbox for teams：チームメンバーの確認                                             |
| Dropbox for teams：チームメンバーの確認と管理                                       |
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
.\tbx.exe dropbox team member update batch invisible -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team member update batch invisible -file /PATH/TO/DATA_FILE.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション | 説明                   | デフォルト |
|------------|------------------------|------------|
| `-file`    | データファイルへのパス |            |
| `-peer`    | アカウントの別名       | default    |

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

視認性を変えるためのメンバーリスト

| 列    | 説明                     | 例               |
|-------|--------------------------|------------------|
| email | メンバーのメールアドレス | taro@example.com |

最初の行はヘッダ行です. プログラムは、ヘッダのないファイルを受け入れます.
```
email
taro@example.com
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

| 列                      | 説明                                                                                            |
|-------------------------|-------------------------------------------------------------------------------------------------|
| status                  | 処理の状態                                                                                      |
| reason                  | 失敗またはスキップの理由                                                                        |
| input.Email             | メンバーのメールアドレス                                                                        |
| result.team_member_id   | チームにおけるメンバーのID                                                                      |
| result.email            | ユーザーのメールアドレス                                                                        |
| result.email_verified   | trueの場合、ユーザーのメールアドレスはユーザーによって所有されていることが確認されています.     |
| result.status           | チームにおけるメンバーのステータス(active/invited/suspended/removed)                            |
| result.given_name       | 名                                                                                              |
| result.surname          | 名字                                                                                            |
| result.familiar_name    | ロケール依存の名前                                                                              |
| result.display_name     | ユーザーのDropboxアカウントの表示名称                                                           |
| result.abbreviated_name | ユーザーの省略名称                                                                              |
| result.member_folder_id | ユーザールートフォルダの名前空間ID.                                                             |
| result.external_id      | このユーザーに関連づけられた外部ID                                                              |
| result.account_id       | ユーザーのアカウントID                                                                          |
| result.persistent_id    | ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.                             |
| result.joined_on        | メンバーがチームに参加した日時.                                                                 |
| result.invited_on       | ユーザーがチームに招待された日付と時間                                                          |
| result.role             | ユーザーのチームでの役割 (team_admin, user_management_admin, support_admin, または member_only) |
| result.tag              | 処理のタグ                                                                                      |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...


