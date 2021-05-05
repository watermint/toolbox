---
layout: release
title: Changes of Release 78
lang: en
---

# Changes between `Release 78` to `Release 79`

# Commands added


| Command                   | Title                             |
|---------------------------|-----------------------------------|
| dev stage gui             | GUI proof of concept              |
| file archive local        | Archive local files               |
| group member batch add    | Bulk add members into groups      |
| group member batch delete | Delete members from groups        |
| group member batch update | Add or delete members from groups |
| team report activity      | Activities report                 |
| team report devices       | Devices report                    |
| team report membership    | Membership report                 |
| team report storage       | Storage report                    |



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
  	... // 11 identical fields
  }
```
# Command spec changed: `team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "Filter links by visibility (public/team_only/password)",
- 			Default:  "public",
+ 			Default:  "all",
  			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{
  				"options": []interface{}{
+ 					string("all"),
  					string("public"),
  					string("team_only"),
  					... // 3 identical elements
  				},
  			},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
