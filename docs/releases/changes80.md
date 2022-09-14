---
layout: release
title: Changes of Release 79
lang: en
---

# Changes between `Release 79` to `Release 80`

# Commands added


| Command                   | Title                                                  |
|---------------------------|--------------------------------------------------------|
| member folder replication | Replicate a folder to another member's personal folder |
| member update invisible   | Enable directory restriction to members                |
| member update visible     | Disable directory restriction to members               |
| teamfolder member add     | Batch adding users/groups to team folders              |
| teamfolder member delete  | Batch removing users/groups from team folders          |



# Command spec changed: `dev stage scoped`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Individual",
  			Desc:     "Account alias for individual",
- 			Default:  "&{Individual [files.content.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{string("files.content.read")},
  		},
  		&{
  			Name:     "Team",
  			Desc:     "Account alias for team",
- 			Default:  "&{Team [members.read team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("members.read"), string("team_info.read")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev stage teamfolder`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("files.content.read"), string("files.content.write"), string("groups.write"), string("sharing.read"), ...},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `group member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [groups.read groups.write] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("groups.read"), string("groups.write")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `group member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [groups.read groups.write] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("groups.read"), string("groups.write")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `group member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [groups.read groups.write] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("groups.read"), string("groups.write")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `member file permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Team member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [files.permanent_delete team_data.member members.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{string("files.permanent_delete"), string("team_data.member"), string("members.read")},
  		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team report activity`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `team report devices`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `team report membership`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `team report storage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
