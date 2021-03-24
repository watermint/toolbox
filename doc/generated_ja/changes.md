# `リリース 88` から `リリース 89` までの変更点

# 追加されたコマンド


| コマンド             | タイトル                                                      |
|----------------------|---------------------------------------------------------------|
| file paper append    | Append the content to the end of the existing Paper doc       |
| file paper create    | Create new Paper in the path                                  |
| file paper overwrite | Overwrite existing Paper document                             |
| file paper prepend   | Append the content to the beginning of the existing Paper doc |



# コマンド仕様の変更: `config disable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `config enable`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `config features`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_audit`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_file`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_info`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_mgmt`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect user_file`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev benchmark local`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev benchmark upload`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build catalogue`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build doc`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build license`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build preflight`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build readme`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci artifact connect`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci artifact up`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci auth connect`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci auth export`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci auth import`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev diag endpoint`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev diag throughput`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev kvs dump`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev release candidate`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev release publish`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay approve`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay bundle`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay recipe`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay remote`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev spec diff`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
- 			Name:     "FilePath",
+ 			Name:     "DocLang",
- 			Desc:     "File path to output",
+ 			Desc:     "Document language",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
- 			Name:     "Lang",
+ 			Name:     "FilePath",
- 			Desc:     "Language",
+ 			Desc:     "File path to output",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Release1", Desc: "Release name 1", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Release2", Desc: "Release name 2", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev spec doc`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage dbxfs`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage gmail`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage griddata`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "In", Desc: "Test input grid data"}},
  	GridDataOutput: {&{Name: "Out", Desc: "Output grid data"}},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage gui`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage scoped`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage teamfolder`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage upload_append`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test echo`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test kvsfootprint`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test monkey`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test recipe`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test resources`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test setup teamsharedlink`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{Name: "Query", Desc: "Query", TypeName: "string"},
+ 		&{Name: "Seed", Desc: "Shared link seed value", Default: "0", TypeName: "int"},
  		&{Name: "Visibility", Desc: "Visibility", Default: "random", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util anonymise`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util curl`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util image jpeg`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util wait`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file archive local`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file compare account`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file compare local`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file copy`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file dispatch local`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file download`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file export doc`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file import batch url`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file import url`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file info`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock acquire`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock batch acquire`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock batch release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file merge`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file mount list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file move`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file restore`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file search content`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file search name`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file size`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file sync down`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file sync online`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file sync up`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file watch`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest delete closed`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest delete url`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group batch delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group folder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member batch add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member batch delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member batch update`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group rename`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `image info`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history archive`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history ship`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job log jobid`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job log kind`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job log last`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `license`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member clear externalid`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member detach`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member folder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member folder replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member invite`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member quota list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member quota update`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member quota usage`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member reinvite`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update email`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update externalid`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update invisible`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update profile`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update visible`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana team list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana team project list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana team task list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana workspace list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana workspace project list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github content get`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github content put`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github issue list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github profile`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release asset download`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release asset list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release asset upload`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release draft`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github tag create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter batch add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label rename`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message label add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message label delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message processed list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail thread list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet append`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "Data", Desc: "Input data file"}},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet clear`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet export`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "Exported sheet data"}},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet import`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "Data", Desc: "Input data file"}},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets spreadsheet create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services slack conversation list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedfolder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedfolder member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity batch user`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity daily event`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity event`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity user`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team content member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team content mount list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team content policy list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team device list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team device unlink`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team diag explorer`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"File": "business_file", "Info": "business_info", "Mgmt": "business_management", "Peer": "business_file"},
  	Services:       {"dropbox_business"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 5 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team feature`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team filerequest clone`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team filerequest list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team info`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team linkedapp list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace file size`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report activity`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report devices`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report membership`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report storage`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink delete links`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink delete member`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink update password`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink update visibility`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder archive`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder batch archive`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder batch replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file size`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder member add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	Name:  "add",
  	Title: "Batch adding users/groups to team folders",
  	Desc: (
  		"""
  		This command will do (1) create new team folders or new sub-folders if the team folder does not exist. The command does not (2) change access inheritance setting of any folders, (3) create a group if that not exist. This command is designed to be idempotent. You can safely retry if any errors happen on the operation. The command will not report an error to keep idempotence. For example, the command will not report an error like, the member already have access to the folder.
+ 		
+ 		Example:
+ 		
+ 		* Sales (team folder, editor access for the group "Sales")
+ 			* Sydney (viewer access for individual account sydney@example.com)
+ 			* Tokyo (editor access for the group "Tokyo Deal Desk")
+ 				* Monthly (viewer access for individual account success@example.com)
+ 		* Marketing (team folder, editor access for the group "Marketing")
+ 			* Sydney (editor access for the group "Sydney Sales")
+ 			* Tokyo (viewer access for the group "Tokyo Sales")
+ 		
+ 		1. Prepare CSV like below
+ 		
+ 		```
+ 		Sales,,editor,Sales
+ 		Sales,Sydney,editor,sydney@example.com
+ 		Sales,Tokyo,editor,Tokyo Deal Desk
+ 		Sales,Tokyo/Monthly,viewer,success@example.com
+ 		Marketing,,editor,Marketing
+ 		Marketing,Sydney,editor,Sydney Sales
+ 		Marketing,Tokyo,viewer,Tokyo Sales
+ 		```
+ 		
+ 		2. Then run the command like below
+ 		
+ 		```
+ 		tbx teamfolder member add -file /PATH/TO/DATA.csv
+ 		```
+ 		
+ 		Note: the command will create a team folder if not exist. But the command will not a group if not found. Groups must exist before run this command.
  		"""
  	),
  	Remarks: "(Irreversible operation)",
  	Path:    "teamfolder member add",
  	... // 14 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder partial replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder policy list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util date today`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util datetime now`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util decode base_32`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util decode base_64`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util encode base_32`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util encode base_64`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util qrcode create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util qrcode wifi`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util time now`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util unixtime format`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Format", Desc: "Time format", Default: "iso8601", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Precision", Desc: "Time precision (second/ms/ns)", Default: "second", TypeName: "essentials.model.mo_string.select_string", ...},
+ 		&{Name: "Time", Desc: "Unix Time", Default: "0", TypeName: "int"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util unixtime now`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx sheet export`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "Export data"}},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx sheet import`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "Data", Desc: "Input data file"}},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx sheet list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `version`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
