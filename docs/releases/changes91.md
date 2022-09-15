---
layout: release
title: Changes of Release 90
lang: en
---

# Changes between `Release 90` to `Release 91`

# Commands added


| Command           | Title                           |
|-------------------|---------------------------------|
| dev build info    | Generate build information file |
| dev build package | Package a build                 |
| dev release doc   | Generate release documents      |
| util git clone    | Clone git repository            |



# Command spec changed: `dev build license`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DestPath", Desc: "Dest path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
- 		&{
- 			Name:     "SourcePath",
- 			Desc:     "Path to licenses (go-licenses output folder)",
- 			TypeName: "essentials.model.mo_path.file_system_path_impl",
- 			TypeAttr: map[string]any{"shouldExist": bool(false)},
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		... // 4 identical entries
  		"Info": "business_info",
  		"Mgmt": "business_management",
+ 		"Peer": "github_public",
  	},
  	Services: {"dropbox", "dropbox_business", "github"},
  	IsSecret: true,
  	... // 12 identical fields
  }
```
# Command spec changed: `team diag explorer`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"File": "business_file",
  		"Info": "business_info",
  		"Mgmt": "business_management",
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: true,
  	... // 12 identical fields
  }
```
