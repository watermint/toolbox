# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)

![watermint toolbox](resources/watermint-toolbox-256x256.png)

Tools for Dropbox and Dropbox Business.

# Licensing & Disclaimers

watermint toolbox is licensed under the MIT license.
Please see LICENSE.md or LICENSE.txt for more detail.

Please carefully note:

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

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

   file          File operation                   
   group         Group management                 
   license       Show license information         
   member        Team member management           
   sharedfolder  Shared folder                    
   sharedlink    Shared Link of Personal account  
   team          Dropbox Business Team            
   teamfolder    Team folder management           

```

## Commands

| Command                                                                         | Description                                                   |
|---------------------------------------------------------------------------------|---------------------------------------------------------------|
| [file compare account](doc/generated/file-compare-account.md)                   | Compare files of two accounts                                 |
| [file compare local](doc/generated/file-compare-local.md)                       | Compare local folders and Dropbox folders                     |
| [file copy](doc/generated/file-copy.md)                                         | Copy files                                                    |
| [file delete](doc/generated/file-delete.md)                                     | Delete file or folder                                         |
| [file download](doc/generated/file-download.md)                                 | Download a file from Dropbox                                  |
| [file export doc](doc/generated/file-export-doc.md)                             | Export document                                               |
| [file import batch url](doc/generated/file-import-batch-url.md)                 | Batch import files from URL                                   |
| [file import url](doc/generated/file-import-url.md)                             | Import file from the URL                                      |
| [file list](doc/generated/file-list.md)                                         | List files and folders                                        |
| [file merge](doc/generated/file-merge.md)                                       | Merge paths                                                   |
| [file move](doc/generated/file-move.md)                                         | Move files                                                    |
| [file replication](doc/generated/file-replication.md)                           | Replicate file content to the other account                   |
| [file restore](doc/generated/file-restore.md)                                   | Restore files under given path                                |
| [file sync preflight up](doc/generated/file-sync-preflight-up.md)               | Upstream sync preflight check                                 |
| [file sync up](doc/generated/file-sync-up.md)                                   | Upstream sync with Dropbox                                    |
| [file upload](doc/generated/file-upload.md)                                     | Upload file                                                   |
| [file watch](doc/generated/file-watch.md)                                       | Watch file activities                                         |
| [group batch delete](doc/generated/group-batch-delete.md)                       | Delete groups                                                 |
| [group delete](doc/generated/group-delete.md)                                   | Delete group                                                  |
| [group list](doc/generated/group-list.md)                                       | List group(s)                                                 |
| [group member list](doc/generated/group-member-list.md)                         | List members of groups                                        |
| [job history archive](doc/generated/job-history-archive.md)                     | Archive jobs                                                  |
| [job history delete](doc/generated/job-history-delete.md)                       | Delete old job history                                        |
| [job history list](doc/generated/job-history-list.md)                           | Show job history                                              |
| [job history ship](doc/generated/job-history-ship.md)                           | Ship Job logs to Dropbox path                                 |
| [license](doc/generated/license.md)                                             | Show license information                                      |
| [member delete](doc/generated/member-delete.md)                                 | Delete members                                                |
| [member detach](doc/generated/member-detach.md)                                 | Convert Dropbox Business accounts to a Basic account          |
| [member invite](doc/generated/member-invite.md)                                 | Invite member(s)                                              |
| [member list](doc/generated/member-list.md)                                     | List team member(s)                                           |
| [member quota list](doc/generated/member-quota-list.md)                         | List team member quota                                        |
| [member quota update](doc/generated/member-quota-update.md)                     | Update team member quota                                      |
| [member quota usage](doc/generated/member-quota-usage.md)                       | List team member storage usage                                |
| [member reinvite](doc/generated/member-reinvite.md)                             | {"key":"recipe.member.reinvite.title","params":{}}            |
| [member replication](doc/generated/member-replication.md)                       | Replicate team member files                                   |
| [member update email](doc/generated/member-update-email.md)                     | Member email operation                                        |
| [member update externalid](doc/generated/member-update-externalid.md)           | Update External ID of team members                            |
| [member update profile](doc/generated/member-update-profile.md)                 | Update member profile                                         |
| [sharedfolder list](doc/generated/sharedfolder-list.md)                         | List shared folder(s)                                         |
| [sharedfolder member list](doc/generated/sharedfolder-member-list.md)           | List shared folder member(s)                                  |
| [sharedlink create](doc/generated/sharedlink-create.md)                         | Create shared link                                            |
| [sharedlink delete](doc/generated/sharedlink-delete.md)                         | Remove shared links                                           |
| [sharedlink list](doc/generated/sharedlink-list.md)                             | List of shared link(s)                                        |
| [team activity daily event](doc/generated/team-activity-daily-event.md)         | Report activities by day                                      |
| [team activity event](doc/generated/team-activity-event.md)                     | Event log                                                     |
| [team activity user](doc/generated/team-activity-user.md)                       | Activities log per user                                       |
| [team device list](doc/generated/team-device-list.md)                           | List all devices/sessions in the team                         |
| [team device unlink](doc/generated/team-device-unlink.md)                       | Unlink device sessions                                        |
| [team diag explorer](doc/generated/team-diag-explorer.md)                       | Report while team information                                 |
| [team feature](doc/generated/team-feature.md)                                   | Team feature                                                  |
| [team filerequest list](doc/generated/team-filerequest-list.md)                 | List all file requests in the team                            |
| [team info](doc/generated/team-info.md)                                         | Team information                                              |
| [team linkedapp list](doc/generated/team-linkedapp-list.md)                     | List linked applications                                      |
| [team namespace file list](doc/generated/team-namespace-file-list.md)           | List all files and folders of the team namespaces             |
| [team namespace file size](doc/generated/team-namespace-file-size.md)           | List all files and folders of the team namespaces             |
| [team namespace list](doc/generated/team-namespace-list.md)                     | List all namespaces of the team                               |
| [team namespace member list](doc/generated/team-namespace-member-list.md)       | List members of shared folders and team folders in the team   |
| [team sharedlink list](doc/generated/team-sharedlink-list.md)                   | List of shared links                                          |
| [team sharedlink update expiry](doc/generated/team-sharedlink-update-expiry.md) | Update expiration date of public shared links within the team |
| [teamfolder archive](doc/generated/teamfolder-archive.md)                       | Archive team folder                                           |
| [teamfolder batch archive](doc/generated/teamfolder-batch-archive.md)           | Archiving team folders                                        |
| [teamfolder batch permdelete](doc/generated/teamfolder-batch-permdelete.md)     | Permanently delete team folders                               |
| [teamfolder batch replication](doc/generated/teamfolder-batch-replication.md)   | Batch replication of team folders                             |
| [teamfolder file list](doc/generated/teamfolder-file-list.md)                   | List files in team folders                                    |
| [teamfolder file size](doc/generated/teamfolder-file-size.md)                   | Calculate size of team folders                                |
| [teamfolder list](doc/generated/teamfolder-list.md)                             | List team folder(s)                                           |
| [teamfolder permdelete](doc/generated/teamfolder-permdelete.md)                 | Permanently delete team folder                                |
| [teamfolder replication](doc/generated/teamfolder-replication.md)               | Replicate a team folder to the other team                     |

