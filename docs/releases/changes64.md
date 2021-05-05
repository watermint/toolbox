---
layout: release
title: Changes of Release 63
lang: en
---

# Changes between `Release 63` to `Release 64`

# Commands added


| Command                            | Title                                                           |
|------------------------------------|-----------------------------------------------------------------|
| config disable                     | Disable a feature.                                              |
| config enable                      | Enable a feature.                                               |
| config features                    | List available optional features.                               |
| dev ci artifact connect            | Connect to Dropbox for uploading artifact from CI               |
| dev ci auth connect                | Authenticate for generating end to end testing                  |
| dev ci auth export                 | Export auth tokens of end to end test                           |
| dev ci auth import                 | Import auth tokens of end to end test from environment variable |
| file dispatch local                | Dispatch local files                                            |
| services github issue list         | List issues of the public/private GitHub repository             |
| services github profile            | Get the authenticated user                                      |
| services github release asset list | List assets of GitHub Release                                   |
| services github release asset up   | Upload assets file into the GitHub Release                      |
| services github release draft      | Create release draft                                            |
| services github release list       | List releases                                                   |
| services github tag create         | Create a tag on the repository                                  |
| version                            | Show version                                                    |



# Commands deleted


| Command     | Title                                          |
|-------------|------------------------------------------------|
| dev ci auth | Authenticate for generating end to end testing |



# Command spec changed: `connect business_audit`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `connect business_file`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `connect business_info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `connect business_mgmt`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `connect user_file`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev async`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 		&{
+ 			Name:     "RunConcurrently",
+ 			Desc:     "run concurrently",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: rows

```
  &dc_recipe.Report{
  	Name: "rows",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups and their members.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `dev ci artifact up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev ci artifact up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path to upload",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path to upload",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: {&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, ...},
  }
