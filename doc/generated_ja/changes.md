# `リリース 63` から `リリース 64` までの変更点

# 追加されたコマンド

| コマンド            | タイトル             |
|---------------------|----------------------|
| file dispatch local | Dispatch local files |
| version             | Show version         |



# コマンド仕様の変更: `dev async`



## 変更されたレポート: rows

```
  &rc_doc.Report{
  	Name: "rows",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
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
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `dev diag procmon`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev diag procmon",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -repository-path /LOCAL/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev release publish",
- 	CliArgs:         "",
+ 	CliArgs:         "-artifact-path /LOCAL/PATH/TO/ARTIFACT",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test monkey",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/PROCESS",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file download",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/OF/FILE -local-path /LOCAL/PATH/TO/DOWNLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file export doc",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/FILE -local-path /LOCAL/PATH/TO/EXPORT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
  		&{Name: "server_modified", Desc: "The last time the file was modified on Dropbox."},
- 		&{
- 			Name: "revision",
- 			Desc: "A unique identifier for the current revision of a file.",
- 		},
  		&{Name: "size", Desc: "If this folder is a shared folder mount point, the ID of the shared folder mounted at this location."},
- 		&{Name: "content_hash", Desc: "A hash of the file content."},
  		&{Name: "export_name", Desc: "File name for export file."},
  		&{Name: "export_size", Desc: "File size of export file."},
- 		&{Name: "export_hash", Desc: "Content hash of export file."},
  	},
  }

```

# コマンド仕様の変更: `file import batch url`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.url", Desc: "Url to download"},
  		&{Name: "input.path", Desc: "Path to store file (use path given by `-path` when the record is empty)"},
- 		&{Name: "result.id", Desc: "A unique identifier for the file."},
  		&{Name: "result.tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "result.path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file import url`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file list`



## 変更されたレポート: file_list

```
  &rc_doc.Report{
  	Name: "file_list",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "file restore",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file sync preflight up`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync preflight up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file sync up",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file upload`



## 変更されたレポート: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

## 変更されたレポート: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "result.client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `file watch`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file watch",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/WATCH",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "filerequest create",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `group add`



## 変更されたレポート: added_group

```
  &rc_doc.Report{
  	Name: "added_group",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group batch delete`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.name", Desc: "Group name"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group list`



## 変更されたレポート: group

```
  &rc_doc.Report{
  	Name: "group",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group member add`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "Email address of the member"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group member delete`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.member_email", Desc: "Email address of the member"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `group member list`



## 変更されたレポート: group_member

```
  &rc_doc.Report{
  	Name: "group_member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `group rename`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.new_name", Desc: "New group name"},
  		&{Name: "result.group_name", Desc: "Name of a group"},
- 		&{Name: "result.group_id", Desc: "A group's identifier"},
  		&{Name: "result.group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "result.group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "result.member_count", Desc: "The number of members in the group."},
  	},
  }

```

# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job history ship",
- 	CliArgs:         "",
+ 	CliArgs:         "-dropbox-path /DROPBOX/PATH/TO/UPLOAD",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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

# コマンド仕様の変更: `job loop`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job loop",
- 	CliArgs:         "",
+ 	CliArgs:         `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 17:58:38"`,
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `job run`



