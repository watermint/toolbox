---
layout: release
title: リリースの変更点 92
lang: ja
---

# `リリース 92` から `リリース 93` までの変更点

# 追加されたコマンド


| コマンド                 | タイトル                                               |
|--------------------------|--------------------------------------------------------|
| file restore all         | 指定されたパス以下をリストアします                     |
| team content member size | チームフォルダや共有フォルダのメンバー数をカウントする |



# 削除されたコマンド


| コマンド      | タイトル                                |
|---------------|-----------------------------------------|
| file download | Dropboxからファイルをダウンロードします |
| file restore  | 指定されたパス以下をリストアします      |



# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Left":  "user_full",
+ 		"Left":  "dropbox_scoped_individual",
- 		"Right": "user_full",
+ 		"Right": "dropbox_scoped_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Left",
  			Desc:     "アカウントの別名 (左)",
  			Default:  "left",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "LeftPath", Desc: "アカウントのルートからのパス (左)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Right",
  			Desc:     "アカウントの別名 (右)",
  			Default:  "right",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "RightPath", Desc: "アカウントのルートからのパス (右)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file compare local`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file copy`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "Src", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file delete`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file export doc`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "エクスポートするDropbox上のドキュメントパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Format", Desc: "エクスポート書式", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "保存先ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file import batch url`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Path", Desc: "インポート先のパス", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file import url`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "インポート先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.write")},
  		},
  		&{Name: "Url", Desc: "URL", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file list`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "削除済みファイルを含める", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file merge`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DryRun", Desc: "リハーサルを行います", Default: "true", TypeName: "bool", ...},
  		&{Name: "From", Desc: "統合するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "KeepEmptyFolder", Desc: "統合後に空となったフォルダを維持する", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "To", Desc: "統合するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "WithinSameNamespace", Desc: "ネームスペースを超えないように制御します. \xe3\x81"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file mount list`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("sharing.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## 変更されたレポート: mount

```
  &dc_recipe.Report{
  	Name: "mount",
  	Desc: "このレポートは共有フォルダの一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 10 identical elements
  		&{Name: "policy_viewer_info", Desc: "だれが閲覧社情報を有効化・無効化できるか"},
  		&{Name: "owner_team_name", Desc: "このフォルダを所有するチームの名前"},
+ 		&{Name: "access_inheritance", Desc: "アクセス継承タイプ"},
  	},
  }
```
# コマンド仕様の変更: `file move`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "Src", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Dst": "user_full",
+ 		"Dst": "dropbox_scoped_individual",
- 		"Src": "user_full",
+ 		"Src": "dropbox_scoped_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Dst",
  			Desc:     "アカウントの別名 (宛先)",
  			Default:  "dst",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.metadata.read"), string("files.content.write")},
  		},
  		&{Name: "DstPath", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Src",
  			Desc:     "アカウントの別名 (元)",
  			Default:  "src",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "SrcPath", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file search content`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "指定されたファイルカテゴリに検索を限定しま\xe3"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file search name`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "指定されたファイルカテゴリに検索を限定しま\xe3"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file size`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Depth", Desc: "すべてのファイルとフォルダの深さのフォルダ\xe3"..., Default: "2", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Path", Desc: "スキャンするパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file sync down`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 5 identical elements
  		&{Name: "NameNamePrefix", Desc: "名前によるフィルター. 名前の前方一致による\xe3\x83"...},
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "SkipExisting", Desc: "既存ファイルをスキップします. 上書きしません.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file sync online`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "NameNamePrefix", Desc: "名前によるフィルター. 名前の前方一致による\xe3\x83"...},
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "SkipExisting", Desc: "既存ファイルをスキップします. 上書きしません.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Src", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file sync up`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 7 identical elements
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{Name: "Overwrite", Desc: "ターゲットパス上に既存のファイルが存在する\xe5"..., Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file watch`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "監視対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Recursive", Desc: "パスを再起的に監視", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder list`



## 変更されたレポート: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "このレポートは共有フォルダの一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 8 identical elements
  		&{Name: "policy_viewer_info", Desc: "だれが閲覧社情報を有効化・無効化できるか"},
  		&{Name: "owner_team_name", Desc: "このフォルダを所有するチームの名前"},
+ 		&{Name: "access_inheritance", Desc: "アクセス継承タイプ"},
  	},
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "event",
  	Title:   "イベントログ",
- 	Desc:    "",
+ 	Desc:    "リリース91以降では、`-start-time`または`-end-time`を`-24h`（24時間）または`-10m`（10分）のようなフォーマットで現在からの相対的な時間として解析します.\nもし、1時間ごとにイベントを取得した\xe3"...,
  	Remarks: "",
  	Path:    "team activity event",
  	... // 18 identical fields
  }
```
# コマンド仕様の変更: `team content member list`



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
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...},
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
+ 				string("files.metadata.read"), string("groups.read"), string("members.read"),
+ 				string("sharing.read"), string("team_data.member"),
+ 				string("team_data.team_space"), string("team_info.read"),
+ 			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
