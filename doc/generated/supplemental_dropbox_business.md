# Member management commands

## Information commands

Below commands are to retrieve information about team members.

| Command                                     | Description                    |
|---------------------------------------------|--------------------------------|
| [member list](member-list.md)               | List team member(s)            |
| [member folder list](member-folder-list.md) | Find folders for each member   |
| [member quota list](member-quota-list.md)   | List team member quota         |
| [member quota usage](member-quota-usage.md) | List team member storage usage |
| [team activity user](team-activity-user.md) | Activities log per user        |

## Basic management commands

Below commands are for managing team member accounts. Those commands are for a bulk operation by a CSV file.

| Command                                               | Description                                          |
|-------------------------------------------------------|------------------------------------------------------|
| [member invite](member-invite.md)                     | Invite member(s)                                     |
| [member delete](member-delete.md)                     | Delete members                                       |
| [member detach](member-detach.md)                     | Convert Dropbox Business accounts to a Basic account |
| [member reinvite](member-reinvite.md)                 | Reinvite invited status members to the team          |
| [member update email](member-update-email.md)         | Member email operation                               |
| [member update profile](member-update-profile.md)     | Update member profile                                |
| [member update visible](member-update-visible.md)     | Disable directory restriction to members             |
| [member update invisible](member-update-invisible.md) | Enable directory restriction to members              |
| [member quota update](member-quota-update.md)         | Update team member quota                             |

## Member profile setting commands

