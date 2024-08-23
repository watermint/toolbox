---
layout: release
title: リリースの変更点 135
lang: ja
---

# `リリース 135` から `リリース 136` までの変更点

# 追加されたコマンド


| コマンド                 | タイトル                                                                     |
|--------------------------|------------------------------------------------------------------------------|
| config license install   | ライセンスキーのインストール                                                 |
| dropbox file restore ext | 特定の拡張子を持つファイルの復元                                             |
| util feed json           | URLからフィードを読み込み、コンテンツをJSONとして出力する。                  |
| util json query          | JSONデータを問い合わせる                                                     |
| util uuid timestamp      | UUIDタイムスタンプの解析                                                     |
| util uuid ulid           | ULID（Universally Unique Lexicographically Sortable Identifier）を生成する。 |
| util uuid v7             | UUID v7 の生成                                                               |
| util uuid version        | UUIDのバージョンとバリアントの解析                                           |



# 削除されたコマンド


| コマンド                         | タイトル                           |
|----------------------------------|------------------------------------|
| util desktop display list        | このマシンのディスプレイを一覧表示 |
| util desktop screenshot interval | 定期的にスクリーンショットを撮る   |
| util desktop screenshot snap     | スクリーンショットを撮る           |



# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dev license issue`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "LifecycleWarningAfter", Desc: "ビルド時刻からこの期間経過後のライフサイク\xe3"..., Default: "31536000", TypeName: "int", ...},
  		&{Name: "Owner", Desc: "ライセンス・リポジトリの所有者", Default: "watermint", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
+ 		&{
+ 			Name:     "RecipeAllowedPrefix",
+ 			Desc:     "レシピの接頭辞",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "RecipesAllowed", Desc: "コンマで区切られたレシピのリスト", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Repository", Desc: "ライセンス・リポジトリ", Default: "toolbox-supplement", TypeName: "string", ...},
  		&{Name: "Scope", Desc: "ライセンス範囲", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release announcement`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "CategoryId", Desc: "お知らせカテゴリーID", Default: "DIC_kwDOBFqRm84CQesd", TypeName: "string", ...},
  		&{Name: "Owner", Desc: "レポジトリの所有者", Default: "watermint", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", Default: "toolbox", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release asset`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "コンテンツパス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repo", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Text", Desc: "テキストコンテンツ", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release asseturl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "SourceOwner", Desc: "ソースリポジトリの所有者", TypeName: "string"},
  		&{Name: "SourceRepo", Desc: "ソースリポジトリ名", TypeName: "string"},
  		... // 3 identical elements
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
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `dev release checkin`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
- 	IsSecret:        false,
+ 	IsSecret:        true,
  	IsConsole:       false,
  	IsExperimental:  false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Branch", Desc: "リポジトリブランチ", Default: "main", TypeName: "string", ...},
  		&{Name: "Owner", Desc: "レポジトリの所有者", Default: "watermint", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repo", Desc: "レポジトリ名", Default: "toolbox", TypeName: "string", ...},
  		&{Name: "SupplementBranch", Desc: "リポジトリブランチ名の補足", Default: "main", TypeName: "string", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
- 			TypeAttr: string("github_public"),
+ 			TypeAttr: string("github"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"ConnGithub": "github_repo",
+ 		"ConnGithub": "github",
- 		"Peer":       "github_repo",
+ 		"Peer":       "github",
  	},
  	Services: {"github"},
  	IsSecret: true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ArtifactPath", Desc: "成果物へのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Branch", Desc: "対象ブランチ", Default: "main", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox file account feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file account filesystem`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file account info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Left": "dropbox_individual", "Right": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file export url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock batch acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock batch release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Dst": "dropbox_individual", "Src": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file restore all`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file revision download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file revision list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file revision restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file share info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder leave`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedlink delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedlink info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sync down`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sync online`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file tag add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file tag delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file tag list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file template apply`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file template capture`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox file watch`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox paper append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox paper create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox paper overwrite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox paper prepend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin group role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin group role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team backup device status`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper count`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content member size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filerequest clone`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filesystem`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group update type`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team insight scan`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team insight scanretry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold revision list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold update desc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold update name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member folder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch invisible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch visible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace summary`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file batch copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file sync batch up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch leave`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder isolate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink cap expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink cap visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink delete links`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink delete member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update password`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder partial replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder sync setting list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder sync setting update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `github content get`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "コンテンツへのパス.", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Ref", Desc: "リファレンス名", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github content put`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "コンテンツへのパス.", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Filter", Desc: "どのような種類の課題を返すかを示します.", Default: "assigned", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Labels", Desc: "カンマで区切られたラベル名のリスト.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Since", Desc: "指定した時間以降に更新された通知のみを表示\xe3"..., TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "State", Desc: "返すべき課題の状態を示す.", Default: "open", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github release asset download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "ダウンロード パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github release asset list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github release asset upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Asset", Desc: "成果物のパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Name", Desc: "リリース名称", TypeName: "string"},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Tag", Desc: "タグ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github release list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `github tag create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Sha1", Desc: "コミットのSHA1ハッシュ", TypeName: "string"},
  		&{Name: "Tag", Desc: "タグ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util monitor client`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# コマンド仕様の変更: `util release install`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AcceptLicenseAgreement", Desc: "対象リリースの使用許諾契約に同意する", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "インストールするパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
- 			TypeAttr: string("github_public"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", Default: "latest", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util tidy pack remote`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
