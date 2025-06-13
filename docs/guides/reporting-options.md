---
layout: page
title: Reporting options
lang: en
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


