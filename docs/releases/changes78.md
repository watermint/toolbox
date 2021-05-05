---
layout: release
title: Changes of Release 77
lang: en
---

# Changes between `Release 77` to `Release 78`

# Commands added


| Command            | Title                        |
|--------------------|------------------------------|
| group folder list  | Find folders of each group   |
| member folder list | Find folders for each member |



# Command spec changed: `team diag explorer`



## Changed report: namespace_member

```
  &dc_recipe.Report{
  	Name: "namespace_member",
  	Desc: "This report shows a list of members of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "entry_is_inherited", Desc: "True if the member has access from a parent folder"},
  		&{Name: "email", Desc: "Email address of user."},
  		&{
  			Name: "display_name",
- 			Desc: "Type of the session (web_session, desktop_client, or mobile_client)",
+ 			Desc: "Team member display name.",
  		},
  		&{Name: "group_name", Desc: "Name of the group"},
  		&{Name: "invitee_email", Desc: "Email address of invitee for this folder"},
  	},
  }
```
# Command spec changed: `team namespace member list`



## Changed report: namespace_member

```
  &dc_recipe.Report{
  	Name: "namespace_member",
  	Desc: "This report shows a list of members of namespaces in the team.",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 3 identical elements
  		&{Name: "entry_is_inherited", Desc: "True if the member has access from a parent folder"},
  		&{Name: "email", Desc: "Email address of user."},
  		&{
  			Name: "display_name",
- 			Desc: "Type of the session (web_session, desktop_client, or mobile_client)",
+ 			Desc: "Team member display name.",
  		},
  		&{Name: "group_name", Desc: "Name of the group"},
  		&{Name: "invitee_email", Desc: "Email address of invitee for this folder"},
  	},
  }
```
