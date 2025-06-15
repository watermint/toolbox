---
layout: release
title: Changes of Release 139
lang: en
---

# Changes between `Release 139` to `Release 140`

# Command spec changed: `dev release publish`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ArtifactPath", Desc: "Path to artifacts", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Branch", Desc: "Target branch", Default: "main", TypeName: "string", ...},
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
+ 		&{
+ 			Name:     "ExecutableName",
+ 			Desc:     "The name of the executable file to be published.",
+ 			Default:  "tbx",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "HomebrewRepoBranch",
+ 			Desc:     "The branch of the Homebrew tap repository to use for publishing.",
+ 			Default:  "master",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "HomebrewRepoName",
+ 			Desc:     "The name of the Homebrew tap repository.",
+ 			Default:  "homebrew-toolbox",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "HomebrewRepoOwner",
+ 			Desc:     "The owner of the Homebrew tap repository.",
+ 			Default:  "watermint",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepoName",
+ 			Desc:     "The name of the repository to publish the release to.",
+ 			Default:  "toolbox",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepoOwner",
+ 			Desc:     "The owner of the repository to publish the release to.",
+ 			Default:  "watermint",
+ 			TypeName: "string",
+ 		},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file tag add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Add a tag to the file/folder",
+ 	Title:   "Add tag to file or folder",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BasePath",
- 			Desc:     "Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a per"...,
+ 			Desc:     "Base path for adding a tag.",
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "Path",
- 			Desc:     "Target path",
+ 			Desc:     "File or folder path to add a tag.",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:     "Tag",
- 			Desc:     "Tag name",
+ 			Desc:     "Tag to add to the file or folder.",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file tag delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dropbox file tag delete",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/TO/",
- 		"TARGET",
+ 		"PROCESS",
  		" -tag TAG_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 9 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BasePath",
