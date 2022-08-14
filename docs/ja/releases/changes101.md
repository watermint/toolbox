---
layout: release
title: リリースの変更点 100
lang: ja
---

# `リリース 100` から `リリース 101` までの変更点

# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BlockBlockSize",
  			Desc:     "一括アップロード時のブロックサイズ",
- 			Default:  "40",
+ 			Default:  "8",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(40),
+ 				"value": float64(8),
  			},
  		},
  		&{Name: "Method", Desc: "アップロード方法", Default: "block", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "NumFiles", Desc: "ファイル数.", Default: "1000", TypeName: "int", ...},
  		... // 7 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
