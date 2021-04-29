# Changes between `Release 89` to `Release 90`

# Command spec changed: `dev build doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Badge", Desc: "Include badges of build status", Default: "true", TypeName: "bool", ...},
- 		&{
- 			Name:     "CommandPath",
- 			Desc:     "Relative path to generate command manuals",
- 			Default:  "doc/generated/",
- 			TypeName: "string",
- 		},
  		&{Name: "DocLang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"},
- 		&{
- 			Name:     "Readme",
- 			Desc:     "Filename of README",
- 			Default:  "README.md",
- 			TypeName: "string",
- 		},
- 		&{
- 			Name:     "Security",
- 			Desc:     "Filename of SECURITY_AND_PRIVACY",
- 			Default:  "SECURITY_AND_PRIVACY.md",
- 			TypeName: "string",
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
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
  	IsSecret: true,
  	... // 11 identical fields
  }
```
