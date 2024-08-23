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

This product itself is experimental and is not subject to the maintained to keep quality of service. The project will try to fix critical bugs and security issues with the best effort. But that is also not guaranteed.

The product will not release any patch release of a certain major releases. The product will apply fixes as next release when those fixes accepted to do.

## Specification changes

The deliverables of this project are stand-alone executable programs. The specification changes will not be applied unless you explicitly upgrade your version of the program.

The following policy will be used to make changes in new version releases.

Command paths, arguments, return values, etc. will be upgraded to be as compatible as possible, but may be discontinued or changed.Addition of arguments, etc.
The general policy is as follows.

* Changes that do not break existing behavior, such as the addition of arguments or changes to messages, will be implemented without notice.
* Commands that are considered infrequently used will be discontinued or moved without notice.
* Changes to other commands will be announced 30-180 days or more in advance.

Changes in specifications will be announced at [Announcements](https://github.com/watermint/toolbox/discussions/categories/announcements). Please refer to [Specification Change](https://toolbox.watermint.org/guides/spec-change.html) for a list of planned specification changes.

## Availability period for each release

In general, new security issues are discovered every day. In order not to leave these security and critical issues unaddressed by continuing to use older watermint toolbox releases, a maximum availability period has been set for release 130 and above. Please see [#815](https://github.com/watermint/toolbox/discussions/815) for more details.

# Announcements

* [#870 Deprecation: util monitor client](https://github.com/watermint/toolbox/discussions/870)
* [#868 Removal of the screenshot commands](https://github.com/watermint/toolbox/discussions/868)
* [#835 Google commands deprecation](https://github.com/watermint/toolbox/discussions/835)
* [#836 Remove binaries that are more than six months old after release](https://github.com/watermint/toolbox/discussions/836)
* [#815 Lifecycle: Availability period for each release](https://github.com/watermint/toolbox/discussions/815)

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

`tbx` have various features. Run without an option for a list of supported commands and options.
You can see available commands and options by running executable without arguments like below.

```
% ./tbx

watermint toolbox xx.x.xxx
==========================

© 2016-2024 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox for teams

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
| [dropbox file sharedfolder leave](docs/commands/dropbox-file-sharedfolder-leave.md)                     | Leave from the shared folder                                  |
| [dropbox file sharedfolder list](docs/commands/dropbox-file-sharedfolder-list.md)                       | List shared folder(s)                                         |
| [dropbox file sharedfolder member add](docs/commands/dropbox-file-sharedfolder-member-add.md)           | Add a member to the shared folder                             |
| [dropbox file sharedfolder member delete](docs/commands/dropbox-file-sharedfolder-member-delete.md)     | Delete a member from the shared folder                        |
| [dropbox file sharedfolder member list](docs/commands/dropbox-file-sharedfolder-member-list.md)         | List shared folder member(s)                                  |
| [dropbox file sharedfolder mount add](docs/commands/dropbox-file-sharedfolder-mount-add.md)             | Add the shared folder to the current user's Dropbox           |
| [dropbox file sharedfolder mount delete](docs/commands/dropbox-file-sharedfolder-mount-delete.md)       | The current user unmounts the designated folder.              |
| [dropbox file sharedfolder mount list](docs/commands/dropbox-file-sharedfolder-mount-list.md)           | List all shared folders the current user mounted              |
| [dropbox file sharedfolder mount mountable](docs/commands/dropbox-file-sharedfolder-mount-mountable.md) | List all shared folders the current user can mount            |
| [dropbox file sharedfolder share](docs/commands/dropbox-file-sharedfolder-share.md)                     | Share a folder                                                |
| [dropbox file sharedfolder unshare](docs/commands/dropbox-file-sharedfolder-unshare.md)                 | Unshare a folder                                              |
| [dropbox file sharedlink create](docs/commands/dropbox-file-sharedlink-create.md)                       | Create shared link                                            |
| [dropbox file sharedlink delete](docs/commands/dropbox-file-sharedlink-delete.md)                       | Remove shared links                                           |
| [dropbox file sharedlink file list](docs/commands/dropbox-file-sharedlink-file-list.md)                 | List files for the shared link                                |
| [dropbox file sharedlink info](docs/commands/dropbox-file-sharedlink-info.md)                           | Get information about the shared link                         |
| [dropbox file sharedlink list](docs/commands/dropbox-file-sharedlink-list.md)                           | List of shared link(s)                                        |
| [dropbox file size](docs/commands/dropbox-file-size.md)                                                 | Storage usage                                                 |
| [dropbox file sync down](docs/commands/dropbox-file-sync-down.md)                                       | Downstream sync with Dropbox                                  |
| [dropbox file sync online](docs/commands/dropbox-file-sync-online.md)                                   | Sync online files                                             |
| [dropbox file sync up](docs/commands/dropbox-file-sync-up.md)                                           | Upstream sync with Dropbox                                    |
| [dropbox file tag add](docs/commands/dropbox-file-tag-add.md)                                           | Add a tag to the file/folder                                  |
| [dropbox file tag delete](docs/commands/dropbox-file-tag-delete.md)                                     | Delete a tag from the file/folder                             |
| [dropbox file tag list](docs/commands/dropbox-file-tag-list.md)                                         | List tags of the path                                         |
| [dropbox file template apply](docs/commands/dropbox-file-template-apply.md)                             | Apply file/folder structure template to the Dropbox path      |
| [dropbox file template capture](docs/commands/dropbox-file-template-capture.md)                         | Capture file/folder structure as template from Dropbox path   |
| [dropbox file watch](docs/commands/dropbox-file-watch.md)                                               | Watch file activities                                         |
| [dropbox paper append](docs/commands/dropbox-paper-append.md)                                           | Append the content to the end of the existing Paper doc       |
| [dropbox paper create](docs/commands/dropbox-paper-create.md)                                           | Create new Paper in the path                                  |
| [dropbox paper overwrite](docs/commands/dropbox-paper-overwrite.md)                                     | Overwrite existing Paper document                             |
| [dropbox paper prepend](docs/commands/dropbox-paper-prepend.md)                                         | Append the content to the beginning of the existing Paper doc |
| [util monitor client](docs/commands/util-monitor-client.md)                                             | Start device monitor client                                   |
| [util tidy pack remote](docs/commands/util-tidy-pack-remote.md)                                         | Package remote folder into the zip file                       |

## Dropbox Sign

| Command                                                                                     | Description                 |
|---------------------------------------------------------------------------------------------|-----------------------------|
| [dropbox sign request list](docs/commands/dropbox-sign-request-list.md)                     | List signature requests     |
| [dropbox sign request signature list](docs/commands/dropbox-sign-request-signature-list.md) | List signatures of requests |

## Dropbox for teams

| Command                                                                                                                     | Description                                                                         |
|-----------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| [dropbox team activity batch user](docs/commands/dropbox-team-activity-batch-user.md)                                       | Scan activities for multiple users                                                  |
| [dropbox team activity daily event](docs/commands/dropbox-team-activity-daily-event.md)                                     | Report activities by day                                                            |
| [dropbox team activity event](docs/commands/dropbox-team-activity-event.md)                                                 | Event log                                                                           |
| [dropbox team activity user](docs/commands/dropbox-team-activity-user.md)                                                   | Activities log per user                                                             |
| [dropbox team admin group role add](docs/commands/dropbox-team-admin-group-role-add.md)                                     | Add the role to members of the group                                                |
| [dropbox team admin group role delete](docs/commands/dropbox-team-admin-group-role-delete.md)                               | Delete the role from all members except of members of the exception group           |
| [dropbox team admin list](docs/commands/dropbox-team-admin-list.md)                                                         | List admin roles of members                                                         |
| [dropbox team admin role add](docs/commands/dropbox-team-admin-role-add.md)                                                 | Add a new role to the member                                                        |
| [dropbox team admin role clear](docs/commands/dropbox-team-admin-role-clear.md)                                             | Remove all admin roles from the member                                              |
| [dropbox team admin role delete](docs/commands/dropbox-team-admin-role-delete.md)                                           | Remove a role from the member                                                       |
| [dropbox team admin role list](docs/commands/dropbox-team-admin-role-list.md)                                               | List admin roles of the team                                                        |
| [dropbox team backup device status](docs/commands/dropbox-team-backup-device-status.md)                                     | Dropbox Backup device status change in the specified period                         |
| [dropbox team content legacypaper count](docs/commands/dropbox-team-content-legacypaper-count.md)                           | Count number of Paper documents per member                                          |
| [dropbox team content legacypaper export](docs/commands/dropbox-team-content-legacypaper-export.md)                         | Export entire team member Paper documents into local path                           |
| [dropbox team content legacypaper list](docs/commands/dropbox-team-content-legacypaper-list.md)                             | List team member Paper documents                                                    |
| [dropbox team content member list](docs/commands/dropbox-team-content-member-list.md)                                       | List team folder & shared folder members                                            |
| [dropbox team content member size](docs/commands/dropbox-team-content-member-size.md)                                       | Count number of members of team folders and shared folders                          |
| [dropbox team content mount list](docs/commands/dropbox-team-content-mount-list.md)                                         | List all mounted/unmounted shared folders of team members.                          |
| [dropbox team content policy list](docs/commands/dropbox-team-content-policy-list.md)                                       | List policies of team folders and shared folders in the team                        |
| [dropbox team device list](docs/commands/dropbox-team-device-list.md)                                                       | List all devices/sessions in the team                                               |
| [dropbox team device unlink](docs/commands/dropbox-team-device-unlink.md)                                                   | Unlink device sessions                                                              |
| [dropbox team feature](docs/commands/dropbox-team-feature.md)                                                               | Team feature                                                                        |
| [dropbox team filerequest list](docs/commands/dropbox-team-filerequest-list.md)                                             | List all file requests in the team                                                  |
| [dropbox team filesystem](docs/commands/dropbox-team-filesystem.md)                                                         | Identify team's file system version                                                 |
| [dropbox team group add](docs/commands/dropbox-team-group-add.md)                                                           | Create new group                                                                    |
| [dropbox team group batch add](docs/commands/dropbox-team-group-batch-add.md)                                               | Bulk adding groups                                                                  |
| [dropbox team group batch delete](docs/commands/dropbox-team-group-batch-delete.md)                                         | Delete groups                                                                       |
| [dropbox team group clear externalid](docs/commands/dropbox-team-group-clear-externalid.md)                                 | Clear an external ID of a group                                                     |
| [dropbox team group delete](docs/commands/dropbox-team-group-delete.md)                                                     | Delete group                                                                        |
| [dropbox team group folder list](docs/commands/dropbox-team-group-folder-list.md)                                           | List folders of each group                                                          |
| [dropbox team group list](docs/commands/dropbox-team-group-list.md)                                                         | List group(s)                                                                       |
| [dropbox team group member add](docs/commands/dropbox-team-group-member-add.md)                                             | Add a member to the group                                                           |
| [dropbox team group member batch add](docs/commands/dropbox-team-group-member-batch-add.md)                                 | Bulk add members into groups                                                        |
| [dropbox team group member batch delete](docs/commands/dropbox-team-group-member-batch-delete.md)                           | Delete members from groups                                                          |
| [dropbox team group member batch update](docs/commands/dropbox-team-group-member-batch-update.md)                           | Add or delete members from groups                                                   |
| [dropbox team group member delete](docs/commands/dropbox-team-group-member-delete.md)                                       | Delete a member from the group                                                      |
| [dropbox team group member list](docs/commands/dropbox-team-group-member-list.md)                                           | List members of groups                                                              |
| [dropbox team group rename](docs/commands/dropbox-team-group-rename.md)                                                     | Rename the group                                                                    |
| [dropbox team group update type](docs/commands/dropbox-team-group-update-type.md)                                           | Update group management type                                                        |
| [dropbox team info](docs/commands/dropbox-team-info.md)                                                                     | Team information                                                                    |
| [dropbox team insight scan](docs/commands/dropbox-team-insight-scan.md)                                                     | Scans team data for analysis                                                        |
| [dropbox team legalhold add](docs/commands/dropbox-team-legalhold-add.md)                                                   | Creates new legal hold policy.                                                      |
| [dropbox team legalhold list](docs/commands/dropbox-team-legalhold-list.md)                                                 | Retrieve existing policies                                                          |
| [dropbox team legalhold member batch update](docs/commands/dropbox-team-legalhold-member-batch-update.md)                   | Update member list of legal hold policy                                             |
| [dropbox team legalhold member list](docs/commands/dropbox-team-legalhold-member-list.md)                                   | List members of the legal hold                                                      |
| [dropbox team legalhold release](docs/commands/dropbox-team-legalhold-release.md)                                           | Releases a legal hold by Id                                                         |
| [dropbox team legalhold revision list](docs/commands/dropbox-team-legalhold-revision-list.md)                               | List revisions of the legal hold policy                                             |
| [dropbox team legalhold update desc](docs/commands/dropbox-team-legalhold-update-desc.md)                                   | Update description of the legal hold policy                                         |
| [dropbox team legalhold update name](docs/commands/dropbox-team-legalhold-update-name.md)                                   | Update name of the legal hold policy                                                |
| [dropbox team linkedapp list](docs/commands/dropbox-team-linkedapp-list.md)                                                 | List linked applications                                                            |
| [dropbox team member batch delete](docs/commands/dropbox-team-member-batch-delete.md)                                       | Delete members                                                                      |
| [dropbox team member batch detach](docs/commands/dropbox-team-member-batch-detach.md)                                       | Convert Dropbox for teams accounts to a Basic account                               |
| [dropbox team member batch invite](docs/commands/dropbox-team-member-batch-invite.md)                                       | Invite member(s)                                                                    |
| [dropbox team member batch reinvite](docs/commands/dropbox-team-member-batch-reinvite.md)                                   | Reinvite invited status members to the team                                         |
| [dropbox team member batch suspend](docs/commands/dropbox-team-member-batch-suspend.md)                                     | Bulk suspend members                                                                |
| [dropbox team member batch unsuspend](docs/commands/dropbox-team-member-batch-unsuspend.md)                                 | Bulk unsuspend members                                                              |
| [dropbox team member clear externalid](docs/commands/dropbox-team-member-clear-externalid.md)                               | Clear external_id of members                                                        |
| [dropbox team member feature](docs/commands/dropbox-team-member-feature.md)                                                 | List member feature settings                                                        |
| [dropbox team member file lock all release](docs/commands/dropbox-team-member-file-lock-all-release.md)                     | Release all locks under the path of the member                                      |
| [dropbox team member file lock list](docs/commands/dropbox-team-member-file-lock-list.md)                                   | List locks of the member under the path                                             |
| [dropbox team member file lock release](docs/commands/dropbox-team-member-file-lock-release.md)                             | Release the lock of the path as the member                                          |
| [dropbox team member file permdelete](docs/commands/dropbox-team-member-file-permdelete.md)                                 | Permanently delete the file or folder at a given path of the team member.           |
| [dropbox team member folder list](docs/commands/dropbox-team-member-folder-list.md)                                         | List folders for each member                                                        |
| [dropbox team member folder replication](docs/commands/dropbox-team-member-folder-replication.md)                           | Replicate a folder to another member's personal folder                              |
| [dropbox team member list](docs/commands/dropbox-team-member-list.md)                                                       | List team member(s)                                                                 |
| [dropbox team member quota batch update](docs/commands/dropbox-team-member-quota-batch-update.md)                           | Update team member quota                                                            |
| [dropbox team member quota list](docs/commands/dropbox-team-member-quota-list.md)                                           | List team member quota                                                              |
| [dropbox team member quota usage](docs/commands/dropbox-team-member-quota-usage.md)                                         | List team member storage usage                                                      |
| [dropbox team member replication](docs/commands/dropbox-team-member-replication.md)                                         | Replicate team member files                                                         |
| [dropbox team member suspend](docs/commands/dropbox-team-member-suspend.md)                                                 | Suspend a member                                                                    |
| [dropbox team member unsuspend](docs/commands/dropbox-team-member-unsuspend.md)                                             | Unsuspend a member                                                                  |
| [dropbox team member update batch email](docs/commands/dropbox-team-member-update-batch-email.md)                           | Member email operation                                                              |
| [dropbox team member update batch externalid](docs/commands/dropbox-team-member-update-batch-externalid.md)                 | Update External ID of team members                                                  |
| [dropbox team member update batch invisible](docs/commands/dropbox-team-member-update-batch-invisible.md)                   | Enable directory restriction to members                                             |
| [dropbox team member update batch profile](docs/commands/dropbox-team-member-update-batch-profile.md)                       | Update member profile                                                               |
| [dropbox team member update batch visible](docs/commands/dropbox-team-member-update-batch-visible.md)                       | Disable directory restriction to members                                            |
| [dropbox team namespace file list](docs/commands/dropbox-team-namespace-file-list.md)                                       | List all files and folders of the team namespaces                                   |
| [dropbox team namespace file size](docs/commands/dropbox-team-namespace-file-size.md)                                       | List all files and folders of the team namespaces                                   |
| [dropbox team namespace list](docs/commands/dropbox-team-namespace-list.md)                                                 | List all namespaces of the team                                                     |
| [dropbox team namespace member list](docs/commands/dropbox-team-namespace-member-list.md)                                   | List members of shared folders and team folders in the team                         |
| [dropbox team namespace summary](docs/commands/dropbox-team-namespace-summary.md)                                           | Report team namespace status summary.                                               |
| [dropbox team runas file batch copy](docs/commands/dropbox-team-runas-file-batch-copy.md)                                   | Batch copy files/folders as a member                                                |
| [dropbox team runas file list](docs/commands/dropbox-team-runas-file-list.md)                                               | List files and folders run as a member                                              |
| [dropbox team runas file sync batch up](docs/commands/dropbox-team-runas-file-sync-batch-up.md)                             | Batch sync up that run as members                                                   |
| [dropbox team runas sharedfolder batch leave](docs/commands/dropbox-team-runas-sharedfolder-batch-leave.md)                 | Batch leave from shared folders as a member                                         |
| [dropbox team runas sharedfolder batch share](docs/commands/dropbox-team-runas-sharedfolder-batch-share.md)                 | Batch share folders for members                                                     |
| [dropbox team runas sharedfolder batch unshare](docs/commands/dropbox-team-runas-sharedfolder-batch-unshare.md)             | Batch unshare folders for members                                                   |
| [dropbox team runas sharedfolder isolate](docs/commands/dropbox-team-runas-sharedfolder-isolate.md)                         | Unshare owned shared folders and leave from external shared folders run as a member |
| [dropbox team runas sharedfolder list](docs/commands/dropbox-team-runas-sharedfolder-list.md)                               | List shared folders run as the member                                               |
| [dropbox team runas sharedfolder member batch add](docs/commands/dropbox-team-runas-sharedfolder-member-batch-add.md)       | Batch add members to member's shared folders                                        |
| [dropbox team runas sharedfolder member batch delete](docs/commands/dropbox-team-runas-sharedfolder-member-batch-delete.md) | Batch delete members from member's shared folders                                   |
| [dropbox team runas sharedfolder mount add](docs/commands/dropbox-team-runas-sharedfolder-mount-add.md)                     | Add the shared folder to the specified member's Dropbox                             |
| [dropbox team runas sharedfolder mount delete](docs/commands/dropbox-team-runas-sharedfolder-mount-delete.md)               | The specified user unmounts the designated folder.                                  |
| [dropbox team runas sharedfolder mount list](docs/commands/dropbox-team-runas-sharedfolder-mount-list.md)                   | List all shared folders the specified member mounted                                |
| [dropbox team runas sharedfolder mount mountable](docs/commands/dropbox-team-runas-sharedfolder-mount-mountable.md)         | List all shared folders the member can mount                                        |
| [dropbox team sharedlink cap expiry](docs/commands/dropbox-team-sharedlink-cap-expiry.md)                                   | Set expiry cap to shared links in the team                                          |
| [dropbox team sharedlink cap visibility](docs/commands/dropbox-team-sharedlink-cap-visibility.md)                           | Set visibility cap to shared links in the team                                      |
| [dropbox team sharedlink delete links](docs/commands/dropbox-team-sharedlink-delete-links.md)                               | Batch delete shared links                                                           |
| [dropbox team sharedlink delete member](docs/commands/dropbox-team-sharedlink-delete-member.md)                             | Delete all shared links of the member                                               |
| [dropbox team sharedlink list](docs/commands/dropbox-team-sharedlink-list.md)                                               | List of shared links                                                                |
| [dropbox team sharedlink update expiry](docs/commands/dropbox-team-sharedlink-update-expiry.md)                             | Update expiration date of public shared links within the team                       |
| [dropbox team sharedlink update password](docs/commands/dropbox-team-sharedlink-update-password.md)                         | Set or update shared link passwords                                                 |
| [dropbox team sharedlink update visibility](docs/commands/dropbox-team-sharedlink-update-visibility.md)                     | Update visibility of shared links                                                   |
| [dropbox team teamfolder add](docs/commands/dropbox-team-teamfolder-add.md)                                                 | Add team folder to the team                                                         |
| [dropbox team teamfolder archive](docs/commands/dropbox-team-teamfolder-archive.md)                                         | Archive team folder                                                                 |
| [dropbox team teamfolder batch archive](docs/commands/dropbox-team-teamfolder-batch-archive.md)                             | Archiving team folders                                                              |
| [dropbox team teamfolder batch permdelete](docs/commands/dropbox-team-teamfolder-batch-permdelete.md)                       | Permanently delete team folders                                                     |
| [dropbox team teamfolder batch replication](docs/commands/dropbox-team-teamfolder-batch-replication.md)                     | Batch replication of team folders                                                   |
| [dropbox team teamfolder file list](docs/commands/dropbox-team-teamfolder-file-list.md)                                     | List files in team folders                                                          |
| [dropbox team teamfolder file lock all release](docs/commands/dropbox-team-teamfolder-file-lock-all-release.md)             | Release all locks under the path of the team folder                                 |
| [dropbox team teamfolder file lock list](docs/commands/dropbox-team-teamfolder-file-lock-list.md)                           | List locks in the team folder                                                       |
| [dropbox team teamfolder file lock release](docs/commands/dropbox-team-teamfolder-file-lock-release.md)                     | Release lock of the path in the team folder                                         |
| [dropbox team teamfolder file size](docs/commands/dropbox-team-teamfolder-file-size.md)                                     | Calculate size of team folders                                                      |
| [dropbox team teamfolder list](docs/commands/dropbox-team-teamfolder-list.md)                                               | List team folder(s)                                                                 |
| [dropbox team teamfolder member add](docs/commands/dropbox-team-teamfolder-member-add.md)                                   | Batch adding users/groups to team folders                                           |
| [dropbox team teamfolder member delete](docs/commands/dropbox-team-teamfolder-member-delete.md)                             | Batch removing users/groups from team folders                                       |
| [dropbox team teamfolder member list](docs/commands/dropbox-team-teamfolder-member-list.md)                                 | List team folder members                                                            |
| [dropbox team teamfolder partial replication](docs/commands/dropbox-team-teamfolder-partial-replication.md)                 | Partial team folder replication to the other team                                   |
| [dropbox team teamfolder permdelete](docs/commands/dropbox-team-teamfolder-permdelete.md)                                   | Permanently delete team folder                                                      |
| [dropbox team teamfolder policy list](docs/commands/dropbox-team-teamfolder-policy-list.md)                                 | List policies of team folders                                                       |
| [dropbox team teamfolder replication](docs/commands/dropbox-team-teamfolder-replication.md)                                 | Replicate a team folder to the other team                                           |
| [dropbox team teamfolder sync setting list](docs/commands/dropbox-team-teamfolder-sync-setting-list.md)                     | List team folder sync settings                                                      |
| [dropbox team teamfolder sync setting update](docs/commands/dropbox-team-teamfolder-sync-setting-update.md)                 | Batch update team folder sync settings                                              |

## Figma

| Command                                                                   | Description                           |
|---------------------------------------------------------------------------|---------------------------------------|
| [figma account info](docs/commands/figma-account-info.md)                 | Retrieve current user information     |
| [figma file export all page](docs/commands/figma-file-export-all-page.md) | Export all files/pages under the team |
| [figma file export frame](docs/commands/figma-file-export-frame.md)       | Export all frames of the Figma file   |
| [figma file export node](docs/commands/figma-file-export-node.md)         | Export Figma document Node            |
| [figma file export page](docs/commands/figma-file-export-page.md)         | Export all pages of the Figma file    |
| [figma file info](docs/commands/figma-file-info.md)                       | Show information of the figma file    |
| [figma file list](docs/commands/figma-file-list.md)                       | List files in the Figma Project       |
| [figma project list](docs/commands/figma-project-list.md)                 | List projects of the team             |

## GitHub

| Command                                                                         | Description                                         |
|---------------------------------------------------------------------------------|-----------------------------------------------------|
| [dev release checkin](docs/commands/dev-release-checkin.md)                     | Check in the new release                            |
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

| Command                                                                                                       | Description                                                              |
|---------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|
| [config auth delete](docs/commands/config-auth-delete.md)                                                     | Delete existing auth credential                                          |
| [config auth list](docs/commands/config-auth-list.md)                                                         | List all auth credentials                                                |
| [config feature disable](docs/commands/config-feature-disable.md)                                             | Disable a feature.                                                       |
| [config feature enable](docs/commands/config-feature-enable.md)                                               | Enable a feature.                                                        |
| [config feature list](docs/commands/config-feature-list.md)                                                   | List available optional features.                                        |
| [config license install](docs/commands/config-license-install.md)                                             | Install a license key                                                    |
| [config license list](docs/commands/config-license-list.md)                                                   | List available license keys                                              |
| [dropbox team insight report teamfoldermember](docs/commands/dropbox-team-insight-report-teamfoldermember.md) | Report team folder members                                               |
| [license](docs/commands/license.md)                                                                           | Show license information                                                 |
| [local file template apply](docs/commands/local-file-template-apply.md)                                       | Apply file/folder structure template to the local path                   |
| [local file template capture](docs/commands/local-file-template-capture.md)                                   | Capture file/folder structure as template from local path                |
| [log api job](docs/commands/log-api-job.md)                                                                   | Show statistics of the API log of the job specified by the job ID        |
| [log api name](docs/commands/log-api-name.md)                                                                 | Show statistics of the API log of the job specified by the job name      |
| [log cat curl](docs/commands/log-cat-curl.md)                                                                 | Format capture logs as `curl` sample                                     |
| [log cat job](docs/commands/log-cat-job.md)                                                                   | Retrieve logs of specified Job ID                                        |
| [log cat kind](docs/commands/log-cat-kind.md)                                                                 | Concatenate and print logs of specified log kind                         |
| [log cat last](docs/commands/log-cat-last.md)                                                                 | Print the last job log files                                             |
| [log job archive](docs/commands/log-job-archive.md)                                                           | Archive jobs                                                             |
| [log job delete](docs/commands/log-job-delete.md)                                                             | Delete old job history                                                   |
| [log job list](docs/commands/log-job-list.md)                                                                 | Show job history                                                         |
| [util archive unzip](docs/commands/util-archive-unzip.md)                                                     | Extract the zip archive file                                             |
| [util archive zip](docs/commands/util-archive-zip.md)                                                         | Compress target files into the zip archive                               |
| [util cert selfsigned](docs/commands/util-cert-selfsigned.md)                                                 | Generate self-signed certificate and key                                 |
| [util database exec](docs/commands/util-database-exec.md)                                                     | Execute query on SQLite3 database file                                   |
| [util database query](docs/commands/util-database-query.md)                                                   | Query SQLite3 database                                                   |
| [util date today](docs/commands/util-date-today.md)                                                           | Display current date                                                     |
| [util datetime now](docs/commands/util-datetime-now.md)                                                       | Display current date/time                                                |
| [util decode base32](docs/commands/util-decode-base32.md)                                                     | Decode text from Base32 (RFC 4648) format                                |
| [util decode base64](docs/commands/util-decode-base64.md)                                                     | Decode text from Base64 (RFC 4648) format                                |
| [util desktop open](docs/commands/util-desktop-open.md)                                                       | Open a file or folder with the default application                       |
| [util encode base32](docs/commands/util-encode-base32.md)                                                     | Encode text into Base32 (RFC 4648) format                                |
| [util encode base64](docs/commands/util-encode-base64.md)                                                     | Encode text into Base64 (RFC 4648) format                                |
| [util feed json](docs/commands/util-feed-json.md)                                                             | Load feed from the URL and output the content as JSON                    |
| [util file hash](docs/commands/util-file-hash.md)                                                             | Print file digest                                                        |
| [util git clone](docs/commands/util-git-clone.md)                                                             | Clone git repository                                                     |
| [util image exif](docs/commands/util-image-exif.md)                                                           | Print EXIF metadata of image file                                        |
| [util image placeholder](docs/commands/util-image-placeholder.md)                                             | Create placeholder image                                                 |
| [util json query](docs/commands/util-json-query.md)                                                           | Query JSON data                                                          |
| [util net download](docs/commands/util-net-download.md)                                                       | Download a file                                                          |
| [util qrcode create](docs/commands/util-qrcode-create.md)                                                     | Create a QR code image file                                              |
| [util qrcode wifi](docs/commands/util-qrcode-wifi.md)                                                         | Generate QR code for WIFI configuration                                  |
| [util table format xlsx](docs/commands/util-table-format-xlsx.md)                                             | Formatting xlsx file into text                                           |
| [util text case down](docs/commands/util-text-case-down.md)                                                   | Print lower case text                                                    |
| [util text case up](docs/commands/util-text-case-up.md)                                                       | Print upper case text                                                    |
| [util text encoding from](docs/commands/util-text-encoding-from.md)                                           | Convert text encoding to UTF-8 text file from specified encoding.        |
| [util text encoding to](docs/commands/util-text-encoding-to.md)                                               | Convert text encoding to specified encoding from UTF-8 text file.        |
| [util text nlp english entity](docs/commands/util-text-nlp-english-entity.md)                                 | Split English text into entities                                         |
| [util text nlp english sentence](docs/commands/util-text-nlp-english-sentence.md)                             | Split English text into sentences                                        |
| [util text nlp english token](docs/commands/util-text-nlp-english-token.md)                                   | Split English text into tokens                                           |
| [util text nlp japanese token](docs/commands/util-text-nlp-japanese-token.md)                                 | Tokenize Japanese text                                                   |
| [util text nlp japanese wakati](docs/commands/util-text-nlp-japanese-wakati.md)                               | Wakati gaki (tokenize Japanese text)                                     |
| [util tidy move dispatch](docs/commands/util-tidy-move-dispatch.md)                                           | Dispatch files                                                           |
| [util tidy move simple](docs/commands/util-tidy-move-simple.md)                                               | Archive local files                                                      |
| [util time now](docs/commands/util-time-now.md)                                                               | Display current time                                                     |
| [util unixtime format](docs/commands/util-unixtime-format.md)                                                 | Time format to convert the unix time (epoch seconds from 1970-01-01)     |
| [util unixtime now](docs/commands/util-unixtime-now.md)                                                       | Display current time in unixtime                                         |
| [util uuid timestamp](docs/commands/util-uuid-timestamp.md)                                                   | Parse UUID timestamp                                                     |
| [util uuid ulid](docs/commands/util-uuid-ulid.md)                                                             | Generate ULID (Universally Unique Lexicographically Sortable Identifier) |
| [util uuid v4](docs/commands/util-uuid-v4.md)                                                                 | Generate UUID v4 (random UUID)                                           |
| [util uuid v7](docs/commands/util-uuid-v7.md)                                                                 | Generate UUID v7                                                         |
| [util uuid version](docs/commands/util-uuid-version.md)                                                       | Parse version and variant of UUID                                        |
| [util video subtitles optimize](docs/commands/util-video-subtitles-optimize.md)                               | Optimize subtitles file                                                  |
| [util xlsx create](docs/commands/util-xlsx-create.md)                                                         | Create an empty spreadsheet                                              |
| [util xlsx sheet export](docs/commands/util-xlsx-sheet-export.md)                                             | Export data from the xlsx file                                           |
| [util xlsx sheet import](docs/commands/util-xlsx-sheet-import.md)                                             | Import data into xlsx file                                               |
| [util xlsx sheet list](docs/commands/util-xlsx-sheet-list.md)                                                 | List sheets of the xlsx file                                             |
| [version](docs/commands/version.md)                                                                           | Show version                                                             |

