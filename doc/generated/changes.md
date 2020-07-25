# Changes between `Release 71` to `Release 72`

# Commands added


| Command                                     | Title                              |
|---------------------------------------------|------------------------------------|
| dev stage gmail                             | Gmail command                      |
| dev stage scoped                            | Dropbox scoped OAuth app test      |
| services google mail filter add             | Add a filter.                      |
| services google mail filter delete          | Delete a filter                    |
| services google mail filter list            | List filters                       |
| services google mail label add              | Add a label                        |
| services google mail label delete           | Delete a label                     |
| services google mail label list             | List email labels                  |
| services google mail label rename           | Rename a label                     |
| services google mail message label add      | Add labels to the message          |
| services google mail message label delete   | Remove labels from the message     |
| services google mail message list           | List messages                      |
| services google mail message processed list | List messages in processed format. |
| services google mail thread list            | List threads                       |



# Command spec changed: `dev doc`


## Command configuration changed


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
# Command spec changed: `dev util curl`


## Command configuration changed


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
# Command spec changed: `team diag explorer`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"File": "business_file",
  		"Info": "business_info",
  		"Mgmt": "business_management",
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
