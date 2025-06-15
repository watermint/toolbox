---
layout: page
title: Commands of Dropbox for teams
lang: en
---

# Member management commands

## Information commands

Below commands are to retrieve information about team members.

| Command                                                                                             | Description                                                                                                        |
|-----------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| [dropbox team member list]({{ site.baseurl }}/commands/dropbox-team-member-list.html)               | Display comprehensive list of all team members with their status, roles, and account details                       |
| [dropbox team member feature]({{ site.baseurl }}/commands/dropbox-team-member-feature.html)         | Display feature settings and capabilities enabled for specific team members, helping understand member permissions |
| [dropbox team member folder list]({{ site.baseurl }}/commands/dropbox-team-member-folder-list.html) | Display all folders in each team member's account, useful for content auditing and storage analysis                |
| [dropbox team member quota list]({{ site.baseurl }}/commands/dropbox-team-member-quota-list.html)   | Display storage quota assignments for all team members, helping monitor and plan storage distribution              |
| [dropbox team member quota usage]({{ site.baseurl }}/commands/dropbox-team-member-quota-usage.html) | Show actual storage usage for each team member compared to their quotas, identifying storage needs                 |
| [dropbox team activity user]({{ site.baseurl }}/commands/dropbox-team-activity-user.html)           | Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities    |

## Basic management commands

Below commands are for managing team member accounts. Those commands are for a bulk operation by a CSV file.

| Command                                                                                                                   | Description                                                                                                      |
|---------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| [dropbox team member batch invite]({{ site.baseurl }}/commands/dropbox-team-member-batch-invite.html)                     | Send batch invitations to new team members, streamlining the onboarding process for multiple users               |
| [dropbox team member batch delete]({{ site.baseurl }}/commands/dropbox-team-member-batch-delete.html)                     | Remove multiple team members in batch, efficiently managing team departures and access revocation                |
| [dropbox team member batch detach]({{ site.baseurl }}/commands/dropbox-team-member-batch-detach.html)                     | Convert multiple team accounts to individual Basic accounts, preserving personal data while removing team access |
| [dropbox team member batch reinvite]({{ site.baseurl }}/commands/dropbox-team-member-batch-reinvite.html)                 | Resend invitations to pending members who haven't joined yet, ensuring all intended members receive access       |
| [dropbox team member update batch email]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-email.html)         | Update email addresses for multiple team members in batch, managing email changes efficiently                    |
| [dropbox team member update batch profile]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-profile.html)     | Update profile information for multiple team members including names and job titles in batch                     |
| [dropbox team member update batch visible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-visible.html)     | Make hidden team members visible in the directory, restoring standard visibility settings                        |
| [dropbox team member update batch invisible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-invisible.html) | Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors               |
| [dropbox team member quota batch update]({{ site.baseurl }}/commands/dropbox-team-member-quota-batch-update.html)         | Modify storage quotas for multiple team members in batch, managing storage allocation efficiently                |

## Member profile setting commands

Member profile commands are for bulk updating member profile information.\nIf you need to update the members' email addresses, use the `member update email` command. The command `member update email` receives a CSV file to bulk update email addresses.\nIf you need to update the member's display name, use the `member update profile` command.

| Command                                                                                                               | Description                                                                                   |
|-----------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| [dropbox team member update batch email]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-email.html)     | Update email addresses for multiple team members in batch, managing email changes efficiently |
| [dropbox team member update batch profile]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-profile.html) | Update profile information for multiple team members including names and job titles in batch  |

## Member storage quota control commands

You can see existing member storage quota setting or usage by the `dropbox team member quota list` and `dropbox team member quota usage` command. If you need to update member quota, use the `dropbox team member quota update` command. The command `dropbox team member quota update` receives CSV input for bulk updating storage quota setting.

| Command                                                                                                           | Description                                                                                           |
|-------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| [dropbox team member quota list]({{ site.baseurl }}/commands/dropbox-team-member-quota-list.html)                 | Display storage quota assignments for all team members, helping monitor and plan storage distribution |
| [dropbox team member quota usage]({{ site.baseurl }}/commands/dropbox-team-member-quota-usage.html)               | Show actual storage usage for each team member compared to their quotas, identifying storage needs    |
| [dropbox team member quota batch update]({{ site.baseurl }}/commands/dropbox-team-member-quota-batch-update.html) | Modify storage quotas for multiple team members in batch, managing storage allocation efficiently     |

