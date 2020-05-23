# `リリース 67` から `リリース 68` までの変更点

# 追加されたコマンド


| コマンド            | タイトル                 |
|---------------------|--------------------------|
| dev util image jpeg | Create dummy image files |



# 削除されたコマンド


| コマンド  | タイトル           |
|-----------|--------------------|
| dev dummy | Create dummy files |



# コマンド仕様の変更: `member delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"},
+ 		&{
+ 			Name:     "TransferDestMember",
+ 			Desc:     "If provided, files from the deleted member account will be transferred to this user.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "TransferNotifyAdminEmailOnError",
+ 			Desc:     "If provided, errors during the transfer process will be sent via email to this user.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
  		&{Name: "WipeData", Desc: "If true, controls if the user's data will be deleted on their linked devices", Default: "true", TypeName: "bool"},
  	},
  }
```
