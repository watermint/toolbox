---
layout: release
title: Changes of Release 69
lang: en
---

# Changes between `Release 69` to `Release 70`

# Commands added


| Command                | Title                         |
|------------------------|-------------------------------|
| dev test kvsfootprint  | Test KVS memory footprint     |
| teamfolder add         | Add team folder to the team   |
| teamfolder member list | List team folder members      |
| teamfolder policy list | List policies of team folders |



# Commands deleted


| Command  | Title                                 |
|----------|---------------------------------------|
| job loop | Run runbook until specified date/time |
| job run  | Run workflow with *.runbook file      |



# Command spec changed: `dev ci artifact up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path to upload", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string", ...},
+ 		&{
+ 			Name:     "Timeout",
+ 			Desc:     "Operation timeout in seconds",
+ 			Default:  "30",
+ 			TypeName: "int",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ChunkSizeKb",
  			Desc:     "Upload chunk size in KB",
- 			Default:  "4096",
+ 			Default:  "65536",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(153600),
  				"min":   float64(1),
- 				"value": float64(4096),
+ 				"value": float64(65536),
  			},
  		},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "FailOnError", Desc: "Returns error when any error happens while the operation. This c"..., Default: "false", TypeName: "bool", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Filter the returned events to a single category. This field is o"...,
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
