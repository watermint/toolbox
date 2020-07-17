# Changes between `Release 71` to `Release 72`

# Commands added


| Command                           | Title                         |
|-----------------------------------|-------------------------------|
| dev stage gmail                   | Gmail command                 |
| dev stage scoped                  | Dropbox scoped OAuth app test |
| services google mail label list   | List email labels             |
| services google mail message list | List messages                 |



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
+ 		&{
+ 			Name:     "DocLang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
  		&{Name: "Filename", Desc: "Filename", Default: "README.md", TypeName: "string", ...},
- 		&{
- 			Name:     "Lang",
- 			Desc:     "Language",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
