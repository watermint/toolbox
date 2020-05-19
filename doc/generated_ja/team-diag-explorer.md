# team diag explorer

チーム全体の情報をレポートします 

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

| Label               | Description                         |
|---------------------|-------------------------------------|
| business_file       | Dropbox Business File access        |
| business_info       | Dropbox Business Information access |
| business_management | Dropbox Business management         |

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
.\tbx.exe team diag explorer 
```

macOS, Linux:
```
$HOME/Desktop/tbx team diag explorer 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option  | Description                         | Default |
|---------|-------------------------------------|---------|
| `-all`  | 追加のレポートを含める              | false   |
| `-file` | Dropbox Business ファイルアクアセス | default |
| `-info` | Dropbox Business 情報アクセス       | default |
| `-mgmt` | Dropbox Business 管理               | default |

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

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: device

このレポートではチーム内の既存セッションとメンバー情報を一覧できます.
The command will generate a report in three different formats. `device.csv`, `device.json`, and `device.xlsx`.

| Column                        | Description                                                            |
|-------------------------------|------------------------------------------------------------------------|
| team_member_id                | チームにおけるメンバーのID                                             |
| email                         | ユーザーのメールアドレス                                               |
| status                        | チームにおけるメンバーのステータス(active/invited/suspended/removed)   |
| given_name                    | 名                                                                     |
| surname                       | 名字                                                                   |
| display_name                  | ユーザーのDropboxアカウントの表示名称                                  |
| device_tag                    | セッションのタイプ (web_session, desktop_client, または mobile_client) |
| id                            | セッションID                                                           |
| user_agent                    | ホストデバイスの情報                                                   |
| os                            | ホストOSの情報                                                         |
| browser                       | Webセッションのブラウザ情報                                            |
| ip_address                    | このセッションの昨秋アクティビティのIPアドレス                         |
| country                       | このセッションの最終アクティビティの国                                 |
| created                       | セッションが作成された日時                                             |
| updated                       | このセッションの最終アクティビティの日時                               |
| expires                       | このセッションが失効する日時                                           |
| host_name                     | デスクトップホストの名称                                               |
| client_type                   | Dropboxデスクトップクライアントタイプ (Windows, macまたはlinux)        |
| client_version                | Dropboxクライアントバージョン                                          |
| platform                      | ホストプラットホームの情報                                             |
| is_delete_on_unlink_supported | アカウントのファイルをリンク解除時に削除を試みます                     |
| device_name                   | デバイス名                                                             |
| os_version                    | ホストOSのバージョン                                                   |
| last_carrier                  | このデバイスで利用された最後のキャリア                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`, ...
## Report: namespace_member

このレポートは名前空間とそのメンバー一覧を出力します.
The command will generate a report in three different formats. `namespace_member.csv`, `namespace_member.json`, and `namespace_member.xlsx`.

| Column             | Description                                                                                          |
|--------------------|------------------------------------------------------------------------------------------------------|
| namespace_name     | 名前空間の名称                                                                                       |
| namespace_type     | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)               |
| entry_access_type  | ユーザーの共有ファイル・フォルダへのアクセスレベル (owner, editor, viewer, または viewer_no_comment) |
| entry_is_inherited | メンバーのアクセス権限が上位フォルダから継承されている場合true                                       |
| email              | ユーザーのメールアドレス                                                                             |
| display_name       | セッションのタイプ (web_session, desktop_client, または mobile_client)                               |
| group_name         | グループ名称                                                                                         |
| invitee_email      | このフォルダに招待されたメールアドレス                                                               |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_member_0000.xlsx`, `namespace_member_0001.xlsx`, `namespace_member_0002.xlsx`, ...
## Report: usage

このレポートはユーザーの現在のストレージ利用容量を出力します.
The command will generate a report in three different formats. `usage.csv`, `usage.json`, and `usage.xlsx`.

| Column     | Description                                             |
|------------|---------------------------------------------------------|
| email      | アカウントのメールアドレス                              |
| used_gb    | このユーザーの合計利用スペース (in GB, 1GB = 1024 MB).  |
| used_bytes | ユーザーの合計利用要領 (bytes).                         |
| allocation | このユーザーの利用容量の付与先 (individual, or team)    |
| allocated  | このユーザーアカウントに確保されている合計容量 (bytes). |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`, ...
## Report: namespace

このレポートはチームの名前空間を一覧します.
The command will generate a report in three different formats. `namespace.csv`, `namespace.json`, and `namespace.xlsx`.

| Column         | Description                                                                            |
|----------------|----------------------------------------------------------------------------------------|
| name           | 名前空間の名称                                                                         |
| namespace_type | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| team_member_id | メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID           |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`, ...
## Report: namespace_file

