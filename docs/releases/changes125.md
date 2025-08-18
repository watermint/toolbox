---
layout: release
title: Changes of Release 124
lang: en
---

# Changes between `Release 124` to `Release 125`

# Commands added


| Command              | Title                           |
|----------------------|---------------------------------|
| dev release asseturl | Update asset URL of the release |



# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BlockBlockSize", Desc: "Block size for batch upload", Default: "24", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name:    "Method",
  			Desc:    "Upload method",
  			Default: "block",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("block"), string("sequential")}},
  		},
  		&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev build package`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BuildPath", Desc: "Full path to the binary", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "DeployPath", Desc: "Deploy destination folder path (remote)", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "DistPath", Desc: "Package destination folder path (local)", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
+ 		&{
+ 			Name:     "ExecutableName",
+ 			Desc:     "Executable file name base",
+ 			Default:  "tbx",
+ 			TypeName: "string",
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release asset`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Owner", Desc: "Repository owner", TypeName: "string"},
  		&{Name: "Path", Desc: "Content path", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repo", Desc: "Repository name", TypeName: "string"},
  		&{Name: "Text", Desc: "Text content", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
+ 			TypeAttr: string("github_public"),
+ 		},
+ 	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
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
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev test setup teamsharedlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "Query", Desc: "Query", TypeName: "string"},
  		&{Name: "Seed", Desc: "Shared link seed value", Default: "0", TypeName: "int", ...},
  		&{
  			Name:    "Visibility",
  			Desc:    "Visibility",
  			Default: "random",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("random"), string("public"), string("team_only"), string("with_expire"), ...}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file paper append`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "Import format (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file paper create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "Import format (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file paper overwrite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "Import format (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file paper prepend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{
  			Name:    "Format",
  			Desc:    "Import format (html/markdown/plain_text)",
  			Default: "markdown",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("markdown"), string("plain_text"), string("html")}},
  		},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file search content`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Category",
  			Desc:    "Restricts search to only the file categories specified (image/do"...,
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "MaxResults", Desc: "Maximum number of entry to return", Default: "25", TypeName: "essentials.model.mo_int.range_int", ...},
  		... // 3 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file search name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Category",
  			Desc:    "Restricts search to only the file categories specified (image/do"...,
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ManagementType",
  			Desc:    "Group management type `company_managed` or `user_managed`",
  			Default: "company_managed",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("company_managed"), string("user_managed")}},
  		},
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "ManagementType",
  			Desc:    "Who is allowed to manage the group (user_managed, company_manage"...,
  			Default: "company_managed",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("company_managed"), string("user_managed")}},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "IncludeExternalGroups", Desc: "Include external groups in the report.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group update type`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "Type",
  			Desc:    "Group type (user_managed/company_managed)",
  			Default: "company_managed",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("user_managed"), string("company_managed")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `job log jobid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "Job ID", TypeName: "string"},
  		&{
  			Name:    "Kind",
  			Desc:    "Kind of log",
  			Default: "toolbox",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `job log kind`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Kind",
  			Desc:    "Log kind.",
  			Default: "toolbox",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `job log last`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Kind",
  			Desc:    "Log kind",
  			Default: "toolbox",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberEmail", Desc: "Filter by member email address. Filter by email address."},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services deepl translate text`



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
  			TypeName: "domain.deepl.api.deepl_conn_impl.conn_deepl_api_impl",
