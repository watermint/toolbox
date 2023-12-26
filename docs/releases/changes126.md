---
layout: release
title: Changes of Release 125
lang: en
---

# Changes between `Release 125` to `Release 126`

# Commands added


| Command                            | Title                                                   |
|------------------------------------|---------------------------------------------------------|
| asana team list                    | List team                                               |
| asana team project list            | List projects of the team                               |
| asana team task list               | List task of the team                                   |
| asana workspace list               | List workspaces                                         |
| asana workspace project list       | List projects of the workspace                          |
| config feature disable             | Disable a feature.                                      |
| config feature enable              | Enable a feature.                                       |
| config feature list                | List available optional features.                       |
| deepl translate text               | Translate text                                          |
| dev info                           | Dev information                                         |
| dev lifecycle planchangepath       | Add plan of changing path to commands                   |
| dev lifecycle planprune            | Add plan of the command discontinuation                 |
| dev placeholder pathchange         | Placeholder command for path change document generation |
| dev placeholder prune              | Placeholder of prune workflow messages                  |
| dropbox file account feature       | List Dropbox account features                           |
| dropbox file account filesystem    | Show Dropbox file system version                        |
| dropbox file account info          | Dropbox account info                                    |
| dropbox sign account info          | Show Dropbox Sign account information                   |
| figma account info                 | Retrieve current user information                       |
| figma file export all page         | Export all files/pages under the team                   |
| figma file export frame            | Export all frames of the Figma file                     |
| figma file export node             | Export Figma document Node                              |
| figma file export page             | Export all pages of the Figma file                      |
| figma file info                    | Show information of the figma file                      |
| figma file list                    | List files in the Figma Project                         |
| figma project list                 | List projects of the team                               |
| github content get                 | Get content metadata of the repository                  |
| github content put                 | Put small text content into the repository              |
| github issue list                  | List issues of the public/private GitHub repository     |
| github profile                     | Get the authenticated user                              |
| github release asset download      | Download assets                                         |
| github release asset list          | List assets of GitHub Release                           |
| github release asset upload        | Upload assets file into the GitHub Release              |
| github release draft               | Create release draft                                    |
| github release list                | List releases                                           |
| github tag create                  | Create a tag on the repository                          |
| google calendar event list         | List Google Calendar events                             |
| google mail filter add             | Add a filter.                                           |
| google mail filter batch add       | Batch adding/deleting labels with query                 |
| google mail filter delete          | Delete a filter                                         |
| google mail filter list            | List filters                                            |
| google mail label add              | Add a label                                             |
| google mail label delete           | Delete a label                                          |
| google mail label list             | List email labels                                       |
| google mail label rename           | Rename a label                                          |
| google mail message label add      | Add labels to the message                               |
| google mail message label delete   | Remove labels from the message                          |
| google mail message list           | List messages                                           |
| google mail message processed list | List messages in processed format.                      |
| google mail message send           | Send a mail                                             |
| google mail sendas add             | Creates a custom "from" send-as alias                   |
| google mail sendas delete          | Deletes the specified send-as alias                     |
| google mail sendas list            | Lists the send-as aliases for the specified account     |
| google mail thread list            | List threads                                            |
| google sheets sheet append         | Append data to a spreadsheet                            |
| google sheets sheet clear          | Clears values from a spreadsheet                        |
| google sheets sheet create         | Create a new sheet                                      |
| google sheets sheet delete         | Delete a sheet from the spreadsheet                     |
| google sheets sheet export         | Export sheet data                                       |
| google sheets sheet import         | Import data into the spreadsheet                        |
| google sheets sheet list           | List sheets of the spreadsheet                          |
| google sheets spreadsheet create   | Create a new spreadsheet                                |
| google translate text              | Translate text                                          |
| log cat job                        | Retrieve logs of specified Job ID                       |
| log cat kind                       | Concatenate and print logs of specified log kind        |
| log cat last                       | Print the last job log files                            |
| log job archive                    | Archive jobs                                            |
| log job delete                     | Delete old job history                                  |
| log job list                       | Show job history                                        |
| log job ship                       | Ship Job logs to Dropbox path                           |
| slack conversation history         | Conversation history                                    |
| slack conversation list            | List channels                                           |



# Commands deleted


