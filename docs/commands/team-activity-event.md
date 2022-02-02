---
layout: command
title: Command
lang: en
---

# team activity event

Event log 

From release 91, the command parses `-start-time` or `-end-time` as the relative duration from now with the format like "-24h" (24 hours) or "-10m" (10 minutes).
If you wanted to retrieve events every hour, then run like:

```
tbx team activity event -start-time -1h -output json > latest_events.json
```

Then, append the latest part to the entire log if you want.

```
cat latest_events.json >> all.json
```

Or more precisely, retrieve events every hour with some overlap.
```
tbx team activity event -start-time -1h5m -output json > latest_events.json
```

Then, concatenate, and de-duplicate overlapped events:
```
cat all.json latest_events.json | sort -u > _all.json && mv _all.json all.json
```

If you prefer CSV format, then use the `jq` command to convert it.
```
cat latest_events.json | jq -r '[.timestamp, .actor[.actor.".tag"].display_name, .actor[.actor.".tag"].email, .event_type.description, .event_category.".tag", .origin.access_method.end_user.".tag", .origin.geo_location.ip_address, .origin.geo_location.country, .origin.geo_location.city, .involve_non_team_member, (.participants | @text), (.context | @text)] | @csv' >> all.csv
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

| Description                                     |
|-------------------------------------------------|
| Dropbox Business: View your team's activity log |

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2022 Takayuki Okazaki
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
.\tbx.exe team activity event 
```

macOS, Linux:
```
$HOME/Desktop/tbx team activity event 
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
| `-verbose`         | Show current operations for more detail.                                                  | false                |
| `-workspace`       | Workspace path                                                                            |                      |

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: event

This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.
The command will generate a report in three different formats. `event.csv`, `event.json`, and `event.xlsx`.

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

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `event_0000.xlsx`, `event_0001.xlsx`, `event_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


