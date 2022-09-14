---
layout: release
title: リリースの変更点 102
lang: ja
---

# `リリース 102` から `リリース 103` までの変更点

# 追加されたコマンド


| コマンド                 | タイトル                                                           |
|--------------------------|--------------------------------------------------------------------|
| dev module list          | 依存モジュール一覧                                                 |
| dev test setup massfiles | テストファイルとしてウィキメディアダンプファイルをアップロードする |
| file revision download   | ファイルリビジョンをダウンロードする                               |
| file revision list       | ファイルリビジョン一覧                                             |
| file revision restore    | ファイルリビジョンを復元する                                       |



# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BlockBlockSize",
  			Desc:     "一括アップロード時のブロックサイズ",
- 			Default:  "12",
+ 			Default:  "16",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(12),
+ 				"value": float64(16),
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
