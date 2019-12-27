# team diag explorer 

Report while team information

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it.
If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

This command use following access type(s) during the operation:
* Dropbox Business File access* Dropbox Business Information access

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team diag explorer 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team diag explorer 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

## Options

| オプション | 説明                                | デフォルト |
|------------|-------------------------------------|------------|
| `-all`     | Include additional reports          | false      |
| `-file`    | Dropbox Business file access        | default    |
| `-info`    | Dropbox Business information access | default    |
| `-mgmt`    | Dropbox Business management         | default    |

Common options:

| オプション      | 説明                                                                                             | デフォルト     |
|-----------------|--------------------------------------------------------------------------------------------------|----------------|
| `-bandwidth-kb` | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒)0の場合、制限を行わない | 0              |
| `-concurrency`  | 指定した並列度で並列処理を行います                                                               | プロセッサー数 |
| `-debug`        | デバッグモードを有効にする                                                                       | false          |
| `-low-memory`   | Low memory footprint mode                                                                        | false          |
| `-proxy`        | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                        |                |
| `-quiet`        | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                              | false          |
| `-secure`       | トークンをファイルに保存しません                                                                 | false          |
| `-workspace`    | ワークスペースへのパス                                                                           |                |

## Authentication

For the first run, `toolbox` will ask you an authentication with your Dropbox account. 
Please copy the link and paste it into your browser. Then proceed to authorization.
After authorization, Dropbox will show you an authorization code.
Please copy that code and paste it to the `toolbox`.

```
watermint toolbox xx.x.xxx
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Testing network connection...
Done

1. Visit the URL for the auth dialog:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

## Network configuration: Proxy

The executable automatically detects your proxy configuration from the environment.
However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port.
Currently, the executable doesn't support proxies which require authentication.

# Result

Report file path will be displayed last line of the command line output.
If you missed command line output, please see path below.
[job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: device 

Report files are generated in three formats, `device.csv`, `device.xlsx` and `device.json`.
But if you run with `-low-memory` option, the command will generate only `device.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`...   

| 列                            | 説明                                                                   |
|-------------------------------|------------------------------------------------------------------------|
| team_member_id                | チームにおけるメンバーのID                                             |
| email                         | ユーザーのメールアドレス                                               |
| status                        | チームにおけるメンバーのステータス(active/invited/suspended/removed)   |
| given_name                    | 名                                                                     |
| surname                       | 名字                                                                   |
| familiar_name                 | ロケール依存の名前                                                     |
| display_name                  | ユーザーのDropboxアカウントの表示名称                                  |
| abbreviated_name              | ユーザーの省略名称                                                     |
| external_id                   | このユーザーに関連づけられた外部ID                                     |
| account_id                    | ユーザーのアカウントID                                                 |
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

## Report: feature 

Report files are generated in three formats, `feature.csv`, `feature.xlsx` and `feature.json`.
But if you run with `-low-memory` option, the command will generate only `feature.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `feature_0000.xlsx`, `feature_0001.xlsx`, `feature_0002.xlsx`...   

| 列                          | 説明                                                   |
|-----------------------------|--------------------------------------------------------|
| upload_api_rate_limit       | 毎月利用可能なアップロードAPIコール回数                |
| upload_api_rate_limit_count | この月に利用されたアップロードAPIコール回数            |
| has_team_shared_dropbox     | このチームが共有されたチームルートを持っているかどうか |
| has_team_file_events        | このチームがファイルイベント機能を持っているかどうか   |
| has_team_selective_sync     | このチームがチーム選択型同期を有効化しているかどうか   |

## Report: file_request 

Report files are generated in three formats, `file_request.csv`, `file_request.xlsx` and `file_request.json`.
But if you run with `-low-memory` option, the command will generate only `file_request.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`...   

| 列                          | 説明                                                                      |
|-----------------------------|---------------------------------------------------------------------------|
| account_id                  | ファイルリクエスト所有者のアカウントID                                    |
| team_member_id              | ファイルリクエスト所有者のチームメンバーとしてのID                        |
| email                       | ファイルリクエスト所有者のメールアドレス                                  |
| status                      | ファイルリクエスト所有者ユーザーの状態 (active/invited/suspended/removed) |
| surname                     | ファイルリクエスト所有者の名字                                            |
| given_name                  | ファイルリクエスト所有者の名                                              |
| file_request_id             | ファイルリクエストID                                                      |
| url                         | ファイルリクエストのURL                                                   |
| title                       | ファイルリクエストのタイトル                                              |
| created                     | このファイルリクエストが作成された日時                                    |
| is_open                     | このファイルリクエストがオープンしているかどうか                          |
| file_count                  | このファイルリクエストが受け取ったファイル数                              |
| destination                 | アップロードされたファイルが置かれるDropboxフォルダのパス                 |
| deadline                    | ファイルリクエストの締め切り.                                             |
| deadline_allow_late_uploads | 設定されている場合、締め切りを超えてもアップロードが許可されている        |

