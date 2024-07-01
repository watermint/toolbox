---
layout: release
title: リリースの変更点 134
lang: ja
---

# `リリース 134` から `リリース 135` までの変更点

# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github_public"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
