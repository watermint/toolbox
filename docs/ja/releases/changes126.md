---
layout: release
title: リリースの変更点 125
lang: ja
---

# `リリース 125` から `リリース 126` までの変更点

# 追加されたコマンド


| コマンド                           | タイトル                                              |
|------------------------------------|-------------------------------------------------------|
| asana team list                    | チームのリスト                                        |
| asana team project list            | チームのプロジェクト一覧                              |
| asana team task list               | チームのタスク一覧                                    |
| asana workspace list               | ワークスペースの一覧                                  |
| asana workspace project list       | ワークスペースのプロジェクト一覧                      |
| config feature disable             | 機能を無効化します.                                   |
| config feature enable              | 機能を有効化します.                                   |
| config feature list                | 利用可能なオプション機能一覧.                         |
| deepl translate text               | テキストを翻訳する                                    |
| dev info                           | 開発情報                                              |
| dev lifecycle planchangepath       | コマンドにパスを変更するプランを追加                  |
| dev lifecycle planprune            | コマンド廃止計画を追加                                |
| dev placeholder pathchange         | パス変更文書生成のためのプレースホルダー・コマンド    |
| dev placeholder prune              | 剪定ワークフローメッセージのプレースホルダ            |
| dropbox file account feature       | Dropboxアカウントの機能一覧                           |
| dropbox file account filesystem    | Dropboxのファイルシステムのバージョンを表示する       |
| dropbox file account info          | Dropboxアカウント情報                                 |
| dropbox sign account info          | Dropbox Signのアカウント情報を表示する                |
| figma account info                 | 現在のユーザー情報を取得する                          |
| figma file export all page         | チーム配下のすべてのファイル/ページをエクスポートする |
| figma file export frame            | Figmaファイルの全フレームを書き出す                   |
| figma file export node             | Figmaドキュメント・ノードの書き出し                   |
| figma file export page             | Figmaファイルの全ページを書き出す                     |
| figma file info                    | figmaファイルの情報を表示する                         |
| figma file list                    | Figmaプロジェクト内のファイル一覧                     |
| figma project list                 | チームのプロジェクト一覧                              |
| github content get                 | レポジトリのコンテンツメタデータを取得します.         |
| github content put                 | レポジトリに小さなテキストコンテンツを格納します      |
| github issue list                  | 公開・プライベートGitHubレポジトリの課題一覧          |
| github profile                     | 認証したユーザーの情報を取得                          |
| github release asset download      | アセットをダウンロードします                          |
| github release asset list          | GitHubリリースの成果物一覧                            |
| github release asset upload        | GitHub リリースへ成果物をアップロードします           |
| github release draft               | リリースの下書きを作成                                |
| github release list                | リリースの一覧                                        |
| github tag create                  | レポジトリにタグを作成します                          |
| google calendar event list         | Googleカレンダーのイベントを一覧表示                  |
| google mail filter add             | フィルターを追加します.                               |
| google mail filter batch add       | クエリによるラベルの一括追加・削除                    |
| google mail filter delete          | フィルタの削除                                        |
| google mail filter list            | フィルターの一覧                                      |
| google mail label add              | ラベルの追加                                          |
| google mail label delete           | ラベルの削除.                                         |
| google mail label list             | ラベルのリスト                                        |
| google mail label rename           | ラベルの名前を変更する                                |
| google mail message label add      | メッセージにラベルを追加                              |
| google mail message label delete   | メッセージからラベルを削除する                        |
| google mail message list           | メッセージの一覧                                      |
| google mail message processed list | 処理された形式でメッセージを一覧表示します.           |
| google mail message send           | メールの送信                                          |
| google mail sendas add             | カスタムの "from" send-asエイリアスの作成             |
| google mail sendas delete          | 指定したsend-asエイリアスを削除する                   |
| google mail sendas list            | 指定されたアカウントの送信エイリアスを一覧表示する    |
| google mail thread list            | スレッド一覧                                          |
| google sheets sheet append         | スプレッドシートにデータを追加する                    |
| google sheets sheet clear          | スプレッドシートから値をクリアする                    |
| google sheets sheet create         | 新規シートの作成                                      |
| google sheets sheet delete         | スプレッドシートからシートを削除する                  |
| google sheets sheet export         | シートデータのエクスポート                            |
| google sheets sheet import         | スプレッドシートにデータをインポート                  |
| google sheets sheet list           | スプレッドシートのシート一覧                          |
| google sheets spreadsheet create   | 新しいスプレッドシートの作成                          |
| google translate text              | テキストを翻訳する                                    |
| log cat job                        | 指定したジョブIDのログを取得する                      |
| log cat kind                       | 指定種別のログを結合して出力します                    |
| log cat last                       | 最後のジョブのログファイルを出力.                     |
| log job archive                    | ジョブのアーカイブ                                    |
| log job delete                     | 古いジョブ履歴の削除                                  |
| log job list                       | ジョブ履歴の表示                                      |
| log job ship                       | ログの転送先Dropboxパス                               |
| slack conversation history         | 会話履歴                                              |
| slack conversation list            | チャネルの一覧                                        |



