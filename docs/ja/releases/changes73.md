---
layout: release
title: Changes of Release 72
lang: ja
---

# `リリース 72` から `リリース 73` までの変更点

# 追加されたコマンド


| コマンド                              | タイトル                             |
|---------------------------------------|--------------------------------------|
| dev benchmark upload                  | アップロードのベンチマーク           |
| dev build catalogue                   | カタログを生成します                 |
| dev build doc                         | ドキュメントを生成                   |
| dev build license                     | LICENSE.txtの生成                    |
| dev build preflight                   | リリースに向けて必要な事前準備を実施 |
| dev build readme                      | README.txtの生成                     |
| dev test async                        | 非同期処理フレームワークテスト       |
| dev test echo                         | テキストのエコー                     |
| file size                             | ストレージの利用量                   |
| file sync down                        | Dropboxと下り方向で同期します        |
| file sync online                      | オンラインファイルを同期します       |
| services asana team list              | チームのリスト                       |
| services asana team project list      | チームのプロジェクト一覧             |
| services asana team task list         | チームのタスク一覧                   |
| services asana workspace list         | ワークスペースの一覧                 |
| services asana workspace project list | ワークスペースのプロジェクト一覧     |



# 削除されたコマンド


| コマンド               | タイトル                                     |
|------------------------|----------------------------------------------|
| dev async              | 非同期処理フレームワークテスト               |
| dev catalogue          | カタログを生成します                         |
| dev doc                | ドキュメントを生成                           |
| dev echo               | テキストのエコー                             |
| dev preflight          | リリースに向けて必要な事前準備を実施         |
| file sync preflight up | 上り方向同期のための事前チェックを実施します |
| file upload            | ファイルのアップロード                       |



# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "アップロードするローカルファイルのパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "PeerName", Desc: "アカウントの別名", Default: "deploy", TypeName: "string", ...},
  		&{Name: "Timeout", Desc: "処理タイムアウト(秒)", Default: "30", TypeName: "int", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称    | 説明 |
|---------|------|
| deleted | パス |



## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{Name: "input.file", Desc: "ローカルファイルのパス"},
- 		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_display",
- 			Desc: "パス (表示目的で大文字小文字を区別する).",
- 		},
- 		&{
- 			Name: "result.client_modified",
- 			Desc: "ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ",
- 		},
- 		&{
- 			Name: "result.server_modified",
- 			Desc: "Dropbox上で最後に更新された日時",
- 		},
- 		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
+ 		&{Name: "input.entry_path", Desc: "パス"},
  	},
  }
```

## 変更されたレポート: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "このレポートはアップロード結果の概要を出力\xe3"...,
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "upload_start", Desc: "アップロード開始日時"},
+ 		&{Name: "start", Desc: "開始時間"},
- 		&{Name: "upload_end", Desc: "アップロード終了日時"},
+ 		&{Name: "end", Desc: "完了時間"},
  		&{Name: "num_bytes", Desc: "合計アップロードサイズ (バイト)"},
  		&{Name: "num_files_error", Desc: "失敗またはエラーが発生したファイル数."},
- 		&{
- 			Name: "num_files_upload",
- 			Desc: "アップロード済みまたはアップロード対象ファイル数",
- 		},
+ 		&{
+ 			Name: "num_files_transferred",
+ 			Desc: "アップロード/ダウンロードされたファイル数.",
+ 		},
  		&{Name: "num_files_skip", Desc: "スキップ対象またはスキップ予定のファイル数"},
+ 		&{Name: "num_folder_created", Desc: "作成されたフォルダ数."},
+ 		&{Name: "num_delete", Desc: "削除されたエントリ数."},
  		&{Name: "num_api_call", Desc: "この処理によって消費される見積アップロードA"...},
  	},
  }
```

