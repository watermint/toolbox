# `リリース 63` から `リリース 64` までの変更点

# 追加されたコマンド

| コマンド                           | タイトル                                                        |
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



# 削除されたコマンド

| コマンド    | タイトル                                       |
|-------------|------------------------------------------------|
| dev ci auth | Authenticate for generating end to end testing |



# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev async`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 		&{
+ 			Name:     "RunConcurrently",
+ 			Desc:     "run concurrently",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: rows

```
  &rc_doc.Report{
  	Name: "rows",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups and their members.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev ci artifact up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path to upload",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path to upload",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

## 変更されたレポート: summary

```
  &rc_doc.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, &{Name: "num_files_upload", Desc: "The number of files uploaded or to upload."}, &{Name: "num_files_skip", Desc: "The number of files skipped or to skip."}, &{Name: "num_api_call", Desc: "The number of estimated upload API call for upload."}},
  }

```

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `dev desktop install`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "InstallerUrl",
+ 			Desc:     "Installer download URL",
+ 			Default:  "https://www.dropbox.com/download?full=1&os=win",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Silent", Desc: "Use Silent installer", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "SilentNoLaunch",
+ 			Desc:     "Use Enterprise installer",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev desktop start`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values:  []*rc_doc.Value{},
  }

```

# コマンド仕様の変更: `dev desktop stop`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "WaitSeconds",
+ 			Desc:     "Try stopping the app after given seconds.",
+ 			Default:  "<",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(60)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev desktop suspendupdate`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Unsuspend",
+ 			Desc:     "True when unsuspend Updater",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "UpdaterName",
+ 			Desc:     "Executable name of Dropbox Updater",
+ 			Default:  "DropboxUpdate.exe",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "UpdaterPath",
+ 			Desc:     "Path to Dropbox Updater",
+ 			Default:  "C:/Program Files (x86)/Dropbox/Update",
+ 			TypeName: "string",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev diag procmon`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev diag procmon",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -repository-path /LOCAL/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path to upload procmon logs",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "ProcmonUrl",
+ 			Desc:     "Process monitor download url",
+ 			Default:  "https://download.sysinternals.com/files/ProcessMonitor.zip",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepositoryPath",
+ 			Desc:     "Procmon Work directory",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "RetainLogs",
+ 			Desc:     "Number of Procmon logs to retain",
+ 			Default:  "\x04",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(10000), "min": float64(0), "value": float64(4)},
+ 		},
+ 		&{
+ 			Name:     "RunUntil",
+ 			Desc:     "Skip run after this date/time",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "Duration for waiting procmon",
+ 			Default:  "܈",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(10), "value": float64(1800)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev doc`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Badge",
+ 			Desc:     "Include badges of build status",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "CommandPath",
+ 			Desc:     "Relative path to generate command manuals",
+ 			Default:  "doc/generated/",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Filename", Desc: "Filename", Default: "README.md", TypeName: "string"},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "MarkdownReadme",
+ 			Desc:     "Generate README as markdown format",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev dummy`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "Dest", Desc: "Dummy file destination", TypeName: "string"},
+ 		&{
+ 			Name:     "MaxEntry",
+ 			Desc:     "Maximum entries",
+ 			Default:  "\x00",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 		&{Name: "Path", Desc: "Path to dummy entry file", TypeName: "string"},
+ 	},
  }

```

# コマンド仕様の変更: `dev echo`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values:  []*rc_doc.Value{&{Name: "Text", Desc: "Text to echo", TypeName: "string"}},
  }

```

# コマンド仕様の変更: `dev preflight`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values:  []*rc_doc.Value{},
  }

```

# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "Path to the test resource location",
+ 			Default:  "test/dev/resource.json",
+ 			TypeName: "string",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev release publish",
- 	CliArgs:         "",
+ 	CliArgs:         "-artifact-path /LOCAL/PATH/TO/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "ArtifactPath",
+ 			Desc:     "Path to artifacts",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "Branch", Desc: "Target branch", Default: "master", TypeName: "string"},
+ 		&{
+ 			Name:     "ConnGithub",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
+ 		},
+ 		&{
+ 			Name:     "SkipTests",
+ 			Desc:     "Skip end to end tests.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "Path to test resource",
+ 			Default:  "test/dev/resource.json",
+ 			TypeName: "string",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev spec diff`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "FilePath",
+ 			Desc:     "File path to output",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Release1",
+ 			Desc:     "Release name 1",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Release2",
+ 			Desc:     "Release name 2",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev spec doc`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "FilePath",
+ 			Desc:     "File path",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Lang",
+ 			Desc:     "Language",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test monkey",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Distribution",
+ 			Desc:     "Number of files/folder distribution",
+ 			Default:  "✐",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(10000)},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "File extensions (comma separated)",
+ 			Default:  "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,bmp,wmi,ini,ai,psd",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Monkey test path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "Monkey test duration in seconds",
+ 			Default:  "\n",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(1), "value": float64(10)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev test recipe`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "Recipe",
+ 			Desc:     "Recipe name to test",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Resource",
+ 			Desc:     "Test resource file path",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Verbose",
+ 			Desc:     "Verbose output for testing",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `dev test resources`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values:  []*rc_doc.Value{},
  }

```

# コマンド仕様の変更: `dev util curl`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "BufferSize",
+ 			Desc:     "Size of buffer",
+ 			Default:  "𐀀",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.097152e+06), "min": float64(1024), "value": float64(65536)},
+ 		},
+ 		&{Name: "Record", Desc: "Capture record(s) for the test", TypeName: "string"},
+ 	},
  }

```

# コマンド仕様の変更: `dev util wait`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "Wait seconds",
+ 			Default:  "\x01",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(604800), "min": float64(1), "value": float64(1)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Left",
+ 			Desc:     "Account alias (left)",
+ 			Default:  "left",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "LeftPath",
+ 			Desc:     "The path from account root (left)",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Right",
+ 			Desc:     "Account alias (right)",
+ 			Default:  "right",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "RightPath",
+ 			Desc:     "The path from account root (right)",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  }

```

## 変更されたレポート: diff

