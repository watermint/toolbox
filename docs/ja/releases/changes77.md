---
layout: release
title: Changes of Release 76
lang: ja
---

# `リリース 76` から `リリース 77` までの変更点

# 追加されたコマンド


| コマンド               | タイトル                                                                                                                                    |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| image info             | 画像ファイルのEXIF情報を表示します                                                                                                          |
| member file permdelete | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します完全に削除については、https://www.dropbox.com/help/40 をご覧ください. |



# 削除されたコマンド


| コマンド       | タイトル                       |
|----------------|--------------------------------|
| dev test async | 非同期処理フレームワークテスト |



# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Path", Desc: "Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "Shard",
+ 			Desc:     "名前空間を分散させる共有フォルダの数.",
+ 			Default:  "1",
+ 			TypeName: "int",
+ 		},
  		&{Name: "SizeMaxKb", Desc: "最大ファイルサイズ (KiB).", Default: "2048", TypeName: "int", ...},
  		&{Name: "SizeMinKb", Desc: "最小ファイルサイズ (KiB).", Default: "0", TypeName: "int", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev stage teamfolder`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Peer",
- 			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [groups.write files.content.write] <nil>}",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: []interface{}{string("groups.write"), string("files.content.write")},
- 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space tea"...,
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
+ 			TypeAttr: []interface{}{
+ 				string("files.content.read"), string("files.content.write"),
+ 				string("groups.write"), string("sharing.read"), string("sharing.write"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services asana workspace list`



## 変更されたレポート: workspaces

```
  &dc_recipe.Report{
  	Name: "workspaces",
  	Desc: "ワークスペース",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "gid", Desc: "リソースのグローバルに一意な識別子を文字列\xe3"...},
  		&{Name: "resource_type", Desc: "このリソースのベースタイプ。"},
  		&{Name: "name", Desc: "ワークスペースの名前。"},
  		&{
  			Name: "is_organization",
- 			Desc: `	ワークスペースが組織であるかどうか。`,
+ 			Desc: "ワークスペースが組織であるかどうか。",
  		},
  	},
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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
