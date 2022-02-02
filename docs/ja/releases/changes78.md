---
layout: release
title: リリースの変更点 77
lang: ja
---

# `リリース 77` から `リリース 78` までの変更点

# 追加されたコマンド


| コマンド           | タイトル                   |
|--------------------|----------------------------|
| group folder list  | 各グループのフォルダを探す |
| member folder list | 各メンバーのフォルダを検索 |



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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```

## 変更されたレポート: namespace_member

```
  &dc_recipe.Report{
  	Name: "namespace_member",
  	Desc: "このレポートは名前空間とそのメンバー一覧を\xe5"...,
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "entry_is_inherited", Desc: "メンバーのアクセス権限が上位フォルダから継\xe6"...},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{
  			Name: "display_name",
  			Desc: strings.Join({
  				"\xe3",
- 				"\x82\xbbッションのタイプ (web_session, desktop_client, また\xe3",
- 				"\x81\xaf mobile_client)",
+ 				"\x83\x81ームメンバーの表示名.",
  			}, ""),
  		},
  		&{Name: "group_name", Desc: "グループ名称"},
  		&{Name: "invitee_email", Desc: "このフォルダに招待されたメールアドレス"},
  	},
  }
```
# コマンド仕様の変更: `team namespace member list`



## 変更されたレポート: namespace_member

```
  &dc_recipe.Report{
  	Name: "namespace_member",
  	Desc: "このレポートは名前空間とそのメンバー一覧を\xe5"...,
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "entry_is_inherited", Desc: "メンバーのアクセス権限が上位フォルダから継\xe6"...},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{
  			Name: "display_name",
  			Desc: strings.Join({
  				"\xe3",
- 				"\x82\xbbッションのタイプ (web_session, desktop_client, また\xe3",
- 				"\x81\xaf mobile_client)",
+ 				"\x83\x81ームメンバーの表示名.",
  			}, ""),
  		},
  		&{Name: "group_name", Desc: "グループ名称"},
  		&{Name: "invitee_email", Desc: "このフォルダに招待されたメールアドレス"},
  	},
  }
```
