# `リリース 74` から `リリース 75` までの変更点

# 追加されたコマンド


| コマンド                         | タイトル                     |
|----------------------------------|------------------------------|
| dev stage teamfolder             | Team folder operation sample |
| services slack conversation list | List channels                |



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
- 		"Peer": "business_file",
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
