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
| services asana team list              | List team                             |
| services asana team project list      | List projects of the team             |
| services asana team task list         | List task of the team                 |
| services asana workspace list         | List workspaces                       |
| services asana workspace project list | List projects of the workspace        |



# Commands deleted


| Command       | Title                                 |
|---------------|---------------------------------------|
| dev async     | Async framework test                  |
| dev catalogue | Generate catalogue                    |
| dev doc       | Document generator                    |
| dev echo      | Echo text                             |
| dev preflight | Process prerequisites for the release |



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
+ 			TypeName: "domain.common.model.mo_string.select_string",
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
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  }
```
# Command spec changed: `team diag explorer`


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
  		&{Name: "Depth", Desc: "Report entry for all files and directories depth directories deep", Default: "1", TypeName: "domain.common.model.mo_int.range_int", ...},
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
  		&{Name: "Depth", Desc: "Depth", Default: "1", TypeName: "domain.common.model.mo_int.range_int", ...},
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
+ 			TypeName: "domain.common.model.mo_string.select_string",
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
+ 			TypeName: "domain.common.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("short"), string("long")}},
+ 		},
  	},
  }
```
