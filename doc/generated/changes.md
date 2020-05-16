# Changes between `Release 66` to `Release 67`

# Commands added

| Command                 | Title                                            |
|-------------------------|--------------------------------------------------|
| dev util anonymise      | Anonymise capture log                            |
| job log jobid           | Retrieve logs of specified Job ID                |
| job log kind            | Concatenate and print logs of specified log kind |
| job log last            | Print the last job log files                     |
| member clear externalid | Clear external_id of members                     |


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
# Command spec changed: `member list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "Include deleted members.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info"},
  	},
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
+ 		&{
+ 			Name: "MemberTypeExternal",
+ 			Desc: "Filter folder members. Keep only members are external (not in the same team). Note: Invited members are marked as external member.",
+ 		},
+ 		&{
+ 			Name: "MemberTypeInternal",
+ 			Desc: "Filter folder members. Keep only members are internal (in the same team). Note: Invited members are marked as external member.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
## Changed report: membership

```
  &dc_recipe.Report{
  	Name: "membership",
  	Desc: "This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "member_name", Desc: "Name of this member"},
  		&{Name: "member_email", Desc: "Email address of this member"},
+ 		&{
+ 			Name: "same_team",
+ 			Desc: "Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not.",
+ 		},
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
