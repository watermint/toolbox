---
layout: release
title: Changes of Release 140
lang: en
---

# Changes between `Release 140` to `Release 141`

# Commands added


| Command                       | Title                                                     |
|-------------------------------|-----------------------------------------------------------|
| dev doc knowledge             | Generate reduced knowledge base                           |
| dev doc msg add               | Add a new message                                         |
| dev doc msg catalogue_options | Generate option descriptions for all recipes in catalogue |
| dev doc msg delete            | Delete a message                                          |
| dev doc msg list              | List messages                                             |
| dev doc msg options           | Generate option descriptions for SelectString fields      |
| dev doc msg translate         | Translation helper                                        |
| dev doc msg update            | Update a message                                          |
| dev doc msg verify            | Verify message template variables consistency             |
| dev doc review approve        | Mark a message as reviewed                                |
| dev doc review batch          | Review and approve messages in batch                      |
| dev doc review list           | List unreviewed messages                                  |
| dev doc review options        | Review missing SelectString option descriptions           |



# Command spec changed: `asana team task list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List task of the team",
+ 	Title:   "List tasks of the team",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```

## Changed report: tasks

```
  &dc_recipe.Report{
  	Name: "tasks",
  	Desc: "The task is the basic object around which many operations in Asa"...,
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "completed", Desc: "True if the task is currently marked complete, false if not."},
  		&{Name: "completed_at", Desc: "The time at which this task was completed, or null if the task i"...},
  		&{
  			Name: "due_at",
  			Desc: strings.Join({
  				"Date and time on which this task is due, or null if the task has",
  				" no due time.",
- 				" ",
  			}, ""),
  		},
  		&{Name: "due_on", Desc: "Date on which this task is due, or null if the task has no due d"...},
  	},
  }
```
# Command spec changed: `config auth delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "Delete existing auth credential",
- 	Desc:    "",
+ 	Desc:    "Remove stored authentication credentials for a specific service account. This is useful when you need to revoke access, change accounts, or clean up old authentication tokens. The command requires both the application key name and peer name to identify the"...,
  	Remarks: "",
  	Path:    "config auth delete",
  	... // 19 identical fields
  }
```

## Changed report: deleted

```
  &dc_recipe.Report{
  	Name:    "deleted",
- 	Desc:    "Auth credential data",
+ 	Desc:    "Authentication credential data",
  	Columns: {&{Name: "key_name", Desc: "Application name"}, &{Name: "scope", Desc: "Auth scope"}, &{Name: "peer_name", Desc: "Peer name"}, &{Name: "description", Desc: "Description"}, ...},
  }
```
# Command spec changed: `config auth list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List all auth credentials",
- 	Desc:    "",
+ 	Desc:    "Display all stored authentication credentials and their details including application names, scopes, peer names, and timestamps. This is useful for auditing access, managing multiple accounts, and understanding which services you're authenticated with.",
  	Remarks: "",
  	Path:    "config auth list",
  	... // 19 identical fields
  }
```

## Changed report: entity

```
  &dc_recipe.Report{
  	Name:    "entity",
- 	Desc:    "Auth credential data",
+ 	Desc:    "Authentication credential data",
  	Columns: {&{Name: "key_name", Desc: "Application name"}, &{Name: "scope", Desc: "Auth scope"}, &{Name: "peer_name", Desc: "Peer name"}, &{Name: "description", Desc: "Description"}, ...},
  }
```
# Command spec changed: `config feature disable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "disable",
  	Title:   "Disable a feature.",
- 	Desc:    "",
+ 	Desc:    "Turn off a specific feature in the watermint toolbox configuration. Features control various aspects of the application's behavior, performance settings, and experimental functionality. Disabling features can help with troubleshooting or reverting to previ"...,
  	Remarks: "",
  	Path:    "config feature disable",
  	... // 19 identical fields
  }
```
# Command spec changed: `config feature enable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "enable",
  	Title:   "Enable a feature.",
- 	Desc:    "",
+ 	Desc:    "Turn on a specific feature in the watermint toolbox configuration. Features control various aspects of the application's behavior, performance settings, and experimental functionality. Enabling features allows you to access new capabilities or modify appli"...,
  	Remarks: "",
  	Path:    "config feature enable",
  	... // 19 identical fields
  }
```
# Command spec changed: `config feature list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List available optional features.",
- 	Desc:    "",
+ 	Desc:    "Display all available optional features in the watermint toolbox with their descriptions, current status, and configuration details. This is useful for understanding what functionality can be enabled or disabled, and for managing feature preferences.",
  	Remarks: "",
  	Path:    "config feature list",
  	... // 19 identical fields
  }
```
# Command spec changed: `config license install`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "install",
  	Title:   "Install a license key",
- 	Desc:    "",
+ 	Desc:    "Install and activate a license key for the watermint toolbox. License keys may be required for certain features, premium functionality, or commercial usage. This command stores the license key securely and validates its authenticity.",
  	Remarks: "",
  	Path:    "config license install",
  	... // 19 identical fields
  }
```
# Command spec changed: `config license list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List available license keys",
- 	Desc:    "",
+ 	Desc:    "Display all installed license keys and their details including expiration dates, enabled features, and status. This is useful for managing multiple licenses, checking license validity, and understanding what functionality is available.",
  	Remarks: "",
  	Path:    "config license list",
  	... // 19 identical fields
  }
