# toolbox

[![Build Status](https://travis-ci.org/watermint/toolbox.svg?branch=master)](https://travis-ci.org/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg?branch=master)](https://coveralls.io/github/watermint/toolbox?branch=master)

Tools for Dropbox and Dropbox Business

# Usage

`tbx` have various features. Run without an option for a list of supported commands and options.
Released package contains binaries for each operating system. Each binary are named like `tbx-[version]-[os]`. e.g. if the binary is for Windows, then the name is like `tbx-41.2.0.0-win.exe`.

You can see available commands and options by running executable without arguments like below.

```bash
% ./tbx-41.2.0.0-mac

Usage: 
./tbx-41.2.0.0-mac COMMAND

Available commands:
  file         File operation
  team         Dropbox Business Team management
  member       Dropbox Business member management
  group        Dropbox Business group management
  

Run './tbx-41.2.0.0-mac COMMAND' for more information on a command.
```

## Available commands

| command                         | description                                                   |
|---------------------------------|---------------------------------------------------------------|
| `auth business audit`           | Ask for authorisation with Dropbox Business audit access      |
| `auth business file`            | Ask for authorisation with Dropbox Business file access       |
| `auth business info`            | Ask for authorisation with Dropbox Business info access       |
| `auth business management`      | Ask for authorisation with Dropbox Business management access |
| `auth user`                     | Ask for authorisation with Dropbox User Full access           |
| `file compare`                  | Compare files between two accounts                            |
| `file copy`                     | Copy files                                                    |
| `file list`                     | List files/folders                                            |
| `file metadata`                 | Report metadata for a file or folder                          |
| `file mirror`                   | Mirror files/folders into another account                     |
| `file move`                     | Copy files                                                    |
| `file save`                     | Save the data from a specified URL into a file                |
| `group list`                    | List groups                                                   |
| `group member list`             | List group members                                            |
| `group remove`                  | Remove group                                                  |
| `license`                       | Show license information                                      |
| `member detach`                 | Convert account into Dropbox Basic                            |
| `member invite`                 | Invite members                                                |
| `member list`                   | List members                                                  |
| `member remove`                 | Remove the member from the team                               |
| `member sync`                   | Sync member information with provided csv                     |
| `member update email`           | Update member email address                                   |
| `sharedfolder list`             | List shared folders                                           |
| `sharedlink create`             | Create shared link                                            |
| `sharedlink list`               | List shared links                                             |
| `sharedlink remove`             | Remove shared link                                            |
| `team app list`                 | List all applications linked to the team members' accounts    |
| `team audit events`             | Export activity logs                                          |
| `team audit sharing`            | Export all sharing information across team                    |
| `team device list`              | List devices or web sessions of the team                      |
| `team device unlink`            | Unlink device                                                 |
| `team feature`                  | List team feature values                                      |
| `team info`                     | Team info                                                     |
| `team namespace file list`      | List files/folders in all namespaces of the team              |
| `team namespace file size`      | Calculate size of namespaces                                  |
| `team namespace list`           | List all namespaces of the team                               |
| `team namespace member list`    | List all namespace members of the team                        |
| `team sharedlink list`          | List all shared links of the team members' accounts           |
| `team sharedlink update expire` | Update all shared link expire date of team members' accounts  |
| `teamfolder file list`          | List files/folders in all team folders of the team            |
| `teamfolder list`               | List all team folder of the team                              |
| `teamfolder mirror`             | Mirror team folders into another team                         |
| `teamfolder size`               | Calculate size of team folder                                 |

## Authentication

If an executable contains registered application keys, then the executable will ask you an authentication to your Dropbox account or a team.
Please open the provided URL, then paste authorisation code.

```
toolbox (version `dev`)
Licensed under MIT License. See https://github.com/watermint/toolbox for more detail
1. Visit the URL for the auth dialog:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

If an executable did not contain registered application keys, then the executable will ask you to create generated token from Dropbox's developer site.

```
toolbox (version 41.2.0.0)
Licensed under MIT License. See https://github.com/watermint/toolbox for more detail
1. Visit the MyApp page (you might have to login first):

https://www.dropbox.com/developers/apps

2. Proceed with "Create App"
3. Choose "Dropbox Business API"
4. Choose "Team information"
5. Enter name of your app
6. Proceed with "Create App"
7. Hit "Generate" button near "Generated access token"
8. Copy generated token

Enter the generated token:
```

The executable store tokens at the file under folder `$HOME/.toolbox/(HASH).secret`. If you don't want to store tokens into the file, then please specify option `-secure`.

## Proxy

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add `-proxy` option, like `-proxy hostname:port`.
Currently, the executable doesn't support proxies which require authentication.

# Build

## Embed App keys/secrets

If you want to include application keys and secrets into the executable, please place JSON file named `toolbox.appkeys` under `resources` folder, then compile binaries.
`toolbox.appkeys` file format is like below:

```JSON
{
  "user_full.key": "xxxxxxxxxxxxxx",
  "user_full.secret": "xxxxxxxxxxxxxx",
  "business_info.key": "xxxxxxxxxxxxxx",
  "business_info.secret": "xxxxxxxxxxxxxx",
  "business_file.key": "xxxxxxxxxxxxxx",
  "business_file.secret": "xxxxxxxxxxxxxx",
  "business_management.key": "xxxxxxxxxxxxxx",
  "business_management.secret": "xxxxxxxxxxxxxx",
  "business_audit.key": "xxxxxxxxxxxxxx",
  "business_audit.secret": "xxxxxxxxxxxxxx"
}
```


## Build script

To build an executable, please use Docker like below.

```bash
$ docker build -t toolbox . && rm -fr /tmp/dist && docker run -v /tmp/dist:/dist:rw --rm toolbox
```

