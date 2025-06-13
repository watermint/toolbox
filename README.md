# watermint toolbox

[![Build](https://github.com/watermint/toolbox/actions/workflows/build.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/build.yml)
[![Test](https://github.com/watermint/toolbox/actions/workflows/test.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/test.yml)
[![CodeQL](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml)
[![Codecov](https://codecov.io/gh/watermint/toolbox/branch/main/graph/badge.svg?token=CrE8reSVvE)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

The watermint toolbox is the multi-purpose utility command-line tool for web services including Dropbox, Figma, GitHub, etc. The purpose of the tool is to provide users of cloud services and system administrators with a way to automate workflows and provide a work-around for some issues.

# Licensing & Disclaimers

watermint toolbox is licensed under the Apache License, Version 2.0.
Please see LICENSE.md or LICENSE.txt for more detail.

Please carefully note:
> Unless required by applicable law or agreed to in writing, Licensor provides the Work (and each Contributor provides its Contributions) on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.

# Built executable

Pre-compiled binaries can be found in [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are building directly from the source, please refer [BUILD.md](BUILD.md).

## Installing using Homebrew on macOS/Linux

First, you need to install Homebrew. Please refer the instruction on [the official site](https://brew.sh/). Then run following commands to install watermint toolbox.
```
brew tap watermint/toolbox
brew install toolbox
```

# Product lifecycle

## Maintenance policy

This product itself is experimental and is not subject to the maintained to keep quality of service. The project will attempt to fix critical bugs and security issues with the best effort. But that is also not guaranteed.\n\nThe product will not release any patch release of a certain major releases. The product will apply fixes as next release when those fixes accepted to do.

## Specification changes

The deliverables of this project are stand-alone executable programs. The specification changes will not be applied unless you explicitly upgrade your version of the program.\n\nThe following policy will be used to make changes in new version releases.\n\nCommand paths, arguments, return values, etc. will be upgraded to be as compatible as possible, but may be discontinued or changed. The general policy is as follows.\n\n* Changes that do not break existing behavior, such as the addition of arguments or changes to messages, will be implemented without notice.\n* Commands that are considered infrequently used will be discontinued or moved without notice.\n* Changes to other commands will be announced 30-180 days or more in advance.\n\nChanges in specifications will be announced at [Announcements](https://github.com/watermint/toolbox/discussions/categories/announcements). Please refer to [Specification Change](https://toolbox.watermint.org/guides/spec-change.html) for a list of planned specification changes.\n

## Availability period for each release

In general, new security issues are discovered every day. To avoid leaving these security and critical issues unaddressed by continuing to use older watermint toolbox releases, a maximum availability period has been set for release 130 and above. Please see [#815](https://github.com/watermint/toolbox/discussions/815) for more details.

# Announcements

* [#886 Releases released after 2024-02-01 will no longer include macOS Intel binaries.](https://github.com/watermint/toolbox/discussions/886)

# Security and privacy

The watermint toolbox is designed to simplify the use of cloud service APIs. It will not use the data in any way that is contrary to your intentions.

The watermint toolbox does not store the data it retrieves via the linked cloud service API on a separate server, contrary to the intent of the specified command.

For example, if you use the watermint toolbox to retrieve data from a cloud service, those data will only be stored on your PC. Furthermore, in the case of commands that upload files or data to a cloud service, they will only be stored in the location specified by your account.

## Data protection

When you use the watermint toolbox to retrieve data from the cloud service API, your data is stored on your PC as report data or log data. More sensitive information, such as the authentication token for the cloud service API, is also stored on your PC.

It is your responsibility to keep these data stored on your PC secure.

Important information such as authentication tokens are obfuscated so that their contents cannot be easily read. However, this obfuscation is not intended to enhance security, but to prevent unintentional operational errors. If a malicious third party copies your token information to another PC, they may be able to access cloud services that you did not intend.

## Use

As previously stated, the watermint toolbox is designed to store data on your PC or in your cloud account. Processes other than your intended operation include data retrieval for release lifecycle management, as outlined below.

The watermint toolbox has the capability to deactivate specific releases that have critical bugs or security issues. This is achieved by retrieving data from a repository hosted on GitHub approximately every 30 days to assess the status of a release. This access does not collect any private data (such as your cloud account information, local files, token, etc.). It merely checks the release status, but as a side effect, your IP address is sent to GitHub when downloading data.

Please be aware that this access information (date, time and IP address) may be used in the future to estimate the usage of each release.

## Sharing

The watermint toolbox project does not currently manage or obtain data including IP addresses, information that only GitHub, the company that hosts the project, has the possibility to access. However, the project may in the future make this information available, and may disclose anonymised release-by-release usage to project members if deemed necessary for the operation of the project.

Any such changes will be announced on the announcement page and this security & privacy policy page at least 30 days before the change takes effect.

# Usage

`tbx` has various features. Run without an option for a list of supported commands and options.\nYou can see available commands and options by running the executable without arguments as shown below.

```
% ./tbx

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox for Teams

Usage:
======

./tbx  command

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
| version | Show version             |       |

```

# Commands

## DeepL

| Command                                                       | Description    |
|---------------------------------------------------------------|----------------|
| [deepl translate text](docs/commands/deepl-translate-text.md) | Translate text |

## Dropbox (Individual account)

| Command                                                                                                 | Description                                                   |
|---------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [dropbox file account feature](docs/commands/dropbox-file-account-feature.md)                           | List Dropbox account features                                 |
| [dropbox file account filesystem](docs/commands/dropbox-file-account-filesystem.md)                     | Show Dropbox file system version                              |
| [dropbox file account info](docs/commands/dropbox-file-account-info.md)                                 | Dropbox account info                                          |
| [dropbox file compare account](docs/commands/dropbox-file-compare-account.md)                           | Compare files of two accounts                                 |
| [dropbox file compare local](docs/commands/dropbox-file-compare-local.md)                               | Compare local folders and Dropbox folders                     |
| [dropbox file copy](docs/commands/dropbox-file-copy.md)                                                 | Copy files                                                    |
| [dropbox file delete](docs/commands/dropbox-file-delete.md)                                             | Delete file or folder                                         |
| [dropbox file export doc](docs/commands/dropbox-file-export-doc.md)                                     | Export document                                               |
| [dropbox file export url](docs/commands/dropbox-file-export-url.md)                                     | Export a document from the URL                                |
| [dropbox file import batch url](docs/commands/dropbox-file-import-batch-url.md)                         | Batch import files from URL                                   |
| [dropbox file import url](docs/commands/dropbox-file-import-url.md)                                     | Import file from the URL                                      |
| [dropbox file info](docs/commands/dropbox-file-info.md)                                                 | Resolve metadata of the path                                  |
| [dropbox file list](docs/commands/dropbox-file-list.md)                                                 | List files and folders                                        |
| [dropbox file lock acquire](docs/commands/dropbox-file-lock-acquire.md)                                 | Lock a file                                                   |
| [dropbox file lock all release](docs/commands/dropbox-file-lock-all-release.md)                         | Release all locks under the specified path                    |
| [dropbox file lock batch acquire](docs/commands/dropbox-file-lock-batch-acquire.md)                     | Lock multiple files                                           |
| [dropbox file lock batch release](docs/commands/dropbox-file-lock-batch-release.md)                     | Release multiple locks                                        |
| [dropbox file lock list](docs/commands/dropbox-file-lock-list.md)                                       | List locks under the specified path                           |
| [dropbox file lock release](docs/commands/dropbox-file-lock-release.md)                                 | Release a lock                                                |
| [dropbox file merge](docs/commands/dropbox-file-merge.md)                                               | Merge paths                                                   |
| [dropbox file move](docs/commands/dropbox-file-move.md)                                                 | Move files                                                    |
| [dropbox file replication](docs/commands/dropbox-file-replication.md)                                   | Replicate file content to the other account                   |
| [dropbox file request create](docs/commands/dropbox-file-request-create.md)                             | Create a file request                                         |
| [dropbox file request delete closed](docs/commands/dropbox-file-request-delete-closed.md)               | Delete all closed file requests on this account.              |
| [dropbox file request delete url](docs/commands/dropbox-file-request-delete-url.md)                     | Delete a file request by the file request URL                 |
| [dropbox file request list](docs/commands/dropbox-file-request-list.md)                                 | List file requests of the individual account                  |
| [dropbox file restore all](docs/commands/dropbox-file-restore-all.md)                                   | Restore files under given path                                |
| [dropbox file restore ext](docs/commands/dropbox-file-restore-ext.md)                                   | Restore files with a specific extension                       |
| [dropbox file revision download](docs/commands/dropbox-file-revision-download.md)                       | Download the file revision                                    |
| [dropbox file revision list](docs/commands/dropbox-file-revision-list.md)                               | List file revisions                                           |
| [dropbox file revision restore](docs/commands/dropbox-file-revision-restore.md)                         | Restore the file revision                                     |
| [dropbox file search content](docs/commands/dropbox-file-search-content.md)                             | Search file content                                           |
| [dropbox file search name](docs/commands/dropbox-file-search-name.md)                                   | Search file name                                              |
| [dropbox file share info](docs/commands/dropbox-file-share-info.md)                                     | Retrieve sharing information of the file                      |
| [dropbox file sharedfolder info](docs/commands/dropbox-file-sharedfolder-info.md)                       | Get shared folder info                                        |
| [dropbox file sharedfolder leave](docs/commands/dropbox-file-sharedfolder-leave.md)                     | Leave the shared folder                                       |
| [dropbox file sharedfolder list](docs/commands/dropbox-file-sharedfolder-list.md)                       | List shared folders                                           |
| [dropbox file sharedfolder member add](docs/commands/dropbox-file-sharedfolder-member-add.md)           | Add a member to the shared folder                             |
| [dropbox file sharedfolder member delete](docs/commands/dropbox-file-sharedfolder-member-delete.md)     | Remove a member from the shared folder                        |
| [dropbox file sharedfolder member list](docs/commands/dropbox-file-sharedfolder-member-list.md)         | List shared folder members                                    |
| [dropbox file sharedfolder mount add](docs/commands/dropbox-file-sharedfolder-mount-add.md)             | Add the shared folder to the current user's Dropbox           |
| [dropbox file sharedfolder mount delete](docs/commands/dropbox-file-sharedfolder-mount-delete.md)       | Unmount the shared folder                                     |
| [dropbox file sharedfolder mount list](docs/commands/dropbox-file-sharedfolder-mount-list.md)           | List all shared folders the current user has mounted          |
| [dropbox file sharedfolder mount mountable](docs/commands/dropbox-file-sharedfolder-mount-mountable.md) | List all shared folders the current user can mount            |
| [dropbox file sharedfolder share](docs/commands/dropbox-file-sharedfolder-share.md)                     | Share a folder                                                |
| [dropbox file sharedfolder unshare](docs/commands/dropbox-file-sharedfolder-unshare.md)                 | Unshare a folder                                              |
| [dropbox file sharedlink create](docs/commands/dropbox-file-sharedlink-create.md)                       | Create shared link                                            |
| [dropbox file sharedlink delete](docs/commands/dropbox-file-sharedlink-delete.md)                       | Remove shared links                                           |
| [dropbox file sharedlink file list](docs/commands/dropbox-file-sharedlink-file-list.md)                 | List files for the shared link                                |
| [dropbox file sharedlink info](docs/commands/dropbox-file-sharedlink-info.md)                           | Get information about the shared link                         |
| [dropbox file sharedlink list](docs/commands/dropbox-file-sharedlink-list.md)                           | List shared links                                             |
| [dropbox file size](docs/commands/dropbox-file-size.md)                                                 | Storage usage                                                 |
| [dropbox file sync down](docs/commands/dropbox-file-sync-down.md)                                       | Downstream sync with Dropbox                                  |
| [dropbox file sync online](docs/commands/dropbox-file-sync-online.md)                                   | Sync online files                                             |
| [dropbox file sync up](docs/commands/dropbox-file-sync-up.md)                                           | Upstream sync with Dropbox                                    |
| [dropbox file tag add](docs/commands/dropbox-file-tag-add.md)                                           | Add tag to file or folder                                     |
| [dropbox file tag delete](docs/commands/dropbox-file-tag-delete.md)                                     | Delete a tag from the file/folder                             |
| [dropbox file tag list](docs/commands/dropbox-file-tag-list.md)                                         | List tags of the path                                         |
| [dropbox file template apply](docs/commands/dropbox-file-template-apply.md)                             | Apply file/folder structure template to the Dropbox path      |
| [dropbox file template capture](docs/commands/dropbox-file-template-capture.md)                         | Capture file/folder structure as template from Dropbox path   |
| [dropbox file watch](docs/commands/dropbox-file-watch.md)                                               | Watch file activities                                         |
| [dropbox paper append](docs/commands/dropbox-paper-append.md)                                           | Append the content to the end of the existing Paper doc       |
| [dropbox paper create](docs/commands/dropbox-paper-create.md)                                           | Create new Paper in the path                                  |
| [dropbox paper overwrite](docs/commands/dropbox-paper-overwrite.md)                                     | Overwrite an existing Paper document                          |
| [dropbox paper prepend](docs/commands/dropbox-paper-prepend.md)                                         | Append the content to the beginning of the existing Paper doc |
| [util tidy pack remote](docs/commands/util-tidy-pack-remote.md)                                         | Package remote folder into the zip file                       |

## Dropbox Sign

| Command                                                                                     | Description                 |
|---------------------------------------------------------------------------------------------|-----------------------------|
| [dropbox sign request list](docs/commands/dropbox-sign-request-list.md)                     | List signature requests     |
| [dropbox sign request signature list](docs/commands/dropbox-sign-request-signature-list.md) | List signatures of requests |

## Dropbox for teams

| Command                                                                                                                     | Description                                                                                                                           |
|-----------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team activity batch user](docs/commands/dropbox-team-activity-batch-user.md)                                       | Scan and retrieve activity logs for multiple team members in batch, useful for compliance auditing and user behavior analysis         |
| [dropbox team activity daily event](docs/commands/dropbox-team-activity-daily-event.md)                                     | Generate daily activity reports showing team events grouped by date, helpful for tracking team usage patterns and security monitoring |
| [dropbox team activity event](docs/commands/dropbox-team-activity-event.md)                                                 | Retrieve detailed team activity event logs with filtering options, essential for security auditing and compliance monitoring          |
| [dropbox team activity user](docs/commands/dropbox-team-activity-user.md)                                                   | Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities                       |
| [dropbox team admin group role add](docs/commands/dropbox-team-admin-group-role-add.md)                                     | Assign admin roles to all members of a specified group, streamlining role management for large teams                                  |
| [dropbox team admin group role delete](docs/commands/dropbox-team-admin-group-role-delete.md)                               | Remove admin roles from all team members except those in a specified exception group, useful for role cleanup and access control      |
| [dropbox team admin list](docs/commands/dropbox-team-admin-list.md)                                                         | Display all team members with their assigned admin roles, helpful for auditing administrative access and permissions                  |
| [dropbox team admin role add](docs/commands/dropbox-team-admin-role-add.md)                                                 | Grant a specific admin role to an individual team member, enabling granular permission management                                     |
| [dropbox team admin role clear](docs/commands/dropbox-team-admin-role-clear.md)                                             | Revoke all administrative privileges from a team member, useful for role transitions or security purposes                             |
| [dropbox team admin role delete](docs/commands/dropbox-team-admin-role-delete.md)                                           | Remove a specific admin role from a team member while preserving other roles, allowing precise permission adjustments                 |
| [dropbox team admin role list](docs/commands/dropbox-team-admin-role-list.md)                                               | Display all available admin roles in the team with their descriptions and permissions                                                 |
| [dropbox team backup device status](docs/commands/dropbox-team-backup-device-status.md)                                     | Track Dropbox Backup status changes for all team devices over a specified period, monitoring backup health and compliance             |
| [dropbox team content legacypaper count](docs/commands/dropbox-team-content-legacypaper-count.md)                           | Calculate the total number of legacy Paper documents owned by each team member, useful for content auditing and migration planning    |
| [dropbox team content legacypaper export](docs/commands/dropbox-team-content-legacypaper-export.md)                         | Export all legacy Paper documents from team members to local storage in HTML or Markdown format for backup or migration               |
| [dropbox team content legacypaper list](docs/commands/dropbox-team-content-legacypaper-list.md)                             | Generate a comprehensive list of all legacy Paper documents across the team with ownership and metadata information                   |
| [dropbox team content member list](docs/commands/dropbox-team-content-member-list.md)                                       | Display all members with access to team folders and shared folders, showing permission levels and folder relationships                |
| [dropbox team content member size](docs/commands/dropbox-team-content-member-size.md)                                       | Calculate member counts for each team folder and shared folder, helping identify heavily accessed content and optimize permissions    |
| [dropbox team content mount list](docs/commands/dropbox-team-content-mount-list.md)                                         | Display mount status of all shared folders for team members, identifying which folders are actively synced to member devices          |
| [dropbox team content policy list](docs/commands/dropbox-team-content-policy-list.md)                                       | Review all access policies and restrictions applied to team folders and shared folders for governance compliance                      |
| [dropbox team device list](docs/commands/dropbox-team-device-list.md)                                                       | Display all devices and active sessions connected to team member accounts with device details and last activity timestamps            |
| [dropbox team device unlink](docs/commands/dropbox-team-device-unlink.md)                                                   | Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access                  |
| [dropbox team feature](docs/commands/dropbox-team-feature.md)                                                               | Display all features and capabilities enabled for your Dropbox team account, including API limits and special features                |
| [dropbox team filerequest list](docs/commands/dropbox-team-filerequest-list.md)                                             | Display all active and closed file requests created by team members, helping track external file collection activities                |
| [dropbox team filesystem](docs/commands/dropbox-team-filesystem.md)                                                         | Identify whether your team uses legacy or modern file system architecture, important for feature compatibility                        |
| [dropbox team group add](docs/commands/dropbox-team-group-add.md)                                                           | Create a new group in your team for organizing members and managing permissions collectively                                          |
| [dropbox team group batch add](docs/commands/dropbox-team-group-batch-add.md)                                               | Create multiple groups at once using batch processing, efficient for large-scale team organization                                    |
| [dropbox team group batch delete](docs/commands/dropbox-team-group-batch-delete.md)                                         | Remove multiple groups from your team in batch, streamlining group cleanup and reorganization                                         |
| [dropbox team group clear externalid](docs/commands/dropbox-team-group-clear-externalid.md)                                 | Remove external ID mappings from groups, useful when disconnecting from external identity providers                                   |
| [dropbox team group delete](docs/commands/dropbox-team-group-delete.md)                                                     | Remove a specific group from your team, automatically removing all member associations                                                |
| [dropbox team group folder list](docs/commands/dropbox-team-group-folder-list.md)                                           | Display all folders accessible by each group, showing group-based content organization and permissions                                |
| [dropbox team group list](docs/commands/dropbox-team-group-list.md)                                                         | Display all groups in your team with member counts and group management types                                                         |
| [dropbox team group member add](docs/commands/dropbox-team-group-member-add.md)                                             | Add individual team members to a specific group for centralized permission management                                                 |
| [dropbox team group member batch add](docs/commands/dropbox-team-group-member-batch-add.md)                                 | Add multiple members to groups efficiently using batch processing, ideal for large team reorganizations                               |
| [dropbox team group member batch delete](docs/commands/dropbox-team-group-member-batch-delete.md)                           | Remove multiple members from groups in batch, streamlining group membership management                                                |
| [dropbox team group member batch update](docs/commands/dropbox-team-group-member-batch-update.md)                           | Update group memberships in bulk by adding or removing members, optimizing group composition changes                                  |
| [dropbox team group member delete](docs/commands/dropbox-team-group-member-delete.md)                                       | Remove a specific member from a group while preserving their other group memberships                                                  |
| [dropbox team group member list](docs/commands/dropbox-team-group-member-list.md)                                           | Display all members belonging to each group, useful for auditing group compositions and access rights                                 |
| [dropbox team group rename](docs/commands/dropbox-team-group-rename.md)                                                     | Change the name of an existing group to better reflect its purpose or organizational changes                                          |
| [dropbox team group update type](docs/commands/dropbox-team-group-update-type.md)                                           | Change how a group is managed (user-managed vs company-managed), affecting who can modify group membership                            |
| [dropbox team info](docs/commands/dropbox-team-info.md)                                                                     | Display essential team account information including team ID and basic team settings                                                  |
| [dropbox team insight scan](docs/commands/dropbox-team-insight-scan.md)                                                     | Perform comprehensive data scanning across your team for analytics and insights generation                                            |
| [dropbox team legalhold add](docs/commands/dropbox-team-legalhold-add.md)                                                   | Create a legal hold policy to preserve specified team content for compliance or litigation purposes                                   |
| [dropbox team legalhold list](docs/commands/dropbox-team-legalhold-list.md)                                                 | Display all active legal hold policies with their details, members, and preservation status                                           |
| [dropbox team legalhold member batch update](docs/commands/dropbox-team-legalhold-member-batch-update.md)                   | Add or remove multiple team members from legal hold policies in batch for efficient compliance management                             |
| [dropbox team legalhold member list](docs/commands/dropbox-team-legalhold-member-list.md)                                   | Display all team members currently under legal hold policies with their preservation status                                           |
| [dropbox team legalhold release](docs/commands/dropbox-team-legalhold-release.md)                                           | Release a legal hold policy and restore normal file operations for affected members and content                                       |
| [dropbox team legalhold revision list](docs/commands/dropbox-team-legalhold-revision-list.md)                               | Display all file revisions preserved under legal hold policies, ensuring comprehensive data retention                                 |
| [dropbox team legalhold update desc](docs/commands/dropbox-team-legalhold-update-desc.md)                                   | Modify the description of an existing legal hold policy to reflect changes in scope or purpose                                        |
| [dropbox team legalhold update name](docs/commands/dropbox-team-legalhold-update-name.md)                                   | Change the name of a legal hold policy for better identification and organization                                                     |
| [dropbox team linkedapp list](docs/commands/dropbox-team-linkedapp-list.md)                                                 | Display all third-party applications linked to team member accounts for security auditing and access control                          |
| [dropbox team member batch delete](docs/commands/dropbox-team-member-batch-delete.md)                                       | Remove multiple team members in batch, efficiently managing team departures and access revocation                                     |
| [dropbox team member batch detach](docs/commands/dropbox-team-member-batch-detach.md)                                       | Convert multiple team accounts to individual Basic accounts, preserving personal data while removing team access                      |
| [dropbox team member batch invite](docs/commands/dropbox-team-member-batch-invite.md)                                       | Send batch invitations to new team members, streamlining the onboarding process for multiple users                                    |
| [dropbox team member batch reinvite](docs/commands/dropbox-team-member-batch-reinvite.md)                                   | Resend invitations to pending members who haven't joined yet, ensuring all intended members receive access                            |
| [dropbox team member batch suspend](docs/commands/dropbox-team-member-batch-suspend.md)                                     | Temporarily suspend multiple team members' access while preserving their data and settings                                            |
| [dropbox team member batch unsuspend](docs/commands/dropbox-team-member-batch-unsuspend.md)                                 | Restore access for multiple suspended team members, reactivating their accounts in batch                                              |
| [dropbox team member clear externalid](docs/commands/dropbox-team-member-clear-externalid.md)                               | Remove external ID mappings from team members, useful when disconnecting from identity management systems                             |
| [dropbox team member feature](docs/commands/dropbox-team-member-feature.md)                                                 | Display feature settings and capabilities enabled for specific team members, helping understand member permissions                    |
| [dropbox team member file lock all release](docs/commands/dropbox-team-member-file-lock-all-release.md)                     | Release all file locks held by a team member under a specified path, resolving editing conflicts                                      |
| [dropbox team member file lock list](docs/commands/dropbox-team-member-file-lock-list.md)                                   | Display all files locked by a specific team member under a given path, identifying potential collaboration blocks                     |
| [dropbox team member file lock release](docs/commands/dropbox-team-member-file-lock-release.md)                             | Release a specific file lock held by a team member, enabling others to edit the file                                                  |
| [dropbox team member file permdelete](docs/commands/dropbox-team-member-file-permdelete.md)                                 | Permanently delete files or folders from a team member's account, bypassing trash for immediate removal                               |
| [dropbox team member folder list](docs/commands/dropbox-team-member-folder-list.md)                                         | Display all folders in each team member's account, useful for content auditing and storage analysis                                   |
| [dropbox team member folder replication](docs/commands/dropbox-team-member-folder-replication.md)                           | Copy folder contents from one team member to another's personal space, facilitating content transfer and backup                       |
| [dropbox team member list](docs/commands/dropbox-team-member-list.md)                                                       | Display comprehensive list of all team members with their status, roles, and account details                                          |
| [dropbox team member quota batch update](docs/commands/dropbox-team-member-quota-batch-update.md)                           | Modify storage quotas for multiple team members in batch, managing storage allocation efficiently                                     |
| [dropbox team member quota list](docs/commands/dropbox-team-member-quota-list.md)                                           | Display storage quota assignments for all team members, helping monitor and plan storage distribution                                 |
| [dropbox team member quota usage](docs/commands/dropbox-team-member-quota-usage.md)                                         | Show actual storage usage for each team member compared to their quotas, identifying storage needs                                    |
| [dropbox team member replication](docs/commands/dropbox-team-member-replication.md)                                         | Replicate all files from one team member's account to another, useful for account transitions or backups                              |
| [dropbox team member suspend](docs/commands/dropbox-team-member-suspend.md)                                                 | Temporarily suspend a team member's access to their account while preserving all data and settings                                    |
| [dropbox team member unsuspend](docs/commands/dropbox-team-member-unsuspend.md)                                             | Restore access for a suspended team member, reactivating their account and all associated permissions                                 |
| [dropbox team member update batch email](docs/commands/dropbox-team-member-update-batch-email.md)                           | Update email addresses for multiple team members in batch, managing email changes efficiently                                         |
| [dropbox team member update batch externalid](docs/commands/dropbox-team-member-update-batch-externalid.md)                 | Set or update external IDs for multiple team members, integrating with identity management systems                                    |
| [dropbox team member update batch invisible](docs/commands/dropbox-team-member-update-batch-invisible.md)                   | Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors                                    |
| [dropbox team member update batch profile](docs/commands/dropbox-team-member-update-batch-profile.md)                       | Update profile information for multiple team members including names and job titles in batch                                          |
| [dropbox team member update batch visible](docs/commands/dropbox-team-member-update-batch-visible.md)                       | Make hidden team members visible in the directory, restoring standard visibility settings                                             |
| [dropbox team namespace file list](docs/commands/dropbox-team-namespace-file-list.md)                                       | Display comprehensive file and folder listings within team namespaces for content inventory and analysis                              |
| [dropbox team namespace file size](docs/commands/dropbox-team-namespace-file-size.md)                                       | Calculate storage usage for files and folders in team namespaces, providing detailed size analytics                                   |
| [dropbox team namespace list](docs/commands/dropbox-team-namespace-list.md)                                                 | Display all team namespaces including team folders and shared spaces with their configurations                                        |
| [dropbox team namespace member list](docs/commands/dropbox-team-namespace-member-list.md)                                   | Show all members with access to each namespace, detailing permissions and access levels                                               |
| [dropbox team namespace summary](docs/commands/dropbox-team-namespace-summary.md)                                           | Generate comprehensive summary reports of team namespace usage, member counts, and storage statistics                                 |
| [dropbox team runas file batch copy](docs/commands/dropbox-team-runas-file-batch-copy.md)                                   | Copy multiple files or folders on behalf of team members, useful for content management and organization                              |
| [dropbox team runas file list](docs/commands/dropbox-team-runas-file-list.md)                                               | List files and folders in a team member's account by running operations as that member                                                |
| [dropbox team runas file sync batch up](docs/commands/dropbox-team-runas-file-sync-batch-up.md)                             | Upload multiple local files to team members' Dropbox accounts in batch, running as those members                                      |
| [dropbox team runas sharedfolder batch leave](docs/commands/dropbox-team-runas-sharedfolder-batch-leave.md)                 | Remove team members from multiple shared folders in batch by running leave operations as those members                                |
| [dropbox team runas sharedfolder batch share](docs/commands/dropbox-team-runas-sharedfolder-batch-share.md)                 | Share multiple folders on behalf of team members in batch, automating folder sharing processes                                        |
| [dropbox team runas sharedfolder batch unshare](docs/commands/dropbox-team-runas-sharedfolder-batch-unshare.md)             | Remove sharing from multiple folders on behalf of team members, managing folder access in bulk                                        |
| [dropbox team runas sharedfolder isolate](docs/commands/dropbox-team-runas-sharedfolder-isolate.md)                         | Remove all shared folder access for a team member and transfer ownership, useful for departing employees                              |
| [dropbox team runas sharedfolder list](docs/commands/dropbox-team-runas-sharedfolder-list.md)                               | Display all shared folders accessible by a team member, running the operation as that member                                          |
| [dropbox team runas sharedfolder member batch add](docs/commands/dropbox-team-runas-sharedfolder-member-batch-add.md)       | Add multiple members to shared folders in batch on behalf of folder owners, streamlining access management                            |
| [dropbox team runas sharedfolder member batch delete](docs/commands/dropbox-team-runas-sharedfolder-member-batch-delete.md) | Remove multiple members from shared folders in batch on behalf of folder owners, managing access efficiently                          |
| [dropbox team runas sharedfolder mount add](docs/commands/dropbox-team-runas-sharedfolder-mount-add.md)                     | Mount shared folders to team members' accounts on their behalf, ensuring proper folder synchronization                                |
| [dropbox team runas sharedfolder mount delete](docs/commands/dropbox-team-runas-sharedfolder-mount-delete.md)               | Unmount shared folders from team members' accounts on their behalf, managing folder synchronization                                   |
| [dropbox team runas sharedfolder mount list](docs/commands/dropbox-team-runas-sharedfolder-mount-list.md)                   | Display all shared folders currently mounted (synced) to a specific team member's account                                             |
| [dropbox team runas sharedfolder mount mountable](docs/commands/dropbox-team-runas-sharedfolder-mount-mountable.md)         | Show all available shared folders that a team member can mount but hasn't synced yet                                                  |
| [dropbox team sharedlink cap expiry](docs/commands/dropbox-team-sharedlink-cap-expiry.md)                                   | Apply expiration date limits to all team shared links for enhanced security and compliance                                            |
| [dropbox team sharedlink cap visibility](docs/commands/dropbox-team-sharedlink-cap-visibility.md)                           | Enforce visibility restrictions on team shared links, controlling public access levels                                                |
| [dropbox team sharedlink delete links](docs/commands/dropbox-team-sharedlink-delete-links.md)                               | Delete multiple shared links in batch for security compliance or access control cleanup                                               |
| [dropbox team sharedlink delete member](docs/commands/dropbox-team-sharedlink-delete-member.md)                             | Remove all shared links created by a specific team member, useful for departing employees                                             |
| [dropbox team sharedlink list](docs/commands/dropbox-team-sharedlink-list.md)                                               | Display comprehensive list of all shared links created by team members with visibility and expiration details                         |
| [dropbox team sharedlink update expiry](docs/commands/dropbox-team-sharedlink-update-expiry.md)                             | Modify expiration dates for existing shared links across the team to enforce security policies                                        |
| [dropbox team sharedlink update password](docs/commands/dropbox-team-sharedlink-update-password.md)                         | Add or change passwords on team shared links in batch for enhanced security protection                                                |
| [dropbox team sharedlink update visibility](docs/commands/dropbox-team-sharedlink-update-visibility.md)                     | Change access levels of existing shared links between public, team-only, and password-protected                                       |
| [dropbox team teamfolder add](docs/commands/dropbox-team-teamfolder-add.md)                                                 | Create a new team folder for centralized team content storage and collaboration                                                       |
| [dropbox team teamfolder archive](docs/commands/dropbox-team-teamfolder-archive.md)                                         | Archive a team folder to make it read-only while preserving all content and access history                                            |
| [dropbox team teamfolder batch archive](docs/commands/dropbox-team-teamfolder-batch-archive.md)                             | Archive multiple team folders in batch, efficiently managing folder lifecycle and compliance                                          |
| [dropbox team teamfolder batch permdelete](docs/commands/dropbox-team-teamfolder-batch-permdelete.md)                       | Permanently delete multiple archived team folders in batch, freeing storage space                                                     |
| [dropbox team teamfolder batch replication](docs/commands/dropbox-team-teamfolder-batch-replication.md)                     | Replicate multiple team folders to another team account in batch for migration or backup                                              |
| [dropbox team teamfolder file list](docs/commands/dropbox-team-teamfolder-file-list.md)                                     | Display all files and subfolders within team folders for content inventory and management                                             |
| [dropbox team teamfolder file lock all release](docs/commands/dropbox-team-teamfolder-file-lock-all-release.md)             | Release all file locks within a team folder path, resolving editing conflicts in bulk                                                 |
| [dropbox team teamfolder file lock list](docs/commands/dropbox-team-teamfolder-file-lock-list.md)                           | Display all locked files within team folders, identifying collaboration bottlenecks                                                   |
| [dropbox team teamfolder file lock release](docs/commands/dropbox-team-teamfolder-file-lock-release.md)                     | Release specific file locks in team folders to enable collaborative editing                                                           |
| [dropbox team teamfolder file size](docs/commands/dropbox-team-teamfolder-file-size.md)                                     | Calculate storage usage for team folders, providing detailed size analytics for capacity planning                                     |
| [dropbox team teamfolder list](docs/commands/dropbox-team-teamfolder-list.md)                                               | Display all team folders with their status, sync settings, and member access information                                              |
| [dropbox team teamfolder member add](docs/commands/dropbox-team-teamfolder-member-add.md)                                   | Add multiple users or groups to team folders in batch, streamlining access provisioning                                               |
| [dropbox team teamfolder member delete](docs/commands/dropbox-team-teamfolder-member-delete.md)                             | Remove multiple users or groups from team folders in batch, managing access revocation efficiently                                    |
| [dropbox team teamfolder member list](docs/commands/dropbox-team-teamfolder-member-list.md)                                 | Display all members with access to each team folder, showing permission levels and access types                                       |
| [dropbox team teamfolder partial replication](docs/commands/dropbox-team-teamfolder-partial-replication.md)                 | Selectively replicate team folder contents to another team, enabling flexible content migration                                       |
| [dropbox team teamfolder permdelete](docs/commands/dropbox-team-teamfolder-permdelete.md)                                   | Permanently delete an archived team folder and all its contents, irreversibly freeing storage                                         |
| [dropbox team teamfolder policy list](docs/commands/dropbox-team-teamfolder-policy-list.md)                                 | Display all access policies and restrictions applied to team folders for governance review                                            |
| [dropbox team teamfolder replication](docs/commands/dropbox-team-teamfolder-replication.md)                                 | Copy an entire team folder with all contents to another team account for migration or backup                                          |
| [dropbox team teamfolder sync setting list](docs/commands/dropbox-team-teamfolder-sync-setting-list.md)                     | Display sync configuration for all team folders, showing default sync behavior for members                                            |
| [dropbox team teamfolder sync setting update](docs/commands/dropbox-team-teamfolder-sync-setting-update.md)                 | Modify sync settings for multiple team folders in batch, controlling automatic synchronization behavior                               |

## Figma

| Command                                                                   | Description                           |
|---------------------------------------------------------------------------|---------------------------------------|
| [figma account info](docs/commands/figma-account-info.md)                 | Retrieve current user information     |
| [figma file export all page](docs/commands/figma-file-export-all-page.md) | Export all files/pages under the team |
| [figma file export frame](docs/commands/figma-file-export-frame.md)       | Export all frames of the Figma file   |
| [figma file export node](docs/commands/figma-file-export-node.md)         | Export Figma document Node            |
| [figma file export page](docs/commands/figma-file-export-page.md)         | Export all pages of the Figma file    |
| [figma file info](docs/commands/figma-file-info.md)                       | Show information of the Figma file    |
| [figma file list](docs/commands/figma-file-list.md)                       | List files in the Figma Project       |
| [figma project list](docs/commands/figma-project-list.md)                 | List projects of the team             |

## GitHub

| Command                                                                         | Description                                         |
|---------------------------------------------------------------------------------|-----------------------------------------------------|
| [github content get](docs/commands/github-content-get.md)                       | Get content metadata of the repository              |
| [github content put](docs/commands/github-content-put.md)                       | Put small text content into the repository          |
| [github issue list](docs/commands/github-issue-list.md)                         | List issues of the public/private GitHub repository |
| [github profile](docs/commands/github-profile.md)                               | Get the authenticated user                          |
| [github release asset download](docs/commands/github-release-asset-download.md) | Download assets                                     |
| [github release asset list](docs/commands/github-release-asset-list.md)         | List assets of GitHub Release                       |
| [github release asset upload](docs/commands/github-release-asset-upload.md)     | Upload assets file into the GitHub Release          |
| [github release draft](docs/commands/github-release-draft.md)                   | Create release draft                                |
| [github release list](docs/commands/github-release-list.md)                     | List releases                                       |
| [github tag create](docs/commands/github-tag-create.md)                         | Create a tag on the repository                      |
| [util release install](docs/commands/util-release-install.md)                   | Download & install watermint toolbox to the path    |

## Utilities

| Command                                                                                                       | Description                                                                                          |
|---------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|
| [config auth delete](docs/commands/config-auth-delete.md)                                                     | Delete existing auth credential                                                                      |
| [config auth list](docs/commands/config-auth-list.md)                                                         | List all auth credentials                                                                            |
| [config feature disable](docs/commands/config-feature-disable.md)                                             | Disable a feature.                                                                                   |
| [config feature enable](docs/commands/config-feature-enable.md)                                               | Enable a feature.                                                                                    |
| [config feature list](docs/commands/config-feature-list.md)                                                   | List available optional features.                                                                    |
| [config license install](docs/commands/config-license-install.md)                                             | Install a license key                                                                                |
| [config license list](docs/commands/config-license-list.md)                                                   | List available license keys                                                                          |
| [dropbox team insight report teamfoldermember](docs/commands/dropbox-team-insight-report-teamfoldermember.md) | Generate detailed reports on team folder membership, showing access patterns and member distribution |
| [license](docs/commands/license.md)                                                                           | Show license information                                                                             |
| [local file template apply](docs/commands/local-file-template-apply.md)                                       | Apply file/folder structure template to the local path                                               |
| [local file template capture](docs/commands/local-file-template-capture.md)                                   | Capture file/folder structure as template from local path                                            |
| [log api job](docs/commands/log-api-job.md)                                                                   | Show statistics of the API log of the job specified by the job ID                                    |
| [log api name](docs/commands/log-api-name.md)                                                                 | Show statistics of the API log of the job specified by the job name                                  |
| [log cat curl](docs/commands/log-cat-curl.md)                                                                 | Format capture logs as `curl` sample                                                                 |
| [log cat job](docs/commands/log-cat-job.md)                                                                   | Retrieve logs of specified Job ID                                                                    |
| [log cat kind](docs/commands/log-cat-kind.md)                                                                 | Concatenate and print logs of specified log kind                                                     |
| [log cat last](docs/commands/log-cat-last.md)                                                                 | Print the last job log files                                                                         |
| [log job archive](docs/commands/log-job-archive.md)                                                           | Archive jobs                                                                                         |
| [log job delete](docs/commands/log-job-delete.md)                                                             | Delete old job history                                                                               |
| [log job list](docs/commands/log-job-list.md)                                                                 | Show job history                                                                                     |
| [util archive unzip](docs/commands/util-archive-unzip.md)                                                     | Extract the zip archive file                                                                         |
| [util archive zip](docs/commands/util-archive-zip.md)                                                         | Compress target files into the zip archive                                                           |
| [util cert selfsigned](docs/commands/util-cert-selfsigned.md)                                                 | Generate self-signed certificate and key                                                             |
| [util database exec](docs/commands/util-database-exec.md)                                                     | Execute query on SQLite3 database file                                                               |
| [util database query](docs/commands/util-database-query.md)                                                   | Query SQLite3 database                                                                               |
| [util date today](docs/commands/util-date-today.md)                                                           | Display current date                                                                                 |
| [util datetime now](docs/commands/util-datetime-now.md)                                                       | Display current date/time                                                                            |
| [util decode base32](docs/commands/util-decode-base32.md)                                                     | Decode text from Base32 (RFC 4648) format                                                            |
| [util decode base64](docs/commands/util-decode-base64.md)                                                     | Decode text from Base64 (RFC 4648) format                                                            |
| [util desktop open](docs/commands/util-desktop-open.md)                                                       | Open a file or folder with the default application                                                   |
| [util encode base32](docs/commands/util-encode-base32.md)                                                     | Encode text into Base32 (RFC 4648) format                                                            |
| [util encode base64](docs/commands/util-encode-base64.md)                                                     | Encode text into Base64 (RFC 4648) format                                                            |
| [util feed json](docs/commands/util-feed-json.md)                                                             | Load feed from the URL and output the content as JSON                                                |
| [util file hash](docs/commands/util-file-hash.md)                                                             | File Hash                                                                                            |
| [util git clone](docs/commands/util-git-clone.md)                                                             | Clone git repository                                                                                 |
| [util image exif](docs/commands/util-image-exif.md)                                                           | Print EXIF metadata of image file                                                                    |
| [util image placeholder](docs/commands/util-image-placeholder.md)                                             | Create placeholder image                                                                             |
| [util json query](docs/commands/util-json-query.md)                                                           | Query JSON data                                                                                      |
| [util net download](docs/commands/util-net-download.md)                                                       | Download a file                                                                                      |
| [util qrcode create](docs/commands/util-qrcode-create.md)                                                     | Create a QR code image file                                                                          |
| [util qrcode wifi](docs/commands/util-qrcode-wifi.md)                                                         | Generate QR code for WIFI configuration                                                              |
| [util table format xlsx](docs/commands/util-table-format-xlsx.md)                                             | Formatting xlsx file into text                                                                       |
| [util text case down](docs/commands/util-text-case-down.md)                                                   | Print lower case text                                                                                |
| [util text case up](docs/commands/util-text-case-up.md)                                                       | Print upper case text                                                                                |
| [util text encoding from](docs/commands/util-text-encoding-from.md)                                           | Convert text encoding to UTF-8 text file from specified encoding.                                    |
| [util text encoding to](docs/commands/util-text-encoding-to.md)                                               | Convert text encoding to specified encoding from UTF-8 text file.                                    |
| [util text nlp english entity](docs/commands/util-text-nlp-english-entity.md)                                 | Split English text into entities                                                                     |
| [util text nlp english sentence](docs/commands/util-text-nlp-english-sentence.md)                             | Split English text into sentences                                                                    |
| [util text nlp english token](docs/commands/util-text-nlp-english-token.md)                                   | Split English text into tokens                                                                       |
| [util text nlp japanese token](docs/commands/util-text-nlp-japanese-token.md)                                 | Tokenize Japanese text                                                                               |
| [util text nlp japanese wakati](docs/commands/util-text-nlp-japanese-wakati.md)                               | Wakachigaki (tokenize Japanese text)                                                                 |
| [util tidy move dispatch](docs/commands/util-tidy-move-dispatch.md)                                           | Dispatch files                                                                                       |
| [util tidy move simple](docs/commands/util-tidy-move-simple.md)                                               | Archive local files                                                                                  |
| [util time now](docs/commands/util-time-now.md)                                                               | Display current time                                                                                 |
| [util unixtime format](docs/commands/util-unixtime-format.md)                                                 | Time format to convert the unix time (epoch seconds from 1970-01-01)                                 |
| [util unixtime now](docs/commands/util-unixtime-now.md)                                                       | Display current time in unixtime                                                                     |
| [util uuid timestamp](docs/commands/util-uuid-timestamp.md)                                                   | UUID Timestamp                                                                                       |
| [util uuid ulid](docs/commands/util-uuid-ulid.md)                                                             | ULID Utility                                                                                         |
| [util uuid v4](docs/commands/util-uuid-v4.md)                                                                 | Generate UUID v4 (random UUID)                                                                       |
| [util uuid v7](docs/commands/util-uuid-v7.md)                                                                 | Generate UUID v7                                                                                     |
| [util uuid version](docs/commands/util-uuid-version.md)                                                       | Parse version and variant of UUID                                                                    |
| [util xlsx create](docs/commands/util-xlsx-create.md)                                                         | Create an empty spreadsheet                                                                          |
| [util xlsx sheet export](docs/commands/util-xlsx-sheet-export.md)                                             | Export data from the xlsx file                                                                       |
| [util xlsx sheet import](docs/commands/util-xlsx-sheet-import.md)                                             | Import data into xlsx file                                                                           |
| [util xlsx sheet list](docs/commands/util-xlsx-sheet-list.md)                                                 | List sheets of the xlsx file                                                                         |
| [version](docs/commands/version.md)                                                                           | Show version                                                                                         |