```
  &rc_doc.Report{
  	Name:    "diff",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing."}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, &{Name: "left_hash", Desc: "Content hash of left file"}, &{Name: "right_path", Desc: "path of right"}, &{Name: "right_kind", Desc: "folder of file"}, &{Name: "right_size", Desc: "size of right file"}, &{Name: "right_hash", Desc: "Content hash of right file"}},
  }

```

# コマンド仕様の変更: `file compare local`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: diff

```
  &rc_doc.Report{
  	Name:    "diff",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing."}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, &{Name: "left_hash", Desc: "Content hash of left file"}, &{Name: "right_path", Desc: "path of right"}, &{Name: "right_kind", Desc: "folder of file"}, &{Name: "right_size", Desc: "size of right file"}, &{Name: "right_hash", Desc: "Content hash of right file"}},
  }

```

## 変更されたレポート: skip

```
  &rc_doc.Report{
  	Name:    "skip",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing."}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, &{Name: "left_hash", Desc: "Content hash of left file"}, &{Name: "right_path", Desc: "path of right"}, &{Name: "right_kind", Desc: "folder of file"}, &{Name: "right_size", Desc: "size of right file"}, &{Name: "right_hash", Desc: "Content hash of right file"}},
  }

```

# コマンド仕様の変更: `file copy`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Destination path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Source path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `file delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to delete",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file download",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/OF/FILE -local-path /LOCAL/PATH/TO/DOWNLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "File path to download",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path to download",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file export doc",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/FILE -local-path /LOCAL/PATH/TO/EXPORT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox document path to export.",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local path to save",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows a result of exporting file.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
  		&{Name: "export_name", Desc: "File name for export file."},
  		&{Name: "export_size", Desc: "File size of export file."},
- 		&{Name: "export_hash", Desc: "Content hash of export file."},
  	},
  }

```

# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to import",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.url", Desc: "Url to download"},
  		&{Name: "input.path", Desc: "Path to store file (use path given by `-path` when the record is empty)"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
  		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file import url`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to import",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Url", Desc: "URL", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
  		&{Name: "revision", Desc: "A unique identifier for the current revision of a file."},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "Include deleted files",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "Include media information",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool"},
+ 	},
  }

```

## 変更されたレポート: file_list

```
  &rc_doc.Report{
  	Name: "file_list",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file merge`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool"},
+ 		&{
+ 			Name:     "From",
+ 			Desc:     "Path for merge",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "KeepEmptyFolder",
+ 			Desc:     "Keep empty folder after merge",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "To",
+ 			Desc:     "Path to merge",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "WithinSameNamespace",
+ 			Desc:     "Do not cross namespace. That is for preserve sharing permission including a shared link",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Destination path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Source path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `file replication`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Account alias (destionation)",
+ 			Default:  "dst",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "DstPath",
+ 			Desc:     "Destination path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Account alias (source)",
+ 			Default:  "src",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "SrcPath",
+ 			Desc:     "Source path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 	},
  }

```

## 変更されたレポート: replication_diff

```
  &rc_doc.Report{
  	Name:    "replication_diff",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing."}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, &{Name: "left_hash", Desc: "Content hash of left file"}, &{Name: "right_path", Desc: "path of right"}, &{Name: "right_kind", Desc: "folder of file"}, &{Name: "right_size", Desc: "size of right file"}, &{Name: "right_hash", Desc: "Content hash of right file"}},
  }

```

# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file restore",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.path", Desc: "Path"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
  		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Restricts search to only the file categories specified (image/document/pdf/spreadsheet/presentation/audio/video/folder/paper/others).",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{
+ 				"options": []interface{}{
+ 					string(""),
+ 					string("image"),
+ 					string("document"),
+ 					string("pdf"),
+ 					string("spreadsheet"),
+ 					string("presentation"),
+ 					string("audio"),
+ 					string("video"),
+ 					string("folder"),
+ 					string("paper"),
+ 					string("others"),
+ 				},
+ 			},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "Restricts search to only the extensions specified.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Scopes the search to a path in the user's Dropbox.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: matches

```
  &rc_doc.Report{
  	Name:    "matches",
- 	Desc:    "",
+ 	Desc:    "This report shows a result of search with highlighted text.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "tag", Desc: "Type of entry"}, &{Name: "path_display", Desc: "Display path"}, &{Name: "highlight_html", Desc: "Highlighted text in HTML"}},
  }

```

# コマンド仕様の変更: `file search name`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Restricts search to only the file categories specified (image/document/pdf/spreadsheet/presentation/audio/video/folder/paper/others).",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{
+ 				"options": []interface{}{
+ 					string(""),
+ 					string("image"),
+ 					string("document"),
+ 					string("pdf"),
+ 					string("spreadsheet"),
+ 					string("presentation"),
+ 					string("audio"),
+ 					string("video"),
+ 					string("folder"),
+ 					string("paper"),
+ 					string("others"),
+ 				},
+ 			},
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "Restricts search to only the extensions specified.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Scopes the search to a path in the user's Dropbox.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: matches

```
  &rc_doc.Report{
  	Name:    "matches",
- 	Desc:    "",
+ 	Desc:    "This report shows a result of search with highlighted text.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "tag", Desc: "Type of entry"}, &{Name: "path_display", Desc: "Display path"}, &{Name: "highlight_html", Desc: "Highlighted text in HTML"}},
  }

```

# コマンド仕様の変更: `file sync preflight up`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync preflight up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Destination Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

## 変更されたレポート: summary

```
  &rc_doc.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, &{Name: "num_files_upload", Desc: "The number of files uploaded or to upload."}, &{Name: "num_files_skip", Desc: "The number of files skipped or to skip."}, &{Name: "num_api_call", Desc: "The number of estimated upload API call for upload."}},
  }

```

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "ChunkSizeKb",
+ 			Desc:     "Upload chunk size in KB",
+ 			Default:  "𥠀",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)},
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Destination Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

## 変更されたレポート: summary

