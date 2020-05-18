# Changes between `Release 66` to `Release 67`

# Commands added

| Command                 | Title                                            |
|-------------------------|--------------------------------------------------|
| dev util anonymise      | Anonymise capture log                            |
| job log jobid           | Retrieve logs of specified Job ID                |
| job log kind            | Concatenate and print logs of specified log kind |
| job log last            | Print the last job log files                     |
| member clear externalid | Clear external_id of members                     |


# Command spec changed: `config disable`


## Command configuration changed


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
# Command spec changed: `config enable`


## Command configuration changed


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
# Command spec changed: `config features`


## Command configuration changed


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
# Command spec changed: `connect business_audit`


## Command configuration changed


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
# Command spec changed: `connect business_file`


## Command configuration changed


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
# Command spec changed: `connect business_info`


## Command configuration changed


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
# Command spec changed: `connect business_mgmt`


## Command configuration changed


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
# Command spec changed: `connect user_file`


## Command configuration changed


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
# Command spec changed: `dev async`


## Command configuration changed


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
# Command spec changed: `dev catalogue`


## Command configuration changed


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
# Command spec changed: `dev ci artifact connect`


## Command configuration changed


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
# Command spec changed: `dev ci artifact up`


## Command configuration changed


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
# Command spec changed: `dev ci auth connect`


## Command configuration changed


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
# Command spec changed: `dev ci auth export`


## Command configuration changed


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
# Command spec changed: `dev ci auth import`


## Command configuration changed


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
# Command spec changed: `dev doc`


## Command configuration changed


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
# Command spec changed: `dev dummy`


## Command configuration changed


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
# Command spec changed: `dev echo`


## Command configuration changed


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
# Command spec changed: `dev kvs dump`


## Command configuration changed


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
# Command spec changed: `dev preflight`


## Command configuration changed


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
# Command spec changed: `dev release candidate`


## Command configuration changed


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
# Command spec changed: `dev release publish`


## Command configuration changed


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
# Command spec changed: `dev spec diff`


## Command configuration changed


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
# Command spec changed: `dev spec doc`


## Command configuration changed


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
# Command spec changed: `dev test monkey`


## Command configuration changed


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
# Command spec changed: `dev test recipe`


## Command configuration changed


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
# Command spec changed: `dev test resources`


## Command configuration changed


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
# Command spec changed: `dev util curl`


## Command configuration changed


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
# Command spec changed: `dev util wait`


## Command configuration changed


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
# Command spec changed: `file compare account`


## Command configuration changed


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
# Command spec changed: `file compare local`


## Command configuration changed


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
# Command spec changed: `file copy`


## Command configuration changed


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
# Command spec changed: `file delete`


## Command configuration changed


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
# Command spec changed: `file dispatch local`


## Command configuration changed


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
# Command spec changed: `file download`


## Command configuration changed


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
# Command spec changed: `file export doc`


## Command configuration changed


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
# Command spec changed: `file import batch url`


## Command configuration changed


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
# Command spec changed: `file import url`


## Command configuration changed


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
# Command spec changed: `file list`


## Command configuration changed


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
# Command spec changed: `file merge`


## Command configuration changed


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
# Command spec changed: `file move`


## Command configuration changed


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
# Command spec changed: `file replication`


## Command configuration changed


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
# Command spec changed: `file restore`


## Command configuration changed


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
# Command spec changed: `file search content`


## Command configuration changed


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
# Command spec changed: `file search name`


## Command configuration changed


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
# Command spec changed: `file sync preflight up`


## Command configuration changed


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
# Command spec changed: `file sync up`


## Command configuration changed


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
# Command spec changed: `file upload`


## Command configuration changed


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
# Command spec changed: `file watch`


## Command configuration changed


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
# Command spec changed: `filerequest create`


## Command configuration changed


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
# Command spec changed: `filerequest delete closed`


## Command configuration changed


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
# Command spec changed: `filerequest delete url`


## Command configuration changed


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
# Command spec changed: `filerequest list`


## Command configuration changed


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
# Command spec changed: `group add`


## Command configuration changed


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
# Command spec changed: `group batch delete`


## Command configuration changed


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
# Command spec changed: `group delete`


## Command configuration changed


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
# Command spec changed: `group list`


## Command configuration changed


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
# Command spec changed: `group member add`


## Command configuration changed


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
# Command spec changed: `group member delete`


## Command configuration changed


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
# Command spec changed: `group member list`