```

## Changed report: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `dev desktop install`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "InstallerUrl",
+ 			Desc:     "Installer download URL",
+ 			Default:  "https://www.dropbox.com/download?full=1&os=win",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Silent", Desc: "Use Silent installer", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "SilentNoLaunch",
+ 			Desc:     "Use Enterprise installer",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev desktop start`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev desktop stop`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "WaitSeconds",
+ 			Desc:     "Try stopping the app after given seconds.",
+ 			Default:  "60",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(60)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev desktop suspendupdate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Unsuspend",
+ 			Desc:     "True when unsuspend Updater",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "UpdaterName",
+ 			Desc:     "Executable name of Dropbox Updater",
+ 			Default:  "DropboxUpdate.exe",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "UpdaterPath",
+ 			Desc:     "Path to Dropbox Updater",
+ 			Default:  "C:/Program Files (x86)/Dropbox/Update",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev diag procmon`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev diag procmon",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -repository-path /LOCAL/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path to upload procmon logs",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "ProcmonUrl",
+ 			Desc:     "Process monitor download url",
+ 			Default:  "https://download.sysinternals.com/files/ProcessMonitor.zip",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepositoryPath",
+ 			Desc:     "Procmon Work directory",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "RetainLogs",
+ 			Desc:     "Number of Procmon logs to retain",
+ 			Default:  "4",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(10000), "min": float64(0), "value": float64(4)},
+ 		},
+ 		&{
+ 			Name:     "RunUntil",
+ 			Desc:     "Skip run after this date/time",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "Duration for waiting procmon",
+ 			Default:  "1800",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(10), "value": float64(1800)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Badge",
+ 			Desc:     "Include badges of build status",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "CommandPath",
+ 			Desc:     "Relative path to generate command manuals",
+ 			Default:  "doc/generated/",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Filename", Desc: "Filename", Default: "README.md", TypeName: "string"},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "MarkdownReadme",
+ 			Desc:     "Generate README as markdown format",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev dummy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Dest", Desc: "Dummy file destination", TypeName: "string"},
+ 		&{
+ 			Name:     "MaxEntry",
+ 			Desc:     "Maximum entries",
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 		&{Name: "Path", Desc: "Path to dummy entry file", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev echo`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{&{Name: "Text", Desc: "Text to echo", TypeName: "string"}},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev preflight`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "Path to the test resource location",
+ 			Default:  "test/dev/resource.json",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release publish`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev release publish",
- 	CliArgs:         "",
+ 	CliArgs:         "-artifact-path /LOCAL/PATH/TO/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ArtifactPath",
+ 			Desc:     "Path to artifacts",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "Branch", Desc: "Target branch", Default: "master", TypeName: "string"},
+ 		&{
+ 			Name:     "ConnGithub",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
+ 		},
+ 		&{
+ 			Name:     "SkipTests",
+ 			Desc:     "Skip end to end tests.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "Path to test resource",
+ 			Default:  "test/dev/resource.json",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev spec diff`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "FilePath",
+ 			Desc:     "File path to output",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Release1",
+ 			Desc:     "Release name 1",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Release2",
+ 			Desc:     "Release name 2",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev spec doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "FilePath",
+ 			Desc:     "File path",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev test monkey`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test monkey",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Distribution",
+ 			Desc:     "Number of files/folder distribution",
+ 			Default:  "10000",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(10000)},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "File extensions (comma separated)",
+ 			Default:  "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"...,
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Monkey test path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "Monkey test duration in seconds",
+ 			Default:  "10",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(1), "value": float64(10)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev test recipe`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "Recipe",
+ 			Desc:     "Recipe name to test",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Resource",
+ 			Desc:     "Test resource file path",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Verbose",
+ 			Desc:     "Verbose output for testing",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev test resources`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev util curl`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "BufferSize",
+ 			Desc:     "Size of buffer",
+ 			Default:  "65536",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.097152e+06), "min": float64(1024), "value": float64(65536)},
+ 		},
+ 		&{Name: "Record", Desc: "Capture record(s) for the test", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev util wait`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "Wait seconds",
+ 			Default:  "1",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(604800), "min": float64(1), "value": float64(1)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `file compare account`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Left",
+ 			Desc:     "Account alias (left)",
+ 			Default:  "left",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "LeftPath",
+ 			Desc:     "The path from account root (left)",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Right",
+ 			Desc:     "Account alias (right)",
+ 			Default:  "right",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "RightPath",
+ 			Desc:     "The path from account root (right)",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: diff

```
  &dc_recipe.Report{
  	Name:    "diff",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: {&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash,"...}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, ...},
  }
```
# Command spec changed: `file compare local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: diff

```
  &dc_recipe.Report{
  	Name:    "diff",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: {&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash,"...}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, ...},
  }
```

## Changed report: skip

```
  &dc_recipe.Report{
  	Name:    "skip",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: {&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash,"...}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, ...},
  }
```
# Command spec changed: `file copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Destination path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Source path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `file delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to delete",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `file download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file download",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/OF/FILE -local-path /LOCAL/PATH/TO/DOWNLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "File path to download",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path to download",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file export doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file export doc",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/FILE -local-path /LOCAL/PATH/TO/EXPORT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox document path to export.",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path to save",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows a result of exporting file.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "If this folder is a shared folder mount point, the ID of the sha"...},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
  		&{Name: "export_name", Desc: "File name for export file."},
  		&{Name: "export_size", Desc: "File size of export file."},
- 		&{Name: "export_hash", Desc: "Content hash of export file."},
  	},
  }
```
# Command spec changed: `file import batch url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to import",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.url", Desc: "Url to download"},
  		&{Name: "input.path", Desc: "Path to store file (use path given by `-path` when the record is"...},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
  		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file import url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to import",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Url", Desc: "URL", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
  		&{Name: "revision", Desc: "A unique identifier for the current revision of a file."},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "Include deleted files",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "Include media information",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: file_list

```
  &dc_recipe.Report{
  	Name: "file_list",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file merge`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool"},
