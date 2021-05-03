---
layout: release
title: Changes of Release 63
lang: ja
---

# `リリース 63` から `リリース 64` までの変更点

# 追加されたコマンド


| コマンド                           | タイトル                                                |
|------------------------------------|---------------------------------------------------------|
| config disable                     | 機能を無効化します.                                     |
| config enable                      | 機能を有効化します.                                     |
| config features                    | 利用可能なオプション機能一覧.                           |
| dev ci artifact connect            | CI成果物をアップロードするためのDropboxアカウントに接続 |
| dev ci auth connect                | エンドツーエンドテストのための認証                      |
| dev ci auth export                 | エンドツーエンドテストのトークンを出力します            |
| dev ci auth import                 | 環境変数はエンドツーエンドトークンをインポートします    |
| file dispatch local                | ローカルファイルを整理します                            |
| services github issue list         | 公開・プライベートGitHubレポジトリの課題一覧            |
| services github profile            | 認証したユーザーの情報を取得                            |
| services github release asset list | GitHubリリースの成果物一覧                              |
| services github release asset up   | GitHub リリースへ成果物をアップロードします             |
| services github release draft      | リリースの下書きを作成                                  |
| services github release list       | リリースの一覧                                          |
| services github tag create         | レポジトリにタグを作成します                            |
| version                            | バージョン情報                                          |



# 削除されたコマンド


| コマンド    | タイトル                           |
|-------------|------------------------------------|
| dev ci auth | エンドツーエンドテストのための認証 |



# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev async`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 		&{
+ 			Name:     "RunConcurrently",
+ 			Desc:     "並列実行",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: rows

```
  &dc_recipe.Report{
  	Name: "rows",
- 	Desc: "",
+ 	Desc: "このレポートはグループとメンバーを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "group_name", Desc: "グループ名称"},
  		&{Name: "group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
  		&{Name: "access_type", Desc: "グループにおけるユーザーの役割 (member/owner)"},
- 		&{Name: "account_id", Desc: "ユーザーアカウントのID"},
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		... // 2 identical elements
  	},
  }
```
# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev ci artifact up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "アップロード先Dropboxパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "アップロードするローカルファイルのパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "PeerName",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "deploy",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```

## 変更されたレポート: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "このレポートはアップロード結果の概要を出力します.",
  	Columns: {&{Name: "upload_start", Desc: "アップロード開始日時"}, &{Name: "upload_end", Desc: "アップロード終了日時"}, &{Name: "num_bytes", Desc: "合計アップロードサイズ (バイト)"}, &{Name: "num_files_error", Desc: "失敗またはエラーが発生したファイル数."}, ...},
  }
