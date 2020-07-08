# `リリース 71` から `リリース 72` までの変更点

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
