# `リリース 69` から `リリース 70` までの変更点

# 追加されたコマンド


| コマンド               | タイトル                  |
|------------------------|---------------------------|
| dev test kvsfootprint  | Test KVS memory footprint |
| teamfolder member list | List team folder members  |



# コマンド仕様の変更: `file sync up`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ChunkSizeKb",
  			Desc:     "Upload chunk size in KB",
- 			Default:  "4096",
+ 			Default:  "65536",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(153600),
  				"min":   float64(1),
- 				"value": float64(4096),
+ 				"value": float64(65536),
  			},
  		},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "FailOnError", Desc: "Returns error when any error happens while the operation. This c"..., Default: "false", TypeName: "bool", ...},
  		... // 2 identical elements
  	},
  }
```