```
  &rc_doc.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, &{Name: "num_files_upload", Desc: "The number of files uploaded or to upload."}, &{Name: "num_files_skip", Desc: "The number of files skipped or to skip."}, &{Name: "num_api_call", Desc: "The number of estimated upload API call for upload."}},
  }

```

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file upload`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "ChunkSizeKb",
+ 			Desc:     "Upload chunk size in KB",
+ 			Default:  "𥠀",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)},
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Destination Dropbox path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Overwrite",
+ 			Desc:     "Overwrite existing file(s)",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

## 変更されたレポート: summary

```
  &rc_doc.Report{
  	Name:    "summary",
- 	Desc:    "",
+ 	Desc:    "This report shows a summary of the upload results.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "upload_start", Desc: "Time of start uploading"}, &{Name: "upload_end", Desc: "Time of finish uploading"}, &{Name: "num_bytes", Desc: "Total upload size (Bytes)"}, &{Name: "num_files_error", Desc: "The number of files failed or got an error."}, &{Name: "num_files_upload", Desc: "The number of files uploaded or to upload."}, &{Name: "num_files_skip", Desc: "The number of files skipped or to skip."}, &{Name: "num_api_call", Desc: "The number of estimated upload API call for upload."}},
  }

```

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.file", Desc: "Local file path"},
  		&{Name: "input.size", Desc: "Local file size"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
- 		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "result.server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "result.revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "result.size", Desc: "The file size in bytes."},
  		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `file watch`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file watch",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/WATCH",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to watch",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "Watch path recursively",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "filerequest create",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "AllowLateUploads",
+ 			Desc:     "If set, allow uploads after the deadline has passed (one_day/two_days/seven_days/thirty_days/always)",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Deadline",
+ 			Desc:     "The deadline for this file request.",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "The path for the folder in the Dropbox where uploaded files will be sent.",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Title", Desc: "The title of the file request", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: file_request

```
  &rc_doc.Report{
  	Name:    "file_request",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, &{Name: "is_open", Desc: "Whether or not the file request is open."}, &{Name: "file_count", Desc: "The number of files this file request has received."}, &{Name: "destination", Desc: "The path for the folder in the Dropbox where uploaded files will be sent."}, &{Name: "deadline", Desc: "The deadline for this file request."}, &{Name: "deadline_allow_late_uploads", Desc: "If set, allow uploads after the deadline has passed."}},
  }

```

# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: deleted

```
  &rc_doc.Report{
  	Name:    "deleted",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, &{Name: "is_open", Desc: "Whether or not the file request is open."}, &{Name: "file_count", Desc: "The number of files this file request has received."}, &{Name: "destination", Desc: "The path for the folder in the Dropbox where uploaded files will be sent."}, &{Name: "deadline", Desc: "The deadline for this file request."}, &{Name: "deadline_allow_late_uploads", Desc: "If set, allow uploads after the deadline has passed."}},
  }

```

# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Force",
+ 			Desc:     "Force delete the file request.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{Name: "Url", Desc: "URL of the file request.", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: deleted

```
  &rc_doc.Report{
  	Name:    "deleted",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, &{Name: "is_open", Desc: "Whether or not the file request is open."}, &{Name: "file_count", Desc: "The number of files this file request has received."}, &{Name: "destination", Desc: "The path for the folder in the Dropbox where uploaded files will be sent."}, &{Name: "deadline", Desc: "The deadline for this file request."}, &{Name: "deadline_allow_late_uploads", Desc: "If set, allow uploads after the deadline has passed."}},
  }

```

# コマンド仕様の変更: `filerequest list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: file_requests

```
  &rc_doc.Report{
  	Name:    "file_requests",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of file requests.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "id", Desc: "The Id of the file request"}, &{Name: "url", Desc: "The URL of the file request"}, &{Name: "title", Desc: "The title of the file request"}, &{Name: "created", Desc: "Date/time of the file request was created."}, &{Name: "is_open", Desc: "Whether or not the file request is open."}, &{Name: "file_count", Desc: "The number of files this file request has received."}, &{Name: "destination", Desc: "The path for the folder in the Dropbox where uploaded files will be sent."}, &{Name: "deadline", Desc: "The deadline for this file request."}, &{Name: "deadline_allow_late_uploads", Desc: "If set, allow uploads after the deadline has passed."}},
  }

```

# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "ManagementType",
+ 			Desc:     "Group management type `company_managed` or `user_managed`",
+ 			Default:  "company_managed",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("company_managed"), string("user_managed")}},
+ 		},
+ 		&{Name: "Name", Desc: "Group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: added_group

```
  &rc_doc.Report{
  	Name: "added_group",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group batch delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for group name list",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.name", Desc: "Group name"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "Name", Desc: "Group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `group list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  }

```

## 変更されたレポート: group

```
  &rc_doc.Report{
  	Name: "group",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "GroupName", Desc: "Group name", TypeName: "string"},
+ 		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "Email address of the member"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group member delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "GroupName", Desc: "Name of the group", TypeName: "string"},
+ 		&{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "Email address of the member"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group member list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  }

```

## 変更されたレポート: group_member

```
  &rc_doc.Report{
  	Name: "group_member",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups and their members.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "CurrentName", Desc: "Current group name", TypeName: "string"},
+ 		&{Name: "NewName", Desc: "New group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.new_name", Desc: "New group name"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "Target days old",
+ 			Default:  "\a",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(7)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "Target days old",
+ 			Default:  "\x1c",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(28)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `job history list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values:  []*rc_doc.Value{},
  }

```

## 変更されたレポート: log

```
  &rc_doc.Report{
  	Name:    "log",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of job histories.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "job_id", Desc: "Job ID"}, &{Name: "app_version", Desc: "App version"}, &{Name: "recipe_name", Desc: "Command"}, &{Name: "time_start", Desc: "Time Started"}, &{Name: "time_finish", Desc: "Time Finished"}},
  }

```

# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job history ship",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Desc:     "Dropbox path to upload",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 11 identical elements
  		&{Name: "result.revision", Desc: "A unique identifier for the current revision of a file."},
  		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "result.shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "result.parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `job loop`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job loop",
- 	CliArgs:         "",
+ 	CliArgs:         `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 17:58:38"`,
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "IntervalSeconds",
+ 			Desc:     "Interval seconds",
+ 			Default:  "´",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(1), "value": float64(180)},
+ 		},
+ 		&{Name: "QuitOnError", Desc: "Quit on error", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Desc:     "Path to runbook",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Until",
+ 			Desc:     "Run until this date/time",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(false)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `job run`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job run",
- 	CliArgs:         "",
+ 	CliArgs:         "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Fork",
+ 			Desc:     "Fork process on running the workflow",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Desc:     "Path to the runbook",
+ 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "TimeoutSeconds",
+ 			Desc:     "Terminate process when given time passed",
+ 			Default:  "\x00",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(0), "value": float64(0)},
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `license`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values:  []*rc_doc.Value{},
  }

