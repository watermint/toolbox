---
layout: release
title: Changes of Release 71
lang: ja
---

# `リリース 71` から `リリース 72` までの変更点

# 追加されたコマンド


| コマンド                                    | タイトル                                    |
|---------------------------------------------|---------------------------------------------|
| dev stage gmail                             | Gmail コマンド                              |
| dev stage scoped                            | Dropboxのスコープ付きOAuthアプリテスト      |
| services google mail filter add             | フィルターを追加します.                     |
| services google mail filter batch add       | クエリによるラベルの一括追加・削除          |
| services google mail filter delete          | フィルタの削除                              |
| services google mail filter list            | フィルターの一覧                            |
| services google mail label add              | ラベルの追加                                |
| services google mail label delete           | ラベルの削除.                               |
| services google mail label list             | ラベルのリスト                              |
| services google mail label rename           | ラベルの名前を変更する                      |
| services google mail message label add      | メッセージにラベルを追加                    |
| services google mail message label delete   | メッセージからラベルを削除する              |
| services google mail message list           | メッセージの一覧                            |
| services google mail message processed list | 処理された形式でメッセージを一覧表示します. |
| services google mail thread list            | スレッド一覧                                |



# コマンド仕様の変更: `dev doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Badge", Desc: "ビルド状態のバッジを含める", Default: "true", TypeName: "bool", ...},
  		&{Name: "CommandPath", Desc: "コマンドマニュアルを作成する相対パス", Default: "doc/generated/", TypeName: "string", ...},
- 		&{
- 			Name:     "Filename",
- 			Desc:     "ファイル名",
- 			Default:  "README.md",
- 			TypeName: "string",
- 		},
  		&{
- 			Name:    "Lang",
+ 			Name:    "DocLang",
  			Desc:    "言語",
  			Default: "",
  			... // 2 identical fields
  		},
+ 		&{
+ 			Name:     "Filename",
+ 			Desc:     "ファイル名",
+ 			Default:  "README.md",
+ 			TypeName: "string",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev util curl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BufferSize", Desc: "バッファのサイズ", Default: "65536", TypeName: "domain.common.model.mo_int.range_int", ...},
  		&{
  			Name:     "Record",
  			Desc:     "テスト用に直接テストレコードを指定",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team diag explorer`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"File": "business_file",
  		"Info": "business_info",
  		"Mgmt": "business_management",
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
