# team diag explorer 

Report whole team information 

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

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

| Option          | Description                                                                      | Default              |
|-----------------|----------------------------------------------------------------------------------|----------------------|
| `-auto-open`    | Auto open URL or artifact folder                                                 | false                |
| `-bandwidth-kb` | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited  | 0                    |
| `-concurrency`  | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`        | Enable debug mode                                                                | false                |
| `-low-memory`   | Low memory footprint mode                                                        | false                |
| `-output`       | Output format (none/text/markdown/json)                                          | text                 |
| `-proxy`        | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`        | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`       | Do not store tokens into a file                                                  | false                |
| `-workspace`    | Workspace path                                                                   |                      |

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

Report files are generated in three formats like below;
* `device.csv`
* `device.xlsx`
* `device.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`...   

| Column                        | Description                                                                          |
|-------------------------------|--------------------------------------------------------------------------------------|
| team_member_id                | ID of user as a member of a team.                                                    |
| email                         | Email address of user.                                                               |
| status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name                    | Also known as a first name                                                           |
| surname                       | Also known as a last name or family name.                                            |
| familiar_name                 | Locale-dependent name                                                                |
| display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| abbreviated_name              | An abbreviated form of the person's name.                                            |
| external_id                   | External ID that a team can attach to the user.                                      |
| account_id                    | A user's account identifier.                                                         |
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

Report files are generated in three formats like below;
* `feature.csv`
* `feature.xlsx`
* `feature.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

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

Report files are generated in three formats like below;
* `file_request.csv`
* `file_request.xlsx`
* `file_request.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`...   

| Column                      | Description                                                                   |
|-----------------------------|-------------------------------------------------------------------------------|
| account_id                  | Account ID of this file request owner.                                        |
| team_member_id              | ID of file request owner user as a member of a team                           |
| email                       | Email address of this file request owner.                                     |
| status                      | The user status of this file request owner (active/invited/suspended/removed) |
| surname                     | Surname of this file request owner.                                           |
| given_name                  | Given name of this file request owner.                                        |
| file_request_id             | The ID of the file request.                                                   |
| url                         | The URL of the file request.                                                  |
| title                       | The title of the file request.                                                |
| created                     | When this file request was created.                                           |
| is_open                     | Whether or not the file request is open.                                      |
| file_count                  | The number of files this file request has received.                           |
| destination                 | The path of the folder in the Dropbox where uploaded files will be sent       |
| deadline                    | The deadline for this file request.                                           |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           |

## Report: group 

Report files are generated in three formats like below;
* `group.csv`
* `group.xlsx`
* `group.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`...   

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_id              | A group's identifier                                                                  |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| group_external_id     |  External ID of group. This is an arbitrary ID that an admin can attach to a group.   |
| member_count          | The number of members in the group.                                                   |

## Report: group_member 

Report files are generated in three formats like below;
* `group_member.csv`
* `group_member.xlsx`
* `group_member.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`...   

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_id              | A group's identifier                                                                  |
| group_name            | Name of a group.                                                                      |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| access_type           | The role that the user has in the group (member/owner)                                |
| account_id            | A user's account identifier                                                           |
| team_member_id        | ID of user as a member of a team.                                                     |
| email                 | Email address of user.                                                                |
| status                | The user's status as a member of a specific team. (active/invited/suspended/removed)  |
| surname               | Also known as a last name or family name.                                             |
| given_name            | Also known as a first name                                                            |

## Report: info 