```

# コマンド仕様の変更: `member delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "WipeData",
+ 			Desc:     "If true, controls if the user's data will be deleted on their linked devices",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.email", Desc: "Email address of the account"}},
  }

```

# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "RevokeTeamShares",
+ 			Desc:     "True for revoke shared folder access which owned by the team",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.email", Desc: "Email address of the account"}},
  }

```

# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "SilentInvite",
+ 			Desc:     "Do not send welcome email (requires SSO + domain verification instead)",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "Given name of the account"},
  		&{Name: "input.surname", Desc: "Surname of the account"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  }

```

## 変更されたレポート: member

```
  &rc_doc.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "member_folder_id", Desc: "The namespace id of the user's root folder."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  	},
  }

```

# コマンド仕様の変更: `member quota list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: member_quota

```
  &rc_doc.Report{
  	Name:    "member_quota",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of custom quota settings for each team members.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "email", Desc: "Email address of user."}, &{Name: "quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set."}},
  }

```

# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "Quota",
+ 			Desc:     "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom quota set.",
+ 			Default:  "\x00",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.email", Desc: "Email address of user."}, &{Name: "input.quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set."}, &{Name: "result.email", Desc: "Email address of user."}, &{Name: "result.quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set."}},
  }

```

# コマンド仕様の変更: `member quota usage`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: usage

```
  &rc_doc.Report{
  	Name:    "usage",
- 	Desc:    "",
+ 	Desc:    "This report shows current storage usage of users.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "email", Desc: "Email address of the account"}, &{Name: "used_gb", Desc: "The user's total space usage (in GB, 1GB = 1024 MB)."}, &{Name: "used_bytes", Desc: "The user's total space usage (bytes)."}, &{Name: "allocation", Desc: "The user's space allocation (individual, or team)"}, &{Name: "allocated", Desc: "The total space allocated to the user's account (bytes)."}},
  }

```

# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "Silent",
+ 			Desc:     "Do not send welcome email (SSO required)",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "input.email", Desc: "Email address of user."},
  		&{Name: "input.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "input.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "input.given_name", Desc: "Also known as a first name"},
  		&{Name: "input.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "input.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "input.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "input.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "input.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "input.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "input.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "input.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "input.tag", Desc: "Operation tag"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Dst",
+ 			Desc:     "Destination team; team file access",
+ 			Default:  "dst",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "Source team; team file access",
+ 			Default:  "src",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.src_email", Desc: "Source account's email address"}, &{Name: "input.dst_email", Desc: "Destination account's email address"}},
  }

```

# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 		&{
+ 			Name:     "UpdateUnverified",
+ 			Desc:     "Update an account which didn't verified email. If an account email unverified, email address change may affect lose invitation to folders.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.from_email", Desc: "Current Email address"},
  		&{Name: "input.to_email", Desc: "New Email address"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.email", Desc: "Email address of team members"},
  		&{Name: "input.external_id", Desc: "External ID of team members"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "Given name of the account"},
  		&{Name: "input.surname", Desc: "Surname of the account"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "result.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "result.member_folder_id",
- 			Desc: "The namespace id of the user's root folder.",
- 		},
- 		&{
- 			Name: "result.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "result.account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "result.persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `sharedfolder list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: shared_folder

```
  &rc_doc.Report{
  	Name: "shared_folder",
- 	Desc: "",
+ 	Desc: "This report shows a list of shared folders.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder.",
- 		},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		... // 5 identical elements
  		&{Name: "policy_member", Desc: "Who can be a member of this shared folder, as set on the folder itself (team, or anyone)"},
  		&{Name: "policy_viewer_info", Desc: "Who can enable/disable viewer info for this shared folder."},
- 		&{Name: "owner_team_id", Desc: "Team ID of the team that owns the folder"},
  		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
  	},
  }

```

# コマンド仕様の変更: `sharedfolder member list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: member

```
  &rc_doc.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members of shared folders.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder.",
- 		},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 2 identical elements
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		&{Name: "is_inherited", Desc: "True if the member has access from a parent folder"},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "invitee_email", Desc: "Email address of invitee for this folder"},
  	},
  }

```

# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Expires",
+ 			Desc:     "Expiration date/time of the shared link",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Password",
+ 			Desc:     "Password",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "TeamOnly",
+ 			Desc:     "Link is accessible only by team members",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: created

```
  &rc_doc.Report{
  	Name:    "created",
- 	Desc:    "",
+ 	Desc:    "THis report shows a list of shared links.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "id", Desc: "A unique identifier for the linked file or folder"}, &{Name: "tag", Desc: "Entry type (file, or folder)"}, &{Name: "url", Desc: "URL of the shared link."}, &{Name: "name", Desc: "The linked file name (including extension)."}, &{Name: "expires", Desc: "Expiration time, if set."}, &{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."}, &{Name: "visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."}},
  }

```

# コマンド仕様の変更: `sharedlink delete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "File or folder path to remove shared link",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "Attempt to remove the file hierarchy",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 	},
  }

```

## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.id", Desc: "A unique identifier for the linked file or folder"},
  		&{Name: "input.tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "input.url", Desc: "URL of the shared link."},
  		... // 4 identical elements
  	},
  }

```

# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedlink file list",
- 	CliArgs:         "",
+ 	CliArgs:         "-url SHAREDLINK_URL",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 6 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Password",
+ 			Desc:     "Password for the shared link",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 		&{
+ 			Name:     "Url",
+ 			Desc:     "Shared link URL",
+ 			TypeName: "domain.dropbox.model.mo_url.url_impl",
+ 		},
+ 	},
  }

