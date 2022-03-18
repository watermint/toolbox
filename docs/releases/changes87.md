---
layout: release
title: Changes of Release 86
lang: en
---

# Changes between `Release 86` to `Release 87`

# Commands added


| Command                           | Title                                 |
|-----------------------------------|---------------------------------------|
| dev test setup teamsharedlink     | Create demo shared links              |
| team sharedlink delete links      | Batch delete shared links             |
| team sharedlink delete member     | Delete all shared links of the member |
| team sharedlink update password   | Set or update shared link passwords   |
| team sharedlink update visibility | Update visibility of shared links     |



# Command spec changed: `dev stage upload_append`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "dropbox_scoped_individual"},
  	Services:       {"dropbox"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 9 identical fields
  }
```
# Command spec changed: `file export doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox document path to export.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
+ 		&{
+ 			Name:     "Format",
+ 			Desc:     "Export format",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "LocalPath", Desc: "Local path to save", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
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
  			Name: "Visibility",
  			Desc: strings.Join({
  				"Filter links by visibility (",
+ 				"all/",
  				"public/team_only/password)",
  			}, ""),
  			Default:  "all",
  			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]any{"options": []any{string("all"), string("public"), string("team_only"), string("password"), ...}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:            "expiry",
  	Title:           "Update expiration date of public shared links within the team",
- 	Desc:            "",
+ 	Desc:            "Note: From Release 87, this command will receive a file to select shared links to update. If you wanted to update the expiry for all shared links in the team, please consider using a combination of `team sharedlink list`. For example, if you are familiar w"...,
  	Remarks:         "(Irreversible operation)",
  	Path:            "team sharedlink update expiry",
- 	CliArgs:         "-days 28",
+ 	CliArgs:         "-file /PATH/TO/DATA_FILE.csv -days 28",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "New expiration date and time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Days", Desc: "Days to the new expiration date", Default: "0", TypeName: "essentials.model.mo_int.range_int", ...},
- 		&{
- 			Name:     "Peer",
- 			Desc:     "Account alias",
- 			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
- 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Path to data file",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
- 		&{
- 			Name:     "Visibility",
- 			Desc:     "Target link visibility",
- 			Default:  "public",
- 			TypeName: "essentials.model.mo_string.select_string",
- 			TypeAttr: map[string]any{
- 				"options": []any{
- 					string("public"), string("team_only"), string("password"),
- 					string("team_and_password"), ...,
- 				},
- 			},
- 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
+ 			TypeAttr: []any{string("members.read"), string("sharing.write"), string("team_data.member")},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Added report(s)


| Name          | Description                               |
|---------------|-------------------------------------------|
| operation_log | This report shows the transaction result. |



## Deleted report(s)


| Name    | Description                                                                      |
|---------|----------------------------------------------------------------------------------|
| skipped | This report shows a list of shared links with the shared link owner team member. |
| updated | This report shows the transaction result.                                        |


