---
layout: release
title: リリースの変更点 140
lang: ja
---

# `リリース 140` から `リリース 141` までの変更点

# 追加されたコマンド


| コマンド                      | タイトル                                               |
|-------------------------------|--------------------------------------------------------|
| dev doc knowledge             | 縮小版ナレッジベースの生成                             |
| dev doc msg add               | 新しいメッセージを追加                                 |
| dev doc msg catalogue_options | カタログ内のすべてのレシピのオプション説明を生成する   |
| dev doc msg delete            | メッセージを削除                                       |
| dev doc msg list              | メッセージ一覧                                         |
| dev doc msg options           | SelectStringフィールドのオプション説明を生成する       |
| dev doc msg translate         | 翻訳ヘルパー                                           |
| dev doc msg update            | メッセージを更新                                       |
| dev doc msg verify            | メッセージテンプレート変数の一貫性を検証する           |
| dev doc review approve        | メッセージをレビュー済みとしてマーク                   |
| dev doc review batch          | メッセージを一括で確認および承認します                 |
| dev doc review list           | 指定した言語の確認状況を一覧表示します                 |
| dev doc review options        | 不足しているSelectStringオプションの説明をレビューする |



# コマンド仕様の変更: `config auth delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "既存の認証クレデンシャルの削除",
- 	Desc:    "",
+ 	Desc:    "特定のサービスアカウントの保存された認証クレデンシャルを削除します。アクセスの取り消し、アカウントの変更、古い認証トークンのクリーンアップが必要な場合に便利です。削除する\xe8"...,
  	Remarks: "",
  	Path:    "config auth delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `config auth list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "すべての認証情報を一覧表示",
- 	Desc:    "",
+ 	Desc:    "保存されているすべての認証クレデンシャルとその詳細（アプリケーション名、スコープ、ピア名、タイムスタンプを含む）を表示します。アクセスの監査、複数アカウントの管理、認証さ\xe3"...,
  	Remarks: "",
  	Path:    "config auth list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `config feature disable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "disable",
  	Title:   "機能を無効化します.",
- 	Desc:    "",
+ 	Desc:    "watermint toolbox設定の特定の機能を無効にします。機能は、アプリケーションの動作、パフォーマンス設定、実験的機能のさまざまな側面を制御します。機能を無効にすることで、トラブルシ\xe3\x83"...,
  	Remarks: "",
  	Path:    "config feature disable",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `config feature enable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "enable",
  	Title:   "機能を有効化します.",
- 	Desc:    "",
+ 	Desc:    "watermint toolbox設定で特定の機能を有効にします。機能はアプリケーションの動作、パフォーマンス設定、実験的機能のさまざまな側面を制御します。機能を有効にすることで、新機能へのア\xe3\x82"...,
  	Remarks: "",
  	Path:    "config feature enable",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `config feature list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "利用可能なオプション機能一覧.",
- 	Desc:    "",
+ 	Desc:    "watermint toolboxで利用可能なすべてのオプション機能を説明、現在のステータス、設定詳細とともに表示します。有効化または無効化できる機能の理解や機能設定の管理に役立ちます。",
  	Remarks: "",
  	Path:    "config feature list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `config license install`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "install",
  	Title:   "ライセンスキーのインストール",
- 	Desc:    "",
+ 	Desc:    "watermint toolboxのライセンスキーをインストールして有効化します。特定の機能、プレミアム機能、商用利用にはライセンスキーが必要な場合があります。このコマンドはライセンスキーを安\xe5\x85"...,
  	Remarks: "",
  	Path:    "config license install",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `config license list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "利用可能なライセンスキーのリスト",
- 	Desc:    "",
+ 	Desc:    "インストールされているすべてのライセンスキーとその詳細（有効期限、有効な機能、ステータス）を表示します。複数のライセンスの管理、ライセンスの有効性確認、利用可能な機能の把\xe6"...,
  	Remarks: "",
  	Path:    "config license list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file account feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "feature",
  	Title:   "Dropboxアカウントの機能一覧",
- 	Desc:    "",
+ 	Desc:    "接続されたDropboxアカウントで有効な機能と機能を取得して表示します。",
  	Remarks: "",
  	Path:    "dropbox file account feature",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file account filesystem`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "filesystem",
  	Title:   "Dropboxのファイルシステムのバージョンを表示する",
- 	Desc:    "",
+ 	Desc:    "アカウントが使用しているファイルシステムのバージョン/タイプ（個人またはチーム）を表示します。",
  	Remarks: "",
  	Path:    "dropbox file account filesystem",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file account info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "Dropboxアカウント情報",
- 	Desc:    "",
+ 	Desc:    "接続されたDropboxアカウントの名前とメールアドレスを含むプロフィール情報を表示します。",
  	Remarks: "",
  	Path:    "dropbox file account info",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file compare account`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "account",
  	Title:   "二つのアカウントのファイルを比較します",
- 	Desc:    "",
+ 	Desc:    "二つの異なるDropboxアカウント間でファイルとフォルダを比較して差分を特定します。",
  	Remarks: "",
  	Path:    "dropbox file compare account",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file compare local`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "local",
  	Title:   "ローカルフォルダとDropboxフォルダの内容を比較"...,
- 	Desc:    "",
+ 	Desc:    "ローカルファイルとフォルダをDropboxの対応するファイルと比較して差分を特定します。",
  	Remarks: "",
  	Path:    "dropbox file compare local",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "copy",
  	Title:   "ファイルをコピーします",
- 	Desc:    "",
+ 	Desc:    "同じDropboxアカウント内でファイルまたはフォルダをある場所から別の場所にコピーします。",
  	Remarks: "",
  	Path:    "dropbox file copy",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "ファイルまたはフォルダは削除します.",
- 	Desc:    "",
+ 	Desc:    "Dropboxからファイルまたはフォルダを完全に削除します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file export doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "doc",
  	Title:   "ドキュメントのエクスポート",
- 	Desc:    "",
+ 	Desc:    "Dropbox PaperドキュメントとGoogle Docsを指定された形式でローカルファイルにエクスポートします。",
  	Remarks: "(試験的実装です)",
  	Path:    "dropbox file export doc",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file export url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "URLからドキュメントをエクスポート",
- 	Desc:    "",
+ 	Desc:    "共有リンクURLからファイルをダウンロードしてDropboxからエクスポートします。",
  	Remarks: "",
  	Path:    "dropbox file export url",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file import batch url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "URLからファイルを一括インポートします",
- 	Desc:    "",
+ 	Desc:    "URLのリストから複数のファイルをダウンロードしてDropboxにインポートします。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file import batch url",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file import url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "URLからファイルをインポートします",
- 	Desc:    "",
+ 	Desc:    "指定されたURLから単一のファイルをダウンロードしてDropboxにインポートします。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file import url",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "パスのメタデータを解決",
- 	Desc:    "",
+ 	Desc:    "指定したパスのファイルまたはフォルダのメタデータとプロパティを取得します。",
  	Remarks: "",
  	Path:    "dropbox file info",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "ファイルとフォルダを一覧します",
- 	Desc:    "",
+ 	Desc:    "指定したパスのファイルとフォルダを一覧表示し、フィルタリングと再帰オプションを提供します。",
  	Remarks: "",
  	Path:    "dropbox file list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "acquire",
  	Title:   "ファイルをロック",
- 	Desc:    "",
+ 	Desc:    "ファイルに排他ロックを取得して他のユーザーが編集できないようにします。",
  	Remarks: "",
  	Path:    "dropbox file lock acquire",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "指定したパスでのすべてのロックを解除する",
- 	Desc:    "",
+ 	Desc:    "現在のユーザーがアカウント全体で保持しているすべてのファイルロックを解放します。",
  	Remarks: "",
  	Path:    "dropbox file lock all release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock batch acquire`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "acquire",
  	Title:   "複数のファイルをロックする",
- 	Desc:    "",
+ 	Desc:    "単一のバッチ操作で複数のファイルにロックを取得します。",
  	Remarks: "",
  	Path:    "dropbox file lock batch acquire",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock batch release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "複数のロックを解除",
