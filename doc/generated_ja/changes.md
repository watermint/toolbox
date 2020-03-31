# `リリース 62` から `リリース 63` までの変更点

# 追加されたコマンド

| コマンド            | タイトル                                                     |
|---------------------|--------------------------------------------------------------|
| team content member | List team folder & shared folder members                     |
| team content policy | List policies of team folders and shared folders in the team |



# コマンド仕様の変更: `sharedfolder list`



## 変更されたレポート: shared_folder

```
  &rc_doc.Report{
  	Name: "shared_folder",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 5 identical elements
  		&{Name: "is_inside_team_folder", Desc: "Whether this folder is inside of a team folder."},
  		&{Name: "is_team_folder", Desc: "Whether this folder is a team folder."},
+ 		&{
+ 			Name: "policy_manage_access",
+ 			Desc: "Who can add and remove members from this shared folder.",
+ 		},
+ 		&{Name: "policy_shared_link", Desc: "Who links can be shared with."},
  		&{Name: "policy_member", Desc: "Who can be a member of this shared folder, as set on the folder itself (team, or anyone)"},
+ 		&{
+ 			Name: "policy_viewer_info",
+ 			Desc: "Who can enable/disable viewer info for this shared folder.",
+ 		},
+ 		&{Name: "owner_team_id", Desc: "Team ID of the team that owns the folder"},
+ 		&{Name: "owner_team_name", Desc: "Team name of the team that owns the folder"},
  	},
  }

```

