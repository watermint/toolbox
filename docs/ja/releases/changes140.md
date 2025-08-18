---
layout: release
title: リリースの変更点 139
lang: ja
---

# `リリース 139` から `リリース 140` までの変更点

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
  		&{Name: "ConnGithub", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
+ 		&{
+ 			Name:     "ExecutableName",
+ 			Desc:     "公開する実行可能ファイルの名前",
+ 			Default:  "tbx",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "HomebrewRepoBranch",
+ 			Desc:     "公開に使用する Homebrew tap リポジトリのブランチです。",
+ 			Default:  "master",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "HomebrewRepoName",
+ 			Desc:     "Homebrew tap リポジトリの名前",
+ 			Default:  "homebrew-toolbox",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "HomebrewRepoOwner",
+ 			Desc:     "Homebrew tap リポジトリのオーナー",
+ 			Default:  "watermint",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepoName",
+ 			Desc:     "リリースを公開するリポジトリの名前です。",
+ 			Default:  "toolbox",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepoOwner",
+ 			Desc:     "リリースを公開するリポジトリの所有者です。",
+ 			Default:  "watermint",
+ 			TypeName: "string",
+ 		},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("events.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("events.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("events.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("events.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin group role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin group role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ExceptionGroup", Desc: "例外グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "MemberRoles", Desc: "メンバーと管理者の役割のマッピング"}},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "ロールID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team backup device status`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndTime", Desc: "データを取得する期間の終了日時 (該当同時刻\xe3\x82"..., TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("events.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "StartTime", Desc: "データを取得する期間の開始日時 (該当同時刻\xe3\x82"..., TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper count`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Format", Desc: "エクスポートファイル形式 (html/markdown)", Default: "html", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Path", Desc: "エクスポートフォルダのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "FilterBy", Desc: "Paperドキュメントのフィルタリング方法（doc_crea"..., Default: "docs_created", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...},
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content member size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "IncludeSubFolders", Desc: "レポートにサブフォルダーを含める.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberNamePrefix", Desc: "メンバーをフィルタリングします. 名前の前方\xe4\xb8"...},
  		&{Name: "MemberNameSuffix", Desc: "メンバーをフィルタリングします. 名前の後方\xe4\xb8"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...},
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team device list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sessions.list"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("sessions.modify"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team feature`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("file_requests.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filesystem`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ManagementType", Desc: "グループ管理タイプ. `company_managed` または `user_m"..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "ManagementType", Desc: "だれがこのグループを管理できるか (user_managed, "..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "グループ名リストのデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "GroupNameSuffix", Desc: "グループ名でフィルタリングします. 名前の後\xe6\x96"...},
  		&{Name: "IncludeExternalGroups", Desc: "レポートに外部のグループを含める.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group update type`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "グループ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "Type", Desc: "グループタイプ（user_managed/company_managed)", Default: "company_managed", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team info`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team linkedapp list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sessions.list"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.delete"),
  				string("members.write"),
  				string("team_info.read"),
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
# コマンド仕様の変更: `dropbox team member batch detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.delete"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "RevokeTeamShares", Desc: "指定した場合にはユーザーからチームが保有す\xe3"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "SilentInvite", Desc: "ウエルカムメールを送信しません (SSOとドメイ\xe3\x83"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch reinvite`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.delete"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "Silent", Desc: "招待メールを送信しません (SSOが必須となります)", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "チームメンバーのメールアドレス.", TypeName: "string"},
  		&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.permanent_delete"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレスでフィルタリングし\xe3"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member folder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "DstMemberEmail", Desc: "コピー先チームメンバーのメールアドレス", TypeName: "string"},
  		&{Name: "DstPath", Desc: "コピー先チームメンバーのパス. ルート (/) パス"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
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
# コマンド仕様の変更: `dropbox team member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "削除済メンバーを含めます.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_data.member"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "Quota", Desc: "カスタムの容量制限 (1TB = 1024GB). 0の場合、容量\xe5"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "UpdateUnverified", Desc: "アカウントのメールアドレスが確認されていな\xe3"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch invisible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch visible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 5 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "Trueの場合、共有フォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "Trueの場合、チームフォルダを含めます", Default: "true", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_data.member"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllColumns", Desc: "全てのカラムを表示します", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("sharing.read"),
  				string("team_data.member"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace summary`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "SkipMemberSummary", Desc: "メンバー名前空間のスキャンをスキップする", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file batch copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file sync batch up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 8 identical elements
  		&{Name: "NameNameSuffix", Desc: "名前によるフィルター. 名前の後方一致による\xe3\x83"...},
  		&{Name: "Overwrite", Desc: "ターゲットパス上に既存のファイルが存在する\xe5"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch leave`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "KeepCopy", Desc: "フォルダから抜ける時にフォルダの内容をコピ\xe3"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("members.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "MemberPolicy", Desc: "この共有フォルダーのメンバーになれる人.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("members.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "SharedLinkPolicy", Desc: "この共有フォルダー内のコンテンツに作成され\xe3"..., Default: "anyone", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "LeaveCopy", Desc: "trueの場合、この共有フォルダのメンバーは、共"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("members.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder isolate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "KeepCopy", Desc: "フォルダから抜ける時にフォルダの内容をコピ\xe3"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Message", Desc: "カスタム招待メッセージ", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "Silent", Desc: "招待メールを送信しない", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "LeaveCopy", Desc: "trueの場合、この共有フォルダのメンバーは、共"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "共有フォルダーのID.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "IncludeMounted", Desc: "マウントされたフォルダーを含む.", Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink cap expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink cap visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "新しい視認性設定", Default: "team_only", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink delete links`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink delete member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "Visibility", Desc: "可視性によるリンクのフィルタリング (all/public/"..., Default: "all", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update password`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "新しい視認性設定", Default: "team_only", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "SyncSetting", Desc: "チームフォルダの同期設定", Default: "default", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "操作バッチサイズ", Default: "100", TypeName: "int", ...},
  		&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "TeamFolder", Desc: "チームフォルダ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("team_data.member"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "TeamFolder", Desc: "チームフォルダ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "TeamFolder", Desc: "チームフォルダ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AdminGroupName", Desc: "管理者操作のための仮グループ名", Default: "watermint-toolbox-admin", TypeName: "string", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 7 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AdminGroupName", Desc: "管理者操作のための仮グループ名", Default: "watermint-toolbox-admin", TypeName: "string", ...},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 7 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...},
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder partial replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "ファイルパス標準を選択します。これは、特にD"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("account_info.write"),
  				string("files.content.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "DstPath", Desc: "チームフォルダからの相対パス (チームフォル\xe3\x83"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "DstTeamFolderName", Desc: "送信先チームフォルダ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "src",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "SrcPath", Desc: "チームフォルダからの相対パス (チームフォル\xe3\x83"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "SrcTeamFolderName", Desc: "転送元のチームフォルダ名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...},
  		&{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder sync setting list`



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
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("team_data.content.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "ScanAll", Desc: "すべての深さのスキャンを実行する（フォルダ\xe6"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "ShowAll", Desc: "スキャンしたフォルダーをすべて表示する", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder sync setting update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("team_data.content.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