## Suspend/unsuspend member commands

There are two types of commands available for suspending/unsuspending members. If you wanted to suspend/unsuspend a member one by one, please use `dropbox team member suspend` or `dropbox team member unsuspend`. Otherwise, if you want to suspend/unsuspend members through a CSV file, please use the `dropbox team member batch suspend` or `dropbox member batch unsuspend` command.

| Command                                                                                                     | Description                                                                                           |
|-------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| [dropbox team member suspend]({{ site.baseurl }}/commands/dropbox-team-member-suspend.html)                 | Temporarily suspend a team member's access to their account while preserving all data and settings    |
| [dropbox team member unsuspend]({{ site.baseurl }}/commands/dropbox-team-member-unsuspend.html)             | Restore access for a suspended team member, reactivating their account and all associated permissions |
| [dropbox team member batch suspend]({{ site.baseurl }}/commands/dropbox-team-member-batch-suspend.html)     | Temporarily suspend multiple team members' access while preserving their data and settings            |
| [dropbox team member batch unsuspend]({{ site.baseurl }}/commands/dropbox-team-member-batch-unsuspend.html) | Restore access for multiple suspended team members, reactivating their accounts in batch              |

## Directory restriction commands

Directory restriction is the Dropbox for teams feature to hide a member from others. Below commands update this setting to hide or unhide members from others.

| Command                                                                                                                   | Description                                                                                        |
|---------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| [dropbox team member update batch visible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-visible.html)     | Make hidden team members visible in the directory, restoring standard visibility settings          |
| [dropbox team member update batch invisible]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-invisible.html) | Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors |

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                                                                             | Description                                                                                                |
|-----------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|
| [dropbox team group add]({{ site.baseurl }}/commands/dropbox-team-group-add.html)                   | Create a new group in your team for organizing members and managing permissions collectively               |
| [dropbox team group batch add]({{ site.baseurl }}/commands/dropbox-team-group-batch-add.html)       | Create multiple groups at once using batch processing, efficient for large-scale team organization         |
| [dropbox team group batch delete]({{ site.baseurl }}/commands/dropbox-team-group-batch-delete.html) | Remove multiple groups from your team in batch, streamlining group cleanup and reorganization              |
| [dropbox team group delete]({{ site.baseurl }}/commands/dropbox-team-group-delete.html)             | Remove a specific group from your team, automatically removing all member associations                     |
| [dropbox team group list]({{ site.baseurl }}/commands/dropbox-team-group-list.html)                 | Display all groups in your team with member counts and group management types                              |
| [dropbox team group rename]({{ site.baseurl }}/commands/dropbox-team-group-rename.html)             | Change the name of an existing group to better reflect its purpose or organizational changes               |
| [dropbox team group update type]({{ site.baseurl }}/commands/dropbox-team-group-update-type.html)   | Change how a group is managed (user-managed vs company-managed), affecting who can modify group membership |

## Group member management commands

You can add/delete/update group members by the below commands. If you want to add/delete/update group members by CSV file, use `dropbox team group member batch add`, `dropbox team group member batch delete`, or `dropbox team group member batch update`.

| Command                                                                                                           | Description                                                                                             |
|-------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| [dropbox team group member add]({{ site.baseurl }}/commands/dropbox-team-group-member-add.html)                   | Add individual team members to a specific group for centralized permission management                   |
| [dropbox team group member delete]({{ site.baseurl }}/commands/dropbox-team-group-member-delete.html)             | Remove a specific member from a group while preserving their other group memberships                    |
| [dropbox team group member list]({{ site.baseurl }}/commands/dropbox-team-group-member-list.html)                 | Display all members belonging to each group, useful for auditing group compositions and access rights   |
| [dropbox team group member batch add]({{ site.baseurl }}/commands/dropbox-team-group-member-batch-add.html)       | Add multiple members to groups efficiently using batch processing, ideal for large team reorganizations |
| [dropbox team group member batch delete]({{ site.baseurl }}/commands/dropbox-team-group-member-batch-delete.html) | Remove multiple members from groups in batch, streamlining group membership management                  |
| [dropbox team group member batch update]({{ site.baseurl }}/commands/dropbox-team-group-member-batch-update.html) | Update group memberships in bulk by adding or removing members, optimizing group composition changes    |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `dropbox team group list`. The command `dropbox team group list` will report the number of members of each group. If it's zero, a group is not currently used to adding permission to folders.\nIf you want to see which folder uses each group, use the command `dropbox team group folder list`. `dropbox team group folder list` will report the group to folder mapping. The report `group_with_no_folders` will show groups with no folders.\nYou can safely remove groups once you check both the number of members and folders. After confirmation, you can bulk delete groups by using the command `dropbox team group batch delete`.

