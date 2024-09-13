---
layout: release
title: Changes of Release 136
lang: en
---

# Changes between `Release 136` to `Release 137`

# Commands added


| Command          | Title                                  |
|------------------|----------------------------------------|
| dev doc markdown | Generate messages from markdown source |



# Command spec changed: `util json query`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Compact", Desc: "Compact output", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "Lines",
- 			Desc:     "Read JSON Lines (https://jsonlines.org/) format",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "Path", Desc: "File path", TypeName: "Path"},
  		&{Name: "Query", Desc: "Query string. ", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
