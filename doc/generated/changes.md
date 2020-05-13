# Changes between `Release 66` to `Release 67`

# Commands added

| Command       | Title                                            |
|---------------|--------------------------------------------------|
| job log jobid | Retrieve logs of specified Job ID                |
| job log kind  | Concatenate and print logs of specified log kind |
| job log last  | Print the last job log files                     |


# Command spec changed: `job history list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to workspace",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  }
```
# Command spec changed: `team content member`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "Filter by folder name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "Filter by folder name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "Filter by folder name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
# Command spec changed: `team content policy`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "Filter by folder name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "Filter by folder name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "Filter by folder name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
