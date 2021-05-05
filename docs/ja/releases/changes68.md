---
layout: release
title: リリースの変更点: 67
lang: ja
---

# `リリース 67` から `リリース 68` までの変更点

# 追加されたコマンド


| コマンド                    | タイトル                                         |
|-----------------------------|--------------------------------------------------|
| dev util image jpeg         | ダミー画像ファイルを作成します                   |
| services github content get | レポジトリのコンテンツメタデータを取得します.    |
| services github content put | レポジトリに小さなテキストコンテンツを格納します |



# 削除されたコマンド


| コマンド  | タイトル                   |
|-----------|----------------------------|
| dev dummy | ダミーファイルを作成します |



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
+ 		&{
+ 			Name:     "TransferDestMember",
+ 			Desc:     "指定された場合は、指定ユーザーに削除するメンバーのコンテンツを転送します.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "TransferNotifyAdminEmailOnError",
+ 			Desc:     "指定された場合は、転送時にエラーが発生した時にこのユーザーにメールを送信します.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
  		&{Name: "WipeData", Desc: "指定した場合にはユーザーのデータがリンクさ\xe3"..., Default: "true", TypeName: "bool", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
