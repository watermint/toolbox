# team device unlink

デバイスのセッションを解除します (非可逆な操作です)

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS      | Path                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

## Auth scopes

| Label         | Description                  |
|---------------|------------------------------|
| business_file | Dropbox Business File access |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
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

# Usage

This document uses the Desktop folder for command example.
## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe team device unlink -file /path/to/data/file.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx team device unlink -file /path/to/data/file.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option              | Description                                | Default |
|---------------------|--------------------------------------------|---------|
| `-delete-on-unlink` | デバイスリンク解除時にファイルを削除します | false   |
| `-file`             | データファイル                             |         |
| `-peer`             | アカウントの別名                           | default |

## Common options:

| Option            | Description                                                                                        | Default        |
|-------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`      | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`  | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage` | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`    | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`          | デバッグモードを有効にする                                                                         | false          |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# File formats

## Format: File

このレポートではチーム内の既存セッションとメンバー情報を一覧できます.

| Column                        | Description                                                            | Example                                    |
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

The first line is a header line. The program will accept file without the header.
```
team_member_id,email,status,given_name,surname,familiar_name,display_name,abbreviated_name,external_id,account_id,device_tag,id,user_agent,os,browser,ip_address,country,created,updated,expires,host_name,client_type,client_version,platform,is_delete_on_unlink_supported,device_name,os_version,last_carrier
dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,john.smith@example.com,active,John,Smith,John Smith,John Smith,JS,,dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,desktop_client,dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,,,,xx.xxx.x.xxx,United States,2019-09-20T23:47:33Z,2019-10-25T04:42:16Z,,nihonbashi,windows,83.4.152,Windows 10 1903,TRUE,,,
```

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: operation_log

このレポートは処理結果を出力します.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                              | Description                                                            |
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

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

