---
layout: release
title: Changes of Release 88
lang: en
---

# Changes between `Release 88` to `Release 89`

# Commands added


| Command                            | Title                                                         |
|------------------------------------|---------------------------------------------------------------|
| dev stage http_range               | HTTP Range request proof of concept                           |
| file export url                    | Export a document from the URL                                |
| file paper append                  | Append the content to the end of the existing Paper doc       |
| file paper create                  | Create new Paper in the path                                  |
| file paper overwrite               | Overwrite existing Paper document                             |
| file paper prepend                 | Append the content to the beginning of the existing Paper doc |
| services google mail sendas add    | Creates a custom "from" send-as alias                         |
| services google mail sendas delete | Deletes the specified send-as alias                           |
| services google mail sendas list   | Lists the send-as aliases for the specified account           |
| sharedlink info                    | Get information about the shared link                         |



# Command spec changed: `config disable`



## Command configuration changed


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
# Command spec changed: `config enable`



## Command configuration changed

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
# Command spec changed: `config features`



## Command configuration changed

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
# Command spec changed: `connect business_audit`



## Command configuration changed

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
# Command spec changed: `connect business_file`



## Command configuration changed

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
# Command spec changed: `connect business_info`



## Command configuration changed

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
# Command spec changed: `connect business_mgmt`



## Command configuration changed

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
# Command spec changed: `connect user_file`



## Command configuration changed

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
# Command spec changed: `dev benchmark local`



## Command configuration changed

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
# Command spec changed: `dev benchmark upload`



## Command configuration changed

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
# Command spec changed: `dev benchmark uploadlink`



## Command configuration changed

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
# Command spec changed: `dev build catalogue`



## Command configuration changed

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
# Command spec changed: `dev build doc`



## Command configuration changed

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
# Command spec changed: `dev build license`



## Command configuration changed

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
# Command spec changed: `dev build preflight`



## Command configuration changed

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
# Command spec changed: `dev build readme`



## Command configuration changed

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
# Command spec changed: `dev ci artifact connect`



## Command configuration changed

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
# Command spec changed: `dev ci artifact up`



## Command configuration changed

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
# Command spec changed: `dev ci auth connect`



## Command configuration changed

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
# Command spec changed: `dev ci auth export`



## Command configuration changed

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
# Command spec changed: `dev ci auth import`



## Command configuration changed

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
# Command spec changed: `dev diag endpoint`



## Command configuration changed

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
# Command spec changed: `dev diag throughput`



## Command configuration changed

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
# Command spec changed: `dev kvs dump`



## Command configuration changed

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
# Command spec changed: `dev release candidate`



## Command configuration changed

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
# Command spec changed: `dev release publish`



## Command configuration changed

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
# Command spec changed: `dev replay approve`



## Command configuration changed

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
# Command spec changed: `dev replay bundle`



## Command configuration changed

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
# Command spec changed: `dev replay recipe`



## Command configuration changed

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
# Command spec changed: `dev replay remote`



## Command configuration changed

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
# Command spec changed: `dev spec diff`



## Command configuration changed

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
# Command spec changed: `dev spec doc`



## Command configuration changed

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
# Command spec changed: `dev stage dbxfs`



## Command configuration changed

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
# Command spec changed: `dev stage gmail`



## Command configuration changed

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
# Command spec changed: `dev stage griddata`



## Command configuration changed

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
# Command spec changed: `dev stage gui`



## Command configuration changed

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
# Command spec changed: `dev stage scoped`



## Command configuration changed

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
# Command spec changed: `dev stage teamfolder`



## Command configuration changed

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
# Command spec changed: `dev stage upload_append`



## Command configuration changed

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
# Command spec changed: `dev test echo`



## Command configuration changed

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
# Command spec changed: `dev test kvsfootprint`



## Command configuration changed

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
# Command spec changed: `dev test monkey`



## Command configuration changed

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
# Command spec changed: `dev test recipe`



## Command configuration changed

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
# Command spec changed: `dev test resources`



## Command configuration changed

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
# Command spec changed: `dev test setup teamsharedlink`



## Command configuration changed

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
# Command spec changed: `dev util anonymise`



## Command configuration changed

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
# Command spec changed: `dev util curl`



## Command configuration changed

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
# Command spec changed: `dev util image jpeg`



