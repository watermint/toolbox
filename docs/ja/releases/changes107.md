---
layout: release
title: リリースの変更点 106
lang: ja
---

# `リリース 106` から `リリース 107` までの変更点

# 追加されたコマンド

| コマンド                            | タイトル                                                                             |
|-------------------------------------|--------------------------------------------------------------------------------------|
| dev build compile                   | ビルドスクリプトの作成                                                               |
| dev build target                    | ターゲットビルドスクリプトの生成                                                     |
| dev kvs benchmark                   | KVSエンジンのベンチマーク                                                            |
| dev stage encoding                  | エンコードテストコマンド（指定したエンコード名でダミーファイルをアップロードします） |
| services google calendar event list | Googleカレンダーのイベントを一覧表示                                                 |
| util archive unzip                  | ZIPアーカイブファイルを解凍する                                                      |
| util archive zip                    | 対象ファイルをZIPアーカイブに圧縮する                                                |
| util database exec                  | SQLite3データベースファイルへのクエリ実行                                            |
| util database query                 | SQLite3データベースへの問い合わせ                                                    |
| util file hash                      | ファイルダイジェストの表示                                                           |
| util image exif                     | 画像ファイルのEXIFメタデータを表示                                                   |
| util monitor client                 | デバイスモニタークライアントを起動する                                               |
| util net download                   | ファイルをダウンロードする                                                           |
| util text case down                 | 小文字のテキストを表示する                                                           |
| util text case up                   | 大文字のテキストを表示する                                                           |
| util text encoding from             | 指定されたエンコーディングからUTF-8テキストファイルに変換します.                     |
| util text encoding to               | UTF-8テキストファイルから指定されたエンコーディングに変換する.                       |

# 削除されたコマンド

| コマンド   | タイトル                           |
|------------|------------------------------------|
| image info | 画像ファイルのEXIF情報を表示します |

