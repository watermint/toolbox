# `リリース 73` から `リリース 74` までの変更点

# 追加されたコマンド


| コマンド            | タイトル                                            |
|---------------------|-----------------------------------------------------|
| dev benchmark local | Create dummy folder structure in local file system. |



# コマンド仕様の変更: `dev ci artifact up`


## 変更されたレポート: deleted

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
## 変更されたレポート: skipped

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
# コマンド仕様の変更: `file size`


## 追加されたレポート


| 名称 | 説明        |
|------|-------------|
| size | Folder size |


## 削除されたレポート


| 名称           | 説明                                      |
|----------------|-------------------------------------------|
| errors         | This report shows the transaction result. |
| namespace_size | Namespace size                            |


# コマンド仕様の変更: `file sync down`


## 変更されたレポート: deleted

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
## 変更されたレポート: skipped

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
# コマンド仕様の変更: `file sync online`


## 変更されたレポート: deleted

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
## 変更されたレポート: skipped

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
# コマンド仕様の変更: `file sync up`


## 変更されたレポート: deleted

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
## 変更されたレポート: skipped

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
# コマンド仕様の変更: `team diag explorer`


## 設定が変更されたコマンド


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
  	... // 7 identical fields
  }
```
## 変更されたレポート: namespace_size

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
# コマンド仕様の変更: `team namespace file size`


## 削除されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


## 変更されたレポート: namespace_size

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
# コマンド仕様の変更: `teamfolder file size`


## 削除されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


## 変更されたレポート: namespace_size

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
