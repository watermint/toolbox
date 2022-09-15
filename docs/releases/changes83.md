---
layout: release
title: Changes of Release 82
lang: en
---

# Changes between `Release 82` to `Release 83`

# Commands added


| Command                        | Title                                                        |
|--------------------------------|--------------------------------------------------------------|
| dev benchmark uploadlink       | Benchmark single file upload with upload temporary link API. |
| file info                      | Resolve metadata of the path                                 |
| teamfolder partial replication | Partial team folder replication to the other team            |



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
- 		"Peer": "business_file",
+ 		"Peer": "business_management",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
