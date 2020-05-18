# `リリース 66` から `リリース 67` までの変更点

# 追加されたコマンド

| コマンド                | タイトル                                         |
|-------------------------|--------------------------------------------------|
| dev util anonymise      | Anonymise capture log                            |
| job log jobid           | Retrieve logs of specified Job ID                |
| job log kind            | Concatenate and print logs of specified log kind |
| job log last            | Print the last job log files                     |
| member clear externalid | Clear external_id of members                     |


# コマンド仕様の変更: `config disable`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `config enable`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `config features`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `connect business_audit`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `connect business_file`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `connect business_info`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `connect business_mgmt`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `connect user_file`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev async`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev catalogue`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev ci artifact connect`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Full": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev ci artifact up`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth connect`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"Audit":  "business_audit",
  		"File":   "business_file",
  		"Full":   "user_full",
+ 		"Github": "github_repo",
  		"Info":   "business_info",
  		"Mgmt":   "business_management",
  	},
- 	Services:  nil,
+ 	Services:  []string{"dropbox", "dropbox_business", "github"},
  	IsSecret:  true,
  	IsConsole: false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth export`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Audit": "business_audit", "File": "business_file", "Full": "user_full", "Info": "business_info", "Mgmt": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox", "dropbox_business"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev ci auth import`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev doc`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Filename", Desc: "Filename", Default: "README.md", TypeName: "string"},
  		&{Name: "Lang", Desc: "Language", TypeName: "domain.common.model.mo_string.opt_string"},
- 		&{
- 			Name:     "MarkdownReadme",
- 			Desc:     "Generate README as markdown format",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  	},
  }
```
# コマンド仕様の変更: `dev dummy`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev echo`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	Values:          []*dc_recipe.Value{&{Name: "Text", Desc: "Text to echo", TypeName: "string"}},
  }
```
# コマンド仕様の変更: `dev kvs dump`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev preflight`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"ConnGithub": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev spec diff`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev spec doc`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev test monkey`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev test recipe`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `dev test resources`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	Values:          []*dc_recipe.Value{},
  }
```
# コマンド仕様の変更: `dev util curl`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	Values:          []*dc_recipe.Value{&{Name: "BufferSize", Desc: "Size of buffer", Default: "65536", TypeName: "domain.common.model.mo_int.range_int", TypeAttr: map[string]interface{}{"max": float64(2.097152e+06), "min": float64(1024), "value": float64(65536)}}, &{Name: "Record", Desc: "Capture record(s) for the test", TypeName: "string"}},
  }
```
# コマンド仕様の変更: `dev util wait`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        true,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file compare account`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Left": "user_full", "Right": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file compare local`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file copy`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file dispatch local`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file download`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file export doc`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file import batch url`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file import url`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file merge`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file move`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Dst": "user_full", "Src": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file restore`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file search content`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file search name`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file sync preflight up`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file sync up`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file upload`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `file watch`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete closed`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete url`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `filerequest list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group add`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group batch delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group member add`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group member delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group member list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `group rename`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `job history archive`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `job history delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `job history list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
- 	Values:          []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to workspace",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  }
```
# コマンド仕様の変更: `job history ship`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `job loop`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `job run`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       true,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `license`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	Values:          []*dc_recipe.Value{},
  }
```
# コマンド仕様の変更: `member delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member detach`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member invite`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "Include deleted members.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info"},
  	},
  }
```
# コマンド仕様の変更: `member quota list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member quota update`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member quota usage`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member reinvite`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Dst": "business_file", "Src": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member update email`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member update externalid`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `member update profile`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github profile`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github release asset download`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github release asset list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github release asset upload`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github release draft`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github release list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
- 	Services:        nil,
+ 	Services:        []string{"github"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `sharedlink create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `sharedlink delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `sharedlink file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `sharedlink list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{"Peer": "user_full"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team activity batch user`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team activity daily event`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team activity event`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team activity user`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_audit"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team content member`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "Filter by folder name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "Filter by folder name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "Filter by folder name. Filter by name match to the suffix.",
+ 		},
+ 		&{
+ 			Name: "MemberTypeExternal",
+ 			Desc: "Filter folder members. Keep only members are external (not in the same team). Note: Invited members are marked as external member.",
+ 		},
+ 		&{
+ 			Name: "MemberTypeInternal",
+ 			Desc: "Filter folder members. Keep only members are internal (in the same team). Note: Invited members are marked as external member.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
## 変更されたレポート: membership

```
  &dc_recipe.Report{
  	Name: "membership",
  	Desc: "This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "member_name", Desc: "Name of this member"},
  		&{Name: "member_email", Desc: "Email address of this member"},
+ 		&{
+ 			Name: "same_team",
+ 			Desc: "Whether the member is in the same team or not. Returns empty if the member is not able to determine whether in the same team or not.",
+ 		},
  	},
  }
```
# コマンド仕様の変更: `team content policy`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "Filter by folder name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "Filter by folder name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "Filter by folder name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
# コマンド仕様の変更: `team device list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team device unlink`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team diag explorer`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"File": "business_file", "Info": "business_info", "Mgmt": "business_management"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team feature`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team filerequest clone`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team filerequest list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team info`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_info"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team linkedapp list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team namespace file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team namespace file size`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team namespace list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team namespace member list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder archive`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch archive`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder batch replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file size`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder permdelete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      map[string]string{"Peer": "business_file"},
- 	Services:        nil,
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `teamfolder replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 6 identical fields
  }
```
# コマンド仕様の変更: `version`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      map[string]string{},
- 	Services:        nil,
+ 	Services:        []string{},
  	IsSecret:        false,
  	IsConsole:       false,
  	IsExperimental:  false,
  	IsIrreversible:  false,
- 	IsTransient:     false,
+ 	IsTransient:     true,
  	Reports:         nil,
  	Feeds:           nil,
  	Values:          []*dc_recipe.Value{},
  }
```