## Command configuration changed

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
# Command spec changed: `dev util wait`



## Command configuration changed

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
# Command spec changed: `file archive local`



## Command configuration changed

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
# Command spec changed: `file compare account`



## Command configuration changed

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
# Command spec changed: `file compare local`



## Command configuration changed

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
# Command spec changed: `file copy`



## Command configuration changed

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
# Command spec changed: `file delete`



## Command configuration changed

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
# Command spec changed: `file dispatch local`



## Command configuration changed

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
# Command spec changed: `file download`



## Command configuration changed

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
# Command spec changed: `file export doc`



## Command configuration changed

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
# Command spec changed: `file import batch url`



## Command configuration changed

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
# Command spec changed: `file import url`



## Command configuration changed

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
# Command spec changed: `file info`



## Command configuration changed

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
# Command spec changed: `file list`



## Command configuration changed

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
# Command spec changed: `file lock acquire`



## Command configuration changed

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
# Command spec changed: `file lock all release`



## Command configuration changed

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
# Command spec changed: `file lock batch acquire`



## Command configuration changed

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
# Command spec changed: `file lock batch release`



## Command configuration changed

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
# Command spec changed: `file lock list`



## Command configuration changed

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
# Command spec changed: `file lock release`



## Command configuration changed

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
# Command spec changed: `file merge`



## Command configuration changed

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
# Command spec changed: `file mount list`



## Command configuration changed

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
# Command spec changed: `file move`



## Command configuration changed

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
# Command spec changed: `file replication`



## Command configuration changed

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
# Command spec changed: `file restore`



## Command configuration changed

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
# Command spec changed: `file search content`



## Command configuration changed

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
# Command spec changed: `file search name`



## Command configuration changed

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
# Command spec changed: `file size`



## Command configuration changed

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
# Command spec changed: `file sync down`



## Command configuration changed

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
# Command spec changed: `file sync online`



## Command configuration changed

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
# Command spec changed: `file sync up`



## Command configuration changed

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
# Command spec changed: `file watch`



## Command configuration changed

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
# Command spec changed: `filerequest create`



## Command configuration changed

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
# Command spec changed: `filerequest delete closed`



## Command configuration changed

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
# Command spec changed: `filerequest delete url`



## Command configuration changed

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
# Command spec changed: `filerequest list`



## Command configuration changed

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
# Command spec changed: `group add`



## Command configuration changed

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
# Command spec changed: `group batch delete`



## Command configuration changed

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
# Command spec changed: `group delete`



## Command configuration changed

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
# Command spec changed: `group folder list`



## Command configuration changed

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
# Command spec changed: `group list`



## Command configuration changed

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
# Command spec changed: `group member add`



## Command configuration changed

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
# Command spec changed: `group member batch add`



## Command configuration changed

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
# Command spec changed: `group member batch delete`



## Command configuration changed

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
# Command spec changed: `group member batch update`



## Command configuration changed

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
# Command spec changed: `group member delete`



## Command configuration changed

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
# Command spec changed: `group member list`



## Command configuration changed

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
# Command spec changed: `group rename`



## Command configuration changed

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
# Command spec changed: `image info`



## Command configuration changed

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
# Command spec changed: `job history archive`



## Command configuration changed

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
# Command spec changed: `job history delete`



## Command configuration changed

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
# Command spec changed: `job history list`



## Command configuration changed

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
# Command spec changed: `job history ship`



## Command configuration changed

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
# Command spec changed: `job log jobid`



## Command configuration changed

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
# Command spec changed: `job log kind`



## Command configuration changed

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
# Command spec changed: `job log last`



## Command configuration changed

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
# Command spec changed: `license`



## Command configuration changed

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
# Command spec changed: `member clear externalid`



## Command configuration changed

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
# Command spec changed: `member delete`



## Command configuration changed

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
# Command spec changed: `member detach`



## Command configuration changed

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
# Command spec changed: `member file lock all release`



## Command configuration changed

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
# Command spec changed: `member file lock list`



## Command configuration changed

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
# Command spec changed: `member file lock release`



## Command configuration changed

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
# Command spec changed: `member file permdelete`



## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 5 identical fields
  	CliArgs:         "-path /DROPBOX/PATH/TO/PERM_DELETE",
  	CliNote:         "",
