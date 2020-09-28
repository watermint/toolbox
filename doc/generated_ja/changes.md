# `リリース 75` から `リリース 76` までの変更点

# 追加されたコマンド


| コマンド          | タイトル      |
|-------------------|---------------|
| dev replay recipe | Replay recipe |



# 削除されたコマンド


| コマンド        | タイトル      |
|-----------------|---------------|
| dev test replay | Replay recipe |



# コマンド仕様の変更: `dev ci artifact up`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path to upload", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string", ...},
  		&{
  			Name:     "Timeout",
  			Desc:     "Operation timeout in seconds",
- 			Default:  "30",
+ 			Default:  "60",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  	},
  }
```
