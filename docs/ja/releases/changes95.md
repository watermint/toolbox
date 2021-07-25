---
layout: release
title: リリースの変更点 94
lang: ja
---

# `リリース 94` から `リリース 95` までの変更点

# コマンド仕様の変更: `dev test setup teamsharedlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "グループ名", TypeName: "string"},
  		&{Name: "NumLinksPerMember", Desc: "メンバーごとに作成するリンク数", Default: "5", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{
  				string("files.content.write"),
- 				string("sharing.write"),
+ 				string("groups.read"),
- 				string("groups.read"),
+ 				string("sharing.write"),
  				string("team_data.member"),
  			},
  		},
  		&{Name: "Query", Desc: "クエリ", TypeName: "string"},
  		&{Name: "Seed", Desc: "シェアードリンクのシード値", Default: "0", TypeName: "int", ...},
  		&{Name: "Visibility", Desc: "ビジビリティ", Default: "random", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file export url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Format", Desc: "エクスポート書式", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "エクスポート先のローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "Password", Desc: "共有リンクのパスワード", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []interface{}{
- 				string("sharing.read"),
+ 				string("files.content.read"),
- 				string("files.content.read"),
+ 				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "ドキュメントのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []interface{}{
- 				string("files.metadata.read"),
+ 				string("files.content.write"),
- 				string("files.content.write"),
+ 				string("files.metadata.read"),
  			},
  		},
  		&{Name: "DstPath", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Src", Desc: "アカウントの別名 (元)", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "SrcPath", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
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
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{
  				string("files.permanent_delete"),
- 				string("team_data.member"),
+ 				string("members.read"),
- 				string("members.read"),
+ 				string("team_data.member"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content mount list`



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
  		... // 2 identical elements
  		&{Name: "MemberNamePrefix", Desc: "メンバーをフィルタリングします. 名前の前方\xe4\xb8"...},
  		&{Name: "MemberNameSuffix", Desc: "メンバーをフィルタリングします. 名前の後方\xe4\xb8"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("members.read"), string("sharing.read"), string("team_data.member"),
+ 				string("team_data.team_space"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content policy list`



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
  		&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...},
  		&{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...},
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("files.metadata.read"), string("groups.read"), string("members.read"),
+ 				string("sharing.read"), string("team_data.member"),
+ 				string("team_data.team_space"), string("team_info.read"),
+ 			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team device list`



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
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("members.read"), string("sessions.list")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team device unlink`



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
  		&{Name: "DeleteOnUnlink", Desc: "デバイスリンク解除時にファイルを削除します", Default: "false", TypeName: "bool", ...},
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("sessions.modify")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team namespace file list`



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
  		... // 5 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team namespace file size`



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
  		... // 6 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file list`



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
  		&{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file size`



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
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AdminGroupName", Desc: "管理者操作のための仮グループ名", Default: "watermint-toolbox-admin", TypeName: "string", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{
+ 				string("files.content.read"),
+ 				string("files.content.write"),
  				string("groups.read"),
  				string("groups.write"),
- 				string("files.content.read"),
- 				string("files.content.write"),
  				string("sharing.read"),
  				string("sharing.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AdminGroupName", Desc: "管理者操作のための仮グループ名", Default: "watermint-toolbox-admin", TypeName: "string", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{
+ 				string("files.content.read"),
+ 				string("files.content.write"),
  				string("groups.read"),
  				string("groups.write"),
- 				string("files.content.read"),
- 				string("files.content.write"),
  				string("sharing.read"),
  				string("sharing.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
