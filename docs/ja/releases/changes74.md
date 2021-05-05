---
layout: release
title: Changes of Release 73
lang: ja
---

# `リリース 73` から `リリース 74` までの変更点

# 追加されたコマンド


| コマンド                | タイトル                                                                               |
|-------------------------|----------------------------------------------------------------------------------------|
| dev benchmark local     | ローカルファイルシステムにダミーのフォルダ構造を作成します.                            |
| file mount list         | マウント/アンマウントされた共有フォルダの一覧                                          |
| team content mount list | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします. |



# コマンド仕様の変更: `dev ci artifact up`



## 変更されたレポート: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "パス",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "パス"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "ファイルシステムの種別"},
+ 		&{Name: "entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.entry_path", Desc: "パス"},
+ 		&{
+ 			Name: "input.entry_shard.file_system_type",
+ 			Desc: "ファイルシステムの種別",
+ 		},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```
# コマンド仕様の変更: `file size`



## 追加されたレポート


| 名称 | 説明             |
|------|------------------|
| size | フォルダのサイズ |



## 削除されたレポート


| 名称           | 説明                                |
|----------------|-------------------------------------|
| errors         | このレポートは処理結果を出力します. |
| namespace_size | 名前空間のサイズ.                   |


# コマンド仕様の変更: `file sync down`



## 変更されたレポート: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "パス",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "パス"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "ファイルシステムの種別"},
+ 		&{Name: "entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.entry_path", Desc: "パス"},
+ 		&{
+ 			Name: "input.entry_shard.file_system_type",
+ 			Desc: "ファイルシステムの種別",
+ 		},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```
# コマンド仕様の変更: `file sync online`



## 変更されたレポート: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "パス",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "パス"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "ファイルシステムの種別"},
+ 		&{Name: "entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.entry_path", Desc: "パス"},
+ 		&{
+ 			Name: "input.entry_shard.file_system_type",
+ 			Desc: "ファイルシステムの種別",
+ 		},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```
# コマンド仕様の変更: `file sync up`



## 変更されたレポート: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "パス",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "entry_path", Desc: "パス"},
+ 		&{Name: "entry_shard.file_system_type", Desc: "ファイルシステムの種別"},
+ 		&{Name: "entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "entry_shard.attributes", Desc: "シャードの属性"},
  	},
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.entry_path", Desc: "パス"},
+ 		&{
+ 			Name: "input.entry_shard.file_system_type",
+ 			Desc: "ファイルシステムの種別",
+ 		},
+ 		&{Name: "input.entry_shard.shard_id", Desc: "シャードID"},
+ 		&{Name: "input.entry_shard.attributes", Desc: "シャードの属性"},
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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```

## 変更されたレポート: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "このレポートは処理結果を出力します.",
+ 	Desc: "名前空間のサイズ.",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "namespace_name", Desc: "名前空間の名称"},
+ 		&{Name: "namespace_id", Desc: "名前空間ID"},
+ 		&{
+ 			Name: "namespace_type",
+ 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
+ 		},
- 		&{Name: "status", Desc: "処理の状態"},
+ 		&{
+ 			Name: "owner_team_member_id",
+ 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
+ 		},
- 		&{Name: "reason", Desc: "失敗またはスキップの理由"},
+ 		&{Name: "path", Desc: "フォルダへのパス"},
- 		&{Name: "input.name", Desc: "名前空間の名称"},
+ 		&{Name: "count_file", Desc: "このフォルダに含まれるファイル数"},
- 		&{
- 			Name: "input.namespace_type",
- 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
- 		},
+ 		&{Name: "count_folder", Desc: "このフォルダに含まれるフォルダ数"},
- 		&{Name: "result.path", Desc: "フォルダへのパス"},
+ 		&{
+ 			Name: "count_descendant",
+ 			Desc: "このフォルダに含まれるファイル・フォルダ数",
+ 		},
- 		&{
- 			Name: "result.count_file",
- 			Desc: "このフォルダに含まれるファイル数",
- 		},
+ 		&{Name: "size", Desc: "フォルダのサイズ"},
- 		&{
- 			Name: "result.count_folder",
- 			Desc: "このフォルダに含まれるフォルダ数",
- 		},
+ 		&{Name: "depth", Desc: "フォルダの深さ"},
- 		&{
- 			Name: "result.count_descendant",
- 			Desc: "このフォルダに含まれるファイル・フォルダ数",
- 		},
+ 		&{
+ 			Name: "mod_time_earliest",
+ 			Desc: "このフォルダまたは子フォルダ内のファイルの最も古い更新日時",
+ 		},
- 		&{Name: "result.size", Desc: "フォルダのサイズ"},
+ 		&{
+ 			Name: "mod_time_latest",
+ 			Desc: "このフォルダまたは子フォルダ内のファイルの最も新しい更新日時",
+ 		},
  		&{
- 			Name: "result.api_complexity",
+ 			Name: "api_complexity",
  			Desc: "APIを用いて操作する場合のフォルダ複雑度の指標",
  		},
  	},
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "フォルダ階層数の指定",
- 			Default:  "1",
+ 			Default:  "3",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(300),
  				"min":   float64(1),
