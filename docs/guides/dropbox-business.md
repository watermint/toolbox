---
layout: page
title: Commands of Dropbox for teams
lang: en
---

# Member management commands

## Information commands

Below commands are to retrieve information about team members.

| Command                                                                                             | Description                    |
|-----------------------------------------------------------------------------------------------------|--------------------------------|
| [dropbox team member list]({{ site.baseurl }}/commands/dropbox-team-member-list.html)               | List team member(s)            |
| [dropbox team member feature]({{ site.baseurl }}/commands/dropbox-team-member-feature.html)         | List member feature settings   |
| [dropbox team member folder list]({{ site.baseurl }}/commands/dropbox-team-member-folder-list.html) | List folders for each member   |
| [dropbox team member quota list]({{ site.baseurl }}/commands/dropbox-team-member-quota-list.html)   | List team member quota         |
| [dropbox team member quota usage]({{ site.baseurl }}/commands/dropbox-team-member-quota-usage.html) | List team member storage usage |
| [dropbox team activity user]({{ site.baseurl }}/commands/dropbox-team-activity-user.html)           | Activities log per user        |

## Basic management commands

Below commands are for managing team member accounts. Those commands are for a bulk operation by a CSV file.

| Command                                                                                                                   | Description                                           |
|---------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| [dropbox team member batch invite]({{ site.baseurl }}/commands/dropbox-team-member-batch-invite.html)                     | Invite member(s)                                      |
| [dropbox team member batch delete]({{ site.baseurl }}/commands/dropbox-team-member-batch-delete.html)                     | Delete members                                        |
| [dropbox team member batch detach]({{ site.baseurl }}/commands/dropbox-team-member-batch-detach.html)                     | Convert Dropbox for teams accounts to a Basic account |
| [dropbox team member batch reinvite]({{ site.baseurl }}/commands/dropbox-team-member-batch-reinvite.html)                 | Reinvite invited status members to the team           |
| [dropbox team member update batch email]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-email.html)         | Member email operation                                |
| [dropbox team member update batch profile]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-profile.html)     | Batch update member profiles                          |
| [dropbox team member update batch visible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-visible.html)     | Disable directory restriction to members              |
| [dropbox team member update batch invisible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-invisible.html) | Enable directory restriction to members               |
| [dropbox team member quota batch update]({{ site.baseurl }}/commands/dropbox-team-member-quota-batch-update.html)         | Update team member quota                              |

## Member profile setting commands

