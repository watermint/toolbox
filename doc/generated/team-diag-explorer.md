# team diag explorer

Report whole team information 

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
| `-all`  | Include additional reports          | false   |
| `-file` | Dropbox Business file access        | default |
| `-info` | Dropbox Business information access | default |
| `-mgmt` | Dropbox Business management         | default |

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

## Report: device

This report shows a list of current existing sessions in the team with team member information.
The command will generate a report in three different formats. `device.csv`, `device.json`, and `device.xlsx`.

| Column                        | Description                                                                          |
|-------------------------------|--------------------------------------------------------------------------------------|
| team_member_id                | ID of user as a member of a team.                                                    |
| email                         | Email address of user.                                                               |
| status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name                    | Also known as a first name                                                           |
| surname                       | Also known as a last name or family name.                                            |
| display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  |
| id                            | The session id.                                                                      |
| user_agent                    | Information on the hosting device.                                                   |
| os                            | Information on the hosting operating system                                          |
| browser                       | Information on the browser used for this web session.                                |
| ip_address                    | The IP address of the last activity from this session.                               |
| country                       | The country from which the last activity from this session was made.                 |
| created                       | The time this session was created.                                                   |
| updated                       | The time of the last activity from this session.                                     |
| expires                       | The time this session expires                                                        |
| host_name                     | Name of the hosting desktop.                                                         |
| client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             |
| client_version                | The Dropbox client version.                                                          |
| platform                      | Information on the hosting platform.                                                 |
| is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             |
| device_name                   | The device name.                                                                     |
| os_version                    | The hosting OS version.                                                              |
| last_carrier                  | Last carrier used by the device.                                                     |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`, ...

## Report: errors

This report shows the transaction result.
The command will generate a report in three different formats. `errors.csv`, `errors.json`, and `errors.xlsx`.

| Column          | Description                            |
|-----------------|----------------------------------------|
| status          | Status of the operation                |
| reason          | Reason of failure or skipped operation |
| input.namespace | Namespace                              |
| input.path      | Path                                   |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...

## Report: feature

This report shows a list of team features and their settings.
The command will generate a report in three different formats. `feature.csv`, `feature.json`, and `feature.xlsx`.

| Column                      | Description                                       |
|-----------------------------|---------------------------------------------------|
| upload_api_rate_limit       | The number of upload API calls allowed per month. |
| upload_api_rate_limit_count | The number of upload API called this month.       |
| has_team_shared_dropbox     | Does this team have a shared team root.           |
| has_team_file_events        | Does this team have file events.                  |
| has_team_selective_sync     | Does this team have team selective sync enabled.  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `feature_0000.xlsx`, `feature_0001.xlsx`, `feature_0002.xlsx`, ...

## Report: file_request

This report shows a list of file requests with the file request owner team member.
The command will generate a report in three different formats. `file_request.csv`, `file_request.json`, and `file_request.xlsx`.

| Column                      | Description                                                                   |
|-----------------------------|-------------------------------------------------------------------------------|
| email                       | Email address of this file request owner.                                     |
| status                      | The user status of this file request owner (active/invited/suspended/removed) |
| surname                     | Surname of this file request owner.                                           |
| given_name                  | Given name of this file request owner.                                        |
| url                         | The URL of the file request.                                                  |
| title                       | The title of the file request.                                                |
| created                     | When this file request was created.                                           |
| is_open                     | Whether or not the file request is open.                                      |
| file_count                  | The number of files this file request has received.                           |
| destination                 | The path of the folder in the Dropbox where uploaded files will be sent       |
| deadline                    | The deadline for this file request.                                           |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`, ...

## Report: group

This report shows a list of groups in the team.
The command will generate a report in three different formats. `group.csv`, `group.json`, and `group.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| member_count          | The number of members in the group.                                                   |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`, ...

## Report: group_member