- 				"value": float64(1),
+ 				"value": float64(3),
  			},
  		},
  		&{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		... // 6 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 削除されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |



## 変更されたレポート: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "このレポートは処理結果を出力します.",
+ 	Desc: "名前空間のサイズ.",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "namespace_name", Desc: "名前空間の名称"},
+ 		&{Name: "namespace_id", Desc: "名前空間ID"},
+ 		&{
+ 			Name: "namespace_type",
+ 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
+ 		},
- 		&{Name: "status", Desc: "処理の状態"},
+ 		&{
+ 			Name: "owner_team_member_id",
+ 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
+ 		},
- 		&{Name: "reason", Desc: "失敗またはスキップの理由"},
+ 		&{Name: "path", Desc: "フォルダへのパス"},
- 		&{Name: "input.name", Desc: "名前空間の名称"},
+ 		&{Name: "count_file", Desc: "このフォルダに含まれるファイル数"},
- 		&{
- 			Name: "input.namespace_type",
- 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
- 		},
+ 		&{Name: "count_folder", Desc: "このフォルダに含まれるフォルダ数"},
- 		&{Name: "result.path", Desc: "フォルダへのパス"},
+ 		&{
+ 			Name: "count_descendant",
+ 			Desc: "このフォルダに含まれるファイル・フォルダ数",
+ 		},
- 		&{
- 			Name: "result.count_file",
- 			Desc: "このフォルダに含まれるファイル数",
- 		},
+ 		&{Name: "size", Desc: "フォルダのサイズ"},
- 		&{
- 			Name: "result.count_folder",
- 			Desc: "このフォルダに含まれるフォルダ数",
- 		},
+ 		&{Name: "depth", Desc: "フォルダの深さ"},
- 		&{
- 			Name: "result.count_descendant",
- 			Desc: "このフォルダに含まれるファイル・フォルダ数",
- 		},
+ 		&{
+ 			Name: "mod_time_earliest",
+ 			Desc: "このフォルダまたは子フォルダ内のファイルの最も古い更新日時",
+ 		},
- 		&{Name: "result.size", Desc: "フォルダのサイズ"},
+ 		&{
+ 			Name: "mod_time_latest",
+ 			Desc: "このフォルダまたは子フォルダ内のファイルの最も新しい更新日時",
+ 		},
  		&{
- 			Name: "result.api_complexity",
+ 			Name: "api_complexity",
  			Desc: "APIを用いて操作する場合のフォルダ複雑度の指標",
  		},
  	},
  }
```
# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "深さ",
- 			Default:  "1",
+ 			Default:  "3",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(2.147483647e+09),
  				"min":   float64(1),
- 				"value": float64(1),
+ 				"value": float64(3),
  			},
  		},
  		&{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 削除されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |



## 変更されたレポート: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "このレポートは処理結果を出力します.",
+ 	Desc: "名前空間のサイズ.",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "namespace_name", Desc: "名前空間の名称"},
+ 		&{Name: "namespace_id", Desc: "名前空間ID"},
+ 		&{
+ 			Name: "namespace_type",
+ 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
+ 		},
- 		&{Name: "status", Desc: "処理の状態"},
+ 		&{
+ 			Name: "owner_team_member_id",
+ 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
+ 		},
- 		&{Name: "reason", Desc: "失敗またはスキップの理由"},
+ 		&{Name: "path", Desc: "フォルダへのパス"},
- 		&{Name: "input.name", Desc: "名前空間の名称"},
+ 		&{Name: "count_file", Desc: "このフォルダに含まれるファイル数"},
- 		&{
- 			Name: "input.namespace_type",
- 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
- 		},
+ 		&{Name: "count_folder", Desc: "このフォルダに含まれるフォルダ数"},
- 		&{Name: "result.path", Desc: "フォルダへのパス"},
+ 		&{
+ 			Name: "count_descendant",
+ 			Desc: "このフォルダに含まれるファイル・フォルダ数",
+ 		},
- 		&{
- 			Name: "result.count_file",
- 			Desc: "このフォルダに含まれるファイル数",
- 		},
+ 		&{Name: "size", Desc: "フォルダのサイズ"},
- 		&{
- 			Name: "result.count_folder",
- 			Desc: "このフォルダに含まれるフォルダ数",
- 		},
+ 		&{Name: "depth", Desc: "フォルダの深さ"},
- 		&{
- 			Name: "result.count_descendant",
- 			Desc: "このフォルダに含まれるファイル・フォルダ数",
- 		},
+ 		&{
+ 			Name: "mod_time_earliest",
+ 			Desc: "このフォルダまたは子フォルダ内のファイルの最も古い更新日時",
+ 		},
- 		&{Name: "result.size", Desc: "フォルダのサイズ"},
+ 		&{
+ 			Name: "mod_time_latest",
+ 			Desc: "このフォルダまたは子フォルダ内のファイルの最も新しい更新日時",
+ 		},
  		&{
- 			Name: "result.api_complexity",
+ 			Name: "api_complexity",
  			Desc: "APIを用いて操作する場合のフォルダ複雑度の指標",
  		},
  	},
  }
```