+ 		&{
+ 			Name:     "From",
+ 			Desc:     "Path for merge",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "KeepEmptyFolder",
+ 			Desc:     "Keep empty folder after merge",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "To",
+ 			Desc:     "Path to merge",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "WithinSameNamespace",
+ 			Desc:     "Do not cross namespace. That is for preserve sharing permission "...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Destination path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Source path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
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
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Account alias (destionation)",
+ 			Default:  "dst",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "DstPath",
+ 			Desc:     "Destination path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Account alias (source)",
+ 			Default:  "src",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "SrcPath",
+ 			Desc:     "Source path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: replication_diff

```
  &dc_recipe.Report{
  	Name:    "replication_diff",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: {&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash,"...}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, ...},
  }
```
# Command spec changed: `file restore`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file restore",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.path", Desc: "Path"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
  		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file search content`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Restricts search to only the file categories specified (image/do"...,
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "Restricts search to only the extensions specified.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Scopes the search to a path in the user's Dropbox.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: matches

```
  &dc_recipe.Report{
  	Name:    "matches",
- 	Desc:    "",
+ 	Desc:    "This report shows a result of search with highlighted text.",
  	Columns: {&{Name: "tag", Desc: "Type of entry"}, &{Name: "path_display", Desc: "Display path"}, &{Name: "highlight_html", Desc: "Highlighted text in HTML"}},
  }
```
# Command spec changed: `file search name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Restricts search to only the file categories specified (image/do"...,
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "Restricts search to only the extensions specified.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Scopes the search to a path in the user's Dropbox.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: matches

```
  &dc_recipe.Report{
  	Name:    "matches",
- 	Desc:    "",
+ 	Desc:    "This report shows a result of search with highlighted text.",
  	Columns: {&{Name: "tag", Desc: "Type of entry"}, &{Name: "path_display", Desc: "Display path"}, &{Name: "highlight_html", Desc: "Highlighted text in HTML"}},
  }
```
# Command spec changed: `file sync preflight up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync preflight up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Destination Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: {&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, ...},
  }
```

## Changed report: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ChunkSizeKb",
+ 			Desc:     "Upload chunk size in KB",
+ 			Default:  "153600",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)},
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Destination Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: {&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, ...},
  }
```

## Changed report: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ChunkSizeKb",
+ 			Desc:     "Upload chunk size in KB",
+ 			Default:  "153600",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)},
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Destination Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Overwrite",
+ 			Desc:     "Overwrite existing file(s)",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```

## Changed report: summary

```
  &dc_recipe.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: {&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, ...},
  }
```

## Changed report: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `file watch`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file watch",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/WATCH",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to watch",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "Watch path recursively",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `filerequest create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "filerequest create",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "AllowLateUploads",
+ 			Desc:     "If set, allow uploads after the deadline has passed (one_day/two"...,
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Deadline",
+ 			Desc:     "The deadline for this file request.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "The path for the folder in the Dropbox where uploaded files will"...,
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Title", Desc: "The title of the file request", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: file_request

```
  &dc_recipe.Report{
  	Name:    "file_request",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: {&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, ...},
  }
```
# Command spec changed: `filerequest delete closed`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: deleted

```
  &dc_recipe.Report{
  	Name:    "deleted",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: {&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, ...},
  }
```
# Command spec changed: `filerequest delete url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Force",
+ 			Desc:     "Force delete the file request.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Url", Desc: "URL of the file request.", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: deleted

```
  &dc_recipe.Report{
  	Name:    "deleted",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: {&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, ...},
  }
```
# Command spec changed: `filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: file_requests

```
  &dc_recipe.Report{
  	Name:    "file_requests",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: {&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, ...},
  }
```
# Command spec changed: `group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "ManagementType",
+ 			Desc:     "Group management type `company_managed` or `user_managed`",
+ 			Default:  "company_managed",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 		&{Name: "Name", Desc: "Group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: added_group

```
  &dc_recipe.Report{
  	Name: "added_group",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }
```
# Command spec changed: `group batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for group name list",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.name", Desc: "Group name"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }
```
# Command spec changed: `group delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Name", Desc: "Group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `group list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: group

```
  &dc_recipe.Report{
  	Name: "group",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }
```
# Command spec changed: `group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "GroupName", Desc: "Group name", TypeName: "string"},
+ 		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "Email address of the member"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }
```
# Command spec changed: `group member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "GroupName", Desc: "Name of the group", TypeName: "string"},
+ 		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "Email address of the member"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }
```
# Command spec changed: `group member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: group_member

```
  &dc_recipe.Report{
  	Name: "group_member",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups and their members.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "CurrentName", Desc: "Current group name", TypeName: "string"},
+ 		&{Name: "NewName", Desc: "New group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.new_name", Desc: "New group name"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }
```
# Command spec changed: `job history archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "Target days old",
+ 			Default:  "7",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(7)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `job history delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "Target days old",
+ 			Default:  "28",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(28)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `job history list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: log

