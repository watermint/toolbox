---
layout: command
title: Command `asana team task list`
lang: en
---

# asana team task list

List tasks of the team 

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
* Asana: https://asana.com/guide/help/fundamentals/settings#gl-apps

## Auth scopes

| Description                                                                                                                                                         |
|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Asana: (1) Access your name and email address. (2) Access your tasks, projects, and workspaces. (3) Create and modify tasks, projects, and comments on your behalf. |

# Authorization

For the first run, `tbx` will ask you an authentication with your Asana (deprecated see [#647](https://github.com/watermint/toolbox/discussions/647)) account.
Press Enter to launch the browser. The service then performs the authorization and the application receives the results. You can close the browser window when you see the authentication success message.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Opening the authorization URL:
https://app.asana.com/-/oauth_authorize?client_id=xxxxxxxxxxxxxxxx&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&scope=default&state=xxxxxxxx

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
.\tbx.exe asana team task list 
```

macOS, Linux:
```
$HOME/Desktop/tbx asana team task list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                   | Description                                                       | Default |
|--------------------------|-------------------------------------------------------------------|---------|
| `-peer`                  | Account alias                                                     | default |
| `-project-name`          | Name or GID of the project Filter by exact match to the name.     |         |
| `-project-name-prefix`   | Name or GID of the project Filter by name match to the prefix.    |         |
| `-project-name-suffix`   | Name or GID of the project Filter by name match to the suffix.    |         |
| `-team-name`             | Name or GID of the team Filter by exact match to the name.        |         |
| `-team-name-prefix`      | Name or GID of the team Filter by name match to the prefix.       |         |
| `-team-name-suffix`      | Name or GID of the team Filter by name match to the suffix.       |         |
| `-workspace-name`        | Name or GID of the workspace. Filter by exact match to the name.  |         |
| `-workspace-name-prefix` | Name or GID of the workspace. Filter by name match to the prefix. |         |
| `-workspace-name-suffix` | Name or GID of the workspace. Filter by name match to the suffix. |         |

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

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: tasks

The task is the basic object around which many operations in Asana are centered.
The command will generate a report in three different formats. `tasks.csv`, `tasks.json`, and `tasks.xlsx`.

| Column        | Description                                                                   |
|---------------|-------------------------------------------------------------------------------|
| gid           | Globally unique identifier of the resource, as a string.                      |
| name          | Name of the task.                                                             |
| resource_type | The base type of this resource.                                               |
| created_at    | The time at which this resource was created.                                  |
| completed     | True if the task is currently marked complete, false if not.                  |
| completed_at  | The time at which this task was completed, or null if the task is incomplete. |
| due_at        | Date and time on which this task is due, or null if the task has no due time. |
| due_on        | Date on which this task is due, or null if the task has no due date.          |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `tasks_0000.xlsx`, `tasks_0001.xlsx`, `tasks_0002.xlsx`, ...


