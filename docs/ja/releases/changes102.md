---
layout: release
title: リリースの変更点 101
lang: ja
---

# `リリース 101` から `リリース 102` までの変更点

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
- 			Default:  "8",
+ 			Default:  "12",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(8),
+ 				"value": float64(12),
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
# コマンド仕様の変更: `group clear externalid`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.name", Desc: "外部IDをクリアするためのグループ名"},
  		&{Name: "result.group_name", Desc: "グループ名称"},
  		&{
  			Name: "result.group_id",
- 			Desc: "グループ ID",
+ 			Desc: "グループID",
  		},
  		&{Name: "result.group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
  		&{Name: "result.group_external_id", Desc: "チームがグループに付加することができる外部ID."},
  		&{Name: "result.member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `util image placeholder`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Color", Desc: "背景色", Default: "white", TypeName: "string", ...},
  		&{
  			Name: "FontPath",
  			Desc: strings.Join({
  				"True",
+ 				" ",
  				"Typeフォントへのパス（テキストを描画する必要",
  				"がある場合は必須）",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "FontSize", Desc: "文字サイズ", Default: "12", TypeName: "int", ...},
  		&{Name: "Height", Desc: "高さ(ピクセル)", Default: "400", TypeName: "int", ...},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
