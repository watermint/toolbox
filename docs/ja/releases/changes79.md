---
layout: release
title: リリースの変更点 78
lang: ja
---

# `リリース 78` から `リリース 79` までの変更点

# 追加されたコマンド


| コマンド                  | タイトル                             |
|---------------------------|--------------------------------------|
| dev stage gui             | GUIコンセプト実証                    |
| file archive local        | ローカルファイルをアーカイブします   |
| group member batch add    | グループにメンバーを一括追加         |
| group member batch delete | グループからメンバーを削除           |
| group member batch update | グループからメンバーを追加または削除 |
| team report activity      | アクティビティ レポート              |
| team report devices       | デバイス レポート空のレポート        |
| team report membership    | メンバーシップ レポート              |
| team report storage       | ストレージ レポート                  |



# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "出力するリンクを可視性にてフィルターします "...,
- 			Default:  "public",
+ 			Default:  "all",
  			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{
  				"options": []interface{}{
+ 					string("all"),
  					string("public"),
  					string("team_only"),
  					... // 3 identical elements
  				},
  			},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
