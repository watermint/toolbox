---
layout: release
title: Changes of Release 101
lang: en
---

# Changes between `Release 101` to `Release 102`

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
- 			Default:  "8",
+ 			Default:  "12",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(8),
+ 				"value": float64(12),
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
