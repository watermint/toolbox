# Changes between `Release 72` to `Release 73`

# Commands added


| Command                               | Title                                 |
|---------------------------------------|---------------------------------------|
| dev benchmark upload                  | Upload benchmark                      |
| dev build catalogue                   | Generate catalogue                    |
| dev build doc                         | Document generator                    |
| dev build license                     | Generate LICENSE.txt                  |
| dev build preflight                   | Process prerequisites for the release |
| dev build readme                      | Generate README.txt                   |
| dev test async                        | Async framework test                  |
| dev test echo                         | Echo text                             |
| file size                             | Storage usage                         |
| file sync down                        | Downstream sync with Dropbox          |
| file sync online                      | Sync online files                     |
| services asana team list              | List team                             |
| services asana team project list      | List projects of the team             |
| services asana team task list         | List task of the team                 |
| services asana workspace list         | List workspaces                       |
| services asana workspace project list | List projects of the workspace        |



# Commands deleted


| Command                | Title                                 |
|------------------------|---------------------------------------|
| dev async              | Async framework test                  |
| dev catalogue          | Generate catalogue                    |
| dev doc                | Document generator                    |
| dev echo               | Echo text                             |
| dev preflight          | Process prerequisites for the release |
| file sync preflight up | Upstream sync preflight check         |
| file upload            | Upload file                           |



# Command spec changed: `dev ci artifact up`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "Local path to upload",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string", ...},
  		&{Name: "Timeout", Desc: "Operation timeout in seconds", Default: "30", TypeName: "int", ...},
  	},
  }
```
## Added report(s)


| Name    | Description |
|---------|-------------|
| deleted | Path        |


## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.file", Desc: "Local file path"},
- 		&{Name: "input.size", Desc: "Local file size"},
- 		&{
- 			Name: "result.name",
- 			Desc: "The last component of the path (including extension).",
- 		},
- 		&{
- 			Name: "result.path_display",
- 			Desc: "The cased path to be used for display purposes only.",
- 		},
- 		&{
- 			Name: "result.client_modified",
- 			Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox.",
- 		},
- 		&{
- 			Name: "result.server_modified",
- 			Desc: "The last time the file was modified on Dropbox.",
- 		},
- 		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
+ 		&{Name: "input.entry_path", Desc: "Path"},
  	},
  }
```
## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "upload_start", Desc: "Time of start uploading"},
+ 		&{Name: "start", Desc: "Time of start"},
- 		&{Name: "upload_end", Desc: "Time of finish uploading"},
+ 		&{Name: "end", Desc: "Time of finish"},
  		&{Name: "num_bytes", Desc: "Total upload size (Bytes)"},
  		&{Name: "num_files_error", Desc: "The number of files failed or got an error."},
- 		&{Name: "num_files_upload", Desc: "The number of files uploaded or to upload."},
+ 		&{Name: "num_files_transferred", Desc: "The number of files uploaded/downloaded."},
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
+ 		&{Name: "num_folder_created", Desc: "Number of created folders."},
+ 		&{Name: "num_delete", Desc: "Number of deleted entry."},
  		&{Name: "num_api_call", Desc: "The number of estimated upload API call for upload."},
  	},
  }
```
## Changed report: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.file", Desc: "Local file path"},
- 		&{Name: "input.size", Desc: "Local file size"},
+ 		&{Name: "input.path", Desc: "Path"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `dev diag endpoint`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "JobId",
  			Desc:     "Job Id to diagnosis",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Path to the workspace",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `dev diag throughput`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "EndpointNamePrefix", Desc: "Filter by endpoint. Filter by name match to the prefix."},
  		&{Name: "EndpointNameSuffix", Desc: "Filter by endpoint. Filter by name match to the suffix."},
  		&{
  			Name:     "JobId",
  			Desc:     "Specify Job ID",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Path to workspace",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "TimeFormat", Desc: "Time format in go's time format", Default: "2006-01-02 15:04:05.999", TypeName: "string", ...},
  	},
  }
```
# Command spec changed: `dev kvs dump`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Path",
  			Desc:     "Path to KVS data",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  	},
  }
