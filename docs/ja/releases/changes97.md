---
layout: release
title: リリースの変更点 96
lang: ja
---

# `リリース 96` から `リリース 97` までの変更点

# 削除されたコマンド


| コマンド                | タイトル                                                |
|-------------------------|---------------------------------------------------------|
| dev ci artifact connect | CI成果物をアップロードするためのDropboxアカウントに接続 |
| dev test kvsfootprint   | KVSのメモリフットプリントをテストします                 |
| dev test monkey         | モンキーテスト                                          |



# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "指定されたファイルカテゴリに検索を限定しま\xe3"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"},
+ 		&{
+ 			Name:     "MaxResults",
+ 			Desc:     "返却するエントリーの最大数",
+ 			Default:  "25",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]any{"max": float64(100000), "min": float64(0), "value": float64(25)},
+ 		},
  		&{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllowLateUploads", Desc: "設定した場合、期限を過ぎてもアップロードを\xe8"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Deadline", Desc: "ファイルリクエストの締め切り.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Path", Desc: "ファイルをアップロードするDropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_requests.write")},
  		},
  		&{Name: "Title", Desc: "ファイルリクエストのタイトル", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_requests.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Force", Desc: "ファイリクエストを強制的に削除する.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_requests.read"), string("file_requests.write")},
  		},
  		&{Name: "Url", Desc: "ファイルリクエストのURL", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_requests.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("files.content.read"), string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sharing.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sharing.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Expires", Desc: "共有リンクの有効期限日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Password", Desc: "パスワード", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sharing.write")},
  		},
  		&{Name: "TeamOnly", Desc: "リンクがチームメンバーのみアクセスできます", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "共有リンクを削除するファイルまたはフォルダ\xe3"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sharing.write")},
  		},
  		&{Name: "Recursive", Desc: "フォルダ階層をたどって削除します", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Password", Desc: "共有リンクのパスワード", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("files.metadata.read"), string("sharing.read")},
  		},
  		&{Name: "Url", Desc: "共有リンクのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"user_file",
+ 				"scoped_individual",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sharing.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.team_space")},
  		},
  		&{Name: "SyncSetting", Desc: "チームフォルダの同期設定", Default: "default", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.team_space")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.team_space")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.team_space")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.team_space")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...},
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("sharing.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("team_data.team_space")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...},
  		&{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...},
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("sharing.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
