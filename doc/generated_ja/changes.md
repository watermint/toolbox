# `リリース 79` から `リリース 80` までの変更点

# 追加されたコマンド

| コマンド                 | タイトル                                      |
|--------------------------|-----------------------------------------------|
| member update invisible  | Enable directory restriction to members       |
| member update visible    | Disable directory restriction to members      |
| teamfolder member add    | Batch adding users/groups to team folders     |
| teamfolder member delete | Batch removing users/groups from team folders |

# コマンド仕様の変更: `dev stage scoped`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Individual",
  			Desc:     "Account alias for individual",
- 			Default:  "&{Individual [files.content.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{
  			Name:     "Team",
  			Desc:     "Account alias for team",
- 			Default:  "&{Team [members.read team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("members.read"), string("team_info.read")},
  		},
  	},
  }
```

# コマンド仕様の変更: `dev stage teamfolder`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write"), string("groups.write"), string("sharing.read"), ...},
  		},
  	},
  }
```

# コマンド仕様の変更: `group member batch add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
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
  			TypeAttr: []interface{}{string("groups.read"), string("groups.write")},
  		},
  	},
  }
```

# コマンド仕様の変更: `group member batch delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
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
  			TypeAttr: []interface{}{string("groups.read"), string("groups.write")},
  		},
  	},
  }
```

# コマンド仕様の変更: `group member batch update`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
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
  			TypeAttr: []interface{}{string("groups.read"), string("groups.write")},
  		},
  	},
  }
```

# コマンド仕様の変更: `member file permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
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
  			TypeAttr: []interface{}{string("files.permanent_delete"), string("team_data.member"), string("members.read")},
  		},
  	},
  }
```

# コマンド仕様の変更: `team diag explorer`

## 設定が変更されたコマンド

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
  	... // 7 identical fields
  }
```

# コマンド仕様の変更: `team report activity`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  }
```

# コマンド仕様の変更: `team report devices`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  }
```

# コマンド仕様の変更: `team report membership`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  }
```

# コマンド仕様の変更: `team report storage`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{Peer [team_info.read] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("team_info.read")},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  }
```
