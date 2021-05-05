---
layout: release
title: リリースの変更点 62
lang: ja
---

# `リリース 62` から `リリース 63` までの変更点

# 追加されたコマンド


| コマンド            | タイトル                                   |
|---------------------|--------------------------------------------|
| team content member | チームフォルダや共有フォルダのメンバー一覧 |
| team content policy | チームフォルダと共有フォルダのポリシー一覧 |



# コマンド仕様の変更: `sharedfolder list`



## 変更されたレポート: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "is_inside_team_folder", Desc: "フォルダがチームフォルダに内包されているか\xe3"...},
  		&{Name: "is_team_folder", Desc: "このフォルダがチームフォルダであるかどうか"},
+ 		&{
+ 			Name: "policy_manage_access",
+ 			Desc: "このフォルダへメンバーを追加したり削除できるユーザー",
+ 		},
+ 		&{
+ 			Name: "policy_shared_link",
+ 			Desc: "このフォルダの共有リンクを誰が利用できるか",
+ 		},
  		&{Name: "policy_member", Desc: "だれがこの共有フォルダのメンバーに参加でき\xe3"...},
+ 		&{
+ 			Name: "policy_viewer_info",
+ 			Desc: "だれが閲覧社情報を有効化・無効化できるか",
+ 		},
+ 		&{
+ 			Name: "owner_team_id",
+ 			Desc: "このフォルダを所有するチームのチームID",
+ 		},
+ 		&{
+ 			Name: "owner_team_name",
+ 			Desc: "このフォルダを所有するチームの名前",
+ 		},
  	},
  }
```