```
  &dc_recipe.Report{
  	Name:    "log",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of job histories.",
  	Columns: {&{Name: "job_id", Desc: "Job ID"}, &{Name: "app_version", Desc: "App version"}, &{Name: "recipe_name", Desc: "Command"}, &{Name: "time_start", Desc: "Time Started"}, ...},
  }
```
# Command spec changed: `job history ship`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job history ship",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path to upload",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 11 identical elements
  		&{Name: "result.revision", Desc: "A unique identifier for the current revision of a file."},
  		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `job loop`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job loop",
- 	CliArgs:         "",
+ 	CliArgs:         `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 17:58:38"`,
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IntervalSeconds",
+ 			Desc:     "Interval seconds",
+ 			Default:  "180",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(1), "value": float64(180)},
+ 		},
+ 		&{Name: "QuitOnError", Desc: "Quit on error", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Desc:     "Path to runbook",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Until",
+ 			Desc:     "Run until this date/time",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(false)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `job run`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job run",
- 	CliArgs:         "",
+ 	CliArgs:         "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Fork",
+ 			Desc:     "Fork process on running the workflow",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Desc:     "Path to the runbook",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "TimeoutSeconds",
+ 			Desc:     "Terminate process when given time passed",
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(0), "value": float64(0)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `license`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         nil,
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "WipeData",
+ 			Desc:     "If true, controls if the user's data will be deleted on their li"...,
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: {&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.email", Desc: "Email address of the account"}},
  }
```
# Command spec changed: `member detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "RevokeTeamShares",
+ 			Desc:     "True for revoke shared folder access which owned by the team",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: {&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.email", Desc: "Email address of the account"}},
  }
```
# Command spec changed: `member invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "SilentInvite",
+ 			Desc:     "Do not send welcome email (requires SSO + domain verification in"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "Given name of the account"},
  		&{Name: "input.surname", Desc: "Surname of the account"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: member

```
  &dc_recipe.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "member_folder_id", Desc: "The namespace id of the user's root folder."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  	},
  }
```
# Command spec changed: `member quota list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: member_quota

```
  &dc_recipe.Report{
  	Name:    "member_quota",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of custom quota settings for each team members.",
  	Columns: {&{Name: "email", Desc: "Email address of user."}, &{Name: "quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom"...}},
  }
```
# Command spec changed: `member quota update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "Quota",
+ 			Desc:     "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom q"...,
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: {&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.email", Desc: "Email address of user."}, &{Name: "input.quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom"...}, ...},
  }
```
# Command spec changed: `member quota usage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: usage

```
  &dc_recipe.Report{
  	Name:    "usage",
- 	Desc:    "",
+ 	Desc:    "This report shows current storage usage of users.",
  	Columns: {&{Name: "email", Desc: "Email address of the account"}, &{Name: "used_gb", Desc: "The user's total space usage (in GB, 1GB = 1024 MB)."}, &{Name: "used_bytes", Desc: "The user's total space usage (bytes)."}, &{Name: "allocation", Desc: "The user's space allocation (individual, or team)"}, ...},
  }
```
# Command spec changed: `member reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "Silent",
+ 			Desc:     "Do not send welcome email (SSO required)",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "input.email", Desc: "Email address of user."},
  		&{Name: "input.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "input.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "input.given_name", Desc: "Also known as a first name"},
  		&{Name: "input.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "input.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "input.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "input.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "input.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "input.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "input.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "input.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "input.tag", Desc: "Operation tag"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Destination team; team file access",
+ 			Default:  "dst",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Source team; team file access",
+ 			Default:  "src",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: {&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.src_email", Desc: "Source account's email address"}, &{Name: "input.dst_email", Desc: "Destination account's email address"}},
  }
```
# Command spec changed: `member update email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "UpdateUnverified",
+ 			Desc:     "Update an account which didn't verified email. If an account ema"...,
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.from_email", Desc: "Current Email address"},
  		&{Name: "input.to_email", Desc: "New Email address"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.email", Desc: "Email address of team members"},
  		&{Name: "input.external_id", Desc: "External ID of team members"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "Given name of the account"},
  		&{Name: "input.surname", Desc: "Surname of the account"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
- 	Desc: "",
+ 	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder.",
- 		},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		... // 5 identical elements
  		&{Name: "policy_member", Desc: "Who can be a member of this shared folder, as set on the folder "...},
  		&{Name: "policy_viewer_info", Desc: "Who can enable/disable viewer info for this shared folder."},
- 		&{Name: "owner_team_id", Desc: "Team ID of the team that owns the folder"},
  		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
  	},
  }
