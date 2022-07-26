---
layout: release
title: リリースの変更点 97
lang: ja
---

# `リリース 97` から `リリース 98` までの変更点

# 追加されたコマンド


| コマンド                     | タイトル                                                         |
|------------------------------|------------------------------------------------------------------|
| group batch add              | グループの一括追加                                               |
| team admin group role add    | グループのメンバーにロールを追加する                             |
| team admin group role delete | 例外グループのメンバーを除くすべてのメンバーからロールを削除する |
| team admin list              | メンバーの管理者権限一覧                                         |
| team admin role add          | メンバーに新しいロールを追加する                                 |
| team admin role clear        | メンバーからすべての管理者ロールを削除する                       |
| team admin role delete       | メンバーからロールを削除する                                     |
| team admin role list         | チームの管理者の役割を列挙                                       |
| util image placeholder       | プレースホルダー画像の作成                                       |



# コマンド仕様の変更: `dev stage scoped`



## 変更されたレポート: member_list

```
  &dc_recipe.Report{
  	Name: "member_list",
  	Desc: "このレポートはメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 12 identical elements
  		&{Name: "persistent_id", Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で"...},
  		&{Name: "joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member clear externalid`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 15 identical elements
  		&{Name: "result.persistent_id", Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で"...},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member folder list`



## 変更されたレポート: member_with_no_folder

```
  &dc_recipe.Report{
  	Name: "member_with_no_folder",
  	Desc: "このレポートはメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "surname", Desc: "名字"},
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
+ 		&{
+ 			Name: "invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  	},
  }
```
# コマンド仕様の変更: `member invite`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 10 identical elements
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member list`



## 変更されたレポート: member

```
  &dc_recipe.Report{
  	Name: "member",
  	Desc: "このレポートはメンバー一覧を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  	},
  }
```
# コマンド仕様の変更: `member reinvite`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "input.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "input.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "input.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "input.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "input.tag", Desc: "処理のタグ"},
  		... // 5 identical elements
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update email`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 9 identical elements
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update externalid`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 9 identical elements
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update invisible`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 15 identical elements
  		&{Name: "result.persistent_id", Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で"...},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update profile`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 10 identical elements
  		&{Name: "result.display_name", Desc: "ユーザーのDropboxアカウントの表示名称"},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `member update visible`



## 変更されたレポート: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "このレポートは処理結果を出力します.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 15 identical elements
  		&{Name: "result.persistent_id", Desc: "ユーザーに付加できる永続ID. 永続IDはSAML認証で"...},
  		&{Name: "result.joined_on", Desc: "メンバーがチームに参加した日時."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "ユーザーがチームに招待された日付と時間",
+ 		},
  		&{Name: "result.role", Desc: "ユーザーのチームでの役割 (team_admin, user_managemen"...},
  		&{Name: "result.tag", Desc: "処理のタグ"},
  	},
  }
```
# コマンド仕様の変更: `teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...},
  		&{Name: "MemberTypeInternal", Desc: "フォルダメンバーによるフィルター. 内部メン\xe3\x83"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
+ 				string("groups.read"),
+ 				string("members.read"),
  				string("sharing.read"),
  				string("team_data.member"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "スキャンのタイムアウト設定. スキャンタイム\xe3\x82"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
