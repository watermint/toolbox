---
layout: release
title: Changes of Release 102
lang: en
---

# Changes between `Release 102` to `Release 103`

# Commands added


| Command                  | Title                                   |
|--------------------------|-----------------------------------------|
| dev module list          | Dependent module list                   |
| dev test setup massfiles | Upload Wikimedia dump file as test file |
| file revision download   | Download the file revision              |
| file revision list       | List file revisions                     |
| file revision restore    | Restore the file revision               |



# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BlockBlockSize",
  			Desc:     "Block size for batch upload",
- 			Default:  "12",
+ 			Default:  "16",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(12),
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
