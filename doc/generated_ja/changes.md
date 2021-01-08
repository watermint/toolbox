# `リリース 79` から `リリース 80` までの変更点

# 追加されたコマンド

| コマンド                | タイトル                                  |
|-------------------------|-------------------------------------------|
| member update invisible | Enable directory restriction to members   |
| member update visible   | Disable directory restriction to members  |
| teamfolder member add   | Batch adding users/groups to team folders |

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
- 		"Peer": "business_info",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
