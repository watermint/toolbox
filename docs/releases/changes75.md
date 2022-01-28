---
layout: release
title: Changes of Release 74
lang: en
---

# Changes between `Release 74` to `Release 75`

# Commands added


| Command                          | Title                        |
|----------------------------------|------------------------------|
| dev stage teamfolder             | Team folder operation sample |
| dev test replay                  | Replay recipe                |
| services slack conversation list | List channels                |



# Command spec changed: `dev benchmark local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {},
  	Services:       {},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 2 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
- 			Name:     "NodeLambda",
+ 			Name:     "NumFiles",
- 			Desc:     "Lambda parameter for nodes",
+ 			Desc:     "Number of files.",
- 			Default:  "100",
+ 			Default:  "1000",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
- 		&{
- 			Name:     "NodeMax",
- 			Desc:     "Maximum number of nodes",
- 			Default:  "1000",
- 			TypeName: "int",
- 		},
- 		&{
- 			Name:     "NodeMin",
- 			Desc:     "Minimum number of nodes",
- 			Default:  "100",
- 			TypeName: "int",
- 		},
  		&{Name: "Path", Desc: "Path to create", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
- 			Name:     "SizeMax",
+ 			Name:     "SizeMaxKb",
- 			Desc:     "Maximum file size",
+ 			Desc:     "Maximum file size (KiB).",
- 			Default:  "2097152",
+ 			Default:  "2048",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  		&{
- 			Name:     "SizeMin",
+ 			Name:     "SizeMinKb",
- 			Desc:     "Minimum file size",
+ 			Desc:     "Minimum file size (KiB).",
  			Default:  "0",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "user_full"},
  	Services:       {"dropbox"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 2 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KiB", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
- 			Name:     "Lambda",
+ 			Name:     "NumFiles",
- 			Desc:     "Node number Lambda",
+ 			Desc:     "Number of files.",
- 			Default:  "100",
+ 			Default:  "1000",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
- 		&{
- 			Name:     "MaxNodes",
- 			Desc:     "Maximum number of nodes",
- 			Default:  "1000",
- 			TypeName: "int",
- 		},
- 		&{
- 			Name:     "MinNodes",
- 			Desc:     "Minimum number of nodes",
- 			Default:  "10",
- 			TypeName: "int",
- 		},
  		&{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "SizeMaxKb",
+ 			Desc:     "Maximum file size (KiB).",
+ 			Default:  "2048",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "SizeMinKb",
+ 			Desc:     "Minimum file size (KiB).",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
