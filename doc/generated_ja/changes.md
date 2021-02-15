# `リリース 84` から `リリース 85` までの変更点

# 追加されたコマンド

| コマンド                                  | タイトル                         |
|-------------------------------------------|----------------------------------|
| services google sheets sheet clear        | Clears values from a spreadsheet |
| services google sheets sheet export       | Export sheet data                |
| services google sheets sheet list         | List sheets of the spreadsheet   |
| services google sheets spreadsheet create | Create a new spreadsheet         |

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
+ 		"Peer":       "github_repo",
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


