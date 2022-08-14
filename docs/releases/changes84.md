---
layout: release
title: Changes of Release 83
lang: en
---

# Changes between `Release 83` to `Release 84`

# Commands added


| Command                          | Title                                               |
|----------------------------------|-----------------------------------------------------|
| file lock acquire                | Lock a file                                         |
| file lock all release            | Release all locks under the specified path          |
| file lock batch acquire          | Lock multiple files                                 |
| file lock batch release          | Release multiple locks                              |
| file lock list                   | List locks under the specified path                 |
| file lock release                | Release a lock                                      |
| member file lock all release     | Release all locks under the path of the member      |
| member file lock list            | List locks of the member under the path             |
| member file lock release         | Release the lock of the path as the member          |
| teamfolder file lock all release | Release all locks under the path of the team folder |
| teamfolder file lock list        | List locks in the team folder                       |
| teamfolder file lock release     | Release lock of the path in the team folder         |



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
- 		"Peer": "business_management",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