# 削除されたコマンド


| コマンド                                    | タイトル                                                 |
|---------------------------------------------|----------------------------------------------------------|
| config disable                              | 機能を無効化します.                                      |
| config enable                               | 機能を有効化します.                                      |
| config features                             | 利用可能なオプション機能一覧.                            |
| job history archive                         | ジョブのアーカイブ                                       |
| job history delete                          | 古いジョブ履歴の削除                                     |
| job history list                            | ジョブ履歴の表示                                         |
| job history ship                            | ログの転送先Dropboxパス                                  |
| job log jobid                               | 指定したジョブIDのログを取得する                         |
| job log kind                                | 指定種別のログを結合して出力します                       |
| job log last                                | 最後のジョブのログファイルを出力.                        |
| services asana team list                    | チームのリスト                                           |
| services asana team project list            | チームのプロジェクト一覧                                 |
| services asana team task list               | チームのタスク一覧                                       |
| services asana workspace list               | ワークスペースの一覧                                     |
| services asana workspace project list       | ワークスペースのプロジェクト一覧                         |
| services deepl translate text               | テキストを翻訳する                                       |
| services dropbox user feature               | 現在のユーザーの機能設定の一覧                           |
| services dropbox user filesystem            | ユーザーのチームのファイルシステムのバージョンを特定する |
| services dropbox user info                  | 現在のアカウント情報を取得する                           |
| services dropboxsign account info           | アカウント情報を取得する                                 |
| services figma account info                 | 現在のユーザー情報を取得する                             |
| services figma file export all page         | チーム配下のすべてのファイル/ページをエクスポートする    |
| services figma file export frame            | Figmaファイルの全フレームを書き出す                      |
| services figma file export node             | Figmaドキュメント・ノードの書き出し                      |
| services figma file export page             | Figmaファイルの全ページを書き出す                        |
| services figma file info                    | figmaファイルの情報を表示する                            |
| services figma file list                    | Figmaプロジェクト内のファイル一覧                        |
| services figma project list                 | チームのプロジェクト一覧                                 |
| services github content get                 | レポジトリのコンテンツメタデータを取得します.            |
| services github content put                 | レポジトリに小さなテキストコンテンツを格納します         |
| services github issue list                  | 公開・プライベートGitHubレポジトリの課題一覧             |
| services github profile                     | 認証したユーザーの情報を取得                             |
| services github release asset download      | アセットをダウンロードします                             |
| services github release asset list          | GitHubリリースの成果物一覧                               |
| services github release asset upload        | GitHub リリースへ成果物をアップロードします              |
| services github release draft               | リリースの下書きを作成                                   |
| services github release list                | リリースの一覧                                           |
| services github tag create                  | レポジトリにタグを作成します                             |
| services google calendar event list         | Googleカレンダーのイベントを一覧表示                     |
| services google mail filter add             | フィルターを追加します.                                  |
| services google mail filter batch add       | クエリによるラベルの一括追加・削除                       |
| services google mail filter delete          | フィルタの削除                                           |
| services google mail filter list            | フィルターの一覧                                         |
| services google mail label add              | ラベルの追加                                             |
| services google mail label delete           | ラベルの削除.                                            |
| services google mail label list             | ラベルのリスト                                           |
| services google mail label rename           | ラベルの名前を変更する                                   |
| services google mail message label add      | メッセージにラベルを追加                                 |
| services google mail message label delete   | メッセージからラベルを削除する                           |
| services google mail message list           | メッセージの一覧                                         |
| services google mail message processed list | 処理された形式でメッセージを一覧表示します.              |
| services google mail message send           | メールの送信                                             |
| services google mail sendas add             | カスタムの "from" send-asエイリアスの作成                |
| services google mail sendas delete          | 指定したsend-asエイリアスを削除する                      |
| services google mail sendas list            | 指定されたアカウントの送信エイリアスを一覧表示する       |
| services google mail thread list            | スレッド一覧                                             |
| services google sheets sheet append         | スプレッドシートにデータを追加する                       |
| services google sheets sheet clear          | スプレッドシートから値をクリアする                       |
| services google sheets sheet create         | 新規シートの作成                                         |
| services google sheets sheet delete         | スプレッドシートからシートを削除する                     |
| services google sheets sheet export         | シートデータのエクスポート                               |
| services google sheets sheet import         | スプレッドシートにデータをインポート                     |
| services google sheets sheet list           | スプレッドシートのシート一覧                             |
| services google sheets spreadsheet create   | 新しいスプレッドシートの作成                             |
| services google translate text              | テキストを翻訳する                                       |
| services slack conversation history         | 会話履歴                                                 |
| services slack conversation list            | チャネルの一覧                                           |



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
# コマンド仕様の変更: `member detach`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name: "detach",
  	Title: strings.Join({
  		"Dropbox ",
- 		"BusinessユーザーをBasicユーザーに変更します",
+ 		"for teamsのアカウントをBasicアカウントに変更す\xe3\x82",
+ 		"\x8b",
  	}, ""),
  	Desc:    "",
  	Remarks: "(非可逆な操作です)",
  	... // 20 identical fields
  }
