---
layout: release
title: リリースの変更点 111
lang: ja
---

# `リリース 111` から `リリース 112` までの変更点

# コマンド仕様の変更: `dev test setup massfiles`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Base", Desc: "Dropboxのベースパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "BatchSize", Desc: "バッチサイズ", Default: "1000", TypeName: "essentials.model.mo_int.range_int", ...},
+ 		&{
+ 			Name:     "CommitConcurrency",
+ 			Desc:     "コミットする同時実行数",
+ 			Default:  "3",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]any{"max": float64(10), "min": float64(1), "value": float64(3)},
+ 		},
  		&{Name: "Offset", Desc: "アップロードオフセット（ページ数省略）", Default: "0", TypeName: "int", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
+ 		&{
+ 			Name:     "ShardSize",
+ 			Desc:     "シャード数（配布するフォルダ/名前空間の数）。ネームスペースを別途設定する必要がある。",
+ 			Default:  "20",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]any{"max": float64(1000), "min": float64(1), "value": float64(20)},
+ 		},
  		&{Name: "Source", Desc: "ソースファイル", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...},
  		&{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...},
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
+ 				string("groups.read"),
  				string("sharing.read"),
  				string("team_data.member"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
