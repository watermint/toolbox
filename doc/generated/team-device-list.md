# team device list 

List all devices/sessions in the team



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
.\tbx.exe team device list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team device list 
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



## device 

Command result report files are generated in `device.csv`, `device.xlsx` and `device.json` format.
Report in `.xlsx` format will be split into several chunks like `device_0000.xlsx`, `device_0001.xlsx`, and `device_0002.xlsx`.   

| Column                        | Description |
|-------------------------------|-------------|
| team_member_id                |             |
| email                         |             |
| status                        |             |
| given_name                    |             |
| surname                       |             |
| familiar_name                 |             |
| display_name                  |             |
| abbreviated_name              |             |
| external_id                   |             |
| account_id                    |             |
| device_tag                    |             |
| id                            |             |
| user_agent                    |             |
| os                            |             |
| browser                       |             |
| ip_address                    |             |
| country                       |             |
| created                       |             |
| updated                       |             |
| expires                       |             |
| host_name                     |             |
| client_type                   |             |
| client_version                |             |
| platform                      |             |
| is_delete_on_unlink_supported |             |
| device_name                   |             |
| os_version                    |             |
| last_carrier                  |             |





