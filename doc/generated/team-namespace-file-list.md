# team namespace file list 

List all files and folders of the team namespaces



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
.\tbx.exe team namespace file list 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team namespace file list 
```



## Options

| Option                   | Description                                                   | Default   |
|--------------------------|---------------------------------------------------------------|-----------|
| `-include-deleted`       | If true, deleted file or folder will be returned              | false     |
| `-include-media-info`    | If true, media info is set for photo and video in json report | false     |
| `-include-member-folder` | If true, include team member folders                          | false     |
| `-include-shared-folder` | If true, include shared folders                               | true      |
| `-include-team-folder`   | If true, include team folders                                 | true      |
| `-name`                  | List only for the folder matched to the name                  |           |
| `-peer`                  | Account alias                                                 | {default} |


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



## Report: namespace_file 

Report files are generated in `namespace_file.csv`, `namespace_file.xlsx` and `namespace_file.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`...   

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| namespace_type          | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)             |
| namespace_id            | The ID of this namespace.                                                                              |
| namespace_name          | The name of this namespace                                                                             |
| namespace_member_email  | If this is a team member or app folder, the email address of the owning team member.                   |
| file_id                 | A unique identifier for the file.                                                                      |
| tag                     | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                    | The last component of the path (including extension).                                                  |
| path_display            | The cased path to be used for display purposes only.                                                   |
| client_modified         | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified         | The last time the file was modified on Dropbox.                                                        |
| revision                | A unique identifier for the current revision of a file.                                                |
| size                    | The file size in bytes.                                                                                |
| content_hash            | A hash of the file content.                                                                            |
| shared_folder_id        | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.   |
| parent_shared_folder_id | Set if the folder is contained by a shared folder.                                                     |



