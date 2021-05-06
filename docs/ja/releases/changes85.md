---
layout: release
title: リリースの変更点 84
lang: ja
---

# `リリース 84` から `リリース 85` までの変更点

# 追加されたコマンド


| コマンド                                  | タイトル                                                               |
|-------------------------------------------|------------------------------------------------------------------------|
| dev stage griddata                        | グリッドデータテスト                                                   |
| services google sheets sheet append       | スプレッドシートにデータを追加する                                     |
| services google sheets sheet clear        | スプレッドシートから値をクリアする                                     |
| services google sheets sheet export       | シートデータのエクスポート                                             |
| services google sheets sheet import       | スプレッドシートにデータをインポート                                   |
| services google sheets sheet list         | スプレッドシートのシート一覧                                           |
| services google sheets spreadsheet create | 新しいスプレッドシートの作成                                           |
| util date today                           | 現在の日付を表示                                                       |
| util datetime now                         | 現在の日時を表示                                                       |
| util decode base_32                       | Base32 (RFC 4648) 形式からテキストをデコードします                     |
| util decode base_64                       | Base64 (RFC 4648) フォーマットからテキストをデコードします             |
| util encode base_32                       | テキストをBase32(RFC 4648)形式にエンコード                             |
| util encode base_64                       | テキストをBase64(RFC 4648)形式にエンコード                             |
| util time now                             | 現在の時刻を表示                                                       |
| util unixtime format                      | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット |
| util unixtime now                         | UNIX時間で現在の時刻を表示する                                         |
| util xlsx create                          | 空のスプレッドシートを作成する                                         |
| util xlsx sheet export                    | xlsxファイルからデータをエクスポート                                   |
| util xlsx sheet import                    | データをxlsxファイルにインポート                                       |
| util xlsx sheet list                      | xlsxファイルのシート一覧                                               |



# コマンド仕様の変更: `config disable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Key", Desc: "機能キー.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `config enable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Key", Desc: "機能キー.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `config features`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev benchmark local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "NumFiles", Desc: "ファイル数.", Default: "1000", TypeName: "int", ...}, &{Name: "Path", Desc: "作成するパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "SizeMaxKb", Desc: "最大ファイルサイズ (KiB).", Default: "2048", TypeName: "int", ...}, &{Name: "SizeMinKb", Desc: "最小ファイルサイズ (KiB).", Default: "0", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev benchmark upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ChunkSizeKb", Desc: "チャンクサイズをKiB単位でアップロード", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "NumFiles", Desc: "ファイル数.", Default: "1000", TypeName: "int", ...}, &{Name: "Path", Desc: "Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "アップロード先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}, &{Name: "SizeKb", Desc: "サイズ(KB)", Default: "1024", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev build catalogue`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev build doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Badge", Desc: "ビルド状態のバッジを含める", Default: "true", TypeName: "bool", ...}, &{Name: "CommandPath", Desc: "コマンドマニュアルを作成する相対パス", Default: "doc/generated/", TypeName: "string", ...}, &{Name: "DocLang", Desc: "言語", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Readme", Desc: "README のファイル名", Default: "README.md", TypeName: "string", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev build license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DestPath", Desc: "出力先パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "SourcePath", Desc: "ライセンスへのパス (go-licenses 出力フォルダ)", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev build preflight`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev build readme`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "README.txtの作成パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev ci artifact connect`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Full", Desc: "アカウントの別名", Default: "deploy", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev ci artifact up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "アップロードするローカルファイルのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "PeerName", Desc: "アカウントの別名", Default: "deploy", TypeName: "string", ...}, &{Name: "Timeout", Desc: "処理タイムアウト(秒)", Default: "60", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev ci auth connect`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Audit", Desc: "Dropbox Business Audit スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "File", Desc: "Dropbox Business member file access スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Full", Desc: "Dropbox user full access スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Github", Desc: "GitHubへのデプロイメントのためのアカウント別名", Default: "deploy", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev ci auth export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Audit", Desc: "Dropbox Business Audit スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "File", Desc: "Dropbox Business member file access スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Full", Desc: "Dropbox user full access スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Info", Desc: "Dropbox Business info スコープで認証", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev ci auth import`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EnvName", Desc: "環境変数名", Default: "TOOLBOX_ENDTOEND_TOKEN", TypeName: "string", ...}, &{Name: "PeerName", Desc: "アカウントの別名", Default: "end_to_end_test", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev diag endpoint`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "JobId", Desc: "検査するJob ID", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev diag throughput`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Bucket", Desc: "バケットサイズ (ミリ秒)", Default: "1000", TypeName: "int", ...}, &{Name: "EndpointName", Desc: "エンドポイントによりフィルター. 名前による\xe5\xae"...}, &{Name: "EndpointNamePrefix", Desc: "エンドポイントによりフィルター. 名前の前方\xe4\xb8"...}, &{Name: "EndpointNameSuffix", Desc: "エンドポイントによりフィルター. 名前の後方\xe4\xb8"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev kvs dump`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "KVSデータへのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev release candidate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```

