# `リリース 62` から `リリース 63` までの変更点

# 追加されたコマンド

| コマンド            | タイトル                                 |
|---------------------|------------------------------------------|
| team content member | List team folder & shared folder members |



# コマンド仕様の変更: `sharedfolder list`



## 変更されたレポート: shared_folder

```
  &rc_doc.Report{
  	Name: "shared_folder",
  	Desc: "",
  	Columns: []*rc_doc.ReportColumn{
  		... // 6 identical elements
  		&{Name: "is_team_folder", Desc: "Whether this folder is a team folder."},
  		&{Name: "policy_member", Desc: "Who can be a member of this shared folder, as set on the folder itself (team, or anyone)"},
+ 		&{Name: "owner_team_id"},
+ 		&{Name: "owner_team_name"},
  	},
  }

```

