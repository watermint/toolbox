# `リリース 82` から `リリース 83` までの変更点

# 追加されたコマンド

| コマンド                       | タイトル                                                     |
|--------------------------------|--------------------------------------------------------------|
| dev benchmark uploadlink       | Benchmark single file upload with upload temporary link API. |
| file info                      | Resolve metadata of the path                                 |
| teamfolder partial replication | Partial team folder replication to the other team            |

# コマンド仕様の変更: `team diag explorer`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"File": "business_file",
  		"Info": "business_info",
  		"Mgmt": "business_management",
- 		"Peer": "business_file",
+ 		"Peer": "business_management",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
