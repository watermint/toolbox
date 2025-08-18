---
layout: release
title: Changes of Release 127
lang: en
---

# Changes between `Release 127` to `Release 128`

# Commands added


| Command                           | Title                                                       |
|-----------------------------------|-------------------------------------------------------------|
| dropbox team backup device status | Dropbox Backup device status change in the specified period |



# Command spec changed: `util desktop screenshot interval`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Count", Desc: "Number of screenshots to take. If the value is less than 1, the "..., Default: "-1", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(1),
+ 				"max":   float64(2),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Interval", Desc: "Interval seconds between screenshots.", Default: "10", TypeName: "int", ...},
  		&{Name: "NamePattern", Desc: "Name pattern of screenshot file. You can use the following place"..., Default: "{% raw %}{{.{% endraw %}Sequence}}_{% raw %}{{.{% endraw %}Timestamp}}.png", TypeName: "string", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util desktop screenshot snap`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(1),
+ 				"max":   float64(2),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Path", Desc: "Path to save the screenshot", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