| Command                                                                                             | Description                                                                                            |
|-----------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| [dropbox team group list]({{ site.baseurl }}/commands/dropbox-team-group-list.html)                 | Display all groups in your team with member counts and group management types                          |
| [dropbox team group folder list]({{ site.baseurl }}/commands/dropbox-team-group-folder-list.html)   | Display all folders accessible by each group, showing group-based content organization and permissions |
| [dropbox team group batch delete]({{ site.baseurl }}/commands/dropbox-team-group-batch-delete.html) | Remove multiple groups from your team in batch, streamlining group cleanup and reorganization          |

# Team content commands

Admins can handle team folders, shared folders or member's folder content through Dropbox Business API. Please be careful to use those commands.
The namespace is a term in the Dropbox API that is used to manage folder permissions or settings. Folder types such as shared folders, team folders, or nested folders in a team folder, member's root folder or member's app folder are all managed as a namespace.\nThe namespace commands can handle all types of folders, including team folders and member's folder. But commands for specific folder types have more features or detailed information in the report.

## Team folder operation commands

You can create, archive or permanently delete team folders by using the below commands. Please consider using `dropbox team teamfolder batch` commands if you need to handle multiple team folders.

| Command                                                                                                                     | Description                                                                                             |
|-----------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| [dropbox team teamfolder add]({{ site.baseurl }}/commands/dropbox-team-teamfolder-add.html)                                 | Create a new team folder for centralized team content storage and collaboration                         |
| [dropbox team teamfolder archive]({{ site.baseurl }}/commands/dropbox-team-teamfolder-archive.html)                         | Archive a team folder to make it read-only while preserving all content and access history              |
| [dropbox team teamfolder batch archive]({{ site.baseurl }}/commands/dropbox-team-teamfolder-batch-archive.html)             | Archive multiple team folders in batch, efficiently managing folder lifecycle and compliance            |
| [dropbox team teamfolder batch permdelete]({{ site.baseurl }}/commands/dropbox-team-teamfolder-batch-permdelete.html)       | Permanently delete multiple archived team folders in batch, freeing storage space                       |
| [dropbox team teamfolder batch replication]({{ site.baseurl }}/commands/dropbox-team-teamfolder-batch-replication.html)     | Replicate multiple team folders to another team account in batch for migration or backup                |
| [dropbox team teamfolder file size]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-size.html)                     | Calculate storage usage for team folders, providing detailed size analytics for capacity planning       |
| [dropbox team teamfolder list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-list.html)                               | Display all team folders with their status, sync settings, and member access information                |
| [dropbox team teamfolder permdelete]({{ site.baseurl }}/commands/dropbox-team-teamfolder-permdelete.html)                   | Permanently delete an archived team folder and all its contents, irreversibly freeing storage           |
| [dropbox team teamfolder policy list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-policy-list.html)                 | Display all access policies and restrictions applied to team folders for governance review              |
| [dropbox team teamfolder sync setting list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-sync-setting-list.html)     | Display sync configuration for all team folders, showing default sync behavior for members              |
| [dropbox team teamfolder sync setting update]({{ site.baseurl }}/commands/dropbox-team-teamfolder-sync-setting-update.html) | Modify sync settings for multiple team folders in batch, controlling automatic synchronization behavior |

## Team folder permission commands

You can bulk add or delete members into team folders or sub-folders of a team folder through the below commands.