```

## 変更されたレポート: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `dev desktop install`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "InstallerUrl",
+ 			Desc:     "インストーラーのダウンロードURL",
+ 			Default:  "https://www.dropbox.com/download?full=1&os=win",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Silent",
+ 			Desc:     "サイレントインストーラーを利用します",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "SilentNoLaunch",
+ 			Desc:     "エンタープライズインストーラーを利用します",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev desktop start`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev desktop stop`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "WaitSeconds",
+ 			Desc:     "指定秒数後にアプリケーションの停止を試みます",
+ 			Default:  "60",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(60)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev desktop suspendupdate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Unsuspend",
+ 			Desc:     "指定すると一時停止を解除します (通常に戻し\xe3\x81"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "UpdaterName",
+ 			Desc:     "Dropbox Updaterの実行ファイル名",
+ 			Default:  "DropboxUpdate.exe",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "UpdaterPath",
+ 			Desc:     "Dropbox Updaterへのパス",
+ 			Default:  "C:/Program Files (x86)/Dropbox/Update",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev diag procmon`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev diag procmon",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -repository-path /LOCAL/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "procmonのログをアップロードするDropboxパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "ProcmonUrl",
+ 			Desc:     "Process MonitorのダウンロードURL",
+ 			Default:  "https://download.sysinternals.com/files/ProcessMonitor.zip",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepositoryPath",
+ 			Desc:     "Process Monitorの作業ディレクトリ",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "RetainLogs",
+ 			Desc:     "維持するProcmonのログ数",
+ 			Default:  "4",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(10000), "min": float64(0), "value": float64(4)},
+ 		},
+ 		&{
+ 			Name:     "RunUntil",
+ 			Desc:     "指定日時以降は実行をスキップ",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "procmonの実行待機時間",
+ 			Default:  "1800",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(10), "value": float64(1800)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Badge",
+ 			Desc:     "ビルド状態のバッジを含める",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "CommandPath",
+ 			Desc:     "コマンドマニュアルを作成する相対パス",
+ 			Default:  "doc/generated/",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Filename",
+ 			Desc:     "ファイル名",
+ 			Default:  "README.md",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "言語",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "MarkdownReadme",
+ 			Desc:     "READMEをMarkdownフォーマットで生成",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev dummy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Dest", Desc: "ダミーファイルの位置", TypeName: "string"},
+ 		&{
+ 			Name:     "MaxEntry",
+ 			Desc:     "最大エントリ数",
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "ダミーファイルエントリへのパス",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev echo`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Text",
+ 			Desc:     "エコーするテキストインポート先のパス",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev preflight`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "テストリソースへのパス",
+ 			Default:  "test/dev/resource.json",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev release publish",
- 	CliArgs:         "",
+ 	CliArgs:         "-artifact-path /LOCAL/PATH/TO/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ArtifactPath",
+ 			Desc:     "成果物へのパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Branch",
+ 			Desc:     "対象ブランチ",
+ 			Default:  "master",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "ConnGithub",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
+ 		},
+ 		&{
+ 			Name:     "SkipTests",
+ 			Desc:     "エンドツーエンドテストをスキップします.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "テストリソースへのパス",
+ 			Default:  "test/dev/resource.json",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev spec diff`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "FilePath",
+ 			Desc:     "出力先ファイルパス",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "言語",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Release1",
+ 			Desc:     "リリース名1",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Release2",
+ 			Desc:     "リリース名2",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev spec doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "FilePath",
+ 			Desc:     "ファイルパス",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "言語",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test monkey",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Distribution",
+ 			Desc:     "ファイル・フォルダの分布数",
+ 			Default:  "10000",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(10000)},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "カンマ区切りの拡張子一覧",
+ 			Default:  "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"...,
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "モンキーテストパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "モンキーテストの実施時間(秒)",
+ 			Default:  "10",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(1), "value": float64(10)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev test recipe`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "All",
+ 			Desc:     "全てのレシピをテストします",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Recipe",
+ 			Desc:     "テストするレシピ名",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Resource",
+ 			Desc:     "テスト用リソースへのパス",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Verbose",
+ 			Desc:     "テスト結果の詳細出力",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev test resources`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev util curl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "BufferSize",
+ 			Desc:     "バッファのサイズ",
+ 			Default:  "65536",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.097152e+06), "min": float64(1024), "value": float64(65536)},
+ 		},
+ 		&{
+ 			Name:     "Record",
+ 			Desc:     "テスト用に直接テストレコードを指定",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev util wait`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "指定秒数待機",
+ 			Default:  "1",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(604800), "min": float64(1), "value": float64(1)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Left",
+ 			Desc:     "アカウントの別名 (左)",
+ 			Default:  "left",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "LeftPath",
+ 			Desc:     "アカウントのルートからのパス (左)",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Right",
+ 			Desc:     "アカウントの別名 (右)",
+ 			Default:  "right",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "RightPath",
+ 			Desc:     "アカウントのルートからのパス (右)",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: diff

```
  &dc_recipe.Report{
  	Name:    "diff",
- 	Desc:    "",
+ 	Desc:    "このレポートはフォルダ間の差分を出力します.",
  	Columns: {&{Name: "diff_type", Desc: "差分のタイプ`file_content_diff`: コンテンツハッシ\xe3"...}, &{Name: "left_path", Desc: "左のパス"}, &{Name: "left_kind", Desc: "フォルダまたはファイル"}, &{Name: "left_size", Desc: "左ファイルのサイズ"}, ...},
  }
```
# コマンド仕様の変更: `file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox上のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "ローカルパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: diff

```
  &dc_recipe.Report{
  	Name:    "diff",
- 	Desc:    "",
+ 	Desc:    "このレポートはフォルダ間の差分を出力します.",
  	Columns: {&{Name: "diff_type", Desc: "差分のタイプ`file_content_diff`: コンテンツハッシ\xe3"...}, &{Name: "left_path", Desc: "左のパス"}, &{Name: "left_kind", Desc: "フォルダまたはファイル"}, &{Name: "left_size", Desc: "左ファイルのサイズ"}, ...},
  }
```

## 変更されたレポート: skip

```
  &dc_recipe.Report{
  	Name:    "skip",
- 	Desc:    "",
+ 	Desc:    "このレポートはフォルダ間の差分を出力します.",
  	Columns: {&{Name: "diff_type", Desc: "差分のタイプ`file_content_diff`: コンテンツハッシ\xe3"...}, &{Name: "left_path", Desc: "左のパス"}, &{Name: "left_kind", Desc: "フォルダまたはファイル"}, &{Name: "left_size", Desc: "左ファイルのサイズ"}, ...},
  }
```
# コマンド仕様の変更: `file copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "宛先のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "元のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "削除対象のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "file download",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/OF/FILE -local-path /LOCAL/PATH/TO/DOWNLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "ダウンロードするファイルパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "保存先ローカルパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートはファイルとフォルダのメタデータを出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
- 		&{
- 			Name: "path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "file export doc",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/FILE -local-path /LOCAL/PATH/TO/EXPORT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "エクスポートするDropbox上のドキュメントパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "保存先ローカルパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートはファイルのエクスポート結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "名称"},
- 		&{
- 			Name: "path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
- 		&{Name: "id", Desc: "ファイルへの一意なID"},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "これが共有フォルダのマウントポイントである\xe5"...},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
  		&{Name: "export_name", Desc: "エクスポートするファイル名."},
  		&{Name: "export_size", Desc: "エクスポートするファイルのサイズ."},
- 		&{
- 			Name: "export_hash",
- 			Desc: "エクスポートするファイルのコンテンツハッシュ.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "インポート先のパス",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.url", Desc: "ダウンロードするURL"},
  		&{Name: "input.path", Desc: "保存先パス (指定しないと`-path`オプションの指\xe5"...},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
  		&{Name: "result.tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "インポート先のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Url", Desc: "URL", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートはファイルとフォルダのメタデータを出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
- 		&{
- 			Name: "path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
  		&{Name: "revision", Desc: "ファイルの現在バージョンの一意な識別子"},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "削除済みファイルを含める",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "メディア情報を含める",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "パス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "再起的に一覧を実行",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: file_list

```
  &dc_recipe.Report{
  	Name: "file_list",
- 	Desc: "",
+ 	Desc: "このレポートはファイルとフォルダのメタデータを出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
- 		&{
- 			Name: "path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DryRun",
+ 			Desc:     "リハーサルを行います",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "From",
+ 			Desc:     "統合するパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "KeepEmptyFolder",
+ 			Desc:     "統合後に空となったフォルダを維持する",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "To",
+ 			Desc:     "統合するパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "WithinSameNamespace",
+ 			Desc:     "ネームスペースを超えないように制御します. \xe3\x81"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "宛先のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "元のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
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
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "アカウントの別名 (宛先)",
+ 			Default:  "dst",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "DstPath",
+ 			Desc:     "宛先のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "アカウントの別名 (元)",
+ 			Default:  "src",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "SrcPath",
+ 			Desc:     "元のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: replication_diff

```
  &dc_recipe.Report{
  	Name:    "replication_diff",
- 	Desc:    "",
+ 	Desc:    "このレポートはフォルダ間の差分を出力します.",
  	Columns: {&{Name: "diff_type", Desc: "差分のタイプ`file_content_diff`: コンテンツハッシ\xe3"...}, &{Name: "left_path", Desc: "左のパス"}, &{Name: "left_kind", Desc: "フォルダまたはファイル"}, &{Name: "left_size", Desc: "左ファイルのサイズ"}, ...},
  }
```
# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "file restore",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "パス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.path", Desc: "パス"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
  		&{Name: "result.tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "指定されたファイルカテゴリに検索を限定しま\xe3"...,
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "指定されたファイル拡張子に検索を限定します.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "検索対象とするユーザーのDropbox上のパス.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: matches

```
  &dc_recipe.Report{
  	Name:    "matches",
- 	Desc:    "",
+ 	Desc:    "このレポートは検索結果とハイライトテキストを出力します.",
  	Columns: {&{Name: "tag", Desc: "エントリーの種別"}, &{Name: "path_display", Desc: "パス"}, &{Name: "highlight_html", Desc: "HTML書式のハイライト済みテキスト"}},
  }
```
# コマンド仕様の変更: `file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "指定されたファイルカテゴリに検索を限定しま\xe3"...,
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "指定されたファイル拡張子に検索を限定します.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "検索対象とするユーザーのDropbox上のパス.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: matches

```
  &dc_recipe.Report{
  	Name:    "matches",
- 	Desc:    "",
+ 	Desc:    "このレポートは検索結果とハイライトテキストを出力します.",
  	Columns: {&{Name: "tag", Desc: "エントリーの種別"}, &{Name: "path_display", Desc: "パス"}, &{Name: "highlight_html", Desc: "HTML書式のハイライト済みテキスト"}},
  }
```
# コマンド仕様の変更: `file sync preflight up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync preflight up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "転送先のDropboxパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "ローカルファイルのパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```

## 変更されたレポート: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "このレポートはアップロード結果の概要を出力します.",
  	Columns: {&{Name: "upload_start", Desc: "アップロード開始日時"}, &{Name: "upload_end", Desc: "アップロード終了日時"}, &{Name: "num_bytes", Desc: "合計アップロードサイズ (バイト)"}, &{Name: "num_files_error", Desc: "失敗またはエラーが発生したファイル数."}, ...},
  }
```

## 変更されたレポート: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ChunkSizeKb",
+ 			Desc:     "アップロードチャンク容量(Kバイト)",
+ 			Default:  "153600",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)},
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "転送先のDropboxパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "ローカルファイルのパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```

## 変更されたレポート: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "このレポートはアップロード結果の概要を出力します.",
  	Columns: {&{Name: "upload_start", Desc: "アップロード開始日時"}, &{Name: "upload_end", Desc: "アップロード終了日時"}, &{Name: "num_bytes", Desc: "合計アップロードサイズ (バイト)"}, &{Name: "num_files_error", Desc: "失敗またはエラーが発生したファイル数."}, ...},
  }
```

## 変更されたレポート: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ChunkSizeKb",
+ 			Desc:     "アップロードチャンク容量(Kバイト)",
+ 			Default:  "153600",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)},
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "転送先のDropboxパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "ローカルファイルのパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Overwrite",
+ 			Desc:     "既存のファイルを上書きします",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```

## 変更されたレポート: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "このレポートはアップロード結果の概要を出力します.",
  	Columns: {&{Name: "upload_start", Desc: "アップロード開始日時"}, &{Name: "upload_end", Desc: "アップロード終了日時"}, &{Name: "num_bytes", Desc: "合計アップロードサイズ (バイト)"}, &{Name: "num_files_error", Desc: "失敗またはエラーが発生したファイル数."}, ...},
  }
```

## 変更されたレポート: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "ローカルファイルのパス"},
  		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.id", Desc: "ファイルへの一意なID"},
- 		&{
- 			Name: "result.tag",
- 			Desc: "エントリーの種別`file`, `folder`, または `deleted`",
- 		},
  		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "result.client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "result.server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "result.revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
  		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `file watch`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file watch",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/WATCH",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "監視対象のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "パスを再起的に監視",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "filerequest create",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "AllowLateUploads",
+ 			Desc:     "設定した場合、期限を過ぎてもアップロードを\xe8"...,
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Deadline",
+ 			Desc:     "ファイルリクエストの締め切り.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "ファイルをアップロードするDropbox上のパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Title",
+ 			Desc:     "ファイルリクエストのタイトル",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: file_request

```
  &dc_recipe.Report{
  	Name:    "file_request",
- 	Desc:    "",
+ 	Desc:    "このレポートはファイルリクエストの一覧を出力します.",
  	Columns: {&{Name: "id", Desc: "ファイルリクエストのID"}, &{Name: "url", Desc: "ファイルリクエストのURL"}, &{Name: "title", Desc: "ファイルリクエストのタイトル"}, &{Name: "created", Desc: "ファイルリクエストが作成された日時."}, ...},
  }
```
# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: deleted

```
  &dc_recipe.Report{
  	Name:    "deleted",
- 	Desc:    "",
+ 	Desc:    "このレポートはファイルリクエストの一覧を出力します.",
  	Columns: {&{Name: "id", Desc: "ファイルリクエストのID"}, &{Name: "url", Desc: "ファイルリクエストのURL"}, &{Name: "title", Desc: "ファイルリクエストのタイトル"}, &{Name: "created", Desc: "ファイルリクエストが作成された日時."}, ...},
  }
```
# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Force",
+ 			Desc:     "ファイリクエストを強制的に削除する.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Url", Desc: "ファイルリクエストのURL", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: deleted

```
  &dc_recipe.Report{
  	Name:    "deleted",
- 	Desc:    "",
+ 	Desc:    "このレポートはファイルリクエストの一覧を出力します.",
  	Columns: {&{Name: "id", Desc: "ファイルリクエストのID"}, &{Name: "url", Desc: "ファイルリクエストのURL"}, &{Name: "title", Desc: "ファイルリクエストのタイトル"}, &{Name: "created", Desc: "ファイルリクエストが作成された日時."}, ...},
  }
```
# コマンド仕様の変更: `filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: file_requests

```
  &dc_recipe.Report{
  	Name:    "file_requests",
- 	Desc:    "",
+ 	Desc:    "このレポートはファイルリクエストの一覧を出力します.",
  	Columns: {&{Name: "id", Desc: "ファイルリクエストのID"}, &{Name: "url", Desc: "ファイルリクエストのURL"}, &{Name: "title", Desc: "ファイルリクエストのタイトル"}, &{Name: "created", Desc: "ファイルリクエストが作成された日時."}, ...},
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ManagementType",
+ 			Desc:     "グループ管理タイプ. `company_managed` または `user_m"...,
+ 			Default:  "company_managed",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: added_group

```
  &dc_recipe.Report{
  	Name: "added_group",
- 	Desc: "",
+ 	Desc: "このレポートはチーム内のグループを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "グループ名称"},
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "グループ名リストのデータファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.name", Desc: "グループ名"},
  		&{Name: "result.group_name", Desc: "グループ名称"},
- 		&{Name: "result.group_id", Desc: "グループID"},
  		&{Name: "result.group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "result.member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: group

```
  &dc_recipe.Report{
  	Name: "group",
- 	Desc: "",
+ 	Desc: "このレポートはチーム内のグループを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "グループ名称"},
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "GroupName", Desc: "グループ名", TypeName: "string"},
+ 		&{
+ 			Name:     "MemberEmail",
+ 			Desc:     "メンバーのメールアドレス",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "メンバーのメールアドレス"},
  		&{Name: "result.group_name", Desc: "グループ名称"},
- 		&{Name: "result.group_id", Desc: "グループID"},
  		&{Name: "result.group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "result.member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "GroupName", Desc: "グループ名称", TypeName: "string"},
+ 		&{
+ 			Name:     "MemberEmail",
+ 			Desc:     "メンバーのメールアドレス",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "メンバーのメールアドレス"},
  		&{Name: "result.group_name", Desc: "グループ名称"},
- 		&{Name: "result.group_id", Desc: "グループID"},
  		&{Name: "result.group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "result.member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: group_member

```
  &dc_recipe.Report{
  	Name: "group_member",
- 	Desc: "",
+ 	Desc: "このレポートはグループとメンバーを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "group_name", Desc: "グループ名称"},
  		&{Name: "group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
  		&{Name: "access_type", Desc: "グループにおけるユーザーの役割 (member/owner)"},
- 		&{Name: "account_id", Desc: "ユーザーアカウントのID"},
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		... // 2 identical elements
  	},
  }
```
# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "CurrentName", Desc: "現在のグループ名", TypeName: "string"},
+ 		&{Name: "NewName", Desc: "新しいグループ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.new_name", Desc: "新しいグループ名"},
  		&{Name: "result.group_name", Desc: "グループ名称"},
- 		&{Name: "result.group_id", Desc: "グループID"},
  		&{Name: "result.group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "result.member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "目標日数",
+ 			Default:  "7",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(7)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "目標日数",
+ 			Default:  "28",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(28)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: log

```
  &dc_recipe.Report{
  	Name:    "log",
- 	Desc:    "",
+ 	Desc:    "ジョブ履歴を一覧するレポートです.",
  	Columns: {&{Name: "job_id", Desc: "ジョブID"}, &{Name: "app_version", Desc: "アプリケーションバージョン"}, &{Name: "recipe_name", Desc: "コマンド"}, &{Name: "time_start", Desc: "開始時刻"}, ...},
  }
```
# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job history ship",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "アップロード先Dropboxパス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 11 identical elements
  		&{Name: "result.revision", Desc: "ファイルの現在バージョンの一意な識別子"},
  		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `job loop`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job loop",
- 	CliArgs:         "",
+ 	CliArgs:         `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 17:58:38"`,
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IntervalSeconds",
+ 			Desc:     "実行間隔",
+ 			Default:  "180",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(1), "value": float64(180)},
+ 		},
+ 		&{
+ 			Name:     "QuitOnError",
+ 			Desc:     "エラー発生時に終了します.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Desc:     "実行するrunbookのパス",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Until",
+ 			Desc:     "指定日時まで実行します.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(false)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job run`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job run",
- 	CliArgs:         "",
+ 	CliArgs:         "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Fork",
+ 			Desc:     "ワークフローを実行する際にプロセスをフォー\xe3"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Desc:     "Runbookへのパス.",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "TimeoutSeconds",
+ 			Desc:     "指定時間を経過したしたためプロセスを終了し\xe3"...,
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(0), "value": float64(0)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "WipeData",
+ 			Desc:     "指定した場合にはユーザーのデータがリンクさ\xe3"...,
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "このレポートは処理結果を出力します.",
  	Columns: {&{Name: "status", Desc: "処理の状態"}, &{Name: "reason", Desc: "失敗またはスキップの理由"}, &{Name: "input.email", Desc: "アカウントのメールアドレス"}},
  }
```
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "RevokeTeamShares",
+ 			Desc:     "指定した場合にはユーザーからチームが保有す\xe3"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "このレポートは処理結果を出力します.",
  	Columns: {&{Name: "status", Desc: "処理の状態"}, &{Name: "reason", Desc: "失敗またはスキップの理由"}, &{Name: "input.email", Desc: "アカウントのメールアドレス"}},
  }
```
# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "SilentInvite",
+ 			Desc:     "ウエルカムメールを送信しません (SSOとドメイ\xe3\x83"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "アカウントの名前"},
  		&{Name: "input.surname", Desc: "アカウントの名字"},
- 		&{Name: "result.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "result.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "result.email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "result.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "result.given_name", Desc: "名"},
  		&{Name: "result.surname", Desc: "名字"},
- 		&{Name: "result.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "result.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "result.account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: member

```
  &dc_recipe.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "このレポートはメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "given_name", Desc: "名"},
  		&{Name: "surname", Desc: "名字"},
- 		&{Name: "familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{Name: "external_id", Desc: "このユーザーに関連づけられた外部ID"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  	},
  }
```
# コマンド仕様の変更: `member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: member_quota

```
  &dc_recipe.Report{
  	Name:    "member_quota",
- 	Desc:    "",
+ 	Desc:    "このレポートはチームメンバーのカスタム容量上限の設定を出力します.",
  	Columns: {&{Name: "email", Desc: "ユーザーのメールアドレス"}, &{Name: "quota", Desc: "カスタムの容量制限GB (1 TB = 1024 GB). 0の場合、容"...}},
  }
```
# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "Quota",
+ 			Desc:     "カスタムの容量制限 (1TB = 1024GB). 0の場合、容量\xe5"...,
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "このレポートは処理結果を出力します.",
  	Columns: {&{Name: "status", Desc: "処理の状態"}, &{Name: "reason", Desc: "失敗またはスキップの理由"}, &{Name: "input.email", Desc: "ユーザーのメールアドレス"}, &{Name: "input.quota", Desc: "カスタムの容量制限GB (1 TB = 1024 GB). 0の場合、容"...}, ...},
  }
```
# コマンド仕様の変更: `member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: usage

```
  &dc_recipe.Report{
  	Name:    "usage",
- 	Desc:    "",
+ 	Desc:    "このレポートはユーザーの現在のストレージ利用容量を出力します.",
  	Columns: {&{Name: "email", Desc: "アカウントのメールアドレス"}, &{Name: "used_gb", Desc: "このユーザーの合計利用スペース (in GB, 1GB = 1024"...}, &{Name: "used_bytes", Desc: "ユーザーの合計利用要領 (bytes)."}, &{Name: "allocation", Desc: "このユーザーの利用容量の付与先 (individual, or team)"}, ...},
  }
```
# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "Silent",
+ 			Desc:     "招待メールを送信しません (SSOが必須となります)",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{Name: "input.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "input.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "input.email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "input.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "input.given_name", Desc: "名"},
  		&{Name: "input.surname", Desc: "名字"},
- 		&{Name: "input.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "input.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "input.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "input.member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{
- 			Name: "input.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "input.account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "input.persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "input.joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "input.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "input.tag", Desc: "処理のタグ"},
- 		&{Name: "result.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "result.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "result.email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "result.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "result.given_name", Desc: "名"},
  		&{Name: "result.surname", Desc: "名字"},
- 		&{Name: "result.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "result.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "result.account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "宛先チーム; チームのファイルアクセス",
+ 			Default:  "dst",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "元チーム; チームのファイルアクセス",
+ 			Default:  "src",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "このレポートは処理結果を出力します.",
  	Columns: {&{Name: "status", Desc: "処理の状態"}, &{Name: "reason", Desc: "失敗またはスキップの理由"}, &{Name: "input.src_email", Desc: "転送元アカウントのメールアドレス"}, &{Name: "input.dst_email", Desc: "転送先アカウントのメールアドレス"}},
  }
```
# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "UpdateUnverified",
+ 			Desc:     "アカウントのメールアドレスが確認されていな\xe3"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.from_email", Desc: "現在のメールアドレス"},
  		&{Name: "input.to_email", Desc: "新しいメールアドレス"},
- 		&{Name: "result.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "result.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "result.email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "result.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "result.given_name", Desc: "名"},
  		&{Name: "result.surname", Desc: "名字"},
- 		&{Name: "result.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "result.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "result.account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.email", Desc: "チームメンバーのメールアドレス"},
  		&{Name: "input.external_id", Desc: "チームメンバーのExternal ID"},
- 		&{Name: "result.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "result.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "result.email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "result.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "result.given_name", Desc: "名"},
  		&{Name: "result.surname", Desc: "名字"},
- 		&{Name: "result.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "result.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "result.account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "アカウントの名前"},
  		&{Name: "input.surname", Desc: "アカウントの名字"},
- 		&{Name: "result.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "result.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "result.email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "result.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "result.given_name", Desc: "名"},
  		&{Name: "result.surname", Desc: "名字"},
- 		&{Name: "result.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "result.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "result.account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
- 	Desc: "",
+ 	Desc: "このレポートは共有フォルダの一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "共有フォルダのID"},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "親共有フォルダのID. このフィールドはフォルダが他の共有フォルダに含まれる場合のみ設定さ\xe3"...,
- 		},
  		&{Name: "name", Desc: "共有フォルダの名称"},
  		&{Name: "access_type", Desc: "ユーザーの共有ファイル・フォルダへのアクセ\xe3"...},
  		... // 5 identical elements
  		&{Name: "policy_member", Desc: "だれがこの共有フォルダのメンバーに参加でき\xe3"...},
  		&{Name: "policy_viewer_info", Desc: "だれが閲覧社情報を有効化・無効化できるか"},
- 		&{
- 			Name: "owner_team_id",
- 			Desc: "このフォルダを所有するチームのチームID",
- 		},
  		&{Name: "owner_team_name", Desc: "このフォルダを所有するチームの名前"},
  	},
  }
```
# コマンド仕様の変更: `sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: member

```
  &dc_recipe.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "このレポートは共有フォルダのメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "共有フォルダのID"},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "親共有フォルダのID. このフィールドはフォルダが他の共有フォルダに含まれる場合のみ設定さ\xe3"...,
- 		},
  		&{Name: "name", Desc: "共有フォルダの名称"},
  		&{Name: "path_lower", Desc: "共有フォルダのフルパス(小文字に変換済み)."},
  		... // 2 identical elements
  		&{Name: "access_type", Desc: "ユーザーの共有ファイル・フォルダへのアクセ\xe3"...},
  		&{Name: "is_inherited", Desc: "メンバーのアクセス権限が上位フォルダから継\xe6"...},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "group_name", Desc: "グループ名称"},
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "invitee_email", Desc: "このフォルダに招待されたメールアドレス"},
  	},
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Expires",
+ 			Desc:     "共有リンクの有効期限日時",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Password",
+ 			Desc:     "パスワード",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "パス",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "TeamOnly",
+ 			Desc:     "リンクがチームメンバーのみアクセスできます",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: created

```
  &dc_recipe.Report{
  	Name:    "created",
- 	Desc:    "",
+ 	Desc:    "このレポートは共有リンクの一覧を出力します.",
  	Columns: {&{Name: "id", Desc: "ファイルまたはフォルダへのリンクのID"}, &{Name: "tag", Desc: "エントリーの種別 (file, または folder)"}, &{Name: "url", Desc: "共有リンクのURL."}, &{Name: "name", Desc: "リンク先ファイル名称"}, ...},
  }
```
# コマンド仕様の変更: `sharedlink delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "共有リンクを削除するファイルまたはフォルダ\xe3"...,
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "フォルダ階層をたどって削除します",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{Name: "input.id", Desc: "ファイルまたはフォルダへのリンクのID"},
  		&{Name: "input.tag", Desc: "エントリーの種別 (file, または folder)"},
  		&{Name: "input.url", Desc: "共有リンクのURL."},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedlink file list",
- 	CliArgs:         "",
+ 	CliArgs:         "-url SHAREDLINK_URL",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Password",
+ 			Desc:     "共有リンクのパスワード",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Url",
+ 			Desc:     "共有リンクのURL",
+ 			TypeName: "domain.dropbox.model.mo_url.url_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: file_list

```
  &dc_recipe.Report{
  	Name: "file_list",
- 	Desc: "",
+ 	Desc: "このレポートはファイルとフォルダのメタデータを出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
- 		&{
- 			Name: "path_lower",
- 			Desc: "パス (すべて小文字に変換). これは常にスラッシュで始まります.",
- 		},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "このファイルを含む共有フォルダのID.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "このレポートは共有リンクの一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "ファイルまたはフォルダへのリンクのID"},
  		&{Name: "tag", Desc: "エントリーの種別 (file, または folder)"},
  		&{Name: "url", Desc: "共有リンクのURL."},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "終了日時 (該当同時刻を含まない).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "メールアドレスリストのファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "開始日時 (該当時刻を含む)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: combined

```
  &dc_recipe.Report{
  	Name:    "combined",
- 	Desc:    "",
+ 	Desc:    "このレポートはDropbox Businessのアクティビティログとほぼ互換性のあるアクティビティレポートを出力します.",
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```

## 変更されたレポート: user

```
  &dc_recipe.Report{
  	Name:    "user",
- 	Desc:    "",
+ 	Desc:    "このレポートはDropbox Businessのアクティビティログとほぼ互換性のあるアクティビティレポートを出力します.",
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "イベントのカテゴリ",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndDate",
+ 			Desc:     "終了日",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{Name: "StartDate", Desc: "開始日", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: event

```
  &dc_recipe.Report{
  	Name:    "event",
- 	Desc:    "",
+ 	Desc:    "このレポートはDropbox Businessのアクティビティログとほぼ互換性のあるアクティビティレポートを出力します.",
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "終了日時 (該当同時刻を含まない).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "開始日時 (該当時刻を含む)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: event

```
  &dc_recipe.Report{
  	Name:    "event",
- 	Desc:    "",
+ 	Desc:    "このレポートはDropbox Businessのアクティビティログとほぼ互換性のあるアクティビティレポートを出力します.",
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "終了日時 (該当同時刻を含まない).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "開始日時 (該当時刻を含む)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: user

```
  &dc_recipe.Report{
  	Name:    "user",
- 	Desc:    "",
+ 	Desc:    "このレポートはDropbox Businessのアクティビティログとほぼ互換性のあるアクティビティレポートを出力します.",
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```

## 変更されたレポート: user_summary

```
  &dc_recipe.Report{
  	Name: "user_summary",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.user", Desc: "ユーザーのメールアドレス"},
- 		&{Name: "result.user", Desc: "ユーザーのメールアドレス"},
  		&{Name: "result.logins", Desc: "ログインのアクティビティ数"},
  		&{Name: "result.devices", Desc: "デバイスのアクティビティ数"},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `team content member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: membership

```
  &dc_recipe.Report{
  	Name:    "membership",
- 	Desc:    "",
+ 	Desc:    "このレポートは共有フォルダまたはチームフォルダと、そのメンバーを一覧できます. フォルダに複数メンバーがいる場合には、メンバーは1行ずつ出力されます.",
  	Columns: {&{Name: "path", Desc: "パス"}, &{Name: "folder_type", Desc: "フォルダの種別. (`team_folder`: チームフォルダま\xe3"...}, &{Name: "owner_team_name", Desc: "このフォルダを所有するチームの名前"}, &{Name: "access_type", Desc: "このフォルダに対するユーザーのアクセスレベル"}, ...},
  }
```

## 変更されたレポート: no_member

```
  &dc_recipe.Report{
  	Name:    "no_member",
- 	Desc:    "",
+ 	Desc:    "このレポートはメンバーのいないフォルダの一覧を出力します.",
  	Columns: {&{Name: "owner_team_name", Desc: "このフォルダを所有するチームの名前"}, &{Name: "path", Desc: "パス"}, &{Name: "folder_type", Desc: "フォルダの種別. (`team_folder`: チームフォルダま\xe3"...}},
  }
```
# コマンド仕様の変更: `team content policy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: policy

```
  &dc_recipe.Report{
  	Name:    "policy",
- 	Desc:    "",
+ 	Desc:    "このレポートでは共有フォルダならびにチームフォルダについて、現在のポリシー設定を一覧できます.",
  	Columns: {&{Name: "path", Desc: "パス"}, &{Name: "is_team_folder", Desc: "チームフォルダまたはチームフォルダ下のフォ\xe3"...}, &{Name: "owner_team_name", Desc: "このフォルダを所有するチームの名前"}, &{Name: "policy_manage_access", Desc: "このフォルダへメンバーを追加したり削除でき\xe3"...}, ...},
  }
```
# コマンド仕様の変更: `team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: device

```
  &dc_recipe.Report{
  	Name: "device",
- 	Desc: "",
+ 	Desc: "このレポートではチーム内の既存セッションとメンバー情報を一覧できます.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "名"},
  		&{Name: "surname", Desc: "名字"},
- 		&{Name: "familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{Name: "external_id", Desc: "このユーザーに関連づけられた外部ID"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
  		&{Name: "device_tag", Desc: "セッションのタイプ (web_session, desktop_client, また"...},
  		&{Name: "id", Desc: "セッションID"},
  		... // 16 identical elements
  	},
  }
```
# コマンド仕様の変更: `team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DeleteOnUnlink",
+ 			Desc:     "デバイスリンク解除時にファイルを削除します",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "input.given_name", Desc: "名"},
  		&{Name: "input.surname", Desc: "名字"},
- 		&{Name: "input.familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "input.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "input.abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "input.external_id",
- 			Desc: "このユーザーに関連づけられた外部ID",
- 		},
- 		&{Name: "input.account_id", Desc: "ユーザーのアカウントID"},
  		&{Name: "input.device_tag", Desc: "セッションのタイプ (web_session, desktop_client, また"...},
  		&{Name: "input.id", Desc: "セッションID"},
  		... // 16 identical elements
  	},
  }
```
# コマンド仕様の変更: `team diag explorer`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "All",
+ 			Desc:     "追加のレポートを含める",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Dropbox Business ファイルアクアセス",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Info",
+ 			Desc:     "Dropbox Business 情報アクセス",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 		&{
+ 			Name:     "Mgmt",
+ 			Desc:     "Dropbox Business 管理",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: device

```
  &dc_recipe.Report{
  	Name: "device",
- 	Desc: "",
+ 	Desc: "このレポートではチーム内の既存セッションとメンバー情報を一覧できます.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "名"},
  		&{Name: "surname", Desc: "名字"},
- 		&{Name: "familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{Name: "external_id", Desc: "このユーザーに関連づけられた外部ID"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
  		&{Name: "device_tag", Desc: "セッションのタイプ (web_session, desktop_client, また"...},
  		&{Name: "id", Desc: "セッションID"},
  		... // 16 identical elements
  	},
  }
```

## 変更されたレポート: feature

```
  &dc_recipe.Report{
  	Name:    "feature",
- 	Desc:    "",
+ 	Desc:    "このレポートはチームの機能と設定を一覧します.",
  	Columns: {&{Name: "upload_api_rate_limit", Desc: "毎月利用可能なアップロードAPIコール回数"}, &{Name: "upload_api_rate_limit_count", Desc: "この月に利用されたアップロードAPIコール回数"}, &{Name: "has_team_shared_dropbox", Desc: "このチームが共有されたチームルートを持って\xe3"...}, &{Name: "has_team_file_events", Desc: "このチームがファイルイベント機能を持ってい\xe3"...}, ...},
  }
```

## 変更されたレポート: file_request

```
  &dc_recipe.Report{
  	Name: "file_request",
- 	Desc: "",
+ 	Desc: "このレポートはチームメンバーのもつファイルリクエストを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{
- 			Name: "account_id",
- 			Desc: "ファイルリクエスト所有者のアカウントID",
- 		},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ファイルリクエスト所有者のチームメンバーとしてのID",
- 		},
  		&{Name: "email", Desc: "ファイルリクエスト所有者のメールアドレス"},
  		&{Name: "status", Desc: "ファイルリクエスト所有者ユーザーの状態 (activ"...},
  		&{Name: "surname", Desc: "ファイルリクエスト所有者の名字"},
  		&{Name: "given_name", Desc: "ファイルリクエスト所有者の名"},
- 		&{Name: "file_request_id", Desc: "ファイルリクエストID"},
  		&{Name: "url", Desc: "ファイルリクエストのURL"},
  		&{Name: "title", Desc: "ファイルリクエストのタイトル"},
  		... // 6 identical elements
  	},
  }
```

## 変更されたレポート: group

```
  &dc_recipe.Report{
  	Name: "group",
- 	Desc: "",
+ 	Desc: "このレポートはチーム内のグループを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "グループ名称"},
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " グループの外部IDこの任意のIDは管理者がグループに付加できます",
- 		},
  		&{Name: "member_count", Desc: "グループ内のメンバー数"},
  	},
  }
