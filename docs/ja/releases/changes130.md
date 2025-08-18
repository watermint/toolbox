---
layout: release
title: リリースの変更点 129
lang: ja
---

# `リリース 129` から `リリース 130` までの変更点

# 追加されたコマンド


| コマンド                 | タイトル                           |
|--------------------------|------------------------------------|
| config license list      | 利用可能なライセンスキーのリスト   |
| dev license issue        | ライセンスの発行                   |
| dev release announcement | お知らせの更新                     |
| dev release checkin      | 新作りリースをチェック             |
| dev test license         | ライセンスが必要なロジックのテスト |



# 削除されたコマンド


| コマンド                      | タイトル                                                           |
|-------------------------------|--------------------------------------------------------------------|
| dev test auth all             | すべてのスコープでのDropboxへの接続テスト                          |
| dev test setup massfiles      | テストファイルとしてウィキメディアダンプファイルをアップロードする |
| dev test setup teamsharedlink | デモ用共有リンクの作成                                             |



# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```

## 追加されたレポート


| 名称          | 説明     |
|---------------|----------|
| announcements | お知らせ |


# コマンド仕様の変更: `google calendar event list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_calendar"},
+ 	ConnScopes:      map[string]string{"Peer": "google_calendar2024"},
  	Services:        {"google_calendar"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail filter add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail filter batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail filter delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail filter list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail label delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail label list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail label rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail message label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail message label delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail message list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail message processed list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail message send`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail sendas add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail sendas delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail sendas list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google mail thread list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets sheet list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `google sheets spreadsheet create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# コマンド仕様の変更: `util desktop screenshot interval`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Count", Desc: "スクリーンショットの枚数。値が1未満の場合、"..., Default: "-1", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(2),
+ 				"max":   float64(1),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Interval", Desc: "スクリーンショットの間隔秒数。", Default: "10", TypeName: "int", ...},
  		&{Name: "NamePattern", Desc: "スクリーンショットファイルの名前パターン。\xe4"..., Default: "{% raw %}{{.{% endraw %}Sequence}}_{% raw %}{{.{% endraw %}Timestamp}}.png", TypeName: "string", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `util desktop screenshot snap`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(2),
+ 				"max":   float64(1),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Path", Desc: "スクリーンショットを保存するパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
