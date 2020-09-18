# Changes between `Release 74` to `Release 75`

# Commands added


| Command                          | Title                        |
|----------------------------------|------------------------------|
| dev stage teamfolder             | Team folder operation sample |
| services slack conversation list | List channels                |



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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
