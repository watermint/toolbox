# Changes between `Release 71` to `Release 72`

# Commands added


| Command                         | Title                         |
|---------------------------------|-------------------------------|
| dev stage gmail                 | Gmail command                 |
| dev stage scoped                | Dropbox scoped OAuth app test |
| services google mail label list | List email labels             |



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
