---
layout: release
title: Changes of Release 104
lang: en
---

# Changes between `Release 104` to `Release 105`

# Commands added

| Command                                 | Title                                                   |
|-----------------------------------------|---------------------------------------------------------|
| team runas sharedfolder batch leave     | Batch leave from shared folders as a member             |
| team runas sharedfolder list            | List shared folders run as the member                   |
| team runas sharedfolder mount add       | Add the shared folder to the specified member's Dropbox |
| team runas sharedfolder mount delete    | The specified user unmounts the designated folder.      |
| team runas sharedfolder mount list      | List all shared folders the specified member mounted    |
| team runas sharedfolder mount mountable | List all shared folders the member can mount            |

# Command spec changed: `dev benchmark upload`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BlockBlockSize",
  			Desc:     "Block size for batch upload",
- 			Default:  "40",
+ 			Default:  "16",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(40),
+ 				"value": float64(16),
  			},
  		},
  		&{Name: "Method", Desc: "Upload method", Default: "block", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...},
  		... // 7 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas file list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas file list",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST",
+ 	CliArgs:         "-member-email MEMBER@DOMAIN -path /DROPBOX/PATH/TO/LIST",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