```

## 変更されたレポート: group_member

```
  &dc_recipe.Report{
  	Name: "group_member",
- 	Desc: "",
+ 	Desc: "このレポートはグループとメンバーを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "group_id", Desc: "グループID"},
  		&{Name: "group_name", Desc: "グループ名称"},
  		&{Name: "group_management_type", Desc: "だれがこのグループを管理できるか (user_managed, "...},
  		&{Name: "access_type", Desc: "グループにおけるユーザーの役割 (member/owner)"},
- 		&{Name: "account_id", Desc: "ユーザーアカウントのID"},
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		... // 2 identical elements
  	},
  }
```

## 変更されたレポート: info

```
  &dc_recipe.Report{
  	Name:    "info",
- 	Desc:    "",
+ 	Desc:    "このレポートはチームの情報を一覧します.",
  	Columns: {&{Name: "name", Desc: "チームの名称"}, &{Name: "team_id", Desc: "チームのID"}, &{Name: "num_licensed_users", Desc: "このチームで利用可能なライセンス数"}, &{Name: "num_provisioned_users", Desc: "招待済みアカウント数 (アクティブメンバーま\xe3\x81"...}, ...},
  }
```

## 変更されたレポート: linked_app

```
  &dc_recipe.Report{
  	Name: "linked_app",
- 	Desc: "",
+ 	Desc: "このレポートは接続済みアプリケーションと利用ユーザーを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "given_name", Desc: "名"},
  		&{Name: "surname", Desc: "名字"},
- 		&{Name: "familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{Name: "external_id", Desc: "このユーザーに関連づけられた外部ID"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
- 		&{Name: "app_id", Desc: "アプリケーションの固有ID"},
  		&{Name: "app_name", Desc: "アプリケーション名称"},
  		&{Name: "is_app_folder", Desc: "アプリケーションが専用フォルダにリンクする\xe3"...},
  		... // 3 identical elements
  	},
  }
