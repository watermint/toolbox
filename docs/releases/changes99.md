---
layout: release
title: Changes of Release 98
lang: en
---

# Changes between `Release 98` to `Release 99`

# Commands added


| Command                                     | Title                                             |
|---------------------------------------------|---------------------------------------------------|
| file share info                             | Retrieve sharing information of the file          |
| sharedfolder member add                     | Add a member to the shared folder                 |
| sharedfolder member delete                  | Delete a member from the shared folder            |
| sharedfolder share                          | Share a folder                                    |
| sharedfolder unshare                        | Unshare a folder                                  |
| team runas file batch copy                  | Batch copy files/folders as a member              |
| team runas file sync batch up               | Batch sync up that run as members                 |
| team runas sharedfolder batch share         | Batch share folders for members                   |
| team runas sharedfolder batch unshare       | Batch unshare folders for members                 |
| team runas sharedfolder member batch add    | Batch add members to member's shared folders      |
| team runas sharedfolder member batch delete | Batch delete members from member's shared folders |



# Command spec changed: `file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...},
+ 		&{
+ 			Name:     "IncludeExplicitSharedMembers",
+ 			Desc:     " If true, the results will include a flag for each file indicating whether or not that file has any explicit members.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
+ 		&{
+ 			Name:     "IncludeMountedFolders",
+ 			Desc:     " If true, the results will include entries under mounted folders which includes app folder, shared folder and team folder.",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
