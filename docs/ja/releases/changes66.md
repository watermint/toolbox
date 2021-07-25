---
layout: release
title: リリースの変更点 65
lang: ja
---

# `リリース 65` から `リリース 66` までの変更点

# 追加されたコマンド


| コマンド                               | タイトル                                       |
|----------------------------------------|------------------------------------------------|
| dev catalogue                          | カタログを生成します                           |
| dev kvs dump                           | KVSデータのダンプ                              |
| services github release asset download | アセットをダウンロードします                   |
| services github release asset upload   | GitHub リリースへ成果物をアップロードします    |
| team filerequest clone                 | ファイルリクエストを入力データに従い複製します |



# 削除されたコマンド


| コマンド                         | タイトル                                            |
|----------------------------------|-----------------------------------------------------|
| dev desktop install              | Dropboxクライアントをインストールします             |
| dev desktop start                | Dropbox Desktopアプリケーションを起動します         |
| dev desktop stop                 | Dropboxデスクトップアプリケーションの停止を試みます |
| dev desktop suspendupdate        | Dropbox Updaterを停止・再開します                   |
| dev diag procmon                 | Process monitorのログを収集します                   |
| services github release asset up | GitHub リリースへ成果物をアップロードします         |
| web                              | Webコンソールの起動                                 |



# コマンド仕様の変更: `config disable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `config enable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `config features`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `dev preflight`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	IsTransient:    false,
  	Reports:        nil,
  	Feeds:          nil,