```
# コマンド仕様の変更: `team activity batch user`



## 変更されたレポート: combined

```
  &dc_recipe.Report{
  	Name: "combined",
  	Desc: strings.Join({
  		"このレポートは",
- 		"Dropbox Busines",
+ 		"、主にDropbox for team",
  		"sのアクティビティログと",
- 		"ほぼ",
  		"互換性のあるアクティビティ\xe3\x83",
- 		"\xacポートを出力します.",
+ 		"\xadグを表示します。",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```

## 変更されたレポート: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: strings.Join({
  		"このレポートは",
- 		"Dropbox Busines",
+ 		"、主にDropbox for team",
  		"sのアクティビティログと",
- 		"ほぼ",
  		"互換性のあるアクティビティ\xe3\x83",
- 		"\xacポートを出力します.",
+ 		"\xadグを表示します。",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
# コマンド仕様の変更: `team activity daily event`



## 変更されたレポート: event

```
  &dc_recipe.Report{
  	Name: "event",
  	Desc: strings.Join({
  		"このレポートは",
- 		"Dropbox Busines",
+ 		"、主にDropbox for team",
  		"sのアクティビティログと",
- 		"ほぼ",
  		"互換性のあるアクティビティ\xe3\x83",
- 		"\xacポートを出力します.",
+ 		"\xadグを表示します。",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
# コマンド仕様の変更: `team activity event`



## 変更されたレポート: event

```
  &dc_recipe.Report{
  	Name: "event",
  	Desc: strings.Join({
  		"このレポートは",
- 		"Dropbox Busines",
+ 		"、主にDropbox for team",
  		"sのアクティビティログと",
- 		"ほぼ",
  		"互換性のあるアクティビティ\xe3\x83",
- 		"\xacポートを出力します.",
+ 		"\xadグを表示します。",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
# コマンド仕様の変更: `team activity user`



## 変更されたレポート: user

```
  &dc_recipe.Report{
  	Name: "user",
  	Desc: strings.Join({
  		"このレポートは",
- 		"Dropbox Busines",
+ 		"、主にDropbox for team",
  		"sのアクティビティログと",
- 		"ほぼ",
  		"互換性のあるアクティビティ\xe3\x83",
- 		"\xacポートを出力します.",
+ 		"\xadグを表示します。",
  	}, ""),
  	Columns: {&{Name: "timestamp", Desc: "このアクションが実行されたDropbox側でのタイム"...}, &{Name: "member", Desc: "ユーザーの表示名"}, &{Name: "member_email", Desc: "ユーザーのメールアドレス"}, &{Name: "event_type", Desc: "実行されたアクションのタイプ"}, ...},
  }
```
