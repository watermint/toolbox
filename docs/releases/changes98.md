---
layout: release
title: Changes of Release 97
lang: en
---

# Changes between `Release 97` to `Release 98`

# Commands added


| Command                      | Title                                                                     |
|------------------------------|---------------------------------------------------------------------------|
| group batch add              | Bulk adding groups                                                        |
| team admin group role add    | Add the role to members of the group                                      |
| team admin group role delete | Delete the role from all members except of members of the exception group |
| team admin list              | List admin roles of members                                               |
| team admin role add          | Add a new role to the member                                              |
| team admin role clear        | Remove all admin roles from the member                                    |
| team admin role delete       | Remove a role from the member                                             |
| team admin role list         | List admin roles of the team                                              |
| util image placeholder       | Create placeholder image                                                  |



# Command spec changed: `dev stage scoped`



## Changed report: member_list

```
  &dc_recipe.Report{
  	Name: "member_list",
  	Desc: "This report shows a list of members.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 12 identical elements
  		&{Name: "persistent_id", Desc: "Persistent ID that a team can attach to the user. The persistent"...},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{Name: "invited_on", Desc: "The date and time the user was invited to the team"},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member clear externalid`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 15 identical elements
  		&{Name: "result.persistent_id", Desc: "Persistent ID that a team can attach to the user. The persistent"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member folder list`



## Changed report: member_with_no_folder

```
  &dc_recipe.Report{
  	Name: "member_with_no_folder",
  	Desc: "This report shows a list of members.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "surname", Desc: "Also known as a last name or family name."},
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
+ 		&{Name: "invited_on", Desc: "The date and time the user was invited to the team"},
  	},
  }
```
# Command spec changed: `member invite`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 10 identical elements
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member list`



## Changed report: member

```
  &dc_recipe.Report{
  	Name: "member",
  	Desc: "This report shows a list of members.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 5 identical elements
  		&{Name: "display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{Name: "invited_on", Desc: "The date and time the user was invited to the team"},
  		&{Name: "role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  	},
  }
```
# Command spec changed: `member reinvite`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 7 identical elements
  		&{Name: "input.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "input.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "input.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "input.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "input.tag", Desc: "Operation tag"},
  		... // 5 identical elements
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update email`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 9 identical elements
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update externalid`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 9 identical elements
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update invisible`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 15 identical elements
  		&{Name: "result.persistent_id", Desc: "Persistent ID that a team can attach to the user. The persistent"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update profile`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 10 identical elements
  		&{Name: "result.display_name", Desc: "A name that can be used directly to represent the name of a user"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `member update visible`



## Changed report: operation_log

```
  &dc_recipe.Report{
  	Name: "operation_log",
  	Desc: "This report shows the transaction result.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 15 identical elements
  		&{Name: "result.persistent_id", Desc: "Persistent ID that a team can attach to the user. The persistent"...},
  		&{Name: "result.joined_on", Desc: "The date and time the user joined as a member of a specific team."},
+ 		&{
+ 			Name: "result.invited_on",
+ 			Desc: "The date and time the user was invited to the team",
+ 		},
  		&{Name: "result.role", Desc: "The user's role in the team (team_admin, user_management_admin, "...},
  		&{Name: "result.tag", Desc: "Operation tag"},
  	},
  }
```
# Command spec changed: `teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...},
  		&{Name: "MemberTypeInternal", Desc: "Filter folder members. Keep only members are internal (in the sa"...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
  				string("files.metadata.read"),
+ 				string("groups.read"),
+ 				string("members.read"),
  				string("sharing.read"),
  				string("team_data.member"),
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