## 変更されたレポート: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{Name: "input.file", Desc: "ローカルファイルのパス"},
- 		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
+ 		&{Name: "input.path", Desc: "パス"},
  		&{Name: "result.name", Desc: "名称"},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `dev diag endpoint`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "JobId",
  			Desc:     "検査するJob ID",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev diag throughput`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "EndpointNamePrefix", Desc: "エンドポイントによりフィルター. 名前の前方\xe4\xb8"...},
  		&{Name: "EndpointNameSuffix", Desc: "エンドポイントによりフィルター. 名前の後方\xe4\xb8"...},
  		&{
  			Name:     "JobId",
  			Desc:     "ジョブIDの指定",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "TimeFormat", Desc: "日時フォーマット (Goの日付フォーマット)", Default: "2006-01-02 15:04:05.999", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev kvs dump`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Path",
  			Desc:     "KVSデータへのパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ArtifactPath",
  			Desc:     "成果物へのパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Branch", Desc: "対象ブランチ", Default: "master", TypeName: "string", ...},
  		&{Name: "ConnGithub", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "FilePath",
  			Desc:     "出力先ファイルパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Lang",
  			Desc:     "言語",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Release1",
  			Desc:     "リリース名1",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Release2",
  			Desc:     "リリース名2",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "FilePath",
  			Desc:     "ファイルパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Lang",
  			Desc:     "言語",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Distribution",
  			Desc:     "ファイル・フォルダの分布数",
  			Default:  "10000",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(10000)},
  		},
  		&{Name: "Extension", Desc: "カンマ区切りの拡張子一覧", Default: "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"..., TypeName: "string", ...},
  		&{Name: "Path", Desc: "モンキーテストパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{
  			Name:     "Seconds",
  			Desc:     "モンキーテストの実施時間(秒)",
  			Default:  "10",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(1), "value": float64(10)},
  		},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{Name: "All", Desc: "全てのレシピをテストします", Default: "false", TypeName: "bool", ...},
  		&{Name: "NoTimeout", Desc: "レシピのテスト時にタイムアウトしません", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Single",
  			Desc:     "テストするレシピ名",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Verbose", Desc: "テスト結果の詳細出力", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev util anonymise`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "JobIdName", Desc: "ジョブID名にてフィルター. 名前による完全一致"...},
  		&{Name: "JobIdNamePrefix", Desc: "ジョブID名にてフィルター. 名前の前方一致によ"...},
  		&{Name: "JobIdNameSuffix", Desc: "ジョブID名にてフィルター. 名前の後方一致によ"...},
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BufferSize",
  			Desc:     "バッファのサイズ",
  			Default:  "65536",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.097152e+06), "min": float64(1024), "value": float64(65536)},
  		},
  		&{
  			Name:     "Record",
  			Desc:     "テスト用に直接テストレコードを指定",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev util image jpeg`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Count",
  			Desc:     "生成するファイル数",
  			Default:  "10",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(32767), "min": float64(1), "value": float64(10)},
  		},
  		&{
  			Name:     "Height",
  			Desc:     "高さ",
  			Default:  "1080",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1), "value": float64(1080)},
  		},
  		&{Name: "NamePrefix", Desc: "ファイル名のプリフィックス", Default: "test_image", TypeName: "string", ...},
  		&{
  			Name:     "Path",
  			Desc:     "ファイルを生成するパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{
  			Name:     "Quality",
  			Desc:     "JPEGの品質",
  			Default:  "75",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(100), "min": float64(1), "value": float64(75)},
  		},
  		&{Name: "Seed", Desc: "乱数のシード", Default: "1", TypeName: "int", ...},
  		&{
  			Name:     "Width",
  			Desc:     "幅",
  			Default:  "1920",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1), "value": float64(1920)},
  		},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Seconds",
  			Desc:     "指定秒数待機",
  			Default:  "1",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(604800), "min": float64(1), "value": float64(1)},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "ローカルパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "ダウンロードするファイルパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "保存先ローカルパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "エクスポートするDropbox上のドキュメントパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "保存先ローカルパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Path",
  			Desc:     "インポート先のパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "削除済みファイルを含める", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "IncludeMediaInfo",
- 			Desc:     "メディア情報を含める",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "指定されたファイルカテゴリに検索を限定しま\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{
  			Name:     "Extension",
  			Desc:     "指定されたファイル拡張子に検索を限定します.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "検索対象とするユーザーのDropbox上のパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "指定されたファイルカテゴリに検索を限定しま\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{
  			Name:     "Extension",
  			Desc:     "指定されたファイル拡張子に検索を限定します.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "検索対象とするユーザーのDropbox上のパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Query", Desc: "検索文字列.", TypeName: "string"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ChunkSizeKb",
  			Desc:     "アップロードチャンク容量(Kバイト)",
  			Default:  "65536",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
  		},
+ 		&{
+ 			Name:     "Delete",
+ 			Desc:     "ローカルで削除されたファイルがある場合はDropboxのファイルを削除します",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "DropboxPath", Desc: "転送先のDropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "ローカルファイルのパス",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name: "NameDisableIgnore",
+ 			Desc: "名前によるフィルター. システムファイルと除外ファイルを処理対象外とします.",
+ 		},
+ 		&{
+ 			Name: "NameName",
+ 			Desc: "名前によるフィルター. 名前による完全一致でフィルター.",
+ 		},
- 		&{
- 			Name:     "FailOnError",
- 			Desc:     "処理でエラーが発生した場合にエラーを返します. このコマンドはこのフラグが指定されない場"...,
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
+ 		&{
+ 			Name: "NameNamePrefix",
+ 			Desc: "名前によるフィルター. 名前の前方一致によるフィルター.",
+ 		},
- 		&{
- 			Name:     "LocalPath",
- 			Desc:     "ローカルファイルのパス",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
- 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
- 		},
+ 		&{
+ 			Name: "NameNameSuffix",
+ 			Desc: "名前によるフィルター. 名前の後方一致によるフィルター.",
+ 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "SkipExisting",
+ 			Desc:     "既存ファイルをスキップします. 上書きしません.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "WorkPath",
+ 			Desc:     "テンポラリパス",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称    | 説明 |
|---------|------|
| deleted | パス |



## 変更されたレポート: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{Name: "input.file", Desc: "ローカルファイルのパス"},
- 		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
- 		&{Name: "result.name", Desc: "名称"},
- 		&{
- 			Name: "result.path_display",
- 			Desc: "パス (表示目的で大文字小文字を区別する).",
- 		},
- 		&{
- 			Name: "result.client_modified",
- 			Desc: "ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ",
- 		},
- 		&{
- 			Name: "result.server_modified",
- 			Desc: "Dropbox上で最後に更新された日時",
- 		},
- 		&{Name: "result.size", Desc: "ファイルサイズ(バイト単位)"},
- 		&{Name: "result.content_hash", Desc: "ファイルコンテンツのハッシュ"},
+ 		&{Name: "input.entry_path", Desc: "パス"},
  	},
  }
```

## 変更されたレポート: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "このレポートはアップロード結果の概要を出力\xe3"...,
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "upload_start", Desc: "アップロード開始日時"},
+ 		&{Name: "start", Desc: "開始時間"},
- 		&{Name: "upload_end", Desc: "アップロード終了日時"},
+ 		&{Name: "end", Desc: "完了時間"},
  		&{Name: "num_bytes", Desc: "合計アップロードサイズ (バイト)"},
  		&{Name: "num_files_error", Desc: "失敗またはエラーが発生したファイル数."},
