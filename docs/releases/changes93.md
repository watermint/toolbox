---
layout: release
title: Changes of Release 92
lang: en
---

# Changes between `Release 92` to `Release 93`

# Commands added


| Command                  | Title                                                      |
|--------------------------|------------------------------------------------------------|
| file restore all         | Restore files under given path                             |
| team content member size | Count number of members of team folders and shared folders |



# Commands deleted


| Command       | Title                          |
|---------------|--------------------------------|
| file download | Download a file from Dropbox   |
| file restore  | Restore files under given path |



# Command spec changed: `file compare account`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Left":  "user_full",
+ 		"Left":  "dropbox_scoped_individual",
- 		"Right": "user_full",
+ 		"Right": "dropbox_scoped_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Left",
  			Desc:     "Account alias (left)",
  			Default:  "left",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "LeftPath", Desc: "The path from account root (left)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Right",
  			Desc:     "Account alias (right)",
  			Default:  "right",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "RightPath", Desc: "The path from account root (right)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file compare local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file export doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox document path to export.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Format", Desc: "Export format", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "Local path to save", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file import batch url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Path", Desc: "Path to import", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file import url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to import", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.write")},
  		},
  		&{Name: "Url", Desc: "URL", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file merge`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool", ...},
  		&{Name: "From", Desc: "Path for merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "KeepEmptyFolder", Desc: "Keep empty folder after merge", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "To", Desc: "Path to merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "WithinSameNamespace", Desc: "Do not cross namespace. That is for preserve sharing permission "..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("sharing.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: mount

```
  &dc_recipe.Report{
  	Name: "mount",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 10 identical elements
  		&{Name: "policy_viewer_info", Desc: "Who can enable/disable viewer info for this shared folder."},
  		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
+ 		&{Name: "access_inheritance", Desc: "Access inheritance type"},
  	},
  }
```
# Command spec changed: `file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Dst": "user_full",
+ 		"Dst": "dropbox_scoped_individual",
- 		"Src": "user_full",
+ 		"Src": "dropbox_scoped_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Dst",
  			Desc:     "Account alias (destionation)",
  			Default:  "dst",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.metadata.read"), string("files.content.write")},
  		},
  		&{Name: "DstPath", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Src",
  			Desc:     "Account alias (source)",
  			Default:  "src",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "SrcPath", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Depth", Desc: "Report an entry for all files and folders depth folders deep", Default: "2", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Path", Desc: "Path to scan", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file sync down`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 5 identical elements
  		&{Name: "NameNamePrefix", Desc: "Filter by name. Filter by name match to the prefix."},
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "SkipExisting", Desc: "Skip existing files. Do not overwrite", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file sync online`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "NameNamePrefix", Desc: "Filter by name. Filter by name match to the prefix."},
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  		&{Name: "SkipExisting", Desc: "Skip existing files. Do not overwrite", Default: "false", TypeName: "bool", ...},
  		&{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 7 identical elements
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{Name: "Overwrite", Desc: "Overwrite existing file on the target path if that exists.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read"), string("files.content.write")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file watch`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to watch", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.read")},
  		},
  		&{Name: "Recursive", Desc: "Watch path recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder list`



## Changed report: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 8 identical elements
  		&{Name: "policy_viewer_info", Desc: "Who can enable/disable viewer info for this shared folder."},
  		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
+ 		&{Name: "access_inheritance", Desc: "Access inheritance type"},
  	},
  }
```
# Command spec changed: `team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "event",
  	Title:   "Event log",
- 	Desc:    "",
+ 	Desc:    "From release 91, the command parses `-start-time` or `-end-time` as the relative duration from now with the format like \"-24h\" (24 hours) or \"-10m\" (10 minutes).\nIf you wanted to retrieve events every hour, then run like:\n\n```\ntbx team activity event -star"...,
  	Remarks: "",
  	Path:    "team activity event",
  	... // 18 identical fields
  }
```
# Command spec changed: `team content member list`



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
  		... // 3 identical elements
  		&{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...},
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{
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
