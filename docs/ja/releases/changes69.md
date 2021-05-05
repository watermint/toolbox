---
layout: release
title: リリースの変更点: 68
lang: ja
---

# `リリース 68` から `リリース 69` までの変更点

# 追加されたコマンド


| コマンド                 | タイトル                                   |
|--------------------------|--------------------------------------------|
| team content member list | チームフォルダや共有フォルダのメンバー一覧 |
| team content policy list | チームフォルダと共有フォルダのポリシー一覧 |



# 削除されたコマンド


| コマンド            | タイトル                                   |
|---------------------|--------------------------------------------|
| team content member | チームフォルダや共有フォルダのメンバー一覧 |
| team content policy | チームフォルダと共有フォルダのポリシー一覧 |



# コマンド仕様の変更: `dev util image jpeg`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "Path", Desc: "ファイルを生成するパス", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "Quality", Desc: "JPEGの品質", Default: "75", TypeName: "domain.common.model.mo_int.range_int", ...},
  		&{
  			Name:     "Seed",
- 			Desc:     "Random seed",
+ 			Desc:     "乱数のシード",
  			Default:  "1",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  		&{Name: "Width", Desc: "幅", Default: "1920", TypeName: "domain.common.model.mo_int.range_int", ...},
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
- 			Default:  "153600",
+ 			Default:  "4096",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(153600),
  				"min":   float64(1),
- 				"value": float64(153600),
+ 				"value": float64(4096),
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
