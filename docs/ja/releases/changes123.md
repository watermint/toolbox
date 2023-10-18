---
layout: release
title: リリースの変更点 122
lang: ja
---

# `リリース 122` から `リリース 123` までの変更点

# 追加されたコマンド


| コマンド                                    | タイトル                                                                                  |
|---------------------------------------------|-------------------------------------------------------------------------------------------|
| config auth delete                          | 既存の認証クレデンシャルの削除                                                            |
| config auth list                            | すべての認証情報を一覧表示                                                                |
| config disable                              | 機能を無効化します.                                                                       |
| config enable                               | 機能を有効化します.                                                                       |
| config features                             | 利用可能なオプション機能一覧.                                                             |
| dev benchmark local                         | ローカルファイルシステムにダミーのフォルダ構造を作成します.                               |
| dev benchmark upload                        | アップロードのベンチマーク                                                                |
| dev benchmark uploadlink                    | アップロードテンポラリリンクAPIを使ったシングルファイルのアップロードをベンチマーク.      |
| dev build catalogue                         | カタログを生成します                                                                      |
| dev build doc                               | ドキュメントを生成                                                                        |
| dev build info                              | ビルド情報ファイルを生成                                                                  |
| dev build license                           | LICENSE.txtの生成                                                                         |
| dev build package                           | ビルドのパッケージ化                                                                      |
| dev build preflight                         | リリースに向けて必要な事前準備を実施                                                      |
| dev build readme                            | README.txtの生成                                                                          |
| dev ci artifact up                          | CI成果物をアップロードします                                                              |
| dev ci auth export                          | CIビルドのためのデプロイトークンデータの書き出し                                          |
| dev diag endpoint                           | エンドポイントを一覧                                                                      |
| dev diag throughput                         | キャプチャログからスループットを評価                                                      |
| dev kvs concurrency                         | KVSエンジンの同時実行テスト                                                               |
| dev kvs dump                                | KVSデータのダンプ                                                                         |
| dev module list                             | 依存モジュール一覧                                                                        |
| dev release candidate                       | リリース候補を検査します                                                                  |
| dev release doc                             | リリースドキュメントの作成                                                                |
| dev release publish                         | リリースを公開します                                                                      |
| dev replay approve                          | リプレイをテストバンドルとして承認する                                                    |
| dev replay bundle                           | すべてのリプレイを実行                                                                    |
| dev replay recipe                           | レシピのリプレイ実行                                                                      |
| dev replay remote                           | リモートリプレイバンドルの実行                                                            |
| dev spec diff                               | 2リリース間の仕様を比較します                                                             |
| dev spec doc                                | 仕様ドキュメントを生成します                                                              |
| dev stage dbxfs                             | Dropboxのファイルシステムのインプリケーションを確認しますキャッシュされたシステムに対して |
| dev stage encoding                          | エンコードテストコマンド（指定したエンコード名でダミーファイルをアップロードします）      |
| dev stage gmail                             | Gmail コマンド                                                                            |
| dev stage griddata                          | グリッドデータテスト                                                                      |
| dev stage gui launch                        | GUIコンセプト実証                                                                         |
| dev stage http_range                        | HTTPレンジリクエストのプルーフオブコンセプト                                              |
| dev stage scoped                            | Dropboxのスコープ付きOAuthアプリテスト                                                    |
| dev stage teamfolder                        | チームフォルダ処理のサンプル                                                              |
| dev stage upload_append                     | 新しいアップロードAPIテスト                                                               |
| dev test auth all                           | すべてのスコープでのDropboxへの接続テスト                                                 |
| dev test echo                               | テキストのエコー                                                                          |
| dev test panic                              | パニック試験                                                                              |
| dev test recipe                             | レシピのテスト                                                                            |
| dev test resources                          | バイナリの品質テスト                                                                      |
| dev test setup massfiles                    | テストファイルとしてウィキメディアダンプファイルをアップロードする                        |
| dev test setup teamsharedlink               | デモ用共有リンクの作成                                                                    |
| dev util anonymise                          | キャプチャログを匿名化します.                                                             |
| dev util curl                               | capture.logからcURLのプレビューを生成します                                               |
| dev util image jpeg                         | ダミー画像ファイルを作成します                                                            |
| dev util wait                               | 指定した秒数待機します                                                                    |
| file compare account                        | 二つのアカウントのファイルを比較します                                                    |
| file compare local                          | ローカルフォルダとDropboxフォルダの内容を比較します                                       |
| file copy                                   | ファイルをコピーします                                                                    |
| file delete                                 | ファイルまたはフォルダは削除します.                                                       |
| file export doc                             | ドキュメントのエクスポート                                                                |
| file export url                             | URLからドキュメントをエクスポート                                                         |
| file import batch url                       | URLからファイルを一括インポートします                                                     |
| file import url                             | URLからファイルをインポートします                                                         |
| file info                                   | パスのメタデータを解決                                                                    |
| file list                                   | ファイルとフォルダを一覧します                                                            |
| file lock acquire                           | ファイルをロック                                                                          |
| file lock all release                       | 指定したパスでのすべてのロックを解除する                                                  |
| file lock batch acquire                     | 複数のファイルをロックする                                                                |
| file lock batch release                     | 複数のロックを解除                                                                        |
| file lock list                              | 指定したパスの下にあるロックを一覧表示します                                              |
| file lock release                           | ロックを解除します                                                                        |
| file merge                                  | フォルダを統合します                                                                      |
| file move                                   | ファイルを移動します                                                                      |
| file paper append                           | 既存のPaperドキュメントの最後にコンテンツを追加する                                       |
| file paper create                           | パスに新しいPaperを作成                                                                   |
| file paper overwrite                        | 既存のPaperドキュメントを上書きする                                                       |
| file paper prepend                          | 既存のPaperドキュメントの先頭にコンテンツを追加する                                       |
| file replication                            | ファイルコンテンツを他のアカウントに複製します                                            |
| file restore all                            | 指定されたパス以下をリストアします                                                        |
| file revision download                      | ファイルリビジョンをダウンロードする                                                      |
| file revision list                          | ファイルリビジョン一覧                                                                    |
| file revision restore                       | ファイルリビジョンを復元する                                                              |
| file search content                         | ファイルコンテンツを検索                                                                  |
| file search name                            | ファイル名を検索                                                                          |
| file share info                             | ファイルの共有情報を取得する                                                              |
| file size                                   | ストレージの利用量                                                                        |
| file sync down                              | Dropboxと下り方向で同期します                                                             |
| file sync online                            | オンラインファイルを同期します                                                            |
| file sync up                                | Dropboxと上り方向で同期します                                                             |
| file tag add                                | ファイル/フォルダーにタグを追加する                                                       |
| file tag delete                             | ファイル/フォルダーからタグを削除する                                                     |
| file tag list                               | パスのタグを一覧                                                                          |
| file template apply local                   | ファイル/フォルダー構造のテンプレートをローカルパスに適用する                             |
| file template apply remote                  | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する                            |
| file template capture local                 | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む                         |
| file template capture remote                | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。                      |
| file watch                                  | ファイルアクティビティを監視                                                              |
| filerequest create                          | ファイルリクエストを作成します                                                            |
| filerequest delete closed                   | このアカウントの全ての閉じられているファイルリクエストを削除します                        |
| filerequest delete url                      | ファイルリクエストのURLを指定して削除                                                     |
| filerequest list                            | 個人アカウントのファイルリクエストを一覧.                                                 |
| group add                                   | グループを作成します                                                                      |
| group batch add                             | グループの一括追加                                                                        |
| group batch delete                          | グループの削除                                                                            |
| group clear externalid                      | グループの外部IDをクリアする                                                              |
| group delete                                | グループを削除します                                                                      |
| group folder list                           | 各グループのフォルダーを一覧表示                                                          |
| group list                                  | グループを一覧                                                                            |
| group member add                            | メンバーをグループに追加                                                                  |
| group member batch add                      | グループにメンバーを一括追加                                                              |
| group member batch delete                   | グループからメンバーを削除                                                                |
| group member batch update                   | グループからメンバーを追加または削除                                                      |
| group member delete                         | メンバーをグループから削除                                                                |
| group member list                           | グループに所属するメンバー一覧を取得します                                                |
| group rename                                | グループの改名                                                                            |
| group update type                           | グループ管理タイプの更新                                                                  |
| job history archive                         | ジョブのアーカイブ                                                                        |
| job history delete                          | 古いジョブ履歴の削除                                                                      |
| job history list                            | ジョブ履歴の表示                                                                          |
| job history ship                            | ログの転送先Dropboxパス                                                                   |
| job log jobid                               | 指定したジョブIDのログを取得する                                                          |
| job log kind                                | 指定種別のログを結合して出力します                                                        |
| job log last                                | 最後のジョブのログファイルを出力.                                                         |
| license                                     | ライセンス情報を表示します                                                                |
| member batch suspend                        | メンバーの一括一時停止                                                                    |
| member batch unsuspend                      | メンバーの一括停止解除                                                                    |
| member clear externalid                     | メンバーのexternal_idを初期化します                                                       |
| member delete                               | メンバーを削除します                                                                      |
| member detach                               | Dropbox BusinessユーザーをBasicユーザーに変更します                                       |
| member feature                              | メンバーの機能設定一覧                                                                    |
| member file lock all release                | メンバーのパスの下にあるすべてのロックを解除します                                        |
| member file lock list                       | パスの下にあるメンバーのロックを一覧表示                                                  |
| member file lock release                    | メンバーとしてパスのロックを解除します                                                    |
| member file permdelete                      | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                    |
| member folder list                          | 各メンバーのフォルダーを一覧表示                                                          |
| member folder replication                   | フォルダを他のメンバーの個人フォルダに複製します                                          |
| member invite                               | メンバーを招待します                                                                      |
| member list                                 | チームメンバーの一覧                                                                      |
| member quota list                           | メンバーの容量制限情報を一覧します                                                        |
| member quota update                         | チームメンバーの容量制限を変更                                                            |
| member quota usage                          | チームメンバーのストレージ利用状況を取得                                                  |
| member reinvite                             | 招待済み状態メンバーをチームに再招待します                                                |
| member replication                          | チームメンバーのファイルを複製します                                                      |
| member suspend                              | メンバーの一時停止処理                                                                    |
| member unsuspend                            | メンバーの一時停止を解除する                                                              |
| member update email                         | メンバーのメールアドレス処理                                                              |
| member update externalid                    | チームメンバーのExternal IDを更新します.                                                  |
| member update invisible                     | メンバーへのディレクトリ制限を有効にします                                                |
| member update profile                       | メンバーのプロフィール変更                                                                |
| member update visible                       | メンバーへのディレクトリ制限を無効にします                                                |
| services asana team list                    | チームのリスト                                                                            |
| services asana team project list            | チームのプロジェクト一覧                                                                  |
| services asana team task list               | チームのタスク一覧                                                                        |
| services asana workspace list               | ワークスペースの一覧                                                                      |
| services asana workspace project list       | ワークスペースのプロジェクト一覧                                                          |
| services dropbox user feature               | 現在のユーザーの機能設定の一覧                                                            |
| services dropbox user info                  | 現在のアカウント情報を取得する                                                            |
| services dropboxsign account info           | アカウント情報を取得する                                                                  |
| services figma account info                 | 現在のユーザー情報を取得する                                                              |
| services figma file export all page         | チーム配下のすべてのファイル/ページをエクスポートする                                     |
| services figma file export frame            | Figmaファイルの全フレームを書き出す                                                       |
| services figma file export node             | Figmaドキュメント・ノードの書き出し                                                       |
| services figma file export page             | Figmaファイルの全ページを書き出す                                                         |
| services figma file info                    | figmaファイルの情報を表示する                                                             |
| services figma file list                    | Figmaプロジェクト内のファイル一覧                                                         |
| services figma project list                 | チームのプロジェクト一覧                                                                  |
| services github content get                 | レポジトリのコンテンツメタデータを取得します.                                             |
| services github content put                 | レポジトリに小さなテキストコンテンツを格納します                                          |
| services github issue list                  | 公開・プライベートGitHubレポジトリの課題一覧                                              |
| services github profile                     | 認証したユーザーの情報を取得                                                              |
| services github release asset download      | アセットをダウンロードします                                                              |
| services github release asset list          | GitHubリリースの成果物一覧                                                                |
| services github release asset upload        | GitHub リリースへ成果物をアップロードします                                               |
| services github release draft               | リリースの下書きを作成                                                                    |
| services github release list                | リリースの一覧                                                                            |
| services github tag create                  | レポジトリにタグを作成します                                                              |
| services google calendar event list         | Googleカレンダーのイベントを一覧表示                                                      |
| services google mail filter add             | フィルターを追加します.                                                                   |
| services google mail filter batch add       | クエリによるラベルの一括追加・削除                                                        |
| services google mail filter delete          | フィルタの削除                                                                            |
| services google mail filter list            | フィルターの一覧                                                                          |
| services google mail label add              | ラベルの追加                                                                              |
| services google mail label delete           | ラベルの削除.                                                                             |
| services google mail label list             | ラベルのリスト                                                                            |
| services google mail label rename           | ラベルの名前を変更する                                                                    |
| services google mail message label add      | メッセージにラベルを追加                                                                  |
| services google mail message label delete   | メッセージからラベルを削除する                                                            |
| services google mail message list           | メッセージの一覧                                                                          |
| services google mail message processed list | 処理された形式でメッセージを一覧表示します.                                               |
| services google mail message send           | メールの送信                                                                              |
| services google mail sendas add             | カスタムの "from" send-asエイリアスの作成                                                 |
| services google mail sendas delete          | 指定したsend-asエイリアスを削除する                                                       |
| services google mail sendas list            | 指定されたアカウントの送信エイリアスを一覧表示する                                        |
| services google mail thread list            | スレッド一覧                                                                              |
| services google sheets sheet append         | スプレッドシートにデータを追加する                                                        |
| services google sheets sheet clear          | スプレッドシートから値をクリアする                                                        |
| services google sheets sheet create         | 新規シートの作成                                                                          |
| services google sheets sheet delete         | スプレッドシートからシートを削除する                                                      |
| services google sheets sheet export         | シートデータのエクスポート                                                                |
| services google sheets sheet import         | スプレッドシートにデータをインポート                                                      |
| services google sheets sheet list           | スプレッドシートのシート一覧                                                              |
| services google sheets spreadsheet create   | 新しいスプレッドシートの作成                                                              |
| services google translate text              | テキストを翻訳する                                                                        |
| services slack conversation history         | 会話履歴                                                                                  |
| services slack conversation list            | チャネルの一覧                                                                            |
| sharedfolder leave                          | 共有フォルダーから退出する.                                                               |
| sharedfolder list                           | 共有フォルダの一覧                                                                        |
| sharedfolder member add                     | 共有フォルダへのメンバーの追加                                                            |
| sharedfolder member delete                  | 共有フォルダからメンバーを削除する                                                        |
| sharedfolder member list                    | 共有フォルダのメンバーを一覧します                                                        |
| sharedfolder mount add                      | 共有フォルダを現在のユーザーのDropboxに追加する                                           |
| sharedfolder mount delete                   | 現在のユーザーが指定されたフォルダーをアンマウントする.                                   |
| sharedfolder mount list                     | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示                          |
| sharedfolder mount mountable                | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします.                 |
| sharedfolder share                          | フォルダの共有                                                                            |
| sharedfolder unshare                        | フォルダの共有解除                                                                        |
| sharedlink create                           | 共有リンクの作成                                                                          |
| sharedlink delete                           | 共有リンクを削除します                                                                    |
| sharedlink file list                        | 共有リンクのファイルを一覧する                                                            |
| sharedlink info                             | 共有リンクの情報取得                                                                      |
| sharedlink list                             | 共有リンクの一覧                                                                          |
| team activity batch user                    | 複数ユーザーのアクティビティを一括取得します                                              |
| team activity daily event                   | アクティビティーを1日ごとに取得します                                                     |
| team activity event                         | イベントログ                                                                              |
| team activity user                          | ユーザーごとのアクティビティ                                                              |
| team admin group role add                   | グループのメンバーにロールを追加する                                                      |
| team admin group role delete                | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                          |
| team admin list                             | メンバーの管理者権限一覧                                                                  |
| team admin role add                         | メンバーに新しいロールを追加する                                                          |
| team admin role clear                       | メンバーからすべての管理者ロールを削除する                                                |
| team admin role delete                      | メンバーからロールを削除する                                                              |
| team admin role list                        | チームの管理者の役割を列挙                                                                |
| team content legacypaper count              | メンバー1人あたりのPaper文書の枚数                                                        |
| team content legacypaper export             | チームメンバー全員のPaper文書をローカルパスにエクスポート.                                |
| team content legacypaper list               | チームメンバーのPaper文書リスト出力                                                       |
| team content member list                    | チームフォルダや共有フォルダのメンバー一覧                                                |
| team content member size                    | チームフォルダや共有フォルダのメンバー数をカウントする                                    |
| team content mount list                     | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.    |
| team content policy list                    | チームフォルダと共有フォルダのポリシー一覧                                                |
| team device list                            | チーム内全てのデバイス/セッションを一覧します                                             |
| team device unlink                          | デバイスのセッションを解除します                                                          |
| team feature                                | チームの機能を出力します                                                                  |
| team filerequest clone                      | ファイルリクエストを入力データに従い複製します                                            |
| team filerequest list                       | チームないのファイルリクエストを一覧します                                                |
| team info                                   | チームの情報                                                                              |
| team insight scan                           | チーム全体の情報をスキャンする                                                            |
| team insight summarize                      | スキャンしたチーム情報をまとめる                                                          |
| team legalhold add                          | 新しいリーガル・ホールド・ポリシーを作成する.                                             |
| team legalhold list                         | 既存のポリシーを取得する                                                                  |
| team legalhold member batch update          | リーガル・ホールド・ポリシーのメンバーリスト更新                                          |
| team legalhold member list                  | リーガルホールドのメンバーをリストアップ                                                  |
| team legalhold release                      | Idによるリーガルホールドを解除する                                                        |
| team legalhold revision list                | リーガル・ホールド・ポリシーのリビジョンをリストアップする                                |
| team legalhold update desc                  | リーガルホールド・ポリシーの説明を更新                                                    |
| team legalhold update name                  | リーガルホールドポリシーの名称を更新                                                      |
| team linkedapp list                         | リンク済みアプリを一覧                                                                    |
| team namespace file list                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                        |
| team namespace file size                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                        |
| team namespace list                         | チーム内すべての名前空間を一覧                                                            |
| team namespace member list                  | チームフォルダ以下のファイル・フォルダを一覧                                              |
| team namespace summary                      | チーム・ネームスペースの状態概要を報告する.                                               |
| team report activity                        | アクティビティ レポート                                                                   |
| team report devices                         | デバイス レポート空のレポート                                                             |
| team report membership                      | メンバーシップ レポート                                                                   |
| team report storage                         | ストレージ レポート                                                                       |
| team runas file batch copy                  | ファイル/フォルダーをメンバーとして一括コピー                                             |
| team runas file list                        | メンバーとして実行するファイルやフォルダーの一覧                                          |
| team runas file sync batch up               | メンバーとして動作する一括同期                                                            |
| team runas sharedfolder batch leave         | 共有フォルダからメンバーとして一括退出                                                    |
| team runas sharedfolder batch share         | メンバーのフォルダを一括で共有                                                            |
| team runas sharedfolder batch unshare       | メンバーのフォルダの共有を一括解除                                                        |
| team runas sharedfolder isolate             | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する.   |
| team runas sharedfolder list                | 共有フォルダーの一覧をメンバーとして実行                                                  |
| team runas sharedfolder member batch add    | メンバーの共有フォルダにメンバーを一括追加                                                |
| team runas sharedfolder member batch delete | メンバーの共有フォルダからメンバーを一括削除                                              |
| team runas sharedfolder mount add           | 指定したメンバーのDropboxに共有フォルダを追加する                                         |
| team runas sharedfolder mount delete        | 指定されたユーザーが指定されたフォルダーをアンマウントする.                               |
| team runas sharedfolder mount list          | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.           |
| team runas sharedfolder mount mountable     | メンバーがマウントできるすべての共有フォルダーをリストアップ.                             |
| team sharedlink cap expiry                  | チーム内の共有リンクに有効期限の上限を設定                                                |
| team sharedlink cap visibility              | チーム内の共有リンクに可視性の上限を設定                                                  |
| team sharedlink delete links                | 共有リンクの一括削除                                                                      |
| team sharedlink delete member               | メンバーの共有リンクをすべて削除                                                          |
| team sharedlink list                        | 共有リンクの一覧                                                                          |
| team sharedlink update expiry               | チーム内の公開されている共有リンクについて有効期限を更新します                            |
| team sharedlink update password             | 共有リンクのパスワードの設定・更新                                                        |
| team sharedlink update visibility           | 共有リンクの可視性の更新                                                                  |
| teamfolder add                              | チームフォルダを追加します                                                                |
| teamfolder archive                          | チームフォルダのアーカイブ                                                                |
| teamfolder batch archive                    | 複数のチームフォルダをアーカイブします                                                    |
| teamfolder batch permdelete                 | 複数のチームフォルダを完全に削除します                                                    |
| teamfolder batch replication                | チームフォルダの一括レプリケーション                                                      |
| teamfolder file list                        | チームフォルダの一覧                                                                      |
| teamfolder file lock all release            | チームフォルダのパスの下にあるすべてのロックを解除する                                    |
| teamfolder file lock list                   | チームフォルダ内のロックを一覧表示                                                        |
| teamfolder file lock release                | チームフォルダ内のパスのロックを解除                                                      |
| teamfolder file size                        | チームフォルダのサイズを計算                                                              |
| teamfolder list                             | チームフォルダの一覧                                                                      |
| teamfolder member add                       | チームフォルダへのユーザー/グループの一括追加                                             |
| teamfolder member delete                    | チームフォルダからのユーザー/グループの一括削除                                           |
| teamfolder member list                      | チームフォルダのメンバー一覧                                                              |
| teamfolder partial replication              | 部分的なチームフォルダの他チームへのレプリケーション                                      |
| teamfolder permdelete                       | チームフォルダを完全に削除します                                                          |
| teamfolder policy list                      | チームフォルダのポリシー一覧                                                              |
| teamfolder replication                      | チームフォルダを他のチームに複製します                                                    |
| teamfolder sync setting list                | チームフォルダーの同期設定を一覧表示                                                      |
| teamfolder sync setting update              | チームフォルダ同期設定の一括更新                                                          |
| teamspace asadmin file list                 | チームスペース内のファイルやフォルダーを一覧表示することができます。                      |
| teamspace asadmin folder add                | チームスペースにトップレベルのフォルダーを作成                                            |
| teamspace asadmin folder delete             | チームスペースのトップレベルフォルダーを削除する                                          |
| teamspace asadmin folder permdelete         | チームスペースのトップレベルフォルダを完全に削除します。                                  |
| teamspace asadmin member list               | トップレベルのフォルダーメンバーをリストアップ                                            |
| teamspace file list                         | チームスペースにあるファイルやフォルダーを一覧表示                                        |
| util archive unzip                          | ZIPアーカイブファイルを解凍する                                                           |
| util archive zip                            | 対象ファイルをZIPアーカイブに圧縮する                                                     |
| util cert selfsigned                        | 自己署名証明書と鍵の生成                                                                  |
| util database exec                          | SQLite3データベースファイルへのクエリ実行                                                 |
| util database query                         | SQLite3データベースへの問い合わせ                                                         |
| util date today                             | 現在の日付を表示                                                                          |
| util datetime now                           | 現在の日時を表示                                                                          |
| util decode base32                          | Base32 (RFC 4648) 形式からテキストをデコードします                                        |
| util decode base64                          | Base64 (RFC 4648) フォーマットからテキストをデコードします                                |
| util encode base32                          | テキストをBase32(RFC 4648)形式にエンコード                                                |
| util encode base64                          | テキストをBase64(RFC 4648)形式にエンコード                                                |
| util file hash                              | ファイルダイジェストの表示                                                                |
| util git clone                              | git リポジトリをクローン                                                                  |
| util image exif                             | 画像ファイルのEXIFメタデータを表示                                                        |
| util image placeholder                      | プレースホルダー画像の作成                                                                |
| util monitor client                         | デバイスモニタークライアントを起動する                                                    |
| util net download                           | ファイルをダウンロードする                                                                |
| util qrcode create                          | QRコード画像ファイルの作成                                                                |
| util qrcode wifi                            | WIFI設定用のQRコードを生成                                                                |
| util release install                        | watermint toolboxをダウンロードし、パスにインストールします。                             |
| util table format xlsx                      | xlsxファイルをテキストに整形する                                                          |
| util text case down                         | 小文字のテキストを表示する                                                                |
| util text case up                           | 大文字のテキストを表示する                                                                |
| util text encoding from                     | 指定されたエンコーディングからUTF-8テキストファイルに変換します.                          |
| util text encoding to                       | UTF-8テキストファイルから指定されたエンコーディングに変換する.                            |
| util text nlp english entity                | 英文をエンティティに分割する                                                              |
| util text nlp english sentence              | 英文を文章に分割する                                                                      |
| util text nlp english token                 | 英文をトークンに分割する                                                                  |
| util text nlp japanese token                | 日本語テキストのトークン化                                                                |
| util text nlp japanese wakati               | 分かち書き(日本語テキストのトークン化)                                                    |
| util tidy move dispatch                     | ファイルを整理                                                                            |
| util tidy move simple                       | ローカルファイルをアーカイブします                                                        |
| util tidy pack remote                       | リモートフォルダをZIPファイルにパッケージする                                             |
| util time now                               | 現在の時刻を表示                                                                          |
| util unixtime format                        | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット                    |
| util unixtime now                           | UNIX時間で現在の時刻を表示する                                                            |
| util uuid v4                                | UUID v4（ランダムUUID）の生成                                                             |
| util video subtitles optimize               | 字幕ファイルの最適化                                                                      |
| util xlsx create                            | 空のスプレッドシートを作成する                                                            |
| util xlsx sheet export                      | xlsxファイルからデータをエクスポート                                                      |
| util xlsx sheet import                      | データをxlsxファイルにインポート                                                          |
| util xlsx sheet list                        | xlsxファイルのシート一覧                                                                  |
| version                                     | バージョン情報                                                                            |



