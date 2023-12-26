---
layout: page
title: Specification changes
lang: en
---

# Specification changes

# Command path changes

If you continue to use your current version, this will not affect you, but changes will be applied in future versions. If a date is specified, the change is applied to versions released after that date.

| Former path                                 | Current path                       | Command description                                 | Date                 |
|---------------------------------------------|------------------------------------|-----------------------------------------------------|----------------------|
| config disable                              | config feature disable             | Disable a feature.                                  | 2024-04-01T00:00:00Z |
| config enable                               | config feature enable              | Enable a feature.                                   | 2024-04-01T00:00:00Z |
| config features                             | config feature list                | List available optional features.                   | 2024-04-01T00:00:00Z |
| services deepl translate text               | deepl translate text               | Translate text                                      | 2024-06-30T15:00:00Z |
| services dropbox user feature               | dropbox file account feature       | List Dropbox account features                       | 2024-06-30T15:00:00Z |
| services dropbox user filesystem            | dropbox file account filesystem    | Show Dropbox file system version                    | 2024-06-30T15:00:00Z |
| services dropbox user info                  | dropbox file account info          | Dropbox account info                                | 2024-06-30T15:00:00Z |
| services figma account info                 | figma account info                 | Retrieve current user information                   | 2024-06-30T15:00:00Z |
| services figma file export all page         | figma file export all page         | Export all files/pages under the team               | 2024-06-30T15:00:00Z |
| services figma file export frame            | figma file export frame            | Export all frames of the Figma file                 | 2024-06-30T15:00:00Z |
| services figma file export node             | figma file export node             | Export Figma document Node                          | 2024-06-30T15:00:00Z |
| services figma file export page             | figma file export page             | Export all pages of the Figma file                  | 2024-06-30T15:00:00Z |
| services figma file info                    | figma file info                    | Show information of the figma file                  | 2024-06-30T15:00:00Z |
| services figma file list                    | figma file list                    | List files in the Figma Project                     | 2024-06-30T15:00:00Z |
| services figma project list                 | figma project list                 | List projects of the team                           | 2024-06-30T15:00:00Z |
| services google calendar event list         | google calendar event list         | List Google Calendar events                         | 2024-06-30T15:00:00Z |
| services google mail filter add             | google mail filter add             | Add a filter.                                       | 2024-06-30T15:00:00Z |
| services google mail filter batch add       | google mail filter batch add       | Batch adding/deleting labels with query             | 2024-06-30T15:00:00Z |
| services google mail filter delete          | google mail filter delete          | Delete a filter                                     | 2024-06-30T15:00:00Z |
| services google mail filter list            | google mail filter list            | List filters                                        | 2024-06-30T15:00:00Z |
| services google mail label add              | google mail label add              | Add a label                                         | 2024-06-30T15:00:00Z |
| services google mail label delete           | google mail label delete           | Delete a label                                      | 2024-06-30T15:00:00Z |
| services google mail label list             | google mail label list             | List email labels                                   | 2024-06-30T15:00:00Z |
| services google mail label rename           | google mail label rename           | Rename a label                                      | 2024-06-30T15:00:00Z |
| services google mail message label add      | google mail message label add      | Add labels to the message                           | 2024-06-30T15:00:00Z |
| services google mail message label delete   | google mail message label delete   | Remove labels from the message                      | 2024-06-30T15:00:00Z |
| services google mail message list           | google mail message list           | List messages                                       | 2024-06-30T15:00:00Z |
| services google mail message processed list | google mail message processed list | List messages in processed format.                  | 2024-06-30T15:00:00Z |
| services google mail sendas add             | google mail sendas add             | Creates a custom "from" send-as alias               | 2024-06-30T15:00:00Z |
| services google mail sendas delete          | google mail sendas delete          | Deletes the specified send-as alias                 | 2024-06-30T15:00:00Z |
| services google mail sendas list            | google mail sendas list            | Lists the send-as aliases for the specified account | 2024-06-30T15:00:00Z |
| services google mail thread list            | google mail thread list            | List threads                                        | 2024-06-30T15:00:00Z |
| services google sheets sheet append         | google sheets sheet append         | Append data to a spreadsheet                        | 2024-06-30T15:00:00Z |
| services google sheets sheet clear          | google sheets sheet clear          | Clears values from a spreadsheet                    | 2024-06-30T15:00:00Z |
| services google sheets sheet create         | google sheets sheet create         | Create a new sheet                                  | 2024-06-30T15:00:00Z |
| services google sheets sheet delete         | google sheets sheet delete         | Delete a sheet from the spreadsheet                 | 2024-06-30T15:00:00Z |
| services google sheets sheet export         | google sheets sheet export         | Export sheet data                                   | 2024-06-30T15:00:00Z |
| services google sheets sheet import         | google sheets sheet import         | Import data into the spreadsheet                    | 2024-06-30T15:00:00Z |
| services google sheets sheet list           | google sheets sheet list           | List sheets of the spreadsheet                      | 2024-06-30T15:00:00Z |
| services google sheets spreadsheet create   | google sheets spreadsheet create   | Create a new spreadsheet                            | 2024-06-30T15:00:00Z |
| job log jobid                               | log cat job                        | Retrieve logs of specified Job ID                   | 2024-04-01T00:00:00Z |
| job log kind                                | log cat kind                       | Concatenate and print logs of specified log kind    | 2024-04-01T00:00:00Z |
| job log last                                | log cat last                       | Print the last job log files                        | 2024-04-01T00:00:00Z |
| job history archive                         | log job archive                    | Archive jobs                                        | 2024-04-01T00:00:00Z |
| job history delete                          | log job delete                     | Delete old job history                              | 2024-04-01T00:00:00Z |
| job history list                            | log job list                       | Show job history                                    | 2024-04-01T00:00:00Z |

# Deprecation

Below commands will be removed in the future release. If you continue to use your current version, this will not affect you, but changes will be applied in future versions. If a date is specified, the change is applied to versions released after that date.

| Path                                | Command description                                   | Date                 |
|-------------------------------------|-------------------------------------------------------|----------------------|
| log job ship                        | Ship Job logs to Dropbox path                         | 2024-02-01T00:00:00Z |
| teamspace asadmin file list         | List files and folders in team space run as admin     | 2024-07-01T00:00:00Z |
| teamspace asadmin folder add        | Create top level folder in the team space             | 2024-07-01T00:00:00Z |
| teamspace asadmin folder delete     | Delete top level folder of the team space             | 2024-07-01T00:00:00Z |
| teamspace asadmin folder permdelete | Permanently delete top level folder of the team space | 2024-07-01T00:00:00Z |
| teamspace file list                 | List files and folders in team space                  | 2024-07-01T00:00:00Z |


