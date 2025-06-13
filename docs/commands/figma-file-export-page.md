---
layout: command
title: Command `figma file export page`
lang: en
---

# figma file export page

Export all pages of the Figma file 

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
* Figma: https://help.figma.com/hc/

## Auth scopes

| Description                                                  |
|--------------------------------------------------------------|
| Figma: Permission to access your Figma files on your behalf. |

# Authorization

For the first run, `tbx` will ask you an authentication with your Figma account.
Press Enter to launch the browser. The service then performs the authorization and the application receives the results. You can close the browser window when you see the authentication success message.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Opening the authorization URL:
https://www.figma.com/oauth?client_id=xxxxxxxxxxxxxxxxxxx&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&scope=file_read&state=xxxxxxxx

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
.\tbx.exe figma file export page -key FILE_KEY -path /LOCAL/PATH/TO/EXPORT
```

macOS, Linux:
```
$HOME/Desktop/tbx figma file export page -key FILE_KEY -path /LOCAL/PATH/TO/EXPORT
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option    | Description                                                   | Default |
|-----------|---------------------------------------------------------------|---------|
| `-format` | Export format (png/jpg/svg/pdf) (Options: jpg, png, svg, pdf) | pdf     |
| `-key`    | File key                                                      |         |
| `-path`   | Output folder path                                            |         |
| `-peer`   | Account alias                                                 | default |
| `-scale`  | Export scale in percent range from 1 to 400 (default 100)     | 100     |

## Common options:

| Option             | Description                                                                                                                                           | Default              |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                                                                             |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                                                                                      | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited                                                                       | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint) (Options: low, normal)                                                                 | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage) (Options: low, normal, unlimited)                                                | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                                                                             | Number of processors |
| `-debug`           | Enable debug mode                                                                                                                                     | false                |
| `-experiment`      | Enable experimental feature(s).                                                                                                                       |                      |
| `-extra`           | Extra parameter file path                                                                                                                             |                      |
| `-lang`            | Display language (Options: auto, en, ja)                                                                                                              | auto                 |
| `-output`          | Output format (none/text/markdown/json) (Options: text, markdown, json, none)                                                                         | text                 |
| `-output-filter`   | Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON. |                      |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.                                                          |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy (Options: default, on_error, none)                                                                                             | default              |
| `-secure`          | Do not store tokens into a file                                                                                                                       | false                |
| `-skip-logging`    | Skip logging in the local storage                                                                                                                     | false                |
| `-verbose`         | Show current operations for more detail.                                                                                                              | false                |
| `-workspace`       | Workspace path                                                                                                                                        |                      |


