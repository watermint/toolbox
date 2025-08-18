---
layout: release
title: リリースの変更点 130
lang: ja
---

# `リリース 130` から `リリース 131` までの変更点

# 追加されたコマンド


| コマンド                            | タイトル                                                |
|-------------------------------------|---------------------------------------------------------|
| dropbox sign request list           | 署名依頼リスト                                          |
| dropbox sign request signature list | リクエストの署名一覧                                    |
| log api job                         | ジョブIDで指定されたジョブのAPIログの統計情報を表示する |
| log api name                        | ジョブ名で指定されたジョブのAPIログの統計情報を表示する |



# 削除されたコマンド


| コマンド     | タイトル                |
|--------------|-------------------------|
| log job ship | ログの転送先Dropboxパス |



# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github_public"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `log cat job`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Id",
  			Desc:     "ジョブID",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Kind", Desc: "ログの種別", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
