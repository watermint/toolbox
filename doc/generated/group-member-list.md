# group member list 

List members of groups



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

* Individual account token: https://help.dropbox.com/ja-jp/installs-integrations/third-party/third-party-apps
* Business token: https://help.dropbox.com/ja-jp/teams-admins/admin/app-integrations

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe group member list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx group member list 
```



## Options

| Option | Description   | Default   |
|--------|---------------|-----------|
| -peer  | Account alias | {default} |


Common options:

| Option       | Description                                                                      | Default              |
|--------------|----------------------------------------------------------------------------------|----------------------|
| -concurrency | Maximum concurrency for running operation                                        | Number of processors |
| -debug       | Enable debug mode                                                                | false                |
| -proxy       | HTTP/HTTPS proxy (hostname:port)                                                 |                      |
| -quiet       | Suppress non-error messages, and make output readable by a machine (JSON format) | false                |
| -secure      | Do not store tokens into a file                                                  | false                |
| -workspace   | Workspace path                                                                   |                      |


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



## group_member 

Command result report files are generated in `group_member.csv`, `group_member.xlsx` and `group_member.json` format.
Report in `.xlsx` format will be split into several chunks like `group_member_0000.xlsx`, `group_member_0001.xlsx`, and `group_member_0002.xlsx`.   

| Column                | Description |
|-----------------------|-------------|
| group_id              |             |
| group_name            |             |
| group_management_type |             |
| access_type           |             |
| account_id            |             |
| team_member_id        |             |
| email                 |             |
| status                |             |
| surname               |             |
| given_name            |             |





