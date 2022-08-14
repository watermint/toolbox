---
layout: release
title: リリースの変更点 91
lang: ja
---

# `リリース 91` から `リリース 92` までの変更点

# 追加されたコマンド


| コマンド                       | タイトル                                   |
|--------------------------------|--------------------------------------------|
| team sharedlink cap expiry     | チーム内の共有リンクに有効期限の上限を設定 |
| team sharedlink cap visibility | チーム内の共有リンクに可視性の上限を設定   |



# 削除されたコマンド


| コマンド               | タイトル                                     |
|------------------------|----------------------------------------------|
| connect business_audit | チーム監査アクセスに接続する                 |
| connect business_file  | チームファイルアクセスに接続する             |
| connect business_info  | チームの情報アクセスに接続する               |
| connect business_mgmt  | チームの管理アクセスに接続する               |
| connect user_file      | ユーザーのファイルアクセスに接続する         |
| dev ci auth export     | エンドツーエンドテストのトークンを出力します |
| team diag explorer     | チーム全体の情報をレポートします             |



# コマンド仕様の変更: `dev build preflight`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Quick", Desc: "クイックモード", Default: "false", TypeName: "bool"},
+ 	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth connect`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 5 identical fields
  	CliArgs:         "",
  	CliNote:         "",
- 	ConnUsePersonal: true,
+ 	ConnUsePersonal: false,
- 	ConnUseBusiness: true,
+ 	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Audit":  "business_audit",
- 		"File":   "business_file",
- 		"Full":   "user_full",
  		"Github": "github_repo",
- 		"Info":   "business_info",
- 		"Mgmt":   "business_management",
  	},
  	Services: []string{
- 		"dropbox",
- 		"dropbox_business",
  		"github",
  	},
  	IsSecret:  true,
  	IsConsole: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Audit",
- 			Desc:     "Dropbox Business Audit スコープで認証",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
- 		},
- 		&{
- 			Name:     "File",
- 			Desc:     "Dropbox Business member file access スコープで認証",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
- 		},
- 		&{
- 			Name:     "Full",
- 			Desc:     "Dropbox user full access スコープで認証",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
- 		},
  		&{Name: "Github", Desc: "GitHubへのデプロイメントのためのアカウント別名", Default: "deploy", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
- 		&{
- 			Name:     "Info",
- 			Desc:     "Dropbox Business info スコープで認証",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
- 		},
- 		&{
- 			Name:     "Mgmt",
- 			Desc:     "Dropbox Business management スコープで認証",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
- 		},
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
  	... // 5 identical fields
  	CliArgs:         "",
  	CliNote:         "",
- 	ConnUsePersonal: true,
+ 	ConnUsePersonal: false,
- 	ConnUseBusiness: true,
+ 	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Audit":  "business_audit",
- 		"File":   "business_file",
- 		"Full":   "user_full",
  		"Github": "github_repo",
- 		"Info":   "business_info",
- 		"Mgmt":   "business_management",
  		"Peer":   "github_public",
  	},
  	Services: []string{
- 		"dropbox",
- 		"dropbox_business",
  		"github",
  	},
  	IsSecret:  true,
  	IsConsole: true,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Peer",
  			Desc:    "アカウントの別名",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("members.read"), string("sharing.read"), string("team_data.member")},
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
  	Name:  "expiry",
  	Title: "チーム内の公開されている共有リンクについて\xe6"...,
  	Desc: (
  		"""
  		注：リリース87以降、このコマンドは、アップデートする共有リンクを選択するためのファイルを受け取ります. チーム内のすべての共有リンクの有効期限を更新したい場合は、`team sharedlink list`の組み合わせをご検討ください. 例えば、[jq](https://stedolan.github.io/jq/)というコマンドに慣れていれば、以下のように同等の操作を行うことができます（すべての公開リンクに対して28日以内に強制失効させる）.
  		
  		```
- 		tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -days 28
+ 		tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -at +720h
  		```
+ 		リリース92以降、このコマンドは引数 `-days` を受け取りません. 相対的な日時を設定したい場合は、`+720h`（720時間＝30日）のように`-at +HOURh`を使用してください.
  		
+ 		コマンド `team sharedlink update` は、共有リンクに値を設定するためのものです. コマンド `team sharedlink cap` は、共有リンクにキャップ値を設定するためのものです. 例：有効期限を2021-05-06に設定して、`team sharedlink update expiry`で設定した場合. このコマンドは、既存のリンクが2021-05-04のように短い有効期限を持っている場合でも、有効期限を2021-05-06に更新します.
  		"""
  	),
  	Remarks: "(非可逆な操作です)",
  	Path:    "team sharedlink update expiry",
  	CliArgs: strings.Join({
  		"-file /PATH/TO/DATA_FILE.csv -",
- 		"days 28",
+ 		"at +720h",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_time.time_impl",
- 			TypeAttr: map[string]any{"optional": bool(true)},
+ 			TypeAttr: map[string]any{"optional": bool(false)},
  		},
- 		&{
- 			Name:     "Days",
- 			Desc:     "新しい有効期限までの日時",
- 			Default:  "0",
- 			TypeName: "essentials.model.mo_int.range_int",
- 			TypeAttr: map[string]any{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
- 		},
  		&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