# コマンド仕様の変更: `config disable`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "config disable",
- 	CliArgs:         "",
+ 	CliArgs:         "-key FEATURE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `config enable`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "config enable",
- 	CliArgs:         "",
+ 	CliArgs:         "-key FEATURE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev benchmark local`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev benchmark local",
- 	CliArgs:         "-path /LOCAL/PATH/TO/PROCESS",
+ 	CliArgs:         `-num-files NUM -path /LOCAL/PATH/TO/PROCESS -size-max-kb NUM -size-min-kb NUM"`,
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev benchmark upload`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev benchmark upload",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/PROCESS",
+ 	CliArgs:         "-num-files NUM -path /DROPBOX/PATH/TO/PROCESS -size-max-kb NUM -size-min-kb NUM",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev benchmark uploadlink",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/UPLOAD",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/UPLOAD -size-kb NUM",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev build package`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev build package",
  	CliArgs: strings.Join({
  		"-build-path /LOCAL/PATH/",
- 		"OF/build -deploy-path /DROPBOX/PATH/TO/deploy -dest-path /LOCAL/",
- 		"PATH/TO/save_package",
+ 		"TO/build -dist-path /LOCAL/PATH/TO/dist -platform PLATFORM_TYPE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BuildPath", Desc: "バイナリへのフルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "DeployPath", Desc: "デプロイ先フォルダパス(リモート)", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
- 			Name:    "DestPath",
+ 			Name:    "DistPath",
  			Desc:    "パッケージの保存先フォルダのパス(ローカル)",
  			Default: "",
  			... // 2 identical fields
  		},
  		&{Name: "Platform", Desc: "win/linux/macなどのプラットフォーム名", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev ci artifact up`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev ci artifact up",
  	CliArgs: strings.Join({
  		"-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF",
  		"/ARTIFACT",
+ 		" -timeout NUM",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev replay approve`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev replay approve",
- 	CliArgs:         "",
+ 	CliArgs:         "-id JOB_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev replay recipe`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev replay recipe",
- 	CliArgs:         "",
+ 	CliArgs:         "-id JOB_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev stage griddata`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev stage griddata",
- 	CliArgs:         "",
+ 	CliArgs:         "-in /LOCAL/PATH/TO/INPUT.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev test echo`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test echo",
- 	CliArgs:         "",
+ 	CliArgs:         "-text VALUE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev test panic`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test panic",
- 	CliArgs:         "",
+ 	CliArgs:         "-panic-type VALUE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `dev test setup teamsharedlink`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "dev test setup teamsharedlink",
- 	CliArgs:         "",
+ 	CliArgs:         "-group GROUP_NAME -num-links-per-member NUM -query QUERY",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file paper append`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper append",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/append",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/APPEND",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file paper create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper create",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/create",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/CREATE",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file paper overwrite`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper overwrite",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/overwrite",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/OVERWRIT",
+ 		"E",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file paper prepend`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper prepend",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/prepend",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/PREPEND",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file revision download`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file revision download",
  	CliArgs: strings.Join({
  		"-local-path /LOCAL/PATH/TO/DOWNLOAD",
+ 		" -revision REVISION",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file revision restore`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file revision restore",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE -revision REVISION",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file search content`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file search content",
- 	CliArgs:         "",
+ 	CliArgs:         "-query QUERY",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file search name`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file search name",
- 	CliArgs:         "",
+ 	CliArgs:         "-query QUERY",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `file share info`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file share info",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/GET_INFO",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Path",
  			Desc:     "ファイル",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `filerequest create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "(非可逆な操作です)",
  	Path:    "filerequest create",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/OF/FILE",
+ 		"_",
  		"REQUEST",
+ 		" -title TITLE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `filerequest delete url`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "filerequest delete url",
- 	CliArgs:         "",
+ 	CliArgs:         "-url URL",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Force", Desc: "ファイリクエストを強制的に削除する.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:     "Url",
  			Desc:     "ファイルリクエストのURL",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.dropbox.model.mo_url.url_impl",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `group add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "group add",
- 	CliArgs:         "",
+ 	CliArgs:         "-name グループ名",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `group member add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "group member add",
- 	CliArgs:         "",
+ 	CliArgs:         "-group-name GROUP_NAME -member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `group member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "group member delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-group-name GROUP_NAME -member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `group rename`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "group rename",
- 	CliArgs:         "",
+ 	CliArgs:         "-current-name CURRENT_NAME -new-name NEW_NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `job log jobid`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job log jobid",
- 	CliArgs:         "",
+ 	CliArgs:         "-id JOB_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `member file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "member file lock all release",
  	CliArgs: strings.Join({
  		"-",
+ 		"member-email VALUE -",
  		"path /DROPBOX/PATH/TO/RELEASE/LOCK",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `member file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member file lock list",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST",
+ 	CliArgs:         "-member-email EMAIL -path /DROPBOX/PATH/TO/LIST_LOCK",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `member file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "member file lock release",
  	CliArgs: strings.Join({
  		"-",
+ 		"member-email VALUE -",
  		"path /DROPBOX/PATH/TO/RELEASE/LOCK",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `member file permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "member file permdelete",
  	CliArgs: strings.Join({
  		"-",
+ 		"member-email EMAIL -",
  		"path /DROPBOX/PATH/TO/",
- 		"PERM_",
  		"DELETE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `member suspend`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member suspend",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `member unsuspend`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member unsuspend",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github content get`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services github content get",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPOSITORY -path PATH",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github content put`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services github content put",
- 	CliArgs:         "-content /LOCAL/PATH/TO/content",
+ 	CliArgs:         " -owner OWNER -repository REPO -path PATH -content /LOCAL/PATH/TO/content -message MSG",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "services github issue list",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github release asset download`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "services github release asset download",
- 	CliArgs:         "-path /LOCAL/PATH/TO/DOWNLOAD",
+ 	CliArgs:         "-owner OWNER -repository REPO -path /LOCAL/PATH/TO/DOWNLOAD -release RELEASE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github release asset list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "services github release asset list",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO -release RELEASE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github release asset upload`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装かつ非可逆な操作です)",
  	Path:            "services github release asset upload",
- 	CliArgs:         "-asset /LOCAL/PATH/TO/assets",
+ 	CliArgs:         "-owner OWNER -repository REPO -release RELEASE -asset /LOCAL/PATH/TO/assets",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github release draft`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "services github release draft",
  	CliArgs: strings.Join({
  		"-",
- 		"body-file /LOCAL/PATH/TO/body.txt",
+ 		"owner OWNER -repository REPO -body-file /LOCAL/PATH/TO/BODY.txt ",
+ 		"-branch BRANCH -name NAME -tag TAG",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github release list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装です)",
  	Path:            "services github release list",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装かつ非可逆な操作です)",
  	Path:            "services github tag create",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO -sha1 SHA -tag TAG",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail filter delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail filter delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-id ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail label add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail label add",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail label delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail label delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail label rename`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail label rename",
- 	CliArgs:         "",
+ 	CliArgs:         "-current-name CURRENT_NAME -new-name NEW_NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail message label add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail message label add",
- 	CliArgs:         "",
+ 	CliArgs:         "-label LABEL -message-id MSG_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail message label delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail message label delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-label LABEL -message-id MSG_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail message send`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "services google mail message send",
- 	CliArgs:         "",
+ 	CliArgs:         "-body /LOCAL/PATH/TO/INPUT.txt -subject SUBJECT -to TO",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail sendas add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail sendas add",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google mail sendas delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail sendas delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet append`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet append",
- 	CliArgs:         "",
+ 	CliArgs:         "-data /LOCAL/PATH/TO/INPUT.csv -id GOOGLE_SHEET_ID -range RANGE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet clear`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet clear",
- 	CliArgs:         "",
+ 	CliArgs:         "-id GOOGLE_SPREADSHEET_ID -range RANGE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet export`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet export",
- 	CliArgs:         "",
+ 	CliArgs:         "-id GOOGLE_SPREADSHEET_ID -range RANGE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet import`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet import",
- 	CliArgs:         "",
+ 	CliArgs:         "-data /LOCAL/PATH/TO/INPUT.csv -id VALUE -range VALUE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google sheets sheet list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet list",
- 	CliArgs:         "",
+ 	CliArgs:         "-id GOOGLE_SPREADSHEET_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `services google sheets spreadsheet create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets spreadsheet create",
- 	CliArgs:         "",
+ 	CliArgs:         "-title TITLE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder leave`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedfolder leave",
- 	CliArgs:         "",
+ 	CliArgs:         "-shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "sharedfolder member add",
  	CliArgs: strings.Join({
  		"-",
- 		"path /SHARED_FOLDER",
+ 		"email EMAIL -path /DROPBOX",
  		"/PATH/TO/ADD",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "sharedfolder member delete",
  	CliArgs: strings.Join({
  		"-",
- 		"path /SHARED_FOLDER",
+ 		"email EMAIL -path /DROPBOX",
  		"/PATH/TO/DELETE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder mount add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedfolder mount add",
- 	CliArgs:         "",
+ 	CliArgs:         "-shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `sharedfolder mount delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedfolder mount delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team activity daily event`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team activity daily event",
- 	CliArgs:         "",
+ 	CliArgs:         "-start-date DATE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team admin group role add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin group role add",
- 	CliArgs:         "",
+ 	CliArgs:         "-group GROUP_NAME -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team admin group role delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin group role delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-exception-group GROUP_NAME -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team admin role add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin role add",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team admin role clear`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin role clear",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team admin role delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin role delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder list",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount add",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL -shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL -shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount list",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team runas sharedfolder mount mountable`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount mountable",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `team sharedlink delete member`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "team sharedlink delete member",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `teamfolder add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(非可逆な操作です)",
  	Path:            "teamfolder add",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `teamfolder archive`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
+ 				string("team_data.content.read"),
+ 				string("team_data.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "teamfolder file lock all release",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/TO/RELEASE",
- 		"/LOCK",
+ 		" -team-folder NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "teamfolder file lock list",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST -team-folder NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `teamfolder file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "teamfolder file lock release",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/TO/RELEASE",
- 		"/LOCK",
+ 		" -team-folder NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `teamfolder replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(試験的実装かつ非可逆な操作です)",
  	Path:            "teamfolder replication",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `util decode base32`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util decode base32",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "パディングなし", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "テキスト",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "デコードするテキスト"}},
  	JsonInput:      {},
  }
```
# コマンド仕様の変更: `util decode base64`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util decode base64",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "パディングなし", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "テキスト",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "デコードするテキスト"}},
  	JsonInput:      {},
  }
```
# コマンド仕様の変更: `util encode base32`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util encode base32",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "パディングなし", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "テキスト",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "エンコードするテキスト"}},
  	JsonInput:      {},
  }
```
# コマンド仕様の変更: `util encode base64`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util encode base64",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "パディングなし", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "テキスト",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "デコードするテキスト"}},
  	JsonInput:      {},
  }
```
# コマンド仕様の変更: `util qrcode create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util qrcode create",
  	CliArgs: strings.Join({
  		"-out /LOCAL/PATH/TO/",
- 		"create_qrcode.png",
+ 		"OUT.png -text /LOCAL/PATH/TO/INPUT.txt",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Out", Desc: "ファイル名付きの出力パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Size", Desc: "画像解像度(ピクセル)", Default: "256", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name:     "Text",
  			Desc:     "テキストデータ",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "テキスト"}},
  	JsonInput:      {},
  }
```
# コマンド仕様の変更: `util qrcode wifi`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util qrcode wifi",
  	CliArgs: strings.Join({
  		"-out /LOCAL/PATH/TO/",
- 		"create_qrcode.png",
+ 		"OUT.png -ssid SSID",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `util xlsx create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util xlsx create",
  	CliArgs: strings.Join({
  		"-file /LOCAL/PATH/TO/",
- 		"create.xlsx",
+ 		"CREATE.xlsx -sheet SHEET_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `util xlsx sheet export`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util xlsx sheet export",
  	CliArgs: strings.Join({
  		"-file /LOCAL/PATH/TO/",
- 		"export.xlsx",
+ 		"EXPORT.xlsx -sheet SHEET_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
# コマンド仕様の変更: `util xlsx sheet import`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util xlsx sheet import",
  	CliArgs: strings.Join({
  		"-",
- 		"file /LOCAL/PATH/TO/import.xlsx",
+ 		"data /LOCAL/PATH/TO/INPUT.csv -file /LOCAL/PATH/TO/TARGET.xlsx -",
+ 		"sheet SHEET_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