```

## 変更されたレポート: member

```
  &dc_recipe.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "このレポートはメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "email_verified", Desc: "trueの場合、ユーザーのメールアドレスはユーザ"...},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "given_name", Desc: "名"},
  		&{Name: "surname", Desc: "名字"},
- 		&{Name: "familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{
- 			Name: "member_folder_id",
- 			Desc: "ユーザールートフォルダの名前空間ID.",
- 		},
- 		&{Name: "external_id", Desc: "このユーザーに関連づけられた外部ID"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で利用する一意なIDです.",
- 		},
  		&{Name: "joined_on", Desc: "メンバーがチームに参加した日時."},
  		&{Name: "role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  	},
  }
```

## 変更されたレポート: member_quota

```
  &dc_recipe.Report{
  	Name:    "member_quota",
- 	Desc:    "",
+ 	Desc:    "このレポートはチームメンバーのカスタム容量上限の設定を出力します.",
  	Columns: {&{Name: "email", Desc: "ユーザーのメールアドレス"}, &{Name: "quota", Desc: "カスタムの容量制限GB (1 TB = 1024 GB). 0の場合、容"...}},
  }
```

## 変更されたレポート: namespace

```
  &dc_recipe.Report{
  	Name: "namespace",
- 	Desc: "",
+ 	Desc: "このレポートはチームの名前空間を一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "名前空間の名称"},
- 		&{Name: "namespace_id", Desc: "名前空間ID"},
  		&{Name: "namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
  		&{Name: "team_member_id", Desc: "メンバーフォルダまたはアプリフォルダである\xe5"...},
  	},
  }
