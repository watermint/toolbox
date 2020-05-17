# file compare account 

Compare files of two accounts 

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe file compare account -left left -left-path /path/to/compare -right right -right-path /path/to/compare
```

macOS, Linux:

```bash
$HOME/Desktop/tbx file compare account -left left -left-path /path/to/compare -right right -right-path /path/to/compare
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

If you want to compare different path in same account, please specify same alias name to `-left` and `-right`.

## Options

| Option        | Description                        | Default |
|---------------|------------------------------------|---------|
| `-left`       | Account alias (left)               | left    |
| `-left-path`  | The path from account root (left)  |         |
| `-right`      | Account alias (right)              | right   |
| `-right-path` | The path from account root (right) |         |

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

## Report: diff 
This report shows a difference between to folders.
Report files are generated in three formats like below;
* `diff.csv`
* `diff.xlsx`
* `diff.json`

But if you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`diff_0000.xlsx`, `diff_0001.xlsx`, `diff_0002.xlsx`...   

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder of file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

