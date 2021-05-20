---
layout: release
title: リリースの変更点 85
lang: ja
---

# `リリース 85` から `リリース 86` までの変更点

# 追加されたコマンド


| コマンド                | タイトル                                                                                  |
|-------------------------|-------------------------------------------------------------------------------------------|
| dev stage dbxfs         | Dropboxのファイルシステムのインプリケーションを確認しますキャッシュされたシステムに対して |
| dev stage upload_append | 新しいアップロードAPIテスト                                                               |
| util qrcode create      | QRコード画像ファイルの作成                                                                |
| util qrcode wifi        | WIFI設定用のQRコードを生成                                                                |



# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: true,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "user_full"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_individual"},
  	Services:        {"dropbox"},
  	IsSecret:        true,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
+ 		&{
+ 			Name:     "BlockBlockSize",
+ 			Desc:     "一括アップロード時のブロックサイズ",
+ 			Default:  "40",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(1000), "min": float64(1), "value": float64(40)},
+ 		},
- 		&{
- 			Name:     "ChunkSizeKb",
- 			Desc:     "チャンクサイズをKiB単位でアップロード",
- 			Default:  "65536",
- 			TypeName: "essentials.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
- 		},
+ 		&{
+ 			Name:     "Method",
+ 			Desc:     "アップロード方法",
+ 			Default:  "block",
+ 			TypeName: "essentials.model.mo_string.select_string",
+ 			TypeAttr: map[string]interface{}{"options": []interface{}{string("block"), string("sequential")}},
+ 		},
  		&{Name: "NumFiles", Desc: "ファイル数.", Default: "1000", TypeName: "int", ...},
  		&{Name: "Path", Desc: "Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{
  			Name:     "Peer",
  			Desc:     "アカウントの別名",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("files.content.write")},
  		},
+ 		&{
+ 			Name:     "PreScan",
+ 			Desc:     "プリスキャンのデスティネーションパス",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
- 		&{
- 			Name:     "Shard",
- 			Desc:     "名前空間を分散させる共有フォルダの数.",
- 			Default:  "1",
- 			TypeName: "int",
- 		},
+ 		&{
+ 			Name:     "SeqChunkSizeKb",
+ 			Desc:     "チャンクサイズをKiB単位でアップロード",
+ 			Default:  "65536",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
+ 		},
  		&{Name: "SizeMaxKb", Desc: "最大ファイルサイズ (KiB).", Default: "2048", TypeName: "int", ...},
  		&{Name: "SizeMinKb", Desc: "最小ファイルサイズ (KiB).", Default: "0", TypeName: "int", ...},
+ 		&{
+ 			Name:     "Verify",
+ 			Desc:     "アップロード後の検証",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev release publish`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ArtifactPath", Desc: "成果物へのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
  		&{
  			Name:     "Branch",
  			Desc:     "対象ブランチ",
- 			Default:  "master",
+ 			Default:  "main",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{Name: "ConnGithub", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "ChunkSizeKb",
- 			Desc:     "アップロードチャンク容量(Kバイト)",
- 			Default:  "65536",
- 			TypeName: "essentials.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(65536)},
- 		},
+ 		&{
+ 			Name:     "BatchSize",
+ 			Desc:     "一括コミットサイズ",
+ 			Default:  "50",
+ 			TypeName: "essentials.model.mo_int.range_int",
+ 			TypeAttr: map[string]interface{}{"max": float64(1000), "min": float64(1), "value": float64(50)},
+ 		},
  		&{Name: "Delete", Desc: "ローカルで削除されたファイルがある場合はDrop"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "DropboxPath", Desc: "転送先のDropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		... // 5 identical elements
  		&{Name: "Overwrite", Desc: "ターゲットパス上に既存のファイルが存在する\xe5"..., Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
- 		&{
- 			Name:     "WorkPath",
- 			Desc:     "テンポラリパス",
- 			TypeName: "essentials.model.mo_string.opt_string",
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
