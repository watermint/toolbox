---
Title: watermint toolbox, user manual
URL: https://toolbox.watermint.org/home.md
---

# watermint toolbox, user manual

watermint toolbox, user manual. This document describes the commands and options available in the watermint toolbox.

---
Title: Common Options for All Commands
URL: https://toolbox.watermint.org/commands/common-options.md
---

## Common command options

**-auth-database**
: Custom path to auth database (default: $HOME/.toolbox/secrets/secrets.db)

**-auto-open**
: Auto open URL or artifact folder. Default: false

**-bandwidth-kb**
: Bandwidth limit in K bytes per sec for upload/download content. 0 for unlimited. Default: 0

**-budget-memory**
: Memory budget (limits some feature to reduce memory footprint). Options: low, normal. Default: normal

**-budget-storage**
: Storage budget (limits logs or some feature to reduce storage usage). Options: low, normal, unlimited. Default: normal

**-concurrency**
: Maximum concurrency for running operation. Default: Number of processors

**-debug**
: Enable debug mode. Default: false

**-experiment**
: Enable experimental feature(s).

**-extra**
: Extra parameter file path

**-lang**
: Display language. Options: auto, en, ja. Default: auto

**-output**
: Output format (none/text/markdown/json). Options: text, markdown, json, none. Default: text

**-output-filter**
: Output filter query (jq syntax). The output of the report is filtered using jq syntax. This option is only applied when the report is output as JSON.

**-proxy**
: HTTP/HTTPS proxy (hostname:port). Please specify `DIRECT` if you want to skip setting proxy.

**-quiet**
: Suppress non-error messages, and make output readable by a machine (JSON format). Default: false

**-retain-job-data**
: Job data retain policy. Options: default, on_error, none. Default: default

**-secure**
: Do not store tokens into a file. Default: false

**-skip-logging**
: Skip logging in the local storage. Default: false

**-verbose**
: Show current operations for more detail.. Default: false

**-workspace**
: Workspace path

## Commands

---
Title: license
URL: https://toolbox.watermint.org/commands/license.md
---

# license

Show license information 

Display detailed license information for the watermint toolbox and all its components. This includes open source licenses, copyright notices, and third-party dependencies used in the application.

# Usage

This document uses the Desktop folder for command example.
```
tbx license 
```

---
Title: version
URL: https://toolbox.watermint.org/commands/version.md
---

# version

Show version 

Display version information for the watermint toolbox including build date, Git commit hash, and component versions. This is useful for troubleshooting, bug reports, and ensuring you have the latest version.

# Usage

This document uses the Desktop folder for command example.
```
tbx version 
```

# Results

## Report: versions

Components and version information.
The command will generate a report in three different formats. `versions.csv`, `versions.json`, and `versions.xlsx`.

| Column    | Description |
|-----------|-------------|
| key       | Key         |
| component | Component   |
| version   | Version     |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `versions_0000.xlsx`, `versions_0001.xlsx`, `versions_0002.xlsx`, ...

---
Title: config auth delete
URL: https://toolbox.watermint.org/commands/config/auth/delete.md
---

# config auth delete

Delete existing auth credential 

Remove stored authentication credentials for a specific service account. This is useful when you need to revoke access, change accounts, or clean up old authentication tokens. The command requires both the application key name and peer name to identify the credential to delete.

# Usage

This document uses the Desktop folder for command example.
```
tbx config auth delete -key-name KEY_NAME -peer-name PEER_NAME
```

## Options:

**-key-name**
: Application key name

**-peer-name**
: Peer name

# Results

## Report: deleted

Authentication credential data
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column      | Description      |
|-------------|------------------|
| key_name    | Application name |
| scope       | Auth scope       |
| peer_name   | Peer name        |
| description | Description      |
| timestamp   | Timestamp        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

---
Title: config auth list
URL: https://toolbox.watermint.org/commands/config/auth/list.md
---

# config auth list

List all auth credentials 

Display all stored authentication credentials and their details including application names, scopes, peer names, and timestamps. This is useful for auditing access, managing multiple accounts, and understanding which services you're authenticated with.

# Usage

This document uses the Desktop folder for command example.
```
tbx config auth list 
```

# Results

## Report: entity

Authentication credential data
The command will generate a report in three different formats. `entity.csv`, `entity.json`, and `entity.xlsx`.

| Column      | Description      |
|-------------|------------------|
| key_name    | Application name |
| scope       | Auth scope       |
| peer_name   | Peer name        |
| description | Description      |
| timestamp   | Timestamp        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `entity_0000.xlsx`, `entity_0001.xlsx`, `entity_0002.xlsx`, ...

---
Title: config feature disable
URL: https://toolbox.watermint.org/commands/config/feature/disable.md
---

# config feature disable

Disable a feature. 

Turn off a specific feature in the watermint toolbox configuration. Features control various aspects of the application's behavior, performance settings, and experimental functionality. Disabling features can help with troubleshooting or reverting to previous behavior.

# Usage

This document uses the Desktop folder for command example.
```
tbx config feature disable -key FEATURE
```

## Options:

**-key**
: Feature key.

---
Title: config feature enable
URL: https://toolbox.watermint.org/commands/config/feature/enable.md
---

# config feature enable

Enable a feature. 

Turn on a specific feature in the watermint toolbox configuration. Features control various aspects of the application's behavior, performance settings, and experimental functionality. Enabling features allows you to access new capabilities or modify application behavior.

# Usage

This document uses the Desktop folder for command example.
```
tbx config feature enable -key FEATURE
```

## Options:

**-key**
: Feature key.

---
Title: config feature list
URL: https://toolbox.watermint.org/commands/config/feature/list.md
---

# config feature list

List available optional features. 

Display all available optional features in the watermint toolbox with their descriptions, current status, and configuration details. This is useful for understanding what functionality can be enabled or disabled, and for managing feature preferences.

# Usage

This document uses the Desktop folder for command example.
```
tbx config feature list 
```

---
Title: config license install
URL: https://toolbox.watermint.org/commands/config/license/install.md
---

# config license install

Install a license key 

Install and activate a license key for the watermint toolbox. License keys may be required for certain features, premium functionality, or commercial usage. This command stores the license key securely and validates its authenticity.

# Usage

This document uses the Desktop folder for command example.
```
tbx config license install -key LICENSE_KEY
```

## Options:

**-key**
: License key

---
Title: config license list
URL: https://toolbox.watermint.org/commands/config/license/list.md
---

# config license list

List available license keys 

Display all installed license keys and their details including expiration dates, enabled features, and status. This is useful for managing multiple licenses, checking license validity, and understanding what functionality is available.

# Usage

This document uses the Desktop folder for command example.
```
tbx config license list 
```

# Results

## Report: keys

License key summary
The command will generate a report in three different formats. `keys.csv`, `keys.json`, and `keys.xlsx`.

| Column           | Description                         |
|------------------|-------------------------------------|
| key              | License key                         |
| expiration_date  | Expiration date                     |
| valid            | True if the license key is valid    |
| licensee_name    | Licensee name                       |
| licensee_email   | Licensee email                      |
| licensed_recipes | Recipes enabled by this license key |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `keys_0000.xlsx`, `keys_0001.xlsx`, `keys_0002.xlsx`, ...

---
Title: log api job
URL: https://toolbox.watermint.org/commands/log/api/job.md
---

# log api job

Show statistics of the API log of the job specified by the job ID 

Analyze and display API call statistics for a specific job execution. This includes request counts, response times, error rates, and endpoint usage patterns. Useful for performance analysis, debugging API issues, and understanding application behavior during command execution.

# Usage

This document uses the Desktop folder for command example.
```
tbx log api job 
```

## Options:

**-full-url**
: Show full URL. Default: false

**-interval-second**
: Interval in seconds for the time series. Default: 3600

**-job-id**
: Job ID

# Results

## Report: latencies

Latency
The command will generate a report in three different formats. `latencies.csv`, `latencies.json`, and `latencies.xlsx`.

| Column     | Description        |
|------------|--------------------|
| url        | URL                |
| code       | Response code      |
| population | Number of requests |
| mean       | Mean               |
| median     | Median             |
| p_50       | Percentile 50      |
| p_70       | Percentile 70      |
| p_90       | Percentile 90      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `latencies_0000.xlsx`, `latencies_0001.xlsx`, `latencies_0002.xlsx`, ...

## Report: population

Number of requests
The command will generate a report in three different formats. `population.csv`, `population.json`, and `population.xlsx`.

| Column     | Description        |
|------------|--------------------|
| url        | URL                |
| code       | Response code      |
| population | Number of requests |
| proportion | Proportion         |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `population_0000.xlsx`, `population_0001.xlsx`, `population_0002.xlsx`, ...

## Report: time_series

Time series summary
The command will generate a report in three different formats. `time_series.csv`, `time_series.json`, and `time_series.xlsx`.

| Column     | Description                              |
|------------|------------------------------------------|
| time       | Time                                     |
| url        | URL                                      |
| code_2xx   | Number of requests with 2xx              |
| code_3xx   | Number of requests with 3xx              |
| code_4xx   | Number of requests with 4xx (except 429) |
| code_429   | Number of requests with 429              |
| code_5xx   | Number of requests with 5xx              |
| code_other | Number of requests with other            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `time_series_0000.xlsx`, `time_series_0001.xlsx`, `time_series_0002.xlsx`, ...

---
Title: log api name
URL: https://toolbox.watermint.org/commands/log/api/name.md
---

# log api name

Show statistics of the API log of the job specified by the job name 

Analyze and display API call statistics for jobs identified by command name rather than job ID. This allows you to aggregate statistics across multiple executions of the same command, helping identify patterns and performance trends over time.

# Usage

This document uses the Desktop folder for command example.
```
tbx log api name -name JOB_NAME
```

## Options:

**-full-url**
: Show full URL. Default: false

**-interval-second**
: Interval in seconds for the time series. Default: 3600

**-name**
: Job command line path (e.g. `dropbox team member list`)

# Results

## Report: latencies

Latency
The command will generate a report in three different formats. `latencies.csv`, `latencies.json`, and `latencies.xlsx`.

| Column     | Description        |
|------------|--------------------|
| url        | URL                |
| code       | Response code      |
| population | Number of requests |
| mean       | Mean               |
| median     | Median             |
| p_50       | Percentile 50      |
| p_70       | Percentile 70      |
| p_90       | Percentile 90      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `latencies_0000.xlsx`, `latencies_0001.xlsx`, `latencies_0002.xlsx`, ...

## Report: population

Number of requests
The command will generate a report in three different formats. `population.csv`, `population.json`, and `population.xlsx`.

| Column     | Description        |
|------------|--------------------|
| url        | URL                |
| code       | Response code      |
| population | Number of requests |
| proportion | Proportion         |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `population_0000.xlsx`, `population_0001.xlsx`, `population_0002.xlsx`, ...

## Report: time_series

Time series summary
The command will generate a report in three different formats. `time_series.csv`, `time_series.json`, and `time_series.xlsx`.

| Column     | Description                              |
|------------|------------------------------------------|
| time       | Time                                     |
| url        | URL                                      |
| code_2xx   | Number of requests with 2xx              |
| code_3xx   | Number of requests with 3xx              |
| code_4xx   | Number of requests with 4xx (except 429) |
| code_429   | Number of requests with 429              |
| code_5xx   | Number of requests with 5xx              |
| code_other | Number of requests with other            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `time_series_0000.xlsx`, `time_series_0001.xlsx`, `time_series_0002.xlsx`, ...

---
Title: log cat curl
URL: https://toolbox.watermint.org/commands/log/cat/curl.md
---

# log cat curl

Format capture logs as `curl` sample 

Convert API request logs into equivalent curl commands that can be executed independently. This is extremely useful for debugging API issues, reproducing requests outside of the toolbox, sharing examples with support, or creating test scripts.

# Usage

This document uses the Desktop folder for command example.
```
tbx log cat curl 
```

## Options:

**-buffer-size**
: Buffer size for each request. Default: 65536

**-record**
: Give a record of capture log file via command line option

---
Title: log cat job
URL: https://toolbox.watermint.org/commands/log/cat/job.md
---

# log cat job

Retrieve logs of specified Job ID 

Extract and display log files for a specific job execution identified by its Job ID. This includes debug logs, API capture logs, error messages, and system information. Essential for troubleshooting failed executions and analyzing job performance.

# Usage

This document uses the Desktop folder for command example.
```
tbx log cat job -id JOB_ID
```

## Options:

**-id**
: Job ID

**-kind**
: Kind of log. Options:.   • toolbox (kind: toolbox).   • capture (kind: capture).   • summary (kind: summary).   • recipe (kind: recipe).   • result (kind: result). Default: toolbox

**-path**
: Path to the workspace

---
Title: log cat kind
URL: https://toolbox.watermint.org/commands/log/cat/kind.md
---

# log cat kind

Concatenate and print logs of specified log kind 

# Usage

This document uses the Desktop folder for command example.
```
tbx log cat kind 
```

## Options:

**-kind**
: Log kind.. Options:.   • toolbox (kind: toolbox).   • capture (kind: capture).   • summary (kind: summary).   • recipe (kind: recipe).   • result (kind: result). Default: toolbox

**-path**
: Path to workspace.

---
Title: log cat last
URL: https://toolbox.watermint.org/commands/log/cat/last.md
---

# log cat last

Print the last job log files 

# Usage

This document uses the Desktop folder for command example.
```
tbx log cat last 
```

## Options:

**-kind**
: Log kind. Options:.   • toolbox (kind: toolbox).   • capture (kind: capture).   • summary (kind: summary).   • recipe (kind: recipe).   • result (kind: result). Default: toolbox

**-path**
: Path to workspace.

---
Title: log job archive
URL: https://toolbox.watermint.org/commands/log/job/archive.md
---

# log job archive

Archive jobs 

# Usage

This document uses the Desktop folder for command example.
```
tbx log job archive 
```

## Options:

**-days**
: Target days old. Default: 7

**-path**
: Path to the workspace

---
Title: log job delete
URL: https://toolbox.watermint.org/commands/log/job/delete.md
---

# log job delete

Delete old job history 

# Usage

This document uses the Desktop folder for command example.
```
tbx log job delete 
```

## Options:

**-days**
: Target days old. Default: 28

**-path**
: Path to the workspace

---
Title: log job list
URL: https://toolbox.watermint.org/commands/log/job/list.md
---

# log job list

Show job history 

# Usage

This document uses the Desktop folder for command example.
```
tbx log job list 
```

## Options:

**-path**
: Path to workspace

# Results

## Report: log

This report shows a list of job histories.
The command will generate a report in three different formats. `log.csv`, `log.json`, and `log.xlsx`.

| Column      | Description   |
|-------------|---------------|
| job_id      | Job ID        |
| app_version | App version   |
| recipe_name | Command       |
| time_start  | Time Started  |
| time_finish | Time Finished |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `log_0000.xlsx`, `log_0001.xlsx`, `log_0002.xlsx`, ...

---
Title: util archive unzip
URL: https://toolbox.watermint.org/commands/util/archive/unzip.md
---

# util archive unzip

Extract the zip archive file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util archive unzip -in /LOCAL/PATH/TO/ARCHIVE.zip -out /LOCAL/PATH/TO/EXTRACT
```

## Options:

**-in**
: Zip archive file path

**-out**
: Path to extract files

---
Title: util archive zip
URL: https://toolbox.watermint.org/commands/util/archive/zip.md
---

# util archive zip

Compress target files into the zip archive 

# Usage

This document uses the Desktop folder for command example.
```
tbx util archive zip -out /LOCAL/PATH/TO/ARCHIVE.zip -target /LOCAL/PATH/TO/COMPRESS
```

## Options:

**-comment**
: Comment for the zip archive

**-out**
: Zip archive file path

**-target**
: Target path to compress

---
Title: util cert selfsigned
URL: https://toolbox.watermint.org/commands/util/cert/selfsigned.md
---

# util cert selfsigned

Generate self-signed certificate and key 

# Usage

This document uses the Desktop folder for command example.
```
tbx util cert selfsigned -out /LOCAL/PATH/TO/GENERATE_CERT_AND_KEY
```

## Options:

**-days**
: Number of validity days of the certificate. Default: 365

**-out**
: Output folder path

---
Title: util database exec
URL: https://toolbox.watermint.org/commands/util/database/exec.md
---

# util database exec

Execute query on SQLite3 database file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util database exec -file /LOCAL/PATH/DATA.sql -sql SQL
```

## Options:

**-file**
: Path to data file

**-sql**
: Query

---
Title: util database query
URL: https://toolbox.watermint.org/commands/util/database/query.md
---

# util database query

Query SQLite3 database 

# Usage

This document uses the Desktop folder for command example.
```
tbx util database query -file /LOCAL/PATH/DATA.sql -sql SQL
```

## Options:

**-file**
: Path to data file

**-result**
: Query result

**-result-format**
: Output format

**-sql**
: Query

# Grid data output of the command

## Grid data output: Result

Query result

---
Title: util date today
URL: https://toolbox.watermint.org/commands/util/date/today.md
---

# util date today

Display current date 

# Usage

This document uses the Desktop folder for command example.
```
tbx util date today 
```

## Options:

**-offset**
: Offset (day). Default: 0

**-utc**
: Use UTC. Default: false

---
Title: util datetime now
URL: https://toolbox.watermint.org/commands/util/datetime/now.md
---

# util datetime now

Display current date/time 

# Usage

This document uses the Desktop folder for command example.
```
tbx util datetime now 
```

## Options:

**-offset-day**
: Offset (day). Default: 0

**-offset-hour**
: Offset (hour). Default: 0

**-offset-min**
: Offset (min). Default: 0

**-offset-sec**
: Offset (sec). Default: 0

**-utc**
: Use UTC. Default: false

---
Title: util decode base32
URL: https://toolbox.watermint.org/commands/util/decode/base32.md
---

# util decode base32

Decode text from Base32 (RFC 4648) format 

# Usage

This document uses the Desktop folder for command example.
```
tbx util decode base32 -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-no-padding**
: No padding. Default: false

**-text**
: Text

# Text inputs

## Text input: Text

Text to decode

---
Title: util decode base64
URL: https://toolbox.watermint.org/commands/util/decode/base64.md
---

# util decode base64

Decode text from Base64 (RFC 4648) format 

# Usage

This document uses the Desktop folder for command example.
```
tbx util decode base64 -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-no-padding**
: No padding. Default: false

**-text**
: Text

# Text inputs

## Text input: Text

Text to decode

---
Title: util desktop open
URL: https://toolbox.watermint.org/commands/util/desktop/open.md
---

# util desktop open

Open a file or folder with the default application 

# Usage

This document uses the Desktop folder for command example.
```
tbx util desktop open -path /LOCAL/PATH/TO/open
```

## Options:

**-path**
: Path to the file or folder to open

---
Title: util encode base32
URL: https://toolbox.watermint.org/commands/util/encode/base32.md
---

# util encode base32

Encode text into Base32 (RFC 4648) format 

# Usage

This document uses the Desktop folder for command example.
```
tbx util encode base32 -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-no-padding**
: No padding. Default: false

**-text**
: Text

# Text inputs

## Text input: Text

Text to encode

---
Title: util encode base64
URL: https://toolbox.watermint.org/commands/util/encode/base64.md
---

# util encode base64

Encode text into Base64 (RFC 4648) format 

# Usage

This document uses the Desktop folder for command example.
```
tbx util encode base64 -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-no-padding**
: No padding. Default: false

**-text**
: Text

# Text inputs

## Text input: Text

Text to encode

---
Title: util feed json
URL: https://toolbox.watermint.org/commands/util/feed/json.md
---

# util feed json

Load feed from the URL and output the content as JSON 

# Usage

This document uses the Desktop folder for command example.
```
tbx util feed json -url URL
```

## Options:

**-compact**
: Compact output. Default: false

**-url**
: URL of the feed

---
Title: util file hash
URL: https://toolbox.watermint.org/commands/util/file/hash.md
---

# util file hash

File Hash 

# Usage

This document uses the Desktop folder for command example.
```
tbx util file hash -file /LOCAL/PATH/TO/DIGEST
```

## Options:

**-algorithm**
: Hash algorithm (md5/sha1/sha256). Options:.   • md5 (algorithm: md5).   • sha1 (algorithm: sha1).   • sha256 (algorithm: sha256). Default: sha1

**-file**
: Path to the file to create digest

---
Title: util git clone
URL: https://toolbox.watermint.org/commands/util/git/clone.md
---

# util git clone

Clone git repository 

# Usage

This document uses the Desktop folder for command example.
```
tbx util git clone -local-path /LOCAL/PATH/TO/CLONE -url https://github.com/username/repository.git
```

## Options:

**-local-path**
: Local path to clone repository

**-reference**
: Reference name

**-remote-name**
: Name of the remote. Default: origin

**-url**
: Git repository URL

---
Title: util image exif
URL: https://toolbox.watermint.org/commands/util/image/exif.md
---

# util image exif

Print EXIF metadata of image file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util image exif -file /LOCAL/PATH/TO/IMG.jpg
```

## Options:

**-file**
: Path to data file

# Results

## Report: metadata

EXIF data
The command will generate a report in three different formats. `metadata.csv`, `metadata.json`, and `metadata.xlsx`.

| Column             | Description                                                                                          |
|--------------------|------------------------------------------------------------------------------------------------------|
| date_time_original | The date and time when the original image data was generated                                         |
| date_time          | The date and time of image creation. In Exif standard, it is the date and time the file was changed. |
| make               | The name of the manufacturer                                                                         |
| model              | The model name or model number                                                                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `metadata_0000.xlsx`, `metadata_0001.xlsx`, `metadata_0002.xlsx`, ...

---
Title: util image placeholder
URL: https://toolbox.watermint.org/commands/util/image/placeholder.md
---

# util image placeholder

Create placeholder image 

# Usage

This document uses the Desktop folder for command example.
```
tbx util image placeholder -path /LOCAL/PATH/TO/save.png
```

## Options:

**-color**
: Background color. Default: white

**-font-path**
: Path to TrueType font (required if you need to draw text)

**-font-size**
: Font size. Default: 12

**-height**
: Height (pixels). Default: 400

**-path**
: Path to export generated image

**-text**
: Text if you need

**-text-align**
: Text alignment. Options:.   • left (textalign: left).   • center (textalign: center).   • right (textalign: right). Default: left

**-text-color**
: Text color. Default: black

**-text-position**
: Text position. Default: center

**-width**
: Width (pixels). Default: 640

---
Title: util json query
URL: https://toolbox.watermint.org/commands/util/json/query.md
---

# util json query

Query JSON data 

Please refer to [jq Manual](https://jqlang.github.io/jq/manual/) for the syntax (some features are not supported).

# Usage

This document uses the Desktop folder for command example.
```
tbx util json query -path /LOCAL/PATH/TO/DATA.json -query QUERY
```

## Options:

**-compact**
: Compact output. Default: false

**-path**
: File path

**-query**
: Query string

# Text inputs

## Text input: Path

The path to the JSON file

---
Title: util net download
URL: https://toolbox.watermint.org/commands/util/net/download.md
---

# util net download

Download a file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util net download -out /LOCAL/PATH/TO/STORE -url URL_TO_DOWNLOAD
```

## Options:

**-out**
: Local path to store

**-url**
: URL to download

---
Title: util qrcode create
URL: https://toolbox.watermint.org/commands/util/qrcode/create.md
---

# util qrcode create

Create a QR code image file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util qrcode create -out /LOCAL/PATH/TO/OUT.png -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-error-correction-level**
: Error correction level (l/m/q/h).. Options:.   • l (errorcorrectionlevel: l).   • m (errorcorrectionlevel: m).   • q (errorcorrectionlevel: q).   • h (errorcorrectionlevel: h). Default: m

**-mode**
: QR code encoding mode. Options:.   • auto (mode: auto).   • numeric (mode: numeric).   • alpha_numeric (mode: alpha_numeric).   • unicode (mode: unicode). Default: auto

**-out**
: Output path with file name

**-size**
: Image resolution (pixels). Default: 256

**-text**
: Text data

# Text inputs

## Text input: Text

Text

---
Title: util qrcode wifi
URL: https://toolbox.watermint.org/commands/util/qrcode/wifi.md
---

# util qrcode wifi

Generate QR code for WIFI configuration 

# Usage

This document uses the Desktop folder for command example.
```
tbx util qrcode wifi -out /LOCAL/PATH/TO/OUT.png -ssid SSID
```

## Options:

**-error-correction-level**
: Error correction level (l/m/q/h).. Options:.   • l (errorcorrectionlevel: l).   • m (errorcorrectionlevel: m).   • q (errorcorrectionlevel: q).   • h (errorcorrectionlevel: h). Default: m

**-hidden**
: `true` if an SSID is hidden. `false` if an SSID is visible.. Options:.   •  (hidden: ).   • true (hidden: true).   • false (hidden: false)

**-mode**
: QR code encoding mode. Options:.   • auto (mode: auto).   • numeric (mode: numeric).   • alpha_numeric (mode: alpha_numeric).   • unicode (mode: unicode). Default: auto

**-network-type**
: Network type.. Options:.   • WPA.   • WEP.   •  (networktype: ). Default: WPA

**-out**
: Output path with file name

**-size**
: Image resolution (pixels). Default: 256

**-ssid**
: Network SSID

---
Title: util release install
URL: https://toolbox.watermint.org/commands/util/release/install.md
---

# util release install

Download & install watermint toolbox to the path 

# Usage

This document uses the Desktop folder for command example.
```
tbx util release install -path /LOCAL/PATH/TO/INSTALL
```

## Options:

**-accept-license-agreement**
: Accept the target release's license agreement. Default: false

**-path**
: Path to install

**-peer**
: Account alias. Default: default

**-release**
: Release tag name. Default: latest

---
Title: util table format xlsx
URL: https://toolbox.watermint.org/commands/util/table/format/xlsx.md
---

# util table format xlsx

Formatting xlsx file into text 

# Usage

This document uses the Desktop folder for command example.
```
tbx util table format xlsx -sheet SHEET_NAME -dest /LOCAL/PATH/TO/out.txt -template /LOCAL/PATH/TO/template.txt -source /LOCAL/PATH/TO/source.xlsx
```

## Options:

**-dest**
: Destination file path

**-position**
: Start position of table. Default: A1

**-sheet**
: Sheet name

**-source**
: Data source xlsx file

**-template**
: Template file

---
Title: util text case down
URL: https://toolbox.watermint.org/commands/util/text/case/down.md
---

# util text case down

Print lower case text 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text case down -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-text**
: Text

# Text inputs

## Text input: Text

Text to change case

---
Title: util text case up
URL: https://toolbox.watermint.org/commands/util/text/case/up.md
---

# util text case up

Print upper case text 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text case up -text /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-text**
: Text

# Text inputs

## Text input: Text

Text to change case

---
Title: util text encoding from
URL: https://toolbox.watermint.org/commands/util/text/encoding/from.md
---

# util text encoding from

Convert text encoding to UTF-8 text file from specified encoding. 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text encoding from -in /LOCAL/PATH/TO/INPUT_FILE -out /LOCAL/PATH/TO/OUTPUT_FILE -encoding ENCODING
```

## Options:

**-encoding**
: Encoding name

**-in**
: Input file path

**-out**
: Output file path

# Text inputs

## Text input: In

Text to change encoding

---
Title: util text encoding to
URL: https://toolbox.watermint.org/commands/util/text/encoding/to.md
---

# util text encoding to

Convert text encoding to specified encoding from UTF-8 text file. 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text encoding to -in /LOCAL/PATH/TO/INPUT_FILE -out /LOCAL/PATH/TO/OUTPUT_FILE -encoding ENCODING
```

## Options:

**-encoding**
: Encoding name

**-in**
: Input file path

**-out**
: Output file path

# Text inputs

## Text input: In

Text to change encoding

---
Title: util text nlp english entity
URL: https://toolbox.watermint.org/commands/util/text/nlp/english/entity.md
---

# util text nlp english entity

Split English text into entities 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text nlp english entity -in /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-ignore-line-break**
: Consider line break as regular white space while tokenizing. Default: false

**-in**
: Input file path

# Text inputs

## Text input: In

English text file to split

---
Title: util text nlp english sentence
URL: https://toolbox.watermint.org/commands/util/text/nlp/english/sentence.md
---

# util text nlp english sentence

Split English text into sentences 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text nlp english sentence -in /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-ignore-line-break**
: Consider line break as regular white space while tokenizing. Default: false

**-in**
: Input file path

# Text inputs

## Text input: In

English text file to split

---
Title: util text nlp english token
URL: https://toolbox.watermint.org/commands/util/text/nlp/english/token.md
---

# util text nlp english token

Split English text into tokens 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text nlp english token -in /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-ignore-line-break**
: Consider line break as regular white space while tokenizing. Default: false

**-in**
: Input file path

# Text inputs

## Text input: In

English text file to split

---
Title: util text nlp japanese token
URL: https://toolbox.watermint.org/commands/util/text/nlp/japanese/token.md
---

# util text nlp japanese token

Tokenize Japanese text 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text nlp japanese token -in /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-dictionary**
: Dictionary name of the token. Options: ipa (dictionary: ipa), uni (dictionary: uni). Default: ipa

**-ignore-line-break**
: Ignore line break. Default: false

**-in**
: Input file path

**-mode**
: Tokenize mode (normal/search/extended). Options:.   • normal (mode: normal).   • search (mode: search).   • extend (mode: extend). Default: normal

**-omit-bos-eos**
: Omit BOS/EOS tokens. Default: false

# Text inputs

## Text input: In

Input file path

---
Title: util text nlp japanese wakati
URL: https://toolbox.watermint.org/commands/util/text/nlp/japanese/wakati.md
---

# util text nlp japanese wakati

Wakachigaki (tokenize Japanese text) 

# Usage

This document uses the Desktop folder for command example.
```
tbx util text nlp japanese wakati -in /LOCAL/PATH/TO/INPUT.txt
```

## Options:

**-dictionary**
: Dictionary name (ipa/uni). Options: ipa (dictionary: ipa), uni (dictionary: uni). Default: ipa

**-ignore-line-break**
: Ignore line break. Default: false

**-in**
: Input file path

**-separator**
: Text separator. Default:  

# Text inputs

## Text input: In

Input text file path

---
Title: util tidy move dispatch
URL: https://toolbox.watermint.org/commands/util/tidy/move/dispatch.md
---

# util tidy move dispatch

Dispatch files (Irreversible operation)

# Usage

This document uses the Desktop folder for command example.
```
tbx util tidy move dispatch -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-preview**
: Preview mode. Default: false

# File formats

## Format: File

Data file for dispatch rules.

| Column              | Description                                   | Example                                   |
|---------------------|-----------------------------------------------|-------------------------------------------|
| suffix              | Source file name suffix                       | .pdf                                      |
| source_path         | Source path                                   | <no value>/Downloads                      |
| source_file_pattern | Source file name pattern (regular expression) | toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2})  |
| dest_path_pattern   | Destination path pattern                      | <no value>/Document/<no value>-<no value> |
| dest_file_pattern   | Destination file name pattern                 | TBX_<no value>-<no value>-<no value>      |

The first line is a header line. The program will accept a file without the header.
```
suffix,source_path,source_file_pattern,dest_path_pattern,dest_file_pattern
.pdf,<no value>/Downloads,toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2}),<no value>/Document/<no value>-<no value>,TBX_<no value>-<no value>-<no value>
```

---
Title: util tidy move simple
URL: https://toolbox.watermint.org/commands/util/tidy/move/simple.md
---

# util tidy move simple

Archive local files 

# Usage

This document uses the Desktop folder for command example.
```
tbx util tidy move simple -dst /LOCAL/DEST -src /LOCAL/SRC
```

## Options:

**-dst**
: The destination folder path. The command will create folders if they do not exist on the path.

**-exclude-folders**
: Exclude folders. Default: false

**-include-system-files**
: Include system files. Default: false

**-preview**
: Preview mode. Default: false

**-src**
: The source folder path.

---
Title: util tidy pack remote
URL: https://toolbox.watermint.org/commands/util/tidy/pack/remote.md
---

# util tidy pack remote

Package remote folder into the zip file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util tidy pack remote -dropbox-path /DROPBOX/PATH/TO/DOWNLOAD -local-path /LOCAL/PATH/TO/STORE.zip
```

## Options:

**-dropbox-path**
: Dropbox path to download

**-local-path**
: Local path to store zip file

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                            | Description                                                                                                          |
|-----------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                            | Status of the operation                                                                                              |
| reason                            | Reason of failure or skipped operation                                                                               |
| input.name                        | The last component of the path (including extension).                                                                |
| input.path_display                | The cased path to be used for display purposes only.                                                                 |
| input.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| input.size                        | The file size in bytes.                                                                                              |
| input.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |
| result.path                       | Path                                                                                                                 |
| result.name                       | File name                                                                                                            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: util time now
URL: https://toolbox.watermint.org/commands/util/time/now.md
---

# util time now

Display current time 

# Usage

This document uses the Desktop folder for command example.
```
tbx util time now 
```

## Options:

**-utc**
: Use UTC. Default: false

---
Title: util unixtime format
URL: https://toolbox.watermint.org/commands/util/unixtime/format.md
---

# util unixtime format

Time format to convert the unix time (epoch seconds from 1970-01-01) 

# Usage

This document uses the Desktop folder for command example.
```
tbx util unixtime format 
```

## Options:

**-format**
: Time format. Options:.   • iso8601 (Format: iso8601).   • rfc1123 (Format: rfc1123).   • rfc1123z (Format: rfc1123z).   • rfc3339 (Format: rfc3339).   • rfc3339_nano (Format: rfc3339_nano).   • rfc822 (Format: rfc822).   • rfc822z (Format: rfc822z). Default: iso8601

**-precision**
: Time precision (second/ms/ns). Options:.   • second (precision: second).   • ms (precision: ms).   • ns (precision: ns). Default: second

**-time**
: Unix Time. Default: 0

---
Title: util unixtime now
URL: https://toolbox.watermint.org/commands/util/unixtime/now.md
---

# util unixtime now

Display current time in unixtime 

# Usage

This document uses the Desktop folder for command example.
```
tbx util unixtime now 
```

## Options:

**-precision**
: Time precision (second/ms/ns). Options:.   • second (precision: second).   • ms (precision: ms).   • ns (precision: ns). Default: second

---
Title: util uuid timestamp
URL: https://toolbox.watermint.org/commands/util/uuid/timestamp.md
---

# util uuid timestamp

UUID Timestamp 

# Usage

This document uses the Desktop folder for command example.
```
tbx util uuid timestamp -uuid UUID
```

## Options:

**-uuid**
: UUID

---
Title: util uuid ulid
URL: https://toolbox.watermint.org/commands/util/uuid/ulid.md
---

# util uuid ulid

ULID Utility 

# Usage

This document uses the Desktop folder for command example.
```
tbx util uuid ulid 
```

---
Title: util uuid v4
URL: https://toolbox.watermint.org/commands/util/uuid/v4.md
---

# util uuid v4

Generate UUID v4 (random UUID) 

# Usage

This document uses the Desktop folder for command example.
```
tbx util uuid v4 
```

## Options:

**-upper-case**
: Output UUID in upper case. Default: false

---
Title: util uuid v7
URL: https://toolbox.watermint.org/commands/util/uuid/v7.md
---

# util uuid v7

Generate UUID v7 

# Usage

This document uses the Desktop folder for command example.
```
tbx util uuid v7 
```

## Options:

**-upper-case**
: Upper case. Default: false

---
Title: util uuid version
URL: https://toolbox.watermint.org/commands/util/uuid/version.md
---

# util uuid version

Parse version and variant of UUID 

# Usage

This document uses the Desktop folder for command example.
```
tbx util uuid version -uuid UUID
```

## Options:

**-uuid**
: UUID

# Results

## Report: metadata

UUID Metadata
The command will generate a report in three different formats. `metadata.csv`, `metadata.json`, and `metadata.xlsx`.

| Column  | Description  |
|---------|--------------|
| uuid    | UUID string  |
| version | UUID version |
| variant | UUID variant |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `metadata_0000.xlsx`, `metadata_0001.xlsx`, `metadata_0002.xlsx`, ...

---
Title: util xlsx create
URL: https://toolbox.watermint.org/commands/util/xlsx/create.md
---

# util xlsx create

Create an empty spreadsheet 

# Usage

This document uses the Desktop folder for command example.
```
tbx util xlsx create -file /LOCAL/PATH/TO/CREATE.xlsx -sheet SHEET_NAME
```

## Options:

**-file**
: Path to data file

**-sheet**
: Sheet name

---
Title: util xlsx sheet export
URL: https://toolbox.watermint.org/commands/util/xlsx/sheet/export.md
---

# util xlsx sheet export

Export data from the xlsx file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util xlsx sheet export -file /LOCAL/PATH/TO/EXPORT.xlsx -sheet SHEET_NAME
```

## Options:

**-data**
: Export data path

**-data-format**
: Output format

**-file**
: Path to data file

**-sheet**
: Sheet name

# Grid data output of the command

## Grid data output: Data

Export data

---
Title: util xlsx sheet import
URL: https://toolbox.watermint.org/commands/util/xlsx/sheet/import.md
---

# util xlsx sheet import

Import data into xlsx file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util xlsx sheet import -data /LOCAL/PATH/TO/INPUT.csv -file /LOCAL/PATH/TO/TARGET.xlsx -sheet SHEET_NAME
```

## Options:

**-create**
: Create a file if not found. Default: false

**-data**
: Data path

**-file**
: Path to data file

**-position**
: Start position to import in A1 notation. Default: `A1`.. Default: A1

**-sheet**
: Sheet name

# Grid data input for the command

## Grid data input: Data

Input data file

---
Title: util xlsx sheet list
URL: https://toolbox.watermint.org/commands/util/xlsx/sheet/list.md
---

# util xlsx sheet list

List sheets of the xlsx file 

# Usage

This document uses the Desktop folder for command example.
```
tbx util xlsx sheet list -file /LOCAL/PATH/TO/process.xlsx
```

## Options:

**-file**
: Path to data file

# Results

## Report: sheets

Sheet
The command will generate a report in three different formats. `sheets.csv`, `sheets.json`, and `sheets.xlsx`.

| Column | Description                           |
|--------|---------------------------------------|
| name   | Name of the sheet                     |
| rows   | Number of rows                        |
| cols   | Number of columns                     |
| hidden | True if the sheet is marked as hidden |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `sheets_0000.xlsx`, `sheets_0001.xlsx`, `sheets_0002.xlsx`, ...

---
Title: deepl translate text
URL: https://toolbox.watermint.org/commands/deepl/translate/text.md
---

# deepl translate text

Translate text 

# Usage

This document uses the Desktop folder for command example.
```
tbx deepl translate text -target-lang TARGET_LANG -text TEXT_TO_TRANSLATE
```

## Options:

**-peer**
: Account alias. Default: default

**-source-lang**
: Source language code (auto detect when omitted)

**-target-lang**
: Target language code

**-text**
: Text to translate

---
Title: dropbox file copy
URL: https://toolbox.watermint.org/commands/dropbox/file/copy.md
---

# dropbox file copy

Copy files 

Copies files or folders from one location to another within the same Dropbox account.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file copy -src /SRC/PATH -dst /DST/PATH
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dst**
: Destination path

**-peer**
: Account alias. Default: default

**-src**
: Source path

---
Title: dropbox file delete
URL: https://toolbox.watermint.org/commands/dropbox/file/delete.md
---

# dropbox file delete

Delete file or folder (Irreversible operation)

Permanently deletes files or folders from Dropbox (irreversible operation).

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file delete -path /PATH/TO/DELETE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path to delete

**-peer**
: Account alias. Default: default

---
Title: dropbox file info
URL: https://toolbox.watermint.org/commands/dropbox/file/info.md
---

# dropbox file info

Resolve metadata of the path 

Retrieves and displays metadata information for a specific file or folder path.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file info -path /DROPBOX/PATH/TO/FILE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path

**-peer**
: Account alias. Default: default

# Results

## Report: metadata

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `metadata.csv`, `metadata.json`, and `metadata.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |
| shared_folder_id            | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.                 |
| parent_shared_folder_id     | ID of shared folder that holds this file.                                                                            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `metadata_0000.xlsx`, `metadata_0001.xlsx`, `metadata_0002.xlsx`, ...

---
Title: dropbox file list
URL: https://toolbox.watermint.org/commands/dropbox/file/list.md
---

# dropbox file list

List files and folders 

Lists files and folders at a given path with options for recursive listing and filtering.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file list -path /path
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-include-deleted**
: Include deleted files. Default: false

**-include-explicit-shared-members**
: If true, the results will include a flag for each file indicating whether or not that file has any explicit members.. Default: false

**-include-mounted-folders**
: If true, the results will include entries under mounted folders which include app folder, shared folder and team folder.. Default: false

**-path**
: Path

**-peer**
: Account alias. Default: default

**-recursive**
: List recursively. Default: false

# Results

## Report: file_list

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `file_list.csv`, `file_list.json`, and `file_list.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| size                        | The file size in bytes.                                                                                              |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_list_0000.xlsx`, `file_list_0001.xlsx`, `file_list_0002.xlsx`, ...

---
Title: dropbox file merge
URL: https://toolbox.watermint.org/commands/dropbox/file/merge.md
---

# dropbox file merge

Merge paths (Irreversible operation)

Merges contents from one folder into another, with options for dry-run and empty folder handling.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file merge -from /from/path -to /path/to
```
Please add `-dry-run=false` option after verifying integrity of expected result.

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dry-run**
: Dry run. Default: true

**-from**
: Source path for merge

**-keep-empty-folder**
: Keep empty folders after merge. Default: false

**-peer**
: Account alias. Default: default

**-to**
: Destination path for merge

**-within-same-namespace**
: Do not cross namespace. This is to preserve sharing permissions including shared links. Default: false

---
Title: dropbox file move
URL: https://toolbox.watermint.org/commands/dropbox/file/move.md
---

# dropbox file move

Move files (Irreversible operation)

Moves files or folders from one location to another within Dropbox (irreversible operation).

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file move -src /SRC/PATH -dst /DST/PATH
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dst**
: Destination path

**-peer**
: Account alias. Default: default

**-src**
: Source path

---
Title: dropbox file replication
URL: https://toolbox.watermint.org/commands/dropbox/file/replication.md
---

# dropbox file replication

Replicate file content to the other account (Irreversible operation)

Replicates files and folders from one Dropbox account to another, mirroring the source structure.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file replication -src source -src-path /path/src -dst dest -dst-path /path/dest
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dst**
: Account alias (destination). Default: dst

**-dst-path**
: Destination path

**-src**
: Account alias (source). Default: src

**-src-path**
: Source path

# Results

## Report: replication_diff

This report shows a difference between two folders.
The command will generate a report in three different formats. `replication_diff.csv`, `replication_diff.json`, and `replication_diff.xlsx`.

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder or file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `replication_diff_0000.xlsx`, `replication_diff_0001.xlsx`, `replication_diff_0002.xlsx`, ...

---
Title: dropbox file size
URL: https://toolbox.watermint.org/commands/dropbox/file/size.md
---

# dropbox file size

Storage usage 

Calculates and reports the size of folders and their contents at specified depth levels.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file size -path /
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-depth**
: Report entries for files and folders up to the specified depth. Default: 2

**-path**
: Path to scan

**-peer**
: Account alias. Default: default

# Results

## Report: size

Folder size
The command will generate a report in three different formats. `size.csv`, `size.json`, and `size.xlsx`.

| Column                 | Description                                                               |
|------------------------|---------------------------------------------------------------------------|
| path                   | Path                                                                      |
| depth                  | Folder depth                                                              |
| size                   | Size in bytes                                                             |
| num_file               | Number of files in this folder and child folders                          |
| num_folder             | Number of folders in this folder and child folders                        |
| mod_time_earliest      | The earliest modification time of a file in this folder or child folders. |
| mod_time_latest        | The latest modification time of a file in this folder or child folders.   |
| operational_complexity | Operational complexity factor                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `size_0000.xlsx`, `size_0001.xlsx`, `size_0002.xlsx`, ...

---
Title: dropbox file watch
URL: https://toolbox.watermint.org/commands/dropbox/file/watch.md
---

# dropbox file watch

Watch file activities 

Monitors a path for changes and outputs file/folder modifications in real-time.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file watch -path /DROPBOX/PATH/TO/WATCH
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path to watch

**-peer**
: Account alias. Default: default

**-recursive**
: Watch path recursively. Default: false

---
Title: dropbox file account feature
URL: https://toolbox.watermint.org/commands/dropbox/file/account/feature.md
---

# dropbox file account feature

List Dropbox account features 

Retrieves and displays the enabled features and capabilities for the connected Dropbox account.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file account feature 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: report

Feature setting for the user
The command will generate a report in three different formats. `report.csv`, `report.json`, and `report.xlsx`.

| Column               | Description                                                                                                                                       |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| paper_as_files       | When this value is true, the user's Paper docs are accessible in Dropbox with the .paper extension and must be accessed via the /files endpoints. |
| file_locking         | When this value is True, the user can lock files in shared folders.                                                                               |
| team_shared_dropbox  | This feature contains information about whether or not the user is part of a team with a shared team root.                                        |
| distinct_member_home | This feature contains information about whether or not the user's home namespace is distinct from their root namespace.                           |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `report_0000.xlsx`, `report_0001.xlsx`, `report_0002.xlsx`, ...

---
Title: dropbox file account filesystem
URL: https://toolbox.watermint.org/commands/dropbox/file/account/filesystem.md
---

# dropbox file account filesystem

Show Dropbox file system version 

Shows the file system version/type being used by the account (individual or team file system).

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file account filesystem 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: file_system

File system version information
The command will generate a report in three different formats. `file_system.csv`, `file_system.json`, and `file_system.xlsx`.

| Column                                      | Description                                                     |
|---------------------------------------------|-----------------------------------------------------------------|
| version                                     | Version of the file system                                      |
| release_year                                | Year of the file system release                                 |
| has_distinct_member_homes                   | True if the team has distinct member home folder                |
| has_team_shared_dropbox                     | True if the team has team shared Dropbox                        |
| is_team_folder_api_supported                | True if team folder API is supported                            |
| is_path_root_required_to_access_team_folder | True if Dropbox-API-Path-Root is required to access team folder |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_system_0000.xlsx`, `file_system_0001.xlsx`, `file_system_0002.xlsx`, ...

---
Title: dropbox file account info
URL: https://toolbox.watermint.org/commands/dropbox/file/account/info.md
---

# dropbox file account info

Dropbox account info 

Displays profile information for the connected Dropbox account including name and email.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file account info 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: profile

This report shows a list of members.
The command will generate a report in three different formats. `profile.csv`, `profile.json`, and `profile.xlsx`.

| Column         | Description                                                                         |
|----------------|-------------------------------------------------------------------------------------|
| email          | Email address of user.                                                              |
| email_verified | Is true if the user's email is verified to be owned by the user.                    |
| given_name     | Also known as a first name                                                          |
| surname        | Also known as a last name or family name.                                           |
| display_name   | A name that can be used directly to represent the name of a user's Dropbox account. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `profile_0000.xlsx`, `profile_0001.xlsx`, `profile_0002.xlsx`, ...

---
Title: dropbox file compare account
URL: https://toolbox.watermint.org/commands/dropbox/file/compare/account.md
---

# dropbox file compare account

Compare files of two accounts 

Compares files and folders between two different Dropbox accounts to identify differences.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file compare account -left left -left-path /path/to/compare -right right -right-path /path/to/compare
```
If you want to compare different paths in same account, please specify same alias name to `-left` and `-right`.

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-left**
: Account alias (left). Default: left

**-left-path**
: The path from account root (left)

**-right**
: Account alias (right). Default: right

**-right-path**
: The path from account root (right)

# Results

## Report: diff

This report shows a difference between two folders.
The command will generate a report in three different formats. `diff.csv`, `diff.json`, and `diff.xlsx`.

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder or file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `diff_0000.xlsx`, `diff_0001.xlsx`, `diff_0002.xlsx`, ...

---
Title: dropbox file compare local
URL: https://toolbox.watermint.org/commands/dropbox/file/compare/local.md
---

# dropbox file compare local

Compare local folders and Dropbox folders 

Compares local files and folders with their Dropbox counterparts to identify differences.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file compare local -local-path /path/to/local -dropbox-path /path/on/dropbox
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dropbox-path**
: Dropbox path

**-local-path**
: Local path

**-peer**
: Account alias. Default: default

# Results

## Report: diff

This report shows a difference between two folders.
The command will generate a report in three different formats. `diff.csv`, `diff.json`, and `diff.xlsx`.

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder or file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `diff_0000.xlsx`, `diff_0001.xlsx`, `diff_0002.xlsx`, ...

## Report: skip

This report shows a difference between two folders.
The command will generate a report in three different formats. `skip.csv`, `skip.json`, and `skip.xlsx`.

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder or file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `skip_0000.xlsx`, `skip_0001.xlsx`, `skip_0002.xlsx`, ...

---
Title: dropbox file export doc
URL: https://toolbox.watermint.org/commands/dropbox/file/export/doc.md
---

# dropbox file export doc

Export document (Experimental)

Exports Dropbox Paper documents and Google Docs to local files in specified formats.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file export doc -dropbox-path /DROPBOX/PATH/TO/FILE -local-path /LOCAL/PATH/TO/EXPORT
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dropbox-path**
: Dropbox document path to export.

**-format**
: Export format

**-local-path**
: Local path to save

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the result of exporting a file.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column          | Description                                                                                            |
|-----------------|--------------------------------------------------------------------------------------------------------|
| name            | The last component of the path (including extension).                                                  |
| path_display    | The cased path to be used for display purposes only.                                                   |
| client_modified | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified | The last time the file was modified on Dropbox.                                                        |
| size            | The file size in bytes.                                                                                |
| export_name     | File name for export file.                                                                             |
| export_size     | File size of export file.                                                                              |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file export url
URL: https://toolbox.watermint.org/commands/dropbox/file/export/url.md
---

# dropbox file export url

Export a document from the URL 

Exports a file from Dropbox by downloading it from a shared link URL.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file export url -local-path /LOCAL/PATH/TO/EXPORT -url DOCUMENT_URL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-format**
: Export format

**-local-path**
: Local path to export

**-password**
: Password for the shared link

**-peer**
: Account alias. Default: default

**-url**
: URL of the document

# Results

## Report: operation_log

This report shows the result of exporting a file.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column          | Description                                                                                            |
|-----------------|--------------------------------------------------------------------------------------------------------|
| name            | The last component of the path (including extension).                                                  |
| path_display    | The cased path to be used for display purposes only.                                                   |
| client_modified | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified | The last time the file was modified on Dropbox.                                                        |
| size            | The file size in bytes.                                                                                |
| export_name     | File name for export file.                                                                             |
| export_size     | File size of export file.                                                                              |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file import url
URL: https://toolbox.watermint.org/commands/dropbox/file/import/url.md
---

# dropbox file import url

Import file from the URL (Irreversible operation)

Imports a single file into Dropbox by downloading from a specified URL.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file import url -url URL -path /path/to/import
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path to import

**-peer**
: Account alias. Default: default

**-url**
: URL

# Results

## Report: operation_log

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file import batch url
URL: https://toolbox.watermint.org/commands/dropbox/file/import/batch/url.md
---

# dropbox file import batch url

Batch import files from URL (Irreversible operation)

Imports multiple files into Dropbox by downloading from a list of URLs.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file import batch url -file /path/to/data/file -path /path/to/import
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Data file

**-path**
: Path to import

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for batch importing files from URLs.

| Column | Description                                                             | Example                           |
|--------|-------------------------------------------------------------------------|-----------------------------------|
| url    | URL to download                                                         | http://example.com/2019/12/26.zip |
| path   | Path to store file (use path given by `-path` when the record is empty) | /backup/2019-12-16.zip            |

The first line is a header line. The program will accept a file without the header.
```
url,path
http://example.com/2019/12/26.zip,/backup/2019-12-16.zip
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.url                          | URL to download                                                                                                      |
| input.path                         | Path to store file (use path given by `-path` when the record is empty)                                              |
| result.tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file lock acquire
URL: https://toolbox.watermint.org/commands/dropbox/file/lock/acquire.md
---

# dropbox file lock acquire

Lock a file 

Acquires an exclusive lock on a file to prevent others from editing it.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file lock acquire -path /DROPBOX/FILE/PATH/TO/LOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File path to lock

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file lock list
URL: https://toolbox.watermint.org/commands/dropbox/file/lock/list.md
---

# dropbox file lock list

List locks under the specified path 

Lists all files that are currently locked, showing lock holder information.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file lock list -path /DROPBOX/PATH/TO/SEARCH/LOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path

**-peer**
: Account alias. Default: default

# Results

## Report: lock

Lock information
The command will generate a report in three different formats. `lock.csv`, `lock.json`, and `lock.xlsx`.

| Column           | Description                                                                                            |
|------------------|--------------------------------------------------------------------------------------------------------|
| tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name             | The last component of the path (including extension).                                                  |
| path_display     | The cased path to be used for display purposes only.                                                   |
| client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified  | The last time the file was modified on Dropbox.                                                        |
| size             | The file size in bytes.                                                                                |
| is_lock_holder   | True if caller holds the file lock                                                                     |
| lock_holder_name | The display name of the lock holder.                                                                   |
| lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `lock_0000.xlsx`, `lock_0001.xlsx`, `lock_0002.xlsx`, ...

---
Title: dropbox file lock release
URL: https://toolbox.watermint.org/commands/dropbox/file/lock/release.md
---

# dropbox file lock release

Release a lock 

Releases the lock on a specific file, allowing others to edit it.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file lock release -path /DROPBOX/FILE/PATH/TO/UNLOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path to the file

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file lock all release
URL: https://toolbox.watermint.org/commands/dropbox/file/lock/all/release.md
---

# dropbox file lock all release

Release all locks under the specified path 

Releases all file locks held by the current user across the account.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file lock all release -path /DROPBOX/PATH/TO/RELEASE/LOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-batch-size**
: Operation batch size. Default: 100

**-path**
: Path to release locks

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file lock batch acquire
URL: https://toolbox.watermint.org/commands/dropbox/file/lock/batch/acquire.md
---

# dropbox file lock batch acquire

Lock multiple files 

Acquires locks on multiple files in a single batch operation.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file lock batch acquire -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-batch-size**
: Operation batch size. Default: 100

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Path

| Column | Description      | Example              |
|--------|------------------|----------------------|
| path   | Path to the file | /Report/2021-02.xlsx |

The first line is a header line. The program will accept a file without the header.
```
path
/Report/2021-02.xlsx
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path to the file                                                                                       |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file lock batch release
URL: https://toolbox.watermint.org/commands/dropbox/file/lock/batch/release.md
---

# dropbox file lock batch release

Release multiple locks 

Releases locks on multiple files in a single batch operation.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file lock batch release -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Path

| Column | Description      | Example              |
|--------|------------------|----------------------|
| path   | Path to the file | /Report/2021-02.xlsx |

The first line is a header line. The program will accept a file without the header.
```
path
/Report/2021-02.xlsx
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path to the file                                                                                       |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file request create
URL: https://toolbox.watermint.org/commands/dropbox/file/request/create.md
---

# dropbox file request create

Create a file request (Irreversible operation)

Creates a file request folder where others can upload files without Dropbox access.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file request create -path /DROPBOX/PATH/OF/FILE_REQUEST -title TITLE
```

## Options:

**-allow-late-uploads**
: If set, allow uploads after the deadline has passed (one_day/two_days/seven_days/thirty_days/always)

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-deadline**
: The deadline for this file request.

**-path**
: The path for the folder in Dropbox where uploaded files will be sent.

**-peer**
: Account alias. Default: default

**-title**
: The title of the file request

# Results

## Report: file_request

This report shows a list of file requests.
The command will generate a report in three different formats. `file_request.csv`, `file_request.json`, and `file_request.xlsx`.

| Column                      | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| id                          | The Id of the file request                                                |
| url                         | The URL of the file request                                               |
| title                       | The title of the file request                                             |
| created                     | Date/time when the file request was created.                              |
| is_open                     | Whether or not the file request is open.                                  |
| file_count                  | The number of files this file request has received.                       |
| destination                 | The path for the folder in the Dropbox where uploaded files will be sent. |
| deadline                    | The deadline for this file request.                                       |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed.                      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`, ...

---
Title: dropbox file request list
URL: https://toolbox.watermint.org/commands/dropbox/file/request/list.md
---

# dropbox file request list

List file requests of the individual account 

Lists all file requests in the account with their status and details.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file request list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: file_requests

This report shows a list of file requests.
The command will generate a report in three different formats. `file_requests.csv`, `file_requests.json`, and `file_requests.xlsx`.

| Column                      | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| id                          | The Id of the file request                                                |
| url                         | The URL of the file request                                               |
| title                       | The title of the file request                                             |
| created                     | Date/time when the file request was created.                              |
| is_open                     | Whether or not the file request is open.                                  |
| file_count                  | The number of files this file request has received.                       |
| destination                 | The path for the folder in the Dropbox where uploaded files will be sent. |
| deadline                    | The deadline for this file request.                                       |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed.                      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_requests_0000.xlsx`, `file_requests_0001.xlsx`, `file_requests_0002.xlsx`, ...

---
Title: dropbox file request delete closed
URL: https://toolbox.watermint.org/commands/dropbox/file/request/delete/closed.md
---

# dropbox file request delete closed

Delete all closed file requests on this account. (Irreversible operation)

Deletes file requests that have been closed and are no longer accepting uploads.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file request delete closed 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: deleted

This report shows a list of file requests.
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                      | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| id                          | The Id of the file request                                                |
| url                         | The URL of the file request                                               |
| title                       | The title of the file request                                             |
| created                     | Date/time when the file request was created.                              |
| is_open                     | Whether or not the file request is open.                                  |
| file_count                  | The number of files this file request has received.                       |
| destination                 | The path for the folder in the Dropbox where uploaded files will be sent. |
| deadline                    | The deadline for this file request.                                       |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed.                      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

---
Title: dropbox file request delete url
URL: https://toolbox.watermint.org/commands/dropbox/file/request/delete/url.md
---

# dropbox file request delete url

Delete a file request by the file request URL (Irreversible operation)

Deletes a specific file request using its URL.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file request delete url -url URL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-force**
: Force delete the file request.. Default: false

**-peer**
: Account alias. Default: default

**-url**
: URL of the file request.

# Results

## Report: deleted

This report shows a list of file requests.
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                      | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| id                          | The Id of the file request                                                |
| url                         | The URL of the file request                                               |
| title                       | The title of the file request                                             |
| created                     | Date/time when the file request was created.                              |
| is_open                     | Whether or not the file request is open.                                  |
| file_count                  | The number of files this file request has received.                       |
| destination                 | The path for the folder in the Dropbox where uploaded files will be sent. |
| deadline                    | The deadline for this file request.                                       |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed.                      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

---
Title: dropbox file restore all
URL: https://toolbox.watermint.org/commands/dropbox/file/restore/all.md
---

# dropbox file restore all

Restore files under given path (Experimental, and Irreversible operation)

Restores all deleted files and folders within a specified path.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file restore all -path /DROPBOX/PATH/TO/RESTORE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.path                         | Path                                                                                                                 |
| result.tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file restore ext
URL: https://toolbox.watermint.org/commands/dropbox/file/restore/ext.md
---

# dropbox file restore ext

Restore files with a specific extension (Experimental, and Irreversible operation)

Restores deleted files matching specific file extensions within a path.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file restore ext -ext EXT -path /DROPBOX/PATH/TO/RESTORE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-ext**
: Extension to restore (e.g. jpg, png, pdf)

**-path**
: Path to restore

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.path                         | Path                                                                                                                 |
| result.tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox file revision download
URL: https://toolbox.watermint.org/commands/dropbox/file/revision/download.md
---

# dropbox file revision download

Download the file revision 

Downloads a specific revision/version of a file from its revision history.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file revision download -local-path /LOCAL/PATH/TO/DOWNLOAD -revision REVISION
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-local-path**
: Local path to store downloaded file

**-peer**
: Account alias. Default: default

**-revision**
: File revision

# Results

## Report: entry

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `entry.csv`, `entry.json`, and `entry.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| id                          | A unique identifier for the file.                                                                                    |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `entry_0000.xlsx`, `entry_0001.xlsx`, `entry_0002.xlsx`, ...

---
Title: dropbox file revision list
URL: https://toolbox.watermint.org/commands/dropbox/file/revision/list.md
---

# dropbox file revision list

List file revisions 

Lists all available revisions for a file showing modification times and sizes.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file revision list -path /DROPBOX/PATH/TO/FILE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File path

**-peer**
: Account alias. Default: default

# Results

## Report: revisions

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `revisions.csv`, `revisions.json`, and `revisions.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| id                          | A unique identifier for the file.                                                                                    |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `revisions_0000.xlsx`, `revisions_0001.xlsx`, `revisions_0002.xlsx`, ...

---
Title: dropbox file revision restore
URL: https://toolbox.watermint.org/commands/dropbox/file/revision/restore.md
---

# dropbox file revision restore

Restore the file revision 

Restores a file to a previous revision from its version history.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file revision restore -path /DROPBOX/PATH/TO/RESTORE -revision REVISION
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File path

**-peer**
: Account alias. Default: default

**-revision**
: File revision

# Results

## Report: entry

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `entry.csv`, `entry.json`, and `entry.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| id                          | A unique identifier for the file.                                                                                    |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `entry_0000.xlsx`, `entry_0001.xlsx`, `entry_0002.xlsx`, ...

---
Title: dropbox file search content
URL: https://toolbox.watermint.org/commands/dropbox/file/search/content.md
---

# dropbox file search content

Search file content 

Searches for files by content with options for file type and category filtering.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file search content -query QUERY
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-category**
: Restricts search to only the file categories specified (image/document/pdf/spreadsheet/presentation/audio/video/folder/paper/others).. Options:.   •  (category: ).   • image (category: image).   • document (category: document).   • pdf (category: pdf).   • spreadsheet (category: spreadsheet).   • presentation (category: presentation).   • audio (category: audio).   • video (category: video).   • folder (category: folder).   • paper (category: paper).   • others (category: others)

**-extension**
: Restricts search to only the extensions specified.

**-max-results**
: Maximum number of entries to return. Default: 25

**-path**
: Scopes the search to a path in the user's Dropbox.

**-peer**
: Account alias. Default: default

**-query**
: The string to search for.

# Results

## Report: matches

This report shows a result of search with highlighted text.
The command will generate a report in three different formats. `matches.csv`, `matches.json`, and `matches.xlsx`.

| Column         | Description              |
|----------------|--------------------------|
| tag            | Type of entry            |
| path_display   | Display path             |
| highlight_html | Highlighted text in HTML |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `matches_0000.xlsx`, `matches_0001.xlsx`, `matches_0002.xlsx`, ...

---
Title: dropbox file search name
URL: https://toolbox.watermint.org/commands/dropbox/file/search/name.md
---

# dropbox file search name

Search file name 

Searches for files and folders by name pattern across the Dropbox account.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file search name -query QUERY
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-category**
: Restricts search to only the file categories specified (image/document/pdf/spreadsheet/presentation/audio/video/folder/paper/others).. Options:.   •  (category: ).   • image (category: image).   • document (category: document).   • pdf (category: pdf).   • spreadsheet (category: spreadsheet).   • presentation (category: presentation).   • audio (category: audio).   • video (category: video).   • folder (category: folder).   • paper (category: paper).   • others (category: others)

**-extension**
: Restricts search to only the extensions specified.

**-path**
: Scopes the search to a path in the user's Dropbox.

**-peer**
: Account alias. Default: default

**-query**
: The string to search for.

# Results

## Report: matches

This report shows a result of search with highlighted text.
The command will generate a report in three different formats. `matches.csv`, `matches.json`, and `matches.xlsx`.

| Column         | Description              |
|----------------|--------------------------|
| tag            | Type of entry            |
| path_display   | Display path             |
| highlight_html | Highlighted text in HTML |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `matches_0000.xlsx`, `matches_0001.xlsx`, `matches_0002.xlsx`, ...

---
Title: dropbox file share info
URL: https://toolbox.watermint.org/commands/dropbox/file/share/info.md
---

# dropbox file share info

Retrieve sharing information of the file 

Retrieves sharing information and permissions for a specific file or folder.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file share info -path /DROPBOX/PATH/TO/GET_INFO
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File path

**-peer**
: Account alias. Default: default

# Results

## Report: metadata

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `metadata.csv`, `metadata.json`, and `metadata.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| id                          | A unique identifier for the file.                                                                                    |
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_lower                  | The lowercased full path in the user's Dropbox. This always starts with a slash.                                     |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |
| shared_folder_id            | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.                 |
| parent_shared_folder_id     | ID of shared folder that holds this file.                                                                            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `metadata_0000.xlsx`, `metadata_0001.xlsx`, `metadata_0002.xlsx`, ...

---
Title: dropbox file sharedfolder info
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/info.md
---

# dropbox file sharedfolder info

Get shared folder info 

Displays detailed information about a specific shared folder including members and permissions.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder info -shared-folder-id NAMESPACE_ID
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

**-shared-folder-id**
: Shared folder ID

# Results

## Report: policies

This report shows a list of shared folders.
The command will generate a report in three different formats. `policies.csv`, `policies.json`, and `policies.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policies_0000.xlsx`, `policies_0001.xlsx`, `policies_0002.xlsx`, ...

---
Title: dropbox file sharedfolder leave
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/leave.md
---

# dropbox file sharedfolder leave

Leave the shared folder 

Removes yourself from a shared folder you've been added to.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder leave -shared-folder-id SHARED_FOLDER_ID
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-keep-copy**
: Keep a copy of the folder's contents upon relinquishing membership.. Default: false

**-peer**
: Account alias. Default: default

**-shared-folder-id**
: The ID for the shared folder.

---
Title: dropbox file sharedfolder list
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/list.md
---

# dropbox file sharedfolder list

List shared folders 

Lists all shared folders you have access to with their sharing details.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: shared_folder

This report shows a list of shared folders.
The command will generate a report in three different formats. `shared_folder.csv`, `shared_folder.json`, and `shared_folder.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_folder_0000.xlsx`, `shared_folder_0001.xlsx`, `shared_folder_0002.xlsx`, ...

---
Title: dropbox file sharedfolder share
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/share.md
---

# dropbox file sharedfolder share

Share a folder 

Creates a shared folder from an existing folder with configurable sharing policies and permissions.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder share -path /DROPBOX/PATH/TO/SHARE
```

## Options:

**-acl-update-policy**
: Who can change a shared folder's access control list (ACL).. Options: owner (aclupdatepolicy: owner), editor (aclupdatepolicy: editor). Default: owner

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-policy**
: Who can be a member of this shared folder.. Options: team (memberpolicy: team), anyone (memberpolicy: anyone). Default: anyone

**-path**
: Path to be shared

**-peer**
: Account alias. Default: default

**-shared-link-policy**
: Who can view shared links in this folder.. Options: anyone (sharedlinkpolicy: anyone), members (sharedlinkpolicy: members). Default: anyone

# Results

## Report: shared

This report shows a list of shared folders.
The command will generate a report in three different formats. `shared.csv`, `shared.json`, and `shared.xlsx`.

| Column                  | Description                                                                                                             |
|-------------------------|-------------------------------------------------------------------------------------------------------------------------|
| shared_folder_id        | The ID of the shared folder.                                                                                            |
| parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| name                    | The name of this shared folder.                                                                                         |
| access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| path_lower              | The lower-cased full path of this shared folder.                                                                        |
| is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| is_team_folder          | Whether this folder is a team folder.                                                                                   |
| policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| policy_shared_link      | Who links can be shared with.                                                                                           |
| policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| owner_team_id           | Team ID of the folder owner team                                                                                        |
| owner_team_name         | Team name of the team that owns the folder                                                                              |
| access_inheritance      | Access inheritance type                                                                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_0000.xlsx`, `shared_0001.xlsx`, `shared_0002.xlsx`, ...

---
Title: dropbox file sharedfolder unshare
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/unshare.md
---

# dropbox file sharedfolder unshare

Unshare a folder 

Stops sharing a folder and optionally leaves a copy for current members.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder unshare -path /DROPBOX/PATH/TO/UNSHARE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-leave-copy**
: If true, members of this shared folder will get a copy of this folder after it's unshared.. Default: false

**-path**
: Path to be unshared

**-peer**
: Account alias. Default: default

---
Title: dropbox file sharedfolder member add
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/member/add.md
---

# dropbox file sharedfolder member add

Add a member to the shared folder 

Adds new members to a shared folder with specified access permissions.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder member add -email EMAIL -path /DROPBOX/PATH/TO/ADD
```

## Options:

**-access-level**
: Access type (viewer/editor). Options:.   • editor (Can edit, comment, and share).   • viewer (Can view and comment).   • viewer_no_comment (Can only view). Default: editor

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-email**
: Email address of the folder member

**-message**
: Custom message for invitation

**-path**
: Path to the shared folder

**-peer**
: Account alias. Default: default

**-silent**
: Do not send invitation email. Default: false

---
Title: dropbox file sharedfolder member delete
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/member/delete.md
---

# dropbox file sharedfolder member delete

Remove a member from the shared folder 

Removes members from a shared folder, revoking their access.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder member delete -email EMAIL -path /DROPBOX/PATH/TO/DELETE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-email**
: Email address of the folder member

**-leave-copy**
: If true, members of this shared folder will get a copy of this folder after it's unshared.. Default: false

**-path**
: Path to the shared folder

**-peer**
: Account alias. Default: default

---
Title: dropbox file sharedfolder member list
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/member/list.md
---

# dropbox file sharedfolder member list

List shared folder members 

Lists all members of a shared folder with their access levels and email addresses.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder member list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: member

This report shows a list of members of shared folders.
The command will generate a report in three different formats. `member.csv`, `member.json`, and `member.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| name                  | The name of this shared folder.                                                                           |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| is_inherited          | True if the member has access from a parent folder                                                        |
| email                 | Email address of user.                                                                                    |
| display_name          | A name that can be used directly to represent the name of a user's Dropbox account.                       |
| group_name            | Name of a group                                                                                           |
| invitee_email         | Email address of invitee for this folder                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...

---
Title: dropbox file sharedfolder mount add
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/mount/add.md
---

# dropbox file sharedfolder mount add

Add the shared folder to the current user's Dropbox 

Mounts a shared folder to your Dropbox, making it appear in your file structure.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder mount add -shared-folder-id SHARED_FOLDER_ID
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

**-shared-folder-id**
: The ID for the shared folder.

# Results

## Report: mount

This report shows a list of shared folders.
The command will generate a report in three different formats. `mount.csv`, `mount.json`, and `mount.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mount_0000.xlsx`, `mount_0001.xlsx`, `mount_0002.xlsx`, ...

---
Title: dropbox file sharedfolder mount delete
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/mount/delete.md
---

# dropbox file sharedfolder mount delete

Unmount the shared folder 

Unmounts a shared folder from your Dropbox without leaving the folder.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder mount delete -shared-folder-id SHARED_FOLDER_ID
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

**-shared-folder-id**
: The ID for the shared folder.

# Results

## Report: mount

This report shows a list of shared folders.
The command will generate a report in three different formats. `mount.csv`, `mount.json`, and `mount.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mount_0000.xlsx`, `mount_0001.xlsx`, `mount_0002.xlsx`, ...

---
Title: dropbox file sharedfolder mount list
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/mount/list.md
---

# dropbox file sharedfolder mount list

List all shared folders the current user has mounted 

Lists all shared folders currently mounted in your Dropbox.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder mount list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: mounts

This report shows a list of shared folders.
The command will generate a report in three different formats. `mounts.csv`, `mounts.json`, and `mounts.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mounts_0000.xlsx`, `mounts_0001.xlsx`, `mounts_0002.xlsx`, ...

---
Title: dropbox file sharedfolder mount mountable
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedfolder/mount/mountable.md
---

# dropbox file sharedfolder mount mountable

List all shared folders the current user can mount 

Lists shared folders that can be mounted but aren't currently in your Dropbox.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedfolder mount mountable 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-include-mounted**
: Include mounted folders.. Default: false

**-peer**
: Account alias. Default: default

# Results

## Report: mountables

This report shows a list of shared folders.
The command will generate a report in three different formats. `mountables.csv`, `mountables.json`, and `mountables.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mountables_0000.xlsx`, `mountables_0001.xlsx`, `mountables_0002.xlsx`, ...

---
Title: dropbox file sharedlink create
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedlink/create.md
---

# dropbox file sharedlink create

Create shared link (Irreversible operation)

Creates a shared link for a file or folder with optional password protection and expiration date.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedlink create -path /path/to/share
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-expires**
: Expiration date/time of the shared link

**-password**
: Password

**-path**
: Path

**-peer**
: Account alias. Default: default

**-team-only**
: Link is accessible only by team members. Default: false

# Results

## Report: created

This report shows a list of shared links.
The command will generate a report in three different formats. `created.csv`, `created.json`, and `created.xlsx`.

| Column     | Description                                                                                                                                                                                                         |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id         | A unique identifier for the linked file or folder                                                                                                                                                                   |
| tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| url        | URL of the shared link.                                                                                                                                                                                             |
| name       | The linked file name (including extension).                                                                                                                                                                         |
| expires    | Expiration time, if set.                                                                                                                                                                                            |
| path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `created_0000.xlsx`, `created_0001.xlsx`, `created_0002.xlsx`, ...

---
Title: dropbox file sharedlink delete
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedlink/delete.md
---

# dropbox file sharedlink delete

Remove shared links (Irreversible operation)

This command will delete shared links based on the path in Dropbox

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedlink delete -path /path/to/delete
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File or folder path to remove shared link

**-peer**
: Account alias. Default: default

**-recursive**
: Remove shared links recursively. Default: false

# Results

## Report: shared_link

This report shows the transaction result.
The command will generate a report in three different formats. `shared_link.csv`, `shared_link.json`, and `shared_link.xlsx`.

| Column           | Description                                                                                                                                                                                                         |
|------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status           | Status of the operation                                                                                                                                                                                             |
| reason           | Reason of failure or skipped operation                                                                                                                                                                              |
| input.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| input.url        | URL of the shared link.                                                                                                                                                                                             |
| input.name       | The linked file name (including extension).                                                                                                                                                                         |
| input.expires    | Expiration time, if set.                                                                                                                                                                                            |
| input.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| input.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...

---
Title: dropbox file sharedlink info
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedlink/info.md
---

# dropbox file sharedlink info

Get information about the shared link 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedlink info -url SHARED_LINK_URL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-password**
: Password of the link if required.

**-peer**
: Account alias. Default: default

**-url**
: URL of the shared link

# Results

## Report: shared_link

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `shared_link.csv`, `shared_link.json`, and `shared_link.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_lower                  | The lowercased full path in the user's Dropbox. This always starts with a slash.                                     |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...

---
Title: dropbox file sharedlink list
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedlink/list.md
---

# dropbox file sharedlink list

List shared links 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedlink list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: shared_link

This report shows a list of shared links.
The command will generate a report in three different formats. `shared_link.csv`, `shared_link.json`, and `shared_link.xlsx`.

| Column     | Description                                                                                                                                                                                                         |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| url        | URL of the shared link.                                                                                                                                                                                             |
| name       | The linked file name (including extension).                                                                                                                                                                         |
| expires    | Expiration time, if set.                                                                                                                                                                                            |
| path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...

---
Title: dropbox file sharedlink file list
URL: https://toolbox.watermint.org/commands/dropbox/file/sharedlink/file/list.md
---

# dropbox file sharedlink file list

List files for the shared link 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sharedlink file list -url SHAREDLINK_URL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-password**
: Password for the shared link

**-peer**
: Account alias. Default: default

**-url**
: Shared link URL

# Results

## Report: file_list

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `file_list.csv`, `file_list.json`, and `file_list.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| size                        | The file size in bytes.                                                                                              |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_list_0000.xlsx`, `file_list_0001.xlsx`, `file_list_0002.xlsx`, ...

---
Title: dropbox file sync down
URL: https://toolbox.watermint.org/commands/dropbox/file/sync/down.md
---

# dropbox file sync down

Downstream sync with Dropbox 

Downloads files from Dropbox to local filesystem with filtering and overwrite options.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sync down -dropbox-path /DROPBOX/PATH/TO/DOWNLOAD -local-path /LOCAL/PATH/TO/SAVE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-delete**
: Delete local file if a file is removed on Dropbox. Default: false

**-dropbox-path**
: Dropbox path

**-local-path**
: Local path

**-name-disable-ignore**
: Filter by name. Filter system file or ignore files.

**-name-name**
: Filter by name. Filter by exact match to the name.

**-name-name-prefix**
: Filter by name. Filter by name match to the prefix.

**-name-name-suffix**
: Filter by name. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

**-skip-existing**
: Skip existing files. Do not overwrite. Default: false

# Results

## Report: deleted

Path
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                       | Description      |
|------------------------------|------------------|
| entry_path                   | Path             |
| entry_shard.file_system_type | File system type |
| entry_shard.shard_id         | Shard ID         |
| entry_shard.attributes       | Shard attributes |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

## Report: downloaded

This report shows the transaction result.
The command will generate a report in three different formats. `downloaded.csv`, `downloaded.json`, and `downloaded.xlsx`.

| Column                            | Description                                                                                                          |
|-----------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                            | Status of the operation                                                                                              |
| reason                            | Reason of failure or skipped operation                                                                               |
| input.name                        | The last component of the path (including extension).                                                                |
| input.path_display                | The cased path to be used for display purposes only.                                                                 |
| input.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| input.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| input.size                        | The file size in bytes.                                                                                              |
| input.content_hash                | A hash of the file content.                                                                                          |
| input.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |
| result.path                       | Path                                                                                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `downloaded_0000.xlsx`, `downloaded_0001.xlsx`, `downloaded_0002.xlsx`, ...

## Report: skipped

This report shows the transaction result.
The command will generate a report in three different formats. `skipped.csv`, `skipped.json`, and `skipped.xlsx`.

| Column                             | Description                            |
|------------------------------------|----------------------------------------|
| status                             | Status of the operation                |
| reason                             | Reason of failure or skipped operation |
| input.entry_path                   | Path                                   |
| input.entry_shard.file_system_type | File system type                       |
| input.entry_shard.shard_id         | Shard ID                               |
| input.entry_shard.attributes       | Shard attributes                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## Report: summary

This report shows a summary of the upload results.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column                | Description                                   |
|-----------------------|-----------------------------------------------|
| start                 | Time of start                                 |
| end                   | Time of finish                                |
| num_bytes             | Total upload size (Bytes)                     |
| num_files_error       | The number of files failed or got an error.   |
| num_files_transferred | The number of files uploaded/downloaded.      |
| num_files_skip        | The number of files skipped or to skip.       |
| num_folder_created    | Number of created folders.                    |
| num_delete            | Number of deleted entries.                    |
| num_api_call          | The number of estimated API calls for upload. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

---
Title: dropbox file sync online
URL: https://toolbox.watermint.org/commands/dropbox/file/sync/online.md
---

# dropbox file sync online

Sync online files (Irreversible operation)

Synchronizes files between two different locations within Dropbox online storage.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sync online -src /DROPBOX/PATH/SRC -dst /DROPBOX/PATH/DST
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-delete**
: Delete file if a file is removed in source path. Default: false

**-dst**
: Destination path

**-name-disable-ignore**
: Filter by name. Filter system file or ignore files.

**-name-name**
: Filter by name. Filter by exact match to the name.

**-name-name-prefix**
: Filter by name. Filter by name match to the prefix.

**-name-name-suffix**
: Filter by name. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

**-skip-existing**
: Skip existing files. Do not overwrite. Default: false

**-src**
: Source path

# Results

## Report: deleted

Path
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                       | Description      |
|------------------------------|------------------|
| entry_path                   | Path             |
| entry_shard.file_system_type | File system type |
| entry_shard.shard_id         | Shard ID         |
| entry_shard.attributes       | Shard attributes |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

## Report: skipped

This report shows the transaction result.
The command will generate a report in three different formats. `skipped.csv`, `skipped.json`, and `skipped.xlsx`.

| Column                             | Description                            |
|------------------------------------|----------------------------------------|
| status                             | Status of the operation                |
| reason                             | Reason of failure or skipped operation |
| input.entry_path                   | Path                                   |
| input.entry_shard.file_system_type | File system type                       |
| input.entry_shard.shard_id         | Shard ID                               |
| input.entry_shard.attributes       | Shard attributes                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## Report: summary

This report shows a summary of the upload results.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column                | Description                                   |
|-----------------------|-----------------------------------------------|
| start                 | Time of start                                 |
| end                   | Time of finish                                |
| num_bytes             | Total upload size (Bytes)                     |
| num_files_error       | The number of files failed or got an error.   |
| num_files_transferred | The number of files uploaded/downloaded.      |
| num_files_skip        | The number of files skipped or to skip.       |
| num_folder_created    | Number of created folders.                    |
| num_delete            | Number of deleted entries.                    |
| num_api_call          | The number of estimated API calls for upload. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

## Report: uploaded

This report shows the transaction result.
The command will generate a report in three different formats. `uploaded.csv`, `uploaded.json`, and `uploaded.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.path                         | Path                                                                                                                 |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.content_hash                | A hash of the file content.                                                                                          |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...

---
Title: dropbox file sync up
URL: https://toolbox.watermint.org/commands/dropbox/file/sync/up.md
---

# dropbox file sync up

Upstream sync with Dropbox (Irreversible operation)

Uploads files from local filesystem to Dropbox with filtering and overwrite options.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file sync up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-batch-size**
: Batch commit size. Default: 50

**-delete**
: Delete Dropbox file if a file is removed locally. Default: false

**-dropbox-path**
: Destination Dropbox path

**-local-path**
: Local file path

**-name-disable-ignore**
: Filter by name. Filter system file or ignore files.

**-name-name**
: Filter by name. Filter by exact match to the name.

**-name-name-prefix**
: Filter by name. Filter by name match to the prefix.

**-name-name-suffix**
: Filter by name. Filter by name match to the suffix.

**-overwrite**
: Overwrite existing file on the target path if that exists.. Default: false

**-peer**
: Account alias. Default: default

# Results

## Report: deleted

Path
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                       | Description      |
|------------------------------|------------------|
| entry_path                   | Path             |
| entry_shard.file_system_type | File system type |
| entry_shard.shard_id         | Shard ID         |
| entry_shard.attributes       | Shard attributes |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

## Report: skipped

This report shows the transaction result.
The command will generate a report in three different formats. `skipped.csv`, `skipped.json`, and `skipped.xlsx`.

| Column                             | Description                            |
|------------------------------------|----------------------------------------|
| status                             | Status of the operation                |
| reason                             | Reason of failure or skipped operation |
| input.entry_path                   | Path                                   |
| input.entry_shard.file_system_type | File system type                       |
| input.entry_shard.shard_id         | Shard ID                               |
| input.entry_shard.attributes       | Shard attributes                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## Report: summary

This report shows a summary of the upload results.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column                | Description                                   |
|-----------------------|-----------------------------------------------|
| start                 | Time of start                                 |
| end                   | Time of finish                                |
| num_bytes             | Total upload size (Bytes)                     |
| num_files_error       | The number of files failed or got an error.   |
| num_files_transferred | The number of files uploaded/downloaded.      |
| num_files_skip        | The number of files skipped or to skip.       |
| num_folder_created    | Number of created folders.                    |
| num_delete            | Number of deleted entries.                    |
| num_api_call          | The number of estimated API calls for upload. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

## Report: uploaded

This report shows the transaction result.
The command will generate a report in three different formats. `uploaded.csv`, `uploaded.json`, and `uploaded.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.path                         | Path                                                                                                                 |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.content_hash                | A hash of the file content.                                                                                          |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...

---
Title: dropbox file tag add
URL: https://toolbox.watermint.org/commands/dropbox/file/tag/add.md
---

# dropbox file tag add

Add tag to file or folder 

Adds a custom tag to a file or folder for organization and categorization.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file tag add -path /DROPBOX/PATH/TO/TARGET -tag TAG_NAME
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File or folder path to add a tag.

**-peer**
: Account alias. Default: default

**-tag**
: Tag to add to the file or folder.

---
Title: dropbox file tag delete
URL: https://toolbox.watermint.org/commands/dropbox/file/tag/delete.md
---

# dropbox file tag delete

Delete a tag from the file/folder 

Removes a specific tag from a file or folder.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file tag delete -path /DROPBOX/PATH/TO/PROCESS -tag TAG_NAME
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: File or folder path to remove a tag.

**-peer**
: Account alias. Default: default

**-tag**
: Tag name

---
Title: dropbox file tag list
URL: https://toolbox.watermint.org/commands/dropbox/file/tag/list.md
---

# dropbox file tag list

List tags of the path 

Lists all tags associated with a specific file or folder path.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file tag list -path /DROPBOX/PATH/TO/TARGET
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Target path

**-peer**
: Account alias. Default: default

# Results

## Report: tags

File tag
The command will generate a report in three different formats. `tags.csv`, `tags.json`, and `tags.xlsx`.

| Column | Description |
|--------|-------------|
| path   | File path   |
| tag    | File tag    |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `tags_0000.xlsx`, `tags_0001.xlsx`, `tags_0002.xlsx`, ...

---
Title: dropbox file template apply
URL: https://toolbox.watermint.org/commands/dropbox/file/template/apply.md
---

# dropbox file template apply

Apply file/folder structure template to the Dropbox path 

Applies a saved file/folder structure template to create directories and files in Dropbox.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file template apply -path /DROPBOX/PATH/TO/APPLY -template /LOCAL/PATH/TO/template.json
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-path**
: Path to apply template

**-peer**
: Account alias. Default: default

**-template**
: Path to template file

---
Title: dropbox file template capture
URL: https://toolbox.watermint.org/commands/dropbox/file/template/capture.md
---

# dropbox file template capture

Capture file/folder structure as template from Dropbox path 

Captures the file/folder structure from a Dropbox path and saves it as a reusable template.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox file template capture -out /LOCAL/PATH/template.json -path /DROPBOX/PATH/TO/CAPTURE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-out**
: Template output path

**-path**
: Capture target path

**-peer**
: Account alias. Default: default

---
Title: dropbox paper append
URL: https://toolbox.watermint.org/commands/dropbox/paper/append.md
---

# dropbox paper append

Append the content to the end of the existing Paper doc 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox paper append -content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/APPEND.paper
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-content**
: Paper content

**-format**
: Import format (html/markdown/plain_text). Options:.   • markdown (Markdown format).   • plain_text (Plain text format).   • html (HTML format). Default: markdown

**-path**
: Path in the user's Dropbox

**-peer**
: Account alias. Default: default

# Results

## Report: created

Create/updated paper data
The command will generate a report in three different formats. `created.csv`, `created.json`, and `created.xlsx`.

| Column         | Description    |
|----------------|----------------|
| paper_revision | Paper revision |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `created_0000.xlsx`, `created_0001.xlsx`, `created_0002.xlsx`, ...

# Text inputs

## Text input: Content

Paper content

---
Title: dropbox paper create
URL: https://toolbox.watermint.org/commands/dropbox/paper/create.md
---

# dropbox paper create

Create new Paper in the path 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox paper create -content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/CREATE.paper
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-content**
: Paper content

**-format**
: Import format (html/markdown/plain_text). Options:.   • markdown (Markdown format).   • plain_text (Plain text format).   • html (HTML format). Default: markdown

**-path**
: Path in the user's Dropbox

**-peer**
: Account alias. Default: default

# Results

## Report: created

Create/updated paper data
The command will generate a report in three different formats. `created.csv`, `created.json`, and `created.xlsx`.

| Column         | Description      |
|----------------|------------------|
| url            | URL of the Paper |
| result_path    | Result path      |
| paper_revision | Paper revision   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `created_0000.xlsx`, `created_0001.xlsx`, `created_0002.xlsx`, ...

# Text inputs

## Text input: Content

Paper content

---
Title: dropbox paper overwrite
URL: https://toolbox.watermint.org/commands/dropbox/paper/overwrite.md
---

# dropbox paper overwrite

Overwrite an existing Paper document 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox paper overwrite -content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/OVERWRITE.paper
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-content**
: Paper content

**-format**
: Import format (html/markdown/plain_text). Options:.   • markdown (Markdown format).   • plain_text (Plain text format).   • html (HTML format). Default: markdown

**-path**
: Path in the user's Dropbox

**-peer**
: Account alias. Default: default

# Results

## Report: created

Create/updated paper data
The command will generate a report in three different formats. `created.csv`, `created.json`, and `created.xlsx`.

| Column         | Description    |
|----------------|----------------|
| paper_revision | Paper revision |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `created_0000.xlsx`, `created_0001.xlsx`, `created_0002.xlsx`, ...

# Text inputs

## Text input: Content

Paper content

---
Title: dropbox paper prepend
URL: https://toolbox.watermint.org/commands/dropbox/paper/prepend.md
---

# dropbox paper prepend

Append the content to the beginning of the existing Paper doc 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox paper prepend -content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/PREPEND.paper
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-content**
: Paper content

**-format**
: Import format (html/markdown/plain_text). Options:.   • markdown (Markdown format).   • plain_text (Plain text format).   • html (HTML format). Default: markdown

**-path**
: Path in the user's Dropbox

**-peer**
: Account alias. Default: default

# Results

## Report: created

Create/updated paper data
The command will generate a report in three different formats. `created.csv`, `created.json`, and `created.xlsx`.

| Column         | Description    |
|----------------|----------------|
| paper_revision | Paper revision |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `created_0000.xlsx`, `created_0001.xlsx`, `created_0002.xlsx`, ...

# Text inputs

## Text input: Content

Paper content

---
Title: dropbox sign request list
URL: https://toolbox.watermint.org/commands/dropbox/sign/request/list.md
---

# dropbox sign request list

List signature requests 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox sign request list 
```

## Options:

**-account-id**
: Which account to return SignatureRequests for. Must be a team member. Use `all` to indicate all team members. Defaults to your account.

**-peer**
: Account alias. Default: default

# Results

## Report: requests

Signature request
The command will generate a report in three different formats. `requests.csv`, `requests.json`, and `requests.xlsx`.

| Column                  | Description                                                                 |
|-------------------------|-----------------------------------------------------------------------------|
| signature_request_id    | The id of the SignatureRequest.                                             |
| requester_email_address | The email address of the initiator of the SignatureRequest.                 |
| title                   | The title the specified Account uses for the SignatureRequest.              |
| subject                 | The subject in the email that was initially sent to the signers.            |
| message                 | The custom message in the email that was initially sent to the signers.     |
| created_at_rfc3339      | Time the signature request was created.                                     |
| expires_at_rfc3339      | The time when the signature request will expire unsigned signatures.        |
| is_complete             | Whether or not the SignatureRequest has been fully executed by all signers. |
| is_declined             | Whether or not the SignatureRequest has been declined by a signer.          |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `requests_0000.xlsx`, `requests_0001.xlsx`, `requests_0002.xlsx`, ...

---
Title: dropbox sign request signature list
URL: https://toolbox.watermint.org/commands/dropbox/sign/request/signature/list.md
---

# dropbox sign request signature list

List signatures of requests 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox sign request signature list 
```

## Options:

**-account-id**
: Which account to return SignatureRequests for. Must be a team member. Use `all` to indicate all team members. Defaults to your account.

**-peer**
: Account alias. Default: default

# Results

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `signatures_0000.xlsx`, `signatures_0001.xlsx`, `signatures_0002.xlsx`, ...

---
Title: dropbox team feature
URL: https://toolbox.watermint.org/commands/dropbox/team/feature.md
---

# dropbox team feature

Display all features and capabilities enabled for your Dropbox team account, including API limits and special features 

Shows team's enabled features, beta access, and API rate limits. Check before using advanced features or planning integrations. Features may vary by subscription level. Useful for troubleshooting feature availability issues.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team feature 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: feature

Team feature
The command will generate a report in three different formats. `feature.csv`, `feature.json`, and `feature.xlsx`.

| Column                      | Description                                       |
|-----------------------------|---------------------------------------------------|
| upload_api_rate_limit       | The number of upload API calls allowed per month. |
| upload_api_rate_limit_count | The number of upload API calls made this month.   |
| has_team_shared_dropbox     | Does this team have a shared team root.           |
| has_team_file_events        | Team supports file events                         |
| has_team_selective_sync     | Team supports selective sync                      |
| has_distinct_member_homes   | Team has distinct member home folders             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `feature_0000.xlsx`, `feature_0001.xlsx`, `feature_0002.xlsx`, ...

---
Title: dropbox team filesystem
URL: https://toolbox.watermint.org/commands/dropbox/team/filesystem.md
---

# dropbox team filesystem

Identify whether your team uses legacy or modern file system architecture, important for feature compatibility 

Determines underlying file system version affecting feature availability and API behavior. Modern file system enables advanced features like native Paper and enhanced performance. Legacy teams may need migration for full feature access.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team filesystem 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: file_system

File system version information
The command will generate a report in three different formats. `file_system.csv`, `file_system.json`, and `file_system.xlsx`.

| Column                                      | Description                                                     |
|---------------------------------------------|-----------------------------------------------------------------|
| version                                     | Version of the file system                                      |
| release_year                                | Year of the file system release                                 |
| has_distinct_member_homes                   | True if the team has distinct member home folder                |
| has_team_shared_dropbox                     | True if the team has team shared Dropbox                        |
| is_team_folder_api_supported                | True if team folder API is supported                            |
| is_path_root_required_to_access_team_folder | True if Dropbox-API-Path-Root is required to access team folder |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_system_0000.xlsx`, `file_system_0001.xlsx`, `file_system_0002.xlsx`, ...

---
Title: dropbox team info
URL: https://toolbox.watermint.org/commands/dropbox/team/info.md
---

# dropbox team info

Display essential team account information including team ID and basic team settings 

Shows fundamental team account details needed for API integrations and support requests. Team ID is required for various administrative operations. Quick way to verify you're connected to the correct team account.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team info 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: info

This report shows a list of team information.
The command will generate a report in three different formats. `info.csv`, `info.json`, and `info.xlsx`.

| Column                      | Description                                                                                                   |
|-----------------------------|---------------------------------------------------------------------------------------------------------------|
| name                        | The name of the team                                                                                          |
| team_id                     | The ID of the team.                                                                                           |
| num_licensed_users          | The number of licenses available to the team.                                                                 |
| num_provisioned_users       | The number of accounts that have been invited or are already active members of the team.                      |
| policy_shared_folder_member | Which shared folders team members can join (from_team_only, or from_anyone)                                   |
| policy_shared_folder_join   | Who can join folders shared by team members (team, or anyone)                                                 |
| policy_shared_link_create   | Who can view shared links owned by team members (default_public, default_team_only, or team_only)             |
| policy_emm_state            | This describes the Enterprise Mobility Management (EMM) state for this team (disabled, optional, or required) |
| policy_office_add_in        | The admin policy around the Dropbox Office Add-In for this team (disabled, or enabled)                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `info_0000.xlsx`, `info_0001.xlsx`, `info_0002.xlsx`, ...

---
Title: dropbox team activity event
URL: https://toolbox.watermint.org/commands/dropbox/team/activity/event.md
---

# dropbox team activity event

Retrieve detailed team activity event logs with filtering options, essential for security auditing and compliance monitoring 

From release 91, the command parses `-start-time` or `-end-time` as the relative duration from now with the format like "-24h" (24 hours) or "-10m" (10 minutes).
If you wanted to retrieve events every hour, then run like:

```
tbx team activity event -start-time -1h -output json > latest_events.json
```

Then, append the latest part to the entire log if you want.

```
cat latest_events.json >> all.json
```

Or more precisely, retrieve events every hour with some overlap.
```
tbx team activity event -start-time -1h5m -output json > latest_events.json
```

Then, concatenate, and de-duplicate overlapped events:
```
cat all.json latest_events.json | sort -u > _all.json && mv _all.json all.json
```

If you prefer CSV format, then use the `jq` command to convert it.
```
cat latest_events.json | jq -r '[.timestamp, .actor[.actor.".tag"].display_name, .actor[.actor.".tag"].email, .event_type.description, .event_category.".tag", .origin.access_method.end_user.".tag", .origin.geo_location.ip_address, .origin.geo_location.country, .origin.geo_location.city, .involve_non_team_member, (.participants | @text), (.context | @text)] | @csv' >> all.csv
```

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team activity event 
```

## Options:

**-category**
: Filter the returned events to a single category. This field is optional.

**-end-time**
: Ending time (exclusive).

**-peer**
: Account alias. Default: default

**-start-time**
: Starting time (inclusive)

# Results

## Report: event

This report shows activity logs mostly compatible with Dropbox for teams' activity logs.
The command will generate a report in three different formats. `event.csv`, `event.json`, and `event.xlsx`.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `event_0000.xlsx`, `event_0001.xlsx`, `event_0002.xlsx`, ...

---
Title: dropbox team activity user
URL: https://toolbox.watermint.org/commands/dropbox/team/activity/user.md
---

# dropbox team activity user

Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities 

Retrieves detailed activity logs for individual team members, including file operations, sharing activities, and login events. Essential for user-specific audits, investigating security incidents, or understanding individual usage patterns. Can filter by activity category for focused analysis.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team activity user 
```

## Options:

**-category**
: Filter the returned events to a single category. This field is optional.

**-end-time**
: Ending time (exclusive).

**-peer**
: Account alias. Default: default

**-start-time**
: Starting time (inclusive)

# Results

## Report: user

This report shows activity logs mostly compatible with Dropbox for teams' activity logs.
The command will generate a report in three different formats. `user.csv`, `user.json`, and `user.xlsx`.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `user_0000.xlsx`, `user_0001.xlsx`, `user_0002.xlsx`, ...

## Report: user_summary

This report shows the transaction result.
The command will generate a report in three different formats. `user_summary.csv`, `user_summary.json`, and `user_summary.xlsx`.

| Column                 | Description                            |
|------------------------|----------------------------------------|
| status                 | Status of the operation                |
| reason                 | Reason of failure or skipped operation |
| input.user             | User email address                     |
| result.logins          | Number of login activities             |
| result.devices         | Number of device activities            |
| result.sharing         | Number of sharing activities           |
| result.file_operations | Number of file operation activities    |
| result.paper           | Number of activities of Paper          |
| result.others          | Number of other category activities    |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `user_summary_0000.xlsx`, `user_summary_0001.xlsx`, `user_summary_0002.xlsx`, ...

---
Title: dropbox team activity batch user
URL: https://toolbox.watermint.org/commands/dropbox/team/activity/batch/user.md
---

# dropbox team activity batch user

Scan and retrieve activity logs for multiple team members in batch, useful for compliance auditing and user behavior analysis 

This command processes a list of user email addresses from a file and retrieves their activity logs within a specified time range. Useful for HR investigations, compliance reporting, or analyzing patterns across specific user groups.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team activity batch user -file /path/to/file.csv
```

## Options:

**-category**
: Filter the returned events to a single category. This field is optional.

**-end-time**
: Ending time (exclusive).

**-file**
: User email address list file

**-peer**
: Account alias. Default: default

**-start-time**
: Starting time (inclusive)

# File formats

## Format: File

Data file for batch retrieving activities of members.

| Column | Description        | Example          |
|--------|--------------------|------------------|
| email  | User email address | john@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
john@example.com
```

# Results

## Report: combined

This report shows activity logs mostly compatible with Dropbox for teams' activity logs.
The command will generate a report in three different formats. `combined.csv`, `combined.json`, and `combined.xlsx`.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `combined_0000.xlsx`, `combined_0001.xlsx`, `combined_0002.xlsx`, ...

## Report: user

This report shows activity logs mostly compatible with Dropbox for teams' activity logs.
The command will generate a report in three different formats. `user.csv`, `user.json`, and `user.xlsx`.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `user_0000.xlsx`, `user_0001.xlsx`, `user_0002.xlsx`, ...

---
Title: dropbox team activity daily event
URL: https://toolbox.watermint.org/commands/dropbox/team/activity/daily/event.md
---

# dropbox team activity daily event

Generate daily activity reports showing team events grouped by date, helpful for tracking team usage patterns and security monitoring 

Aggregates team activity events by day, making it easier to identify trends and anomalies in team behavior. Particularly useful for creating daily security reports, tracking adoption of new features, or identifying unusual activity patterns that might indicate security concerns.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team activity daily event -start-date DATE
```

## Options:

**-category**
: Event category

**-end-date**
: End date

**-peer**
: Account alias. Default: default

**-start-date**
: Start date

# Results

## Report: event

This report shows activity logs mostly compatible with Dropbox for teams' activity logs.
The command will generate a report in three different formats. `event.csv`, `event.json`, and `event.xlsx`.

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `event_0000.xlsx`, `event_0001.xlsx`, `event_0002.xlsx`, ...

---
Title: dropbox team admin list
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/list.md
---

# dropbox team admin list

Display all team members with their assigned admin roles, helpful for auditing administrative access and permissions 

Generates a comprehensive admin audit report showing all members with elevated privileges. Can include non-admin members for complete visibility. Essential for security reviews, compliance audits, and ensuring appropriate access levels across the organization.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin list 
```

## Options:

**-include-non-admin**
: Include non admin members in the report. Default: false

**-member-roles**
: Member to admin-role mappings

**-member-roles-format**
: Output format

**-peer**
: Account alias. Default: default

# Grid data output of the command

## Grid data output: MemberRoles

Member to admin-role mappings

---
Title: dropbox team admin group role add
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/group/role/add.md
---

# dropbox team admin group role add

Assign admin roles to all members of a specified group, streamlining role management for large teams 

Efficiently grants admin privileges to entire groups rather than individual members. Ideal for departmental admin assignments or when onboarding new admin teams. Changes are applied immediately to all current group members.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin group role add -group GROUP_NAME -role-id ROLE_ID
```

## Options:

**-group**
: Group name

**-peer**
: Account alias. Default: default

**-role-id**
: Role ID

# Results

## Report: roles

The user's roles in the team.
The command will generate a report in three different formats. `roles.csv`, `roles.json`, and `roles.xlsx`.

| Column         | Description                         |
|----------------|-------------------------------------|
| team_member_id | Team member ID                      |
| email          | Email address of the member         |
| role_id        | A string containing encoded role ID |
| name           | The role display name.              |
| description    | Role description.                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `roles_0000.xlsx`, `roles_0001.xlsx`, `roles_0002.xlsx`, ...

---
Title: dropbox team admin group role delete
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/group/role/delete.md
---

# dropbox team admin group role delete

Remove admin roles from all team members except those in a specified exception group, useful for role cleanup and access control 

Bulk removes specific admin roles while preserving them for an exception group. Useful for reorganizing admin structures or implementing least-privilege access. The exception group ensures critical admins retain necessary permissions during cleanup operations.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin group role delete -exception-group GROUP_NAME -role-id ROLE_ID
```

## Options:

**-exception-group**
: Exception group name

**-peer**
: Account alias. Default: default

**-role-id**
: Role ID

# Results

## Report: roles

The user's roles in the team.
The command will generate a report in three different formats. `roles.csv`, `roles.json`, and `roles.xlsx`.

| Column         | Description                         |
|----------------|-------------------------------------|
| team_member_id | Team member ID                      |
| email          | Email address of the member         |
| role_id        | A string containing encoded role ID |
| name           | The role display name.              |
| description    | Role description.                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `roles_0000.xlsx`, `roles_0001.xlsx`, `roles_0002.xlsx`, ...

---
Title: dropbox team admin role add
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/role/add.md
---

# dropbox team admin role add

Grant a specific admin role to an individual team member, enabling granular permission management 

Assigns specific admin roles to individual members for precise permission control. Use when promoting team members to admin positions or adjusting responsibilities. The command validates that the member doesn't already have the specified role to prevent duplicates.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin role add -email EMAIL -role-id ROLE_ID
```

## Options:

**-email**
: Email address of the member

**-peer**
: Account alias. Default: default

**-role-id**
: Role ID

# Results

## Report: roles

The user's roles in the team.
The command will generate a report in three different formats. `roles.csv`, `roles.json`, and `roles.xlsx`.

| Column         | Description                         |
|----------------|-------------------------------------|
| team_member_id | Team member ID                      |
| email          | Email address of the member         |
| role_id        | A string containing encoded role ID |
| name           | The role display name.              |
| description    | Role description.                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `roles_0000.xlsx`, `roles_0001.xlsx`, `roles_0002.xlsx`, ...

---
Title: dropbox team admin role clear
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/role/clear.md
---

# dropbox team admin role clear

Revoke all administrative privileges from a team member, useful for role transitions or security purposes 

Completely removes all admin roles from a member in a single operation. Essential for offboarding admins, responding to security incidents, or transitioning members to non-administrative positions. More efficient than removing roles individually.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin role clear -email EMAIL
```

## Options:

**-email**
: Email address of the member

**-peer**
: Account alias. Default: default

---
Title: dropbox team admin role delete
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/role/delete.md
---

# dropbox team admin role delete

Remove a specific admin role from a team member while preserving other roles, allowing precise permission adjustments 

Selectively removes individual admin roles without affecting other permissions. Useful for adjusting responsibilities or implementing role-based access changes. The command verifies the member has the role before attempting removal.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin role delete -email EMAIL -role-id ROLE_ID
```

## Options:

**-email**
: Email address of the member

**-peer**
: Account alias. Default: default

**-role-id**
: Role ID

# Results

## Report: roles

The user's roles in the team.
The command will generate a report in three different formats. `roles.csv`, `roles.json`, and `roles.xlsx`.

| Column         | Description                         |
|----------------|-------------------------------------|
| team_member_id | Team member ID                      |
| email          | Email address of the member         |
| role_id        | A string containing encoded role ID |
| name           | The role display name.              |
| description    | Role description.                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `roles_0000.xlsx`, `roles_0001.xlsx`, `roles_0002.xlsx`, ...

---
Title: dropbox team admin role list
URL: https://toolbox.watermint.org/commands/dropbox/team/admin/role/list.md
---

# dropbox team admin role list

Display all available admin roles in the team with their descriptions and permissions 

Lists all possible admin roles available in your Dropbox team along with their capabilities. Reference this before assigning roles to understand permission implications. Helps ensure team members receive appropriate access levels.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team admin role list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: roles

The user's roles in the team.
The command will generate a report in three different formats. `roles.csv`, `roles.json`, and `roles.xlsx`.

| Column      | Description                         |
|-------------|-------------------------------------|
| role_id     | A string containing encoded role ID |
| name        | The role display name.              |
| description | Role description.                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `roles_0000.xlsx`, `roles_0001.xlsx`, `roles_0002.xlsx`, ...

---
Title: dropbox team backup device status
URL: https://toolbox.watermint.org/commands/dropbox/team/backup/device/status.md
---

# dropbox team backup device status

Track Dropbox Backup status changes for all team devices over a specified period, monitoring backup health and compliance 

Evaluates and reports the latest status of Dropbox Backup per device session from activity logs for a specified time period. If there is no activity during the specified period, it is reported as the value `no_status_update`.
Multiple device sessions may be displayed in the following cases
* If the Dropbox application has been reinstalled.
* If the Dropbox application has not been unlinked (e.g. you initialized the OS without unlinking the Dropbox application).

In that case, please refer to the report `session_info_updated` to see the most recent report. This command does not automatically make this determination, since it is possible that there may be a session with the same hostname by coincidence.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team backup device status -start-time "2024-01-01"
```

## Options:

**-end-time**
: End date/time of the period to retrieve data for (exclusive). If this is not specified, the current time is used.

**-peer**
: Account alias. Default: default

**-start-time**
: Start date/time of the period to retrieve data for (inclusive).

# Results

## Report: devices

Backup feature status of a device
The command will generate a report in three different formats. `devices.csv`, `devices.json`, and `devices.xlsx`.

| Column                      | Description                          |
|-----------------------------|--------------------------------------|
| actor_user_email            | User email                           |
| actor_user_display_name     | User display name                    |
| session_info_ip_address     | IP address                           |
| session_info_host_name      | Host name                            |
| session_info_updated        | Last Date/time of the session update |
| session_info_client_type    | Client type                          |
| session_info_client_version | Client version                       |
| session_info_platform       | Platform                             |
| timestamp                   | Timestamp of the event               |
| latest_status               | Latest status of the device          |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `devices_0000.xlsx`, `devices_0001.xlsx`, `devices_0002.xlsx`, ...

---
Title: dropbox team content legacypaper count
URL: https://toolbox.watermint.org/commands/dropbox/team/content/legacypaper/count.md
---

# dropbox team content legacypaper count

Calculate the total number of legacy Paper documents owned by each team member, useful for content auditing and migration planning 

Provides Paper document counts per member, distinguishing between created and accessed documents. Essential for planning Paper-to-Dropbox migrations, identifying heavy Paper users, and estimating migration scope. Filter options help focus on relevant document sets.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content legacypaper count 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: stats

Paper count
The command will generate a report in three different formats. `stats.csv`, `stats.json`, and `stats.xlsx`.

| Column       | Description                   |
|--------------|-------------------------------|
| member_email | Member email address          |
| created      | Number of created Paper docs  |
| accessed     | Number of accessed Paper docs |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `stats_0000.xlsx`, `stats_0001.xlsx`, `stats_0002.xlsx`, ...

---
Title: dropbox team content legacypaper export
URL: https://toolbox.watermint.org/commands/dropbox/team/content/legacypaper/export.md
---

# dropbox team content legacypaper export

Export all legacy Paper documents from team members to local storage in HTML or Markdown format for backup or migration 

Bulk exports team Paper documents to local storage, preserving content before migrations or for compliance archives. Supports HTML and Markdown formats. Creates organized folder structure by member. Consider available disk space as this may export large amounts of data.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content legacypaper export -path /LOCAL/PATH/TO/EXPORT
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-filter-by**
: Specify how the Paper docs should be filtered (doc_created/doc_accessed).. Options: docs_created (filterby: docs_created), docs_accessed (filterby: docs_accessed). Default: docs_created

**-format**
: Export file format (html/markdown). Options: html (HTML format), markdown (Markdown format). Default: html

**-path**
: Export folder path

**-peer**
: Account alias. Default: default

# Results

## Report: paper

Export data
The command will generate a report in three different formats. `paper.csv`, `paper.json`, and `paper.xlsx`.

| Column         | Description               |
|----------------|---------------------------|
| member_email   | Member email address      |
| paper_doc_id   | Paper Document ID         |
| paper_owner    | Paper owner email address |
| paper_title    | Paper title               |
| paper_revision | Paper revision            |
| export_path    | Export file path          |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `paper_0000.xlsx`, `paper_0001.xlsx`, `paper_0002.xlsx`, ...

---
Title: dropbox team content legacypaper list
URL: https://toolbox.watermint.org/commands/dropbox/team/content/legacypaper/list.md
---

# dropbox team content legacypaper list

Generate a comprehensive list of all legacy Paper documents across the team with ownership and metadata information 

Creates detailed inventory of all Paper documents including titles, owners, and last modified dates. Use for content audits, identifying orphaned documents, or preparing for migrations. Filter by creation or access patterns to focus analysis.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content legacypaper list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-filter-by**
: Specify how the Paper docs should be filtered (doc_created/doc_accessed).. Options: docs_created (filterby: docs_created), docs_accessed (filterby: docs_accessed). Default: docs_created

**-peer**
: Account alias. Default: default

# Results

## Report: paper

Member Paper metadata
The command will generate a report in three different formats. `paper.csv`, `paper.json`, and `paper.xlsx`.

| Column         | Description               |
|----------------|---------------------------|
| member_email   | Member email address      |
| paper_doc_id   | Paper Document ID         |
| paper_owner    | Paper owner email address |
| paper_title    | Paper title               |
| paper_revision | Paper revision            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `paper_0000.xlsx`, `paper_0001.xlsx`, `paper_0002.xlsx`, ...

---
Title: dropbox team content member list
URL: https://toolbox.watermint.org/commands/dropbox/team/content/member/list.md
---

# dropbox team content member list

Display all members with access to team folders and shared folders, showing permission levels and folder relationships 

Maps folder access across the team, showing which members can access specific folders and their permission levels. Invaluable for access reviews, identifying over-privileged accounts, and understanding content exposure. Helps maintain principle of least privilege.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content member list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-member-type-external**
: Filter folder members. Keep only members that are external (not in the same team). Note: Invited members are marked as external member.

**-member-type-internal**
: Filter folder members. Keep only members that are internal (in the same team). Note: Invited members are marked as external member.

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: membership

This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.
The command will generate a report in three different formats. `membership.csv`, `membership.json`, and `membership.xlsx`.

| Column          | Description                                                                                                                          |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| path            | Path                                                                                                                                 |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)                             |
| owner_team_name | Team name of the team that owns the folder                                                                                           |
| access_type     | User's access level for this folder                                                                                                  |
| member_type     | Type of this member (user, group, or invitee)                                                                                        |
| member_name     | Name of this member                                                                                                                  |
| member_email    | Email address of this member                                                                                                         |
| same_team       | Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `membership_0000.xlsx`, `membership_0001.xlsx`, `membership_0002.xlsx`, ...

## Report: no_member

This report shows folders without members.
The command will generate a report in three different formats. `no_member.csv`, `no_member.json`, and `no_member.xlsx`.

| Column          | Description                                                                                              |
|-----------------|----------------------------------------------------------------------------------------------------------|
| owner_team_name | Team name of the team that owns the folder                                                               |
| path            | Path                                                                                                     |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `no_member_0000.xlsx`, `no_member_0001.xlsx`, `no_member_0002.xlsx`, ...

---
Title: dropbox team content member size
URL: https://toolbox.watermint.org/commands/dropbox/team/content/member/size.md
---

# dropbox team content member size

Calculate member counts for each team folder and shared folder, helping identify heavily accessed content and optimize permissions 

Analyzes folder membership density to identify over-shared content. High member counts may indicate security risks or performance issues. Use to prioritize permission reviews and identify candidates for access restriction or folder restructuring.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content member size 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-include-sub-folders**
: Include sub-folders to the report.. Default: false

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: member_count

Folder member count
The command will generate a report in three different formats. `member_count.csv`, `member_count.json`, and `member_count.xlsx`.

| Column                | Description                                                                                                                            |
|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| path                  | Path                                                                                                                                   |
| folder_type           | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)                               |
| owner_team_name       | Team name of the team that owns the folder                                                                                             |
| has_no_inherit        | True if the folder or any sub-folder does not inherit the access permission from the parent folder                                     |
| is_no_inherit         | True if the folder does not inherit the access from the parent folder                                                                  |
| capacity              | Capacity number for adding members. Empty if it's not able to determine by your permission (e.g. a folder contains an external group). |
| count_total           | Total number of members                                                                                                                |
| count_external_groups | Number of external teams' groups                                                                                                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_count_0000.xlsx`, `member_count_0001.xlsx`, `member_count_0002.xlsx`, ...

---
Title: dropbox team content mount list
URL: https://toolbox.watermint.org/commands/dropbox/team/content/mount/list.md
---

# dropbox team content mount list

Display mount status of all shared folders for team members, identifying which folders are actively synced to member devices 

Shows which shared folders are actively syncing to member devices versus cloud-only access. Critical for bandwidth planning, identifying heavy sync users, and troubleshooting sync issues. Helps optimize storage usage on user devices.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content mount list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Filter members. Filter by email address.

**-member-name**
: Filter members. Filter by exact match to the name.

**-member-name-prefix**
: Filter members. Filter by name match to the prefix.

**-member-name-suffix**
: Filter members. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

# Results

## Report: mount

This report shows a list of shared folders.
The command will generate a report in three different formats. `mount.csv`, `mount.json`, and `mount.xlsx`.

| Column                   | Description                                                                                               |
|--------------------------|-----------------------------------------------------------------------------------------------------------|
| team_member_display_name | Team member display name.                                                                                 |
| team_member_email        | Team member email address                                                                                 |
| namespace_id             | Namespace Id                                                                                              |
| namespace_name           | Name of the folder.                                                                                       |
| access_type              | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| mount_path               | Mount path of this folder. The folder is not mounted if this field is empty.                              |
| is_inside_team_folder    | Whether this folder is inside of a team folder.                                                           |
| is_team_folder           | Whether this folder is a team folder.                                                                     |
| policy_manage_access     | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link       | Who links can be shared with.                                                                             |
| policy_member            | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info       | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name          | Team name of the team that owns the folder                                                                |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mount_0000.xlsx`, `mount_0001.xlsx`, `mount_0002.xlsx`, ...

---
Title: dropbox team content policy list
URL: https://toolbox.watermint.org/commands/dropbox/team/content/policy/list.md
---

# dropbox team content policy list

Review all access policies and restrictions applied to team folders and shared folders for governance compliance 

Comprehensive policy audit showing viewer info restrictions, shared link policies, and other governance settings. Essential for compliance verification and ensuring folders meet organizational security requirements. Identifies policy inconsistencies across folders.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team content policy list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: policy

This report shows a list of shared folders and team folders with their current policy settings.
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column               | Description                                                                                              |
|----------------------|----------------------------------------------------------------------------------------------------------|
| path                 | Path                                                                                                     |
| is_team_folder       | `true` if the folder is a team folder, or inside of a team folder                                        |
| owner_team_name      | Team name of the team that owns the folder                                                               |
| policy_manage_access | Who can add and remove members from this shared folder.                                                  |
| policy_shared_link   | Who links can be shared with.                                                                            |
| policy_member        | Who can be a member of this shared folder, taking into account both the folder and the team-wide policy. |
| policy_viewer_info   | Who can enable/disable viewer info for this shared folder.                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...

---
Title: dropbox team device list
URL: https://toolbox.watermint.org/commands/dropbox/team/device/list.md
---

# dropbox team device list

Display all devices and active sessions connected to team member accounts with device details and last activity timestamps 

Complete device inventory showing all connected devices, platforms, and session ages. Critical for security audits, identifying unauthorized devices, and managing device limits. Export data to track device sprawl and plan security policies.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team device list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: device

This report shows a list of current existing sessions in the team with team member information.
The command will generate a report in three different formats. `device.csv`, `device.json`, and `device.xlsx`.

| Column                        | Description                                                                          |
|-------------------------------|--------------------------------------------------------------------------------------|
| team_member_id                | ID of user as a member of a team.                                                    |
| email                         | Email address of user.                                                               |
| status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name                    | Also known as a first name                                                           |
| surname                       | Also known as a last name or family name.                                            |
| display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  |
| id                            | The session id.                                                                      |
| user_agent                    | Information on the hosting device.                                                   |
| os                            | Information on the hosting operating system                                          |
| browser                       | Information on the browser used for this web session.                                |
| ip_address                    | The IP address of the last activity from this session.                               |
| country                       | The country from which the last activity from this session was made.                 |
| created                       | The time this session was created.                                                   |
| updated                       | The time of the last activity from this session.                                     |
| expires                       | The time this session expires (optional)                                             |
| host_name                     | Name of the hosting desktop.                                                         |
| client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             |
| client_version                | The Dropbox client version.                                                          |
| platform                      | Information on the hosting platform.                                                 |
| is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             |
| device_name                   | The device name.                                                                     |
| os_version                    | The hosting OS version.                                                              |
| last_carrier                  | Last carrier used by the device (optional).                                          |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `device_0000.xlsx`, `device_0001.xlsx`, `device_0002.xlsx`, ...

---
Title: dropbox team device unlink
URL: https://toolbox.watermint.org/commands/dropbox/team/device/unlink.md
---

# dropbox team device unlink

Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access (Irreversible operation)

Immediately terminates device sessions, forcing re-authentication. Critical security tool for lost devices, departing employees, or suspicious activity. Device must reconnect and re-sync after unlinking. Consider member communication before bulk unlinking.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team device unlink -file /path/to/data/file.csv
```

## Options:

**-delete-on-unlink**
: Delete files on unlink. Default: false

**-file**
: Data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

This report shows a list of current existing sessions in the team with team member information.

| Column                        | Description                                                                          | Example                                                                                                         |
|-------------------------------|--------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| team_member_id                | ID of user as a member of a team.                                                    | dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                                                       |
| email                         | Email address of user.                                                               | john.smith@example.com                                                                                          |
| status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) | active                                                                                                          |
| given_name                    | Also known as a first name                                                           | John                                                                                                            |
| surname                       | Also known as a last name or family name.                                            | Smith                                                                                                           |
| familiar_name                 | Locale-dependent name                                                                | John Smith                                                                                                      |
| display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  | John Smith                                                                                                      |
| abbreviated_name              | An abbreviated form of the person's name.                                            | JS                                                                                                              |
| external_id                   | External ID that a team can attach to the user (optional)                            | (empty string if not set)                                                                                       |
| account_id                    | A user's account identifier.                                                         | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                                                        |
| device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  | desktop_client                                                                                                  |
| id                            | The session id.                                                                      | dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                                                      |
| user_agent                    | Information on the hosting device.                                                   | Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 |
| os                            | Information on the hosting operating system                                          | Windows                                                                                                         |
| browser                       | Information on the browser used for this web session.                                | Chrome                                                                                                          |
| ip_address                    | The IP address of the last activity from this session.                               | xx.xxx.x.xxx                                                                                                    |
| country                       | The country from which the last activity from this session was made.                 | United States                                                                                                   |
| created                       | The time this session was created.                                                   | 2019-09-20T23:47:33Z                                                                                            |
| updated                       | The time of the last activity from this session.                                     | 2019-10-25T04:42:16Z                                                                                            |
| expires                       | The time this session expires (optional)                                             | 2024-03-22T10:30:56Z                                                                                            |
| host_name                     | Name of the hosting desktop.                                                         | nihonbashi                                                                                                      |
| client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             | windows                                                                                                         |
| client_version                | The Dropbox client version.                                                          | 83.4.152                                                                                                        |
| platform                      | Information on the hosting platform.                                                 | Windows 10 1903                                                                                                 |
| is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             | TRUE                                                                                                            |
| device_name                   | The device name.                                                                     | My Awesome PC                                                                                                   |
| os_version                    | The hosting OS version.                                                              | (empty string if not set)                                                                                       |
| last_carrier                  | Last carrier used by the device (optional).                                          | AT&T                                                                                                            |

The first line is a header line. The program will accept a file without the header.
```
team_member_id,email,status,given_name,surname,familiar_name,display_name,abbreviated_name,external_id,account_id,device_tag,id,user_agent,os,browser,ip_address,country,created,updated,expires,host_name,client_type,client_version,platform,is_delete_on_unlink_supported,device_name,os_version,last_carrier
dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,john.smith@example.com,active,John,Smith,John Smith,John Smith,JS,(empty string if not set),dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,desktop_client,dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",Windows,Chrome,xx.xxx.x.xxx,United States,2019-09-20T23:47:33Z,2019-10-25T04:42:16Z,2024-03-22T10:30:56Z,nihonbashi,windows,83.4.152,Windows 10 1903,TRUE,My Awesome PC,(empty string if not set),AT&T
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                              | Description                                                                          |
|-------------------------------------|--------------------------------------------------------------------------------------|
| status                              | Status of the operation                                                              |
| reason                              | Reason of failure or skipped operation                                               |
| input.team_member_id                | ID of user as a member of a team.                                                    |
| input.email                         | Email address of user.                                                               |
| input.status                        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| input.given_name                    | Also known as a first name                                                           |
| input.surname                       | Also known as a last name or family name.                                            |
| input.display_name                  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| input.device_tag                    | Type of the session (web_session, desktop_client, or mobile_client)                  |
| input.id                            | The session id.                                                                      |
| input.user_agent                    | Information on the hosting device.                                                   |
| input.os                            | Information on the hosting operating system                                          |
| input.browser                       | Information on the browser used for this web session.                                |
| input.ip_address                    | The IP address of the last activity from this session.                               |
| input.country                       | The country from which the last activity from this session was made.                 |
| input.created                       | The time this session was created.                                                   |
| input.updated                       | The time of the last activity from this session.                                     |
| input.expires                       | The time this session expires (optional)                                             |
| input.host_name                     | Name of the hosting desktop.                                                         |
| input.client_type                   | The Dropbox desktop client type (windows, mac, or linux)                             |
| input.client_version                | The Dropbox client version.                                                          |
| input.platform                      | Information on the hosting platform.                                                 |
| input.is_delete_on_unlink_supported | Whether it's possible to delete all of the account files upon unlinking.             |
| input.device_name                   | The device name.                                                                     |
| input.os_version                    | The hosting OS version.                                                              |
| input.last_carrier                  | Last carrier used by the device (optional).                                          |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team filerequest list
URL: https://toolbox.watermint.org/commands/dropbox/team/filerequest/list.md
---

# dropbox team filerequest list

Display all active and closed file requests created by team members, helping track external file collection activities 

Comprehensive view of all file requests across the team. Monitor external data collection, identify abandoned requests, and ensure compliance with data handling policies. Includes request URLs, creators, and submission counts for audit purposes.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team filerequest list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: file_request

This report shows a list of file requests with the file request owner team member.
The command will generate a report in three different formats. `file_request.csv`, `file_request.json`, and `file_request.xlsx`.

| Column                      | Description                                                                   |
|-----------------------------|-------------------------------------------------------------------------------|
| email                       | Email address of this file request owner.                                     |
| status                      | The user status of this file request owner (active/invited/suspended/removed) |
| surname                     | Surname of this file request owner.                                           |
| given_name                  | Given name of this file request owner.                                        |
| url                         | The URL of the file request.                                                  |
| title                       | The title of the file request.                                                |
| created                     | When this file request was created.                                           |
| is_open                     | Whether or not the file request is open.                                      |
| file_count                  | The number of files this file request has received.                           |
| destination                 | The path of the folder in the Dropbox where uploaded files will be sent       |
| deadline                    | The deadline for this file request.                                           |
| deadline_allow_late_uploads | If set, allow uploads after the deadline has passed                           |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_request_0000.xlsx`, `file_request_0001.xlsx`, `file_request_0002.xlsx`, ...

---
Title: dropbox team group add
URL: https://toolbox.watermint.org/commands/dropbox/team/group/add.md
---

# dropbox team group add

Create a new group in your team for organizing members and managing permissions collectively (Irreversible operation)

Creates groups for logical organization of team members. Groups simplify permission management by allowing bulk operations. Consider naming conventions for easy identification. Groups can be company-managed or member-managed depending on governance needs.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group add -name GROUP_NAME
```

## Options:

**-management-type**
: Group management type `company_managed` or `user_managed`. Options: company_managed (Managed by company administrators), user_managed (Managed by individual users). Default: company_managed

**-name**
: Group name

**-peer**
: Account alias. Default: default

# Results

## Report: added_group

This report shows a list of groups in the team.
The command will generate a report in three different formats. `added_group.csv`, `added_group.json`, and `added_group.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `added_group_0000.xlsx`, `added_group_0001.xlsx`, `added_group_0002.xlsx`, ...

---
Title: dropbox team group delete
URL: https://toolbox.watermint.org/commands/dropbox/team/group/delete.md
---

# dropbox team group delete

Remove a specific group from your team, automatically removing all member associations (Irreversible operation)

Permanently deletes a group and removes all member associations. Members retain access through other groups or individual permissions. Cannot be undone - consider archiving group by removing members instead if unsure. Folder permissions using this group are also removed.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group delete -name GROUP_NAME
```

## Options:

**-name**
: Group name

**-peer**
: Account alias. Default: default

---
Title: dropbox team group list
URL: https://toolbox.watermint.org/commands/dropbox/team/group/list.md
---

# dropbox team group list

Display all groups in your team with member counts and group management types 

Complete inventory of team groups showing sizes and management modes. Use to identify empty groups, oversized groups, or groups needing management type changes. Export for regular auditing and compliance documentation.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: group

This report shows a list of groups in the team.
The command will generate a report in three different formats. `group.csv`, `group.json`, and `group.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`, ...

---
Title: dropbox team group rename
URL: https://toolbox.watermint.org/commands/dropbox/team/group/rename.md
---

# dropbox team group rename

Change the name of an existing group to better reflect its purpose or organizational changes (Irreversible operation)

Updates the display name of a group while maintaining all members and permissions. Useful when departments restructure, projects change names, or group purposes evolve. The rename is immediate and affects all references to the group throughout the system.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group rename -current-name CURRENT_NAME -new-name NEW_NAME
```

## Options:

**-current-name**
: Current group name

**-new-name**
: New group name

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                       | Description                                                                           |
|------------------------------|---------------------------------------------------------------------------------------|
| status                       | Status of the operation                                                               |
| reason                       | Reason of failure or skipped operation                                                |
| input.current_name           | Current group name                                                                    |
| input.new_name               | New group name                                                                        |
| result.group_name            | Name of a group                                                                       |
| result.group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| result.member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group batch add
URL: https://toolbox.watermint.org/commands/dropbox/team/group/batch/add.md
---

# dropbox team group batch add

Create multiple groups at once using batch processing, efficient for large-scale team organization 

Bulk creates groups from a data file, ideal for initial setup or reorganizations. Validates all groups before creation to prevent partial failures. Include external IDs for integration with identity management systems. Significantly faster than individual creation.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group batch add -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-management-type**
: Who is allowed to manage the group (user_managed, company_managed, or system_managed). Options: company_managed (Managed by company administrators), user_managed (Managed by individual users). Default: company_managed

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for batch operation to groups.

| Column | Description | Example |
|--------|-------------|---------|
| name   | Group name  | Sales   |

The first line is a header line. The program will accept a file without the header.
```
name
Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                       | Description                                                                           |
|------------------------------|---------------------------------------------------------------------------------------|
| status                       | Status of the operation                                                               |
| reason                       | Reason of failure or skipped operation                                                |
| input.name                   | Group name                                                                            |
| result.group_name            | Name of a group                                                                       |
| result.group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| result.member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group batch delete
URL: https://toolbox.watermint.org/commands/dropbox/team/group/batch/delete.md
---

# dropbox team group batch delete

Remove multiple groups from your team in batch, streamlining group cleanup and reorganization (Irreversible operation)

Efficiently removes multiple groups in a single operation. Useful for organizational restructuring or cleaning up obsolete groups. Members retain individual permissions but lose group-based access. Verify group contents before deletion as this is irreversible.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group batch delete -file /path/to/file.csv
```

## Options:

**-file**
: Data file for group name list

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for batch operation to groups.

| Column | Description | Example |
|--------|-------------|---------|
| name   | Group name  | Sales   |

The first line is a header line. The program will accept a file without the header.
```
name
Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                       | Description                                                                           |
|------------------------------|---------------------------------------------------------------------------------------|
| status                       | Status of the operation                                                               |
| reason                       | Reason of failure or skipped operation                                                |
| input.name                   | Group name                                                                            |
| result.group_name            | Name of a group                                                                       |
| result.group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| result.member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group clear externalid
URL: https://toolbox.watermint.org/commands/dropbox/team/group/clear/externalid.md
---

# dropbox team group clear externalid

Remove external ID mappings from groups, useful when disconnecting from external identity providers 

Removes external ID associations from groups when migrating away from identity providers or changing integration systems. Group functionality remains intact but loses external system mapping. Useful for troubleshooting sync issues with identity providers.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group clear externalid -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Group name

| Column | Description                           | Example |
|--------|---------------------------------------|---------|
| name   | Name of group to clear an external ID | Sales   |

The first line is a header line. The program will accept a file without the header.
```
name
Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                       | Description                                                                           |
|------------------------------|---------------------------------------------------------------------------------------|
| status                       | Status of the operation                                                               |
| reason                       | Reason of failure or skipped operation                                                |
| input.name                   | Name of group to clear an external ID                                                 |
| result.group_name            | Name of a group                                                                       |
| result.group_id              | Group ID                                                                              |
| result.group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| result.group_external_id     | External ID that a team can attach to the group.                                      |
| result.member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group folder list
URL: https://toolbox.watermint.org/commands/dropbox/team/group/folder/list.md
---

# dropbox team group folder list

Display all folders accessible by each group, showing group-based content organization and permissions 

Maps group permissions to folders, revealing content access patterns. Essential for access reviews and understanding permission inheritance. Helps identify over-permissioned groups and optimize folder structures for security.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group folder list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-group-name**
: Filter by group name. Filter by exact match to the name.

**-group-name-prefix**
: Filter by group name. Filter by name match to the prefix.

**-group-name-suffix**
: Filter by group name. Filter by name match to the suffix.

**-include-external-groups**
: Include external groups in the report.. Default: false

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: group_to_folder

Group to folder mapping.
The command will generate a report in three different formats. `group_to_folder.csv`, `group_to_folder.json`, and `group_to_folder.xlsx`.

| Column             | Description                                                                                              |
|--------------------|----------------------------------------------------------------------------------------------------------|
| group_name         | Name of a group                                                                                          |
| group_type         | Who is allowed to manage the group (user_managed, company_managed, or system_managed)                    |
| group_is_same_team | 'true' if a group is in the same team. Otherwise false.                                                  |
| access_type        | Group's access level for this folder                                                                     |
| namespace_name     | The name of this namespace                                                                               |
| path               | Path                                                                                                     |
| folder_type        | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder) |
| owner_team_name    | Team name of the team that owns the folder                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `group_to_folder_0000.xlsx`, `group_to_folder_0001.xlsx`, `group_to_folder_0002.xlsx`, ...

## Report: group_with_no_folders

This report shows a list of groups in the team.
The command will generate a report in three different formats. `group_with_no_folders.csv`, `group_with_no_folders.json`, and `group_with_no_folders.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `group_with_no_folders_0000.xlsx`, `group_with_no_folders_0001.xlsx`, `group_with_no_folders_0002.xlsx`, ...

---
Title: dropbox team group member add
URL: https://toolbox.watermint.org/commands/dropbox/team/group/member/add.md
---

# dropbox team group member add

Add individual team members to a specific group for centralized permission management (Irreversible operation)

Adds members to groups for inherited permissions and simplified management. Changes take effect immediately for folder access. Consider group size limits and performance implications for very large groups.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group member add -group-name GROUP_NAME -member-email EMAIL
```

## Options:

**-group-name**
: Group name

**-member-email**
: Email address of the member

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                       | Description                                                                           |
|------------------------------|---------------------------------------------------------------------------------------|
| status                       | Status of the operation                                                               |
| reason                       | Reason of failure or skipped operation                                                |
| input.group_name             | Name of the group                                                                     |
| input.member_email           | Email address of the member                                                           |
| result.group_name            | Name of a group                                                                       |
| result.group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| result.member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group member delete
URL: https://toolbox.watermint.org/commands/dropbox/team/group/member/delete.md
---

# dropbox team group member delete

Remove a specific member from a group while preserving their other group memberships (Irreversible operation)

Removes an individual member from a single group without affecting their membership in other groups. Use for targeted permission adjustments or when members change departments. The removal takes effect immediately, revoking any inherited permissions from that group.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group member delete -group-name GROUP_NAME -member-email EMAIL
```

## Options:

**-group-name**
: Name of the group

**-member-email**
: Email address of the member

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                       | Description                                                                           |
|------------------------------|---------------------------------------------------------------------------------------|
| status                       | Status of the operation                                                               |
| reason                       | Reason of failure or skipped operation                                                |
| input.group_name             | Name of the group                                                                     |
| input.member_email           | Email address of the member                                                           |
| result.group_name            | Name of a group                                                                       |
| result.group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| result.member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group member list
URL: https://toolbox.watermint.org/commands/dropbox/team/group/member/list.md
---

# dropbox team group member list

Display all members belonging to each group, useful for auditing group compositions and access rights 

Lists all groups with their complete member rosters. Essential for access audits, verifying group compositions, and understanding permission inheritance. Helps identify empty groups, over-privileged groups, or members with unexpected access through group membership.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group member list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: group_member

This report shows a list of groups and their members.
The command will generate a report in three different formats. `group_member.csv`, `group_member.json`, and `group_member.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group.                                                                      |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| access_type           | The role that the user has in the group (member/owner)                                |
| email                 | Email address of user.                                                                |
| status                | The user's status as a member of a specific team. (active/invited/suspended/removed)  |
| surname               | Also known as a last name or family name.                                             |
| given_name            | Also known as a first name                                                            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `group_member_0000.xlsx`, `group_member_0001.xlsx`, `group_member_0002.xlsx`, ...

---
Title: dropbox team group member batch add
URL: https://toolbox.watermint.org/commands/dropbox/team/group/member/batch/add.md
---

# dropbox team group member batch add

Add multiple members to groups efficiently using batch processing, ideal for large team reorganizations (Irreversible operation)

Bulk adds members to groups using a mapping file. Validates all memberships before applying changes. Ideal for onboarding, departmental changes, or permission standardization projects. Handles errors gracefully with detailed reporting.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group member batch add -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Add members into groups

| Column       | Description          | Example          |
|--------------|----------------------|------------------|
| group_name   | Group name           | Sales            |
| member_email | Member email address | taro@example.com |

The first line is a header line. The program will accept a file without the header.
```
group_name,member_email
Sales,taro@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                            |
|-------------------|----------------------------------------|
| status            | Status of the operation                |
| reason            | Reason of failure or skipped operation |
| input.GroupName   | Group name                             |
| input.MemberEmail | Member email address                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group member batch delete
URL: https://toolbox.watermint.org/commands/dropbox/team/group/member/batch/delete.md
---

# dropbox team group member batch delete

Remove multiple members from groups in batch, streamlining group membership management (Irreversible operation)

Bulk removes members from groups using a CSV file mapping. Validates all memberships before making changes. Useful for organizational restructuring, offboarding processes, or cleaning up group memberships. Processes efficiently with detailed error reporting for any issues.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group member batch delete -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Add members into groups

| Column       | Description          | Example          |
|--------------|----------------------|------------------|
| group_name   | Group name           | Sales            |
| member_email | Member email address | taro@example.com |

The first line is a header line. The program will accept a file without the header.
```
group_name,member_email
Sales,taro@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                            |
|-------------------|----------------------------------------|
| status            | Status of the operation                |
| reason            | Reason of failure or skipped operation |
| input.GroupName   | Group name                             |
| input.MemberEmail | Member email address                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group member batch update
URL: https://toolbox.watermint.org/commands/dropbox/team/group/member/batch/update.md
---

# dropbox team group member batch update

Update group memberships in bulk by adding or removing members, optimizing group composition changes (Irreversible operation)

Modifies group memberships in bulk based on a CSV file. Can both add and remove members in a single operation. Ideal for large-scale reorganizations where group compositions need significant updates. Maintains audit trail of all changes made.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group member batch update -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Add members into groups

| Column       | Description          | Example          |
|--------------|----------------------|------------------|
| group_name   | Group name           | Sales            |
| member_email | Member email address | taro@example.com |

The first line is a header line. The program will accept a file without the header.
```
group_name,member_email
Sales,taro@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                            |
|-------------------|----------------------------------------|
| status            | Status of the operation                |
| reason            | Reason of failure or skipped operation |
| input.GroupName   | Group name                             |
| input.MemberEmail | Member email address                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team group update type
URL: https://toolbox.watermint.org/commands/dropbox/team/group/update/type.md
---

# dropbox team group update type

Change how a group is managed (user-managed vs company-managed), affecting who can modify group membership 

Modifies group management settings to control who can add or remove members. Company-managed groups restrict modifications to admins, while user-managed groups allow designated members to manage membership. Critical for implementing proper governance and access control policies.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team group update type -name GROUP_NAME
```

## Options:

**-name**
: Group name

**-peer**
: Account alias. Default: default

**-type**
: Group type (user_managed/company_managed). Options: user_managed (type: user_managed), company_managed (type: company_managed). Default: company_managed

# Results

## Report: group

This report shows a list of groups in the team.
The command will generate a report in three different formats. `group.csv`, `group.json`, and `group.xlsx`.

| Column                | Description                                                                           |
|-----------------------|---------------------------------------------------------------------------------------|
| group_name            | Name of a group                                                                       |
| group_id              | Group ID                                                                              |
| group_management_type | Who is allowed to manage the group (user_managed, company_managed, or system_managed) |
| group_external_id     | External ID that a team can attach to the group.                                      |
| member_count          | The number of members in the group.                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `group_0000.xlsx`, `group_0001.xlsx`, `group_0002.xlsx`, ...

---
Title: dropbox team insight scan
URL: https://toolbox.watermint.org/commands/dropbox/team/insight/scan.md
---

# dropbox team insight scan

Perform comprehensive data scanning across your team for analytics and insights generation 

This command collects various team data, such as files in team folders, permissions and shared links, and stores them in a database.
The collected data can be analysed with commands such as `dropbox team insight report teamfoldermember`, or with database tools that support SQLite in general.

About how long a scan takes:

Scanning a team often takes a long time. Especially if there are a large number of files stored, the time is linearly proportional to the number of files. To increase the scanning speed, it is better to use the `-concurrency` option for parallel processing.
However, too much parallelism will increase the error rate from the Dropbox server, so a balance must be considered. According to the results of a few benchmarks, a parallelism level of 12-24 for the `-concurrency` option seems to be a good choice.
The time required for scanning depends on the response of the Dropbox server, but is around 20-30 hours per 10 million files (with `-concurrency 16`).

During the scan, users might delete, move or add files during that time. The command does not aim to capture all those differences and report exact results, but to provide rough information as quickly as possible.

For database file sizes:

As this command retrieves all metadata, including the team's files, the size of the database increases with the size of those metadata. Benchmark results show that the database size is around 10-12 GB per 10 million files stored in the team. Make sure that the path specified by `-database` has enough space before running.

About scan errors:

The Dropbox server may return an error when running the scan. The command will automatically try to re-run the scan several times, but the error may not be resolved for a certain period of time due to server congestion or condition. In that case, the command stops the re-run and records the scan task in the database where the error occurred.
If you want to re-run a failed scan, use the `dropbox team insight scanretry` command to run the scan again.
If the issue is not resolved after repeated re-runs and you want to analyse only the coverage of the current scan, you need to perform an aggregation task before the analysis. Aggregation tasks can be performed with the `dropbox team insight summary` command.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team insight scan -database /LOCAL/PATH/TO/database
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-database**
: Path to database

**-max-retries**
: Maximum number of retries. Default: 3

**-peer**
: Account alias. Default: default

**-scan-member-folders**
: Scan member folders. Default: false

**-skip-summarize**
: Skip summarize tasks. Default: false

# Results

## Report: errors

Error report
The command will generate a report in three different formats. `errors.csv`, `errors.json`, and `errors.xlsx`.

| Column   | Description    |
|----------|----------------|
| category | Error category |
| message  | Error message  |
| tag      | Error tag      |
| detail   | Error details  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...

---
Title: dropbox team insight report teamfoldermember
URL: https://toolbox.watermint.org/commands/dropbox/team/insight/report/teamfoldermember.md
---

# dropbox team insight report teamfoldermember

Generate detailed reports on team folder membership, showing access patterns and member distribution 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team insight report teamfoldermember -database /LOCAL/PATH/TO/database
```

## Options:

**-database**
: Path to database

# Results

## Report: entry

Team folder member
The command will generate a report in three different formats. `entry.csv`, `entry.json`, and `entry.xlsx`.

| Column              | Description                            |
|---------------------|----------------------------------------|
| team_folder_id      | Team folder ID                         |
| team_folder_name    | Team folder name                       |
| path_display        | Display path                           |
| access_type         | Access type                            |
| is_inherited        | Inherited access                       |
| member_type         | Member type                            |
| same_team           | True if the member is in the same team |
| group_id            | Group ID                               |
| group_name          | Group name                             |
| group_type          | Group management type                  |
| group_member_count  | Group member count                     |
| invitee_email       | Invitee email                          |
| user_team_member_id | User team member ID                    |
| user_email          | User email                             |
| user_display_name   | User display name                      |
| user_account_id     | User account ID                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `entry_0000.xlsx`, `entry_0001.xlsx`, `entry_0002.xlsx`, ...

---
Title: dropbox team legalhold add
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/add.md
---

# dropbox team legalhold add

Create a legal hold policy to preserve specified team content for compliance or litigation purposes 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold add -member /PATH/TO/member_email.csv -name POLICY_NAME
```

## Options:

**-description**
: A description of the legal hold policy.

**-end-date**
: End date of the legal hold policy.

**-member**
: Email of the member or members you want to place a hold on

**-name**
: Policy name.

**-peer**
: Account alias. Default: default

**-start-date**
: Start date of the legal hold policy.

# File formats

## Format: Member

Member email address

| Column | Description               | Example          |
|--------|---------------------------|------------------|
| email  | Team member email address | emma@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
emma@example.com
```

# Results

## Report: policy

Legal hold policy
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column                    | Description                                     |
|---------------------------|-------------------------------------------------|
| id                        | The legal hold id.                              |
| name                      | Policy name.                                    |
| description               | A description of the legal hold policy.         |
| status                    | The current state of the hold.                  |
| start_date                | Start date of the legal hold policy.            |
| end_date                  | End date of the legal hold policy.              |
| activation_time           | The time at which the legal hold was activated. |
| permanently_deleted_users | Number of users permanently removed.            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...

---
Title: dropbox team legalhold list
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/list.md
---

# dropbox team legalhold list

Display all active legal hold policies with their details, members, and preservation status 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold list 
```

## Options:

**-include-released**
: Whether to return holds that were released.. Default: false

**-peer**
: Account alias. Default: default

# Results

## Report: policies

Legal hold policy
The command will generate a report in three different formats. `policies.csv`, `policies.json`, and `policies.xlsx`.

| Column                    | Description                                     |
|---------------------------|-------------------------------------------------|
| id                        | The legal hold id.                              |
| name                      | Policy name.                                    |
| description               | A description of the legal hold policy.         |
| status                    | The current state of the hold.                  |
| start_date                | Start date of the legal hold policy.            |
| end_date                  | End date of the legal hold policy.              |
| activation_time           | The time at which the legal hold was activated. |
| permanently_deleted_users | Number of users permanently removed.            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policies_0000.xlsx`, `policies_0001.xlsx`, `policies_0002.xlsx`, ...

---
Title: dropbox team legalhold release
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/release.md
---

# dropbox team legalhold release

Release a legal hold policy and restore normal file operations for affected members and content 

Ends a legal hold policy and removes preservation requirements. Content becomes subject to normal retention and deletion policies again. Use when litigation concludes or preservation is no longer required. The release is logged for audit purposes but cannot be undone.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold release -policy-id POLICY_ID
```

## Options:

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID

---
Title: dropbox team legalhold member list
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/member/list.md
---

# dropbox team legalhold member list

Display all team members currently under legal hold policies with their preservation status 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold member list -policy-id POLICY_ID
```

## Options:

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID

# Results

## Report: member

This report shows a list of members.
The command will generate a report in three different formats. `member.csv`, `member.json`, and `member.xlsx`.

| Column           | Description                                                                                                          |
|------------------|----------------------------------------------------------------------------------------------------------------------|
| team_member_id   | ID of user as a member of a team.                                                                                    |
| email            | Email address of user.                                                                                               |
| email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| given_name       | Also known as a first name                                                                                           |
| surname          | Also known as a last name or family name.                                                                            |
| familiar_name    | Locale-dependent name                                                                                                |
| display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| abbreviated_name | An abbreviated form of the person's name.                                                                            |
| member_folder_id | The namespace id of the user's root folder.                                                                          |
| external_id      | External ID that a team can attach to the user.                                                                      |
| account_id       | A user's account identifier.                                                                                         |
| persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| invited_on       | The date and time the user was invited to the team                                                                   |
| role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| tag              | Operation tag                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...

---
Title: dropbox team legalhold member batch update
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/member/batch/update.md
---

# dropbox team legalhold member batch update

Add or remove multiple team members from legal hold policies in batch for efficient compliance management 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold member batch update -member /PATH/TO/MEMBER_LIST.csv -policy-id POLICY_ID
```

## Options:

**-member**
: Path to member list file

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID

# File formats

## Format: Member

Member email address

| Column | Description   | Example          |
|--------|---------------|------------------|
| email  | Email address | emma@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
emma@example.com
```

# Results

## Report: policy

This report shows a list of members.
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column           | Description                                                                                                          |
|------------------|----------------------------------------------------------------------------------------------------------------------|
| team_member_id   | ID of user as a member of a team.                                                                                    |
| email            | Email address of user.                                                                                               |
| email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| given_name       | Also known as a first name                                                                                           |
| surname          | Also known as a last name or family name.                                                                            |
| familiar_name    | Locale-dependent name                                                                                                |
| display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| abbreviated_name | An abbreviated form of the person's name.                                                                            |
| member_folder_id | The namespace id of the user's root folder.                                                                          |
| external_id      | External ID that a team can attach to the user.                                                                      |
| account_id       | A user's account identifier.                                                                                         |
| persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| invited_on       | The date and time the user was invited to the team                                                                   |
| role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| tag              | Operation tag                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...

---
Title: dropbox team legalhold revision list
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/revision/list.md
---

# dropbox team legalhold revision list

Display all file revisions preserved under legal hold policies, ensuring comprehensive data retention 

Shows the complete revision history of files under legal hold including all modifications. Tracks file versions preserved by the policy to ensure nothing is lost. Critical for maintaining defensible preservation records and demonstrating compliance with legal requirements.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold revision list -after DATE_TIME -policy-id POLICY_ID
```

## Options:

**-after**
: Get revisions after this specified date and time

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID.

# Results

## Report: revision

Revision
The command will generate a report in three different formats. `revision.csv`, `revision.json`, and `revision.xlsx`.

| Column | Description |
|--------|-------------|

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `revision_0000.xlsx`, `revision_0001.xlsx`, `revision_0002.xlsx`, ...

---
Title: dropbox team legalhold update desc
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/update/desc.md
---

# dropbox team legalhold update desc

Modify the description of an existing legal hold policy to reflect changes in scope or purpose 

Updates the description field of a legal hold policy for better documentation. Useful for adding case references, updating matter details, or clarifying preservation scope. Changes are tracked in the revision history for audit purposes.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold update desc -desc DESCRIPTION -policy-id POLICY_ID
```

## Options:

**-desc**
: New description

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID

# Results

## Report: policy

Legal hold policy
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column                    | Description                                     |
|---------------------------|-------------------------------------------------|
| id                        | The legal hold id.                              |
| name                      | Policy name.                                    |
| description               | A description of the legal hold policy.         |
| status                    | The current state of the hold.                  |
| start_date                | Start date of the legal hold policy.            |
| end_date                  | End date of the legal hold policy.              |
| activation_time           | The time at which the legal hold was activated. |
| permanently_deleted_users | Number of users permanently removed.            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...

---
Title: dropbox team legalhold update name
URL: https://toolbox.watermint.org/commands/dropbox/team/legalhold/update/name.md
---

# dropbox team legalhold update name

Change the name of a legal hold policy for better identification and organization 

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team legalhold update name -name NEW_NAME -policy-id POLICY_ID
```

## Options:

**-name**
: New name

**-peer**
: Account alias. Default: default

**-policy-id**
: Legal hold policy ID

# Results

## Report: policy

Legal hold policy
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column                    | Description                                     |
|---------------------------|-------------------------------------------------|
| id                        | The legal hold id.                              |
| name                      | Policy name.                                    |
| description               | A description of the legal hold policy.         |
| status                    | The current state of the hold.                  |
| start_date                | Start date of the legal hold policy.            |
| end_date                  | End date of the legal hold policy.              |
| activation_time           | The time at which the legal hold was activated. |
| permanently_deleted_users | Number of users permanently removed.            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...

---
Title: dropbox team linkedapp list
URL: https://toolbox.watermint.org/commands/dropbox/team/linkedapp/list.md
---

# dropbox team linkedapp list

Display all third-party applications linked to team member accounts for security auditing and access control 

Lists all third-party applications with access to team members' Dropbox accounts. Essential for security audits, identifying unauthorized apps, and managing OAuth integrations. Shows which members use which apps, helping enforce application policies and identify potential security risks.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team linkedapp list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: linked_app

This report shows a list of linked apps with the users of the apps.
The command will generate a report in three different formats. `linked_app.csv`, `linked_app.json`, and `linked_app.xlsx`.

| Column        | Description                                                                          |
|---------------|--------------------------------------------------------------------------------------|
| email         | Email address of user.                                                               |
| status        | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name    | Also known as a first name                                                           |
| surname       | Also known as a last name or family name.                                            |
| display_name  | A name that can be used directly to represent the name of a user's Dropbox account.  |
| app_name      | The application name.                                                                |
| is_app_folder | Whether the linked application uses a dedicated folder.                              |
| publisher     | The application publisher name.                                                      |
| publisher_url | The publisher's URL.                                                                 |
| linked        | The time this application was linked                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `linked_app_0000.xlsx`, `linked_app_0001.xlsx`, `linked_app_0002.xlsx`, ...

---
Title: dropbox team member feature
URL: https://toolbox.watermint.org/commands/dropbox/team/member/feature.md
---

# dropbox team member feature

Display feature settings and capabilities enabled for specific team members, helping understand member permissions 

Shows which features and capabilities are enabled for team members. Useful for troubleshooting access issues, verifying feature rollouts, and understanding member capabilities. Helps identify why certain members can or cannot access specific functionality.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member feature 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: features

Member features
The command will generate a report in three different formats. `features.csv`, `features.json`, and `features.xlsx`.

| Column         | Description                                                                                                                                       |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| email          | Email address of the team member                                                                                                                  |
| paper_as_files | When this value is true, the user's Paper docs are accessible in Dropbox with the .paper extension and must be accessed via the /files endpoints. |
| file_locking   | When this value is True, the user can lock files in shared folders.                                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `features_0000.xlsx`, `features_0001.xlsx`, `features_0002.xlsx`, ...

---
Title: dropbox team member list
URL: https://toolbox.watermint.org/commands/dropbox/team/member/list.md
---

# dropbox team member list

Display comprehensive list of all team members with their status, roles, and account details 

Provides complete team roster including active, suspended, and optionally deleted members. Shows email addresses, names, roles, and account status. Fundamental for team audits, license management, and understanding team composition. Export for HR or compliance reporting.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member list 
```

## Options:

**-include-deleted**
: Include deleted members.. Default: false

**-peer**
: Account alias. Default: default

# Results

## Report: member

This report shows a list of members.
The command will generate a report in three different formats. `member.csv`, `member.json`, and `member.xlsx`.

| Column         | Description                                                                                    |
|----------------|------------------------------------------------------------------------------------------------|
| email          | Email address of user.                                                                         |
| email_verified | Is true if the user's email is verified to be owned by the user.                               |
| status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| given_name     | Also known as a first name                                                                     |
| surname        | Also known as a last name or family name.                                                      |
| display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| joined_on      | The date and time the user joined as a member of a specific team.                              |
| invited_on     | The date and time the user was invited to the team                                             |
| role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...

---
Title: dropbox team member replication
URL: https://toolbox.watermint.org/commands/dropbox/team/member/replication.md
---

# dropbox team member replication

Replicate all files from one team member's account to another, useful for account transitions or backups (Irreversible operation)

Creates complete copies of member data between accounts, preserving folder structures and sharing where possible. Essential for role transitions, creating backups, or merging accounts. Requires sufficient storage in destination account. Consider using batch processing for multiple replications.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member replication -file /path/to/file.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dst**
: Destination team; team file access. Default: dst

**-file**
: Data file

**-src**
: Source team; team file access. Default: src

# File formats

## Format: File

Data file for replicating member contents.

| Column    | Description                         | Example                |
|-----------|-------------------------------------|------------------------|
| src_email | Source account's email address      | john@example.net       |
| dst_email | Destination account's email address | john.smith@example.com |

The first line is a header line. The program will accept a file without the header.
```
src_email,dst_email
john@example.net,john.smith@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column          | Description                            |
|-----------------|----------------------------------------|
| status          | Status of the operation                |
| reason          | Reason of failure or skipped operation |
| input.src_email | Source account's email address         |
| input.dst_email | Destination account's email address    |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member suspend
URL: https://toolbox.watermint.org/commands/dropbox/team/member/suspend.md
---

# dropbox team member suspend

Temporarily suspend a team member's access to their account while preserving all data and settings 

Immediately blocks member access while maintaining all data, settings, and group memberships. Use for security incidents, policy violations, or temporary leaves. Choose whether to keep data on linked devices. Member can be unsuspended later with full access restored.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member suspend -email EMAIL
```

## Options:

**-email**
: Member's email address

**-keep-data**
: Keep the user's data on their linked devices. Default: false

**-peer**
: Account alias. Default: default

---
Title: dropbox team member unsuspend
URL: https://toolbox.watermint.org/commands/dropbox/team/member/unsuspend.md
---

# dropbox team member unsuspend

Restore access for a suspended team member, reactivating their account and all associated permissions 

Reactivates a suspended member's account, restoring full access to data and team resources. All previous permissions, group memberships, and settings are preserved. Use when suspension reasons are resolved or members return from leave.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member unsuspend -email EMAIL
```

## Options:

**-email**
: Member's email address

**-peer**
: Account alias. Default: default

---
Title: dropbox team member batch delete
URL: https://toolbox.watermint.org/commands/dropbox/team/member/batch/delete.md
---

# dropbox team member batch delete

Remove multiple team members in batch, efficiently managing team departures and access revocation (Irreversible operation)

Bulk removes team members while preserving their data through transfers. Requires specifying destination member for file transfers and admin notification email. Ideal for layoffs, department closures, or mass offboarding. Optionally wipes data from linked devices for security.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member batch delete -file /PATH/TO/member_list.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

**-transfer-dest-member**
: If provided, files from the deleted member account will be transferred to this user.

**-transfer-notify-admin-email-on-error**
: If provided, errors during the transfer process will be sent via email to this user.

**-wipe-data**
: If true, controls if the user's data will be deleted on their linked devices. Default: true

# File formats

## Format: File

Data file for deleting team members.

| Column | Description                  | Example          |
|--------|------------------------------|------------------|
| email  | Email address of the account | john@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
john@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column      | Description                            |
|-------------|----------------------------------------|
| status      | Status of the operation                |
| reason      | Reason of failure or skipped operation |
| input.email | Email address of the account           |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member batch detach
URL: https://toolbox.watermint.org/commands/dropbox/team/member/batch/detach.md
---

# dropbox team member batch detach

Convert multiple team accounts to individual Basic accounts, preserving personal data while removing team access (Irreversible operation)

Bulk converts team members to personal Dropbox Basic accounts. Members retain their files but lose team features and shared folder access. Useful for contractors ending engagements or when downsizing teams. Consider data retention policies before detaching.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member batch detach -file /PATH/TO/member_list.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

**-revoke-team-shares**
: True to revoke shared folder access owned by the team. Default: false

# File formats

## Format: File

Data file for converting team members into Dropbox Basic account.

| Column | Description                  | Example          |
|--------|------------------------------|------------------|
| email  | Email address of the account | john@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
john@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column      | Description                            |
|-------------|----------------------------------------|
| status      | Status of the operation                |
| reason      | Reason of failure or skipped operation |
| input.email | Email address of the account           |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member batch invite
URL: https://toolbox.watermint.org/commands/dropbox/team/member/batch/invite.md
---

# dropbox team member batch invite

Send batch invitations to new team members, streamlining the onboarding process for multiple users (Irreversible operation)

Sends team invitations to multiple email addresses from a CSV file. Supports silent invites for SSO environments. Ideal for onboarding new departments, acquisitions, or seasonal workers. Validates email formats and checks for existing members before sending.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member batch invite -file /PATH/TO/member_list.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

**-silent-invite**
: Do not send welcome email (requires SSO + domain verification instead). Default: false

# File formats

## Format: File

Data file for inviting team members.

| Column     | Description                  | Example          |
|------------|------------------------------|------------------|
| email      | Email address of the account | john@example.com |
| given_name | Given name of the account    | John             |
| surname    | Surname of the account       | Smith            |

The first line is a header line. The program will accept a file without the header.
```
email,given_name,surname
john@example.com,John,Smith
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                | Description                                                                                    |
|-----------------------|------------------------------------------------------------------------------------------------|
| status                | Status of the operation                                                                        |
| reason                | Reason of failure or skipped operation                                                         |
| input.email           | Email address of the account                                                                   |
| input.given_name      | Given name of the account                                                                      |
| input.surname         | Surname of the account                                                                         |
| result.email          | Email address of user.                                                                         |
| result.email_verified | Is true if the user's email is verified to be owned by the user.                               |
| result.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| result.given_name     | Also known as a first name                                                                     |
| result.surname        | Also known as a last name or family name.                                                      |
| result.display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| result.joined_on      | The date and time the user joined as a member of a specific team.                              |
| result.invited_on     | The date and time the user was invited to the team                                             |
| result.role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| result.tag            | Operation tag                                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member batch reinvite
URL: https://toolbox.watermint.org/commands/dropbox/team/member/batch/reinvite.md
---

# dropbox team member batch reinvite

Resend invitations to pending members who haven't joined yet, ensuring all intended members receive access (Irreversible operation)

Resends invitations to all members with pending status. Useful when initial invites expire, get lost in spam, or after resolving email delivery issues. Can send silently for SSO environments. Helps ensure complete team onboarding.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member batch reinvite 
```

## Options:

**-peer**
: Account alias. Default: default

**-silent**
: Do not send welcome email (SSO required). Default: false

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                | Description                                                                                    |
|-----------------------|------------------------------------------------------------------------------------------------|
| status                | Status of the operation                                                                        |
| reason                | Reason of failure or skipped operation                                                         |
| input.email           | Email address of user.                                                                         |
| input.email_verified  | Is true if the user's email is verified to be owned by the user.                               |
| input.status          | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| input.given_name      | Also known as a first name                                                                     |
| input.surname         | Also known as a last name or family name.                                                      |
| input.display_name    | A name that can be used directly to represent the name of a user's Dropbox account.            |
| input.joined_on       | The date and time the user joined as a member of a specific team.                              |
| input.invited_on      | The date and time the user was invited to the team                                             |
| input.role            | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| input.tag             | Operation tag                                                                                  |
| result.email          | Email address of user.                                                                         |
| result.email_verified | Is true if the user's email is verified to be owned by the user.                               |
| result.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| result.given_name     | Also known as a first name                                                                     |
| result.surname        | Also known as a last name or family name.                                                      |
| result.display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| result.joined_on      | The date and time the user joined as a member of a specific team.                              |
| result.invited_on     | The date and time the user was invited to the team                                             |
| result.role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| result.tag            | Operation tag                                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member batch suspend
URL: https://toolbox.watermint.org/commands/dropbox/team/member/batch/suspend.md
---

# dropbox team member batch suspend

Temporarily suspend multiple team members' access while preserving their data and settings 

Bulk suspends team members, blocking access while preserving all data and settings. Use for extended leaves, security investigations, or temporary access restrictions. Option to keep or remove data from devices. Members can be unsuspended later with full access restored.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member batch suspend -file /PATH/TO/member_list.csv
```

## Options:

**-file**
: Path to data file

**-keep-data**
: Keep the user's data on their linked devices. Default: false

**-peer**
: Account alias. Default: default

# File formats

## Format: File

User selector data

| Column | Description            | Example          |
|--------|------------------------|------------------|
| email  | Member's email address | john@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
john@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column      | Description                            |
|-------------|----------------------------------------|
| status      | Status of the operation                |
| reason      | Reason of failure or skipped operation |
| input.email | Member's email address                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member batch unsuspend
URL: https://toolbox.watermint.org/commands/dropbox/team/member/batch/unsuspend.md
---

# dropbox team member batch unsuspend

Restore access for multiple suspended team members, reactivating their accounts in batch 

Bulk reactivates suspended team members, restoring full access to their accounts and data. Use when members return from leave, investigations conclude, or access restrictions lift. All previous permissions and group memberships are restored automatically.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member batch unsuspend -file /PATH/TO/member_list.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

User selector data

| Column | Description            | Example          |
|--------|------------------------|------------------|
| email  | Member's email address | john@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
john@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column      | Description                            |
|-------------|----------------------------------------|
| status      | Status of the operation                |
| reason      | Reason of failure or skipped operation |
| input.email | Member's email address                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member clear externalid
URL: https://toolbox.watermint.org/commands/dropbox/team/member/clear/externalid.md
---

# dropbox team member clear externalid

Remove external ID mappings from team members, useful when disconnecting from identity management systems 

Bulk removes external IDs from team members listed in a CSV file. Essential when migrating between identity providers, cleaning up after SCIM disconnection, or resolving ID conflicts. Does not affect member access, only removes the external identifier mapping.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member clear externalid -file /PATH/TO/member_list.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Email addresses of team members

| Column | Description                 | Example          |
|--------|-----------------------------|------------------|
| email  | Email address of the member | john@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
john@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                                          |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                              |
| reason                  | Reason of failure or skipped operation                                                                               |
| input.email             | Email address of the member                                                                                          |
| result.team_member_id   | ID of user as a member of a team.                                                                                    |
| result.email            | Email address of user.                                                                                               |
| result.email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| result.status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| result.given_name       | Also known as a first name                                                                                           |
| result.surname          | Also known as a last name or family name.                                                                            |
| result.familiar_name    | Locale-dependent name                                                                                                |
| result.display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| result.abbreviated_name | An abbreviated form of the person's name.                                                                            |
| result.member_folder_id | The namespace id of the user's root folder.                                                                          |
| result.external_id      | External ID that a team can attach to the user.                                                                      |
| result.account_id       | A user's account identifier.                                                                                         |
| result.persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| result.joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| result.invited_on       | The date and time the user was invited to the team                                                                   |
| result.role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| result.tag              | Operation tag                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member file permdelete
URL: https://toolbox.watermint.org/commands/dropbox/team/member/file/permdelete.md
---

# dropbox team member file permdelete

Permanently delete files or folders from a team member's account, bypassing trash for immediate removal (Experimental, and Irreversible operation)

Permanently deletes specified files or folders without possibility of recovery. Use with extreme caution for removing sensitive data, complying with data retention policies, or freeing storage. Cannot be undone - ensure proper authorization before use.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member file permdelete -member-email EMAIL -path /DROPBOX/PATH/TO/DELETE
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Team member email address

**-path**
: Path to delete

**-peer**
: Account alias. Default: default

---
Title: dropbox team member file lock list
URL: https://toolbox.watermint.org/commands/dropbox/team/member/file/lock/list.md
---

# dropbox team member file lock list

Display all files locked by a specific team member under a given path, identifying potential collaboration blocks 

Lists all files currently locked by a specific member within a path. Helps identify collaboration bottlenecks, troubleshoot editing conflicts, and audit file access patterns. Useful for understanding why team members cannot edit certain files.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member file lock list -member-email EMAIL -path /DROPBOX/PATH/TO/LIST_LOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Member email address

**-path**
: Path

**-peer**
: Account alias. Default: default

# Results

## Report: lock

Lock information
The command will generate a report in three different formats. `lock.csv`, `lock.json`, and `lock.xlsx`.

| Column           | Description                                                                                            |
|------------------|--------------------------------------------------------------------------------------------------------|
| tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name             | The last component of the path (including extension).                                                  |
| path_display     | The cased path to be used for display purposes only.                                                   |
| client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified  | The last time the file was modified on Dropbox.                                                        |
| size             | The file size in bytes.                                                                                |
| is_lock_holder   | True if caller holds the file lock                                                                     |
| lock_holder_name | The display name of the lock holder.                                                                   |
| lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `lock_0000.xlsx`, `lock_0001.xlsx`, `lock_0002.xlsx`, ...

---
Title: dropbox team member file lock release
URL: https://toolbox.watermint.org/commands/dropbox/team/member/file/lock/release.md
---

# dropbox team member file lock release

Release a specific file lock held by a team member, enabling others to edit the file 

Releases a single file lock held by a member, allowing others to edit. Use when specific files are blocking team collaboration or when lock holders are unavailable. More precise than bulk release when only specific files need unlocking.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member file lock release -member-email VALUE -path /DROPBOX/PATH/TO/RELEASE/LOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Member email address

**-path**
: Path to release lock

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member file lock all release
URL: https://toolbox.watermint.org/commands/dropbox/team/member/file/lock/all/release.md
---

# dropbox team member file lock all release

Release all file locks held by a team member under a specified path, resolving editing conflicts 

Bulk releases all file locks held by a member within a specified folder path. Essential when members leave unexpectedly or during system issues. Processes in batches for efficiency. Consider notifying affected users as their unsaved changes in locked files may be lost.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member file lock all release -member-email VALUE -path /DROPBOX/PATH/TO/RELEASE/LOCK
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-batch-size**
: Batch operation size. Default: 100

**-member-email**
: Member email address

**-path**
: Path to release lock

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member folder list
URL: https://toolbox.watermint.org/commands/dropbox/team/member/folder/list.md
---

# dropbox team member folder list

Display all folders in each team member's account, useful for content auditing and storage analysis 

Enumerates folders across team members' personal spaces. Filter by folder name to focus results. Essential for understanding content distribution, auditing member storage, and planning migrations or cleanups.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member folder list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-member-email**
: Filter by member email address. Filter by email address.

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: member_to_folder

Member to folder mapping.
The command will generate a report in three different formats. `member_to_folder.csv`, `member_to_folder.json`, and `member_to_folder.xlsx`.

| Column          | Description                                                                                              |
|-----------------|----------------------------------------------------------------------------------------------------------|
| member_name     | Team member display name.                                                                                |
| member_email    | Email address of the member                                                                              |
| access_type     | User's access level for this folder                                                                      |
| namespace_name  | The name of this namespace                                                                               |
| path            | Path                                                                                                     |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder) |
| owner_team_name | Team name of the team that owns the folder                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_to_folder_0000.xlsx`, `member_to_folder_0001.xlsx`, `member_to_folder_0002.xlsx`, ...

## Report: member_with_no_folder

This report shows a list of members.
The command will generate a report in three different formats. `member_with_no_folder.csv`, `member_with_no_folder.json`, and `member_with_no_folder.xlsx`.

| Column       | Description                                                                          |
|--------------|--------------------------------------------------------------------------------------|
| email        | Email address of user.                                                               |
| status       | The user's status as a member of a specific team. (active/invited/suspended/removed) |
| given_name   | Also known as a first name                                                           |
| surname      | Also known as a last name or family name.                                            |
| display_name | A name that can be used directly to represent the name of a user's Dropbox account.  |
| invited_on   | The date and time the user was invited to the team                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_with_no_folder_0000.xlsx`, `member_with_no_folder_0001.xlsx`, `member_with_no_folder_0002.xlsx`, ...

---
Title: dropbox team member folder replication
URL: https://toolbox.watermint.org/commands/dropbox/team/member/folder/replication.md
---

# dropbox team member folder replication

Copy folder contents from one team member to another's personal space, facilitating content transfer and backup (Irreversible operation)

Copies complete folder hierarchies between members' personal spaces, preserving structure. Ideal for creating backups, transitioning responsibilities, or setting up new members with standard folder structures. Monitor available storage before large replications.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member folder replication -dst-path /DROPBOX/PATH/OF/DST -src-path /DROPBOX/PATH/OF/SRC -dst-member-email DST_MEMBER@email.address -src-member-email SRC_MEMBER@email.address
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dst-member-email**
: Destination team member email address

**-dst-path**
: The path for the destination team member. Note the root (/) path is not allowed. You should choose any folder under the root.

**-peer**
: Account alias. Default: default

**-src-member-email**
: Source team member email address

**-src-path**
: The path of the source team member

---
Title: dropbox team member quota list
URL: https://toolbox.watermint.org/commands/dropbox/team/member/quota/list.md
---

# dropbox team member quota list

Display storage quota assignments for all team members, helping monitor and plan storage distribution 

Shows current storage quota settings for all team members, distinguishing between default and custom quotas. Identifies members with special storage needs or restrictions. Use for capacity planning and ensuring fair storage distribution across teams.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member quota list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: member_quota

This report shows a list of custom quota settings for each team member.
The command will generate a report in three different formats. `member_quota.csv`, `member_quota.json`, and `member_quota.xlsx`.

| Column | Description                                                                 |
|--------|-----------------------------------------------------------------------------|
| email  | Email address of user.                                                      |
| quota  | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_quota_0000.xlsx`, `member_quota_0001.xlsx`, `member_quota_0002.xlsx`, ...

---
Title: dropbox team member quota usage
URL: https://toolbox.watermint.org/commands/dropbox/team/member/quota/usage.md
---

# dropbox team member quota usage

Show actual storage usage for each team member compared to their quotas, identifying storage needs 

Displays current storage consumption versus allocated quotas for each member. Highlights members approaching limits, underutilizing space, or needing quota adjustments. Critical for proactive storage management and preventing work disruptions due to full quotas.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member quota usage 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

# Results

## Report: usage

This report shows current storage usage of users.
The command will generate a report in three different formats. `usage.csv`, `usage.json`, and `usage.xlsx`.

| Column     | Description                                              |
|------------|----------------------------------------------------------|
| email      | Email address of the account                             |
| used_gb    | The user's total space usage (in GB, 1GB = 1024 MB).     |
| used_bytes | The user's total space usage (bytes).                    |
| allocation | The user's space allocation (individual, or team)        |
| allocated  | The total space allocated to the user's account (bytes). |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `usage_0000.xlsx`, `usage_0001.xlsx`, `usage_0002.xlsx`, ...

---
Title: dropbox team member quota batch update
URL: https://toolbox.watermint.org/commands/dropbox/team/member/quota/batch/update.md
---

# dropbox team member quota batch update

Modify storage quotas for multiple team members in batch, managing storage allocation efficiently 

Bulk updates storage quotas for team members using a CSV file. Set custom quotas based on roles, departments, or usage patterns. Use 0 to remove custom quotas and revert to team defaults. Essential for storage governance and cost management.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member quota batch update -file /path/to/file.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

**-quota**
: Custom quota in GB (1TB = 1024GB). 0 if the user has no custom quota set.. Default: 0

# File formats

## Format: File

This report shows a list of custom quota settings for each team member.

| Column | Description                                                                 | Example          |
|--------|-----------------------------------------------------------------------------|------------------|
| email  | Email address of user.                                                      | john@example.com |
| quota  | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. | 50               |

The first line is a header line. The program will accept a file without the header.
```
email,quota
john@example.com,50
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column       | Description                                                                 |
|--------------|-----------------------------------------------------------------------------|
| status       | Status of the operation                                                     |
| reason       | Reason of failure or skipped operation                                      |
| input.email  | Email address of user.                                                      |
| input.quota  | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. |
| result.email | Email address of user.                                                      |
| result.quota | Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member update batch email
URL: https://toolbox.watermint.org/commands/dropbox/team/member/update/batch/email.md
---

# dropbox team member update batch email

Update email addresses for multiple team members in batch, managing email changes efficiently (Irreversible operation)

Bulk updates member email addresses using a CSV mapping file. Essential for domain migrations, name changes, or correcting email errors. Validates new addresses and preserves all member data and permissions. Option to update unverified emails with caution.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member update batch email -file /path/to/data/file.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

**-update-unverified**
: Update an account which hasn't verified its email. If an account email is unverified, changing the email address may cause loss of invitation to folders.. Default: false

# File formats

## Format: File

Data file for updating team member email addresses.

| Column     | Description           | Example                |
|------------|-----------------------|------------------------|
| from_email | Current Email address | john@example.com       |
| to_email   | New Email address     | john.smith@example.net |

The first line is a header line. The program will accept a file without the header.
```
from_email,to_email
john@example.com,john.smith@example.net
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                | Description                                                                                    |
|-----------------------|------------------------------------------------------------------------------------------------|
| status                | Status of the operation                                                                        |
| reason                | Reason of failure or skipped operation                                                         |
| input.from_email      | Current Email address                                                                          |
| input.to_email        | New Email address                                                                              |
| result.email          | Email address of user.                                                                         |
| result.email_verified | Is true if the user's email is verified to be owned by the user.                               |
| result.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| result.given_name     | Also known as a first name                                                                     |
| result.surname        | Also known as a last name or family name.                                                      |
| result.display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| result.joined_on      | The date and time the user joined as a member of a specific team.                              |
| result.invited_on     | The date and time the user was invited to the team                                             |
| result.role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| result.tag            | Operation tag                                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member update batch externalid
URL: https://toolbox.watermint.org/commands/dropbox/team/member/update/batch/externalid.md
---

# dropbox team member update batch externalid

Set or update external IDs for multiple team members, integrating with identity management systems (Irreversible operation)

Maps external identity system IDs to Dropbox team members in bulk. Critical for SCIM integration, SSO setup, or syncing with HR systems. Ensures consistent identity mapping across platforms. Updates existing IDs or sets new ones as needed.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member update batch externalid -file /path/to/file.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for updating member external id.

| Column      | Description                   | Example          |
|-------------|-------------------------------|------------------|
| email       | Email address of team members | john@example.com |
| external_id | External ID of team members   | 0123456789       |

The first line is a header line. The program will accept a file without the header.
```
email,external_id
john@example.com,0123456789
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                | Description                                                                                    |
|-----------------------|------------------------------------------------------------------------------------------------|
| status                | Status of the operation                                                                        |
| reason                | Reason of failure or skipped operation                                                         |
| input.email           | Email address of team members                                                                  |
| input.external_id     | External ID of team members                                                                    |
| result.email          | Email address of user.                                                                         |
| result.email_verified | Is true if the user's email is verified to be owned by the user.                               |
| result.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| result.given_name     | Also known as a first name                                                                     |
| result.surname        | Also known as a last name or family name.                                                      |
| result.display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| result.joined_on      | The date and time the user joined as a member of a specific team.                              |
| result.invited_on     | The date and time the user was invited to the team                                             |
| result.role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| result.tag            | Operation tag                                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member update batch invisible
URL: https://toolbox.watermint.org/commands/dropbox/team/member/update/batch/invisible.md
---

# dropbox team member update batch invisible

Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors (Irreversible operation)

Bulk hides members from team directory searches and listings. Useful for executives, security personnel, or external contractors who need access but shouldn't appear in directories. Hidden members retain all functionality but enhanced privacy.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member update batch invisible -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Member list for changing visibility

| Column | Description          | Example          |
|--------|----------------------|------------------|
| email  | Member email address | taro@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
taro@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                                          |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                              |
| reason                  | Reason of failure or skipped operation                                                                               |
| input.Email             | Member email address                                                                                                 |
| result.team_member_id   | ID of user as a member of a team.                                                                                    |
| result.email            | Email address of user.                                                                                               |
| result.email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| result.status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| result.given_name       | Also known as a first name                                                                                           |
| result.surname          | Also known as a last name or family name.                                                                            |
| result.familiar_name    | Locale-dependent name                                                                                                |
| result.display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| result.abbreviated_name | An abbreviated form of the person's name.                                                                            |
| result.member_folder_id | The namespace id of the user's root folder.                                                                          |
| result.external_id      | External ID that a team can attach to the user.                                                                      |
| result.account_id       | A user's account identifier.                                                                                         |
| result.persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| result.joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| result.invited_on       | The date and time the user was invited to the team                                                                   |
| result.role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| result.tag              | Operation tag                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member update batch profile
URL: https://toolbox.watermint.org/commands/dropbox/team/member/update/batch/profile.md
---

# dropbox team member update batch profile

Update profile information for multiple team members including names and job titles in batch (Irreversible operation)

Bulk updates member profile information including given names and surnames. Ideal for standardizing name formats, correcting widespread errors, or updating after organizational changes. Maintains consistency across team directories and improves searchability.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member update batch profile -file /path/to/data/file.csv
```

## Options:

**-file**
: Data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for batch profile updates.

| Column     | Description                  | Example          |
|------------|------------------------------|------------------|
| email      | Email address of the account | john@example.com |
| given_name | Given name of the account    | John             |
| surname    | Surname of the account       | Smith            |

The first line is a header line. The program will accept a file without the header.
```
email,given_name,surname
john@example.com,John,Smith
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                | Description                                                                                    |
|-----------------------|------------------------------------------------------------------------------------------------|
| status                | Status of the operation                                                                        |
| reason                | Reason of failure or skipped operation                                                         |
| input.email           | Email address of the account                                                                   |
| input.given_name      | Given name of the account                                                                      |
| input.surname         | Surname of the account                                                                         |
| result.email          | Email address of user.                                                                         |
| result.email_verified | Is true if the user's email is verified to be owned by the user.                               |
| result.status         | The user's status as a member of a specific team. (active/invited/suspended/removed)           |
| result.given_name     | Also known as a first name                                                                     |
| result.surname        | Also known as a last name or family name.                                                      |
| result.display_name   | A name that can be used directly to represent the name of a user's Dropbox account.            |
| result.joined_on      | The date and time the user joined as a member of a specific team.                              |
| result.invited_on     | The date and time the user was invited to the team                                             |
| result.role           | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only) |
| result.tag            | Operation tag                                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team member update batch visible
URL: https://toolbox.watermint.org/commands/dropbox/team/member/update/batch/visible.md
---

# dropbox team member update batch visible

Make hidden team members visible in the directory, restoring standard visibility settings (Irreversible operation)

Bulk restores visibility for previously hidden members in team directories. Use when privacy requirements change, contractors become employees, or to correct visibility errors. Members become searchable and appear in team listings again.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team member update batch visible -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Member list for changing visibility

| Column | Description          | Example          |
|--------|----------------------|------------------|
| email  | Member email address | taro@example.com |

The first line is a header line. The program will accept a file without the header.
```
email
taro@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                                          |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                              |
| reason                  | Reason of failure or skipped operation                                                                               |
| input.Email             | Member email address                                                                                                 |
| result.team_member_id   | ID of user as a member of a team.                                                                                    |
| result.email            | Email address of user.                                                                                               |
| result.email_verified   | Is true if the user's email is verified to be owned by the user.                                                     |
| result.status           | The user's status as a member of a specific team. (active/invited/suspended/removed)                                 |
| result.given_name       | Also known as a first name                                                                                           |
| result.surname          | Also known as a last name or family name.                                                                            |
| result.familiar_name    | Locale-dependent name                                                                                                |
| result.display_name     | A name that can be used directly to represent the name of a user's Dropbox account.                                  |
| result.abbreviated_name | An abbreviated form of the person's name.                                                                            |
| result.member_folder_id | The namespace id of the user's root folder.                                                                          |
| result.external_id      | External ID that a team can attach to the user.                                                                      |
| result.account_id       | A user's account identifier.                                                                                         |
| result.persistent_id    | Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication. |
| result.joined_on        | The date and time the user joined as a member of a specific team.                                                    |
| result.invited_on       | The date and time the user was invited to the team                                                                   |
| result.role             | The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)                       |
| result.tag              | Operation tag                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team namespace list
URL: https://toolbox.watermint.org/commands/dropbox/team/namespace/list.md
---

# dropbox team namespace list

Display all team namespaces including team folders and shared spaces with their configurations 

Enumerates all namespace types in the team including ownership, paths, and access levels. Provides comprehensive view of team's folder architecture. Use for understanding organizational structure, planning migrations, or auditing folder governance.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team namespace list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: namespace

This report shows a list of namespaces in the team.
The command will generate a report in three different formats. `namespace.csv`, `namespace.json`, and `namespace.xlsx`.

| Column         | Description                                                                                |
|----------------|--------------------------------------------------------------------------------------------|
| name           | The name of this namespace                                                                 |
| namespace_type | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_0000.xlsx`, `namespace_0001.xlsx`, `namespace_0002.xlsx`, ...

---
Title: dropbox team namespace summary
URL: https://toolbox.watermint.org/commands/dropbox/team/namespace/summary.md
---

# dropbox team namespace summary

Generate comprehensive summary reports of team namespace usage, member counts, and storage statistics 

Aggregates namespace data to show overall team structure, storage distribution, and access patterns. Provides high-level insights into how team content is organized across different namespace types. Useful for capacity planning and organizational assessments.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team namespace summary 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

**-skip-member-summary**
: Skip scanning member namespaces. Default: false

# Results

## Report: folder_without_parent

Folders without parent folder.
The command will generate a report in three different formats. `folder_without_parent.csv`, `folder_without_parent.json`, and `folder_without_parent.xlsx`.

| Column                  | Description                                                                                                             |
|-------------------------|-------------------------------------------------------------------------------------------------------------------------|
| shared_folder_id        | The ID of the shared folder.                                                                                            |
| parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| name                    | The name of this shared folder.                                                                                         |
| access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| path_lower              | The lower-cased full path of this shared folder.                                                                        |
| is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| is_team_folder          | Whether this folder is a team folder.                                                                                   |
| policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| policy_shared_link      | Who links can be shared with.                                                                                           |
| policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| owner_team_id           | Team ID of the folder owner team                                                                                        |
| owner_team_name         | Team name of the team that owns the folder                                                                              |
| access_inheritance      | Access inheritance type                                                                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `folder_without_parent_0000.xlsx`, `folder_without_parent_0001.xlsx`, `folder_without_parent_0002.xlsx`, ...

## Report: member

Member namespace summary
The command will generate a report in three different formats. `member.csv`, `member.json`, and `member.xlsx`.

| Column              | Description                                                  |
|---------------------|--------------------------------------------------------------|
| email               | Member email address                                         |
| total_namespaces    | Number of total namespaces (excluding member root namespace) |
| mounted_namespaces  | Number of mounted folders                                    |
| owner_namespaces    | Number of shared folders owned by this member                |
| team_folders        | Number of team folders                                       |
| inside_team_folders | Number of folders inside team folders                        |
| external_folders    | Number of folders shared by a user outside the team          |
| app_folders         | Number of app folders                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `member_0000.xlsx`, `member_0001.xlsx`, `member_0002.xlsx`, ...

## Report: team

Team namespace summary.
The command will generate a report in three different formats. `team.csv`, `team.json`, and `team.xlsx`.

| Column          | Description          |
|-----------------|----------------------|
| namespace_type  | Type of namespace    |
| namespace_count | Number of namespaces |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `team_0000.xlsx`, `team_0001.xlsx`, `team_0002.xlsx`, ...

## Report: team_folder

Team folder summary.
The command will generate a report in three different formats. `team_folder.csv`, `team_folder.json`, and `team_folder.xlsx`.

| Column                | Description                                  |
|-----------------------|----------------------------------------------|
| name                  | Team folder name                             |
| num_namespaces_inside | Number of namespaces inside this team folder |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`, ...

---
Title: dropbox team namespace file list
URL: https://toolbox.watermint.org/commands/dropbox/team/namespace/file/list.md
---

# dropbox team namespace file list

Display comprehensive file and folder listings within team namespaces for content inventory and analysis 

Lists all files and folders within team namespaces with filtering options. Include or exclude deleted items, member folders, shared folders, and team folders. Essential for content audits, migration planning, and understanding data distribution across namespace types.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team namespace file list 
```

## Options:

**-folder-name**
: List only for the folder matched to the name. Filter by exact match to the name.

**-folder-name-prefix**
: List only for the folder matched to the name. Filter by name match to the prefix.

**-folder-name-suffix**
: List only for the folder matched to the name. Filter by name match to the suffix.

**-include-deleted**
: If true, deleted file or folder will be returned. Default: false

**-include-member-folder**
: If true, include team member folders. Default: false

**-include-shared-folder**
: If true, include shared folders. Default: true

**-include-team-folder**
: If true, include team folders. Default: true

**-peer**
: Account alias. Default: default

# Results

## Report: errors

This report shows the transaction result.
The command will generate a report in three different formats. `errors.csv`, `errors.json`, and `errors.xlsx`.

| Column          | Description                            |
|-----------------|----------------------------------------|
| status          | Status of the operation                |
| reason          | Reason of failure or skipped operation |
| input.namespace | Namespace                              |
| input.path      | Path                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...

## Report: namespace_file

This report shows a list of namespaces in the team.
The command will generate a report in three different formats. `namespace_file.csv`, `namespace_file.json`, and `namespace_file.xlsx`.

| Column                 | Description                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------|
| namespace_type         | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)             |
| namespace_name         | The name of this namespace                                                                             |
| namespace_member_email | If this is a team member or app folder, the email address of the owning team member.                   |
| tag                    | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                   | The last component of the path (including extension).                                                  |
| path_display           | The cased path to be used for display purposes only.                                                   |
| client_modified        | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified        | The last time the file was modified on Dropbox.                                                        |
| size                   | The file size in bytes.                                                                                |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`, ...

---
Title: dropbox team namespace file size
URL: https://toolbox.watermint.org/commands/dropbox/team/namespace/file/size.md
---

# dropbox team namespace file size

Calculate storage usage for files and folders in team namespaces, providing detailed size analytics 

Analyzes storage consumption across team namespaces with configurable depth scanning. Shows size distribution by namespace type (team, shared, member, app folders). Critical for storage optimization, identifying large folders, and planning archival strategies.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team namespace file size 
```

## Options:

**-depth**
: Report entry for all files and directories depth directories deep. Default: 3

**-folder-name**
: List only for the folder matched to the name. Filter by exact match to the name.

**-folder-name-prefix**
: List only for the folder matched to the name. Filter by name match to the prefix.

**-folder-name-suffix**
: List only for the folder matched to the name. Filter by name match to the suffix.

**-include-app-folder**
: If true, include app folders. Default: false

**-include-member-folder**
: If true, include team member folders. Default: false

**-include-shared-folder**
: If true, include shared folders. Default: true

**-include-team-folder**
: If true, include team folders. Default: true

**-peer**
: Account alias. Default: default

# Results

## Report: namespace_size

Namespace size in bytes
The command will generate a report in three different formats. `namespace_size.csv`, `namespace_size.json`, and `namespace_size.xlsx`.

| Column               | Description                                                                                |
|----------------------|--------------------------------------------------------------------------------------------|
| namespace_name       | The name of this namespace                                                                 |
| namespace_id         | The ID of this namespace.                                                                  |
| namespace_type       | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| owner_team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |
| path                 | Path to the folder                                                                         |
| count_file           | Number of files under the folder                                                           |
| count_folder         | Number of folders under the folder                                                         |
| count_descendant     | Number of files and folders under the folder                                               |
| size                 | Size of the folder                                                                         |
| depth                | Namespace depth                                                                            |
| mod_time_earliest    | Earliest modification time in namespace                                                    |
| mod_time_latest      | Latest modification time in namespace                                                      |
| api_complexity       | Folder complexity index for API operations                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`, ...

---
Title: dropbox team namespace member list
URL: https://toolbox.watermint.org/commands/dropbox/team/namespace/member/list.md
---

# dropbox team namespace member list

Show all members with access to each namespace, detailing permissions and access levels 

Maps namespace access showing which members can access which folders and their permission levels. Reveals access patterns, over-privileged namespaces, and helps ensure appropriate access controls. Essential for security audits and access reviews.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team namespace member list 
```

## Options:

**-all-columns**
: Show all columns. Default: false

**-peer**
: Account alias. Default: default

# Results

## Report: namespace_member

This report shows a list of members of namespaces in the team.
The command will generate a report in three different formats. `namespace_member.csv`, `namespace_member.json`, and `namespace_member.xlsx`.

| Column             | Description                                                                                               |
|--------------------|-----------------------------------------------------------------------------------------------------------|
| namespace_name     | The name of this namespace                                                                                |
| namespace_type     | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)                |
| entry_access_type  | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| entry_is_inherited | True if the member has access from a parent folder                                                        |
| email              | Email address of user.                                                                                    |
| display_name       | Team member display name.                                                                                 |
| group_name         | Name of the group                                                                                         |
| invitee_email      | Email address of invitee for this folder                                                                  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_member_0000.xlsx`, `namespace_member_0001.xlsx`, `namespace_member_0002.xlsx`, ...

---
Title: dropbox team runas file list
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/file/list.md
---

# dropbox team runas file list

List files and folders in a team member's account by running operations as that member 

Allows admins to view file listings in member accounts without member credentials. Essential for investigating issues, auditing content, or helping members locate files. All actions are logged for security. Use responsibly and follow privacy policies.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas file list -member-email MEMBER@DOMAIN -path /DROPBOX/PATH/TO/LIST
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-include-deleted**
: Include deleted files. Default: false

**-include-explicit-shared-members**
: If true, the results will include a flag for each file indicating whether or not that file has any explicit members.. Default: false

**-include-mounted-folders**
: If true, the results will include entries under mounted folders which include app folder, shared folder and team folder.. Default: false

**-member-email**
: Email address of the member

**-path**
: Path

**-peer**
: Account alias. Default: default

**-recursive**
: List recursively. Default: false

# Results

## Report: file_list

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `file_list.csv`, `file_list.json`, and `file_list.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| size                        | The file size in bytes.                                                                                              |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `file_list_0000.xlsx`, `file_list_0001.xlsx`, `file_list_0002.xlsx`, ...

---
Title: dropbox team runas file batch copy
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/file/batch/copy.md
---

# dropbox team runas file batch copy

Copy multiple files or folders on behalf of team members, useful for content management and organization (Irreversible operation)

Admin tool to copy files between member accounts without their credentials. Useful for distributing templates, recovering deleted content, or setting up new members. Operates with admin privileges while maintaining audit trails. Requires appropriate admin permissions.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas file batch copy -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Mapping between source and destination paths

| Column       | Description               | Example          |
|--------------|---------------------------|------------------|
| member_email | Team member email address | emma@example.com |
| src_path     | Source path               | /report          |
| dst_path     | Destination path          | /backup/report   |

The first line is a header line. The program will accept a file without the header.
```
member_email,src_path,dst_path
emma@example.com,/report,/backup/report
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column             | Description                            |
|--------------------|----------------------------------------|
| status             | Status of the operation                |
| reason             | Reason of failure or skipped operation |
| input.member_email | Team member email address              |
| input.src_path     | Source path                            |
| input.dst_path     | Destination path                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team runas file sync batch up
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/file/sync/batch/up.md
---

# dropbox team runas file sync batch up

Upload multiple local files to team members' Dropbox accounts in batch, running as those members (Irreversible operation)

Admin bulk upload tool for distributing files to multiple member accounts simultaneously. Ideal for deploying templates, policies, or required documents. Maintains consistent file distribution across teams. All uploads are tracked for compliance.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas file sync batch up -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-batch-size**
: Batch commit size. Default: 250

**-delete**
: Delete Dropbox file if a file is removed locally. Default: false

**-exit-on-failure**
: Exit the program on failure. Default: false

**-file**
: Path to data file

**-name-disable-ignore**
: Name for the sync batch operation. Filter system file or ignore files.

**-name-name**
: Name for the sync batch operation. Filter by exact match to the name.

**-name-name-prefix**
: Name for the sync batch operation. Filter by name match to the prefix.

**-name-name-suffix**
: Name for the sync batch operation. Filter by name match to the suffix.

**-overwrite**
: Overwrite existing files if they exist.. Default: false

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Mapping of local files to Dropbox paths for batch upload.

| Column       | Description                               | Example                         |
|--------------|-------------------------------------------|---------------------------------|
| member_email | Email address of the Dropbox team member. | user@example.com                |
| local_path   | Local file path to upload.                | /Users/alice/Documents/file.txt |
| dropbox_path | Destination path in Dropbox.              | /Team Folder/Project/file.txt   |

The first line is a header line. The program will accept a file without the header.
```
member_email,local_path,dropbox_path
user@example.com,/Users/alice/Documents/file.txt,/Team Folder/Project/file.txt
```

# Results

## Report: deleted

Path
The command will generate a report in three different formats. `deleted.csv`, `deleted.json`, and `deleted.xlsx`.

| Column                       | Description      |
|------------------------------|------------------|
| entry_path                   | Path             |
| entry_shard.file_system_type | File system type |
| entry_shard.shard_id         | Shard ID         |
| entry_shard.attributes       | Shard attributes |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column             | Description                               |
|--------------------|-------------------------------------------|
| status             | Status of the operation                   |
| reason             | Reason of failure or skipped operation    |
| input.member_email | Email address of the Dropbox team member. |
| input.local_path   | Local file path to upload.                |
| input.dropbox_path | Destination path in Dropbox.              |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

## Report: skipped

This report shows the transaction result.
The command will generate a report in three different formats. `skipped.csv`, `skipped.json`, and `skipped.xlsx`.

| Column                             | Description                            |
|------------------------------------|----------------------------------------|
| status                             | Status of the operation                |
| reason                             | Reason of failure or skipped operation |
| input.entry_path                   | Path                                   |
| input.entry_shard.file_system_type | File system type                       |
| input.entry_shard.shard_id         | Shard ID                               |
| input.entry_shard.attributes       | Shard attributes                       |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...

## Report: summary

This report shows a summary of the upload results.
The command will generate a report in three different formats. `summary.csv`, `summary.json`, and `summary.xlsx`.

| Column                | Description                                   |
|-----------------------|-----------------------------------------------|
| start                 | Time of start                                 |
| end                   | Time of finish                                |
| num_bytes             | Total upload size (Bytes)                     |
| num_files_error       | The number of files failed or got an error.   |
| num_files_transferred | The number of files uploaded/downloaded.      |
| num_files_skip        | The number of files skipped or to skip.       |
| num_folder_created    | Number of created folders.                    |
| num_delete            | Number of deleted entries.                    |
| num_api_call          | The number of estimated API calls for upload. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...

## Report: uploaded

This report shows the transaction result.
The command will generate a report in three different formats. `uploaded.csv`, `uploaded.json`, and `uploaded.xlsx`.

| Column                             | Description                                                                                                          |
|------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| status                             | Status of the operation                                                                                              |
| reason                             | Reason of failure or skipped operation                                                                               |
| input.path                         | Path                                                                                                                 |
| result.name                        | The last component of the path (including extension).                                                                |
| result.path_display                | The cased path to be used for display purposes only.                                                                 |
| result.client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| result.server_modified             | The last time the file was modified on Dropbox.                                                                      |
| result.size                        | The file size in bytes.                                                                                              |
| result.content_hash                | A hash of the file content.                                                                                          |
| result.has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder isolate
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/isolate.md
---

# dropbox team runas sharedfolder isolate

Remove all shared folder access for a team member and transfer ownership, useful for departing employees (Irreversible operation)

Emergency admin action to remove all members from a shared folder except its owner. Use for security incidents, data breaches, or when folder content needs immediate access restriction. Preserves folder structure while eliminating external access risks.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder isolate -member-email EMAIL
```

## Options:

**-base-path**
: Base path of the shared folder to isolate.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-keep-copy**
: Keep a copy after isolation.. Default: false

**-member-email**
: Email address of the member to isolate.

**-peer**
: Account alias. Default: default

# Results

## Report: isolated

This report shows the transaction result.
The command will generate a report in three different formats. `isolated.csv`, `isolated.json`, and `isolated.xlsx`.

| Column                      | Description                                                                                               |
|-----------------------------|-----------------------------------------------------------------------------------------------------------|
| status                      | Status of the operation                                                                                   |
| reason                      | Reason of failure or skipped operation                                                                    |
| input.shared_folder_id      | The ID of the shared folder.                                                                              |
| input.name                  | The name of this shared folder.                                                                           |
| input.access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| input.path_lower            | The lower-cased full path of this shared folder.                                                          |
| input.is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| input.is_team_folder        | Whether this folder is a team folder.                                                                     |
| input.policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| input.policy_shared_link    | Who links can be shared with.                                                                             |
| input.policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| input.policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| input.policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| input.owner_team_name       | Team name of the team that owns the folder                                                                |
| input.access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `isolated_0000.xlsx`, `isolated_0001.xlsx`, `isolated_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder list
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/list.md
---

# dropbox team runas sharedfolder list

Display all shared folders accessible by a team member, running the operation as that member 

Admin view of member's shared folder access including permission levels and folder details. Essential for access audits, investigating over-sharing, or troubleshooting permission issues. Helps ensure appropriate access levels and identify security risks.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder list -member-email EMAIL
```

## Options:

**-base-path**
: Base path of the shared folder to list.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Email address of the member to list.

**-peer**
: Account alias. Default: default

# Results

## Report: shared_folder

This report shows a list of shared folders.
The command will generate a report in three different formats. `shared_folder.csv`, `shared_folder.json`, and `shared_folder.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_folder_0000.xlsx`, `shared_folder_0001.xlsx`, `shared_folder_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder batch leave
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/batch/leave.md
---

# dropbox team runas sharedfolder batch leave

Remove team members from multiple shared folders in batch by running leave operations as those members 

Admin tool to remove members from multiple shared folders without their interaction. Useful for access cleanup, security responses, or organizational changes. Operates as the member would, maintaining proper audit trails. Cannot remove folder owners.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder batch leave -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Base path of the shared folder to leave.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-keep-copy**
: Keep a copy of the folder after leaving.. Default: false

**-peer**
: Account alias. Default: default

# File formats

## Format: File

List of member folders for batch operations.

| Column       | Description                  | Example                      |
|--------------|------------------------------|------------------------------|
| member_email | Email address of the member. | member@example.com           |
| path         | Path to the member's folder. | /Team Folder/Shared/file.txt |

The first line is a header line. The program will accept a file without the header.
```
member_email,path
member@example.com,/Team Folder/Shared/file.txt
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                         | Description                                                                                                             |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| status                         | Status of the operation                                                                                                 |
| reason                         | Reason of failure or skipped operation                                                                                  |
| input.member_email             | Email address of the member.                                                                                            |
| input.path                     | Path to the member's folder.                                                                                            |
| result.shared_folder_id        | The ID of the shared folder.                                                                                            |
| result.parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| result.name                    | The name of this shared folder.                                                                                         |
| result.access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| result.path_lower              | The lower-cased full path of this shared folder.                                                                        |
| result.is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| result.is_team_folder          | Whether this folder is a team folder.                                                                                   |
| result.policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| result.policy_shared_link      | Who links can be shared with.                                                                                           |
| result.policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| result.policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| result.policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| result.owner_team_id           | Team ID of the folder owner team                                                                                        |
| result.owner_team_name         | Team name of the team that owns the folder                                                                              |
| result.access_inheritance      | Access inheritance type                                                                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder batch share
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/batch/share.md
---

# dropbox team runas sharedfolder batch share

Share multiple folders on behalf of team members in batch, automating folder sharing processes 

Admin batch tool for creating shared folders on behalf of members. Streamlines folder sharing for new projects or team reorganizations. Sets appropriate permissions and sends invitations. All sharing actions are logged for security compliance.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder batch share -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-acl-update-policy**
: Access control update policy.. Options: owner (aclupdatepolicy: owner), editor (aclupdatepolicy: editor). Default: owner

**-base-path**
: Base path of the shared folder to share.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-member-policy**
: Policy for shared folder members.. Options: team (memberpolicy: team), anyone (memberpolicy: anyone). Default: anyone

**-peer**
: Account alias. Default: default

**-shared-link-policy**
: Policy for shared links.. Options: anyone (sharedlinkpolicy: anyone), members (sharedlinkpolicy: members). Default: anyone

# File formats

## Format: File

List of member folders for batch operations.

| Column       | Description                  | Example                      |
|--------------|------------------------------|------------------------------|
| member_email | Email address of the member. | member@example.com           |
| path         | Path to the member's folder. | /Team Folder/Shared/file.txt |

The first line is a header line. The program will accept a file without the header.
```
member_email,path
member@example.com,/Team Folder/Shared/file.txt
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                         | Description                                                                                                             |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| status                         | Status of the operation                                                                                                 |
| reason                         | Reason of failure or skipped operation                                                                                  |
| input.member_email             | Email address of the member.                                                                                            |
| input.path                     | Path to the member's folder.                                                                                            |
| result.shared_folder_id        | The ID of the shared folder.                                                                                            |
| result.parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| result.name                    | The name of this shared folder.                                                                                         |
| result.access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| result.path_lower              | The lower-cased full path of this shared folder.                                                                        |
| result.is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| result.is_team_folder          | Whether this folder is a team folder.                                                                                   |
| result.policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| result.policy_shared_link      | Who links can be shared with.                                                                                           |
| result.policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| result.policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| result.policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| result.owner_team_id           | Team ID of the folder owner team                                                                                        |
| result.owner_team_name         | Team name of the team that owns the folder                                                                              |
| result.access_inheritance      | Access inheritance type                                                                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder batch unshare
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/batch/unshare.md
---

# dropbox team runas sharedfolder batch unshare

Remove sharing from multiple folders on behalf of team members, managing folder access in bulk 

Admin tool to revoke folder sharing in bulk for security or compliance. Removes sharing while preserving folder contents for the owner. Critical for incident response or preventing data leaks. All unshare actions create audit records.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder batch unshare -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Base path of the shared folder to unshare.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-leave-copy**
: Leave a copy after unsharing.. Default: false

**-peer**
: Account alias. Default: default

# File formats

## Format: File

List of member folders for batch operations.

| Column       | Description                  | Example                      |
|--------------|------------------------------|------------------------------|
| member_email | Email address of the member. | member@example.com           |
| path         | Path to the member's folder. | /Team Folder/Shared/file.txt |

The first line is a header line. The program will accept a file without the header.
```
member_email,path
member@example.com,/Team Folder/Shared/file.txt
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                         | Description                                                                                                             |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| status                         | Status of the operation                                                                                                 |
| reason                         | Reason of failure or skipped operation                                                                                  |
| input.member_email             | Email address of the member.                                                                                            |
| input.path                     | Path to the member's folder.                                                                                            |
| result.shared_folder_id        | The ID of the shared folder.                                                                                            |
| result.parent_shared_folder_id | The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder. |
| result.name                    | The name of this shared folder.                                                                                         |
| result.access_type             | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)               |
| result.path_lower              | The lower-cased full path of this shared folder.                                                                        |
| result.is_inside_team_folder   | Whether this folder is inside of a team folder.                                                                         |
| result.is_team_folder          | Whether this folder is a team folder.                                                                                   |
| result.policy_manage_access    | Who can add and remove members from this shared folder.                                                                 |
| result.policy_shared_link      | Who links can be shared with.                                                                                           |
| result.policy_member_folder    | Who can be a member of this shared folder, as set on the folder itself.                                                 |
| result.policy_member           | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                                |
| result.policy_viewer_info      | Who can enable/disable viewer info for this shared folder.                                                              |
| result.owner_team_id           | Team ID of the folder owner team                                                                                        |
| result.owner_team_name         | Team name of the team that owns the folder                                                                              |
| result.access_inheritance      | Access inheritance type                                                                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder member batch add
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/member/batch/add.md
---

# dropbox team runas sharedfolder member batch add

Add multiple members to shared folders in batch on behalf of folder owners, streamlining access management 

Admin tool to bulk add members to specific shared folders with defined permissions. Efficient for project kickoffs, team expansions, or access standardization. Validates member emails and permissions before applying changes. Creates comprehensive audit trail.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder member batch add -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Base path of the shared folder to add members.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-message**
: Message to send to new members.

**-peer**
: Account alias. Default: default

**-silent**
: Add members silently without notification.. Default: false

# File formats

## Format: File

Details of the member to add.

| Column         | Description                          | Example                      |
|----------------|--------------------------------------|------------------------------|
| member_email   | Email address of the member to add.  | member@example.com           |
| path           | Path to the shared folder.           | /Team Folder/Shared/file.txt |
| access_level   | Access level to grant to the member. | editor                       |
| group_or_email | Group name or email address to add.  | group@example.com            |

The first line is a header line. The program will accept a file without the header.
```
member_email,path,access_level,group_or_email
member@example.com,/Team Folder/Shared/file.txt,editor,group@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column               | Description                            |
|----------------------|----------------------------------------|
| status               | Status of the operation                |
| reason               | Reason of failure or skipped operation |
| input.member_email   | Email address of the member to add.    |
| input.path           | Path to the shared folder.             |
| input.access_level   | Access level to grant to the member.   |
| input.group_or_email | Group name or email address to add.    |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder member batch delete
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/member/batch/delete.md
---

# dropbox team runas sharedfolder member batch delete

Remove multiple members from shared folders in batch on behalf of folder owners, managing access efficiently 

Admin bulk removal of members from shared folders for security or reorganization. Preserves folder content while revoking access for specified members. Essential for quick security responses or access cleanup. Cannot remove folder owner.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder member batch delete -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Base path of the shared folder to remove members.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-leave-copy**
: Leave a copy after removing member.. Default: false

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Details of the member to remove.

| Column         | Description                            | Example                      |
|----------------|----------------------------------------|------------------------------|
| member_email   | Email address of the member to remove. | member@example.com           |
| path           | Path to the shared folder.             | /Team Folder/Shared/file.txt |
| group_or_email | Group name or email address to remove. | group@example.com            |

The first line is a header line. The program will accept a file without the header.
```
member_email,path,group_or_email
member@example.com,/Team Folder/Shared/file.txt,group@example.com
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column               | Description                            |
|----------------------|----------------------------------------|
| status               | Status of the operation                |
| reason               | Reason of failure or skipped operation |
| input.member_email   | Email address of the member to remove. |
| input.path           | Path to the shared folder.             |
| input.group_or_email | Group name or email address to remove. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder mount add
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/mount/add.md
---

# dropbox team runas sharedfolder mount add

Mount shared folders to team members' accounts on their behalf, ensuring proper folder synchronization 

Admin action to mount shared folders in member accounts when they cannot do it themselves. Useful for troubleshooting sync issues, helping non-technical users, or ensuring critical folders are properly mounted. Operates as if the member performed the action.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder mount add -member-email EMAIL -shared-folder-id SHARED_FOLDER_ID
```

## Options:

**-base-path**
: Base path of the shared folder to mount.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Email address of the member

**-peer**
: Account alias. Default: default

**-shared-folder-id**
: Shared folder ID

# Results

## Report: mount

This report shows a list of shared folders.
The command will generate a report in three different formats. `mount.csv`, `mount.json`, and `mount.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mount_0000.xlsx`, `mount_0001.xlsx`, `mount_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder mount delete
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/mount/delete.md
---

# dropbox team runas sharedfolder mount delete

Unmount shared folders from team members' accounts on their behalf, managing folder synchronization 

Admin tool to unmount shared folders from member accounts without removing access. Useful for troubleshooting sync issues, managing local storage, or temporarily removing folders from sync. Member retains access and can remount later.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder mount delete -member-email EMAIL -shared-folder-id SHARED_FOLDER_ID
```

## Options:

**-base-path**
: Base path of the shared folder to unmount.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Email address of the member

**-peer**
: Account alias. Default: default

**-shared-folder-id**
: The ID for the shared folder.

# Results

## Report: mount

This report shows a list of shared folders.
The command will generate a report in three different formats. `mount.csv`, `mount.json`, and `mount.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mount_0000.xlsx`, `mount_0001.xlsx`, `mount_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder mount list
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/mount/list.md
---

# dropbox team runas sharedfolder mount list

Display all shared folders currently mounted (synced) to a specific team member's account 

Admin view of which shared folders are actively mounted (syncing) in a member's account. Helps diagnose sync issues, understand storage usage, or verify proper folder access. Distinguishes between mounted and unmounted but accessible folders.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder mount list -member-email EMAIL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Member email address

**-peer**
: Account alias. Default: default

# Results

## Report: mounts

This report shows a list of shared folders.
The command will generate a report in three different formats. `mounts.csv`, `mounts.json`, and `mounts.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mounts_0000.xlsx`, `mounts_0001.xlsx`, `mounts_0002.xlsx`, ...

---
Title: dropbox team runas sharedfolder mount mountable
URL: https://toolbox.watermint.org/commands/dropbox/team/runas/sharedfolder/mount/mountable.md
---

# dropbox team runas sharedfolder mount mountable

Show all available shared folders that a team member can mount but hasn't synced yet 

Lists shared folders accessible to a member but not currently synced to their device. Useful for identifying available folders, helping members find content, or understanding why certain folders aren't appearing locally. Shows potential sync options.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team runas sharedfolder mount mountable -member-email EMAIL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-include-mounted**
: Include mounted folders.. Default: false

**-member-email**
: Member email address

**-peer**
: Account alias. Default: default

# Results

## Report: mountables

This report shows a list of shared folders.
The command will generate a report in three different formats. `mountables.csv`, `mountables.json`, and `mountables.xlsx`.

| Column                | Description                                                                                               |
|-----------------------|-----------------------------------------------------------------------------------------------------------|
| shared_folder_id      | The ID of the shared folder.                                                                              |
| name                  | The name of this shared folder.                                                                           |
| access_type           | The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment) |
| path_lower            | The lower-cased full path of this shared folder.                                                          |
| is_inside_team_folder | Whether this folder is inside of a team folder.                                                           |
| is_team_folder        | Whether this folder is a team folder.                                                                     |
| policy_manage_access  | Who can add and remove members from this shared folder.                                                   |
| policy_shared_link    | Who links can be shared with.                                                                             |
| policy_member_folder  | Who can be a member of this shared folder, as set on the folder itself.                                   |
| policy_member         | Who can be a member of this shared folder, as set on the folder itself (team, or anyone)                  |
| policy_viewer_info    | Who can enable/disable viewer info for this shared folder.                                                |
| owner_team_name       | Team name of the team that owns the folder                                                                |
| access_inheritance    | Access inheritance type                                                                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `mountables_0000.xlsx`, `mountables_0001.xlsx`, `mountables_0002.xlsx`, ...

---
Title: dropbox team sharedlink list
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/list.md
---

# dropbox team sharedlink list

Display comprehensive list of all shared links created by team members with visibility and expiration details 

Comprehensive inventory of all team shared links showing URLs, visibility settings, expiration dates, and creators. Essential for security audits, identifying risky links, and understanding external sharing patterns. Filter by various criteria for focused analysis.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-peer**
: Account alias. Default: default

**-visibility**
: Filter links by visibility (all/public/team_only/password). Options:.   • all (Visibility option: all).   • public (Anyone with the link can access).   • team_only (Only team members can access).   • password (Password protected access).   • team_and_password (Team members only with password).   • shared_folder_only (Only shared folder members can access). Default: all

# Results

## Report: shared_link

This report shows a list of shared links with the shared link owner team member.
The command will generate a report in three different formats. `shared_link.csv`, `shared_link.json`, and `shared_link.xlsx`.

| Column     | Description                                                                                                                                                                                                         |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| url        | URL of the shared link.                                                                                                                                                                                             |
| name       | The linked file name (including extension).                                                                                                                                                                         |
| expires    | Expiration time, if set.                                                                                                                                                                                            |
| path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| email      | Email address of user.                                                                                                                                                                                              |
| status     | The user's status as a member of a specific team. (active/invited/suspended/removed)                                                                                                                                |
| surname    | Surname of the link owner                                                                                                                                                                                           |
| given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `shared_link_0000.xlsx`, `shared_link_0001.xlsx`, `shared_link_0002.xlsx`, ...

---
Title: dropbox team sharedlink cap expiry
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/cap/expiry.md
---

# dropbox team sharedlink cap expiry

Apply expiration date limits to all team shared links for enhanced security and compliance (Irreversible operation)

Applies expiration dates to existing shared links without them. Essential for security compliance and reducing exposure of perpetual links. Can target links by age or apply blanket expiration policy. Helps prevent unauthorized long-term access to shared content.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink cap expiry -at "+72h" -file /PATH/TO/shared_link_list.csv
```

## Options:

**-at**
: New expiry date/time

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Target shared link

| Column | Description     | Example                                  |
|--------|-----------------|------------------------------------------|
| url    | Shared link URL | https://www.dropbox.com/scl/fo/fir9vjelf |

The first line is a header line. The program will accept a file without the header.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team sharedlink cap visibility
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/cap/visibility.md
---

# dropbox team sharedlink cap visibility

Enforce visibility restrictions on team shared links, controlling public access levels (Irreversible operation)

Modifies shared link visibility settings to enforce team security policies. Can restrict public links to team-only or password-protected access. Critical for preventing data leaks and ensuring links comply with organizational security requirements.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink cap visibility -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-new-visibility**
: New visibility setting. Options: team_only (newvisibility: team_only). Default: team_only

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Target shared link

| Column | Description     | Example                                  |
|--------|-----------------|------------------------------------------|
| url    | Shared link URL | https://www.dropbox.com/scl/fo/fir9vjelf |

The first line is a header line. The program will accept a file without the header.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team sharedlink delete links
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/delete/links.md
---

# dropbox team sharedlink delete links

Delete multiple shared links in batch for security compliance or access control cleanup (Irreversible operation)

Bulk deletes shared links based on criteria like age, visibility, or path patterns. Use for security remediation, removing obsolete links, or enforcing new sharing policies. Permanent action that immediately revokes access through deleted links.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink delete links -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Target shared link

| Column | Description     | Example                                  |
|--------|-----------------|------------------------------------------|
| url    | Shared link URL | https://www.dropbox.com/scl/fo/fir9vjelf |

The first line is a header line. The program will accept a file without the header.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team sharedlink delete member
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/delete/member.md
---

# dropbox team sharedlink delete member

Remove all shared links created by a specific team member, useful for departing employees (Irreversible operation)

Removes all shared links created by a specific member, regardless of content location. Essential for secure offboarding, responding to compromised accounts, or enforcing immediate access revocation. Cannot be undone, so use with appropriate authorization.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink delete member -member-email EMAIL
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-member-email**
: Member email address

**-peer**
: Account alias. Default: default

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team sharedlink update expiry
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/update/expiry.md
---

# dropbox team sharedlink update expiry

Modify expiration dates for existing shared links across the team to enforce security policies (Irreversible operation)

Modifies expiration dates for existing shared links to enforce new security policies or extend access for legitimate use cases. Can target specific links or apply bulk updates. Helps maintain balance between security and usability.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink update expiry -file /PATH/TO/DATA_FILE.csv -at +720h
```

## Options:

**-at**
: New expiration date and time

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Target shared link

| Column | Description     | Example                                  |
|--------|-----------------|------------------------------------------|
| url    | Shared link URL | https://www.dropbox.com/scl/fo/fir9vjelf |

The first line is a header line. The program will accept a file without the header.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team sharedlink update password
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/update/password.md
---

# dropbox team sharedlink update password

Add or change passwords on team shared links in batch for enhanced security protection (Irreversible operation)

Applies password protection to existing shared links or updates current passwords. Critical for securing sensitive content shared externally. Can target vulnerable links or apply passwords based on content sensitivity. Notify link recipients of new requirements.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink update password -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Shared link / password pair list

| Column   | Description          | Example                                  |
|----------|----------------------|------------------------------------------|
| url      | Shared link URL      | https://www.dropbox.com/scl/fo/fir9vjelf |
| password | Shared link password | STRONG_PASSWORD                          |

The first line is a header line. The program will accept a file without the header.
```
url,password
https://www.dropbox.com/scl/fo/fir9vjelf,STRONG_PASSWORD
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team sharedlink update visibility
URL: https://toolbox.watermint.org/commands/dropbox/team/sharedlink/update/visibility.md
---

# dropbox team sharedlink update visibility

Change access levels of existing shared links between public, team-only, and password-protected (Irreversible operation)

Updates shared link visibility from public to team-only or other restricted settings. Essential for reducing external exposure and meeting compliance requirements. Can target links by current visibility level or content location. Changes take effect immediately.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team sharedlink update visibility -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-file**
: Path to data file

**-new-visibility**
: New visibility setting. Options: public (newvisibility: public), team_only (newvisibility: team_only). Default: team_only

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Target shared link

| Column | Description     | Example                                  |
|--------|-----------------|------------------------------------------|
| url    | Shared link URL | https://www.dropbox.com/scl/fo/fir9vjelf |

The first line is a header line. The program will accept a file without the header.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column            | Description                                                                                                                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| status            | Status of the operation                                                                                                                                                                                             |
| reason            | Reason of failure or skipped operation                                                                                                                                                                              |
| input.url         | Shared link URL                                                                                                                                                                                                     |
| result.tag        | Entry type (file, or folder)                                                                                                                                                                                        |
| result.url        | URL of the shared link.                                                                                                                                                                                             |
| result.name       | The linked file name (including extension).                                                                                                                                                                         |
| result.expires    | Expiration time, if set.                                                                                                                                                                                            |
| result.path_lower | The lowercased full path in the user's Dropbox.                                                                                                                                                                     |
| result.visibility | The current visibility of the link after considering the shared links policies of the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder). |
| result.email      | Email address of user.                                                                                                                                                                                              |
| result.surname    | Surname of the link owner                                                                                                                                                                                           |
| result.given_name | Given name of the link owner                                                                                                                                                                                        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder add
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/add.md
---

# dropbox team teamfolder add

Create a new team folder for centralized team content storage and collaboration (Irreversible operation)

Creates new team folders with defined access controls and sync settings. Set up departmental folders, project spaces, or archive locations. Configure initial permissions and determine whether content syncs to member devices by default.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder add -name NAME
```

## Options:

**-name**
: Team folder name

**-peer**
: Account alias. Default: default

**-sync-setting**
: Sync setting for the team folder. Options: default (syncsetting: default), not_synced (syncsetting: not_synced). Default: default

# Results

## Report: added

This report shows a list of team folders in the team.
The command will generate a report in three different formats. `added.csv`, `added.json`, and `added.xlsx`.

| Column                 | Description                                                                                |
|------------------------|--------------------------------------------------------------------------------------------|
| name                   | The name of the team folder.                                                               |
| status                 | The status of the team folder (active, archived, or archive_in_progress)                   |
| is_team_shared_dropbox | True if the team has team shared Dropbox                                                   |
| sync_setting           | The sync setting applied to this team folder (default, not_synced, or not_synced_inactive) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `added_0000.xlsx`, `added_0001.xlsx`, `added_0002.xlsx`, ...

---
Title: dropbox team teamfolder archive
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/archive.md
---

# dropbox team teamfolder archive

Archive a team folder to make it read-only while preserving all content and access history (Irreversible operation)

Converts active team folders to archived status, making them read-only while preserving all content and permissions. Use for completed projects, historical records, or compliance requirements. Archived folders can be reactivated if needed.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder archive -name TEAMFOLDER_NAME
```

## Options:

**-name**
: Team folder name

**-peer**
: Account alias. Default: default

---
Title: dropbox team teamfolder list
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/list.md
---

# dropbox team teamfolder list

Display all team folders with their status, sync settings, and member access information 

Comprehensive list of all team folders showing names, status (active/archived), sync settings, and access levels. Fundamental for team folder governance, planning reorganizations, and understanding team structure. Export for documentation or analysis.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder list 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: team_folder

This report shows a list of team folders in the team.
The command will generate a report in three different formats. `team_folder.csv`, `team_folder.json`, and `team_folder.xlsx`.

| Column                 | Description                                                                                |
|------------------------|--------------------------------------------------------------------------------------------|
| name                   | The name of the team folder.                                                               |
| status                 | The status of the team folder (active, archived, or archive_in_progress)                   |
| is_team_shared_dropbox | True if the team has team shared Dropbox                                                   |
| sync_setting           | The sync setting applied to this team folder (default, not_synced, or not_synced_inactive) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `team_folder_0000.xlsx`, `team_folder_0001.xlsx`, `team_folder_0002.xlsx`, ...

---
Title: dropbox team teamfolder permdelete
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/permdelete.md
---

# dropbox team teamfolder permdelete

Permanently delete an archived team folder and all its contents, irreversibly freeing storage (Irreversible operation)

Irreversibly deletes a team folder and all contained files. Use only with proper authorization and after confirming no critical data remains. Essential for compliance with data retention policies or removing sensitive content. This action cannot be undone.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder permdelete -name TEAMFOLDER_NAME
```

## Options:

**-name**
: Team folder name

**-peer**
: Account alias. Default: default

---
Title: dropbox team teamfolder replication
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/replication.md
---

# dropbox team teamfolder replication

Copy an entire team folder with all contents to another team account for migration or backup (Experimental, and Irreversible operation)

Creates an exact duplicate of a team folder preserving structure, permissions, and content. Use for creating backups, setting up test environments, or preparing for major changes. Consider available storage and replication time for large folders.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder replication -name NAME
```

## Options:

**-dst-peer-name**
: Destination team account alias. Default: dst

**-name**
: Team folder name

**-src-peer-name**
: Source team account alias. Default: src

# Results

## Report: verification

This report shows a difference between two folders.
The command will generate a report in three different formats. `verification.csv`, `verification.json`, and `verification.xlsx`.

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder or file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `verification_0000.xlsx`, `verification_0001.xlsx`, `verification_0002.xlsx`, ...

---
Title: dropbox team teamfolder batch archive
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/batch/archive.md
---

# dropbox team teamfolder batch archive

Archive multiple team folders in batch, efficiently managing folder lifecycle and compliance (Irreversible operation)

Bulk archives team folders based on criteria like age, name patterns, or activity levels. Streamlines folder lifecycle management and helps maintain organized team spaces. Preserves all content while preventing new modifications.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder batch archive -file /path/to/file.csv
```

## Options:

**-file**
: Data file for a list of team folder names

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for batch creating team folders.

| Column | Description         | Example |
|--------|---------------------|---------|
| name   | Name of team folder | Sales   |

The first line is a header line. The program will accept a file without the header.
```
name
Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                        | Description                                                                                |
|-------------------------------|--------------------------------------------------------------------------------------------|
| status                        | Status of the operation                                                                    |
| reason                        | Reason of failure or skipped operation                                                     |
| input.name                    | Name of team folder                                                                        |
| result.name                   | The name of the team folder.                                                               |
| result.status                 | The status of the team folder (active, archived, or archive_in_progress)                   |
| result.is_team_shared_dropbox | True if the team has team shared Dropbox                                                   |
| result.sync_setting           | The sync setting applied to this team folder (default, not_synced, or not_synced_inactive) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder batch permdelete
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/batch/permdelete.md
---

# dropbox team teamfolder batch permdelete

Permanently delete multiple archived team folders in batch, freeing storage space (Irreversible operation)

Permanently deletes multiple team folders and all their contents without possibility of recovery. Use only with proper authorization for removing obsolete data, complying with retention policies, or emergency cleanup. This action cannot be undone.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder batch permdelete -file /path/to/file.csv
```

## Options:

**-file**
: Data file for a list of team folder names

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Data file for batch creating team folders.

| Column | Description         | Example |
|--------|---------------------|---------|
| name   | Name of team folder | Sales   |

The first line is a header line. The program will accept a file without the header.
```
name
Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column     | Description                            |
|------------|----------------------------------------|
| status     | Status of the operation                |
| reason     | Reason of failure or skipped operation |
| input.name | Name of team folder                    |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder batch replication
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/batch/replication.md
---

# dropbox team teamfolder batch replication

Replicate multiple team folders to another team account in batch for migration or backup (Irreversible operation)

Creates copies of multiple team folders with their complete structures and permissions. Useful for creating backups, setting up parallel environments, or preparing for migrations. Consider storage implications before large replications.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder batch replication -file TEAMFOLDER_NAME_LIST.csv
```

## Options:

**-dst-peer-name**
: Destination team account alias. Default: dst

**-file**
: Data file for a list of team folder names

**-src-peer-name**
: Source team account alias. Default: src

# File formats

## Format: File

Data file for batch creating team folders.

| Column | Description         | Example |
|--------|---------------------|---------|
| name   | Name of team folder | Sales   |

The first line is a header line. The program will accept a file without the header.
```
name
Sales
```

# Results

## Report: verification

This report shows a difference between two folders.
The command will generate a report in three different formats. `verification.csv`, `verification.json`, and `verification.xlsx`.

| Column     | Description                                                                                                                                                                            |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| diff_type  | Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing. |
| left_path  | path of left                                                                                                                                                                           |
| left_kind  | folder or file                                                                                                                                                                         |
| left_size  | size of left file                                                                                                                                                                      |
| left_hash  | Content hash of left file                                                                                                                                                              |
| right_path | path of right                                                                                                                                                                          |
| right_kind | folder or file                                                                                                                                                                         |
| right_size | size of right file                                                                                                                                                                     |
| right_hash | Content hash of right file                                                                                                                                                             |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `verification_0000.xlsx`, `verification_0001.xlsx`, `verification_0002.xlsx`, ...

---
Title: dropbox team teamfolder file list
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/file/list.md
---

# dropbox team teamfolder file list

Display all files and subfolders within team folders for content inventory and management 

Enumerates all files in team folders with details like size, modification dates, and paths. Essential for content audits, migration planning, and understanding data distribution. Can filter by file types or patterns for targeted analysis.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder file list 
```

## Options:

**-folder-name**
: List only for the folder matched to the name. Filter by exact match to the name.

**-folder-name-prefix**
: List only for the folder matched to the name. Filter by name match to the prefix.

**-folder-name-suffix**
: List only for the folder matched to the name. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

# Results

## Report: errors

This report shows the transaction result.
The command will generate a report in three different formats. `errors.csv`, `errors.json`, and `errors.xlsx`.

| Column          | Description                            |
|-----------------|----------------------------------------|
| status          | Status of the operation                |
| reason          | Reason of failure or skipped operation |
| input.namespace | Namespace                              |
| input.path      | Path                                   |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...

## Report: namespace_file

This report shows a list of namespaces in the team.
The command will generate a report in three different formats. `namespace_file.csv`, `namespace_file.json`, and `namespace_file.xlsx`.

| Column                 | Description                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------|
| namespace_type         | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)             |
| namespace_name         | The name of this namespace                                                                             |
| namespace_member_email | If this is a team member or app folder, the email address of the owning team member.                   |
| tag                    | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name                   | The last component of the path (including extension).                                                  |
| path_display           | The cased path to be used for display purposes only.                                                   |
| client_modified        | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified        | The last time the file was modified on Dropbox.                                                        |
| size                   | The file size in bytes.                                                                                |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_file_0000.xlsx`, `namespace_file_0001.xlsx`, `namespace_file_0002.xlsx`, ...

---
Title: dropbox team teamfolder file size
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/file/size.md
---

# dropbox team teamfolder file size

Calculate storage usage for team folders, providing detailed size analytics for capacity planning 

Analyzes storage consumption within team folders showing size distribution and largest files. Essential for capacity planning, identifying candidates for archival, and understanding storage costs. Helps optimize team folder usage and plan for growth.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder file size 
```

## Options:

**-depth**
: Depth. Default: 3

**-folder-name**
: List only folders matching the name. Filter by exact match to the name.

**-folder-name-prefix**
: List only folders matching the name. Filter by name match to the prefix.

**-folder-name-suffix**
: List only folders matching the name. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

# Results

## Report: namespace_size

Namespace size in bytes
The command will generate a report in three different formats. `namespace_size.csv`, `namespace_size.json`, and `namespace_size.xlsx`.

| Column               | Description                                                                                |
|----------------------|--------------------------------------------------------------------------------------------|
| namespace_name       | The name of this namespace                                                                 |
| namespace_id         | The ID of this namespace.                                                                  |
| namespace_type       | The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder) |
| owner_team_member_id | If this is a team member or app folder, the ID of the owning team member.                  |
| path                 | Path to the folder                                                                         |
| count_file           | Number of files under the folder                                                           |
| count_folder         | Number of folders under the folder                                                         |
| count_descendant     | Number of files and folders under the folder                                               |
| size                 | Size of the folder                                                                         |
| depth                | Namespace depth                                                                            |
| mod_time_earliest    | Earliest modification time in namespace                                                    |
| mod_time_latest      | Latest modification time in namespace                                                      |
| api_complexity       | Folder complexity index for API operations                                                 |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `namespace_size_0000.xlsx`, `namespace_size_0001.xlsx`, `namespace_size_0002.xlsx`, ...

---
Title: dropbox team teamfolder file lock list
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/file/lock/list.md
---

# dropbox team teamfolder file lock list

Display all locked files within team folders, identifying collaboration bottlenecks 

Lists all currently locked files in team folders with lock holder information and lock duration. Helps identify collaboration bottlenecks, stale locks, and users who may need assistance. Essential for maintaining smooth team workflows.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder file lock list -path /DROPBOX/PATH/TO/LIST -team-folder NAME
```

## Options:

**-path**
: Path

**-peer**
: Account alias. Default: default

**-team-folder**
: Team folder name

# Results

## Report: lock

Lock information
The command will generate a report in three different formats. `lock.csv`, `lock.json`, and `lock.xlsx`.

| Column           | Description                                                                                            |
|------------------|--------------------------------------------------------------------------------------------------------|
| tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| name             | The last component of the path (including extension).                                                  |
| path_display     | The cased path to be used for display purposes only.                                                   |
| client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| server_modified  | The last time the file was modified on Dropbox.                                                        |
| size             | The file size in bytes.                                                                                |
| is_lock_holder   | True if caller holds the file lock                                                                     |
| lock_holder_name | The display name of the lock holder.                                                                   |
| lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `lock_0000.xlsx`, `lock_0001.xlsx`, `lock_0002.xlsx`, ...

---
Title: dropbox team teamfolder file lock release
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/file/lock/release.md
---

# dropbox team teamfolder file lock release

Release specific file locks in team folders to enable collaborative editing 

Releases individual file locks in team folders when specific files are blocking work. More precise than bulk release when only certain files need unlocking. Useful for resolving urgent editing conflicts without affecting other locked files.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder file lock release -path /DROPBOX/PATH/TO/RELEASE -team-folder NAME
```

## Options:

**-path**
: Path to release lock

**-peer**
: Account alias. Default: default

**-team-folder**
: Team folder name

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder file lock all release
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/file/lock/all/release.md
---

# dropbox team teamfolder file lock all release

Release all file locks within a team folder path, resolving editing conflicts in bulk 

Bulk releases all file locks within specified team folders. Use when multiple locks are blocking team productivity or after system issues. Notifies lock holders when possible. May cause loss of unsaved changes in locked files.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder file lock all release -path /DROPBOX/PATH/TO/RELEASE -team-folder NAME
```

## Options:

**-batch-size**
: Operation batch size. Default: 100

**-path**
: Path to release lock

**-peer**
: Account alias. Default: default

**-team-folder**
: Team folder name

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                  | Description                                                                                            |
|-------------------------|--------------------------------------------------------------------------------------------------------|
| status                  | Status of the operation                                                                                |
| reason                  | Reason of failure or skipped operation                                                                 |
| input.path              | Path                                                                                                   |
| result.tag              | Type of entry. `file`, `folder`, or `deleted`                                                          |
| result.client_modified  | For files, this is the modification time set by the desktop client when the file was added to Dropbox. |
| result.server_modified  | The last time the file was modified on Dropbox.                                                        |
| result.size             | The file size in bytes.                                                                                |
| result.is_lock_holder   | True if caller holds the file lock                                                                     |
| result.lock_holder_name | The display name of the lock holder.                                                                   |
| result.lock_created     | The timestamp when the lock was created.                                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder member add
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/member/add.md
---

# dropbox team teamfolder member add

Add multiple users or groups to team folders in batch, streamlining access provisioning (Irreversible operation)

Grants access to team folders for individuals or groups with defined permissions (view/edit). Use for onboarding, project assignments, or expanding access. Group additions efficiently manage permissions through group membership rather than individual assignments.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder member add -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-admin-group-name**
: Temporary group name for admin operation. Default: watermint-toolbox-admin

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Team folder and member list for adding access. Each row can have one member and the one folder. If you want to add two or more members to the folder, please create rows for those members. Similarly, if you want to add a member to two or more folders, please create rows for those folders.

| Column                     | Description                                                                                                  | Example |
|----------------------------|--------------------------------------------------------------------------------------------------------------|---------|
| team_folder_name           | Team folder name                                                                                             | Sales   |
| path                       | Relative path from the team folder root. Leave empty if you want to add a member to root of the team folder. | Report  |
| access_type                | Access type (viewer/editor)                                                                                  | editor  |
| group_name_or_member_email | Group name or member email address                                                                           | Sales   |

The first line is a header line. The program will accept a file without the header.
```
team_folder_name,path,access_type,group_name_or_member_email
Sales,Report,editor,Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                           | Description                                                                                                  |
|----------------------------------|--------------------------------------------------------------------------------------------------------------|
| status                           | Status of the operation                                                                                      |
| reason                           | Reason of failure or skipped operation                                                                       |
| input.team_folder_name           | Team folder name                                                                                             |
| input.path                       | Relative path from the team folder root. Leave empty if you want to add a member to root of the team folder. |
| input.access_type                | Access type (viewer/editor)                                                                                  |
| input.group_name_or_member_email | Group name or member email address                                                                           |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder member delete
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/member/delete.md
---

# dropbox team teamfolder member delete

Remove multiple users or groups from team folders in batch, managing access revocation efficiently (Irreversible operation)

Revokes team folder access for specific members or entire groups. Essential for offboarding, project completion, or security responses. Removal is immediate and affects all folder contents. Consider data retention needs before removing members with edit access.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder member delete -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-admin-group-name**
: Temporary group name for admin operation. Default: watermint-toolbox-admin

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Team folder and member list for removing access. Each row can have one member and one folder. If you want to remove two or more members from the folder, please create rows for those members. Similarly, if you want to remove a member from two or more folders, please create rows for those folders.

| Column                     | Description                                                                                                           | Example |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------|---------|
| team_folder_name           | Team folder name                                                                                                      | Sales   |
| path                       | Relative path from the team folder root. Leave empty if you want to remove a member from the root of the team folder. | Report  |
| group_name_or_member_email | Group name or member email address                                                                                    | Sales   |

The first line is a header line. The program will accept a file without the header.
```
team_folder_name,path,group_name_or_member_email
Sales,Report,Sales
```

# Results

## Report: operation_log

This report shows the transaction result.
The command will generate a report in three different formats. `operation_log.csv`, `operation_log.json`, and `operation_log.xlsx`.

| Column                           | Description                                                                                                           |
|----------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| status                           | Status of the operation                                                                                               |
| reason                           | Reason of failure or skipped operation                                                                                |
| input.team_folder_name           | Team folder name                                                                                                      |
| input.path                       | Relative path from the team folder root. Leave empty if you want to remove a member from the root of the team folder. |
| input.group_name_or_member_email | Group name or member email address                                                                                    |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...

---
Title: dropbox team teamfolder member list
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/member/list.md
---

# dropbox team teamfolder member list

Display all members with access to each team folder, showing permission levels and access types 

Shows complete membership for all team folders including permission levels and whether access is direct or through groups. Critical for access audits, security reviews, and understanding who can access sensitive content. Identifies over-privileged access.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder member list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-member-type-external**
: Filter folder members. Keep only members that are external (not in the same team). Note: Invited members are marked as external member.

**-member-type-internal**
: Filter folder members. Keep only members that are internal (in the same team). Note: Invited members are marked as external member.

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: membership

This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.
The command will generate a report in three different formats. `membership.csv`, `membership.json`, and `membership.xlsx`.

| Column          | Description                                                                                                                          |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| path            | Path                                                                                                                                 |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)                             |
| owner_team_name | Team name of the team that owns the folder                                                                                           |
| access_type     | User's access level for this folder                                                                                                  |
| member_type     | Type of this member (user, group, or invitee)                                                                                        |
| member_name     | Name of this member                                                                                                                  |
| member_email    | Email address of this member                                                                                                         |
| same_team       | Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not. |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `membership_0000.xlsx`, `membership_0001.xlsx`, `membership_0002.xlsx`, ...

## Report: no_member

This report shows folders without members.
The command will generate a report in three different formats. `no_member.csv`, `no_member.json`, and `no_member.xlsx`.

| Column          | Description                                                                                              |
|-----------------|----------------------------------------------------------------------------------------------------------|
| owner_team_name | Team name of the team that owns the folder                                                               |
| path            | Path                                                                                                     |
| folder_type     | Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `no_member_0000.xlsx`, `no_member_0001.xlsx`, `no_member_0002.xlsx`, ...

---
Title: dropbox team teamfolder partial replication
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/partial/replication.md
---

# dropbox team teamfolder partial replication

Selectively replicate team folder contents to another team, enabling flexible content migration (Irreversible operation)

Copies selected subfolders or files from team folders rather than entire structures. Useful for creating targeted backups, extracting project deliverables, or migrating specific content. More efficient than full replication when only portions are needed.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder partial replication -src-team-folder-name SRC_TEAMFOLDER_NAME -src-path /REL/PATH/SRC -dst-team-folder-name DST_TEAMFOLDER_NAME -dst-path /REL/PATH/DST
```

## Options:

**-base-path**
: Base path for partial replication. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-dst**
: Destination account alias. Default: dst

**-dst-path**
: Destination path

**-dst-team-folder-name**
: Destination team folder name

**-src**
: Peer name for the src team. Default: src

**-src-path**
: Relative path from the team folder (please specify '/' for the team folder root)

**-src-team-folder-name**
: Source team folder name

---
Title: dropbox team teamfolder policy list
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/policy/list.md
---

# dropbox team teamfolder policy list

Display all access policies and restrictions applied to team folders for governance review 

Shows all policies governing team folder behavior including sync defaults, sharing restrictions, and access controls. Helps understand why folders behave certain ways and ensures policy compliance. Reference before creating new folders or modifying settings.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder policy list 
```

## Options:

**-base-path**
: Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a personal folder with your username will be selected. This is convenient for referencing or uploading files in your personal folder, as you don't need to include the folder name with your username in the path. On the other hand, if you specify `root`, you can access all folders with permissions. On the other hand, when accessing your personal folder, you need to specify a path that includes the name of your personal folder.. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-folder-name**
: Filter by folder name. Filter by exact match to the name.

**-folder-name-prefix**
: Filter by folder name. Filter by name match to the prefix.

**-folder-name-suffix**
: Filter by folder name. Filter by name match to the suffix.

**-peer**
: Account alias. Default: default

**-scan-timeout**
: Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEAMFOLDER_NAME/:ERROR-SCAN-TIMEOUT:/SUBFOLDER_NAME`.. Options: short (scantimeout: short), long (scantimeout: long). Default: short

# Results

## Report: policy

This report shows a list of shared folders and team folders with their current policy settings.
The command will generate a report in three different formats. `policy.csv`, `policy.json`, and `policy.xlsx`.

| Column               | Description                                                                                              |
|----------------------|----------------------------------------------------------------------------------------------------------|
| path                 | Path                                                                                                     |
| is_team_folder       | `true` if the folder is a team folder, or inside of a team folder                                        |
| owner_team_name      | Team name of the team that owns the folder                                                               |
| policy_manage_access | Who can add and remove members from this shared folder.                                                  |
| policy_shared_link   | Who links can be shared with.                                                                            |
| policy_member        | Who can be a member of this shared folder, taking into account both the folder and the team-wide policy. |
| policy_viewer_info   | Who can enable/disable viewer info for this shared folder.                                               |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `policy_0000.xlsx`, `policy_0001.xlsx`, `policy_0002.xlsx`, ...

---
Title: dropbox team teamfolder sync setting list
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/sync/setting/list.md
---

# dropbox team teamfolder sync setting list

Display sync configuration for all team folders, showing default sync behavior for members 

Shows current sync settings for all team folders indicating whether they automatically sync to new members' devices. Helps understand bandwidth impact, storage requirements, and ensures appropriate content distribution policies.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder sync setting list 
```

## Options:

**-peer**
: Account alias. Default: default

**-scan-all**
: Perform a scan for all depths (can take considerable time depending on folder structure). Default: false

**-show-all**
: Show all scanned folders. Default: false

# Results

## Report: folders

This report shows a list of metadata of files or folders in the path.
The command will generate a report in three different formats. `folders.csv`, `folders.json`, and `folders.xlsx`.

| Column                      | Description                                                                                                          |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------|
| id                          | A unique identifier for the file.                                                                                    |
| tag                         | Type of entry. `file`, `folder`, or `deleted`                                                                        |
| name                        | The last component of the path (including extension).                                                                |
| path_lower                  | The lowercased full path in the user's Dropbox. This always starts with a slash.                                     |
| path_display                | The cased path to be used for display purposes only.                                                                 |
| client_modified             | For files, this is the modification time set by the desktop client when the file was added to Dropbox.               |
| server_modified             | The last time the file was modified on Dropbox.                                                                      |
| revision                    | A unique identifier for the current revision of a file.                                                              |
| size                        | The file size in bytes.                                                                                              |
| content_hash                | A hash of the file content.                                                                                          |
| has_explicit_shared_members | If true, the results will include a flag for each file indicating whether or not that file has any explicit members. |
| shared_folder_id            | If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.                 |
| parent_shared_folder_id     | ID of shared folder that holds this file.                                                                            |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `folders_0000.xlsx`, `folders_0001.xlsx`, `folders_0002.xlsx`, ...

## Report: settings

Folder settings
The command will generate a report in three different formats. `settings.csv`, `settings.json`, and `settings.xlsx`.

| Column       | Description                                               |
|--------------|-----------------------------------------------------------|
| team_folder  | Team folder name                                          |
| path         | Path (Relative to the team folder. Blank for first level) |
| sync_setting | Sync setting                                              |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `settings_0000.xlsx`, `settings_0001.xlsx`, `settings_0002.xlsx`, ...

---
Title: dropbox team teamfolder sync setting update
URL: https://toolbox.watermint.org/commands/dropbox/team/teamfolder/sync/setting/update.md
---

# dropbox team teamfolder sync setting update

Modify sync settings for multiple team folders in batch, controlling automatic synchronization behavior 

Modifies sync behavior for team folders between automatic sync to all members or manual sync selection. Use to reduce storage usage on devices, manage bandwidth, or ensure critical folders sync automatically. Apply changes during low-activity periods.

# Usage

This document uses the Desktop folder for command example.
```
tbx dropbox team teamfolder sync setting update -file /PATH/TO/DATA_FILE.csv
```

## Options:

**-file**
: Path to data file

**-peer**
: Account alias. Default: default

# File formats

## Format: File

Sync settings for team folders

| Column       | Description                       | Example         |
|--------------|-----------------------------------|-----------------|
| path         | Path to the target folder         | /Sales/Forecast |
| sync_setting | Sync setting (default/not_synced) | not_synced      |

The first line is a header line. The program will accept a file without the header.
```
path,sync_setting
/Sales/Forecast,not_synced
```

# Results

## Report: updated

This report shows the transaction result.
The command will generate a report in three different formats. `updated.csv`, `updated.json`, and `updated.xlsx`.

| Column                | Description                                                              |
|-----------------------|--------------------------------------------------------------------------|
| status                | Status of the operation                                                  |
| reason                | Reason of failure or skipped operation                                   |
| input.path            | Path to the target folder                                                |
| input.sync_setting    | Sync setting (default/not_synced)                                        |
| result.team_folder_id | Team folder ID                                                           |
| result.name           | The name of the team folder.                                             |
| result.status         | The status of the team folder (active, archived, or archive_in_progress) |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `updated_0000.xlsx`, `updated_0001.xlsx`, `updated_0002.xlsx`, ...

---
Title: figma account info
URL: https://toolbox.watermint.org/commands/figma/account/info.md
---

# figma account info

Retrieve current user information 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma account info 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: user

User information
The command will generate a report in three different formats. `user.csv`, `user.json`, and `user.xlsx`.

| Column  | Description                              |
|---------|------------------------------------------|
| id      | Unique stable id of the user             |
| handle  | Name of the user                         |
| img_url | URL link to the user's profile image     |
| email   | Email associated with the user's account |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `user_0000.xlsx`, `user_0001.xlsx`, `user_0002.xlsx`, ...

---
Title: figma file info
URL: https://toolbox.watermint.org/commands/figma/file/info.md
---

# figma file info

Show information of the Figma file 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma file info -key FILE_KEY
```

## Options:

**-all-nodes**
: Include all node information. Default: false

**-key**
: File key

**-peer**
: Account alias. Default: default

# Results

## Report: document

Figma Document
The command will generate a report in three different formats. `document.csv`, `document.json`, and `document.xlsx`.

| Column       | Description                      |
|--------------|----------------------------------|
| name         | Name of the document             |
| role         | Your role                        |
| lastModified | Last modified timestamp          |
| editorType   | Figma editor type (figma/figjam) |
| version      | Version of the document          |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `document_0000.xlsx`, `document_0001.xlsx`, `document_0002.xlsx`, ...

## Report: node

Node of the Figma document
The command will generate a report in three different formats. `node.csv`, `node.json`, and `node.xlsx`.

| Column              | Description                                            |
|---------------------|--------------------------------------------------------|
| id                  | Node ID                                                |
| type                | Type of the node                                       |
| name                | Name of the node                                       |
| absoluteBoundingBox | Bounding box of the node in absolute space coordinates |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `node_0000.xlsx`, `node_0001.xlsx`, `node_0002.xlsx`, ...

---
Title: figma file list
URL: https://toolbox.watermint.org/commands/figma/file/list.md
---

# figma file list

List files in the Figma Project 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma file list -project-id PROJECT_ID. Use `services figma project list` command to retrieve PROJECT_IDs on your team.
```

## Options:

**-peer**
: Account alias. Default: default

**-project-id**
: Project ID

# Results

## Report: files

Figma file
The command will generate a report in three different formats. `files.csv`, `files.json`, and `files.xlsx`.

| Column       | Description             |
|--------------|-------------------------|
| key          | Figma file key          |
| name         | Name of the document    |
| thumbnailUrl | Thumbnail URL           |
| lastModified | Last modified timestamp |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `files_0000.xlsx`, `files_0001.xlsx`, `files_0002.xlsx`, ...

---
Title: figma file export frame
URL: https://toolbox.watermint.org/commands/figma/file/export/frame.md
---

# figma file export frame

Export all frames of the Figma file 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma file export frame -key FILE_KEY -path /LOCAL/PATH/TO/EXPORT
```

## Options:

**-format**
: Export format (png/jpg/svg/pdf). Options:.   • jpg (Format: jpg).   • png (Format: png).   • svg (Format: svg).   • pdf (PDF document format). Default: pdf

**-key**
: File key

**-path**
: Output folder path

**-peer**
: Account alias. Default: default

**-scale**
: Export scale in percent range from 1 to 400 (default 100). Default: 100

---
Title: figma file export node
URL: https://toolbox.watermint.org/commands/figma/file/export/node.md
---

# figma file export node

Export Figma document Node 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma file export node -key FILE_KEY -id NODE_ID -path /LOCAL/PATH/TO/EXPORT
```

## Options:

**-format**
: Export format (png/jpg/svg/pdf). Options:.   • jpg (Format: jpg).   • png (Format: png).   • svg (Format: svg).   • pdf (PDF document format). Default: pdf

**-id**
: Node ID

**-key**
: File Key

**-path**
: Output folder path

**-peer**
: Account alias. Default: default

**-scale**
: Export scale in percent range from 1 to 400 (default 100). Default: 100

---
Title: figma file export page
URL: https://toolbox.watermint.org/commands/figma/file/export/page.md
---

# figma file export page

Export all pages of the Figma file 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma file export page -key FILE_KEY -path /LOCAL/PATH/TO/EXPORT
```

## Options:

**-format**
: Export format (png/jpg/svg/pdf). Options:.   • jpg (Format: jpg).   • png (Format: png).   • svg (Format: svg).   • pdf (PDF document format). Default: pdf

**-key**
: File key

**-path**
: Output folder path

**-peer**
: Account alias. Default: default

**-scale**
: Export scale in percent range from 1 to 400 (default 100). Default: 100

---
Title: figma file export all page
URL: https://toolbox.watermint.org/commands/figma/file/export/all/page.md
---

# figma file export all page

Export all files/pages under the team 

This command exports all pages for the files below the team. However, if the same file already exists in the export destination, the watermint toolbox compares the timestamps and downloads only if there are updates. Also, if the page does not contain any content, the process is skipped.

# Usage

This document uses the Desktop folder for command example.
```
tbx figma file export all page -path /LOCAL/PATH/TO/EXPORT -team-id TEAM_ID
```

## Options:

**-format**
: Export format (png/jpg/svg/pdf). Options:.   • jpg (Format: jpg).   • png (Format: png).   • svg (Format: svg).   • pdf (PDF document format). Default: pdf

**-path**
: Output folder path

**-peer**
: Account alias. Default: default

**-scale**
: Export scale in percent range from 1 to 400 (default 100). Default: 100

**-team-id**
: Team ID. To obtain a team id, navigate to a team page of a team you are a part of. The team id will be present in the URL after the word team and before your team name.

---
Title: figma project list
URL: https://toolbox.watermint.org/commands/figma/project/list.md
---

# figma project list

List projects of the team 

# Usage

This document uses the Desktop folder for command example.
```
tbx figma project list -team-id TEAM_ID. To obtain a team id, navigate to a team page of a team you are a part of. The team id will be present in the URL after the word team and before your team name.
```

## Options:

**-peer**
: Account alias. Default: default

**-team-id**
: Team ID. To obtain a team id, navigate to a team page of a team you are a part of. The team id will be present in the URL after the word team and before your team name.

# Results

## Report: projects

Figma project
The command will generate a report in three different formats. `projects.csv`, `projects.json`, and `projects.xlsx`.

| Column | Description         |
|--------|---------------------|
| id     | Figma Project ID    |
| name   | Name of the project |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `projects_0000.xlsx`, `projects_0001.xlsx`, `projects_0002.xlsx`, ...

---
Title: github profile
URL: https://toolbox.watermint.org/commands/github/profile.md
---

# github profile

Get the authenticated user (Experimental)

# Usage

This document uses the Desktop folder for command example.
```
tbx github profile 
```

## Options:

**-peer**
: Account alias. Default: default

# Results

## Report: user

GitHub user profile
The command will generate a report in three different formats. `user.csv`, `user.json`, and `user.xlsx`.

| Column | Description      |
|--------|------------------|
| login  | Login user name  |
| name   | Name of the user |
| url    | URL of the user  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `user_0000.xlsx`, `user_0001.xlsx`, `user_0002.xlsx`, ...

---
Title: github content get
URL: https://toolbox.watermint.org/commands/github/content/get.md
---

# github content get

Get content metadata of the repository 

# Usage

This document uses the Desktop folder for command example.
```
tbx github content get -owner OWNER -repository REPOSITORY -path PATH
```

## Options:

**-owner**
: Owner of the repository

**-path**
: Path to the content

**-peer**
: Account alias. Default: default

**-ref**
: Name of reference

**-repository**
: Name of the repository

# Results

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

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `content_0000.xlsx`, `content_0001.xlsx`, `content_0002.xlsx`, ...

---
Title: github content put
URL: https://toolbox.watermint.org/commands/github/content/put.md
---

# github content put

Put small text content into the repository 

# Usage

This document uses the Desktop folder for command example.
```
tbx github content put  -owner OWNER -repository REPO -path PATH -content /LOCAL/PATH/TO/content -message MSG
```

## Options:

**-branch**
: Name of the branch

**-content**
: Path to a content file

**-message**
: Commit message

**-owner**
: Owner of the repository

**-path**
: Path to the content

**-peer**
: Account alias. Default: default

**-repository**
: Name of the repository

# Results

## Report: commit

Commit information
The command will generate a report in three different formats. `commit.csv`, `commit.json`, and `commit.xlsx`.

| Column | Description        |
|--------|--------------------|
| sha    | SHA1 of the commit |
| url    | URL of the commit  |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `commit_0000.xlsx`, `commit_0001.xlsx`, `commit_0002.xlsx`, ...

---
Title: github issue list
URL: https://toolbox.watermint.org/commands/github/issue/list.md
---

# github issue list

List issues of the public/private GitHub repository (Experimental)

# Usage

This document uses the Desktop folder for command example.
```
tbx github issue list -owner OWNER -repository REPO
```

## Options:

**-filter**
: Indicates which sorts of issues to return.. Options:.   • assigned (filter: assigned).   • created (filter: created).   • mentioned (filter: mentioned).   • subscribed (filter: subscribed).   • repos (filter: repos).   • all (filter: all). Default: assigned

**-labels**
: A list of comma separated label names.

**-owner**
: Owner of the repository

**-peer**
: Account alias. Default: default

**-repository**
: Repository name

**-since**
: Only show notifications updated after the given time.

**-state**
: Indicates the state of the issues to return.. Options:.   • open (Open issues only).   • closed (Closed issues only).   • all (All issues). Default: open

# Results

## Report: issues

GitHub Issue
The command will generate a report in three different formats. `issues.csv`, `issues.json`, and `issues.xlsx`.

| Column | Description      |
|--------|------------------|
| number | Issue number     |
| url    | URL of the issue |
| title  | Title            |
| state  | Issue state      |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `issues_0000.xlsx`, `issues_0001.xlsx`, `issues_0002.xlsx`, ...

---
Title: github release draft
URL: https://toolbox.watermint.org/commands/github/release/draft.md
---

# github release draft

Create release draft (Experimental, and Irreversible operation)

# Usage

This document uses the Desktop folder for command example.
```
tbx github release draft -owner OWNER -repository REPO -body-file /LOCAL/PATH/TO/BODY.txt -branch BRANCH -name NAME -tag TAG
```

## Options:

**-body-file**
: File path to body text. The file must be encoded in UTF-8 without BOM.

**-branch**
: Name of the target branch

**-name**
: Name of the release

**-owner**
: Owner of the repository

**-peer**
: Account alias. Default: default

**-repository**
: Name of the repository

**-tag**
: Name of the tag

# Results

## Report: release

Release on GitHub
The command will generate a report in three different formats. `release.csv`, `release.json`, and `release.xlsx`.

| Column   | Description        |
|----------|--------------------|
| id       | Release ID         |
| tag_name | Release tag name   |
| name     | Release name       |
| draft    | Release is a draft |
| url      | URL of the release |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `release_0000.xlsx`, `release_0001.xlsx`, `release_0002.xlsx`, ...

---
Title: github release list
URL: https://toolbox.watermint.org/commands/github/release/list.md
---

# github release list

List releases (Experimental)

# Usage

This document uses the Desktop folder for command example.
```
tbx github release list -owner OWNER -repository REPO
```

## Options:

**-owner**
: Repository owner

**-peer**
: Account alias. Default: default

**-repository**
: Repository name

# Results

## Report: releases

Release on GitHub
The command will generate a report in three different formats. `releases.csv`, `releases.json`, and `releases.xlsx`.

| Column   | Description        |
|----------|--------------------|
| tag_name | Release tag name   |
| name     | Release name       |
| draft    | Release is a draft |
| url      | URL of the release |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `releases_0000.xlsx`, `releases_0001.xlsx`, `releases_0002.xlsx`, ...

---
Title: github release asset download
URL: https://toolbox.watermint.org/commands/github/release/asset/download.md
---

# github release asset download

Download assets (Experimental)

# Usage

This document uses the Desktop folder for command example.
```
tbx github release asset download -owner OWNER -repository REPO -path /LOCAL/PATH/TO/DOWNLOAD -release RELEASE
```

## Options:

**-owner**
: Owner of the repository

**-path**
: Path to download

**-peer**
: Account alias. Default: default

**-release**
: Release tag name

**-repository**
: Name of the repository

# Results

## Report: downloads

This report shows the transaction result.
The command will generate a report in three different formats. `downloads.csv`, `downloads.json`, and `downloads.xlsx`.

| Column     | Description                            |
|------------|----------------------------------------|
| status     | Status of the operation                |
| reason     | Reason of failure or skipped operation |
| input.file | File path                              |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `downloads_0000.xlsx`, `downloads_0001.xlsx`, `downloads_0002.xlsx`, ...

---
Title: github release asset list
URL: https://toolbox.watermint.org/commands/github/release/asset/list.md
---

# github release asset list

List assets of GitHub Release (Experimental)

# Usage

This document uses the Desktop folder for command example.
```
tbx github release asset list -owner OWNER -repository REPO -release RELEASE
```

## Options:

**-owner**
: Owner of the repository

**-peer**
: Account alias. Default: default

**-release**
: Release tag name

**-repository**
: Name of the repository

# Results

## Report: assets

GitHub Release assets
The command will generate a report in three different formats. `assets.csv`, `assets.json`, and `assets.xlsx`.

| Column         | Description         |
|----------------|---------------------|
| name           | Name of the asset   |
| size           | Size of the asset   |
| state          | State of the asset  |
| download_count | Number of downloads |
| download_url   | Download URL        |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `assets_0000.xlsx`, `assets_0001.xlsx`, `assets_0002.xlsx`, ...

---
Title: github release asset upload
URL: https://toolbox.watermint.org/commands/github/release/asset/upload.md
---

# github release asset upload

Upload assets file into the GitHub Release (Experimental, and Irreversible operation)

# Usage

This document uses the Desktop folder for command example.
```
tbx github release asset upload -owner OWNER -repository REPO -release RELEASE -asset /LOCAL/PATH/TO/assets
```

## Options:

**-asset**
: Path to assets

**-owner**
: Owner of the repository

**-peer**
: Account alias. Default: default

**-release**
: Release tag name

**-repository**
: Name of the repository

# Results

## Report: uploads

This report shows the transaction result.
The command will generate a report in three different formats. `uploads.csv`, `uploads.json`, and `uploads.xlsx`.

| Column                | Description                            |
|-----------------------|----------------------------------------|
| status                | Status of the operation                |
| reason                | Reason of failure or skipped operation |
| input.file            | File path                              |
| result.name           | Name of the asset                      |
| result.size           | Size of the asset                      |
| result.state          | State of the asset                     |
| result.download_count | Number of downloads                    |
| result.download_url   | Download URL                           |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `uploads_0000.xlsx`, `uploads_0001.xlsx`, `uploads_0002.xlsx`, ...

---
Title: github tag create
URL: https://toolbox.watermint.org/commands/github/tag/create.md
---

# github tag create

Create a tag on the repository (Experimental, and Irreversible operation)

# Usage

This document uses the Desktop folder for command example.
```
tbx github tag create -owner OWNER -repository REPO -sha1 SHA -tag TAG
```

## Options:

**-owner**
: Owner of the repository

**-peer**
: Account alias. Default: default

**-repository**
: Name of the repository

**-sha1**
: SHA1 hash of the commit

**-tag**
: Tag name

# Results

## Report: created

This report shows the transaction result.
The command will generate a report in three different formats. `created.csv`, `created.json`, and `created.xlsx`.

| Column           | Description                            |
|------------------|----------------------------------------|
| status           | Status of the operation                |
| reason           | Reason of failure or skipped operation |
| input.owner      | Owner of the repository                |
| input.repository | Name of the repository                 |
| input.tag        | Tag name                               |
| input.sha_1      | SHA1 hash of the commit                |
| result.tag       | Tag name                               |
| result.sha       | SHA1 sum of the commit                 |
| result.message   | Message of the commit                  |
| result.url       | URL of the tag                         |

In case of a report becomes large, a report in `.xlsx` format will be split into several chunks like follows; `created_0000.xlsx`, `created_0001.xlsx`, `created_0002.xlsx`, ...

---
Title: local file template apply
URL: https://toolbox.watermint.org/commands/local/file/template/apply.md
---

# local file template apply

Apply file/folder structure template to the local path 

# Usage

This document uses the Desktop folder for command example.
```
tbx local file template apply -path /LOCAL/PATH/TO/APPLY -template /LOCAL/PATH/TO/template.json
```

## Options:

**-path**
: Path to apply template

**-template**
: Path to template file

---
Title: local file template capture
URL: https://toolbox.watermint.org/commands/local/file/template/capture.md
---

# local file template capture

Capture file/folder structure as template from local path 

# Usage

This document uses the Desktop folder for command example.
```
tbx local file template capture -out /LOCAL/PATH/template.json -path /LOCAL/PATH/TO/CAPTURE
```

## Options:

**-out**
: Template output path

**-path**
: Capture target path

## Additional documents

- [Path variables](https://toolbox.watermint.org/guides/path-variables.md)
---
Title: Path variables
URL: https://toolbox.watermint.org/guides/path-variables.md
---

# Path variables

Path variables are predefined variables which will be replaced on runtime. For example, if you specify a path with the variable like `{{.DropboxPersonal}}/Pictures`, then the path will be replaced with actual path to Personal Dropbox's folder. But the tool does not guarantee the existence or accuracy.

| Path variable                  | Description                                                                                    |
|--------------------------------|------------------------------------------------------------------------------------------------|
| {{.DropboxPersonal}}           | Path to Dropbox Personal account root folder.                                                  |
| {{.DropboxBusiness}}           | Path to Dropbox for teams account root folder.                                                 |
| {{.DropboxBusinessOrPersonal}} | Path to Dropbox for teams account root folder, or Personal Dropbox account if it is not found. |
| {{.DropboxPersonalOrBusiness}} | Path to Dropbox Personal account root folder, or Business Dropbox account if it is not found.  |
| {{.Home}}                      | The home folder of the current user.                                                           |
| {{.Username}}                  | The name of the current user.                                                                  |
| {{.Hostname}}                  | The host name of the current computer.                                                         |
| {{.ExecPath}}                  | Path to this program.                                                                          |
| {{.Rand8}}                     | Randomized 8 digit number leading with 0.                                                      |
| {{.Year}}                      | Current local year with format 'yyyy' like 2021.                                               |
| {{.Month}}                     | Current local month with format 'mm' like 01.                                                  |
| {{.Day}}                       | Current local day with format 'dd' like 05.                                                    |
| {{.Date}}                      | Current local date with format yyyy-mm-dd.                                                     |
| {{.Time}}                      | Current local time with format HH-MM-SS.                                                       |
| {{.DateUTC}}                   | Current UTC date with format yyyy-mm-dd.                                                       |
| {{.TimeUTC}}                   | Current UTC time with format HH-MM-SS.                                                         |

- [Experimental features](https://toolbox.watermint.org/guides/experimental-features.md)
---
Title: Experimental features
URL: https://toolbox.watermint.org/guides/experimental-features.md
---

# Experimental features

The experimental feature switch is for testing or accessing early access features. You can enable those features with the option `-experiment`. If you want to specify multiple features, please select those features joined with a comma. (e.g. `-experiment feature1,feature2`).

| name                                     | Description                                                                                                                                                                                             |
|------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| batch_balance                            | Execute batch from the largest batch                                                                                                                                                                    |
| batch_non_durable                        | Use non-durable batch framework                                                                                                                                                                         |
| batch_random                             | Execute batch with random batchId order.                                                                                                                                                                |
| batch_sequential                         | Execute batch sequentially in same batchId.                                                                                                                                                             |
| congestion_window_aggressive             | Apply aggressive initial congestion window size                                                                                                                                                         |
| congestion_window_no_limit               | Do not limit concurrency with the congestion window.                                                                                                                                                    |
| dbx_auth_course_grained_scope            | Requests all Dropbox authorization scopes instead of command-defined ones. This is used as a workaround in case the program does not work properly with the authorization scope defined in the command. |
| dbx_auth_redirect                        | Use redirect processing for authorization process to Dropbox                                                                                                                                            |
| dbx_client_conditioner_error100          | Simulate server errors. 100% of requests will fail with server errors.                                                                                                                                  |
| dbx_client_conditioner_error20           | Simulate server errors. 20% of requests will fail with server errors.                                                                                                                                   |
| dbx_client_conditioner_error40           | Simulate server errors. 40% of requests will fail with server errors.                                                                                                                                   |
| dbx_client_conditioner_narrow100         | Simulate rate limit errors. 100% of requests will fail with rate limitation.                                                                                                                            |
| dbx_client_conditioner_narrow20          | Simulate rate limit errors. 20% of requests will fail with rate limitation.                                                                                                                             |
| dbx_client_conditioner_narrow40          | Simulate rate limit errors. 40% of requests will fail with rate limitation.                                                                                                                             |
| dbx_disable_auto_path_root               | Disable auto path root. When disabled, if a user's home namespace is distinct from their root namespace, the user's home namespace will be used as default for all API calls.                           |
| dbx_download_block                       | Download files divided into blocks (improve concurrency)                                                                                                                                                |
| file_sync_disable_reduce_create_folder   | Disable reduce create_folder on syncing file systems. That will create empty folder while syncing folders.                                                                                              |
| legacy_local_to_dbx_connector            | Use legacy local to dropbox sync connector                                                                                                                                                              |
| use_no_cache_dbxfs                       | Use non-cache dropbox file system                                                                                                                                                                       |
| kvs_badger                               | Use Badger as KVS engine                                                                                                                                                                                |
| kvs_badger_turnstile                     | Use Badger as KVS engine with turnstile                                                                                                                                                                 |
| kvs_bitcask                              | Use Bitcask as KVS engine                                                                                                                                                                               |
| kvs_bitcask_turnstile                    | Use Bitcask as the key-value store with turnstile                                                                                                                                                       |
| kvs_sqlite                               | Use Sqlite3 as KVS engine                                                                                                                                                                               |
| kvs_sqlite_turnstile                     | Use SQLite as the key-value store with turnstile                                                                                                                                                        |
| profile_cpu                              | Enable CPU profiler                                                                                                                                                                                     |
| profile_memory                           | Enable memory profiler                                                                                                                                                                                  |
| report_all_columns                       | Show all columns defined as data structure.                                                                                                                                                             |
| suppress_progress                        | Suppress progress indicators                                                                                                                                                                            |
| validate_network_connection_on_bootstrap | Validate network connection on bootstrap                                                                                                                                                                |

- [Troubleshooting](https://toolbox.watermint.org/guides/troubleshooting.md)
---
Title: Troubleshooting
URL: https://toolbox.watermint.org/guides/troubleshooting.md
---

# Firewall or proxy server settings

The tool automatically detects proxy configuration from the system. However, that may fail or cause misconfiguration. In those cases, please use the `-proxy` option to specify proxy server hostname and port number like `-proxy 192.168.1.1:8080` (for proxy server 192.168.1.1, and the port number 8080). 

Note: This tool does not support proxy servers with any authentication such as Basic authentication or NTLM.

# Performance issue

If the command feels slow or stalled, please try re-run with an option `-verbose`. That will show more detailed progress. But in most cases, the cause is simply you have a larger data to process. Otherwise, you already hit a rate limit from API servers. If you want to see rate limit status, please see capture logs and debug for more details. 

The tool automatically adjusts concurrency to avoid additional limitation from API servers. If you want to see current concurrency, please run the command like below. That will show a current window size (maximum concurrency) per endpoint. The debug message "WaiterStatus" reports current concurrency and window sizes. The map "runners" is for operations currently waiting for a result from API servers. The map "window" is for window size for each endpoint. The map "concurrency" is for current concurrency per endpoint. The below example indicates for the endpoint "https://api.dropboxapi.com/2/file_requests/create", the tool does not allow call that endpoint with the concurrency greater than one. That means it requires operation one by one, and there is no easy workaround to speed up operations.
```
tbx job log last -quiet | jq 'select(.msg == "WaiterStatus")' 
{
  "level": "DEBUG",
  "time": "2020-11-10T14:55:57.501+0900",
  "name": "z951.z960.z112064",
  "caller": "nw_congestion/congestion.go:310",
  "msg": "WaiterStatus",
  "goroutine": "gr:284877",
  "runners": {
    "gr:1": {
      "key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create",
      "go_routine_id": "gr:1",
      "running_since": "2020-11-10T14:55:56.124899+09:00"
    }
  },
  "numRunners": 1,
  "waiters": [],
  "numWaiters": 0,
  "window": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 1,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url/check_job_status": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search/continue_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/users/get_current_account": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_reference/get": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/delete_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/get_metadata": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/sharing/list_mountable_folders": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/download": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/export": 4
  },
  "concurrency": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 1
  }
}
```

# Garbled output

If the tool output garbled, please stop the tool with Ctrl+C. This issue usually happens when your console does not have a font to display it. Then, please try to change the font that supports your language. Or, please try the option `-lang en` to overwrite language setting of the tool to English.

In PowerShell, you can change the font with (1) right-click on the title bar, (2) click properties, (3) then choose the font tab, (4) then change to an appropriate font like "MS Gothic."

# Log files

By default, log files are stored under the path "%USERPROFILE%\.toolbox\jobs" (e.g. `C:\Users\USERNAME\.toolbox\jobs`) on windows, or "$HOME/.toolbox/jobs" in Linux or macOS (e.g. `/Users/USERNAME/.toolbox/jobs`). Log files contain information such as (1) Runtime information, e.g. OS type/version/environment variables, (2) Runtime options to the tool (including a copy of input data files), (3) Account information of services such as Dropbox, (4) Request and response data to API servers, (5) Data in services such as file name, metadata, id, URL etc. (depends on the command).

Those logs do not contain password, credentials, or API token. But API tokens are stored under the path "%USERPROFILE%\.toolbox\secrets" (e.g. `C:\Users\USERNAME\.toolbox\secrets`) on windows, or "$HOME/.toolbox/secrets" in Linux or macOS (e.g. `/Users/USERNAME/.toolbox/secrets`). These secrets folder files are obfuscated but please do not share these files to anyone including a service provider support such as Dropbox support.

## Log format

There are several folders and files stored under the `jobs` folder. First, the job folder will be created every run with a name (internally called Job Id) with the format "yyyyMMdd-HHmmSS.xxx". The first part "yyyyMMdd-HHmmSS" is for local date/time of the command start. The second part ".xxx" is the sequential or random three-character ID to avoid conflict with a concurrent run.

Under the job folder, there are subfolders (1) `logs`: runtime logs including request/response data, parameters, or debug information, (2) `reports`: reports folder is for managing generated reports, (3) `kvs`: KVS folder is for runtime database folder. 

On troubleshooting, files under `logs` are essential to understand what happened in runtime. The tool generates several types of logs. Those logs are JSON Lines format. Note: JSON Lines is a format that separate data with line separators. Please see [JSON Lines](https://jsonlines.org/) for more detail about the specification.

Some logs are compressed with gzip format. If a log is compressed, then the file has a suffix '.gz'. Additionally, logs such as capture logs and toolbox logs are divided by certain size. If you want to analyze logs, please consider using `job log` commands. For example, `job log last -quiet` will report toolbox logs of the latest job with decompressed and concatenated.

## Debug logs

The tool will record all debug information into debug logs that have a prefix "toolbox". All records have a source code file name and line at the operation. If you find some suspicious error, then go to the source code and debug it. Some troubleshooting requires statistical analysis such as for performance tuning or out of memory. It is better to work with a tool such as `grep` or [jq](https://stedolan.github.io/jq/). 

If you want to see heap size data in time series, please run the command like below. Then you can see time + heap size in CSV format.
```
tbx job log last -quiet | jq -r 'select(.msg == "Heap stats") | [.time, .HeapInuse] | @csv'
"2020-11-10T14:55:45.725+0900",18604032
"2020-11-10T14:55:50.725+0900",15130624
"2020-11-10T14:55:55.725+0900",17408000
"2020-11-10T14:56:00.725+0900",17014784
"2020-11-10T14:56:05.726+0900",19193856
"2020-11-10T14:56:10.725+0900",19136512
"2020-11-10T14:56:15.726+0900",16637952
"2020-11-10T14:56:20.725+0900",16678912
"2020-11-10T14:56:25.727+0900",16678912
"2020-11-10T14:56:30.730+0900",16678912
"2020-11-10T14:56:35.726+0900",16678912
```

## API transaction logs

The tool will record API requests and responses into capture logs that have a prefix "capture". This capture logs do not contain requests and responses of OAuth. Additionally, API token strings are replaced with `<secret>`.

- [Commands of Dropbox for teams](https://toolbox.watermint.org/guides/dropbox-business.md)
---
Title: Commands of Dropbox for teams
URL: https://toolbox.watermint.org/guides/dropbox-business.md
---

# Member management commands

## Information commands

Below commands are to retrieve information about team members.

| Command                                                               | Description                                                                                                        |
|-----------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| [dropbox team member list](dropbox-team-member-list.md)               | Display comprehensive list of all team members with their status, roles, and account details                       |
| [dropbox team member feature](dropbox-team-member-feature.md)         | Display feature settings and capabilities enabled for specific team members, helping understand member permissions |
| [dropbox team member folder list](dropbox-team-member-folder-list.md) | Display all folders in each team member's account, useful for content auditing and storage analysis                |
| [dropbox team member quota list](dropbox-team-member-quota-list.md)   | Display storage quota assignments for all team members, helping monitor and plan storage distribution              |
| [dropbox team member quota usage](dropbox-team-member-quota-usage.md) | Show actual storage usage for each team member compared to their quotas, identifying storage needs                 |
| [dropbox team activity user](dropbox-team-activity-user.md)           | Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities    |

## Basic management commands

Below commands are for managing team member accounts. Those commands are for a bulk operation by a CSV file.

| Command                                                                                     | Description                                                                                                      |
|---------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| [dropbox team member batch invite](dropbox-team-member-batch-invite.md)                     | Send batch invitations to new team members, streamlining the onboarding process for multiple users               |
| [dropbox team member batch delete](dropbox-team-member-batch-delete.md)                     | Remove multiple team members in batch, efficiently managing team departures and access revocation                |
| [dropbox team member batch detach](dropbox-team-member-batch-detach.md)                     | Convert multiple team accounts to individual Basic accounts, preserving personal data while removing team access |
| [dropbox team member batch reinvite](dropbox-team-member-batch-reinvite.md)                 | Resend invitations to pending members who haven't joined yet, ensuring all intended members receive access       |
| [dropbox team member update batch email](dropbox-team-member-update-batch-email.md)         | Update email addresses for multiple team members in batch, managing email changes efficiently                    |
| [dropbox team member update batch profile](dropbox-team-member-update-batch-profile.md)     | Update profile information for multiple team members including names and job titles in batch                     |
| [dropbox team member update batch visible](dropbox-team-member-update-batch-visible.md)     | Make hidden team members visible in the directory, restoring standard visibility settings                        |
| [dropbox team member update batch invisible](dropbox-team-member-update-batch-invisible.md) | Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors               |
| [dropbox team member quota batch update](dropbox-team-member-quota-batch-update.md)         | Modify storage quotas for multiple team members in batch, managing storage allocation efficiently                |

## Member profile setting commands

Member profile commands are for bulk updating member profile information.\nIf you need to update the members' email addresses, use the `member update email` command. The command `member update email` receives a CSV file to bulk update email addresses.\nIf you need to update the member's display name, use the `member update profile` command.

| Command                                                                                 | Description                                                                                   |
|-----------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| [dropbox team member update batch email](dropbox-team-member-update-batch-email.md)     | Update email addresses for multiple team members in batch, managing email changes efficiently |
| [dropbox team member update batch profile](dropbox-team-member-update-batch-profile.md) | Update profile information for multiple team members including names and job titles in batch  |

## Member storage quota control commands

You can see existing member storage quota setting or usage by the `dropbox team member quota list` and `dropbox team member quota usage` command. If you need to update member quota, use the `dropbox team member quota update` command. The command `dropbox team member quota update` receives CSV input for bulk updating storage quota setting.

| Command                                                                             | Description                                                                                           |
|-------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| [dropbox team member quota list](dropbox-team-member-quota-list.md)                 | Display storage quota assignments for all team members, helping monitor and plan storage distribution |
| [dropbox team member quota usage](dropbox-team-member-quota-usage.md)               | Show actual storage usage for each team member compared to their quotas, identifying storage needs    |
| [dropbox team member quota batch update](dropbox-team-member-quota-batch-update.md) | Modify storage quotas for multiple team members in batch, managing storage allocation efficiently     |

## Suspend/unsuspend member commands

There are two types of commands available for suspending/unsuspending members. If you wanted to suspend/unsuspend a member one by one, please use `dropbox team member suspend` or `dropbox team member unsuspend`. Otherwise, if you want to suspend/unsuspend members through a CSV file, please use the `dropbox team member batch suspend` or `dropbox member batch unsuspend` command.

| Command                                                                       | Description                                                                                           |
|-------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| [dropbox team member suspend](dropbox-team-member-suspend.md)                 | Temporarily suspend a team member's access to their account while preserving all data and settings    |
| [dropbox team member unsuspend](dropbox-team-member-unsuspend.md)             | Restore access for a suspended team member, reactivating their account and all associated permissions |
| [dropbox team member batch suspend](dropbox-team-member-batch-suspend.md)     | Temporarily suspend multiple team members' access while preserving their data and settings            |
| [dropbox team member batch unsuspend](dropbox-team-member-batch-unsuspend.md) | Restore access for multiple suspended team members, reactivating their accounts in batch              |

## Directory restriction commands

Directory restriction is the Dropbox for teams feature to hide a member from others. Below commands update this setting to hide or unhide members from others.

| Command                                                                                     | Description                                                                                        |
|---------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| [dropbox team member update batch visible](dropbox-team-member-update-batch-visible.md)     | Make hidden team members visible in the directory, restoring standard visibility settings          |
| [dropbox team member update batch invisible](dropbox-team-member-update-batch-invisible.md) | Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors |

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                                               | Description                                                                                                |
|-----------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|
| [dropbox team group add](dropbox-team-group-add.md)                   | Create a new group in your team for organizing members and managing permissions collectively               |
| [dropbox team group batch add](dropbox-team-group-batch-add.md)       | Create multiple groups at once using batch processing, efficient for large-scale team organization         |
| [dropbox team group batch delete](dropbox-team-group-batch-delete.md) | Remove multiple groups from your team in batch, streamlining group cleanup and reorganization              |
| [dropbox team group delete](dropbox-team-group-delete.md)             | Remove a specific group from your team, automatically removing all member associations                     |
| [dropbox team group list](dropbox-team-group-list.md)                 | Display all groups in your team with member counts and group management types                              |
| [dropbox team group rename](dropbox-team-group-rename.md)             | Change the name of an existing group to better reflect its purpose or organizational changes               |
| [dropbox team group update type](dropbox-team-group-update-type.md)   | Change how a group is managed (user-managed vs company-managed), affecting who can modify group membership |

## Group member management commands

You can add/delete/update group members by the below commands. If you want to add/delete/update group members by CSV file, use `dropbox team group member batch add`, `dropbox team group member batch delete`, or `dropbox team group member batch update`.

| Command                                                                             | Description                                                                                             |
|-------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| [dropbox team group member add](dropbox-team-group-member-add.md)                   | Add individual team members to a specific group for centralized permission management                   |
| [dropbox team group member delete](dropbox-team-group-member-delete.md)             | Remove a specific member from a group while preserving their other group memberships                    |
| [dropbox team group member list](dropbox-team-group-member-list.md)                 | Display all members belonging to each group, useful for auditing group compositions and access rights   |
| [dropbox team group member batch add](dropbox-team-group-member-batch-add.md)       | Add multiple members to groups efficiently using batch processing, ideal for large team reorganizations |
| [dropbox team group member batch delete](dropbox-team-group-member-batch-delete.md) | Remove multiple members from groups in batch, streamlining group membership management                  |
| [dropbox team group member batch update](dropbox-team-group-member-batch-update.md) | Update group memberships in bulk by adding or removing members, optimizing group composition changes    |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `dropbox team group list`. The command `dropbox team group list` will report the number of members of each group. If it's zero, a group is not currently used to adding permission to folders.\nIf you want to see which folder uses each group, use the command `dropbox team group folder list`. `dropbox team group folder list` will report the group to folder mapping. The report `group_with_no_folders` will show groups with no folders.\nYou can safely remove groups once you check both the number of members and folders. After confirmation, you can bulk delete groups by using the command `dropbox team group batch delete`.

| Command                                                               | Description                                                                                            |
|-----------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| [dropbox team group list](dropbox-team-group-list.md)                 | Display all groups in your team with member counts and group management types                          |
| [dropbox team group folder list](dropbox-team-group-folder-list.md)   | Display all folders accessible by each group, showing group-based content organization and permissions |
| [dropbox team group batch delete](dropbox-team-group-batch-delete.md) | Remove multiple groups from your team in batch, streamlining group cleanup and reorganization          |

# Team content commands

Admins can handle team folders, shared folders or member's folder content through Dropbox Business API. Please be careful to use those commands.
The namespace is a term in the Dropbox API that is used to manage folder permissions or settings. Folder types such as shared folders, team folders, or nested folders in a team folder, member's root folder or member's app folder are all managed as a namespace.\nThe namespace commands can handle all types of folders, including team folders and member's folder. But commands for specific folder types have more features or detailed information in the report.

## Team folder operation commands

You can create, archive or permanently delete team folders by using the below commands. Please consider using `dropbox team teamfolder batch` commands if you need to handle multiple team folders.

| Command                                                                                       | Description                                                                                             |
|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| [dropbox team teamfolder add](dropbox-team-teamfolder-add.md)                                 | Create a new team folder for centralized team content storage and collaboration                         |
| [dropbox team teamfolder archive](dropbox-team-teamfolder-archive.md)                         | Archive a team folder to make it read-only while preserving all content and access history              |
| [dropbox team teamfolder batch archive](dropbox-team-teamfolder-batch-archive.md)             | Archive multiple team folders in batch, efficiently managing folder lifecycle and compliance            |
| [dropbox team teamfolder batch permdelete](dropbox-team-teamfolder-batch-permdelete.md)       | Permanently delete multiple archived team folders in batch, freeing storage space                       |
| [dropbox team teamfolder batch replication](dropbox-team-teamfolder-batch-replication.md)     | Replicate multiple team folders to another team account in batch for migration or backup                |
| [dropbox team teamfolder file size](dropbox-team-teamfolder-file-size.md)                     | Calculate storage usage for team folders, providing detailed size analytics for capacity planning       |
| [dropbox team teamfolder list](dropbox-team-teamfolder-list.md)                               | Display all team folders with their status, sync settings, and member access information                |
| [dropbox team teamfolder permdelete](dropbox-team-teamfolder-permdelete.md)                   | Permanently delete an archived team folder and all its contents, irreversibly freeing storage           |
| [dropbox team teamfolder policy list](dropbox-team-teamfolder-policy-list.md)                 | Display all access policies and restrictions applied to team folders for governance review              |
| [dropbox team teamfolder sync setting list](dropbox-team-teamfolder-sync-setting-list.md)     | Display sync configuration for all team folders, showing default sync behavior for members              |
| [dropbox team teamfolder sync setting update](dropbox-team-teamfolder-sync-setting-update.md) | Modify sync settings for multiple team folders in batch, controlling automatic synchronization behavior |

## Team folder permission commands

You can bulk add or delete members into team folders or sub-folders of a team folder through the below commands.

| Command                                                                           | Description                                                                                        |
|-----------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| [dropbox team teamfolder member list](dropbox-team-teamfolder-member-list.md)     | Display all members with access to each team folder, showing permission levels and access types    |
| [dropbox team teamfolder member add](dropbox-team-teamfolder-member-add.md)       | Add multiple users or groups to team folders in batch, streamlining access provisioning            |
| [dropbox team teamfolder member delete](dropbox-team-teamfolder-member-delete.md) | Remove multiple users or groups from team folders in batch, managing access revocation efficiently |

## Team folder & shared folder commands

The below commands are for both team folders and shared folders of the team.\nIf you wanted to know who actually use specific folders, please consider using the command `dropbox team content mount list`. Mount is a status a user add a shared folder to his/her Dropbox account.

| Command                                                                 | Description                                                                                                                        |
|-------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team content member list](dropbox-team-content-member-list.md) | Display all members with access to team folders and shared folders, showing permission levels and folder relationships             |
| [dropbox team content member size](dropbox-team-content-member-size.md) | Calculate member counts for each team folder and shared folder, helping identify heavily accessed content and optimize permissions |
| [dropbox team content mount list](dropbox-team-content-mount-list.md)   | Display mount status of all shared folders for team members, identifying which folders are actively synced to member devices       |
| [dropbox team content policy list](dropbox-team-content-policy-list.md) | Review all access policies and restrictions applied to team folders and shared folders for governance compliance                   |

## Namespace commands

| Command                                                                     | Description                                                                                              |
|-----------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| [dropbox team namespace list](dropbox-team-namespace-list.md)               | Display all team namespaces including team folders and shared spaces with their configurations           |
| [dropbox team namespace summary](dropbox-team-namespace-summary.md)         | Generate comprehensive summary reports of team namespace usage, member counts, and storage statistics    |
| [dropbox team namespace file list](dropbox-team-namespace-file-list.md)     | Display comprehensive file and folder listings within team namespaces for content inventory and analysis |
| [dropbox team namespace file size](dropbox-team-namespace-file-size.md)     | Calculate storage usage for files and folders in team namespaces, providing detailed size analytics      |
| [dropbox team namespace member list](dropbox-team-namespace-member-list.md) | Show all members with access to each namespace, detailing permissions and access levels                  |

## Team file request commands

| Command                                                           | Description                                                                                                            |
|-------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [dropbox team filerequest list](dropbox-team-filerequest-list.md) | Display all active and closed file requests created by team members, helping track external file collection activities |

## Member file commands

| Command                                                                       | Description                                                                                             |
|-------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| [dropbox team member file permdelete](dropbox-team-member-file-permdelete.md) | Permanently delete files or folders from a team member's account, bypassing trash for immediate removal |

## Team insight

Team Insight is a feature of Dropbox Business that provides a view of team folder summary.

| Command                                                                                         | Description                                                                                               |
|-------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [dropbox team insight scan](dropbox-team-insight-scan.md)                                       | Perform comprehensive data scanning across your team for analytics and insights generation                |
| [dropbox team insight scanretry](dropbox-team-insight-scanretry.md)                             | Re-run failed or incomplete scans to ensure complete data collection for team insights                    |
| [dropbox team insight summarize](dropbox-team-insight-summarize.md)                             | Generate summary reports from scanned team data, providing actionable insights on team usage and patterns |
| [dropbox team insight report teamfoldermember](dropbox-team-insight-report-teamfoldermember.md) | Generate detailed reports on team folder membership, showing access patterns and member distribution      |

# Team shared link commands

The team shared link commands are capable of listing all shared links in the team or update/delete specified shared links.

| Command                                                                                   | Description                                                                                                   |
|-------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| [dropbox team sharedlink list](dropbox-team-sharedlink-list.md)                           | Display comprehensive list of all shared links created by team members with visibility and expiration details |
| [dropbox team sharedlink cap expiry](dropbox-team-sharedlink-cap-expiry.md)               | Apply expiration date limits to all team shared links for enhanced security and compliance                    |
| [dropbox team sharedlink cap visibility](dropbox-team-sharedlink-cap-visibility.md)       | Enforce visibility restrictions on team shared links, controlling public access levels                        |
| [dropbox team sharedlink update expiry](dropbox-team-sharedlink-update-expiry.md)         | Modify expiration dates for existing shared links across the team to enforce security policies                |
| [dropbox team sharedlink update password](dropbox-team-sharedlink-update-password.md)     | Add or change passwords on team shared links in batch for enhanced security protection                        |
| [dropbox team sharedlink update visibility](dropbox-team-sharedlink-update-visibility.md) | Change access levels of existing shared links between public, team-only, and password-protected               |
| [dropbox team sharedlink delete links](dropbox-team-sharedlink-delete-links.md)           | Delete multiple shared links in batch for security compliance or access control cleanup                       |
| [dropbox team sharedlink delete member](dropbox-team-sharedlink-delete-member.md)         | Remove all shared links created by a specific team member, useful for departing employees                     |

## Difference between `dropbox team sharedlink cap` and `dropbox team sharedlink update`

Commands `dropbox team sharedlink update` is for setting a value to the shared links. Commands `dropbox team sharedlink cap` is for putting a cap value to the shared links.\nFor example: if you set expiry by `dropbox team sharedlink update expiry` with the expiration date 2021-05-06. The command will update the expiry to 2021-05-06 even if the existing link has a shorter expiration date like 2021-05-04.\nOn the other hand, `dropbox team sharedlink cap expiry` updates links when the link has a longer expiration date, like 2021-05-07.\n\nSimilarly, the command `dropbox team sharedlink cap visibility` will restrict visibility only when the link has less protected visibility. For example, if you want to ensure shared links without passwords are restricted to the team only. `dropbox team sharedlink cap visibility` will update visibility to the team only when a link is public and has no password.

## Example (list links):

List all public links in the team\n\n\n\nResults are stored in CSV, xlsx, and JSON format. You can modify the report for updating shared links.\nIf you are familiar with the command jq, you can create CSV file directly like below.\n\n\n\nList links filtered by link owner email address:\n\n\n

## Example (delete links):

Delete all link that listed in the CSV file\n\n\n\nIf you are familiar with jq command, you can send data directly from the pipe like below (pass single dash `-` to the `-file` option to read from standard input).\n\nInvalid argument: team sharedlink delete links -file -n
Error: <no value>

watermint toolbox 140.8.313
===========================

© 2016-2025 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox for teams

Usage:
======

tbx  command

Available commands:
===================

| Command | Description              | Notes |
|---------|--------------------------|-------|
| asana   | Asana commands           |       |
| config  | CLI configuration        |       |
| deepl   | DeepL commands           |       |
| dropbox | Dropbox commands         |       |
| figma   | Figma commands           |       |
| github  | GitHub commands          |       |
| license | Show license information |       |
| local   | Commands for local PC    |       |
| log     | Log utilities            |       |
| slack   | Slack commands           |       |
| util    | Utilities                |       |
| version | Show version             |       |\n

# File lock title

Dropbox Business file lock information

## File lock member title

| Command                                                                                   | Description                                                                                                       |
|-------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| [dropbox team member file lock all release](dropbox-team-member-file-lock-all-release.md) | Release all file locks held by a team member under a specified path, resolving editing conflicts                  |
| [dropbox team member file lock list](dropbox-team-member-file-lock-list.md)               | Display all files locked by a specific team member under a given path, identifying potential collaboration blocks |
| [dropbox team member file lock release](dropbox-team-member-file-lock-release.md)         | Release a specific file lock held by a team member, enabling others to edit the file                              |

## File lock team folder title

| Command                                                                                           | Description                                                                               |
|---------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| [dropbox team teamfolder file list](dropbox-team-teamfolder-file-list.md)                         | Display all files and subfolders within team folders for content inventory and management |
| [dropbox team teamfolder file lock all release](dropbox-team-teamfolder-file-lock-all-release.md) | Release all file locks within a team folder path, resolving editing conflicts in bulk     |
| [dropbox team teamfolder file lock list](dropbox-team-teamfolder-file-lock-list.md)               | Display all locked files within team folders, identifying collaboration bottlenecks       |
| [dropbox team teamfolder file lock release](dropbox-team-teamfolder-file-lock-release.md)         | Release specific file locks in team folders to enable collaborative editing               |

# Activities log commands

The team activity log commands can export activity logs by certain filters or perspectives.

| Command                                                                   | Description                                                                                                                           |
|---------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team activity batch user](dropbox-team-activity-batch-user.md)   | Scan and retrieve activity logs for multiple team members in batch, useful for compliance auditing and user behavior analysis         |
| [dropbox team activity daily event](dropbox-team-activity-daily-event.md) | Generate daily activity reports showing team events grouped by date, helpful for tracking team usage patterns and security monitoring |
| [dropbox team activity event](dropbox-team-activity-event.md)             | Retrieve detailed team activity event logs with filtering options, essential for security auditing and compliance monitoring          |
| [dropbox team activity user](dropbox-team-activity-user.md)               | Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities                       |

# Connected applications and devices commands

The below commands can retrieve information about connected devices or applications in the team.

| Command                                                                   | Description                                                                                                                |
|---------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| [dropbox team device list](dropbox-team-device-list.md)                   | Display all devices and active sessions connected to team member accounts with device details and last activity timestamps |
| [dropbox team device unlink](dropbox-team-device-unlink.md)               | Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access       |
| [dropbox team linkedapp list](dropbox-team-linkedapp-list.md)             | Display all third-party applications linked to team member accounts for security auditing and access control               |
| [dropbox team backup device status](dropbox-team-backup-device-status.md) | Track Dropbox Backup status changes for all team devices over a specified period, monitoring backup health and compliance  |

# Other usecases

## External ID

External ID is the attribute that is not shown in any user interface of Dropbox. This attribute is for keeping a relationship between Dropbox and identity source (e.g. Active Directory, HR database) by identity management software such as Dropbox AD Connector. If you are using Dropbox AD Connector and you built a new Active Directory tree. You may need to clear existing external IDs to disconnect relationships with the old Active Directory tree and the new tree.\nIf you skip clear external IDs, Dropbox AD Connector may unintentionally delete accounts during configuring to the new tree.\nIf you want to see existing external IDs, use the `dropbox team member list` command. But the command will not include external ID by default. Please add the option `-experiment report_all_columns` like below.\n\n\n

| Command                                                                                       | Description                                                                                               |
|-----------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [dropbox team member list](dropbox-team-member-list.md)                                       | Display comprehensive list of all team members with their status, roles, and account details              |
| [dropbox team member clear externalid](dropbox-team-member-clear-externalid.md)               | Remove external ID mappings from team members, useful when disconnecting from identity management systems |
| [dropbox team member update batch externalid](dropbox-team-member-update-batch-externalid.md) | Set or update external IDs for multiple team members, integrating with identity management systems        |
| [dropbox team group list](dropbox-team-group-list.md)                                         | Display all groups in your team with member counts and group management types                             |
| [dropbox team group clear externalid](dropbox-team-group-clear-externalid.md)                 | Remove external ID mappings from groups, useful when disconnecting from external identity providers       |

## Data migration helper commands

Data migration helper commands copies member folders or team folders to another account or team. Please test before using those commands before actual data migration.

| Command                                                                                       | Description                                                                                                     |
|-----------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|
| [dropbox team member folder replication](dropbox-team-member-folder-replication.md)           | Copy folder contents from one team member to another's personal space, facilitating content transfer and backup |
| [dropbox team member replication](dropbox-team-member-replication.md)                         | Replicate all files from one team member's account to another, useful for account transitions or backups        |
| [dropbox team teamfolder partial replication](dropbox-team-teamfolder-partial-replication.md) | Selectively replicate team folder contents to another team, enabling flexible content migration                 |
| [dropbox team teamfolder replication](dropbox-team-teamfolder-replication.md)                 | Copy an entire team folder with all contents to another team account for migration or backup                    |

## Team info commands

| Command                                               | Description                                                                                                            |
|-------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [dropbox team feature](dropbox-team-feature.md)       | Display all features and capabilities enabled for your Dropbox team account, including API limits and special features |
| [dropbox team filesystem](dropbox-team-filesystem.md) | Identify whether your team uses legacy or modern file system architecture, important for feature compatibility         |
| [dropbox team info](dropbox-team-info.md)             | Display essential team account information including team ID and basic team settings                                   |

# Paper commands

## Legacy paper commands

Commands for a team's legacy Paper content. Please see the [official guide](https://developers.dropbox.com/paper-migration-guide) for more details about legacy Paper and migration

| Command                                                                               | Description                                                                                                                        |
|---------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team content legacypaper count](dropbox-team-content-legacypaper-count.md)   | Calculate the total number of legacy Paper documents owned by each team member, useful for content auditing and migration planning |
| [dropbox team content legacypaper list](dropbox-team-content-legacypaper-list.md)     | Generate a comprehensive list of all legacy Paper documents across the team with ownership and metadata information                |
| [dropbox team content legacypaper export](dropbox-team-content-legacypaper-export.md) | Export all legacy Paper documents from team members to local storage in HTML or Markdown format for backup or migration            |

# Team admin commands

Below commands are for managing team admins.

| Command                                                                         | Description                                                                                                                      |
|---------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|
| [dropbox team admin list](dropbox-team-admin-list.md)                           | Display all team members with their assigned admin roles, helpful for auditing administrative access and permissions             |
| [dropbox team admin role add](dropbox-team-admin-role-add.md)                   | Grant a specific admin role to an individual team member, enabling granular permission management                                |
| [dropbox team admin role clear](dropbox-team-admin-role-clear.md)               | Revoke all administrative privileges from a team member, useful for role transitions or security purposes                        |
| [dropbox team admin role delete](dropbox-team-admin-role-delete.md)             | Remove a specific admin role from a team member while preserving other roles, allowing precise permission adjustments            |
| [dropbox team admin role list](dropbox-team-admin-role-list.md)                 | Display all available admin roles in the team with their descriptions and permissions                                            |
| [dropbox team admin group role add](dropbox-team-admin-group-role-add.md)       | Assign admin roles to all members of a specified group, streamlining role management for large teams                             |
| [dropbox team admin group role delete](dropbox-team-admin-group-role-delete.md) | Remove admin roles from all team members except those in a specified exception group, useful for role cleanup and access control |

# Commands that run as a team member

You can run a command as a team member. For example, you can upload a file into member's folder by using `dropbox team runas file sync batch up`.

| Command                                                                                                       | Description                                                                                                  |
|---------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------|
| [dropbox team runas file list](dropbox-team-runas-file-list.md)                                               | List files and folders in a team member's account by running operations as that member                       |
| [dropbox team runas file batch copy](dropbox-team-runas-file-batch-copy.md)                                   | Copy multiple files or folders on behalf of team members, useful for content management and organization     |
| [dropbox team runas file sync batch up](dropbox-team-runas-file-sync-batch-up.md)                             | Upload multiple local files to team members' Dropbox accounts in batch, running as those members             |
| [dropbox team runas sharedfolder list](dropbox-team-runas-sharedfolder-list.md)                               | Display all shared folders accessible by a team member, running the operation as that member                 |
| [dropbox team runas sharedfolder isolate](dropbox-team-runas-sharedfolder-isolate.md)                         | Remove all shared folder access for a team member and transfer ownership, useful for departing employees     |
| [dropbox team runas sharedfolder mount add](dropbox-team-runas-sharedfolder-mount-add.md)                     | Mount shared folders to team members' accounts on their behalf, ensuring proper folder synchronization       |
| [dropbox team runas sharedfolder mount delete](dropbox-team-runas-sharedfolder-mount-delete.md)               | Unmount shared folders from team members' accounts on their behalf, managing folder synchronization          |
| [dropbox team runas sharedfolder mount list](dropbox-team-runas-sharedfolder-mount-list.md)                   | Display all shared folders currently mounted (synced) to a specific team member's account                    |
| [dropbox team runas sharedfolder mount mountable](dropbox-team-runas-sharedfolder-mount-mountable.md)         | Show all available shared folders that a team member can mount but hasn't synced yet                         |
| [dropbox team runas sharedfolder batch leave](dropbox-team-runas-sharedfolder-batch-leave.md)                 | Remove team members from multiple shared folders in batch by running leave operations as those members       |
| [dropbox team runas sharedfolder batch share](dropbox-team-runas-sharedfolder-batch-share.md)                 | Share multiple folders on behalf of team members in batch, automating folder sharing processes               |
| [dropbox team runas sharedfolder batch unshare](dropbox-team-runas-sharedfolder-batch-unshare.md)             | Remove sharing from multiple folders on behalf of team members, managing folder access in bulk               |
| [dropbox team runas sharedfolder member batch add](dropbox-team-runas-sharedfolder-member-batch-add.md)       | Add multiple members to shared folders in batch on behalf of folder owners, streamlining access management   |
| [dropbox team runas sharedfolder member batch delete](dropbox-team-runas-sharedfolder-member-batch-delete.md) | Remove multiple members from shared folders in batch on behalf of folder owners, managing access efficiently |

# Legal hold

With legal holds, admins can place a legal hold on members of their team and view and export all the content that's been created or modified by those members.

| Command                                                                                     | Description                                                                                               |
|---------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [dropbox team legalhold add](dropbox-team-legalhold-add.md)                                 | Create a legal hold policy to preserve specified team content for compliance or litigation purposes       |
| [dropbox team legalhold list](dropbox-team-legalhold-list.md)                               | Display all active legal hold policies with their details, members, and preservation status               |
| [dropbox team legalhold member batch update](dropbox-team-legalhold-member-batch-update.md) | Add or remove multiple team members from legal hold policies in batch for efficient compliance management |
| [dropbox team legalhold member list](dropbox-team-legalhold-member-list.md)                 | Display all team members currently under legal hold policies with their preservation status               |
| [dropbox team legalhold release](dropbox-team-legalhold-release.md)                         | Release a legal hold policy and restore normal file operations for affected members and content           |
| [dropbox team legalhold revision list](dropbox-team-legalhold-revision-list.md)             | Display all file revisions preserved under legal hold policies, ensuring comprehensive data retention     |
| [dropbox team legalhold update desc](dropbox-team-legalhold-update-desc.md)                 | Modify the description of an existing legal hold policy to reflect changes in scope or purpose            |
| [dropbox team legalhold update name](dropbox-team-legalhold-update-name.md)                 | Change the name of a legal hold policy for better identification and organization                         |

# Notes:

Dropbox Business footnote information

- [Specification changes](https://toolbox.watermint.org/guides/spec-change.md)
---
Title: Specification changes
URL: https://toolbox.watermint.org/guides/spec-change.md
---

# Spec Change Section

# Path change

Details about path changes in the spec

| CLI path (from) | CLI path (to) | Description of path change | Prune after build date |
|-----------------|---------------|----------------------------|------------------------|

# Prune change

Details about prune changes in the spec

| CLI path for prune                                                                     | Description of prune                 | Prune after build date |
|----------------------------------------------------------------------------------------|--------------------------------------|------------------------|
| [util text nlp english entity](https://github.com/watermint/toolbox/discussions/905)   | Split English text into entities     | 2025-07-30T15:00:00Z   |
| [util text nlp english sentence](https://github.com/watermint/toolbox/discussions/905) | Split English text into sentences    | 2025-07-30T15:00:00Z   |
| [util text nlp english token](https://github.com/watermint/toolbox/discussions/905)    | Split English text into tokens       | 2025-07-30T15:00:00Z   |
| [util text nlp japanese token](https://github.com/watermint/toolbox/discussions/905)   | Tokenize Japanese text               | 2025-07-30T15:00:00Z   |
| [util text nlp japanese wakati](https://github.com/watermint/toolbox/discussions/905)  | Wakachigaki (tokenize Japanese text) | 2025-07-30T15:00:00Z   |

- [Reporting options](https://toolbox.watermint.org/guides/reporting-options.md)
---
Title: Reporting options
URL: https://toolbox.watermint.org/guides/reporting-options.md
---

# Reporting options

The watermint toolbox creates reports from data obtained from services via the API. The reports differ depending on the command.

* Commands that output data from the API as reports.
* Commands that output data from the API as reports after processing.

The command output format is designed for core use cases. To avoid confusion, the command omits irrelevant/low priority fields.
You can output the abbreviated data using the method shown below, or create a report with your preferred format by using output filter.

# Hidden columns

In the CSV and xlsx reports produced by the command, some columns may be omitted. This includes, for example, internally used IDs and data of little relevance.

```
$ ./tbx dropbox file account info

watermint toolbox xxx.x.xxx
===========================

© 2016-2024 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

| email                 | email_verified | given_name | surname | display_name |
|-----------------------|----------------|------------|---------|--------------|
| xxx@xxxxxxxxxxxxx.xxx | true           | xxxx       | xxxx    | xxxxxxxx     |
```

If you want to output this type of data, you can add the `-experiment report_all_columns` option to output all defined columns.

```
$ ./tbx dropbox file account info -experiment report_all_columns

watermint toolbox xxx.x.xxx
===========================

© 2016-2024 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

| team_member_id                            | email                 | email_verified | status | given_name | surname | familiar_name | display_name | abbreviated_name | member_folder_id | external_id | account_id                               | persistent_id | joined_on | invited_on | role | tag |
|-------------------------------------------|-----------------------|----------------|--------|------------|---------|---------------|--------------|------------------|------------------|-------------|------------------------------------------|---------------|-----------|------------|------|-----|
| dbmid:xxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx | xxx@xxxxxxxxxxxxx.xxx | true           |        | xxxx       | xxxx    | xxxxxxxx      | xxxxxxx     | xxxx             |                  |             | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx |               |           |            |      |     |
```

Even if you use this option, some information may not be output. If you need more detailed information, try the following output filters.

# Output filter option

This feature allows you to filter the output of the command.
This is useful if you want to process the output in a specific format.

> NOTE: This filter is applied to all output reports. It will not work as intended for commands that output multiple report formats.
> When processing multiple report formats, please use the `util json query` command to process each of the output JSON files.
> 

In addition, in some cases, data in JSON format contains more data.
If you want to retrieve such hidden data, this option will help you to extract it as a report.

For example, the command [dropbox team member list](https://toolbox.watermint.org/commands/dropbox-team-member-list.html) returns a list of team members.
JSON report contains raw data from the Dropbox API.
If you want to extract only the email address and the verification status of the team members, you can use the output filter option.

```
$ ./tbx dropbox team member list -output json --output-filter "[.profile.email, .profile.email_verified]"
["sugito@example.com", true]
["kajiwara@example.com", true]
["takimoto@example.com", false]
["ueno@example.com", true]
["tomioka@example.com", false]
```

Then, if you want to format this data as a CSV, you can use the `@csv` filter like this (adding `| @csv` at the end):

```
$ ./tbx dropbox team member list -output json --output-filter "[.profile.email, .profile.email_verified] | @csv"
"sugito@example.com",true
"kajiwara@example.com",true
"takimoto@example.com",false
"ueno@example.com",true
"tomioka@example.com",false
```

In case you want to test the output filter, you can run the command first without the output filter option.\nThe command will generate the raw JSON output.\nThen, you can test the query with the command [util json query](https://toolbox.watermint.org/commands/util-json-query.html) to test the query.\n

- [Authentication Guide](https://toolbox.watermint.org/guides/troubleshooting.md)
---
Title: Authentication Guide
URL: https://toolbox.watermint.org/guides/troubleshooting.md
---

# Authentication Overview

Authentication Overview

The watermint toolbox requires proper authentication to access Dropbox services. The toolbox supports multiple authentication methods and securely manages tokens for seamless operation.

Key authentication concepts:
- OAuth 2.0 flow for secure authorization
- Token-based authentication for API access
- Secure token storage in local database
- Automatic token refresh when possible
- Support for multiple account configurations

# Dropbox Authentication

Dropbox Authentication

The toolbox uses OAuth 2.0 to authenticate with Dropbox. This process involves:

1. Initial Authorization:
   - Run any command that requires Dropbox access
   - The toolbox will open a browser window to Dropbox authorization page
   - Sign in to your Dropbox account and grant permissions
   - The toolbox receives an authorization code and exchanges it for access tokens

2. Supported Account Types:
   - Personal Dropbox accounts
   - Dropbox Business accounts (with team member access)
   - Dropbox Business admin accounts (with full team access)

3. Required Permissions:
   - File access permissions (read/write as needed)
   - Team information access (for business accounts)
   - User information access for account identification

4. Authentication Flow:
   - Commands automatically detect when authentication is needed
   - Browser-based OAuth flow ensures secure credential handling
   - No passwords or API keys need to be manually entered

# Token Management

Token Management

The toolbox securely manages authentication tokens:

1. Token Storage:
   - Tokens are stored in encrypted local database
   - Default location: $HOME/.toolbox/secrets/secrets.db
   - Custom database path can be specified with -auth-database flag

2. Token Lifecycle:
   - Access tokens are automatically refreshed when possible
   - Expired tokens trigger re-authentication flow
   - Tokens are associated with specific account configurations

3. Multiple Accounts:
   - Support for multiple Dropbox accounts
   - Each account maintains separate token storage
   - Account selection via command-line flags or configuration

4. Token Security:
   - Tokens are encrypted at rest
   - No tokens are logged or exposed in command output
   - Secure deletion when accounts are removed

5. Managing Tokens:
   - Use 'config auth list' to view configured accounts
   - Use 'config auth delete' to remove account configurations
   - Re-authentication is automatic when tokens are invalid

# Authentication Troubleshooting

Authentication Troubleshooting

Common authentication issues and solutions:

1. Browser Not Opening:
   - Check if you're running in a headless environment
   - Use -auto-open=false flag to disable automatic browser opening
   - Copy the authorization URL manually if needed

2. Permission Denied Errors:
   - Verify you have necessary permissions on your Dropbox account
   - For business accounts, ensure you have appropriate team member access
   - Re-authenticate if permissions have changed

3. Token Expired Errors:
   - Run the command again to trigger automatic re-authentication
   - Check if your account has been suspended or permissions revoked
   - Clear old tokens with 'config auth delete' if needed

4. Database Access Issues:
   - Ensure the secrets database directory is writable
   - Check file permissions on the database file
   - Use -auth-database flag to specify alternative location

5. Network Connectivity:
   - Verify internet connection for OAuth flow
   - Check if corporate firewalls are blocking access
   - Ensure access to dropbox.com and dropboxapi.com domains

6. Multiple Account Conflicts:
   - Use 'config auth list' to see all configured accounts
   - Remove conflicting accounts with 'config auth delete'
   - Specify explicit account in command if needed

# Security Best Practices

Security Best Practices

Follow these security practices when using authentication:

1. Token Protection:
   - Never share your secrets database file
   - Use appropriate file permissions on the database
   - Regularly review and clean up unused account configurations

2. Account Access:
   - Use dedicated service accounts for automation
   - Regularly review OAuth app authorizations in your Dropbox account
   - Revoke access for unused or suspicious applications

3. Environment Security:
   - Use secure workstations for authentication
   - Avoid authenticating on shared or public computers
   - Clear browser history after authentication if using public computers

4. Network Security:
   - Use secure networks for authentication
   - Avoid public WiFi for initial authentication
   - Consider using VPN for additional security

5. Monitoring:
   - Regularly review Dropbox account activity logs
   - Monitor for unexpected API usage
   - Set up alerts for unusual account activity

6. Backup and Recovery:
   - Keep secure backups of important data
   - Have a recovery plan if authentication is compromised
   - Know how to revoke and re-establish authentication

- [Error Handling Guide](https://toolbox.watermint.org/guides/troubleshooting.md)
---
Title: Error Handling Guide
URL: https://toolbox.watermint.org/guides/troubleshooting.md
---

# Common Errors and Solutions

Common Errors and Solutions

This section covers the most frequently encountered errors and their solutions:

1. Command Not Found:
   - Ensure you're using the correct command syntax
   - Use 'help' or '--help' to see available commands
   - Check for typos in command names

2. Invalid Arguments:
   - Verify required arguments are provided
   - Check argument formats (paths, emails, etc.)
   - Use quotes for arguments containing spaces

3. Configuration Issues:
   - Check if required configuration files exist
   - Verify configuration file permissions
   - Use default configurations when in doubt

4. Output Directory Issues:
   - Ensure output directories exist and are writable
   - Check disk space availability
   - Avoid paths with special characters

5. General Troubleshooting Steps:
   - Run with -debug flag for detailed logging
   - Check the command documentation
   - Verify your environment meets requirements
   - Try with minimal arguments first

# Network Errors

Network Errors

Network connectivity issues and solutions:

1. Connection Timeout:
   - Check internet connectivity
   - Verify DNS resolution for dropbox.com and dropboxapi.com
   - Try again after network issues are resolved
   - Consider using -bandwidth-kb flag to limit transfer speed

2. SSL/TLS Errors:
   - Update your system certificates
   - Check if corporate firewalls are interfering
   - Verify system time is accurate

3. Proxy Issues:
   - Configure system proxy settings if needed
   - Check proxy authentication requirements
   - Test direct connection when possible

4. DNS Resolution Failures:
   - Try using different DNS servers
   - Check /etc/hosts file for conflicts
   - Verify network configuration

5. Intermittent Connection Issues:
   - Implement retry logic with delays
   - Use smaller batch sizes for large operations
   - Monitor network stability during transfers

6. Corporate Network Restrictions:
   - Work with IT team to whitelist required domains
   - Request access to necessary ports (443, 80)
   - Consider using mobile hotspot for testing

# Authentication Errors

Authentication Errors

Authentication-related errors and solutions:

1. Token Expired/Invalid:
   - Run the command again to trigger re-authentication
   - Use 'config auth delete' to remove old credentials
   - Ensure system time is accurate

2. Permission Denied:
   - Verify account has necessary permissions
   - Check if account is suspended or restricted
   - For business accounts, ensure proper team member access

3. OAuth Flow Failures:
   - Try different browser or incognito mode
   - Clear browser cookies for dropbox.com
   - Check if ad blockers are interfering

4. Browser Not Opening:
   - Use -auto-open=false and copy URL manually
   - Check if running in headless environment
   - Verify default browser configuration

5. Multiple Account Conflicts:
   - Use 'config auth list' to see configured accounts
   - Remove conflicting accounts with 'config auth delete'
   - Specify account explicitly in commands

6. Database Access Issues:
   - Check permissions on secrets database file
   - Verify database directory is writable
   - Use -auth-database flag for custom location

# File System Errors

File System Errors

Local file system related errors and solutions:

1. Permission Denied:
   - Check file and directory permissions
   - Ensure user has read/write access as needed
   - Use sudo cautiously and only when necessary

2. Disk Space Issues:
   - Check available disk space with df -h
   - Clean up temporary files and logs
   - Use -budget-storage=low flag to reduce storage usage

3. Path Not Found:
   - Verify file and directory paths exist
   - Use absolute paths when relative paths fail
   - Check for typos in path names

4. File Lock Issues:
   - Close applications that might have files open
   - Check for running processes using files
   - Wait and retry if files are temporarily locked

5. Character Encoding Issues:
   - Ensure file names use valid character encoding
   - Avoid special characters in file names
   - Use UTF-8 encoding for text files

6. Symlink and Junction Issues:
   - Verify symlinks point to valid targets
   - Check permissions on symlink targets
   - Consider using direct paths instead of symlinks

7. Long Path Names:
   - Keep path lengths reasonable (< 260 characters on Windows)
   - Use shorter directory and file names
   - Move operations closer to root directory

# Rate Limit Errors

Rate Limit Errors

API rate limiting errors and solutions:

1. Too Many Requests (429 Error):
   - Wait before retrying (toolbox handles this automatically)
   - Reduce concurrency with -concurrency flag
   - Use smaller batch sizes for bulk operations

2. Daily/Hourly Limits:
   - Spread operations across multiple days
   - Monitor API usage patterns
   - Consider using multiple accounts for large operations

3. Burst Limit Exceeded:
   - Add delays between operations
   - Use batch operations when available
   - Avoid rapid sequential API calls

4. Team Rate Limits:
   - Coordinate with other team members using the API
   - Implement organization-wide rate limiting policies
   - Monitor team-wide API usage

5. Optimization Strategies:
   - Use list operations instead of individual file requests
   - Cache results to avoid repeated API calls
   - Combine multiple operations into single requests where possible

6. Monitoring and Planning:
   - Track API usage patterns
   - Plan large operations during off-peak hours
   - Set up alerts for approaching rate limits

# API Errors

API Errors

General API errors and solutions:

1. 400 Bad Request:
   - Verify request parameters are correct
   - Check data formats (dates, emails, paths)
   - Ensure required fields are provided

2. 401 Unauthorized:
   - Check authentication credentials
   - Verify token hasn't expired
   - Ensure proper authorization scope

3. 403 Forbidden:
   - Verify account permissions
   - Check if feature is available for account type
   - Ensure API access hasn't been restricted

4. 404 Not Found:
   - Verify file/folder paths exist
   - Check if items have been moved or deleted
   - Ensure proper path formatting

5. 409 Conflict:
   - Handle concurrent modification conflicts
   - Retry with updated information
   - Resolve conflicts manually if needed

6. 500 Internal Server Error:
   - Retry the operation after a delay
   - Check Dropbox status page for service issues
   - Contact support if error persists

7. Service Unavailable (503):
   - Wait and retry (temporary service issues)
   - Check for scheduled maintenance
   - Use exponential backoff for retries

# Debug Techniques

Debug Techniques

Advanced debugging techniques for troubleshooting:

1. Enable Debug Logging:
   - Use -debug flag for verbose output
   - Check log files for detailed information
   - Look for specific error messages and codes

2. Test with Minimal Parameters:
   - Start with simplest possible command
   - Add parameters one by one to isolate issues
   - Use default values when possible

3. Environment Verification:
   - Check system requirements
   - Verify environment variables
   - Test on different machines if available

4. Network Diagnostics:
   - Use ping/traceroute to test connectivity
   - Check firewall and proxy settings
   - Monitor network traffic during operations

5. API Testing:
   - Use API testing tools to verify endpoints
   - Check API responses manually
   - Verify request formats and parameters

6. Log Analysis:
   - Review application logs systematically
   - Look for patterns in error messages
   - Check timestamps for sequence of events

7. Isolation Testing:
   - Test with different accounts
   - Try operations on different files/folders
   - Use minimal test data sets

8. Community Resources:
   - Search documentation and FAQ
   - Check community forums and discussions
   - Report bugs with detailed reproduction steps

- [Best Practices Guide](https://toolbox.watermint.org/guides/troubleshooting.md)
---
Title: Best Practices Guide
URL: https://toolbox.watermint.org/guides/troubleshooting.md
---

# General Best Practices

General Best Practices

Follow these general guidelines for effective toolbox usage:

1. Command Preparation:
   - Always read command documentation before use
   - Test commands with sample data first
   - Use --help flag to understand available options
   - Verify required permissions and prerequisites

2. Data Backup:
   - Create backups before major operations
   - Test restore procedures regularly
   - Use version control for important files
   - Document backup and recovery procedures

3. Error Handling:
   - Enable debug logging for troubleshooting (-debug flag)
   - Keep logs of important operations
   - Implement proper error checking in scripts
   - Have rollback procedures for critical operations

4. Resource Management:
   - Monitor disk space before large operations
   - Use appropriate concurrency settings (-concurrency flag)
   - Manage memory usage with -budget-memory flag
   - Clean up temporary files and logs regularly

5. Documentation:
   - Document custom workflows and procedures
   - Keep track of configuration changes
   - Maintain inventory of automated scripts
   - Document troubleshooting steps for common issues

6. Testing:
   - Test commands in development environment first
   - Use small datasets for initial testing
   - Validate results before processing large batches
   - Implement automated testing for critical workflows

# Performance Optimization

Performance Optimization

Optimize toolbox performance with these techniques:

1. Concurrency Management:
   - Adjust -concurrency flag based on system resources
   - Higher concurrency for I/O intensive operations
   - Lower concurrency for CPU intensive operations
   - Monitor system resources during operations

2. Bandwidth Optimization:
   - Use -bandwidth-kb flag to limit network usage
   - Schedule large transfers during off-peak hours
   - Consider network conditions and limitations
   - Monitor transfer speeds and adjust accordingly

3. Memory Management:
   - Use -budget-memory=low for memory-constrained environments
   - Process data in smaller chunks for large datasets
   - Monitor memory usage during operations
   - Clear caches and temporary data regularly

4. Storage Optimization:
   - Use -budget-storage=low to reduce storage usage
   - Clean up logs and temporary files regularly
   - Use appropriate output formats (avoid verbose formats when not needed)
   - Compress data when possible

5. Batch Operations:
   - Group similar operations together
   - Use batch commands when available
   - Minimize API calls with efficient operations
   - Process multiple items in single commands

6. Caching Strategies:
   - Leverage local caching for frequently accessed data
   - Avoid redundant API calls
   - Use incremental operations when possible
   - Cache authentication tokens properly

7. Network Optimization:
   - Use stable, high-speed network connections
   - Avoid wireless connections for large transfers
   - Consider geographic proximity to servers
   - Implement retry logic with exponential backoff

# Security Best Practices

Security Best Practices

Maintain security while using the toolbox:

1. Authentication Security:
   - Use strong, unique passwords for accounts
   - Enable two-factor authentication when available
   - Regularly review and rotate credentials
   - Use dedicated service accounts for automation

2. Token Management:
   - Protect authentication database files
   - Use appropriate file permissions (600 or 700)
   - Avoid sharing authentication databases
   - Regularly audit configured accounts

3. Data Protection:
   - Encrypt sensitive data at rest and in transit
   - Use secure protocols (HTTPS, SSH) for all communications
   - Implement proper access controls
   - Regular security audits of data access

4. Environment Security:
   - Use secure workstations for operations
   - Keep systems updated with security patches
   - Use anti-virus and anti-malware protection
   - Secure physical access to systems

5. Network Security:
   - Use VPN for remote access
   - Avoid public WiFi for sensitive operations
   - Implement network segmentation where appropriate
   - Monitor network traffic for anomalies

6. Audit and Monitoring:
   - Log all significant operations
   - Monitor account activity regularly
   - Set up alerts for unusual activity
   - Maintain audit trails for compliance

7. Incident Response:
   - Have incident response procedures
   - Know how to revoke access quickly
   - Maintain contact information for security teams
   - Practice incident response scenarios

# Automation Best Practices

Automation Best Practices

Best practices for automating toolbox operations:

1. Script Development:
   - Use version control for all scripts
   - Implement proper error handling and logging
   - Add comments and documentation
   - Use configuration files for parameters

2. Scheduling and Execution:
   - Use cron jobs or task schedulers appropriately
   - Implement proper locking to prevent concurrent runs
   - Set up monitoring and alerting for failures
   - Use appropriate user accounts for automation

3. Parameter Management:
   - Use configuration files instead of hardcoded values
   - Implement parameter validation
   - Use environment variables for sensitive data
   - Provide default values where appropriate

4. Error Handling:
   - Implement comprehensive error checking
   - Use appropriate exit codes
   - Log errors with sufficient detail
   - Implement retry logic with backoff

5. Testing:
   - Test scripts in development environment
   - Use test data for validation
   - Implement automated testing where possible
   - Validate results automatically

6. Monitoring:
   - Log all significant operations
   - Monitor script execution times
   - Set up alerts for failures
   - Track resource usage

7. Maintenance:
   - Regular review and updates of scripts
   - Monitor for deprecated features
   - Keep dependencies updated
   - Document maintenance procedures

# Data Management Best Practices

Data Management Best Practices

Effective data management strategies:

1. Data Organization:
   - Use consistent naming conventions
   - Organize data in logical folder structures
   - Implement proper file and folder hierarchy
   - Use metadata and tags effectively

2. Data Validation:
   - Verify data integrity before and after operations
   - Use checksums for critical data
   - Implement data validation rules
   - Test with sample data before processing

3. Backup Strategies:
   - Implement regular automated backups
   - Test backup restoration procedures
   - Use multiple backup locations (3-2-1 rule)
   - Document backup and recovery procedures

4. Data Lifecycle Management:
   - Define data retention policies
   - Implement automated archiving
   - Clean up old and unnecessary data
   - Monitor storage usage trends

5. Data Synchronization:
   - Use incremental sync when possible
   - Verify sync operations regularly
   - Handle conflicts appropriately
   - Monitor sync performance and errors

6. Data Quality:
   - Implement data quality checks
   - Clean and normalize data regularly
   - Remove duplicates and inconsistencies
   - Validate data formats and standards

7. Compliance and Governance:
   - Follow data governance policies
   - Ensure compliance with regulations
   - Implement proper access controls
   - Maintain audit trails for data operations

# Team Collaboration Best Practices

Team Collaboration Best Practices

Effective team collaboration with toolbox:

1. Account Management:
   - Use dedicated service accounts for shared operations
   - Implement proper access controls and permissions
   - Regular review of account access and privileges
   - Document account usage and responsibilities

2. Configuration Management:
   - Use centralized configuration management
   - Version control for shared configurations
   - Implement configuration validation
   - Document configuration changes

3. Workflow Coordination:
   - Define clear workflows and procedures
   - Implement proper change management
   - Use communication channels for coordination
   - Schedule operations to avoid conflicts

4. Knowledge Sharing:
   - Document procedures and best practices
   - Conduct regular training sessions
   - Share troubleshooting experiences
   - Maintain knowledge base and FAQ

5. Quality Assurance:
   - Implement peer review processes
   - Use testing and validation procedures
   - Define quality standards and metrics
   - Regular audits of processes and results

6. Communication:
   - Establish clear communication protocols
   - Use collaboration tools effectively
   - Document decisions and changes
   - Regular team meetings and updates

7. Incident Management:
   - Define incident response procedures
   - Establish escalation paths
   - Maintain contact information
   - Conduct post-incident reviews

# Maintenance and Updates

Maintenance and Updates

Keep your toolbox installation and workflows updated:

1. Software Updates:
   - Regularly check for toolbox updates
   - Test updates in development environment first
   - Keep dependencies updated
   - Monitor for security updates

2. Configuration Maintenance:
   - Regular review of configurations
   - Update deprecated settings
   - Optimize performance settings
   - Clean up unused configurations

3. Data Maintenance:
   - Regular cleanup of logs and temporary files
   - Archive old data appropriately
   - Optimize storage usage
   - Validate data integrity periodically

4. Documentation Updates:
   - Keep documentation current with changes
   - Update procedures and workflows
   - Review and update troubleshooting guides
   - Maintain version history

5. Performance Monitoring:
   - Monitor system performance regularly
   - Track operation times and success rates
   - Identify and address bottlenecks
   - Optimize based on usage patterns

6. Security Maintenance:
   - Regular security audits
   - Update security configurations
   - Review access permissions
   - Monitor for security vulnerabilities

7. Backup and Recovery:
   - Test backup and recovery procedures
   - Update recovery documentation
   - Verify backup integrity
   - Practice disaster recovery scenarios

8. Training and Skills:
   - Keep team skills current
   - Provide training on new features
   - Share knowledge and best practices
   - Encourage continuous learning

- [Guide to advanced reporting features including hidden columns, output filtering, and JSON processing](https://toolbox.watermint.org/guides/reporting-options.md)
---
Title: Guide to advanced reporting features including hidden columns, output filtering, and JSON processing
URL: https://toolbox.watermint.org/guides/reporting-options.md
---

# Reporting Guide

# Advanced Reporting Features

Many commands in watermint toolbox generate reports with tabular data. There are three powerful features that can extend and modify command output for advanced reporting needs:

## 1. Show All Columns with -experimental report_all_columns

By default, commands display only the most commonly used columns to keep output readable. However, many commands have additional hidden columns that can be revealed using the experimental flag.

### Usage
```bash
go run . [command] -experimental report_all_columns
```

### Example: Dropbox Team Member List

**Standard output (limited columns):**
```bash
go run . dropbox team member list
```
Shows: Email, Status, Role, etc.

**Extended output (all columns):**
```bash
go run . dropbox team member list -experimental report_all_columns
```
Shows: Email, Status, Role, Team Member ID, Account ID, External ID, Profile, Date Joined, Groups, etc.

This reveals internal IDs and metadata that are often needed for automation and integration scenarios.

## 2. Filter Output with -output-filter

The output filter option allows you to run queries on command results using a SQL-like syntax. This requires `-output json` to be specified.

### Usage
```bash
go run . [command] -output json -output-filter "query expression"
```

### Query Syntax
- `select`: Choose columns to display
- `where`: Filter rows based on conditions
- `order by`: Sort results
- `limit`: Limit number of results

### Examples

**Filter active team members:**
```bash
go run . dropbox team member list -output json -output-filter "select email, status where status = 'active'"
```

**Find members by domain:**
```bash
go run . dropbox team member list -output json -output-filter "select email, team_member_id where email like '%@company.com'"
```

**Get top 10 largest files:**
```bash
go run . dropbox file list -output json -output-filter "select name, size order by size desc limit 10"
```

## 3. Post-Process with util json query

The `util json query` command provides jq-like functionality for processing JSON output from other commands. It's particularly useful for complex data transformations.

### Usage
```bash
go run . [command] -output json | go run . util json query -query "jq expression"
```

### Common Patterns

**Extract specific fields:**
```bash
go run . dropbox team member list -output json | go run . util json query -query ".[] | {email, team_member_id}"
```

**Group and count:**
```bash
go run . dropbox team member list -output json | go run . util json query -query "group_by(.status) | map({status: .[0].status, count: length})"
```

**Filter and transform:**
```bash
go run . dropbox file list -output json | go run . util json query -query ".[] | select(.size > 1000000) | .name"
```

## Complete Example: Team Member ID and Email Report

A common administrative task is to generate a report of all team member IDs with email addresses. Here's how to accomplish this using the three features:

### Step 1: Get all columns including team_member_id
```bash
go run . dropbox team member list -experimental report_all_columns -output json
```

### Step 2: Filter to get only team_member_id and email
```bash
go run . dropbox team member list -experimental report_all_columns -output json -output-filter "select team_member_id, email"
```

### Step 3: Alternative using util json query
```bash
go run . dropbox team member list -experimental report_all_columns -output json | go run . util json query -query ".[] | {team_member_id, email}"
```

### Step 4: Export to CSV for spreadsheet use
```bash
go run . dropbox team member list -experimental report_all_columns -output csv -output-filter "select team_member_id, email" > team_members.csv
```

## Best Practices

### 1. Discover Available Columns
Always start with `-experimental report_all_columns` to see what data is available:
```bash
go run . [command] -experimental report_all_columns | head -5
```

### 2. Use Output Filter for Simple Queries
For straightforward filtering and column selection, `-output-filter` is more efficient:
```bash
# Good for simple cases
go run . dropbox team member list -output json -output-filter "select email where status = 'active'"
```

### 3. Use util json query for Complex Processing
For complex transformations, aggregations, or when you need jq's full power:
```bash
# Good for complex cases
go run . dropbox team member list -output json | go run . util json query -query "group_by(.role) | map({role: .[0].role, members: map(.email)})"
```

### 4. Combine with Standard Tools
These features integrate well with standard Unix tools:
```bash
# Count active members
go run . dropbox team member list -output json -output-filter "select email where status = 'active'" | jq length

# Sort and save
go run . dropbox team member list -output csv -output-filter "select email, team_member_id order by email" > sorted_members.csv
```

## Performance Considerations

- `-experimental report_all_columns` may return more data and take slightly longer
- `-output-filter` is processed server-side and is generally faster for large datasets
- `util json query` processes data client-side and is better for complex transformations
- For very large datasets, consider using `-output-filter` to reduce data transfer

## Troubleshooting

### Common Issues

**Column not found in output-filter:**
```bash
# First check available columns
go run . [command] -experimental report_all_columns | head -1
```

**JSON parsing errors:**
```bash
# Ensure -output json is specified
go run . [command] -output json -output-filter "..."
```

**Complex jq expressions:**
```bash
# Test expressions incrementally
go run . [command] -output json | go run . util json query -query ".[0]"  # First record
go run . [command] -output json | go run . util json query -query ".[] | keys"  # Available fields
```

These advanced reporting features provide powerful ways to extract, filter, and transform data from watermint toolbox commands, enabling sophisticated reporting and automation workflows.