- 	ConnUsePersonal: true,
+ 	ConnUsePersonal: false,
- 	ConnUseBusiness: false,
+ 	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Team member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_scoped_",
- 				"individual",
+ 				"team",
  			}, ""),
  			TypeAttr: []any{string("files.permanent_delete"), string("team_data.member"), string("members.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# Command spec changed: `member folder list`



## Command configuration changed

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
# Command spec changed: `member folder replication`



## Command configuration changed

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
# Command spec changed: `member invite`



## Command configuration changed

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
# Command spec changed: `member list`



## Command configuration changed

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
# Command spec changed: `member quota list`



## Command configuration changed

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
# Command spec changed: `member quota update`



## Command configuration changed

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
# Command spec changed: `member quota usage`



## Command configuration changed

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
# Command spec changed: `member reinvite`



## Command configuration changed

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
# Command spec changed: `member replication`



## Command configuration changed

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
# Command spec changed: `member update email`



## Command configuration changed

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
# Command spec changed: `member update externalid`



## Command configuration changed

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
# Command spec changed: `member update invisible`



## Command configuration changed

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
# Command spec changed: `member update profile`



## Command configuration changed

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
# Command spec changed: `member update visible`



## Command configuration changed

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
# Command spec changed: `services asana team list`



## Command configuration changed

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
# Command spec changed: `services asana team project list`



## Command configuration changed

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
# Command spec changed: `services asana team task list`



## Command configuration changed

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
# Command spec changed: `services asana workspace list`



## Command configuration changed

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
# Command spec changed: `services asana workspace project list`



## Command configuration changed

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
# Command spec changed: `services github content get`



## Command configuration changed

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
# Command spec changed: `services github content put`



## Command configuration changed

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
# Command spec changed: `services github issue list`



## Command configuration changed

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
# Command spec changed: `services github profile`



## Command configuration changed

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
# Command spec changed: `services github release asset download`



## Command configuration changed

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
# Command spec changed: `services github release asset list`



## Command configuration changed

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
# Command spec changed: `services github release asset upload`



## Command configuration changed

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
# Command spec changed: `services github release draft`



## Command configuration changed

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
# Command spec changed: `services github release list`



## Command configuration changed

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
# Command spec changed: `services github tag create`



## Command configuration changed

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
# Command spec changed: `services google mail filter add`



## Command configuration changed

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
# Command spec changed: `services google mail filter batch add`



## Command configuration changed

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
# Command spec changed: `services google mail filter delete`



## Command configuration changed

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
# Command spec changed: `services google mail filter list`



## Command configuration changed

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
# Command spec changed: `services google mail label add`



## Command configuration changed

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
# Command spec changed: `services google mail label delete`



## Command configuration changed

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
# Command spec changed: `services google mail label list`



## Command configuration changed

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
# Command spec changed: `services google mail label rename`



## Command configuration changed

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
# Command spec changed: `services google mail message label add`



## Command configuration changed

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
# Command spec changed: `services google mail message label delete`



## Command configuration changed

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
# Command spec changed: `services google mail message list`



## Command configuration changed

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
# Command spec changed: `services google mail message processed list`



## Command configuration changed

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
# Command spec changed: `services google mail thread list`



## Command configuration changed

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
# Command spec changed: `services google sheets sheet append`



## Command configuration changed

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
# Command spec changed: `services google sheets sheet clear`



## Command configuration changed

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
# Command spec changed: `services google sheets sheet export`



## Command configuration changed

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
# Command spec changed: `services google sheets sheet import`



## Command configuration changed

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
# Command spec changed: `services google sheets sheet list`



## Command configuration changed

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
# Command spec changed: `services google sheets spreadsheet create`



## Command configuration changed

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
# Command spec changed: `services slack conversation list`



## Command configuration changed

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
# Command spec changed: `sharedfolder list`



## Command configuration changed

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
# Command spec changed: `sharedfolder member list`



## Command configuration changed

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
# Command spec changed: `sharedlink create`



## Command configuration changed

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
# Command spec changed: `sharedlink delete`



## Command configuration changed

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
# Command spec changed: `sharedlink file list`



## Command configuration changed

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
# Command spec changed: `sharedlink list`



## Command configuration changed

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
# Command spec changed: `team activity batch user`



## Command configuration changed

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
# Command spec changed: `team activity daily event`



## Command configuration changed

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
# Command spec changed: `team activity event`



## Command configuration changed

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
# Command spec changed: `team activity user`



## Command configuration changed

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
# Command spec changed: `team content member list`



## Command configuration changed

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
# Command spec changed: `team content mount list`



## Command configuration changed

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
# Command spec changed: `team content policy list`



## Command configuration changed

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
# Command spec changed: `team device list`



## Command configuration changed

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
# Command spec changed: `team device unlink`



## Command configuration changed

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
# Command spec changed: `team diag explorer`



## Command configuration changed

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
# Command spec changed: `team feature`



## Command configuration changed

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
# Command spec changed: `team filerequest clone`



## Command configuration changed

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
# Command spec changed: `team filerequest list`



## Command configuration changed

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
# Command spec changed: `team info`



## Command configuration changed

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
# Command spec changed: `team linkedapp list`



## Command configuration changed

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
# Command spec changed: `team namespace file list`



## Command configuration changed

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
# Command spec changed: `team namespace file size`



## Command configuration changed

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
# Command spec changed: `team namespace list`



## Command configuration changed

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
# Command spec changed: `team namespace member list`



## Command configuration changed

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
# Command spec changed: `team report activity`



## Command configuration changed

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
# Command spec changed: `team report devices`



## Command configuration changed

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
# Command spec changed: `team report membership`



## Command configuration changed

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
# Command spec changed: `team report storage`



## Command configuration changed

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
# Command spec changed: `team sharedlink delete links`



## Command configuration changed

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
# Command spec changed: `team sharedlink delete member`



## Command configuration changed

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
# Command spec changed: `team sharedlink list`



## Command configuration changed

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
# Command spec changed: `team sharedlink update expiry`



## Command configuration changed

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
# Command spec changed: `team sharedlink update password`



## Command configuration changed

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
# Command spec changed: `team sharedlink update visibility`



## Command configuration changed

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
# Command spec changed: `teamfolder add`



## Command configuration changed

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
# Command spec changed: `teamfolder archive`



## Command configuration changed

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
# Command spec changed: `teamfolder batch archive`



## Command configuration changed

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
# Command spec changed: `teamfolder batch permdelete`



## Command configuration changed

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
# Command spec changed: `teamfolder batch replication`



## Command configuration changed

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
# Command spec changed: `teamfolder file list`



## Command configuration changed

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
# Command spec changed: `teamfolder file lock all release`



## Command configuration changed

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
# Command spec changed: `teamfolder file lock list`



## Command configuration changed

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
# Command spec changed: `teamfolder file lock release`



## Command configuration changed

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
# Command spec changed: `teamfolder file size`



## Command configuration changed

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
# Command spec changed: `teamfolder list`



## Command configuration changed

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
# Command spec changed: `teamfolder member add`



## Command configuration changed

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
# Command spec changed: `teamfolder member delete`



## Command configuration changed

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
# Command spec changed: `teamfolder member list`



## Command configuration changed

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
# Command spec changed: `teamfolder partial replication`



## Command configuration changed

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
# Command spec changed: `teamfolder permdelete`



## Command configuration changed

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
# Command spec changed: `teamfolder policy list`



## Command configuration changed

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
# Command spec changed: `teamfolder replication`



## Command configuration changed

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
# Command spec changed: `util date today`



## Command configuration changed

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
# Command spec changed: `util datetime now`



## Command configuration changed

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
# Command spec changed: `util decode base_32`



## Command configuration changed

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
# Command spec changed: `util decode base_64`



## Command configuration changed

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
# Command spec changed: `util encode base_32`



## Command configuration changed

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
# Command spec changed: `util encode base_64`



## Command configuration changed

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
# Command spec changed: `util qrcode create`



## Command configuration changed

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
# Command spec changed: `util qrcode wifi`



## Command configuration changed

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
# Command spec changed: `util time now`



## Command configuration changed

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
# Command spec changed: `util unixtime format`



## Command configuration changed

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
# Command spec changed: `util unixtime now`



## Command configuration changed

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
# Command spec changed: `util xlsx create`



## Command configuration changed

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
# Command spec changed: `util xlsx sheet export`



## Command configuration changed

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
# Command spec changed: `util xlsx sheet import`



## Command configuration changed

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
# Command spec changed: `util xlsx sheet list`



## Command configuration changed

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
# Command spec changed: `version`



## Command configuration changed

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
