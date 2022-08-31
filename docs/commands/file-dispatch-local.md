---
layout: command
title: Command
lang: en
---

# file dispatch local

Dispatch local files (Irreversible operation)

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe file dispatch local -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx file dispatch local -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option     | Description       | Default |
|------------|-------------------|---------|
| `-file`    | Path to data file |         |
| `-preview` | Preview mode      | false   |

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

# File formats

## Format: File

Data file for dispatch rules.

| Column              | Description                                   | Example                                   |
|---------------------|-----------------------------------------------|-------------------------------------------|
| suffix              | Source file name suffix                       | .pdf                                      |
| source_path         | Source path                                   | <no value>/Downloads                      |
| source_file_pattern | Source file name pattern (regular expression) | toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2})  |
| dest_path_pattern   | Destination path pattern                      | <no value>/Document/<no value>-<no value> |
| dest_file_pattern   | Destination file name pattern                 | TBX_<no value>-<no value>-<no value>      |

The first line is a header line. The program will accept a file without the header.
```
suffix,source_path,source_file_pattern,dest_path_pattern,dest_file_pattern
.pdf,<no value>/Downloads,toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2}),<no value>/Document/<no value>-<no value>,TBX_<no value>-<no value>-<no value>
```

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


