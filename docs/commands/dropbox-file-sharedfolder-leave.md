---
layout: command
title: Command `dropbox file sharedfolder leave`
lang: en
---

# dropbox file sharedfolder leave

Leave the shared folder 

Upon success, the current user will no longer have access to the folder. Please use `dropbox file sharedfolder list` command to find the shared_folder_id of the folder you want to leave.

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS      | Path                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox (Individual account): https://help.dropbox.com/installs-integrations/third-party/third-party-apps

## Auth scopes

| Description                                                                                          |
|------------------------------------------------------------------------------------------------------|
| Dropbox: View basic information about your Dropbox account such as your username, email, and country |
| Dropbox: View your Dropbox sharing settings and collaborators                                        |
| Dropbox: View and manage your Dropbox sharing settings and collaborators                             |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account.
Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the application.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:\n\nhttps://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx\n\n2. Click 'Allow' (you might have to login first):\n3. Copy the authorization code:
Enter the authorization code
```

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe dropbox file sharedfolder leave -shared-folder-id SHARED_FOLDER_ID
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox file sharedfolder leave -shared-folder-id SHARED_FOLDER_ID
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option              | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Default |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-base-path`        | Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder. | root    |
| `-keep-copy`        | Keep a copy of the folder's contents upon relinquishing membership.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | false   |
| `-peer`             | Account alias                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | default |
| `-shared-folder-id` | The ID for the shared folder.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |         |

## Common options:

| Option             | Description                                                                                                                                           | Default              |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                                                                             |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                                                                                      | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited                                                                       | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint)                                                                                        | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage)                                                                                  | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                                                                             | Number of processors |
| `-debug`           | Enable debug mode                                                                                                                                     | false                |
| `-experiment`      | Enable experimental feature(s).                                                                                                                       |                      |
| `-extra`           | Extra parameter file path                                                                                                                             |                      |
| `-lang`            | Display language                                                                                                                                      | auto                 |
| `-output`          | Output format (none/text/markdown/json)                                                                                                               | text                 |
| `-output-filter`   | Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON. |                      |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.                                                          |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy                                                                                                                                | default              |
| `-secure`          | Do not store tokens into a file                                                                                                                       | false                |
| `-skip-logging`    | Skip logging in the local storage                                                                                                                     | false                |
| `-verbose`         | Show current operations for more detail.                                                                                                              | false                |
| `-workspace`       | Workspace path                                                                                                                                        |                      |


