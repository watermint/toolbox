---
layout: release
title: Changes of Release 109
lang: en
---

# Changes between `Release 109` to `Release 110`

# Commands added


| Command                         | Title                                            |
|---------------------------------|--------------------------------------------------|
| dev ci auth export              | Export deploy token data for CI build            |
| services hellosign account info | Retrieve account information                     |
| util release install            | Download & install watermint toolbox to the path |



# Commands deleted


| Command             | Title                                                           |
|---------------------|-----------------------------------------------------------------|
| dev ci auth connect | Authenticate for generating end to end testing                  |
| dev ci auth import  | Import auth tokens of end to end test from environment variable |



# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "PreScan", Desc: "Pre-scan destination path", Default: "false", TypeName: "bool", ...},
  		&{Name: "SeqChunkSizeKb", Desc: "Upload chunk size in KiB", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...},
  		... // 3 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev benchmark uploadlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "SizeKb", Desc: "Size in KB", Default: "1024", TypeName: "int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Github": "github_repo",
  		"Peer":   "github_public",
  	},
  	Services: {"github"},
  	IsSecret: true,
  	... // 12 identical fields
  }
```
# Command spec changed: `dev stage dbxfs`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to scan", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev stage encoding`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Encoding", Desc: "Encoding", TypeName: "string"},
  		&{Name: "Name", Desc: "Name of the file", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev stage http_range`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox file path to download", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path to store", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev stage scoped`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Individual": "dropbox_scoped_individual",
+ 		"Individual": "dropbox_individual",
- 		"Team":       "dropbox_scoped_team",
+ 		"Team":       "dropbox_team",
  	},
  	Services: {"dropbox", "dropbox_business"},
  	IsSecret: true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Team", Desc: "Account alias for team", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev stage teamfolder`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `dev stage upload_append`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Upload path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev test auth all`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `dev test setup massfiles`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Base", Desc: "Dropbox base path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "BatchSize", Desc: "Batch size", Default: "1000", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Offset", Desc: "Upload offset (skip # pages)", Default: "0", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Source", Desc: "Source file", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 5 identical fields
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
  				... // 2 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
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
# Command spec changed: `file compare account`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Left":  "dropbox_scoped_individual",
+ 		"Left":  "dropbox_individual",
- 		"Right": "dropbox_scoped_individual",
+ 		"Right": "dropbox_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "left",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "LeftPath", Desc: "The path from account root (left)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "right",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox document path to export.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Format", Desc: "Export format", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "LocalPath", Desc: "Local path to save", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
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
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "URL of the document", TypeName: "domain.dropbox.model.mo_url.url_impl"},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Path", Desc: "Path to import", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to import", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Url", Desc: "URL", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.metadata.read"),
  			},
  		},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "IncludeMountedFolders", Desc: " If true, the results will include entries under mounted folders"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file lock acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "File path to lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Path to release locks", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file lock batch acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file lock batch release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.metadata.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to the file", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool", ...},
  		&{Name: "From", Desc: "Path for merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "KeepEmptyFolder", Desc: "Keep empty folder after merge", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "To", Desc: "Path to merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "WithinSameNamespace", Desc: "Do not cross namespace. That is for preserve sharing permission "..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{Name: "Format", Desc: "Import format (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{Name: "Format", Desc: "Import format (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{Name: "Format", Desc: "Import format (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Content", Desc: "Paper content", TypeName: "Content"},
  		&{Name: "Format", Desc: "Import format (html/markdown/plain_text)", Default: "markdown", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "Path in the user's Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
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
- 		"Dst": "dropbox_scoped_individual",
+ 		"Dst": "dropbox_individual",
- 		"Src": "dropbox_scoped_individual",
+ 		"Src": "dropbox_individual",
  	},
  	Services: {"dropbox"},
  	IsSecret: false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  				string("files.metadata.read"),
  			},
  		},
  		&{Name: "DstPath", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "src",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "SrcPath", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file restore all`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file revision download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "LocalPath", Desc: "Local path to store downloaded file", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  			},
  		},
  		&{Name: "Revision", Desc: "File revision", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file revision list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "File path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.metadata.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file revision restore`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "File path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "Revision", Desc: "File revision", TypeName: "string"},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "MaxResults", Desc: "Maximum number of entry to return", Default: "25", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `file share info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "File", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  			},
  		},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Depth", Desc: "Report an entry for all files and folders depth folders deep", Default: "2", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Path", Desc: "Path to scan", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 5 identical elements
  		&{Name: "NameNamePrefix", Desc: "Filter by name. Filter by name match to the prefix."},
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "NameNamePrefix", Desc: "Filter by name. Filter by name match to the prefix."},
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 7 identical elements
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{Name: "Overwrite", Desc: "Overwrite existing file on the target path if that exists.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "Path to watch", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  			},
  		},
  		&{Name: "Recursive", Desc: "Watch path recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `filerequest create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllowLateUploads", Desc: "If set, allow uploads after the deadline has passed (one_day/two"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Deadline", Desc: "The deadline for this file request.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Path", Desc: "The path for the folder in the Dropbox where uploaded files will"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.write"),
  			},
  		},
  		&{Name: "Title", Desc: "The title of the file request", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `filerequest delete closed`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `filerequest delete url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Force", Desc: "Force delete the file request.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.read"),
  				string("file_requests.write"),
  			},
  		},
  		&{Name: "Url", Desc: "URL of the file request.", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("file_requests.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ManagementType", Desc: "Group management type `company_managed` or `user_managed`", Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "ManagementType", Desc: "Who is allowed to manage the group (user_managed, company_manage"..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file for group name list", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `group list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("groups.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `job history ship`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("files.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member batch suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member batch unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.delete"),
  				string("members.write"),
+ 				string("team_info.read"),
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
# Command spec changed: `member detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.delete"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RevokeTeamShares", Desc: "True for revoke shared folder access which owned by the team", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("account_info.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BatchSize", Desc: "Batch operation size", Default: "100", TypeName: "int", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.content.write"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.content.write"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `member folder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DstMemberEmail", Desc: "Destination team member email address", TypeName: "string"},
  		&{Name: "DstPath", Desc: "The path for the destination team member. Note the root (/) path"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
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
# Command spec changed: `member invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SilentInvite", Desc: "Do not send welcome email (requires SSO + domain verification in"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "Include deleted members.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member quota list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member quota update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Quota", Desc: "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom q"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member quota usage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `member reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.delete"),
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Silent", Desc: "Do not send welcome email (SSO required)", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# Command spec changed: `member suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Member's email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member update email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "UpdateUnverified", Desc: "Update an account which didn't verified email. If an account ema"..., Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member update externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member update invisible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member update profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `member update visible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services asana team list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "WorkspaceName", Desc: "Name or GID of the workspace. Filter by exact match to the name."},
  		&{Name: "WorkspaceNamePrefix", Desc: "Name or GID of the workspace. Filter by name match to the prefix."},
  		&{Name: "WorkspaceNameSuffix", Desc: "Name or GID of the workspace. Filter by name match to the suffix."},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services asana team project list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "TeamName", Desc: "Name or GID of the team Filter by exact match to the name."},
  		&{Name: "TeamNamePrefix", Desc: "Name or GID of the team Filter by name match to the prefix."},
  		... // 4 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services asana team task list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "ProjectName", Desc: "Name or GID of the project Filter by exact match to the name."},
  		&{Name: "ProjectNamePrefix", Desc: "Name or GID of the project Filter by name match to the prefix."},
  		... // 7 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services asana workspace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services asana workspace project list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "asana"},
  	Services:       {"asana"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default <nil> default}",
+ 			Default:  "default",
  			TypeName: "domain.asana.api.as_conn_impl.conn_asana_api",
  			TypeAttr: []any{string("default")},
  		},
  		&{Name: "WorkspaceName", Desc: "Name or GID of the workspace. Filter by exact match to the name."},
  		&{Name: "WorkspaceNamePrefix", Desc: "Name or GID of the workspace. Filter by name match to the prefix."},
  		&{Name: "WorkspaceNameSuffix", Desc: "Name or GID of the workspace. Filter by name match to the suffix."},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services dropbox user feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `services google sheets sheet append`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Data", Desc: "Data file path", TypeName: "Data"},
  		&{Name: "Id", Desc: "Spreadsheet Id", TypeName: "string"},
  		&{Name: "InputRaw", Desc: "Raw input", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Range", Desc: "The range the values cover, in A1 notation. This is a string lik"..., TypeName: "string"},
  	},
  	GridDataInput:  {&{Name: "Data", Desc: "Input data file"}},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets sheet clear`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "Spreadsheet ID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Range", Desc: "The range the values cover, in A1 notation. This is a string lik"..., TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets sheet create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Cols", Desc: "Number of columns", Default: "26", TypeName: "int", ...},
  		&{Name: "Id", Desc: "Spreadsheet ID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Rows", Desc: "Number of rows", Default: "1000", TypeName: "int", ...},
  		&{Name: "Title", Desc: "Sheet title", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets sheet delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "Spreadsheet Id", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "SheetId", Desc: "Sheet ID (Please use `services google sheets sheet list` to see "..., TypeName: "string"},
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
  		... // 2 identical elements
  		&{Name: "DateTimeRender", Desc: "How dates, times, and durations should be represented in the out"..., Default: "serial", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Id", Desc: "Spreadsheet ID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets.readonly] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets.readonly")},
  		},
  		&{Name: "Range", Desc: "The range the values cover, in A1 notation. This is a string lik"..., TypeName: "string"},
  		&{Name: "ValueRender", Desc: "How values should be represented in the output.", Default: "formatted", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "Exported sheet data"}},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets sheet import`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Data", Desc: "Data file path", TypeName: "Data"},
  		&{Name: "Id", Desc: "Spreadsheet Id", TypeName: "string"},
  		&{Name: "InputRaw", Desc: "Raw input", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Range", Desc: "The range the values cover, in A1 notation. This is a string lik"..., TypeName: "string"},
  	},
  	GridDataInput:  {&{Name: "Data", Desc: "Input data file"}},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets sheet list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "Spreadsheet ID", TypeName: "string"},
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets.readonly] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets.readonly")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `services google sheets spreadsheet create`



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
- 			Default:  "&{default [https://www.googleapis.com/auth/spreadsheets] <nil>}",
+ 			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_sheets",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/spreadsheets")},
  		},
  		&{Name: "Title", Desc: "Title of the spreadsheet", TypeName: "string"},
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
  	... // 9 identical fields
  	ConnScopes:     {"Peer": "slack"},
  	Services:       {"slack"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 10 identical fields
  }
