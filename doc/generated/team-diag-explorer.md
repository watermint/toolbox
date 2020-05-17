# team diag explorer 

Report whole team information 

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

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option  | Description                         | Default |
|---------|-------------------------------------|---------|
| `-all`  | Include additional reports          | false   |
| `-file` | Dropbox Business file access        | default |
| `-info` | Dropbox Business information access | default |
| `-mgmt` | Dropbox Business management         | default |

Common options:

| Option            | Description                                                                      | Default              |
|-------------------|----------------------------------------------------------------------------------|----------------------|
| `-auto-open`      | Auto open URL or artifact folder                                                 | false                |
| `-bandwidth-kb`   | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited  | 0                    |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                   | normal               |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)             | normal               |
| `-concurrency`    | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`          | Enable debug mode                                                                | false                |
| `-output`         | Output format (none/text/markdown/json)                                          | text                 |
| `-proxy`          | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`          | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`         | Do not store tokens into a file                                                  | false                |
| `-workspace`      | Workspace path                                                                   |                      |

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: device 
This report shows a list of current existing sessions in the team with team member information.
Report files are generated in three formats like below;
* `device.csv`
* `device.xlsx`
* `device.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`...   

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

## Report: feature 
This report shows a list of team features and their settings.
Report files are generated in three formats like below;
* `feature.csv`
* `feature.xlsx`
* `feature.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`feature_0000.xlsx`, `feature_0001.xlsx`, `feature_0002.xlsx`...   

| Column                      | Description                                       |
|-----------------------------|---------------------------------------------------|
| upload_api_rate_limit       | The number of upload API calls allowed per month. |
| upload_api_rate_limit_count | The number of upload API called this month.       |
| has_team_shared_dropbox     | Does this team have a shared team root.           |
| has_team_file_events        | Does this team have file events.                  |
| has_team_selective_sync     | Does this team have team selective sync enabled.  |

## Report: file_request 
This report shows a list of file requests with the file request owner team member.
Report files are generated in three formats like below;
* `file_request.csv`
* `file_request.xlsx`
* `file_request.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`...   

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

## Report: group 
This report shows a list of groups in the team.
Report files are generated in three formats like below;
* `group.csv`
* `group.xlsx`
* `group.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`...   

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| member_count          | The number of members in the group.                                                   |

## Report: group_member 
This report shows a list of groups and their members.
Report files are generated in three formats like below;
* `group_member.csv`
* `group_member.xlsx`
* `group_member.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`...   

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group.                                                                      |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| access_type           | The role that the user has in the group (member/owner)                                |
| email                 | Email address of user.                                                                |
| status                | The user's status as a member of a specific team. (active/invited/suspended/removed)  |
| surname               | Also known as a last name or family name.                                             |
| given_name            | Also known as a first name                                                            |

## Report: info 
This report shows a list of team information.
Report files are generated in three formats like below;
* `info.csv`
* `info.xlsx`
* `info.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`info_0000.xlsx`, `info_0001.xlsx`, `info_0002.xlsx`...   

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

## Report: linked_app 
This report shows a list of linked app with the user of the app.
Report files are generated in three formats like below;
* `linked_app.csv`
* `linked_app.xlsx`
* `linked_app.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`...   

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

## Report: member 
This report shows a list of members.
Report files are generated in three formats like below;
* `member.csv`
* `member.xlsx`
* `member.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`...   

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

## Report: member_quota 
This report shows a list of custom quota settings for each team members.
Report files are generated in three formats like below;
* `member_quota.csv`
* `member_quota.xlsx`
* `member_quota.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`...   

| Column | Description                                                                 |
|--------|-----------------------------------------------------------------------------|
| email  | Email address of user.                                                      |
| quota  | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. |

## Report: namespace 
This report shows a list of namespaces in the team.
Report files are generated in three formats like below;
* `namespace.csv`
* `namespace.xlsx`
* `namespace.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`...   

| Column         | Description                                                                                |
|----------------|--------------------------------------------------------------------------------------------|
| name           | The name of this namespace                                                                 |
| namespace_type | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |

## Report: namespace_file 
This report shows a list of namespaces in the team.
Report files are generated in three formats like below;
* `namespace_file.csv`
* `namespace_file.xlsx`
* `namespace_file.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`...   

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

## Report: namespace_member 
This report shows a list of members of namespaces in the team.
Report files are generated in three formats like below;
* `namespace_member.csv`
* `namespace_member.xlsx`
* `namespace_member.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_member_0000.xlsx`, `namespace_member_0001.xlsx`, `namespace_member_0002.xlsx`...   

| Column             | Description                                                                                               |
|--------------------|-----------------------------------------------------------------------------------------------------------|
| namespace_name     | The name of this namespace                                                                                |
| namespace_type     | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)                |
| entry_access_type  | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| entry_is_inherited | True if the member has access from a parent folder                                                        |
| email              | Email address of user.                                                                                    |
| display_name       | Type of the session (web_session, desktop_client, or mobile_client)                                       |
| group_name         | Name of the group                                                                                         |
| invitee_email      | Email address of invitee for this folder                                                                  |

## Report: namespace_size 
This report shows the transaction result.
Report files are generated in three formats like below;
* `namespace_size.csv`
* `namespace_size.xlsx`
* `namespace_size.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`...   

| Column                  | Description                                                                                |
|-------------------------|--------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                    |
| reason                  | Reason of failure or skipped operation                                                     |
| input.name              | The name of this namespace                                                                 |
| input.namespace_type    | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| result.path             | Path to the folder                                                                         |
| result.count_file       | Number of files under the folder                                                           |
| result.count_folder     | Number of folders under the folder                                                         |
| result.count_descendant | Number of files and folders under the folder                                               |
| result.size             | Size of the folder                                                                         |
| result.api_complexity   | Folder complexity index for API operations                                                 |

## Report: shared_link 
This report shows a list of shared links with the shared link owner team member.
Report files are generated in three formats like below;
* `shared_link.csv`
* `shared_link.xlsx`
* `shared_link.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`...   

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

## Report: team_folder 
This report shows a list of team folders in the team.
Report files are generated in three formats like below;
* `team_folder.csv`
* `team_folder.xlsx`
* `team_folder.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`...   

| Column                 | Description                                                                                |
|------------------------|--------------------------------------------------------------------------------------------|
| name                   | The name of the team folder.                                                               |
| status                 | The status of the team folder (active, archived, or archive_in_progress)                   |
| is_team_shared_dropbox |                                                                                            |
| sync_setting           | The sync setting applied to this team folder (default, not_synced, or not_synced_inactive) |

## Report: usage 
This report shows current storage usage of users.
Report files are generated in three formats like below;
* `usage.csv`
* `usage.xlsx`
* `usage.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`...   

| Column     | Description                                              |
|------------|----------------------------------------------------------|
| email      | Email address of the account                             |
| used_gb    | The user's total space usage (in GB, 1GB = 1024 MB).     |
| used_bytes | The user's total space usage (bytes).                    |
| allocation | The user's space allocation (individual, or team)        |
| allocated  | The total space allocated to the user's account (bytes). |

