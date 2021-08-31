---
layout: page
title: Dropbox Business commands
lang: en
---

# Member management commands

## Information commands

Below commands are to retrieve information about team members.

| Command                                                                   | Description                    |
|---------------------------------------------------------------------------|--------------------------------|
| [member list]({{ site.baseurl }}/commands/member-list.html)               | List team member(s)            |
| [member feature]({{ site.baseurl }}/commands/member-feature.html)         | List member feature settings   |
| [member folder list]({{ site.baseurl }}/commands/member-folder-list.html) | Find folders for each member   |
| [member quota list]({{ site.baseurl }}/commands/member-quota-list.html)   | List team member quota         |
| [member quota usage]({{ site.baseurl }}/commands/member-quota-usage.html) | List team member storage usage |
| [team activity user]({{ site.baseurl }}/commands/team-activity-user.html) | Activities log per user        |

## Basic management commands

Below commands are for managing team member accounts. Those commands are for a bulk operation by a CSV file.

| Command                                                                             | Description                                          |
|-------------------------------------------------------------------------------------|------------------------------------------------------|
| [member invite]({{ site.baseurl }}/commands/member-invite.html)                     | Invite member(s)                                     |
| [member delete]({{ site.baseurl }}/commands/member-delete.html)                     | Delete members                                       |
| [member detach]({{ site.baseurl }}/commands/member-detach.html)                     | Convert Dropbox Business accounts to a Basic account |
| [member reinvite]({{ site.baseurl }}/commands/member-reinvite.html)                 | Reinvite invited status members to the team          |
| [member update email]({{ site.baseurl }}/commands/member-update-email.html)         | Member email operation                               |
| [member update profile]({{ site.baseurl }}/commands/member-update-profile.html)     | Update member profile                                |
| [member update visible]({{ site.baseurl }}/commands/member-update-visible.html)     | Disable directory restriction to members             |
| [member update invisible]({{ site.baseurl }}/commands/member-update-invisible.html) | Enable directory restriction to members              |
| [member quota update]({{ site.baseurl }}/commands/member-quota-update.html)         | Update team member quota                             |

## Member profile setting commands

