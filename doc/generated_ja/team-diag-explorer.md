# team diag explorer

チーム全体の情報をレポートします 

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

| ラベル              | 説明                              |
|---------------------|-----------------------------------|
| business_file       | Dropbox Business ファイルアクセス |
| business_info       | Dropbox Business 情報アクセス     |
| business_management | Dropbox Business 管理             |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します. リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
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
.\tbx.exe team diag explorer 
```

macOS, Linux:
```
$HOME/Desktop/tbx team diag explorer 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション | 説明                                | デフォルト |
|------------|-------------------------------------|------------|
| `-all`     | 追加のレポートを含める              | false      |
| `-file`    | Dropbox Business ファイルアクアセス | default    |
| `-info`    | Dropbox Business 情報アクセス       | default    |
| `-mgmt`    | Dropbox Business 管理               | default    |

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
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: device

このレポートではチーム内の既存セッションとメンバー情報を一覧できます.
このコマンドはレポートを3種類の書式で出力します. `device.csv`, `device.json`, ならびに `device.xlsx`.

| 列                            | 説明                                                                   |
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

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`, ...
## レポート: errors

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `errors.csv`, `errors.json`, ならびに `errors.xlsx`.

| 列              | 説明                     |
|-----------------|--------------------------|
| status          | 処理の状態               |
| reason          | 失敗またはスキップの理由 |
| input.namespace | 名前空間                 |
| input.path      | パス                     |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...
## レポート: feature

このレポートはチームの機能と設定を一覧します.
このコマンドはレポートを3種類の書式で出力します. `feature.csv`, `feature.json`, ならびに `feature.xlsx`.

| 列                          | 説明                                                   |
|-----------------------------|--------------------------------------------------------|
| upload_api_rate_limit       | 毎月利用可能なアップロードAPIコール回数                |
| upload_api_rate_limit_count | この月に利用されたアップロードAPIコール回数            |
| has_team_shared_dropbox     | このチームが共有されたチームルートを持っているかどうか |
| has_team_file_events        | このチームがファイルイベント機能を持っているかどうか   |
| has_team_selective_sync     | このチームがチーム選択型同期を有効化しているかどうか   |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `feature_0000.xlsx`, `feature_0001.xlsx`, `feature_0002.xlsx`, ...
## レポート: file_request

このレポートはチームメンバーのもつファイルリクエストを一覧します.
このコマンドはレポートを3種類の書式で出力します. `file_request.csv`, `file_request.json`, ならびに `file_request.xlsx`.

| 列                          | 説明                                                                      |
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

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`, ...
## レポート: group

このレポートはチーム内のグループを一覧します.
このコマンドはレポートを3種類の書式で出力します. `group.csv`, `group.json`, ならびに `group.xlsx`.

| 列                    | 説明                                                                                    |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_name            | グループ名称                                                                            |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| member_count          | グループ内のメンバー数                                                                  |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`, ...
## レポート: group_member

このレポートはグループとメンバーを一覧します.
このコマンドはレポートを3種類の書式で出力します. `group_member.csv`, `group_member.json`, ならびに `group_member.xlsx`.

| 列                    | 説明                                                                                    |
|-----------------------|-----------------------------------------------------------------------------------------|
| group_name            | グループ名称                                                                            |
| group_management_type | だれがこのグループを管理できるか (user_managed, company_managed, または system_managed) |
| access_type           | グループにおけるユーザーの役割 (member/owner)                                           |
| email                 | ユーザーのメールアドレス                                                                |
| status                | チームにおけるメンバーのステータス(active/invited/suspended/removed)                    |
| surname               | 名字                                                                                    |
| given_name            | 名                                                                                      |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`, ...
## レポート: info

このレポートはチームの情報を一覧します.
このコマンドはレポートを3種類の書式で出力します. `info.csv`, `info.json`, ならびに `info.xlsx`.

| 列                          | 説明                                                                                                     |
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

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `info_0000.xlsx`, `info_0001.xlsx`, `info_0002.xlsx`, ...
## レポート: linked_app

このレポートは接続済みアプリケーションと利用ユーザーを一覧します.
このコマンドはレポートを3種類の書式で出力します. `linked_app.csv`, `linked_app.json`, ならびに `linked_app.xlsx`.

