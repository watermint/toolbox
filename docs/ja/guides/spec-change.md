---
layout: page
title: 仕様変更
lang: ja
---

# 仕様変更

# コマンドパスの変更

現在のバージョンを使い続けても影響はありませんが、変更は将来のバージョンに適用されます。日付が指定されている場合、その日付以降にリリースされたバージョンに変更が適用される。

| 旧パス                                                                                              | 現在のパス                         | コマンドの説明                                        | 日付                 |
|-----------------------------------------------------------------------------------------------------|------------------------------------|-------------------------------------------------------|----------------------|
| [config disable](https://github.com/watermint/toolbox/discussions/795)                              | config feature disable             | 機能を無効化します.                                   | 2024-04-01T00:00:00Z |
| [config enable](https://github.com/watermint/toolbox/discussions/795)                               | config feature enable              | 機能を有効化します.                                   | 2024-04-01T00:00:00Z |
| [config features](https://github.com/watermint/toolbox/discussions/795)                             | config feature list                | 利用可能なオプション機能一覧.                         | 2024-04-01T00:00:00Z |
| [services deepl translate text](https://github.com/watermint/toolbox/discussions/797)               | deepl translate text               | テキストを翻訳する                                    | 2024-06-30T15:00:00Z |
| [services dropbox user feature](https://github.com/watermint/toolbox/discussions/797)               | dropbox file account feature       | Dropboxアカウントの機能一覧                           | 2024-06-30T15:00:00Z |
| [services dropbox user filesystem](https://github.com/watermint/toolbox/discussions/797)            | dropbox file account filesystem    | Dropboxのファイルシステムのバージョンを表示する       | 2024-06-30T15:00:00Z |
| [services dropbox user info](https://github.com/watermint/toolbox/discussions/797)                  | dropbox file account info          | Dropboxアカウント情報                                 | 2024-06-30T15:00:00Z |
| [services figma account info](https://github.com/watermint/toolbox/discussions/797)                 | figma account info                 | 現在のユーザー情報を取得する                          | 2024-06-30T15:00:00Z |
| [services figma file export all page](https://github.com/watermint/toolbox/discussions/797)         | figma file export all page         | チーム配下のすべてのファイル/ページをエクスポートする | 2024-06-30T15:00:00Z |
| [services figma file export frame](https://github.com/watermint/toolbox/discussions/797)            | figma file export frame            | Figmaファイルの全フレームを書き出す                   | 2024-06-30T15:00:00Z |
| [services figma file export node](https://github.com/watermint/toolbox/discussions/797)             | figma file export node             | Figmaドキュメント・ノードの書き出し                   | 2024-06-30T15:00:00Z |
| [services figma file export page](https://github.com/watermint/toolbox/discussions/797)             | figma file export page             | Figmaファイルの全ページを書き出す                     | 2024-06-30T15:00:00Z |
| [services figma file info](https://github.com/watermint/toolbox/discussions/797)                    | figma file info                    | figmaファイルの情報を表示する                         | 2024-06-30T15:00:00Z |
| [services figma file list](https://github.com/watermint/toolbox/discussions/797)                    | figma file list                    | Figmaプロジェクト内のファイル一覧                     | 2024-06-30T15:00:00Z |
| [services figma project list](https://github.com/watermint/toolbox/discussions/797)                 | figma project list                 | チームのプロジェクト一覧                              | 2024-06-30T15:00:00Z |
| [services google calendar event list](https://github.com/watermint/toolbox/discussions/797)         | google calendar event list         | Googleカレンダーのイベントを一覧表示                  | 2024-06-30T15:00:00Z |
| [services google mail filter add](https://github.com/watermint/toolbox/discussions/797)             | google mail filter add             | フィルターを追加します.                               | 2024-06-30T15:00:00Z |
| [services google mail filter batch add](https://github.com/watermint/toolbox/discussions/797)       | google mail filter batch add       | クエリによるラベルの一括追加・削除                    | 2024-06-30T15:00:00Z |
| [services google mail filter delete](https://github.com/watermint/toolbox/discussions/797)          | google mail filter delete          | フィルタの削除                                        | 2024-06-30T15:00:00Z |
| [services google mail filter list](https://github.com/watermint/toolbox/discussions/797)            | google mail filter list            | フィルターの一覧                                      | 2024-06-30T15:00:00Z |
| [services google mail label add](https://github.com/watermint/toolbox/discussions/797)              | google mail label add              | ラベルの追加                                          | 2024-06-30T15:00:00Z |
| [services google mail label delete](https://github.com/watermint/toolbox/discussions/797)           | google mail label delete           | ラベルの削除.                                         | 2024-06-30T15:00:00Z |
| [services google mail label list](https://github.com/watermint/toolbox/discussions/797)             | google mail label list             | ラベルのリスト                                        | 2024-06-30T15:00:00Z |
| [services google mail label rename](https://github.com/watermint/toolbox/discussions/797)           | google mail label rename           | ラベルの名前を変更する                                | 2024-06-30T15:00:00Z |
| [services google mail message label add](https://github.com/watermint/toolbox/discussions/797)      | google mail message label add      | メッセージにラベルを追加                              | 2024-06-30T15:00:00Z |
| [services google mail message label delete](https://github.com/watermint/toolbox/discussions/797)   | google mail message label delete   | メッセージからラベルを削除する                        | 2024-06-30T15:00:00Z |
| [services google mail message list](https://github.com/watermint/toolbox/discussions/797)           | google mail message list           | メッセージの一覧                                      | 2024-06-30T15:00:00Z |
| [services google mail message processed list](https://github.com/watermint/toolbox/discussions/797) | google mail message processed list | 処理された形式でメッセージを一覧表示します.           | 2024-06-30T15:00:00Z |
| [services google mail sendas add](https://github.com/watermint/toolbox/discussions/797)             | google mail sendas add             | カスタムの "from" send-asエイリアスの作成             | 2024-06-30T15:00:00Z |
| [services google mail sendas delete](https://github.com/watermint/toolbox/discussions/797)          | google mail sendas delete          | 指定したsend-asエイリアスを削除する                   | 2024-06-30T15:00:00Z |
| [services google mail sendas list](https://github.com/watermint/toolbox/discussions/797)            | google mail sendas list            | 指定されたアカウントの送信エイリアスを一覧表示する    | 2024-06-30T15:00:00Z |
| [services google mail thread list](https://github.com/watermint/toolbox/discussions/797)            | google mail thread list            | スレッド一覧                                          | 2024-06-30T15:00:00Z |
| [services google sheets sheet append](https://github.com/watermint/toolbox/discussions/797)         | google sheets sheet append         | スプレッドシートにデータを追加する                    | 2024-06-30T15:00:00Z |
| [services google sheets sheet clear](https://github.com/watermint/toolbox/discussions/797)          | google sheets sheet clear          | スプレッドシートから値をクリアする                    | 2024-06-30T15:00:00Z |
| [services google sheets sheet create](https://github.com/watermint/toolbox/discussions/797)         | google sheets sheet create         | 新規シートの作成                                      | 2024-06-30T15:00:00Z |
| [services google sheets sheet delete](https://github.com/watermint/toolbox/discussions/797)         | google sheets sheet delete         | スプレッドシートからシートを削除する                  | 2024-06-30T15:00:00Z |
| [services google sheets sheet export](https://github.com/watermint/toolbox/discussions/797)         | google sheets sheet export         | シートデータのエクスポート                            | 2024-06-30T15:00:00Z |
| [services google sheets sheet import](https://github.com/watermint/toolbox/discussions/797)         | google sheets sheet import         | スプレッドシートにデータをインポート                  | 2024-06-30T15:00:00Z |
| [services google sheets sheet list](https://github.com/watermint/toolbox/discussions/797)           | google sheets sheet list           | スプレッドシートのシート一覧                          | 2024-06-30T15:00:00Z |
| [services google sheets spreadsheet create](https://github.com/watermint/toolbox/discussions/797)   | google sheets spreadsheet create   | 新しいスプレッドシートの作成                          | 2024-06-30T15:00:00Z |
| [job log jobid](https://github.com/watermint/toolbox/discussions/794)                               | log cat job                        | 指定したジョブIDのログを取得する                      | 2024-04-01T00:00:00Z |
| [job log kind](https://github.com/watermint/toolbox/discussions/794)                                | log cat kind                       | 指定種別のログを結合して出力します                    | 2024-04-01T00:00:00Z |
| [job log last](https://github.com/watermint/toolbox/discussions/794)                                | log cat last                       | 最後のジョブのログファイルを出力.                     | 2024-04-01T00:00:00Z |
| [job history archive](https://github.com/watermint/toolbox/discussions/794)                         | log job archive                    | ジョブのアーカイブ                                    | 2024-04-01T00:00:00Z |
| [job history delete](https://github.com/watermint/toolbox/discussions/794)                          | log job delete                     | 古いジョブ履歴の削除                                  | 2024-04-01T00:00:00Z |
| [job history list](https://github.com/watermint/toolbox/discussions/794)                            | log job list                       | ジョブ履歴の表示                                      | 2024-04-01T00:00:00Z |

# 非推奨

以下のコマンドは将来のリリースで削除される予定です。現在のバージョンを使い続けても影響はありませんが、変更は将来のバージョンに適用されます。日付が指定されている場合、その日付以降にリリースされたバージョンに変更が適用される。

| パス                                                                                        | コマンドの説明                                                       | 日付                 |
|---------------------------------------------------------------------------------------------|----------------------------------------------------------------------|----------------------|
| log job ship                                                                                | ログの転送先Dropboxパス                                              | 2024-02-01T00:00:00Z |
| [teamspace asadmin file list](https://github.com/watermint/toolbox/discussions/796)         | チームスペース内のファイルやフォルダーを一覧表示することができます。 | 2024-07-01T00:00:00Z |
| [teamspace asadmin folder add](https://github.com/watermint/toolbox/discussions/796)        | チームスペースにトップレベルのフォルダーを作成                       | 2024-07-01T00:00:00Z |
| [teamspace asadmin folder delete](https://github.com/watermint/toolbox/discussions/796)     | チームスペースのトップレベルフォルダーを削除する                     | 2024-07-01T00:00:00Z |
| [teamspace asadmin folder permdelete](https://github.com/watermint/toolbox/discussions/796) | チームスペースのトップレベルフォルダを完全に削除します。             | 2024-07-01T00:00:00Z |
| [teamspace file list](https://github.com/watermint/toolbox/discussions/796)                 | チームスペースにあるファイルやフォルダーを一覧表示                   | 2024-07-01T00:00:00Z |