```
# Command spec changed: `sharedfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: member

```
  &dc_recipe.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder.",
- 		},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 2 identical elements
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		&{Name: "is_inherited", Desc: "True if the member has access from a parent folder"},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "invitee_email", Desc: "Email address of invitee for this folder"},
  	},
  }
```
# Command spec changed: `sharedlink create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Expires",
+ 			Desc:     "Expiration date/time of the shared link",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Password",
+ 			Desc:     "Password",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "TeamOnly",
+ 			Desc:     "Link is accessible only by team members",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: created

```
  &dc_recipe.Report{
  	Name:    "created",
- 	Desc:    "",
+ 	Desc:    "THis report shows a list of shared links.",
  	Columns: {&{Name: "id", Desc: "A unique identifier for the linked file or folder"}, &{Name: "tag", Desc: "Entry type (file, or folder)"}, &{Name: "url", Desc: "URL of the shared link."}, &{Name: "name", Desc: "The linked file name (including extension)."}, ...},
  }
```
# Command spec changed: `sharedlink delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "File or folder path to remove shared link",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "Attempt to remove the file hierarchy",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.id", Desc: "A unique identifier for the linked file or folder"},
  		&{Name: "input.tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "input.url", Desc: "URL of the shared link."},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `sharedlink file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedlink file list",
- 	CliArgs:         "",
+ 	CliArgs:         "-url SHAREDLINK_URL",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Password",
+ 			Desc:     "Password for the shared link",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Url",
+ 			Desc:     "Shared link URL",
+ 			TypeName: "domain.dropbox.model.mo_url.url_impl",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: file_list

```
  &dc_recipe.Report{
  	Name: "file_list",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }
```
# Command spec changed: `sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "THis report shows a list of shared links.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the linked file or folder"},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `team activity batch user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Filter the returned events to a single category. This field is o"...,
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "Ending time (exclusive).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "User email address list file",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "Starting time (inclusive)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: combined

```
  &dc_recipe.Report{
  	Name:    "combined",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```

## Changed report: user

```
  &dc_recipe.Report{
  	Name:    "user",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `team activity daily event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Event category",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndDate",
+ 			Desc:     "End date",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{Name: "StartDate", Desc: "Start date", TypeName: "string"},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: event

```
  &dc_recipe.Report{
  	Name:    "event",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Filter the returned events to a single category. This field is o"...,
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "Ending time (exclusive).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "Starting time (inclusive)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: event

```
  &dc_recipe.Report{
  	Name:    "event",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```
# Command spec changed: `team activity user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Filter the returned events to a single category. This field is o"...,
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "Ending time (exclusive).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "Starting time (inclusive)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: user

```
  &dc_recipe.Report{
  	Name:    "user",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: {&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, ...},
  }
```

## Changed report: user_summary

```
  &dc_recipe.Report{
  	Name: "user_summary",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.user", Desc: "User email address"},
- 		&{Name: "result.user", Desc: "User email address"},
  		&{Name: "result.logins", Desc: "Number of login activities"},
  		&{Name: "result.devices", Desc: "Number of device activities"},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `team content member`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: membership

```
  &dc_recipe.Report{
  	Name:    "membership",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.",
  	Columns: {&{Name: "path", Desc: "Path"}, &{Name: "folder_type", Desc: "Type of the folder. (`team_folder`: a team folder or in a team f"...}, &{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"}, &{Name: "access_type", Desc: "User's access level for this folder"}, ...},
  }
```

## Changed report: no_member

```
  &dc_recipe.Report{
  	Name:    "no_member",
- 	Desc:    "",
+ 	Desc:    "This report shows folders without members.",
  	Columns: {&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"}, &{Name: "path", Desc: "Path"}, &{Name: "folder_type", Desc: "Type of the folder. (`team_folder`: a team folder or in a team f"...}},
  }
```
# Command spec changed: `team content policy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: policy

```
  &dc_recipe.Report{
  	Name:    "policy",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of shared folders and team folders with their current policy settings.",
  	Columns: {&{Name: "path", Desc: "Path"}, &{Name: "is_team_folder", Desc: "`true` if the folder is a team folder, or inside of a team folder"}, &{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"}, &{Name: "policy_manage_access", Desc: "Who can add and remove members from this shared folder."}, ...},
  }
```
# Command spec changed: `team device list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: device

```
  &dc_recipe.Report{
  	Name: "device",
- 	Desc: "",
+ 	Desc: "This report shows a list of current existing sessions in the team with team member information.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }
```
# Command spec changed: `team device unlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DeleteOnUnlink",
+ 			Desc:     "Delete files on unlink",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "input.given_name", Desc: "Also known as a first name"},
  		&{Name: "input.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "input.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{
- 			Name: "input.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "input.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
  		&{Name: "input.device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "input.id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }
```
# Command spec changed: `team diag explorer`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "All",
+ 			Desc:     "Include additional reports",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Dropbox Business file access",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Info",
+ 			Desc:     "Dropbox Business information access",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 		&{
+ 			Name:     "Mgmt",
+ 			Desc:     "Dropbox Business management",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: device

```
  &dc_recipe.Report{
  	Name: "device",
- 	Desc: "",
+ 	Desc: "This report shows a list of current existing sessions in the team with team member information.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }
```

## Changed report: feature

```
  &dc_recipe.Report{
  	Name:    "feature",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team features and their settings.",
  	Columns: {&{Name: "upload_api_rate_limit", Desc: "The number of upload API calls allowed per month."}, &{Name: "upload_api_rate_limit_count", Desc: "The number of upload API called this month."}, &{Name: "has_team_shared_dropbox", Desc: "Does this team have a shared team root."}, &{Name: "has_team_file_events", Desc: "Does this team have file events."}, ...},
  }
```

## Changed report: file_request

```
  &dc_recipe.Report{
  	Name: "file_request",
- 	Desc: "",
+ 	Desc: "This report shows a list of file requests with the file request owner team member.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "account_id", Desc: "Account ID of this file request owner."},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ID of file request owner user as a member of a team",
- 		},
  		&{Name: "email", Desc: "Email address of this file request owner."},
  		&{Name: "status", Desc: "The user status of this file request owner (active/invited/suspe"...},
  		&{Name: "surname", Desc: "Surname of this file request owner."},
  		&{Name: "given_name", Desc: "Given name of this file request owner."},
- 		&{Name: "file_request_id", Desc: "The ID of the file request."},
  		&{Name: "url", Desc: "The URL of the file request."},
  		&{Name: "title", Desc: "The title of the file request."},
  		... // 6 identical elements
  	},
  }
```

## Changed report: group

```
  &dc_recipe.Report{
  	Name: "group",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }
```

## Changed report: group_member

```
  &dc_recipe.Report{
  	Name: "group_member",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups and their members.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_manage"...},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		... // 2 identical elements
  	},
  }
```

## Changed report: info

```
  &dc_recipe.Report{
  	Name:    "info",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team information.",
  	Columns: {&{Name: "name", Desc: "The name of the team"}, &{Name: "team_id", Desc: "The ID of the team."}, &{Name: "num_licensed_users", Desc: "The number of licenses available to the team."}, &{Name: "num_provisioned_users", Desc: "The number of accounts that have been invited or are already act"...}, ...},
  }
```

## Changed report: linked_app

```
  &dc_recipe.Report{
  	Name: "linked_app",
- 	Desc: "",
+ 	Desc: "This report shows a list of linked app with the user of the app.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "app_id", Desc: "The application unique id."},
  		&{Name: "app_name", Desc: "The application name."},
  		&{Name: "is_app_folder", Desc: "Whether the linked application uses a dedicated folder."},
  		... // 3 identical elements
  	},
  }
```

## Changed report: member

```
  &dc_recipe.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "member_folder_id", Desc: "The namespace id of the user's root folder."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  	},
  }
```

## Changed report: member_quota

```
  &dc_recipe.Report{
  	Name:    "member_quota",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of custom quota settings for each team members.",
  	Columns: {&{Name: "email", Desc: "Email address of user."}, &{Name: "quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom"...}},
  }
```

## Changed report: namespace

```
  &dc_recipe.Report{
  	Name: "namespace",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning tea"...},
  	},
  }
```

## Changed report: namespace_file

```
  &dc_recipe.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the"...},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "Set if the folder is contained by a shared folder.",
- 		},
  	},
  }
