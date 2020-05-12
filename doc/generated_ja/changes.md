# `リリース 66` から `リリース 67` までの変更点

# 追加されたコマンド

| コマンド      | タイトル                                         |
|---------------|--------------------------------------------------|
| job log jobid | Retrieve logs of specified Job ID                |
| job log kind  | Concatenate and print logs of specified log kind |
| job log last  | Print the last job log files                     |


# コマンド仕様の変更: `job history list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Path",
+ 			Desc:     "Path to workspace",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 	},
  }
```
