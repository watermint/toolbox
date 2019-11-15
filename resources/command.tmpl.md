# {{.Command}} 

{{.CommandTitle}}

{{.CommandDesc}}

{{if .UseAuth}}
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
{{ if .UseAuthPersonal }}* Individual account: https://help.dropbox.com/ja-jp/installs-integrations/third-party/third-party-apps{{end}}{{ if .UseAuthBusiness }}* Dropbox Business: https://help.dropbox.com/ja-jp/teams-admins/admin/app-integrations{{end}}

This command use following access type(s) during the operation:
{{ range $scope := .AuthScopes }}{{ with (eq $scope "business_info")}}* Dropbox Business Information access{{end}}{{ with (eq $scope "business_file")}}* Dropbox Business File access{{end}}{{ with (eq $scope "business_mgmt")}}* Dropbox Business management{{end}}{{ with (eq $scope "business_audit")}}* Dropbox Business Auditing{{end}}{{ with (eq $scope "user_file")}}* Dropbox Full access{{end}}{{end}}
{{end}}

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe {{.Command}} {{.CommandArgs}}
```

macOS, Linux:

```bash
$HOME/Desktop/tbx {{.Command}} {{.CommandArgs}}
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity.
Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue.
Then please proceed "System Preference", then open "Security & Privacy",
select "General" tab. You may find the message like:

> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk.
At second run, please hit button "Open" on the dialogue.

{{.CommandNote}}

## Options

{{.Options}}

Common options:

{{.CommonOptions}}

{{if .UseAuth}}
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
{{end}}

## Network configuration: Proxy

The executable automatically detects your proxy configuration from the environment.
However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port.
Currently, the executable doesn't support proxies which require authentication.

{{if .ReportAvailable }}
# Result

Report file path will be displayed last line of the command line output.
If you missed command line output, please see path below.
[job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

{{ $reports := .Reports }}
{{ range $name := .ReportNames }}
## Report: {{ $name }} 

Report files are generated in `{{$name}}.csv`, `{{$name}}.xlsx` and `{{$name}}.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `{{$name}}_0000.xlsx`, `{{$name}}_0001.xlsx`, `{{$name}}_0002.xlsx`...   

{{ index $reports $name }}
{{end}}
{{end}}