# teamfolder member list

List team folder members 

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
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

# Usage

This document uses the Desktop folder for command example.
## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe teamfolder member list 
```

macOS, Linux:
```
$HOME/Desktop/tbx teamfolder member list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                  | Description                                                                                                                                                                        | Default |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-folder-name`          | Filter by folder name. Filter by exact match to the name.                                                                                                                          |         |
| `-folder-name-prefix`   | Filter by folder name. Filter by name match to the prefix.                                                                                                                         |         |
| `-folder-name-suffix`   | Filter by folder name. Filter by name match to the suffix.                                                                                                                         |         |
| `-member-type-external` | Filter folder members. Keep only members are external (not in the same team). Note: Invited members are marked as external member.                                                 |         |
| `-member-type-internal` | Filter folder members. Keep only members are internal (in the same team). Note: Invited members are marked as external member.                                                     |         |
| `-peer`                 | Account alias                                                                                                                                                                      | default |
| `-scan-timeout`         | Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`. | short   |

## Common options:

| Option            | Description                                                                               | Default              |
|-------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auto-open`      | Auto open URL or artifact folder                                                          | false                |
| `-bandwidth-kb`   | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited           | 0                    |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                            | normal               |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)                      | normal               |
| `-concurrency`    | Maximum concurrency for running operation                                                 | Number of processors |
| `-debug`          | Enable debug mode                                                                         | false                |
| `-experiment`     | Enable experimental feature(s).                                                           |                      |
| `-lang`           | Display language                                                                          | auto                 |
| `-output`         | Output format (none/text/markdown/json)                                                   | text                 |
| `-proxy`          | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy. |                      |
| `-quiet`          | Suppress non-error messages, and make output readable by a machine (JSON format)          | false                |
| `-secure`         | Do not store tokens into a file                                                           | false                |
| `-verbose`        | Show current operations for more detail.                                                  | false                |
| `-workspace`      | Workspace path                                                                            |                      |

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: membership

This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.
The command will generate a report in three different formats. `membership.csv`, `membership.json`, and `membership.xlsx`.

| Column          | Description                                                                                                                          |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| path            | Path                                                                                                                                 |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)                             |
| owner_team_name | Team name of the team that owns the folder                                                                                           |
| access_type     | User's access level for this folder                                                                                                  |
| member_type     | Type of this member (user, group, or invitee)                                                                                        |
| member_name     | Name of this member                                                                                                                  |
| member_email    | Email address of this member                                                                                                         |
| same_team       | Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `membership_0000.xlsx`, `membership_0001.xlsx`, `membership_0002.xlsx`, ...

## Report: no_member

This report shows folders without members.
The command will generate a report in three different formats. `no_member.csv`, `no_member.json`, and `no_member.xlsx`.

| Column          | Description                                                                                              |
|-----------------|----------------------------------------------------------------------------------------------------------|
| owner_team_name | Team name of the team that owns the folder                                                               |
| path            | Path                                                                                                     |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `no_member_0000.xlsx`, `no_member_0001.xlsx`, `no_member_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

