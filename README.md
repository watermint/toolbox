# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=svg)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/watermint/toolbox)](https://goreportcard.com/report/github.com/watermint/toolbox)

Tools for Dropbox and Dropbox Business.

# Licensing & Disclaimers

watermint toolbox is licensed under the MIT license. Please see LICENSE.md or LICENSE.txt for more detail.

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

| command                      | description                                                   |
|------------------------------|---------------------------------------------------------------|
| `file compare account`       | Compare files of two account                                  |
| `file copy`                  | Copy files                                                    |
| `file import batch url`      | Batch import files from URL                                   |
| `file import url`            | Import file from the URL                                      |
| `file list`                  | List files and folders                                        |
| `file merge`                 | Merge paths                                                   |
| `file move`                  | Move files                                                    |
| `file replication`           | Replicate file content to the other account                   |
| `group list`                 | List group(s)                                                 |
| `group member list`          | List members of groups                                        |
| `group remove`               | Remove group                                                  |
| `license`                    | Show license information                                      |
| `member detach`              | Convert Dropbox Business accounts to a Basic account          |
| `member invite`              | Invite member(s)                                              |
| `member list`                | List team member(s)                                           |
| `member quota list`          | List team member quota                                        |
| `member quota usage`         | List team member storage usage                                |
| `member remove`              | Remove members                                                |
| `member update email`        | Member email operation                                        |
| `member update profile`      | Update member profile                                         |
| `sharedfolder list`          | List shared folder(s)                                         |
| `sharedfolder member list`   | List shared folder member(s)                                  |
| `sharedlink create`          | Create shared link                                            |
| `sharedlink list`            | List of shared link(s)                                        |
| `sharedlink remove`          | Remove shared links                                           |
| `team activity`              | Team activity log                                             |
| `team device list`           | List all devices/sessions in the team                         |
| `team device unlink`         | Unlink device sessions                                        |
| `team feature`               | Team feature                                                  |
| `team filerequest list`      | List all file requests in the team                            |
| `team info`                  | Team information                                              |
| `team linkedapp list`        | List linked applications                                      |
| `team namespace file list`   | List all files and folders of the team namespaces             |
| `team namespace file size`   | List all files and folders of the team namespaces             |
| `team namespace list`        | List all namespaces of the team                               |
| `team namespace member list` | List members of shared folders and team folders in the team   |
| `team sharedlink cap expiry` | Update expiration date of public shared links within the team |
| `team sharedlink list`       | List of shared links                                          |
| `teamfolder archive`         | Archive team folder                                           |
| `teamfolder list`            | List team folder(s)                                           |
| `teamfolder permdelete`      | Permanently delete team folder                                |
| `teamfolder replication`     | Replicate a team folder to the other team                     |
| `web`                        | Launch web console (experimental)                             |


## Authentication

If an executable contains registered application keys, then the executable will ask you an authentication to your Dropbox account or a team.
Please open the provided URL, then paste authorisation code.

```
toolbox xx.x.xxx
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Testing network connection...
Done

1. Visit the URL for the auth dialog:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

The executable store tokens at the file under folder `$HOME/.toolbox/secrets/(HASH).secret`. If you don't want to store tokens into the file, then please specify option `-secure`.

## Proxy

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add `-proxy` option, like `-proxy hostname:port`.
Currently, the executable doesn't support proxies which require authentication.
