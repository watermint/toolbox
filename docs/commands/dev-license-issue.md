---
layout: command
title: Command `dev license issue`
lang: en
---

# dev license issue

Issue a license 

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
* GitHub: https://developer.github.com/apps/managing-oauth-apps/deleting-an-oauth-app/

## Auth scopes

| Description                                                                                                                                                                                                                                                                                                                                                    |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GitHub: Grants full access to repositories, including private repositories. That includes read/write access to code, commit statuses, repository and organization projects, invitations, collaborators, adding team memberships, deployment statuses, and repository webhooks for repositories and organizations. Also grants ability to manage user projects. |

# Authorization

For the first run, `tbx` will ask you an authentication with your GitHub account.
Press the Enter key to launch the browser. The service then performs the authorization and tbx receives the results. You can close the browser window when you see the authentication success message.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Opening the authorization URL:
https://github.com/login/oauth/authorize?client_id=xxxxxxxxxxxxxxxxxxxx&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&scope=repo&state=xxxxxxxx

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
.\tbx.exe dev license issue -licensee-email LICENSEE_EMAIL -licensee-name LICENSEE_NAME
```

macOS, Linux:
```
$HOME/Desktop/tbx dev license issue -licensee-email LICENSEE_EMAIL -licensee-name LICENSEE_NAME
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                       | Description                                                         | Default            |
|------------------------------|---------------------------------------------------------------------|--------------------|
| `-app-name`                  | Application name                                                    | watermint toolbox  |
| `-branch`                    | License repository branch                                           | main               |
| `-expiration`                | License expiration date                                             |                    |
| `-licensee-email`            | Licensee email                                                      |                    |
| `-licensee-name`             | Licensee name                                                       |                    |
| `-lifecycle-available-after` | Lifecycle available after this period from the build time (seconds) | 94608000           |
| `-lifecycle-warning-after`   | Lifecycle warning after this period from the build time (seconds)   | 31536000           |
| `-owner`                     | License repository owner                                            | watermint          |
| `-peer`                      | Account alias                                                       | default            |
| `-recipe-allowed-prefix`     | Recipe allowed prefix                                               |                    |
| `-recipes-allowed`           | Comma separated list of recipes allowed                             |                    |
| `-repository`                | License repository                                                  | toolbox-supplement |
| `-scope`                     | License scope                                                       |                    |

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
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy.                                                             |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy                                                                                                                                | default              |
| `-secure`          | Do not store tokens into a file                                                                                                                       | false                |
| `-skip-logging`    | Skip logging in the local storage                                                                                                                     | false                |
| `-verbose`         | Show current operations for more detail.                                                                                                              | false                |
| `-workspace`       | Workspace path                                                                                                                                        |                      |