- 		&{
- 			Name: "num_files_upload",
- 			Desc: "アップロード済みまたはアップロード対象ファイル数",
- 		},
+ 		&{
+ 			Name: "num_files_transferred",
+ 			Desc: "アップロード/ダウンロードされたファイル数.",
+ 		},
  		&{Name: "num_files_skip", Desc: "スキップ対象またはスキップ予定のファイル数"},
+ 		&{Name: "num_folder_created", Desc: "作成されたフォルダ数."},
+ 		&{Name: "num_delete", Desc: "削除されたエントリ数."},
  		&{Name: "num_api_call", Desc: "この処理によって消費される見積アップロードA"...},
  	},
  }
```

## 変更されたレポート: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "処理の状態"},
  		&{Name: "reason", Desc: "失敗またはスキップの理由"},
- 		&{Name: "input.file", Desc: "ローカルファイルのパス"},
- 		&{Name: "input.size", Desc: "ローカルファイルのサイズ"},
+ 		&{Name: "input.path", Desc: "パス"},
  		&{Name: "result.name", Desc: "名称"},
  		&{Name: "result.path_display", Desc: "パス (表示目的で大文字小文字を区別する)."},
  		... // 4 identical elements
  	},
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "AllowLateUploads",
  			Desc:     "設定した場合、期限を過ぎてもアップロードを\xe8"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Deadline", Desc: "ファイルリクエストの締め切り.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Path", Desc: "ファイルをアップロードするDropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ManagementType",
  			Desc:     "グループ管理タイプ. `company_managed` または `user_m"...,
  			Default:  "company_managed",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("company_managed"), string("user_managed")}},
  		},
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Days",
  			Desc:     "目標日数",
  			Default:  "7",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(7)},
  		},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Days",
  			Desc:     "目標日数",
  			Default:  "28",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(28)},
  		},
  	},
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
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job log jobid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "ジョブID", TypeName: "string"},
  		&{
  			Name:     "Kind",
  			Desc:     "ログの種別",
  			Default:  "toolbox",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job log kind`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Kind",
  			Desc:     "ログの種別.",
  			Default:  "toolbox",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job log last`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Kind",
  			Desc:     "ログの種別",
  			Default:  "toolbox",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{
  			Name:     "Path",
  			Desc:     "ワークスペースへのパス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
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
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...},
  		&{
  			Name:     "TransferDestMember",
  			Desc:     "指定された場合は、指定ユーザーに削除するメ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "TransferNotifyAdminEmailOnError",
  			Desc:     "指定された場合は、転送時にエラーが発生した\xe6"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "WipeData", Desc: "指定した場合にはユーザーのデータがリンクさ\xe3"..., Default: "true", TypeName: "bool", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...},
  		&{
  			Name:     "Quota",
  			Desc:     "カスタムの容量制限 (1TB = 1024GB). 0の場合、容量\xe5"...,
  			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github content get`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "コンテンツへのパス.", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{
  			Name:     "Ref",
  			Desc:     "リファレンス名",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github content put`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Branch",
  			Desc:     "ブランチ名",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Content",
  			Desc:     "コンテンツファイルへのパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Message", Desc: "コミットメッセージ", TypeName: "string"},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		... // 3 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release asset download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			Name:     "Path",
  			Desc:     "ダウンロード パス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release asset upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Asset",
  			Desc:     "成果物のパス",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BodyFile",
  			Desc:     "本文テキストファイルへのパスファイルはBOMな\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Branch", Desc: "対象ブランチ名", TypeName: "string"},
  		&{Name: "Name", Desc: "リリース名称", TypeName: "string"},
  		... // 4 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail filter add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AddLabelIfNotExist", Desc: "ラベルが存在しない場合はラベルを作成します.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "AddLabels",
  			Desc:     "','で区切られたメッセージに追加するラベルの\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaExcludeChats", Desc: "チャットを除外するかどうか", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "CriteriaFrom",
  			Desc:     "送信者の表示名またはメールアドレス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaHasAttachment", Desc: "添付ファイルがあるメッセージ.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "CriteriaNegatedQuery",
  			Desc:     "指定されたクエリにマッチしないメッセージの\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaNoAttachment", Desc: "添付ファイルがないメッセージ.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "CriteriaQuery",
  			Desc:     "指定されたクエリにマッチするメッセージのみ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaSize", Desc: "すべてのヘッダと添付ファイルを含む、RFC822 メ"..., Default: "0", TypeName: "int", ...},
  		&{
  			Name:     "CriteriaSizeComparison",
  			Desc:     "メッセージのサイズをどのようにバイト数で表\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "CriteriaTo",
  			Desc:     "受信者の表示名またはメールアドレス. to\"、\"cc\"\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Forward",
  			Desc:     "メッセージの転送先となるメールアドレス.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{
  			Name:     "RemoveLabels",
  			Desc:     "','で区切られたメッセージから削除するラベル\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail filter batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AddLabelIfNotExist", Desc: "ラベルが存在しない場合はラベルを作成します.", Default: "false", TypeName: "bool", ...},
  		&{
- 			Name: "ApplyToInboxMessages",
+ 			Name: "ApplyToExistingMessages",
  			Desc: strings.Join({
- 				"INBOX内",
  				"\xe3",
- 				"\x81\xaeクエリを満たす",
+ 				"\x82\xafエリを満たす既存の",
  				"メッセージにラベルを適用します.",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ColorBackground",
  			Desc:     "背景色.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:     "ColorText",
  			Desc:     "テキストの色.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:     "LabelListVisibility",
  			Desc:     "Gmail ウェブインタフェースのラベルリストのラ\xe3"...,
  			Default:  "labelShow",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("labelHide"), string("labelShow"), string("labelShowIfUnread")}},
  		},
  		&{
  			Name:     "MessageListVisibility",
  			Desc:     "Gmail ウェブインターフェースのメッセージリス\xe3"...,
  			Default:  "show",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("hide"), string("show")}},
  		},
  		&{Name: "Name", Desc: "ラベル名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail message list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Format",
  			Desc:     "メッセージを返すフォーマット. ",
  			Default:  "metadata",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "SPAMやTRASHからのメッセージを結果に含める.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Labels",
  			Desc:     "指定されたラベルにすべて一致するラベルを持\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "MaxResults", Desc: "返すメッセージの最大数.", Default: "20", TypeName: "int", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{
  			Name:     "Query",
  			Desc:     "指定されたクエリにマッチするメッセージのみ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail message processed list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Format",
  			Desc:     "メッセージを返すフォーマット. ",
  			Default:  "metadata",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "SPAMやTRASHからのメッセージを結果に含める.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Labels",
  			Desc:     "指定されたラベルにすべて一致するラベルを持\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "MaxResults", Desc: "返すメッセージの最大数.", Default: "20", TypeName: "int", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{
  			Name:     "Query",
  			Desc:     "指定されたクエリにマッチするメッセージのみ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Expires", Desc: "共有リンクの有効期限日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Password",
  			Desc:     "パスワード",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "TeamOnly", Desc: "リンクがチームメンバーのみアクセスできます", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Password",
  			Desc:     "共有リンクのパスワード",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Url", Desc: "共有リンクのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "File", Desc: "メールアドレスリストのファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		... // 2 identical elements
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "イベントのカテゴリ",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "EndDate",
  			Desc:     "終了日",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "string"},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "一つのイベントカテゴリのみを返すようなフィ\xe3"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "スキャンのタイムアウト設定. スキャンタイムアウトした場合、チームフォルダのサブフォルダ"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "スキャンのタイムアウト設定. スキャンタイムアウトした場合、チームフォルダのサブフォルダ"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team diag explorer`



## 追加されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |


# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前による完全一致でフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の前方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の後方一致によるフィルター.",
+ 		},
  		&{Name: "IncludeDeleted", Desc: "指定された場合、削除済みのファイルやフォル\xe3"..., Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "IncludeMediaInfo",
- 			Desc:     "指定された場合、JSONレポートに写真や動画ファイルへのメデイア情報を含めます",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "IncludeMemberFolder", Desc: "指定された場合、チームメンバーのフォルダを\xe5"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
- 		&{
- 			Name:     "Name",
- 			Desc:     "指定された名前に一致するフォルダのみを一覧します",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |


# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "フォルダ階層数の指定",
  			Default:  "1",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
- 				"max":   float64(2.147483647e+09),
+ 				"max":   float64(300),
  				"min":   float64(1),
  				"value": float64(1),
  			},
  		},
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前による完全一致でフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の前方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の後方一致によるフィルター.",
+ 		},
  		&{Name: "IncludeAppFolder", Desc: "Trueの場合、アプリフォルダを含めます", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeMemberFolder", Desc: "Trueの場合、チームメンバーフォルダを含めます", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
- 		&{
- 			Name:     "Name",
- 			Desc:     "指定された名前に一致するフォルダのみを一覧します",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |


# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "出力するリンクを可視性にてフィルターします "...,
  			Default:  "public",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("public"), string("team_only"), string("password"), string("team_and_password"), ...}},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "新しい有効期限の日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Days",
  			Desc:     "新しい有効期限までの日時",
  			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "対象となるリンクの公開範囲",
  			Default:  "public",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("public"), string("team_only"), string("password"), string("team_and_password"), ...}},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "SyncSetting",
  			Desc:     "チームフォルダの同期設定",
  			Default:  "default",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("default"), string("not_synced")}},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前による完全一致でフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の前方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の後方一致によるフィルター.",
+ 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |


# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "深さ",
  			Default:  "1",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
  		},
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前による完全一致でフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の前方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "名前に一致するフォルダのみをリストアップします. 名前の後方一致によるフィルター.",
+ 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称   | 説明                                |
|--------|-------------------------------------|
| errors | このレポートは処理結果を出力します. |


# コマンド仕様の変更: `teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "スキャンのタイムアウト設定. スキャンタイムアウトした場合、チームフォルダのサブフォルダ"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "スキャンのタイムアウト設定. スキャンタイムアウトした場合、チームフォルダのサブフォルダ"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
