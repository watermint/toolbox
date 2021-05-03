---
layout: release
title: Changes of Release 81
lang: ja
---

# `リリース 81` から `リリース 82` までの変更点

# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "NameNamePrefix", Desc: "名前によるフィルター. 名前の前方一致による\xe3\x83"...},
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
- 		&{
- 			Name:     "Peer",
- 			Desc:     "アカウントの別名",
- 			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
- 		},
  		&{
- 			Name: "SkipExisting",
+ 			Name: "Overwrite",
  			Desc: strings.Join({
- 				"既存ファイルをスキップします. 上書きしませ\xe3\x82",
- 				"\x93",
+ 				"ターゲットパス上に既存のファイルが存在する\xe5",
+ 				"\xa0\xb4合に上書きします",
  				".",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
  		&{Name: "WorkPath", Desc: "テンポラリパス", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