```

## 変更されたレポート: namespace_file

```
  &dc_recipe.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "このレポートはチームの名前空間を一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
- 		&{Name: "namespace_id", Desc: "名前空間ID"},
  		&{Name: "namespace_name", Desc: "名前空間の名称"},
  		&{Name: "namespace_member_email", Desc: "これがチームメンバーフォルダまたはアプリフ\xe3"...},
- 		&{Name: "file_id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "設定されている場合、共有フォルダに内包されています.",
- 		},
  	},
  }
```

## 変更されたレポート: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.name", Desc: "名前空間の名称"},
- 		&{Name: "input.namespace_id", Desc: "名前空間ID"},
  		&{Name: "input.namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
- 		},
- 		&{Name: "result.namespace_name", Desc: "名前空間の名称"},
- 		&{Name: "result.namespace_id", Desc: "名前空間ID"},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
- 		},
  		&{Name: "result.path", Desc: "フォルダへのパス"},
  		&{Name: "result.count_file", Desc: "このフォルダに含まれるファイル数"},
  		... // 4 identical elements
  	},
  }
```

## 変更されたレポート: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "ファイルまたはフォルダへのリンクのID",
- 		},
  		&{Name: "tag", Desc: "エントリーの種別 (file, または folder)"},
  		&{Name: "url", Desc: "共有リンクのURL."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "パス (すべて小文字に変換)."},
  		&{Name: "visibility", Desc: "共有リンクの開示範囲"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		... // 2 identical elements
  	},
  }