Report files are generated in three formats like below;
* `info.csv`
* `info.xlsx`
* `info.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

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

Report files are generated in three formats like below;
* `linked_app.csv`
* `linked_app.xlsx`
* `linked_app.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`...   

| Column           | Description                                                                          |
|------------------|--------------------------------------------------------------------------------------|
| team_member_id   | ID of user as a member of a team.                                                    |
| email            | Email address of user.                                                               |
| status           | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name       | Also known as a first name                                                           |
| surname          | Also known as a last name or family name.                                            |
| familiar_name    | Locale-dependent name                                                                |
| display_name     | A name that can be used directly to represent the name of a user's Dropbox account.  |
| abbreviated_name | An abbreviated form of the person's name.                                            |
| external_id      | External ID that a team can attach to the user.                                      |
| account_id       | A user's account identifier.                                                         |
| app_id           | The application unique id.                                                           |
| app_name         | The application name.                                                                |
| is_app_folder    | Whether the linked application uses a dedicated folder.                              |
| publisher        | The publisher's URL.                                                                 |
| publisher_url    | The application publisher name.                                                      |
| linked           | The time this application was linked                                                 |

## Report: member 

Report files are generated in three formats like below;
* `member.csv`
* `member.xlsx`
* `member.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`...   

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

## Report: member_quota 

Report files are generated in three formats like below;
* `member_quota.csv`
* `member_quota.xlsx`
* `member_quota.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`...   

| Column | Description                                                                 |
|--------|-----------------------------------------------------------------------------|
| email  | Email address of user.                                                      |
| quota  | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. |

## Report: namespace 

Report files are generated in three formats like below;
* `namespace.csv`
* `namespace.xlsx`
* `namespace.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`...   

| Column         | Description                                                                                |
|----------------|--------------------------------------------------------------------------------------------|
| name           | The name of this namespace                                                                 |
| namespace_id   | The ID of this namespace.                                                                  |
| namespace_type | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |

## Report: namespace_file 

Report files are generated in three formats like below;
* `namespace_file.csv`
* `namespace_file.xlsx`
* `namespace_file.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`...   

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| namespace_type          | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)             |
| namespace_id            | The ID of this namespace.                                                                              |
| namespace_name          | The name of this namespace                                                                             |
| namespace_member_email  | If this is a team member or app folder, the email address of the owning team member.                   |
| file_id                 | A unique identifier for the file.                                                                      |
| tag                     | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                    | The last component of the path (including extension).                                                  |
| path_display            | The cased path to be used for display purposes only.                                                   |
| client_modified         | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified         | The last time the file was modified on Dropbox.                                                        |
| revision                | A unique identifier for the current revision of a file.                                                |
| size                    | The file size in bytes.                                                                                |
| content_hash            | A hash of the file content.                                                                            |
| shared_folder_id        | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.   |
| parent_shared_folder_id | Set if the folder is contained by a shared folder.                                                     |

## Report: namespace_size 

Report files are generated in three formats like below;
* `namespace_size.csv`
* `namespace_size.xlsx`
* `namespace_size.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`...   

| Column                      | Description                                                                                |
|-----------------------------|--------------------------------------------------------------------------------------------|
| status                      | Status of the operation                                                                    |
| reason                      | Reason of failure or skipped operation                                                     |
| input.name                  | The name of this namespace                                                                 |
| input.namespace_id          | The ID of this namespace.                                                                  |
| input.namespace_type        | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| input.team_member_id        | If this is a team member or app folder, the ID of the owning team member.                  |
| result.namespace_name       | The name of this namespace                                                                 |
| result.namespace_id         | The ID of this namespace.                                                                  |
| result.namespace_type       | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| result.owner_team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |
| result.path                 | Path to the folder                                                                         |
| result.count_file           | Number of files under the folder                                                           |
| result.count_folder         | Number of folders under the folder                                                         |
| result.count_descendant     | Number of files and folders under the folder                                               |
| result.size                 | Size of the folder                                                                         |
| result.api_complexity       | Folder complexity index for API operations                                                 |

## Report: shared_link 

Report files are generated in three formats like below;
* `shared_link.csv`
* `shared_link.xlsx`
* `shared_link.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`...   

| Column         | Description                                                                                                                                                                                                             |
|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| shared_link_id | A unique identifier for the linked file or folder                                                                                                                                                                       |
| tag            | Entry type (file, or folder)                                                                                                                                                                                            |
| url            | URL of the shared link.                                                                                                                                                                                                 |
| name           | The linked file name (including extension).                                                                                                                                                                             |
| expires        | Expiration time, if set.                                                                                                                                                                                                |
| path_lower     | The lowercased full path in the user's Dropbox.                                                                                                                                                                         |
| visibility     | The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| account_id     | A user's account identifier.                                                                                                                                                                                            |
| team_member_id | ID of user as a member of a team.                                                                                                                                                                                       |
| email          | Email address of user.                                                                                                                                                                                                  |
| status         | The user's status as a member of a specific team. (active/invited/suspended/removed)                                                                                                                                    |
| surname        | Surname of the link owner                                                                                                                                                                                               |
| given_name     | Given name of the link owner                                                                                                                                                                                            |

## Report: usage 

Report files are generated in three formats like below;
* `usage.csv`
* `usage.xlsx`
* `usage.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`...   

| Column     | Description                                              |
|------------|----------------------------------------------------------|
| email      | Email address of the account                             |
| used_gb    | The user's total space usage (in GB, 1GB = 1024 MB).     |
| used_bytes | The user's total space usage (bytes).                    |
| allocation | The user's space allocation (individual, or team)        |
| allocated  | The total space allocated to the user's account (bytes). |

