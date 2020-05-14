# `リリース 66` から `リリース 67` までの変更点

# 追加されたコマンド

| コマンド      | タイトル                                         |
|---------------|--------------------------------------------------|
| job log jobid | Retrieve logs of specified Job ID                |
| job log kind  | Concatenate and print logs of specified log kind |
| job log last  | Print the last job log files                     |


# コマンド仕様の変更: `job history list`


## 設定が変更されたコマンド


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
# コマンド仕様の変更: `member list`


## 設定が変更されたコマンド


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
# コマンド仕様の変更: `team content member`


## 設定が変更されたコマンド


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
+ 			Desc: "Filter folder members. Keep only members are external (not in the same team).",
+ 		},
+ 		&{
+ 			Name: "MemberTypeInternal",
+ 			Desc: "Filter folder members. Keep only members are internal (in the same team).",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
## 変更されたレポート: membership

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
# コマンド仕様の変更: `team content policy`


## 設定が変更されたコマンド


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
