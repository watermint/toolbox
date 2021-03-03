# `リリース 85` から `リリース 86` までの変更点

# 追加されたコマンド

| コマンド                | タイトル            |
|-------------------------|---------------------|
| dev stage upload_append | New upload API test |



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
+ 			Default:  "50",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(1000), "min": float64(1), "value": float64(50)},
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
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  }
```
