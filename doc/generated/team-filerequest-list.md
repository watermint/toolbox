# team filerequest list 

List all file requests in the team 

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team filerequest list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team filerequest list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option  | Description   | Default |
|---------|---------------|---------|
| `-peer` | Account alias | default |

Common options:

| Option            | Description                                                                      | Default              |
|-------------------|----------------------------------------------------------------------------------|----------------------|
| `-auto-open`      | Auto open URL or artifact folder                                                 | false                |
| `-bandwidth-kb`   | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited  | 0                    |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                   | normal               |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)             | normal               |
| `-concurrency`    | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`          | Enable debug mode                                                                | false                |
| `-output`         | Output format (none/text/markdown/json)                                          | text                 |
| `-proxy`          | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`          | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`         | Do not store tokens into a file                                                  | false                |
| `-workspace`      | Workspace path                                                                   |                      |

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: file_request 
This report shows a list of file requests with the file request owner team member.
Report files are generated in three formats like below;
* `file_request.csv`
* `file_request.xlsx`
* `file_request.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`...   

| Column                      | Description                                                                   |
|-----------------------------|-------------------------------------------------------------------------------|
| email                       | Email address of this file request owner.                                     |
| status                      | The user status of this file request owner (active/invited/suspended/removed) |
| surname                     | Surname of this file request owner.                                           |
| given_name                  | Given name of this file request owner.                                        |
| url                         | The URL of the file request.                                                  |
| title                       | The title of the file request.                                                |
| created                     | When this file request was created.                                           |
| is_open                     | Whether or not the file request is open.                                      |
| file_count                  | The number of files this file request has received.                           |
| destination                 | The path of the folder in the Dropbox where uploaded files will be sent       |
| deadline                    | The deadline for this file request.                                           |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           |

