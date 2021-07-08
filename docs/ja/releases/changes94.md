---
layout: release
title: リリースの変更点 93
lang: ja
---

# `リリース 93` から `リリース 94` までの変更点

# 追加されたコマンド


| コマンド                                    | タイトル                                                                                  |
|---------------------------------------------|-------------------------------------------------------------------------------------------|
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
| dev ci artifact connect                     | CI成果物をアップロードするためのDropboxアカウントに接続                                   |
| dev ci artifact up                          | CI成果物をアップロードします                                                              |
| dev ci auth connect                         | エンドツーエンドテストのための認証                                                        |
| dev ci auth import                          | 環境変数はエンドツーエンドトークンをインポートします                                      |
| dev diag endpoint                           | エンドポイントを一覧                                                                      |
| dev diag throughput                         | キャプチャログからスループットを評価                                                      |
| dev kvs dump                                | KVSデータのダンプ                                                                         |
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
| dev stage gmail                             | Gmail コマンド                                                                            |
| dev stage griddata                          | グリッドデータテスト                                                                      |
| dev stage gui launch                        | GUI proof of concept                                                                      |
| dev stage http_range                        | HTTPレンジリクエストのプルーフオブコンセプト                                              |
| dev stage scoped                            | Dropboxのスコープ付きOAuthアプリテスト                                                    |
| dev stage teamfolder                        | チームフォルダ処理のサンプル                                                              |
| dev stage upload_append                     | 新しいアップロードAPIテスト                                                               |
| dev test auth all                           | Test for connect to Dropbox with all scopes                                               |
| dev test echo                               | テキストのエコー                                                                          |
| dev test kvsfootprint                       | KVSのメモリフットプリントをテストします                                                   |
| dev test monkey                             | モンキーテスト                                                                            |
| dev test panic                              | Panic test                                                                                |
| dev test recipe                             | レシピのテスト                                                                            |
| dev test resources                          | バイナリの品質テスト                                                                      |
| dev test setup teamsharedlink               | デモ用共有リンクの作成                                                                    |
| dev util anonymise                          | キャプチャログを匿名化します.                                                             |
| dev util curl                               | capture.logからcURLのプレビューを生成します                                               |
| dev util image jpeg                         | ダミー画像ファイルを作成します                                                            |
| dev util wait                               | 指定した秒数待機します                                                                    |
| file archive local                          | ローカルファイルをアーカイブします                                                        |
| file compare account                        | 二つのアカウントのファイルを比較します                                                    |
| file compare local                          | ローカルフォルダとDropboxフォルダの内容を比較します                                       |
| file copy                                   | ファイルをコピーします                                                                    |
| file delete                                 | ファイルまたはフォルダは削除します.                                                       |
| file dispatch local                         | ローカルファイルを整理します                                                              |
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
| file mount list                             | マウント/アンマウントされた共有フォルダの一覧                                             |
| file move                                   | ファイルを移動します                                                                      |
| file paper append                           | 既存のPaperドキュメントの最後にコンテンツを追加する                                       |
| file paper create                           | パスに新しいPaperを作成                                                                   |
| file paper overwrite                        | 既存のPaperドキュメントを上書きする                                                       |
| file paper prepend                          | 既存のPaperドキュメントの先頭にコンテンツを追加する                                       |
| file replication                            | ファイルコンテンツを他のアカウントに複製します                                            |
| file restore all                            | 指定されたパス以下をリストアします                                                        |
| file search content                         | ファイルコンテンツを検索                                                                  |
| file search name                            | ファイル名を検索                                                                          |
| file size                                   | ストレージの利用量                                                                        |
| file sync down                              | Dropboxと下り方向で同期します                                                             |
| file sync online                            | オンラインファイルを同期します                                                            |
| file sync up                                | Dropboxと上り方向で同期します                                                             |
| file watch                                  | ファイルアクティビティを監視                                                              |
| filerequest create                          | ファイルリクエストを作成します                                                            |
| filerequest delete closed                   | このアカウントの全ての閉じられているファイルリクエストを削除します                        |
| filerequest delete url                      | ファイルリクエストのURLを指定して削除                                                     |
| filerequest list                            | 個人アカウントのファイルリクエストを一覧.                                                 |
| group add                                   | グループを作成します                                                                      |
| group batch delete                          | グループの削除                                                                            |
| group delete                                | グループを削除します                                                                      |
| group folder list                           | 各グループのフォルダを探す                                                                |
| group list                                  | グループを一覧                                                                            |
| group member add                            | メンバーをグループに追加                                                                  |
| group member batch add                      | グループにメンバーを一括追加                                                              |
| group member batch delete                   | グループからメンバーを削除                                                                |
| group member batch update                   | グループからメンバーを追加または削除                                                      |
| group member delete                         | メンバーをグループから削除                                                                |
| group member list                           | グループに所属するメンバー一覧を取得します                                                |
| group rename                                | グループの改名                                                                            |
| image info                                  | 画像ファイルのEXIF情報を表示します                                                        |
| job history archive                         | ジョブのアーカイブ                                                                        |
| job history delete                          | 古いジョブ履歴の削除                                                                      |
| job history list                            | ジョブ履歴の表示                                                                          |
| job history ship                            | ログの転送先Dropboxパス                                                                   |
| job log jobid                               | 指定したジョブIDのログを取得する                                                          |
| job log kind                                | 指定種別のログを結合して出力します                                                        |
| job log last                                | 最後のジョブのログファイルを出力.                                                         |
| license                                     | ライセンス情報を表示します                                                                |
| member clear externalid                     | メンバーのexternal_idを初期化します                                                       |
| member delete                               | メンバーを削除します                                                                      |
| member detach                               | Dropbox BusinessユーザーをBasicユーザーに変更します                                       |
| member file lock all release                | メンバーのパスの下にあるすべてのロックを解除します                                        |
| member file lock list                       | パスの下にあるメンバーのロックを一覧表示                                                  |
| member file lock release                    | メンバーとしてパスのロックを解除します                                                    |
| member file permdelete                      | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                    |
| member folder list                          | 各メンバーのフォルダを検索                                                                |
| member folder replication                   | フォルダを他のメンバーの個人フォルダに複製します                                          |
| member invite                               | メンバーを招待します                                                                      |
| member list                                 | チームメンバーの一覧                                                                      |
| member quota list                           | メンバーの容量制限情報を一覧します                                                        |
| member quota update                         | チームメンバーの容量制限を変更                                                            |
| member quota usage                          | チームメンバーのストレージ利用状況を取得                                                  |
| member reinvite                             | 招待済み状態メンバーをチームに再招待します                                                |
| member replication                          | チームメンバーのファイルを複製します                                                      |
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
| services google mail message send           | Send a mail                                                                               |
| services google mail sendas add             | カスタムの "from" send-asエイリアスの作成                                                 |
| services google mail sendas delete          | 指定したsend-asエイリアスを削除する                                                       |
| services google mail sendas list            | 指定されたアカウントの送信エイリアスを一覧表示する                                        |
| services google mail thread list            | スレッド一覧                                                                              |
| services google sheets sheet append         | スプレッドシートにデータを追加する                                                        |
| services google sheets sheet clear          | スプレッドシートから値をクリアする                                                        |
| services google sheets sheet export         | シートデータのエクスポート                                                                |
| services google sheets sheet import         | スプレッドシートにデータをインポート                                                      |
| services google sheets sheet list           | スプレッドシートのシート一覧                                                              |
| services google sheets spreadsheet create   | 新しいスプレッドシートの作成                                                              |
| services slack conversation list            | チャネルの一覧                                                                            |
| sharedfolder list                           | 共有フォルダの一覧                                                                        |
| sharedfolder member list                    | 共有フォルダのメンバーを一覧します                                                        |
| sharedlink create                           | 共有リンクの作成                                                                          |
| sharedlink delete                           | 共有リンクを削除します                                                                    |
| sharedlink file list                        | 共有リンクのファイルを一覧する                                                            |
| sharedlink info                             | 共有リンクの情報取得                                                                      |
| sharedlink list                             | 共有リンクの一覧                                                                          |
| team activity batch user                    | 複数ユーザーのアクティビティを一括取得します                                              |
| team activity daily event                   | アクティビティーを1日ごとに取得します                                                     |
| team activity event                         | イベントログ                                                                              |
| team activity user                          | ユーザーごとのアクティビティ                                                              |
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
| team linkedapp list                         | リンク済みアプリを一覧                                                                    |
| team namespace file list                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                        |
| team namespace file size                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                        |
| team namespace list                         | チーム内すべての名前空間を一覧                                                            |
| team namespace member list                  | チームフォルダ以下のファイル・フォルダを一覧                                              |
| team report activity                        | アクティビティ レポート                                                                   |
| team report devices                         | デバイス レポート空のレポート                                                             |
| team report membership                      | メンバーシップ レポート                                                                   |
| team report storage                         | ストレージ レポート                                                                       |
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
| util date today                             | 現在の日付を表示                                                                          |
| util datetime now                           | 現在の日時を表示                                                                          |
| util decode base_32                         | Base32 (RFC 4648) 形式からテキストをデコードします                                        |
| util decode base_64                         | Base64 (RFC 4648) フォーマットからテキストをデコードします                                |
| util encode base_32                         | テキストをBase32(RFC 4648)形式にエンコード                                                |
| util encode base_64                         | テキストをBase64(RFC 4648)形式にエンコード                                                |
| util git clone                              | git リポジトリをクローン                                                                  |
| util qrcode create                          | QRコード画像ファイルの作成                                                                |
| util qrcode wifi                            | WIFI設定用のQRコードを生成                                                                |
| util time now                               | 現在の時刻を表示                                                                          |
| util unixtime format                        | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット                    |
| util unixtime now                           | UNIX時間で現在の時刻を表示する                                                            |
| util xlsx create                            | 空のスプレッドシートを作成する                                                            |
| util xlsx sheet export                      | xlsxファイルからデータをエクスポート                                                      |
| util xlsx sheet import                      | データをxlsxファイルにインポート                                                          |
| util xlsx sheet list                        | xlsxファイルのシート一覧                                                                  |
| version                                     | バージョン情報                                                                            |