```
# Command spec changed: `dev release publish`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ArtifactPath",
  			Desc:     "Path to artifacts",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Branch", Desc: "Target branch", Default: "master", TypeName: "string", ...},
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...},
  	},
  }
```
# Command spec changed: `dev spec diff`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "FilePath",
  			Desc:     "File path to output",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Lang",
  			Desc:     "Language",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Release1",
  			Desc:     "Release name 1",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Release2",
  			Desc:     "Release name 2",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `dev spec doc`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "FilePath",
  			Desc:     "File path",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Lang",
  			Desc:     "Language",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `dev test monkey`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Distribution",
  			Desc:     "Number of files/folder distribution",
  			Default:  "10000",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(10000)},
  		},
  		&{Name: "Extension", Desc: "File extensions (comma separated)", Default: "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"..., TypeName: "string", ...},
  		&{Name: "Path", Desc: "Monkey test path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{
  			Name:     "Seconds",
  			Desc:     "Monkey test duration in seconds",
  			Default:  "10",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(86400), "min": float64(1), "value": float64(10)},
  		},
  	},
  }
```
# Command spec changed: `dev test recipe`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool", ...},
  		&{Name: "NoTimeout", Desc: "Do not timeout running recipe tests", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Single",
  			Desc:     "Recipe name to test",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Verbose", Desc: "Verbose output for testing", Default: "false", TypeName: "bool", ...},
  	},
  }
```
# Command spec changed: `dev util anonymise`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "JobIdName", Desc: "Filter by job id name Filter by exact match to the name."},
  		&{Name: "JobIdNamePrefix", Desc: "Filter by job id name Filter by name match to the prefix."},
  		&{Name: "JobIdNameSuffix", Desc: "Filter by job id name Filter by name match to the suffix."},
  		&{
  			Name:     "Path",
  			Desc:     "Path to the workspace",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `dev util curl`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BufferSize",
  			Desc:     "Size of buffer",
  			Default:  "65536",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.097152e+06), "min": float64(1024), "value": float64(65536)},
  		},
  		&{
  			Name:     "Record",
  			Desc:     "Capture record(s) for the test",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `dev util image jpeg`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Count",
  			Desc:     "Number of files to generate",
  			Default:  "10",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(32767), "min": float64(1), "value": float64(10)},
  		},
  		&{
  			Name:     "Height",
  			Desc:     "Height",
  			Default:  "1080",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1), "value": float64(1080)},
  		},
  		&{Name: "NamePrefix", Desc: "Filename prefix", Default: "test_image", TypeName: "string", ...},
  		&{
  			Name:     "Path",
  			Desc:     "Path to generate files",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{
  			Name:     "Quality",
  			Desc:     "Quality of jpeg",
  			Default:  "75",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(100), "min": float64(1), "value": float64(75)},
  		},
  		&{Name: "Seed", Desc: "Random seed", Default: "1", TypeName: "int", ...},
  		&{
  			Name:     "Width",
  			Desc:     "Width",
  			Default:  "1920",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(65535), "min": float64(1), "value": float64(1920)},
  		},
  	},
  }
```
# Command spec changed: `dev util wait`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Seconds",
  			Desc:     "Wait seconds",
  			Default:  "1",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(604800), "min": float64(1), "value": float64(1)},
  		},
  	},
  }
```
# Command spec changed: `file compare local`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "Local path",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  }
```
# Command spec changed: `file download`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "File path to download", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "Local path to download",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  	},
  }
```
# Command spec changed: `file export doc`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "DropboxPath", Desc: "Dropbox document path to export.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "LocalPath",
  			Desc:     "Local path to save",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
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
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{
  			Name:     "Path",
  			Desc:     "Path to import",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
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
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "IncludeMediaInfo",
- 			Desc:     "Include media information",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Restricts search to only the file categories specified (image/do"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{
  			Name:     "Extension",
  			Desc:     "Restricts search to only the extensions specified.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Scopes the search to a path in the user's Dropbox.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
  	},
  }