## 設定が変更されたコマンド



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job run",
- 	CliArgs:         "",
+ 	CliArgs:         "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# コマンド仕様の変更: `member invite`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "Given name of the account"},
  		&{Name: "input.surname", Desc: "Surname of the account"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member list`



## 変更されたレポート: member

```
  &rc_doc.Report{
  	Name: "member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "member_folder_id", Desc: "The namespace id of the user's root folder."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  	},
  }

```

# コマンド仕様の変更: `member reinvite`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "input.email", Desc: "Email address of user."},
  		&{Name: "input.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "input.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "input.given_name", Desc: "Also known as a first name"},
  		&{Name: "input.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "input.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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
  		&{Name: "input.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "input.tag", Desc: "Operation tag"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member update email`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.from_email", Desc: "Current Email address"},
  		&{Name: "input.to_email", Desc: "New Email address"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member update externalid`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "input.email", Desc: "Email address of team members"},
  		&{Name: "input.external_id", Desc: "External ID of team members"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `member update profile`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "input.given_name", Desc: "Given name of the account"},
  		&{Name: "input.surname", Desc: "Surname of the account"},
- 		&{Name: "result.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "result.email", Desc: "Email address of user."},
  		&{Name: "result.email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "result.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "result.given_name", Desc: "Also known as a first name"},
  		&{Name: "result.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "result.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }

```

# コマンド仕様の変更: `sharedfolder list`



## 変更されたレポート: shared_folder

```
  &rc_doc.Report{
  	Name: "shared_folder",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder.",
- 		},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		... // 5 identical elements
  		&{Name: "policy_member", Desc: "Who can be a member of this shared folder, as set on the folder itself (team, or anyone)"},
  		&{Name: "policy_viewer_info", Desc: "Who can enable/disable viewer info for this shared folder."},
- 		&{Name: "owner_team_id", Desc: "Team ID of the team that owns the folder"},
  		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
  	},
  }

```

# コマンド仕様の変更: `sharedfolder member list`



## 変更されたレポート: member

```
  &rc_doc.Report{
  	Name: "member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
- 		&{
- 			Name: "parent_shared_folder_id",
- 			Desc: "The ID of the parent shared folder. This field is present only if the folder is contained within another shared folder.",
- 		},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "path_lower", Desc: "The lower-cased full path of this shared folder."},
  		... // 2 identical elements
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		&{Name: "is_inherited", Desc: "True if the member has access from a parent folder"},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "invitee_email", Desc: "Email address of invitee for this folder"},
  	},
  }

```

# コマンド仕様の変更: `sharedlink delete`



## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.id", Desc: "A unique identifier for the linked file or folder"},
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
- 	CliArgs:         "",
+ 	CliArgs:         "-url SHAREDLINK_URL",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  }

```

## 変更されたレポート: file_list

```
  &rc_doc.Report{
  	Name: "file_list",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
- 		&{
- 			Name: "path_lower",
- 			Desc: "The lowercased full path in the user's Dropbox. This always starts with a slash.",
- 		},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `sharedlink list`



## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the linked file or folder"},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 4 identical elements
  	},
  }