```
# Command spec changed: `sharedfolder leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "KeepCopy", Desc: "Keep a copy of the folder's contents upon relinquishing membership.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "The ID for the shared folder.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Message", Desc: "Custom message for invitation", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Shared folder path of the member", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "Silent", Desc: "Do not send invitation email", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the folder member", TypeName: "string"},
  		&{Name: "LeaveCopy", Desc: "If true, members of this shared folder will get a copy of this f"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "Shared folder path of the member", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "The ID for the shared folder.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "The ID for the shared folder.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeMounted", Desc: "Include mounted folders.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AclUpdatePolicy", Desc: "Who can change a shared folder's access control list (ACL).", Default: "owner", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "MemberPolicy", Desc: "Who can be a member of this shared folder.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "Path to be shared", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "SharedLinkPolicy", Desc: "Who can view shared links in this folder.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "LeaveCopy", Desc: "If true, members of this shared folder will get a copy of this f"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "Path to be unshared", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.read"),
  				string("sharing.read"),
  				string("sharing.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedlink create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Expires", Desc: "Expiration date/time of the shared link", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "Password", Desc: "Password", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "TeamOnly", Desc: "Link is accessible only by team members", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedlink delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Path", Desc: "File or folder path to remove shared link", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.write"),
  			},
  		},
  		&{Name: "Recursive", Desc: "Attempt to remove the file hierarchy", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedlink file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Password", Desc: "Password for the shared link", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.metadata.read"),
  				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "Shared link URL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedlink info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Password", Desc: "Password of the link if required.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  		&{Name: "Url", Desc: "URL of the shared link", TypeName: "domain.dropbox.model.mo_url.url_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("sharing.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team activity batch user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("events.read"),
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team activity daily event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("events.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("events.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team activity user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("events.read"),
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin group role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Group", Desc: "Group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin group role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ExceptionGroup", Desc: "Exception group name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("groups.read"),
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "MemberRoles", Desc: "Member to admin-role mappings"}},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin role clear`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Email", Desc: "Email address of the member", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("members.write"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "RoleId", Desc: "Role ID", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team admin role list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team content legacypaper count`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FilterBy", Desc: "Specify how the Paper docs should be filtered (doc_created/doc_a"..., Default: "docs_created", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Format", Desc: "Export file format (html/markdown)", Default: "html", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Path", Desc: "Export folder path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "FilterBy", Desc: "Specify how the Paper docs should be filtered (doc_created/doc_a"..., Default: "docs_created", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team content member size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team content mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "MemberNamePrefix", Desc: "Filter members. Filter by name match to the prefix."},
  		&{Name: "MemberNameSuffix", Desc: "Filter members. Filter by name match to the suffix."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team device list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sessions.list"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("sessions.modify"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team filerequest clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("file_requests.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team linkedapp list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sessions.list"),
+ 				string("team_info.read"),
  			},
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
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team namespace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team namespace member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team namespace summary`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team report activity`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `team report devices`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `team report membership`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `team report storage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `team runas file batch copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas file sync batch up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 7 identical elements
  		&{Name: "NameNameSuffix", Desc: "Filter by name. Filter by name match to the suffix."},
  		&{Name: "Overwrite", Desc: "Overwrite existing file on the target path if that exists.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("members.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder batch leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "KeepCopy", Desc: "Keep a copy of the folder's contents upon relinquishing membership.", Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AclUpdatePolicy", Desc: "Who can add and remove members of this shared folder.", Default: "owner", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "MemberPolicy", Desc: "Who can be a member of this shared folder.", Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SharedLinkPolicy", Desc: "The policy to apply to shared links created for content inside t"..., Default: "anyone", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder batch unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "LeaveCopy", Desc: "If true, members of this shared folder will get a copy of this f"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder isolate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `team runas sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Message", Desc: "Custom message for invitation", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 4 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Silent", Desc: "Do not send invitation email", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "LeaveCopy", Desc: "If true, members of this shared folder will get a copy of this f"..., Default: "false", TypeName: "bool", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 4 identical elements
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "The ID for the shared folder.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 3 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "SharedFolderId", Desc: "The ID for the shared folder.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team runas sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeMounted", Desc: "Include mounted folders.", Default: "false", TypeName: "bool", ...},
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				... // 2 identical elements
  				string("team_data.member"),
  				string("team_data.team_space"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink cap expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "New visibility setting", Default: "team_only", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink delete links`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink delete member`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.read"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  		&{Name: "Visibility", Desc: "Filter links by visibility (all/public/team_only/password)", Default: "all", TypeName: "essentials.model.mo_string.select_string", ...},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
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
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink update password`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "NewVisibility", Desc: "New visibility setting", Default: "team_only", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("members.read"),
  				string("sharing.write"),
  				string("team_data.member"),
+ 				string("team_info.read"),
  			},
  		},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.content.read"),
- 				string("team_data.content.write"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder batch archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder batch permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
  				string("team_info.read"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder batch replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder partial replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
  				string("team_info.read"),
  			},
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
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `teamfolder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
- 		"Dst": "dropbox_scoped_team",
+ 		"Dst": "dropbox_team",
- 		"Src": "dropbox_scoped_team",
+ 		"Src": "dropbox_team",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 12 identical fields
  }
```
# Command spec changed: `util monitor client`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MonitorInterval", Desc: "Monitoring interval (seconds)", Default: "10", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "Name", Desc: "Client name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{
+ 				string("account_info.read"),
  				string("files.content.write"),
  			},
  		},
  		&{Name: "SyncInterval", Desc: "Sync to Dropbox interval (seconds)", Default: "3600", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{Name: "SyncPath", Desc: "Path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
