---
layout: release
title: Changes of Release 81
lang: en
---

# Changes between `Release 81` to `Release 82`

# Command spec changed: `file sync up`



## Command configuration changed


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
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `team diag explorer`



## Command configuration changed


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
  	... // 11 identical fields
  }
```