```

# コマンド仕様の変更: `team activity user`



## 変更されたレポート: user_summary

```
  &rc_doc.Report{
  	Name: "user_summary",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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

# コマンド仕様の変更: `team device list`



## 変更されたレポート: device

```
  &rc_doc.Report{
  	Name: "device",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }

```

# コマンド仕様の変更: `team device unlink`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 5 identical elements
  		&{Name: "input.given_name", Desc: "Also known as a first name"},
  		&{Name: "input.surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "input.familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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

# コマンド仕様の変更: `team diag explorer`



## 変更されたレポート: device

```
  &rc_doc.Report{
  	Name: "device",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 3 identical elements
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
  		&{Name: "device_tag", Desc: "Type of the session (web_session, desktop_client, or mobile_client)"},
  		&{Name: "id", Desc: "The session id."},
  		... // 16 identical elements
  	},
  }

```

## 変更されたレポート: file_request

```
  &rc_doc.Report{
  	Name: "file_request",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "account_id", Desc: "Account ID of this file request owner."},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ID of file request owner user as a member of a team",
- 		},
  		&{Name: "email", Desc: "Email address of this file request owner."},
  		&{Name: "status", Desc: "The user status of this file request owner (active/invited/suspended/removed)"},
  		&{Name: "surname", Desc: "Surname of this file request owner."},
  		&{Name: "given_name", Desc: "Given name of this file request owner."},
- 		&{Name: "file_request_id", Desc: "The ID of the file request."},
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
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "group_name", Desc: "Name of a group"},
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
- 		&{
- 			Name: "group_external_id",
- 			Desc: " External ID of group. This is an arbitrary ID that an admin can attach to a group.",
- 		},
  		&{Name: "member_count", Desc: "The number of members in the group."},
  	},
  }

```

## 変更されたレポート: group_member

```
  &rc_doc.Report{
  	Name: "group_member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "group_id", Desc: "A group's identifier"},
  		&{Name: "group_name", Desc: "Name of a group."},
  		&{Name: "group_management_type", Desc: "Who is allowed to manage the group (user_managed, company_managed, or system_managed)"},
  		&{Name: "access_type", Desc: "The role that the user has in the group (member/owner)"},
- 		&{Name: "account_id", Desc: "A user's account identifier"},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

## 変更されたレポート: linked_app

```
  &rc_doc.Report{
  	Name: "linked_app",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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

## 変更されたレポート: member

```
  &rc_doc.Report{
  	Name: "member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "email_verified", Desc: "Is true if the user's email is verified to be owned by the user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
- 		&{Name: "abbreviated_name", Desc: "An abbreviated form of the person's name."},
- 		&{Name: "member_folder_id", Desc: "The namespace id of the user's root folder."},
- 		&{Name: "external_id", Desc: "External ID that a team can attach to the user."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{
- 			Name: "persistent_id",
- 			Desc: "Persistent ID that a team can attach to the user. The persistent ID is unique ID to be used for SAML authentication.",
- 		},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, support_admin, or member_only)"},
  	},
  }

```

## 変更されたレポート: namespace

```
  &rc_doc.Report{
  	Name: "namespace",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning team member."},
  	},
  }

```

## 変更されたレポート: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the owning team member."},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

## 変更されたレポート: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
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

## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `team filerequest list`



## 変更されたレポート: file_request

```
  &rc_doc.Report{
  	Name: "file_request",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "account_id", Desc: "Account ID of this file request owner."},
- 		&{
- 			Name: "team_member_id",
- 			Desc: "ID of file request owner user as a member of a team",
- 		},
  		&{Name: "email", Desc: "Email address of this file request owner."},
  		&{Name: "status", Desc: "The user status of this file request owner (active/invited/suspended/removed)"},
  		&{Name: "surname", Desc: "Surname of this file request owner."},
  		&{Name: "given_name", Desc: "Given name of this file request owner."},
- 		&{Name: "file_request_id", Desc: "The ID of the file request."},
  		&{Name: "url", Desc: "The URL of the file request."},
  		&{Name: "title", Desc: "The title of the file request."},
  		... // 6 identical elements
  	},
  }

```

# コマンド仕様の変更: `team linkedapp list`



## 変更されたレポート: linked_app

```
  &rc_doc.Report{
  	Name: "linked_app",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		&{Name: "given_name", Desc: "Also known as a first name"},
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
- 		&{Name: "familiar_name", Desc: "Locale-dependent name"},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user's Dropbox account."},
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

# コマンド仕様の変更: `team namespace file list`



## 変更されたレポート: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the owning team member."},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `team namespace file size`



## 変更されたレポート: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
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

# コマンド仕様の変更: `team namespace list`



## 変更されたレポート: namespace

```
  &rc_doc.Report{
  	Name: "namespace",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning team member."},
  	},
  }

```

# コマンド仕様の変更: `team namespace member list`



## 変更されたレポート: namespace_member

```
  &rc_doc.Report{
  	Name: "namespace_member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "entry_access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		... // 5 identical elements
  	},
  }

```

# コマンド仕様の変更: `team sharedlink list`



## 変更されたレポート: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{
- 			Name: "shared_link_id",
- 			Desc: "A unique identifier for the linked file or folder",
- 		},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 2 identical elements
  		&{Name: "path_lower", Desc: "The lowercased full path in the user's Dropbox."},
  		&{Name: "visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."},
- 		&{Name: "account_id", Desc: "A user's account identifier."},
- 		&{Name: "team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "email", Desc: "Email address of user."},
  		&{Name: "status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `team sharedlink update expiry`



## 変更されたレポート: updated

```
  &rc_doc.Report{
  	Name: "updated",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
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
  		&{Name: "input.visibility", Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder)."},
- 		&{Name: "input.account_id", Desc: "A user's account identifier."},
- 		&{Name: "input.team_member_id", Desc: "ID of user as a member of a team."},
  		&{Name: "input.email", Desc: "Email address of user."},
  		&{Name: "input.status", Desc: "The user's status as a member of a specific team. (active/invited/suspended/removed)"},
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
- 			Desc: "The current visibility of the link after considering the shared links policies of the the team (in case the link's owner is part of a team) and the shared folder (in case the linked file is part of a shared folder).",
- 		},
  	},
  }

```

# コマンド仕様の変更: `teamfolder batch archive`



## 変更されたレポート: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "Name of team folder"},
- 		&{Name: "result.team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "result.name", Desc: "The name of the team folder."},
  		&{Name: "result.status", Desc: "The status of the team folder (active, archived, or archive_in_progress)"},
  		... // 2 identical elements
  	},
  }

```

# コマンド仕様の変更: `teamfolder file list`



## 変更されたレポート: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
  		&{Name: "namespace_member_email", Desc: "If this is a team member or app folder, the email address of the owning team member."},
- 		&{Name: "file_id", Desc: "A unique identifier for the file."},
  		&{Name: "tag", Desc: "Type of entry. `file`, `folder`, or `deleted`"},
  		&{Name: "name", Desc: "The last component of the path (including extension)."},
  		&{Name: "path_display", Desc: "The cased path to be used for display purposes only."},
  		&{Name: "client_modified", Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox."},
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

# コマンド仕様の変更: `teamfolder file size`



## 変更されたレポート: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
  		&{Name: "input.name", Desc: "The name of this namespace"},
- 		&{Name: "input.namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "input.namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
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

# コマンド仕様の変更: `teamfolder list`



## 変更されたレポート: team_folder

```
  &rc_doc.Report{
  	Name: "team_folder",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "name", Desc: "The name of the team folder."},
  		&{Name: "status", Desc: "The status of the team folder (active, archived, or archive_in_progress)"},
  		... // 2 identical elements
  	},
  }

```

