---
layout: release
title: リリースの変更点 108
lang: ja
---

# `リリース 108` から `リリース 109` までの変更点

# 追加されたコマンド

| コマンド                            | タイトル                             |
|-------------------------------------|--------------------------------------|
| services google sheets sheet create | 新規シートの作成                     |
| services google sheets sheet delete | スプレッドシートからシートを削除する |

# コマンド仕様の変更: `util date today`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{Name: "Offset", Desc: "オフセット（日）", Default: "0", TypeName: "int"},
  		&{Name: "Utc", Desc: "UTC(協定世界時)を使用する", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

# コマンド仕様の変更: `util datetime now`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "OffsetDay",
+ 			Desc:     "オフセット（日）",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "OffsetHour",
+ 			Desc:     "オフセット（時間）",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "OffsetMin",
+ 			Desc:     "オフセット（分）",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "OffsetSec",
+ 			Desc:     "オフセット（秒）",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
  		&{Name: "Utc", Desc: "UTC(協定世界時)を使用する", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