```

## 変更されたレポート: file_list

```
  &rc_doc.Report{
  	Name: "file_list",
- 	Desc: "",
+ 	Desc: "This report shows a list of metadata of files or folders in the path.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "ID of shared folder that holds this file.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `sharedlink list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "THis report shows a list of shared links.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the linked file or folder"},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 4 identical elements
  	},
  }

```

# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Filter the returned events to a single category. This field is optional.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "Ending time (exclusive).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "User email address list file",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "Starting time (inclusive)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  }

```

## 変更されたレポート: combined

```
  &rc_doc.Report{
  	Name:    "combined",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, &{Name: "category", Desc: "Category of the events in event audit log."}, &{Name: "access_method", Desc: "The method that was used to perform the action."}, &{Name: "ip_address", Desc: "IP Address."}, &{Name: "country", Desc: "Country code."}, &{Name: "city", Desc: "City name"}, &{Name: "involve_non_team_members", Desc: "True if the action involved a non team member either as the actor or as one of the affected users."}, &{Name: "participants", Desc: "Zero or more users and/or groups that are affected by the action."}, &{Name: "context", Desc: "The user or team on whose behalf the actor performed the action."}, &{Name: "assets", Desc: "Zero or more content assets involved in the action."}, &{Name: "other_info", Desc: "The variable event schema applicable to this type of action."}},
  }

```

## 変更されたレポート: user

```
  &rc_doc.Report{
  	Name:    "user",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, &{Name: "category", Desc: "Category of the events in event audit log."}, &{Name: "access_method", Desc: "The method that was used to perform the action."}, &{Name: "ip_address", Desc: "IP Address."}, &{Name: "country", Desc: "Country code."}, &{Name: "city", Desc: "City name"}, &{Name: "involve_non_team_members", Desc: "True if the action involved a non team member either as the actor or as one of the affected users."}, &{Name: "participants", Desc: "Zero or more users and/or groups that are affected by the action."}, &{Name: "context", Desc: "The user or team on whose behalf the actor performed the action."}, &{Name: "assets", Desc: "Zero or more content assets involved in the action."}, &{Name: "other_info", Desc: "The variable event schema applicable to this type of action."}},
  }

```

# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Event category",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndDate",
+ 			Desc:     "End date",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{Name: "StartDate", Desc: "Start date", TypeName: "string"},
+ 	},
  }

```

## 変更されたレポート: event

```
  &rc_doc.Report{
  	Name:    "event",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, &{Name: "category", Desc: "Category of the events in event audit log."}, &{Name: "access_method", Desc: "The method that was used to perform the action."}, &{Name: "ip_address", Desc: "IP Address."}, &{Name: "country", Desc: "Country code."}, &{Name: "city", Desc: "City name"}, &{Name: "involve_non_team_members", Desc: "True if the action involved a non team member either as the actor or as one of the affected users."}, &{Name: "participants", Desc: "Zero or more users and/or groups that are affected by the action."}, &{Name: "context", Desc: "The user or team on whose behalf the actor performed the action."}, &{Name: "assets", Desc: "Zero or more content assets involved in the action."}, &{Name: "other_info", Desc: "The variable event schema applicable to this type of action."}},
  }

```

# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Filter the returned events to a single category. This field is optional.",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "Ending time (exclusive).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "Starting time (inclusive)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  }

```

## 変更されたレポート: event

```
  &rc_doc.Report{
  	Name:    "event",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, &{Name: "category", Desc: "Category of the events in event audit log."}, &{Name: "access_method", Desc: "The method that was used to perform the action."}, &{Name: "ip_address", Desc: "IP Address."}, &{Name: "country", Desc: "Country code."}, &{Name: "city", Desc: "City name"}, &{Name: "involve_non_team_members", Desc: "True if the action involved a non team member either as the actor or as one of the affected users."}, &{Name: "participants", Desc: "Zero or more users and/or groups that are affected by the action."}, &{Name: "context", Desc: "The user or team on whose behalf the actor performed the action."}, &{Name: "assets", Desc: "Zero or more content assets involved in the action."}, &{Name: "other_info", Desc: "The variable event schema applicable to this type of action."}},
  }

```

# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Category",
+ 			Desc:     "Filter the returned events to a single category. This field is optional.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "EndTime",
+ 			Desc:     "Ending time (exclusive).",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Desc:     "Starting time (inclusive)",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 	},
  }

```

## 変更されたレポート: user

```
  &rc_doc.Report{
  	Name:    "user",
- 	Desc:    "",
+ 	Desc:    "This report shows an activity logs with mostly compatible with Dropbox Business's activity logs.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "timestamp", Desc: "The Dropbox timestamp representing when the action was taken."}, &{Name: "member", Desc: "User display name"}, &{Name: "member_email", Desc: "User email address"}, &{Name: "event_type", Desc: "The particular type of action taken."}, &{Name: "category", Desc: "Category of the events in event audit log."}, &{Name: "access_method", Desc: "The method that was used to perform the action."}, &{Name: "ip_address", Desc: "IP Address."}, &{Name: "country", Desc: "Country code."}, &{Name: "city", Desc: "City name"}, &{Name: "involve_non_team_members", Desc: "True if the action involved a non team member either as the actor or as one of the affected users."}, &{Name: "participants", Desc: "Zero or more users and/or groups that are affected by the action."}, &{Name: "context", Desc: "The user or team on whose behalf the actor performed the action."}, &{Name: "assets", Desc: "Zero or more content assets involved in the action."}, &{Name: "other_info", Desc: "The variable event schema applicable to this type of action."}},
  }

```

## 変更されたレポート: user_summary

```
  &rc_doc.Report{
  	Name: "user_summary",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.user", Desc: "User email address"},
- 		&{Name: "result.user", Desc: "User email address"},
  		&{Name: "result.logins", Desc: "Number of login activities"},
  		&{Name: "result.devices", Desc: "Number of device activities"},
  		... // 4 identical elements
  	},
  }

```

# コマンド仕様の変更: `team content member`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: membership

```
  &rc_doc.Report{
  	Name:    "membership",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of shared folders and team folders with their members. If a folder has multiple members, then members are listed with rows.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "path", Desc: "Path"}, &{Name: "folder_type", Desc: "Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)"}, &{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"}, &{Name: "access_type", Desc: "User's access level for this folder"}, &{Name: "member_type", Desc: "Type of this member (user, group, or invitee)"}, &{Name: "member_name", Desc: "Name of this member"}, &{Name: "member_email", Desc: "Email address of this member"}},
  }

```

## 変更されたレポート: no_member

