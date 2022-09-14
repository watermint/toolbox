---
layout: release
title: リリースの変更点 109
lang: ja
---

# `リリース 109` から `リリース 110` までの変更点

# 追加されたコマンド


| コマンド                        | タイトル                                                      |
|---------------------------------|---------------------------------------------------------------|
| dev ci auth export              | CIビルドのためのデプロイトークンデータの書き出し              |
| services hellosign account info | アカウント情報を取得する                                      |
| util release install            | watermint toolboxをダウンロードし、パスにインストールします。 |



# 削除されたコマンド


| コマンド            | タイトル                                             |
|---------------------|------------------------------------------------------|
| dev ci auth connect | エンドツーエンドテストのための認証                   |
| dev ci auth import  | 環境変数はエンドツーエンドトークンをインポートします |



# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "NumFiles", Desc: "ファイル数.", Default: "1000", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "PreScan", Desc: "プリスキャンのデスティネーションパス", Default: "false", TypeName: "bool", ...},
  		&{Name: "SeqChunkSizeKb", Desc: "チャンクサイズをKiB単位でアップロード", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...},
  		... // 3 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "アップロード先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "SizeKb", Desc: "サイズ(KB)", Default: "1024", TypeName: "int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Github": "github_repo",
  		"Peer":   "github_public",
  	},
  	Services: {"github"},
  	IsSecret: true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `dev stage dbxfs`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "スキャンするパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev stage encoding`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Encoding", Desc: "エンコーディング", TypeName: "string"},
  		&{Name: "Name", Desc: "ファイル名", TypeName: "string"},
  		&{Name: "Path", Desc: "アップロード先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev stage http_range`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "ダウンロードするDropboxファイルのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "保存先のローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev stage scoped`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Individual": "dropbox_scoped_individual",
+ 		"Individual": "dropbox_individual",
- 		"Team":       "dropbox_scoped_team",
+ 		"Team":       "dropbox_team",
  	},
  	Services: {"dropbox", "dropbox_business"},
  	IsSecret: true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Team", Desc: "チーム向けのアカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev stage teamfolder`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `dev stage upload_append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "アップロードパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev test auth all`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `dev test setup massfiles`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Base", Desc: "Dropboxのベースパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "BatchSize", Desc: "バッチサイズ", Default: "1000", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Offset", Desc: "アップロードオフセット（ページ数省略）", Default: "0", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Source", Desc: "ソースファイル", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev test setup teamsharedlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "グループ名", TypeName: "string"},
  		&{Name: "NumLinksPerMember", Desc: "メンバーごとに作成するリンク数", Default: "5", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
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
# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Left":  "dropbox_scoped_individual",
+ 		"Left":  "dropbox_individual",
- 		"Right": "dropbox_scoped_individual",
+ 		"Right": "dropbox_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "left",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "LeftPath", Desc: "アカウントのルートからのパス (左)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "right",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "エクスポートするDropbox上のドキュメントパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Format", Desc: "エクスポート書式", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "保存先ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Format", Desc: "エクスポート書式", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "エクスポート先のローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Password", Desc: "共有リンクのパスワード", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "ドキュメントのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Path", Desc: "インポート先のパス", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "インポート先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Url", Desc: "URL", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.metadata.read"),
  			},
  		},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "IncludeMountedFolders", Desc: " Trueの場合は、マウントされたフォルダ（appフ\xe3\x82"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file lock acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "ロックするファイルパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "操作バッチサイズ", Default: "100", TypeName: "int", ...},
  		&{Name: "Path", Desc: "ロックを解除するためのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file lock batch acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "操作バッチサイズ", Default: "100", TypeName: "int", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file lock batch release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.metadata.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "ファイルへのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DryRun", Desc: "リハーサルを行います", Default: "true", TypeName: "bool", ...},
  		&{Name: "From", Desc: "統合するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "KeepEmptyFolder", Desc: "統合後に空となったフォルダを維持する", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "To", Desc: "統合するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "WithinSameNamespace", Desc: "ネームスペースを超えないように制御します. \xe3\x81"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Src", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file paper append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{Name: "Format", Desc: "入力フォーマット (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file paper create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{Name: "Format", Desc: "入力フォーマット (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file paper overwrite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{Name: "Format", Desc: "入力フォーマット (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file paper prepend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{Name: "Format", Desc: "入力フォーマット (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
- 		"Dst": "dropbox_scoped_individual",
+ 		"Dst": "dropbox_individual",
- 		"Src": "dropbox_scoped_individual",
+ 		"Src": "dropbox_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  				string("files.metadata.read"),
  			},
  		},
  		&{Name: "DstPath", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "src",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "SrcPath", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file restore all`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file revision download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "LocalPath", Desc: "ダウンロードしたファイルを保存するローカル\xe3"..., TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  			},
  		},
  		&{Name: "Revision", Desc: "ファイルリビジョン", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file revision list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "ファイルパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file revision restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "ファイルパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Revision", Desc: "ファイルリビジョン", TypeName: "string"},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "MaxResults", Desc: "返却するエントリーの最大数", Default: "25", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "指定されたファイルカテゴリに検索を限定しま\xe3"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file share info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "ファイル", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  			},
  		},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Depth", Desc: "すべてのファイルとフォルダの深さのフォルダ\xe3"..., Default: "2", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Path", Desc: "スキャンするパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 5 identical elements
  		&{Name: "NameNamePrefix", Desc: "名前によるフィルター. 名前の前方一致による\xe3\x83"...},
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "NameNamePrefix", Desc: "名前によるフィルター. 名前の前方一致による\xe3\x83"...},
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 7 identical elements
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{Name: "Overwrite", Desc: "ターゲットパス上に既存のファイルが存在する\xe5"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "監視対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Recursive", Desc: "パスを再起的に監視", Default: "false", TypeName: "bool", ...},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
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
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Force", Desc: "ファイリクエストを強制的に削除する.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.read"),
  				string("file_requests.write"),
  			},
  		},
  		&{Name: "Url", Desc: "ファイルリクエストのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ManagementType", Desc: "グループ管理タイプ. `company_managed` または `user_m"..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "ManagementType", Desc: "だれがこのグループを管理できるか (user_managed, "..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "グループ名リストのデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "GroupName", Desc: "グループ名", TypeName: "string"},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "GroupName", Desc: "グループ名称", TypeName: "string"},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "CurrentName", Desc: "現在のグループ名", TypeName: "string"},
  		&{Name: "NewName", Desc: "新しいグループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member batch suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "KeepData", Desc: "リンク先のデバイスにユーザーのデータを保持\xe3"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member batch unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.delete"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "TransferDestMember", Desc: "指定された場合は、指定ユーザーに削除するメ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "TransferNotifyAdminEmailOnError", Desc: "指定された場合は、転送時にエラーが発生した\xe6"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "WipeData", Desc: "指定した場合にはユーザーのデータがリンクさ\xe3"..., Default: "true", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.delete"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RevokeTeamShares", Desc: "指定した場合にはユーザーからチームが保有す\xe3"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("account_info.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "バッチ処理サイズ", Default: "100", TypeName: "int", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.content.write"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.content.write"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "チームメンバーのメールアドレス.", TypeName: "string"},
  		&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.permanent_delete"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `member folder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DstMemberEmail", Desc: "コピー先チームメンバーのメールアドレス", TypeName: "string"},
  		&{Name: "DstPath", Desc: "コピー先チームメンバーのパス. ルート (/) パス"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SrcMemberEmail", Desc: "送信元チームメンバーのメールアドレス.", TypeName: "string"},
  		&{Name: "SrcPath", Desc: "コピー元メンバーのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SilentInvite", Desc: "ウエルカムメールを送信しません (SSOとドメイ\xe3\x83"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "削除済メンバーを含めます.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Quota", Desc: "カスタムの容量制限 (1TB = 1024GB). 0の場合、容量\xe5"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.delete"),
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Silent", Desc: "招待メールを送信しません (SSOが必須となります)", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `member suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "KeepData", Desc: "リンク先のデバイスにユーザーのデータを保持\xe3"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "UpdateUnverified", Desc: "アカウントのメールアドレスが確認されていな\xe3"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member update invisible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member update visible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services asana team list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "WorkspaceName", Desc: "ワークスペースの名前または GID。 名前による\xe5\xae"...},
  		&{Name: "WorkspaceNamePrefix", Desc: "ワークスペースの名前または GID。 名前の前方\xe4\xb8"...},
  		&{Name: "WorkspaceNameSuffix", Desc: "ワークスペースの名前または GID。 名前の後方\xe4\xb8"...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services asana team project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "TeamName", Desc: "チーム名またはGID 名前による完全一致でフィル"...},
  		&{Name: "TeamNamePrefix", Desc: "チーム名またはGID 名前の前方一致によるフィル"...},
  		... // 4 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services asana team task list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "ProjectName", Desc: "プロジェクトの名前またはGID 名前による完全一"...},
  		&{Name: "ProjectNamePrefix", Desc: "プロジェクトの名前またはGID 名前の前方一致に"...},
  		... // 7 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services asana workspace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services asana workspace project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "WorkspaceName", Desc: "ワークスペースの名前または GID。 名前による\xe5\xae"...},
  		&{Name: "WorkspaceNamePrefix", Desc: "ワークスペースの名前または GID。 名前の前方\xe4\xb8"...},
  		&{Name: "WorkspaceNameSuffix", Desc: "ワークスペースの名前または GID。 名前の後方\xe4\xb8"...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services dropbox user feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Data", Desc: "データファイルのパス", TypeName: "Data"},
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{Name: "InputRaw", Desc: "Raw入力", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Range", Desc: "値がカバーする範囲をA1表記で表します. これは"..., TypeName: "string"},
  	},
  	GridDataInput:  {&{Name: "Data", Desc: "入力データファイル"}},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Range", Desc: "値がカバーする範囲をA1表記で表します. これは"..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Cols", Desc: "カラム数", Default: "26", TypeName: "int", ...},
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Rows", Desc: "行数", Default: "1000", TypeName: "int", ...},
  		&{Name: "Title", Desc: "シート名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "SheetId", Desc: "シートID (シートIDは `services google sheets sheet list` "..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "DateTimeRender", Desc: "日付、時間、および期間を出力でどのように表\xe7"..., Default: "serial", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets.readonly] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets.readonly")},
  		},
  		&{Name: "Range", Desc: "値がカバーする範囲をA1表記で表します. これは"..., TypeName: "string"},
  		&{Name: "ValueRender", Desc: "値を出力でどのように表現すべきか.", Default: "formatted", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "書き出したシートデータ"}},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Data", Desc: "データファイルのパス", TypeName: "Data"},
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{Name: "InputRaw", Desc: "Raw入力", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Range", Desc: "値がカバーする範囲をA1表記で表します. これは"..., TypeName: "string"},
  	},
  	GridDataInput:  {&{Name: "Data", Desc: "入力データファイル"}},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets.readonly] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets.readonly")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google sheets spreadsheet create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Title", Desc: "スプレッドシートのタイトル", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services slack conversation list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "slack"},
  	Services:       {"slack"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder leave`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "KeepCopy", Desc: "フォルダから抜ける時にフォルダの内容をコピ\xe3"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Message", Desc: "カスタム招待メッセージ", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "メンバーの共有フォルダのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "Silent", Desc: "招待メールを送信しない", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "フォルダメンバーのメールアドレス", TypeName: "string"},
  		&{Name: "LeaveCopy", Desc: "trueの場合、この共有フォルダのメンバーは、共"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "メンバーの共有フォルダのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder mount delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeMounted", Desc: "マウントされたフォルダーを含む.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AclUpdatePolicy", Desc: "共有フォルダーのアクセスコントロールリスト\xef"..., Default: "owner", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "MemberPolicy", Desc: "この共有フォルダーのメンバーになれる人.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "共有するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedLinkPolicy", Desc: "このフォルダー内の共有リンクを閲覧できる人.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "LeaveCopy", Desc: "trueの場合、この共有フォルダのメンバーは、共"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "共有解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
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
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "共有リンクを削除するファイルまたはフォルダ\xe3"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Password", Desc: "共有リンクのパスワード", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "共有リンクのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Password", Desc: "必要に応じてリンクのパスワードを指定.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "一つのイベントカテゴリのみを返すようなフィ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "File", Desc: "メールアドレスリストのファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("events.read"),
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "イベントのカテゴリ", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndDate", Desc: "終了日", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("events.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "一つのイベントカテゴリのみを返すようなフィ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("events.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "一つのイベントカテゴリのみを返すようなフィ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("events.read"),
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin group role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin group role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ExceptionGroup", Desc: "例外グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeNonAdmin", Desc: "管理者以外のメンバーをレポートに含める", Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberRoles", Desc: "メンバーと管理者の役割のマッピング", TypeName: "MemberRoles"},
  		&{Name: "MemberRolesFormat", Desc: "出力フォーマット"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "MemberRoles", Desc: "メンバーと管理者の役割のマッピング"}},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin role clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team admin role list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content legacypaper count`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content legacypaper export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FilterBy", Desc: "Paperドキュメントのフィルタリング方法（doc_crea"..., Default: "docs_created", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Format", Desc: "エクスポートファイル形式 (html/markdown)", Default: "html", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "エクスポートフォルダのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content legacypaper list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FilterBy", Desc: "Paperドキュメントのフィルタリング方法（doc_crea"..., Default: "docs_created", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team content member size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team content mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "MemberNamePrefix", Desc: "メンバーをフィルタリングします. 名前の前方\xe4\xb8"...},
  		&{Name: "MemberNameSuffix", Desc: "メンバーをフィルタリングします. 名前の後方\xe4\xb8"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sessions.list"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DeleteOnUnlink", Desc: "デバイスリンク解除時にファイルを削除します", Default: "false", TypeName: "bool", ...},
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("sessions.modify"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team filerequest clone`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("file_requests.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sessions.list"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team namespace summary`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team runas file batch copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas file sync batch up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 7 identical elements
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{Name: "Overwrite", Desc: "ターゲットパス上に既存のファイルが存在する\xe5"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder batch leave`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "KeepCopy", Desc: "フォルダから抜ける時にフォルダの内容をコピ\xe3"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder batch share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AclUpdatePolicy", Desc: "この共有フォルダのメンバーを追加・削除でき\xe3"..., Default: "owner", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "MemberPolicy", Desc: "この共有フォルダーのメンバーになれる人.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SharedLinkPolicy", Desc: "この共有フォルダー内のコンテンツに作成され\xe3"..., Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder batch unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "LeaveCopy", Desc: "trueの場合、この共有フォルダのメンバーは、共"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder isolate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Message", Desc: "カスタム招待メッセージ", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 4 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Silent", Desc: "招待メールを送信しない", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "LeaveCopy", Desc: "trueの場合、この共有フォルダのメンバーは、共"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 4 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeMounted", Desc: "マウントされたフォルダーを含む.", Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink cap expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "新しい有効期限の日付/時間", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink cap visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "新しい視認性設定", Default: "team_only", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink delete links`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink delete member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Visibility", Desc: "可視性によるリンクのフィルタリング (all/public/"..., Default: "all", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "新しい有効期限の日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink update password`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink update visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "新しい視認性設定", Default: "team_only", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.content.read"),
- 				string("team_data.content.write"),
  				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
  				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder partial replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
  				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `util monitor client`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MonitorInterval", Desc: "モニタリング間隔（秒）", Default: "10", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Name", Desc: "クライアント名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "SyncInterval", Desc: "Dropboxへの同期間隔（秒）", Default: "3600", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "SyncPath", Desc: "アップロード先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
