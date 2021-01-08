# Changes between `Release 79` to `Release 80`

# Commands added

| Command                  | Title                                       |
|--------------------------|---------------------------------------------|
| member update invisible  | Enable directory restriction to members     |
| member update visible    | Disable directory restriction to members    |
| teamfolder member add    | Batch adding users/groups to team folders   |
| teamfolder member delete | Batch removing users/groups to team folders |

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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
