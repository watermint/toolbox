---
layout: release
title: Changes of Release 75
lang: ja
---

# `リリース 75` から `リリース 76` までの変更点

# 追加されたコマンド


| コマンド           | タイトル                               |
|--------------------|----------------------------------------|
| dev replay approve | リプレイをテストバンドルとして承認する |
| dev replay bundle  | すべてのリプレイを実行                 |
| dev replay recipe  | レシピのリプレイ実行                   |
| dev replay remote  | リモートリプレイバンドルの実行         |



# 削除されたコマンド


| コマンド        | タイトル             |
|-----------------|----------------------|
| dev test replay | レシピのリプレイ実行 |



# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "アップロードするローカルファイルのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "PeerName", Desc: "アカウントの別名", Default: "deploy", TypeName: "string", ...},
  		&{
  			Name:     "Timeout",
  			Desc:     "処理タイムアウト(秒)",
- 			Default:  "30",
+ 			Default:  "60",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Days", Desc: "目標日数", Default: "7", TypeName: "essentials.model.mo_int.range_int", ...},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "ワークスペースへのパス.",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Days", Desc: "目標日数", Default: "28", TypeName: "essentials.model.mo_int.range_int", ...},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "ワークスペースへのパス.",
+ 			TypeName: "essentials.model.mo_string.opt_string",
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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
