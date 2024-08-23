---
layout: command
title: Command `dropbox team insight scan`
lang: en
---

# dropbox team insight scan

Scans team data for analysis 

This command collects various team data, such as files in team folders, permissions and shared links, and stores them in a database.
The collected data can be analysed with commands such as `dropbox team insight report teamfoldermember`, or with database tools that support SQLite in general.

About how long a scan takes:.

Scanning a team often takes a long time. Especially if there are a large number of files stored, the time is linearly proportional to the number of files. To increase the scanning speed, it is better to use the `-concurrency` option for parallel processing.
However, too much parallelism will increase the error rate from the Dropbox server, so a balance must be considered. According to the results of a few benchmarks, a parallelism level of 12-24 for the `-concurrency` option seems to be a good choice.
The time required for scanning depends on the response of the Dropbox server, but is around 20-30 hours per 10 million files (with `-concurrency 16`).

During the scan, users might delete, move or add files during that time. The command does not aim to capture all those differences and report exact results, but to provide rough information as quickly as possible.

For database file sizes:.

As this command retrieves all metadata, including the team's files, the size of the database increases with the size of those metadata. Benchmark results show that the database size is around 10-12 GB per 10 million files stored in the team. Make sure that the path specified by `-database` has enough space before running.

About scan errors:.

The Dropbox server may return an error when running the scan. The command will automatically try to re-run the scan several times, but the error may not be resolved for a certain period of time due to server congestion or condition. In that case, the command stops the re-run and records the scan task in the database where the error occurred.
If you want to re-run a failed scan, use the `dropbox team insight scanretry` command to run the scan again.
If the issue is not resolved after repeated re-runs and you want to analyse only the coverage of the current scan, you need to perform an aggregation task before the analysis. Aggregation tasks can be performed with the `dropbox team insight summary` command.

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

| Description |
|-------------|

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account.
Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2024 Takayuki Okazaki
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
.\tbx.exe dropbox team insight scan -database /LOCAL/PATH/TO/database
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team insight scan -database /LOCAL/PATH/TO/database
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                 | Description               | Default |
|------------------------|---------------------------|---------|
| `-database`            | Path to database          |         |
| `-max-retries`         | Maximum number of retries | 3       |
| `-peer`                | Account alias             | default |
| `-scan-member-folders` | Scan member folders       | false   |
| `-skip-summarize`      | Skip summarize tasks      | false   |

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

## Report: errors

Error report
The command will generate a report in three different formats. `errors.csv`, `errors.json`, and `errors.xlsx`.

| Column   | Description    |
|----------|----------------|
| category | Error category |
| message  | Error message  |
| tag      | Error tag      |
| detail   | Error details  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