```

## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
- 		&{Name: "result.namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "result.namespace_id", Desc: "The ID of this namespace."},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
  		&{Name: "result.path", Desc: "Path to the folder"},
  		&{Name: "result.count_file", Desc: "Number of files under the folder"},
  		... // 4 identical elements
  	},
  }
```

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "This report shows a list of shared links with the shared link owner team member.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "visibility", Desc: "The current visibility of the link after considering the shared "...},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		... // 2 identical elements
  	},
  }
```

## Changed report: usage

```
  &dc_recipe.Report{
  	Name:    "usage",
- 	Desc:    "",
+ 	Desc:    "This report shows current storage usage of users.",
  	Columns: {&{Name: "email", Desc: "Email address of the account"}, &{Name: "used_gb", Desc: "The user's total space usage (in GB, 1GB = 1024 MB)."}, &{Name: "used_bytes", Desc: "The user's total space usage (bytes)."}, &{Name: "allocation", Desc: "The user's space allocation (individual, or team)"}, ...},
  }
```
# Command spec changed: `team feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: feature

```
  &dc_recipe.Report{
  	Name:    "feature",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team features and their settings.",
  	Columns: {&{Name: "upload_api_rate_limit", Desc: "The number of upload API calls allowed per month."}, &{Name: "upload_api_rate_limit_count", Desc: "The number of upload API called this month."}, &{Name: "has_team_shared_dropbox", Desc: "Does this team have a shared team root."}, &{Name: "has_team_file_events", Desc: "Does this team have file events."}, ...},
  }
```
# Command spec changed: `team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: file_request

