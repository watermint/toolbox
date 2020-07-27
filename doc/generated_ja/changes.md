# `リリース 72` から `リリース 73` までの変更点

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
  		&{Name: "DocLang", Desc: "Language", TypeName: "domain.common.model.mo_string.opt_string"},
  		&{
- 			Name:     "Filename",
+ 			Name:     "Readme",
- 			Desc:     "Filename",
+ 			Desc:     "Filename of README",
  			Default:  "README.md",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
+ 		&{
+ 			Name:     "Security",
+ 			Desc:     "Filename of SECURITY_AND_PRIVACY",
+ 			Default:  "SECURITY_AND_PRIVACY.md",
+ 			TypeName: "string",
+ 		},
  	},
  }
```
