# `リリース 85` から `リリース 86` までの変更点

# 追加されたコマンド


| コマンド                | タイトル                                           |
|-------------------------|----------------------------------------------------|
| dev stage dbxfs         | Verify Dropbox File System impl. for cached system |
| dev stage upload_append | New upload API test                                |
| util qrcode create      | Create a QR code image file                        |
| util qrcode wifi        | Generate QR code for WIFI configuration            |



# コマンド仕様の変更: `dev benchmark upload`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "BlockBlockSize",
+ 			Desc:     "Block size for batch upload",
+ 			Default:  "40",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(1000), "min": float64(1), "value": float64(40)},
+ 		},
- 		&{
- 			Name:     "ChunkSizeKb",
- 			Desc:     "Upload chunk size in KiB",
- 			Default:  "65536",
- 			TypeName: "essentials.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
- 		},
+ 		&{
+ 			Name:     "Method",
+ 			Desc:     "Upload method",
+ 			Default:  "block",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("block"), string("sequential")}},
+ 		},
  		&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.write")},
  		},
+ 		&{
+ 			Name:     "PreScan",
+ 			Desc:     "Pre-scan destination path",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
- 		&{
- 			Name:     "Shard",
- 			Desc:     "Number of shared folders to distribute namespace",
- 			Default:  "1",
- 			TypeName: "int",
- 		},
+ 		&{
+ 			Name:     "SeqChunkSizeKb",
+ 			Desc:     "Upload chunk size in KiB",
+ 			Default:  "65536",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
+ 		},
  		&{Name: "SizeMaxKb", Desc: "Maximum file size (KiB).", Default: "2048", TypeName: "int", ...},
  		&{Name: "SizeMinKb", Desc: "Minimum file size (KiB).", Default: "0", TypeName: "int", ...},
+ 		&{Name: "Verify", Desc: "Verify after upload", Default: "false", TypeName: "bool"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  }
```

# コマンド仕様の変更: `dev release publish`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ArtifactPath", Desc: "Path to artifacts", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
  			Name:     "Branch",
  			Desc:     "Target branch",
- 			Default:  "master",
+ 			Default:  "main",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  }
```

# コマンド仕様の変更: `file sync up`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "ChunkSizeKb",
- 			Desc:     "Upload chunk size in KB",
- 			Default:  "65536",
- 			TypeName: "essentials.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
- 		},
+ 		&{
+ 			Name:     "BatchSize",
+ 			Desc:     "Batch commit size",
+ 			Default:  "50",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(1000), "min": float64(1), "value": float64(50)},
+ 		},
  		&{Name: "Delete", Desc: "Delete Dropbox file if a file removed locally", Default: "false", TypeName: "bool", ...},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 5 identical elements
  		&{Name: "Overwrite", Desc: "Overwrite existing file on the target path if that exists.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
- 		&{
- 			Name:     "WorkPath",
- 			Desc:     "Temporary path",
- 			TypeName: "essentials.model.mo_string.opt_string",
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  }
```
