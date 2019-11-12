# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=svg)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/watermint/toolbox)](https://goreportcard.com/report/github.com/watermint/toolbox)

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
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox Business

Usage:
./tbx  command

Available commands:
   file          File operation
   group         Group management (Dropbox Business)
   license       Show license information
   member        Team member management (Dropbox Business)
   sharedfolder  Shared folder
   sharedlink    Shared Link of Personal account
   team          Dropbox Business Team
   teamfolder    Team folder management (Dropbox Business)
   web           Launch web console (experimental)
```

## Commands

| Command                                                                         | Description                                                   |
|---------------------------------------------------------------------------------|---------------------------------------------------------------|
| [file compare account](doc/generated/file-compare-account.md)                   | Compare files of two accounts                                 |
| [file compare local](doc/generated/file-compare-local.md)                       | Compare local folders and Dropbox folders                     |
| [file copy](doc/generated/file-copy.md)                                         | Copy files                                                    |
| [file import batch url](doc/generated/file-import-batch-url.md)                 | Batch import files from URL                                   |
| [file import url](doc/generated/file-import-url.md)                             | Import file from the URL                                      |
| [file list](doc/generated/file-list.md)                                         | List files and folders                                        |
| [file merge](doc/generated/file-merge.md)                                       | Merge paths                                                   |
| [file move](doc/generated/file-move.md)                                         | Move files                                                    |
| [file replication](doc/generated/file-replication.md)                           | Replicate file content to the other account                   |
| [file upload](doc/generated/file-upload.md)                                     | Upload file                                                   |
| [group delete](doc/generated/group-delete.md)                                   | Delete group                                                  |
| [group list](doc/generated/group-list.md)                                       | List group(s)                                                 |
| [group member list](doc/generated/group-member-list.md)                         | List members of groups                                        |
| [license](doc/generated/license.md)                                             | Show license information                                      |
| [member delete](doc/generated/member-delete.md)                                 | Delete members                                                |
| [member detach](doc/generated/member-detach.md)                                 | Convert Dropbox Business accounts to a Basic account          |
| [member invite](doc/generated/member-invite.md)                                 | Invite member(s)                                              |
| [member list](doc/generated/member-list.md)                                     | List team member(s)                                           |
| [member quota list](doc/generated/member-quota-list.md)                         | List team member quota                                        |
| [member quota usage](doc/generated/member-quota-usage.md)                       | List team member storage usage                                |
| [member update email](doc/generated/member-update-email.md)                     | Member email operation                                        |
| [member update profile](doc/generated/member-update-profile.md)                 | Update member profile                                         |
| [sharedfolder list](doc/generated/sharedfolder-list.md)                         | List shared folder(s)                                         |
| [sharedfolder member list](doc/generated/sharedfolder-member-list.md)           | List shared folder member(s)                                  |
| [sharedlink create](doc/generated/sharedlink-create.md)                         | Create shared link                                            |
| [sharedlink delete](doc/generated/sharedlink-delete.md)                         | Remove shared links                                           |
| [sharedlink list](doc/generated/sharedlink-list.md)                             | List of shared link(s)                                        |
| [team activity daily event](doc/generated/team-activity-daily-event.md)         | Report activities by day                                      |
| [team activity event](doc/generated/team-activity-event.md)                     | Event log                                                     |
| [team device list](doc/generated/team-device-list.md)                           | List all devices/sessions in the team                         |
| [team device unlink](doc/generated/team-device-unlink.md)                       | Unlink device sessions                                        |
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
| [teamfolder list](doc/generated/teamfolder-list.md)                             | List team folder(s)                                           |
| [teamfolder permdelete](doc/generated/teamfolder-permdelete.md)                 | Permanently delete team folder                                |
| [teamfolder replication](doc/generated/teamfolder-replication.md)               | Replicate a team folder to the other team                     |
| [web](doc/generated/web.md)                                                     | Launch web console (experimental)                             |

