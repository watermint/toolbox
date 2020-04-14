# `リリース 64` から `リリース 65` までの変更点

# コマンド仕様の変更: `dev ci auth import`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*rc_doc.Value{
  		&{
  			Name:     "EnvName",
  			Desc:     "Environment variable name",
- 			Default:  "TOOLBOX_ENDTOEND",
+ 			Default:  "TOOLBOX_ENDTOEND_TOKEN",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "PeerName", Desc: "Account alias", Default: "end_to_end_test", TypeName: "string"},
  	},
  }

```

# コマンド仕様の変更: `dev desktop stop`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*rc_doc.Value{
  		&{
  			Name:     "WaitSeconds",
  			Desc:     "Try stopping the app after given seconds.",
- 			Default:  "60",
+ 			Default:  "0",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(2.147483647e+09),
  				"min":   float64(0),
- 				"value": float64(60),
+ 				"value": float64(0),
  			},
  		},
  	},
  }

```

# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*rc_doc.Value{
  		... // 2 identical elements
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo"},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool"},
  		&{
  			Name:     "TestResource",
  			Desc:     "Path to test resource",
- 			Default:  "test/dev/resource.json",
+ 			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }

```

# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*rc_doc.Value{
  		&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KB", Default: "153600", TypeName: "domain.common.model.mo_int.range_int", TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)}},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
+ 		&{
+ 			Name:     "FailOnError",
+ 			Desc:     "Returns error when any error happens while the operation. This command will not return any error when this flag is not enabled. All errors are written in the report.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "LocalPath", Desc: "Local file path", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"},
  	},
  }

```