| Command                                                                                                         | Description                                                                                        |
|-----------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| [dropbox team teamfolder member list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-member-list.html)     | Display all members with access to each team folder, showing permission levels and access types    |
| [dropbox team teamfolder member add]({{ site.baseurl }}/commands/dropbox-team-teamfolder-member-add.html)       | Add multiple users or groups to team folders in batch, streamlining access provisioning            |
| [dropbox team teamfolder member delete]({{ site.baseurl }}/commands/dropbox-team-teamfolder-member-delete.html) | Remove multiple users or groups from team folders in batch, managing access revocation efficiently |

## Team folder & shared folder commands

The below commands are for both team folders and shared folders of the team.\nIf you wanted to know who actually use specific folders, please consider using the command `dropbox team content mount list`. Mount is a status a user add a shared folder to his/her Dropbox account.

| Command                                                                                               | Description                                                                                                                        |
|-------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team content member list]({{ site.baseurl }}/commands/dropbox-team-content-member-list.html) | Display all members with access to team folders and shared folders, showing permission levels and folder relationships             |
| [dropbox team content member size]({{ site.baseurl }}/commands/dropbox-team-content-member-size.html) | Calculate member counts for each team folder and shared folder, helping identify heavily accessed content and optimize permissions |
| [dropbox team content mount list]({{ site.baseurl }}/commands/dropbox-team-content-mount-list.html)   | Display mount status of all shared folders for team members, identifying which folders are actively synced to member devices       |
| [dropbox team content policy list]({{ site.baseurl }}/commands/dropbox-team-content-policy-list.html) | Review all access policies and restrictions applied to team folders and shared folders for governance compliance                   |

## Namespace commands

| Command                                                                                                   | Description                                                                                              |
|-----------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| [dropbox team namespace list]({{ site.baseurl }}/commands/dropbox-team-namespace-list.html)               | Display all team namespaces including team folders and shared spaces with their configurations           |
| [dropbox team namespace summary]({{ site.baseurl }}/commands/dropbox-team-namespace-summary.html)         | Generate comprehensive summary reports of team namespace usage, member counts, and storage statistics    |
| [dropbox team namespace file list]({{ site.baseurl }}/commands/dropbox-team-namespace-file-list.html)     | Display comprehensive file and folder listings within team namespaces for content inventory and analysis |
| [dropbox team namespace file size]({{ site.baseurl }}/commands/dropbox-team-namespace-file-size.html)     | Calculate storage usage for files and folders in team namespaces, providing detailed size analytics      |
| [dropbox team namespace member list]({{ site.baseurl }}/commands/dropbox-team-namespace-member-list.html) | Show all members with access to each namespace, detailing permissions and access levels                  |

## Team file request commands

| Command                                                                                         | Description                                                                                                            |
|-------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [dropbox team filerequest list]({{ site.baseurl }}/commands/dropbox-team-filerequest-list.html) | Display all active and closed file requests created by team members, helping track external file collection activities |

## Member file commands

| Command                                                                                                     | Description                                                                                             |
|-------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| [dropbox team member file permdelete]({{ site.baseurl }}/commands/dropbox-team-member-file-permdelete.html) | Permanently delete files or folders from a team member's account, bypassing trash for immediate removal |

## Team insight

Team Insight is a feature of Dropbox Business that provides a view of team folder summary.

| Command                                                                                                                       | Description                                                                                               |
|-------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [dropbox team insight scan]({{ site.baseurl }}/commands/dropbox-team-insight-scan.html)                                       | Perform comprehensive data scanning across your team for analytics and insights generation                |
| [dropbox team insight scanretry]({{ site.baseurl }}/commands/dropbox-team-insight-scanretry.html)                             | Re-run failed or incomplete scans to ensure complete data collection for team insights                    |
| [dropbox team insight summarize]({{ site.baseurl }}/commands/dropbox-team-insight-summarize.html)                             | Generate summary reports from scanned team data, providing actionable insights on team usage and patterns |
| [dropbox team insight report teamfoldermember]({{ site.baseurl }}/commands/dropbox-team-insight-report-teamfoldermember.html) | Generate detailed reports on team folder membership, showing access patterns and member distribution      |

# Team shared link commands

The team shared link commands are capable of listing all shared links in the team or update/delete specified shared links.