```
# Command spec changed: `file search name`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Restricts search to only the file categories specified (image/do"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}},
  		},
  		&{
  			Name:     "Extension",
  			Desc:     "Restricts search to only the extensions specified.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Scopes the search to a path in the user's Dropbox.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Query", Desc: "The string to search for.", TypeName: "string"},
  	},
  }
```
# Command spec changed: `file sync up`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ChunkSizeKb",
  			Desc:     "Upload chunk size in KB",
  			Default:  "65536",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
  		},
+ 		&{
+ 			Name:     "Delete",
+ 			Desc:     "Delete Dropbox file if a file removed locally",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
+ 		&{
+ 			Name:     "LocalPath",
+ 			Desc:     "Local file path",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
+ 		},
+ 		&{
+ 			Name: "NameDisableIgnore",
+ 			Desc: "Filter by name. Filter system file or ignore files.",
+ 		},
+ 		&{Name: "NameName", Desc: "Filter by name. Filter by exact match to the name."},
- 		&{
- 			Name:     "FailOnError",
- 			Desc:     "Returns error when any error happens while the operation. This command will not return any error when this flag is not enabled. "...,
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
+ 		&{
+ 			Name: "NameNamePrefix",
+ 			Desc: "Filter by name. Filter by name match to the prefix.",
+ 		},
- 		&{
- 			Name:     "LocalPath",
- 			Desc:     "Local file path",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
- 			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
- 		},
+ 		&{
+ 			Name: "NameNameSuffix",
+ 			Desc: "Filter by name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "SkipExisting",
+ 			Desc:     "Skip existing files. Do not overwrite",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "WorkPath",
+ 			Desc:     "Temporary path",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  	},
  }
```
## Added report(s)


| Name    | Description |
|---------|-------------|
| deleted | Path        |


## Changed report: skipped

```
  &dc_recipe.Report{
  	Name: "skipped",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.file", Desc: "Local file path"},
- 		&{Name: "input.size", Desc: "Local file size"},
- 		&{
- 			Name: "result.name",
- 			Desc: "The last component of the path (including extension).",
- 		},
- 		&{
- 			Name: "result.path_display",
- 			Desc: "The cased path to be used for display purposes only.",
- 		},
- 		&{
- 			Name: "result.client_modified",
- 			Desc: "For files, this is the modification time set by the desktop client when the file was added to Dropbox.",
- 		},
- 		&{
- 			Name: "result.server_modified",
- 			Desc: "The last time the file was modified on Dropbox.",
- 		},
- 		&{Name: "result.size", Desc: "The file size in bytes."},
- 		&{Name: "result.content_hash", Desc: "A hash of the file content."},
+ 		&{Name: "input.entry_path", Desc: "Path"},
  	},
  }
```
## Changed report: summary

```
  &dc_recipe.Report{
  	Name: "summary",
  	Desc: "This report shows a summary of the upload results.",
  	Columns: []*dc_recipe.ReportColumn{
- 		&{Name: "upload_start", Desc: "Time of start uploading"},
+ 		&{Name: "start", Desc: "Time of start"},
- 		&{Name: "upload_end", Desc: "Time of finish uploading"},
+ 		&{Name: "end", Desc: "Time of finish"},
  		&{Name: "num_bytes", Desc: "Total upload size (Bytes)"},
  		&{Name: "num_files_error", Desc: "The number of files failed or got an error."},
- 		&{Name: "num_files_upload", Desc: "The number of files uploaded or to upload."},
+ 		&{Name: "num_files_transferred", Desc: "The number of files uploaded/downloaded."},
  		&{Name: "num_files_skip", Desc: "The number of files skipped or to skip."},
+ 		&{Name: "num_folder_created", Desc: "Number of created folders."},
+ 		&{Name: "num_delete", Desc: "Number of deleted entry."},
  		&{Name: "num_api_call", Desc: "The number of estimated upload API call for upload."},
  	},
  }