```

## Changed report: keys

```
  &dc_recipe.Report{
  	Name: "keys",
  	Desc: "License key summary",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "licensee_name", Desc: "Licensee name"},
  		&{Name: "licensee_email", Desc: "Licensee email"},
  		&{
  			Name: "licensed_recipes",
  			Desc: strings.Join({
  				"Recipes enabled by this licen",
- 				"c",
+ 				"s",
  				"e key",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dev build package`



## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
  		&{Name: "num_folder_created", Desc: "Number of created folders."},
  		&{
  			Name: "num_delete",
- 			Desc: "Number of deleted entry.",
+ 			Desc: "Number of deleted entries.",
  		},
  		&{
  			Name: "num_api_call",
  			Desc: strings.Join({
  				"The number of estimated ",
- 				"upload API call",
+ 				"API calls",
  				" for upload.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dev ci artifact up`



## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
  		&{Name: "num_folder_created", Desc: "Number of created folders."},
  		&{
  			Name: "num_delete",
- 			Desc: "Number of deleted entry.",
+ 			Desc: "Number of deleted entries.",
  		},
  		&{
  			Name: "num_api_call",
  			Desc: strings.Join({
  				"The number of estimated ",
- 				"upload API call",
+ 				"API calls",
  				" for upload.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dev diag endpoint`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "JobId",
- 			Desc:     "Job Id to diagnosis",
+ 			Desc:     "Job ID to diagnose",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: report

```
  &dc_recipe.Report{
  	Name: "report",
  	Desc: "Endpoint statistics",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "endpoint", Desc: "Endpoint URL"},
  		&{Name: "count", Desc: "Number of requests"},
  		&{
  			Name: "count_success",
- 			Desc: "Number of success requests",
+ 			Desc: "Number of successful requests",
  		},
  		&{Name: "count_failure", Desc: "Number of failed requests"},
  	},
  }
```
# Command spec changed: `dev diag throughput`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "JobId", Desc: "Specify Job ID", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Path", Desc: "Path to workspace", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "TimeFormat",
- 			Desc:     "Time format in go's time format",
+ 			Desc:     "Time format in Go time format",
  			Default:  "2006-01-02 15:04:05.999",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: report

```
  &dc_recipe.Report{
  	Name: "report",
  	Desc: "Throughput",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "time", Desc: "Timestamp"},
  		&{Name: "concurrency", Desc: "Concurrency."},
  		&{
  			Name: "success_concurrency",
  			Desc: strings.Join({
  				"Number of concurrent requests of success",
+ 				"ful operations",
  			}, ""),
  		},
  		&{
  			Name: "success_sent",
  			Desc: strings.Join({
  				"Sum of sent bytes of success",
+ 				"ful",
  				" requests in the bucket in bytes",
  			}, ""),
  		},
  		&{
  			Name: "success_received",
  			Desc: strings.Join({
  				"Sum of received bytes of success",
+ 				"ful",
  				" requests in the bucket in bytes",
  			}, ""),
  		},
  		&{Name: "failure_concurrency", Desc: "Number of concurrent requests of failure"},
  		&{Name: "failure_sent", Desc: "Sum of sent bytes of failed requests in the bucket in bytes"},
  		&{Name: "failure_received", Desc: "Sum of received bytes of failed requests in the bucket in bytes"},
  	},
  }
```
# Command spec changed: `dev replay approve`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Id",
- 			Desc:     "Job Id.",
+ 			Desc:     "Job ID.",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "Name", Desc: "Extra name of the approved recipe", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "ReplayPath", Desc: "Replay repository path. Fall back to the environment variable `T"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "WorkspacePath", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev replay recipe`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Id",
- 			Desc:     "Job Id.",
+ 			Desc:     "Job ID.",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev replay remote`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ReplayUrl",
- 			Desc:     "Replay bundle shared link url",
+ 			Desc:     "Replay bundle shared link URL",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev util anonymise`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "JobIdName",
  			Desc: strings.Join({
  				"Filter by job ",
- 				"id",
+ 				"ID",
  				" name Filter by exact match to the name.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "JobIdNamePrefix",
  			Desc: strings.Join({
  				"Filter by job ",
- 				"id",
+ 				"ID",
  				" name Filter by name match to the prefix.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "JobIdNameSuffix",
  			Desc: strings.Join({
  				"Filter by job ",
- 				"id",
+ 				"ID",
  				" name Filter by name match to the suffix.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev util image jpeg`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "NamePrefix", Desc: "Filename prefix", Default: "test_image", TypeName: "string", ...},
  		&{Name: "Path", Desc: "Path to generate files", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			Name:     "Quality",
- 			Desc:     "Quality of jpeg",
+ 			Desc:     "Quality of JPEG",
  			Default:  "75",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{"max": float64(100), "min": float64(1), "value": float64(75)},
  		},
  		&{Name: "Seed", Desc: "Random seed", Default: "1", TypeName: "int", ...},
  		&{Name: "Width", Desc: "Width", Default: "1920", TypeName: "essentials.model.mo_int.range_int", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file account feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "feature",
  	Title:   "List Dropbox account features",
- 	Desc:    "",
+ 	Desc:    "Retrieves and displays the enabled features and capabilities for the connected Dropbox account.",
  	Remarks: "",
  	Path:    "dropbox file account feature",
  	... // 19 identical fields
  }
```

## Changed report: report

```
  &dc_recipe.Report{
  	Name: "report",
  	Desc: "Feature setting for the user",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "paper_as_files", Desc: "When this value is true, the user's Paper docs are accessible in"...},
  		&{Name: "file_locking", Desc: "When this value is True, the user can lock files in shared folders."},
  		&{Name: "team_shared_dropbox", Desc: "This feature contains information about whether or not the user "...},
  		&{
  			Name: "distinct_member_home",
  			Desc: strings.Join({
+ 				"T",
  				"his feature contains information about whether or not the user's",
  				" home namespace is distinct from their root namespace.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file account filesystem`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "filesystem",
  	Title:   "Show Dropbox file system version",
- 	Desc:    "",
+ 	Desc:    "Shows the file system version/type being used by the account (individual or team file system).",
  	Remarks: "",
  	Path:    "dropbox file account filesystem",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file account info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "Dropbox account info",
- 	Desc:    "",
+ 	Desc:    "Displays profile information for the connected Dropbox account including name and email.",
  	Remarks: "",
  	Path:    "dropbox file account info",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file compare account`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "account",
  	Title:   "Compare files of two accounts",
- 	Desc:    "",
+ 	Desc:    "Compares files and folders between two different Dropbox accounts to identify differences.",
  	Remarks: "",
  	Path:    "dropbox file compare account",
  	CliArgs: "-left left -left-path /path/to/compare -right right -right-path "...,
  	CliNote: strings.Join({
  		"If you want to compare different path",
+ 		"s",
  		" in same account, please specify same alias name to `-left` and ",
  		"`-right`.",
  	}, ""),
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	... // 15 identical fields
  }
```

## Changed report: diff

```
  &dc_recipe.Report{
  	Name: "diff",
  	Desc: strings.Join({
  		"This report shows a difference between t",
+ 		"w",
  		"o folders.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "left_hash", Desc: "Content hash of left file"},
  		&{Name: "right_path", Desc: "path of right"},
  		&{
  			Name: "right_kind",
- 			Desc: "folder of file",
+ 			Desc: "folder or file",
  		},
  		&{Name: "right_size", Desc: "size of right file"},
  		&{Name: "right_hash", Desc: "Content hash of right file"},
  	},
  }
```
# Command spec changed: `dropbox file compare local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "local",
  	Title:   "Compare local folders and Dropbox folders",
- 	Desc:    "",
+ 	Desc:    "Compares local files and folders with their Dropbox counterparts to identify differences.",
  	Remarks: "",
  	Path:    "dropbox file compare local",
  	... // 19 identical fields
  }
```

## Changed report: diff

```
  &dc_recipe.Report{
  	Name: "diff",
  	Desc: strings.Join({
  		"This report shows a difference between t",
+ 		"w",
  		"o folders.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "left_hash", Desc: "Content hash of left file"},
  		&{Name: "right_path", Desc: "path of right"},
  		&{
  			Name: "right_kind",
- 			Desc: "folder of file",
+ 			Desc: "folder or file",
  		},
  		&{Name: "right_size", Desc: "size of right file"},
  		&{Name: "right_hash", Desc: "Content hash of right file"},
  	},
  }
```

## Changed report: skip

```
  &dc_recipe.Report{
  	Name: "skip",
  	Desc: strings.Join({
  		"This report shows a difference between t",
+ 		"w",
  		"o folders.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "left_hash", Desc: "Content hash of left file"},
  		&{Name: "right_path", Desc: "path of right"},
  		&{
  			Name: "right_kind",
- 			Desc: "folder of file",
+ 			Desc: "folder or file",
  		},
  		&{Name: "right_size", Desc: "size of right file"},
  		&{Name: "right_hash", Desc: "Content hash of right file"},
  	},
  }
```
# Command spec changed: `dropbox file copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "copy",
  	Title:   "Copy files",
- 	Desc:    "",
+ 	Desc:    "Copies files or folders from one location to another within the same Dropbox account.",
  	Remarks: "",
  	Path:    "dropbox file copy",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "Delete file or folder",
- 	Desc:    "",
+ 	Desc:    "Permanently deletes files or folders from Dropbox (irreversible operation).",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file export doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "doc",
  	Title:   "Export document",
- 	Desc:    "",
+ 	Desc:    "Exports Dropbox Paper documents and Google Docs to local files in specified formats.",
  	Remarks: "(Experimental)",
  	Path:    "dropbox file export doc",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: strings.Join({
  		"This report shows ",
- 		"a",
+ 		"the",
  		" result of exporting",
+ 		" a",
  		" file.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
  		&{
  			Name: "size",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
+ 			Desc: "The file size in bytes.",
  		},
  		&{Name: "export_name", Desc: "File name for export file."},
  		&{Name: "export_size", Desc: "File size of export file."},
  	},
  }
```
# Command spec changed: `dropbox file export url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Export a document from the URL",
- 	Desc:    "",
+ 	Desc:    "Exports a file from Dropbox by downloading it from a shared link URL.",
  	Remarks: "",
  	Path:    "dropbox file export url",
  	CliArgs: strings.Join({
  		"-local-path /LOCAL/PATH/TO/",
- 		"export",
+ 		"EXPORT",
  		" -url DOCUMENT_URL",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 16 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: strings.Join({
  		"This report shows ",
- 		"a",
+ 		"the",
  		" result of exporting",
+ 		" a",
  		" file.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
  		&{
  			Name: "size",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
+ 			Desc: "The file size in bytes.",
  		},
  		&{Name: "export_name", Desc: "File name for export file."},
  		&{Name: "export_size", Desc: "File size of export file."},
  	},
  }
```
# Command spec changed: `dropbox file import batch url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Batch import files from URL",
- 	Desc:    "",
+ 	Desc:    "Imports multiple files into Dropbox by downloading from a list of URLs.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file import batch url",
  	... // 19 identical fields
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
  			Name: "input.url",
- 			Desc: "Url to download",
+ 			Desc: "URL to download",
  		},
  		&{Name: "input.path", Desc: "Path to store file (use path given by `-path` when the record is"...},
  		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		... // 6 identical elements
  	},
  }
```
# Command spec changed: `dropbox file import url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Import file from the URL",
- 	Desc:    "",
+ 	Desc:    "Imports a single file into Dropbox by downloading from a specified URL.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file import url",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "Resolve metadata of the path",
- 	Desc:    "",
+ 	Desc:    "Retrieves and displays metadata information for a specific file or folder path.",
  	Remarks: "",
  	Path:    "dropbox file info",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List files and folders",
- 	Desc:    "",
+ 	Desc:    "Lists files and folders at a given path with options for recursive listing and filtering.",
  	Remarks: "",
  	Path:    "dropbox file list",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...},
  		&{
  			Name: "IncludeExplicitSharedMembers",
  			Desc: strings.Join({
- 				" ",
  				"If true, the results will include a flag for each file indicatin",
  				"g whether or not that file has any explicit members.",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "IncludeMountedFolders",
  			Desc: strings.Join({
- 				" ",
  				"If true, the results will include entries under mounted folders ",
  				"which include",
- 				"s",
  				" app folder, shared folder and team folder.",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file lock acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "acquire",
  	Title:   "Lock a file",
- 	Desc:    "",
+ 	Desc:    "Acquires an exclusive lock on a file to prevent others from editing it.",
  	Remarks: "",
  	Path:    "dropbox file lock acquire",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "Release all locks under the specified path",
- 	Desc:    "",
+ 	Desc:    "Releases all file locks held by the current user across the account.",
  	Remarks: "",
  	Path:    "dropbox file lock all release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file lock batch acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "acquire",
  	Title:   "Lock multiple files",
- 	Desc:    "",
+ 	Desc:    "Acquires locks on multiple files in a single batch operation.",
  	Remarks: "",
  	Path:    "dropbox file lock batch acquire",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file lock batch release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "Release multiple locks",
- 	Desc:    "",
+ 	Desc:    "Releases locks on multiple files in a single batch operation.",
  	Remarks: "",
  	Path:    "dropbox file lock batch release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List locks under the specified path",
- 	Desc:    "",
+ 	Desc:    "Lists all files that are currently locked, showing lock holder information.",
  	Remarks: "",
  	Path:    "dropbox file lock list",
  	... // 19 identical fields
  }
```

## Changed report: lock

```
  &dc_recipe.Report{
  	Name: "lock",
  	Desc: "Lock information",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "Release a lock",
- 	Desc:    "",
+ 	Desc:    "Releases the lock on a specific file, allowing others to edit it.",
  	Remarks: "",
  	Path:    "dropbox file lock release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file merge`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "merge",
  	Title:   "Merge paths",
- 	Desc:    "",
+ 	Desc:    "Merges contents from one folder into another, with options for dry-run and empty folder handling.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file merge",
  	CliArgs: "-from /from/path -to /path/to",
  	CliNote: strings.Join({
  		"Please add `-dry-run=false` option after verify",
+ 		"ing",
  		" integrity of expected result.",
  	}, ""),
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "From",
- 			Desc:     "Path for merge",
+ 			Desc:     "Source path for merge",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "KeepEmptyFolder",
- 			Desc:     "Keep empty folder after merge",
+ 			Desc:     "Keep empty folders after merge",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:     "To",
- 			Desc:     "Path to merge",
+ 			Desc:     "Destination path for merge",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "WithinSameNamespace",
  			Desc: strings.Join({
  				"Do not cross namespace. Th",
- 				"at is for",
+ 				"is is to",
  				" preserve sharing permission",
- 				" including a shared link",
+ 				"s including shared links",
  			}, ""),
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
# Command spec changed: `dropbox file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "move",
  	Title:   "Move files",
- 	Desc:    "",
+ 	Desc:    "Moves files or folders from one location to another within Dropbox (irreversible operation).",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file move",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "replication",
  	Title: "Replicate file content to the other account",
  	Desc: strings.Join({
- 		"This command will replicate files/folders. But it does not inclu",
- 		"de sharing permissions. The command replicates only for folder c",
- 		"ontents of given path",
+ 		"Replicates files and folders from one Dropbox account to another",
+ 		", mirroring the source structure",
  		".",
  	}, ""),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file replication",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			Name:     "Dst",
- 			Desc:     "Account alias (destionation)",
+ 			Desc:     "Account alias (destination)",
  			Default:  "dst",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{string("account_info.read"), string("files.content.write"), string("files.metadata.read")},
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

## Changed report: replication_diff

```
  &dc_recipe.Report{
  	Name: "replication_diff",
  	Desc: strings.Join({
  		"This report shows a difference between t",
+ 		"w",
  		"o folders.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "left_hash", Desc: "Content hash of left file"},
  		&{Name: "right_path", Desc: "path of right"},
  		&{
  			Name: "right_kind",
- 			Desc: "folder of file",
+ 			Desc: "folder or file",
  		},
  		&{Name: "right_size", Desc: "size of right file"},
  		&{Name: "right_hash", Desc: "Content hash of right file"},
  	},
  }
```
# Command spec changed: `dropbox file request create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create a file request",
- 	Desc:    "",
+ 	Desc:    "Creates a file request folder where others can upload files without Dropbox access.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file request create",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AllowLateUploads", Desc: "If set, allow uploads after the deadline has passed (one_day/two"..., TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Deadline", Desc: "The deadline for this file request.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name: "Path",
  			Desc: strings.Join({
  				"The path for the folder in ",
- 				"the ",
  				"Dropbox where uploaded files will be sent.",
  			}, ""),
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Title", Desc: "The title of the file request", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: file_request

```
  &dc_recipe.Report{
  	Name: "file_request",
  	Desc: "This report shows a list of file requests.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "id", Desc: "The Id of the file request"},
  		&{Name: "url", Desc: "The URL of the file request"},
  		&{Name: "title", Desc: "The title of the file request"},
  		&{
  			Name: "created",
  			Desc: strings.Join({
  				"Date/time ",
- 				"of",
+ 				"when",
  				" the file request was created.",
  			}, ""),
  		},
  		&{Name: "is_open", Desc: "Whether or not the file request is open."},
  		&{Name: "file_count", Desc: "The number of files this file request has received."},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `dropbox file request delete closed`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "closed",
  	Title:   "Delete all closed file requests on this account.",
- 	Desc:    "",
+ 	Desc:    "Deletes file requests that have been closed and are no longer accepting uploads.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file request delete closed",
  	... // 19 identical fields
  }
```

## Changed report: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "This report shows a list of file requests.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "id", Desc: "The Id of the file request"},
  		&{Name: "url", Desc: "The URL of the file request"},
  		&{Name: "title", Desc: "The title of the file request"},
  		&{
  			Name: "created",
  			Desc: strings.Join({
  				"Date/time ",
- 				"of",
+ 				"when",
  				" the file request was created.",
  			}, ""),
  		},
  		&{Name: "is_open", Desc: "Whether or not the file request is open."},
  		&{Name: "file_count", Desc: "The number of files this file request has received."},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `dropbox file request delete url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Delete a file request by the file request URL",
- 	Desc:    "",
+ 	Desc:    "Deletes a specific file request using its URL.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file request delete url",
  	... // 19 identical fields
  }
```

## Changed report: deleted

```
  &dc_recipe.Report{
  	Name: "deleted",
  	Desc: "This report shows a list of file requests.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "id", Desc: "The Id of the file request"},
  		&{Name: "url", Desc: "The URL of the file request"},
  		&{Name: "title", Desc: "The title of the file request"},
  		&{
  			Name: "created",
  			Desc: strings.Join({
  				"Date/time ",
- 				"of",
+ 				"when",
  				" the file request was created.",
  			}, ""),
  		},
  		&{Name: "is_open", Desc: "Whether or not the file request is open."},
  		&{Name: "file_count", Desc: "The number of files this file request has received."},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `dropbox file request list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List file requests of the individual account",
- 	Desc:    "",
+ 	Desc:    "Lists all file requests in the account with their status and details.",
  	Remarks: "",
  	Path:    "dropbox file request list",
  	... // 19 identical fields
  }
```

## Changed report: file_requests

```
  &dc_recipe.Report{
  	Name: "file_requests",
  	Desc: "This report shows a list of file requests.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "id", Desc: "The Id of the file request"},
  		&{Name: "url", Desc: "The URL of the file request"},
  		&{Name: "title", Desc: "The title of the file request"},
  		&{
  			Name: "created",
  			Desc: strings.Join({
  				"Date/time ",
- 				"of",
+ 				"when",
  				" the file request was created.",
  			}, ""),
  		},
  		&{Name: "is_open", Desc: "Whether or not the file request is open."},
  		&{Name: "file_count", Desc: "The number of files this file request has received."},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `dropbox file restore all`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "all",
  	Title:   "Restore files under given path",
- 	Desc:    "",
+ 	Desc:    "Restores all deleted files and folders within a specified path.",
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "dropbox file restore all",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file restore ext`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "ext",
  	Title:   "Restore files with a specific extension",
- 	Desc:    "",
+ 	Desc:    "Restores deleted files matching specific file extensions within a path.",
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "dropbox file restore ext",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file revision download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "download",
  	Title:   "Download the file revision",
- 	Desc:    "",
+ 	Desc:    "Downloads a specific revision/version of a file from its revision history.",
  	Remarks: "",
  	Path:    "dropbox file revision download",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file revision list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List file revisions",
- 	Desc:    "",
+ 	Desc:    "Lists all available revisions for a file showing modification times and sizes.",
  	Remarks: "",
  	Path:    "dropbox file revision list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file revision restore`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "restore",
  	Title:   "Restore the file revision",
- 	Desc:    "",
+ 	Desc:    "Restores a file to a previous revision from its version history.",
  	Remarks: "",
  	Path:    "dropbox file revision restore",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file search content`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "content",
  	Title:   "Search file content",
- 	Desc:    "",
+ 	Desc:    "Searches for files by content with options for file type and category filtering.",
  	Remarks: "",
  	Path:    "dropbox file search content",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string_internal", TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}}},
  		&{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name: "MaxResults",
  			Desc: strings.Join({
  				"Maximum number of entr",
- 				"y",
+ 				"ies",
  				" to return",
  			}, ""),
  			Default:  "25",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{"max": float64(100000), "min": float64(0), "value": float64(25)},
  		},
  		&{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file search name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "name",
  	Title:   "Search file name",
- 	Desc:    "",
+ 	Desc:    "Searches for files and folders by name pattern across the Dropbox account.",
  	Remarks: "",
  	Path:    "dropbox file search name",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file share info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "Retrieve sharing information of the file",
- 	Desc:    "",
+ 	Desc:    "Retrieves sharing information and permissions for a specific file or folder.",
  	Remarks: "",
  	Path:    "dropbox file share info",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			Name:     "Path",
- 			Desc:     "File",
+ 			Desc:     "File path",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "Get shared folder info",
- 	Desc:    "",
+ 	Desc:    "Displays detailed information about a specific shared folder including members and permissions.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder info",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:     "SharedFolderId",
- 			Desc:     "Namespace ID",
+ 			Desc:     "Shared folder ID",
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

## Changed report: policies

```
  &dc_recipe.Report{
  	Name: "policies",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "leave",
- 	Title: "Leave from the shared folder",
+ 	Title: "Leave the shared folder",
  	Desc: strings.Join({
- 		"Upon success, the current user will no longer have access to the",
- 		" folder. Please use `dropbox file sharedfolder list` command to ",
- 		"find the shared_folder_id of the folder you want to leave",
+ 		"Removes yourself from a shared folder you've been added to",
  		".",
  	}, ""),
  	Remarks: "",
  	Path:    "dropbox file sharedfolder leave",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List shared folder(s)",
+ 	Title:   "List shared folders",
- 	Desc:    "",
+ 	Desc:    "Lists all shared folders you have access to with their sharing details.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder list",
  	... // 19 identical fields
  }
```

## Changed report: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Add a member to the shared folder",
- 	Desc:    "",
+ 	Desc:    "Adds new members to a shared folder with specified access permissions.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder member add",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Email", Desc: "Email address of the folder member", TypeName: "string"},
  		&{Name: "Message", Desc: "Custom message for invitation", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
  			Name:     "Path",
- 			Desc:     "Shared folder path of the member",
+ 			Desc:     "Path to the shared folder",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Silent", Desc: "Do not send invitation email", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"Delet",
+ 		"Remov",
  		"e a member from the shared folder",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Removes members from a shared folder, revoking their access.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder member delete",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Email", Desc: "Email address of the folder member", TypeName: "string"},
  		&{Name: "LeaveCopy", Desc: "If true, members of this shared folder will get a copy of this f"..., Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Path",
- 			Desc:     "Shared folder path of the member",
+ 			Desc:     "Path to the shared folder",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List shared folder member(s)",
+ 	Title:   "List shared folder members",
- 	Desc:    "",
+ 	Desc:    "Lists all members of a shared folder with their access levels and email addresses.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder member list",
  	... // 19 identical fields
  }
```

## Changed report: member

```
  &dc_recipe.Report{
  	Name: "member",
  	Desc: "This report shows a list of members of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		&{Name: "is_inside_team_folder", Desc: "Whether this folder is inside of a team folder."},
  		... // 7 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Add the shared folder to the current user's Dropbox",
- 	Desc:    "",
+ 	Desc:    "Mounts a shared folder to your Dropbox, making it appear in your file structure.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount add",
  	... // 19 identical fields
  }
```

## Changed report: mount

```
  &dc_recipe.Report{
  	Name: "mount",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "delete",
- 	Title: "The current user unmounts the designated folder.",
+ 	Title: "Unmount the shared folder",
  	Desc: strings.Join({
  		"U",
- 		"pon success, the current user cannot access the folder unless ad",
- 		"ding the folder again. Please use `dropbox file sharedfolder mou",
- 		"nt list` command to find the shared_folder_id of the folder you ",
- 		"want to delete",
+ 		"nmounts a shared folder from your Dropbox without leaving the fo",
+ 		"lder",
  		".",
  	}, ""),
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount delete",
  	... // 19 identical fields
  }
```

## Changed report: mount

```
  &dc_recipe.Report{
  	Name: "mount",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
  		"List all shared folders the current user ",
+ 		"has ",
  		"mounted",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Lists all shared folders currently mounted in your Dropbox.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount list",
  	... // 19 identical fields
  }
```

## Changed report: mounts

```
  &dc_recipe.Report{
  	Name: "mounts",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "mountable",
  	Title:   "List all shared folders the current user can mount",
- 	Desc:    "",
+ 	Desc:    "Lists shared folders that can be mounted but aren't currently in your Dropbox.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount mountable",
  	... // 19 identical fields
  }
```

## Changed report: mountables

```
  &dc_recipe.Report{
  	Name: "mountables",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "share",
  	Title:   "Share a folder",
- 	Desc:    "",
+ 	Desc:    "Creates a shared folder from an existing folder with configurable sharing policies and permissions.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder share",
  	... // 19 identical fields
  }
```

## Changed report: shared

```
  &dc_recipe.Report{
  	Name: "shared",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 10 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sharedfolder unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "unshare",
  	Title:   "Unshare a folder",
- 	Desc:    "",
+ 	Desc:    "Stops sharing a folder and optionally leaves a copy for current members.",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder unshare",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file sharedlink create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create shared link",
- 	Desc:    "",
+ 	Desc:    "Creates a shared link for a file or folder with optional password protection and expiration date.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file sharedlink create",
  	... // 19 identical fields
  }
```

## Changed report: created

```
  &dc_recipe.Report{
  	Name: "created",
  	Desc: strings.Join({
  		"T",
- 		"H",
+ 		"h",
  		"is report shows a list of shared links.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "expires", Desc: "Expiration time, if set."},
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file sharedlink delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "delete",
  	Title: "Remove shared links",
  	Desc: strings.Join({
  		"This command will delete shared links based on the path in ",
- 		"the ",
  		"Dropbox",
  	}, ""),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file sharedlink delete",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Path", Desc: "File or folder path to remove shared link", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:     "Recursive",
- 			Desc:     "Attempt to remove the file hierarchy",
+ 			Desc:     "Remove shared links recursively",
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

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "input.expires", Desc: "Expiration time, if set."},
  		&{Name: "input.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "input.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List of shared link(s)",
+ 	Title:   "List shared links",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
  	Desc: strings.Join({
  		"T",
- 		"H",
+ 		"h",
  		"is report shows a list of shared links.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "expires", Desc: "Expiration time, if set."},
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "size",
  	Title:   "Storage usage",
- 	Desc:    "",
+ 	Desc:    "Calculates and reports the size of folders and their contents at specified depth levels.",
  	Remarks: "",
  	Path:    "dropbox file size",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			Name: "Depth",
  			Desc: strings.Join({
  				"Report ",
- 				"an entry for all files and folders depth folders deep",
+ 				"entries for files and folders up to the specified depth",
  			}, ""),
  			Default:  "2",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{"max": float64(300), "min": float64(1), "value": float64(2)},
  		},
  		&{Name: "Path", Desc: "Path to scan", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: size

```
  &dc_recipe.Report{
  	Name: "size",
  	Desc: "Folder size",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "path", Desc: "Path"},
  		&{
  			Name: "depth",
- 			Desc: "Folder depth.",
+ 			Desc: "Folder depth",
  		},
  		&{Name: "size", Desc: "Size in bytes"},
  		&{Name: "num_file", Desc: "Number of files in this folder and child folders"},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `dropbox file sync down`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "down",
  	Title:   "Downstream sync with Dropbox",
- 	Desc:    "",
+ 	Desc:    "Downloads files from Dropbox to local filesystem with filtering and overwrite options.",
  	Remarks: "",
  	Path:    "dropbox file sync down",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			Name: "Delete",
  			Desc: strings.Join({
  				"Delete local file if a file ",
+ 				"is ",
  				"removed on Dropbox",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
  		&{Name: "num_folder_created", Desc: "Number of created folders."},
  		&{
  			Name: "num_delete",
- 			Desc: "Number of deleted entry.",
+ 			Desc: "Number of deleted entries.",
  		},
  		&{
  			Name: "num_api_call",
  			Desc: strings.Join({
  				"The number of estimated ",
- 				"upload API call",
+ 				"API calls",
  				" for upload.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file sync online`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "online",
  	Title:   "Sync online files",
- 	Desc:    "",
+ 	Desc:    "Synchronizes files between two different locations within Dropbox online storage.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file sync online",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			Name: "Delete",
  			Desc: strings.Join({
  				"Delete file if a file ",
+ 				"is ",
  				"removed in source path",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "NameDisableIgnore", Desc: "Filter by name. Filter system file or ignore files."},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
  		&{Name: "num_folder_created", Desc: "Number of created folders."},
  		&{
  			Name: "num_delete",
- 			Desc: "Number of deleted entry.",
+ 			Desc: "Number of deleted entries.",
  		},
  		&{
  			Name: "num_api_call",
  			Desc: strings.Join({
  				"The number of estimated ",
- 				"upload API call",
+ 				"API calls",
  				" for upload.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "up",
  	Title:   "Upstream sync with Dropbox",
- 	Desc:    "",
+ 	Desc:    "Uploads files from local filesystem to Dropbox with filtering and overwrite options.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox file sync up",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "BatchSize", Desc: "Batch commit size", Default: "50", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name: "Delete",
  			Desc: strings.Join({
  				"Delete Dropbox file if a file ",
+ 				"is ",
  				"removed locally",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "LocalPath", Desc: "Local file path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
  		&{Name: "num_folder_created", Desc: "Number of created folders."},
  		&{
  			Name: "num_delete",
- 			Desc: "Number of deleted entry.",
+ 			Desc: "Number of deleted entries.",
  		},
  		&{
  			Name: "num_api_call",
  			Desc: strings.Join({
  				"The number of estimated ",
- 				"upload API call",
+ 				"API calls",
  				" for upload.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox file tag add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Add tag to file or folder",
- 	Desc:    "",
+ 	Desc:    "Adds a custom tag to a file or folder for organization and categorization.",
  	Remarks: "",
  	Path:    "dropbox file tag add",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BasePath",
- 			Desc:     "Base path for adding a tag.",
+ 			Desc:     "Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a per"...,
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "Path", Desc: "File or folder path to add a tag.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Tag", Desc: "Tag to add to the file or folder.", TypeName: "string"},
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
  	Name:    "delete",
  	Title:   "Delete a tag from the file/folder",
- 	Desc:    "",
+ 	Desc:    "Removes a specific tag from a file or folder.",
  	Remarks: "",
  	Path:    "dropbox file tag delete",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BasePath",
- 			Desc:     "Base path for removing a tag.",
+ 			Desc:     "Choose the file path standard. This is an option for Dropbox for Teams in particular. If you are using the personal version of Dropbox, it basically doesn't matter what you choose. In Dropbox for Teams, if you select `home` in the updated team space, a per"...,
  			Default:  "root",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string("root"), string("home")}},
  		},
  		&{Name: "Path", Desc: "File or folder path to remove a tag.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Tag", Desc: "Tag name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file tag list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "List tags of the path",
- 	Desc:    "",
+ 	Desc:    "Lists all tags associated with a specific file or folder path.",
  	Remarks: "",
  	Path:    "dropbox file tag list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file template apply`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "apply",
  	Title:   "Apply file/folder structure template to the Dropbox path",
- 	Desc:    "",
+ 	Desc:    "Applies a saved file/folder structure template to create directories and files in Dropbox.",
  	Remarks: "",
  	Path:    "dropbox file template apply",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file template capture`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "capture",
  	Title:   "Capture file/folder structure as template from Dropbox path",
- 	Desc:    "",
+ 	Desc:    "Captures the file/folder structure from a Dropbox path and saves it as a reusable template.",
  	Remarks: "",
  	Path:    "dropbox file template capture",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox file watch`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "watch",
  	Title:   "Watch file activities",
- 	Desc:    "",
+ 	Desc:    "Monitors a path for changes and outputs file/folder modifications in real-time.",
  	Remarks: "",
  	Path:    "dropbox file watch",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox paper overwrite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "overwrite",
  	Title: strings.Join({
  		"Overwrite ",
+ 		"an ",
  		"existing Paper document",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team activity batch user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "user",
  	Title: strings.Join({
  		"Scan a",
- 		"ctivities for multiple user",
+ 		"nd retrieve activity logs for multiple team members in batch, us",
+ 		"eful for compliance auditing and user behavior analysi",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "This command processes a list of user email addresses from a file and retrieves their activity logs within a specified time range. Useful for HR investigations, compliance reporting, or analyzing patterns across specific user groups.",
  	Remarks: "",
  	Path:    "dropbox team activity batch user",
  	... // 19 identical fields
  }
```

## Changed report: combined

```
  &dc_recipe.Report{
  	Name: "combined",
  	Desc: strings.Join({
  		"This report shows a",
- 		"n a",
  		"ctivity logs",
- 		" with",
  		" mostly compatible with Dropbox for teams'",
- 		"s",
  		" activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```

## Changed report: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: strings.Join({
  		"This report shows a",
- 		"n a",
  		"ctivity logs",
- 		" with",
  		" mostly compatible with Dropbox for teams'",
- 		"s",
  		" activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `dropbox team activity daily event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "event",
- 	Title:   "Report activities by day",
+ 	Title:   "Generate daily activity reports showing team events grouped by date, helpful for tracking team usage patterns and security monitoring",
- 	Desc:    "",
+ 	Desc:    "Aggregates team activity events by day, making it easier to identify trends and anomalies in team behavior. Particularly useful for creating daily security reports, tracking adoption of new features, or identifying unusual activity patterns that might indi"...,
  	Remarks: "",
  	Path:    "dropbox team activity daily event",
  	... // 19 identical fields
  }
```

## Changed report: event

```
  &dc_recipe.Report{
  	Name: "event",
  	Desc: strings.Join({
  		"This report shows a",
- 		"n a",
  		"ctivity logs",
- 		" with",
  		" mostly compatible with Dropbox for teams'",
- 		"s",
  		" activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `dropbox team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "event",
- 	Title:   "Event log",
+ 	Title:   "Retrieve detailed team activity event logs with filtering options, essential for security auditing and compliance monitoring",
  	Desc:    "From release 91, the command parses `-start-time` or `-end-time`"...,
  	Remarks: "",
  	... // 20 identical fields
  }
```

## Changed report: event

```
  &dc_recipe.Report{
  	Name: "event",
  	Desc: strings.Join({
  		"This report shows a",
- 		"n a",
  		"ctivity logs",
- 		" with",
  		" mostly compatible with Dropbox for teams'",
- 		"s",
  		" activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `dropbox team activity user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "user",
- 	Title:   "Activities log per user",
+ 	Title:   "Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities",
- 	Desc:    "",
+ 	Desc:    "Retrieves detailed activity logs for individual team members, including file operations, sharing activities, and login events. Essential for user-specific audits, investigating security incidents, or understanding individual usage patterns. Can filter by a"...,
  	Remarks: "",
  	Path:    "dropbox team activity user",
  	... // 19 identical fields
  }
```

## Changed report: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: strings.Join({
  		"This report shows a",
- 		"n a",
  		"ctivity logs",
- 		" with",
  		" mostly compatible with Dropbox for teams'",
- 		"s",
  		" activity logs.",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `dropbox team admin group role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
  		"A",
- 		"dd the role to members of the group",
+ 		"ssign admin roles to all members of a specified group, streamlin",
+ 		"ing role management for large teams",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Efficiently grants admin privileges to entire groups rather than individual members. Ideal for departmental admin assignments or when onboarding new admin teams. Changes are applied immediately to all current group members.",
  	Remarks: "",
  	Path:    "dropbox team admin group role add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team admin group role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"Delete the role from all members except of members of the except",
- 		"ion group",
+ 		"Remove admin roles from all team members except those in a speci",
+ 		"fied exception group, useful for role cleanup and access control",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk removes specific admin roles while preserving them for an exception group. Useful for reorganizing admin structures or implementing least-privilege access. The exception group ensures critical admins retain necessary permissions during cleanup operati"...,
  	Remarks: "",
  	Path:    "dropbox team admin group role delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team admin list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List admin roles of members",
+ 	Title:   "Display all team members with their assigned admin roles, helpful for auditing administrative access and permissions",
- 	Desc:    "",
+ 	Desc:    "Generates a comprehensive admin audit report showing all members with elevated privileges. Can include non-admin members for complete visibility. Essential for security reviews, compliance audits, and ensuring appropriate access levels across the organizat"...,
  	Remarks: "",
  	Path:    "dropbox team admin list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team admin role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Add a new role to the member",
+ 	Title:   "Grant a specific admin role to an individual team member, enabling granular permission management",
- 	Desc:    "",
+ 	Desc:    "Assigns specific admin roles to individual members for precise permission control. Use when promoting team members to admin positions or adjusting responsibilities. The command validates that the member doesn't already have the specified role to prevent du"...,
  	Remarks: "",
  	Path:    "dropbox team admin role add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team admin role clear`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "clear",
  	Title: strings.Join({
  		"Re",
- 		"move all admin roles from the member",
+ 		"voke all administrative privileges from a team member, useful fo",
+ 		"r role transitions or security purposes",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Completely removes all admin roles from a member in a single operation. Essential for offboarding admins, responding to security incidents, or transitioning members to non-administrative positions. More efficient than removing roles individually.",
  	Remarks: "",
  	Path:    "dropbox team admin role clear",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team admin role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
- 	Title:   "Remove a role from the member",
+ 	Title:   "Remove a specific admin role from a team member while preserving other roles, allowing precise permission adjustments",
- 	Desc:    "",
+ 	Desc:    "Selectively removes individual admin roles without affecting other permissions. Useful for adjusting responsibilities or implementing role-based access changes. The command verifies the member has the role before attempting removal.",
  	Remarks: "",
  	Path:    "dropbox team admin role delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team admin role list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List admin roles of the team",
+ 	Title:   "Display all available admin roles in the team with their descriptions and permissions",
- 	Desc:    "",
+ 	Desc:    "Lists all possible admin roles available in your Dropbox team along with their capabilities. Reference this before assigning roles to understand permission implications. Helps ensure team members receive appropriate access levels.",
  	Remarks: "",
  	Path:    "dropbox team admin role list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team backup device status`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "status",
  	Title: strings.Join({
- 		"Dropbox Backup device status change in the specified period",
+ 		"Track Dropbox Backup status changes for all team devices over a ",
+ 		"specified period, monitoring backup health and compliance",
  	}, ""),
  	Desc: strings.Join({
  		... // 339 identical bytes
  		"d.\n* If the Dropbox application has not been unlinked (e.g. you ",
  		"initialized the OS without unlinking the Dropbox application).\n\n",
- 		"i",
+ 		"I",
  		"n that case, please refer to the report `session_info_updated` t",
  		"o see the most recent report. This command does not automaticall",
  		... // 114 identical bytes
  	}, ""),
  	Remarks: "",
  	Path:    "dropbox team backup device status",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper count`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "count",
  	Title: strings.Join({
  		"C",
- 		"ount number of Paper documents per member",
+ 		"alculate the total number of legacy Paper documents owned by eac",
+ 		"h team member, useful for content auditing and migration plannin",
+ 		"g",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Provides Paper document counts per member, distinguishing between created and accessed documents. Essential for planning Paper-to-Dropbox migrations, identifying heavy Paper users, and estimating migration scope. Filter options help focus on relevant docum"...,
  	Remarks: "",
  	Path:    "dropbox team content legacypaper count",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "export",
  	Title: strings.Join({
  		"Export ",
- 		"entire team member Paper documents into local path",
+ 		"all legacy Paper documents from team members to local storage in",
+ 		" HTML or Markdown format for backup or migration",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk exports team Paper documents to local storage, preserving content before migrations or for compliance archives. Supports HTML and Markdown formats. Creates organized folder structure by member. Consider available disk space as this may export large am"...,
  	Remarks: "",
  	Path:    "dropbox team content legacypaper export",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List team member Paper documents",
+ 		"Generate a comprehensive list of all legacy Paper documents acro",
+ 		"ss the team with ownership and metadata information",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Creates detailed inventory of all Paper documents including titles, owners, and last modified dates. Use for content audits, identifying orphaned documents, or preparing for migrations. Filter by creation or access patterns to focus analysis.",
  	Remarks: "",
  	Path:    "dropbox team content legacypaper list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team content member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List team folder & shared folder member",
+ 		"Display all members with access to team folders and shared folde",
+ 		"rs, showing permission levels and folder relationship",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Maps folder access across the team, showing which members can access specific folders and their permission levels. Invaluable for access reviews, identifying over-privileged accounts, and understanding content exposure. Helps maintain principle of least pr"...,
  	Remarks: "",
  	Path:    "dropbox team content member list",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."},
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{
  			Name: "MemberTypeExternal",
  			Desc: strings.Join({
  				"Filter folder members. Keep only members ",
+ 				"that ",
  				"are external (not in the same team). Note: Invited members are m",
  				"arked as external member.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "MemberTypeInternal",
  			Desc: strings.Join({
  				"Filter folder members. Keep only members ",
+ 				"that ",
  				"are internal (in the same team). Note: Invited members are marke",
  				"d as external member.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	Name: "size",
  	Title: strings.Join({
  		"C",
- 		"ount number of members of team folders and shared folder",
+ 		"alculate member counts for each team folder and shared folder, h",
+ 		"elping identify heavily accessed content and optimize permission",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Analyzes folder membership density to identify over-shared content. High member counts may indicate security risks or performance issues. Use to prioritize permission reviews and identify candidates for access restriction or folder restructuring.",
  	Remarks: "",
  	Path:    "dropbox team content member size",
  	... // 19 identical fields
  }
```

## Changed report: member_count

```
  &dc_recipe.Report{
  	Name: "member_count",
  	Desc: "Folder member count",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "has_no_inherit", Desc: "True if the folder or any sub-folder does not inherit the access"...},
  		&{Name: "is_no_inherit", Desc: "True if the folder does not inherit the access from the parent f"...},
  		&{
  			Name: "capacity",
  			Desc: strings.Join({
  				"Capacity number ",
- 				"to add",
+ 				"for adding",
  				" members. Empty if it's not able to determine by your permission",
  				" (e.g. a folder contains an external group).",
  			}, ""),
  		},
  		&{Name: "count_total", Desc: "Total number of members"},
  		&{
  			Name: "count_external_groups",
- 			Desc: "Number of external teams' group",
+ 			Desc: "Number of external teams' groups",
  		},
  	},
  }
```
# Command spec changed: `dropbox team content mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List all mounted/unmounted shared folders of team members.",
+ 		"Display mount status of all shared folders for team members, ide",
+ 		"ntifying which folders are actively synced to member devices",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Shows which shared folders are actively syncing to member devices versus cloud-only access. Critical for bandwidth planning, identifying heavy sync users, and troubleshooting sync issues. Helps optimize storage usage on user devices.",
  	Remarks: "",
  	Path:    "dropbox team content mount list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team content policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List policies of team folders and shared folders in the team",
+ 		"Review all access policies and restrictions applied to team fold",
+ 		"ers and shared folders for governance compliance",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Comprehensive policy audit showing viewer info restrictions, shared link policies, and other governance settings. Essential for compliance verification and ensuring folders meet organizational security requirements. Identifies policy inconsistencies across"...,
  	Remarks: "",
  	Path:    "dropbox team content policy list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team device list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List all devices/sessions in the team",
+ 		"Display all devices and active sessions connected to team member",
+ 		" accounts with device details and last activity timestamps",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Complete device inventory showing all connected devices, platforms, and session ages. Critical for security audits, identifying unauthorized devices, and managing device limits. Export data to track device sprawl and plan security policies.",
  	Remarks: "",
  	Path:    "dropbox team device list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team device unlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "unlink",
- 	Title:   "Unlink device sessions",
+ 	Title:   "Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access",
- 	Desc:    "",
+ 	Desc:    "Immediately terminates device sessions, forcing re-authentication. Critical security tool for lost devices, departing employees, or suspicious activity. Device must reconnect and re-sync after unlinking. Consider member communication before bulk unlinking.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team device unlink",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "feature",
- 	Title:   "Team feature",
+ 	Title:   "Display all features and capabilities enabled for your Dropbox team account, including API limits and special features",
- 	Desc:    "",
+ 	Desc:    "Shows team's enabled features, beta access, and API rate limits. Check before using advanced features or planning integrations. Features may vary by subscription level. Useful for troubleshooting feature availability issues.",
  	Remarks: "",
  	Path:    "dropbox team feature",
  	... // 19 identical fields
  }
```

## Changed report: feature

```
  &dc_recipe.Report{
  	Name: "feature",
  	Desc: "Team feature",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "upload_api_rate_limit", Desc: "The number of upload API calls allowed per month."},
  		&{
  			Name: "upload_api_rate_limit_count",
  			Desc: strings.Join({
  				"The number of upload API call",
- 				"ed",
+ 				"s made",
  				" this month.",
  			}, ""),
  		},
  		&{Name: "has_team_shared_dropbox", Desc: "Does this team have a shared team root."},
  		&{Name: "has_team_file_events", Desc: "Team supports file events"},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `dropbox team filerequest clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "clone",
  	Title: strings.Join({
- 		"Clone file requests by given data",
+ 		"Duplicate existing file requests with customized settings, usefu",
+ 		"l for creating similar requests across team members",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Creates new file requests based on existing templates with modified settings. Streamlines standardized collection processes like monthly reports or recurring submissions. Preserves folder structure while allowing customization per recipient.",
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "dropbox team filerequest clone",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List all file requests in the team",
+ 		"Display all active and closed file requests created by team memb",
+ 		"ers, helping track external file collection activities",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Comprehensive view of all file requests across the team. Monitor external data collection, identify abandoned requests, and ensure compliance with data handling policies. Includes request URLs, creators, and submission counts for audit purposes.",
  	Remarks: "",
  	Path:    "dropbox team filerequest list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team filesystem`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "filesystem",
  	Title: strings.Join({
  		"Identify ",
+ 		"whe",
  		"t",
- 		"eam's file system version",
+ 		"her your team uses legacy or modern file system architecture, im",
+ 		"portant for feature compatibility",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Determines underlying file system version affecting feature availability and API behavior. Modern file system enables advanced features like native Paper and enhanced performance. Legacy teams may need migration for full feature access.",
  	Remarks: "",
  	Path:    "dropbox team filesystem",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Create new group",
+ 	Title:   "Create a new group in your team for organizing members and managing permissions collectively",
- 	Desc:    "",
+ 	Desc:    "Creates groups for logical organization of team members. Groups simplify permission management by allowing bulk operations. Consider naming conventions for easy identification. Groups can be company-managed or member-managed depending on governance needs.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Bulk adding groups",
+ 	Title:   "Create multiple groups at once using batch processing, efficient for large-scale team organization",
- 	Desc:    "",
+ 	Desc:    "Bulk creates groups from a data file, ideal for initial setup or reorganizations. Validates all groups before creation to prevent partial failures. Include external IDs for integration with identity management systems. Significantly faster than individual "...,
  	Remarks: "",
  	Path:    "dropbox team group batch add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
- 	Title:   "Delete groups",
+ 	Title:   "Remove multiple groups from your team in batch, streamlining group cleanup and reorganization",
- 	Desc:    "",
+ 	Desc:    "Efficiently removes multiple groups in a single operation. Useful for organizational restructuring or cleaning up obsolete groups. Members retain individual permissions but lose group-based access. Verify group contents before deletion as this is irreversi"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group batch delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
- 	Title:   "Clear an external ID of a group",
+ 	Title:   "Remove external ID mappings from groups, useful when disconnecting from external identity providers",
- 	Desc:    "",
+ 	Desc:    "Removes external ID associations from groups when migrating away from identity providers or changing integration systems. Group functionality remains intact but loses external system mapping. Useful for troubleshooting sync issues with identity providers.",
  	Remarks: "",
  	Path:    "dropbox team group clear externalid",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "delete",
- 	Title: "Delete group",
+ 	Title: "Remove a specific group from your team, automatically removing all member associations",
  	Desc: strings.Join({
- 		"This command does not confirm whether the group used in existing",
- 		" folders",
+ 		"Permanently deletes a group and removes all member associations.",
+ 		" Members retain access through other groups or individual permis",
+ 		"sions. Cannot be undone - consider archiving group by removing m",
+ 		"embers instead if unsure. Folder permissions using this group ar",
+ 		"e also removed.",
  	}, ""),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List folders of each group",
+ 	Title:   "Display all folders accessible by each group, showing group-based content organization and permissions",
- 	Desc:    "",
+ 	Desc:    "Maps group permissions to folders, revealing content access patterns. Essential for access reviews and understanding permission inheritance. Helps identify over-permissioned groups and optimize folder structures for security.",
  	Remarks: "",
  	Path:    "dropbox team group folder list",
  	... // 19 identical fields
  }
```

## Changed report: group_to_folder

```
  &dc_recipe.Report{
  	Name: "group_to_folder",
  	Desc: "Group to folder mapping.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
  		&{Name: "group_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
  		&{
  			Name: "group_is_same_team",
  			Desc: strings.Join({
  				"'true' if a group is in ",
+ 				"the ",
  				"same team. Otherwise false.",
  			}, ""),
  		},
  		&{Name: "access_type", Desc: "Group's access level for this folder"},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `dropbox team group list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List group(s)",
+ 	Title:   "Display all groups in your team with member counts and group management types",
- 	Desc:    "",
+ 	Desc:    "Complete inventory of team groups showing sizes and management modes. Use to identify empty groups, oversized groups, or groups needing management type changes. Export for regular auditing and compliance documentation.",
  	Remarks: "",
  	Path:    "dropbox team group list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Add a member to the group",
+ 	Title:   "Add individual team members to a specific group for centralized permission management",
- 	Desc:    "",
+ 	Desc:    "Adds members to groups for inherited permissions and simplified management. Changes take effect immediately for folder access. Consider group size limits and performance implications for very large groups.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group member add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Bulk add members into groups",
+ 	Title:   "Add multiple members to groups efficiently using batch processing, ideal for large team reorganizations",
- 	Desc:    "",
+ 	Desc:    "Bulk adds members to groups using a mapping file. Validates all memberships before applying changes. Ideal for onboarding, departmental changes, or permission standardization projects. Handles errors gracefully with detailed reporting.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group member batch add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
- 	Title:   "Delete members from groups",
+ 	Title:   "Remove multiple members from groups in batch, streamlining group membership management",
- 	Desc:    "",
+ 	Desc:    "Bulk removes members from groups using a CSV file mapping. Validates all memberships before making changes. Useful for organizational restructuring, offboarding processes, or cleaning up group memberships. Processes efficiently with detailed error reportin"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group member batch delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "update",
  	Title: strings.Join({
- 		"Add or delete members from group",
+ 		"Update group memberships in bulk by adding or removing members, ",
+ 		"optimizing group composition change",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Modifies group memberships in bulk based on a CSV file. Can both add and remove members in a single operation. Ideal for large-scale reorganizations where group compositions need significant updates. Maintains audit trail of all changes made.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group member batch update",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
- 	Title:   "Delete a member from the group",
+ 	Title:   "Remove a specific member from a group while preserving their other group memberships",
- 	Desc:    "",
+ 	Desc:    "Removes an individual member from a single group without affecting their membership in other groups. Use for targeted permission adjustments or when members change departments. The removal takes effect immediately, revoking any inherited permissions from t"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group member delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List members of groups",
+ 	Title:   "Display all members belonging to each group, useful for auditing group compositions and access rights",
- 	Desc:    "",
+ 	Desc:    "Lists all groups with their complete member rosters. Essential for access audits, verifying group compositions, and understanding permission inheritance. Helps identify empty groups, over-privileged groups, or members with unexpected access through group m"...,
  	Remarks: "",
  	Path:    "dropbox team group member list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "rename",
- 	Title:   "Rename the group",
+ 	Title:   "Change the name of an existing group to better reflect its purpose or organizational changes",
- 	Desc:    "",
+ 	Desc:    "Updates the display name of a group while maintaining all members and permissions. Useful when departments restructure, projects change names, or group purposes evolve. The rename is immediate and affects all references to the group throughout the system.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team group rename",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team group update type`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "type",
- 	Title:   "Update group management type",
+ 	Title:   "Change how a group is managed (user-managed vs company-managed), affecting who can modify group membership",
- 	Desc:    "",
+ 	Desc:    "Modifies group management settings to control who can add or remove members. Company-managed groups restrict modifications to admins, while user-managed groups allow designated members to manage membership. Critical for implementing proper governance and a"...,
  	Remarks: "",
  	Path:    "dropbox team group update type",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "info",
- 	Title:   "Team information",
+ 	Title:   "Display essential team account information including team ID and basic team settings",
- 	Desc:    "",
+ 	Desc:    "Shows fundamental team account details needed for API integrations and support requests. Team ID is required for various administrative operations. Quick way to verify you're connected to the correct team account.",
  	Remarks: "",
  	Path:    "dropbox team info",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team insight report teamfoldermember`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "teamfoldermember",
- 	Title:   "Report team folder members",
+ 	Title:   "Generate detailed reports on team folder membership, showing access patterns and member distribution",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team insight scan`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "scan",
- 	Title: "Scans team data for analysis",
+ 	Title: "Perform comprehensive data scanning across your team for analytics and insights generation",
  	Desc: strings.Join({
  		... // 194 identical bytes
  		"pbox team insight report teamfoldermember`, or with database too",
  		"ls that support SQLite in general.\n\nAbout how long a scan takes:",
- 		".",
  		"\n\nScanning a team often takes a long time. Especially if there a",
  		"re a large number of files stored, the time is linearly proporti",
  		... // 645 identical bytes
  		" those differences and report exact results, but to provide roug",
  		"h information as quickly as possible.\n\n\nFor database file sizes:",
- 		".",
  		"\n\nAs this command retrieves all metadata, including the team's f",
  		"iles, the size of the database increases with the size of those ",
  		... // 90 identical bytes
  		"files stored in the team. Make sure that the path specified by `",
  		"-database` has enough space before running.\n\n\nAbout scan errors:",
- 		".",
  		"\n\nThe Dropbox server may return an error when running the scan. ",
  		"The command will automatically try to re-run the scan several ti",
  		... // 586 identical bytes
  	}, ""),
  	Remarks: "",
  	Path:    "dropbox team insight scan",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team insight scanretry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "scanretry",
  	Title: strings.Join({
  		"Re",
- 		"try scan for errors on the last scan",
+ 		"-run failed or incomplete scans to ensure complete data collecti",
+ 		"on for team insights",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team insight summarize`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "summarize",
  	Title: strings.Join({
- 		"Summarize team data for analysi",
+ 		"Generate summary reports from scanned team data, providing actio",
+ 		"nable insights on team usage and pattern",
  		"s",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team legalhold add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Creates new legal hold policy.",
+ 	Title:   "Create a legal hold policy to preserve specified team content for compliance or litigation purposes",
  	Desc:    "",
  	Remarks: "",
  	... // 13 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Description", Desc: "A description of the legal hold policy.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "EndDate", Desc: "End date of the legal hold policy.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{
  			Name: "Member",
  			Desc: strings.Join({
- 				"e",
+ 				"E",
  				"mail of the member or members you want to place a hold on",
  			}, ""),
  			Default:  "",
  			TypeName: "infra.feed.fd_file_impl.row_feed",
  			TypeAttr: nil,
  		},
  		&{Name: "Name", Desc: "Policy name.", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{Name: "StartDate", Desc: "Start date of the legal hold policy.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team legalhold list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "Retrieve existing policies",
+ 	Title:   "Display all active legal hold policies with their details, members, and preservation status",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team legalhold member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "update",
  	Title: strings.Join({
- 		"Update member list of legal hold policy",
+ 		"Add or remove multiple team members from legal hold policies in ",
+ 		"batch for efficient compliance management",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team legalhold member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List members of the legal hold",
+ 	Title:   "Display all team members currently under legal hold policies with their preservation status",
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team legalhold release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "release",
- 	Title:   "Releases a legal hold by Id",
+ 	Title:   "Release a legal hold policy and restore normal file operations for affected members and content",
- 	Desc:    "",
+ 	Desc:    "Ends a legal hold policy and removes preservation requirements. Content becomes subject to normal retention and deletion policies again. Use when litigation concludes or preservation is no longer required. The release is logged for audit purposes but canno"...,
  	Remarks: "",
  	Path:    "dropbox team legalhold release",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team legalhold revision list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List revisions under legal hold",
+ 	Title:   "Display all file revisions preserved under legal hold policies, ensuring comprehensive data retention",
- 	Desc:    "",
+ 	Desc:    "Shows the complete revision history of files under legal hold including all modifications. Tracks file versions preserved by the policy to ensure nothing is lost. Critical for maintaining defensible preservation records and demonstrating compliance with le"...,
  	Remarks: "",
  	Path:    "dropbox team legalhold revision list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team legalhold update desc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "desc",
  	Title: strings.Join({
- 		"Updat",
+ 		"Modify th",
  		"e description of ",
- 		"the legal hold policy",
+ 		"an existing legal hold policy to reflect changes in scope or pur",
+ 		"pose",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Updates the description field of a legal hold policy for better documentation. Useful for adding case references, updating matter details, or clarifying preservation scope. Changes are tracked in the revision history for audit purposes.",
  	Remarks: "",
  	Path:    "dropbox team legalhold update desc",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team legalhold update name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "name",
  	Title: strings.Join({
- 		"Update name of the legal hold policy",
+ 		"Change the name of a legal hold policy for better identification",
+ 		" and organization",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `dropbox team linkedapp list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List linked applications",
+ 	Title:   "Display all third-party applications linked to team member accounts for security auditing and access control",
- 	Desc:    "",
+ 	Desc:    "Lists all third-party applications with access to team members' Dropbox accounts. Essential for security audits, identifying unauthorized apps, and managing OAuth integrations. Shows which members use which apps, helping enforce application policies and id"...,
  	Remarks: "",
  	Path:    "dropbox team linkedapp list",
  	... // 19 identical fields
  }
```

## Changed report: linked_app

```
  &dc_recipe.Report{
  	Name: "linked_app",
  	Desc: strings.Join({
  		"This report shows a list of linked app",
+ 		"s",
  		" with the user",
+ 		"s",
  		" of the app",
+ 		"s",
  		".",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "app_name", Desc: "The application name."},
  		&{Name: "is_app_folder", Desc: "Whether the linked application uses a dedicated folder."},
  		&{
  			Name: "publisher",
- 			Desc: "The publisher's URL.",
+ 			Desc: "The application publisher name.",
  		},
  		&{
  			Name: "publisher_url",
- 			Desc: "The application publisher name.",
+ 			Desc: "The publisher's URL.",
  		},
  		&{Name: "linked", Desc: "The time this application was linked"},
  	},
  }
```
# Command spec changed: `dropbox team member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
- 	Title:   "Delete members",
+ 	Title:   "Remove multiple team members in batch, efficiently managing team departures and access revocation",
- 	Desc:    "",
+ 	Desc:    "Bulk removes team members while preserving their data through transfers. Requires specifying destination member for file transfers and admin notification email. Ideal for layoffs, department closures, or mass offboarding. Optionally wipes data from linked "...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member batch delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member batch detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "detach",
  	Title: strings.Join({
  		"Convert ",
- 		"Dropbox for teams accounts to a Basic account",
+ 		"multiple team accounts to individual Basic accounts, preserving ",
+ 		"personal data while removing team access",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk converts team members to personal Dropbox Basic accounts. Members retain their files but lose team features and shared folder access. Useful for contractors ending engagements or when downsizing teams. Consider data retention policies before detaching.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member batch detach",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name: "RevokeTeamShares",
  			Desc: strings.Join({
  				"True ",
- 				"for",
+ 				"to",
  				" revoke shared folder access",
- 				" which",
  				" owned by the team",
  			}, ""),
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
# Command spec changed: `dropbox team member batch invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "invite",
- 	Title:   "Invite member(s)",
+ 	Title:   "Send batch invitations to new team members, streamlining the onboarding process for multiple users",
- 	Desc:    "",
+ 	Desc:    "Sends team invitations to multiple email addresses from a CSV file. Supports silent invites for SSO environments. Ideal for onboarding new departments, acquisitions, or seasonal workers. Validates email formats and checks for existing members before sending.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member batch invite",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member batch reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "reinvite",
  	Title: strings.Join({
  		"Re",
- 		"invite invited status members to the team",
+ 		"send invitations to pending members who haven't joined yet, ensu",
+ 		"ring all intended members receive access",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Resends invitations to all members with pending status. Useful when initial invites expire, get lost in spam, or after resolving email delivery issues. Can send silently for SSO environments. Helps ensure complete team onboarding.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member batch reinvite",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member batch suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "suspend",
- 	Title:   "Bulk suspend members",
+ 	Title:   "Temporarily suspend multiple team members' access while preserving their data and settings",
- 	Desc:    "",
+ 	Desc:    "Bulk suspends team members, blocking access while preserving all data and settings. Use for extended leaves, security investigations, or temporary access restrictions. Option to keep or remove data from devices. Members can be unsuspended later with full a"...,
  	Remarks: "",
  	Path:    "dropbox team member batch suspend",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member batch unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "unsuspend",
- 	Title:   "Bulk unsuspend members",
+ 	Title:   "Restore access for multiple suspended team members, reactivating their accounts in batch",
- 	Desc:    "",
+ 	Desc:    "Bulk reactivates suspended team members, restoring full access to their accounts and data. Use when members return from leave, investigations conclude, or access restrictions lift. All previous permissions and group memberships are restored automatically.",
  	Remarks: "",
  	Path:    "dropbox team member batch unsuspend",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
- 	Title:   "Clear external_id of members",
+ 	Title:   "Remove external ID mappings from team members, useful when disconnecting from identity management systems",
- 	Desc:    "",
+ 	Desc:    "Bulk removes external IDs from team members listed in a CSV file. Essential when migrating between identity providers, cleaning up after SCIM disconnection, or resolving ID conflicts. Does not affect member access, only removes the external identifier mapp"...,
  	Remarks: "",
  	Path:    "dropbox team member clear externalid",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "feature",
- 	Title:   "List member feature settings",
+ 	Title:   "Display feature settings and capabilities enabled for specific team members, helping understand member permissions",
- 	Desc:    "",
+ 	Desc:    "Shows which features and capabilities are enabled for team members. Useful for troubleshooting access issues, verifying feature rollouts, and understanding member capabilities. Helps identify why certain members can or cannot access specific functionality.",
  	Remarks: "",
  	Path:    "dropbox team member feature",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "release",
  	Title: strings.Join({
  		"Release all ",
+ 		"fi",
  		"l",
- 		"ocks under the path of the member",
+ 		"e locks held by a team member under a specified path, resolving ",
+ 		"editing conflicts",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk releases all file locks held by a member within a specified folder path. Essential when members leave unexpectedly or during system issues. Processes in batches for efficiency. Consider notifying affected users as their unsaved changes in locked files"...,
  	Remarks: "",
  	Path:    "dropbox team member file lock all release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team member file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List locks of the member under the path",
+ 		"Display all files locked by a specific team member under a given",
+ 		" path, identifying potential collaboration blocks",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Lists all files currently locked by a specific member within a path. Helps identify collaboration bottlenecks, troubleshoot editing conflicts, and audit file access patterns. Useful for understanding why team members cannot edit certain files.",
  	Remarks: "",
  	Path:    "dropbox team member file lock list",
  	... // 19 identical fields
  }
```

## Changed report: lock

```
  &dc_recipe.Report{
  	Name: "lock",
  	Desc: "Lock information",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team member file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "release",
  	Title: strings.Join({
  		"Release ",
- 		"the lock of the path as the member",
+ 		"a specific file lock held by a team member, enabling others to e",
+ 		"dit the file",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Releases a single file lock held by a member, allowing others to edit. Use when specific files are blocking team collaboration or when lock holders are unavailable. More precise than bulk release when only specific files need unlocking.",
  	Remarks: "",
  	Path:    "dropbox team member file lock release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team member file permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "permdelete",
  	Title: strings.Join({
  		"Permanently delete ",
- 		"the file or folder at a given path of the team member.",
+ 		"files or folders from a team member's account, bypassing trash f",
+ 		"or immediate removal",
  	}, ""),
  	Desc: strings.Join({
  		"P",
- 		"lease see https://www.dropbox.com/help/40 for more detail about ",
- 		"permanent deletion",
+ 		"ermanently deletes specified files or folders without possibilit",
+ 		"y of recovery. Use with extreme caution for removing sensitive d",
+ 		"ata, complying with data retention policies, or freeing storage.",
+ 		" Cannot be undone - ensure proper authorization before use",
  		".",
  	}, ""),
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "dropbox team member file permdelete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List folders for each member",
+ 	Title:   "Display all folders in each team member's account, useful for content auditing and storage analysis",
- 	Desc:    "",
+ 	Desc:    "Enumerates folders across team members' personal spaces. Filter by folder name to focus results. Essential for understanding content distribution, auditing member storage, and planning migrations or cleanups.",
  	Remarks: "",
  	Path:    "dropbox team member folder list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member folder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "replication",
  	Title: strings.Join({
- 		"Replicate a folder to another member's personal folder",
+ 		"Copy folder contents from one team member to another's personal ",
+ 		"space, facilitating content transfer and backup",
  	}, ""),
  	Desc: strings.Join({
- 		"This command will replicate files under the source folder to the",
- 		" destination folder. The source folder can be a source member's ",
- 		"personal folder, a shared folder, or a team folder. But that mus",
- 		"t be mounted and accessible. This command will overwrite a file ",
- 		"if the file already exists on the destination path. \nThis comman",
- 		"d is the one-way copy from source path in a source member, to de",
- 		"stination path in destination member. That means the command wil",
- 		"l not delete the file on the destination path, which deleted on ",
- 		"the source path",
+ 		"Copies complete folder hierarchies between members' personal spa",
+ 		"ces, preserving structure. Ideal for creating backups, transitio",
+ 		"ning responsibilities, or setting up new members with standard f",
+ 		"older structures. Monitor available storage before large replica",
+ 		"tions",
  		".",
  	}, ""),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member folder replication",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List team member(s)",
+ 	Title:   "Display comprehensive list of all team members with their status, roles, and account details",
- 	Desc:    "",
+ 	Desc:    "Provides complete team roster including active, suspended, and optionally deleted members. Shows email addresses, names, roles, and account status. Fundamental for team audits, license management, and understanding team composition. Export for HR or compli"...,
  	Remarks: "",
  	Path:    "dropbox team member list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member quota batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "update",
- 	Title:   "Update team member quota",
+ 	Title:   "Modify storage quotas for multiple team members in batch, managing storage allocation efficiently",
- 	Desc:    "",
+ 	Desc:    "Bulk updates storage quotas for team members using a CSV file. Set custom quotas based on roles, departments, or usage patterns. Use 0 to remove custom quotas and revert to team defaults. Essential for storage governance and cost management.",
  	Remarks: "",
  	Path:    "dropbox team member quota batch update",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member quota list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List team member quota",
+ 	Title:   "Display storage quota assignments for all team members, helping monitor and plan storage distribution",
- 	Desc:    "",
+ 	Desc:    "Shows current storage quota settings for all team members, distinguishing between default and custom quotas. Identifies members with special storage needs or restrictions. Use for capacity planning and ensuring fair storage distribution across teams.",
  	Remarks: "",
  	Path:    "dropbox team member quota list",
  	... // 19 identical fields
  }
```

## Changed report: member_quota

```
  &dc_recipe.Report{
  	Name: "member_quota",
  	Desc: strings.Join({
  		"This report shows a list of custom quota settings for each team ",
  		"member",
- 		"s",
  		".",
  	}, ""),
  	Columns: {&{Name: "email", Desc: "Email address of user."}, &{Name: "quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom"...}},
  }
```
# Command spec changed: `dropbox team member quota usage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "usage",
- 	Title:   "List team member storage usage",
+ 	Title:   "Show actual storage usage for each team member compared to their quotas, identifying storage needs",
- 	Desc:    "",
+ 	Desc:    "Displays current storage consumption versus allocated quotas for each member. Highlights members approaching limits, underutilizing space, or needing quota adjustments. Critical for proactive storage management and preventing work disruptions due to full q"...,
  	Remarks: "",
  	Path:    "dropbox team member quota usage",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "replication",
- 	Title:   "Replicate team member files",
+ 	Title:   "Replicate all files from one team member's account to another, useful for account transitions or backups",
- 	Desc:    "",
+ 	Desc:    "Creates complete copies of member data between accounts, preserving folder structures and sharing where possible. Essential for role transitions, creating backups, or merging accounts. Requires sufficient storage in destination account. Consider using batc"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member replication",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "suspend",
- 	Title:   "Suspend a member",
+ 	Title:   "Temporarily suspend a team member's access to their account while preserving all data and settings",
- 	Desc:    "",
+ 	Desc:    "Immediately blocks member access while maintaining all data, settings, and group memberships. Use for security incidents, policy violations, or temporary leaves. Choose whether to keep data on linked devices. Member can be unsuspended later with full acces"...,
  	Remarks: "",
  	Path:    "dropbox team member suspend",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "unsuspend",
- 	Title:   "Unsuspend a member",
+ 	Title:   "Restore access for a suspended team member, reactivating their account and all associated permissions",
- 	Desc:    "",
+ 	Desc:    "Reactivates a suspended member's account, restoring full access to data and team resources. All previous permissions, group memberships, and settings are preserved. Use when suspension reasons are resolved or members return from leave.",
  	Remarks: "",
  	Path:    "dropbox team member unsuspend",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member update batch email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "email",
- 	Title:   "Member email operation",
+ 	Title:   "Update email addresses for multiple team members in batch, managing email changes efficiently",
- 	Desc:    "",
+ 	Desc:    "Bulk updates member email addresses using a CSV mapping file. Essential for domain migrations, name changes, or correcting email errors. Validates new addresses and preserves all member data and permissions. Option to update unverified emails with caution.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member update batch email",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{
  			Name: "UpdateUnverified",
  			Desc: strings.Join({
  				"Update an account which ",
- 				"didn't verified email. If an account email unverified, email add",
- 				"ress change may affect lose",
+ 				"hasn't verified its email. If an account email is unverified, ch",
+ 				"anging the email address may cause loss of",
  				" invitation to folders.",
  			}, ""),
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
# Command spec changed: `dropbox team member update batch externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "externalid",
  	Title: strings.Join({
- 		"Update External ID of team member",
+ 		"Set or update external IDs for multiple team members, integratin",
+ 		"g with identity management system",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Maps external identity system IDs to Dropbox team members in bulk. Critical for SCIM integration, SSO setup, or syncing with HR systems. Ensures consistent identity mapping across platforms. Updates existing IDs or sets new ones as needed.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member update batch externalid",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member update batch invisible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "invisible",
  	Title: strings.Join({
- 		"Enable directory restriction to membe",
+ 		"Hide team members from the directory listing, enhancing privacy ",
+ 		"for sensitive roles or contracto",
  		"rs",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk hides members from team directory searches and listings. Useful for executives, security personnel, or external contractors who need access but shouldn't appear in directories. Hidden members retain all functionality but enhanced privacy.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member update batch invisible",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member update batch profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "profile",
- 	Title:   "Batch update member profiles",
+ 	Title:   "Update profile information for multiple team members including names and job titles in batch",
- 	Desc:    "",
+ 	Desc:    "Bulk updates member profile information including given names and surnames. Ideal for standardizing name formats, correcting widespread errors, or updating after organizational changes. Maintains consistency across team directories and improves searchability.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member update batch profile",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team member update batch visible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "visible",
  	Title: strings.Join({
- 		"Disable directory restriction to member",
+ 		"Make hidden team members visible in the directory, restoring sta",
+ 		"ndard visibility setting",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk restores visibility for previously hidden members in team directories. Use when privacy requirements change, contractors become employees, or to correct visibility errors. Members become searchable and appear in team listings again.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team member update batch visible",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team namespace file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List all files and folders of the team namespace",
+ 		"Display comprehensive file and folder listings within team names",
+ 		"paces for content inventory and analysi",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Lists all files and folders within team namespaces with filtering options. Include or exclude deleted items, member folders, shared folders, and team folders. Essential for content audits, migration planning, and understanding data distribution across name"...,
  	Remarks: "",
  	Path:    "dropbox team namespace file list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "size",
  	Title: strings.Join({
- 		"List all files and folders of the team namespace",
+ 		"Calculate storage usage for files and folders in team namespaces",
+ 		", providing detailed size analytic",
  		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Analyzes storage consumption across team namespaces with configurable depth scanning. Shows size distribution by namespace type (team, shared, member, app folders). Critical for storage optimization, identifying large folders, and planning archival strateg"...,
  	Remarks: "",
  	Path:    "dropbox team namespace file size",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...},
  		&{Name: "IncludeAppFolder", Desc: "If true, include app folders", Default: "false", TypeName: "bool", ...},
  		&{
  			Name: "IncludeMemberFolder",
  			Desc: strings.Join({
- 				"i",
+ 				"I",
  				"f true, include team member folders",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team namespace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List all namespaces of the team",
+ 	Title:   "Display all team namespaces including team folders and shared spaces with their configurations",
- 	Desc:    "",
+ 	Desc:    "Enumerates all namespace types in the team including ownership, paths, and access levels. Provides comprehensive view of team's folder architecture. Use for understanding organizational structure, planning migrations, or auditing folder governance.",
  	Remarks: "",
  	Path:    "dropbox team namespace list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team namespace member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List members of shared folders and team folders in the team",
+ 		"Show all members with access to each namespace, detailing permis",
+ 		"sions and access levels",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Maps namespace access showing which members can access which folders and their permission levels. Reveals access patterns, over-privileged namespaces, and helps ensure appropriate access controls. Essential for security audits and access reviews.",
  	Remarks: "",
  	Path:    "dropbox team namespace member list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team namespace summary`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "summary",
  	Title: strings.Join({
- 		"Report team namespace status summary.",
+ 		"Generate comprehensive summary reports of team namespace usage, ",
+ 		"member counts, and storage statistics",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Aggregates namespace data to show overall team structure, storage distribution, and access patterns. Provides high-level insights into how team content is organized across different namespace types. Useful for capacity planning and organizational assessments.",
  	Remarks: "",
  	Path:    "dropbox team namespace summary",
  	... // 19 identical fields
  }
```

## Changed report: folder_without_parent

```
  &dc_recipe.Report{
  	Name: "folder_without_parent",
  	Desc: "Folders without parent folder.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 10 identical elements
  	},
  }
```
# Command spec changed: `dropbox team report activity`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "activity",
- 	Title:   "Activities report",
+ 	Title:   "Generate detailed activity reports covering all team operations, useful for compliance and usage analysis",
- 	Desc:    "",
+ 	Desc:    "Creates comprehensive activity reports showing team member actions, file operations, sharing events, and administrative changes. Customizable date ranges and filters. Essential for security monitoring, compliance reporting, and understanding team collabora"...,
  	Remarks: "",
  	Path:    "dropbox team report activity",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team report devices`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "devices",
- 	Title:   "Devices report",
+ 	Title:   "Create comprehensive device usage reports showing all connected devices, platforms, and access patterns",
- 	Desc:    "",
+ 	Desc:    "Shows all devices connected to team accounts including type, OS, sync status, and last activity. Critical for security audits, identifying unauthorized devices, and managing device policies. Helps enforce security standards and investigate access anomalies.",
  	Remarks: "",
  	Path:    "dropbox team report devices",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team report membership`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "membership",
- 	Title:   "Membership report",
+ 	Title:   "Generate team membership reports including member status, roles, and account statistics over time",
- 	Desc:    "",
+ 	Desc:    "Provides membership analytics including active users, growth trends, and role distributions. Track team expansion, monitor license usage, and identify inactive accounts. Useful for budget planning and optimizing team size.",
  	Remarks: "",
  	Path:    "dropbox team report membership",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team report storage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "storage",
- 	Title:   "Storage report",
+ 	Title:   "Create detailed storage usage reports showing team consumption, trends, and member distribution",
- 	Desc:    "",
+ 	Desc:    "Provides comprehensive storage analytics including total usage, growth trends, and per-member consumption. Identifies storage hogs, helps predict future needs, and supports capacity planning. Export data for budgeting and resource allocation decisions.",
  	Remarks: "",
  	Path:    "dropbox team report storage",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team runas file batch copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "copy",
  	Title: strings.Join({
- 		"Batch copy files/folders as a member",
+ 		"Copy multiple files or folders on behalf of team members, useful",
+ 		" for content management and organization",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin tool to copy files between member accounts without their credentials. Useful for distributing templates, recovering deleted content, or setting up new members. Operates with admin privileges while maintaining audit trails. Requires appropriate admin "...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team runas file batch copy",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team runas file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
  		"List files and folders ",
- 		"run as a",
+ 		"in a team member's account by running operations as that",
  		" member",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Allows admins to view file listings in member accounts without member credentials. Essential for investigating issues, auditing content, or helping members locate files. All actions are logged for security. Use responsibly and follow privacy policies.",
  	Remarks: "",
  	Path:    "dropbox team runas file list",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...},
  		&{
  			Name: "IncludeExplicitSharedMembers",
  			Desc: strings.Join({
- 				" ",
  				"If true, the results will include a flag for each file indicatin",
  				"g whether or not that file has any explicit members.",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "IncludeMountedFolders",
  			Desc: strings.Join({
- 				" ",
  				"If true, the results will include entries under mounted folders ",
  				"which include",
- 				"s",
  				" app folder, shared folder and team folder.",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 2 identical elements
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
- 		"Batch upstream sync with Dropbox",
+ 		"Upload multiple local files to team members' Dropbox accounts in",
+ 		" batch, running as those members",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin bulk upload tool for distributing files to multiple member accounts simultaneously. Ideal for deploying templates, policies, or required documents. Maintains consistent file distribution across teams. All uploads are tracked for compliance.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team runas file sync batch up",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BasePath", Desc: "Choose the file path standard. This is an option for Dropbox for"..., Default: "root", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "BatchSize", Desc: "Batch commit size", Default: "250", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name: "Delete",
  			Desc: strings.Join({
  				"Delete Dropbox file if a file ",
+ 				"is ",
  				"removed locally",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "ExitOnFailure", Desc: "Exit the program on failure", Default: "false", TypeName: "bool", ...},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		... // 6 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
  		&{Name: "num_folder_created", Desc: "Number of created folders."},
  		&{
  			Name: "num_delete",
- 			Desc: "Number of deleted entry.",
+ 			Desc: "Number of deleted entries.",
  		},
  		&{
  			Name: "num_api_call",
  			Desc: strings.Join({
  				"The number of estimated ",
- 				"upload API call",
+ 				"API calls",
  				" for upload.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "leave",
- 	Title:   "Leave shared folders in batch",
+ 	Title:   "Remove team members from multiple shared folders in batch by running leave operations as those members",
- 	Desc:    "",
+ 	Desc:    "Admin tool to remove members from multiple shared folders without their interaction. Useful for access cleanup, security responses, or organizational changes. Operates as the member would, maintaining proper audit trails. Cannot remove folder owners.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder batch leave",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "result.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "result.parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		&{
  			Name: "result.name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "result.access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "result.path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 10 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "share",
- 	Title:   "Share shared folders in batch",
+ 	Title:   "Share multiple folders on behalf of team members in batch, automating folder sharing processes",
- 	Desc:    "",
+ 	Desc:    "Admin batch tool for creating shared folders on behalf of members. Streamlines folder sharing for new projects or team reorganizations. Sets appropriate permissions and sends invitations. All sharing actions are logged for security compliance.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder batch share",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "result.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "result.parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		&{
  			Name: "result.name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "result.access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "result.path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 10 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "unshare",
- 	Title:   "Unshare shared folders in batch",
+ 	Title:   "Remove sharing from multiple folders on behalf of team members, managing folder access in bulk",
- 	Desc:    "",
+ 	Desc:    "Admin tool to revoke folder sharing in bulk for security or compliance. Removes sharing while preserving folder contents for the owner. Critical for incident response or preventing data leaks. All unshare actions create audit records.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder batch unshare",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "result.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "result.parent_shared_folder_id", Desc: "The ID of the parent shared folder. This field is present only i"...},
  		&{
  			Name: "result.name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "result.access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "result.path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 10 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder isolate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "isolate",
  	Title: strings.Join({
- 		"Isolate member from shared folder",
+ 		"Remove all shared folder access for a team member and transfer o",
+ 		"wnership, useful for departing employees",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Emergency admin action to remove all members from a shared folder except its owner. Use for security incidents, data breaches, or when folder content needs immediate access restriction. Preserves folder structure while eliminating external access risks.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team runas sharedfolder isolate",
  	... // 19 identical fields
  }
```

## Changed report: isolated

```
  &dc_recipe.Report{
  	Name: "isolated",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "input.name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "input.access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "input.path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List shared folders",
+ 	Title:   "Display all shared folders accessible by a team member, running the operation as that member",
- 	Desc:    "",
+ 	Desc:    "Admin view of member's shared folder access including permission levels and folder details. Essential for access audits, investigating over-sharing, or troubleshooting permission issues. Helps ensure appropriate access levels and identify security risks.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder list",
  	... // 19 identical fields
  }
```

## Changed report: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
  		"Add m",
+ 		"ultipl",
  		"e",
+ 		" me",
  		"mbers to shared folders in batch",
+ 		" on behalf of folder owners, streamlining access management",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin tool to bulk add members to specific shared folders with defined permissions. Efficient for project kickoffs, team expansions, or access standardization. Validates member emails and permissions before applying changes. Creates comprehensive audit trail.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder member batch add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
  		"Remove m",
+ 		"ultipl",
  		"e",
+ 		" me",
  		"mbers from shared folders ",
+ 		"in batch on behalf of folder owners, managing access effic",
  		"i",
- 		"n batch",
+ 		"ently",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin bulk removal of members from shared folders for security or reorganization. Preserves folder content while revoking access for specified members. Essential for quick security responses or access cleanup. Cannot remove folder owner.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder member batch delete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
  		"Mount ",
+ 		"sh",
  		"a",
+ 		"red",
  		" ",
- 		"shared folder as another member",
+ 		"folders to team members' accounts on their behalf, ensuring prop",
+ 		"er folder synchronization",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin action to mount shared folders in member accounts when they cannot do it themselves. Useful for troubleshooting sync issues, helping non-technical users, or ensuring critical folders are properly mounted. Operates as if the member performed the action.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount add",
  	... // 19 identical fields
  }
```

## Changed report: mount

```
  &dc_recipe.Report{
  	Name: "mount",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"The specified user unmounts the designated folder.",
+ 		"Unmount shared folders from team members' accounts on their beha",
+ 		"lf, managing folder synchronization",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin tool to unmount shared folders from member accounts without removing access. Useful for troubleshooting sync issues, managing local storage, or temporarily removing folders from sync. Member retains access and can remount later.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount delete",
  	... // 19 identical fields
  }
```

## Changed report: mount

```
  &dc_recipe.Report{
  	Name: "mount",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"List",
+ 		"Display",
  		" all shared folders ",
- 		"the specified member mounted",
+ 		"currently mounted (synced) to a specific team member's account",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Admin view of which shared folders are actively mounted (syncing) in a member's account. Helps diagnose sync issues, understand storage usage, or verify proper folder access. Distinguishes between mounted and unmounted but accessible folders.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount list",
  	... // 19 identical fields
  }
```

## Changed report: mounts

```
  &dc_recipe.Report{
  	Name: "mounts",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "mountable",
  	Title: strings.Join({
- 		"List all shared folders the member can moun",
+ 		"Show all available shared folders that a team member can mount b",
+ 		"ut hasn't synced ye",
  		"t",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Lists shared folders accessible to a member but not currently synced to their device. Useful for identifying available folders, helping members find content, or understanding why certain folders aren't appearing locally. Shows potential sync options.",
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount mountable",
  	... // 19 identical fields
  }
```

## Changed report: mountables

```
  &dc_recipe.Report{
  	Name: "mountables",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{
  			Name: "name",
- 			Desc: "The name of the this shared folder.",
+ 			Desc: "The name of this shared folder.",
  		},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 9 identical elements
  	},
  }
```
# Command spec changed: `dropbox team sharedlink cap expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "expiry",
  	Title: strings.Join({
- 		"Set expiry cap to shared links in the team",
+ 		"Apply expiration date limits to all team shared links for enhanc",
+ 		"ed security and compliance",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Applies expiration dates to existing shared links without them. Essential for security compliance and reducing exposure of perpetual links. Can target links by age or apply blanket expiration policy. Helps prevent unauthorized long-term access to shared co"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink cap expiry",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team sharedlink cap visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "visibility",
  	Title: strings.Join({
- 		"Set visibility cap to shared links in the team",
+ 		"Enforce visibility restrictions on team shared links, controllin",
+ 		"g public access levels",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Modifies shared link visibility settings to enforce team security policies. Can restrict public links to team-only or password-protected access. Critical for preventing data leaks and ensuring links comply with organizational security requirements.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink cap visibility",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team sharedlink delete links`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "links",
- 	Title:   "Batch delete shared links",
+ 	Title:   "Delete multiple shared links in batch for security compliance or access control cleanup",
- 	Desc:    "",
+ 	Desc:    "Bulk deletes shared links based on criteria like age, visibility, or path patterns. Use for security remediation, removing obsolete links, or enforcing new sharing policies. Permanent action that immediately revokes access through deleted links.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink delete links",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team sharedlink delete member`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "member",
  	Title: strings.Join({
- 		"Delet",
+ 		"Remov",
  		"e all shared links ",
- 		"of the member",
+ 		"created by a specific team member, useful for departing employee",
+ 		"s",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Removes all shared links created by a specific member, regardless of content location. Essential for secure offboarding, responding to compromised accounts, or enforcing immediate access revocation. Cannot be undone, so use with appropriate authorization.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink delete member",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List of shared links",
+ 	Title:   "Display comprehensive list of all shared links created by team members with visibility and expiration details",
- 	Desc:    "",
+ 	Desc:    "Comprehensive inventory of all team shared links showing URLs, visibility settings, expiration dates, and creators. Essential for security audits, identifying risky links, and understanding external sharing patterns. Filter by various criteria for focused "...,
  	Remarks: "",
  	Path:    "dropbox team sharedlink list",
  	... // 19 identical fields
  }
```

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
  	Desc: "This report shows a list of shared links with the shared link ow"...,
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "expires", Desc: "Expiration time, if set."},
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `dropbox team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "expiry",
  	Title: strings.Join({
- 		"Update expiration date of public shared links within the team",
+ 		"Modify expiration dates for existing shared links across the tea",
+ 		"m to enforce security policies",
  	}, ""),
  	Desc: (
  		"""
- 		Note: From Release 87, this command will receive a file to select shared links to update. If you wanted to update the expiry for all shared links in the team, please consider using a combination of `dropbox team sharedlink list`. For example, if you are familiar with the command [jq](https://stedolan.github.io/jq/), then you can do an equivalent operation as like below (force expiry within 28 days for every public link).
- 		
- 		```
- 		tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -at +720h
- 		```
- 		From Release 92, the command will not receive the argument `-days`. If you want to set a relative date/time, please use `-at +HOURh` like `+720h` (720 hours = 30 days).
- 		
- 		Commands `dropbox team sharedlink update` is for setting a value to the shared links. Commands `dropbox team sharedlink cap` is for putting a cap value to the shared links. For example: if you set expiry by `dropbox team sharedlink update expiry` with the expiration date 2021-05-06. The command will update the expiry to 2021-05-06 even if the existing link has a shorter expiration date like 2021-05-04.
+ 		Modifies expiration dates for existing shared links to enforce new security policies or extend access for legitimate use cases. Can target specific links or apply bulk updates. Helps maintain balance between security and usability.
  		"""
  	),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink update expiry",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team sharedlink update password`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "password",
  	Title: strings.Join({
- 		"Set or update shared link passwords",
+ 		"Add or change passwords on team shared links in batch for enhanc",
+ 		"ed security protection",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Applies password protection to existing shared links or updates current passwords. Critical for securing sensitive content shared externally. Can target vulnerable links or apply passwords based on content sensitivity. Notify link recipients of new require"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink update password",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team sharedlink update visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "visibility",
  	Title: strings.Join({
- 		"Update visibility of shared links",
+ 		"Change access levels of existing shared links between public, te",
+ 		"am-only, and password-protected",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Updates shared link visibility from public to team-only or other restricted settings. Essential for reducing external exposure and meeting compliance requirements. Can target links by current visibility level or content location. Changes take effect immedi"...,
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team sharedlink update visibility",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
  		&{Name: "result.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{
  			Name: "result.visibility",
  			Desc: strings.Join({
  				"The current visibility of the link after considering the shared ",
  				"links policies of the t",
- 				"he t",
  				"eam (in case the link's owner is part of a team) and the shared ",
  				"folder (in case the linked file is part of a shared folder).",
  			}, ""),
  		},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.surname", Desc: "Surname of the link owner"},
  		&{Name: "result.given_name", Desc: "Given name of the link owner"},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
- 	Title:   "Add team folder to the team",
+ 	Title:   "Create a new team folder for centralized team content storage and collaboration",
- 	Desc:    "",
+ 	Desc:    "Creates new team folders with defined access controls and sync settings. Set up departmental folders, project spaces, or archive locations. Configure initial permissions and determine whether content syncs to member devices by default.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "archive",
- 	Title:   "Archive team folder",
+ 	Title:   "Archive a team folder to make it read-only while preserving all content and access history",
- 	Desc:    "",
+ 	Desc:    "Converts active team folders to archived status, making them read-only while preserving all content and permissions. Use for completed projects, historical records, or compliance requirements. Archived folders can be reactivated if needed.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder archive",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder batch archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "archive",
- 	Title:   "Archiving team folders",
+ 	Title:   "Archive multiple team folders in batch, efficiently managing folder lifecycle and compliance",
- 	Desc:    "",
+ 	Desc:    "Bulk archives team folders based on criteria like age, name patterns, or activity levels. Streamlines folder lifecycle management and helps maintain organized team spaces. Preserves all content while preventing new modifications.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder batch archive",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder batch permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "permdelete",
- 	Title:   "Permanently delete team folders",
+ 	Title:   "Permanently delete multiple archived team folders in batch, freeing storage space",
- 	Desc:    "",
+ 	Desc:    "Permanently deletes multiple team folders and all their contents without possibility of recovery. Use only with proper authorization for removing obsolete data, complying with retention policies, or emergency cleanup. This action cannot be undone.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder batch permdelete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder batch replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "replication",
  	Title: strings.Join({
- 		"Batch replication of team folders",
+ 		"Replicate multiple team folders to another team account in batch",
+ 		" for migration or backup",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Creates copies of multiple team folders with their complete structures and permissions. Useful for creating backups, setting up parallel environments, or preparing for migrations. Consider storage implications before large replications.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder batch replication",
  	... // 19 identical fields
  }
```

## Changed report: verification

```
  &dc_recipe.Report{
  	Name: "verification",
  	Desc: strings.Join({
  		"This report shows a difference between t",
+ 		"w",
  		"o folders.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "left_hash", Desc: "Content hash of left file"},
  		&{Name: "right_path", Desc: "path of right"},
  		&{
  			Name: "right_kind",
- 			Desc: "folder of file",
+ 			Desc: "folder or file",
  		},
  		&{Name: "right_size", Desc: "size of right file"},
  		&{Name: "right_hash", Desc: "Content hash of right file"},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List files in team folders",
+ 	Title:   "Display all files and subfolders within team folders for content inventory and management",
- 	Desc:    "",
+ 	Desc:    "Enumerates all files in team folders with details like size, modification dates, and paths. Essential for content audits, migration planning, and understanding data distribution. Can filter by file types or patterns for targeted analysis.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder file list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "release",
  	Title: strings.Join({
  		"Release all ",
+ 		"fi",
  		"l",
- 		"ocks under the path of the team folder",
+ 		"e locks within a team folder path, resolving editing conflicts i",
+ 		"n bulk",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Bulk releases all file locks within specified team folders. Use when multiple locks are blocking team productivity or after system issues. Notifies lock holders when possible. May cause loss of unsaved changes in locked files.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder file lock all release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List locks in the team folder",
+ 	Title:   "Display all locked files within team folders, identifying collaboration bottlenecks",
- 	Desc:    "",
+ 	Desc:    "Lists all currently locked files in team folders with lock holder information and lock duration. Helps identify collaboration bottlenecks, stale locks, and users who may need assistance. Essential for maintaining smooth team workflows.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder file lock list",
  	... // 19 identical fields
  }
```

## Changed report: lock

```
  &dc_recipe.Report{
  	Name: "lock",
  	Desc: "Lock information",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 6 identical elements
  		&{Name: "is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "release",
  	Title: strings.Join({
  		"Release ",
- 		"lock of the path in the team folder",
+ 		"specific file locks in team folders to enable collaborative edit",
+ 		"ing",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Releases individual file locks in team folders when specific files are blocking work. More precise than bulk release when only certain files need unlocking. Useful for resolving urgent editing conflicts without affecting other locked files.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder file lock release",
  	... // 19 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "result.is_lock_holder", Desc: "True if caller holds the file lock"},
  		&{Name: "result.lock_holder_name", Desc: "The display name of the lock holder."},
  		&{
  			Name: "result.lock_created",
  			Desc: strings.Join({
  				"The timestamp ",
- 				"of",
+ 				"when",
  				" the lock was created.",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "size",
- 	Title:   "Calculate size of team folders",
+ 	Title:   "Calculate storage usage for team folders, providing detailed size analytics for capacity planning",
- 	Desc:    "",
+ 	Desc:    "Analyzes storage consumption within team folders showing size distribution and largest files. Essential for capacity planning, identifying candidates for archival, and understanding storage costs. Helps optimize team folder usage and plan for growth.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder file size",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Depth", Desc: "Depth", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name: "FolderName",
  			Desc: strings.Join({
  				"List only fo",
+ 				"lde",
  				"r",
+ 				"s",
  				" ",
+ 				"ma",
  				"t",
- 				"he folder matched to",
+ 				"ching",
  				" the name. Filter by exact match to the name.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "FolderNamePrefix",
  			Desc: strings.Join({
  				"List only fo",
+ 				"lde",
  				"r",
+ 				"s",
  				" ",
+ 				"ma",
  				"t",
- 				"he folder matched to",
+ 				"ching",
  				" the name. Filter by name match to the prefix.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "FolderNameSuffix",
  			Desc: strings.Join({
  				"List only fo",
+ 				"lde",
  				"r",
+ 				"s",
  				" ",
+ 				"ma",
  				"t",
- 				"he folder matched to",
+ 				"ching",
  				" the name. Filter by name match to the suffix.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List team folder(s)",
+ 	Title:   "Display all team folders with their status, sync settings, and member access information",
- 	Desc:    "",
+ 	Desc:    "Comprehensive list of all team folders showing names, status (active/archived), sync settings, and access levels. Fundamental for team folder governance, planning reorganizations, and understanding team structure. Export for documentation or analysis.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
- 		"Batch adding users/groups to team folders",
+ 		"Add multiple users or groups to team folders in batch, streamlin",
+ 		"ing access provisioning",
  	}, ""),
  	Desc: (
  		"""
- 		This command will do (1) create new team folders or new sub-folders if the team folder does not exist. The command does not (2) change access inheritance setting of any folders, (3) create a group if that not exist. This command is designed to be idempotent. You can safely retry if any errors happen on the operation. The command will not report an error to keep idempotence. For example, the command will not report an error like, the member already have access to the folder.
- 		
- 		Example:
- 		
- 		* Sales (team folder, editor access for the group "Sales")
- 			* Sydney (viewer access for individual account sydney@example.com)
- 			* Tokyo (editor access for the group "Tokyo Deal Desk")
- 				* Monthly (viewer access for individual account success@example.com)
- 		* Marketing (team folder, editor access for the group "Marketing")
- 			* Sydney (editor access for the group "Sydney Sales")
- 			* Tokyo (viewer access for the group "Tokyo Sales")
- 		
- 		1. Prepare CSV like below
- 		
- 		```
- 		Sales,,editor,Sales
- 		Sales,Sydney,editor,sydney@example.com
- 		Sales,Tokyo,editor,Tokyo Deal Desk
- 		Sales,Tokyo/Monthly,viewer,success@example.com
- 		Marketing,,editor,Marketing
- 		Marketing,Sydney,editor,Sydney Sales
- 		Marketing,Tokyo,viewer,Tokyo Sales
- 		```
- 		
- 		2. Then run the command like below
- 		
- 		```
- 		tbx teamfolder member add -file /PATH/TO/DATA.csv
- 		```
- 		
- 		Note: the command will create a team folder if not exist. But the command will not a group if not found. Groups must exist before run this command.
+ 		Grants access to team folders for individuals or groups with defined permissions (view/edit). Use for onboarding, project assignments, or expanding access. Group additions efficiently manage permissions through group membership rather than individual assignments.
  		"""
  	),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder member add",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"Batch removing users/groups from team folders",
+ 		"Remove multiple users or groups from team folders in batch, mana",
+ 		"ging access revocation efficiently",
  	}, ""),
  	Desc: strings.Join({
- 		"The command does not (1) change access inheritance setting of an",
- 		"y folders, (2) remove a group, (3) unshare a nested folder. For ",
- 		"(3), that means the nested folder stays the same setting (e.g. s",
- 		"hared link policy for the folder). This command is designed to b",
- 		"e idempotent. You can safely retry if any errors happen on the o",
- 		"peration. The command will not report an error to keep idempoten",
- 		"ce. For example, the command will not report an error like, (1) ",
- 		"the member already lose access to the folder, (2) the folder is ",
- 		"not found",
+ 		"Revokes team folder access for specific members or entire groups",
+ 		". Essential for offboarding, project completion, or security res",
+ 		"ponses. Removal is immediate and affects all folder contents. Co",
+ 		"nsider data retention needs before removing members with edit ac",
+ 		"cess",
  		".",
  	}, ""),
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder member delete",
  	... // 19 identical fields
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
  		&{Name: "input.team_folder_name", Desc: "Team folder name"},
  		&{
  			Name: "input.path",
  			Desc: strings.Join({
  				"Relative path from the team folder root. Leave empty if you want",
  				" to ",
- 				"add a member to",
+ 				"remove a member from the",
  				" root of the team folder.",
  			}, ""),
  		},
  		&{Name: "input.group_name_or_member_email", Desc: "Group name or member email address"},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List team folder members",
+ 	Title:   "Display all members with access to each team folder, showing permission levels and access types",
- 	Desc:    "",
+ 	Desc:    "Shows complete membership for all team folders including permission levels and whether access is direct or through groups. Critical for access audits, security reviews, and understanding who can access sensitive content. Identifies over-privileged access.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder member list",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."},
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{
  			Name: "MemberTypeExternal",
  			Desc: strings.Join({
  				"Filter folder members. Keep only members ",
+ 				"that ",
  				"are external (not in the same team). Note: Invited members are m",
  				"arked as external member.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "MemberTypeInternal",
  			Desc: strings.Join({
  				"Filter folder members. Keep only members ",
+ 				"that ",
  				"are internal (in the same team). Note: Invited members are marke",
  				"d as external member.",
  			}, ""),
  			Default:  "",
  			TypeName: "",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
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
  	Name: "replication",
  	Title: strings.Join({
- 		"Partial team folder replication to the other team",
+ 		"Selectively replicate team folder contents to another team, enab",
+ 		"ling flexible content migration",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Copies selected subfolders or files from team folders rather than entire structures. Useful for creating targeted backups, extracting project deliverables, or migrating specific content. More efficient than full replication when only portions are needed.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder partial replication",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "permdelete",
- 	Title:   "Permanently delete team folder",
+ 	Title:   "Permanently delete an archived team folder and all its contents, irreversibly freeing storage",
- 	Desc:    "",
+ 	Desc:    "Irreversibly deletes a team folder and all contained files. Use only with proper authorization and after confirming no critical data remains. Essential for compliance with data retention policies or removing sensitive content. This action cannot be undone.",
  	Remarks: "(Irreversible operation)",
  	Path:    "dropbox team teamfolder permdelete",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List policies of team folders",
+ 	Title:   "Display all access policies and restrictions applied to team folders for governance review",
- 	Desc:    "",
+ 	Desc:    "Shows all policies governing team folder behavior including sync defaults, sharing restrictions, and access controls. Helps understand why folders behave certain ways and ensures policy compliance. Reference before creating new folders or modifying settings.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder policy list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "replication",
  	Title: strings.Join({
- 		"Replicate a team folder to the other team",
+ 		"Copy an entire team folder with all contents to another team acc",
+ 		"ount for migration or backup",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Creates an exact duplicate of a team folder preserving structure, permissions, and content. Use for creating backups, setting up test environments, or preparing for major changes. Consider available storage and replication time for large folders.",
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "dropbox team teamfolder replication",
  	... // 19 identical fields
  }
```

## Changed report: verification

```
  &dc_recipe.Report{
  	Name: "verification",
  	Desc: strings.Join({
  		"This report shows a difference between t",
+ 		"w",
  		"o folders.",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		... // 4 identical elements
  		&{Name: "left_hash", Desc: "Content hash of left file"},
  		&{Name: "right_path", Desc: "path of right"},
  		&{
  			Name: "right_kind",
- 			Desc: "folder of file",
+ 			Desc: "folder or file",
  		},
  		&{Name: "right_size", Desc: "size of right file"},
  		&{Name: "right_hash", Desc: "Content hash of right file"},
  	},
  }
```
# Command spec changed: `dropbox team teamfolder sync setting list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "List team folder sync settings",
+ 	Title:   "Display sync configuration for all team folders, showing default sync behavior for members",
- 	Desc:    "",
+ 	Desc:    "Shows current sync settings for all team folders indicating whether they automatically sync to new members' devices. Helps understand bandwidth impact, storage requirements, and ensures appropriate content distribution policies.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder sync setting list",
  	... // 19 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder sync setting update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "update",
  	Title: strings.Join({
- 		"Batch update team folder sync settings",
+ 		"Modify sync settings for multiple team folders in batch, control",
+ 		"ling automatic synchronization behavior",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Modifies sync behavior for team folders between automatic sync to all members or manual sync selection. Use to reduce storage usage on devices, manage bandwidth, or ensure critical folders sync automatically. Apply changes during low-activity periods.",
  	Remarks: "",
  	Path:    "dropbox team teamfolder sync setting update",
  	... // 19 identical fields
  }
```
# Command spec changed: `figma file info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name: "info",
  	Title: strings.Join({
  		"Show information of the ",
- 		"f",
+ 		"F",
  		"igma file",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `figma file list`



## Changed report: files

```
  &dc_recipe.Report{
  	Name: "files",
  	Desc: "Figma file",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "key", Desc: "Figma file key"},
  		&{
  			Name: "name",
- 			Desc: "Name fo the document",
+ 			Desc: "Name of the document",
  		},
  		&{Name: "thumbnailUrl", Desc: "Thumbnail URL"},
  		&{Name: "lastModified", Desc: "Last modified timestamp"},
  	},
  }
```
# Command spec changed: `github profile`



## Changed report: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: "GitHub user profile",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "login", Desc: "Login user name"},
  		&{Name: "name", Desc: "Name of the user"},
  		&{
  			Name: "url",
- 			Desc: "Url of the user",
+ 			Desc: "URL of the user",
  		},
  	},
  }
```
# Command spec changed: `github release draft`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "BodyFile",
  			Desc: strings.Join({
  				"File path to body text. T",
- 				"He file must",
+ 				"he file must be",
  				" encoded in UTF-8 without BOM.",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]any{"shouldExist": bool(false)},
  		},
  		&{Name: "Branch", Desc: "Name of the target branch", TypeName: "string"},
  		&{Name: "Name", Desc: "Name of the release", TypeName: "string"},
  		... // 4 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `license`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "license",
  	Title:   "Show license information",
- 	Desc:    "",
+ 	Desc:    "Display detailed license information for the watermint toolbox and all its components. This includes open source licenses, copyright notices, and third-party dependencies used in the application.",
  	Remarks: "",
  	Path:    "license",
  	... // 19 identical fields
  }
```
# Command spec changed: `log api job`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "job",
  	Title:   "Show statistics of the API log of the job specified by the job ID",
- 	Desc:    "",
+ 	Desc:    "Analyze and display API call statistics for a specific job execution. This includes request counts, response times, error rates, and endpoint usage patterns. Useful for performance analysis, debugging API issues, and understanding application behavior duri"...,
  	Remarks: "",
  	Path:    "log api job",
  	... // 19 identical fields
  }
```
# Command spec changed: `log api name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "name",
  	Title:   "Show statistics of the API log of the job specified by the job name",
- 	Desc:    "",
+ 	Desc:    "Analyze and display API call statistics for jobs identified by command name rather than job ID. This allows you to aggregate statistics across multiple executions of the same command, helping identify patterns and performance trends over time.",
  	Remarks: "",
  	Path:    "log api name",
  	... // 19 identical fields
  }
```
# Command spec changed: `log cat curl`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "curl",
  	Title:   "Format capture logs as `curl` sample",
- 	Desc:    "",
+ 	Desc:    "Convert API request logs into equivalent curl commands that can be executed independently. This is extremely useful for debugging API issues, reproducing requests outside of the toolbox, sharing examples with support, or creating test scripts.",
  	Remarks: "",
  	Path:    "log cat curl",
  	... // 19 identical fields
  }
```
# Command spec changed: `log cat job`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "job",
  	Title:   "Retrieve logs of specified Job ID",
- 	Desc:    "",
+ 	Desc:    "Extract and display log files for a specific job execution identified by its Job ID. This includes debug logs, API capture logs, error messages, and system information. Essential for troubleshooting failed executions and analyzing job performance.",
  	Remarks: "",
  	Path:    "log cat job",
  	... // 19 identical fields
  }
```
# Command spec changed: `util encode base64`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 20 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
  	TextInput: []*dc_recipe.DocTextInput{
  		&{
  			Name: "Text",
- 			Desc: "Text to decode",
+ 			Desc: "Text to encode",
  		},
  	},
  	JsonInput: {},
  }
```
# Command spec changed: `util git clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "LocalPath", Desc: "Local path to clone repository", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Reference", Desc: "Reference name", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "RemoteName", Desc: "Name of the remote", Default: "origin", TypeName: "string", ...},
  		&{
  			Name:     "Url",
- 			Desc:     "Git repository url",
+ 			Desc:     "Git repository URL",
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
# Command spec changed: `util image placeholder`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Color", Desc: "Background color", Default: "white", TypeName: "string", ...},
  		&{
  			Name: "FontPath",
  			Desc: strings.Join({
  				"Path to True",
- 				" ",
  				"Type font (required if you need to draw",
- 				" a",
  				" text)",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "FontSize", Desc: "Font size", Default: "12", TypeName: "int", ...},
  		&{
  			Name:     "Height",
- 			Desc:     "Height (pixel)",
+ 			Desc:     "Height (pixels)",
  			Default:  "400",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path to export generated image", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Text", Desc: "Text if you need", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "TextAlign", Desc: "Text alignment", Default: "left", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "TextColor", Desc: "Text color", Default: "black", TypeName: "string", ...},
  		&{Name: "TextPosition", Desc: "Text position", Default: "center", TypeName: "string", ...},
  		&{
  			Name:     "Width",
- 			Desc:     "Width (pixel)",
+ 			Desc:     "Width (pixels)",
  			Default:  "640",
  			TypeName: "int",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util json query`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Compact", Desc: "Compact output", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "File path", TypeName: "Path"},
  		&{
  			Name:     "Query",
- 			Desc:     "Query string. ",
+ 			Desc:     "Query string",
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
# Command spec changed: `util qrcode create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ErrorCorrectionLevel", Desc: "Error correction level (l/m/q/h).", Default: "m", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Mode", Desc: "QR code encoding mode", Default: "auto", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Out", Desc: "Output path with file name", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			Name:     "Size",
- 			Desc:     "Image resolution (pixel)",
+ 			Desc:     "Image resolution (pixels)",
  			Default:  "256",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{"max": float64(32767), "min": float64(25), "value": float64(256)},
  		},
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
  		&{Name: "ErrorCorrectionLevel", Desc: "Error correction level (l/m/q/h).", Default: "m", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{
  			Name: "Hidden",
  			Desc: strings.Join({
  				"`true` if a",
+ 				"n",
  				" SSID is hidden. `false` if a",
+ 				"n",
  				" SSID is visible.",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_string.select_string_internal",
  			TypeAttr: map[string]any{"options": []any{string(""), string("true"), string("false")}},
  		},
  		&{Name: "Mode", Desc: "QR code encoding mode", Default: "auto", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "NetworkType", Desc: "Network type.", Default: "WPA", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Out", Desc: "Output path with file name", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			Name:     "Size",
- 			Desc:     "Image resolution (pixel)",
+ 			Desc:     "Image resolution (pixels)",
  			Default:  "256",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{"max": float64(32767), "min": float64(25), "value": float64(256)},
  		},
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
  		&{
  			Name: "AcceptLicenseAgreement",
  			Desc: strings.Join({
  				"Accept t",
- 				"o t",
  				"he target release's license agreement",
  			}, ""),
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path to install", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_public", ...},
  		&{Name: "Release", Desc: "Release tag name", Default: "latest", TypeName: "string", ...},
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
  	Name: "wakati",
  	Title: strings.Join({
  		"Waka",
- 		"ti ",
+ 		"chi",
  		"gaki (tokenize Japanese text)",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
# Command spec changed: `util tidy move simple`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "Dst",
  			Desc: strings.Join({
  				"The destination folder path. The command will create folders if ",
+ 				"they do ",
  				"not exist on the path.",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]any{"shouldExist": bool(false)},
  		},
  		&{Name: "ExcludeFolders", Desc: "Exclude folders", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeSystemFiles", Desc: "Include system files", Default: "false", TypeName: "bool", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util xlsx sheet import`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Create", Desc: "Create a file if not found", Default: "false", TypeName: "bool", ...},
  		&{Name: "Data", Desc: "Data path", TypeName: "Data"},
  		&{Name: "File", Desc: "Path to data file", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			Name: "Position",
  			Desc: strings.Join({
  				"Start position to import in A1 notation. Default",
+ 				":",
  				" `A1`.",
  			}, ""),
  			Default:  "A1",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "Sheet", Desc: "Sheet name", TypeName: "string"},
  	},
  	GridDataInput:  {&{Name: "Data", Desc: "Input data file"}},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util xlsx sheet list`



## Changed report: sheets

```
  &dc_recipe.Report{
  	Name: "sheets",
  	Desc: "Sheet",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "Name of the sheet"},
  		&{Name: "rows", Desc: "Number of rows"},
  		&{Name: "cols", Desc: "Number of columns"},
  		&{
  			Name: "hidden",
  			Desc: strings.Join({
  				"True ",
- 				"when if the sheet",
+ 				"if the sheet is",
  				" marked as hidden",
  			}, ""),
  		},
  	},
  }
```
# Command spec changed: `version`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "version",
  	Title:   "Show version",
- 	Desc:    "",
+ 	Desc:    "Display version information for the watermint toolbox including build date, Git commit hash, and component versions. This is useful for troubleshooting, bug reports, and ensuring you have the latest version.",
  	Remarks: "",
  	Path:    "version",
  	... // 19 identical fields
  }
```
