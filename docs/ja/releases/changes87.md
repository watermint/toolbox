---
layout: release
title: リリースの変更点 86
lang: ja
---

# `リリース 86` から `リリース 87` までの変更点

# 追加されたコマンド


| コマンド                          | タイトル                           |
|-----------------------------------|------------------------------------|
| dev test setup teamsharedlink     | デモ用共有リンクの作成             |
| team sharedlink delete links      | 共有リンクの一括削除               |
| team sharedlink delete member     | メンバーの共有リンクをすべて削除   |
| team sharedlink update password   | 共有リンクのパスワードの設定・更新 |
| team sharedlink update visibility | 共有リンクの可視性の更新           |



# コマンド仕様の変更: `dev stage upload_append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "dropbox_scoped_individual"},
  	Services:       {"dropbox"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 10 identical fields
  }
```
# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "エクスポートするDropbox上のドキュメントパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
+ 		&{
+ 			Name:     "Format",
+ 			Desc:     "エクスポート書式",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "LocalPath", Desc: "保存先ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
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
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name: "Visibility",
  			Desc: strings.Join({
  				"\xe5",
- 				"\x87\xba力するリンクを可視性にてフィルターします (",
+ 				"\x8f\xaf視性によるリンクのフィルタリング (all/",
  				"public/team_only/password)",
  			}, ""),
  			Default:  "all",
  			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]any{"options": []any{string("all"), string("public"), string("team_only"), string("password"), ...}},
  		},
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
  	Name:            "expiry",
  	Title:           "チーム内の公開されている共有リンクについて\xe6"...,
- 	Desc:            "",
+ 	Desc:            "注：リリース87以降、このコマンドは、アップデートする共有リンクを選択するためのファイルを受け取ります. チーム内のすべての共有リンクの有効期限を更新したい場合は、`team sharedlink l"...,
  	Remarks:         "(非可逆な操作です)",
  	Path:            "team sharedlink update expiry",
- 	CliArgs:         "-days 28",
+ 	CliArgs:         "-file /PATH/TO/DATA_FILE.csv -days 28",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "新しい有効期限の日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Days", Desc: "新しい有効期限までの日時", Default: "0", TypeName: "essentials.model.mo_int.range_int", ...},
- 		&{
- 			Name:     "Peer",
- 			Desc:     "アカウントの別名",
- 			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
- 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "データファイルへのパス",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
- 		&{
- 			Name:     "Visibility",
- 			Desc:     "対象となるリンクの公開範囲",
- 			Default:  "public",
- 			TypeName: "essentials.model.mo_string.select_string",
- 			TypeAttr: map[string]any{
- 				"options": []any{
- 					string("public"), string("team_only"), string("password"),
- 					string("team_and_password"), ...,
- 				},
- 			},
- 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "アカウントの別名",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
+ 			TypeAttr: []any{string("members.read"), string("sharing.write"), string("team_data.member")},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称          | 説明                                |
|---------------|-------------------------------------|
| operation_log | このレポートは処理結果を出力します. |



## 削除されたレポート


| 名称    | 説明                                                                    |
|---------|-------------------------------------------------------------------------|
| skipped | このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します. |
| updated | このレポートは処理結果を出力します.                                     |