| Command                                                                                                                 | Description                                                                                                   |
|-------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| [dropbox team sharedlink list]({{ site.baseurl }}/commands/dropbox-team-sharedlink-list.html)                           | Display comprehensive list of all shared links created by team members with visibility and expiration details |
| [dropbox team sharedlink cap expiry]({{ site.baseurl }}/commands/dropbox-team-sharedlink-cap-expiry.html)               | Apply expiration date limits to all team shared links for enhanced security and compliance                    |
| [dropbox team sharedlink cap visibility]({{ site.baseurl }}/commands/dropbox-team-sharedlink-cap-visibility.html)       | Enforce visibility restrictions on team shared links, controlling public access levels                        |
| [dropbox team sharedlink update expiry]({{ site.baseurl }}/commands/dropbox-team-sharedlink-update-expiry.html)         | Modify expiration dates for existing shared links across the team to enforce security policies                |
| [dropbox team sharedlink update password]({{ site.baseurl }}/commands/dropbox-team-sharedlink-update-password.html)     | Add or change passwords on team shared links in batch for enhanced security protection                        |
| [dropbox team sharedlink update visibility]({{ site.baseurl }}/commands/dropbox-team-sharedlink-update-visibility.html) | Change access levels of existing shared links between public, team-only, and password-protected               |
| [dropbox team sharedlink delete links]({{ site.baseurl }}/commands/dropbox-team-sharedlink-delete-links.html)           | Delete multiple shared links in batch for security compliance or access control cleanup                       |
| [dropbox team sharedlink delete member]({{ site.baseurl }}/commands/dropbox-team-sharedlink-delete-member.html)         | Remove all shared links created by a specific team member, useful for departing employees                     |

## Difference between `dropbox team sharedlink cap` and `dropbox team sharedlink update`

Commands `dropbox team sharedlink update` is for setting a value to the shared links. Commands `dropbox team sharedlink cap` is for putting a cap value to the shared links.\nFor example: if you set expiry by `dropbox team sharedlink update expiry` with the expiration date 2021-05-06. The command will update the expiry to 2021-05-06 even if the existing link has a shorter expiration date like 2021-05-04.\nOn the other hand, `dropbox team sharedlink cap expiry` updates links when the link has a longer expiration date, like 2021-05-07.\n\nSimilarly, the command `dropbox team sharedlink cap visibility` will restrict visibility only when the link has less protected visibility. For example, if you want to ensure shared links without passwords are restricted to the team only. `dropbox team sharedlink cap visibility` will update visibility to the team only when a link is public and has no password.

## Example (list links):

List all public links in the team\n\n\n\nResults are stored in CSV, xlsx, and JSON format. You can modify the report for updating shared links.\nIf you are familiar with the command jq, you can create CSV file directly like below.\n\n\n\nList links filtered by link owner email address:\n\n\n

## Example (delete links):

Delete all link that listed in the CSV file\n\n\n\nIf you are familiar with jq command, you can send data directly from the pipe like below (pass single dash `-` to the `-file` option to read from standard input).\n\nInvalid argument: team sharedlink delete links -file -n
Error: <no value>

watermint toolbox 140.8.313
===========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox for teams

Usage:
======

tbx  command

Available commands:
===================

| Command | Description              | Notes |
|---------|--------------------------|-------|
| asana   | Asana commands           |       |
| config  | CLI configuration        |       |
| deepl   | DeepL commands           |       |
| dropbox | Dropbox commands         |       |
| figma   | Figma commands           |       |
| github  | GitHub commands          |       |
| license | Show license information |       |
| local   | Commands for local PC    |       |
| log     | Log utilities            |       |
| slack   | Slack commands           |       |
| util    | Utilities                |       |
| version | Show version             |       |\n

# File lock title

Dropbox Business file lock information

## File lock member title

| Command                                                                                                                 | Description                                                                                                       |
|-------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| [dropbox team member file lock all release]({{ site.baseurl }}/commands/dropbox-team-member-file-lock-all-release.html) | Release all file locks held by a team member under a specified path, resolving editing conflicts                  |
| [dropbox team member file lock list]({{ site.baseurl }}/commands/dropbox-team-member-file-lock-list.html)               | Display all files locked by a specific team member under a given path, identifying potential collaboration blocks |
| [dropbox team member file lock release]({{ site.baseurl }}/commands/dropbox-team-member-file-lock-release.html)         | Release a specific file lock held by a team member, enabling others to edit the file                              |

