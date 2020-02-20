# team activity batch user 

Scan activities for multiple users 

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

This command use following access type(s) during the operation:
* Dropbox Business Auditing

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team activity batch user 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team activity batch user 
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
| `-user-list`  | User email address list file                                             |         |

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

# File formats

## Format: UserList 

| Column | Description        | Value example    |
|--------|--------------------|------------------|
| email  | User email address | john@example.com |

The first line is a header line. The program will accept file without the header.

```csv
email
john@example.com
```

# Authorization

For the first run, `tbx` will ask you an authentication with your Dropbox account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.

```

watermint toolbox xx.x.xxx
==========================

© 2016-2020 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code

```

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: combined 

Report files are generated in three formats like below;
* `combined.csv`
* `combined.xlsx`
* `combined.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`combined_0000.xlsx`, `combined_0001.xlsx`, `combined_0002.xlsx`...   

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

## Report: user 

Report files are generated in three formats like below;
* `user.csv`
* `user.xlsx`
* `user.json`

But if you run with `-low-memory` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows;
`user_0000.xlsx`, `user_0001.xlsx`, `user_0002.xlsx`...   

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