```

## 変更されたレポート: usage

```
  &dc_recipe.Report{
  	Name:    "usage",
- 	Desc:    "",
+ 	Desc:    "このレポートはユーザーの現在のストレージ利用容量を出力します.",
  	Columns: {&{Name: "email", Desc: "アカウントのメールアドレス"}, &{Name: "used_gb", Desc: "このユーザーの合計利用スペース (in GB, 1GB = 1024"...}, &{Name: "used_bytes", Desc: "ユーザーの合計利用要領 (bytes)."}, &{Name: "allocation", Desc: "このユーザーの利用容量の付与先 (individual, or team)"}, ...},
  }
```
# コマンド仕様の変更: `team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: feature

```
  &dc_recipe.Report{
  	Name:    "feature",
- 	Desc:    "",
+ 	Desc:    "このレポートはチームの機能と設定を一覧します.",
  	Columns: {&{Name: "upload_api_rate_limit", Desc: "毎月利用可能なアップロードAPIコール回数"}, &{Name: "upload_api_rate_limit_count", Desc: "この月に利用されたアップロードAPIコール回数"}, &{Name: "has_team_shared_dropbox", Desc: "このチームが共有されたチームルートを持って\xe3"...}, &{Name: "has_team_file_events", Desc: "このチームがファイルイベント機能を持ってい\xe3"...}, ...},
  }
```
# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: file_request

