# `リリース 89` から `リリース 90` までの変更点

# コマンド仕様の変更: `dev build doc`



## 設定が変更されたコマンド


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
# コマンド仕様の変更: `member file permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "permdelete",
  	Title: strings.Join({
  		"Permanently delete the file or folder at a given path of the tea",
  		"m member",
- 		". Please see https://www.dropbox.com/help/40 for more detail abo",
- 		"ut permanent deletion",
  		".",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Please see https://www.dropbox.com/help/40 for more detail about permanent deletion.",
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "member file permdelete",
  	... // 18 identical fields
  }
```
