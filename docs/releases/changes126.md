---
layout: release
title: Changes of Release 125
lang: en
---

# Changes between `Release 125` to `Release 126`

# Commands added


| Command                    | Title                                                   |
|----------------------------|---------------------------------------------------------|
| config feature disable     | Disable a feature.                                      |
| config feature enable      | Enable a feature.                                       |
| config feature list        | List available optional features.                       |
| dev info                   | Dev information                                         |
| dev placeholder pathchange | Placeholder command for path change document generation |
| dev placeholder prune      | Placeholder of prune workflow messages                  |
| log cat job                | Retrieve logs of specified Job ID                       |
| log cat kind               | Concatenate and print logs of specified log kind        |
| log cat last               | Print the last job log files                            |
| log job archive            | Archive jobs                                            |
| log job delete             | Delete old job history                                  |
| log job list               | Show job history                                        |
| log job ship               | Ship Job logs to Dropbox path                           |



# Commands deleted


| Command             | Title                                            |
|---------------------|--------------------------------------------------|
| config disable      | Disable a feature.                               |
| config enable       | Enable a feature.                                |
| config features     | List available optional features.                |
| job history archive | Archive jobs                                     |
| job history delete  | Delete old job history                           |
| job history list    | Show job history                                 |
| job history ship    | Ship Job logs to Dropbox path                    |
| job log jobid       | Retrieve logs of specified Job ID                |
| job log kind        | Concatenate and print logs of specified log kind |
| job log last        | Print the last job log files                     |



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