```
  &dc_recipe.Report{
  	Name: "file_request",
- 	Desc: "",
+ 	Desc: "This report shows a list of file requests with the file request owner team member.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "account_id", Desc: "Account ID of this file request owner."},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ID of file request owner user as a member of a team",
- 		},
  		&{Name: "email", Desc: "Email address of this file request owner."},
  		&{Name: "status", Desc: "The user status of this file request owner (active/invited/suspe"...},
  		&{Name: "surname", Desc: "Surname of this file request owner."},
  		&{Name: "given_name", Desc: "Given name of this file request owner."},
- 		&{Name: "file_request_id", Desc: "The ID of the file request."},
  		&{Name: "url", Desc: "The URL of the file request."},
  		&{Name: "title", Desc: "The title of the file request."},
  		... // 6 identical elements
  	},
  }
```
# Command spec changed: `team info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: info

```
  &dc_recipe.Report{
  	Name:    "info",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team information.",
  	Columns: {&{Name: "name", Desc: "The name of the team"}, &{Name: "team_id", Desc: "The ID of the team."}, &{Name: "num_licensed_users", Desc: "The number of licenses available to the team."}, &{Name: "num_provisioned_users", Desc: "The number of accounts that have been invited or are already act"...}, ...},
  }
```
# Command spec changed: `team linkedapp list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: linked_app

```
  &dc_recipe.Report{
  	Name: "linked_app",
- 	Desc: "",
+ 	Desc: "This report shows a list of linked app with the user of the app.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "app_id", Desc: "The application unique id."},
  		&{Name: "app_name", Desc: "The application name."},
  		&{Name: "is_app_folder", Desc: "Whether the linked application uses a dedicated folder."},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `team namespace file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "If true, deleted file or folder will be returned",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "If true, media info is set for photo and video in json report",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "If true, include team member folders",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "If true, include shared folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "If true, include team folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Desc:     "List only for the folder matched to the name",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: namespace_file

