---
layout: command
title: Command
lang: en
---

# services google sheets sheet import

Import data into the spreadsheet 

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
* Google: https://support.google.com/accounts/answer/3466521

## Auth scopes

| Description                                                                    |
|--------------------------------------------------------------------------------|
| Google Sheets: See, edit, create, and delete your spreadsheets in Google Drive |

# Authorization

For the first run, `tbx` will ask you an authentication with your Google account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2022 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://accounts.google.com/o/oauth2/auth?client_id=xxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&state=xxxxxxxx

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
.\tbx.exe services google sheets sheet import -data /LOCAL/PATH/TO/INPUT.csv -id VALUE -range VALUE
```

macOS, Linux:
```
$HOME/Desktop/tbx services google sheets sheet import -data /LOCAL/PATH/TO/INPUT.csv -id VALUE -range VALUE
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option       | Description                                                                                                                                                                                                                                                                                                                                                                          | Default                                                         |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|
| `-data`      | Data file path                                                                                                                                                                                                                                                                                                                                                                       |                                                                 |
| `-id`        | Spreadsheet Id                                                                                                                                                                                                                                                                                                                                                                       |                                                                 |
| `-input-raw` | Raw input                                                                                                                                                                                                                                                                                                                                                                            | false                                                           |
| `-peer`      | Account alias                                                                                                                                                                                                                                                                                                                                                                        | &{default [https://www.googleapis.com/auth/spreadsheets] <nil>} |
| `-range`     | The range the values cover, in A1 notation. This is a string like Sheet1!A1:B2, that refers to a group of cells in the spreadsheet, and is typically used in formulas. `Sheet1!A1:B2` refers to the first two cells in the top two rows of Sheet1. `A1:B2` refers to the first two cells in the top two rows of the first visible sheet. `Sheet1` refers to all the cells in Sheet1. |                                                                 |

## Common options:

| Option             | Description                                                                               | Default              |
|--------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auto-open`       | Auto open URL or artifact folder                                                          | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited           | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint)                            | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage)                      | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                 | Number of processors |
| `-debug`           | Enable debug mode                                                                         | false                |
| `-experiment`      | Enable experimental feature(s).                                                           |                      |
| `-extra`           | Extra parameter file path                                                                 |                      |
| `-lang`            | Display language                                                                          | auto                 |
| `-output`          | Output format (none/text/markdown/json)                                                   | text                 |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy. |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)          | false                |
| `-retain-job-data` | Job data retain policy                                                                    | default              |
| `-secure`          | Do not store tokens into a file                                                           | false                |
| `-skip-logging`    | Skip logging in the local storage                                                         | false                |
| `-verbose`         | Show current operations for more detail.                                                  | false                |
| `-workspace`       | Workspace path                                                                            |                      |

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: updated

Updated value
The command will generate a report in three different formats. `updated.csv`, `updated.json`, and `updated.xlsx`.

| Column          | Description               |
|-----------------|---------------------------|
| spreadsheet_id  | Spreadsheet Id            |
| updated_range   | Updated range             |
| updated_rows    | Number of updated rows    |
| updated_columns | Number of updated columns |
| updated_cells   | Number of updated cells   |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `updated_0000.xlsx`, `updated_0001.xlsx`, `updated_0002.xlsx`, ...

# Grid data input for the command

## Grid data input: Data

Input data file

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


