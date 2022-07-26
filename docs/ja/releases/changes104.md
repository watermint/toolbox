---
layout: release
title: リリースの変更点 103
lang: ja
---

# `リリース 103` から `リリース 104` までの変更点

# 追加されたコマンド


| コマンド                     | タイトル                                                                  |
|------------------------------|---------------------------------------------------------------------------|
| sharedfolder leave           | 共有フォルダーから退出する.                                               |
| sharedfolder mount add       | 共有フォルダを現在のユーザーのDropboxに追加する                           |
| sharedfolder mount delete    | 現在のユーザーが指定されたフォルダーをアンマウントする.                   |
| sharedfolder mount list      | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示          |
| sharedfolder mount mountable | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします. |
| team namespace summary       | チーム・ネームスペースの状態概要を報告する.                               |
| team runas file list         | メンバーとして実行するファイルやフォルダーの一覧                          |



# 削除されたコマンド


| コマンド        | タイトル                                      |
|-----------------|-----------------------------------------------|
| file mount list | マウント/アンマウントされた共有フォルダの一覧 |



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
- 			Default:  "16",
+ 			Default:  "40",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(16),
+ 				"value": float64(40),
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
# コマンド仕様の変更: `member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレスでフィルタリングし\xe3"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
+ 				string("groups.read"),
  				string("members.read"),
  				string("sharing.read"),
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
# コマンド仕様の変更: `sharedfolder list`



## 変更されたレポート: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "このレポートは共有フォルダの一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "shared_folder_id", Desc: "共有フォルダのID"},
  		&{Name: "name", Desc: "共有フォルダの名称"},
  		&{Name: "access_type", Desc: "ユーザーの共有ファイル・フォルダへのアクセ\xe3"...},
  		... // 9 identical elements
  	},
  }
```
