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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "RunConcurrently",
+ 			Desc:     "false",
+ 			Default:  "run concurrently",
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
+ 			Default:  "Dropbox path to upload",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local path to upload",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "PeerName", Desc: "deploy", Default: "Account alias", TypeName: "string"},
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
+ 			Desc:     "https://www.dropbox.com/download?full=1&os=win",
+ 			Default:  "Installer download URL",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Silent", Desc: "false", Default: "Use Silent installer", TypeName: "bool"},
+ 		&{
+ 			Name:     "SilentNoLaunch",
+ 			Desc:     "false",
+ 			Default:  "Use Enterprise installer",
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
+ 			Desc:     "0",
+ 			Default:  "Try stopping the app after given seconds.",
+ 			TypeName: "int",
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
+ 			Desc:     "false",
+ 			Default:  "True when unsuspend Updater",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "UpdaterName",
+ 			Desc:     "DropboxUpdate.exe",
+ 			Default:  "Executable name of Dropbox Updater",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "UpdaterPath",
+ 			Desc:     "C:/Program Files (x86)/Dropbox/Update",
+ 			Default:  "Path to Dropbox Updater",
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
+ 			Default:  "Dropbox path to upload procmon logs",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "ProcmonUrl",
+ 			Desc:     "https://download.sysinternals.com/files/ProcessMonitor.zip",
+ 			Default:  "Process monitor download url",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "RepositoryPath",
+ 			Default:  "Procmon Work directory",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "RetainLogs",
+ 			Desc:     "4",
+ 			Default:  "Number of Procmon logs to retain",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "RunUntil",
+ 			Default:  "Skip run after this date/time",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "1800",
+ 			Default:  "Duration for waiting procmon",
+ 			TypeName: "int",
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
+ 			Desc:     "true",
+ 			Default:  "Include badges of build status",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "CommandPath",
+ 			Desc:     "doc/generated/",
+ 			Default:  "Relative path to generate command manuals",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Filename", Desc: "README.md", Default: "Filename", TypeName: "string"},
+ 		&{Name: "Lang", Default: "Language", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "MarkdownReadme",
+ 			Desc:     "false",
+ 			Default:  "Generate README as markdown format",
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
+ 		&{Name: "Dest", Default: "Dummy file destination", TypeName: "string"},
+ 		&{Name: "MaxEntry", Desc: "0", Default: "Maximum entries", TypeName: "int"},
+ 		&{Name: "Path", Default: "Path to dummy entry file", TypeName: "string"},
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
+ 	Values:  []*rc_doc.Value{&{Name: "Text", Default: "Text to echo", TypeName: "string"}},
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
+ 			Desc:     "test/dev/resource.json",
+ 			Default:  "Path to the test resource location",
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
+ 			Default:  "Path to artifacts",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "Branch", Desc: "master", Default: "Target branch", TypeName: "string"},
+ 		&{
+ 			Name:     "ConnGithub",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "TestResource",
+ 			Desc:     "test/dev/resource.json",
+ 			Default:  "Path to test resource",
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
+ 		&{Name: "FilePath", Default: "File path to output", TypeName: "reflect.rtype"},
+ 		&{Name: "Lang", Default: "Language", TypeName: "reflect.rtype"},
+ 		&{Name: "Release1", Default: "Release name 1", TypeName: "reflect.rtype"},
+ 		&{Name: "Release2", Default: "Release name 2", TypeName: "reflect.rtype"},
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
+ 		&{Name: "FilePath", Default: "File path", TypeName: "reflect.rtype"},
+ 		&{Name: "Lang", Default: "Language", TypeName: "reflect.rtype"},
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
+ 			Desc:     "10000",
+ 			Default:  "Number of files/folder distribution",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "Extension",
+ 			Desc:     "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,bmp,wmi,ini,ai,psd",
+ 			Default:  "File extensions (comma separated)",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Path", Default: "Monkey test path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Seconds",
+ 			Desc:     "10",
+ 			Default:  "Monkey test duration in seconds",
+ 			TypeName: "int",
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
+ 		&{Name: "All", Desc: "false", Default: "Test all recipes", TypeName: "bool"},
+ 		&{Name: "Recipe", Default: "Recipe name to test", TypeName: "reflect.rtype"},
+ 		&{Name: "Resource", Default: "Test resource file path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Verbose",
+ 			Desc:     "false",
+ 			Default:  "Verbose output for testing",
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
+ 		&{Name: "BufferSize", Desc: "65536", Default: "Size of buffer", TypeName: "int"},
+ 		&{Name: "Record", Default: "Capture record(s) for the test", TypeName: "string"},
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
+ 	Values:  []*rc_doc.Value{&{Name: "Seconds", Desc: "1", Default: "Wait seconds", TypeName: "int"}},
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
+ 			Desc:     "left",
+ 			Default:  "Account alias (left)",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "LeftPath",
+ 			Default:  "The path from account root (left)",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Right",
+ 			Desc:     "right",
+ 			Default:  "Account alias (right)",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "RightPath",
+ 			Default:  "The path from account root (right)",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "DropboxPath", Default: "Dropbox path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local path",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Dst", Default: "Destination path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Src", Default: "Source path", TypeName: "reflect.rtype"},
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
+ 		&{Name: "Path", Default: "Path to delete", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "DropboxPath", Default: "File path to download", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local path to download",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Dropbox document path to export.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local path to save",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{Name: "Path", Default: "Path to import", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Path", Default: "Path to import", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Url", Default: "URL", TypeName: "string"},
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
+ 			Desc:     "false",
+ 			Default:  "Include deleted files",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "false",
+ 			Default:  "Include media information",
+ 			TypeName: "bool",
+ 		},
+ 		&{Name: "Path", Default: "Path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Recursive", Desc: "false", Default: "List recursively", TypeName: "bool"},
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
+ 		&{Name: "DryRun", Desc: "true", Default: "Dry run", TypeName: "bool"},
+ 		&{Name: "From", Default: "Path for merge", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "KeepEmptyFolder",
+ 			Desc:     "false",
+ 			Default:  "Keep empty folder after merge",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "To", Default: "Path to merge", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "WithinSameNamespace",
+ 			Desc:     "false",
+ 			Default:  "Do not cross namespace. That is for preserve sharing permission including a shared link",
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
+ 		&{Name: "Dst", Default: "Destination path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Src", Default: "Source path", TypeName: "reflect.rtype"},
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
+ 			Desc:     "dst",
+ 			Default:  "Account alias (destionation)",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "DstPath", Default: "Destination path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "src",
+ 			Default:  "Account alias (source)",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "SrcPath", Default: "Source path", TypeName: "reflect.rtype"},
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
+ 		&{Name: "Path", Default: "Path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Restricts search to only the file categories specified (image/document/pdf/spreadsheet/presentation/audio/video/folder/paper/others).",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Restricts search to only the extensions specified.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Default:  "Scopes the search to a path in the user's Dropbox.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Query", Default: "The string to search for.", TypeName: "string"},
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
+ 			Default:  "Restricts search to only the file categories specified (image/document/pdf/spreadsheet/presentation/audio/video/folder/paper/others).",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Restricts search to only the extensions specified.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Default:  "Scopes the search to a path in the user's Dropbox.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Query", Default: "The string to search for.", TypeName: "string"},
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
+ 			Default:  "Destination Dropbox path",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local file path",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "153600",
+ 			Default:  "Upload chunk size in KB",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Default:  "Destination Dropbox path",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local file path",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "153600",
+ 			Default:  "Upload chunk size in KB",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "DropboxPath",
+ 			Default:  "Destination Dropbox path",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Default:  "Local file path",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "Overwrite",
+ 			Desc:     "false",
+ 			Default:  "Overwrite existing file(s)",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Path", Default: "Path to watch", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "false",
+ 			Default:  "Watch path recursively",
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
+ 			Default:  "If set, allow uploads after the deadline has passed (one_day/two_days/seven_days/thirty_days/always)",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Deadline",
+ 			Default:  "The deadline for this file request.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Path",
+ 			Default:  "The path for the folder in the Dropbox where uploaded files will be sent.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Title", Default: "The title of the file request", TypeName: "string"},
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "false",
+ 			Default:  "Force delete the file request.",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Url", Default: "URL of the file request.", TypeName: "string"},
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "company_managed",
+ 			Default:  "Group management type `company_managed` or `user_managed`",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("company_managed"), string("user_managed")}},
+ 		},
+ 		&{Name: "Name", Default: "Group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Data file for group name list",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Name", Default: "Group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "GroupName", Default: "Group name", TypeName: "string"},
+ 		&{Name: "MemberEmail", Default: "Email address of the member", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "GroupName", Default: "Name of the group", TypeName: "string"},
+ 		&{Name: "MemberEmail", Default: "Email address of the member", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "CurrentName", Default: "Current group name", TypeName: "string"},
+ 		&{Name: "NewName", Default: "New group name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 	Values:  []*rc_doc.Value{&{Name: "Days", Desc: "7", Default: "Target days old", TypeName: "int"}},
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
+ 	Values:  []*rc_doc.Value{&{Name: "Days", Desc: "28", Default: "Target days old", TypeName: "int"}},
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
+ 			Default:  "Dropbox path to upload",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "180",
+ 			Default:  "Interval seconds",
+ 			TypeName: "int",
+ 		},
+ 		&{Name: "QuitOnError", Desc: "false", Default: "Quit on error", TypeName: "bool"},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Default:  "Path to runbook",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{Name: "Until", Default: "Run until this date/time", TypeName: "reflect.rtype"},
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
+ 			Desc:     "false",
+ 			Default:  "Fork process on running the workflow",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "RunbookPath",
+ 			Default:  "Path to the runbook",
+ 			TypeName: "reflect.rtype",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name:     "TimeoutSeconds",
+ 			Desc:     "0",
+ 			Default:  "Terminate process when given time passed",
+ 			TypeName: "int",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "WipeData",
+ 			Desc:     "true",
+ 			Default:  "If true, controls if the user's data will be deleted on their linked devices",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "RevokeTeamShares",
+ 			Desc:     "false",
+ 			Default:  "True for revoke shared folder access which owned by the team",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "SilentInvite",
+ 			Desc:     "false",
+ 			Default:  "Do not send welcome email (requires SSO + domain verification instead)",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Quota",
+ 			Desc:     "0",
+ 			Default:  "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom quota set.",
+ 			TypeName: "int",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Silent",
+ 			Desc:     "false",
+ 			Default:  "Do not send welcome email (SSO required)",
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
+ 			Desc:     "dst",
+ 			Default:  "Destination team; team file access",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Src",
+ 			Desc:     "src",
+ 			Default:  "Source team; team file access",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "UpdateUnverified",
+ 			Desc:     "false",
+ 			Default:  "Update an account which didn't verified email. If an account email unverified, email address change may affect lose invitation to folders.",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Expiration date/time of the shared link",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Password", Default: "Password", TypeName: "reflect.rtype"},
+ 		&{Name: "Path", Default: "Path", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "TeamOnly",
+ 			Desc:     "false",
+ 			Default:  "Link is accessible only by team members",
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
+ 			Default:  "File or folder path to remove shared link",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Recursive",
+ 			Desc:     "false",
+ 			Default:  "Attempt to remove the file hierarchy",
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
+ 			Default:  "Password for the shared link",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "Url", Default: "Shared link URL", TypeName: "reflect.rtype"},
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Filter the returned events to a single category. This field is optional.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "EndTime", Default: "Ending time (exclusive).", TypeName: "reflect.rtype"},
+ 		&{Name: "File", Default: "User email address list file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Default:  "Starting time (inclusive)",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Category", Default: "Event category", TypeName: "reflect.rtype"},
+ 		&{Name: "EndDate", Default: "End date", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "StartDate", Default: "Start date", TypeName: "string"},
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
+ 			Default:  "Filter the returned events to a single category. This field is optional.",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "EndTime", Default: "Ending time (exclusive).", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Default:  "Starting time (inclusive)",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Filter the returned events to a single category. This field is optional.",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{Name: "EndTime", Default: "Ending time (exclusive).", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "StartTime",
+ 			Default:  "Starting time (inclusive)",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "false",
+ 			Default:  "Delete files on unlink",
+ 			TypeName: "bool",
+ 		},
+ 		&{Name: "File", Default: "Data file", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "false",
+ 			Default:  "Include additional reports",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Desc:     "default",
+ 			Default:  "Dropbox Business file access",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Info",
+ 			Desc:     "default",
+ 			Default:  "Dropbox Business information access",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Mgmt",
+ 			Desc:     "default",
+ 			Default:  "Dropbox Business management",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "false",
+ 			Default:  "If true, deleted file or folder will be returned",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMediaInfo",
+ 			Desc:     "false",
+ 			Default:  "If true, media info is set for photo and video in json report",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "false",
+ 			Default:  "If true, include team member folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "true",
+ 			Default:  "If true, include shared folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "true",
+ 			Default:  "If true, include team folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Default:  "List only for the folder matched to the name",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "1",
+ 			Default:  "Report entry for all files and directories depth directories deep",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "IncludeAppFolder",
+ 			Desc:     "false",
+ 			Default:  "If true, include app folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMemberFolder",
+ 			Desc:     "false",
+ 			Default:  "if true, include team member folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeSharedFolder",
+ 			Desc:     "true",
+ 			Default:  "If true, include shared folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeTeamFolder",
+ 			Desc:     "true",
+ 			Default:  "If true, include team folders",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "Name",
+ 			Default:  "List only for the folder matched to the name",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "AllColumns", Desc: "false", Default: "Show all columns", TypeName: "bool"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "public",
+ 			Default:  "Filter links by visibility (public/team_only/password)",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "At", Default: "New expiration date and time", TypeName: "reflect.rtype"},
+ 		&{
+ 			Name:     "Days",
+ 			Desc:     "0",
+ 			Default:  "Days to the new expiration date",
+ 			TypeName: "int",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Visibility",
+ 			Desc:     "public",
+ 			Default:  "Target link visibility",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Name", Default: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Data file for a list of team folder names",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Default:  "Data file for a list of team folder names",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "dst",
+ 			Default:  "Destination team account alias",
+ 			TypeName: "string",
+ 		},
+ 		&{
+ 			Name:     "File",
+ 			Default:  "Data file for a list of team folder names",
+ 			TypeName: "reflect.rtype",
+ 		},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "src",
+ 			Default:  "Source team account alias",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Depth", Desc: "1", Default: "Depth", TypeName: "int"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 		&{Name: "Name", Default: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "default",
+ 			Default:  "Account alias",
+ 			TypeName: "reflect.rtype",
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
+ 			Desc:     "dst",
+ 			Default:  "Destination team account alias",
+ 			TypeName: "string",
+ 		},
+ 		&{Name: "Name", Default: "Team folder name", TypeName: "string"},
+ 		&{
+ 			Name:     "SrcPeerName",
+ 			Desc:     "src",
+ 			Default:  "Source team account alias",
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
+ 	Values:  []*rc_doc.Value{&{Name: "Port", Desc: "7800", Default: "Port number", TypeName: "int"}},
  }

```