- 	Desc:    "",
+ 	Desc:    "単一のバッチ操作で複数のファイルのロックを解放します。",
  	Remarks: "",
  	Path:    "dropbox file lock batch release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "指定したパスの下にあるロックを一覧表示します",
- 	Desc:    "",
+ 	Desc:    "現在のユーザーが保持しているすべてのファイルロックを一覧表示します。",
  	Remarks: "",
  	Path:    "dropbox file lock list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "ロックを解除します",
- 	Desc:    "",
+ 	Desc:    "特定のファイルのロックを解除し、他のユーザーが編集できるようにします。",
  	Remarks: "",
  	Path:    "dropbox file lock release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file merge`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "merge",
  	Title:   "フォルダを統合します",
- 	Desc:    "",
+ 	Desc:    "あるフォルダの内容を別のフォルダにマージし、ドライランと空フォルダの処理オプションを提供します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file merge",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file move`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "move",
  	Title:   "ファイルを移動します",
- 	Desc:    "",
+ 	Desc:    "同じDropboxアカウント内でファイルまたはフォルダをある場所から別の場所に移動します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file move",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "ファイルリクエストを作成します",
- 	Desc:    "",
+ 	Desc:    "Dropboxアクセス権限がないユーザーがファイルをアップロードできるファイルリクエストフォルダを作成します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file request create",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request delete closed`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "closed",
  	Title:   "このアカウントの全ての閉じられているファイ\xe3"...,
- 	Desc:    "",
+ 	Desc:    "クローズされ、アップロードを受け付けなくなったファイルリクエストを削除します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file request delete closed",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request delete url`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "ファイルリクエストのURLを指定して削除",
- 	Desc:    "",
+ 	Desc:    "URLを使用して特定のファイルリクエストを削除します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file request delete url",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file request list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "個人アカウントのファイルリクエストを一覧.",
- 	Desc:    "",
+ 	Desc:    "アカウント内のすべてのファイルリクエストをステータスと詳細とともにリスト表示します。",
  	Remarks: "",
  	Path:    "dropbox file request list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file restore all`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "all",
  	Title:   "指定されたパス以下をリストアします",
- 	Desc:    "",
+ 	Desc:    "指定されたパス内のすべての削除されたファイルとフォルダを復元します。",
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "dropbox file restore all",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file restore ext`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "ext",
  	Title:   "特定の拡張子を持つファイルの復元",
- 	Desc:    "",
+ 	Desc:    "パス内の特定のファイル拡張子に一致する削除されたファイルを復元します。",
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "dropbox file restore ext",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file revision download`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "download",
  	Title:   "ファイルリビジョンをダウンロードする",
- 	Desc:    "",
+ 	Desc:    "ファイルのリビジョン履歴から特定のリビジョン/バージョンをダウンロードします。",
  	Remarks: "",
  	Path:    "dropbox file revision download",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file revision list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "ファイルリビジョン一覧",
- 	Desc:    "",
+ 	Desc:    "ファイルの利用可能なすべてのリビジョンを変更時刻とサイズとともにリスト表示します。",
  	Remarks: "",
  	Path:    "dropbox file revision list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file revision restore`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "restore",
  	Title:   "ファイルリビジョンを復元する",
- 	Desc:    "",
+ 	Desc:    "ファイルをバージョン履歴から以前のリビジョンに復元します。",
  	Remarks: "",
  	Path:    "dropbox file revision restore",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file search content`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "content",
  	Title:   "ファイルコンテンツを検索",
- 	Desc:    "",
+ 	Desc:    "ファイルの内容を検索して、指定されたキーワードやフレーズを含むファイルを見つけます。",
  	Remarks: "",
  	Path:    "dropbox file search content",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file search name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "name",
  	Title:   "ファイル名を検索",
- 	Desc:    "",
+ 	Desc:    "ファイル名とフォルダ名を検索して、指定されたパターンに一致するアイテムを見つけます。",
  	Remarks: "",
  	Path:    "dropbox file search name",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file share info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "ファイルの共有情報を取得する",
- 	Desc:    "",
+ 	Desc:    "ファイルまたはフォルダの共有情報と権限の詳細を表示します。",
  	Remarks: "",
  	Path:    "dropbox file share info",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "info",
  	Title:   "共有フォルダ情報の取得",
- 	Desc:    "",
+ 	Desc:    "特定の共有フォルダの詳細情報とメンバーを表示します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder info",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "共有フォルダの一覧",
- 	Desc:    "",
+ 	Desc:    "アクセス可能なすべての共有フォルダとその共有詳細をリスト表示します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "共有フォルダへのメンバーの追加",
- 	Desc:    "",
+ 	Desc:    "共有フォルダに新しいメンバーを追加し、指定されたアクセス権限を付与します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder member add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "共有フォルダからメンバーを削除する",
- 	Desc:    "",
+ 	Desc:    "共有フォルダからメンバーを削除し、そのアクセス権を取り消します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder member delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "共有フォルダのメンバーを一覧します",
- 	Desc:    "",
+ 	Desc:    "共有フォルダのすべてのメンバーとそのアクセスレベル、メールアドレスをリスト表示します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder member list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "共有フォルダを現在のユーザーのDropboxに追加する",
- 	Desc:    "",
+ 	Desc:    "共有フォルダをDropboxにマウントし、ファイル構造に表示されるようにします。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "現在のユーザーがマウントしているすべての共\xe6"...,
- 	Desc:    "",
+ 	Desc:    "Dropboxにマウントされているすべての共有フォルダをリスト表示します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "mountable",
  	Title:   "現在のユーザーがマウントできるすべての共有\xe3"...,
- 	Desc:    "",
+ 	Desc:    "マウント可能だが現在Dropboxにマウントされていない共有フォルダをリスト表示します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder mount mountable",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "share",
  	Title:   "フォルダの共有",
- 	Desc:    "",
+ 	Desc:    "既存のフォルダから設定可能な共有ポリシーと権限で共有フォルダを作成します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder share",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedfolder unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "unshare",
  	Title:   "フォルダの共有解除",
- 	Desc:    "",
+ 	Desc:    "フォルダの共有を停止し、オプションで現在のメンバーにコピーを残します。",
  	Remarks: "",
  	Path:    "dropbox file sharedfolder unshare",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sharedlink create`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "共有リンクの作成",
- 	Desc:    "",
+ 	Desc:    "オプションのパスワード保護と有効期限でファイルまたはフォルダの共有リンクを作成します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file sharedlink create",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "size",
  	Title:   "ストレージの利用量",
- 	Desc:    "",
+ 	Desc:    "指定した深さレベルでフォルダとその内容のサイズを計算してレポートします。",
  	Remarks: "",
  	Path:    "dropbox file size",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sync down`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "down",
  	Title:   "Dropboxと下り方向で同期します",
- 	Desc:    "",
+ 	Desc:    "フィルタリングと上書きオプションでDropboxからローカルファイルシステムにファイルをダウンロードします。",
  	Remarks: "",
  	Path:    "dropbox file sync down",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sync online`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "online",
  	Title:   "オンラインファイルを同期します",
- 	Desc:    "",
+ 	Desc:    "Dropboxオンラインストレージ内の2つの異なる場所間でファイルを同期します。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file sync online",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file sync up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "up",
  	Title:   "Dropboxと上り方向で同期します",
- 	Desc:    "",
+ 	Desc:    "フィルタリングと上書きオプションでローカルファイルシステムからDropboxにファイルをアップロードします。",
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox file sync up",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file tag add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "ファイル/フォルダーにタグを追加する",
- 	Desc:    "",
+ 	Desc:    "Dropboxのファイルまたはフォルダーにカスタムタグを追加します。タグはコンテンツの整理と分類に役立ち、検索と管理を容易にします。同じファイルまたはフォルダーに複数のタグを追加で"...,
  	Remarks: "",
  	Path:    "dropbox file tag add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file tag delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "ファイル/フォルダーからタグを削除する",
