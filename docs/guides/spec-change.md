---
layout: page
title: Specification changes
lang: en
---

# Specification changes

# Command path changes

If you continue to use your current version, this will not affect you, but changes will be applied in future versions. If a date is specified, the change is applied to versions released after that date.

| Former path         | Current path           | Command description                              | Date                 |
|---------------------|------------------------|--------------------------------------------------|----------------------|
| job history archive | log job archive        | Archive jobs                                     | 2024-04-01T00:00:00Z |
| job history delete  | log job delete         | Delete old job history                           | 2024-04-01T00:00:00Z |
| job history list    | log job list           | Show job history                                 | 2024-04-01T00:00:00Z |
| job log jobid       | log cat job            | Retrieve logs of specified Job ID                | 2024-04-01T00:00:00Z |
| job log kind        | log cat kind           | Concatenate and print logs of specified log kind | 2024-04-01T00:00:00Z |
| job log last        | log cat last           | Print the last job log files                     | 2024-04-01T00:00:00Z |
| config disable      | config feature disable | Disable a feature.                               | 2024-04-01T00:00:00Z |
| config enable       | config feature enable  | Enable a feature.                                | 2024-04-01T00:00:00Z |
| config features     | config feature list    | List available optional features.                | 2024-04-01T00:00:00Z |

# Deprecation

Below commands will be removed in the future release. If you continue to use your current version, this will not affect you, but changes will be applied in future versions. If a date is specified, the change is applied to versions released after that date.

| Path         | Command description           | Date                 |
|--------------|-------------------------------|----------------------|
| log job ship | Ship Job logs to Dropbox path | 2024-01-01T00:00:00Z |


