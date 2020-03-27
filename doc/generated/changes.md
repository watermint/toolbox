# Changes between `Release 62` to `Release 63`

# Commands added

| Command             | Title                                    |
|---------------------|------------------------------------------|
| team content member | List team folder & shared folder members |



# Command spec changed: `sharedfolder list`



## Changed report: shared_folder

```
  &rc_doc.Report{
  	Name: "shared_folder",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 6 identical elements
  		&{Name: "is_team_folder", Desc: "Whether this folder is a team folder."},
  		&{Name: "policy_member", Desc: "Who can be a member of this shared folder, as set on the folder itself (team, or anyone)"},
+ 		&{Name: "owner_team_id", Desc: "Team ID of the team that owns the folder"},
+ 		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
  	},
  }

```

