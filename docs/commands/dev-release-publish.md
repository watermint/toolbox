---
layout: command
title: Command `dev release publish`
lang: en
---

# dev release publish

Publish release 

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
* GitHub: https://developer.github.com/apps/managing-oauth-apps/deleting-an-oauth-app/

## Auth scopes

| Description                                                                                                                                                                                                                                                                                                                                                    |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GitHub: Grants full access to repositories, including private repositories. That includes read/write access to code, commit statuses, repository and organization projects, invitations, collaborators, adding team memberships, deployment statuses, and repository webhooks for repositories and organizations. Also grants ability to manage user projects. |

# Authorization

For the first run, `tbx` will ask you an authentication with your GitHub account.
Press Enter to launch the browser. The service then performs the authorization and the application receives the results. You can close the browser window when you see the authentication success message.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Opening the authorization URL:
https://github.com/login/oauth/authorize?client_id=xxxxxxxxxxxxxxxxxxxx&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&scope=repo&state=xxxxxxxx

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
.\tbx.exe dev release publish -artifact-path /LOCAL/PATH/TO/ARTIFACT
```

macOS, Linux:
```
$HOME/Desktop/tbx dev release publish -artifact-path /LOCAL/PATH/TO/ARTIFACT
```

Note for macOS Catalina 10.15 or above: macOS verifies Developer identity. Currently, `tbx` is not ready for it. Please select "Cancel" on the first dialogue. Then please proceed "System Preference", then open "Security & Privacy", select "General" tab.
You may find the message like:
> "tbx" was blocked from use because it is not from an identified developer.

And you may find the button "Allow Anyway". Please hit the button with your risk. At second run, please hit button "Open" on the dialogue.

## Options:

| Option                  | Description                                                      | Default          |
|-------------------------|------------------------------------------------------------------|------------------|
| `-artifact-path`        | Path to artifacts                                                |                  |
| `-branch`               | Target branch                                                    | main             |
| `-conn-github`          | Account alias                                                    | default          |
| `-executable-name`      | The name of the executable file to be published.                 | tbx              |
| `-homebrew-repo-branch` | The branch of the Homebrew tap repository to use for publishing. | master           |
| `-homebrew-repo-name`   | The name of the Homebrew tap repository.                         | homebrew-toolbox |
| `-homebrew-repo-owner`  | The owner of the Homebrew tap repository.                        | watermint        |
| `-repo-name`            | The name of the repository to publish the release to.            | toolbox          |
| `-repo-owner`           | The owner of the repository to publish the release to.           | watermint        |
| `-skip-tests`           | Skip end to end tests.                                           | false            |

## Common options:

| Option             | Description                                                                                                                                           | Default              |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `-auth-database`   | Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)                                                                             |                      |
| `-auto-open`       | Auto open URL or artifact folder                                                                                                                      | false                |
| `-bandwidth-kb`    | Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited                                                                       | 0                    |
| `-budget-memory`   | Memory budget (limits some feature to reduce memory footprint)                                                                                        | normal               |
| `-budget-storage`  | Storage budget (limits logs or some feature to reduce storage usage)                                                                                  | normal               |
| `-concurrency`     | Maximum concurrency for running operation                                                                                                             | Number of processors |
| `-debug`           | Enable debug mode                                                                                                                                     | false                |
| `-experiment`      | Enable experimental feature(s).                                                                                                                       |                      |
| `-extra`           | Extra parameter file path                                                                                                                             |                      |
| `-lang`            | Display language                                                                                                                                      | auto                 |
| `-output`          | Output format (none/text/markdown/json)                                                                                                               | text                 |
| `-output-filter`   | Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON. |                      |
| `-proxy`           | HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.                                                          |                      |
| `-quiet`           | Suppress non-error messages, and make output readable by a machine (JSON format)                                                                      | false                |
| `-retain-job-data` | Job data retain policy                                                                                                                                | default              |
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

## Report: commit

Commit information
The command will generate a report in three different formats. `commit.csv`, `commit.json`, and `commit.xlsx`.

| Column | Description        |
|--------|--------------------|
| sha    | SHA1 of the commit |
| url    | URL of the commit  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `commit_0000.xlsx`, `commit_0001.xlsx`, `commit_0002.xlsx`, ...

## Report: content

Content metadata
The command will generate a report in three different formats. `content.csv`, `content.json`, and `content.xlsx`.

| Column | Description     |
|--------|-----------------|
| type   | Type of content |
| name   | Name            |
| path   | Path            |
| sha    | SHA1            |
| size   | Size            |
| target | Symlink target  |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `content_0000.xlsx`, `content_0001.xlsx`, `content_0002.xlsx`, ...

## Report: result

Recipe test result
The command will generate a report in three different formats. `result.csv`, `result.json`, and `result.xlsx`.

| Column          | Description                   |
|-----------------|-------------------------------|
| path            | Path to the recipe            |
| name            | Name of the recipe            |
| skip            | True if the test skipped      |
| timeout_enabled | True if timeout mode enabled  |
| use_mock        | True if use mock mode enabled |
| timeout         | True if the test timeout      |
| duration        | Test duration in milliseconds |
| no_error        | True if no error received     |
| error           | Error                         |

If you run with `-budget-memory low` option, the command will generate only JSON format report.

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `result_0000.xlsx`, `result_0001.xlsx`, `result_0002.xlsx`, ...


