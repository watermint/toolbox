# team activity event 

Event log 

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team activity event 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team activity event 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option        | Description                                                              | Default |
|---------------|--------------------------------------------------------------------------|---------|
| `-category`   | Filter the returned events to a single category. This field is optional. |         |
| `-end-time`   | Ending time (exclusive).                                                 |         |
| `-peer`       | Account alias                                                            | default |
| `-start-time` | Starting time (inclusive)                                                |         |

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

## Report: event 
This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.
Report files are generated in three formats like below;
* `event.csv`
* `event.xlsx`
* `event.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`event_0000.xlsx`, `event_0001.xlsx`, `event_0002.xlsx`...   

| Column                   | Description                                                                                        |
|--------------------------|----------------------------------------------------------------------------------------------------|
| timestamp                | The Dropbox timestamp representing when the action was taken.                                      |
| member                   | User display name                                                                                  |
| member_email             | User email address                                                                                 |
| event_type               | The particular type of action taken.                                                               |
| category                 | Category of the events in event audit log.                                                         |
| access_method            | The method that was used to perform the action.                                                    |
| ip_address               | IP Address.                                                                                        |
| country                  | Country code.                                                                                      |
| city                     | City name                                                                                          |
| involve_non_team_members | True if the action involved a non team member either as the actor or as one of the affected users. |
| participants             | Zero or more users and/or groups that are affected by the action.                                  |
| context                  | The user or team on whose behalf the actor performed the action.                                   |
| assets                   | Zero or more content assets involved in the action.                                                |
| other_info               | The variable event schema applicable to this type of action.                                       |

