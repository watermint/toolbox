# `リリース 84` から `リリース 85` までの変更点

# 追加されたコマンド

| コマンド                                  | タイトル                       |
|-------------------------------------------|--------------------------------|
| services google sheets sheet list         | List sheets of the spreadsheet |
| services google sheets spreadsheet create | Create a new spreadsheet       |

# コマンド仕様の変更: `dev release candidate`

## 追加されたレポート

| 名称   | 説明               |
|--------|--------------------|
| result | Recipe test result |

# コマンド仕様の変更: `dev release publish`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
  		"ConnGithub": "github_repo",
+ 		"Peer":       "github_repo",
  	},
  	Services: {"github"},
  	IsSecret: true,
  	... // 7 identical fields
  }
```

## 追加されたレポート

| 名称   | 説明               |
|--------|--------------------|
| commit | Commit information |
| result | Recipe test result |

# コマンド仕様の変更: `dev test recipe`

## 追加されたレポート

| 名称   | 説明               |
|--------|--------------------|
| result | Recipe test result |

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