- 			TypeAttr: string(""),
+ 			TypeAttr: string("deepl"),
  		},
  		&{Name: "SourceLang", Desc: "Source language code (auto detect when omitted)", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "TargetLang", Desc: "Target language code", TypeName: "string"},
  		&{Name: "Text", Desc: "Text to translate", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services dropboxsign account info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AccountId", Desc: "Account ID", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropboxsign.api.hs_conn_impl.conn_hello_sign_api",
- 			TypeAttr: nil,
+ 			TypeAttr: string("dropbox_sign"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma account info`



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
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma file export all page`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "Export format (png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Path", Desc: "Output folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "Export scale in percent range from 1 to 400 (default 100)", Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "TeamId", Desc: "Team ID. To obtain a team id, navigate to a team page of a team "..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma file export frame`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "Export format (png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Key", Desc: "File key", TypeName: "string"},
  		&{Name: "Path", Desc: "Output folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "Export scale in percent range from 1 to 400 (default 100)", Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma file export node`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "Export format (png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Id", Desc: "Node ID", TypeName: "string"},
  		&{Name: "Key", Desc: "File Key", TypeName: "string"},
  		&{Name: "Path", Desc: "Output folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "Export scale in percent range from 1 to 400 (default 100)", Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma file export page`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "Export format (png/jpg/svg/pdf)",
  			Default: "pdf",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("jpg"), string("png"), string("svg"), string("pdf")}},
  		},
  		&{Name: "Key", Desc: "File key", TypeName: "string"},
  		&{Name: "Path", Desc: "Output folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "Scale", Desc: "Export scale in percent range from 1 to 400 (default 100)", Default: "100", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma file info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllNodes", Desc: "Include all node information", Default: "false", TypeName: "bool", ...},
  		&{Name: "Key", Desc: "File key", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma file list`



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
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "ProjectId", Desc: "Project ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services figma project list`



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
  			TypeName: "domain.figma.api.fg_conn_impl.conn_figma_api",
- 			TypeAttr: nil,
+ 			TypeAttr: []any{string("file_read")},
  		},
  		&{Name: "TeamId", Desc: "Team ID. To obtain a team id, navigate to a team page of a team "..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github content get`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to the content", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Ref", Desc: "Name of reference", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github content put`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to the content", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Filter",
  			Desc:    "Indicates which sorts of issues to return.",
  			Default: "assigned",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("assigned"), string("created"), string("mentioned"), string("subscribed"), ...}},
  		},
  		&{Name: "Labels", Desc: "A list of comma separated label names.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "Repository name", TypeName: "string"},
  		&{Name: "Since", Desc: "Only show notifications updated after the given time.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name:    "State",
  			Desc:    "Indicates the state of the issues to return.",
  			Default: "open",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("open"), string("closed"), string("all")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github profile`



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
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github release asset download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to download", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github release asset list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github release asset upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Asset", Desc: "Path to assets", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github release draft`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Name", Desc: "Name of the release", TypeName: "string"},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  		&{Name: "Tag", Desc: "Name of the tag", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github release list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Repository owner", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "Repository name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services github tag create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: nil,
+ 			TypeAttr: string("github_repo"),
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  		&{Name: "Sha1", Desc: "SHA1 hash of the commit", TypeName: "string"},
  		&{Name: "Tag", Desc: "Tag name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google mail label add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ColorBackground",
  			Desc:    "The background color.",
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:    "ColorText",
  			Desc:    "The text color.",
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:    "LabelListVisibility",
  			Desc:    "The visibility of the label in the label list in the Gmail web i"...,
  			Default: "labelShow",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("labelHide"), string("labelShow"), string("labelShowIfUnread")}},
  		},
  		&{
  			Name:    "MessageListVisibility",
  			Desc:    "The visibility of messages with this label in the message list i"...,
  			Default: "show",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("hide"), string("show")}},
  		},
  		&{Name: "Name", Desc: "Name of the label", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google mail message list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "The format to return the message in. ",
  			Default: "metadata",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Labels", Desc: "Only return messages with labels that match all of the specified"..., TypeName: "essentials.model.mo_string.opt_string"},
  		... // 4 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google mail message processed list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "The format to return the message in. ",
  			Default: "metadata",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Labels", Desc: "Only return messages with labels that match all of the specified"..., TypeName: "essentials.model.mo_string.opt_string"},
  		... // 4 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets sheet export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Data", Desc: "Path to export.", TypeName: "Data"},
  		&{Name: "DataFormat", Desc: "Output format"},
  		&{
  			Name:    "DateTimeRender",
  			Desc:    "How dates, times, and durations should be represented in the out"...,
  			Default: "serial",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("serial"), string("formatted")}},
  		},
  		&{Name: "Id", Desc: "Spreadsheet ID", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_sheets", ...},
  		&{Name: "Range", Desc: "The range the values cover, in A1 notation. This is a string lik"..., TypeName: "string"},
  		&{
  			Name:    "ValueRender",
  			Desc:    "How values should be represented in the output.",
  			Default: "formatted",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("formatted"), string("formatted"), string("formula")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "Exported sheet data"}},
  	... // 2 identical fields
  }