```
## Changed report: uploaded

```
  &dc_recipe.Report{
  	Name: "uploaded",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "status", Desc: "Status of the operation"},
  		&{Name: "reason", Desc: "Reason of failure or skipped operation"},
- 		&{Name: "input.file", Desc: "Local file path"},
- 		&{Name: "input.size", Desc: "Local file size"},
+ 		&{Name: "input.path", Desc: "Path"},
  		&{Name: "result.name", Desc: "The last component of the path (including extension)."},
  		&{Name: "result.path_display", Desc: "The cased path to be used for display purposes only."},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `filerequest create`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "AllowLateUploads",
  			Desc:     "If set, allow uploads after the deadline has passed (one_day/two"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Deadline", Desc: "The deadline for this file request.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Path", Desc: "The path for the folder in the Dropbox where uploaded files will"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `group add`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ManagementType",
  			Desc:     "Group management type `company_managed` or `user_managed`",
  			Default:  "company_managed",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("company_managed"), string("user_managed")}},
  		},
  		&{Name: "Name", Desc: "Group name", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Days",
  			Desc:     "Target days old",
  			Default:  "7",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(7)},
  		},
  	},
  }
```
# Command spec changed: `job history delete`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Days",
  			Desc:     "Target days old",
  			Default:  "28",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(3650), "min": float64(1), "value": float64(28)},
  		},
  	},
  }
```
# Command spec changed: `job history list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Path",
  			Desc:     "Path to workspace",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `job log jobid`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Id", Desc: "Job ID", TypeName: "string"},
  		&{
  			Name:     "Kind",
  			Desc:     "Kind of log",
  			Default:  "toolbox",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Path to the workspace",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `job log kind`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Kind",
  			Desc:     "Log kind.",
  			Default:  "toolbox",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Path to workspace.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `job log last`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Kind",
  			Desc:     "Log kind",
  			Default:  "toolbox",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("toolbox"), string("capture"), string("summary"), string("recipe"), ...}},
  		},
  		&{
  			Name:     "Path",
  			Desc:     "Path to workspace.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  	},
  }
```
# Command spec changed: `member delete`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...},
  		&{
  			Name:     "TransferDestMember",
  			Desc:     "If provided, files from the deleted member account will be trans"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "TransferNotifyAdminEmailOnError",
  			Desc:     "If provided, errors during the transfer process will be sent via"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "WipeData", Desc: "If true, controls if the user's data will be deleted on their li"..., Default: "true", TypeName: "bool", ...},
  	},
  }
```
# Command spec changed: `member quota update`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...},
  		&{
  			Name:     "Quota",
  			Desc:     "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom q"...,
  			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
  		},
  	},
  }
```
# Command spec changed: `services github content get`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to the content", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{
  			Name:     "Ref",
  			Desc:     "Name of reference",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  }
```
# Command spec changed: `services github content put`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Branch",
  			Desc:     "Name of the branch",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Content",
  			Desc:     "Path to a content file",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Message", Desc: "Commit message", TypeName: "string"},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		... // 3 identical elements
  	},
  }
```
# Command spec changed: `services github release asset download`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			Name:     "Path",
  			Desc:     "Path to download",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  }
```
# Command spec changed: `services github release asset upload`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Asset",
  			Desc:     "Path to assets",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `services github release draft`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BodyFile",
  			Desc:     "File path to body text. THe file must encoded in UTF-8 without BOM.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_path.file_system_path_impl",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]interface{}{"shouldExist": bool(false)},
  		},
  		&{Name: "Branch", Desc: "Name of the target branch", TypeName: "string"},
  		&{Name: "Name", Desc: "Name of the release", TypeName: "string"},
  		... // 4 identical elements
  	},
  }
