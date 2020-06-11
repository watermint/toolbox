# `リリース 68` から `リリース 69` までの変更点

# 追加されたコマンド


| コマンド                 | タイトル                                                     |
|--------------------------|--------------------------------------------------------------|
| team content member list | List team folder & shared folder members                     |
| team content policy list | List policies of team folders and shared folders in the team |



# 削除されたコマンド


| コマンド            | タイトル                                                     |
|---------------------|--------------------------------------------------------------|
| team content member | List team folder & shared folder members                     |
| team content policy | List policies of team folders and shared folders in the team |



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
- 			Default:  "153600",
+ 			Default:  "4096",
  			TypeName: "domain.common.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
  				"max":   float64(153600),
  				"min":   float64(1),
- 				"value": float64(153600),
+ 				"value": float64(4096),
  			},
  		},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "FailOnError", Desc: "Returns error when any error happens while the operation. This command will not return any error when this flag is not enabled. All errors are written in the report.", Default: "false", TypeName: "bool"},
  		... // 2 identical elements
  	},
  }
```