Member profile commands are for bulk updating member profile information.
If you need to update the members' email addresses, use the 'member update email` command. The command 'member update email` receives a CSV file to bulk update email addresses.
If you need to update the member's display name, use the 'member update profile` command.

| Command                                                                         | Description            |
|---------------------------------------------------------------------------------|------------------------|
| [member update email]({{ site.baseurl }}/commands/member-update-email.html)     | Member email operation |
| [member update profile]({{ site.baseurl }}/commands/member-update-profile.html) | Update member profile  |

## Member storage quota control commands

You can see existing member storage quota setting or usage by the `member quota list` and `member quota usage` command. If you need to update member quota, use the `member quota update` command. The command `member quota update` receives CSV input for bulk updating storage quota setting.

| Command                                                                     | Description                    |
|-----------------------------------------------------------------------------|--------------------------------|
| [member quota list]({{ site.baseurl }}/commands/member-quota-list.html)     | List team member quota         |
| [member quota usage]({{ site.baseurl }}/commands/member-quota-usage.html)   | List team member storage usage |
| [member quota update]({{ site.baseurl }}/commands/member-quota-update.html) | Update team member quota       |

## There are two types of commands available for suspending/unsuspending members. If you wanted to suspend/unsuspend a member one by one, please use `member suspend` or `member unsuspend`. Otherwise, if you want to suspend/unsuspend members through a CSV file, please use the `member batch suspend` or `member batch unsuspend` command.

Suspend/unsuspend a member

| Command                                                                           | Description            |
|-----------------------------------------------------------------------------------|------------------------|
| [member suspend]({{ site.baseurl }}/commands/member-suspend.html)                 | Suspend a member       |
| [member unsuspend]({{ site.baseurl }}/commands/member-unsuspend.html)             | Unsuspend a member     |
| [member batch suspend]({{ site.baseurl }}/commands/member-batch-suspend.html)     | Bulk suspend members   |
| [member batch unsuspend]({{ site.baseurl }}/commands/member-batch-unsuspend.html) | Bulk unsuspend members |

## Directory restriction commands

Directory restriction is the Dropbox Business feature to hide a member from others. Below commands update this setting to hide or unhide members from others.

| Command                                                                             | Description                              |
|-------------------------------------------------------------------------------------|------------------------------------------|
| [member update visible]({{ site.baseurl }}/commands/member-update-visible.html)     | Disable directory restriction to members |
| [member update invisible]({{ site.baseurl }}/commands/member-update-invisible.html) | Enable directory restriction to members  |

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                                                   | Description      |
|---------------------------------------------------------------------------|------------------|
| [group add]({{ site.baseurl }}/commands/group-add.html)                   | Create new group |
| [group delete]({{ site.baseurl }}/commands/group-delete.html)             | Delete group     |
| [group batch delete]({{ site.baseurl }}/commands/group-batch-delete.html) | Delete groups    |
| [group list]({{ site.baseurl }}/commands/group-list.html)                 | List group(s)    |
| [group rename]({{ site.baseurl }}/commands/group-rename.html)             | Rename the group |

## Group member management commands

You can add/delete/update group members by the below commands. If you want to add/delete/update group members by CSV file, use `group member batch add`, `group member batch delete`, or `group member batch delete`.

| Command                                                                                 | Description                       |
|-----------------------------------------------------------------------------------------|-----------------------------------|
| [group member add]({{ site.baseurl }}/commands/group-member-add.html)                   | Add a member to the group         |
| [group member delete]({{ site.baseurl }}/commands/group-member-delete.html)             | Delete a member from the group    |
| [group member list]({{ site.baseurl }}/commands/group-member-list.html)                 | List members of groups            |
| [group member batch add]({{ site.baseurl }}/commands/group-member-batch-add.html)       | Bulk add members into groups      |
| [group member batch delete]({{ site.baseurl }}/commands/group-member-batch-delete.html) | Delete members from groups        |
| [group member batch update]({{ site.baseurl }}/commands/group-member-batch-update.html) | Add or delete members from groups |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `group list`. The command `group list` will report the number of members of each group. If it's zero, a group is not currently used to adding permission to folders.
If you want to see which folder uses each group, use the command `group folder list`. `group folder list` will report the group to folder mapping. The report `group_with_no_folders` will show groups with no folders.
You can safely remove groups once if you check both the number of members and folders. After confirmation, you can bulk delete groups by using the command `group batch delete`.

| Command                                                                   | Description                |
|---------------------------------------------------------------------------|----------------------------|
| [group list]({{ site.baseurl }}/commands/group-list.html)                 | List group(s)              |
| [group folder list]({{ site.baseurl }}/commands/group-folder-list.html)   | Find folders of each group |
| [group batch delete]({{ site.baseurl }}/commands/group-batch-delete.html) | Delete groups              |

# Team content commands

Admins' can handle team folders, shared folders or member's folder content thru Dropbox Business API. Please be careful to use those commands.
The namespace is the term in Dropbox API that is for manage folder permissions or settings. Folder types such as shared folders, team folders, or nested folder in a team folder, member's root folder or member's app folder are all managed as a namespace.
The namespace commands can handle all types of folders, including team folders and member's folder. But commands for specific folder types have more features or detailed information in the report.

## Team folder operation commands

You can create, archive or permanently delete team folders by using the below commands. Please consider using `teamfolder batch` commands if you need to handle multiple team folders.

| Command                                                                                       | Description                       |
|-----------------------------------------------------------------------------------------------|-----------------------------------|
| [teamfolder list]({{ site.baseurl }}/commands/teamfolder-list.html)                           | List team folder(s)               |
| [teamfolder policy list]({{ site.baseurl }}/commands/teamfolder-policy-list.html)             | List policies of team folders     |
| [teamfolder file size]({{ site.baseurl }}/commands/teamfolder-file-size.html)                 | Calculate size of team folders    |
| [teamfolder add]({{ site.baseurl }}/commands/teamfolder-add.html)                             | Add team folder to the team       |
| [teamfolder archive]({{ site.baseurl }}/commands/teamfolder-archive.html)                     | Archive team folder               |
| [teamfolder permdelete]({{ site.baseurl }}/commands/teamfolder-permdelete.html)               | Permanently delete team folder    |
| [teamfolder batch archive]({{ site.baseurl }}/commands/teamfolder-batch-archive.html)         | Archiving team folders            |
| [teamfolder batch permdelete]({{ site.baseurl }}/commands/teamfolder-batch-permdelete.html)   | Permanently delete team folders   |
| [teamfolder batch replication]({{ site.baseurl }}/commands/teamfolder-batch-replication.html) | Batch replication of team folders |

## Team folder permission commands

You can bulk add or delete members into team folders or sub-folders of a team folder through the below commands.

| Command                                                                               | Description                                   |
|---------------------------------------------------------------------------------------|-----------------------------------------------|
| [teamfolder member list]({{ site.baseurl }}/commands/teamfolder-member-list.html)     | List team folder members                      |
| [teamfolder member add]({{ site.baseurl }}/commands/teamfolder-member-add.html)       | Batch adding users/groups to team folders     |
| [teamfolder member delete]({{ site.baseurl }}/commands/teamfolder-member-delete.html) | Batch removing users/groups from team folders |

## Team folder & shared folder commands

The below commands are for both team folders and shared folders of the team.
If you wanted to know who are actually uses specific folders, please consider using the command `team content mount list`. Mount is a status a user add a shared folder to his/her Dropbox account.

| Command                                                                               | Description                                                  |
|---------------------------------------------------------------------------------------|--------------------------------------------------------------|
| [team content member list]({{ site.baseurl }}/commands/team-content-member-list.html) | List team folder & shared folder members                     |
| [team content member size]({{ site.baseurl }}/commands/team-content-member-size.html) | Count number of members of team folders and shared folders   |
| [team content mount list]({{ site.baseurl }}/commands/team-content-mount-list.html)   | List all mounted/unmounted shared folders of team members.   |
| [team content policy list]({{ site.baseurl }}/commands/team-content-policy-list.html) | List policies of team folders and shared folders in the team |

## Namespace commands

| Command                                                                                   | Description                                                 |
|-------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| [team namespace list]({{ site.baseurl }}/commands/team-namespace-list.html)               | List all namespaces of the team                             |
| [team namespace file list]({{ site.baseurl }}/commands/team-namespace-file-list.html)     | List all files and folders of the team namespaces           |
| [team namespace file size]({{ site.baseurl }}/commands/team-namespace-file-size.html)     | List all files and folders of the team namespaces           |
| [team namespace member list]({{ site.baseurl }}/commands/team-namespace-member-list.html) | List members of shared folders and team folders in the team |

## Team file request commands

| Command                                                                         | Description                        |
|---------------------------------------------------------------------------------|------------------------------------|
| [team filerequest list]({{ site.baseurl }}/commands/team-filerequest-list.html) | List all file requests in the team |

## Member file commands

| Command                                                                           | Description                                                               |
|-----------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [member file permdelete]({{ site.baseurl }}/commands/member-file-permdelete.html) | Permanently delete the file or folder at a given path of the team member. |

# Team shared link commands

The team shared link commands are capable of listing all shared links in the team or update/delete specified shared links.

| Command                                                                                                 | Description                                                   |
|---------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [team sharedlink list]({{ site.baseurl }}/commands/team-sharedlink-list.html)                           | List of shared links                                          |
| [team sharedlink cap expiry]({{ site.baseurl }}/commands/team-sharedlink-cap-expiry.html)               | Set expiry cap to shared links in the team                    |
| [team sharedlink cap visibility]({{ site.baseurl }}/commands/team-sharedlink-cap-visibility.html)       | Set visibility cap to shared links in the team                |
| [team sharedlink update expiry]({{ site.baseurl }}/commands/team-sharedlink-update-expiry.html)         | Update expiration date of public shared links within the team |
| [team sharedlink update password]({{ site.baseurl }}/commands/team-sharedlink-update-password.html)     | Set or update shared link passwords                           |
| [team sharedlink update visibility]({{ site.baseurl }}/commands/team-sharedlink-update-visibility.html) | Update visibility of shared links                             |
| [team sharedlink delete links]({{ site.baseurl }}/commands/team-sharedlink-delete-links.html)           | Batch delete shared links                                     |
| [team sharedlink delete member]({{ site.baseurl }}/commands/team-sharedlink-delete-member.html)         | Delete all shared links of the member                         |

## Difference between `team sharedlink cap` and `team sharedlink update`

Commands `team sharedlink update` is for setting a value to the shared links. Commands `team sharedlink cap` is for putting a cap value to the shared links.
For example: if you set expiry by `team sharedlink update expiry` with the expiration date 2021-05-06. The command will update the expiry to 2021-05-06 even if the existing link has a shorter expiration date like 2021-05-04.
On the other hand, `team sharedlink cap expiry` updates links when the link has a longer expiration date, like 2021-05-07.

Similarly, the command `team sharedlink cap visibility` will restrict visibility only when the link has less protected visibility. For example, if you want to ensure shared links without password to the team only. `team sharedlink cap visibility` will update visibility to the team only when a link is public and has no password.

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

# File lock

File lock commands are capable of listing current file locks or releasing file locks as admin.

## File lock commands for members

| Command                                                                                       | Description                                    |
|-----------------------------------------------------------------------------------------------|------------------------------------------------|
| [member file lock all release]({{ site.baseurl }}/commands/member-file-lock-all-release.html) | Release all locks under the path of the member |
| [member file lock list]({{ site.baseurl }}/commands/member-file-lock-list.html)               | List locks of the member under the path        |
| [member file lock release]({{ site.baseurl }}/commands/member-file-lock-release.html)         | Release the lock of the path as the member     |

## File lock commands for team folders

| Command                                                                                               | Description                                         |
|-------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| [teamfolder file list]({{ site.baseurl }}/commands/teamfolder-file-list.html)                         | List files in team folders                          |
| [teamfolder file lock all release]({{ site.baseurl }}/commands/teamfolder-file-lock-all-release.html) | Release all locks under the path of the team folder |
| [teamfolder file lock list]({{ site.baseurl }}/commands/teamfolder-file-lock-list.html)               | List locks in the team folder                       |
| [teamfolder file lock release]({{ site.baseurl }}/commands/teamfolder-file-lock-release.html)         | Release lock of the path in the team folder         |

# Activities log commands

The team activities log commands can export activities log by a certain filter or perspective.

| Command                                                                                 | Description                        |
|-----------------------------------------------------------------------------------------|------------------------------------|
| [team activity batch user]({{ site.baseurl }}/commands/team-activity-batch-user.html)   | Scan activities for multiple users |
| [team activity daily event]({{ site.baseurl }}/commands/team-activity-daily-event.html) | Report activities by day           |
| [team activity event]({{ site.baseurl }}/commands/team-activity-event.html)             | Event log                          |
| [team activity user]({{ site.baseurl }}/commands/team-activity-user.html)               | Activities log per user            |

# Connected applications and devices commands

The below commands can retrieve information about connected devices or applications in the team.

| Command                                                                     | Description                           |
|-----------------------------------------------------------------------------|---------------------------------------|
| [team device list]({{ site.baseurl }}/commands/team-device-list.html)       | List all devices/sessions in the team |
| [team device unlink]({{ site.baseurl }}/commands/team-device-unlink.html)   | Unlink device sessions                |
| [team linkedapp list]({{ site.baseurl }}/commands/team-linkedapp-list.html) | List linked applications              |

# Other usecases

## External ID

External ID is the attribute that is not shown in any user interface of Dropbox. This attribute is for keep a relationship between Dropbox and identity source (e.g. Active Directory, HR database) by identity management software such as Dropbox AD Connector. In case if you are using Dropbox AD Connector and you built a new Active Directory tree. You may need to clear existing external IDs to disconnect relationships with the old Active Directory tree and the new tree.
If you skip clear external IDs, Dropbox AD Connector may unintentionally delete accounts during configuring to the new tree.
If you want to see existing external IDs, use the `member list` command. But the command will not include external ID by default. Please consider using [jq](https://stedolan.github.io/jq/) command and run like below.

```
tbx member list -output json | jq -r '[.profile.email, .profile.external_id] | @csv'
```

| Command                                                                               | Description                        |
|---------------------------------------------------------------------------------------|------------------------------------|
| [member list]({{ site.baseurl }}/commands/member-list.html)                           | List team member(s)                |
| [member clear externalid]({{ site.baseurl }}/commands/member-clear-externalid.html)   | Clear external_id of members       |
| [member update externalid]({{ site.baseurl }}/commands/member-update-externalid.html) | Update External ID of team members |

## Data migration helper commands

Data migration helper commands copies member folders or team folders to another account or team. Please test before using those commands before actual data migration.

| Command                                                                                           | Description                                            |
|---------------------------------------------------------------------------------------------------|--------------------------------------------------------|
| [member folder replication]({{ site.baseurl }}/commands/member-folder-replication.html)           | Replicate a folder to another member's personal folder |
| [member replication]({{ site.baseurl }}/commands/member-replication.html)                         | Replicate team member files                            |
| [teamfolder partial replication]({{ site.baseurl }}/commands/teamfolder-partial-replication.html) | Partial team folder replication to the other team      |
| [teamfolder replication]({{ site.baseurl }}/commands/teamfolder-replication.html)                 | Replicate a team folder to the other team              |

## Team info commands

| Command                                                       | Description      |
|---------------------------------------------------------------|------------------|
| [team feature]({{ site.baseurl }}/commands/team-feature.html) | Team feature     |
| [team info]({{ site.baseurl }}/commands/team-info.html)       | Team information |

# Paper commands

## Legacy paper commands

Commands for a team's legacy Paper content. Please see [official guide](https://developers.dropbox.com/paper-migration-guide) more detail about legacy Paper and migration

| Command                                                                                             | Description                                               |
|-----------------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| [team content legacypaper count]({{ site.baseurl }}/commands/team-content-legacypaper-count.html)   | Count number of Paper documents per member                |
| [team content legacypaper list]({{ site.baseurl }}/commands/team-content-legacypaper-list.html)     | List team member Paper documents                          |
| [team content legacypaper export]({{ site.baseurl }}/commands/team-content-legacypaper-export.html) | Export entire team member Paper documents into local path |

# Notes:

Dropbox Business commands require admin permissions to execute them. Auth tokens must not share with anyone, including Dropbox support.


