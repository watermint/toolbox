# team device unlink

デバイスのセッションを解除します (非可逆な操作です)

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
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

## 認可スコープ

| 説明                         |
|------------------------------|
| Dropbox Business File access |

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
.\tbx.exe team device unlink -file /path/to/data/file.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx team device unlink -file /path/to/data/file.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション          | 説明                                       | デフォルト |
|---------------------|--------------------------------------------|------------|
| `-delete-on-unlink` | デバイスリンク解除時にファイルを削除します | false      |
| `-file`             | データファイル                             |            |
| `-peer`             | アカウントの別名                           | default    |

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

# ファイル書式

## 書式: File

このレポートではチーム内の既存セッションとメンバー情報を一覧できます.

| 列                            | 説明                                                                   | 例                                         |
|-------------------------------|------------------------------------------------------------------------|--------------------------------------------|
| team_member_id                | チームにおけるメンバーのID                                             | dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx  |
| email                         | ユーザーのメールアドレス                                               | john.smith@example.com                     |
| status                        | チームにおけるメンバーのステータス(active/invited/suspended/removed)   | active                                     |
| given_name                    | 名                                                                     | John                                       |
| surname                       | 名字                                                                   | Smith                                      |
| familiar_name                 | ロケール依存の名前                                                     | John Smith                                 |
| display_name                  | ユーザーのDropboxアカウントの表示名称                                  | John Smith                                 |
| abbreviated_name              | ユーザーの省略名称                                                     | JS                                         |
| external_id                   | このユーザーに関連づけられた外部ID                                     |                                            |
| account_id                    | ユーザーのアカウントID                                                 | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   |
| device_tag                    | セッションのタイプ (web_session, desktop_client, または mobile_client) | desktop_client                             |
| id                            | セッションID                                                           | dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx |
| user_agent                    | ホストデバイスの情報                                                   |                                            |
| os                            | ホストOSの情報                                                         |                                            |
| browser                       | Webセッションのブラウザ情報                                            |                                            |
| ip_address                    | このセッションの昨秋アクティビティのIPアドレス                         | xx.xxx.x.xxx                               |
| country                       | このセッションの最終アクティビティの国                                 | United States                              |
| created                       | セッションが作成された日時                                             | 2019-09-20T23:47:33Z                       |
| updated                       | このセッションの最終アクティビティの日時                               | 2019-10-25T04:42:16Z                       |
| expires                       | このセッションが失効する日時                                           |                                            |
| host_name                     | デスクトップホストの名称                                               | nihonbashi                                 |
| client_type                   | Dropboxデスクトップクライアントタイプ (Windows, macまたはlinux)        | windows                                    |
| client_version                | Dropboxクライアントバージョン                                          | 83.4.152                                   |
| platform                      | ホストプラットホームの情報                                             | Windows 10 1903                            |
| is_delete_on_unlink_supported | アカウントのファイルをリンク解除時に削除を試みます                     | TRUE                                       |
| device_name                   | デバイス名                                                             |                                            |
| os_version                    | ホストOSのバージョン                                                   |                                            |
| last_carrier                  | このデバイスで利用された最後のキャリア                                 |                                            |

最初の行はヘッダ行です. プログラムは、ヘッダのないファイルを受け入れます.
```
team_member_id,email,status,given_name,surname,familiar_name,display_name,abbreviated_name,external_id,account_id,device_tag,id,user_agent,os,browser,ip_address,country,created,updated,expires,host_name,client_type,client_version,platform,is_delete_on_unlink_supported,device_name,os_version,last_carrier
dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,john.smith@example.com,active,John,Smith,John Smith,John Smith,JS,,dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,desktop_client,dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,,,,xx.xxx.x.xxx,United States,2019-09-20T23:47:33Z,2019-10-25T04:42:16Z,,nihonbashi,windows,83.4.152,Windows 10 1903,TRUE,,,
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

| 列                                  | 説明                                                                   |
|-------------------------------------|------------------------------------------------------------------------|
| status                              | 処理の状態                                                             |
| reason                              | 失敗またはスキップの理由                                               |
| input.team_member_id                | チームにおけるメンバーのID                                             |
| input.email                         | ユーザーのメールアドレス                                               |
| input.status                        | チームにおけるメンバーのステータス(active/invited/suspended/removed)   |
| input.given_name                    | 名                                                                     |
| input.surname                       | 名字                                                                   |
| input.display_name                  | ユーザーのDropboxアカウントの表示名称                                  |
| input.device_tag                    | セッションのタイプ (web_session, desktop_client, または mobile_client) |
| input.id                            | セッションID                                                           |
| input.user_agent                    | ホストデバイスの情報                                                   |
| input.os                            | ホストOSの情報                                                         |
| input.browser                       | Webセッションのブラウザ情報                                            |
| input.ip_address                    | このセッションの昨秋アクティビティのIPアドレス                         |
| input.country                       | このセッションの最終アクティビティの国                                 |
| input.created                       | セッションが作成された日時                                             |
| input.updated                       | このセッションの最終アクティビティの日時                               |
| input.expires                       | このセッションが失効する日時                                           |
| input.host_name                     | デスクトップホストの名称                                               |
| input.client_type                   | Dropboxデスクトップクライアントタイプ (Windows, macまたはlinux)        |
| input.client_version                | Dropboxクライアントバージョン                                          |
| input.platform                      | ホストプラットホームの情報                                             |
| input.is_delete_on_unlink_supported | アカウントのファイルをリンク解除時に削除を試みます                     |
| input.device_name                   | デバイス名                                                             |
| input.os_version                    | ホストOSのバージョン                                                   |
| input.last_carrier                  | このデバイスで利用された最後のキャリア                                 |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

