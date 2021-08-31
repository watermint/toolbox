---
layout: release
title: Changes of Release 95
lang: en
---

# Changes between `Release 95` to `Release 96`

# Commands added


| Command                         | Title                                                     |
|---------------------------------|-----------------------------------------------------------|
| member feature                  | List member feature settings                              |
| services dropbox user feature   | List feature settings for current user                    |
| team content legacypaper count  | Count number of Paper documents per member                |
| team content legacypaper export | Export entire team member Paper documents into local path |
| team content legacypaper list   | List team member Paper documents                          |



# Command spec changed: `member replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "business_file",
+ 		"Dst": "dropbox_scoped_team",
- 		"Src": "business_file",
+ 		"Src": "dropbox_scoped_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Dst",
  			Desc:     "Destination team; team file access",
  			Default:  "dst",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{},
  		},
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Src",
  			Desc:     "Source team; team file access",
  			Default:  "src",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github issue list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Filter",
+ 			Desc:     "Indicates which sorts of issues to return.",
+ 			Default:  "assigned",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{
+ 				"options": []interface{}{
+ 					string("assigned"), string("created"), string("mentioned"),
+ 					string("subscribed"), ...,
+ 				},
+ 			},
+ 		},
+ 		&{
+ 			Name:     "Labels",
+ 			Desc:     "A list of comma separated label names.",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "Repository", Desc: "Repository name", TypeName: "string"},
+ 		&{
+ 			Name:     "Since",
+ 			Desc:     "Only show notifications updated after the given time.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "State",
+ 			Desc:     "Indicates the state of the issues to return.",
+ 			Default:  "open",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("open"), string("closed"), string("all")}},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team filerequest clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("file_requests.read"), string("members.read"), string("team_data.member")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
