---
layout: release
title: Changes of Release 94
lang: en
---

# Changes between `Release 94` to `Release 95`

# Commands added


| Command                | Title                  |
|------------------------|------------------------|
| member batch suspend   | Bulk suspend members   |
| member batch unsuspend | Bulk unsuspend members |
| member suspend         | Suspend a member       |
| member unsuspend       | Unsuspend a member     |



# Command spec changed: `dev test setup teamsharedlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "Group name", TypeName: "string"},
  		&{Name: "NumLinksPerMember", Desc: "Number of links to create per member", Default: "5", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.content.write"),
- 				string("sharing.write"),
+ 				string("groups.read"),
- 				string("groups.read"),
+ 				string("sharing.write"),
  				string("team_data.member"),
  			},
  		},
  		&{Name: "Query", Desc: "Query", TypeName: "string"},
  		&{Name: "Seed", Desc: "Shared link seed value", Default: "0", TypeName: "int", ...},
  		&{Name: "Visibility", Desc: "Visibility", Default: "random", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file export url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Format", Desc: "Export format", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "Local path to export", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Password", Desc: "Password for the shared link", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
- 				string("sharing.read"),
+ 				string("files.content.read"),
- 				string("files.content.read"),
+ 				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "URL of the document", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
- 				string("files.metadata.read"),
+ 				string("files.content.write"),
- 				string("files.content.write"),
+ 				string("files.metadata.read"),
  			},
  		},
  		&{Name: "DstPath", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Src", Desc: "Account alias (source)", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "SrcPath", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member file permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Team member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.permanent_delete"),
- 				string("team_data.member"),
+ 				string("members.read"),
- 				string("members.read"),
+ 				string("team_data.member"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content mount list`



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
  		... // 2 identical elements
  		&{Name: "MemberNamePrefix", Desc: "Filter members. Filter by name match to the prefix."},
  		&{Name: "MemberNameSuffix", Desc: "Filter members. Filter by name match to the suffix."},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("members.read"), string("sharing.read"), string("team_data.member"),
+ 				string("team_data.team_space"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content policy list`



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
  		&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."},
  		&{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."},
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("groups.read"), string("members.read"),
+ 				string("sharing.read"), string("team_data.member"),
+ 				string("team_data.team_space"), string("team_info.read"),
+ 			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team device list`



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
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("members.read"), string("sessions.list")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team device unlink`



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
  		&{Name: "DeleteOnUnlink", Desc: "Delete files on unlink", Default: "false", TypeName: "bool", ...},
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("sessions.modify")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team namespace file list`



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
  		... // 5 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team namespace file size`



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
  		... // 6 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder file list`



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
  		&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...},
  		&{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder file size`



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
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{
  			Name:    "Peer",
  			Desc:    "Account alias",
  			Default: "default",
  			TypeName: strings.Join({
  				"domain.dropbox.api.dbx_conn_impl.conn_",
- 				"business_file",
+ 				"scoped_team",
  			}, ""),
- 			TypeAttr: nil,
+ 			TypeAttr: []any{
+ 				string("files.metadata.read"), string("members.read"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AdminGroupName", Desc: "Temporary group name for admin operation", Default: "watermint-toolbox-admin", TypeName: "string", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("groups.read"),
- 				string("groups.write"),
  				string("files.content.read"),
  				string("files.content.write"),
+ 				string("groups.read"),
+ 				string("groups.write"),
  				string("sharing.read"),
  				string("sharing.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AdminGroupName", Desc: "Temporary group name for admin operation", Default: "watermint-toolbox-admin", TypeName: "string", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("groups.read"),
- 				string("groups.write"),
  				string("files.content.read"),
  				string("files.content.write"),
+ 				string("groups.read"),
+ 				string("groups.write"),
  				string("sharing.read"),
  				string("sharing.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
