---
layout: release
title: リリースの変更点 104
lang: ja
---

# `リリース 104` から `リリース 105` までの変更点

# 追加されたコマンド


| コマンド                                | タイトル                                                                        |
|-----------------------------------------|---------------------------------------------------------------------------------|
| team runas sharedfolder batch leave     | 共有フォルダからメンバーとして一括退出                                          |
| team runas sharedfolder list            | 共有フォルダーの一覧をメンバーとして実行                                        |
| team runas sharedfolder mount add       | 指定したメンバーのDropboxに共有フォルダを追加する                               |
| team runas sharedfolder mount delete    | 指定されたユーザーが指定されたフォルダーをアンマウントする.                     |
| team runas sharedfolder mount list      | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします. |
| team runas sharedfolder mount mountable | メンバーがマウントできるすべての共有フォルダーをリストアップ.                   |



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
- 			Default:  "40",
+ 			Default:  "16",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(40),
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
# コマンド仕様の変更: `team runas file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas file list",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST",
+ 	CliArgs:         "-member-email MEMBER@DOMAIN -path /DROPBOX/PATH/TO/LIST",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 16 identical fields
  }
```