このレポートはチームの名前空間を一覧します.
The command will generate a report in three different formats. `namespace_file.csv`, `namespace_file.json`, and `namespace_file.xlsx`.

| Column                 | Description                                                                                     |
|------------------------|-------------------------------------------------------------------------------------------------|
| namespace_type         | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)          |
| namespace_name         | 名前空間の名称                                                                                  |
| namespace_member_email | これがチームメンバーフォルダまたはアプリフォルダの場合、所有するチームメンバーのメールアドレス. |
| tag                    | エントリーの種別`file`, `folder`, または `deleted`                                              |
| name                   | 名称                                                                                            |
| path_display           | パス (表示目的で大文字小文字を区別する).                                                        |
| client_modified        | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ                                    |
| server_modified        | Dropbox上で最後に更新された日時                                                                 |
| size                   | ファイルサイズ(バイト単位)                                                                      |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`, ...
## Report: shared_link

このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します.
The command will generate a report in three different formats. `shared_link.csv`, `shared_link.json`, and `shared_link.xlsx`.

| Column     | Description                                                          |
|------------|----------------------------------------------------------------------|
| tag        | エントリーの種別 (file, または folder)                               |
| url        | 共有リンクのURL.                                                     |
| name       | リンク先ファイル名称                                                 |
| expires    | 有効期限 (設定されている場合)                                        |
| path_lower | パス (すべて小文字に変換).                                           |
| visibility | 共有リンクの開示範囲                                                 |
| email      | ユーザーのメールアドレス                                             |
| status     | チームにおけるメンバーのステータス(active/invited/suspended/removed) |
| surname    | リンク所有者の名字                                                   |
| given_name | リンク所有者の名                                                     |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...
## Report: member

このレポートはメンバー一覧を出力します.
The command will generate a report in three different formats. `member.csv`, `member.json`, and `member.xlsx`.

| Column         | Description                                                                                     |
|----------------|-------------------------------------------------------------------------------------------------|
| email          | ユーザーのメールアドレス                                                                        |
| email_verified | trueの場合、ユーザーのメールアドレスはユーザーによって所有されていることが確認されています.     |
| status         | チームにおけるメンバーのステータス(active/invited/suspended/removed)                            |
| given_name     | 名                                                                                              |
| surname        | 名字                                                                                            |
| display_name   | ユーザーのDropboxアカウントの表示名称                                                           |
| joined_on      | メンバーがチームに参加した日時.                                                                 |
| role           | ユーザーのチームでの役割 (team_admin, user_management_admin, support_admin, または member_only) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...
## Report: group

このレポートはチーム内のグループを一覧します.
The command will generate a report in three different formats. `group.csv`, `group.json`, and `group.xlsx`.

| Column                | Description                                                                             |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_name            | グループ名称                                                                            |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| member_count          | グループ内のメンバー数                                                                  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`, ...
## Report: member_quota

このレポートはチームメンバーのカスタム容量上限の設定を出力します.
The command will generate a report in three different formats. `member_quota.csv`, `member_quota.json`, and `member_quota.xlsx`.

| Column | Description                                                         |
|--------|---------------------------------------------------------------------|
| email  | ユーザーのメールアドレス                                            |
| quota  | カスタムの容量制限GB (1 TB = 1024 GB). 0の場合、容量制限をしません. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`, ...
## Report: team_folder

このレポートはチーム内のチームフォルダを一覧します.
The command will generate a report in three different formats. `team_folder.csv`, `team_folder.json`, and `team_folder.xlsx`.

| Column                 | Description                                                                          |
|------------------------|--------------------------------------------------------------------------------------|
| name                   | チームフォルダの名称                                                                 |
| status                 | チームフォルダの状態 (active, archived, または archive_in_progress)                  |
| is_team_shared_dropbox |                                                                                      |
| sync_setting           | チームフォルダに設定された同期設定 (default, not_synced, または not_synced_inactive) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`, ...
## Report: feature

このレポートはチームの機能と設定を一覧します.
The command will generate a report in three different formats. `feature.csv`, `feature.json`, and `feature.xlsx`.

| Column                      | Description                                            |
|-----------------------------|--------------------------------------------------------|
| upload_api_rate_limit       | 毎月利用可能なアップロードAPIコール回数                |
| upload_api_rate_limit_count | この月に利用されたアップロードAPIコール回数            |
| has_team_shared_dropbox     | このチームが共有されたチームルートを持っているかどうか |
| has_team_file_events        | このチームがファイルイベント機能を持っているかどうか   |
| has_team_selective_sync     | このチームがチーム選択型同期を有効化しているかどうか   |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `feature_0000.xlsx`, `feature_0001.xlsx`, `feature_0002.xlsx`, ...
## Report: namespace_size

