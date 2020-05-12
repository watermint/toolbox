# Changes between `Release 66` to `Release 67`

# Commands added

| Command       | Title                                            |
|---------------|--------------------------------------------------|
| job log jobid | Retrieve logs of specified Job ID                |
| job log kind  | Concatenate and print logs of specified log kind |
| job log last  | Print the last job log files                     |


# Command spec changed: `job history list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to workspace",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  }
```