```
  &dc_recipe.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the"...},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "Set if the folder is contained by a shared folder.",
- 		},
  	},
  }
```
# Command spec changed: `team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Depth",
+ 			Desc:     "Report entry for all files and directories depth directories deep",
+ 			Default:  "1",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
+ 		},
+ 		&{
+ 			Name:     "IncludeAppFolder",
+ 			Desc:     "If true, include app folders",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "if true, include team member folders",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "If true, include shared folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "If true, include team folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Desc:     "List only for the folder matched to the name",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
- 		&{Name: "result.namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "result.namespace_id", Desc: "The ID of this namespace."},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
  		&{Name: "result.path", Desc: "Path to the folder"},
  		&{Name: "result.count_file", Desc: "Number of files under the folder"},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `team namespace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: namespace

```
  &dc_recipe.Report{
  	Name: "namespace",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning tea"...},
  	},
  }
```
# Command spec changed: `team namespace member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "AllColumns", Desc: "Show all columns", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: namespace_member

```
  &dc_recipe.Report{
  	Name: "namespace_member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
  		&{Name: "entry_access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		... // 5 identical elements
  	},
  }
```
# Command spec changed: `team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "Filter links by visibility (public/team_only/password)",
+ 			Default:  "public",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: shared_link

```
  &dc_recipe.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "This report shows a list of shared links with the shared link owner team member.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "visibility", Desc: "The current visibility of the link after considering the shared "...},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "At",
+ 			Desc:     "New expiration date and time",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "Days to the new expiration date",
+ 			Default:  "0",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "Target link visibility",
+ 			Default:  "public",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{...}},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: skipped

```
  &dc_recipe.Report{
  	Name:    "skipped",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of shared links with the shared link owner team member.",
  	Columns: {&{Name: "tag", Desc: "Entry type (file, or folder)"}, &{Name: "url", Desc: "URL of the shared link."}, &{Name: "name", Desc: "The linked file name (including extension)."}, &{Name: "expires", Desc: "Expiration time, if set."}, ...},
  }
```

## Changed report: updated

```
  &dc_recipe.Report{
  	Name: "updated",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{
- 			Name: "input.shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "input.tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "input.url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "input.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "input.visibility", Desc: "The current visibility of the link after considering the shared "...},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
- 		&{Name: "input.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "input.email", Desc: "Email address of user."},
  		&{Name: "input.status", Desc: "The user's status as a member of a specific team. (active/invite"...},
  		&{Name: "input.surname", Desc: "Surname of the link owner"},
  		&{Name: "input.given_name", Desc: "Given name of the link owner"},
- 		&{Name: "result.id", Desc: "A unique identifier for the linked file or folder"},
- 		&{Name: "result.tag", Desc: "Entry type (file, or folder)"},
- 		&{Name: "result.url", Desc: "URL of the shared link."},
- 		&{Name: "result.name", Desc: "The linked file name (including extension)."},
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox.",
- 		},
- 		&{
- 			Name: "result.visibility",
- 			Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part"...,
- 		},
  	},
  }
```
# Command spec changed: `teamfolder archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder batch archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for a list of team folder names",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "Name of team folder"},
- 		&{Name: "result.team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "result.name", Desc: "The name of the team folder."},
  		&{Name: "result.status", Desc: "The status of the team folder (active, archived, or archive_in_p"...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `teamfolder batch permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for a list of team folder names",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: {&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.name", Desc: "Name of team folder"}},
  }
```
# Command spec changed: `teamfolder batch replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DstPeerName",
+ 			Desc:     "Destination team account alias",
+ 			Default:  "dst",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for a list of team folder names",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "Source team account alias",
+ 			Default:  "src",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: verification

```
  &dc_recipe.Report{
  	Name:    "verification",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: {&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash,"...}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, ...},
  }
```
# Command spec changed: `teamfolder file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: namespace_file

```
  &dc_recipe.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the"...},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop clie"...},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "Set if the folder is contained by a shared folder.",
- 		},
  	},
  }
```
# Command spec changed: `teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Depth",
+ 			Desc:     "Depth",
+ 			Default:  "1",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: namespace_size

```
  &dc_recipe.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_fold"...},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
- 		&{Name: "result.namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "result.namespace_id", Desc: "The ID of this namespace."},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
  		&{Name: "result.path", Desc: "Path to the folder"},
  		&{Name: "result.count_file", Desc: "Number of files under the folder"},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `teamfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: team_folder

```
  &dc_recipe.Report{
  	Name: "team_folder",
- 	Desc: "",
+ 	Desc: "This report shows a list of team folders in the team.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "name", Desc: "The name of the team folder."},
  		&{Name: "status", Desc: "The status of the team folder (active, archived, or archive_in_p"...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `teamfolder permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "DstPeerName",
+ 			Desc:     "Destination team account alias",
+ 			Default:  "dst",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "Source team account alias",
+ 			Default:  "src",
+ 			TypeName: "string",
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```

## Changed report: verification

```
  &dc_recipe.Report{
  	Name:    "verification",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: {&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash,"...}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, ...},
  }
```
# Command spec changed: `web`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Port",
+ 			Desc:     "Port number",
+ 			Default:  "7800",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1024), "value": float64(7800)},
+ 		},
+ 	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