## Report: group 

Report files are generated in three formats, `group.csv`, `group.xlsx` and `group.json`.
But if you run with `-low-memory` option, the command will generate only `group.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`...   

| 列                    | 説明                                                                                    |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_name            | グループ名称                                                                            |
| group_id              | グループID                                                                              |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| group_external_id     |  グループの外部IDこの任意のIDは管理者がグループに付加できます                           |
| member_count          | グループ内のメンバー数                                                                  |

## Report: group_member 

Report files are generated in three formats, `group_member.csv`, `group_member.xlsx` and `group_member.json`.
But if you run with `-low-memory` option, the command will generate only `group_member.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`...   

| 列                    | 説明                                                                                    |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_id              | グループID                                                                              |
| group_name            | グループ名称                                                                            |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| access_type           | グループにおけるユーザーの役割 (member/owner)                                           |
| account_id            | ユーザーアカウントのID                                                                  |
| team_member_id        | チームにおけるメンバーのID                                                              |
| email                 | ユーザーのメールアドレス                                                                |
| status                | チームにおけるメンバーのステータス(active/invited/suspended/removed)                    |
| surname               | 名字                                                                                    |
| given_name            | 名                                                                                      |

## Report: info 

Report files are generated in three formats, `info.csv`, `info.xlsx` and `info.json`.
But if you run with `-low-memory` option, the command will generate only `info.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `info_0000.xlsx`, `info_0001.xlsx`, `info_0002.xlsx`...   

| 列                          | 説明                                                                                                          |
|-----------------------------|---------------------------------------------------------------------------------------------------------------|
| name                        | チームの名称                                                                                                  |
| team_id                     | The ID of the team.                                                                                           |
| num_licensed_users          | このチームで利用可能なライセンス数                                                                            |
| num_provisioned_users       | The number of accounts that have been invited or are already active members of the team.                      |
| policy_shared_folder_member | Which shared folders team members can join (from_team_only, or from_anyone)                                   |
| policy_shared_folder_join   | Who can join folders shared by team members (team, or anyone)                                                 |
| policy_shared_link_create   | Who can view shared links owned by team members (default_public, default_team_only, or team_only)             |
| policy_emm_state            | This describes the Enterprise Mobility Management (EMM) state for this team (disabled, optional, or required) |
| policy_office_add_in        | The admin policy around the Dropbox Office Add-In for this team (disabled, or enabled)                        |

## Report: linked_app 

Report files are generated in three formats, `linked_app.csv`, `linked_app.xlsx` and `linked_app.json`.
But if you run with `-low-memory` option, the command will generate only `linked_app.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`...   

| 列               | 説明                                                                 |
|------------------|----------------------------------------------------------------------|
| team_member_id   | チームにおけるメンバーのID                                           |
| email            | ユーザーのメールアドレス                                             |
| status           | チームにおけるメンバーのステータス(active/invited/suspended/removed) |
| given_name       | 名                                                                   |
| surname          | 名字                                                                 |
| familiar_name    | ロケール依存の名前                                                   |
| display_name     | ユーザーのDropboxアカウントの表示名称                                |
| abbreviated_name | ユーザーの省略名称                                                   |
| external_id      | このユーザーに関連づけられた外部ID                                   |
| account_id       | ユーザーのアカウントID                                               |
| app_id           | アプリケーションの固有ID                                             |
| app_name         | アプリケーション名称                                                 |
| is_app_folder    | アプリケーションが専用フォルダにリンクするかどうか                   |
| publisher        | パブリッシャーのURL                                                  |
| publisher_url    | アプリケーションパブリッシャーの名前                                 |
| linked           | アプリケーションがリンクされた日時                                   |

## Report: member 

Report files are generated in three formats, `member.csv`, `member.xlsx` and `member.json`.
But if you run with `-low-memory` option, the command will generate only `member.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`...   

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

## Report: member_quota 

Report files are generated in three formats, `member_quota.csv`, `member_quota.xlsx` and `member_quota.json`.
But if you run with `-low-memory` option, the command will generate only `member_quota.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`...   

| 列    | 説明                                                                |
|-------|---------------------------------------------------------------------|
| email | ユーザーのメールアドレス                                            |
| quota | カスタムの容量制限GB (1 TB = 1024 GB). 0の場合、容量制限をしません. |

## Report: namespace 

Report files are generated in three formats, `namespace.csv`, `namespace.xlsx` and `namespace.json`.
But if you run with `-low-memory` option, the command will generate only `namespace.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`...   

| 列             | 説明                                                                                   |
|----------------|----------------------------------------------------------------------------------------|
| name           | 名前空間の名称                                                                         |
| namespace_id   | 名前空間ID                                                                             |
| namespace_type | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| team_member_id | メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID           |

## Report: namespace_file 

Report files are generated in three formats, `namespace_file.csv`, `namespace_file.xlsx` and `namespace_file.json`.
But if you run with `-low-memory` option, the command will generate only `namespace_file.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`...   

| 列                      | 説明                                                                                            |
|-------------------------|-------------------------------------------------------------------------------------------------|
| namespace_type          | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)          |
| namespace_id            | 名前空間ID                                                                                      |
| namespace_name          | 名前空間の名称                                                                                  |
| namespace_member_email  | これがチームメンバーフォルダまたはアプリフォルダの場合、所有するチームメンバーのメールアドレス. |
| file_id                 | ファイルへの一意なID                                                                            |
| tag                     | エントリーの種別`file`, `folder`, または `deleted`                                              |
| name                    | 名称                                                                                            |
| path_display            | パス (表示目的で大文字小文字を区別する).                                                        |
| client_modified         | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ                                    |
| server_modified         | Dropbox上で最後に更新された日時                                                                 |
| revision                | ファイルの現在バージョンの一意な識別子                                                          |
| size                    | ファイルサイズ(バイト単位)                                                                      |
| content_hash            | ファイルコンテンツのハッシュ                                                                    |
| shared_folder_id        | これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダのID。      |
| parent_shared_folder_id | 設定されている場合、共有フォルダに内包されています.                                             |

## Report: namespace_size 

Report files are generated in three formats, `namespace_size.csv`, `namespace_size.xlsx` and `namespace_size.json`.
But if you run with `-low-memory` option, the command will generate only `namespace_size.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`...   

| 列                          | 説明                                                                                   |
|-----------------------------|----------------------------------------------------------------------------------------|
| status                      | 処理の状態                                                                             |
| reason                      | 失敗またはスキップの理由                                                               |
| input.name                  | 名前空間の名称                                                                         |
| input.namespace_id          | 名前空間ID                                                                             |
| input.namespace_type        | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| input.team_member_id        | メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID           |
| result.namespace_name       | 名前空間の名称                                                                         |
| result.namespace_id         | 名前空間ID                                                                             |
| result.namespace_type       | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| result.owner_team_member_id | メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID           |
| result.path                 | フォルダへのパス                                                                       |
| result.count_file           | このフォルダに含まれるファイル数                                                       |
| result.count_folder         | このフォルダに含まれるフォルダ数                                                       |
| result.count_descendant     | このフォルダに含まれるファイル・フォルダ数                                             |
| result.size                 | フォルダのサイズ                                                                       |
| result.api_complexity       | APIを用いて操作する場合のフォルダ複雑度の指標                                          |

## Report: shared_link 

Report files are generated in three formats, `shared_link.csv`, `shared_link.xlsx` and `shared_link.json`.
But if you run with `-low-memory` option, the command will generate only `shared_link.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`...   

| 列             | 説明                                                                 |
|----------------|----------------------------------------------------------------------|
| shared_link_id | ファイルまたはフォルダへのリンクのID                                 |
| tag            | エントリーの種別 (file, または folder)                               |
| url            | 共有リンクのURL.                                                     |
| name           | リンク先ファイル名称                                                 |
| expires        | 有効期限 (設定されている場合)                                        |
| path_lower     | パス (すべて小文字に変換).                                           |
| visibility     | 共有リンクの開示範囲                                                 |
| account_id     | ユーザーのアカウントID                                               |
| team_member_id | チームにおけるメンバーのID                                           |
| email          | ユーザーのメールアドレス                                             |
| status         | チームにおけるメンバーのステータス(active/invited/suspended/removed) |
| surname        | リンク所有者の名字                                                   |
| given_name     | リンク所有者の名                                                     |

## Report: usage 

Report files are generated in three formats, `usage.csv`, `usage.xlsx` and `usage.json`.
But if you run with `-low-memory` option, the command will generate only `usage.json}}` report.
In case of a report become large, a report in `.xlsx` format will be split into several chunks
like `usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`...   

| 列         | 説明                                                     |
|------------|----------------------------------------------------------|
| email      | アカウントのメールアドレス                               |
| used_gb    | The user's total space usage (in GB, 1GB = 1024 MB).     |
| used_bytes | The user's total space usage (bytes).                    |
| allocation | The user's space allocation (individual, or team)        |
| allocated  | The total space allocated to the user's account (bytes). |

