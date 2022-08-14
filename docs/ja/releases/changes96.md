---
layout: release
title: リリースの変更点 95
lang: ja
---

# `リリース 95` から `リリース 96` までの変更点

# 追加されたコマンド


| コマンド                        | タイトル                                                   |
|---------------------------------|------------------------------------------------------------|
| member feature                  | メンバーの機能設定一覧                                     |
| services dropbox user feature   | 現在のユーザーの機能設定の一覧                             |
| team content legacypaper count  | メンバー1人あたりのPaper文書の枚数                         |
| team content legacypaper export | チームメンバー全員のPaper文書をローカルパスにエクスポート. |
| team content legacypaper list   | チームメンバーのPaper文書リスト出力                        |



# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "business_file",
+ 		"Dst": "dropbox_scoped_team",
- 		"Src": "business_file",
+ 		"Src": "dropbox_scoped_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Dst",
  			Desc:    "宛先チーム; チームのファイルアクセス",
  			Default: "dst",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{},
  		},
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "Src",
  			Desc:    "元チーム; チームのファイルアクセス",
  			Default: "src",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Filter",
+ 			Desc:     "どのような種類の課題を返すかを示します.",
+ 			Default:  "assigned",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]any{
+ 				"options": []any{
+ 					string("assigned"), string("created"), string("mentioned"),
+ 					string("subscribed"), ...,
+ 				},
+ 			},
+ 		},
+ 		&{
+ 			Name:     "Labels",
+ 			Desc:     "カンマで区切られたラベル名のリスト.",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Since",
+ 			Desc:     "指定した時間以降に更新された通知のみを表示します.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]any{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "State",
+ 			Desc:     "返すべき課題の状態を示す.",
+ 			Default:  "open",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]any{"options": []any{string("open"), string("closed"), string("all")}},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team filerequest clone`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_requests.read"), string("members.read"), string("team_data.member")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("members.read"), string("sessions.list")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.member")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team namespace member list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllColumns", Desc: "全てのカラムを表示します", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sharing.read"), string("team_data.member"), string("team_info.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
