---
layout: release
title: リリースの変更点 128
lang: ja
---

# `リリース 128` から `リリース 129` までの変更点

# 追加されたコマンド


| コマンド                                            | タイトル                                                                                |
|-----------------------------------------------------|-----------------------------------------------------------------------------------------|
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
| dropbox file revision download                      | ファイルリビジョンをダウンロードする                                                    |
| dropbox file revision list                          | ファイルリビジョン一覧                                                                  |
| dropbox file revision restore                       | ファイルリビジョンを復元する                                                            |
| dropbox file search content                         | ファイルコンテンツを検索                                                                |
| dropbox file search name                            | ファイル名を検索                                                                        |
| dropbox file share info                             | ファイルの共有情報を取得する                                                            |
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
| local file template apply                           | ファイル/フォルダー構造のテンプレートをローカルパスに適用する                           |
| local file template capture                         | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む                       |



# 削除されたコマンド


| コマンド                                    | タイトル                                                                                |
|---------------------------------------------|-----------------------------------------------------------------------------------------|
| file compare account                        | 二つのアカウントのファイルを比較します                                                  |
| file compare local                          | ローカルフォルダとDropboxフォルダの内容を比較します                                     |
| file copy                                   | ファイルをコピーします                                                                  |
| file delete                                 | ファイルまたはフォルダは削除します.                                                     |
| file export doc                             | ドキュメントのエクスポート                                                              |
| file export url                             | URLからドキュメントをエクスポート                                                       |
| file import batch url                       | URLからファイルを一括インポートします                                                   |
| file import url                             | URLからファイルをインポートします                                                       |
| file info                                   | パスのメタデータを解決                                                                  |
| file list                                   | ファイルとフォルダを一覧します                                                          |
| file lock acquire                           | ファイルをロック                                                                        |
| file lock all release                       | 指定したパスでのすべてのロックを解除する                                                |
| file lock batch acquire                     | 複数のファイルをロックする                                                              |
| file lock batch release                     | 複数のロックを解除                                                                      |
| file lock list                              | 指定したパスの下にあるロックを一覧表示します                                            |
| file lock release                           | ロックを解除します                                                                      |
| file merge                                  | フォルダを統合します                                                                    |
| file move                                   | ファイルを移動します                                                                    |
| file paper append                           | 既存のPaperドキュメントの最後にコンテンツを追加する                                     |
| file paper create                           | パスに新しいPaperを作成                                                                 |
| file paper overwrite                        | 既存のPaperドキュメントを上書きする                                                     |
| file paper prepend                          | 既存のPaperドキュメントの先頭にコンテンツを追加する                                     |
| file replication                            | ファイルコンテンツを他のアカウントに複製します                                          |
| file restore all                            | 指定されたパス以下をリストアします                                                      |
| file revision download                      | ファイルリビジョンをダウンロードする                                                    |
| file revision list                          | ファイルリビジョン一覧                                                                  |
| file revision restore                       | ファイルリビジョンを復元する                                                            |
| file search content                         | ファイルコンテンツを検索                                                                |
| file search name                            | ファイル名を検索                                                                        |
| file share info                             | ファイルの共有情報を取得する                                                            |
| file size                                   | ストレージの利用量                                                                      |
| file sync down                              | Dropboxと下り方向で同期します                                                           |
| file sync online                            | オンラインファイルを同期します                                                          |
| file sync up                                | Dropboxと上り方向で同期します                                                           |
| file tag add                                | ファイル/フォルダーにタグを追加する                                                     |
| file tag delete                             | ファイル/フォルダーからタグを削除する                                                   |
| file tag list                               | パスのタグを一覧                                                                        |
| file template apply local                   | ファイル/フォルダー構造のテンプレートをローカルパスに適用する                           |
| file template apply remote                  | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する                          |
| file template capture local                 | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む                       |
| file template capture remote                | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。                    |
| file watch                                  | ファイルアクティビティを監視                                                            |
| filerequest create                          | ファイルリクエストを作成します                                                          |
| filerequest delete closed                   | このアカウントの全ての閉じられているファイルリクエストを削除します                      |
| filerequest delete url                      | ファイルリクエストのURLを指定して削除                                                   |
| filerequest list                            | 個人アカウントのファイルリクエストを一覧.                                               |
| group add                                   | グループを作成します                                                                    |
| group batch add                             | グループの一括追加                                                                      |
| group batch delete                          | グループの削除                                                                          |
| group clear externalid                      | グループの外部IDをクリアする                                                            |
| group delete                                | グループを削除します                                                                    |
| group folder list                           | 各グループのフォルダーを一覧表示                                                        |
| group list                                  | グループを一覧                                                                          |
| group member add                            | メンバーをグループに追加                                                                |
| group member batch add                      | グループにメンバーを一括追加                                                            |
| group member batch delete                   | グループからメンバーを削除                                                              |
| group member batch update                   | グループからメンバーを追加または削除                                                    |
| group member delete                         | メンバーをグループから削除                                                              |
| group member list                           | グループに所属するメンバー一覧を取得します                                              |
| group rename                                | グループの改名                                                                          |
| group update type                           | グループ管理タイプの更新                                                                |
| member batch suspend                        | メンバーの一括一時停止                                                                  |
| member batch unsuspend                      | メンバーの一括停止解除                                                                  |
| member clear externalid                     | メンバーのexternal_idを初期化します                                                     |
| member delete                               | メンバーを削除します                                                                    |
| member detach                               | Dropbox for teamsのアカウントをBasicアカウントに変更する                                |
| member feature                              | メンバーの機能設定一覧                                                                  |
| member file lock all release                | メンバーのパスの下にあるすべてのロックを解除します                                      |
| member file lock list                       | パスの下にあるメンバーのロックを一覧表示                                                |
| member file lock release                    | メンバーとしてパスのロックを解除します                                                  |
| member file permdelete                      | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                  |
| member folder list                          | 各メンバーのフォルダーを一覧表示                                                        |
| member folder replication                   | フォルダを他のメンバーの個人フォルダに複製します                                        |
| member invite                               | メンバーを招待します                                                                    |
| member list                                 | チームメンバーの一覧                                                                    |
| member quota list                           | メンバーの容量制限情報を一覧します                                                      |
| member quota update                         | チームメンバーの容量制限を変更                                                          |
| member quota usage                          | チームメンバーのストレージ利用状況を取得                                                |
| member reinvite                             | 招待済み状態メンバーをチームに再招待します                                              |
| member replication                          | チームメンバーのファイルを複製します                                                    |
| member suspend                              | メンバーの一時停止処理                                                                  |
| member unsuspend                            | メンバーの一時停止を解除する                                                            |
| member update email                         | メンバーのメールアドレス処理                                                            |
| member update externalid                    | チームメンバーのExternal IDを更新します.                                                |
| member update invisible                     | メンバーへのディレクトリ制限を有効にします                                              |
| member update profile                       | メンバーのプロフィール変更                                                              |
| member update visible                       | メンバーへのディレクトリ制限を無効にします                                              |
| sharedfolder leave                          | 共有フォルダーから退出する.                                                             |
| sharedfolder list                           | 共有フォルダの一覧                                                                      |
| sharedfolder member add                     | 共有フォルダへのメンバーの追加                                                          |
| sharedfolder member delete                  | 共有フォルダからメンバーを削除する                                                      |
| sharedfolder member list                    | 共有フォルダのメンバーを一覧します                                                      |
| sharedfolder mount add                      | 共有フォルダを現在のユーザーのDropboxに追加する                                         |
| sharedfolder mount delete                   | 現在のユーザーが指定されたフォルダーをアンマウントする.                                 |
| sharedfolder mount list                     | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示                        |
| sharedfolder mount mountable                | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします.               |
| sharedfolder share                          | フォルダの共有                                                                          |
| sharedfolder unshare                        | フォルダの共有解除                                                                      |
| sharedlink create                           | 共有リンクの作成                                                                        |
| sharedlink delete                           | 共有リンクを削除します                                                                  |
| sharedlink file list                        | 共有リンクのファイルを一覧する                                                          |
| sharedlink info                             | 共有リンクの情報取得                                                                    |
| sharedlink list                             | 共有リンクの一覧                                                                        |
| team activity batch user                    | 複数ユーザーのアクティビティを一括取得します                                            |
| team activity daily event                   | アクティビティーを1日ごとに取得します                                                   |
| team activity event                         | イベントログ                                                                            |
| team activity user                          | ユーザーごとのアクティビティ                                                            |
| team admin group role add                   | グループのメンバーにロールを追加する                                                    |
| team admin group role delete                | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                        |
| team admin list                             | メンバーの管理者権限一覧                                                                |
| team admin role add                         | メンバーに新しいロールを追加する                                                        |
| team admin role clear                       | メンバーからすべての管理者ロールを削除する                                              |
| team admin role delete                      | メンバーからロールを削除する                                                            |
| team admin role list                        | チームの管理者の役割を列挙                                                              |
| team content legacypaper count              | メンバー1人あたりのPaper文書の枚数                                                      |
| team content legacypaper export             | チームメンバー全員のPaper文書をローカルパスにエクスポート.                              |
| team content legacypaper list               | チームメンバーのPaper文書リスト出力                                                     |
| team content member list                    | チームフォルダや共有フォルダのメンバー一覧                                              |
| team content member size                    | チームフォルダや共有フォルダのメンバー数をカウントする                                  |
| team content mount list                     | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.  |
| team content policy list                    | チームフォルダと共有フォルダのポリシー一覧                                              |
| team device list                            | チーム内全てのデバイス/セッションを一覧します                                           |
| team device unlink                          | デバイスのセッションを解除します                                                        |
| team feature                                | チームの機能を出力します                                                                |
| team filerequest clone                      | ファイルリクエストを入力データに従い複製します                                          |
| team filerequest list                       | チームないのファイルリクエストを一覧します                                              |
| team filesystem                             | チームのファイルシステムのバージョンを特定する                                          |
| team info                                   | チームの情報                                                                            |
| team insight scan                           | チーム全体の情報をスキャンする                                                          |
| team insight summarize                      | スキャンしたチーム情報をまとめる                                                        |
| team legalhold add                          | 新しいリーガル・ホールド・ポリシーを作成する.                                           |
| team legalhold list                         | 既存のポリシーを取得する                                                                |
| team legalhold member batch update          | リーガル・ホールド・ポリシーのメンバーリスト更新                                        |
| team legalhold member list                  | リーガルホールドのメンバーをリストアップ                                                |
| team legalhold release                      | Idによるリーガルホールドを解除する                                                      |
| team legalhold revision list                | リーガル・ホールド・ポリシーのリビジョンをリストアップする                              |
| team legalhold update desc                  | リーガルホールド・ポリシーの説明を更新                                                  |
| team legalhold update name                  | リーガルホールドポリシーの名称を更新                                                    |
| team linkedapp list                         | リンク済みアプリを一覧                                                                  |
| team namespace file list                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| team namespace file size                    | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| team namespace list                         | チーム内すべての名前空間を一覧                                                          |
| team namespace member list                  | チームフォルダ以下のファイル・フォルダを一覧                                            |
| team namespace summary                      | チーム・ネームスペースの状態概要を報告する.                                             |
| team report activity                        | アクティビティ レポート                                                                 |
| team report devices                         | デバイス レポート空のレポート                                                           |
| team report membership                      | メンバーシップ レポート                                                                 |
| team report storage                         | ストレージ レポート                                                                     |
| team runas file batch copy                  | ファイル/フォルダーをメンバーとして一括コピー                                           |
| team runas file list                        | メンバーとして実行するファイルやフォルダーの一覧                                        |
| team runas file sync batch up               | メンバーとして動作する一括同期                                                          |
| team runas sharedfolder batch leave         | 共有フォルダからメンバーとして一括退出                                                  |
| team runas sharedfolder batch share         | メンバーのフォルダを一括で共有                                                          |
| team runas sharedfolder batch unshare       | メンバーのフォルダの共有を一括解除                                                      |
| team runas sharedfolder isolate             | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| team runas sharedfolder list                | 共有フォルダーの一覧をメンバーとして実行                                                |
| team runas sharedfolder member batch add    | メンバーの共有フォルダにメンバーを一括追加                                              |
| team runas sharedfolder member batch delete | メンバーの共有フォルダからメンバーを一括削除                                            |
| team runas sharedfolder mount add           | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| team runas sharedfolder mount delete        | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| team runas sharedfolder mount list          | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| team runas sharedfolder mount mountable     | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| team sharedlink cap expiry                  | チーム内の共有リンクに有効期限の上限を設定                                              |
| team sharedlink cap visibility              | チーム内の共有リンクに可視性の上限を設定                                                |
| team sharedlink delete links                | 共有リンクの一括削除                                                                    |
| team sharedlink delete member               | メンバーの共有リンクをすべて削除                                                        |
| team sharedlink list                        | 共有リンクの一覧                                                                        |
| team sharedlink update expiry               | チーム内の公開されている共有リンクについて有効期限を更新します                          |
| team sharedlink update password             | 共有リンクのパスワードの設定・更新                                                      |
| team sharedlink update visibility           | 共有リンクの可視性の更新                                                                |
| teamfolder add                              | チームフォルダを追加します                                                              |
| teamfolder archive                          | チームフォルダのアーカイブ                                                              |
| teamfolder batch archive                    | 複数のチームフォルダをアーカイブします                                                  |
| teamfolder batch permdelete                 | 複数のチームフォルダを完全に削除します                                                  |
| teamfolder batch replication                | チームフォルダの一括レプリケーション                                                    |
| teamfolder file list                        | チームフォルダの一覧                                                                    |
| teamfolder file lock all release            | チームフォルダのパスの下にあるすべてのロックを解除する                                  |
| teamfolder file lock list                   | チームフォルダ内のロックを一覧表示                                                      |
| teamfolder file lock release                | チームフォルダ内のパスのロックを解除                                                    |
| teamfolder file size                        | チームフォルダのサイズを計算                                                            |
| teamfolder list                             | チームフォルダの一覧                                                                    |
| teamfolder member add                       | チームフォルダへのユーザー/グループの一括追加                                           |
| teamfolder member delete                    | チームフォルダからのユーザー/グループの一括削除                                         |
| teamfolder member list                      | チームフォルダのメンバー一覧                                                            |
| teamfolder partial replication              | 部分的なチームフォルダの他チームへのレプリケーション                                    |
| teamfolder permdelete                       | チームフォルダを完全に削除します                                                        |
| teamfolder policy list                      | チームフォルダのポリシー一覧                                                            |
| teamfolder replication                      | チームフォルダを他のチームに複製します                                                  |
| teamfolder sync setting list                | チームフォルダーの同期設定を一覧表示                                                    |
| teamfolder sync setting update              | チームフォルダ同期設定の一括更新                                                        |



