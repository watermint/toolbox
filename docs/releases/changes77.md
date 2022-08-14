---
layout: release
title: Changes of Release 76
lang: en
---

# Changes between `Release 76` to `Release 77`

# Commands added


| Command                | Title                                                                                                                                                          |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| image info             | Show EXIF information of the image file                                                                                                                        |
| member file permdelete | Permanently delete the file or folder at a given path of the team member. Please see https://www.dropbox.com/help/40 for more detail about permanent deletion. |



# Commands deleted


| Command        | Title                |
|----------------|----------------------|
| dev test async | Async framework test |



# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "Shard",
+ 			Desc:     "Number of shared folders to distribute namespace",
+ 			Default:  "1",
+ 			TypeName: "int",
+ 		},
  		&{Name: "SizeMaxKb", Desc: "Maximum file size (KiB).", Default: "2048", TypeName: "int", ...},
  		&{Name: "SizeMinKb", Desc: "Minimum file size (KiB).", Default: "0", TypeName: "int", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev stage teamfolder`



## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Peer",
- 			Desc:     "Account alias",
- 			Default:  "&{Peer [groups.write files.content.write] <nil>}",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: []any{string("groups.write"), string("files.content.write")},
- 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space tea"...,
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
+ 			TypeAttr: []any{
+ 				string("files.content.read"), string("files.content.write"),
+ 				string("groups.write"), string("sharing.read"), string("sharing.write"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
+ 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `services asana workspace list`



## Changed report: workspaces

```
  &dc_recipe.Report{
  	Name: "workspaces",
  	Desc: "Workspace",
  	Columns: []*dc_recipe.ReportColumn{
  		&{Name: "gid", Desc: "Globally unique identifier of the resource, as a string."},
  		&{Name: "resource_type", Desc: "The base type of this resource."},
  		&{Name: "name", Desc: "The name of the workspace."},
  		&{
  			Name: "is_organization",
  			Desc: strings.Join({
- 				`	`,
  				"Whether the workspace is an organization.",
  			}, ""),
  		},
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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 11 identical fields
  }
```
