# Building watermint toolbox

# Prerequisites

* Register your own application on Dropbox's developer site for each required application type.
* Place application keys into the file.

## Desktop development environment

Place JSON file named `toolbox.appkeys` under `resources` folder, then run or compile binaries.
`toolbox.appkeys` file format is like below:

```JSON
{
  "user_full.key": "xxxxxxxxxxxxxx",
  "user_full.secret": "xxxxxxxxxxxxxx",
  "business_info.key": "xxxxxxxxxxxxxx",
  "business_info.secret": "xxxxxxxxxxxxxx",
  "business_file.key": "xxxxxxxxxxxxxx",
  "business_file.secret": "xxxxxxxxxxxxxx",
  "business_management.key": "xxxxxxxxxxxxxx",
  "business_management.secret": "xxxxxxxxxxxxxx",
  "business_audit.key": "xxxxxxxxxxxxxx",
  "business_audit.secret": "xxxxxxxxxxxxxx"
}
```

## On CI environment (CircleCI)

Configure environment variable `TOOLBOX_APPKEYS` like below

```
{"user_full.key":"xxxxxxxxxxxxxx","user_full.secret":"xxxxxxxxxxxxxx","business_info.key":"xxxxxxxxxxxxxx","business_info.secret":"xxxxxxxxxxxxxx","business_file.key":"xxxxxxxxxxxxxx","business_file.secret":"xxxxxxxxxxxxxx","business_management.key":"xxxxxxxxxxxxxx","business_management.secret":"xxxxxxxxxxxxxx","business_audit.key":"xxxxxxxxxxxxxx","business_audit.secret":"xxxxxxxxxxxxxx"}
```

# Docker build

To build an executable, please use Docker like below.

```bash
$ docker build -t toolbox . && rm -fr /tmp/dist && docker run -v /tmp/dist:/dist:rw --rm toolbox
```
