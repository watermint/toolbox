# Changes between `Release 88` to `Release 89`

# Command spec changed: `dev spec diff`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
- 			Name:     "FilePath",
+ 			Name:     "DocLang",
- 			Desc:     "File path to output",
+ 			Desc:     "Document language",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
- 			Name:     "Lang",
+ 			Name:     "FilePath",
- 			Desc:     "Language",
+ 			Desc:     "File path to output",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Release1", Desc: "Release name 1", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Release2", Desc: "Release name 2", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  }
```
