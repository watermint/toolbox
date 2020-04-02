# dev ci artifact up 

Upload CI artifact 

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT
```

macOS, Linux:

```bash
$HOME/Desktop/tbx dev ci artifact up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option          | Description            | Default |
|-----------------|------------------------|---------|
| `-dropbox-path` | Dropbox path to upload |         |
| `-local-path`   | Local path to upload   |         |
| `-peer-name`    | Account alias          |         |

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

## Report: skipped 
This report shows the transaction result.
Report files are generated in three formats like below;
* `skipped.csv`
* `skipped.xlsx`
* `skipped.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`...   

| Column                 | Description                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------|
| status                 | Status of the operation                                                                                |
| reason                 | Reason of failure or skipped operation                                                                 |
| input.file             | Local file path                                                                                        |
| input.size             | Local file size                                                                                        |
| result.name            | The last component of the path (including extension).                                                  |
| result.path_display    | The cased path to be used for display purposes only.                                                   |
| result.client_modified | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified | The last time the file was modified on Dropbox.                                                        |
| result.size            | The file size in bytes.                                                                                |

## Report: summary 
This report shows a summary of the upload results.
Report files are generated in three formats like below;
* `summary.csv`
* `summary.xlsx`
* `summary.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`...   

| Column           | Description                                         |
|------------------|-----------------------------------------------------|
| upload_start     | Time of start uploading                             |
| upload_end       | Time of finish uploading                            |
| num_bytes        | Total upload size (Bytes)                           |
| num_files_error  | The number of files failed or got an error.         |
| num_files_upload | The number of files uploaded or to upload.          |
| num_files_skip   | The number of files skipped or to skip.             |
| num_api_call     | The number of estimated upload API call for upload. |

## Report: uploaded 
This report shows the transaction result.
Report files are generated in three formats like below;
* `uploaded.csv`
* `uploaded.xlsx`
* `uploaded.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`...   

| Column                 | Description                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------|
| status                 | Status of the operation                                                                                |
| reason                 | Reason of failure or skipped operation                                                                 |
| input.file             | Local file path                                                                                        |
| input.size             | Local file size                                                                                        |
| result.name            | The last component of the path (including extension).                                                  |
| result.path_display    | The cased path to be used for display purposes only.                                                   |
| result.client_modified | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified | The last time the file was modified on Dropbox.                                                        |
| result.size            | The file size in bytes.                                                                                |

