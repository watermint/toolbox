# dev stage scoped

Dropbox scoped OAuth app test 

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
* Dropbox (Individual account): https://help.dropbox.com/installs-integrations/third-party/third-party-apps
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

## Auth scopes

| Label                     | Description                       |
|---------------------------|-----------------------------------|
| dropbox_scoped_individual | Dropbox Individual account access |
| dropbox_scoped_team       | Dropbox team access               |

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
.\tbx.exe dev stage scoped 
```

macOS, Linux:
```
$HOME/Desktop/tbx dev stage scoped 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option        | Description                  | Default                                     |
|---------------|------------------------------|---------------------------------------------|
| `-individual` | Account alias for individual | &{Individual [files.content.read] <nil>}    |
| `-team`       | Account alias for team       | &{Team [members.read team_info.read] <nil>} |

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

## Report: file_list

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `file_list.csv`, `file_list.json`, and `file_list.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| id                      | A unique identifier for the file.                                                                      |
| tag                     | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                    | The last component of the path (including extension).                                                  |
| path_lower              | The lowercased full path in the user's Dropbox. This always starts with a slash.                       |
| path_display            | The cased path to be used for display purposes only.                                                   |
| client_modified         | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified         | The last time the file was modified on Dropbox.                                                        |
| revision                | A unique identifier for the current revision of a file.                                                |
| size                    | The file size in bytes.                                                                                |
| content_hash            | A hash of the file content.                                                                            |
| shared_folder_id        | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.   |
| parent_shared_folder_id | ID of shared folder that holds this file.                                                              |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `file_list_0000.xlsx`, `file_list_0001.xlsx`, `file_list_0002.xlsx`, ...

## Report: member_list

This report shows a list of members.
The command will generate a report in three different formats. `member_list.csv`, `member_list.json`, and `member_list.xlsx`.

| Column           | Description                                                                                                          |
|------------------|----------------------------------------------------------------------------------------------------------------------|
| team_member_id   | ID of user as a member of a team.                                                                                    |
| email            | Email address of user.                                                                                               |
| email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| given_name       | Also known as a first name                                                                                           |
| surname          | Also known as a last name or family name.                                                                            |
| familiar_name    | Locale-dependent name                                                                                                |
| display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| abbreviated_name | An abbreviated form of the person's name.                                                                            |
| member_folder_id | The namespace id of the user's root folder.                                                                          |
| external_id      | External ID that a team can attach to the user.                                                                      |
| account_id       | A user's account identifier.                                                                                         |
| persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| tag              | Operation tag                                                                                                        |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `member_list_0000.xlsx`, `member_list_0001.xlsx`, `member_list_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

