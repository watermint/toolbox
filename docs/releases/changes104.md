---
layout: release
title: Changes of Release 103
lang: en
---

# Changes between `Release 103` to `Release 104`

# Commands added


| Command                      | Title                                               |
|------------------------------|-----------------------------------------------------|
| sharedfolder leave           | Leave from the shared folder                        |
| sharedfolder mount add       | Add the shared folder to the current user's Dropbox |
| sharedfolder mount delete    | The current user unmounts the designated folder.    |
| sharedfolder mount list      | List all shared folders the current user mounted    |
| sharedfolder mount mountable | List all shared folders the current user can mount  |
| team namespace summary       | Report team namespace status summary.               |
| team runas file list         | List files and folders run as a member              |



# Commands deleted


| Command         | Title                                 |
|-----------------|---------------------------------------|
| file mount list | List mounted/unmounted shared folders |



# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "BlockBlockSize",
  			Desc:     "Block size for batch upload",
- 			Default:  "16",
+ 			Default:  "40",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
  				"max":   float64(1000),
  				"min":   float64(1),
- 				"value": float64(16),
+ 				"value": float64(40),
  			},
  		},
  		&{Name: "Method", Desc: "Upload method", Default: "block", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...},
  		... // 7 identical elements
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
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."},
  		&{Name: "MemberEmail", Desc: "Filter by member email address. Filter by email address."},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
+ 				string("groups.read"),
  				string("members.read"),
  				string("sharing.read"),
  				... // 2 identical elements
  			},
  		},
  		&{Name: "ScanTimeout", Desc: "Scan timeout mode. If the scan timeouts, the path of a subfolder"..., Default: "short", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `sharedfolder list`



## Changed report: shared_folder

```
  &dc_recipe.Report{
  	Name: "shared_folder",
  	Desc: "This report shows a list of shared folders.",
  	Columns: []*dc_recipe.ReportColumn{
+ 		&{Name: "shared_folder_id", Desc: "The ID of the shared folder."},
  		&{Name: "name", Desc: "The name of the this shared folder."},
  		&{Name: "access_type", Desc: "The current user's access level for this shared file/folder (own"...},
  		... // 9 identical elements
  	},
  }
```
