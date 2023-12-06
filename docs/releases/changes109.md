---
layout: release
title: Changes of Release 108
lang: en
---

# Changes between `Release 108` to `Release 109`

# Commands added


| Command                             | Title                               |
|-------------------------------------|-------------------------------------|
| services google sheets sheet create | Create a new sheet                  |
| services google sheets sheet delete | Delete a sheet from the spreadsheet |



# Command spec changed: `util date today`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{Name: "Offset", Desc: "Offset (day)", Default: "0", TypeName: "int"},
  		&{Name: "Utc", Desc: "Use UTC", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util datetime now`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{Name: "OffsetDay", Desc: "Offset (day)", Default: "0", TypeName: "int"},
+ 		&{Name: "OffsetHour", Desc: "Offset (hour)", Default: "0", TypeName: "int"},
+ 		&{Name: "OffsetMin", Desc: "Offset (min)", Default: "0", TypeName: "int"},
+ 		&{Name: "OffsetSec", Desc: "Offset (sec)", Default: "0", TypeName: "int"},
  		&{Name: "Utc", Desc: "Use UTC", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