This report shows a list of groups and their members.
The command will generate a report in three different formats. `group_member.csv`, `group_member.json`, and `group_member.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group.                                                                      |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| access_type           | The role that the user has in the group (member/owner)                                |
| email                 | Email address of user.                                                                |
| status                | The user's status as a member of a specific team. (active/invited/suspended/removed)  |
| surname               | Also known as a last name or family name.                                             |
| given_name            | Also known as a first name                                                            |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`, ...

## Report: info

This report shows a list of team information.
The command will generate a report in three different formats. `info.csv`, `info.json`, and `info.xlsx`.

| Column                      | Description                                                                                                   |
|-----------------------------|---------------------------------------------------------------------------------------------------------------|
| name                        | The name of the team                                                                                          |
| team_id                     | The ID of the team.                                                                                           |
| num_licensed_users          | The number of licenses available to the team.                                                                 |
| num_provisioned_users       | The number of accounts that have been invited or are already active members of the team.                      |
| policy_shared_folder_member | Which shared folders team members can join (from_team_only, or from_anyone)                                   |
| policy_shared_folder_join   | Who can join folders shared by team members (team, or anyone)                                                 |
| policy_shared_link_create   | Who can view shared links owned by team members (default_public, default_team_only, or team_only)             |
| policy_emm_state            | This describes the Enterprise Mobility Management (EMM) state for this team (disabled, optional, or required) |
| policy_office_add_in        | The admin policy around the Dropbox Office Add-In for this team (disabled, or enabled)                        |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `info_0000.xlsx`, `info_0001.xlsx`, `info_0002.xlsx`, ...

## Report: linked_app

This report shows a list of linked app with the user of the app.
The command will generate a report in three different formats. `linked_app.csv`, `linked_app.json`, and `linked_app.xlsx`.

| Column        | Description                                                                          |
|---------------|--------------------------------------------------------------------------------------|
| email         | Email address of user.                                                               |
| status        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name    | Also known as a first name                                                           |
| surname       | Also known as a last name or family name.                                            |
| display_name  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| app_name      | The application name.                                                                |
| is_app_folder | Whether the linked application uses a dedicated folder.                              |
| publisher     | The publisher's URL.                                                                 |
| publisher_url | The application publisher name.                                                      |
| linked        | The time this application was linked                                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`, ...

## Report: member

This report shows a list of members.
The command will generate a report in three different formats. `member.csv`, `member.json`, and `member.xlsx`.

| Column         | Description                                                                                    |
|----------------|------------------------------------------------------------------------------------------------|
| email          | Email address of user.                                                                         |
| email_verified | Is true if the user's email is verified to be owned by the user.                               |
| status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| given_name     | Also known as a first name                                                                     |
| surname        | Also known as a last name or family name.                                                      |
| display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| joined_on      | The date and time the user joined as a member of a specific team.                              |
| role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...

## Report: member_quota

This report shows a list of custom quota settings for each team members.
The command will generate a report in three different formats. `member_quota.csv`, `member_quota.json`, and `member_quota.xlsx`.

| Column | Description                                                                 |
|--------|-----------------------------------------------------------------------------|
| email  | Email address of user.                                                      |
| quota  | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`, ...

## Report: namespace

This report shows a list of namespaces in the team.
The command will generate a report in three different formats. `namespace.csv`, `namespace.json`, and `namespace.xlsx`.

| Column         | Description                                                                                |
|----------------|--------------------------------------------------------------------------------------------|
| name           | The name of this namespace                                                                 |
| namespace_type | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`, ...

## Report: namespace_file

This report shows a list of namespaces in the team.
The command will generate a report in three different formats. `namespace_file.csv`, `namespace_file.json`, and `namespace_file.xlsx`.

| Column                 | Description                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------|
| namespace_type         | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)             |
| namespace_name         | The name of this namespace                                                                             |
| namespace_member_email | If this is a team member or app folder, the email address of the owning team member.                   |
| tag                    | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                   | The last component of the path (including extension).                                                  |
| path_display           | The cased path to be used for display purposes only.                                                   |
| client_modified        | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified        | The last time the file was modified on Dropbox.                                                        |
| size                   | The file size in bytes.                                                                                |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`, ...