## Command configuration changed


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
# Command spec changed: `group rename`


## Command configuration changed


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
# Command spec changed: `job history archive`


## Command configuration changed


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
# Command spec changed: `job history delete`


## Command configuration changed


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
# Command spec changed: `job history list`


## Command configuration changed


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
# Command spec changed: `job history ship`


## Command configuration changed


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
# Command spec changed: `job loop`


## Command configuration changed


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
# Command spec changed: `job run`


## Command configuration changed


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
# Command spec changed: `license`


## Command configuration changed


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
# Command spec changed: `member delete`


## Command configuration changed


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
# Command spec changed: `member detach`


## Command configuration changed


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
# Command spec changed: `member invite`


## Command configuration changed


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
# Command spec changed: `member list`


## Command configuration changed


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
# Command spec changed: `member quota list`


## Command configuration changed


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
# Command spec changed: `member quota update`


## Command configuration changed


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
# Command spec changed: `member quota usage`


## Command configuration changed


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
# Command spec changed: `member reinvite`


## Command configuration changed


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
# Command spec changed: `member replication`


## Command configuration changed


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
# Command spec changed: `member update email`


## Command configuration changed


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
# Command spec changed: `member update externalid`


## Command configuration changed


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
# Command spec changed: `member update profile`


## Command configuration changed


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
# Command spec changed: `services github issue list`


## Command configuration changed


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
# Command spec changed: `services github profile`


## Command configuration changed


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
# Command spec changed: `services github release asset download`


## Command configuration changed


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
# Command spec changed: `services github release asset list`


## Command configuration changed


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
# Command spec changed: `services github release asset upload`


## Command configuration changed


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
# Command spec changed: `services github release draft`


## Command configuration changed


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
# Command spec changed: `services github release list`


## Command configuration changed


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
# Command spec changed: `services github tag create`


## Command configuration changed


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
# Command spec changed: `sharedfolder list`


## Command configuration changed


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
# Command spec changed: `sharedfolder member list`


## Command configuration changed


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
# Command spec changed: `sharedlink create`


## Command configuration changed


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
# Command spec changed: `sharedlink delete`


## Command configuration changed


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
# Command spec changed: `sharedlink file list`


## Command configuration changed


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
# Command spec changed: `sharedlink list`


## Command configuration changed


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
# Command spec changed: `team activity batch user`


## Command configuration changed


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
# Command spec changed: `team activity daily event`


## Command configuration changed


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
# Command spec changed: `team activity event`


## Command configuration changed


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
# Command spec changed: `team activity user`


## Command configuration changed


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
# Command spec changed: `team content member`


## Command configuration changed


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
## Changed report: membership

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
# Command spec changed: `team content policy`


## Command configuration changed


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
# Command spec changed: `team device list`


## Command configuration changed


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
# Command spec changed: `team device unlink`


## Command configuration changed


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
# Command spec changed: `team diag explorer`


## Command configuration changed


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
# Command spec changed: `team feature`


## Command configuration changed


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
# Command spec changed: `team filerequest clone`


## Command configuration changed


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
# Command spec changed: `team filerequest list`


## Command configuration changed


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
# Command spec changed: `team info`


## Command configuration changed


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
# Command spec changed: `team linkedapp list`


## Command configuration changed


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
# Command spec changed: `team namespace file list`


## Command configuration changed


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
# Command spec changed: `team namespace file size`


## Command configuration changed


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
# Command spec changed: `team namespace list`


## Command configuration changed


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
# Command spec changed: `team namespace member list`


## Command configuration changed


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
# Command spec changed: `team sharedlink list`


## Command configuration changed


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
# Command spec changed: `team sharedlink update expiry`


## Command configuration changed


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
# Command spec changed: `teamfolder archive`


## Command configuration changed


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
# Command spec changed: `teamfolder batch archive`


## Command configuration changed


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
# Command spec changed: `teamfolder batch permdelete`


## Command configuration changed


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
# Command spec changed: `teamfolder batch replication`


## Command configuration changed


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
# Command spec changed: `teamfolder file list`


## Command configuration changed


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
# Command spec changed: `teamfolder file size`


## Command configuration changed


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
# Command spec changed: `teamfolder list`


## Command configuration changed


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
# Command spec changed: `teamfolder permdelete`


## Command configuration changed


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
# Command spec changed: `teamfolder replication`


## Command configuration changed


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
# Command spec changed: `version`


## Command configuration changed


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
