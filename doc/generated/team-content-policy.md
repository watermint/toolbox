# team content policy 

List policies of team folders and shared folders in the team 

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team content policy 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team content policy 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options

| Option                | Description                                                | Default |
|-----------------------|------------------------------------------------------------|---------|
| `-folder-name`        | Filter by folder name. Filter by exact match to the name.  |         |
| `-folder-name-prefix` | Filter by folder name. Filter by name match to the prefix. |         |
| `-folder-name-suffix` | Filter by folder name. Filter by name match to the suffix. |         |
| `-peer`               | Account alias                                              | default |

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

## Report: policy 
This report shows a list of shared folders and team folders with their current policy settings.
Report files are generated in three formats like below;
* `policy.csv`
* `policy.xlsx`
* `policy.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`...   

| Column               | Description                                                                                              |
|----------------------|----------------------------------------------------------------------------------------------------------|
| path                 | Path                                                                                                     |
| is_team_folder       | `true` if the folder is a team folder, or inside of a team folder                                        |
| owner_team_name      | Team name of the team that owns the folder                                                               |
| policy_manage_access | Who can add and remove members from this shared folder.                                                  |
| policy_shared_link   | Who links can be shared with.                                                                            |
| policy_member        | Who can be a member of this shared folder, taking into account both the folder and the team-wide policy. |
| policy_viewer_info   | Who can enable/disable viewer info for this shared folder.                                               |