Member profile commands are for bulk updating member profile information.
If you need to update the members' email addresses, use the 'member update email` command. The command 'member update email` receives a CSV file to bulk update email addresses.
If you need to update the member's display name, use the 'member update profile` command.

| Command                                           | Description            |
|---------------------------------------------------|------------------------|
| [member update email](member-update-email.md)     | Member email operation |
| [member update profile](member-update-profile.md) | Update member profile  |

## Member storage quota control commands

You can see existing member storage quota setting or usage by the `member quota list` and `member quota usage` command. If you need to update member quota, use the `member quota update` command. The command `member quota update` receives CSV input for bulk updating storage quota setting.

| Command                                       | Description                    |
|-----------------------------------------------|--------------------------------|
| [member quota list](member-quota-list.md)     | List team member quota         |
| [member quota usage](member-quota-usage.md)   | List team member storage usage |
| [member quota update](member-quota-update.md) | Update team member quota       |

## Directory restriction commands

Directory restriction is the Dropbox Business feature to hide a member from others. Below commands update this setting to hide or unhide members from others.

| Command                                               | Description                              |
|-------------------------------------------------------|------------------------------------------|
| [member update visible](member-update-visible.md)     | Disable directory restriction to members |
| [member update invisible](member-update-invisible.md) | Enable directory restriction to members  |

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                     | Description      |
|---------------------------------------------|------------------|
| [group add](group-add.md)                   | Create new group |
| [group delete](group-delete.md)             | Delete group     |
| [group batch delete](group-batch-delete.md) | Delete groups    |
| [group list](group-list.md)                 | List group(s)    |
| [group rename](group-rename.md)             | Rename the group |

## Group member management commands

You can add/delete/update group members by the below commands. If you want to add/delete/update group members by CSV file, use `group member batch add`, `group member batch delete`, or `group member batch delete`.

| Command                                                   | Description                       |
|-----------------------------------------------------------|-----------------------------------|
| [group member add](group-member-add.md)                   | Add a member to the group         |
| [group member delete](group-member-delete.md)             | Delete a member from the group    |
| [group member list](group-member-list.md)                 | List members of groups            |
| [group member batch add](group-member-batch-add.md)       | Bulk add members into groups      |
| [group member batch delete](group-member-batch-delete.md) | Delete members from groups        |
| [group member batch update](group-member-batch-update.md) | Add or delete members from groups |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `group list`. The command `group list` will report the number of members of each group. If it's zero, a group is not currently used to adding permission to folders.
If you want to see which folder uses each group, use the command `group folder list`. `group folder list` will report the group to folder mapping. The report `group_with_no_folders` will show groups with no folders.
You can safely remove groups once if you check both the number of members and folders. After confirmation, you can bulk delete groups by using the command `group batch delete`.

| Command                                     | Description                |
|---------------------------------------------|----------------------------|
| [group list](group-list.md)                 | List group(s)              |
| [group folder list](group-folder-list.md)   | Find folders of each group |
| [group batch delete](group-batch-delete.md) | Delete groups              |

# Team content commands

Admins' can handle team folders, shared folders or member's folder content thru Dropbox Business API. Please be careful to use those commands.
The namespace is the term in Dropbox API that is for manage folder permissions or settings. Folder types such as shared folders, team folders, or nested folder in a team folder, member's root folder or member's app folder are all managed as a namespace.
The namespace commands can handle all types of folders, including team folders and member's folder. But commands for specific folder types have more features or detailed information in the report.

## Team folder operation commands

You can create, archive or permanently delete team folders by using the below commands. Please consider using `teamfolder batch` commands if you need to handle multiple team folders.

| Command                                                         | Description                       |
|-----------------------------------------------------------------|-----------------------------------|
| [teamfolder list](teamfolder-list.md)                           | List team folder(s)               |
| [teamfolder policy list](teamfolder-policy-list.md)             | List policies of team folders     |
| [teamfolder file size](teamfolder-file-size.md)                 | Calculate size of team folders    |
| [teamfolder add](teamfolder-add.md)                             | Add team folder to the team       |
| [teamfolder archive](teamfolder-archive.md)                     | Archive team folder               |
| [teamfolder permdelete](teamfolder-permdelete.md)               | Permanently delete team folder    |
| [teamfolder batch archive](teamfolder-batch-archive.md)         | Archiving team folders            |
| [teamfolder batch permdelete](teamfolder-batch-permdelete.md)   | Permanently delete team folders   |
| [teamfolder batch replication](teamfolder-batch-replication.md) | Batch replication of team folders |

## Team folder permission commands

You can bulk add or delete members into team folders or sub-folders of a team folder through the below commands.

| Command                                                 | Description                                   |
|---------------------------------------------------------|-----------------------------------------------|
| [teamfolder member list](teamfolder-member-list.md)     | List team folder members                      |
| [teamfolder member add](teamfolder-member-add.md)       | Batch adding users/groups to team folders     |
| [teamfolder member delete](teamfolder-member-delete.md) | Batch removing users/groups from team folders |

## Team folder & shared folder commands

The below commands are for both team folders and shared folders of the team.
If you wanted to know who are actually uses specific folders, please consider using the command `team content mount list`. Mount is a status a user add a shared folder to his/her Dropbox account.

| Command                                                 | Description                                                  |
|---------------------------------------------------------|--------------------------------------------------------------|
| [team content member list](team-content-member-list.md) | List team folder & shared folder members                     |
| [team content mount list](team-content-mount-list.md)   | List all mounted/unmounted shared folders of team members.   |
| [team content policy list](team-content-policy-list.md) | List policies of team folders and shared folders in the team |

## Namespace commands

| Command                                                     | Description                                                 |
|-------------------------------------------------------------|-------------------------------------------------------------|
| [team namespace list](team-namespace-list.md)               | List all namespaces of the team                             |
| [team namespace file list](team-namespace-file-list.md)     | List all files and folders of the team namespaces           |
| [team namespace file size](team-namespace-file-size.md)     | List all files and folders of the team namespaces           |
| [team namespace member list](team-namespace-member-list.md) | List members of shared folders and team folders in the team |

## Team file request commands

| Command                                           | Description                        |
|---------------------------------------------------|------------------------------------|
| [team filerequest list](team-filerequest-list.md) | List all file requests in the team |

## Member file commands

| Command                                             | Description                                                               |
|-----------------------------------------------------|---------------------------------------------------------------------------|
| [member file permdelete](member-file-permdelete.md) | Permanently delete the file or folder at a given path of the team member. |

# Team shared link commands

The team shared link commands are capable of listing all shared links in the team or update/delete specified shared links.

| Command                                                                   | Description                                                   |
|---------------------------------------------------------------------------|---------------------------------------------------------------|
| [team sharedlink list](team-sharedlink-list.md)                           | List of shared links                                          |
| [team sharedlink update expiry](team-sharedlink-update-expiry.md)         | Update expiration date of public shared links within the team |
| [team sharedlink update password](team-sharedlink-update-password.md)     | Set or update shared link passwords                           |
| [team sharedlink update visibility](team-sharedlink-update-visibility.md) | Update visibility of shared links                             |
| [team sharedlink delete links](team-sharedlink-delete-links.md)           | Batch delete shared links                                     |
| [team sharedlink delete member](team-sharedlink-delete-member.md)         | Delete all shared links of the member                         |

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

| Command                                                         | Description                                    |
|-----------------------------------------------------------------|------------------------------------------------|
| [member file lock all release](member-file-lock-all-release.md) | Release all locks under the path of the member |
| [member file lock list](member-file-lock-list.md)               | List locks of the member under the path        |
| [member file lock release](member-file-lock-release.md)         | Release the lock of the path as the member     |

## File lock commands for team folders

| Command                                                                 | Description                                         |
|-------------------------------------------------------------------------|-----------------------------------------------------|
| [teamfolder file list](teamfolder-file-list.md)                         | List files in team folders                          |
| [teamfolder file lock all release](teamfolder-file-lock-all-release.md) | Release all locks under the path of the team folder |
| [teamfolder file lock list](teamfolder-file-lock-list.md)               | List locks in the team folder                       |
| [teamfolder file lock release](teamfolder-file-lock-release.md)         | Release lock of the path in the team folder         |

# Activities log commands

The team activities log commands can export activities log by a certain filter or perspective.

| Command                                                   | Description                        |
|-----------------------------------------------------------|------------------------------------|
| [team activity batch user](team-activity-batch-user.md)   | Scan activities for multiple users |
| [team activity daily event](team-activity-daily-event.md) | Report activities by day           |
| [team activity event](team-activity-event.md)             | Event log                          |
| [team activity user](team-activity-user.md)               | Activities log per user            |

# Connected applications and devices commands

The below commands can retrieve information about connected devices or applications in the team.

| Command                                       | Description                           |
|-----------------------------------------------|---------------------------------------|
| [team device list](team-device-list.md)       | List all devices/sessions in the team |
| [team device unlink](team-device-unlink.md)   | Unlink device sessions                |
| [team linkedapp list](team-linkedapp-list.md) | List linked applications              |

# Other usecases

## External ID

External ID is the attribute that is not shown in any user interface of Dropbox. This attribute is for keep a relationship between Dropbox and identity source (e.g. Active Directory, HR database) by identity management software such as Dropbox AD Connector. In case if you are using Dropbox AD Connector and you built a new Active Directory tree. You may need to clear existing external IDs to disconnect relationships with the old Active Directory tree and the new tree.
If you skip clear external IDs, Dropbox AD Connector may unintentionally delete accounts during configuring to the new tree.
If you want to see existing external IDs, use the `member list` command. But the command will not include external ID by default. Please consider using [jq](https://stedolan.github.io/jq/) command and run like below.

```
tbx member list -output json | jq -r '[.profile.email, .profile.external_id] | @csv'
```

| Command                                                 | Description                        |
|---------------------------------------------------------|------------------------------------|
| [member list](member-list.md)                           | List team member(s)                |
| [member clear externalid](member-clear-externalid.md)   | Clear external_id of members       |
| [member update externalid](member-update-externalid.md) | Update External ID of team members |

## Data migration helper commands

Data migration helper commands copies member folders or team folders to another account or team. Please test before using those commands before actual data migration.

| Command                                                             | Description                                            |
|---------------------------------------------------------------------|--------------------------------------------------------|
| [member folder replication](member-folder-replication.md)           | Replicate a folder to another member's personal folder |
| [member replication](member-replication.md)                         | Replicate team member files                            |
| [teamfolder partial replication](teamfolder-partial-replication.md) | Partial team folder replication to the other team      |
| [teamfolder replication](teamfolder-replication.md)                 | Replicate a team folder to the other team              |

## Team info commands

| Command                         | Description      |
|---------------------------------|------------------|
| [team feature](team-feature.md) | Team feature     |
| [team info](team-info.md)       | Team information |

# Notes:

Dropbox Business commands require admin permissions to execute them. Auth tokens must not share with anyone, including Dropbox support.

