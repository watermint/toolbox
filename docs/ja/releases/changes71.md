---
layout: release
title: リリースの変更点: 70
lang: ja
---

# `リリース 70` から `リリース 71` までの変更点

# 追加されたコマンド


| コマンド            | タイトル                             |
|---------------------|--------------------------------------|
| dev diag endpoint   | エンドポイントを一覧                 |
| dev diag throughput | キャプチャログからスループットを評価 |



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
  	... // 11 identical fields
  }
```
