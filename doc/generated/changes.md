# Changes between `Release 63` to `Release 64`

# Commands added

| Command             | Title                |
|---------------------|----------------------|
| file dispatch local | Dispatch local files |



# Command spec changed: `dev async`



## Changed report: rows

```
  &rc_doc.Report{
  	Name: "rows",
  	Desc: "",
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

# Command spec changed: `dev ci artifact up`



## Changed report: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
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

## Changed report: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
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

# Command spec changed: `file download`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `file export doc`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `file import batch url`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `file import url`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `file list`



## Changed report: file_list

```
  &rc_doc.Report{
  	Name: "file_list",
  	Desc: "",
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

# Command spec changed: `file restore`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `file sync preflight up`



## Changed report: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
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

## Changed report: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
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

# Command spec changed: `file sync up`



## Changed report: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
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

## Changed report: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
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

# Command spec changed: `file upload`



## Changed report: skipped

```
  &rc_doc.Report{
  	Name: "skipped",
  	Desc: "",
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

## Changed report: uploaded

```
  &rc_doc.Report{
  	Name: "uploaded",
  	Desc: "",
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

# Command spec changed: `group add`



## Changed report: added_group

```
  &rc_doc.Report{
  	Name: "added_group",
  	Desc: "",
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

# Command spec changed: `group batch delete`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `group list`



## Changed report: group

```
  &rc_doc.Report{
  	Name: "group",
  	Desc: "",
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

# Command spec changed: `group member add`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `group member delete`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `group member list`



## Changed report: group_member

```
  &rc_doc.Report{
  	Name: "group_member",
  	Desc: "",
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

# Command spec changed: `group rename`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `job history ship`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `member invite`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `member list`



## Changed report: member

```
  &rc_doc.Report{
  	Name: "member",
  	Desc: "",
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

# Command spec changed: `member reinvite`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `member update email`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `member update externalid`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `member update profile`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `sharedfolder list`



## Changed report: shared_folder

```
  &rc_doc.Report{
  	Name: "shared_folder",
  	Desc: "",
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

# Command spec changed: `sharedfolder member list`



## Changed report: member

```
  &rc_doc.Report{
  	Name: "member",
  	Desc: "",
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

# Command spec changed: `sharedlink delete`



## Changed report: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
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

# Command spec changed: `sharedlink file list`



## Changed report: file_list

```
  &rc_doc.Report{
  	Name: "file_list",
  	Desc: "",
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

# Command spec changed: `sharedlink list`



## Changed report: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "id", Desc: "A unique identifier for the linked file or folder"},
  		&{Name: "tag", Desc: "Entry type (file, or folder)"},
  		&{Name: "url", Desc: "URL of the shared link."},
  		... // 4 identical elements
  	},
  }

```

# Command spec changed: `team activity user`



## Changed report: user_summary

```
  &rc_doc.Report{
  	Name: "user_summary",
  	Desc: "",
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

# Command spec changed: `team device list`



## Changed report: device

```
  &rc_doc.Report{
  	Name: "device",
  	Desc: "",
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

# Command spec changed: `team device unlink`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `team diag explorer`



## Changed report: device

```
  &rc_doc.Report{
  	Name: "device",
  	Desc: "",
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

## Changed report: file_request

```
  &rc_doc.Report{
  	Name: "file_request",
  	Desc: "",
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

## Changed report: group

```
  &rc_doc.Report{
  	Name: "group",
  	Desc: "",
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

## Changed report: group_member

```
  &rc_doc.Report{
  	Name: "group_member",
  	Desc: "",
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

## Changed report: linked_app

```
  &rc_doc.Report{
  	Name: "linked_app",
  	Desc: "",
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

## Changed report: member

```
  &rc_doc.Report{
  	Name: "member",
  	Desc: "",
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

## Changed report: namespace

```
  &rc_doc.Report{
  	Name: "namespace",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning team member."},
  	},
  }

```

## Changed report: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
  	Desc: "",
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

## Changed report: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
  	Desc: "",
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

## Changed report: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
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

# Command spec changed: `team filerequest list`



## Changed report: file_request

```
  &rc_doc.Report{
  	Name: "file_request",
  	Desc: "",
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

# Command spec changed: `team linkedapp list`



## Changed report: linked_app

```
  &rc_doc.Report{
  	Name: "linked_app",
  	Desc: "",
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

# Command spec changed: `team namespace file list`



## Changed report: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
  	Desc: "",
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

# Command spec changed: `team namespace file size`



## Changed report: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
  	Desc: "",
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

# Command spec changed: `team namespace list`



## Changed report: namespace

```
  &rc_doc.Report{
  	Name: "namespace",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "team_member_id", Desc: "If this is a team member or app folder, the ID of the owning team member."},
  	},
  }

```

# Command spec changed: `team namespace member list`



## Changed report: namespace_member

```
  &rc_doc.Report{
  	Name: "namespace_member",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		&{Name: "namespace_name", Desc: "The name of this namespace"},
- 		&{Name: "namespace_id", Desc: "The ID of this namespace."},
  		&{Name: "namespace_type", Desc: "The type of this namespace (app_folder, shared_folder, team_folder, or team_member_folder)"},
  		&{Name: "entry_access_type", Desc: "The current user's access level for this shared file/folder (owner, editor, viewer, or viewer_no_comment)"},
  		... // 5 identical elements
  	},
  }

```

# Command spec changed: `team sharedlink list`



## Changed report: shared_link

```
  &rc_doc.Report{
  	Name: "shared_link",
  	Desc: "",
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

# Command spec changed: `team sharedlink update expiry`



## Changed report: updated

```
  &rc_doc.Report{
  	Name: "updated",
  	Desc: "",
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

# Command spec changed: `teamfolder batch archive`



## Changed report: operation_log

```
  &rc_doc.Report{
  	Name: "operation_log",
  	Desc: "",
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

# Command spec changed: `teamfolder file list`



## Changed report: namespace_file

```
  &rc_doc.Report{
  	Name: "namespace_file",
  	Desc: "",
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

# Command spec changed: `teamfolder file size`



## Changed report: namespace_size

```
  &rc_doc.Report{
  	Name: "namespace_size",
  	Desc: "",
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

# Command spec changed: `teamfolder list`



## Changed report: team_folder

```
  &rc_doc.Report{
  	Name: "team_folder",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
- 		&{Name: "team_folder_id", Desc: "The ID of the team folder."},
  		&{Name: "name", Desc: "The name of the team folder."},
  		&{Name: "status", Desc: "The status of the team folder (active, archived, or archive_in_progress)"},
  		... // 2 identical elements
  	},
  }

```