```
# Command spec changed: `services google mail filter add`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "AddLabels",
  			Desc:     "List of labels to add to the message, separated by ','.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaExcludeChats", Desc: "Whether the response should exclude chats.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "CriteriaFrom",
  			Desc:     "The sender's display name or email address.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaHasAttachment", Desc: "Messages that have any attachment.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "CriteriaNegatedQuery",
  			Desc:     "Only return messages not matching the specified query.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaNoAttachment", Desc: "Messages that does not have any attachment.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "CriteriaQuery",
  			Desc:     "Only return messages matching the specified query.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "CriteriaSize", Desc: "The size of the entire RFC822 message in bytes, including all he"..., Default: "0", TypeName: "int", ...},
  		&{
  			Name:     "CriteriaSizeComparison",
  			Desc:     "How the message size in bytes should be in relation to the size "...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "CriteriaTo",
  			Desc:     "The recipient's display name or email address. Includes recipien"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Forward",
  			Desc:     "Email address that the message should be forwarded to.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{
  			Name:     "RemoveLabels",
  			Desc:     "List of labels to remove from the message, separated by ','.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  }
```
# Command spec changed: `services google mail filter batch add`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...},
  		&{
- 			Name:     "ApplyToInboxMessages",
+ 			Name:     "ApplyToExistingMessages",
- 			Desc:     "Apply labels to messages satisfy query in INBOX.",
+ 			Desc:     "Apply labels to existing messages that satisfy the query.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  }
```
# Command spec changed: `services google mail label add`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "ColorBackground",
  			Desc:     "The background color.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:     "ColorText",
  			Desc:     "The text color.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}},
  		},
  		&{
  			Name:     "LabelListVisibility",
  			Desc:     "The visibility of the label in the label list in the Gmail web i"...,
  			Default:  "labelShow",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("labelHide"), string("labelShow"), string("labelShowIfUnread")}},
  		},
  		&{
  			Name:     "MessageListVisibility",
  			Desc:     "The visibility of messages with this label in the message list i"...,
  			Default:  "show",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("hide"), string("show")}},
  		},
  		&{Name: "Name", Desc: "Name of the label", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  }
```
# Command spec changed: `services google mail message list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Format",
  			Desc:     "The format to return the message in. ",
  			Default:  "metadata",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Labels",
  			Desc:     "Only return messages with labels that match all of the specified"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "MaxResults", Desc: "Maximum number of messages to return.", Default: "20", TypeName: "int", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{
  			Name:     "Query",
  			Desc:     "Only return messages matching the specified query.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  }
```
# Command spec changed: `services google mail message processed list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Format",
  			Desc:     "The format to return the message in. ",
  			Default:  "metadata",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("full"), string("metadata"), string("minimal"), string("raw")}},
  		},
  		&{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Labels",
  			Desc:     "Only return messages with labels that match all of the specified"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "MaxResults", Desc: "Maximum number of messages to return.", Default: "20", TypeName: "int", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{
  			Name:     "Query",
  			Desc:     "Only return messages matching the specified query.",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
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
  	Values: []*dc_recipe.Value{
  		&{Name: "Expires", Desc: "Expiration date/time of the shared link", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Password",
  			Desc:     "Password",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "TeamOnly", Desc: "Link is accessible only by team members", Default: "false", TypeName: "bool", ...},
  	},
  }
```
# Command spec changed: `sharedlink file list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Password",
  			Desc:     "Password for the shared link",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Url", Desc: "Shared link URL", TypeName: "domain.dropbox.model.mo_url.url_impl"},
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
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Filter the returned events to a single category. This field is o"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "File", Desc: "User email address list file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		... // 2 identical elements
  	},
  }
