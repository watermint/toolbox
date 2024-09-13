---
layout: release
title: リリースの変更点 127
lang: ja
---

# `リリース 127` から `リリース 128` までの変更点

# 追加されたコマンド


| コマンド                          | タイトル                                                              |
|-----------------------------------|-----------------------------------------------------------------------|
| dropbox team backup device status | Dropbox バックアップ デバイスのステータスが指定期間内に変更された場合 |



# コマンド仕様の変更: `util desktop screenshot interval`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Count", Desc: "スクリーンショットの枚数。値が1未満の場合、"..., Default: "-1", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(1),
+ 				"max":   float64(2),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Interval", Desc: "スクリーンショットの間隔秒数。", Default: "10", TypeName: "int", ...},
  		&{Name: "NamePattern", Desc: "スクリーンショットファイルの名前パターン。\xe4"..., Default: "{% raw %}{{.{% endraw %}Sequence}}_{% raw %}{{.{% endraw %}Timestamp}}.png", TypeName: "string", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util desktop screenshot snap`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(1),
+ 				"max":   float64(2),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Path", Desc: "スクリーンショットを保存するパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
