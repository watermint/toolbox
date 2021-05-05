---
layout: release
title: Changes of Release 74
lang: ja
---

# `リリース 74` から `リリース 75` までの変更点

# 追加されたコマンド


| コマンド                         | タイトル                     |
|----------------------------------|------------------------------|
| dev stage teamfolder             | チームフォルダ処理のサンプル |
| dev test replay                  | レシピのリプレイ実行         |
| services slack conversation list | チャネルの一覧               |



# コマンド仕様の変更: `dev benchmark local`



## 設定が変更されたコマンド


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
- 			Desc:     "ノードのラムダパラメーター",
+ 			Desc:     "ファイル数.",
- 			Default:  "100",
+ 			Default:  "1000",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
- 		&{Name: "NodeMax", Desc: "最大ノード数", Default: "1000", TypeName: "int"},
- 		&{Name: "NodeMin", Desc: "最小ノード数", Default: "100", TypeName: "int"},
  		&{Name: "Path", Desc: "作成するパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
- 			Name:     "SizeMax",
+ 			Name:     "SizeMaxKb",
- 			Desc:     "最大ファイルサイズ",
+ 			Desc:     "最大ファイルサイズ (KiB).",
- 			Default:  "2097152",
+ 			Default:  "2048",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  		&{
- 			Name:     "SizeMin",
+ 			Name:     "SizeMinKb",
- 			Desc:     "最小ファイルサイズ",
+ 			Desc:     "最小ファイルサイズ (KiB).",
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
# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


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
  		&{Name: "ChunkSizeKb", Desc: "チャンクサイズをKiB単位でアップロード", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
- 			Name:     "Lambda",
+ 			Name:     "NumFiles",
- 			Desc:     "ノード数のλ",
+ 			Desc:     "ファイル数.",
- 			Default:  "100",
+ 			Default:  "1000",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
- 		&{Name: "MaxNodes", Desc: "最大ノード数", Default: "1000", TypeName: "int"},
- 		&{Name: "MinNodes", Desc: "最小ノード数", Default: "10", TypeName: "int"},
  		&{Name: "Path", Desc: "Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "SizeMaxKb",
+ 			Desc:     "最大ファイルサイズ (KiB).",
+ 			Default:  "2048",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "SizeMinKb",
+ 			Desc:     "最小ファイルサイズ (KiB).",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