## File lock team folder title

| Command                                                                                                                         | Description                                                                               |
|---------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| [dropbox team teamfolder file list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-list.html)                         | Display all files and subfolders within team folders for content inventory and management |
| [dropbox team teamfolder file lock all release]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-lock-all-release.html) | Release all file locks within a team folder path, resolving editing conflicts in bulk     |
| [dropbox team teamfolder file lock list]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-lock-list.html)               | Display all locked files within team folders, identifying collaboration bottlenecks       |
| [dropbox team teamfolder file lock release]({{ site.baseurl }}/commands/dropbox-team-teamfolder-file-lock-release.html)         | Release specific file locks in team folders to enable collaborative editing               |

# Activities log commands

The team activity log commands can export activity logs by certain filters or perspectives.

| Command                                                                                                 | Description                                                                                                                           |
|---------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team activity batch user]({{ site.baseurl }}/commands/dropbox-team-activity-batch-user.html)   | Scan and retrieve activity logs for multiple team members in batch, useful for compliance auditing and user behavior analysis         |
| [dropbox team activity daily event]({{ site.baseurl }}/commands/dropbox-team-activity-daily-event.html) | Generate daily activity reports showing team events grouped by date, helpful for tracking team usage patterns and security monitoring |
| [dropbox team activity event]({{ site.baseurl }}/commands/dropbox-team-activity-event.html)             | Retrieve detailed team activity event logs with filtering options, essential for security auditing and compliance monitoring          |
| [dropbox team activity user]({{ site.baseurl }}/commands/dropbox-team-activity-user.html)               | Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities                       |

# Connected applications and devices commands

The below commands can retrieve information about connected devices or applications in the team.

| Command                                                                                                 | Description                                                                                                                |
|---------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| [dropbox team device list]({{ site.baseurl }}/commands/dropbox-team-device-list.html)                   | Display all devices and active sessions connected to team member accounts with device details and last activity timestamps |
| [dropbox team device unlink]({{ site.baseurl }}/commands/dropbox-team-device-unlink.html)               | Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access       |
| [dropbox team linkedapp list]({{ site.baseurl }}/commands/dropbox-team-linkedapp-list.html)             | Display all third-party applications linked to team member accounts for security auditing and access control               |
| [dropbox team backup device status]({{ site.baseurl }}/commands/dropbox-team-backup-device-status.html) | Track Dropbox Backup status changes for all team devices over a specified period, monitoring backup health and compliance  |

# Other usecases

## External ID

External ID is the attribute that is not shown in any user interface of Dropbox. This attribute is for keeping a relationship between Dropbox and identity source (e.g. Active Directory, HR database) by identity management software such as Dropbox AD Connector. If you are using Dropbox AD Connector and you built a new Active Directory tree. You may need to clear existing external IDs to disconnect relationships with the old Active Directory tree and the new tree.\nIf you skip clear external IDs, Dropbox AD Connector may unintentionally delete accounts during configuring to the new tree.\nIf you want to see existing external IDs, use the `dropbox team member list` command. But the command will not include external ID by default. Please add the option `-experiment report_all_columns` like below.\n\n\n

| Command                                                                                                                     | Description                                                                                               |
|-----------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [dropbox team member list]({{ site.baseurl }}/commands/dropbox-team-member-list.html)                                       | Display comprehensive list of all team members with their status, roles, and account details              |
| [dropbox team member clear externalid]({{ site.baseurl }}/commands/dropbox-team-member-clear-externalid.html)               | Remove external ID mappings from team members, useful when disconnecting from identity management systems |
| [dropbox team member update batch externalid]({{ site.baseurl }}/commands/dropbox-team-member-update-batch-externalid.html) | Set or update external IDs for multiple team members, integrating with identity management systems        |
| [dropbox team group list]({{ site.baseurl }}/commands/dropbox-team-group-list.html)                                         | Display all groups in your team with member counts and group management types                             |
| [dropbox team group clear externalid]({{ site.baseurl }}/commands/dropbox-team-group-clear-externalid.html)                 | Remove external ID mappings from groups, useful when disconnecting from external identity providers       |

