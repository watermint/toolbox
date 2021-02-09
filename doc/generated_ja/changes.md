# `リリース 83` から `リリース 84` までの変更点

# 追加されたコマンド

| コマンド  | タイトル    |
|-----------|-------------|
| file lock | Lock a file |

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
- 		"Peer": "business_management",
+ 		"Peer": "business_file",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