## 追加されたレポート


| 名称   | 説明             |
|--------|------------------|
| result | レシピテスト結果 |


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
  	... // 5 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ArtifactPath", Desc: "成果物へのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "対象ブランチ", Default: "master", TypeName: "string", ...}, &{Name: "ConnGithub", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "SkipTests", Desc: "エンドツーエンドテストをスキップします.", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```

## 追加されたレポート


| 名称   | 説明             |
|--------|------------------|
| commit | コミット情報     |
| result | レシピテスト結果 |


# コマンド仕様の変更: `dev replay approve`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "ジョブID.", TypeName: "string"}, &{Name: "Name", Desc: "承認されたレシピの追加名", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "ReplayPath", Desc: "リプレイのリポジトリパス指定されていない場\xe5"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "WorkspacePath", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev replay bundle`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "PeerName", Desc: "配置用アカウントのエイリアス", Default: "deploy", TypeName: "string", ...}, &{Name: "ReplayPath", Desc: "リプレイのリポジトリパス指定されていない場\xe5"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "ResultsPath", Desc: "テストの失敗結果の送信先パス", Default: "/watermint-toolbox-logs/{% raw %}{{.{% endraw %}Date}}-{% raw %}{{.{% endraw %}Time}}/{% raw %}{{.{% endraw %}Random}}", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl", ...}, &{Name: "Timeout", Desc: "テスト結果のアップロード操作のタイムアウト", Default: "60", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev replay recipe`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "ジョブID.", TypeName: "string"}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev replay remote`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ReplayUrl", Desc: "リプレイバンドル共有リンクURL", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev spec diff`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FilePath", Desc: "出力先ファイルパス", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Lang", Desc: "言語", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Release1", Desc: "リリース名1", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Release2", Desc: "リリース名2", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev spec doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FilePath", Desc: "ファイルパス", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Lang", Desc: "言語", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev stage gmail`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーID. 特別な値 'me' は、認証されたユーザ"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev stage gui`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev stage scoped`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Individual", Desc: "個人向けのアカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}, &{Name: "Team", Desc: "チーム向けのアカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev stage teamfolder`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev test echo`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Text", Desc: "エコーするテキストインポート先のパス", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev test kvsfootprint`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Count", Desc: "テスト回数", Default: "1", TypeName: "int", ...}, &{Name: "Duplicate", Desc: "重複レコードを作成します", Default: "1", TypeName: "int", ...}, &{Name: "NumEntries", Desc: "書き込みするエントリ数", Default: "1", TypeName: "int", ...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev test monkey`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Distribution", Desc: "ファイル・フォルダの分布数", Default: "10000", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Extension", Desc: "カンマ区切りの拡張子一覧", Default: "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"..., TypeName: "string", ...}, &{Name: "Path", Desc: "モンキーテストパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev test recipe`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "All", Desc: "全てのレシピをテストします", Default: "false", TypeName: "bool", ...}, &{Name: "NoTimeout", Desc: "レシピのテスト時にタイムアウトしません", Default: "false", TypeName: "bool", ...}, &{Name: "Single", Desc: "テストするレシピ名", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Verbose", Desc: "テスト結果の詳細出力", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```

## 追加されたレポート


| 名称   | 説明             |
|--------|------------------|
| result | レシピテスト結果 |


