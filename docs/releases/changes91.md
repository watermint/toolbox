---
layout: release
title: Changes of Release 90
lang: en
---

# Changes between `Release 90` to `Release 91`

# Commands added


| Command         | Title                      |
|-----------------|----------------------------|
| dev release doc | Generate release documents |



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
  	IsSecret: true,
  	... // 11 identical fields
  }
```
