---
layout: command
title: Command `log api job`
lang: en
---

# log api job

Show statistics of the API log of the job specified by the job ID 

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe log api job 
```

macOS, Linux:
```
$HOME/Desktop/tbx log api job 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option             | Description                             | Default |
|--------------------|-----------------------------------------|---------|
| `-full-url`        | Show full URL                           | false   |
| `-interval-second` | Interval in seconds for the time series | 3600    |
| `-job-id`          | Job ID                                  |         |

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

## Report: latencies

Latency
The command will generate a report in three different formats. `latencies.csv`, `latencies.json`, and `latencies.xlsx`.

| Column     | Description        |
|------------|--------------------|
| url        | URL                |
| code       | Response code      |
| population | Number of requests |
| mean       | Mean               |
| median     | Median             |
| p_50       | Percentile 50      |
| p_70       | Percentile 70      |
| p_90       | Percentile 90      |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `latencies_0000.xlsx`, `latencies_0001.xlsx`, `latencies_0002.xlsx`, ...

## Report: population

Number of requests
The command will generate a report in three different formats. `population.csv`, `population.json`, and `population.xlsx`.

| Column     | Description        |
|------------|--------------------|
| url        | URL                |
| code       | Response code      |
| population | Number of requests |
| proportion | Proportion         |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `population_0000.xlsx`, `population_0001.xlsx`, `population_0002.xlsx`, ...

## Report: time_series

Time series summary
The command will generate a report in three different formats. `time_series.csv`, `time_series.json`, and `time_series.xlsx`.

| Column     | Description                              |
|------------|------------------------------------------|
| time       | Time                                     |
| url        | URL                                      |
| code_2xx   | Number of requests with 2xx              |
| code_3xx   | Number of requests with 3xx              |
| code_4xx   | Number of requests with 4xx (except 429) |
| code_429   | Number of requests with 429              |
| code_5xx   | Number of requests with 5xx              |
| code_other | Number of requests with other            |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `time_series_0000.xlsx`, `time_series_0001.xlsx`, `time_series_0002.xlsx`, ...


