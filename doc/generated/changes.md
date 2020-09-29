# Changes between `Release 75` to `Release 76`

# Commands added


| Command            | Title                             |
|--------------------|-----------------------------------|
| dev replay approve | Approve the replay as test bundle |
| dev replay bundle  | Run all replays                   |
| dev replay recipe  | Replay recipe                     |
| dev replay remote  | Run remote replay bundle          |



# Commands deleted


| Command         | Title         |
|-----------------|---------------|
| dev test replay | Replay recipe |



# Command spec changed: `dev ci artifact up`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path to upload", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string", ...},
  		&{
  			Name:     "Timeout",
  			Desc:     "Operation timeout in seconds",
- 			Default:  "30",
+ 			Default:  "60",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `job history archive`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Days", Desc: "Target days old", Default: "7", TypeName: "essentials.model.mo_int.range_int", ...},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to the workspace",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  	},
  }
```
# Command spec changed: `job history delete`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Days", Desc: "Target days old", Default: "28", TypeName: "essentials.model.mo_int.range_int", ...},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to the workspace",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  	},
  }
```
