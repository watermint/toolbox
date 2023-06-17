---
layout: release
title: Changes of Release 89
lang: en
---

# Changes between `Release 89` to `Release 90`

# Command spec changed: `dev build doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
# Command spec changed: `member file permdelete`



## Command configuration changed


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
  	... // 19 identical fields
  }
```
