# `リリース 76` から `リリース 77` までの変更点

# 追加されたコマンド


| コマンド               | タイトル                                                                                                                                                       |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| image info             | Show EXIF information of the image file                                                                                                                        |
| member file permdelete | Permanently delete the file or folder at a given path of the team member. Please see https://www.dropbox.com/help/40 for more detail about permanent deletion. |



# コマンド仕様の変更: `dev benchmark upload`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
+ 		&{
+ 			Name:     "Shard",
+ 			Desc:     "Number of shared folders to distribute namespace",
+ 			Default:  "1",
+ 			TypeName: "int",
+ 		},
  		&{Name: "SizeMaxKb", Desc: "Maximum file size (KiB).", Default: "2048", TypeName: "int", ...},
  		&{Name: "SizeMinKb", Desc: "Minimum file size (KiB).", Default: "0", TypeName: "int", ...},
  	},
  }
```
# コマンド仕様の変更: `dev stage teamfolder`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Peer",
- 			Desc:     "Account alias",
- 			Default:  "&{Peer [groups.write files.content.write] <nil>}",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: []interface{}{string("groups.write"), string("files.content.write")},
- 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space tea"...,
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
+ 			TypeAttr: []interface{}{
+ 				string("files.content.read"), string("files.content.write"),
+ 				string("groups.write"), string("sharing.read"), string("sharing.write"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
+ 		},
  	},
  }
```
# コマンド仕様の変更: `services asana workspace list`


## 変更されたレポート: workspaces

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
- 			Desc: `	Whether the workspace is an organization.`,
+ 			Desc: "Whether the workspace is an organization.",
  		},
  	},
  }
```
