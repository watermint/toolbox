---
layout: command
title: Command `dropbox sign request signature list`
lang: en
---

# dropbox sign request signature list

List signatures of requests 

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
* Dropbox Sign: https://faq.hellosign.com/hc/en-us/articles/360035403131-HelloSign-API-accounts-and-how-to-find-your-API-key

## Auth scopes

| Description                                         |
|-----------------------------------------------------|
| Dropbox Sign: Process request as authenticated user |

# Authorization

For the first run, `tbx` will ask you an authentication with your DropboxSign account.
Log in to Dropbox Sign and copy the API key of your application from API Integration. Enter the copied API key into the application.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Please enter your credential(s).
```

# Installation

Please download the pre-compiled binary from [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are using Windows, please download the zip file like `tbx-xx.x.xxx-win.zip`. Then, extract the archive and place `tbx.exe` on the Desktop folder. 
The watermint toolbox can run from any path in the system if allowed by the system. But the instruction samples are using the Desktop folder. Please replace the path if you placed the binary other than the Desktop folder.

# Usage

This document uses the Desktop folder for command example.

## Run

Windows:
```
cd $HOME\Desktop
.\tbx.exe dropbox sign request signature list 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox sign request signature list 
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option        | Description                                                                                                                             | Default |
|---------------|-----------------------------------------------------------------------------------------------------------------------------------------|---------|
| `-account-id` | Which account to return SignatureRequests for. Must be a team member. Use `all` to indicate all team members. Defaults to your account. |         |
| `-peer`       | Account alias                                                                                                                           | default |

## Common options:

| Option             | Description                                                                                                                                           | Default              |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                                                                             |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                                                                                      | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited                                                                       | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint) (Options: low, normal)                                                                 | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage) (Options: low, normal, unlimited)                                                | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                                                                             | Number of processors |
| `-debug`           | Enable debug mode                                                                                                                                     | false                |
| `-experiment`      | Enable experimental feature(s).                                                                                                                       |                      |
| `-extra`           | Extra parameter file path                                                                                                                             |                      |
| `-lang`            | Display language (Options: auto, en, ja)                                                                                                              | auto                 |
| `-output`          | Output format (none/text/markdown/json) (Options: text, markdown, json, none)                                                                         | text                 |
| `-output-filter`   | Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON. |                      |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.                                                          |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy (Options: default, on_error, none)                                                                                             | default              |
| `-secure`          | Do not store tokens into a file                                                                                                                       | false                |
| `-skip-logging`    | Skip logging in the local storage                                                                                                                     | false                |
| `-verbose`         | Show current operations for more detail.                                                                                                              | false                |
| `-workspace`       | Workspace path                                                                                                                                        |                      |

# Results

Report file path will be displayed last line of the command line output. If you missed the command line output, please see path below. [job-id] will be the date/time of the run. Please see the latest job-id.

| OS      | Path pattern                                | Example                                                |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## Report: signatures

Signature data of a request.
The command will generate a report in three different formats. `signatures.csv`, `signatures.json`, and `signatures.xlsx`.

| Column                  | Description                                                                    |
|-------------------------|--------------------------------------------------------------------------------|
| signature_request_id    | The id of the SignatureRequest.                                                |
| signature_id            | Signature identifier.                                                          |
| requester_email_address | The email address of the initiator of the SignatureRequest.                    |
| title                   | The title the specified Account uses for the SignatureRequest.                 |
| subject                 | The subject in the email that was initially sent to the signers.               |
| message                 | The custom message in the email that was initially sent to the signers.        |
| created_at_rfc3339      | Time the signature request was created.                                        |
| expires_at_rfc3339      | The time when the signature request will expire unsigned signatures.           |
| is_complete             | Whether or not the SignatureRequest has been fully executed by all signers.    |
| is_declined             | Whether or not the SignatureRequest has been declined by a signer.             |
| signer_email_address    | The email address of the signer.                                               |
| signer_name             | The name of the signer.                                                        |
| signer_role             | The role of the signer.                                                        |
| order                   | If signer order is assigned this is the 0-based index for this signer.         |
| status_code             | The current status of the signature. eg: awaiting_signature, signed, declined. |
| decline_reason          | The reason provided by the signer for declining the request.                   |
| signed_at_rfc3339       | Time that the document was signed or empty.                                    |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `signatures_0000.xlsx`, `signatures_0001.xlsx`, `signatures_0002.xlsx`, ...


