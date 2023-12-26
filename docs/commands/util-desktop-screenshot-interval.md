---
layout: command
title: Command `util desktop screenshot interval`
lang: en
---

# util desktop screenshot interval

Take screenshots at regular intervals 

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe util desktop screenshot interval -path /LOCAL/PATH/TO/SCREENSHOT/DIR -interval INTERVAL_SECONDS
```

macOS, Linux:
```
$HOME/Desktop/tbx util desktop screenshot interval -path /LOCAL/PATH/TO/SCREENSHOT/DIR -interval INTERVAL_SECONDS
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Default                          |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| `-count`             | Number of screenshots to take. If the value is less than 1, the screenshot is taken until the process is killed.                                                                                                                                                                                                                                                                                                                                                                                                                                       | -1                               |
| `-display-id`        | Display ID to take screenshot. To get the display ID, run `util desktop display list` command.                                                                                                                                                                                                                                                                                                                                                                                                                                                         | 0                                |
| `-interval`          | Interval seconds between screenshots.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | 10                               |
| `-name-pattern`      | Name pattern of screenshot file. You can use the following placeholders:`<no value>` .. date (yyyy-MM-dd), `<no value>` .. date in UTC (yyyy-MM-dd), `<no value>` .. display height, `<no value>` .. display ID, `<no value>` .. display width, `<no value>` .. display horizontal offset, `<no value>` .. display vertical offset, `<no value>` .. 5 digit sequence number, `<no value>` .. time (HH-mm-ss), `<no value>` .. time in UTC (HH-mm-ss), `<no value>` .. timestamp (yyyyMMdd-HHmmss), `<no value>` .. timestamp in UTC (yyyyMMdd-HHmmss). | {% raw %}{{.{% endraw %}Sequence}}_{% raw %}{{.{% endraw %}Timestamp}}.png |
| `-path`              | Path to the folder to save screenshots.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |                                  |
| `-skip-if-no-change` | Skip taking screenshot if the screen is not changed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | false                            |

## Common options:

| Option             | Description                                                                               | Default              |
|--------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                 |                      |
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

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


