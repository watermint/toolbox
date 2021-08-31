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



# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Filter",
+ 			Desc:     "どのような種類の課題を返すかを示します.",
+ 			Default:  "assigned",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{
+ 				"options": []interface{}{
+ 					string("assigned"), string("created"), string("mentioned"),
+ 					string("subscribed"), ...,
+ 				},
+ 			},
+ 		},
+ 		&{
+ 			Name:     "Labels",
+ 			Desc:     "カンマで区切られたラベル名のリスト.",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Since",
+ 			Desc:     "指定した時間以降に更新された通知のみを表示します.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "State",
+ 			Desc:     "返すべき課題の状態を示す.",
+ 			Default:  "open",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("open"), string("closed"), string("all")}},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
