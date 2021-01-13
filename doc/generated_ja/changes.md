# `リリース 81` から `リリース 82` までの変更点

# コマンド仕様の変更: `file sync up`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "NameNamePrefix", Desc: "Filter by name. Filter by name match to the prefix."},
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
- 		&{
- 			Name:     "Peer",
- 			Desc:     "Account alias",
- 			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
- 		},
  		&{
- 			Name:     "SkipExisting",
+ 			Name:     "Overwrite",
- 			Desc:     "Skip existing files. Do not overwrite",
+ 			Desc:     "Overwrite existing file on the target path if that exists.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
  		&{Name: "WorkPath", Desc: "Temporary path", TypeName: "essentials.model.mo_string.opt_string"},
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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
