---
layout: release
title: リリースの変更点: 66
lang: ja
---

# `リリース 66` から `リリース 67` までの変更点

# 追加されたコマンド


| コマンド                | タイトル                            |
|-------------------------|-------------------------------------|
| dev util anonymise      | キャプチャログを匿名化します.       |
| job log jobid           | 指定したジョブIDのログを取得する    |
| job log kind            | 指定種別のログを結合して出力します  |
| job log last            | 最後のジョブのログファイルを出力.   |
| member clear externalid | メンバーのexternal_idを初期化します |



# コマンド仕様の変更: `config disable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `config enable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `config features`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev async`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev catalogue`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev ci artifact connect`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Full": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth connect`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"Audit":  "business_audit",
  		"File":   "business_file",
  		"Full":   "user_full",
+ 		"Github": "github_repo",
  		"Info":   "business_info",
  		"Mgmt":   "business_management",
  	},
- 	Services:  nil,
+ 	Services:  []string{"dropbox", "dropbox_business", "github"},
  	IsSecret:  true,
  	IsConsole: false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Audit": "business_audit", "File": "business_file", "Full": "user_full", "Info": "business_info", ...},
- 	Services:        nil,
+ 	Services:        []string{"dropbox", "dropbox_business"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Filename", Desc: "ファイル名", Default: "README.md", TypeName: "string", ...},
  		&{Name: "Lang", Desc: "言語", TypeName: "domain.common.model.mo_string.opt_string"},
- 		&{
- 			Name:     "MarkdownReadme",
- 			Desc:     "READMEをMarkdownフォーマットで生成",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev dummy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev echo`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	... // 5 identical fields
  }
```
# コマンド仕様の変更: `dev kvs dump`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev preflight`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 5 identical fields
  	CliArgs:         "",
  	CliNote:         "",
- 	ConnUsePersonal: false,
+ 	ConnUsePersonal: true,
- 	ConnUseBusiness: false,
+ 	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
+ 		"Audit":  "business_audit",
+ 		"File":   "business_file",
+ 		"Full":   "user_full",
+ 		"Github": "github_repo",
+ 		"Info":   "business_info",
+ 		"Mgmt":   "business_management",
  	},
- 	Services:  nil,
+ 	Services:  []string{"dropbox", "dropbox_business", "github"},
  	IsSecret:  true,
  	IsConsole: true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"ConnGithub": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev spec diff`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev spec doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev test recipe`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `dev test resources`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	... // 5 identical fields
  }
```
# コマンド仕様の変更: `dev util curl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	... // 5 identical fields
  }
```
# コマンド仕様の変更: `dev util wait`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Left": "user_full", "Right": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file dispatch local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Dst": "user_full", "Src": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file sync preflight up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file watch`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `job history list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
- 	Values:          []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "ワークスペースへのパス.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `job loop`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `job run`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	... // 5 identical fields
  }
```
# コマンド仕様の変更: `member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "削除済メンバーを含めます.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "business_file", "Src": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github release asset download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github release asset list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github release asset upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github release list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedlink delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team content member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "フォルダ名によるフィルター. 名前による完全一致でフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "フォルダ名によるフィルター. 名前の前方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "フォルダ名によるフィルター. 名前の後方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "MemberTypeExternal",
+ 			Desc: "フォルダメンバーによるフィルター. 外部メンバーのみを残します (同じチームにいないメンバ\xe3"...,
+ 		},
+ 		&{
+ 			Name: "MemberTypeInternal",
+ 			Desc: "フォルダメンバーによるフィルター. 内部メンバーのみを残します (同じチームのメンバー)注意"...,
+ 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## 変更されたレポート: membership

```
  &dc_recipe.Report{
  	Name: "membership",
  	Desc: "このレポートは共有フォルダまたはチームフォ\xe3"...,
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "member_name", Desc: "このメンバーの名前"},
  		&{Name: "member_email", Desc: "このメンバーのメールアドレス"},
+ 		&{
+ 			Name: "same_team",
+ 			Desc: "メンバーが同じチームかどうか. もしメンバーが同じチームかどうか判定できない場合は空白を"...,
+ 		},
  	},
  }
```
# コマンド仕様の変更: `team content policy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "フォルダ名によるフィルター. 名前による完全一致でフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "フォルダ名によるフィルター. 名前の前方一致によるフィルター.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "フォルダ名によるフィルター. 名前の後方一致によるフィルター.",
+ 		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
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
+ 		"Peer": "business_file",
  	},
- 	Services:  nil,
+ 	Services:  []string{"dropbox_business"},
  	IsSecret:  false,
  	IsConsole: false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team filerequest clone`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 6 identical fields
  	CliNote:         "",
  	ConnUsePersonal: false,
- 	ConnUseBusiness: false,
+ 	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
+ 		"DstFile": "business_file",
+ 		"DstMgmt": "business_management",
+ 		"SrcFile": "business_file",
+ 		"SrcMgmt": "business_management",
  	},
- 	Services:  nil,
+ 	Services:  []string{"dropbox_business"},
  	IsSecret:  false,
  	IsConsole: false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 6 identical fields
  	CliNote:         "",
  	ConnUsePersonal: false,
- 	ConnUseBusiness: false,
+ 	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
+ 		"DstFile": "business_file",
+ 		"DstMgmt": "business_management",
+ 		"SrcFile": "business_file",
+ 		"SrcMgmt": "business_management",
  	},
- 	Services:  nil,
+ 	Services:  []string{"dropbox_business"},
  	IsSecret:  false,
  	IsConsole: false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `version`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	... // 5 identical fields
  }
```
