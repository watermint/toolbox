# team content member

チームフォルダや共有フォルダのメンバー一覧 

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
.\tbx.exe team content member 
```

macOS, Linux:
```
$HOME/Desktop/tbx team content member 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                  | Description                                                                                                                        | Default |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-folder-name`          | Filter by folder name. Filter by exact match to the name.                                                                          |         |
| `-folder-name-prefix`   | Filter by folder name. Filter by name match to the prefix.                                                                         |         |
| `-folder-name-suffix`   | Filter by folder name. Filter by name match to the suffix.                                                                         |         |
| `-member-type-external` | Filter folder members. Keep only members are external (not in the same team). Note: Invited members are marked as external member. |         |
| `-member-type-internal` | Filter folder members. Keep only members are internal (in the same team). Note: Invited members are marked as external member.     |         |
| `-peer`                 | アカウントの別名                                                                                                                   | default |

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

## Report: no_member

このレポートはメンバーのいないフォルダの一覧を出力します.
The command will generate a report in three different formats. `no_member.csv`, `no_member.json`, and `no_member.xlsx`.

| Column          | Description                                                                                                      |
|-----------------|------------------------------------------------------------------------------------------------------------------|
| owner_team_name | このフォルダを所有するチームの名前                                                                               |
| path            | パス                                                                                                             |
| folder_type     | フォルダの種別. (`team_folder`: チームフォルダまたはチームフォルダ以下のフォルダ, `shared_folder`: 共有フォルダ) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `no_member_0000.xlsx`, `no_member_0001.xlsx`, `no_member_0002.xlsx`, ...
## Report: membership

このレポートは共有フォルダまたはチームフォルダと、そのメンバーを一覧できます. フォルダに複数メンバーがいる場合には、メンバーは1行ずつ出力されます.
The command will generate a report in three different formats. `membership.csv`, `membership.json`, and `membership.xlsx`.

| Column          | Description                                                                                                                          |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| path            | パス                                                                                                                                 |
| folder_type     | フォルダの種別. (`team_folder`: チームフォルダまたはチームフォルダ以下のフォルダ, `shared_folder`: 共有フォルダ)                     |
| owner_team_name | このフォルダを所有するチームの名前                                                                                                   |
| access_type     | このフォルダに対するユーザーのアクセスレベル                                                                                         |
| member_type     | メンバーの種類 (user, group または invitee)                                                                                          |
| member_name     | このメンバーの名前                                                                                                                   |
| member_email    | このメンバーのメールアドレス                                                                                                         |
| same_team       | Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `membership_0000.xlsx`, `membership_0001.xlsx`, `membership_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