| 列            | 説明                                                                 |
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

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`, ...
## レポート: member

このレポートはメンバー一覧を出力します.
このコマンドはレポートを3種類の書式で出力します. `member.csv`, `member.json`, ならびに `member.xlsx`.

| 列             | 説明                                                                                            |
|----------------|-------------------------------------------------------------------------------------------------|
| email          | ユーザーのメールアドレス                                                                        |
| email_verified | trueの場合、ユーザーのメールアドレスはユーザーによって所有されていることが確認されています.     |
| status         | チームにおけるメンバーのステータス(active/invited/suspended/removed)                            |
| given_name     | 名                                                                                              |
| surname        | 名字                                                                                            |
| display_name   | ユーザーのDropboxアカウントの表示名称                                                           |
| joined_on      | メンバーがチームに参加した日時.                                                                 |
| role           | ユーザーのチームでの役割 (team_admin, user_management_admin, support_admin, または member_only) |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...
## レポート: member_quota

このレポートはチームメンバーのカスタム容量上限の設定を出力します.
このコマンドはレポートを3種類の書式で出力します. `member_quota.csv`, `member_quota.json`, ならびに `member_quota.xlsx`.

| 列    | 説明                                                                |
|-------|---------------------------------------------------------------------|
| email | ユーザーのメールアドレス                                            |
| quota | カスタムの容量制限GB (1 TB = 1024 GB). 0の場合、容量制限をしません. |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`, ...
## レポート: namespace

このレポートはチームの名前空間を一覧します.
このコマンドはレポートを3種類の書式で出力します. `namespace.csv`, `namespace.json`, ならびに `namespace.xlsx`.

| 列             | 説明                                                                                   |
|----------------|----------------------------------------------------------------------------------------|
| name           | 名前空間の名称                                                                         |
| namespace_type | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder) |
| team_member_id | メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID           |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`, ...
## レポート: namespace_file

このレポートはチームの名前空間を一覧します.
このコマンドはレポートを3種類の書式で出力します. `namespace_file.csv`, `namespace_file.json`, ならびに `namespace_file.xlsx`.

| 列                     | 説明                                                                                            |
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

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`, ...
## レポート: namespace_member

このレポートは名前空間とそのメンバー一覧を出力します.
このコマンドはレポートを3種類の書式で出力します. `namespace_member.csv`, `namespace_member.json`, ならびに `namespace_member.xlsx`.

| 列                 | 説明                                                                                                 |
|--------------------|------------------------------------------------------------------------------------------------------|
| namespace_name     | 名前空間の名称                                                                                       |
| namespace_type     | 名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)               |
| entry_access_type  | ユーザーの共有ファイル・フォルダへのアクセスレベル (owner, editor, viewer, または viewer_no_comment) |
| entry_is_inherited | メンバーのアクセス権限が上位フォルダから継承されている場合true                                       |
| email              | ユーザーのメールアドレス                                                                             |
| display_name       | セッションのタイプ (web_session, desktop_client, または mobile_client)                               |
| group_name         | グループ名称                                                                                         |
| invitee_email      | このフォルダに招待されたメールアドレス                                                               |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `namespace_member_0000.xlsx`, `namespace_member_0001.xlsx`, `namespace_member_0002.xlsx`, ...
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
## レポート: shared_link

このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します.
このコマンドはレポートを3種類の書式で出力します. `shared_link.csv`, `shared_link.json`, ならびに `shared_link.xlsx`.

| 列         | 説明                                                                 |
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

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...
## レポート: team_folder

このレポートはチーム内のチームフォルダを一覧します.
このコマンドはレポートを3種類の書式で出力します. `team_folder.csv`, `team_folder.json`, ならびに `team_folder.xlsx`.

| 列                     | 説明                                                                                 |
|------------------------|--------------------------------------------------------------------------------------|
| name                   | チームフォルダの名称                                                                 |
| status                 | チームフォルダの状態 (active, archived, または archive_in_progress)                  |
| is_team_shared_dropbox |                                                                                      |
| sync_setting           | チームフォルダに設定された同期設定 (default, not_synced, または not_synced_inactive) |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`, ...
## レポート: usage

このレポートはユーザーの現在のストレージ利用容量を出力します.
このコマンドはレポートを3種類の書式で出力します. `usage.csv`, `usage.json`, ならびに `usage.xlsx`.

| 列         | 説明                                                    |
|------------|---------------------------------------------------------|
| email      | アカウントのメールアドレス                              |
| used_gb    | このユーザーの合計利用スペース (in GB, 1GB = 1024 MB).  |
| used_bytes | ユーザーの合計利用要領 (bytes).                         |
| allocation | このユーザーの利用容量の付与先 (individual, or team)    |
| allocated  | このユーザーアカウントに確保されている合計容量 (bytes). |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

