---
layout: release
title: リリースの変更点 140
lang: ja
---

# `リリース 140` から `リリース 141` までの変更点

# 削除されたコマンド


| コマンド                                            | タイトル                                                                                |
|-----------------------------------------------------|-----------------------------------------------------------------------------------------|
| asana team list                                     | チームのリスト                                                                          |
| asana team project list                             | チームのプロジェクト一覧                                                                |
| asana team task list                                | チームのタスク一覧                                                                      |
| asana workspace list                                | ワークスペースの一覧                                                                    |
| asana workspace project list                        | ワークスペースのプロジェクト一覧                                                        |
| config auth delete                                  | 既存の認証クレデンシャルの削除                                                          |
| config auth list                                    | すべての認証情報を一覧表示                                                              |
| config feature disable                              | 機能を無効化します.                                                                     |
| config feature enable                               | 機能を有効化します.                                                                     |
| config feature list                                 | 利用可能なオプション機能一覧.                                                           |
| config license install                              | ライセンスキーのインストール                                                            |
| config license list                                 | 利用可能なライセンスキーのリスト                                                        |
| deepl translate text                                | テキストを翻訳する                                                                      |
| dev benchmark local                                 | ローカルファイルシステムにダミーのフォルダ構造を作成します.                             |
| dev benchmark upload                                | アップロードのベンチマーク                                                              |
| dev benchmark uploadlink                            | アップロードテンポラリリンクAPIを使ったシングルファイルのアップロードをベンチマーク.    |
| dev build catalogue                                 | カタログを生成します                                                                    |
| dev build doc                                       | ドキュメントを生成                                                                      |
| dev build info                                      | ビルド情報ファイルを生成                                                                |
| dev build license                                   | LICENSE.txtの生成                                                                       |
| dev build package                                   | ビルドのパッケージ化                                                                    |
| dev build preflight                                 | リリースに向けて必要な事前準備を実施                                                    |
| dev build readme                                    | README.txtの生成                                                                        |
| dev ci artifact up                                  | CI成果物をアップロードします                                                            |
| dev ci auth export                                  | CIビルドのためのデプロイトークンデータの書き出し                                        |
| dev diag endpoint                                   | エンドポイントを一覧                                                                    |
| dev diag throughput                                 | キャプチャログからスループットを評価                                                    |
| dev doc markdown                                    | マークダウンソースからメッセージを生成する                                              |
| dev info                                            | 開発情報                                                                                |
| dev kvs concurrency                                 | KVSエンジンの同時実行テスト                                                             |
| dev kvs dump                                        | KVSデータのダンプ                                                                       |
| dev license issue                                   | ライセンスの発行                                                                        |
| dev lifecycle assets                                | 非推奨資産の削除                                                                        |
| dev lifecycle planchangepath                        | コマンドにパスを変更するプランを追加                                                    |
| dev lifecycle planprune                             | コマンド廃止計画を追加                                                                  |
| dev module list                                     | 依存モジュール一覧                                                                      |
| dev placeholder pathchange                          | パス変更文書生成のためのプレースホルダー・コマンド                                      |
| dev placeholder prune                               | 剪定ワークフローメッセージのプレースホルダ                                              |
| dev release announcement                            | お知らせの更新                                                                          |
| dev release asset                                   | ファイルをリポジトリにコミットする                                                      |
| dev release asseturl                                | リリースのアセットURLを更新                                                             |
| dev release candidate                               | リリース候補を検査します                                                                |
| dev release checkin                                 | 新作りリースをチェック                                                                  |
| dev release doc                                     | リリースドキュメントの作成                                                              |
| dev release publish                                 | リリースを公開します                                                                    |
| dev replay approve                                  | リプレイをテストバンドルとして承認する                                                  |
| dev replay bundle                                   | すべてのリプレイを実行                                                                  |
| dev replay recipe                                   | レシピのリプレイ実行                                                                    |
| dev replay remote                                   | リモートリプレイバンドルの実行                                                          |
| dev spec diff                                       | 2リリース間の仕様を比較します                                                           |
| dev spec doc                                        | 仕様ドキュメントを生成します                                                            |
| dev test echo                                       | テキストのエコー                                                                        |
| dev test license                                    | ライセンスが必要なロジックのテスト                                                      |
| dev test panic                                      | パニック試験                                                                            |
| dev test recipe                                     | レシピのテスト                                                                          |
| dev test resources                                  | バイナリの品質テスト                                                                    |
| dev util anonymise                                  | キャプチャログを匿名化します.                                                           |
| dev util image jpeg                                 | ダミー画像ファイルを作成します                                                          |
| dev util wait                                       | 指定した秒数待機します                                                                  |
| dropbox file account feature                        | Dropboxアカウントの機能一覧                                                             |
| dropbox file account filesystem                     | Dropboxのファイルシステムのバージョンを表示する                                         |
| dropbox file account info                           | Dropboxアカウント情報                                                                   |
| dropbox file compare account                        | 二つのアカウントのファイルを比較します                                                  |
| dropbox file compare local                          | ローカルフォルダとDropboxフォルダの内容を比較します                                     |
| dropbox file copy                                   | ファイルをコピーします                                                                  |
| dropbox file delete                                 | ファイルまたはフォルダは削除します.                                                     |
| dropbox file export doc                             | ドキュメントのエクスポート                                                              |
| dropbox file export url                             | URLからドキュメントをエクスポート                                                       |
| dropbox file import batch url                       | URLからファイルを一括インポートします                                                   |
| dropbox file import url                             | URLからファイルをインポートします                                                       |
| dropbox file info                                   | パスのメタデータを解決                                                                  |
| dropbox file list                                   | ファイルとフォルダを一覧します                                                          |
| dropbox file lock acquire                           | ファイルをロック                                                                        |
| dropbox file lock all release                       | 指定したパスでのすべてのロックを解除する                                                |
| dropbox file lock batch acquire                     | 複数のファイルをロックする                                                              |
| dropbox file lock batch release                     | 複数のロックを解除                                                                      |
| dropbox file lock list                              | 指定したパスの下にあるロックを一覧表示します                                            |
| dropbox file lock release                           | ロックを解除します                                                                      |
| dropbox file merge                                  | フォルダを統合します                                                                    |
| dropbox file move                                   | ファイルを移動します                                                                    |
| dropbox file replication                            | ファイルコンテンツを他のアカウントに複製します                                          |
| dropbox file request create                         | ファイルリクエストを作成します                                                          |
| dropbox file request delete closed                  | このアカウントの全ての閉じられているファイルリクエストを削除します                      |
| dropbox file request delete url                     | ファイルリクエストのURLを指定して削除                                                   |
| dropbox file request list                           | 個人アカウントのファイルリクエストを一覧.                                               |
| dropbox file restore all                            | 指定されたパス以下をリストアします                                                      |
| dropbox file restore ext                            | 特定の拡張子を持つファイルの復元                                                        |
| dropbox file revision download                      | ファイルリビジョンをダウンロードする                                                    |
| dropbox file revision list                          | ファイルリビジョン一覧                                                                  |
| dropbox file revision restore                       | ファイルリビジョンを復元する                                                            |
| dropbox file search content                         | ファイルコンテンツを検索                                                                |
| dropbox file search name                            | ファイル名を検索                                                                        |
| dropbox file share info                             | ファイルの共有情報を取得する                                                            |
| dropbox file sharedfolder info                      | 共有フォルダ情報の取得                                                                  |
| dropbox file sharedfolder leave                     | 共有フォルダーから退出する.                                                             |
| dropbox file sharedfolder list                      | 共有フォルダの一覧                                                                      |
| dropbox file sharedfolder member add                | 共有フォルダへのメンバーの追加                                                          |
| dropbox file sharedfolder member delete             | 共有フォルダからメンバーを削除する                                                      |
| dropbox file sharedfolder member list               | 共有フォルダのメンバーを一覧します                                                      |
| dropbox file sharedfolder mount add                 | 共有フォルダを現在のユーザーのDropboxに追加する                                         |
| dropbox file sharedfolder mount delete              | 現在のユーザーが指定されたフォルダーをアンマウントする.                                 |
| dropbox file sharedfolder mount list                | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示                        |
| dropbox file sharedfolder mount mountable           | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします.               |
| dropbox file sharedfolder share                     | フォルダの共有                                                                          |
| dropbox file sharedfolder unshare                   | フォルダの共有解除                                                                      |
| dropbox file sharedlink create                      | 共有リンクの作成                                                                        |
| dropbox file sharedlink delete                      | 共有リンクを削除します                                                                  |
| dropbox file sharedlink file list                   | 共有リンクのファイルを一覧する                                                          |
| dropbox file sharedlink info                        | 共有リンクの情報取得                                                                    |
| dropbox file sharedlink list                        | 共有リンクの一覧                                                                        |
| dropbox file size                                   | ストレージの利用量                                                                      |
| dropbox file sync down                              | Dropboxと下り方向で同期します                                                           |
| dropbox file sync online                            | オンラインファイルを同期します                                                          |
| dropbox file sync up                                | Dropboxと上り方向で同期します                                                           |
| dropbox file tag add                                | ファイル/フォルダーにタグを追加する                                                     |
| dropbox file tag delete                             | ファイル/フォルダーからタグを削除する                                                   |
| dropbox file tag list                               | パスのタグを一覧                                                                        |
| dropbox file template apply                         | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する                          |
| dropbox file template capture                       | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。                    |
| dropbox file watch                                  | ファイルアクティビティを監視                                                            |
| dropbox paper append                                | 既存のPaperドキュメントの最後にコンテンツを追加する                                     |
| dropbox paper create                                | パスに新しいPaperを作成                                                                 |
| dropbox paper overwrite                             | 既存のPaperドキュメントを上書きする                                                     |
| dropbox paper prepend                               | 既存のPaperドキュメントの先頭にコンテンツを追加する                                     |
| dropbox sign account info                           | Dropbox Signのアカウント情報を表示する                                                  |
| dropbox sign request list                           | 署名依頼リスト                                                                          |
| dropbox sign request signature list                 | リクエストの署名一覧                                                                    |
| dropbox team activity batch user                    | 複数ユーザーのアクティビティを一括取得します                                            |
| dropbox team activity daily event                   | アクティビティーを1日ごとに取得します                                                   |
| dropbox team activity event                         | イベントログ                                                                            |
| dropbox team activity user                          | ユーザーごとのアクティビティ                                                            |
| dropbox team admin group role add                   | グループのメンバーにロールを追加する                                                    |
| dropbox team admin group role delete                | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                        |
| dropbox team admin list                             | メンバーの管理者権限一覧                                                                |
| dropbox team admin role add                         | メンバーに新しいロールを追加する                                                        |
| dropbox team admin role clear                       | メンバーからすべての管理者ロールを削除する                                              |
| dropbox team admin role delete                      | メンバーからロールを削除する                                                            |
| dropbox team admin role list                        | チームの管理者の役割を列挙                                                              |
| dropbox team backup device status                   | Dropbox バックアップ デバイスのステータスが指定期間内に変更された場合                   |
| dropbox team content legacypaper count              | メンバー1人あたりのPaper文書の枚数                                                      |
| dropbox team content legacypaper export             | チームメンバー全員のPaper文書をローカルパスにエクスポート.                              |
| dropbox team content legacypaper list               | チームメンバーのPaper文書リスト出力                                                     |
| dropbox team content member list                    | チームフォルダや共有フォルダのメンバー一覧                                              |
| dropbox team content member size                    | チームフォルダや共有フォルダのメンバー数をカウントする                                  |
| dropbox team content mount list                     | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.  |
| dropbox team content policy list                    | チームフォルダと共有フォルダのポリシー一覧                                              |
| dropbox team device list                            | チーム内全てのデバイス/セッションを一覧します                                           |
| dropbox team device unlink                          | デバイスのセッションを解除します                                                        |
| dropbox team feature                                | チームの機能を出力します                                                                |
| dropbox team filerequest clone                      | ファイルリクエストを入力データに従い複製します                                          |
| dropbox team filerequest list                       | チームないのファイルリクエストを一覧します                                              |
| dropbox team filesystem                             | チームのファイルシステムのバージョンを特定する                                          |
| dropbox team group add                              | グループを作成します                                                                    |
| dropbox team group batch add                        | グループの一括追加                                                                      |
| dropbox team group batch delete                     | グループの削除                                                                          |
| dropbox team group clear externalid                 | グループの外部IDをクリアする                                                            |
| dropbox team group delete                           | グループを削除します                                                                    |
| dropbox team group folder list                      | 各グループのフォルダーを一覧表示                                                        |
| dropbox team group list                             | グループを一覧                                                                          |
| dropbox team group member add                       | メンバーをグループに追加                                                                |
| dropbox team group member batch add                 | グループにメンバーを一括追加                                                            |
| dropbox team group member batch delete              | グループからメンバーを削除                                                              |
| dropbox team group member batch update              | グループからメンバーを追加または削除                                                    |
| dropbox team group member delete                    | メンバーをグループから削除                                                              |
| dropbox team group member list                      | グループに所属するメンバー一覧を取得します                                              |
| dropbox team group rename                           | グループの改名                                                                          |
| dropbox team group update type                      | グループ管理タイプの更新                                                                |
| dropbox team info                                   | チームの情報                                                                            |
| dropbox team insight report teamfoldermember        | チームフォルダーメンバーを報告                                                          |
| dropbox team insight scan                           | チームデータをスキャンして分析                                                          |
| dropbox team insight scanretry                      | 前回のスキャンでエラーがあった場合、スキャンを再試行する                                |
| dropbox team insight summarize                      | 分析のためにチームデータをまとめる                                                      |
| dropbox team legalhold add                          | 新しいリーガル・ホールド・ポリシーを作成する.                                           |
| dropbox team legalhold list                         | 既存のポリシーを取得する                                                                |
| dropbox team legalhold member batch update          | リーガル・ホールド・ポリシーのメンバーリスト更新                                        |
| dropbox team legalhold member list                  | リーガルホールドのメンバーをリストアップ                                                |
| dropbox team legalhold release                      | Idによるリーガルホールドを解除する                                                      |
| dropbox team legalhold revision list                | リーガル・ホールド・ポリシーのリビジョンをリストアップする                              |
| dropbox team legalhold update desc                  | リーガルホールド・ポリシーの説明を更新                                                  |
| dropbox team legalhold update name                  | リーガルホールドポリシーの名称を更新                                                    |
| dropbox team linkedapp list                         | リンク済みアプリを一覧                                                                  |
| dropbox team member batch delete                    | メンバーを削除します                                                                    |
| dropbox team member batch detach                    | Dropbox for teamsのアカウントをBasicアカウントに変更する                                |
| dropbox team member batch invite                    | メンバーを招待します                                                                    |
| dropbox team member batch reinvite                  | 招待済み状態メンバーをチームに再招待します                                              |
| dropbox team member batch suspend                   | メンバーの一括一時停止                                                                  |
| dropbox team member batch unsuspend                 | メンバーの一括停止解除                                                                  |
| dropbox team member clear externalid                | メンバーのexternal_idを初期化します                                                     |
| dropbox team member feature                         | メンバーの機能設定一覧                                                                  |
| dropbox team member file lock all release           | メンバーのパスの下にあるすべてのロックを解除します                                      |
| dropbox team member file lock list                  | パスの下にあるメンバーのロックを一覧表示                                                |
| dropbox team member file lock release               | メンバーとしてパスのロックを解除します                                                  |
| dropbox team member file permdelete                 | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                  |
| dropbox team member folder list                     | 各メンバーのフォルダーを一覧表示                                                        |
| dropbox team member folder replication              | フォルダを他のメンバーの個人フォルダに複製します                                        |
| dropbox team member list                            | チームメンバーの一覧                                                                    |
| dropbox team member quota batch update              | チームメンバーの容量制限を変更                                                          |
| dropbox team member quota list                      | メンバーの容量制限情報を一覧します                                                      |
| dropbox team member quota usage                     | チームメンバーのストレージ利用状況を取得                                                |
| dropbox team member replication                     | チームメンバーのファイルを複製します                                                    |
| dropbox team member suspend                         | メンバーの一時停止処理                                                                  |
| dropbox team member unsuspend                       | メンバーの一時停止を解除する                                                            |
| dropbox team member update batch email              | メンバーのメールアドレス処理                                                            |
| dropbox team member update batch externalid         | チームメンバーのExternal IDを更新します.                                                |
| dropbox team member update batch invisible          | メンバーへのディレクトリ制限を有効にします                                              |
| dropbox team member update batch profile            | メンバーのプロフィール変更                                                              |
| dropbox team member update batch visible            | メンバーへのディレクトリ制限を無効にします                                              |
| dropbox team namespace file list                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| dropbox team namespace file size                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| dropbox team namespace list                         | チーム内すべての名前空間を一覧                                                          |
| dropbox team namespace member list                  | チームフォルダ以下のファイル・フォルダを一覧                                            |
| dropbox team namespace summary                      | チーム・ネームスペースの状態概要を報告する.                                             |
| dropbox team report activity                        | アクティビティ レポート                                                                 |
| dropbox team report devices                         | デバイス レポート空のレポート                                                           |
| dropbox team report membership                      | メンバーシップ レポート                                                                 |
| dropbox team report storage                         | ストレージ レポート                                                                     |
| dropbox team runas file batch copy                  | ファイル/フォルダーをメンバーとして一括コピー                                           |
| dropbox team runas file list                        | メンバーとして実行するファイルやフォルダーの一覧                                        |
| dropbox team runas file sync batch up               | メンバーとして動作する一括同期                                                          |
| dropbox team runas sharedfolder batch leave         | 共有フォルダからメンバーとして一括退出                                                  |
| dropbox team runas sharedfolder batch share         | メンバーのフォルダを一括で共有                                                          |
| dropbox team runas sharedfolder batch unshare       | メンバーのフォルダの共有を一括解除                                                      |
| dropbox team runas sharedfolder isolate             | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| dropbox team runas sharedfolder list                | 共有フォルダーの一覧をメンバーとして実行                                                |
| dropbox team runas sharedfolder member batch add    | メンバーの共有フォルダにメンバーを一括追加                                              |
| dropbox team runas sharedfolder member batch delete | メンバーの共有フォルダからメンバーを一括削除                                            |
| dropbox team runas sharedfolder mount add           | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| dropbox team runas sharedfolder mount delete        | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| dropbox team runas sharedfolder mount list          | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| dropbox team runas sharedfolder mount mountable     | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| dropbox team sharedlink cap expiry                  | チーム内の共有リンクに有効期限の上限を設定                                              |
| dropbox team sharedlink cap visibility              | チーム内の共有リンクに可視性の上限を設定                                                |
| dropbox team sharedlink delete links                | 共有リンクの一括削除                                                                    |
| dropbox team sharedlink delete member               | メンバーの共有リンクをすべて削除                                                        |
| dropbox team sharedlink list                        | 共有リンクの一覧                                                                        |
| dropbox team sharedlink update expiry               | チーム内の公開されている共有リンクについて有効期限を更新します                          |
| dropbox team sharedlink update password             | 共有リンクのパスワードの設定・更新                                                      |
| dropbox team sharedlink update visibility           | 共有リンクの可視性の更新                                                                |
| dropbox team teamfolder add                         | チームフォルダを追加します                                                              |
| dropbox team teamfolder archive                     | チームフォルダのアーカイブ                                                              |
| dropbox team teamfolder batch archive               | 複数のチームフォルダをアーカイブします                                                  |
| dropbox team teamfolder batch permdelete            | 複数のチームフォルダを完全に削除します                                                  |
| dropbox team teamfolder batch replication           | チームフォルダの一括レプリケーション                                                    |
| dropbox team teamfolder file list                   | チームフォルダの一覧                                                                    |
| dropbox team teamfolder file lock all release       | チームフォルダのパスの下にあるすべてのロックを解除する                                  |
| dropbox team teamfolder file lock list              | チームフォルダ内のロックを一覧表示                                                      |
| dropbox team teamfolder file lock release           | チームフォルダ内のパスのロックを解除                                                    |
| dropbox team teamfolder file size                   | チームフォルダのサイズを計算                                                            |
| dropbox team teamfolder list                        | チームフォルダの一覧                                                                    |
| dropbox team teamfolder member add                  | チームフォルダへのユーザー/グループの一括追加                                           |
| dropbox team teamfolder member delete               | チームフォルダからのユーザー/グループの一括削除                                         |
| dropbox team teamfolder member list                 | チームフォルダのメンバー一覧                                                            |
| dropbox team teamfolder partial replication         | 部分的なチームフォルダの他チームへのレプリケーション                                    |
| dropbox team teamfolder permdelete                  | チームフォルダを完全に削除します                                                        |
| dropbox team teamfolder policy list                 | チームフォルダのポリシー一覧                                                            |
| dropbox team teamfolder replication                 | チームフォルダを他のチームに複製します                                                  |
| dropbox team teamfolder sync setting list           | チームフォルダーの同期設定を一覧表示                                                    |
| dropbox team teamfolder sync setting update         | チームフォルダ同期設定の一括更新                                                        |
| figma account info                                  | 現在のユーザー情報を取得する                                                            |
| figma file export all page                          | チーム配下のすべてのファイル/ページをエクスポートする                                   |
| figma file export frame                             | Figmaファイルの全フレームを書き出す                                                     |
| figma file export node                              | Figmaドキュメント・ノードの書き出し                                                     |
| figma file export page                              | Figmaファイルの全ページを書き出す                                                       |
| figma file info                                     | figmaファイルの情報を表示する                                                           |
| figma file list                                     | Figmaプロジェクト内のファイル一覧                                                       |
| figma project list                                  | チームのプロジェクト一覧                                                                |
| github content get                                  | レポジトリのコンテンツメタデータを取得します.                                           |
| github content put                                  | レポジトリに小さなテキストコンテンツを格納します                                        |
| github issue list                                   | 公開・プライベートGitHubレポジトリの課題一覧                                            |
| github profile                                      | 認証したユーザーの情報を取得                                                            |
| github release asset download                       | アセットをダウンロードします                                                            |
| github release asset list                           | GitHubリリースの成果物一覧                                                              |
| github release asset upload                         | GitHub リリースへ成果物をアップロードします                                             |
| github release draft                                | リリースの下書きを作成                                                                  |
| github release list                                 | リリースの一覧                                                                          |
| github tag create                                   | レポジトリにタグを作成します                                                            |
| license                                             | ライセンス情報を表示します                                                              |
| local file template apply                           | ファイル/フォルダー構造のテンプレートをローカルパスに適用する                           |
| local file template capture                         | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む                       |
| log api job                                         | ジョブIDで指定されたジョブのAPIログの統計情報を表示する                                 |
| log api name                                        | ジョブ名で指定されたジョブのAPIログの統計情報を表示する                                 |
| log cat curl                                        | キャプチャログを `curl` サンプルとしてフォーマットする                                  |
| log cat job                                         | 指定したジョブIDのログを取得する                                                        |
| log cat kind                                        | 指定種別のログを結合して出力します                                                      |
| log cat last                                        | 最後のジョブのログファイルを出力.                                                       |
| log job archive                                     | ジョブのアーカイブ                                                                      |
| log job delete                                      | 古いジョブ履歴の削除                                                                    |
| log job list                                        | ジョブ履歴の表示                                                                        |
| slack conversation history                          | 会話履歴                                                                                |
| slack conversation list                             | チャネルの一覧                                                                          |
| util archive unzip                                  | ZIPアーカイブファイルを解凍する                                                         |
| util archive zip                                    | 対象ファイルをZIPアーカイブに圧縮する                                                   |
| util cert selfsigned                                | 自己署名証明書と鍵の生成                                                                |
| util database exec                                  | SQLite3データベースファイルへのクエリ実行                                               |
| util database query                                 | SQLite3データベースへの問い合わせ                                                       |
| util date today                                     | 現在の日付を表示                                                                        |
| util datetime now                                   | 現在の日時を表示                                                                        |
| util decode base32                                  | Base32 (RFC 4648) 形式からテキストをデコードします                                      |
| util decode base64                                  | Base64 (RFC 4648) フォーマットからテキストをデコードします                              |
| util desktop open                                   | デフォルトのアプリケーションでファイルやフォルダを開く                                  |
| util encode base32                                  | テキストをBase32(RFC 4648)形式にエンコード                                              |
| util encode base64                                  | テキストをBase64(RFC 4648)形式にエンコード                                              |
| util feed json                                      | URLからフィードを読み込み、コンテンツをJSONとして出力する。                             |
| util file hash                                      | ファイルダイジェストの表示                                                              |
| util git clone                                      | git リポジトリをクローン                                                                |
| util image exif                                     | 画像ファイルのEXIFメタデータを表示                                                      |
| util image placeholder                              | プレースホルダー画像の作成                                                              |
| util json query                                     | JSONデータを問い合わせる                                                                |
| util net download                                   | ファイルをダウンロードする                                                              |
| util qrcode create                                  | QRコード画像ファイルの作成                                                              |
| util qrcode wifi                                    | WIFI設定用のQRコードを生成                                                              |
| util release install                                | watermint toolboxをダウンロードし、パスにインストールします。                           |
| util table format xlsx                              | xlsxファイルをテキストに整形する                                                        |
| util text case down                                 | 小文字のテキストを表示する                                                              |
| util text case up                                   | 大文字のテキストを表示する                                                              |
| util text encoding from                             | 指定されたエンコーディングからUTF-8テキストファイルに変換します.                        |
| util text encoding to                               | UTF-8テキストファイルから指定されたエンコーディングに変換する.                          |
| util text nlp english entity                        | 英文をエンティティに分割する                                                            |
| util text nlp english sentence                      | 英文を文章に分割する                                                                    |
| util text nlp english token                         | 英文をトークンに分割する                                                                |
| util text nlp japanese token                        | 日本語テキストのトークン化                                                              |
| util text nlp japanese wakati                       | 分かち書き(日本語テキストのトークン化)                                                  |
| util tidy move dispatch                             | ファイルを整理                                                                          |
| util tidy move simple                               | ローカルファイルをアーカイブします                                                      |
| util tidy pack remote                               | リモートフォルダをZIPファイルにパッケージする                                           |
| util time now                                       | 現在の時刻を表示                                                                        |
| util unixtime format                                | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット                  |
| util unixtime now                                   | UNIX時間で現在の時刻を表示する                                                          |
| util uuid timestamp                                 | UUIDタイムスタンプの解析                                                                |
| util uuid ulid                                      | ULID（Universally Unique Lexicographically Sortable Identifier）を生成する。            |
| util uuid v4                                        | UUID v4（ランダムUUID）の生成                                                           |
| util uuid v7                                        | UUID v7 の生成                                                                          |
| util uuid version                                   | UUIDのバージョンとバリアントの解析                                                      |
| util xlsx create                                    | 空のスプレッドシートを作成する                                                          |
| util xlsx sheet export                              | xlsxファイルからデータをエクスポート                                                    |
| util xlsx sheet import                              | データをxlsxファイルにインポート                                                        |
| util xlsx sheet list                                | xlsxファイルのシート一覧                                                                |
| version                                             | バージョン情報                                                                          |



