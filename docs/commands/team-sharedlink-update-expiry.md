---
layout: command
title: Command
lang: en
---

# team sharedlink update expiry

Update expiration date of public shared links within the team (Irreversible operation)

Note: From Release 87, this command will receive a file to select shared links to update. If you wanted to update the expiry for all shared links in the team, please consider using a combination of `team sharedlink list`. For example, if you are familiar with the command [jq](https://stedolan.github.io/jq/), then you can do an equivalent operation as like below (force expiry within 28 days for every public link).

```
tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -days 28
```

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
* Dropbox Business: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## Auth scopes

| Description                                                                       |
|-----------------------------------------------------------------------------------|
| Dropbox Business: View your team membership                                       |
| Dropbox Business: View and manage your Dropbox sharing settings and collaborators |
| Dropbox Business: View structure of your team's and members' folders              |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2021 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
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
.\tbx.exe team sharedlink update expiry -file /PATH/TO/DATA_FILE.csv -days 28
```

macOS, Linux:
```
$HOME/Desktop/tbx team sharedlink update expiry -file /PATH/TO/DATA_FILE.csv -days 28
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option  | Description                     | Default |
|---------|---------------------------------|---------|
| `-at`   | New expiration date and time    |         |
| `-days` | Days to the new expiration date | 0       |
| `-file` | Path to data file               |         |
| `-peer` | Account alias                   | default |

## Common options:

| Option            | Description                                                                               | Default              |
|-------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auto-open`      | Auto open URL or artifact folder                                                          | false                |
| `-bandwidth-kb`   | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited           | 0                    |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                            | normal               |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)                      | normal               |
| `-concurrency`    | Maximum concurrency for running operation                                                 | Number of processors |
| `-debug`          | Enable debug mode                                                                         | false                |
| `-experiment`     | Enable experimental feature(s).                                                           |                      |
| `-lang`           | Display language                                                                          | auto                 |
| `-output`         | Output format (none/text/markdown/json)                                                   | text                 |
| `-proxy`          | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy. |                      |
| `-quiet`          | Suppress non-error messages, and make output readable by a machine (JSON format)          | false                |
| `-secure`         | Do not store tokens into a file                                                           | false                |
| `-verbose`        | Show current operations for more detail.                                                  | false                |
| `-workspace`      | Workspace path                                                                            |                      |

# File formats

## Format: File

Target shared link

| Column | Description     | Example                                  |
|--------|-----------------|------------------------------------------|
| url    | Shared link URL | https://www.dropbox.com/scl/fo/fir9vjelf |

The first line is a header line. The program will accept a file without the header.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                             |
|-------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                                 |
| reason            | Reason of failure or skipped operation                                                                                                                                                                                  |
| input.url         | Shared link URL                                                                                                                                                                                                         |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                            |
| result.url        | URL of the shared link.                                                                                                                                                                                                 |
| result.name       | The linked file name (including extension).                                                                                                                                                                             |
| result.expires    | Expiration time, if set.                                                                                                                                                                                                |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                         |
| result.visibility | The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                                  |
| result.surname    | Surname of the link owner                                                                                                                                                                                               |
| result.given_name | Given name of the link owner                                                                                                                                                                                            |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