| Command                                     | Title                                               |
|---------------------------------------------|-----------------------------------------------------|
| config disable                              | Disable a feature.                                  |
| config enable                               | Enable a feature.                                   |
| config features                             | List available optional features.                   |
| job history archive                         | Archive jobs                                        |
| job history delete                          | Delete old job history                              |
| job history list                            | Show job history                                    |
| job history ship                            | Ship Job logs to Dropbox path                       |
| job log jobid                               | Retrieve logs of specified Job ID                   |
| job log kind                                | Concatenate and print logs of specified log kind    |
| job log last                                | Print the last job log files                        |
| services asana team list                    | List team                                           |
| services asana team project list            | List projects of the team                           |
| services asana team task list               | List task of the team                               |
| services asana workspace list               | List workspaces                                     |
| services asana workspace project list       | List projects of the workspace                      |
| services deepl translate text               | Translate text                                      |
| services dropbox user feature               | List feature settings for current user              |
| services dropbox user filesystem            | Identify user's team file system version            |
| services dropbox user info                  | Retrieve current account info                       |
| services dropboxsign account info           | Retrieve account information                        |
| services figma account info                 | Retrieve current user information                   |
| services figma file export all page         | Export all files/pages under the team               |
| services figma file export frame            | Export all frames of the Figma file                 |
| services figma file export node             | Export Figma document Node                          |
| services figma file export page             | Export all pages of the Figma file                  |
| services figma file info                    | Show information of the figma file                  |
| services figma file list                    | List files in the Figma Project                     |
| services figma project list                 | List projects of the team                           |
| services github content get                 | Get content metadata of the repository              |
| services github content put                 | Put small text content into the repository          |
| services github issue list                  | List issues of the public/private GitHub repository |
| services github profile                     | Get the authenticated user                          |
| services github release asset download      | Download assets                                     |
| services github release asset list          | List assets of GitHub Release                       |
| services github release asset upload        | Upload assets file into the GitHub Release          |
| services github release draft               | Create release draft                                |
| services github release list                | List releases                                       |
| services github tag create                  | Create a tag on the repository                      |
| services google calendar event list         | List Google Calendar events                         |
| services google mail filter add             | Add a filter.                                       |
| services google mail filter batch add       | Batch adding/deleting labels with query             |
| services google mail filter delete          | Delete a filter                                     |
| services google mail filter list            | List filters                                        |
| services google mail label add              | Add a label                                         |
| services google mail label delete           | Delete a label                                      |
| services google mail label list             | List email labels                                   |
| services google mail label rename           | Rename a label                                      |
| services google mail message label add      | Add labels to the message                           |
| services google mail message label delete   | Remove labels from the message                      |
| services google mail message list           | List messages                                       |
| services google mail message processed list | List messages in processed format.                  |
| services google mail message send           | Send a mail                                         |
| services google mail sendas add             | Creates a custom "from" send-as alias               |
| services google mail sendas delete          | Deletes the specified send-as alias                 |
| services google mail sendas list            | Lists the send-as aliases for the specified account |
| services google mail thread list            | List threads                                        |
| services google sheets sheet append         | Append data to a spreadsheet                        |
| services google sheets sheet clear          | Clears values from a spreadsheet                    |
| services google sheets sheet create         | Create a new sheet                                  |
| services google sheets sheet delete         | Delete a sheet from the spreadsheet                 |
| services google sheets sheet export         | Export sheet data                                   |
| services google sheets sheet import         | Import data into the spreadsheet                    |
| services google sheets sheet list           | List sheets of the spreadsheet                      |
| services google sheets spreadsheet create   | Create a new spreadsheet                            |
| services google translate text              | Translate text                                      |
| services slack conversation history         | Conversation history                                |
| services slack conversation list            | List channels                                       |



# Command spec changed: `dev benchmark local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev benchmark local",
  	CliArgs: strings.Join({
  		"-num-files NUM -path /LOCAL/PATH/TO/PROCESS -size-max-kb NUM -si",
  		"ze-min-kb NUM",
- 		`"`,
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 16 identical fields
  }
```
# Command spec changed: `dev build catalogue`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Importer",
+ 			Desc:     "Importer type",
+ 			Default:  "default",
+ 			TypeName: "essentials.model.mo_string.select_string_internal",
+ 			TypeAttr: map[string]any{"options": []any{string("default"), string("enhanced")}},
+ 		},
+ 	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "detach",
  	Title: strings.Join({
  		"Convert Dropbox ",
- 		"Busines",
+ 		"for team",
  		"s accounts to a Basic account",
  	}, ""),
  	Desc:    "",
  	Remarks: "(Irreversible operation)",
  	... // 20 identical fields
  }
```
# Command spec changed: `team activity batch user`



## Changed report: combined

```
  &dc_recipe.Report{
  	Name: "combined",
  	Desc: strings.Join({
  		"This report shows an activity logs with mostly compatible with D",
  		"ropbox ",
- 		"Busines",
+ 		"for team",
  		"s's activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```

## Changed report: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: strings.Join({
  		"This report shows an activity logs with mostly compatible with D",
  		"ropbox ",
- 		"Busines",
+ 		"for team",
  		"s's activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `team activity daily event`



## Changed report: event

```
  &dc_recipe.Report{
  	Name: "event",
  	Desc: strings.Join({
  		"This report shows an activity logs with mostly compatible with D",
  		"ropbox ",
- 		"Busines",
+ 		"for team",
  		"s's activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `team activity event`



## Changed report: event

```
  &dc_recipe.Report{
  	Name: "event",
  	Desc: strings.Join({
  		"This report shows an activity logs with mostly compatible with D",
  		"ropbox ",
- 		"Busines",
+ 		"for team",
  		"s's activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `team activity user`



## Changed report: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: strings.Join({
  		"This report shows an activity logs with mostly compatible with D",
  		"ropbox ",
- 		"Busines",
+ 		"for team",
  		"s's activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
