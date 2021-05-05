---
layout: release
title: リリースの変更点 88
lang: ja
---

# `リリース 88` から `リリース 89` までの変更点

# 追加されたコマンド


| コマンド                           | タイトル                                            |
|------------------------------------|-----------------------------------------------------|
| dev stage http_range               | HTTP Range request proof of concept                 |
| file export url                    | URLからドキュメントをエクスポート                   |
| file paper append                  | 既存のPaperドキュメントの最後にコンテンツを追加する |
| file paper create                  | パスに新しいPaperを作成                             |
| file paper overwrite               | 既存のPaperドキュメントを上書きする                 |
| file paper prepend                 | 既存のPaperドキュメントの先頭にコンテンツを追加する |
| services google mail sendas add    | Creates a custom "from" send-as alias               |
| services google mail sendas delete | Deletes the specified send-as alias                 |
| services google mail sendas list   | Lists the send-as aliases for the specified account |
| sharedlink info                    | 共有リンクの情報取得                                |



# コマンド仕様の変更: `config disable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `config enable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `config features`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev benchmark local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build catalogue`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build preflight`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev build readme`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci artifact connect`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci auth connect`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci auth export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev ci auth import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev diag endpoint`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev diag throughput`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev kvs dump`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay approve`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay bundle`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay recipe`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev replay remote`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev spec diff`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
- 			Name:     "FilePath",
+ 			Name:     "DocLang",
- 			Desc:     "出力先ファイルパス",
+ 			Desc:     "ドキュメント言語",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
- 			Name:     "Lang",
+ 			Name:     "FilePath",
- 			Desc:     "言語",
+ 			Desc:     "出力先ファイルパス",
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Release1", Desc: "リリース名1", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Release2", Desc: "リリース名2", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev spec doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage dbxfs`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage gmail`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage griddata`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "In", Desc: "テスト入力グリッドデータ"}},
  	GridDataOutput: {&{Name: "Out", Desc: "出力グリッドデータ"}},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage gui`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage scoped`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage teamfolder`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev stage upload_append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test echo`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test kvsfootprint`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test recipe`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test resources`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev test setup teamsharedlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  		&{Name: "Query", Desc: "クエリ", TypeName: "string"},
+ 		&{
+ 			Name:     "Seed",
+ 			Desc:     "シェアードリンクのシード値",
+ 			Default:  "0",
+ 			TypeName: "int",
+ 		},
  		&{Name: "Visibility", Desc: "ビジビリティ", Default: "random", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util anonymise`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util curl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util image jpeg`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `dev util wait`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file archive local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file dispatch local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock batch acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock batch release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file sync down`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file sync online`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `file watch`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `image info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job log jobid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job log kind`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `job log last`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member file permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 5 identical fields
  	CliArgs:         "-path /DROPBOX/PATH/TO/PERM_DELETE",
  	CliNote:         "",
- 	ConnUsePersonal: true,
+ 	ConnUsePersonal: false,
- 	ConnUseBusiness: false,
+ 	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_business"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "MemberEmail", Desc: "チームメンバーのメールアドレス.", TypeName: "string"},
  		&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []interface{}{string("files.permanent_delete"), string("team_data.member"), string("members.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member folder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update invisible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `member update visible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana team list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana team project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana team task list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana workspace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services asana workspace project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github content get`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github content put`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release asset download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release asset list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release asset upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github release list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services github tag create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail filter list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail label rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message label delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail message processed list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google mail thread list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet append`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "Data", Desc: "入力データファイル"}},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "書き出したシートデータ"}},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "Data", Desc: "入力データファイル"}},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets sheet list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services google sheets spreadsheet create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `services slack conversation list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team content mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team content policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team diag explorer`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     {"File": "business_file", "Info": "business_info", "Mgmt": "business_management", "Peer": "business_file"},
  	Services:       {"dropbox_business"},
- 	IsSecret:       false,
+ 	IsSecret:       true,
  	IsConsole:      false,
  	IsExperimental: false,
  	... // 5 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team filerequest clone`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink delete links`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink delete member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink update password`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `team sharedlink update visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:  "add",
  	Title: "チームフォルダへのユーザー/グループの一括追加",
  	Desc: strings.Join({
  		... // 605 identical bytes
  		"\x81\xaf、「メンバーはすでにそのフォルダにアクセ\xe3\x82",
  		"\xb9している」というようなエラーを報告しません",
+ 		".\n\n例:\n\n* Sales（チームフォルダ、グループ「Sales\xe3",
+ 		"\x80\x8dの編集者アクセス)\n\t* Sydney (個人アカウント syd",
+ 		"ney@example.com の閲覧アクセス)\n\t* Tokyo (グループ \"T",
+ 		"okyo Deal Desk\" の編集者アクセス)\n\t\t* Monthly (個人ア",
+ 		"カウント success@example.com の閲覧アクセス )\n* Marke",
+ 		`ting (チームフォルダ、グループ "Marketing "の編集`,
+ 		"者アクセス)\n\t* Sydney (グループ \"Sydney Sales\" の編\xe9\x9b",
+ 		"\x86者アクセス)\n\t* Tokyo (グループ \"Tokyo Sales \"のビュ",
+ 		"ーアアクセス)\n\n1. 次のようなCSVファイルを準備\xe3",
+ 		"\x81\x97ます\n\n```\nSales,,editor,Sales\nSales,Sydney,editor,sydney@exa",
+ 		"mple.com\nSales,Tokyo,editor,Tokyo Deal Desk\nSales,Tokyo/Monthly,",
+ 		"viewer,success@example.com\nMarketing,,editor,Marketing\nMarketing",
+ 		",Sydney,editor,Sydney Sales\nMarketing,Tokyo,viewer,Tokyo Sales\n`",
+ 		"``\n\n2. その後、以下のようにコマンドを実行しま",
+ 		"す.\n\n```\ntbx teamfolder member add -file /PATH/TO/DATA.csv\n```\n",
+ 		"\n注: このコマンドは、チームフォルダが存在し\xe3",
+ 		"\x81\xaaい場合には、チームフォルダを作成します. し",
+ 		"かし、このコマンドは、グループが見つからな\xe3",
+ 		"\x81\x84場合には、グループを作成しません. このコマ",
+ 		"ンドを実行する前に、グループが存在している\xe5",
+ 		"\xbf\x85要があります",
  		... // 1 identical byte
  	}, ""),
  	Remarks: "(非可逆な操作です)",
  	Path:    "teamfolder member add",
  	... // 14 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder partial replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util date today`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util datetime now`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util decode base_32`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util decode base_64`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util encode base_32`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util encode base_64`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util qrcode create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util qrcode wifi`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util time now`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util unixtime format`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Format", Desc: "時間フォーマット", Default: "iso8601", TypeName: "essentials.model.mo_string.select_string", ...},
  		&{Name: "Precision", Desc: "時間精度 (second/ms/ns)", Default: "second", TypeName: "essentials.model.mo_string.select_string", ...},
+ 		&{Name: "Time", Desc: "Unix 時間", Default: "0", TypeName: "int"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util unixtime now`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx sheet export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {&{Name: "Data", Desc: "データの書き出し"}},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx sheet import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {&{Name: "Data", Desc: "入力データファイル"}},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `util xlsx sheet list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
# コマンド仕様の変更: `version`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 19 identical fields
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      nil,
+ 	TextInput:      []*dc_recipe.DocTextInput{},
- 	JsonInput:      nil,
+ 	JsonInput:      []*dc_recipe.DocJsonInput{},
  }
```