## Report: namespace_member

This report shows a list of members of namespaces in the team.
The command will generate a report in three different formats. `namespace_member.csv`, `namespace_member.json`, and `namespace_member.xlsx`.

| Column             | Description                                                                                               |
|--------------------|-----------------------------------------------------------------------------------------------------------|
| namespace_name     | The name of this namespace                                                                                |
| namespace_type     | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)                |
| entry_access_type  | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| entry_is_inherited | True if the member has access from a parent folder                                                        |
| email              | Email address of user.                                                                                    |
| display_name       | Team member display name.                                                                                 |
| group_name         | Name of the group                                                                                         |
| invitee_email      | Email address of invitee for this folder                                                                  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_member_0000.xlsx`, `namespace_member_0001.xlsx`, `namespace_member_0002.xlsx`, ...

## Report: namespace_size

Namespace size
The command will generate a report in three different formats. `namespace_size.csv`, `namespace_size.json`, and `namespace_size.xlsx`.

| Column               | Description                                                                                |
|----------------------|--------------------------------------------------------------------------------------------|
| namespace_name       | The name of this namespace                                                                 |
| namespace_id         | The ID of this namespace.                                                                  |
| namespace_type       | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| owner_team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |
| path                 | Path to the folder                                                                         |
| count_file           | Number of files under the folder                                                           |
| count_folder         | Number of folders under the folder                                                         |
| count_descendant     | Number of files and folders under the folder                                               |
| size                 | Size of the folder                                                                         |
| depth                | Folder depth                                                                               |
| mod_time_earliest    | The earliest modification time of a file in this folder or child folders.                  |
| mod_time_latest      | The latest modification time of a file in this folder or child folders                     |
| api_complexity       | Folder complexity index for API operations                                                 |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`, ...

## Report: shared_link

This report shows a list of shared links with the shared link owner team member.
The command will generate a report in three different formats. `shared_link.csv`, `shared_link.json`, and `shared_link.xlsx`.

| Column     | Description                                                                                                                                                                                                             |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| tag        | Entry type (file, or folder)                                                                                                                                                                                            |
| url        | URL of the shared link.                                                                                                                                                                                                 |
| name       | The linked file name (including extension).                                                                                                                                                                             |
| expires    | Expiration time, if set.                                                                                                                                                                                                |
| path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                         |
| visibility | The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| email      | Email address of user.                                                                                                                                                                                                  |
| status     | The user's status as a member of a specific team. (active/invited/suspended/removed)                                                                                                                                    |
| surname    | Surname of the link owner                                                                                                                                                                                               |
| given_name | Given name of the link owner                                                                                                                                                                                            |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...

## Report: team_folder

This report shows a list of team folders in the team.
The command will generate a report in three different formats. `team_folder.csv`, `team_folder.json`, and `team_folder.xlsx`.

| Column                 | Description                                                                                |
|------------------------|--------------------------------------------------------------------------------------------|
| name                   | The name of the team folder.                                                               |
| status                 | The status of the team folder (active, archived, or archive_in_progress)                   |
| is_team_shared_dropbox |                                                                                            |
| sync_setting           | The sync setting applied to this team folder (default, not_synced, or not_synced_inactive) |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`, ...

## Report: usage

This report shows current storage usage of users.
The command will generate a report in three different formats. `usage.csv`, `usage.json`, and `usage.xlsx`.

| Column     | Description                                              |
|------------|----------------------------------------------------------|
| email      | Email address of the account                             |
| used_gb    | The user's total space usage (in GB, 1GB = 1024 MB).     |
| used_bytes | The user's total space usage (bytes).                    |
| allocation | The user's space allocation (individual, or team)        |
| allocated  | The total space allocated to the user's account (bytes). |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