```
  &rc_doc.Report{
  	Name:    "no_member",
- 	Desc:    "",
+ 	Desc:    "This report shows folders without members.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"}, &{Name: "path", Desc: "Path"}, &{Name: "folder_type", Desc: "Type of the folder. (`team_folder`: a team folder or in a team folder, `shared_folder`: a shared folder)"}},
  }

```

# コマンド仕様の変更: `team content policy`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: policy

```
  &rc_doc.Report{
  	Name:    "policy",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of shared folders and team folders with their current policy settings.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "path", Desc: "Path"}, &{Name: "is_team_folder", Desc: "`true` if the folder is a team folder, or inside of a team folder"}, &{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"}, &{Name: "policy_manage_access", Desc: "Who can add and remove members from this shared folder."}, &{Name: "policy_shared_link", Desc: "Who links can be shared with."}, &{Name: "policy_member", Desc: "Who can be a member of this shared folder, taking into account both the folder and the team-wide policy."}, &{Name: "policy_viewer_info", Desc: "Who can enable/disable viewer info for this shared folder."}},
  }

```

# コマンド仕様の変更: `team device list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: device

```
  &rc_doc.Report{
  	Name: "device",
- 	Desc: "",
+ 	Desc: "This report shows a list of current existing sessions in the team with team member information.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }

```

# コマンド仕様の変更: `team device unlink`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DeleteOnUnlink",
+ 			Desc:     "Delete files on unlink",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 5 identical elements
  		&{Name: "input.given_name", Desc: "Also known as a first name"},
  		&{Name: "input.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "input.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{
- 			Name: "input.abbreviated_name",
- 			Desc: "An abbreviated form of the person's name.",
- 		},
- 		&{
- 			Name: "input.external_id",
- 			Desc: "External ID that a team can attach to the user.",
- 		},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
  		&{Name: "input.device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "input.id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }

```

# コマンド仕様の変更: `team diag explorer`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "All",
+ 			Desc:     "Include additional reports",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Dropbox Business file access",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Info",
+ 			Desc:     "Dropbox Business information access",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 		&{
+ 			Name:     "Mgmt",
+ 			Desc:     "Dropbox Business management",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
+ 		},
+ 	},
  }

```

## 変更されたレポート: device

```
  &rc_doc.Report{
  	Name: "device",
- 	Desc: "",
+ 	Desc: "This report shows a list of current existing sessions in the team with team member information.",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }

```

## 変更されたレポート: feature

```
  &rc_doc.Report{
  	Name:    "feature",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team features and their settings.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "upload_api_rate_limit", Desc: "The number of upload API calls allowed per month."}, &{Name: "upload_api_rate_limit_count", Desc: "The number of upload API called this month."}, &{Name: "has_team_shared_dropbox", Desc: "Does this team have a shared team root."}, &{Name: "has_team_file_events", Desc: "Does this team have file events."}, &{Name: "has_team_selective_sync", Desc: "Does this team have team selective sync enabled."}},
  }

```

## 変更されたレポート: file_request

```
  &rc_doc.Report{
  	Name: "file_request",
- 	Desc: "",
+ 	Desc: "This report shows a list of file requests with the file request owner team member.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "account_id", Desc: "Account ID of this file request owner."},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ID of file request owner user as a member of a team",
- 		},
  		&{Name: "email", Desc: "Email address of this file request owner."},
  		&{Name: "status", Desc: "The user status of this file request owner (active/invited/suspended/removed)"},
  		&{Name: "surname", Desc: "Surname of this file request owner."},
  		&{Name: "given_name", Desc: "Given name of this file request owner."},
- 		&{Name: "file_request_id", Desc: "The ID of the file request."},
  		&{Name: "url", Desc: "The URL of the file request."},
  		&{Name: "title", Desc: "The title of the file request."},
  		... // 6 identical elements
  	},
  }

```

## 変更されたレポート: group

```
  &rc_doc.Report{
  	Name: "group",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }

```

## 変更されたレポート: group_member

```
  &rc_doc.Report{
  	Name: "group_member",
- 	Desc: "",
+ 	Desc: "This report shows a list of groups and their members.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

## 変更されたレポート: info

```
  &rc_doc.Report{
  	Name:    "info",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team information.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "name", Desc: "The name of the team"}, &{Name: "team_id", Desc: "The ID of the team."}, &{Name: "num_licensed_users", Desc: "The number of licenses available to the team."}, &{Name: "num_provisioned_users", Desc: "The number of accounts that have been invited or are already active members of the team."}, &{Name: "policy_shared_folder_member", Desc: "Which shared folders team members can join (from_team_only, or from_anyone)"}, &{Name: "policy_shared_folder_join", Desc: "Who can join folders shared by team members (team, or anyone)"}, &{Name: "policy_shared_link_create", Desc: "Who can view shared links owned by team members (default_public, default_team_only, or team_only)"}, &{Name: "policy_emm_state", Desc: "This describes the Enterprise Mobility Management (EMM) state for this team (disabled, optional, or required)"}, &{Name: "policy_office_add_in", Desc: "The admin policy around the Dropbox Office Add-In for this team (disabled, or enabled)"}},
  }

```

## 変更されたレポート: linked_app

```
  &rc_doc.Report{
  	Name: "linked_app",
- 	Desc: "",
+ 	Desc: "This report shows a list of linked app with the user of the app.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "app_id", Desc: "The application unique id."},
  		&{Name: "app_name", Desc: "The application name."},
  		&{Name: "is_app_folder", Desc: "Whether the linked application uses a dedicated folder."},
  		... // 3 identical elements
  	},
  }

```

## 変更されたレポート: member

```
  &rc_doc.Report{
  	Name: "member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "member_folder_id", Desc: "The namespace id of the user's root folder."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  	},
  }

```

## 変更されたレポート: member_quota

```
  &rc_doc.Report{
  	Name:    "member_quota",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of custom quota settings for each team members.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "email", Desc: "Email address of user."}, &{Name: "quota", Desc: "Custom quota in GB (1 TB = 1024 GB). 0 if the user has no custom quota set."}},
  }

```

## 変更されたレポート: namespace

```
  &rc_doc.Report{
  	Name: "namespace",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning team member."},
  	},
  }

