# Changes between `Release 72` to `Release 73`

# Commands added


| Command                               | Title                          |
|---------------------------------------|--------------------------------|
| services asana team list              | List team                      |
| services asana team project list      | List projects of the team      |
| services asana team task list         | List task of the team          |
| services asana workspace list         | List workspaces                |
| services asana workspace project list | List projects of the workspace |



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
# Command spec changed: `services google mail filter batch add`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...},
  		&{
- 			Name:     "ApplyToInboxMessages",
+ 			Name:     "ApplyToExistingMessages",
- 			Desc:     "Apply labels to messages satisfy query in INBOX.",
+ 			Desc:     "Apply labels to existing messages that satisfy the query.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  }
```
