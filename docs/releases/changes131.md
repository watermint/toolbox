---
layout: release
title: Changes of Release 130
lang: en
---

# Changes between `Release 130` to `Release 131`

# Commands added


| Command                             | Title                                                               |
|-------------------------------------|---------------------------------------------------------------------|
| dropbox sign request list           | List signature requests                                             |
| dropbox sign request signature list | List signatures of requests                                         |
| log api job                         | Show statistics of the API log of the job specified by the job ID   |
| log api name                        | Show statistics of the API log of the job specified by the job name |



# Commands deleted


| Command      | Title                         |
|--------------|-------------------------------|
| log job ship | Ship Job logs to Dropbox path |



# Command spec changed: `log cat job`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Id",
  			Desc:     "Job ID",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Kind", Desc: "Kind of log", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