## Data migration helper commands

Data migration helper commands copies member folders or team folders to another account or team. Please test before using those commands before actual data migration.

| Command                                                                                                                     | Description                                                                                                     |
|-----------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| [dropbox team member folder replication]({{ site.baseurl }}/commands/dropbox-team-member-folder-replication.html)           | Copy folder contents from one team member to another's personal space, facilitating content transfer and backup |
| [dropbox team member replication]({{ site.baseurl }}/commands/dropbox-team-member-replication.html)                         | Replicate all files from one team member's account to another, useful for account transitions or backups        |
| [dropbox team teamfolder partial replication]({{ site.baseurl }}/commands/dropbox-team-teamfolder-partial-replication.html) | Selectively replicate team folder contents to another team, enabling flexible content migration                 |
| [dropbox team teamfolder replication]({{ site.baseurl }}/commands/dropbox-team-teamfolder-replication.html)                 | Copy an entire team folder with all contents to another team account for migration or backup                    |

## Team info commands

| Command                                                                             | Description                                                                                                            |
|-------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [dropbox team feature]({{ site.baseurl }}/commands/dropbox-team-feature.html)       | Display all features and capabilities enabled for your Dropbox team account, including API limits and special features |
| [dropbox team filesystem]({{ site.baseurl }}/commands/dropbox-team-filesystem.html) | Identify whether your team uses legacy or modern file system architecture, important for feature compatibility         |
| [dropbox team info]({{ site.baseurl }}/commands/dropbox-team-info.html)             | Display essential team account information including team ID and basic team settings                                   |

# Paper commands

## Legacy paper commands

Commands for a team's legacy Paper content. Please see the [official guide](https://developers.dropbox.com/paper-migration-guide) for more details about legacy Paper and migration

| Command                                                                                                             | Description                                                                                                                        |
|---------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team content legacypaper count]({{ site.baseurl }}/commands/dropbox-team-content-legacypaper-count.html)   | Calculate the total number of legacy Paper documents owned by each team member, useful for content auditing and migration planning |
| [dropbox team content legacypaper list]({{ site.baseurl }}/commands/dropbox-team-content-legacypaper-list.html)     | Generate a comprehensive list of all legacy Paper documents across the team with ownership and metadata information                |
| [dropbox team content legacypaper export]({{ site.baseurl }}/commands/dropbox-team-content-legacypaper-export.html) | Export all legacy Paper documents from team members to local storage in HTML or Markdown format for backup or migration            |

# Team admin commands

Below commands are for managing team admins.

| Command                                                                                                       | Description                                                                                                                      |
|---------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team admin list]({{ site.baseurl }}/commands/dropbox-team-admin-list.html)                           | Display all team members with their assigned admin roles, helpful for auditing administrative access and permissions             |
| [dropbox team admin role add]({{ site.baseurl }}/commands/dropbox-team-admin-role-add.html)                   | Grant a specific admin role to an individual team member, enabling granular permission management                                |
| [dropbox team admin role clear]({{ site.baseurl }}/commands/dropbox-team-admin-role-clear.html)               | Revoke all administrative privileges from a team member, useful for role transitions or security purposes                        |
| [dropbox team admin role delete]({{ site.baseurl }}/commands/dropbox-team-admin-role-delete.html)             | Remove a specific admin role from a team member while preserving other roles, allowing precise permission adjustments            |
| [dropbox team admin role list]({{ site.baseurl }}/commands/dropbox-team-admin-role-list.html)                 | Display all available admin roles in the team with their descriptions and permissions                                            |
| [dropbox team admin group role add]({{ site.baseurl }}/commands/dropbox-team-admin-group-role-add.html)       | Assign admin roles to all members of a specified group, streamlining role management for large teams                             |
| [dropbox team admin group role delete]({{ site.baseurl }}/commands/dropbox-team-admin-group-role-delete.html) | Remove admin roles from all team members except those in a specified exception group, useful for role cleanup and access control |

# Commands that run as a team member

You can run a command as a team member. For example, you can upload a file into member's folder by using `dropbox team runas file sync batch up`.

