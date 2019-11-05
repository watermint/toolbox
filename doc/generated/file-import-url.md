# file import url 

Import file from the URL

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
.\tbx.exe file import url 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx file import url 
```

## Options

| Option  | Description    | Default   |
|---------|----------------|-----------|
| `-path` | Path to import |           |
| `-peer` | Account alias  | {default} |
| `-url`  | URL            |           |

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

## Network configuration: Proxy

The executable automatically detects your proxy configuration from the environment.
However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port.
Currently, the executable doesn't support proxies which require authentication.

# Result

Report file path will be displayed last line of the command line output.
If you missed command line output, please see path below.
[job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## Report: import_url 

Report files are generated in `import_url.csv`, `import_url.xlsx` and `import_url.json` format.
In case of a report become large, report in `.xlsx` format will be split into several chunks
like `import_url_0000.xlsx`, `import_url_0001.xlsx`, `import_url_0002.xlsx`...   

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| id                      | A unique identifier for the file.                                                                      |
| tag                     | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                    | The last component of the path (including extension).                                                  |
| path_lower              | The lowercased full path in the user's Dropbox. This always starts with a slash.                       |
| path_display            | The cased path to be used for display purposes only.                                                   |
| client_modified         | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified         | The last time the file was modified on Dropbox.                                                        |
| revision                | A unique identifier for the current revision of a file.                                                |
| size                    | The file size in bytes.                                                                                |
| content_hash            | A hash of the file content.                                                                            |
| shared_folder_id        | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.   |
| parent_shared_folder_id |                                                                                                        |

