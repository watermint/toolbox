---
layout: release
title: Changes of Release 110
lang: en
---

# Changes between `Release 110` to `Release 111`

# Commands added


| Command                             | Title                                                       |
|-------------------------------------|-------------------------------------------------------------|
| file tag add                        | Add a tag to the file/folder                                |
| file tag delete                     | Delete a tag from the file/folder                           |
| file tag list                       | List tags of the path                                       |
| file template apply local           | Apply file/folder structure template to the local path      |
| file template apply remote          | Apply file/folder structure template to the Dropbox path    |
| file template capture local         | Capture file/folder structure as template from local path   |
| file template capture remote        | Capture file/folder structure as template from Dropbox path |
| services dropbox user info          | Retrieve current account info                               |
| teamspace asadmin file list         | List files and folders in team space run as admin           |
| teamspace asadmin folder add        | Create top level folder in the team space                   |
| teamspace asadmin folder delete     | Delete top level folder of the team space                   |
| teamspace asadmin folder permdelete | Permanently delete top level folder of the team space       |
| teamspace asadmin member list       | List top level folder members                               |
| teamspace file list                 | List files and folders in team space                        |
| util tidy move dispatch             | Dispatch files                                              |
| util tidy move simple               | Archive local files                                         |
| util tidy pack remote               | Package remote folder into the zip file                     |



# Commands deleted


| Command             | Title                |
|---------------------|----------------------|
| file archive local  | Archive local files  |
| file dispatch local | Dispatch local files |



# Command spec changed: `config auth delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "KeyName", Desc: "Application key name", TypeName: "string"},
  		&{Name: "PeerName", Desc: "Peer name", TypeName: "string"},
- 		&{
- 			Name:     "Scope",
- 			Desc:     "Auth scope",
- 			TypeName: "essentials.model.mo_string.opt_string",
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## Added report(s)


| Name    | Description          |
|---------|----------------------|
| deleted | Auth credential data |