```
  &dc_recipe.Report{
  	Name: "file_request",
- 	Desc: "",
+ 	Desc: "このレポートはチームメンバーのもつファイルリクエストを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{
- 			Name: "account_id",
- 			Desc: "ファイルリクエスト所有者のアカウントID",
- 		},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ファイルリクエスト所有者のチームメンバーとしてのID",
- 		},
  		&{Name: "email", Desc: "ファイルリクエスト所有者のメールアドレス"},
  		&{Name: "status", Desc: "ファイルリクエスト所有者ユーザーの状態 (activ"...},
  		&{Name: "surname", Desc: "ファイルリクエスト所有者の名字"},
  		&{Name: "given_name", Desc: "ファイルリクエスト所有者の名"},
- 		&{Name: "file_request_id", Desc: "ファイルリクエストID"},
  		&{Name: "url", Desc: "ファイルリクエストのURL"},
  		&{Name: "title", Desc: "ファイルリクエストのタイトル"},
  		... // 6 identical elements
  	},
  }
```
# コマンド仕様の変更: `team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: info

```
  &dc_recipe.Report{
  	Name:    "info",
- 	Desc:    "",
+ 	Desc:    "このレポートはチームの情報を一覧します.",
  	Columns: {&{Name: "name", Desc: "チームの名称"}, &{Name: "team_id", Desc: "チームのID"}, &{Name: "num_licensed_users", Desc: "このチームで利用可能なライセンス数"}, &{Name: "num_provisioned_users", Desc: "招待済みアカウント数 (アクティブメンバーま\xe3\x81"...}, ...},
  }
```
# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: linked_app

```
  &dc_recipe.Report{
  	Name: "linked_app",
- 	Desc: "",
+ 	Desc: "このレポートは接続済みアプリケーションと利用ユーザーを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "given_name", Desc: "名"},
  		&{Name: "surname", Desc: "名字"},
- 		&{Name: "familiar_name", Desc: "ロケール依存の名前"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
- 		&{Name: "abbreviated_name", Desc: "ユーザーの省略名称"},
- 		&{Name: "external_id", Desc: "このユーザーに関連づけられた外部ID"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
- 		&{Name: "app_id", Desc: "アプリケーションの固有ID"},
  		&{Name: "app_name", Desc: "アプリケーション名称"},
  		&{Name: "is_app_folder", Desc: "アプリケーションが専用フォルダにリンクする\xe3"...},
  		... // 3 identical elements
  	},
  }
```
# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "指定された場合、削除済みのファイルやフォル\xe3"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "指定された場合、JSONレポートに写真や動画ファ"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "指定された場合、チームメンバーのフォルダを\xe5"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "Trueの場合、共有フォルダを含めます",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "Trueの場合、チームフォルダを含めます",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Desc:     "指定された名前に一致するフォルダのみを一覧\xe3"...,
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: namespace_file

```
  &dc_recipe.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "このレポートはチームの名前空間を一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
- 		&{Name: "namespace_id", Desc: "名前空間ID"},
  		&{Name: "namespace_name", Desc: "名前空間の名称"},
  		&{Name: "namespace_member_email", Desc: "これがチームメンバーフォルダまたはアプリフ\xe3"...},
- 		&{Name: "file_id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "設定されている場合、共有フォルダに内包されています.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Depth",
+ 			Desc:     "フォルダ階層数の指定",
+ 			Default:  "1",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
+ 		},
+ 		&{
+ 			Name:     "IncludeAppFolder",
+ 			Desc:     "Trueの場合、アプリフォルダを含めます",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "Trueの場合、チームメンバーフォルダを含めます",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "Trueの場合、共有フォルダを含めます",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "Trueの場合、チームフォルダを含めます",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Desc:     "指定された名前に一致するフォルダのみを一覧\xe3"...,
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.name", Desc: "名前空間の名称"},
- 		&{Name: "input.namespace_id", Desc: "名前空間ID"},
  		&{Name: "input.namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
- 		},
- 		&{Name: "result.namespace_name", Desc: "名前空間の名称"},
- 		&{Name: "result.namespace_id", Desc: "名前空間ID"},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
- 		},
  		&{Name: "result.path", Desc: "フォルダへのパス"},
  		&{Name: "result.count_file", Desc: "このフォルダに含まれるファイル数"},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: namespace

```
  &dc_recipe.Report{
  	Name: "namespace",
- 	Desc: "",
+ 	Desc: "このレポートはチームの名前空間を一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "名前空間の名称"},
- 		&{Name: "namespace_id", Desc: "名前空間ID"},
  		&{Name: "namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
  		&{Name: "team_member_id", Desc: "メンバーフォルダまたはアプリフォルダである\xe5"...},
  	},
  }
```
# コマンド仕様の変更: `team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "AllColumns",
+ 			Desc:     "全てのカラムを表示します",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: namespace_member

```
  &dc_recipe.Report{
  	Name: "namespace_member",
- 	Desc: "",
+ 	Desc: "このレポートは名前空間とそのメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_name", Desc: "名前空間の名称"},
- 		&{Name: "namespace_id", Desc: "名前空間ID"},
  		&{Name: "namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
  		&{Name: "entry_access_type", Desc: "ユーザーの共有ファイル・フォルダへのアクセ\xe3"...},
  		... // 5 identical elements
  	},
  }
```
# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "出力するリンクを可視性にてフィルターします "...,
+ 			Default:  "public",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "ファイルまたはフォルダへのリンクのID",
- 		},
  		&{Name: "tag", Desc: "エントリーの種別 (file, または folder)"},
  		&{Name: "url", Desc: "共有リンクのURL."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "パス (すべて小文字に変換)."},
  		&{Name: "visibility", Desc: "共有リンクの開示範囲"},
