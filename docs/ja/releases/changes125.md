---
layout: release
title: リリースの変更点 124
lang: ja
---

# `リリース 124` から `リリース 125` までの変更点

# 追加されたコマンド


| コマンド             | タイトル                    |
|----------------------|-----------------------------|
| dev release asseturl | リリースのアセットURLを更新 |



# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BlockBlockSize", Desc: "一括アップロード時のブロックサイズ", Default: "24", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name:    "Method",
  			Desc:    "アップロード方法",
  			Default: "block",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("block"), string("sequential")}},
  		},
  		&{Name: "NumFiles", Desc: "ファイル数.", Default: "1000", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev build package`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BuildPath", Desc: "バイナリへのフルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "DeployPath", Desc: "デプロイ先フォルダパス(リモート)", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "DistPath", Desc: "パッケージの保存先フォルダのパス(ローカル)", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
+ 		&{
+ 			Name:     "ExecutableName",
+ 			Desc:     "実行ファイル名ベース",
+ 			Default:  "tbx",
+ 			TypeName: "string",
+ 		},
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
  	... // 17 identical fields
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
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repo", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Text", Desc: "テキストコンテンツ", TypeName: "string"},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
+ 			TypeAttr: string("github_public"),
+ 		},
+ 	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ArtifactPath", Desc: "成果物へのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Branch", Desc: "対象ブランチ", Default: "main", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "Query", Desc: "クエリ", TypeName: "string"},
  		&{Name: "Seed", Desc: "シェアードリンクのシード値", Default: "0", TypeName: "int", ...},
  		&{
  			Name:    "Visibility",
  			Desc:    "ビジビリティ",
  			Default: "random",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("random"), string("public"), string("team_only"), string("with_expire"), ...}},
  		},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "入力フォーマット (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "入力フォーマット (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "入力フォーマット (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paperのコンテンツ", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "入力フォーマット (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "ユーザーのDropbox内のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Category",
  			Desc:    "指定されたファイルカテゴリに検索を限定しま\xe3"...,
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "MaxResults", Desc: "返却するエントリーの最大数", Default: "25", TypeName: "essentials.model.mo_int.range_int", ...},
  		... // 3 identical elements
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Category",
  			Desc:    "指定されたファイルカテゴリに検索を限定しま\xe3"...,
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"},
  		... // 2 identical elements
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ManagementType",
  			Desc:    "グループ管理タイプ. `company_managed` または `user_m"...,
  			Default: "company_managed",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("company_managed"), string("user_managed")}},
  		},
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "ManagementType",
  			Desc:    "だれがこのグループを管理できるか (user_managed, "...,
  			Default: "company_managed",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("company_managed"), string("user_managed")}},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "IncludeExternalGroups", Desc: "レポートに外部のグループを含める.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group update type`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "Type",
  			Desc:    "グループタイプ（user_managed/company_managed)",
  			Default: "company_managed",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("user_managed"), string("company_managed")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job log jobid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "ジョブID", TypeName: "string"},
  		&{
  			Name:    "Kind",
  			Desc:    "ログの種別",
  			Default: "toolbox",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job log kind`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Kind",
  			Desc:    "ログの種別.",
  			Default: "toolbox",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job log last`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Kind",
  			Desc:    "ログの種別",
  			Default: "toolbox",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレスでフィルタリングし\xe3"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services deepl translate text`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.deepl.api.deepl_conn_impl.conn_deepl_api_impl",
- 			TypeAttr: string(""),
+ 			TypeAttr: string("deepl"),
  		},
  		&{Name: "SourceLang", Desc: "ソース言語コード（省略時は自動検出）", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "TargetLang", Desc: "対象言語コード", TypeName: "string"},
  		&{Name: "Text", Desc: "翻訳するテキスト", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services dropboxsign account info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AccountId", Desc: "アカウントID", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropboxsign.api.hs_conn_impl.conn_hello_sign_api",
- 			TypeAttr: nil,
+ 			TypeAttr: string("dropbox_sign"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma account info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma file export all page`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "書き出し形式（png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Path", Desc: "出力フォルダーパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "エクスポートスケールを1～400の範囲でパーセン"..., Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "TeamId", Desc: "チームID. チームIDを取得するには、自分が所属\xe3"..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma file export frame`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "書き出し形式（png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Key", Desc: "ファイルキー", TypeName: "string"},
  		&{Name: "Path", Desc: "出力フォルダーパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "エクスポートスケールを1～400の範囲でパーセン"..., Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma file export node`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "書き出し形式（png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Id", Desc: "ノードID", TypeName: "string"},
  		&{Name: "Key", Desc: "ファイルキー", TypeName: "string"},
  		&{Name: "Path", Desc: "出力フォルダーパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "エクスポートスケールを1～400の範囲でパーセン"..., Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma file export page`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "書き出し形式（png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Key", Desc: "ファイルキー", TypeName: "string"},
  		&{Name: "Path", Desc: "出力フォルダーパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "エクスポートスケールを1～400の範囲でパーセン"..., Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma file info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllNodes", Desc: "すべてのノード情報を含む", Default: "false", TypeName: "bool", ...},
  		&{Name: "Key", Desc: "ファイルキー", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "ProjectId", Desc: "プロジェクトID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services figma project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "TeamId", Desc: "チームID. チームIDを取得するには、自分が所属\xe3"..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github content get`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "コンテンツへのパス.", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Ref", Desc: "リファレンス名", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github content put`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Filter",
  			Desc:    "どのような種類の課題を返すかを示します.",
  			Default: "assigned",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("assigned"), string("created"), string("mentioned"), string("subscribed"), ...}},
  		},
  		&{Name: "Labels", Desc: "カンマで区切られたラベル名のリスト.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Since", Desc: "指定した時間以降に更新された通知のみを表示\xe3"..., TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name:    "State",
  			Desc:    "返すべき課題の状態を示す.",
  			Default: "open",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("open"), string("closed"), string("all")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release asset download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{Name: "Path", Desc: "ダウンロード パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release asset list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release asset upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Asset", Desc: "成果物のパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Release", Desc: "リリースタグ名", TypeName: "string"},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  		&{Name: "Tag", Desc: "タグ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github release list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
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
# コマンド仕様の変更: `services google mail label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ColorBackground",
  			Desc:    "背景色.",
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:    "ColorText",
  			Desc:    "テキストの色.",
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:    "LabelListVisibility",
  			Desc:    "Gmail ウェブインタフェースのラベルリストのラ\xe3"...,
  			Default: "labelShow",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("labelHide"), string("labelShow"), string("labelShowIfUnread")}},
  		},
  		&{
  			Name:    "MessageListVisibility",
  			Desc:    "Gmail ウェブインターフェースのメッセージリス\xe3"...,
  			Default: "show",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("hide"), string("show")}},
  		},
  		&{Name: "Name", Desc: "ラベル名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail message list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "メッセージを返すフォーマット. ",
  			Default: "metadata",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "SPAMやTRASHからのメッセージを結果に含める.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Labels", Desc: "指定されたラベルにすべて一致するラベルを持\xe3"..., TypeName: "essentials.model.mo_string.opt_string"},
  		... // 4 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail message processed list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "メッセージを返すフォーマット. ",
  			Default: "metadata",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "SPAMやTRASHからのメッセージを結果に含める.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Labels", Desc: "指定されたラベルにすべて一致するラベルを持\xe3"..., TypeName: "essentials.model.mo_string.opt_string"},
  		... // 4 identical elements
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
  		&{Name: "Data", Desc: "エクスポート先のパス.", TypeName: "Data"},
  		&{Name: "DataFormat", Desc: "出力フォーマット"},
  		&{
  			Name:    "DateTimeRender",
  			Desc:    "日付、時間、および期間を出力でどのように表\xe7"...,
  			Default: "serial",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("serial"), string("formatted")}},
  		},
  		&{Name: "Id", Desc: "スプレッドシートID", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_sheets", ...},
  		&{Name: "Range", Desc: "値がカバーする範囲をA1表記で表します. これは"..., TypeName: "string"},
  		&{
  			Name:    "ValueRender",
  			Desc:    "値を出力でどのように表現すべきか.",
  			Default: "formatted",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("formatted"), string("formatted"), string("formula")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "書き出したシートデータ"}},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services slack conversation history`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "After", Desc: "これ以降のメッセージを取得する.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Channel", Desc: "チャンネルID（C1234567890のようなもの）.", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "Peer",
+ 			Default:  "default",
  			TypeName: "domain.slack.api.work_conn_impl.conn_slack_api",
  			TypeAttr: []any{string("channels:history")},
  		},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
- 			Default:  "Peer",
+ 			Default:  "default",
  			TypeName: "domain.slack.api.work_conn_impl.conn_slack_api",
  			TypeAttr: []any{string("channels:read")},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "AccessLevel",
  			Desc:    "アクセス権限 (viewer/editor)",
  			Default: "editor",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("editor"), string("viewer"), string("viewer_no_comment")}},
  		},
  		&{Name: "Email", Desc: "フォルダメンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Message", Desc: "カスタム招待メッセージ", TypeName: "essentials.model.mo_string.opt_string"},
  		... // 3 identical elements
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "AclUpdatePolicy",
  			Desc:    "共有フォルダーのアクセスコントロールリスト\xef"...,
  			Default: "owner",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("owner"), string("editor")}},
  		},
  		&{
  			Name:    "MemberPolicy",
  			Desc:    "この共有フォルダーのメンバーになれる人.",
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("team"), string("anyone")}},
  		},
  		&{Name: "Path", Desc: "共有するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:    "SharedLinkPolicy",
  			Desc:    "このフォルダー内の共有リンクを閲覧できる人.",
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("anyone"), string("members")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "FilterBy",
  			Desc:    "Paperドキュメントのフィルタリング方法（doc_crea"...,
  			Default: "docs_created",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("docs_created"), string("docs_accessed")}},
  		},
  		&{
  			Name:    "Format",
  			Desc:    "エクスポートファイル形式 (html/markdown)",
  			Default: "html",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("html"), string("markdown")}},
  		},
  		&{Name: "Path", Desc: "エクスポートフォルダのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "FilterBy",
  			Desc:    "Paperドキュメントのフィルタリング方法（doc_crea"...,
  			Default: "docs_created",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("docs_created"), string("docs_accessed")}},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `team content member size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSubFolders", Desc: "レポートにサブフォルダーを含める.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "AclUpdatePolicy",
  			Desc:    "この共有フォルダのメンバーを追加・削除でき\xe3"...,
  			Default: "owner",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("owner"), string("editor")}},
  		},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "MemberPolicy",
  			Desc:    "この共有フォルダーのメンバーになれる人.",
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("team"), string("anyone")}},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "SharedLinkPolicy",
  			Desc:    "この共有フォルダー内のコンテンツに作成され\xe3"...,
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("anyone"), string("members")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "NewVisibility",
  			Desc:    "新しい視認性設定",
  			Default: "team_only",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("team_only")}},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "Visibility",
  			Desc:    "可視性によるリンクのフィルタリング (all/public/"...,
  			Default: "all",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("all"), string("public"), string("team_only"), string("password"), ...}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "NewVisibility",
  			Desc:    "新しい視認性設定",
  			Default: "team_only",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("public"), string("team_only")}},
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "SyncSetting",
  			Desc:    "チームフォルダの同期設定",
  			Default: "default",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("default"), string("not_synced")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamspace asadmin member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util file hash`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Algorithm",
  			Desc:    "ハッシュアルゴリズム(md5/sha1/sha256)",
  			Default: "sha1",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("md5"), string("sha1"), string("sha256")}},
  		},
  		&{Name: "File", Desc: "ダイジェストを作成するファイルへのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util image placeholder`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "Path", Desc: "生成された画像をエクスポートするパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Text", Desc: "必要に応じてテキスト", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:    "TextAlign",
  			Desc:    "テキストの配置",
  			Default: "left",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("left"), string("center"), string("right")}},
  		},
  		&{Name: "TextColor", Desc: "文字色", Default: "black", TypeName: "string", ...},
  		&{Name: "TextPosition", Desc: "テキストの位置", Default: "center", TypeName: "string", ...},
  		&{Name: "Width", Desc: "幅(ピクセル)", Default: "640", TypeName: "int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util qrcode create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ErrorCorrectionLevel",
  			Desc:    "誤差補正レベル（l/m/q/h）.",
  			Default: "m",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("l"), string("m"), string("q"), string("h")}},
  		},
  		&{
  			Name:    "Mode",
  			Desc:    "QRコードのエンコードモード",
  			Default: "auto",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("auto"), string("numeric"), string("alpha_numeric"), string("unicode")}},
  		},
  		&{Name: "Out", Desc: "ファイル名付きの出力パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Size", Desc: "画像解像度(ピクセル)", Default: "256", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Text", Desc: "テキストデータ", TypeName: "Text"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util qrcode wifi`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ErrorCorrectionLevel",
  			Desc:    "誤差補正レベル（l/m/q/h）.",
  			Default: "m",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("l"), string("m"), string("q"), string("h")}},
  		},
  		&{
  			Name:    "Hidden",
  			Desc:    "SSIDが隠されている場合は、`true`となります. SSID"...,
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("true"), string("false")}},
  		},
  		&{
  			Name:    "Mode",
  			Desc:    "QRコードのエンコードモード",
  			Default: "auto",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("auto"), string("numeric"), string("alpha_numeric"), string("unicode")}},
  		},
  		&{
  			Name:    "NetworkType",
  			Desc:    "ネットワークの種類.",
  			Default: "WPA",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("WPA"), string("WEP"), string("")}},
  		},
  		&{Name: "Out", Desc: "ファイル名付きの出力パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Size", Desc: "画像解像度(ピクセル)", Default: "256", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Ssid", Desc: "ネットワークSSID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util release install`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AcceptLicenseAgreement", Desc: "対象リリースの使用許諾契約に同意する", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "インストールするパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
+ 			TypeAttr: string("github_public"),
+ 		},
  		&{Name: "Release", Desc: "リリースタグ名", Default: "latest", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util text nlp japanese token`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Dictionary",
  			Desc:    "トークンの辞書名",
  			Default: "ipa",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("ipa"), string("uni")}},
  		},
  		&{Name: "IgnoreLineBreak", Desc: "改行を無視する", Default: "false", TypeName: "bool", ...},
  		&{Name: "In", Desc: "入力ファイルのパス", TypeName: "In"},
  		&{
  			Name:    "Mode",
  			Desc:    "トークン化モード(normal/search/extended)",
  			Default: "normal",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("normal"), string("search"), string("extend")}},
  		},
  		&{Name: "OmitBosEos", Desc: "BOS/EOSトークンの省略", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util text nlp japanese wakati`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Dictionary",
  			Desc:    "辞書名（ipa/uni）",
  			Default: "ipa",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("ipa"), string("uni")}},
  		},
  		&{Name: "IgnoreLineBreak", Desc: "改行を無視する", Default: "false", TypeName: "bool", ...},
  		&{Name: "In", Desc: "入力ファイルのパス", TypeName: "In"},
  		&{Name: "Separator", Desc: "テキストセパレータ", Default: " ", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util unixtime format`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "時間フォーマット",
  			Default: "iso8601",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("iso8601"), string("rfc1123"), string("rfc1123z"), string("rfc3339"), ...}},
  		},
  		&{
  			Name:    "Precision",
  			Desc:    "時間精度 (second/ms/ns)",
  			Default: "second",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("second"), string("ms"), string("ns")}},
  		},
  		&{Name: "Time", Desc: "Unix 時間", Default: "0", TypeName: "int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util unixtime now`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Precision",
  			Desc:    "時間精度 (second/ms/ns)",
  			Default: "second",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("second"), string("ms"), string("ns")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
