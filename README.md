# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![codecov](https://codecov.io/gh/watermint/toolbox/branch/master/graph/badge.svg)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

Set of tool commands for Dropbox and Dropbox Business.

# Licensing & Disclaimers

watermint toolbox is licensed under the MIT license.
Please see LICENSE.md or LICENSE.txt for more detail.

Please carefully note:
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

# Built executable

Pre-compiled binaries can be found in [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are building directly from the source, please refer [BUILD.md](BUILD.md).

## Installing using Homebrew on macOS

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

© 2016-2020 Takayuki Okazaki
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
| connect      | Connect to the account          |       |
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

| Command                                                                 | Description                                                                                                                                                    |
|-------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [connect user_file](doc/generated/connect-user_file.md)                 | Connect to user file access                                                                                                                                    |
| [file compare account](doc/generated/file-compare-account.md)           | Compare files of two accounts                                                                                                                                  |
| [file compare local](doc/generated/file-compare-local.md)               | Compare local folders and Dropbox folders                                                                                                                      |
| [file copy](doc/generated/file-copy.md)                                 | Copy files                                                                                                                                                     |
| [file delete](doc/generated/file-delete.md)                             | Delete file or folder                                                                                                                                          |
| [file download](doc/generated/file-download.md)                         | Download a file from Dropbox                                                                                                                                   |
| [file export doc](doc/generated/file-export-doc.md)                     | Export document                                                                                                                                                |
| [file import batch url](doc/generated/file-import-batch-url.md)         | Batch import files from URL                                                                                                                                    |
| [file import url](doc/generated/file-import-url.md)                     | Import file from the URL                                                                                                                                       |
| [file info](doc/generated/file-info.md)                                 | Resolve metadata of the path                                                                                                                                   |
| [file list](doc/generated/file-list.md)                                 | List files and folders                                                                                                                                         |
| [file lock acquire](doc/generated/file-lock-acquire.md)                 | Lock a file                                                                                                                                                    |
| [file lock all release](doc/generated/file-lock-all-release.md)         | Release all locks under the specified path                                                                                                                     |
| [file lock batch acquire](doc/generated/file-lock-batch-acquire.md)     | Lock multiple files                                                                                                                                            |
| [file lock batch release](doc/generated/file-lock-batch-release.md)     | Release multiple locks                                                                                                                                         |
| [file lock list](doc/generated/file-lock-list.md)                       | List locks under the specified path                                                                                                                            |
| [file lock release](doc/generated/file-lock-release.md)                 | Release a lock                                                                                                                                                 |
| [file merge](doc/generated/file-merge.md)                               | Merge paths                                                                                                                                                    |
| [file mount list](doc/generated/file-mount-list.md)                     | List mounted/unmounted shared folders                                                                                                                          |
| [file move](doc/generated/file-move.md)                                 | Move files                                                                                                                                                     |
| [file replication](doc/generated/file-replication.md)                   | Replicate file content to the other account                                                                                                                    |
| [file restore](doc/generated/file-restore.md)                           | Restore files under given path                                                                                                                                 |
| [file search content](doc/generated/file-search-content.md)             | Search file content                                                                                                                                            |
| [file search name](doc/generated/file-search-name.md)                   | Search file name                                                                                                                                               |
| [file size](doc/generated/file-size.md)                                 | Storage usage                                                                                                                                                  |
| [file sync down](doc/generated/file-sync-down.md)                       | Downstream sync with Dropbox                                                                                                                                   |
| [file sync online](doc/generated/file-sync-online.md)                   | Sync online files                                                                                                                                              |
| [file sync up](doc/generated/file-sync-up.md)                           | Upstream sync with Dropbox                                                                                                                                     |
| [file watch](doc/generated/file-watch.md)                               | Watch file activities                                                                                                                                          |
| [filerequest create](doc/generated/filerequest-create.md)               | Create a file request                                                                                                                                          |
| [filerequest delete closed](doc/generated/filerequest-delete-closed.md) | Delete all closed file requests on this account.                                                                                                               |
| [filerequest delete url](doc/generated/filerequest-delete-url.md)       | Delete a file request by the file request URL                                                                                                                  |
| [filerequest list](doc/generated/filerequest-list.md)                   | List file requests of the individual account                                                                                                                   |
| [job history ship](doc/generated/job-history-ship.md)                   | Ship Job logs to Dropbox path                                                                                                                                  |
| [member file permdelete](doc/generated/member-file-permdelete.md)       | Permanently delete the file or folder at a given path of the team member. Please see https://www.dropbox.com/help/40 for more detail about permanent deletion. |
| [sharedfolder list](doc/generated/sharedfolder-list.md)                 | List shared folder(s)                                                                                                                                          |
| [sharedfolder member list](doc/generated/sharedfolder-member-list.md)   | List shared folder member(s)                                                                                                                                   |
| [sharedlink create](doc/generated/sharedlink-create.md)                 | Create shared link                                                                                                                                             |
| [sharedlink delete](doc/generated/sharedlink-delete.md)                 | Remove shared links                                                                                                                                            |
| [sharedlink file list](doc/generated/sharedlink-file-list.md)           | List files for the shared link                                                                                                                                 |
| [sharedlink list](doc/generated/sharedlink-list.md)                     | List of shared link(s)                                                                                                                                         |

## Dropbox Business

| Command                                                                               | Description                                                   |
|---------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [connect business_audit](doc/generated/connect-business_audit.md)                     | Connect to the team audit access                              |
| [connect business_file](doc/generated/connect-business_file.md)                       | Connect to the team file access                               |
| [connect business_info](doc/generated/connect-business_info.md)                       | Connect to the team info access                               |
| [connect business_mgmt](doc/generated/connect-business_mgmt.md)                       | Connect to the team management access                         |
| [group add](doc/generated/group-add.md)                                               | Create new group                                              |
| [group batch delete](doc/generated/group-batch-delete.md)                             | Delete groups                                                 |
| [group delete](doc/generated/group-delete.md)                                         | Delete group                                                  |
| [group folder list](doc/generated/group-folder-list.md)                               | Find folders of each group                                    |
| [group list](doc/generated/group-list.md)                                             | List group(s)                                                 |
| [group member add](doc/generated/group-member-add.md)                                 | Add a member to the group                                     |
| [group member batch add](doc/generated/group-member-batch-add.md)                     | Bulk add members into groups                                  |
| [group member batch delete](doc/generated/group-member-batch-delete.md)               | Delete members from groups                                    |
| [group member batch update](doc/generated/group-member-batch-update.md)               | Add or delete members from groups                             |
| [group member delete](doc/generated/group-member-delete.md)                           | Delete a member from the group                                |
| [group member list](doc/generated/group-member-list.md)                               | List members of groups                                        |
| [group rename](doc/generated/group-rename.md)                                         | Rename the group                                              |
| [member clear externalid](doc/generated/member-clear-externalid.md)                   | Clear external_id of members                                  |
| [member delete](doc/generated/member-delete.md)                                       | Delete members                                                |
| [member detach](doc/generated/member-detach.md)                                       | Convert Dropbox Business accounts to a Basic account          |
| [member file lock all release](doc/generated/member-file-lock-all-release.md)         | Release all locks under the path of the member                |
| [member file lock list](doc/generated/member-file-lock-list.md)                       | List locks of the member under the path                       |
| [member file lock release](doc/generated/member-file-lock-release.md)                 | Release the lock of the path as the member                    |
| [member folder list](doc/generated/member-folder-list.md)                             | Find folders for each member                                  |
| [member folder replication](doc/generated/member-folder-replication.md)               | Replicate a folder to another member's personal folder        |
| [member invite](doc/generated/member-invite.md)                                       | Invite member(s)                                              |
| [member list](doc/generated/member-list.md)                                           | List team member(s)                                           |
| [member quota list](doc/generated/member-quota-list.md)                               | List team member quota                                        |
| [member quota update](doc/generated/member-quota-update.md)                           | Update team member quota                                      |
| [member quota usage](doc/generated/member-quota-usage.md)                             | List team member storage usage                                |
| [member reinvite](doc/generated/member-reinvite.md)                                   | Reinvite invited status members to the team                   |
| [member replication](doc/generated/member-replication.md)                             | Replicate team member files                                   |
| [member update email](doc/generated/member-update-email.md)                           | Member email operation                                        |
| [member update externalid](doc/generated/member-update-externalid.md)                 | Update External ID of team members                            |
| [member update invisible](doc/generated/member-update-invisible.md)                   | Enable directory restriction to members                       |
| [member update profile](doc/generated/member-update-profile.md)                       | Update member profile                                         |
| [member update visible](doc/generated/member-update-visible.md)                       | Disable directory restriction to members                      |
| [team activity batch user](doc/generated/team-activity-batch-user.md)                 | Scan activities for multiple users                            |
| [team activity daily event](doc/generated/team-activity-daily-event.md)               | Report activities by day                                      |
| [team activity event](doc/generated/team-activity-event.md)                           | Event log                                                     |
| [team activity user](doc/generated/team-activity-user.md)                             | Activities log per user                                       |
| [team content member list](doc/generated/team-content-member-list.md)                 | List team folder & shared folder members                      |
| [team content mount list](doc/generated/team-content-mount-list.md)                   | List all mounted/unmounted shared folders of team members.    |
| [team content policy list](doc/generated/team-content-policy-list.md)                 | List policies of team folders and shared folders in the team  |
| [team device list](doc/generated/team-device-list.md)                                 | List all devices/sessions in the team                         |
| [team device unlink](doc/generated/team-device-unlink.md)                             | Unlink device sessions                                        |
| [team diag explorer](doc/generated/team-diag-explorer.md)                             | Report whole team information                                 |
| [team feature](doc/generated/team-feature.md)                                         | Team feature                                                  |
| [team filerequest list](doc/generated/team-filerequest-list.md)                       | List all file requests in the team                            |
| [team info](doc/generated/team-info.md)                                               | Team information                                              |
| [team linkedapp list](doc/generated/team-linkedapp-list.md)                           | List linked applications                                      |
| [team namespace file list](doc/generated/team-namespace-file-list.md)                 | List all files and folders of the team namespaces             |
| [team namespace file size](doc/generated/team-namespace-file-size.md)                 | List all files and folders of the team namespaces             |
| [team namespace list](doc/generated/team-namespace-list.md)                           | List all namespaces of the team                               |
| [team namespace member list](doc/generated/team-namespace-member-list.md)             | List members of shared folders and team folders in the team   |
| [team sharedlink list](doc/generated/team-sharedlink-list.md)                         | List of shared links                                          |
| [team sharedlink update expiry](doc/generated/team-sharedlink-update-expiry.md)       | Update expiration date of public shared links within the team |
| [teamfolder add](doc/generated/teamfolder-add.md)                                     | Add team folder to the team                                   |
| [teamfolder archive](doc/generated/teamfolder-archive.md)                             | Archive team folder                                           |
| [teamfolder batch archive](doc/generated/teamfolder-batch-archive.md)                 | Archiving team folders                                        |
| [teamfolder batch permdelete](doc/generated/teamfolder-batch-permdelete.md)           | Permanently delete team folders                               |
| [teamfolder batch replication](doc/generated/teamfolder-batch-replication.md)         | Batch replication of team folders                             |
| [teamfolder file list](doc/generated/teamfolder-file-list.md)                         | List files in team folders                                    |
| [teamfolder file lock all release](doc/generated/teamfolder-file-lock-all-release.md) | Release all locks under the path of the team folder           |
| [teamfolder file lock list](doc/generated/teamfolder-file-lock-list.md)               | List locks in the team folder                                 |
| [teamfolder file lock release](doc/generated/teamfolder-file-lock-release.md)         | Release lock of the path in the team folder                   |
| [teamfolder file size](doc/generated/teamfolder-file-size.md)                         | Calculate size of team folders                                |
| [teamfolder list](doc/generated/teamfolder-list.md)                                   | List team folder(s)                                           |
| [teamfolder member add](doc/generated/teamfolder-member-add.md)                       | Batch adding users/groups to team folders                     |
| [teamfolder member delete](doc/generated/teamfolder-member-delete.md)                 | Batch removing users/groups from team folders                 |
| [teamfolder member list](doc/generated/teamfolder-member-list.md)                     | List team folder members                                      |
| [teamfolder partial replication](doc/generated/teamfolder-partial-replication.md)     | Partial team folder replication to the other team             |
| [teamfolder permdelete](doc/generated/teamfolder-permdelete.md)                       | Permanently delete team folder                                |
| [teamfolder policy list](doc/generated/teamfolder-policy-list.md)                     | List policies of team folders                                 |
| [teamfolder replication](doc/generated/teamfolder-replication.md)                     | Replicate a team folder to the other team                     |

## GitHub

| Command                                                                                           | Description                                         |
|---------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| [services github content get](doc/generated/services-github-content-get.md)                       | Get content metadata of the repository              |
| [services github content put](doc/generated/services-github-content-put.md)                       | Put small text content into the repository          |
| [services github issue list](doc/generated/services-github-issue-list.md)                         | List issues of the public/private GitHub repository |
| [services github profile](doc/generated/services-github-profile.md)                               | Get the authenticated user                          |
| [services github release asset download](doc/generated/services-github-release-asset-download.md) | Download assets                                     |
| [services github release asset list](doc/generated/services-github-release-asset-list.md)         | List assets of GitHub Release                       |
| [services github release asset upload](doc/generated/services-github-release-asset-upload.md)     | Upload assets file into the GitHub Release          |
| [services github release draft](doc/generated/services-github-release-draft.md)                   | Create release draft                                |
| [services github release list](doc/generated/services-github-release-list.md)                     | List releases                                       |
| [services github tag create](doc/generated/services-github-tag-create.md)                         | Create a tag on the repository                      |

## Google Gmail

| Command                                                                                                     | Description                             |
|-------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| [services google mail filter add](doc/generated/services-google-mail-filter-add.md)                         | Add a filter.                           |
| [services google mail filter batch add](doc/generated/services-google-mail-filter-batch-add.md)             | Batch adding/deleting labels with query |
| [services google mail filter delete](doc/generated/services-google-mail-filter-delete.md)                   | Delete a filter                         |
| [services google mail filter list](doc/generated/services-google-mail-filter-list.md)                       | List filters                            |
| [services google mail label add](doc/generated/services-google-mail-label-add.md)                           | Add a label                             |
| [services google mail label delete](doc/generated/services-google-mail-label-delete.md)                     | Delete a label                          |
| [services google mail label list](doc/generated/services-google-mail-label-list.md)                         | List email labels                       |
| [services google mail label rename](doc/generated/services-google-mail-label-rename.md)                     | Rename a label                          |
| [services google mail message label add](doc/generated/services-google-mail-message-label-add.md)           | Add labels to the message               |
| [services google mail message label delete](doc/generated/services-google-mail-message-label-delete.md)     | Remove labels from the message          |
| [services google mail message list](doc/generated/services-google-mail-message-list.md)                     | List messages                           |
| [services google mail message processed list](doc/generated/services-google-mail-message-processed-list.md) | List messages in processed format.      |
| [services google mail thread list](doc/generated/services-google-mail-thread-list.md)                       | List threads                            |

## Google Sheets

| Command                                                                                                 | Description                      |
|---------------------------------------------------------------------------------------------------------|----------------------------------|
| [services google sheets sheet append](doc/generated/services-google-sheets-sheet-append.md)             | Append data to a spreadsheet     |
| [services google sheets sheet clear](doc/generated/services-google-sheets-sheet-clear.md)               | Clears values from a spreadsheet |
| [services google sheets sheet export](doc/generated/services-google-sheets-sheet-export.md)             | Export sheet data                |
| [services google sheets sheet import](doc/generated/services-google-sheets-sheet-import.md)             | Import data into the spreadsheet |
| [services google sheets sheet list](doc/generated/services-google-sheets-sheet-list.md)                 | List sheets of the spreadsheet   |
| [services google sheets spreadsheet create](doc/generated/services-google-sheets-spreadsheet-create.md) | Create a new spreadsheet         |

## Asana

| Command                                                                                         | Description                    |
|-------------------------------------------------------------------------------------------------|--------------------------------|
| [services asana team list](doc/generated/services-asana-team-list.md)                           | List team                      |
| [services asana team project list](doc/generated/services-asana-team-project-list.md)           | List projects of the team      |
| [services asana team task list](doc/generated/services-asana-team-task-list.md)                 | List task of the team          |
| [services asana workspace list](doc/generated/services-asana-workspace-list.md)                 | List workspaces                |
| [services asana workspace project list](doc/generated/services-asana-workspace-project-list.md) | List projects of the workspace |

## Slack

| Command                                                                               | Description   |
|---------------------------------------------------------------------------------------|---------------|
| [services slack conversation list](doc/generated/services-slack-conversation-list.md) | List channels |

## Utilities

| Command                                                     | Description                                      |
|-------------------------------------------------------------|--------------------------------------------------|
| [config disable](doc/generated/config-disable.md)           | Disable a feature.                               |
| [config enable](doc/generated/config-enable.md)             | Enable a feature.                                |
| [config features](doc/generated/config-features.md)         | List available optional features.                |
| [file archive local](doc/generated/file-archive-local.md)   | Archive local files                              |
| [file dispatch local](doc/generated/file-dispatch-local.md) | Dispatch local files                             |
| [job history archive](doc/generated/job-history-archive.md) | Archive jobs                                     |
| [job history delete](doc/generated/job-history-delete.md)   | Delete old job history                           |
| [job history list](doc/generated/job-history-list.md)       | Show job history                                 |
| [job log jobid](doc/generated/job-log-jobid.md)             | Retrieve logs of specified Job ID                |
| [job log kind](doc/generated/job-log-kind.md)               | Concatenate and print logs of specified log kind |
| [job log last](doc/generated/job-log-last.md)               | Print the last job log files                     |
| [license](doc/generated/license.md)                         | Show license information                         |
| [util decode base_32](doc/generated/util-decode-base_32.md) | Decode text from Base32 (RFC 4648) format        |
| [util decode base_64](doc/generated/util-decode-base_64.md) | Decode text from Base64 (RFC 4648) format        |
| [util encode base_32](doc/generated/util-encode-base_32.md) | Encode text into Base32 (RFC 4648) format        |
| [util encode base_64](doc/generated/util-encode-base_64.md) | Encode text into Base64 (RFC 4648) format        |
| [version](doc/generated/version.md)                         | Show version                                     |

