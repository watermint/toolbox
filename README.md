# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![codecov](https://codecov.io/gh/watermint/toolbox/branch/main/graph/badge.svg?token=CrE8reSVvE)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

The multi-purpose utility command-line tool for web services including Dropbox, Dropbox Business, Google, GitHub, etc.

# Licensing & Disclaimers

watermint toolbox is licensed under the MIT license.
Please see LICENSE.md or LICENSE.txt for more detail.

Please carefully note:
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

# Built executable

Pre-compiled binaries can be found in [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are building directly from the source, please refer [BUILD.md](BUILD.md).

## Installing using Homebrew on macOS/Linux

First, you need to install Homebrew. Please refer the instruction on [the official site](https://brew.sh/). Then run following commands to install watermint toolbox.
```
brew tap watermint/toolbox
brew install toolbox
```

# Security and privacy

## Information Not Collected 

The watermint toolbox does not collect any information to third-party servers.

The watermint toolbox is for interacting with the services like Dropbox with your account. There is no third-party account involved. The Commands stores API token, logs, files, or reports on your PC storage.

## Sensitive data

Most sensitive data, such as API token, are saved on your PC storage in obfuscated & made restricted access. However, it's your responsibility to keep those data secret. 
Please do not share files, especially the `secrets` folder under toolbox work path (`C:\Users\<your user name>\.toolbox`, or `$HOME/.toolbox` by default).

# Usage

`tbx` have various features. Run without an option for a list of supported commands and options.
You can see available commands and options by running executable without arguments like below.

```
% ./tbx

watermint toolbox xx.x.xxx
==========================

© 2016-2023 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox Business

Usage:
======

./tbx  command

Available commands:
===================

| Command      | Description                     | Notes |
|--------------|---------------------------------|-------|
| config       | watermint toolbox configuration |       |
| file         | File operation                  |       |
| filerequest  | File request operation          |       |
| group        | Group management                |       |
| license      | Show license information        |       |
| member       | Team member management          |       |
| sharedfolder | Shared folder                   |       |
| sharedlink   | Shared Link of Personal account |       |
| team         | Dropbox Business Team           |       |
| teamfolder   | Team folder management          |       |
| version      | Show version                    |       |

```

# Commands

## Dropbox (Individual account)

| Command                                                                         | Description                                                   |
|---------------------------------------------------------------------------------|---------------------------------------------------------------|
| [file compare account](docs/commands/file-compare-account.md)                   | Compare files of two accounts                                 |
| [file compare local](docs/commands/file-compare-local.md)                       | Compare local folders and Dropbox folders                     |
| [file copy](docs/commands/file-copy.md)                                         | Copy files                                                    |
| [file delete](docs/commands/file-delete.md)                                     | Delete file or folder                                         |
| [file export doc](docs/commands/file-export-doc.md)                             | Export document                                               |
| [file export url](docs/commands/file-export-url.md)                             | Export a document from the URL                                |
| [file import batch url](docs/commands/file-import-batch-url.md)                 | Batch import files from URL                                   |
| [file import url](docs/commands/file-import-url.md)                             | Import file from the URL                                      |
| [file info](docs/commands/file-info.md)                                         | Resolve metadata of the path                                  |
| [file list](docs/commands/file-list.md)                                         | List files and folders                                        |
| [file lock acquire](docs/commands/file-lock-acquire.md)                         | Lock a file                                                   |
| [file lock all release](docs/commands/file-lock-all-release.md)                 | Release all locks under the specified path                    |
| [file lock batch acquire](docs/commands/file-lock-batch-acquire.md)             | Lock multiple files                                           |
| [file lock batch release](docs/commands/file-lock-batch-release.md)             | Release multiple locks                                        |
| [file lock list](docs/commands/file-lock-list.md)                               | List locks under the specified path                           |
| [file lock release](docs/commands/file-lock-release.md)                         | Release a lock                                                |
| [file merge](docs/commands/file-merge.md)                                       | Merge paths                                                   |
| [file move](docs/commands/file-move.md)                                         | Move files                                                    |
| [file paper append](docs/commands/file-paper-append.md)                         | Append the content to the end of the existing Paper doc       |
| [file paper create](docs/commands/file-paper-create.md)                         | Create new Paper in the path                                  |
| [file paper overwrite](docs/commands/file-paper-overwrite.md)                   | Overwrite existing Paper document                             |
| [file paper prepend](docs/commands/file-paper-prepend.md)                       | Append the content to the beginning of the existing Paper doc |
| [file replication](docs/commands/file-replication.md)                           | Replicate file content to the other account                   |
| [file restore all](docs/commands/file-restore-all.md)                           | Restore files under given path                                |
| [file revision download](docs/commands/file-revision-download.md)               | Download the file revision                                    |
| [file revision list](docs/commands/file-revision-list.md)                       | List file revisions                                           |
| [file revision restore](docs/commands/file-revision-restore.md)                 | Restore the file revision                                     |
| [file search content](docs/commands/file-search-content.md)                     | Search file content                                           |
| [file search name](docs/commands/file-search-name.md)                           | Search file name                                              |
| [file share info](docs/commands/file-share-info.md)                             | Retrieve sharing information of the file                      |
| [file size](docs/commands/file-size.md)                                         | Storage usage                                                 |
| [file sync down](docs/commands/file-sync-down.md)                               | Downstream sync with Dropbox                                  |
| [file sync online](docs/commands/file-sync-online.md)                           | Sync online files                                             |
| [file sync up](docs/commands/file-sync-up.md)                                   | Upstream sync with Dropbox                                    |
| [file tag add](docs/commands/file-tag-add.md)                                   | Add a tag to the file/folder                                  |
| [file tag delete](docs/commands/file-tag-delete.md)                             | Delete a tag from the file/folder                             |
| [file tag list](docs/commands/file-tag-list.md)                                 | List tags of the path                                         |
| [file template apply remote](docs/commands/file-template-apply-remote.md)       | Apply file/folder structure template to the Dropbox path      |
| [file template capture remote](docs/commands/file-template-capture-remote.md)   | Capture file/folder structure as template from Dropbox path   |
| [file watch](docs/commands/file-watch.md)                                       | Watch file activities                                         |
| [filerequest create](docs/commands/filerequest-create.md)                       | Create a file request                                         |
| [filerequest delete closed](docs/commands/filerequest-delete-closed.md)         | Delete all closed file requests on this account.              |
| [filerequest delete url](docs/commands/filerequest-delete-url.md)               | Delete a file request by the file request URL                 |
| [filerequest list](docs/commands/filerequest-list.md)                           | List file requests of the individual account                  |
| [job history ship](docs/commands/job-history-ship.md)                           | Ship Job logs to Dropbox path                                 |
| [services dropbox user feature](docs/commands/services-dropbox-user-feature.md) | List feature settings for current user                        |
| [services dropbox user info](docs/commands/services-dropbox-user-info.md)       | Retrieve current account info                                 |
| [sharedfolder leave](docs/commands/sharedfolder-leave.md)                       | Leave from the shared folder                                  |
| [sharedfolder list](docs/commands/sharedfolder-list.md)                         | List shared folder(s)                                         |
| [sharedfolder member add](docs/commands/sharedfolder-member-add.md)             | Add a member to the shared folder                             |
| [sharedfolder member delete](docs/commands/sharedfolder-member-delete.md)       | Delete a member from the shared folder                        |
| [sharedfolder member list](docs/commands/sharedfolder-member-list.md)           | List shared folder member(s)                                  |
| [sharedfolder mount add](docs/commands/sharedfolder-mount-add.md)               | Add the shared folder to the current user's Dropbox           |
| [sharedfolder mount delete](docs/commands/sharedfolder-mount-delete.md)         | The current user unmounts the designated folder.              |
| [sharedfolder mount list](docs/commands/sharedfolder-mount-list.md)             | List all shared folders the current user mounted              |
| [sharedfolder mount mountable](docs/commands/sharedfolder-mount-mountable.md)   | List all shared folders the current user can mount            |
| [sharedfolder share](docs/commands/sharedfolder-share.md)                       | Share a folder                                                |
| [sharedfolder unshare](docs/commands/sharedfolder-unshare.md)                   | Unshare a folder                                              |
| [sharedlink create](docs/commands/sharedlink-create.md)                         | Create shared link                                            |
| [sharedlink delete](docs/commands/sharedlink-delete.md)                         | Remove shared links                                           |
| [sharedlink file list](docs/commands/sharedlink-file-list.md)                   | List files for the shared link                                |
| [sharedlink info](docs/commands/sharedlink-info.md)                             | Get information about the shared link                         |
| [sharedlink list](docs/commands/sharedlink-list.md)                             | List of shared link(s)                                        |
| [teamspace file list](docs/commands/teamspace-file-list.md)                     | List files and folders in team space                          |
| [util monitor client](docs/commands/util-monitor-client.md)                     | Start device monitor client                                   |
| [util tidy pack remote](docs/commands/util-tidy-pack-remote.md)                 | Package remote folder into the zip file                       |

## Dropbox Business

| Command                                                                                                     | Description                                                                         |
|-------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| [group add](docs/commands/group-add.md)                                                                     | Create new group                                                                    |
| [group batch add](docs/commands/group-batch-add.md)                                                         | Bulk adding groups                                                                  |
| [group batch delete](docs/commands/group-batch-delete.md)                                                   | Delete groups                                                                       |
| [group clear externalid](docs/commands/group-clear-externalid.md)                                           | Clear an external ID of a group                                                     |
| [group delete](docs/commands/group-delete.md)                                                               | Delete group                                                                        |
| [group folder list](docs/commands/group-folder-list.md)                                                     | List folders of each group                                                          |
| [group list](docs/commands/group-list.md)                                                                   | List group(s)                                                                       |
| [group member add](docs/commands/group-member-add.md)                                                       | Add a member to the group                                                           |
| [group member batch add](docs/commands/group-member-batch-add.md)                                           | Bulk add members into groups                                                        |
| [group member batch delete](docs/commands/group-member-batch-delete.md)                                     | Delete members from groups                                                          |
| [group member batch update](docs/commands/group-member-batch-update.md)                                     | Add or delete members from groups                                                   |
| [group member delete](docs/commands/group-member-delete.md)                                                 | Delete a member from the group                                                      |
| [group member list](docs/commands/group-member-list.md)                                                     | List members of groups                                                              |
| [group rename](docs/commands/group-rename.md)                                                               | Rename the group                                                                    |
| [group update type](docs/commands/group-update-type.md)                                                     | Update group management type                                                        |
| [member batch suspend](docs/commands/member-batch-suspend.md)                                               | Bulk suspend members                                                                |
| [member batch unsuspend](docs/commands/member-batch-unsuspend.md)                                           | Bulk unsuspend members                                                              |
| [member clear externalid](docs/commands/member-clear-externalid.md)                                         | Clear external_id of members                                                        |
| [member delete](docs/commands/member-delete.md)                                                             | Delete members                                                                      |
| [member detach](docs/commands/member-detach.md)                                                             | Convert Dropbox Business accounts to a Basic account                                |
| [member feature](docs/commands/member-feature.md)                                                           | List member feature settings                                                        |
| [member file lock all release](docs/commands/member-file-lock-all-release.md)                               | Release all locks under the path of the member                                      |
| [member file lock list](docs/commands/member-file-lock-list.md)                                             | List locks of the member under the path                                             |
| [member file lock release](docs/commands/member-file-lock-release.md)                                       | Release the lock of the path as the member                                          |
| [member file permdelete](docs/commands/member-file-permdelete.md)                                           | Permanently delete the file or folder at a given path of the team member.           |
| [member folder list](docs/commands/member-folder-list.md)                                                   | List folders for each member                                                        |
| [member folder replication](docs/commands/member-folder-replication.md)                                     | Replicate a folder to another member's personal folder                              |
| [member invite](docs/commands/member-invite.md)                                                             | Invite member(s)                                                                    |
| [member list](docs/commands/member-list.md)                                                                 | List team member(s)                                                                 |
| [member quota list](docs/commands/member-quota-list.md)                                                     | List team member quota                                                              |
| [member quota update](docs/commands/member-quota-update.md)                                                 | Update team member quota                                                            |
| [member quota usage](docs/commands/member-quota-usage.md)                                                   | List team member storage usage                                                      |
| [member reinvite](docs/commands/member-reinvite.md)                                                         | Reinvite invited status members to the team                                         |
| [member replication](docs/commands/member-replication.md)                                                   | Replicate team member files                                                         |
| [member suspend](docs/commands/member-suspend.md)                                                           | Suspend a member                                                                    |
| [member unsuspend](docs/commands/member-unsuspend.md)                                                       | Unsuspend a member                                                                  |
| [member update email](docs/commands/member-update-email.md)                                                 | Member email operation                                                              |
| [member update externalid](docs/commands/member-update-externalid.md)                                       | Update External ID of team members                                                  |
| [member update invisible](docs/commands/member-update-invisible.md)                                         | Enable directory restriction to members                                             |
| [member update profile](docs/commands/member-update-profile.md)                                             | Update member profile                                                               |
| [member update visible](docs/commands/member-update-visible.md)                                             | Disable directory restriction to members                                            |
| [team activity batch user](docs/commands/team-activity-batch-user.md)                                       | Scan activities for multiple users                                                  |
| [team activity daily event](docs/commands/team-activity-daily-event.md)                                     | Report activities by day                                                            |
| [team activity event](docs/commands/team-activity-event.md)                                                 | Event log                                                                           |
| [team activity user](docs/commands/team-activity-user.md)                                                   | Activities log per user                                                             |
| [team admin group role add](docs/commands/team-admin-group-role-add.md)                                     | Add the role to members of the group                                                |
| [team admin group role delete](docs/commands/team-admin-group-role-delete.md)                               | Delete the role from all members except of members of the exception group           |
| [team admin list](docs/commands/team-admin-list.md)                                                         | List admin roles of members                                                         |
| [team admin role add](docs/commands/team-admin-role-add.md)                                                 | Add a new role to the member                                                        |
| [team admin role clear](docs/commands/team-admin-role-clear.md)                                             | Remove all admin roles from the member                                              |
| [team admin role delete](docs/commands/team-admin-role-delete.md)                                           | Remove a role from the member                                                       |
| [team admin role list](docs/commands/team-admin-role-list.md)                                               | List admin roles of the team                                                        |
| [team content legacypaper count](docs/commands/team-content-legacypaper-count.md)                           | Count number of Paper documents per member                                          |
| [team content legacypaper export](docs/commands/team-content-legacypaper-export.md)                         | Export entire team member Paper documents into local path                           |
| [team content legacypaper list](docs/commands/team-content-legacypaper-list.md)                             | List team member Paper documents                                                    |
| [team content member list](docs/commands/team-content-member-list.md)                                       | List team folder & shared folder members                                            |
| [team content member size](docs/commands/team-content-member-size.md)                                       | Count number of members of team folders and shared folders                          |
| [team content mount list](docs/commands/team-content-mount-list.md)                                         | List all mounted/unmounted shared folders of team members.                          |
| [team content policy list](docs/commands/team-content-policy-list.md)                                       | List policies of team folders and shared folders in the team                        |
| [team device list](docs/commands/team-device-list.md)                                                       | List all devices/sessions in the team                                               |
| [team device unlink](docs/commands/team-device-unlink.md)                                                   | Unlink device sessions                                                              |
| [team feature](docs/commands/team-feature.md)                                                               | Team feature                                                                        |
| [team filerequest list](docs/commands/team-filerequest-list.md)                                             | List all file requests in the team                                                  |
| [team info](docs/commands/team-info.md)                                                                     | Team information                                                                    |
| [team legalhold add](docs/commands/team-legalhold-add.md)                                                   | Creates new legal hold policy.                                                      |
| [team legalhold list](docs/commands/team-legalhold-list.md)                                                 | Retrieve existing policies                                                          |
| [team legalhold member batch update](docs/commands/team-legalhold-member-batch-update.md)                   | Update member list of legal hold policy                                             |
| [team legalhold member list](docs/commands/team-legalhold-member-list.md)                                   | List members of the legal hold                                                      |
| [team legalhold release](docs/commands/team-legalhold-release.md)                                           | Releases a legal hold by Id                                                         |
| [team legalhold revision list](docs/commands/team-legalhold-revision-list.md)                               | List revisions of the legal hold policy                                             |
| [team legalhold update desc](docs/commands/team-legalhold-update-desc.md)                                   | Update description of the legal hold policy                                         |
| [team legalhold update name](docs/commands/team-legalhold-update-name.md)                                   | Update name of the legal hold policy                                                |
| [team linkedapp list](docs/commands/team-linkedapp-list.md)                                                 | List linked applications                                                            |
| [team namespace file list](docs/commands/team-namespace-file-list.md)                                       | List all files and folders of the team namespaces                                   |
| [team namespace file size](docs/commands/team-namespace-file-size.md)                                       | List all files and folders of the team namespaces                                   |
| [team namespace list](docs/commands/team-namespace-list.md)                                                 | List all namespaces of the team                                                     |
| [team namespace member list](docs/commands/team-namespace-member-list.md)                                   | List members of shared folders and team folders in the team                         |
| [team namespace summary](docs/commands/team-namespace-summary.md)                                           | Report team namespace status summary.                                               |
| [team runas file batch copy](docs/commands/team-runas-file-batch-copy.md)                                   | Batch copy files/folders as a member                                                |
| [team runas file list](docs/commands/team-runas-file-list.md)                                               | List files and folders run as a member                                              |
| [team runas file sync batch up](docs/commands/team-runas-file-sync-batch-up.md)                             | Batch sync up that run as members                                                   |
| [team runas sharedfolder batch leave](docs/commands/team-runas-sharedfolder-batch-leave.md)                 | Batch leave from shared folders as a member                                         |
| [team runas sharedfolder batch share](docs/commands/team-runas-sharedfolder-batch-share.md)                 | Batch share folders for members                                                     |
| [team runas sharedfolder batch unshare](docs/commands/team-runas-sharedfolder-batch-unshare.md)             | Batch unshare folders for members                                                   |
| [team runas sharedfolder isolate](docs/commands/team-runas-sharedfolder-isolate.md)                         | Unshare owned shared folders and leave from external shared folders run as a member |
| [team runas sharedfolder list](docs/commands/team-runas-sharedfolder-list.md)                               | List shared folders run as the member                                               |
| [team runas sharedfolder member batch add](docs/commands/team-runas-sharedfolder-member-batch-add.md)       | Batch add members to member's shared folders                                        |
| [team runas sharedfolder member batch delete](docs/commands/team-runas-sharedfolder-member-batch-delete.md) | Batch delete members from member's shared folders                                   |
| [team runas sharedfolder mount add](docs/commands/team-runas-sharedfolder-mount-add.md)                     | Add the shared folder to the specified member's Dropbox                             |
| [team runas sharedfolder mount delete](docs/commands/team-runas-sharedfolder-mount-delete.md)               | The specified user unmounts the designated folder.                                  |
| [team runas sharedfolder mount list](docs/commands/team-runas-sharedfolder-mount-list.md)                   | List all shared folders the specified member mounted                                |
| [team runas sharedfolder mount mountable](docs/commands/team-runas-sharedfolder-mount-mountable.md)         | List all shared folders the member can mount                                        |
| [team sharedlink cap expiry](docs/commands/team-sharedlink-cap-expiry.md)                                   | Set expiry cap to shared links in the team                                          |
| [team sharedlink cap visibility](docs/commands/team-sharedlink-cap-visibility.md)                           | Set visibility cap to shared links in the team                                      |
| [team sharedlink delete links](docs/commands/team-sharedlink-delete-links.md)                               | Batch delete shared links                                                           |
| [team sharedlink delete member](docs/commands/team-sharedlink-delete-member.md)                             | Delete all shared links of the member                                               |
| [team sharedlink list](docs/commands/team-sharedlink-list.md)                                               | List of shared links                                                                |
| [team sharedlink update expiry](docs/commands/team-sharedlink-update-expiry.md)                             | Update expiration date of public shared links within the team                       |
| [team sharedlink update password](docs/commands/team-sharedlink-update-password.md)                         | Set or update shared link passwords                                                 |
| [team sharedlink update visibility](docs/commands/team-sharedlink-update-visibility.md)                     | Update visibility of shared links                                                   |
| [teamfolder add](docs/commands/teamfolder-add.md)                                                           | Add team folder to the team                                                         |
| [teamfolder archive](docs/commands/teamfolder-archive.md)                                                   | Archive team folder                                                                 |
| [teamfolder batch archive](docs/commands/teamfolder-batch-archive.md)                                       | Archiving team folders                                                              |
| [teamfolder batch permdelete](docs/commands/teamfolder-batch-permdelete.md)                                 | Permanently delete team folders                                                     |
| [teamfolder batch replication](docs/commands/teamfolder-batch-replication.md)                               | Batch replication of team folders                                                   |
| [teamfolder file list](docs/commands/teamfolder-file-list.md)                                               | List files in team folders                                                          |
| [teamfolder file lock all release](docs/commands/teamfolder-file-lock-all-release.md)                       | Release all locks under the path of the team folder                                 |
| [teamfolder file lock list](docs/commands/teamfolder-file-lock-list.md)                                     | List locks in the team folder                                                       |
| [teamfolder file lock release](docs/commands/teamfolder-file-lock-release.md)                               | Release lock of the path in the team folder                                         |
| [teamfolder file size](docs/commands/teamfolder-file-size.md)                                               | Calculate size of team folders                                                      |
| [teamfolder list](docs/commands/teamfolder-list.md)                                                         | List team folder(s)                                                                 |
| [teamfolder member add](docs/commands/teamfolder-member-add.md)                                             | Batch adding users/groups to team folders                                           |
| [teamfolder member delete](docs/commands/teamfolder-member-delete.md)                                       | Batch removing users/groups from team folders                                       |
| [teamfolder member list](docs/commands/teamfolder-member-list.md)                                           | List team folder members                                                            |
| [teamfolder partial replication](docs/commands/teamfolder-partial-replication.md)                           | Partial team folder replication to the other team                                   |
| [teamfolder permdelete](docs/commands/teamfolder-permdelete.md)                                             | Permanently delete team folder                                                      |
| [teamfolder policy list](docs/commands/teamfolder-policy-list.md)                                           | List policies of team folders                                                       |
| [teamfolder replication](docs/commands/teamfolder-replication.md)                                           | Replicate a team folder to the other team                                           |
| [teamfolder sync setting list](docs/commands/teamfolder-sync-setting-list.md)                               | List team folder sync settings                                                      |
| [teamfolder sync setting update](docs/commands/teamfolder-sync-setting-update.md)                           | Batch update team folder sync settings                                              |
| [teamspace asadmin file list](docs/commands/teamspace-asadmin-file-list.md)                                 | List files and folders in team space run as admin                                   |
| [teamspace asadmin folder add](docs/commands/teamspace-asadmin-folder-add.md)                               | Create top level folder in the team space                                           |
| [teamspace asadmin folder delete](docs/commands/teamspace-asadmin-folder-delete.md)                         | Delete top level folder of the team space                                           |
| [teamspace asadmin folder permdelete](docs/commands/teamspace-asadmin-folder-permdelete.md)                 | Permanently delete top level folder of the team space                               |

## Figma

| Command                                                                                     | Description                           |
|---------------------------------------------------------------------------------------------|---------------------------------------|
| [services figma account info](docs/commands/services-figma-account-info.md)                 | Retrieve current user information     |
| [services figma file export all page](docs/commands/services-figma-file-export-all-page.md) | Export all files/pages under the team |
| [services figma file export frame](docs/commands/services-figma-file-export-frame.md)       | Export all frames of the Figma file   |
| [services figma file export node](docs/commands/services-figma-file-export-node.md)         | Export Figma document Node            |
| [services figma file export page](docs/commands/services-figma-file-export-page.md)         | Export all pages of the Figma file    |
| [services figma file info](docs/commands/services-figma-file-info.md)                       | Show information of the figma file    |
| [services figma file list](docs/commands/services-figma-file-list.md)                       | List files in the Figma Project       |
| [services figma project list](docs/commands/services-figma-project-list.md)                 | List projects of the team             |

## GitHub

| Command                                                                                           | Description                                         |
|---------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| [services github content get](docs/commands/services-github-content-get.md)                       | Get content metadata of the repository              |
| [services github content put](docs/commands/services-github-content-put.md)                       | Put small text content into the repository          |
| [services github issue list](docs/commands/services-github-issue-list.md)                         | List issues of the public/private GitHub repository |
| [services github profile](docs/commands/services-github-profile.md)                               | Get the authenticated user                          |
| [services github release asset download](docs/commands/services-github-release-asset-download.md) | Download assets                                     |
| [services github release asset list](docs/commands/services-github-release-asset-list.md)         | List assets of GitHub Release                       |
| [services github release asset upload](docs/commands/services-github-release-asset-upload.md)     | Upload assets file into the GitHub Release          |
| [services github release draft](docs/commands/services-github-release-draft.md)                   | Create release draft                                |
| [services github release list](docs/commands/services-github-release-list.md)                     | List releases                                       |
| [services github tag create](docs/commands/services-github-tag-create.md)                         | Create a tag on the repository                      |
| [util release install](docs/commands/util-release-install.md)                                     | Download & install watermint toolbox to the path    |

## Google Calendar

| Command                                                                                     | Description                 |
|---------------------------------------------------------------------------------------------|-----------------------------|
| [services google calendar event list](docs/commands/services-google-calendar-event-list.md) | List Google Calendar events |

## Google Gmail

| Command                                                                                                     | Description                                         |
|-------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| [services google mail filter add](docs/commands/services-google-mail-filter-add.md)                         | Add a filter.                                       |
| [services google mail filter batch add](docs/commands/services-google-mail-filter-batch-add.md)             | Batch adding/deleting labels with query             |
| [services google mail filter delete](docs/commands/services-google-mail-filter-delete.md)                   | Delete a filter                                     |
| [services google mail filter list](docs/commands/services-google-mail-filter-list.md)                       | List filters                                        |
| [services google mail label add](docs/commands/services-google-mail-label-add.md)                           | Add a label                                         |
| [services google mail label delete](docs/commands/services-google-mail-label-delete.md)                     | Delete a label                                      |
| [services google mail label list](docs/commands/services-google-mail-label-list.md)                         | List email labels                                   |
| [services google mail label rename](docs/commands/services-google-mail-label-rename.md)                     | Rename a label                                      |
| [services google mail message label add](docs/commands/services-google-mail-message-label-add.md)           | Add labels to the message                           |
| [services google mail message label delete](docs/commands/services-google-mail-message-label-delete.md)     | Remove labels from the message                      |
| [services google mail message list](docs/commands/services-google-mail-message-list.md)                     | List messages                                       |
| [services google mail message processed list](docs/commands/services-google-mail-message-processed-list.md) | List messages in processed format.                  |
| [services google mail sendas add](docs/commands/services-google-mail-sendas-add.md)                         | Creates a custom "from" send-as alias               |
| [services google mail sendas delete](docs/commands/services-google-mail-sendas-delete.md)                   | Deletes the specified send-as alias                 |
| [services google mail sendas list](docs/commands/services-google-mail-sendas-list.md)                       | Lists the send-as aliases for the specified account |
| [services google mail thread list](docs/commands/services-google-mail-thread-list.md)                       | List threads                                        |

## Google Sheets

| Command                                                                                                 | Description                         |
|---------------------------------------------------------------------------------------------------------|-------------------------------------|
| [services google sheets sheet append](docs/commands/services-google-sheets-sheet-append.md)             | Append data to a spreadsheet        |
| [services google sheets sheet clear](docs/commands/services-google-sheets-sheet-clear.md)               | Clears values from a spreadsheet    |
| [services google sheets sheet create](docs/commands/services-google-sheets-sheet-create.md)             | Create a new sheet                  |
| [services google sheets sheet delete](docs/commands/services-google-sheets-sheet-delete.md)             | Delete a sheet from the spreadsheet |
| [services google sheets sheet export](docs/commands/services-google-sheets-sheet-export.md)             | Export sheet data                   |
| [services google sheets sheet import](docs/commands/services-google-sheets-sheet-import.md)             | Import data into the spreadsheet    |
| [services google sheets sheet list](docs/commands/services-google-sheets-sheet-list.md)                 | List sheets of the spreadsheet      |
| [services google sheets spreadsheet create](docs/commands/services-google-sheets-spreadsheet-create.md) | Create a new spreadsheet            |

## Utilities

| Command                                                                     | Description                                                          |
|-----------------------------------------------------------------------------|----------------------------------------------------------------------|
| [config auth delete](docs/commands/config-auth-delete.md)                   | Delete existing auth credential                                      |
| [config auth list](docs/commands/config-auth-list.md)                       | List all auth credentials                                            |
| [config disable](docs/commands/config-disable.md)                           | Disable a feature.                                                   |
| [config enable](docs/commands/config-enable.md)                             | Enable a feature.                                                    |
| [config features](docs/commands/config-features.md)                         | List available optional features.                                    |
| [file template apply local](docs/commands/file-template-apply-local.md)     | Apply file/folder structure template to the local path               |
| [file template capture local](docs/commands/file-template-capture-local.md) | Capture file/folder structure as template from local path            |
| [job history archive](docs/commands/job-history-archive.md)                 | Archive jobs                                                         |
| [job history delete](docs/commands/job-history-delete.md)                   | Delete old job history                                               |
| [job history list](docs/commands/job-history-list.md)                       | Show job history                                                     |
| [job log jobid](docs/commands/job-log-jobid.md)                             | Retrieve logs of specified Job ID                                    |
| [job log kind](docs/commands/job-log-kind.md)                               | Concatenate and print logs of specified log kind                     |
| [job log last](docs/commands/job-log-last.md)                               | Print the last job log files                                         |
| [license](docs/commands/license.md)                                         | Show license information                                             |
| [util archive unzip](docs/commands/util-archive-unzip.md)                   | Extract the zip archive file                                         |
| [util archive zip](docs/commands/util-archive-zip.md)                       | Compress target files into the zip archive                           |
| [util cert selfsigned](docs/commands/util-cert-selfsigned.md)               | Generate self-signed certificate and key                             |
| [util database exec](docs/commands/util-database-exec.md)                   | Execute query on SQLite3 database file                               |
| [util database query](docs/commands/util-database-query.md)                 | Query SQLite3 database                                               |
| [util date today](docs/commands/util-date-today.md)                         | Display current date                                                 |
| [util datetime now](docs/commands/util-datetime-now.md)                     | Display current date/time                                            |
| [util decode base32](docs/commands/util-decode-base32.md)                   | Decode text from Base32 (RFC 4648) format                            |
| [util decode base64](docs/commands/util-decode-base64.md)                   | Decode text from Base64 (RFC 4648) format                            |
| [util encode base32](docs/commands/util-encode-base32.md)                   | Encode text into Base32 (RFC 4648) format                            |
| [util encode base64](docs/commands/util-encode-base64.md)                   | Encode text into Base64 (RFC 4648) format                            |
| [util file hash](docs/commands/util-file-hash.md)                           | Print file digest                                                    |
| [util git clone](docs/commands/util-git-clone.md)                           | Clone git repository                                                 |
| [util image exif](docs/commands/util-image-exif.md)                         | Print EXIF metadata of image file                                    |
| [util image placeholder](docs/commands/util-image-placeholder.md)           | Create placeholder image                                             |
| [util net download](docs/commands/util-net-download.md)                     | Download a file                                                      |
| [util qrcode create](docs/commands/util-qrcode-create.md)                   | Create a QR code image file                                          |
| [util qrcode wifi](docs/commands/util-qrcode-wifi.md)                       | Generate QR code for WIFI configuration                              |
| [util table format xlsx](docs/commands/util-table-format-xlsx.md)           | Formatting xlsx file into text                                       |
| [util text case down](docs/commands/util-text-case-down.md)                 | Print lower case text                                                |
| [util text case up](docs/commands/util-text-case-up.md)                     | Print upper case text                                                |
| [util text encoding from](docs/commands/util-text-encoding-from.md)         | Convert text encoding to UTF-8 text file from specified encoding.    |
| [util text encoding to](docs/commands/util-text-encoding-to.md)             | Convert text encoding to specified encoding from UTF-8 text file.    |
| [util tidy move dispatch](docs/commands/util-tidy-move-dispatch.md)         | Dispatch files                                                       |
| [util tidy move simple](docs/commands/util-tidy-move-simple.md)             | Archive local files                                                  |
| [util time now](docs/commands/util-time-now.md)                             | Display current time                                                 |
| [util unixtime format](docs/commands/util-unixtime-format.md)               | Time format to convert the unix time (epoch seconds from 1970-01-01) |
| [util unixtime now](docs/commands/util-unixtime-now.md)                     | Display current time in unixtime                                     |
| [util xlsx create](docs/commands/util-xlsx-create.md)                       | Create an empty spreadsheet                                          |
| [util xlsx sheet export](docs/commands/util-xlsx-sheet-export.md)           | Export data from the xlsx file                                       |
| [util xlsx sheet import](docs/commands/util-xlsx-sheet-import.md)           | Import data into xlsx file                                           |
| [util xlsx sheet list](docs/commands/util-xlsx-sheet-list.md)               | List sheets of the xlsx file                                         |
| [version](docs/commands/version.md)                                         | Show version                                                         |

