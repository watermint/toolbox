---
layout: release
title: リリースの変更点 125
lang: ja
---

# `リリース 125` から `リリース 126` までの変更点

# 追加されたコマンド


| コマンド                   | タイトル                                           |
|----------------------------|----------------------------------------------------|
| config feature disable     | 機能を無効化します.                                |
| config feature enable      | 機能を有効化します.                                |
| config feature list        | 利用可能なオプション機能一覧.                      |
| dev info                   | 開発情報                                           |
| dev placeholder pathchange | パス変更文書生成のためのプレースホルダー・コマンド |
| dev placeholder prune      | 剪定ワークフローメッセージのプレースホルダ         |
| log cat job                | 指定したジョブIDのログを取得する                   |
| log cat kind               | 指定種別のログを結合して出力します                 |
| log cat last               | 最後のジョブのログファイルを出力.                  |
| log job archive            | ジョブのアーカイブ                                 |
| log job delete             | 古いジョブ履歴の削除                               |
| log job list               | ジョブ履歴の表示                                   |
| log job ship               | ログの転送先Dropboxパス                            |



# 削除されたコマンド


| コマンド            | タイトル                           |
|---------------------|------------------------------------|
| config disable      | 機能を無効化します.                |
| config enable       | 機能を有効化します.                |
| config features     | 利用可能なオプション機能一覧.      |
| job history archive | ジョブのアーカイブ                 |
| job history delete  | 古いジョブ履歴の削除               |
| job history list    | ジョブ履歴の表示                   |
| job history ship    | ログの転送先Dropboxパス            |
| job log jobid       | 指定したジョブIDのログを取得する   |
| job log kind        | 指定種別のログを結合して出力します |
| job log last        | 最後のジョブのログファイルを出力.  |



# コマンド仕様の変更: `dev benchmark local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev benchmark local",
  	CliArgs: strings.Join({
  		"-num-files NUM -path /LOCAL/PATH/TO/PROCESS -size-max-kb NUM -si",
  		"ze-min-kb NUM",
- 		`"`,
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 16 identical fields
  }
```
# コマンド仕様の変更: `dev build catalogue`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
- 	Values:  []*dc_recipe.Value{},
+ 	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "Importer",
+ 			Desc:     "インポーター・タイプ",
+ 			Default:  "default",
+ 			TypeName: "essentials.model.mo_string.select_string_internal",
+ 			TypeAttr: map[string]any{"options": []any{string("default"), string("enhanced")}},
+ 		},
+ 	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
