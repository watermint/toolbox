---
layout: command
title: services google mail filter batch add
---

# services google mail filter batch add

Batch adding/deleting labels with query 

# Security

`watermint toolbox` stores credentials into the file system. That is located at below path:

| OS      | Path                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

Please do not share those files to anyone including Dropbox support.
You can delete those files after use if you want to remove it. If you want to make sure removal of credentials, revoke application access from setting or the admin console.

Please see below help article for more detail:
* Google: https://support.google.com/accounts/answer/3466521

## Auth scopes

| Description                                      |
|--------------------------------------------------|
| Gmail: View and modify but not delete your email |
| Gmail: Manage your basic mail settings           |

# Authorization

For the first run, `tbx` will ask you an authentication with your Google account. Please copy the link and paste it into your browser. Then proceed to authorization. After authorization, Dropbox will show you an authorization code. Please copy that code and paste it to the `tbx`.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2021 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

1. Visit the URL for the auth dialogue:

https://accounts.google.com/o/oauth2/auth?client_id=xxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe services google mail filter batch add -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:
```
$HOME/Desktop/tbx services google mail filter batch add -file /PATH/TO/DATA_FILE.csv
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                        | Description                                                                                    | Default |
|-------------------------------|------------------------------------------------------------------------------------------------|---------|
| `-add-label-if-not-exist`     | Create a label if it is not exist.                                                             | false   |
| `-apply-to-existing-messages` | Apply labels to existing messages that satisfy the query.                                      | false   |
| `-file`                       | Path to data file                                                                              |         |
| `-peer`                       | Account alias                                                                                  | default |
| `-user-id`                    | The user's email address. The special value me can be used to indicate the authenticated user. | me      |

## Common options:

| Option            | Description                                                                               | Default              |
|-------------------|-------------------------------------------------------------------------------------------|----------------------|
| `-auto-open`      | Auto open URL or artifact folder                                                          | false                |
| `-bandwidth-kb`   | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited           | 0                    |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                            | normal               |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)                      | normal               |
| `-concurrency`    | Maximum concurrency for running operation                                                 | Number of processors |
| `-debug`          | Enable debug mode                                                                         | false                |
| `-experiment`     | Enable experimental feature(s).                                                           |                      |
| `-lang`           | Display language                                                                          | auto                 |
| `-output`         | Output format (none/text/markdown/json)                                                   | text                 |
| `-proxy`          | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want skip setting proxy. |                      |
| `-quiet`          | Suppress non-error messages, and make output readable by a machine (JSON format)          | false                |
| `-secure`         | Do not store tokens into a file                                                           | false                |
| `-verbose`        | Show current operations for more detail.                                                  | false                |
| `-workspace`      | Workspace path                                                                            |                      |

# File formats

## Format: File

Filter data to add.

| Column        | Description                                        | Example          |
|---------------|----------------------------------------------------|------------------|
| query         | Only return messages matching the specified query. | from:@google.com |
| add_labels    | Label names to add, separated by ';'               | my_label         |
| delete_labels | Label names to delete, separated by ';'            | my_label,INBOX   |

The first line is a header line. The program will accept a file without the header.
```
query,add_labels,delete_labels
from:@google.com,my_label,"my_label,INBOX"
```

# Results

Report file path will be displayed last line of the command line output. If you missed command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: filters

This report shows the transaction result.
The command will generate a report in three different formats. `filters.csv`, `filters.json`, and `filters.xlsx`.

| Column                | Description                                                         |
|-----------------------|---------------------------------------------------------------------|
| status                | Status of the operation                                             |
| reason                | Reason of failure or skipped operation                              |
| input.query           | Only return messages matching the specified query.                  |
| input.add_labels      | Label names to add, separated by ';'                                |
| input.delete_labels   | Label names to delete, separated by ';'                             |
| result.id             | Filter Id                                                           |
| result.criteria_query | Filter criteria: Only return messages matching the specified query. |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `filters_0000.xlsx`, `filters_0001.xlsx`, `filters_0002.xlsx`, ...

## Report: messages

Message resource
The command will generate a report in three different formats. `messages.csv`, `messages.json`, and `messages.xlsx`.

| Column    | Description                                  |
|-----------|----------------------------------------------|
| id        | The immutable ID of the message.             |
| thread_id | The ID of the thread the message belongs to. |
| date      | Date                                         |
| subject   | Subject                                      |
| to        | To                                           |
| cc        | Cc                                           |
| from      | From                                         |
| reply_to  | Reply-To                                     |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report become large, a report in `.xlsx` format will be split into several chunks like follows; `messages_0000.xlsx`, `messages_0001.xlsx`, `messages_0002.xlsx`, ...

# Proxy configuration

The executable automatically detects your proxy configuration from the environment. However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port. Currently, the executable doesn't support proxies which require authentication.