```

## 変更されたレポート: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the owning team member."},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "Set if the folder is contained by a shared folder.",
- 		},
  	},
  }

```

## 変更されたレポート: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
- 		&{Name: "result.namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "result.namespace_id", Desc: "The ID of this namespace."},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
  		&{Name: "result.path", Desc: "Path to the folder"},
  		&{Name: "result.count_file", Desc: "Number of files under the folder"},
  		... // 4 identical elements
  	},
  }

```

## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "This report shows a list of shared links with the shared link owner team member.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

## 変更されたレポート: usage

```
  &rc_doc.Report{
  	Name:    "usage",
- 	Desc:    "",
+ 	Desc:    "This report shows current storage usage of users.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "email", Desc: "Email address of the account"}, &{Name: "used_gb", Desc: "The user's total space usage (in GB, 1GB = 1024 MB)."}, &{Name: "used_bytes", Desc: "The user's total space usage (bytes)."}, &{Name: "allocation", Desc: "The user's space allocation (individual, or team)"}, &{Name: "allocated", Desc: "The total space allocated to the user's account (bytes)."}},
  }

```

# コマンド仕様の変更: `team feature`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  }

```

## 変更されたレポート: feature

```
  &rc_doc.Report{
  	Name:    "feature",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team features and their settings.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "upload_api_rate_limit", Desc: "The number of upload API calls allowed per month."}, &{Name: "upload_api_rate_limit_count", Desc: "The number of upload API called this month."}, &{Name: "has_team_shared_dropbox", Desc: "Does this team have a shared team root."}, &{Name: "has_team_file_events", Desc: "Does this team have file events."}, &{Name: "has_team_selective_sync", Desc: "Does this team have team selective sync enabled."}},
  }

```

# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: file_request

```
  &rc_doc.Report{
  	Name: "file_request",
- 	Desc: "",
+ 	Desc: "This report shows a list of file requests with the file request owner team member.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "account_id", Desc: "Account ID of this file request owner."},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ID of file request owner user as a member of a team",
- 		},
  		&{Name: "email", Desc: "Email address of this file request owner."},
  		&{Name: "status", Desc: "The user status of this file request owner (active/invited/suspended/removed)"},
  		&{Name: "surname", Desc: "Surname of this file request owner."},
  		&{Name: "given_name", Desc: "Given name of this file request owner."},
- 		&{Name: "file_request_id", Desc: "The ID of the file request."},
  		&{Name: "url", Desc: "The URL of the file request."},
  		&{Name: "title", Desc: "The title of the file request."},
  		... // 6 identical elements
  	},
  }

```

# コマンド仕様の変更: `team info`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
+ 		},
+ 	},
  }

```

## 変更されたレポート: info

```
  &rc_doc.Report{
  	Name:    "info",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of team information.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "name", Desc: "The name of the team"}, &{Name: "team_id", Desc: "The ID of the team."}, &{Name: "num_licensed_users", Desc: "The number of licenses available to the team."}, &{Name: "num_provisioned_users", Desc: "The number of accounts that have been invited or are already active members of the team."}, &{Name: "policy_shared_folder_member", Desc: "Which shared folders team members can join (from_team_only, or from_anyone)"}, &{Name: "policy_shared_folder_join", Desc: "Who can join folders shared by team members (team, or anyone)"}, &{Name: "policy_shared_link_create", Desc: "Who can view shared links owned by team members (default_public, default_team_only, or team_only)"}, &{Name: "policy_emm_state", Desc: "This describes the Enterprise Mobility Management (EMM) state for this team (disabled, optional, or required)"}, &{Name: "policy_office_add_in", Desc: "The admin policy around the Dropbox Office Add-In for this team (disabled, or enabled)"}},
  }

```

# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: linked_app

```
  &rc_doc.Report{
  	Name: "linked_app",
- 	Desc: "",
+ 	Desc: "This report shows a list of linked app with the user of the app.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "app_id", Desc: "The application unique id."},
  		&{Name: "app_name", Desc: "The application name."},
  		&{Name: "is_app_folder", Desc: "Whether the linked application uses a dedicated folder."},
  		... // 3 identical elements
  	},
  }

```

# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "IncludeDeleted",
+ 			Desc:     "If true, deleted file or folder will be returned",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "If true, media info is set for photo and video in json report",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "If true, include team member folders",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "If true, include shared folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "If true, include team folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Desc:     "List only for the folder matched to the name",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the owning team member."},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "Set if the folder is contained by a shared folder.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Depth",
+ 			Desc:     "Report entry for all files and directories depth directories deep",
+ 			Default:  "\x01",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
+ 		},
+ 		&{
+ 			Name:     "IncludeAppFolder",
+ 			Desc:     "If true, include app folders",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "if true, include team member folders",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "If true, include shared folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "If true, include team folders",
+ 			Default:  "true",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Desc:     "List only for the folder matched to the name",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
- 		&{Name: "result.namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "result.namespace_id", Desc: "The ID of this namespace."},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
  		&{Name: "result.path", Desc: "Path to the folder"},
  		&{Name: "result.count_file", Desc: "Number of files under the folder"},
  		... // 4 identical elements
  	},
  }

```

# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: namespace

```
  &rc_doc.Report{
  	Name: "namespace",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning team member."},
  	},
  }

```

# コマンド仕様の変更: `team namespace member list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "AllColumns", Desc: "Show all columns", Default: "false", TypeName: "bool"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: namespace_member

```
  &rc_doc.Report{
  	Name: "namespace_member",
- 	Desc: "",
+ 	Desc: "This report shows a list of members of namespaces in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "entry_access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		... // 5 identical elements
  	},
  }

```

# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "Filter links by visibility (public/team_only/password)",
+ 			Default:  "public",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{
+ 				"options": []interface{}{
+ 					string("public"),
+ 					string("team_only"),
+ 					string("password"),
+ 					string("team_and_password"),
+ 					string("shared_folder_only"),
+ 				},
+ 			},
+ 		},
+ 	},
  }