- 	Desc:    "",
+ 	Desc:    "Dropboxのファイルまたはフォルダーから特定のタグを削除します。この操作はタグの関連付けのみを削除し、ファイルまたはフォルダー自体には影響しません。古いまたは不正確なタグをクリ"...,
  	Remarks: "",
  	Path:    "dropbox file tag delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file tag list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "パスのタグを一覧",
- 	Desc:    "",
+ 	Desc:    "Dropboxの特定のファイルまたはフォルダーに関連付けられたすべてのタグを表示します。このコマンドは、コンテンツの整理と分類のために適用されたタグを確認するのに役立ちます。出力に"...,
  	Remarks: "",
  	Path:    "dropbox file tag list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file template apply`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "apply",
  	Title:   "Dropboxのパスにファイル/フォルダー構造のテン\xe3\x83"...,
- 	Desc:    "",
+ 	Desc:    "保存されたファイル/フォルダ構造テンプレートを適用してDropboxにディレクトリとファイルを作成します。",
  	Remarks: "",
  	Path:    "dropbox file template apply",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file template capture`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "capture",
  	Title:   "Dropboxのパスからファイル/フォルダ構造をテン\xe3\x83"...,
- 	Desc:    "",
+ 	Desc:    "Dropboxパスからファイル/フォルダ構造をキャプチャして再利用可能なテンプレートとして保存します。",
  	Remarks: "",
  	Path:    "dropbox file template capture",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox file watch`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "watch",
  	Title:   "ファイルアクティビティを監視",
