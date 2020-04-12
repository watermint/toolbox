# services github release asset up 

Upload assets file into the GitHub Release (Experimental)

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe services github release asset up -asset /LOCAL/PATH/TO/assets
```

macOS, Linux:

```bash
$HOME/Desktop/tbx services github release asset up -asset /LOCAL/PATH/TO/assets
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option        | Description             | Default          |
|---------------|-------------------------|------------------|
| `-asset`      | Path to assets          |                  |
| `-owner`      | Owner of the repository |                  |
| `-peer`       | Account alias           | &{default <nil>} |
| `-release`    | Release tag name        |                  |
| `-repository` | Name of the repository  |                  |

Common options:

| Option          | Description                                                                      | Default              |
|-----------------|----------------------------------------------------------------------------------|----------------------|
| `-auto-open`    | Auto open URL or artifact folder                                                 | false                |
| `-bandwidth-kb` | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited  | 0                    |
| `-concurrency`  | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`        | Enable debug mode                                                                | false                |
| `-low-memory`   | Low memory footprint mode                                                        | false                |
| `-output`       | Output format (none/text/markdown/json)                                          | text                 |
| `-proxy`        | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`        | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`       | Do not store tokens into a file                                                  | false                |
| `-workspace`    | Workspace path                                                                   |                      |

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: uploads 
This report shows the transaction result.
Report files are generated in three formats like below;
* `uploads.csv`
* `uploads.xlsx`
* `uploads.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`uploads_0000.xlsx`, `uploads_0001.xlsx`, `uploads_0002.xlsx`...   

| Column                | Description                            |
|-----------------------|----------------------------------------|
| status                | Status of the operation                |
| reason                | Reason of failure or skipped operation |
| input.file            | File path                              |
| result.name           | Name of the asset                      |
| result.size           | Size of the asset                      |
| result.state          | State of the asset                     |
| result.download_count | Number of downloads                    |

