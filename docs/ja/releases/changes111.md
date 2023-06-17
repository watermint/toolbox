---
layout: release
title: リリースの変更点 110
lang: ja
---

# `リリース 110` から `リリース 111` までの変更点

# 追加されたコマンド


| コマンド                            | タイトル                                                             |
|-------------------------------------|----------------------------------------------------------------------|
| file tag add                        | ファイル/フォルダーにタグを追加する                                  |
| file tag delete                     | ファイル/フォルダーからタグを削除する                                |
| file tag list                       | パスのタグを一覧                                                     |
| file template apply local           | ファイル/フォルダー構造のテンプレートをローカルパスに適用する        |
| file template apply remote          | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する       |
| file template capture local         | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む    |
| file template capture remote        | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。 |
| services dropbox user info          | 現在のアカウント情報を取得する                                       |
| teamspace asadmin file list         | チームスペース内のファイルやフォルダーを一覧表示することができます。 |
| teamspace asadmin folder add        | チームスペースにトップレベルのフォルダーを作成                       |
| teamspace asadmin folder delete     | チームスペースのトップレベルフォルダーを削除する                     |
| teamspace asadmin folder permdelete | チームスペースのトップレベルフォルダを完全に削除します。             |
| teamspace asadmin member list       | トップレベルのフォルダーメンバーをリストアップ                       |
| teamspace file list                 | チームスペースにあるファイルやフォルダーを一覧表示                   |
| util tidy move dispatch             | ファイルを整理                                                       |
| util tidy move simple               | ローカルファイルをアーカイブします                                   |
| util tidy pack remote               | リモートフォルダをZIPファイルにパッケージする                        |



# 削除されたコマンド


| コマンド            | タイトル                           |
|---------------------|------------------------------------|
| file archive local  | ローカルファイルをアーカイブします |
| file dispatch local | ローカルファイルを整理します       |



# コマンド仕様の変更: `config auth delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "KeyName", Desc: "アプリケーションキー名", TypeName: "string"},
  		&{Name: "PeerName", Desc: "ピア名", TypeName: "string"},
- 		&{
- 			Name:     "Scope",
- 			Desc:     "認証スコープ",
- 			TypeName: "essentials.model.mo_string.opt_string",
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## 追加されたレポート


| 名称    | 説明                     |
|---------|--------------------------|
| deleted | 認証クレデンシャルデータ |


