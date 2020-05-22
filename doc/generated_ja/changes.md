# `リリース 67` から `リリース 68` までの変更点

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
+ 		&{
+ 			Name:     "TransferDestMember",
+ 			Desc:     "If provided, files from the deleted member account will be transferred to this user.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "TransferNotifyAdminEmailOnError",
+ 			Desc:     "If provided, errors during the transfer process will be sent via email to this user.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
  		&{Name: "WipeData", Desc: "If true, controls if the user's data will be deleted on their linked devices", Default: "true", TypeName: "bool"},
  	},
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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: []string{"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