- 	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "テストリソースへのパス",
- 			Default:  "test/dev/resource.json",
- 			TypeName: "string",
- 		},
- 	},
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	IsTransient:    false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "ConnGithub", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "テストリソースへのパス",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
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
- 		&{
- 			Name:     "Recipe",
- 			Desc:     "テストするレシピ名",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
+ 		&{
+ 			Name:     "NoTimeout",
+ 			Desc:     "レシピのテスト時にタイムアウトしません",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{
- 			Name:     "Resource",
+ 			Name:     "Single",
- 			Desc:     "テスト用リソースへのパス",
+ 			Desc:     "テストするレシピ名",
  			Default:  "",
  			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Verbose", Desc: "テスト結果の詳細出力", Default: "false", TypeName: "bool", ...},
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
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `dev util wait`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "ファイルまたはフォルダは削除します.",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file delete",
  	CliArgs: "-path /PATH/TO/DELETE",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file dispatch local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "local",
  	Title:   "ローカルファイルを整理します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file dispatch local",
  	CliArgs: "-file /PATH/TO/DATA_FILE.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "URLからファイルを一括インポートします",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file import batch url",
  	CliArgs: "-file /path/to/data/file -path /path/to/import",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "URLからファイルをインポートします",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file import url",
  	CliArgs: "-url URL -path /path/to/import",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "merge",
  	Title:   "フォルダを統合します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file merge",
  	CliArgs: "-from /from/path -to /path/to",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "move",
  	Title:   "ファイルを移動します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file move",
  	CliArgs: "-src /SRC/PATH -dst /DST/PATH",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "ファイルコンテンツを他のアカウントに複製し\xe3"...,
  	Desc:    "このコマンドはファイルとフォルダを複製しま\xe3"...,
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file replication",
  	CliArgs: "-src source -src-path /path/src -dst dest -dst-path /path/dest",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "restore",
  	Title:   "指定されたパス以下をリストアします",
  	Desc:    "",
- 	Remarks: "(試験的実装です)",
+ 	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "file restore",
  	CliArgs: "-path /DROPBOX/PATH/TO/RESTORE",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "up",
  	Title:   "Dropboxと上り方向で同期します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file sync up",
  	CliArgs: "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF"...,
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "upload",
  	Title:   "ファイルのアップロード",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "file upload",
  	CliArgs: "-local-path /PATH/TO/UPLOAD -dropbox-path /DROPBOX/PATH",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "ファイルリクエストを作成します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "filerequest create",
  	CliArgs: "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "closed",
  	Title:   "このアカウントの全ての閉じられているファイ\xe3"...,
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "filerequest delete closed",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "ファイルリクエストのURLを指定して削除",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "filerequest delete url",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "グループを作成します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "group add",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "メンバーをグループに追加",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "group member add",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "rename",
  	Title:   "グループの改名",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "group rename",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `job loop`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "loop",
  	Title:   "指定日時までrunbookを実行します.",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(試験的実装です)",
  	Path:    "job loop",
  	CliArgs: `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 `...,
  	... // 4 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	IsTransient:    false,
  	... // 7 identical fields
  }
```
# コマンド仕様の変更: `job run`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "run",
  	Title:   "*.runbookoファイルにてワークフローを実行します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(試験的実装です)",
  	Path:    "job run",
  	CliArgs: "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	... // 4 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	IsTransient:    false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Fork",
- 			Desc:     "ワークフローを実行する際にプロセスをフォークする.",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "RunbookPath", Desc: "Runbookへのパス.", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
- 		&{
- 			Name:     "TimeoutSeconds",
- 			Desc:     "指定時間を経過したしたためプロセスを終了します.",
- 			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(0), "value": float64(0)},
- 		},
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
  	Name:    "delete",
  	Title:   "メンバーを削除します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member delete",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "detach",
  	Title:   "Dropbox BusinessユーザーをBasicユーザーに変更します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member detach",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "invite",
  	Title:   "メンバーを招待します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member invite",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "reinvite",
  	Title:   "招待済み状態メンバーをチームに再招待します",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member reinvite",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "email",
  	Title:   "メンバーのメールアドレス処理",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member update email",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
  	Title:   "チームメンバーのExternal IDを更新します.",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member update externalid",
  	CliArgs: "-file /path/to/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "profile",
  	Title:   "メンバーのプロフィール変更",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "member update profile",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `services github profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `services github release asset list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```

## 変更されたレポート: assets

```
  &dc_recipe.Report{
  	Name: "assets",
  	Desc: "GitHub リリースの成果物",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "state", Desc: "アセットの状態"},
  		&{Name: "download_count", Desc: "ダウンロード数"},
+ 		&{Name: "download_url", Desc: "ダウンロードURL"},
  	},
  }
```
# コマンド仕様の変更: `services github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "draft",
  	Title:   "リリースの下書きを作成",
  	Desc:    "",
- 	Remarks: "(試験的実装です)",
+ 	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "services github release draft",
  	CliArgs: "-body-file /LOCAL/PATH/TO/body.txt",
  	... // 4 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github release list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: true,
  	... // 8 identical fields
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "共有リンクの作成",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "sharedlink create",
  	CliArgs: "-path /path/to/share",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team diag explorer`



## 追加されたレポート


| 名称             | 説明                                                  |
|------------------|-------------------------------------------------------|
| namespace_member | このレポートは名前空間とそのメンバー一覧を出力します. |
| team_folder      | このレポートはチーム内のチームフォルダを一覧します.   |


# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "Name",
  			Desc:     "指定された名前に一致するフォルダのみを一覧\xe3"...,
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "Name",
  			Desc:     "指定された名前に一致するフォルダのみを一覧\xe3"...,
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "チームフォルダのアーカイブ",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "teamfolder archive",
  	CliArgs: "-name チームフォルダ名",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "複数のチームフォルダをアーカイブします",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(非可逆な操作です)",
  	Path:    "teamfolder batch archive",
  	CliArgs: "-file /path/to/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "チームフォルダを他のチームに複製します",
  	Desc:    "",
- 	Remarks: "(非可逆な操作です)",
+ 	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "teamfolder replication",
  	CliArgs: "",
  	... // 5 identical fields
  	IsSecret:       false,
  	IsConsole:      false,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: true,
  	IsTransient:    false,
  	... // 7 identical fields
  }
```
