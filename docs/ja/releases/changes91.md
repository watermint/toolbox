---
layout: release
title: Changes of Release 90
lang: ja
---

# `リリース 90` から `リリース 91` までの変更点

# 追加されたコマンド


| コマンド        | タイトル                   |
|-----------------|----------------------------|
| dev release doc | Generate release documents |



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
+ 		"Peer": "business_info",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: true,
  	... // 11 identical fields
  }
```
