---
layout: command
title: Command `dropbox file request delete url`
lang: en
---

# dropbox file request delete url

Delete a file request by the file request URL (Irreversible operation)

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
| Dropbox: View your Dropbox file requests                                                             |
| Dropbox: View and manage your Dropbox file requests                                                  |

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
.\tbx.exe dropbox file request delete url -url URL
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox file request delete url -url URL
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Default |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-base-path` | Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder. | root    |
| `-force`     | Force delete the file request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | false   |
| `-peer`      | Account alias                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | default |
| `-url`       | URL of the file request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |         |

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

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: deleted

This report shows a list of file requests.
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                      | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| id                          | The Id of the file request                                                |
| url                         | The URL of the file request                                               |
| title                       | The title of the file request                                             |
| created                     | Date/time when the file request was created.                              |
| is_open                     | Whether or not the file request is open.                                  |
| file_count                  | The number of files this file request has received.                       |
| destination                 | The path for the folder in the Dropbox where uploaded files will be sent. |
| deadline                    | The deadline for this file request.                                       |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed.                      |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...


