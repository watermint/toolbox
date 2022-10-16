---
layout: release
title: Changes of Release 111
lang: en
---

# Changes between `Release 111` to `Release 112`

# Command spec changed: `dev test setup massfiles`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Base", Desc: "Dropbox base path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "BatchSize", Desc: "Batch size", Default: "1000", TypeName: "essentials.model.mo_int.range_int", ...},
+ 		&{
+ 			Name:     "CommitConcurrency",
+ 			Desc:     "Number of concurrency to commit",
+ 			Default:  "3",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]any{"max": float64(10), "min": float64(1), "value": float64(3)},
+ 		},
  		&{Name: "Offset", Desc: "Upload offset (skip # pages)", Default: "0", TypeName: "int", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
+ 		&{
+ 			Name:     "ShardSize",
+ 			Desc:     "Number of shards (number of folder/namespaces to distribute). Need to setup namespaces separately.",
+ 			Default:  "20",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]any{"max": float64(1000), "min": float64(1), "value": float64(20)},
+ 		},
  		&{Name: "Source", Desc: "Source file", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."},
  		&{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."},
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
+ 				string("groups.read"),
  				string("sharing.read"),
  				string("team_data.member"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