```

## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
- 	Desc: "",
+ 	Desc: "This report shows a list of shared links with the shared link owner team member.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "At",
+ 			Desc:     "New expiration date and time",
+ 			TypeName: "domain.dropbox.model.mo_time.time_impl",
+ 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 		},
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "Days to the new expiration date",
+ 			Default:  "\x00",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "Target link visibility",
+ 			Default:  "public",
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{
+ 				"options": []interface{}{
+ 					string("public"),
+ 					string("team_only"),
+ 					string("password"),
+ 					string("team_and_password"),
+ 					string("shared_folder_only"),
+ 				},
+ 			},
+ 		},
+ 	},
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name:    "skipped",
- 	Desc:    "",
+ 	Desc:    "This report shows a list of shared links with the shared link owner team member.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "tag", Desc: "Entry type (file, or folder)"}, &{Name: "url", Desc: "URL of the shared link."}, &{Name: "name", Desc: "The linked file name (including extension)."}, &{Name: "expires", Desc: "Expiration time, if set."}, &{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."}, &{Name: "visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."}, &{Name: "email", Desc: "Email address of user."}, &{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"}, &{Name: "surname", Desc: "Surname of the link owner"}, &{Name: "given_name", Desc: "Given name of the link owner"}},
  }

```

## 変更されたレポート: updated

```
  &rc_doc.Report{
  	Name: "updated",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{
- 			Name: "input.shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "input.tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "input.url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "input.path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "input.visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
- 		&{Name: "input.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "input.email", Desc: "Email address of user."},
  		&{Name: "input.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "input.surname", Desc: "Surname of the link owner"},
  		&{Name: "input.given_name", Desc: "Given name of the link owner"},
- 		&{Name: "result.id", Desc: "A unique identifier for the linked file or folder"},
- 		&{Name: "result.tag", Desc: "Entry type (file, or folder)"},
- 		&{Name: "result.url", Desc: "URL of the shared link."},
- 		&{Name: "result.name", Desc: "The linked file name (including extension)."},
  		&{Name: "result.expires", Desc: "Expiration time, if set."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox.",
- 		},
- 		&{
- 			Name: "result.visibility",
- 			Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder).",
- 		},
  	},
  }

```

# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for a list of team folder names",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "Name of team folder"},
- 		&{Name: "result.team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "result.name", Desc: "The name of the team folder."},
  		&{Name: "result.status", Desc: "The status of the team folder (active, archived, or archive_in_progress)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `teamfolder batch permdelete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for a list of team folder names",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name:    "operation_log",
- 	Desc:    "",
+ 	Desc:    "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "status", Desc: "Status of the operation"}, &{Name: "reason", Desc: "Reason of failure or skipped operation"}, &{Name: "input.name", Desc: "Name of team folder"}},
  }

```

# コマンド仕様の変更: `teamfolder batch replication`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DstPeerName",
+ 			Desc:     "Destination team account alias",
+ 			Default:  "dst",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "Data file for a list of team folder names",
+ 			TypeName: "infra.feed.fd_file_impl.row_feed",
+ 		},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "Source team account alias",
+ 			Default:  "src",
+ 			TypeName: "string",
+ 		},
+ 	},
  }

```

## 変更されたレポート: verification

```
  &rc_doc.Report{
  	Name:    "verification",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing."}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, &{Name: "left_hash", Desc: "Content hash of left file"}, &{Name: "right_path", Desc: "path of right"}, &{Name: "right_kind", Desc: "folder of file"}, &{Name: "right_size", Desc: "size of right file"}, &{Name: "right_hash", Desc: "Content hash of right file"}},
  }

```

# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
- 	Desc: "",
+ 	Desc: "This report shows a list of namespaces in the team.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the owning team member."},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "The file size in bytes."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
- 		&{
- 			Name: "shared_folder_id",
- 			Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location.",
- 		},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "Set if the folder is contained by a shared folder.",
- 		},
  	},
  }

```

# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Depth",
+ 			Desc:     "Depth",
+ 			Default:  "\x01",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
- 	Desc: "",
+ 	Desc: "This report shows the transaction result.",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{
- 			Name: "input.team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
- 		&{Name: "result.namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "result.namespace_id", Desc: "The ID of this namespace."},
- 		&{
- 			Name: "result.namespace_type",
- 			Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)",
- 		},
- 		&{
- 			Name: "result.owner_team_member_id",
- 			Desc: "If this is a team member or app folder, the ID of the owning team member.",
- 		},
  		&{Name: "result.path", Desc: "Path to the folder"},
  		&{Name: "result.count_file", Desc: "Number of files under the folder"},
  		... // 4 identical elements
  	},
  }

```

# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

## 変更されたレポート: team_folder

```
  &rc_doc.Report{
  	Name: "team_folder",
- 	Desc: "",
+ 	Desc: "This report shows a list of team folders in the team.",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "name", Desc: "The name of the team folder."},
  		&{Name: "status", Desc: "The status of the team folder (active, archived, or archive_in_progress)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "default",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 		},
+ 	},
  }

```

# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "DstPeerName",
+ 			Desc:     "Destination team account alias",
+ 			Default:  "dst",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "Source team account alias",
+ 			Default:  "src",
+ 			TypeName: "string",
+ 		},
+ 	},
  }

```

## 変更されたレポート: verification

```
  &rc_doc.Report{
  	Name:    "verification",
- 	Desc:    "",
+ 	Desc:    "This report shows a difference between to folders.",
  	Columns: []*rc_doc.ReportColumn{&{Name: "diff_type", Desc: "Type of difference. `file_content_diff`: different content hash, `{left|right}_file_missing`: left or right file missing, `{left|right}_folder_missing`: left or right folder missing."}, &{Name: "left_path", Desc: "path of left"}, &{Name: "left_kind", Desc: "folder or file"}, &{Name: "left_size", Desc: "size of left file"}, &{Name: "left_hash", Desc: "Content hash of left file"}, &{Name: "right_path", Desc: "path of right"}, &{Name: "right_kind", Desc: "folder of file"}, &{Name: "right_size", Desc: "size of right file"}, &{Name: "right_hash", Desc: "Content hash of right file"}},
  }

```

# コマンド仕様の変更: `web`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  nil,
+ 	Values: []*rc_doc.Value{
+ 		&{
+ 			Name:     "Port",
+ 			Desc:     "Port number",
+ 			Default:  "Ṹ",
+ 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1024), "value": float64(7800)},
+ 		},
+ 	},
  }

```