- 			Desc:     "Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a per"...,
+ 			Desc:     "Base path for removing a tag.",
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "Path",
- 			Desc:     "Target path",
+ 			Desc:     "File or folder path to remove a tag.",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Tag", Desc: "Tag name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team activity batch user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "File", Desc: "User email address list file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("events.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team activity daily event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Event category", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndDate", Desc: "End date", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("events.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("events.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team activity user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("events.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin group role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin group role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ExceptionGroup", Desc: "Exception group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeNonAdmin", Desc: "Include non admin members in the report", Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberRoles", Desc: "Member to admin-role mappings", TypeName: "MemberRoles"},
  		&{Name: "MemberRolesFormat", Desc: "Output format"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "MemberRoles", Desc: "Member to admin-role mappings"}},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin role clear`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team admin role list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team backup device status`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndTime", Desc: "End date/time of the period to retrieve data for (exclusive). If"..., TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("events.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "StartTime", Desc: "Start date/time of the period to retrieve data for (inclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper count`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Format", Desc: "Export file format (html/markdown)", Default: "html", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Path", Desc: "Export folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "FilterBy", Desc: "Specify how the Paper docs should be filtered (doc_created/doc_a"..., Default: "docs_created", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...},
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content member size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "IncludeSubFolders", Desc: "Include sub-folders to the report.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberNamePrefix", Desc: "Filter members. Filter by name match to the prefix."},
  		&{Name: "MemberNameSuffix", Desc: "Filter members. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team content policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."},
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team device list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sessions.list"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team device unlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DeleteOnUnlink", Desc: "Delete files on unlink", Default: "false", TypeName: "bool", ...},
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("sessions.modify"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: feature

```
  &dc_recipe.Report{
  	Name: "feature",
- 	Desc: "This report shows a list of team features and their settings.",
+ 	Desc: "Team feature",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "upload_api_rate_limit", Desc: "The number of upload API calls allowed per month."},
  		&{Name: "upload_api_rate_limit_count", Desc: "The number of upload API called this month."},
  		&{Name: "has_team_shared_dropbox", Desc: "Does this team have a shared team root."},
  		&{
  			Name: "has_team_file_events",
- 			Desc: "Does this team have file events.",
+ 			Desc: "Team supports file events",
  		},
  		&{
  			Name: "has_team_selective_sync",
- 			Desc: "Does this team have team selective sync enabled.",
+ 			Desc: "Team supports selective sync",
  		},
  		&{
  			Name: "has_distinct_member_homes",
  			Desc: strings.Join({
- 				"Does this team have team member folder.",
+ 				"Team has distinct member home folders",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("file_requests.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team filesystem`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ManagementType", Desc: "Group management type `company_managed` or `user_managed`", Default: "company_managed", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "ManagementType", Desc: "Who is allowed to manage the group (user_managed, company_manage"..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file for group name list", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "GroupNameSuffix", Desc: "Filter by group name. Filter by name match to the suffix."},
  		&{Name: "IncludeExternalGroups", Desc: "Include external groups in the report.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "GroupName", Desc: "Group name", TypeName: "string"},
  		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "GroupName", Desc: "Name of the group", TypeName: "string"},
  		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "CurrentName", Desc: "Current group name", TypeName: "string"},
  		&{Name: "NewName", Desc: "New group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team group update type`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("groups.read"),
  				string("groups.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "Type", Desc: "Group type (user_managed/company_managed)", Default: "company_managed", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team legalhold revision list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List revisions of the legal hold policy",
+ 	Title:   "List revisions under legal hold",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "After", Desc: "Get revisions after this specified date and time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:     "PolicyId",
- 			Desc:     "Legal hold policy ID",
+ 			Desc:     "Legal hold policy ID.",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team legalhold update desc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dropbox team legalhold update desc",
  	CliArgs: strings.Join({
  		"-desc ",
- 		"NEW_",
  		"DESCRIPTION -policy-id POLICY_ID",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 16 identical fields
  }
```
# Command spec changed: `dropbox team linkedapp list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sessions.list"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.delete"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "TransferDestMember", Desc: "If provided, files from the deleted member account will be trans"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "TransferNotifyAdminEmailOnError", Desc: "If provided, errors during the transfer process will be sent via"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "WipeData", Desc: "If true, controls if the user's data will be deleted on their li"..., Default: "true", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member batch detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.delete"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "RevokeTeamShares", Desc: "True for revoke shared folder access which owned by the team", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member batch invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "SilentInvite", Desc: "Do not send welcome email (requires SSO + domain verification in"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member batch reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.delete"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "Silent", Desc: "Do not send welcome email (SSO required)", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member batch suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "KeepData", Desc: "Keep the user's data on their linked devices", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member batch unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member file permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "Team member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.permanent_delete"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "MemberEmail", Desc: "Filter by member email address. Filter by email address."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member folder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "DstMemberEmail", Desc: "Destination team member email address", TypeName: "string"},
  		&{Name: "DstPath", Desc: "The path for the destination team member. Note the root (/) path"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "SrcMemberEmail", Desc: "Source team member email address", TypeName: "string"},
  		&{Name: "SrcPath", Desc: "The path of the source team member", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "Include deleted members.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_data.member"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member quota batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "Quota", Desc: "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom q"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member quota list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Member's email address", TypeName: "string"},
  		&{Name: "KeepData", Desc: "Keep the user's data on their linked devices", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Member's email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member update batch email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "UpdateUnverified", Desc: "Update an account which didn't verified email. If an account ema"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member update batch externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member update batch invisible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member update batch profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "profile",
- 	Title:   "Update member profile",
+ 	Title:   "Batch update member profiles",
  	Desc:    "",
  	Remarks: "(Irreversible operation)",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team member update batch visible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("members.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team namespace file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 5 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "Namespace size",
+ 	Desc: "Namespace size in bytes",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "count_descendant", Desc: "Number of files and folders under the folder"},
  		&{Name: "size", Desc: "Size of the folder"},
  		&{
  			Name: "depth",
- 			Desc: "Folder depth",
+ 			Desc: "Namespace depth",
  		},
  		&{
  			Name: "mod_time_earliest",
  			Desc: strings.Join({
- 				"The e",
+ 				"E",
  				"arliest modification time ",
- 				"of a file in this folder or child folders.",
+ 				"in namespace",
  			}, ""),
  		},
  		&{
  			Name: "mod_time_latest",
  			Desc: strings.Join({
- 				"The l",
+ 				"L",
  				"atest modification time ",
- 				"of a f",
  				"i",
- 				"le i",
  				"n ",
- 				"this folder or child folders",
+ 				"namespace",
  			}, ""),
  		},
  		&{Name: "api_complexity", Desc: "Folder complexity index for API operations"},
  	},
  }
```
# Command spec changed: `dropbox team namespace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_data.member"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team namespace member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllColumns", Desc: "Show all columns", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("sharing.read"),
  				string("team_data.member"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team namespace summary`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "SkipMemberSummary", Desc: "Skip scanning member namespaces", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team report activity`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team report devices`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team report membership`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team report storage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas file batch copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas file sync batch up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "up",
  	Title: strings.Join({
  		"Batch ",
- 		"sync up that run as members",
+ 		"upstream sync with Dropbox",
  	}, ""),
  	Desc:    "",
  	Remarks: "(Irreversible operation)",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "ExitOnFailure", Desc: "Exit the program on failure", Default: "false", TypeName: "bool", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name: "NameDisableIgnore",
  			Desc: strings.Join({
- 				"Filter by name",
+ 				"Name for the sync batch operation",
  				". Filter system file or ignore files.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "NameName",
  			Desc: strings.Join({
- 				"Filter by name",
+ 				"Name for the sync batch operation",
  				". Filter by exact match to the name.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "NameNamePrefix",
  			Desc: strings.Join({
- 				"Filter by name",
+ 				"Name for the sync batch operation",
  				". Filter by name match to the prefix.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "NameNameSuffix",
  			Desc: strings.Join({
- 				"Filter by name",
+ 				"Name for the sync batch operation",
  				". Filter by name match to the suffix.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "Overwrite",
  			Desc: strings.Join({
  				"Overwrite existing file",
- 				" on the target path if that exists",
+ 				"s if they exist",
  				".",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{
  			Name: "input.member_email",
- 			Desc: "The email address of the member",
+ 			Desc: "Email address of the Dropbox team member.",
  		},
  		&{
  			Name: "input.local_path",
- 			Desc: "Local file path",
+ 			Desc: "Local file path to upload.",
  		},
  		&{
  			Name: "input.dropbox_path",
- 			Desc: "Destination Dropbox path",
+ 			Desc: "Destination path in Dropbox.",
  		},
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "leave",
- 	Title:   "Batch leave from shared folders as a member",
+ 	Title:   "Leave shared folders in batch",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to leave",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name: "KeepCopy",
  			Desc: strings.Join({
  				"Keep a copy of the folder",
- 				"'s contents upon relinquishing membership",
+ 				" after leaving",
  				".",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("members.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{
  			Name: "input.member_email",
- 			Desc: "Member email address",
+ 			Desc: "Email address of the member.",
  		},
  		&{
  			Name: "input.path",
- 			Desc: "Path to share",
+ 			Desc: "Path to the member's folder.",
  		},
  		&{Name: "result.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "result.parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		... // 13 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "share",
- 	Title:   "Batch share folders for members",
+ 	Title:   "Share shared folders in batch",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "AclUpdatePolicy",
- 			Desc:     "Who can add and remove members of this shared folder.",
+ 			Desc:     "Access control update policy.",
  			Default:  "owner",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("owner"), string("editor")}},
  		},
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to share",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name: "MemberPolicy",
  			Desc: strings.Join({
- 				"Who can be a member of this shared folder",
+ 				"Policy for shared folder members",
  				".",
  			}, ""),
  			Default:  "anyone",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("team"), string("anyone")}},
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("members.read"),
  				... // 4 identical elements
  			},
  		},
  		&{
  			Name:     "SharedLinkPolicy",
- 			Desc:     "The policy to apply to shared links created for content inside this shared folder.",
+ 			Desc:     "Policy for shared links.",
  			Default:  "anyone",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("anyone"), string("members")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{
  			Name: "input.member_email",
- 			Desc: "Member email address",
+ 			Desc: "Email address of the member.",
  		},
  		&{
  			Name: "input.path",
- 			Desc: "Path to share",
+ 			Desc: "Path to the member's folder.",
  		},
  		&{Name: "result.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "result.parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		... // 13 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "unshare",
- 	Title:   "Batch unshare folders for members",
+ 	Title:   "Unshare shared folders in batch",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to unshare",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "LeaveCopy",
- 			Desc:     "If true, members of this shared folder will get a copy of this folder after it's unshared. ",
+ 			Desc:     "Leave a copy after unsharing.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("members.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{
  			Name: "input.member_email",
- 			Desc: "Member email address",
+ 			Desc: "Email address of the member.",
  		},
  		&{
  			Name: "input.path",
- 			Desc: "Path to share",
+ 			Desc: "Path to the member's folder.",
  		},
  		&{Name: "result.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "result.parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		... // 13 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder isolate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "isolate",
  	Title: strings.Join({
- 		"Unshare owned shared folders and leave from external shared fold",
- 		"ers run as a memb",
+ 		"Isolate member from shared fold",
  		"er",
  	}, ""),
  	Desc:    "",
  	Remarks: "(Irreversible operation)",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to isolate",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "KeepCopy",
- 			Desc:     "Keep a copy of the folder's contents upon relinquishing membership.",
+ 			Desc:     "Keep a copy after isolation.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "MemberEmail",
- 			Desc:     "Member email address",
+ 			Desc:     "Email address of the member to isolate.",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List shared folders run as the member",
+ 	Title:   "List shared folders",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to list",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "MemberEmail",
- 			Desc:     "Member email address",
+ 			Desc:     "Email address of the member to list.",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
- 		"Batch add members to member's shared folders",
+ 		"Add members to shared folders in batch",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to add members",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Message",
- 			Desc:     "Custom message for invitation",
+ 			Desc:     "Message to send to new members.",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{
  			Name:     "Silent",
- 			Desc:     "Do not send invitation email",
+ 			Desc:     "Add members silently without notification.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{
  			Name: "input.member_email",
- 			Desc: "Team member email address",
+ 			Desc: "Email address of the member to add.",
  		},
  		&{
  			Name: "input.path",
- 			Desc: "Shared folder path of the member",
+ 			Desc: "Path to the shared folder.",
  		},
  		&{
  			Name: "input.access_level",
- 			Desc: "Access type (viewer/editor)",
+ 			Desc: "Access level to grant to the member.",
  		},
  		&{
  			Name: "input.group_or_email",
  			Desc: strings.Join({
  				"Group name or ",
- 				"member email address",
+ 				"email address to add.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"Batch delete members from member's shared folders",
+ 		"Remove members from shared folders in batch",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to remove members",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name: "LeaveCopy",
  			Desc: strings.Join({
- 				"If true, members of this shared folder will get a copy of this f",
- 				"older after it's unshared",
+ 				"Leave a copy after removing member",
  				".",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{
  			Name: "input.member_email",
- 			Desc: "Team member email address",
+ 			Desc: "Email address of the member to remove.",
  		},
  		&{
  			Name: "input.path",
- 			Desc: "Shared folder path of the member",
+ 			Desc: "Path to the shared folder.",
  		},
  		&{
  			Name: "input.group_or_email",
  			Desc: strings.Join({
  				"Group name or ",
- 				"member email address",
+ 				"email address to remove.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
- 		"Add the shared folder to the specified member's Dropbox",
+ 		"Mount a shared folder as another member",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to mount",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "MemberEmail",
- 			Desc:     "Member email address",
+ 			Desc:     "Email address of the member",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 4 identical elements
  			},
  		},
  		&{
  			Name:     "SharedFolderId",
- 			Desc:     "The ID for the shared folder.",
+ 			Desc:     "Shared folder ID",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder",
+ 				"Base path of the shared folder to unmount",
  				".",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "MemberEmail",
- 			Desc:     "Member email address",
+ 			Desc:     "Email address of the member",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "The ID for the shared folder.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "IncludeMounted", Desc: "Include mounted folders.", Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink cap expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "New expiry date/time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink cap visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "New visibility setting", Default: "team_only", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink delete links`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink delete member`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "Visibility", Desc: "Filter links by visibility (all/public/team_only/password)", Default: "all", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "New expiration date and time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink update password`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink update visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "New visibility setting", Default: "team_only", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("members.read"),
  				string("sharing.write"),
  				... // 2 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  		&{Name: "SyncSetting", Desc: "Sync setting for the team folder", Default: "default", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...},
  		&{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("team_data.member"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.write"),
  				string("files.metadata.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("members.read"),
  				... // 3 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "Namespace size",
+ 	Desc: "Namespace size in bytes",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "count_descendant", Desc: "Number of files and folders under the folder"},
  		&{Name: "size", Desc: "Size of the folder"},
  		&{
  			Name: "depth",
- 			Desc: "Folder depth",
+ 			Desc: "Namespace depth",
  		},
  		&{
  			Name: "mod_time_earliest",
  			Desc: strings.Join({
- 				"The e",
+ 				"E",
  				"arliest modification time ",
- 				"of a file in this folder or child folders.",
+ 				"in namespace",
  			}, ""),
  		},
  		&{
  			Name: "mod_time_latest",
  			Desc: strings.Join({
- 				"The l",
+ 				"L",
  				"atest modification time ",
- 				"of a f",
  				"i",
- 				"le i",
  				"n ",
- 				"this folder or child folders",
+ 				"namespace",
  			}, ""),
  		},
  		&{Name: "api_complexity", Desc: "Folder complexity index for API operations"},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 7 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
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
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 7 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...},
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 5 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder partial replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BasePath",
  			Desc: strings.Join({
- 				"Choose the file path standard. This is an option for Dropbox for",
- 				" Teams in particular. If you are using the personal version of D",
- 				"ropbox, it basically doesn't matter what you choose. In Dropbox ",
- 				"for Teams, if you select `home` in the updated team space, a per",
- 				"sonal folder with your username will be selected. This is conven",
- 				"ient for referencing or uploading files in your personal folder,",
- 				" as you don't need to include the folder name with your username",
- 				" in the path. On the other hand, if you specify `root`, you can ",
- 				"access all folders with permissions. On the other hand, when acc",
- 				"essing your personal folder, you need to specify a path that inc",
- 				"ludes the name of your personal folder.",
+ 				"Base path for partial replication",
  			}, ""),
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{
  			Name:     "Dst",
- 			Desc:     "Peer name for the destination team",
+ 			Desc:     "Destination account alias",
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("account_info.write"),
  				string("files.content.read"),
  				... // 4 identical elements
  			},
  		},
  		&{
  			Name:     "DstPath",
- 			Desc:     "Relative path from the team folder (please specify '/' for the team folder root)",
+ 			Desc:     "Destination path",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "DstTeamFolderName", Desc: "Destination team folder name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "src",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "SrcPath", Desc: "Relative path from the team folder (please specify '/' for the t"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "SrcTeamFolderName", Desc: "Source team folder name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."},
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("groups.read"),
  				... // 4 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder sync setting list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("team_data.content.read"),
  				... // 3 identical elements
  			},
  		},
  		&{Name: "ScanAll", Desc: "Perform a scan for all depths (can take considerable time depend"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "ShowAll", Desc: "Show all scanned folders", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder sync setting update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("team_data.content.read"),
  				... // 4 identical elements
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github release draft`



## Changed report: release

```
  &dc_recipe.Report{
  	Name: "release",
  	Desc: "Release on GitHub",
  	Columns: []*dc_recipe.ReportColumn{
  		&{
  			Name: "id",
- 			Desc: "Id of the release",
+ 			Desc: "Release ID",
  		},
  		&{
  			Name: "tag_name",
- 			Desc: "Tag name",
+ 			Desc: "Release tag name",
  		},
  		&{
  			Name: "name",
- 			Desc: "Name of the release",
+ 			Desc: "Release name",
  		},
  		&{
  			Name: "draft",
- 			Desc: "True when the release is draft.",
+ 			Desc: "Release is a draft",
  		},
  		&{Name: "url", Desc: "URL of the release"},
  	},
  }
```
# Command spec changed: `github release list`



## Changed report: releases

```
  &dc_recipe.Report{
  	Name: "releases",
  	Desc: "Release on GitHub",
  	Columns: []*dc_recipe.ReportColumn{
  		&{
  			Name: "tag_name",
- 			Desc: "Tag name",
+ 			Desc: "Release tag name",
  		},
  		&{
  			Name: "name",
- 			Desc: "Name of the release",
+ 			Desc: "Release name",
  		},
  		&{
  			Name: "draft",
- 			Desc: "True when the release is draft.",
+ 			Desc: "Release is a draft",
  		},
  		&{Name: "url", Desc: "URL of the release"},
  	},
  }
```
# Command spec changed: `util file hash`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "hash",
- 	Title:   "Print file digest",
+ 	Title:   "File Hash",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `util git clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util git clone",
  	CliArgs: strings.Join({
  		"-local-path /LOCAL/PATH/TO/",
- 		"clone -url https://git.repository.url",
+ 		"CLONE -url https://github.com/username/repository.git",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 9 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "LocalPath",
- 			Desc:     "Local path to clone",
+ 			Desc:     "Local path to clone repository",
  			Default:  "",
  			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]any{"shouldExist": bool(false)},
  		},
  		&{Name: "Reference", Desc: "Reference name", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "RemoteName", Desc: "Name of the remote", Default: "origin", TypeName: "string", ...},
  		&{Name: "Url", Desc: "Git repository url", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util uuid timestamp`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "timestamp",
- 	Title:   "Parse UUID timestamp",
+ 	Title:   "UUID Timestamp",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `util uuid ulid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "ulid",
- 	Title:   "Generate ULID (Universally Unique Lexicographically Sortable Identifier)",
+ 	Title:   "ULID Utility",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `util uuid v4`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "UpperCase",
- 			Desc:     "Use upper case characters",
+ 			Desc:     "Output UUID in upper case",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
