# Changes between `Release 70` to `Release 71`

# Commands added


| Command             | Title                                 |
|---------------------|---------------------------------------|
| dev diag endpoint   | List endpoints                        |
| dev diag throughput | Evaluate throughput from capture logs |



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
