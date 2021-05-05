---
layout: release
title: リリースの変更点: 90
lang: ja
---

# `リリース 90` から `リリース 91` までの変更点

# 追加されたコマンド


| コマンド          | タイトル                   |
|-------------------|----------------------------|
| dev build info    | ビルド情報ファイルを生成   |
| dev build package | ビルドのパッケージ化       |
| dev release doc   | リリースドキュメントの作成 |
| util git clone    | git リポジトリをクローン   |



# コマンド仕様の変更: `dev build license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DestPath", Desc: "出力先パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
- 		&{
- 			Name:     "SourcePath",
- 			Desc:     "ライセンスへのパス (go-licenses 出力フォルダ)",
- 			TypeName: "essentials.model.mo_path.file_system_path_impl",
- 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		... // 4 identical entries
  		"Info": "business_info",
  		"Mgmt": "business_management",
+ 		"Peer": "github_public",
  	},
  	Services: {"dropbox", "dropbox_business", "github"},
  	IsSecret: true,
  	... // 11 identical fields
  }
```
