---
layout: release
title: リリースの変更点: 69
lang: ja
---

# `リリース 69` から `リリース 70` までの変更点

# 追加されたコマンド


| コマンド               | タイトル                                |
|------------------------|-----------------------------------------|
| dev test kvsfootprint  | KVSのメモリフットプリントをテストします |
| teamfolder add         | チームフォルダを追加します              |
| teamfolder member list | チームフォルダのメンバー一覧            |
| teamfolder policy list | チームフォルダのポリシー一覧            |



# 削除されたコマンド


| コマンド | タイトル                                       |
|----------|------------------------------------------------|
| job loop | 指定日時までrunbookを実行します.               |
| job run  | *.runbookoファイルにてワークフローを実行します |



# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "アップロードするローカルファイルのパス", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "PeerName", Desc: "アカウントの別名", Default: "deploy", TypeName: "string", ...},
+ 		&{
+ 			Name:     "Timeout",
+ 			Desc:     "処理タイムアウト(秒)",
+ 			Default:  "30",
+ 			TypeName: "int",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ChunkSizeKb",
  			Desc:     "アップロードチャンク容量(Kバイト)",
- 			Default:  "4096",
+ 			Default:  "65536",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(153600),
  				"min":   float64(1),
- 				"value": float64(4096),
+ 				"value": float64(65536),
  			},
  		},
  		&{Name: "DropboxPath", Desc: "転送先のDropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "FailOnError", Desc: "処理でエラーが発生した場合にエラーを返しま\xe3"..., Default: "false", TypeName: "bool", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
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
