# {{.Command}} 

{{.CommandTitle}} {{.CommandRemarks}}

{{.CommandDesc}}

{{if .UseAuth}}
# {{ msg "doc.command.security.header" }}

{{ msg "doc.command.security.credential_location" }}

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

{{ msg "doc.command.security.credential_remarks" }}
{{ msg "doc.command.security.how_to_remove_it" }}

{{ msg "doc.command.security.how_to_help_center" }}
{{ if .UseAuthPersonal }}* {{ msg "doc.command.security.how_to_help_center_individual" }}: https://help.dropbox.com/installs-integrations/third-party/third-party-apps{{end}}{{ if .UseAuthBusiness }}* {{ msg "doc.command.security.how_to_help_center_business" }}: https://help.dropbox.com/teams-admins/admin/app-integrations{{end}}

{{ msg "doc.command.security.scopes" }}
{{ range $scope := .AuthScopes }}{{ with (eq $scope "business_info")}}* Dropbox Business Information access{{end}}{{ with (eq $scope "business_file")}}* Dropbox Business File access{{end}}{{ with (eq $scope "business_mgmt")}}* Dropbox Business management{{end}}{{ with (eq $scope "business_audit")}}* Dropbox Business Auditing{{end}}{{ with (eq $scope "user_file")}}* Dropbox Full access{{end}}{{end}}
{{end}}

# {{ msg "doc.command.usage.header" }}

{{ msg "doc.command.usage.usage_remarks" }}

## {{ msg "doc.command.usage.run.header" }}

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe {{.Command}} {{.CommandArgs}}
```

macOS, Linux:

```bash
$HOME/Desktop/tbx {{.Command}} {{.CommandArgs}}
```

{{ msg "doc.command.usage.run.catalina_remarks1" }}
{{ msg "doc.command.usage.run.catalina_remarks2" }}

{{ msg "doc.command.usage.run.catalina_remarks3" }}

{{.CommandNote}}

## {{ msg "doc.command.usage.options" }}

{{.Options}}

{{ msg "doc.command.usage.options.common_options" }}

{{.CommonOptions}}

{{if .FeedAvailable }}
# {{ msg "doc.command.file_formats.header" }}

{{ $feeds := .Feeds }}
{{ $feedSamples := .FeedSamples }}
{{ range $name := .FeedNames }}
## {{ msg "doc.command.file_formats.format.header" }}: {{ $name }} 

{{ index $feeds $name }}

{{ msg "doc.command.file_formats.format.header_line" }}

```csv
{{ index $feedSamples $name }}```
{{end}}
{{end}}

{{if .UseAuth}}
# {{ msg "doc.command.auth.header" }}

{{ msg "doc.command.auth.description" }}

{{ .AuthExample }}
{{end}}

# {{ msg "doc.command.proxy.header" }}

{{ msg "doc.command.proxy.description" }}

{{if .ReportAvailable }}
# {{ msg "doc.command.report.header" }}

{{ msg "doc.command.report.file_location" }}

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

{{ $reports := .Reports }}
{{ $reportDesc := .ReportDesc }}
{{ range $name := .ReportNames }}
## {{ msg "doc.command.report.report.header" }}: {{ $name }} 
{{ index $reportDesc $name }}
{{ msg "doc.command.report.report.format_description" }}
* `{{$name}}.csv`
* `{{$name}}.xlsx`
* `{{$name}}.json`

{{ msg "doc.command.report.report.low_memory_option" }}

{{ msg "doc.command.report.report.xlsx_remarks" }}
`{{$name}}_0000.xlsx`, `{{$name}}_0001.xlsx`, `{{$name}}_0002.xlsx`...   

{{ index $reports $name }}
{{end}}
{{end}}