# コマンド仕様の変更: `dev test resources`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev util anonymise`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "JobIdName", Desc: "ジョブID名にてフィルター. 名前による完全一致"...}, &{Name: "JobIdNamePrefix", Desc: "ジョブID名にてフィルター. 名前の前方一致によ"...}, &{Name: "JobIdNameSuffix", Desc: "ジョブID名にてフィルター. 名前の後方一致によ"...}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev util curl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BufferSize", Desc: "バッファのサイズ", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Record", Desc: "テスト用に直接テストレコードを指定", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev util image jpeg`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Count", Desc: "生成するファイル数", Default: "10", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Height", Desc: "高さ", Default: "1080", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "NamePrefix", Desc: "ファイル名のプリフィックス", Default: "test_image", TypeName: "string", ...}, &{Name: "Path", Desc: "ファイルを生成するパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `dev util wait`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Seconds", Desc: "指定秒数待機", Default: "1", TypeName: "essentials.model.mo_int.range_int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file archive local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "宛先フォルダのパス. このコマンドは、パス上\xe3\x81"..., TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "ExcludeFolders", Desc: "フォルダを除外する", Default: "false", TypeName: "bool", ...}, &{Name: "IncludeSystemFiles", Desc: "システムファイルを含める", Default: "false", TypeName: "bool", ...}, &{Name: "Preview", Desc: "プレビューモード", Default: "false", TypeName: "bool", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Left", Desc: "アカウントの別名 (左)", Default: "left", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "LeftPath", Desc: "アカウントのルートからのパス (左)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Right", Desc: "アカウントの別名 (右)", Default: "right", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "RightPath", Desc: "アカウントのルートからのパス (右)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Src", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file dispatch local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Preview", Desc: "プレビューモード", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "ダウンロードするファイルパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "保存先ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "エクスポートするDropbox上のドキュメントパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "保存先ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Path", Desc: "インポート先のパス", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "インポート先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "URL", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "IncludeDeleted", Desc: "削除済みファイルを含める", Default: "false", TypeName: "bool", ...}, &{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "再起的に一覧を実行", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file lock acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "ロックするファイルパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "操作バッチサイズ", Default: "100", TypeName: "int", ...}, &{Name: "Path", Desc: "ロックを解除するためのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file lock batch acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "操作バッチサイズ", Default: "100", TypeName: "int", ...}, &{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file lock batch release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "ファイルへのパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DryRun", Desc: "リハーサルを行います", Default: "true", TypeName: "bool", ...}, &{Name: "From", Desc: "統合するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "KeepEmptyFolder", Desc: "統合後に空となったフォルダを維持する", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Src", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "アカウントの別名 (宛先)", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "DstPath", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Src", Desc: "アカウントの別名 (元)", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "SrcPath", Desc: "元のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "指定されたファイルカテゴリに検索を限定しま\xe3"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}}, &{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "指定されたファイルカテゴリに検索を限定しま\xe3"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}}, &{Name: "Extension", Desc: "指定されたファイル拡張子に検索を限定します.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "検索対象とするユーザーのDropbox上のパス.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "すべてのファイルとフォルダの深さのフォルダ\xe3"..., Default: "2", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "スキャンするパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file sync down`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Delete", Desc: "Dropbox上のファイルを削除した場合、ローカルフ"..., Default: "false", TypeName: "bool", ...}, &{Name: "DropboxPath", Desc: "Dropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "ローカルパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "NameDisableIgnore", Desc: "名前によるフィルター. システムファイルと除\xe5\xa4"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file sync online`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Delete", Desc: "ソースパスでファイルが削除された場合はファ\xe3"..., Default: "false", TypeName: "bool", ...}, &{Name: "Dst", Desc: "宛先のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "NameDisableIgnore", Desc: "名前によるフィルター. システムファイルと除\xe5\xa4"...}, &{Name: "NameName", Desc: "名前によるフィルター. 名前による完全一致で\xe3\x83"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ChunkSizeKb", Desc: "アップロードチャンク容量(Kバイト)", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Delete", Desc: "ローカルで削除されたファイルがある場合はDrop"..., Default: "false", TypeName: "bool", ...}, &{Name: "DropboxPath", Desc: "転送先のDropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "ローカルファイルのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `file watch`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "監視対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "パスを再起的に監視", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `filerequest create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AllowLateUploads", Desc: "設定した場合、期限を過ぎてもアップロードを\xe8"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Deadline", Desc: "ファイルリクエストの締め切り.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Path", Desc: "ファイルをアップロードするDropbox上のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `filerequest delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `filerequest delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Force", Desc: "ファイリクエストを強制的に削除する.", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "ファイルリクエストのURL", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ManagementType", Desc: "グループ管理タイプ. `company_managed` または `user_m"..., Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Name", Desc: "グループ名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "グループ名リストのデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "グループ名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...}, &{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...}, &{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...}, &{Name: "GroupName", Desc: "グループ名でフィルタリングします. 名前によ\xe3\x82"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "GroupName", Desc: "グループ名", TypeName: "string"}, &{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "GroupName", Desc: "グループ名称", TypeName: "string"}, &{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "CurrentName", Desc: "現在のグループ名", TypeName: "string"}, &{Name: "NewName", Desc: "新しいグループ名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `image info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "画像へのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job history archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Days", Desc: "目標日数", Default: "7", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job history delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Days", Desc: "目標日数", Default: "28", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job history list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job history ship`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "アップロード先Dropboxパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job log jobid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "ジョブID", TypeName: "string"}, &{Name: "Kind", Desc: "ログの種別", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job log kind`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Kind", Desc: "ログの種別.", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `job log last`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Kind", Desc: "ログの種別", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "ワークスペースへのパス.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "TransferDestMember", Desc: "指定された場合は、指定ユーザーに削除するメ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "TransferNotifyAdminEmailOnError", Desc: "指定された場合は、転送時にエラーが発生した\xe6"..., TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "RevokeTeamShares", Desc: "指定した場合にはユーザーからチームが保有す\xe3"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "バッチ処理サイズ", Default: "100", TypeName: "int", ...}, &{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"}, &{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"}, &{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "メンバーのメールアドレス", TypeName: "string"}, &{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member file permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "チームメンバーのメールアドレス.", TypeName: "string"}, &{Name: "Path", Desc: "削除対象のパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...}, &{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...}, &{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...}, &{Name: "MemberEmail", Desc: "メンバーのメールアドレスでフィルタリングし\xe3"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member folder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstMemberEmail", Desc: "コピー先チームメンバーのメールアドレス", TypeName: "string"}, &{Name: "DstPath", Desc: "コピー先チームメンバーのパス. ルート (/) パス"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "SrcMemberEmail", Desc: "送信元チームメンバーのメールアドレス.", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "SilentInvite", Desc: "ウエルカムメールを送信しません (SSOとドメイ\xe3\x83"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "IncludeDeleted", Desc: "削除済メンバーを含めます.", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member quota update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "Quota", Desc: "カスタムの容量制限 (1TB = 1024GB). 0の場合、容量\xe5"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "Silent", Desc: "招待メールを送信しません (SSOが必須となります)", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "宛先チーム; チームのファイルアクセス", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Src", Desc: "元チーム; チームのファイルアクセス", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member update email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "UpdateUnverified", Desc: "アカウントのメールアドレスが確認されていな\xe3"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member update externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member update invisible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member update profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `member update visible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services asana team list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "WorkspaceName", Desc: "ワークスペースの名前または GID。 名前による\xe5\xae"...}, &{Name: "WorkspaceNamePrefix", Desc: "ワークスペースの名前または GID。 名前の前方\xe4\xb8"...}, &{Name: "WorkspaceNameSuffix", Desc: "ワークスペースの名前または GID。 名前の後方\xe4\xb8"...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services asana team project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "TeamName", Desc: "チーム名またはGID 名前による完全一致でフィル"...}, &{Name: "TeamNamePrefix", Desc: "チーム名またはGID 名前の前方一致によるフィル"...}, &{Name: "TeamNameSuffix", Desc: "チーム名またはGID 名前の後方一致によるフィル"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services asana team task list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "ProjectName", Desc: "プロジェクトの名前またはGID 名前による完全一"...}, &{Name: "ProjectNamePrefix", Desc: "プロジェクトの名前またはGID 名前の前方一致に"...}, &{Name: "ProjectNameSuffix", Desc: "プロジェクトの名前またはGID 名前の後方一致に"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services asana workspace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services asana workspace project list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "WorkspaceName", Desc: "ワークスペースの名前または GID。 名前による\xe5\xae"...}, &{Name: "WorkspaceNamePrefix", Desc: "ワークスペースの名前または GID。 名前の前方\xe4\xb8"...}, &{Name: "WorkspaceNameSuffix", Desc: "ワークスペースの名前または GID。 名前の後方\xe4\xb8"...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github content get`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Path", Desc: "コンテンツへのパス.", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Ref", Desc: "リファレンス名", TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github content put`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Branch", Desc: "ブランチ名", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Content", Desc: "コンテンツファイルへのパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Message", Desc: "コミットメッセージ", TypeName: "string"}, &{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github issue list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github release asset download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Path", Desc: "ダウンロード パス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "リリースタグ名", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github release asset list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "リリースタグ名", TypeName: "string"}, &{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github release asset upload`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Asset", Desc: "成果物のパス", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "リリースタグ名", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github release draft`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BodyFile", Desc: "本文テキストファイルへのパスファイルはBOMな\xe3"..., TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "対象ブランチ名", TypeName: "string"}, &{Name: "Name", Desc: "リリース名称", TypeName: "string"}, &{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github release list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services github tag create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "レポジトリの所有者", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "レポジトリ名", TypeName: "string"}, &{Name: "Sha1", Desc: "コミットのSHA1ハッシュ", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail filter add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "ラベルが存在しない場合はラベルを作成します.", Default: "false", TypeName: "bool", ...}, &{Name: "AddLabels", Desc: "','で区切られたメッセージに追加するラベルの\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "CriteriaExcludeChats", Desc: "チャットを除外するかどうか", Default: "false", TypeName: "bool", ...}, &{Name: "CriteriaFrom", Desc: "送信者の表示名またはメールアドレス.", TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail filter batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "ラベルが存在しない場合はラベルを作成します.", Default: "false", TypeName: "bool", ...}, &{Name: "ApplyToExistingMessages", Desc: "クエリを満たす既存のメッセージにラベルを適\xe7"..., Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail filter delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "フィルターID", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail filter list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ColorBackground", Desc: "背景色.", TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}}}, &{Name: "ColorText", Desc: "テキストの色.", TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}}}, &{Name: "LabelListVisibility", Desc: "Gmail ウェブインタフェースのラベルリストのラ\xe3"..., Default: "labelShow", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "MessageListVisibility", Desc: "Gmail ウェブインターフェースのメッセージリス\xe3"..., Default: "show", TypeName: "essentials.model.mo_string.select_string", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail label delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "ラベル名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail label list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail label rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "CurrentName", Desc: "現在のラベル名", TypeName: "string"}, &{Name: "NewName", Desc: "新しいラベル名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail message label add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "ラベルが存在しない場合はラベルを作成します.", Default: "false", TypeName: "bool", ...}, &{Name: "Label", Desc: "このメッセージを追加するラベル名.", TypeName: "string"}, &{Name: "MessageId", Desc: "メッセージの不変ID", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail message label delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Label", Desc: "このメッセージを削除するラベル名.", TypeName: "string"}, &{Name: "MessageId", Desc: "メッセージの不変ID", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail message list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Format", Desc: "メッセージを返すフォーマット. ", Default: "metadata", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "IncludeSpamTrash", Desc: "SPAMやTRASHからのメッセージを結果に含める.", Default: "false", TypeName: "bool", ...}, &{Name: "Labels", Desc: "指定されたラベルにすべて一致するラベルを持\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "MaxResults", Desc: "返すメッセージの最大数.", Default: "20", TypeName: "int", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail message processed list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Format", Desc: "メッセージを返すフォーマット. ", Default: "metadata", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "IncludeSpamTrash", Desc: "SPAMやTRASHからのメッセージを結果に含める.", Default: "false", TypeName: "bool", ...}, &{Name: "Labels", Desc: "指定されたラベルにすべて一致するラベルを持\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "MaxResults", Desc: "返すメッセージの最大数.", Default: "20", TypeName: "int", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services google mail thread list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "ユーザーのメールアドレス. 特別な値meは、認証"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `services slack conversation list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "Peer", TypeName: "domain.slack.api.work_conn_impl.conn_slack_api", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Expires", Desc: "共有リンクの有効期限日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Password", Desc: "パスワード", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `sharedlink delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "共有リンクを削除するファイルまたはフォルダ\xe3"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "フォルダ階層をたどって削除します", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `sharedlink file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Password", Desc: "共有リンクのパスワード", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "共有リンクのURL", TypeName: "domain.dropbox.model.mo_url.url_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "一つのイベントカテゴリのみを返すようなフィ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "File", Desc: "メールアドレスリストのファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "イベントのカテゴリ", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndDate", Desc: "終了日", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartDate", Desc: "開始日", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "一つのイベントカテゴリのみを返すようなフィ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "一つのイベントカテゴリのみを返すようなフィ\xe3"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "終了日時 (該当同時刻を含まない).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartTime", Desc: "開始日時 (該当時刻を含む)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...}, &{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...}, &{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...}, &{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team content mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "メンバーをフィルタリングします. メールアド\xe3\x83"...}, &{Name: "MemberName", Desc: "メンバーをフィルタリングします. 名前による\xe5\xae"...}, &{Name: "MemberNamePrefix", Desc: "メンバーをフィルタリングします. 名前の前方\xe4\xb8"...}, &{Name: "MemberNameSuffix", Desc: "メンバーをフィルタリングします. 名前の後方\xe4\xb8"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team content policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...}, &{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...}, &{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DeleteOnUnlink", Desc: "デバイスリンク解除時にファイルを削除します", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "データファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team diag explorer`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "All", Desc: "追加のレポートを含める", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Dropbox Business ファイルアクアセス", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Info", Desc: "Dropbox Business 情報アクセス", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}, &{Name: "Mgmt", Desc: "Dropbox Business 管理", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team filerequest clone`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "IncludeDeleted", Desc: "指定された場合、削除済みのファイルやフォル\xe3"..., Default: "false", TypeName: "bool", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "フォルダ階層数の指定", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AllColumns", Desc: "全てのカラムを表示します", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "終了日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "開始日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Visibility", Desc: "出力するリンクを可視性にてフィルターします "..., Default: "all", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "At", Desc: "新しい有効期限の日時", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Days", Desc: "新しい有効期限までの日時", Default: "0", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Visibility", Desc: "対象となるリンクの公開範囲", Default: "public", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "SyncSetting", Desc: "チームフォルダの同期設定", Default: "default", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstPeerName", Desc: "宛先チームのアカウント別名", Default: "dst", TypeName: "string", ...}, &{Name: "File", Desc: "チームフォルダ名のデータファイル", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "SrcPeerName", Desc: "元チームのアカウント別名", Default: "src", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "操作バッチサイズ", Default: "100", TypeName: "int", ...}, &{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "チームフォルダ名", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "パス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "チームフォルダ名", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "ロックを解除するパス", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "チームフォルダ名", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "深さ", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "FolderName", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNamePrefix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, &{Name: "FolderNameSuffix", Desc: "名前に一致するフォルダのみをリストアップし\xe3"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AdminGroupName", Desc: "管理者操作のための仮グループ名", Default: "watermint-toolbox-admin", TypeName: "string", ...}, &{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AdminGroupName", Desc: "管理者操作のための仮グループ名", Default: "watermint-toolbox-admin", TypeName: "string", ...}, &{Name: "File", Desc: "データファイルへのパス", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...}, &{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...}, &{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...}, &{Name: "MemberTypeExternal", Desc: "フォルダメンバーによるフィルター. 外部メン\xe3\x83"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder partial replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "宛先チームの別名", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "DstPath", Desc: "チームフォルダからの相対パス (チームフォル\xe3\x83"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "DstTeamFolderName", Desc: "送信先チームフォルダ名", TypeName: "string"}, &{Name: "Src", Desc: "転送元チームの別名", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "フォルダ名によるフィルター. 名前による完全\xe4\xb8"...}, &{Name: "FolderNamePrefix", Desc: "フォルダ名によるフィルター. 名前の前方一致\xe3\x81"...}, &{Name: "FolderNameSuffix", Desc: "フォルダ名によるフィルター. 名前の後方一致\xe3\x81"...}, &{Name: "Peer", Desc: "アカウントの別名", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstPeerName", Desc: "宛先チームのアカウント別名", Default: "dst", TypeName: "string", ...}, &{Name: "Name", Desc: "チームフォルダ名", TypeName: "string"}, &{Name: "SrcPeerName", Desc: "元チームのアカウント別名", Default: "src", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# コマンド仕様の変更: `version`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
