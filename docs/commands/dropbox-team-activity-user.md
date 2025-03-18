---
layout: command
title: Command `dropbox team activity user`
lang: en
---

# dropbox team activity user

Activities log per user 

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
* Dropbox for teams: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## Auth scopes

| Description                                                                                              |
|----------------------------------------------------------------------------------------------------------|
| Dropbox for teams: View and edit basic information about your Dropbox account such as your profile photo |
| Dropbox for teams: View your team's activity log                                                         |
| Dropbox for teams: View your team membership                                                             |
| Dropbox for teams: View basic information about your team including names, user count, and team settings |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account.
Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
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
.\tbx.exe dropbox team activity user 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team activity user 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option        | Description                                                              | Default |
|---------------|--------------------------------------------------------------------------|---------|
| `-category`   | Filter the returned events to a single category. This field is optional. |         |
| `-end-time`   | Ending time (exclusive).                                                 |         |
| `-peer`       | Account alias                                                            | default |
| `-start-time` | Starting time (inclusive)                                                |         |

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

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: user

This report shows an activity logs with mostly compatible with Dropbox for teams's activity logs.
The command will generate a report in three different formats. `user.csv`, `user.json`, and `user.xlsx`.

| Column                   | Description                                                                                        |
|--------------------------|----------------------------------------------------------------------------------------------------|
| timestamp                | The Dropbox timestamp representing when the action was taken.                                      |
| member                   | User display name                                                                                  |
| member_email             | User email address                                                                                 |
| event_type               | The particular type of action taken.                                                               |
| category                 | Category of the events in event audit log.                                                         |
| access_method            | The method that was used to perform the action.                                                    |
| ip_address               | IP Address.                                                                                        |
| country                  | Country code.                                                                                      |
| city                     | City name                                                                                          |
| involve_non_team_members | True if the action involved a non team member either as the actor or as one of the affected users. |
| participants             | Zero or more users and/or groups that are affected by the action.                                  |
| context                  | The user or team on whose behalf the actor performed the action.                                   |
| assets                   | Zero or more content assets involved in the action.                                                |
| other_info               | The variable event schema applicable to this type of action.                                       |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `user_0000.xlsx`, `user_0001.xlsx`, `user_0002.xlsx`, ...

## Report: user_summary

This report shows the transaction result.
The command will generate a report in three different formats. `user_summary.csv`, `user_summary.json`, and `user_summary.xlsx`.

| Column                 | Description                            |
|------------------------|----------------------------------------|
| status                 | Status of the operation                |
| reason                 | Reason of failure or skipped operation |
| input.user             | User email address                     |
| result.logins          | Number of login activities             |
| result.devices         | Number of device activities            |
| result.sharing         | Number of sharing activities           |
| result.file_operations | Number of file operation activities    |
| result.paper           | Number of activities of Paper          |
| result.others          | Number of other category activities    |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `user_summary_0000.xlsx`, `user_summary_0001.xlsx`, `user_summary_0002.xlsx`, ...