このレポートは処理結果を出力します.
The command will generate a report in three different formats. `namespace_size.csv`, `namespace_size.json`, and `namespace_size.xlsx`.

| Column                  | Description                                                                            |
|-------------------------|----------------------------------------------------------------------------------------|
| status                  | 処理の状態                                                                             |
| reason                  | 失敗またはスキップの理由                                                               |
| input.name              | 名前空間の名称                                                                         |
| input.namespace_type    | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| result.path             | フォルダへのパス                                                                       |
| result.count_file       | このフォルダに含まれるファイル数                                                       |
| result.count_folder     | このフォルダに含まれるフォルダ数                                                       |
| result.count_descendant | このフォルダに含まれるファイル・フォルダ数                                             |
| result.size             | フォルダのサイズ                                                                       |
| result.api_complexity   | APIを用いて操作する場合のフォルダ複雑度の指標                                          |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`, ...
## Report: group_member

このレポートはグループとメンバーを一覧します.
The command will generate a report in three different formats. `group_member.csv`, `group_member.json`, and `group_member.xlsx`.

| Column                | Description                                                                             |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_name            | グループ名称                                                                            |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| access_type           | グループにおけるユーザーの役割 (member/owner)                                           |
| email                 | ユーザーのメールアドレス                                                                |
| status                | チームにおけるメンバーのステータス(active/invited/suspended/removed)                    |
| surname               | 名字                                                                                    |
| given_name            | 名                                                                                      |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`, ...
## Report: file_request

このレポートはチームメンバーのもつファイルリクエストを一覧します.
The command will generate a report in three different formats. `file_request.csv`, `file_request.json`, and `file_request.xlsx`.

| Column                      | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| email                       | ファイルリクエスト所有者のメールアドレス                                  |
| status                      | ファイルリクエスト所有者ユーザーの状態 (active/invited/suspended/removed) |
| surname                     | ファイルリクエスト所有者の名字                                            |
| given_name                  | ファイルリクエスト所有者の名                                              |
| url                         | ファイルリクエストのURL                                                   |
| title                       | ファイルリクエストのタイトル                                              |
| created                     | このファイルリクエストが作成された日時                                    |
| is_open                     | このファイルリクエストがオープンしているかどうか                          |
| file_count                  | このファイルリクエストが受け取ったファイル数                              |
| destination                 | アップロードされたファイルが置かれるDropboxフォルダのパス                 |
| deadline                    | ファイルリクエストの締め切り.                                             |
| deadline_allow_late_uploads | 設定されている場合、締め切りを超えてもアップロードが許可されている        |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`, ...
## Report: linked_app

このレポートは接続済みアプリケーションと利用ユーザーを一覧します.
The command will generate a report in three different formats. `linked_app.csv`, `linked_app.json`, and `linked_app.xlsx`.

| Column        | Description                                                          |
|---------------|----------------------------------------------------------------------|
| email         | ユーザーのメールアドレス                                             |
| status        | チームにおけるメンバーのステータス(active/invited/suspended/removed) |
| given_name    | 名                                                                   |
| surname       | 名字                                                                 |
| display_name  | ユーザーのDropboxアカウントの表示名称                                |
| app_name      | アプリケーション名称                                                 |
| is_app_folder | アプリケーションが専用フォルダにリンクするかどうか                   |
| publisher     | パブリッシャーのURL                                                  |
| publisher_url | アプリケーションパブリッシャーの名前                                 |
| linked        | アプリケーションがリンクされた日時                                   |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`, ...
## Report: info

このレポートはチームの情報を一覧します.
The command will generate a report in three different formats. `info.csv`, `info.json`, and `info.xlsx`.

| Column                      | Description                                                                                              |
|-----------------------------|----------------------------------------------------------------------------------------------------------|
| name                        | チームの名称                                                                                             |
| team_id                     | チームのID                                                                                               |
| num_licensed_users          | このチームで利用可能なライセンス数                                                                       |
| num_provisioned_users       | 招待済みアカウント数 (アクティブメンバーまたは招待済み)                                                  |
| policy_shared_folder_member | チームメンバーが参加できる共有フォルダ (from_team_onlyまたはfrom_anyone)                                 |
| policy_shared_folder_join   | チームメンバーが共有できる範囲 (teamまたは anyone)                                                       |
| policy_shared_link_create   | チームメンバーが所有する共有リンクを誰が閲覧可能か (default_public, default_team_only, または team_only) |
| policy_emm_state            | Enterprise Mobility Management (EMM) のチームに対する状態 (disabled, optional, or required)              |
| policy_office_add_in        | Dropbox Office アドインについてこのチームに対する管理者ポリシー (disabled, or enabled)                   |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `info_0000.xlsx`, `info_0001.xlsx`, `info_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