```
# Command spec changed: `team activity daily event`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Event category",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "EndDate",
  			Desc:     "End date",
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartDate", Desc: "Start date", TypeName: "string"},
  	},
  }
```
# Command spec changed: `team activity event`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Filter the returned events to a single category. This field is o"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  }
```
# Command spec changed: `team activity user`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Category",
  			Desc:     "Filter the returned events to a single category. This field is o"...,
  			Default:  "",
- 			TypeName: "domain.common.model.mo_string.opt_string",
+ 			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...},
  		&{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  	},
  }
```
# Command spec changed: `team content member list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEA"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  }
```
# Command spec changed: `team content policy list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEA"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  }
```
# Command spec changed: `team diag explorer`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"File": "business_file",
  		"Info": "business_info",
  		"Mgmt": "business_management",
- 		"Peer": "business_file",
+ 		"Peer": "business_management",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
## Added report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# Command spec changed: `team namespace file list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "List only for the folder matched to the name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "IncludeDeleted", Desc: "If true, deleted file or folder will be returned", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "IncludeMediaInfo",
- 			Desc:     "If true, media info is set for photo and video in json report",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "IncludeMemberFolder", Desc: "If true, include team member folders", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
- 		&{
- 			Name:     "Name",
- 			Desc:     "List only for the folder matched to the name",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  }
```
## Added report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# Command spec changed: `team namespace file size`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "Report entry for all files and directories depth directories deep",
  			Default:  "1",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{
- 				"max":   float64(2.147483647e+09),
+ 				"max":   float64(300),
  				"min":   float64(1),
  				"value": float64(1),
  			},
  		},
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "List only for the folder matched to the name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "IncludeAppFolder", Desc: "If true, include app folders", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeMemberFolder", Desc: "if true, include team member folders", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
- 		&{
- 			Name:     "Name",
- 			Desc:     "List only for the folder matched to the name",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  }
```
## Added report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# Command spec changed: `team sharedlink list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "Filter links by visibility (public/team_only/password)",
  			Default:  "public",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("public"), string("team_only"), string("password"), string("team_and_password"), ...}},
  		},
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
  	Values: []*dc_recipe.Value{
  		&{Name: "At", Desc: "New expiration date and time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}},
  		&{
  			Name:     "Days",
  			Desc:     "Days to the new expiration date",
  			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "Target link visibility",
  			Default:  "public",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("public"), string("team_only"), string("password"), string("team_and_password"), ...}},
  		},
  	},
  }
```
# Command spec changed: `teamfolder add`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "SyncSetting",
  			Desc:     "Sync setting for the team folder",
  			Default:  "default",
- 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{"options": []interface{}{string("default"), string("not_synced")}},
  		},
  	},
  }
```
# Command spec changed: `teamfolder file list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "List only for the folder matched to the name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  }
```
## Added report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# Command spec changed: `teamfolder file size`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Depth",
  			Desc:     "Depth",
  			Default:  "1",
- 			TypeName: "domain.common.model.mo_int.range_int",
+ 			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(1), "value": float64(1)},
  		},
+ 		&{
+ 			Name: "FolderName",
+ 			Desc: "List only for the folder matched to the name. Filter by exact match to the name.",
+ 		},
+ 		&{
+ 			Name: "FolderNamePrefix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the prefix.",
+ 		},
+ 		&{
+ 			Name: "FolderNameSuffix",
+ 			Desc: "List only for the folder matched to the name. Filter by name match to the suffix.",
+ 		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  }
```
## Added report(s)


| Name   | Description                               |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# Command spec changed: `teamfolder member list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 4 identical elements
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEA"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  }
```
# Command spec changed: `teamfolder policy list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
+ 		&{
+ 			Name:     "ScanTimeout",
+ 			Desc:     "Scan timeout mode. If the scan timeouts, the path of a subfolder of the team folder will be replaced with a dummy path like `TEA"...,
+ 			Default:  "short",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  }
```
