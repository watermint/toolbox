# file compare account 

Compare files of two accounts

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it.
If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Individual account: https://help.dropbox.com/ja-jp/installs-integrations/third-party/third-party-apps

This command use following access type(s) during the operation:
* Dropbox Full access

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

If you want to compare different path in same account, please specify same alias name to `-left` and `-right`.

## Options

| Option        | Description                        | Default   |
|---------------|------------------------------------|-----------|
| `-left`       | Account alias (left)               | {default} |
| `-left-path`  | The path from account root (left)  |           |
| `-right`      | Account alias (right)              | {default} |
| `-right-path` | The path from account root (right) |           |

Common options:

| Option         | Description                                                                      | Default              |
|----------------|----------------------------------------------------------------------------------|----------------------|
| `-concurrency` | Maximum concurrency for running operation                                        | Number of processors |
| `-debug`       | Enable debug mode                                                                | false                |
| `-proxy`       | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| `-quiet`       | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| `-secure`      | Do not store tokens into a file                                                  | false                |
| `-workspace`   | Workspace path                                                                   |                      |

## Authentication

For the first run, `toolbox` will ask you an authentication with your Dropbox account. 
Please copy the link and paste it into your browser. Then proceed to authorization.
After authorization, Dropbox will show you an authorization code.
Please copy that code and paste it to the `toolbox`.

```
watermint toolbox xx.x.xxx
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Testing network connection...
Done

1. Visit the URL for the auth dialog:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

# Result

Report file path will be displayed last line of the command line output.
If you missed command line output, please see path below.
[job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: diff 

Report files are generated in `diff.csv`, `diff.xlsx` and `diff.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `diff_0000.xlsx`, `diff_0001.xlsx`, `diff_0002.xlsx`...   

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

