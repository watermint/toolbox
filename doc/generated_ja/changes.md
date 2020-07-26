# `リリース 71` から `リリース 72` までの変更点

# 追加されたコマンド


| コマンド                                    | タイトル                                |
|---------------------------------------------|-----------------------------------------|
| dev stage gmail                             | Gmail command                           |
| dev stage scoped                            | Dropbox scoped OAuth app test           |
| services google mail filter add             | Add a filter.                           |
| services google mail filter batch add       | Batch adding/deleting labels with query |
| services google mail filter delete          | Delete a filter                         |
| services google mail filter list            | List filters                            |
| services google mail label add              | Add a label                             |
| services google mail label delete           | Delete a label                          |
| services google mail label list             | List email labels                       |
| services google mail label rename           | Rename a label                          |
| services google mail message label add      | Add labels to the message               |
| services google mail message label delete   | Remove labels from the message          |
| services google mail message list           | List messages                           |
| services google mail message processed list | List messages in processed format.      |
| services google mail thread list            | List threads                            |



# コマンド仕様の変更: `dev doc`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Badge", Desc: "Include badges of build status", Default: "true", TypeName: "bool", ...},
  		&{Name: "CommandPath", Desc: "Relative path to generate command manuals", Default: "doc/generated/", TypeName: "string", ...},
- 		&{Name: "Filename", Desc: "Filename", Default: "README.md", TypeName: "string"},
  		&{
- 			Name:    "Lang",
+ 			Name:    "DocLang",
  			Desc:    "Language",
  			Default: "",
  			... // 2 identical fields
  		},
+ 		&{Name: "Filename", Desc: "Filename", Default: "README.md", TypeName: "string"},
  	},
  }
```
# コマンド仕様の変更: `dev util curl`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BufferSize", Desc: "Size of buffer", Default: "65536", TypeName: "domain.common.model.mo_int.range_int", ...},
  		&{
  			Name:     "Record",
  			Desc:     "Capture record(s) for the test",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