Member profile commands are for bulk updating member profile information.
If you need to update the members' email addresses, use the 'member update email` command. The command 'member update email` receives a CSV file to bulk update email addresses.
If you need to update the member's display name, use the 'member update profile` command.

| Command                                                                                                               | Description                  |
|-----------------------------------------------------------------------------------------------------------------------|------------------------------|
| [dropbox team member update batch email]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-email.html)     | Member email operation       |
| [dropbox team member update batch profile]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-profile.html) | Batch update member profiles |

## Member storage quota control commands

You can see existing member storage quota setting or usage by the `dropbox team member quota list` and `dropbox team member quota usage` command. If you need to update member quota, use the `dropbox team member quota update` command. The command `dropbox team member quota update` receives CSV input for bulk updating storage quota setting.

| Command                                                                                                           | Description                    |
|-------------------------------------------------------------------------------------------------------------------|--------------------------------|
| [dropbox team member quota list]({{ site.baseurl }}/commands/dropbox-team-member-quota-list.html)                 | List team member quota         |
| [dropbox team member quota usage]({{ site.baseurl }}/commands/dropbox-team-member-quota-usage.html)               | List team member storage usage |
| [dropbox team member quota batch update]({{ site.baseurl }}/commands/dropbox-team-member-quota-batch-update.html) | Update team member quota       |

## There are two types of commands available for suspending/unsuspending members. If you wanted to suspend/unsuspend a member one by one, please use `dropbox team member suspend` or `dropbox team member unsuspend`. Otherwise, if you want to suspend/unsuspend members through a CSV file, please use the `dropbox team member batch suspend` or `dropbox member batch unsuspend` command.

Suspend/unsuspend a member

| Command                                                                                                     | Description            |
|-------------------------------------------------------------------------------------------------------------|------------------------|
| [dropbox team member suspend]({{ site.baseurl }}/commands/dropbox-team-member-suspend.html)                 | Suspend a member       |
| [dropbox team member unsuspend]({{ site.baseurl }}/commands/dropbox-team-member-unsuspend.html)             | Unsuspend a member     |
| [dropbox team member batch suspend]({{ site.baseurl }}/commands/dropbox-team-member-batch-suspend.html)     | Bulk suspend members   |
| [dropbox team member batch unsuspend]({{ site.baseurl }}/commands/dropbox-team-member-batch-unsuspend.html) | Bulk unsuspend members |

## Directory restriction commands

Directory restriction is the Dropbox for teams feature to hide a member from others. Below commands update this setting to hide or unhide members from others.

| Command                                                                                                                   | Description                              |
|---------------------------------------------------------------------------------------------------------------------------|------------------------------------------|
| [dropbox team member update batch visible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-visible.html)     | Disable directory restriction to members |
| [dropbox team member update batch invisible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-invisible.html) | Enable directory restriction to members  |

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                                                                             | Description                  |
|-----------------------------------------------------------------------------------------------------|------------------------------|
| [dropbox team group add]({{ site.baseurl }}/commands/dropbox-team-group-add.html)                   | Create new group             |
| [dropbox team group batch add]({{ site.baseurl }}/commands/dropbox-team-group-batch-add.html)       | Bulk adding groups           |
| [dropbox team group batch delete]({{ site.baseurl }}/commands/dropbox-team-group-batch-delete.html) | Delete groups                |
| [dropbox team group delete]({{ site.baseurl }}/commands/dropbox-team-group-delete.html)             | Delete group                 |
| [dropbox team group list]({{ site.baseurl }}/commands/dropbox-team-group-list.html)                 | List group(s)                |
| [dropbox team group rename]({{ site.baseurl }}/commands/dropbox-team-group-rename.html)             | Rename the group             |
| [dropbox team group update type]({{ site.baseurl }}/commands/dropbox-team-group-update-type.html)   | Update group management type |

## Group member management commands

You can add/delete/update group members by the below commands. If you want to add/delete/update group members by CSV file, use `dropbox team group member batch add`, `dropbox team group member batch delete`, or `dropbox team group member batch delete`.

| Command                                                                                                           | Description                       |
|-------------------------------------------------------------------------------------------------------------------|-----------------------------------|
| [dropbox team group member add]({{ site.baseurl }}/commands/dropbox-team-group-member-add.html)                   | Add a member to the group         |
| [dropbox team group member delete]({{ site.baseurl }}/commands/dropbox-team-group-member-delete.html)             | Delete a member from the group    |
| [dropbox team group member list]({{ site.baseurl }}/commands/dropbox-team-group-member-list.html)                 | List members of groups            |
| [dropbox team group member batch add]({{ site.baseurl }}/commands/dropbox-team-group-member-batch-add.html)       | Bulk add members into groups      |
| [dropbox team group member batch delete]({{ site.baseurl }}/commands/dropbox-team-group-member-batch-delete.html) | Delete members from groups        |
| [dropbox team group member batch update]({{ site.baseurl }}/commands/dropbox-team-group-member-batch-update.html) | Add or delete members from groups |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `dropbox team group list`. The command `dropbox team group list` will report the number of members of each group. If it's zero, a group is not currently used to adding permission to folders.
If you want to see which folder uses each group, use the command `dropbox team group folder list`. `dropbox team group folder list` will report the group to folder mapping. The report `group_with_no_folders` will show groups with no folders.
You can safely remove groups once if you check both the number of members and folders. After confirmation, you can bulk delete groups by using the command `dropbox team group batch delete`.

| Command                                                                                             | Description                |
|-----------------------------------------------------------------------------------------------------|----------------------------|
| [dropbox team group list]({{ site.baseurl }}/commands/dropbox-team-group-list.html)                 | List group(s)              |
| [dropbox team group folder list]({{ site.baseurl }}/commands/dropbox-team-group-folder-list.html)   | List folders of each group |
| [dropbox team group batch delete]({{ site.baseurl }}/commands/dropbox-team-group-batch-delete.html) | Delete groups              |

# Team content commands

Admins' can handle team folders, shared folders or member's folder content thru Dropbox Business API. Please be careful to use those commands.
The namespace is the term in Dropbox API that is for manage folder permissions or settings. Folder types such as shared folders, team folders, or nested folder in a team folder, member's root folder or member's app folder are all managed as a namespace.
The namespace commands can handle all types of folders, including team folders and member's folder. But commands for specific folder types have more features or detailed information in the report.

## Team folder operation commands

You can create, archive or permanently delete team folders by using the below commands. Please consider using `dropbox team teamfolder batch` commands if you need to handle multiple team folders.

| Command                                                                                                                     | Description                            |
|-----------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| [dropbox team teamfolder add]({{ site.baseurl }}/commands/dropbox-team-teamfolder-add.html)                                 | Add team folder to the team            |
| [dropbox team teamfolder archive]({{ site.baseurl }}/commands/dropbox-team-teamfolder-archive.html)                         | Archive team folder                    |
| [dropbox team teamfolder batch archive]({{ site.baseurl }}/commands/dropbox-team-teamfolder-batch-archive.html)             | Archiving team folders                 |
| [dropbox team teamfolder batch permdelete]({{ site.baseurl }}/commands/dropbox-team-teamfolder-batch-permdelete.html)       | Permanently delete team folders        |
| [dropbox team teamfolder batch replication]({{ site.baseurl }}/commands/dropbox-team-teamfolder-batch-replication.html)     | Batch replication of team folders      |
| [dropbox team teamfolder file size]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-size.html)                     | Calculate size of team folders         |
| [dropbox team teamfolder list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-list.html)                               | List team folder(s)                    |
| [dropbox team teamfolder permdelete]({{ site.baseurl }}/commands/dropbox-team-teamfolder-permdelete.html)                   | Permanently delete team folder         |
| [dropbox team teamfolder policy list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-policy-list.html)                 | List policies of team folders          |
| [dropbox team teamfolder sync setting list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-sync-setting-list.html)     | List team folder sync settings         |
| [dropbox team teamfolder sync setting update]({{ site.baseurl }}/commands/dropbox-team-teamfolder-sync-setting-update.html) | Batch update team folder sync settings |

## Team folder permission commands

You can bulk add or delete members into team folders or sub-folders of a team folder through the below commands.

| Command                                                                                                         | Description                                   |
|-----------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| [dropbox team teamfolder member list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-member-list.html)     | List team folder members                      |
| [dropbox team teamfolder member add]({{ site.baseurl }}/commands/dropbox-team-teamfolder-member-add.html)       | Batch adding users/groups to team folders     |
| [dropbox team teamfolder member delete]({{ site.baseurl }}/commands/dropbox-team-teamfolder-member-delete.html) | Batch removing users/groups from team folders |

## Team folder & shared folder commands

The below commands are for both team folders and shared folders of the team.
If you wanted to know who are actually uses specific folders, please consider using the command `dropbox team content mount list`. Mount is a status a user add a shared folder to his/her Dropbox account.

| Command                                                                                               | Description                                                  |
|-------------------------------------------------------------------------------------------------------|--------------------------------------------------------------|
| [dropbox team content member list]({{ site.baseurl }}/commands/dropbox-team-content-member-list.html) | List team folder & shared folder members                     |
| [dropbox team content member size]({{ site.baseurl }}/commands/dropbox-team-content-member-size.html) | Count number of members of team folders and shared folders   |
| [dropbox team content mount list]({{ site.baseurl }}/commands/dropbox-team-content-mount-list.html)   | List all mounted/unmounted shared folders of team members.   |
| [dropbox team content policy list]({{ site.baseurl }}/commands/dropbox-team-content-policy-list.html) | List policies of team folders and shared folders in the team |

## Namespace commands

| Command                                                                                                   | Description                                                 |
|-----------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| [dropbox team namespace list]({{ site.baseurl }}/commands/dropbox-team-namespace-list.html)               | List all namespaces of the team                             |
| [dropbox team namespace summary]({{ site.baseurl }}/commands/dropbox-team-namespace-summary.html)         | Report team namespace status summary.                       |
| [dropbox team namespace file list]({{ site.baseurl }}/commands/dropbox-team-namespace-file-list.html)     | List all files and folders of the team namespaces           |
| [dropbox team namespace file size]({{ site.baseurl }}/commands/dropbox-team-namespace-file-size.html)     | List all files and folders of the team namespaces           |
| [dropbox team namespace member list]({{ site.baseurl }}/commands/dropbox-team-namespace-member-list.html) | List members of shared folders and team folders in the team |

## Team file request commands

| Command                                                                                         | Description                        |
|-------------------------------------------------------------------------------------------------|------------------------------------|
| [dropbox team filerequest list]({{ site.baseurl }}/commands/dropbox-team-filerequest-list.html) | List all file requests in the team |

## Member file commands

| Command                                                                                                     | Description                                                               |
|-------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [dropbox team member file permdelete]({{ site.baseurl }}/commands/dropbox-team-member-file-permdelete.html) | Permanently delete the file or folder at a given path of the team member. |

## Team insight

Team Insight is a feature of Dropbox Business that provides a view of team folder summary.

| Command                                                                                                                       | Description                            |
|-------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| [dropbox team insight scan]({{ site.baseurl }}/commands/dropbox-team-insight-scan.html)                                       | Scans team data for analysis           |
| [dropbox team insight scanretry]({{ site.baseurl }}/commands/dropbox-team-insight-scanretry.html)                             | Retry scan for errors on the last scan |
| [dropbox team insight summarize]({{ site.baseurl }}/commands/dropbox-team-insight-summarize.html)                             | Summarize team data for analysis       |
| [dropbox team insight report teamfoldermember]({{ site.baseurl }}/commands/dropbox-team-insight-report-teamfoldermember.html) | Report team folder members             |

# Team shared link commands

The team shared link commands are capable of listing all shared links in the team or update/delete specified shared links.

| Command                                                                                                                 | Description                                                   |
|-------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [dropbox team sharedlink list]({{ site.baseurl }}/commands/dropbox-team-sharedlink-list.html)                           | List of shared links                                          |
| [dropbox team sharedlink cap expiry]({{ site.baseurl }}/commands/dropbox-team-sharedlink-cap-expiry.html)               | Set expiry cap to shared links in the team                    |
| [dropbox team sharedlink cap visibility]({{ site.baseurl }}/commands/dropbox-team-sharedlink-cap-visibility.html)       | Set visibility cap to shared links in the team                |
| [dropbox team sharedlink update expiry]({{ site.baseurl }}/commands/dropbox-team-sharedlink-update-expiry.html)         | Update expiration date of public shared links within the team |
| [dropbox team sharedlink update password]({{ site.baseurl }}/commands/dropbox-team-sharedlink-update-password.html)     | Set or update shared link passwords                           |
| [dropbox team sharedlink update visibility]({{ site.baseurl }}/commands/dropbox-team-sharedlink-update-visibility.html) | Update visibility of shared links                             |
| [dropbox team sharedlink delete links]({{ site.baseurl }}/commands/dropbox-team-sharedlink-delete-links.html)           | Batch delete shared links                                     |
| [dropbox team sharedlink delete member]({{ site.baseurl }}/commands/dropbox-team-sharedlink-delete-member.html)         | Delete all shared links of the member                         |

## Difference between `dropbox team sharedlink cap` and `dropbox team sharedlink update`

Commands `dropbox team sharedlink update` is for setting a value to the shared links. Commands `dropbox team sharedlink cap` is for putting a cap value to the shared links.
For example: if you set expiry by `dropbox team sharedlink update expiry` with the expiration date 2021-05-06. The command will update the expiry to 2021-05-06 even if the existing link has a shorter expiration date like 2021-05-04.
On the other hand, `dropbox team sharedlink cap expiry` updates links when the link has a longer expiration date, like 2021-05-07.

Similarly, the command `dropbox team sharedlink cap visibility` will restrict visibility only when the link has less protected visibility. For example, if you want to ensure shared links without password to the team only. `dropbox team sharedlink cap visibility` will update visibility to the team only when a link is public and has no password.

## Example (list links):

List all public links in the team

```
tbx team sharedlink list -visibility public
```

Results are stored in CSV, xlsx, and JSON format. You can modify the report for updating shared links.
If you are familiar with the command jq, then they can create CSV file directly like below.

```
tbx team sharedlink list -output json | jq '.sharedlink.url' > all_links.csv
```

List links filtered by link owner email address:

```
tbx team sharedlink list -output json | jq 'select(.member.profile.email == "username@example.com") | .sharedlink.url'
```

## Example (delete links):

Delete all link that listed in the CSV file

```
tbx team sharedlink delete links -file /PATH/TO/DATA.csv
```

If you are familiar with jq command, then they can send data directly from the pipe like below (pass single dash `-` to the `-file` option to read from standard input).

```
tbx team sharedlink list -visibility public -output json | tbx team sharedlink delete links -file -
```

# File lock title

Dropbox Business file lock information

## File lock member title

| Command                                                                                                                 | Description                                    |
|-------------------------------------------------------------------------------------------------------------------------|------------------------------------------------|
| [dropbox team member file lock all release]({{ site.baseurl }}/commands/dropbox-team-member-file-lock-all-release.html) | Release all locks under the path of the member |
| [dropbox team member file lock list]({{ site.baseurl }}/commands/dropbox-team-member-file-lock-list.html)               | List locks of the member under the path        |
| [dropbox team member file lock release]({{ site.baseurl }}/commands/dropbox-team-member-file-lock-release.html)         | Release the lock of the path as the member     |

## File lock team folder title

| Command                                                                                                                         | Description                                         |
|---------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| [dropbox team teamfolder file list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-list.html)                         | List files in team folders                          |
| [dropbox team teamfolder file lock all release]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-lock-all-release.html) | Release all locks under the path of the team folder |
| [dropbox team teamfolder file lock list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-lock-list.html)               | List locks in the team folder                       |
| [dropbox team teamfolder file lock release]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-lock-release.html)         | Release lock of the path in the team folder         |

# Activities log commands

The team activities log commands can export activities log by a certain filter or perspective.

| Command                                                                                                 | Description                        |
|---------------------------------------------------------------------------------------------------------|------------------------------------|
| [dropbox team activity batch user]({{ site.baseurl }}/commands/dropbox-team-activity-batch-user.html)   | Scan activities for multiple users |
| [dropbox team activity daily event]({{ site.baseurl }}/commands/dropbox-team-activity-daily-event.html) | Report activities by day           |
| [dropbox team activity event]({{ site.baseurl }}/commands/dropbox-team-activity-event.html)             | Event log                          |
| [dropbox team activity user]({{ site.baseurl }}/commands/dropbox-team-activity-user.html)               | Activities log per user            |

# Connected applications and devices commands

The below commands can retrieve information about connected devices or applications in the team.

| Command                                                                                                 | Description                                                 |
|---------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| [dropbox team device list]({{ site.baseurl }}/commands/dropbox-team-device-list.html)                   | List all devices/sessions in the team                       |
| [dropbox team device unlink]({{ site.baseurl }}/commands/dropbox-team-device-unlink.html)               | Unlink device sessions                                      |
| [dropbox team linkedapp list]({{ site.baseurl }}/commands/dropbox-team-linkedapp-list.html)             | List linked applications                                    |
| [dropbox team backup device status]({{ site.baseurl }}/commands/dropbox-team-backup-device-status.html) | Dropbox Backup device status change in the specified period |

# Other usecases

## External ID

External ID is the attribute that is not shown in any user interface of Dropbox. This attribute is for keep a relationship between Dropbox and identity source (e.g. Active Directory, HR database) by identity management software such as Dropbox AD Connector. In case if you are using Dropbox AD Connector and you built a new Active Directory tree. You may need to clear existing external IDs to disconnect relationships with the old Active Directory tree and the new tree.
If you skip clear external IDs, Dropbox AD Connector may unintentionally delete accounts during configuring to the new tree.
If you want to see existing external IDs, use the `dropbox team member list` command. But the command will not include external ID by default. Please add the option `-experiment report_all_columns` like below.

```
tbx member list -experiment report_all_columns
```

| Command                                                                                                                     | Description                        |
|-----------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| [dropbox team member list]({{ site.baseurl }}/commands/dropbox-team-member-list.html)                                       | List team member(s)                |
| [dropbox team member clear externalid]({{ site.baseurl }}/commands/dropbox-team-member-clear-externalid.html)               | Clear external_id of members       |
| [dropbox team member update batch externalid]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-externalid.html) | Update External ID of team members |
| [dropbox team group list]({{ site.baseurl }}/commands/dropbox-team-group-list.html)                                         | List group(s)                      |
| [dropbox team group clear externalid]({{ site.baseurl }}/commands/dropbox-team-group-clear-externalid.html)                 | Clear an external ID of a group    |

## Data migration helper commands

Data migration helper commands copies member folders or team folders to another account or team. Please test before using those commands before actual data migration.

| Command                                                                                                                     | Description                                            |
|-----------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------|
| [dropbox team member folder replication]({{ site.baseurl }}/commands/dropbox-team-member-folder-replication.html)           | Replicate a folder to another member's personal folder |
| [dropbox team member replication]({{ site.baseurl }}/commands/dropbox-team-member-replication.html)                         | Replicate team member files                            |
| [dropbox team teamfolder partial replication]({{ site.baseurl }}/commands/dropbox-team-teamfolder-partial-replication.html) | Partial team folder replication to the other team      |
| [dropbox team teamfolder replication]({{ site.baseurl }}/commands/dropbox-team-teamfolder-replication.html)                 | Replicate a team folder to the other team              |

## Team info commands

| Command                                                                             | Description                         |
|-------------------------------------------------------------------------------------|-------------------------------------|
| [dropbox team feature]({{ site.baseurl }}/commands/dropbox-team-feature.html)       | Team feature                        |
| [dropbox team filesystem]({{ site.baseurl }}/commands/dropbox-team-filesystem.html) | Identify team's file system version |
| [dropbox team info]({{ site.baseurl }}/commands/dropbox-team-info.html)             | Team information                    |

# Paper commands

## Legacy paper commands

Commands for a team's legacy Paper content. Please see [official guide](https://developers.dropbox.com/paper-migration-guide) more detail about legacy Paper and migration

| Command                                                                                                             | Description                                               |
|---------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| [dropbox team content legacypaper count]({{ site.baseurl }}/commands/dropbox-team-content-legacypaper-count.html)   | Count number of Paper documents per member                |
| [dropbox team content legacypaper list]({{ site.baseurl }}/commands/dropbox-team-content-legacypaper-list.html)     | List team member Paper documents                          |
| [dropbox team content legacypaper export]({{ site.baseurl }}/commands/dropbox-team-content-legacypaper-export.html) | Export entire team member Paper documents into local path |

# Team admin commands

Below commands are for managing team admins.

| Command                                                                                                       | Description                                                               |
|---------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [dropbox team admin list]({{ site.baseurl }}/commands/dropbox-team-admin-list.html)                           | List admin roles of members                                               |
| [dropbox team admin role add]({{ site.baseurl }}/commands/dropbox-team-admin-role-add.html)                   | Add a new role to the member                                              |
| [dropbox team admin role clear]({{ site.baseurl }}/commands/dropbox-team-admin-role-clear.html)               | Remove all admin roles from the member                                    |
| [dropbox team admin role delete]({{ site.baseurl }}/commands/dropbox-team-admin-role-delete.html)             | Remove a role from the member                                             |
| [dropbox team admin role list]({{ site.baseurl }}/commands/dropbox-team-admin-role-list.html)                 | List admin roles of the team                                              |
| [dropbox team admin group role add]({{ site.baseurl }}/commands/dropbox-team-admin-group-role-add.html)       | Add the role to members of the group                                      |
| [dropbox team admin group role delete]({{ site.baseurl }}/commands/dropbox-team-admin-group-role-delete.html) | Delete the role from all members except of members of the exception group |

# Commands that run as a team member

You can run a command as a team member. For example, you can upload a file into member's folder by using `dropbox team runas file sync batch up`.

| Command                                                                                                                                     | Description                                          |
|---------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------|
| [dropbox team runas file list]({{ site.baseurl }}/commands/dropbox-team-runas-file-list.html)                                               | List files and folders run as a member               |
| [dropbox team runas file batch copy]({{ site.baseurl }}/commands/dropbox-team-runas-file-batch-copy.html)                                   | Batch copy files/folders as a member                 |
| [dropbox team runas file sync batch up]({{ site.baseurl }}/commands/dropbox-team-runas-file-sync-batch-up.html)                             | Batch upstream sync with Dropbox                     |
| [dropbox team runas sharedfolder list]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-list.html)                               | List shared folders                                  |
| [dropbox team runas sharedfolder isolate]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-isolate.html)                         | Isolate member from shared folder                    |
| [dropbox team runas sharedfolder mount add]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-add.html)                     | Mount a shared folder as another member              |
| [dropbox team runas sharedfolder mount delete]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-delete.html)               | The specified user unmounts the designated folder.   |
| [dropbox team runas sharedfolder mount list]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-list.html)                   | List all shared folders the specified member mounted |
| [dropbox team runas sharedfolder mount mountable]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-mountable.html)         | List all shared folders the member can mount         |
| [dropbox team runas sharedfolder batch leave]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-batch-leave.html)                 | Leave shared folders in batch                        |
| [dropbox team runas sharedfolder batch share]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-batch-share.html)                 | Share shared folders in batch                        |
| [dropbox team runas sharedfolder batch unshare]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-batch-unshare.html)             | Unshare shared folders in batch                      |
| [dropbox team runas sharedfolder member batch add]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-member-batch-add.html)       | Add members to shared folders in batch               |
| [dropbox team runas sharedfolder member batch delete]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-member-batch-delete.html) | Remove members from shared folders in batch          |

# Legal hold

With legal holds, admins can place a legal hold on members of their team and view and export all the content that's been created or modified by those members.

| Command                                                                                                                   | Description                                 |
|---------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| [dropbox team legalhold add]({{ site.baseurl }}/commands/dropbox-team-legalhold-add.html)                                 | Creates new legal hold policy.              |
| [dropbox team legalhold list]({{ site.baseurl }}/commands/dropbox-team-legalhold-list.html)                               | Retrieve existing policies                  |
| [dropbox team legalhold member batch update]({{ site.baseurl }}/commands/dropbox-team-legalhold-member-batch-update.html) | Update member list of legal hold policy     |
| [dropbox team legalhold member list]({{ site.baseurl }}/commands/dropbox-team-legalhold-member-list.html)                 | List members of the legal hold              |
| [dropbox team legalhold release]({{ site.baseurl }}/commands/dropbox-team-legalhold-release.html)                         | Releases a legal hold by Id                 |
| [dropbox team legalhold revision list]({{ site.baseurl }}/commands/dropbox-team-legalhold-revision-list.html)             | List revisions under legal hold             |
| [dropbox team legalhold update desc]({{ site.baseurl }}/commands/dropbox-team-legalhold-update-desc.html)                 | Update description of the legal hold policy |
| [dropbox team legalhold update name]({{ site.baseurl }}/commands/dropbox-team-legalhold-update-name.html)                 | Update name of the legal hold policy        |

# Notes:

Dropbox Business footnote information