# コマンド仕様の変更: `dev lifecycle planchangepath`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev lifecycle planchangepath",
  	CliArgs: strings.Join({
  		"-announce-url URL -compatibility-file /LOCAL/PATH/TO/compat.json",
  		" -",
+ 		"message-file /LOCAL/PATH/TO/messages.json -",
  		`date "2020-04-01 17:58:38" -current-path RECIPE -former-path REC`,
  		"IPE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 9 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AnnounceUrl", Desc: "アナウンスURL", TypeName: "string"},
  		&{Name: "Compact", Desc: "コンパクトな出力を生成する", Default: "false", TypeName: "bool", ...},
  		&{Name: "CompatibilityFile", Desc: "互換性ファイル", Default: "catalogue/catalogue_compatibility.json", TypeName: "essentials.model.mo_path.file_system_path_impl", ...},
+ 		&{
+ 			Name:     "CurrentBase",
+ 			Desc:     "現在のレシピのベースパス",
+ 			Default:  "citron",
+ 			TypeName: "string",
+ 		},
  		&{Name: "CurrentPath", Desc: "現在のCLIパス", TypeName: "string"},
  		&{Name: "Date", Desc: "発効日", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
+ 		&{
+ 			Name:     "FormerBase",
+ 			Desc:     "旧レシピのベースパス",
+ 			Default:  "recipe",
+ 			TypeName: "string",
+ 		},
  		&{Name: "FormerPath", Desc: "旧CLIパス", TypeName: "string"},
+ 		&{
+ 			Name:     "MessageFile",
+ 			Desc:     "メッセージファイルのパス",
+ 			Default:  "resources/messages/en/messages.json",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]any{"shouldExist": bool(false)},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