- 		&{Name: "account_id", Desc: "ユーザーのアカウントID"},
- 		&{Name: "team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		... // 2 identical elements
  	},
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "At",
+ 			Desc:     "新しい有効期限の日時",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "新しい有効期限までの日時",
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "対象となるリンクの公開範囲",
+ 			Default:  "public",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name:    "skipped",
- 	Desc:    "",
+ 	Desc:    "このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します.",
  	Columns: {&{Name: "tag", Desc: "エントリーの種別 (file, または folder)"}, &{Name: "url", Desc: "共有リンクのURL."}, &{Name: "name", Desc: "リンク先ファイル名称"}, &{Name: "expires", Desc: "有効期限 (設定されている場合)"}, ...},
  }
```

## 変更されたレポート: updated

```
  &dc_recipe.Report{
  	Name: "updated",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{
- 			Name: "input.shared_link_id",
- 			Desc: "ファイルまたはフォルダへのリンクのID",
- 		},
  		&{Name: "input.tag", Desc: "エントリーの種別 (file, または folder)"},
  		&{Name: "input.url", Desc: "共有リンクのURL."},
  		... // 2 identical elements
  		&{Name: "input.path_lower", Desc: "パス (すべて小文字に変換)."},
  		&{Name: "input.visibility", Desc: "共有リンクの開示範囲"},
- 		&{Name: "input.account_id", Desc: "ユーザーのアカウントID"},
- 		&{Name: "input.team_member_id", Desc: "チームにおけるメンバーのID"},
  		&{Name: "input.email", Desc: "ユーザーのメールアドレス"},
  		&{Name: "input.status", Desc: "チームにおけるメンバーのステータス(active/invit"...},
  		&{Name: "input.surname", Desc: "リンク所有者の名字"},
  		&{Name: "input.given_name", Desc: "リンク所有者の名"},
- 		&{Name: "result.id", Desc: "ファイルまたはフォルダへのリンクのID"},
- 		&{Name: "result.tag", Desc: "エントリーの種別 (file, または folder)"},
- 		&{Name: "result.url", Desc: "共有リンクのURL."},
- 		&{Name: "result.name", Desc: "リンク先ファイル名称"},
  		&{Name: "result.expires", Desc: "有効期限 (設定されている場合)"},
- 		&{Name: "result.path_lower", Desc: "パス (すべて小文字に変換)."},
- 		&{Name: "result.visibility", Desc: "共有リンクの開示範囲"},
  	},
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "チームフォルダ名のデータファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.name", Desc: "チームフォルダ名"},
- 		&{Name: "result.team_folder_id", Desc: "チームフォルダのID"},
  		&{Name: "result.name", Desc: "チームフォルダの名称"},
  		&{Name: "result.status", Desc: "チームフォルダの状態 (active, archived, または arch"...},
  		... // 2 identical elements
  	},
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "チームフォルダ名のデータファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "このレポートは処理結果を出力します.",
  	Columns: {&{Name: "status", Desc: "処理の状態"}, &{Name: "reason", Desc: "失敗またはスキップの理由"}, &{Name: "input.name", Desc: "チームフォルダ名"}},
  }
```
# コマンド仕様の変更: `teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DstPeerName",
+ 			Desc:     "宛先チームのアカウント別名",
+ 			Default:  "dst",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "チームフォルダ名のデータファイル",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "元チームのアカウント別名",
+ 			Default:  "src",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: verification

```
  &dc_recipe.Report{
  	Name:    "verification",
- 	Desc:    "",
+ 	Desc:    "このレポートはフォルダ間の差分を出力します.",
  	Columns: {&{Name: "diff_type", Desc: "差分のタイプ`file_content_diff`: コンテンツハッシ\xe3"...}, &{Name: "left_path", Desc: "左のパス"}, &{Name: "left_kind", Desc: "フォルダまたはファイル"}, &{Name: "left_size", Desc: "左ファイルのサイズ"}, ...},
  }
```
# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: namespace_file

```
  &dc_recipe.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "このレポートはチームの名前空間を一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
- 		&{Name: "namespace_id", Desc: "名前空間ID"},
  		&{Name: "namespace_name", Desc: "名前空間の名称"},
  		&{Name: "namespace_member_email", Desc: "これがチームメンバーフォルダまたはアプリフ\xe3"...},
- 		&{Name: "file_id", Desc: "ファイルへの一意なID"},
  		&{Name: "tag", Desc: "エントリーの種別`file`, `folder`, または `deleted`"},
  		&{Name: "name", Desc: "名称"},
  		&{Name: "path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		&{Name: "client_modified", Desc: "ファイルの場合、更新日時はクライアントPC上\xe3\x81"...},
  		&{Name: "server_modified", Desc: "Dropbox上で最後に更新された日時"},
- 		&{
- 			Name: "revision",
- 			Desc: "ファイルの現在バージョンの一意な識別子",
- 		},
  		&{Name: "size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "content_hash", Desc: "ファイルコンテンツのハッシュ"},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "これが共有フォルダのマウントポイントである場合、ここにマウントされている共有フォルダ\xe3\x81"...,
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "設定されている場合、共有フォルダに内包されています.",
- 		},
  	},
  }
```
# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Depth",
+ 			Desc:     "深さ",
+ 			Default:  "1",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
  		&{Name: "input.name", Desc: "名前空間の名称"},
- 		&{Name: "input.namespace_id", Desc: "名前空間ID"},
  		&{Name: "input.namespace_type", Desc: "名前異空間のタイプ (app_folder, shared_folder, team_fol"...},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
- 		},
- 		&{Name: "result.namespace_name", Desc: "名前空間の名称"},
- 		&{Name: "result.namespace_id", Desc: "名前空間ID"},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "名前異空間のタイプ (app_folder, shared_folder, team_folder, または team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "メンバーフォルダまたはアプリフォルダである場合、その所有者チームメンバーのID",
- 		},
  		&{Name: "result.path", Desc: "フォルダへのパス"},
  		&{Name: "result.count_file", Desc: "このフォルダに含まれるファイル数"},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: team_folder

```
  &dc_recipe.Report{
  	Name: "team_folder",
- 	Desc: "",
+ 	Desc: "このレポートはチーム内のチームフォルダを一覧します.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_folder_id", Desc: "チームフォルダのID"},
  		&{Name: "name", Desc: "チームフォルダの名称"},
  		&{Name: "status", Desc: "チームフォルダの状態 (active, archived, または arch"...},
  		... // 2 identical elements
  	},
  }
```
# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DstPeerName",
+ 			Desc:     "宛先チームのアカウント別名",
+ 			Default:  "dst",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "元チームのアカウント別名",
+ 			Default:  "src",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: verification

```
  &dc_recipe.Report{
  	Name:    "verification",
- 	Desc:    "",
+ 	Desc:    "このレポートはフォルダ間の差分を出力します.",
  	Columns: {&{Name: "diff_type", Desc: "差分のタイプ`file_content_diff`: コンテンツハッシ\xe3"...}, &{Name: "left_path", Desc: "左のパス"}, &{Name: "left_kind", Desc: "フォルダまたはファイル"}, &{Name: "left_size", Desc: "左ファイルのサイズ"}, ...},
  }
```
# コマンド仕様の変更: `web`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Port",
+ 			Desc:     "ポート番号",
+ 			Default:  "7800",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1024), "value": float64(7800)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