- 	Desc:    "",
+ 	Desc:    "パスの変更を監視し、ファイル/フォルダの変更をリアルタイムで出力します。",
  	Remarks: "",
  	Path:    "dropbox file watch",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity batch user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "user",
  	Title: strings.Join({
  		"複数\xe3\x83",
- 		"\xa6ーザーのアクティビティを一括取得します",
+ 		"\x81ームメンバーのアクティビティログをバッチ取",
+ 		"得し、コンプライアンス監査やユーザー行動分\xe6",
+ 		"\x9e\x90に活用",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "ファイルからユーザーのメールアドレスリストを読み込み、指定された期間内のアクティビティログを取得します。人事調査、コンプライアンスレポート、特定のユーザーグループのパター\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team activity batch user",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity daily event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "event",
  	Title: strings.Join({
+ 		"日別",
  		"アクティビティ\xe3\x83",
- 		"\xbcを1日ごとに取得します",
+ 		"\xacポートを生成し、チーム利用パターンとセキュ",
+ 		"リティ監視に活用",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "チームのアクティビティイベントを日別に集計し、チーム行動の傾向や異常を特定しやすくします。日次セキュリティレポートの作成、新機能の採用状況の追跡、セキュリティ上の懸念を示\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team activity daily event",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity event`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "event",
- 	Title:   "イベントログ",
+ 	Title:   "詳細なチームアクティビティイベントログをフィルタリングオプション付きで取得、セキュリティ監査とコンプライアンス監視に必須",
  	Desc:    "リリース91以降では、`-start-time`または`-end-time`\xe3\x82"...,
  	Remarks: "",
  	... // 20 identical fields
  }
```
# コマンド仕様の変更: `dropbox team activity user`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "user",
  	Title: strings.Join({
- 		"ユーザーごとの",
+ 		"特定チームメンバーのアクティビティログを取\xe5",
+ 		"\xbe\x97、ファイル操作、ログイン、共有",
  		"アクティビティ",
+ 		"を表示",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "個々のチームメンバーの詳細なアクティビティログを取得します。ファイル操作、共有アクティビティ、ログインイベントを含みます。ユーザー固有の監査、セキュリティインシデントの調\xe6"...,
  	Remarks: "",
  	Path:    "dropbox team activity user",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin group role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
+ 		"指定",
  		"グループの",
+ 		"全",
  		"メンバーに",
- 		"ロールを追加する",
+ 		"管理者ロールを割り当て、大規模チームのロー\xe3",
+ 		"\x83\xab管理を効率化",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "個々のメンバーではなくグループ全体に管理者権限を効率的に付与します。部門管理者の割り当てや新しい管理チームのオンボーディング時に最適です。変更は現在のグループメンバー全員\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team admin group role add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin group role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"例外グループのメンバーを除くすべてのメンバ\xe3",
- 		"\x83\xbcからロールを削除する",
+ 		"指定した例外グループを除く全チームメンバー\xe3",
+ 		"\x81\x8bら管理者ロールを削除、ロールのクリーンア\xe3\x83",
+ 		"\x83プとアクセス制御に便利",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "特定の管理者ロールを一括削除しながら、例外グループには保持します。管理者構造の再編成や最小権限アクセスの実装に便利です。例外グループにより、クリーンアップ操作中も重要な管\xe7"...,
  	Remarks: "",
  	Path:    "dropbox team admin group role delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"メンバーの管理者権限一覧",
+ 		"割り当てられた管理者ロールを持つ全チームメ\xe3",
+ 		"\x83\xb3バーを表示、管理アクセスと権限の監査に有\xe7\x94",
+ 		"\xa8",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "昇格された権限を持つすべてのメンバーを示す包括的な管理者監査レポートを生成します。完全な可視性のために非管理者メンバーを含めることができます。セキュリティレビュー、コンプ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team admin list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
- 		"メンバーに新しいロールを追加する",
+ 		"個々のチームメンバーに特定の管理者ロールを\xe4",
+ 		"\xbb\x98与、きめ細かな権限管理を実現",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "個々のメンバーに特定の管理者ロールを割り当て、正確な権限制御を行います。チームメンバーを管理職に昇進させたり、責任を調整したりする際に使用します。コマンドは、重複を防ぐた\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team admin role add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role clear`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "clear",
  	Title: strings.Join({
  		"\xe3\x83",
- 		"\xa1ンバーからすべての管理者ロールを削除する",
+ 		"\x81ームメンバーから全管理者権限を取り消し、ロ",
+ 		"ール移行やセキュリティ目的に有用",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "メンバーからすべての管理者ロールを一度に完全に削除します。管理者のオフボーディング、セキュリティインシデントへの対応、またはメンバーの非管理職への移行に不可欠です。個別に\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team admin role clear",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
+ 		"他",
  		"\xe3",
- 		"\x83\xa1ンバーからロールを削除する",
+ 		"\x81\xaeロールを保持したままチームメンバーから特\xe5\xae",
+ 		"\x9aの管理者ロールを削除、正確な権限調整が可能",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "他の権限に影響を与えることなく、個々の管理者ロールを選択的に削除します。責任の調整やロールベースのアクセス変更の実装に便利です。コマンドは削除を試みる前にメンバーがロール\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team admin role delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team admin role list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
  		"チーム\xe3\x81",
- 		"\xae管理者の役割を列挙",
+ 		"\xa7利用可能なすべての管理者ロールとその説明・",
+ 		"権限を表示",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "Dropboxチームで利用可能なすべての管理者ロールとその機能を一覧表示します。ロールを割り当てる前に参照して、権限の影響を理解してください。チームメンバーが適切なアクセスレベルを"...,
  	Remarks: "",
  	Path:    "dropbox team admin role list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper count`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "count",
  	Title:   "メンバー1人あたりのPaper文書の枚数",
- 	Desc:    "",
+ 	Desc:    "メンバーごとのPaperドキュメント数を提供し、作成されたドキュメントとアクセスされたドキュメントを区別します。PaperからDropboxへの移行計画、ヘビーユーザーの特定、移行範囲の見積も\xe3\x82"...,
  	Remarks: "",
  	Path:    "dropbox team content legacypaper count",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper export`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "export",
  	Title:   "チームメンバー全員のPaper文書をローカルパス\xe3\x81"...,
- 	Desc:    "",
+ 	Desc:    "チームのPaperドキュメントをローカルストレージに一括エクスポートし、移行前またはコンプライアンスアーカイブのためにコンテンツを保存します。HTMLおよびMarkdown形式をサポート。メン\xe3\x83"...,
  	Remarks: "",
  	Path:    "dropbox team content legacypaper export",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content legacypaper list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームメンバーのPaper文書リスト出力",
- 	Desc:    "",
+ 	Desc:    "タイトル、所有者、最終更新日を含むすべてのPaperドキュメントの詳細なインベントリを作成します。コンテンツ監査、孤立したドキュメントの特定、または移行の準備に使用します。作成\xe3\x81"...,
  	Remarks: "",
  	Path:    "dropbox team content legacypaper list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダや共有フォルダのメンバー一覧",
- 	Desc:    "",
+ 	Desc:    "チーム全体のフォルダアクセスをマッピングし、特定のフォルダにアクセスできるメンバーとその権限レベルを表示します。アクセスレビュー、過剰な権限を持つアカウントの特定、コンテ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team content member list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content member size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "size",
  	Title:   "チームフォルダや共有フォルダのメンバー数を\xe3"...,
- 	Desc:    "",
+ 	Desc:    "フォルダメンバーシップの密度を分析して、過度に共有されているコンテンツを特定します。メンバー数が多いとセキュリティリスクやパフォーマンスの問題を示す可能性があります。権限\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team content member size",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームメンバーのマウント済み/アンマウント済"...,
- 	Desc:    "",
+ 	Desc:    "共有フォルダがメンバーのデバイスにアクティブに同期されているか、クラウドのみのアクセスかを表示します。帯域幅の計画、ヘビー同期ユーザーの特定、同期問題のトラブルシューティ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team content mount list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team content policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダと共有フォルダのポリシー一覧",
- 	Desc:    "",
+ 	Desc:    "閲覧者情報の制限、共有リンクポリシー、その他のガバナンス設定を示す包括的なポリシー監査。コンプライアンスの検証と、フォルダが組織のセキュリティ要件を満たしていることを確認\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team content policy list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team device list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
  		"\xe3\x83",
- 		"\x81ーム内全てのデバイス/セッションを一覧しま\xe3\x81",
- 		"\x99",
+ 		"\x87バイス詳細と最終アクティビティタイムスタン",
+ 		"プ付きで、チームメンバーアカウントに接続さ\xe3",
+ 		"\x82\x8cた全デバイスとアクティブセッションを表示",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "接続されたすべてのデバイス、プラットフォーム、セッション期間を示す完全なデバイスインベントリです。セキュリティ監査、未承認デバイスの特定、デバイス制限の管理に重要です。デ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team device list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team device unlink`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "unlink",
  	Title: strings.Join({
- 		"デバイスのセッションを解除します",
+ 		"紛失・盗難デバイスの保護やアクセス取り消し\xe3",
+ 		"\x81\xab必須、チームメンバーアカウントからデバイ\xe3\x82",
+ 		"\xb9をリモート切断",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "デバイスセッションを即座に終了し、再認証を強制します。紛失デバイス、退職者、疑わしい活動に対する重要なセキュリティツールです。リンク解除後はデバイスの再接続と再同期が必要\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team device unlink",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "feature",
  	Title: strings.Join({
- 		"チームの機能を出力します",
+ 		"APIリミットや特殊機能を含む、Dropboxチームアカ",
+ 		"ウントで有効なすべての機能と性能を表示",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "チームの有効な機能、ベータアクセス、APIレート制限を表示します。高度な機能を使用したり統合を計画したりする前に確認してください。機能はサブスクリプションレベルによって異なる\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team feature",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filerequest clone`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "clone",
  	Title:   "ファイルリクエストを入力データに従い複製し\xe3"...,
- 	Desc:    "",
+ 	Desc:    "既存のテンプレートに基づいて設定を変更した新しいファイルリクエストを作成します。月次レポートや定期的な提出などの標準化された収集プロセスを効率化します。受信者ごとのカスタ\xe3"...,
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "dropbox team filerequest clone",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filerequest list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームないのファイルリクエストを一覧します",
- 	Desc:    "",
+ 	Desc:    "チーム全体のすべてのファイルリクエストの包括的なビュー。外部データ収集の監視、放棄されたリクエストの特定、データ処理ポリシーへの準拠の確保。監査目的でリクエストURL、作成者\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team filerequest list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team filesystem`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "filesystem",
  	Title:   "チームのファイルシステムのバージョンを特定\xe3"...,
- 	Desc:    "",
+ 	Desc:    "機能の利用可能性とAPIの動作に影響する基礎となるファイルシステムのバージョンを決定します。最新のファイルシステムは、ネイティブPaperやパフォーマンスの向上などの高度な機能を有\xe5\x8a"...,
  	Remarks: "",
  	Path:    "dropbox team filesystem",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "グループを作成します",
- 	Desc:    "",
+ 	Desc:    "チームメンバーの論理的な編成のためのグループを作成します。グループは一括操作を可能にすることで権限管理を簡素化します。識別しやすい命名規則を検討してください。ガバナンスの\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "グループの一括追加",
- 	Desc:    "",
+ 	Desc:    "データファイルからグループを一括作成します。初期設定や組織再編に最適です。部分的な失敗を防ぐため、作成前にすべてのグループを検証します。アイデンティティ管理システムとの統\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team group batch add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "グループの削除",
- 	Desc:    "",
+ 	Desc:    "複数のグループを単一の操作で効率的に削除します。組織の再構築や古いグループのクリーンアップに便利です。メンバーは個別の権限を保持しますが、グループベースのアクセスを失いま\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group batch delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
  	Title:   "グループの外部IDをクリアする",
- 	Desc:    "",
+ 	Desc:    "アイデンティティプロバイダーから移行する場合や統合システムを変更する場合に、グループから外部IDの関連付けを削除します。グループの機能はそのまま残りますが、外部システムのマ\xe3\x83"...,
  	Remarks: "",
  	Path:    "dropbox team group clear externalid",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:  "delete",
  	Title: "グループを削除します",
  	Desc: strings.Join({
  		"\xe3",
- 		"\x81\x93のコマンドはグループがフォルダなどで利用\xe3\x81",
- 		"\x95れているかどうかを確認しない点ご注意くださ",
- 		"い",
+ 		"\x82\xb0ループを完全に削除し、すべてのメンバーの\xe9\x96",
+ 		"\xa2連付けを削除します。メンバーは他のグループ",
+ 		"または個別の権限を通じてアクセスを保持しま\xe3",
+ 		"\x81\x99。元に戻すことはできません - 不確かな場合\xe3\x81",
+ 		"\xaf、代わりにメンバーを削除してグループをアー",
+ 		"カイブすることを検討してください。このグル\xe3",
+ 		"\x83\xbcプを使用するフォルダ権限も削除されます。",
  	}, ""),
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "各グループのフォルダーを一覧表示",
- 	Desc:    "",
+ 	Desc:    "グループの権限をフォルダにマッピングし、コンテンツアクセスパターンを明らかにします。アクセスレビューと権限の継承の理解に不可欠です。過剰な権限を持つグループを特定し、セキ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team group folder list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "グループを一覧",
+ 	Title:   "メンバー数とグループ管理タイプを含むチーム内の全グループを表示",
- 	Desc:    "",
+ 	Desc:    "サイズと管理モードを示すチームグループの完全なインベントリです。空のグループ、過大なグループ、管理タイプの変更が必要なグループを特定するために使用します。定期的な監査とコ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team group list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "メンバーをグループに追加",
- 	Desc:    "",
+ 	Desc:    "継承された権限と簡素化された管理のためにメンバーをグループに追加します。変更はフォルダアクセスに即座に反映されます。非常に大きなグループの場合は、グループサイズの制限とパ\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group member add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "グループにメンバーを一括追加",
- 	Desc:    "",
+ 	Desc:    "マッピングファイルを使用してメンバーをグループに一括追加します。変更を適用する前にすべてのメンバーシップを検証します。オンボーディング、部門変更、または権限標準化プロジェ\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group member batch add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "グループからメンバーを削除",
- 	Desc:    "",
+ 	Desc:    "CSVファイルマッピングを使用してグループからメンバーを一括削除します。変更を行う前にすべてのメンバーシップを検証します。組織の再構築、オフボーディングプロセス、またはグルー\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group member batch delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "update",
  	Title:   "グループからメンバーを追加または削除",
- 	Desc:    "",
+ 	Desc:    "CSVファイルに基づいてグループメンバーシップを一括変更します。単一の操作でメンバーの追加と削除の両方が可能です。グループ構成に大幅な更新が必要な大規模な再編成に最適です。行\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group member batch update",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "メンバーをグループから削除",
- 	Desc:    "",
+ 	Desc:    "他のグループのメンバーシップに影響を与えることなく、単一のグループから個々のメンバーを削除します。対象を絞った権限調整や、メンバーが部門を変更する場合に使用します。削除は\xe5"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group member delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "グループに所属するメンバー一覧を取得します",
- 	Desc:    "",
+ 	Desc:    "すべてのグループとその完全なメンバー名簿を一覧表示します。アクセス監査、グループ構成の確認、権限継承の理解に不可欠です。空のグループ、過剰な権限を持つグループ、またはグル\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team group member list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group rename`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "rename",
  	Title:   "グループの改名",
- 	Desc:    "",
+ 	Desc:    "すべてのメンバーと権限を維持しながら、グループの表示名を更新します。部門が再構築されたり、プロジェクト名が変更されたり、グループの目的が進化したりする場合に便利です。名前\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team group rename",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team group update type`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "type",
  	Title:   "グループ管理タイプの更新",
- 	Desc:    "",
+ 	Desc:    "メンバーの追加や削除を誰が行えるかを制御するためにグループ管理設定を変更します。会社管理グループは変更を管理者に制限し、ユーザー管理グループは指定されたメンバーがメンバー\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team group update type",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team info`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "info",
- 	Title:   "チームの情報",
+ 	Title:   "チームIDと基本チーム設定を含む必須チームアカウント情報を表示",
- 	Desc:    "",
+ 	Desc:    "API統合とサポート要求に必要な基本的なチームアカウントの詳細を表示します。チームIDは様々な管理操作に必要です。正しいチームアカウントに接続していることを確認する簡単な方法で\xe3\x81"...,
  	Remarks: "",
  	Path:    "dropbox team info",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "Idによるリーガルホールドを解除する",
- 	Desc:    "",
+ 	Desc:    "リーガルホールドポリシーを終了し、保存要件を削除します。コンテンツは再び通常の保持および削除ポリシーの対象となります。訴訟が終結した場合や保存が不要になった場合に使用しま\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team legalhold release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold revision list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "リーガル・ホールド・ポリシーのリビジョンを\xe3"...,
- 	Desc:    "",
+ 	Desc:    "すべての変更を含む、リーガルホールド下のファイルの完全なリビジョン履歴を表示します。ポリシーによって保存されたファイルバージョンを追跡し、何も失われないようにします。防御\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team legalhold revision list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team legalhold update desc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "desc",
  	Title:   "リーガルホールド・ポリシーの説明を更新",
- 	Desc:    "",
+ 	Desc:    "より良い文書化のためにリーガルホールドポリシーの説明フィールドを更新します。ケース参照の追加、案件詳細の更新、または保存範囲の明確化に便利です。変更は監査目的でリビジョン\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team legalhold update desc",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team linkedapp list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "リンク済みアプリを一覧",
- 	Desc:    "",
+ 	Desc:    "チームメンバーのDropboxアカウントにアクセスできるすべてのサードパーティアプリケーションを一覧表示します。セキュリティ監査、不正なアプリの特定、OAuth統合の管理に不可欠です。ど\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team linkedapp list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "メンバーを削除します",
- 	Desc:    "",
+ 	Desc:    "転送を通じてデータを保持しながら、チームメンバーを一括削除します。ファイル転送の宛先メンバーと管理者通知メールの指定が必要です。レイオフ、部門閉鎖、または大量のオフボーデ\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member batch delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "detach",
  	Title:   "Dropbox for teamsのアカウントをBasicアカウントに変"...,
- 	Desc:    "",
+ 	Desc:    "チームメンバーを個人のDropbox Basicアカウントに一括変換します。メンバーはファイルを保持しますが、チーム機能と共有フォルダへのアクセスを失います。契約終了の請負業者やチームの縮"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member batch detach",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch invite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "invite",
  	Title:   "メンバーを招待します",
- 	Desc:    "",
+ 	Desc:    "CSVファイルから複数のメールアドレスにチーム招待を送信します。SSO環境用のサイレント招待をサポート。新しい部門のオンボーディング、買収、または季節労働者に最適です。送信前にメ\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member batch invite",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch reinvite`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "reinvite",
  	Title:   "招待済み状態メンバーをチームに再招待します",
- 	Desc:    "",
+ 	Desc:    "保留中のステータスを持つすべてのメンバーに招待を再送信します。初回の招待が期限切れになったり、スパムで失われたり、メール配信の問題を解決した後に便利です。SSO環境では無音で\xe9"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member batch reinvite",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "suspend",
  	Title:   "メンバーの一括一時停止",
- 	Desc:    "",
+ 	Desc:    "チームメンバーを一括停止し、すべてのデータと設定を保持しながらアクセスをブロックします。長期休暇、セキュリティ調査、または一時的なアクセス制限に使用します。デバイスからデ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member batch suspend",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member batch unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "unsuspend",
  	Title:   "メンバーの一括停止解除",
- 	Desc:    "",
+ 	Desc:    "停止されたチームメンバーを一括再アクティブ化し、アカウントとデータへの完全なアクセスを復元します。メンバーが休暇から戻ったり、調査が終了したり、アクセス制限が解除されたり\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member batch unsuspend",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member clear externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
  	Title:   "メンバーのexternal_idを初期化します",
- 	Desc:    "",
+ 	Desc:    "CSVファイルにリストされたチームメンバーから外部IDを一括削除します。アイデンティティプロバイダー間の移行、SCIM切断後のクリーンアップ、またはIDの競合の解決に不可欠です。メンバ\xe3\x83"...,
  	Remarks: "",
  	Path:    "dropbox team member clear externalid",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member feature`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "feature",
  	Title:   "メンバーの機能設定一覧",
- 	Desc:    "",
+ 	Desc:    "チームメンバーに対して有効になっている機能と機能を表示します。アクセスの問題のトラブルシューティング、機能のロールアウトの確認、メンバーの機能の理解に便利です。特定のメン\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member feature",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "メンバーのパスの下にあるすべてのロックを解\xe9"...,
- 	Desc:    "",
+ 	Desc:    "指定されたフォルダパス内でメンバーが保持しているすべてのファイルロックを一括解除します。メンバーが予期せず退社した場合やシステムの問題が発生した場合に不可欠です。効率性の\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member file lock all release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "パスの下にあるメンバーのロックを一覧表示",
- 	Desc:    "",
+ 	Desc:    "パス内で特定のメンバーが現在ロックしているすべてのファイルを一覧表示します。コラボレーションのボトルネックの特定、編集の競合のトラブルシューティング、ファイルアクセスパタ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member file lock list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "メンバーとしてパスのロックを解除します",
- 	Desc:    "",
+ 	Desc:    "メンバーが保持している単一のファイルロックを解除し、他のユーザーが編集できるようにします。特定のファイルがチームのコラボレーションをブロックしている場合や、ロック保持者が\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team member file lock release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member folder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "各メンバーのフォルダーを一覧表示",
- 	Desc:    "",
+ 	Desc:    "チームメンバーの個人スペース全体のフォルダを列挙します。フォルダ名でフィルタリングして結果に焦点を当てます。コンテンツの配布の理解、メンバーストレージの監査、移行またはク\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member folder list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "チームメンバーの一覧",
+ 	Title:   "ステータス、ロール、アカウント詳細を含む全チームメンバーの包括的リストを表示",
- 	Desc:    "",
+ 	Desc:    "チーム全体の包括的なメンバーリストを表示し、ステータス、ロール、アカウントの詳細を含みます。メンバー監査、組織構造の理解、アクセス管理の計画に不可欠です。削除されたメンバ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota batch update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "update",
  	Title:   "チームメンバーの容量制限を変更",
- 	Desc:    "",
+ 	Desc:    "CSVファイルを使用してチームメンバーのストレージクォータを一括更新します。ロール、部門、または使用パターンに基づいてカスタムクォータを設定します。0を使用してカスタムクォータ"...,
  	Remarks: "",
  	Path:    "dropbox team member quota batch update",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "メンバーの容量制限情報を一覧します",
- 	Desc:    "",
+ 	Desc:    "すべてのチームメンバーの現在のストレージクォータ設定を表示し、デフォルトとカスタムクォータを区別します。特別なストレージニーズや制限を持つメンバーを特定します。容量計画と\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member quota list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member quota usage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "usage",
  	Title:   "チームメンバーのストレージ利用状況を取得",
- 	Desc:    "",
+ 	Desc:    "各メンバーの現在のストレージ消費量と割り当てられたクォータを表示します。制限に近づいているメンバー、スペースを十分に活用していないメンバー、またはクォータ調整が必要なメン\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member quota usage",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "チームメンバーのファイルを複製します",
- 	Desc:    "",
+ 	Desc:    "アカウント間でメンバーデータの完全なコピーを作成し、可能な限りフォルダ構造と共有を保持します。ロールの移行、バックアップの作成、またはアカウントの統合に不可欠です。宛先ア\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member replication",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member suspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "suspend",
  	Title:   "メンバーの一時停止処理",
- 	Desc:    "",
+ 	Desc:    "すべてのデータ、設定、グループメンバーシップを維持しながら、メンバーアクセスを即座にブロックします。セキュリティインシデント、ポリシー違反、または一時的な休暇に使用します\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member suspend",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member unsuspend`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "unsuspend",
  	Title:   "メンバーの一時停止を解除する",
- 	Desc:    "",
+ 	Desc:    "停止されたメンバーのアカウントを再アクティブ化し、データとチームリソースへの完全なアクセスを復元します。以前のすべての権限、グループメンバーシップ、および設定が保持されま\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team member unsuspend",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch email`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "email",
  	Title:   "メンバーのメールアドレス処理",
- 	Desc:    "",
+ 	Desc:    "CSVマッピングファイルを使用してメンバーのメールアドレスを一括更新します。ドメイン移行、名前変更、またはメールエラーの修正に不可欠です。新しいアドレスを検証し、すべてのメン\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member update batch email",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch externalid`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
  	Title:   "チームメンバーのExternal IDを更新します.",
- 	Desc:    "",
+ 	Desc:    "外部アイデンティティシステムIDをDropboxチームメンバーに一括マッピングします。SCIM統合、SSOセットアップ、またはHRシステムとの同期に重要です。プラットフォーム間で一貫したアイデン\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member update batch externalid",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch invisible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "invisible",
  	Title:   "メンバーへのディレクトリ制限を有効にします",
- 	Desc:    "",
+ 	Desc:    "チームディレクトリの検索とリストからメンバーを一括で非表示にします。アクセスは必要だがディレクトリに表示されるべきではない役員、セキュリティ担当者、または外部請負業者に便\xe5"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member update batch invisible",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch profile`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "profile",
  	Title:   "メンバーのプロフィール変更",
- 	Desc:    "",
+ 	Desc:    "名前と姓を含むメンバープロファイル情報を一括更新します。名前形式の標準化、広範なエラーの修正、または組織変更後の更新に最適です。チームディレクトリ全体で一貫性を維持し、検\xe7"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member update batch profile",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team member update batch visible`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "visible",
  	Title:   "メンバーへのディレクトリ制限を無効にします",
- 	Desc:    "",
+ 	Desc:    "チームディレクトリで以前非表示だったメンバーの可視性を一括復元します。プライバシー要件が変更されたり、請負業者が従業員になったり、可視性エラーを修正する場合に使用します。\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team member update batch visible",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チーム内全ての名前空間でのファイル・フォル\xe3"...,
- 	Desc:    "",
+ 	Desc:    "フィルタリングオプションを使用してチームネームスペース内のすべてのファイルとフォルダを一覧表示します。削除されたアイテム、メンバーフォルダ、共有フォルダ、チームフォルダを\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team namespace file list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "size",
  	Title:   "チーム内全ての名前空間でのファイル・フォル\xe3"...,
- 	Desc:    "",
+ 	Desc:    "設定可能な深度スキャンでチームネームスペース全体のストレージ消費を分析します。ネームスペースタイプ（チーム、共有、メンバー、アプリフォルダ）別のサイズ分布を表示します。ス\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team namespace file size",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チーム内すべての名前空間を一覧",
- 	Desc:    "",
+ 	Desc:    "所有権、パス、アクセスレベルを含むチーム内のすべてのネームスペースタイプを列挙します。チームのフォルダアーキテクチャの包括的なビューを提供します。組織構造の理解、移行の計\xe7"...,
  	Remarks: "",
  	Path:    "dropbox team namespace list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダ以下のファイル・フォルダを一覧",
- 	Desc:    "",
+ 	Desc:    "どのメンバーがどのフォルダにアクセスでき、その権限レベルを示すネームスペースアクセスをマッピングします。アクセスパターン、過剰な権限を持つネームスペースを明らかにし、適切\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team namespace member list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team namespace summary`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "summary",
  	Title:   "チーム・ネームスペースの状態概要を報告する.",
- 	Desc:    "",
+ 	Desc:    "ネームスペースデータを集約して、全体的なチーム構造、ストレージ分布、アクセスパターンを表示します。チームコンテンツがさまざまなネームスペースタイプ間でどのように編成されて\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team namespace summary",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report activity`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "activity",
  	Title: strings.Join({
- 		"アクティビティ レポート",
+ 		"全チーム操作をカバーする詳細アクティビティ\xe3",
+ 		"\x83\xacポートを生成、コンプライアンスと使用分析\xe3\x81",
+ 		"\xab有用",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "すべてのチーム操作をカバーする詳細なアクティビティレポートを生成します。コンプライアンスと使用分析に有用です。期間、ユーザー、アクティビティタイプでフィルタリングできます\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team report activity",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report devices`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "devices",
  	Title:   "デバイス レポート空のレポート",
- 	Desc:    "",
+ 	Desc:    "タイプ、OS、同期ステータス、最後のアクティビティを含む、チームアカウントに接続されているすべてのデバイスを表示します。セキュリティ監査、不正なデバイスの特定、デバイスポリ\xe3\x82"...,
  	Remarks: "",
  	Path:    "dropbox team report devices",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report membership`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "membership",
  	Title:   "メンバーシップ レポート",
- 	Desc:    "",
+ 	Desc:    "アクティブユーザー、成長傾向、ロール分布を含むメンバーシップ分析を提供します。チームの拡大を追跡し、ライセンス使用量を監視し、非アクティブなアカウントを特定します。予算計\xe7"...,
  	Remarks: "",
  	Path:    "dropbox team report membership",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team report storage`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "storage",
- 	Title:   "ストレージ レポート",
+ 	Title:   "チーム消費、トレンド、メンバー分布を示す詳細ストレージ使用レポートを作成",
- 	Desc:    "",
+ 	Desc:    "チームのストレージ消費、使用傾向、メンバー別分布を示す詳細なストレージ使用レポートを作成します。容量計画、コスト管理、使用量の最適化に重要です。ストレージクォータの調整や\xe8"...,
  	Remarks: "",
  	Path:    "dropbox team report storage",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file batch copy`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "copy",
  	Title:   "ファイル/フォルダーをメンバーとして一括コピー",
- 	Desc:    "",
+ 	Desc:    "メンバーの資格情報なしでメンバーアカウント間でファイルをコピーする管理者ツール。テンプレートの配布、削除されたコンテンツの回復、または新しいメンバーの設定に便利です。監査\xe8"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team runas file batch copy",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "メンバーとして実行するファイルやフォルダー\xe3"...,
- 	Desc:    "",
+ 	Desc:    "管理者がメンバーの資格情報なしでメンバーアカウントのファイル一覧を表示できるようにします。問題の調査、コンテンツの監査、またはメンバーがファイルを見つけるのを支援するため\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas file list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas file sync batch up`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "up",
  	Title:   "メンバーとして動作する一括同期",
- 	Desc:    "",
+ 	Desc:    "複数のメンバーアカウントに同時にファイルを配布するための管理者一括アップロードツール。テンプレート、ポリシー、または必要なドキュメントの展開に最適です。チーム全体で一貫し\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team runas file sync batch up",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch leave`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "leave",
  	Title:   "共有フォルダからメンバーとして一括退出",
- 	Desc:    "",
+ 	Desc:    "メンバーの操作なしに複数の共有フォルダからメンバーを削除する管理者ツール。アクセスのクリーンアップ、セキュリティ対応、または組織変更に便利です。適切な監査証跡を維持しなが\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder batch leave",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch share`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "share",
  	Title:   "メンバーのフォルダを一括で共有",
- 	Desc:    "",
+ 	Desc:    "メンバーに代わって共有フォルダを作成する管理者バッチツール。新しいプロジェクトやチーム再編成のフォルダ共有を効率化します。適切な権限を設定し、招待を送信します。すべての共\xe6"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder batch share",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder batch unshare`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "unshare",
  	Title:   "メンバーのフォルダの共有を一括解除",
- 	Desc:    "",
+ 	Desc:    "セキュリティまたはコンプライアンスのためにフォルダ共有を一括で取り消す管理者ツール。所有者のフォルダ内容を保持しながら共有を削除します。インシデント対応やデータ漏洩の防止\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder batch unshare",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder isolate`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "isolate",
  	Title:   "所有する共有フォルダの共有を解除し、メンバ\xe3"...,
- 	Desc:    "",
+ 	Desc:    "所有者を除く共有フォルダからすべてのメンバーを削除する緊急管理者アクション。セキュリティインシデント、データ侵害、またはフォルダコンテンツに即座のアクセス制限が必要な場合\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team runas sharedfolder isolate",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "共有フォルダーの一覧をメンバーとして実行",
- 	Desc:    "",
+ 	Desc:    "権限レベルとフォルダの詳細を含むメンバーの共有フォルダアクセスの管理者ビュー。アクセス監査、過剰共有の調査、または権限の問題のトラブルシューティングに不可欠です。適切なア\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder member batch add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "メンバーの共有フォルダにメンバーを一括追加",
- 	Desc:    "",
+ 	Desc:    "定義された権限で特定の共有フォルダにメンバーを一括追加する管理者ツール。プロジェクト開始、チーム拡張、またはアクセス標準化に効率的です。変更を適用する前にメンバーのメール\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder member batch add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder member batch delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "メンバーの共有フォルダからメンバーを一括削除",
- 	Desc:    "",
+ 	Desc:    "セキュリティまたは再編成のために共有フォルダからメンバーを管理者が一括削除。指定されたメンバーのアクセスを取り消しながらフォルダコンテンツを保持します。迅速なセキュリティ\xe5"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder member batch delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "指定したメンバーのDropboxに共有フォルダを追加"...,
- 	Desc:    "",
+ 	Desc:    "メンバーが自分でできない場合に、メンバーアカウントに共有フォルダをマウントする管理者アクション。同期の問題のトラブルシューティング、技術的でないユーザーの支援、または重要\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount delete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "指定されたユーザーが指定されたフォルダーを\xe3"...,
- 	Desc:    "",
+ 	Desc:    "アクセスを削除せずにメンバーアカウントから共有フォルダをマウント解除する管理者ツール。同期の問題のトラブルシューティング、ローカルストレージの管理、または同期からフォルダ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount delete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "指定されたメンバーがマウントしているすべて\xe3"...,
- 	Desc:    "",
+ 	Desc:    "メンバーのアカウントでアクティブにマウント（同期）されている共有フォルダの管理者ビュー。同期の問題の診断、ストレージ使用量の理解、または適切なフォルダアクセスの確認に役立\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team runas sharedfolder mount mountable`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "mountable",
  	Title:   "メンバーがマウントできるすべての共有フォル\xe3"...,
- 	Desc:    "",
+ 	Desc:    "メンバーがアクセスできるが、現在デバイスに同期されていない共有フォルダを一覧表示します。利用可能なフォルダの特定、メンバーがコンテンツを見つけるのを支援、または特定のフォ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team runas sharedfolder mount mountable",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink cap expiry`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "expiry",
  	Title:   "チーム内の共有リンクに有効期限の上限を設定",
- 	Desc:    "",
+ 	Desc:    "有効期限のない既存の共有リンクに有効期限を適用します。セキュリティコンプライアンスと永続的なリンクの露出削減に不可欠です。経過時間でリンクをターゲットにしたり、一括有効期\xe9"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team sharedlink cap expiry",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink cap visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "visibility",
  	Title:   "チーム内の共有リンクに可視性の上限を設定",
- 	Desc:    "",
+ 	Desc:    "チームセキュリティポリシーを実施するために共有リンクの可視性設定を変更します。パブリックリンクをチームのみまたはパスワード保護されたアクセスに制限できます。データ漏洩を防\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team sharedlink cap visibility",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink delete links`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "links",
  	Title:   "共有リンクの一括削除",
- 	Desc:    "",
+ 	Desc:    "経過時間、可視性、またはパスパターンなどの基準に基づいて共有リンクを一括削除します。セキュリティの修復、古いリンクの削除、または新しい共有ポリシーの実施に使用します。削除\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team sharedlink delete links",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink delete member`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "member",
  	Title:   "メンバーの共有リンクをすべて削除",
- 	Desc:    "",
+ 	Desc:    "コンテンツの場所に関係なく、特定のメンバーが作成したすべての共有リンクを削除します。安全なオフボーディング、侵害されたアカウントへの対応、または即座のアクセス取り消しの実\xe6"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team sharedlink delete member",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "共有リンクの一覧",
- 	Desc:    "",
+ 	Desc:    "URL、可視性設定、有効期限、作成者を示すすべてのチーム共有リンクの包括的なインベントリ。セキュリティ監査、リスクのあるリンクの特定、外部共有パターンの理解に不可欠です。焦点\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team sharedlink list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update password`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "password",
  	Title:   "共有リンクのパスワードの設定・更新",
- 	Desc:    "",
+ 	Desc:    "既存の共有リンクにパスワード保護を適用するか、現在のパスワードを更新します。外部で共有される機密コンテンツを保護するために重要です。脆弱なリンクをターゲットにしたり、コン\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team sharedlink update password",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team sharedlink update visibility`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "visibility",
  	Title:   "共有リンクの可視性の更新",
- 	Desc:    "",
+ 	Desc:    "共有リンクの可視性をパブリックからチームのみまたは他の制限された設定に更新します。外部への露出を減らし、コンプライアンス要件を満たすために不可欠です。現在の可視性レベルま\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team sharedlink update visibility",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder add`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
- 		"チームフォルダを追加します",
+ 		"集約されたチームコンテンツストレージとコラ\xe3",
+ 		"\x83\x9cレーション用の新しいチームフォルダを作成",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "集中化されたチームコンテンツストレージとコラボレーション用の新しいチームフォルダを作成します。プロジェクトベースの作業、部門別フォルダ、共有リソースに最適です。作成時に同\xe6"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder add",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "チームフォルダのアーカイブ",
- 	Desc:    "",
+ 	Desc:    "アクティブなチームフォルダをアーカイブステータスに変換し、すべてのコンテンツと権限を保持しながら読み取り専用にします。完了したプロジェクト、履歴記録、またはコンプライアン\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder archive",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder batch archive`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "複数のチームフォルダをアーカイブします",
- 	Desc:    "",
+ 	Desc:    "経過時間、名前パターン、またはアクティビティレベルなどの基準に基づいてチームフォルダを一括アーカイブします。フォルダライフサイクル管理を効率化し、整理されたチームスペース\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder batch archive",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder batch permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "permdelete",
  	Title:   "複数のチームフォルダを完全に削除します",
- 	Desc:    "",
+ 	Desc:    "複数のチームフォルダとそのすべてのコンテンツを回復の可能性なしに永久に削除します。古いデータの削除、保持ポリシーへの準拠、または緊急クリーンアップのために適切な承認を得て\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder batch permdelete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder batch replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "チームフォルダの一括レプリケーション",
- 	Desc:    "",
+ 	Desc:    "複数のチームフォルダを完全な構造と権限とともにコピーを作成します。バックアップの作成、並列環境の設定、または移行の準備に便利です。大規模なレプリケーションの前にストレージ\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder batch replication",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダの一覧",
- 	Desc:    "",
+ 	Desc:    "サイズ、変更日、パスなどの詳細を含むチームフォルダ内のすべてのファイルを列挙します。コンテンツ監査、移行計画、データ分布の理解に不可欠です。対象を絞った分析のためにファイ\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder file list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock all release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "チームフォルダのパスの下にあるすべてのロッ\xe3"...,
- 	Desc:    "",
+ 	Desc:    "指定されたチームフォルダ内のすべてのファイルロックを一括解除します。複数のロックがチームの生産性をブロックしている場合やシステムの問題の後に使用します。可能な場合はロック\xe4"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder file lock all release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダ内のロックを一覧表示",
- 	Desc:    "",
+ 	Desc:    "ロック保持者情報とロック期間を含むチームフォルダ内の現在ロックされているすべてのファイルを一覧表示します。コラボレーションのボトルネック、古いロック、支援が必要な可能性の\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder file lock list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file lock release`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "release",
  	Title:   "チームフォルダ内のパスのロックを解除",
- 	Desc:    "",
+ 	Desc:    "特定のファイルが作業をブロックしている場合に、チームフォルダ内の個々のファイルロックを解除します。特定のファイルのみのロック解除が必要な場合、一括解除よりも精密です。他の\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder file lock release",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder file size`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "size",
  	Title:   "チームフォルダのサイズを計算",
- 	Desc:    "",
+ 	Desc:    "チームフォルダ内のストレージ消費を分析し、サイズ分布と最大のファイルを表示します。容量計画、アーカイブ候補の特定、ストレージコストの理解に不可欠です。チームフォルダの使用\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder file size",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
- 	Title:   "チームフォルダの一覧",
+ 	Title:   "ステータス、同期設定、メンバーアクセス情報を含む全チームフォルダを表示",
- 	Desc:    "",
+ 	Desc:    "ステータス、同期設定、メンバーアクセス情報を含むすべてのチームフォルダを表示します。フォルダ管理、アクセスレビュー、組織構造の理解に不可欠です。アーカイブされたフォルダや\xe9"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder member list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダのメンバー一覧",
- 	Desc:    "",
+ 	Desc:    "権限レベルとアクセスが直接かグループ経由かを含む、すべてのチームフォルダの完全なメンバーシップを表示します。アクセス監査、セキュリティレビュー、機密コンテンツにアクセスで\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder member list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder partial replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "部分的なチームフォルダの他チームへのレプリ\xe3"...,
- 	Desc:    "",
+ 	Desc:    "チームフォルダから全体の構造ではなく選択したサブフォルダまたはファイルをコピーします。ターゲットバックアップの作成、プロジェクト成果物の抽出、または特定のコンテンツの移行\xe3"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder partial replication",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder permdelete`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "permdelete",
  	Title:   "チームフォルダを完全に削除します",
- 	Desc:    "",
+ 	Desc:    "チームフォルダとすべての含まれるファイルを不可逆的に削除します。重要なデータが残っていないことを確認した後、適切な承認を得てのみ使用してください。データ保持ポリシーへの準\xe6"...,
  	Remarks: "(非可逆な操作です)",
  	Path:    "dropbox team teamfolder permdelete",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder policy list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダのポリシー一覧",
- 	Desc:    "",
+ 	Desc:    "同期のデフォルト、共有制限、アクセス制御を含むチームフォルダの動作を管理するすべてのポリシーを表示します。フォルダが特定の方法で動作する理由を理解し、ポリシーへの準拠を確\xe4"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder policy list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder replication`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "チームフォルダを他のチームに複製します",
- 	Desc:    "",
+ 	Desc:    "構造、権限、コンテンツを保持してチームフォルダの正確な複製を作成します。バックアップの作成、テスト環境の設定、または大きな変更の準備に使用します。大きなフォルダの場合は利\xe7"...,
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "dropbox team teamfolder replication",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder sync setting list`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "list",
  	Title:   "チームフォルダーの同期設定を一覧表示",
- 	Desc:    "",
+ 	Desc:    "すべてのチームフォルダの現在の同期設定を表示し、新しいメンバーのデバイスに自動的に同期するかどうかを示します。帯域幅への影響、ストレージ要件を理解し、適切なコンテンツ配布\xe3"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder sync setting list",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `dropbox team teamfolder sync setting update`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "update",
  	Title:   "チームフォルダ同期設定の一括更新",
- 	Desc:    "",
+ 	Desc:    "チームフォルダの同期動作をすべてのメンバーへの自動同期または手動同期選択の間で変更します。デバイスのストレージ使用量を削減したり、帯域幅を管理したり、重要なフォルダが自動\xe7"...,
  	Remarks: "",
  	Path:    "dropbox team teamfolder sync setting update",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `license`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "license",
  	Title:   "ライセンス情報を表示します",
- 	Desc:    "",
+ 	Desc:    "watermint toolboxとその全コンポーネントの詳細なライセンス情報を表示します。これにはオープンソースライセンス、著作権表示、およびアプリケーションで使用されているサードパーティ依\xe5\xad"...,
  	Remarks: "",
  	Path:    "license",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `log api job`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "job",
  	Title:   "ジョブIDで指定されたジョブのAPIログの統計情\xe5\xa0"...,
- 	Desc:    "",
+ 	Desc:    "特定のジョブ実行のAPI呼び出し統計を分析し表示します。リクエスト数、レスポンス時間、エラー率、エンドポイント使用パターンが含まれます。パフォーマンス分析、APIの問題のデバッグ\xe3"...,
  	Remarks: "",
  	Path:    "log api job",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `log api name`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "name",
  	Title:   "ジョブ名で指定されたジョブのAPIログの統計情\xe5"...,
- 	Desc:    "",
+ 	Desc:    "ジョブIDではなくコマンド名で識別されるジョブのAPI呼び出し統計を分析し表示します。同じコマンドの複数実行にわたって統計を集約し、時間の経過とともにパターンやパフォーマンスト\xe3\x83"...,
  	Remarks: "",
  	Path:    "log api name",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `log cat curl`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "curl",
  	Title:   "キャプチャログを `curl` サンプルとしてフォー\xe3\x83"...,
- 	Desc:    "",
+ 	Desc:    "APIリクエストログを独立して実行可能な同等のcurlコマンドに変換します。APIの問題のデバッグ、toolbox外でのリクエストの再現、サポートとの例の共有、テストスクリプトの作成に非常に役\xe7\xab"...,
  	Remarks: "",
  	Path:    "log cat curl",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `log cat job`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "job",
  	Title:   "指定したジョブIDのログを取得する",
- 	Desc:    "",
+ 	Desc:    "ジョブIDで識別される特定のジョブ実行のログファイルを抽出して表示します。デバッグログ、APIキャプチャログ、エラーメッセージ、システム情報が含まれます。失敗した実行のトラブル\xe3\x82"...,
  	Remarks: "",
  	Path:    "log cat job",
  	... // 19 identical fields
  }
```
# コマンド仕様の変更: `version`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "version",
  	Title:   "バージョン情報",
- 	Desc:    "",
+ 	Desc:    "ビルド日時、Gitコミットハッシュ、コンポーネントバージョンを含むwatermint toolboxのバージョン情報を表示します。トラブルシューティング、バグレポート、最新バージョンの確認に便利で\xe3\x81"...,
  	Remarks: "",
  	Path:    "version",
  	... // 19 identical fields
  }
```
