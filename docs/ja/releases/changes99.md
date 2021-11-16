---
layout: release
title: リリースの変更点 98
lang: ja
---

# `リリース 98` から `リリース 99` までの変更点

# 追加されたコマンド


| コマンド                                    | タイトル                                      |
|---------------------------------------------|-----------------------------------------------|
| file share info                             | ファイルの共有情報を取得する                  |
| sharedfolder member add                     | 共有フォルダへのメンバーの追加                |
| sharedfolder member delete                  | 共有フォルダからメンバーを削除する            |
| sharedfolder share                          | フォルダの共有                                |
| sharedfolder unshare                        | フォルダの共有解除                            |
| team runas file batch copy                  | ファイル/フォルダーをメンバーとして一括コピー |
| team runas file sync batch up               | メンバーとして動作する一括同期                |
| team runas sharedfolder batch share         | メンバーのフォルダを一括で共有                |
| team runas sharedfolder batch unshare       | メンバーのフォルダの共有を一括解除            |
| team runas sharedfolder member batch add    | メンバーの共有フォルダにメンバーを一括追加    |
| team runas sharedfolder member batch delete | メンバーの共有フォルダからメンバーを一括削除  |



# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "削除済みファイルを含める", Default: "false", TypeName: "bool", ...},
+ 		&{
+ 			Name:     "IncludeExplicitSharedMembers",
+ 			Desc:     " trueの場合、結果には、各ファイルに明示的なメンバーがいるかどうかを示すフラグが含まれま"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMountedFolders",
+ 			Desc:     " Trueの場合は、マウントされたフォルダ（appフォルダ、sharedフォルダ、teamフォルダ）のエント\xe3\x83"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
