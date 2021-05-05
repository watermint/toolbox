---
layout: release
title: Changes of Release 79
lang: ja
---

# `リリース 79` から `リリース 80` までの変更点

# 追加されたコマンド


| コマンド                  | タイトル                                         |
|---------------------------|--------------------------------------------------|
| member folder replication | フォルダを他のメンバーの個人フォルダに複製します |
| member update invisible   | メンバーへのディレクトリ制限を有効にします       |
| member update visible     | メンバーへのディレクトリ制限を無効にします       |
| teamfolder member add     | チームフォルダへのユーザー/グループの一括追加    |
| teamfolder member delete  | チームフォルダからのユーザー/グループの一括削除  |



# コマンド仕様の変更: `dev stage scoped`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Individual",
  			Desc:     "個人向けのアカウントの別名",
- 			Default:  "&{Individual [files.content.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{
  			Name:     "Team",
  			Desc:     "チーム向けのアカウントの別名",
- 			Default:  "&{Team [members.read team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("members.read"), string("team_info.read")},
  		},
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
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write"), string("groups.write"), string("sharing.read"), ...},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [groups.read groups.write] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("groups.read"), string("groups.write")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [groups.read groups.write] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("groups.read"), string("groups.write")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [groups.read groups.write] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("groups.read"), string("groups.write")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member file permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "チームメンバーのメールアドレス.", TypeName: "string"},
  		&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [files.permanent_delete team_data.member members.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []interface{}{string("files.permanent_delete"), string("team_data.member"), string("members.read")},
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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
