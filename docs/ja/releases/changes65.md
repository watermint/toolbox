---
layout: release
title: Changes of Release 64
lang: ja
---

# `リリース 64` から `リリース 65` までの変更点

# コマンド仕様の変更: `dev ci auth import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "EnvName",
  			Desc:     "環境変数名",
- 			Default:  "TOOLBOX_ENDTOEND",
+ 			Default:  "TOOLBOX_ENDTOEND_TOKEN",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "PeerName", Desc: "アカウントの別名", Default: "end_to_end_test", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev desktop stop`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "WaitSeconds",
  			Desc:     "指定秒数後にアプリケーションの停止を試みます",
- 			Default:  "60",
+ 			Default:  "0",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(2.147483647e+09),
  				"min":   float64(0),
- 				"value": float64(60),
+ 				"value": float64(0),
  			},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "ConnGithub", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "TestResource",
  			Desc:     "テストリソースへのパス",
- 			Default:  "test/dev/resource.json",
+ 			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
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
  		&{Name: "ChunkSizeKb", Desc: "アップロードチャンク容量(Kバイト)", Default: "153600", TypeName: "domain.common.model.mo_int.range_int", ...},
  		&{Name: "DropboxPath", Desc: "転送先のDropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
+ 		&{
+ 			Name:     "FailOnError",
+ 			Desc:     "処理でエラーが発生した場合にエラーを返します. このコマンドはこのフラグが指定されない場"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "LocalPath", Desc: "ローカルファイルのパス", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
