# teamfolder member add

Batch adding users/groups to team folders (Irreversible operation)

This command will do (1) create new team folders or new sub-folders if the team folder does not exist. The command does
not (2) change access inheritance setting of any folders, (3) create a group if that not exist. This command is designed
to be idempotent. You can safely retry if any errors happen on the operation. The command will not report an error to
keep idempotence. For example, the command will not report an error like, the member already have access to the folder.

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS      | Path                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support. You can delete those files after use if you want to
remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

## Auth scopes

| Label               | Description         |
|---------------------|---------------------|
| dropbox_scoped_team | Dropbox team access |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it
into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code.
Please copy that code and paste it to the `tbx`.
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
.\tbx.exe teamfolder member add -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx teamfolder member add -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please
select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "
General" tab. You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "
Open" on the dialogue.

## Options:

| Option              | Description                              | Default                 |
|---------------------|------------------------------------------|-------------------------|
| `-admin-group-name` | Temporary group name for admin operation | watermint-toolbox-admin |
| `-file`             | Path to data file                        |                         |
| `-peer`             | Account alias                            | default                 |

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

# File formats

## Format: File

Team folder and member list for adding access. Each row can have one member and the one folder. If you want to add two
or more members to the folder, please create rows for those members. Similarly, if you want to add a member to two or
more folders, please create rows for those folders.

| Column                     | Description                                                                                                  | Example |
|----------------------------|--------------------------------------------------------------------------------------------------------------|---------|
| team_folder_name           | Team folder name                                                                                             | Sales   |
| path                       | Relative path from the team folder root. Leave empty if you want to add a member to root of the team folder. | Report  |
| access_type                | Access type (viewer/editor)                                                                                  | editor  |
| group_name_or_member_email | Group name or member email address                                                                           | Sales   |

The first line is a header line. The program will accept a file without the header.
```
team_folder_name,path,access_type,group_name_or_member_email
Sales,Report,editor,Sales
```

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see
path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: operation_log

This report shows the transaction result. The command will generate a report in three different
formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                           | Description                                                                                                  |
|----------------------------------|--------------------------------------------------------------------------------------------------------------|
| status                           | Status of the operation                                                                                      |
| reason                           | Reason of failure or skipped operation                                                                       |
| input.team_folder_name           | Team folder name                                                                                             |
| input.path                       | Relative path from the team folder root. Leave empty if you want to add a member to root of the team folder. |
| input.access_type                | Access type (viewer/editor)                                                                                  |
| input.group_name_or_member_email | Group name or member email address                                                                           |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like
follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you
want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't
support proxies which require authentication.

