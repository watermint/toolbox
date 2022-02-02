---
layout: release
title: Changes of Release 73
lang: en
---

# Changes between `Release 73` to `Release 74`

# Commands added


| Command                 | Title                                                      |
|-------------------------|------------------------------------------------------------|
| dev benchmark local     | Create dummy folder structure in local file system.        |
| file mount list         | List mounted/unmounted shared folders                      |
| team content mount list | List all mounted/unmounted shared folders of team members. |



# Command spec changed: `dev ci artifact up`



## Changed report: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "Path",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "Path"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.entry_path", Desc: "Path"},
+ 		&{Name: "input.entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```
# Command spec changed: `file size`



## Added report(s)


| Name | Description |
|------|-------------|
| size | Folder size |



## Deleted report(s)


| Name           | Description                               |
|----------------|-------------------------------------------|
| errors         | This report shows the transaction result. |
| namespace_size | Namespace size                            |


# Command spec changed: `file sync down`



## Changed report: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "Path",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "Path"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.entry_path", Desc: "Path"},
+ 		&{Name: "input.entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```
# Command spec changed: `file sync online`



## Changed report: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "Path",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "Path"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.entry_path", Desc: "Path"},
+ 		&{Name: "input.entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```
# Command spec changed: `file sync up`



## Changed report: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "Path",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "Path"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "entry_shard.attributes", Desc: "Shard attributes"},
  	},
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.entry_path", Desc: "Path"},
+ 		&{Name: "input.entry_shard.file_system_type", Desc: "File system type"},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "Shard ID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "Shard attributes"},
  	},
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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```

## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "This report shows the transaction result.",
+ 	Desc: "Namespace size",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "namespace_name", Desc: "The name of this namespace"},
+ 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
+ 		&{
+ 			Name: "namespace_type",
+ 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
+ 		},
- 		&{Name: "status", Desc: "Status of the operation"},
+ 		&{
+ 			Name: "owner_team_member_id",
+ 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
+ 		},
- 		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
+ 		&{Name: "path", Desc: "Path to the folder"},
- 		&{Name: "input.name", Desc: "The name of this namespace"},
+ 		&{Name: "count_file", Desc: "Number of files under the folder"},
- 		&{
- 			Name: "input.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
+ 		&{Name: "count_folder", Desc: "Number of folders under the folder"},
- 		&{Name: "result.path", Desc: "Path to the folder"},
+ 		&{Name: "count_descendant", Desc: "Number of files and folders under the folder"},
- 		&{Name: "result.count_file", Desc: "Number of files under the folder"},
+ 		&{Name: "size", Desc: "Size of the folder"},
- 		&{Name: "result.count_folder", Desc: "Number of folders under the folder"},
+ 		&{Name: "depth", Desc: "Folder depth"},
- 		&{
- 			Name: "result.count_descendant",
- 			Desc: "Number of files and folders under the folder",
- 		},
+ 		&{
+ 			Name: "mod_time_earliest",
+ 			Desc: "The earliest modification time of a file in this folder or child folders.",
+ 		},
- 		&{Name: "result.size", Desc: "Size of the folder"},
+ 		&{
+ 			Name: "mod_time_latest",
+ 			Desc: "The latest modification time of a file in this folder or child folders",
+ 		},
  		&{
- 			Name: "result.api_complexity",
+ 			Name: "api_complexity",
  			Desc: "Folder complexity index for API operations",
  		},
  	},
  }
```
# Command spec changed: `team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "Report entry for all files and directories depth directories deep",
- 			Default:  "1",
+ 			Default:  "3",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(300),
  				"min":   float64(1),
- 				"value": float64(1),
+ 				"value": float64(3),
  			},
  		},
  		&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...},
  		&{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		... // 6 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Deleted report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |



## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "This report shows the transaction result.",
+ 	Desc: "Namespace size",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "namespace_name", Desc: "The name of this namespace"},
+ 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
+ 		&{
+ 			Name: "namespace_type",
+ 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
+ 		},
- 		&{Name: "status", Desc: "Status of the operation"},
+ 		&{
+ 			Name: "owner_team_member_id",
+ 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
+ 		},
- 		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
+ 		&{Name: "path", Desc: "Path to the folder"},
- 		&{Name: "input.name", Desc: "The name of this namespace"},
+ 		&{Name: "count_file", Desc: "Number of files under the folder"},
- 		&{
- 			Name: "input.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
+ 		&{Name: "count_folder", Desc: "Number of folders under the folder"},
- 		&{Name: "result.path", Desc: "Path to the folder"},
+ 		&{Name: "count_descendant", Desc: "Number of files and folders under the folder"},
- 		&{Name: "result.count_file", Desc: "Number of files under the folder"},
+ 		&{Name: "size", Desc: "Size of the folder"},
- 		&{Name: "result.count_folder", Desc: "Number of folders under the folder"},
+ 		&{Name: "depth", Desc: "Folder depth"},
- 		&{
- 			Name: "result.count_descendant",
- 			Desc: "Number of files and folders under the folder",
- 		},
+ 		&{
+ 			Name: "mod_time_earliest",
+ 			Desc: "The earliest modification time of a file in this folder or child folders.",
+ 		},
- 		&{Name: "result.size", Desc: "Size of the folder"},
+ 		&{
+ 			Name: "mod_time_latest",
+ 			Desc: "The latest modification time of a file in this folder or child folders",
+ 		},
  		&{
- 			Name: "result.api_complexity",
+ 			Name: "api_complexity",
  			Desc: "Folder complexity index for API operations",
  		},
  	},
  }
```
# Command spec changed: `teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "Depth",
- 			Default:  "1",
+ 			Default:  "3",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(2.147483647e+09),
  				"min":   float64(1),
- 				"value": float64(1),
+ 				"value": float64(3),
  			},
  		},
  		&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...},
  		&{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Deleted report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |



## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "This report shows the transaction result.",
+ 	Desc: "Namespace size",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "namespace_name", Desc: "The name of this namespace"},
+ 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
+ 		&{
+ 			Name: "namespace_type",
+ 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
+ 		},
- 		&{Name: "status", Desc: "Status of the operation"},
+ 		&{
+ 			Name: "owner_team_member_id",
+ 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
+ 		},
- 		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
+ 		&{Name: "path", Desc: "Path to the folder"},
- 		&{Name: "input.name", Desc: "The name of this namespace"},
+ 		&{Name: "count_file", Desc: "Number of files under the folder"},
- 		&{
- 			Name: "input.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
+ 		&{Name: "count_folder", Desc: "Number of folders under the folder"},
- 		&{Name: "result.path", Desc: "Path to the folder"},
+ 		&{Name: "count_descendant", Desc: "Number of files and folders under the folder"},
- 		&{Name: "result.count_file", Desc: "Number of files under the folder"},
+ 		&{Name: "size", Desc: "Size of the folder"},
- 		&{Name: "result.count_folder", Desc: "Number of folders under the folder"},
+ 		&{Name: "depth", Desc: "Folder depth"},
- 		&{
- 			Name: "result.count_descendant",
- 			Desc: "Number of files and folders under the folder",
- 		},
+ 		&{
+ 			Name: "mod_time_earliest",
+ 			Desc: "The earliest modification time of a file in this folder or child folders.",
+ 		},
- 		&{Name: "result.size", Desc: "Size of the folder"},
+ 		&{
+ 			Name: "mod_time_latest",
+ 			Desc: "The latest modification time of a file in this folder or child folders",
+ 		},
  		&{
- 			Name: "result.api_complexity",
+ 			Name: "api_complexity",
  			Desc: "Folder complexity index for API operations",
  		},
  	},
  }
```