| Command                                                                                                                                     | Description                                                                                                  |
|---------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------|
| [dropbox team runas file list]({{ site.baseurl }}/commands/dropbox-team-runas-file-list.html)                                               | List files and folders in a team member's account by running operations as that member                       |
| [dropbox team runas file batch copy]({{ site.baseurl }}/commands/dropbox-team-runas-file-batch-copy.html)                                   | Copy multiple files or folders on behalf of team members, useful for content management and organization     |
| [dropbox team runas file sync batch up]({{ site.baseurl }}/commands/dropbox-team-runas-file-sync-batch-up.html)                             | Upload multiple local files to team members' Dropbox accounts in batch, running as those members             |
| [dropbox team runas sharedfolder list]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-list.html)                               | Display all shared folders accessible by a team member, running the operation as that member                 |
| [dropbox team runas sharedfolder isolate]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-isolate.html)                         | Remove all shared folder access for a team member and transfer ownership, useful for departing employees     |
| [dropbox team runas sharedfolder mount add]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-add.html)                     | Mount shared folders to team members' accounts on their behalf, ensuring proper folder synchronization       |
| [dropbox team runas sharedfolder mount delete]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-delete.html)               | Unmount shared folders from team members' accounts on their behalf, managing folder synchronization          |
| [dropbox team runas sharedfolder mount list]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-list.html)                   | Display all shared folders currently mounted (synced) to a specific team member's account                    |
| [dropbox team runas sharedfolder mount mountable]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-mount-mountable.html)         | Show all available shared folders that a team member can mount but hasn't synced yet                         |
| [dropbox team runas sharedfolder batch leave]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-batch-leave.html)                 | Remove team members from multiple shared folders in batch by running leave operations as those members       |
| [dropbox team runas sharedfolder batch share]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-batch-share.html)                 | Share multiple folders on behalf of team members in batch, automating folder sharing processes               |
| [dropbox team runas sharedfolder batch unshare]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-batch-unshare.html)             | Remove sharing from multiple folders on behalf of team members, managing folder access in bulk               |
| [dropbox team runas sharedfolder member batch add]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-member-batch-add.html)       | Add multiple members to shared folders in batch on behalf of folder owners, streamlining access management   |
| [dropbox team runas sharedfolder member batch delete]({{ site.baseurl }}/commands/dropbox-team-runas-sharedfolder-member-batch-delete.html) | Remove multiple members from shared folders in batch on behalf of folder owners, managing access efficiently |

# Legal hold

With legal holds, admins can place a legal hold on members of their team and view and export all the content that's been created or modified by those members.

| Command                                                                                                                   | Description                                                                                               |
|---------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [dropbox team legalhold add]({{ site.baseurl }}/commands/dropbox-team-legalhold-add.html)                                 | Create a legal hold policy to preserve specified team content for compliance or litigation purposes       |
| [dropbox team legalhold list]({{ site.baseurl }}/commands/dropbox-team-legalhold-list.html)                               | Display all active legal hold policies with their details, members, and preservation status               |
| [dropbox team legalhold member batch update]({{ site.baseurl }}/commands/dropbox-team-legalhold-member-batch-update.html) | Add or remove multiple team members from legal hold policies in batch for efficient compliance management |
| [dropbox team legalhold member list]({{ site.baseurl }}/commands/dropbox-team-legalhold-member-list.html)                 | Display all team members currently under legal hold policies with their preservation status               |
| [dropbox team legalhold release]({{ site.baseurl }}/commands/dropbox-team-legalhold-release.html)                         | Release a legal hold policy and restore normal file operations for affected members and content           |
| [dropbox team legalhold revision list]({{ site.baseurl }}/commands/dropbox-team-legalhold-revision-list.html)             | Display all file revisions preserved under legal hold policies, ensuring comprehensive data retention     |
| [dropbox team legalhold update desc]({{ site.baseurl }}/commands/dropbox-team-legalhold-update-desc.html)                 | Modify the description of an existing legal hold policy to reflect changes in scope or purpose            |
| [dropbox team legalhold update name]({{ site.baseurl }}/commands/dropbox-team-legalhold-update-name.html)                 | Change the name of a legal hold policy for better identification and organization                         |

# Notes:

Dropbox Business footnote information


