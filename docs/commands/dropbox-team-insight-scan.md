---
layout: command
title: Command `dropbox team insight scan`
lang: en
---

# dropbox team insight scan

Perform comprehensive data scanning across your team for analytics and insights generation 

This command collects various team data, such as files in team folders, permissions and shared links, and stores them in a database.
The collected data can be analysed with commands such as `dropbox team insight report teamfoldermember`, or with database tools that support SQLite in general.

About how long a scan takes:

Scanning a team often takes a long time. Especially if there are a large number of files stored, the time is linearly proportional to the number of files. To increase the scanning speed, it is better to use the `-concurrency` option for parallel processing.
However, too much parallelism will increase the error rate from the Dropbox server, so a balance must be considered. According to the results of a few benchmarks, a parallelism level of 12-24 for the `-concurrency` option seems to be a good choice.
The time required for scanning depends on the response of the Dropbox server, but is around 20-30 hours per 10 million files (with `-concurrency 16`).

During the scan, users might delete, move or add files during that time. The command does not aim to capture all those differences and report exact results, but to provide rough information as quickly as possible.

For database file sizes:

As this command retrieves all metadata, including the team's files, the size of the database increases with the size of those metadata. Benchmark results show that the database size is around 10-12 GB per 10 million files stored in the team. Make sure that the path specified by `-database` has enough space before running.

About scan errors:

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

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-database**
: Path to database

**-max-retries**
: Maximum number of retries. Default: 3

**-peer**
: Account alias. Default: default

**-scan-member-folders**
: Scan member folders. Default: false

**-skip-summarize**
: Skip summarize tasks. Default: false

## Common options:

**-auth-database**
: Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)

**-auto-open**
: Auto open URL or artifact folder. Default: false

**-bandwidth-kb**
: Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited. Default: 0

**-budget-memory**
: Memory budget (limits some feature to reduce memory footprint). Options: low, normal. Default: normal

**-budget-storage**
: Storage budget (limits logs or some feature to reduce storage usage). Options: low, normal, unlimited. Default: normal

**-concurrency**
: Maximum concurrency for running operation. Default: Number of processors

**-debug**
: Enable debug mode. Default: false

**-experiment**
: Enable experimental feature(s).

**-extra**
: Extra parameter file path

**-lang**
: Display language. Options: auto, en, ja. Default: auto

**-output**
: Output format (none/text/markdown/json). Options: text, markdown, json, none. Default: text

**-output-filter**
: Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON.

**-proxy**
: HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.

**-quiet**
: Suppress non-error messages, and make output readable by a machine (JSON format). Default: false

**-retain-job-data**
: Job data retain policy. Options: default, on_error, none. Default: default

**-secure**
: Do not store tokens into a file. Default: false

**-skip-logging**
: Skip logging in the local storage. Default: false

**-verbose**
: Show current operations for more detail.. Default: false

**-workspace**
: Workspace path

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...


