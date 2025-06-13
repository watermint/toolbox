---
layout: release
title: リリースの変更点 131
lang: ja
---

# `リリース 131` から `リリース 132` までの変更点

# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
