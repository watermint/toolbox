---
layout: command
title: Command `dev test coverage list`
lang: en
---

# dev test coverage list

Test Coverage List 

Analyze and list packages with test coverage below threshold

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe dev test coverage list 
```

macOS, Linux:
```
$HOME/Desktop/tbx dev test coverage list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

**-max-package**
: Maximum packages to display. Default: 30

**-min-package**
: Minimum packages to display. Default: 10

**-threshold**
: Coverage threshold percentage. Default: 50

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

## Report: coverage_report

{"key":"recipe.dev.test.coverage.coverage_report.desc","params":{}}
The command will generate a report in three different formats. `coverage_report.csv`, `coverage_report.json`, and `coverage_report.xlsx`.

| Column     | Description                                                                    |
|------------|--------------------------------------------------------------------------------|
| package    | {"key":"recipe.dev.test.coverage.coverage_report.package.desc","params":{}}    |
| coverage   | {"key":"recipe.dev.test.coverage.coverage_report.coverage.desc","params":{}}   |
| statements | {"key":"recipe.dev.test.coverage.coverage_report.statements.desc","params":{}} |
| no_test    | {"key":"recipe.dev.test.coverage.coverage_report.no_test.desc","params":{}}    |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `coverage_report_0000.xlsx`, `coverage_report_0001.xlsx`, `coverage_report_0002.xlsx`, ...

## Report: summary_report

{"key":"recipe.dev.test.coverage.coverage_report.desc","params":{}}
The command will generate a report in three different formats. `summary_report.csv`, `summary_report.json`, and `summary_report.xlsx`.

| Column     | Description                                                                    |
|------------|--------------------------------------------------------------------------------|
| package    | {"key":"recipe.dev.test.coverage.coverage_report.package.desc","params":{}}    |
| coverage   | {"key":"recipe.dev.test.coverage.coverage_report.coverage.desc","params":{}}   |
| statements | {"key":"recipe.dev.test.coverage.coverage_report.statements.desc","params":{}} |
| no_test    | {"key":"recipe.dev.test.coverage.coverage_report.no_test.desc","params":{}}    |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `summary_report_0000.xlsx`, `summary_report_0001.xlsx`, `summary_report_0002.xlsx`, ...