```
# Command spec changed: `services slack conversation history`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "After", Desc: "Retrieve messages after this.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Channel", Desc: "Channel ID (like C1234567890).", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "Peer",
+ 			Default:  "default",
  			TypeName: "domain.slack.api.work_conn_impl.conn_slack_api",
  			TypeAttr: []any{string("channels:history")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services slack conversation list`



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
- 			Default:  "Peer",
+ 			Default:  "default",
  			TypeName: "domain.slack.api.work_conn_impl.conn_slack_api",
  			TypeAttr: []any{string("channels:read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "AccessLevel",
  			Desc:    "Access type (viewer/editor)",
  			Default: "editor",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("editor"), string("viewer"), string("viewer_no_comment")}},
  		},
  		&{Name: "Email", Desc: "Email address of the folder member", TypeName: "string"},
  		&{Name: "Message", Desc: "Custom message for invitation", TypeName: "essentials.model.mo_string.opt_string"},
  		... // 3 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "AclUpdatePolicy",
  			Desc:    "Who can change a shared folder's access control list (ACL).",
  			Default: "owner",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("owner"), string("editor")}},
  		},
  		&{
  			Name:    "MemberPolicy",
  			Desc:    "Who can be a member of this shared folder.",
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("team"), string("anyone")}},
  		},
  		&{Name: "Path", Desc: "Path to be shared", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:    "SharedLinkPolicy",
  			Desc:    "Who can view shared links in this folder.",
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("anyone"), string("members")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content legacypaper export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "FilterBy",
  			Desc:    "Specify how the Paper docs should be filtered (doc_created/doc_a"...,
  			Default: "docs_created",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("docs_created"), string("docs_accessed")}},
  		},
  		&{
  			Name:    "Format",
  			Desc:    "Export file format (html/markdown)",
  			Default: "html",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("html"), string("markdown")}},
  		},
  		&{Name: "Path", Desc: "Export folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content legacypaper list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "FilterBy",
  			Desc:    "Specify how the Paper docs should be filtered (doc_created/doc_a"...,
  			Default: "docs_created",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("docs_created"), string("docs_accessed")}},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content member size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSubFolders", Desc: "Include sub-folders to the report.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder batch share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "AclUpdatePolicy",
  			Desc:    "Who can add and remove members of this shared folder.",
  			Default: "owner",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("owner"), string("editor")}},
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "MemberPolicy",
  			Desc:    "Who can be a member of this shared folder.",
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("team"), string("anyone")}},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "SharedLinkPolicy",
  			Desc:    "The policy to apply to shared links created for content inside t"...,
  			Default: "anyone",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("anyone"), string("members")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink cap visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "NewVisibility",
  			Desc:    "New visibility setting",
  			Default: "team_only",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("team_only")}},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "Visibility",
  			Desc:    "Filter links by visibility (all/public/team_only/password)",
  			Default: "all",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("all"), string("public"), string("team_only"), string("password"), ...}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink update visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:    "NewVisibility",
  			Desc:    "New visibility setting",
  			Default: "team_only",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("public"), string("team_only")}},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "SyncSetting",
  			Desc:    "Sync setting for the team folder",
  			Default: "default",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("default"), string("not_synced")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamspace asadmin member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name:    "ScanTimeout",
  			Desc:    "Scan timeout mode. If the scan timeouts, the path of a subfolder"...,
  			Default: "short",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("short"), string("long")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util file hash`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Algorithm",
  			Desc:    "Hash algorithm (md5/sha1/sha256)",
  			Default: "sha1",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("md5"), string("sha1"), string("sha256")}},
  		},
  		&{Name: "File", Desc: "Path to the file to create digest", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util image placeholder`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "Path", Desc: "Path to export generated image", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Text", Desc: "Text if you need", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:    "TextAlign",
  			Desc:    "Text alignment",
  			Default: "left",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("left"), string("center"), string("right")}},
  		},
  		&{Name: "TextColor", Desc: "Text color", Default: "black", TypeName: "string", ...},
  		&{Name: "TextPosition", Desc: "Text position", Default: "center", TypeName: "string", ...},
  		&{Name: "Width", Desc: "Width (pixel)", Default: "640", TypeName: "int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util qrcode create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ErrorCorrectionLevel",
  			Desc:    "Error correction level (l/m/q/h).",
  			Default: "m",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("l"), string("m"), string("q"), string("h")}},
  		},
  		&{
  			Name:    "Mode",
  			Desc:    "QR code encoding mode",
  			Default: "auto",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("auto"), string("numeric"), string("alpha_numeric"), string("unicode")}},
  		},
  		&{Name: "Out", Desc: "Output path with file name", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Size", Desc: "Image resolution (pixel)", Default: "256", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Text", Desc: "Text data", TypeName: "Text"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util qrcode wifi`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "ErrorCorrectionLevel",
  			Desc:    "Error correction level (l/m/q/h).",
  			Default: "m",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("l"), string("m"), string("q"), string("h")}},
  		},
  		&{
  			Name:    "Hidden",
  			Desc:    "`true` if a SSID is hidden. `false` if a SSID is visible.",
  			Default: "",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string(""), string("true"), string("false")}},
  		},
  		&{
  			Name:    "Mode",
  			Desc:    "QR code encoding mode",
  			Default: "auto",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("auto"), string("numeric"), string("alpha_numeric"), string("unicode")}},
  		},
  		&{
  			Name:    "NetworkType",
  			Desc:    "Network type.",
  			Default: "WPA",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("WPA"), string("WEP"), string("")}},
  		},
  		&{Name: "Out", Desc: "Output path with file name", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Size", Desc: "Image resolution (pixel)", Default: "256", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Ssid", Desc: "Network SSID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util release install`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AcceptLicenseAgreement", Desc: "Accept to the target release's license agreement", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "Path to install", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
+ 			TypeAttr: string("github_public"),
+ 		},
  		&{Name: "Release", Desc: "Release tag name", Default: "latest", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util text nlp japanese token`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Dictionary",
  			Desc:    "Dictionary name of the token",
  			Default: "ipa",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("ipa"), string("uni")}},
  		},
  		&{Name: "IgnoreLineBreak", Desc: "Ignore line break", Default: "false", TypeName: "bool", ...},
  		&{Name: "In", Desc: "Input file path", TypeName: "In"},
  		&{
  			Name:    "Mode",
  			Desc:    "Tokenize mode (normal/search/extended)",
  			Default: "normal",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("normal"), string("search"), string("extend")}},
  		},
  		&{Name: "OmitBosEos", Desc: "Omit BOS/EOS tokens", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util text nlp japanese wakati`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Dictionary",
  			Desc:    "Dictionary name (ipa/uni)",
  			Default: "ipa",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("ipa"), string("uni")}},
  		},
  		&{Name: "IgnoreLineBreak", Desc: "Ignore line break", Default: "false", TypeName: "bool", ...},
  		&{Name: "In", Desc: "Input file path", TypeName: "In"},
  		&{Name: "Separator", Desc: "Text separator", Default: " ", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util unixtime format`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Format",
  			Desc:    "Time format",
  			Default: "iso8601",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("iso8601"), string("rfc1123"), string("rfc1123z"), string("rfc3339"), ...}},
  		},
  		&{
  			Name:    "Precision",
  			Desc:    "Time precision (second/ms/ns)",
  			Default: "second",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("second"), string("ms"), string("ns")}},
  		},
  		&{Name: "Time", Desc: "Unix Time", Default: "0", TypeName: "int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util unixtime now`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:    "Precision",
  			Desc:    "Time precision (second/ms/ns)",
  			Default: "second",
  			TypeName: strings.Join({
  				"essentials.model.mo_string.select_string",
+ 				"_internal",
  			}, ""),
  			TypeAttr: map[string]any{"options": []any{string("second"), string("ms"), string("ns")}},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
