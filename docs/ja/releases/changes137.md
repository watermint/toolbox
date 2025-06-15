---
layout: release
title: リリースの変更点 136
lang: ja
---

# `リリース 136` から `リリース 137` までの変更点

# 追加されたコマンド


| コマンド         | タイトル                                   |
|------------------|--------------------------------------------|
| dev doc markdown | マークダウンソースからメッセージを生成する |



# コマンド仕様の変更: `util json query`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Compact", Desc: "コンパクトな出力", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "Lines",
- 			Desc:     "JSON Lines (https://jsonlines.org/) フォーマットの読み込み",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "Path", Desc: "ファイルパス", TypeName: "Path"},
  		&{Name: "Query", Desc: "クエリー文字列。 ", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
