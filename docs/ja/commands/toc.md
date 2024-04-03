---
layout: page
title: コマンド
lang: ja
---

# コマンド

## Dropbox (個人アカウント)

| コマンド                                                                                                                   | 説明                                                                      |
|----------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [dropbox file account feature]({{ site.baseurl }}/ja/commands/dropbox-file-account-feature.html)                           | Dropboxアカウントの機能一覧                                               |
| [dropbox file account filesystem]({{ site.baseurl }}/ja/commands/dropbox-file-account-filesystem.html)                     | Dropboxのファイルシステムのバージョンを表示する                           |
| [dropbox file account info]({{ site.baseurl }}/ja/commands/dropbox-file-account-info.html)                                 | Dropboxアカウント情報                                                     |
| [dropbox file compare account]({{ site.baseurl }}/ja/commands/dropbox-file-compare-account.html)                           | 二つのアカウントのファイルを比較します                                    |
| [dropbox file compare local]({{ site.baseurl }}/ja/commands/dropbox-file-compare-local.html)                               | ローカルフォルダとDropboxフォルダの内容を比較します                       |
| [dropbox file copy]({{ site.baseurl }}/ja/commands/dropbox-file-copy.html)                                                 | ファイルをコピーします                                                    |
| [dropbox file delete]({{ site.baseurl }}/ja/commands/dropbox-file-delete.html)                                             | ファイルまたはフォルダは削除します.                                       |
| [dropbox file export doc]({{ site.baseurl }}/ja/commands/dropbox-file-export-doc.html)                                     | ドキュメントのエクスポート                                                |
| [dropbox file export url]({{ site.baseurl }}/ja/commands/dropbox-file-export-url.html)                                     | URLからドキュメントをエクスポート                                         |
| [dropbox file import batch url]({{ site.baseurl }}/ja/commands/dropbox-file-import-batch-url.html)                         | URLからファイルを一括インポートします                                     |
| [dropbox file import url]({{ site.baseurl }}/ja/commands/dropbox-file-import-url.html)                                     | URLからファイルをインポートします                                         |
| [dropbox file info]({{ site.baseurl }}/ja/commands/dropbox-file-info.html)                                                 | パスのメタデータを解決                                                    |
| [dropbox file list]({{ site.baseurl }}/ja/commands/dropbox-file-list.html)                                                 | ファイルとフォルダを一覧します                                            |
| [dropbox file lock acquire]({{ site.baseurl }}/ja/commands/dropbox-file-lock-acquire.html)                                 | ファイルをロック                                                          |
| [dropbox file lock all release]({{ site.baseurl }}/ja/commands/dropbox-file-lock-all-release.html)                         | 指定したパスでのすべてのロックを解除する                                  |
| [dropbox file lock batch acquire]({{ site.baseurl }}/ja/commands/dropbox-file-lock-batch-acquire.html)                     | 複数のファイルをロックする                                                |
| [dropbox file lock batch release]({{ site.baseurl }}/ja/commands/dropbox-file-lock-batch-release.html)                     | 複数のロックを解除                                                        |
| [dropbox file lock list]({{ site.baseurl }}/ja/commands/dropbox-file-lock-list.html)                                       | 指定したパスの下にあるロックを一覧表示します                              |
| [dropbox file lock release]({{ site.baseurl }}/ja/commands/dropbox-file-lock-release.html)                                 | ロックを解除します                                                        |
| [dropbox file merge]({{ site.baseurl }}/ja/commands/dropbox-file-merge.html)                                               | フォルダを統合します                                                      |
| [dropbox file move]({{ site.baseurl }}/ja/commands/dropbox-file-move.html)                                                 | ファイルを移動します                                                      |
| [dropbox file replication]({{ site.baseurl }}/ja/commands/dropbox-file-replication.html)                                   | ファイルコンテンツを他のアカウントに複製します                            |
| [dropbox file request create]({{ site.baseurl }}/ja/commands/dropbox-file-request-create.html)                             | ファイルリクエストを作成します                                            |
| [dropbox file request delete closed]({{ site.baseurl }}/ja/commands/dropbox-file-request-delete-closed.html)               | このアカウントの全ての閉じられているファイルリクエストを削除します        |
| [dropbox file request delete url]({{ site.baseurl }}/ja/commands/dropbox-file-request-delete-url.html)                     | ファイルリクエストのURLを指定して削除                                     |
| [dropbox file request list]({{ site.baseurl }}/ja/commands/dropbox-file-request-list.html)                                 | 個人アカウントのファイルリクエストを一覧.                                 |
| [dropbox file restore all]({{ site.baseurl }}/ja/commands/dropbox-file-restore-all.html)                                   | 指定されたパス以下をリストアします                                        |
| [dropbox file revision download]({{ site.baseurl }}/ja/commands/dropbox-file-revision-download.html)                       | ファイルリビジョンをダウンロードする                                      |
| [dropbox file revision list]({{ site.baseurl }}/ja/commands/dropbox-file-revision-list.html)                               | ファイルリビジョン一覧                                                    |
| [dropbox file revision restore]({{ site.baseurl }}/ja/commands/dropbox-file-revision-restore.html)                         | ファイルリビジョンを復元する                                              |
| [dropbox file search content]({{ site.baseurl }}/ja/commands/dropbox-file-search-content.html)                             | ファイルコンテンツを検索                                                  |
| [dropbox file search name]({{ site.baseurl }}/ja/commands/dropbox-file-search-name.html)                                   | ファイル名を検索                                                          |
| [dropbox file share info]({{ site.baseurl }}/ja/commands/dropbox-file-share-info.html)                                     | ファイルの共有情報を取得する                                              |
| [dropbox file sharedfolder leave]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-leave.html)                     | 共有フォルダーから退出する.                                               |
| [dropbox file sharedfolder list]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-list.html)                       | 共有フォルダの一覧                                                        |
| [dropbox file sharedfolder member add]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-member-add.html)           | 共有フォルダへのメンバーの追加                                            |
| [dropbox file sharedfolder member delete]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-member-delete.html)     | 共有フォルダからメンバーを削除する                                        |
| [dropbox file sharedfolder member list]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-member-list.html)         | 共有フォルダのメンバーを一覧します                                        |
| [dropbox file sharedfolder mount add]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-mount-add.html)             | 共有フォルダを現在のユーザーのDropboxに追加する                           |
| [dropbox file sharedfolder mount delete]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-mount-delete.html)       | 現在のユーザーが指定されたフォルダーをアンマウントする.                   |
| [dropbox file sharedfolder mount list]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-mount-list.html)           | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示          |
| [dropbox file sharedfolder mount mountable]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-mount-mountable.html) | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします. |
| [dropbox file sharedfolder share]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-share.html)                     | フォルダの共有                                                            |
| [dropbox file sharedfolder unshare]({{ site.baseurl }}/ja/commands/dropbox-file-sharedfolder-unshare.html)                 | フォルダの共有解除                                                        |
| [dropbox file sharedlink create]({{ site.baseurl }}/ja/commands/dropbox-file-sharedlink-create.html)                       | 共有リンクの作成                                                          |
| [dropbox file sharedlink delete]({{ site.baseurl }}/ja/commands/dropbox-file-sharedlink-delete.html)                       | 共有リンクを削除します                                                    |
| [dropbox file sharedlink file list]({{ site.baseurl }}/ja/commands/dropbox-file-sharedlink-file-list.html)                 | 共有リンクのファイルを一覧する                                            |
| [dropbox file sharedlink info]({{ site.baseurl }}/ja/commands/dropbox-file-sharedlink-info.html)                           | 共有リンクの情報取得                                                      |
| [dropbox file sharedlink list]({{ site.baseurl }}/ja/commands/dropbox-file-sharedlink-list.html)                           | 共有リンクの一覧                                                          |
| [dropbox file size]({{ site.baseurl }}/ja/commands/dropbox-file-size.html)                                                 | ストレージの利用量                                                        |
| [dropbox file sync down]({{ site.baseurl }}/ja/commands/dropbox-file-sync-down.html)                                       | Dropboxと下り方向で同期します                                             |
| [dropbox file sync online]({{ site.baseurl }}/ja/commands/dropbox-file-sync-online.html)                                   | オンラインファイルを同期します                                            |
| [dropbox file sync up]({{ site.baseurl }}/ja/commands/dropbox-file-sync-up.html)                                           | Dropboxと上り方向で同期します                                             |
| [dropbox file tag add]({{ site.baseurl }}/ja/commands/dropbox-file-tag-add.html)                                           | ファイル/フォルダーにタグを追加する                                       |
| [dropbox file tag delete]({{ site.baseurl }}/ja/commands/dropbox-file-tag-delete.html)                                     | ファイル/フォルダーからタグを削除する                                     |
| [dropbox file tag list]({{ site.baseurl }}/ja/commands/dropbox-file-tag-list.html)                                         | パスのタグを一覧                                                          |
| [dropbox file template apply]({{ site.baseurl }}/ja/commands/dropbox-file-template-apply.html)                             | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する            |
| [dropbox file template capture]({{ site.baseurl }}/ja/commands/dropbox-file-template-capture.html)                         | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。      |
| [dropbox file watch]({{ site.baseurl }}/ja/commands/dropbox-file-watch.html)                                               | ファイルアクティビティを監視                                              |
| [dropbox paper append]({{ site.baseurl }}/ja/commands/dropbox-paper-append.html)                                           | 既存のPaperドキュメントの最後にコンテンツを追加する                       |
| [dropbox paper create]({{ site.baseurl }}/ja/commands/dropbox-paper-create.html)                                           | パスに新しいPaperを作成                                                   |
| [dropbox paper overwrite]({{ site.baseurl }}/ja/commands/dropbox-paper-overwrite.html)                                     | 既存のPaperドキュメントを上書きする                                       |
| [dropbox paper prepend]({{ site.baseurl }}/ja/commands/dropbox-paper-prepend.html)                                         | 既存のPaperドキュメントの先頭にコンテンツを追加する                       |
| [log job ship]({{ site.baseurl }}/ja/commands/log-job-ship.html)                                                           | ログの転送先Dropboxパス                                                   |
| [teamspace file list]({{ site.baseurl }}/ja/commands/teamspace-file-list.html)                                             | チームスペースにあるファイルやフォルダーを一覧表示                        |
| [util monitor client]({{ site.baseurl }}/ja/commands/util-monitor-client.html)                                             | デバイスモニタークライアントを起動する                                    |
| [util tidy pack remote]({{ site.baseurl }}/ja/commands/util-tidy-pack-remote.html)                                         | リモートフォルダをZIPファイルにパッケージする                             |

## チーム向けDropbox

| コマンド                                                                                                                                       | 説明                                                                                    |
|------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| [dropbox team activity batch user]({{ site.baseurl }}/ja/commands/dropbox-team-activity-batch-user.html)                                       | 複数ユーザーのアクティビティを一括取得します                                            |
| [dropbox team activity daily event]({{ site.baseurl }}/ja/commands/dropbox-team-activity-daily-event.html)                                     | アクティビティーを1日ごとに取得します                                                   |
| [dropbox team activity event]({{ site.baseurl }}/ja/commands/dropbox-team-activity-event.html)                                                 | イベントログ                                                                            |
| [dropbox team activity user]({{ site.baseurl }}/ja/commands/dropbox-team-activity-user.html)                                                   | ユーザーごとのアクティビティ                                                            |
| [dropbox team admin group role add]({{ site.baseurl }}/ja/commands/dropbox-team-admin-group-role-add.html)                                     | グループのメンバーにロールを追加する                                                    |
| [dropbox team admin group role delete]({{ site.baseurl }}/ja/commands/dropbox-team-admin-group-role-delete.html)                               | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                        |
| [dropbox team admin list]({{ site.baseurl }}/ja/commands/dropbox-team-admin-list.html)                                                         | メンバーの管理者権限一覧                                                                |
| [dropbox team admin role add]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-add.html)                                                 | メンバーに新しいロールを追加する                                                        |
| [dropbox team admin role clear]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-clear.html)                                             | メンバーからすべての管理者ロールを削除する                                              |
| [dropbox team admin role delete]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-delete.html)                                           | メンバーからロールを削除する                                                            |
| [dropbox team admin role list]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-list.html)                                               | チームの管理者の役割を列挙                                                              |
| [dropbox team backup device status]({{ site.baseurl }}/ja/commands/dropbox-team-backup-device-status.html)                                     | Dropbox バックアップ デバイスのステータスが指定期間内に変更された場合                   |
| [dropbox team content legacypaper count]({{ site.baseurl }}/ja/commands/dropbox-team-content-legacypaper-count.html)                           | メンバー1人あたりのPaper文書の枚数                                                      |
| [dropbox team content legacypaper export]({{ site.baseurl }}/ja/commands/dropbox-team-content-legacypaper-export.html)                         | チームメンバー全員のPaper文書をローカルパスにエクスポート.                              |
| [dropbox team content legacypaper list]({{ site.baseurl }}/ja/commands/dropbox-team-content-legacypaper-list.html)                             | チームメンバーのPaper文書リスト出力                                                     |
| [dropbox team content member list]({{ site.baseurl }}/ja/commands/dropbox-team-content-member-list.html)                                       | チームフォルダや共有フォルダのメンバー一覧                                              |
| [dropbox team content member size]({{ site.baseurl }}/ja/commands/dropbox-team-content-member-size.html)                                       | チームフォルダや共有フォルダのメンバー数をカウントする                                  |
| [dropbox team content mount list]({{ site.baseurl }}/ja/commands/dropbox-team-content-mount-list.html)                                         | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.  |
| [dropbox team content policy list]({{ site.baseurl }}/ja/commands/dropbox-team-content-policy-list.html)                                       | チームフォルダと共有フォルダのポリシー一覧                                              |
| [dropbox team device list]({{ site.baseurl }}/ja/commands/dropbox-team-device-list.html)                                                       | チーム内全てのデバイス/セッションを一覧します                                           |
| [dropbox team device unlink]({{ site.baseurl }}/ja/commands/dropbox-team-device-unlink.html)                                                   | デバイスのセッションを解除します                                                        |
| [dropbox team feature]({{ site.baseurl }}/ja/commands/dropbox-team-feature.html)                                                               | チームの機能を出力します                                                                |
| [dropbox team filerequest list]({{ site.baseurl }}/ja/commands/dropbox-team-filerequest-list.html)                                             | チームないのファイルリクエストを一覧します                                              |
| [dropbox team filesystem]({{ site.baseurl }}/ja/commands/dropbox-team-filesystem.html)                                                         | チームのファイルシステムのバージョンを特定する                                          |
| [dropbox team group add]({{ site.baseurl }}/ja/commands/dropbox-team-group-add.html)                                                           | グループを作成します                                                                    |
| [dropbox team group batch add]({{ site.baseurl }}/ja/commands/dropbox-team-group-batch-add.html)                                               | グループの一括追加                                                                      |
| [dropbox team group batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-batch-delete.html)                                         | グループの削除                                                                          |
| [dropbox team group clear externalid]({{ site.baseurl }}/ja/commands/dropbox-team-group-clear-externalid.html)                                 | グループの外部IDをクリアする                                                            |
| [dropbox team group delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-delete.html)                                                     | グループを削除します                                                                    |
| [dropbox team group folder list]({{ site.baseurl }}/ja/commands/dropbox-team-group-folder-list.html)                                           | 各グループのフォルダーを一覧表示                                                        |
| [dropbox team group list]({{ site.baseurl }}/ja/commands/dropbox-team-group-list.html)                                                         | グループを一覧                                                                          |
| [dropbox team group member add]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-add.html)                                             | メンバーをグループに追加                                                                |
| [dropbox team group member batch add]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-batch-add.html)                                 | グループにメンバーを一括追加                                                            |
| [dropbox team group member batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-batch-delete.html)                           | グループからメンバーを削除                                                              |
| [dropbox team group member batch update]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-batch-update.html)                           | グループからメンバーを追加または削除                                                    |
| [dropbox team group member delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-delete.html)                                       | メンバーをグループから削除                                                              |
| [dropbox team group member list]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-list.html)                                           | グループに所属するメンバー一覧を取得します                                              |
| [dropbox team group rename]({{ site.baseurl }}/ja/commands/dropbox-team-group-rename.html)                                                     | グループの改名                                                                          |
| [dropbox team group update type]({{ site.baseurl }}/ja/commands/dropbox-team-group-update-type.html)                                           | グループ管理タイプの更新                                                                |
| [dropbox team info]({{ site.baseurl }}/ja/commands/dropbox-team-info.html)                                                                     | チームの情報                                                                            |
| [dropbox team insight scan]({{ site.baseurl }}/ja/commands/dropbox-team-insight-scan.html)                                                     | チームデータをスキャンして分析                                                          |
| [dropbox team legalhold add]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-add.html)                                                   | 新しいリーガル・ホールド・ポリシーを作成する.                                           |
| [dropbox team legalhold list]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-list.html)                                                 | 既存のポリシーを取得する                                                                |
| [dropbox team legalhold member batch update]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-member-batch-update.html)                   | リーガル・ホールド・ポリシーのメンバーリスト更新                                        |
| [dropbox team legalhold member list]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-member-list.html)                                   | リーガルホールドのメンバーをリストアップ                                                |
| [dropbox team legalhold release]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-release.html)                                           | Idによるリーガルホールドを解除する                                                      |
| [dropbox team legalhold revision list]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-revision-list.html)                               | リーガル・ホールド・ポリシーのリビジョンをリストアップする                              |
| [dropbox team legalhold update desc]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-update-desc.html)                                   | リーガルホールド・ポリシーの説明を更新                                                  |
| [dropbox team legalhold update name]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-update-name.html)                                   | リーガルホールドポリシーの名称を更新                                                    |
| [dropbox team linkedapp list]({{ site.baseurl }}/ja/commands/dropbox-team-linkedapp-list.html)                                                 | リンク済みアプリを一覧                                                                  |
| [dropbox team member batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-delete.html)                                       | メンバーを削除します                                                                    |
| [dropbox team member batch detach]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-detach.html)                                       | Dropbox for teamsのアカウントをBasicアカウントに変更する                                |
| [dropbox team member batch invite]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-invite.html)                                       | メンバーを招待します                                                                    |
| [dropbox team member batch reinvite]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-reinvite.html)                                   | 招待済み状態メンバーをチームに再招待します                                              |
| [dropbox team member batch suspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-suspend.html)                                     | メンバーの一括一時停止                                                                  |
| [dropbox team member batch unsuspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-unsuspend.html)                                 | メンバーの一括停止解除                                                                  |
| [dropbox team member clear externalid]({{ site.baseurl }}/ja/commands/dropbox-team-member-clear-externalid.html)                               | メンバーのexternal_idを初期化します                                                     |
| [dropbox team member feature]({{ site.baseurl }}/ja/commands/dropbox-team-member-feature.html)                                                 | メンバーの機能設定一覧                                                                  |
| [dropbox team member file lock all release]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-lock-all-release.html)                     | メンバーのパスの下にあるすべてのロックを解除します                                      |
| [dropbox team member file lock list]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-lock-list.html)                                   | パスの下にあるメンバーのロックを一覧表示                                                |
| [dropbox team member file lock release]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-lock-release.html)                             | メンバーとしてパスのロックを解除します                                                  |
| [dropbox team member file permdelete]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-permdelete.html)                                 | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                  |
| [dropbox team member folder list]({{ site.baseurl }}/ja/commands/dropbox-team-member-folder-list.html)                                         | 各メンバーのフォルダーを一覧表示                                                        |
| [dropbox team member folder replication]({{ site.baseurl }}/ja/commands/dropbox-team-member-folder-replication.html)                           | フォルダを他のメンバーの個人フォルダに複製します                                        |
| [dropbox team member list]({{ site.baseurl }}/ja/commands/dropbox-team-member-list.html)                                                       | チームメンバーの一覧                                                                    |
| [dropbox team member quota batch update]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-batch-update.html)                           | チームメンバーの容量制限を変更                                                          |
| [dropbox team member quota list]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-list.html)                                           | メンバーの容量制限情報を一覧します                                                      |
| [dropbox team member quota usage]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-usage.html)                                         | チームメンバーのストレージ利用状況を取得                                                |
| [dropbox team member replication]({{ site.baseurl }}/ja/commands/dropbox-team-member-replication.html)                                         | チームメンバーのファイルを複製します                                                    |
| [dropbox team member suspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-suspend.html)                                                 | メンバーの一時停止処理                                                                  |
| [dropbox team member unsuspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-unsuspend.html)                                             | メンバーの一時停止を解除する                                                            |
| [dropbox team member update batch email]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-email.html)                           | メンバーのメールアドレス処理                                                            |
| [dropbox team member update batch externalid]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-externalid.html)                 | チームメンバーのExternal IDを更新します.                                                |
| [dropbox team member update batch invisible]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-invisible.html)                   | メンバーへのディレクトリ制限を有効にします                                              |
| [dropbox team member update batch profile]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-profile.html)                       | メンバーのプロフィール変更                                                              |
| [dropbox team member update batch visible]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-visible.html)                       | メンバーへのディレクトリ制限を無効にします                                              |
| [dropbox team namespace file list]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-file-list.html)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [dropbox team namespace file size]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-file-size.html)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [dropbox team namespace list]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-list.html)                                                 | チーム内すべての名前空間を一覧                                                          |
| [dropbox team namespace member list]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-member-list.html)                                   | チームフォルダ以下のファイル・フォルダを一覧                                            |
| [dropbox team namespace summary]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-summary.html)                                           | チーム・ネームスペースの状態概要を報告する.                                             |
| [dropbox team runas file batch copy]({{ site.baseurl }}/ja/commands/dropbox-team-runas-file-batch-copy.html)                                   | ファイル/フォルダーをメンバーとして一括コピー                                           |
| [dropbox team runas file list]({{ site.baseurl }}/ja/commands/dropbox-team-runas-file-list.html)                                               | メンバーとして実行するファイルやフォルダーの一覧                                        |
| [dropbox team runas file sync batch up]({{ site.baseurl }}/ja/commands/dropbox-team-runas-file-sync-batch-up.html)                             | メンバーとして動作する一括同期                                                          |
| [dropbox team runas sharedfolder batch leave]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-batch-leave.html)                 | 共有フォルダからメンバーとして一括退出                                                  |
| [dropbox team runas sharedfolder batch share]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-batch-share.html)                 | メンバーのフォルダを一括で共有                                                          |
| [dropbox team runas sharedfolder batch unshare]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-batch-unshare.html)             | メンバーのフォルダの共有を一括解除                                                      |
| [dropbox team runas sharedfolder isolate]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-isolate.html)                         | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| [dropbox team runas sharedfolder list]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-list.html)                               | 共有フォルダーの一覧をメンバーとして実行                                                |
| [dropbox team runas sharedfolder member batch add]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-member-batch-add.html)       | メンバーの共有フォルダにメンバーを一括追加                                              |
| [dropbox team runas sharedfolder member batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-member-batch-delete.html) | メンバーの共有フォルダからメンバーを一括削除                                            |
| [dropbox team runas sharedfolder mount add]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-add.html)                     | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| [dropbox team runas sharedfolder mount delete]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-delete.html)               | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| [dropbox team runas sharedfolder mount list]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-list.html)                   | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| [dropbox team runas sharedfolder mount mountable]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-mountable.html)         | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| [dropbox team sharedlink cap expiry]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-cap-expiry.html)                                   | チーム内の共有リンクに有効期限の上限を設定                                              |
| [dropbox team sharedlink cap visibility]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-cap-visibility.html)                           | チーム内の共有リンクに可視性の上限を設定                                                |
| [dropbox team sharedlink delete links]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-delete-links.html)                               | 共有リンクの一括削除                                                                    |
| [dropbox team sharedlink delete member]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-delete-member.html)                             | メンバーの共有リンクをすべて削除                                                        |
| [dropbox team sharedlink list]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-list.html)                                               | 共有リンクの一覧                                                                        |
| [dropbox team sharedlink update expiry]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-update-expiry.html)                             | チーム内の公開されている共有リンクについて有効期限を更新します                          |
| [dropbox team sharedlink update password]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-update-password.html)                         | 共有リンクのパスワードの設定・更新                                                      |
| [dropbox team sharedlink update visibility]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-update-visibility.html)                     | 共有リンクの可視性の更新                                                                |
| [dropbox team teamfolder add]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-add.html)                                                 | チームフォルダを追加します                                                              |
| [dropbox team teamfolder archive]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-archive.html)                                         | チームフォルダのアーカイブ                                                              |
| [dropbox team teamfolder batch archive]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-batch-archive.html)                             | 複数のチームフォルダをアーカイブします                                                  |
| [dropbox team teamfolder batch permdelete]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-batch-permdelete.html)                       | 複数のチームフォルダを完全に削除します                                                  |
| [dropbox team teamfolder batch replication]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-batch-replication.html)                     | チームフォルダの一括レプリケーション                                                    |
| [dropbox team teamfolder file list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-list.html)                                     | チームフォルダの一覧                                                                    |
| [dropbox team teamfolder file lock all release]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-lock-all-release.html)             | チームフォルダのパスの下にあるすべてのロックを解除する                                  |
| [dropbox team teamfolder file lock list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-lock-list.html)                           | チームフォルダ内のロックを一覧表示                                                      |
| [dropbox team teamfolder file lock release]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-lock-release.html)                     | チームフォルダ内のパスのロックを解除                                                    |
| [dropbox team teamfolder file size]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-size.html)                                     | チームフォルダのサイズを計算                                                            |
| [dropbox team teamfolder list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-list.html)                                               | チームフォルダの一覧                                                                    |
| [dropbox team teamfolder member add]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-member-add.html)                                   | チームフォルダへのユーザー/グループの一括追加                                           |
| [dropbox team teamfolder member delete]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-member-delete.html)                             | チームフォルダからのユーザー/グループの一括削除                                         |
| [dropbox team teamfolder member list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-member-list.html)                                 | チームフォルダのメンバー一覧                                                            |
| [dropbox team teamfolder partial replication]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-partial-replication.html)                 | 部分的なチームフォルダの他チームへのレプリケーション                                    |
| [dropbox team teamfolder permdelete]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-permdelete.html)                                   | チームフォルダを完全に削除します                                                        |
| [dropbox team teamfolder policy list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-policy-list.html)                                 | チームフォルダのポリシー一覧                                                            |
| [dropbox team teamfolder replication]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-replication.html)                                 | チームフォルダを他のチームに複製します                                                  |
| [dropbox team teamfolder sync setting list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-sync-setting-list.html)                     | チームフォルダーの同期設定を一覧表示                                                    |
| [dropbox team teamfolder sync setting update]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-sync-setting-update.html)                 | チームフォルダ同期設定の一括更新                                                        |
| [teamspace asadmin file list]({{ site.baseurl }}/ja/commands/teamspace-asadmin-file-list.html)                                                 | チームスペース内のファイルやフォルダーを一覧表示することができます。                    |
| [teamspace asadmin folder add]({{ site.baseurl }}/ja/commands/teamspace-asadmin-folder-add.html)                                               | チームスペースにトップレベルのフォルダーを作成                                          |
| [teamspace asadmin folder delete]({{ site.baseurl }}/ja/commands/teamspace-asadmin-folder-delete.html)                                         | チームスペースのトップレベルフォルダーを削除する                                        |
| [teamspace asadmin folder permdelete]({{ site.baseurl }}/ja/commands/teamspace-asadmin-folder-permdelete.html)                                 | チームスペースのトップレベルフォルダを完全に削除します。                                |

## DeepL

| コマンド                                                                         | 説明               |
|----------------------------------------------------------------------------------|--------------------|
| [deepl translate text]({{ site.baseurl }}/ja/commands/deepl-translate-text.html) | テキストを翻訳する |

## Figma

| コマンド                                                                                     | 説明                                                  |
|----------------------------------------------------------------------------------------------|-------------------------------------------------------|
| [figma account info]({{ site.baseurl }}/ja/commands/figma-account-info.html)                 | 現在のユーザー情報を取得する                          |
| [figma file export all page]({{ site.baseurl }}/ja/commands/figma-file-export-all-page.html) | チーム配下のすべてのファイル/ページをエクスポートする |
| [figma file export frame]({{ site.baseurl }}/ja/commands/figma-file-export-frame.html)       | Figmaファイルの全フレームを書き出す                   |
| [figma file export node]({{ site.baseurl }}/ja/commands/figma-file-export-node.html)         | Figmaドキュメント・ノードの書き出し                   |
| [figma file export page]({{ site.baseurl }}/ja/commands/figma-file-export-page.html)         | Figmaファイルの全ページを書き出す                     |
| [figma file info]({{ site.baseurl }}/ja/commands/figma-file-info.html)                       | figmaファイルの情報を表示する                         |
| [figma file list]({{ site.baseurl }}/ja/commands/figma-file-list.html)                       | Figmaプロジェクト内のファイル一覧                     |
| [figma project list]({{ site.baseurl }}/ja/commands/figma-project-list.html)                 | チームのプロジェクト一覧                              |

## GitHub

| コマンド                                                                                           | 説明                                                          |
|----------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [dev release checkin]({{ site.baseurl }}/ja/commands/dev-release-checkin.html)                     | 新作りリースをチェック                                        |
| [github content get]({{ site.baseurl }}/ja/commands/github-content-get.html)                       | レポジトリのコンテンツメタデータを取得します.                 |
| [github content put]({{ site.baseurl }}/ja/commands/github-content-put.html)                       | レポジトリに小さなテキストコンテンツを格納します              |
| [github issue list]({{ site.baseurl }}/ja/commands/github-issue-list.html)                         | 公開・プライベートGitHubレポジトリの課題一覧                  |
| [github profile]({{ site.baseurl }}/ja/commands/github-profile.html)                               | 認証したユーザーの情報を取得                                  |
| [github release asset download]({{ site.baseurl }}/ja/commands/github-release-asset-download.html) | アセットをダウンロードします                                  |
| [github release asset list]({{ site.baseurl }}/ja/commands/github-release-asset-list.html)         | GitHubリリースの成果物一覧                                    |
| [github release asset upload]({{ site.baseurl }}/ja/commands/github-release-asset-upload.html)     | GitHub リリースへ成果物をアップロードします                   |
| [github release draft]({{ site.baseurl }}/ja/commands/github-release-draft.html)                   | リリースの下書きを作成                                        |
| [github release list]({{ site.baseurl }}/ja/commands/github-release-list.html)                     | リリースの一覧                                                |
| [github tag create]({{ site.baseurl }}/ja/commands/github-tag-create.html)                         | レポジトリにタグを作成します                                  |
| [util release install]({{ site.baseurl }}/ja/commands/util-release-install.html)                   | watermint toolboxをダウンロードし、パスにインストールします。 |

## Google Calendar

| コマンド                                                                                     | 説明                                 |
|----------------------------------------------------------------------------------------------|--------------------------------------|
| [google calendar event list]({{ site.baseurl }}/ja/commands/google-calendar-event-list.html) | Googleカレンダーのイベントを一覧表示 |

## Google Gmail

| コマンド                                                                                                     | 説明                                               |
|--------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| [google mail filter add]({{ site.baseurl }}/ja/commands/google-mail-filter-add.html)                         | フィルターを追加します.                            |
| [google mail filter batch add]({{ site.baseurl }}/ja/commands/google-mail-filter-batch-add.html)             | クエリによるラベルの一括追加・削除                 |
| [google mail filter delete]({{ site.baseurl }}/ja/commands/google-mail-filter-delete.html)                   | フィルタの削除                                     |
| [google mail filter list]({{ site.baseurl }}/ja/commands/google-mail-filter-list.html)                       | フィルターの一覧                                   |
| [google mail label add]({{ site.baseurl }}/ja/commands/google-mail-label-add.html)                           | ラベルの追加                                       |
| [google mail label delete]({{ site.baseurl }}/ja/commands/google-mail-label-delete.html)                     | ラベルの削除.                                      |
| [google mail label list]({{ site.baseurl }}/ja/commands/google-mail-label-list.html)                         | ラベルのリスト                                     |
| [google mail label rename]({{ site.baseurl }}/ja/commands/google-mail-label-rename.html)                     | ラベルの名前を変更する                             |
| [google mail message label add]({{ site.baseurl }}/ja/commands/google-mail-message-label-add.html)           | メッセージにラベルを追加                           |
| [google mail message label delete]({{ site.baseurl }}/ja/commands/google-mail-message-label-delete.html)     | メッセージからラベルを削除する                     |
| [google mail message list]({{ site.baseurl }}/ja/commands/google-mail-message-list.html)                     | メッセージの一覧                                   |
| [google mail message processed list]({{ site.baseurl }}/ja/commands/google-mail-message-processed-list.html) | 処理された形式でメッセージを一覧表示します.        |
| [google mail sendas add]({{ site.baseurl }}/ja/commands/google-mail-sendas-add.html)                         | カスタムの "from" send-asエイリアスの作成          |
| [google mail sendas delete]({{ site.baseurl }}/ja/commands/google-mail-sendas-delete.html)                   | 指定したsend-asエイリアスを削除する                |
| [google mail sendas list]({{ site.baseurl }}/ja/commands/google-mail-sendas-list.html)                       | 指定されたアカウントの送信エイリアスを一覧表示する |
| [google mail thread list]({{ site.baseurl }}/ja/commands/google-mail-thread-list.html)                       | スレッド一覧                                       |

## Google Sheets

| コマンド                                                                                                 | 説明                                 |
|----------------------------------------------------------------------------------------------------------|--------------------------------------|
| [google sheets sheet append]({{ site.baseurl }}/ja/commands/google-sheets-sheet-append.html)             | スプレッドシートにデータを追加する   |
| [google sheets sheet clear]({{ site.baseurl }}/ja/commands/google-sheets-sheet-clear.html)               | スプレッドシートから値をクリアする   |
| [google sheets sheet create]({{ site.baseurl }}/ja/commands/google-sheets-sheet-create.html)             | 新規シートの作成                     |
| [google sheets sheet delete]({{ site.baseurl }}/ja/commands/google-sheets-sheet-delete.html)             | スプレッドシートからシートを削除する |
| [google sheets sheet export]({{ site.baseurl }}/ja/commands/google-sheets-sheet-export.html)             | シートデータのエクスポート           |
| [google sheets sheet import]({{ site.baseurl }}/ja/commands/google-sheets-sheet-import.html)             | スプレッドシートにデータをインポート |
| [google sheets sheet list]({{ site.baseurl }}/ja/commands/google-sheets-sheet-list.html)                 | スプレッドシートのシート一覧         |
| [google sheets spreadsheet create]({{ site.baseurl }}/ja/commands/google-sheets-spreadsheet-create.html) | 新しいスプレッドシートの作成         |

## ユーティリティー

| コマンド                                                                                                                         | 説明                                                                   |
|----------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [config auth delete]({{ site.baseurl }}/ja/commands/config-auth-delete.html)                                                     | 既存の認証クレデンシャルの削除                                         |
| [config auth list]({{ site.baseurl }}/ja/commands/config-auth-list.html)                                                         | すべての認証情報を一覧表示                                             |
| [config feature disable]({{ site.baseurl }}/ja/commands/config-feature-disable.html)                                             | 機能を無効化します.                                                    |
| [config feature enable]({{ site.baseurl }}/ja/commands/config-feature-enable.html)                                               | 機能を有効化します.                                                    |
| [config feature list]({{ site.baseurl }}/ja/commands/config-feature-list.html)                                                   | 利用可能なオプション機能一覧.                                          |
| [dropbox team insight report teamfoldermember]({{ site.baseurl }}/ja/commands/dropbox-team-insight-report-teamfoldermember.html) | チームフォルダーメンバーを報告                                         |
| [license]({{ site.baseurl }}/ja/commands/license.html)                                                                           | ライセンス情報を表示します                                             |
| [local file template apply]({{ site.baseurl }}/ja/commands/local-file-template-apply.html)                                       | ファイル/フォルダー構造のテンプレートをローカルパスに適用する          |
| [local file template capture]({{ site.baseurl }}/ja/commands/local-file-template-capture.html)                                   | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む      |
| [log cat curl]({{ site.baseurl }}/ja/commands/log-cat-curl.html)                                                                 | キャプチャログを `curl` サンプルとしてフォーマットする                 |
| [log cat job]({{ site.baseurl }}/ja/commands/log-cat-job.html)                                                                   | 指定したジョブIDのログを取得する                                       |
| [log cat kind]({{ site.baseurl }}/ja/commands/log-cat-kind.html)                                                                 | 指定種別のログを結合して出力します                                     |
| [log cat last]({{ site.baseurl }}/ja/commands/log-cat-last.html)                                                                 | 最後のジョブのログファイルを出力.                                      |
| [log job archive]({{ site.baseurl }}/ja/commands/log-job-archive.html)                                                           | ジョブのアーカイブ                                                     |
| [log job delete]({{ site.baseurl }}/ja/commands/log-job-delete.html)                                                             | 古いジョブ履歴の削除                                                   |
| [log job list]({{ site.baseurl }}/ja/commands/log-job-list.html)                                                                 | ジョブ履歴の表示                                                       |
| [util archive unzip]({{ site.baseurl }}/ja/commands/util-archive-unzip.html)                                                     | ZIPアーカイブファイルを解凍する                                        |
| [util archive zip]({{ site.baseurl }}/ja/commands/util-archive-zip.html)                                                         | 対象ファイルをZIPアーカイブに圧縮する                                  |
| [util cert selfsigned]({{ site.baseurl }}/ja/commands/util-cert-selfsigned.html)                                                 | 自己署名証明書と鍵の生成                                               |
| [util database exec]({{ site.baseurl }}/ja/commands/util-database-exec.html)                                                     | SQLite3データベースファイルへのクエリ実行                              |
| [util database query]({{ site.baseurl }}/ja/commands/util-database-query.html)                                                   | SQLite3データベースへの問い合わせ                                      |
| [util date today]({{ site.baseurl }}/ja/commands/util-date-today.html)                                                           | 現在の日付を表示                                                       |
| [util datetime now]({{ site.baseurl }}/ja/commands/util-datetime-now.html)                                                       | 現在の日時を表示                                                       |
| [util decode base32]({{ site.baseurl }}/ja/commands/util-decode-base32.html)                                                     | Base32 (RFC 4648) 形式からテキストをデコードします                     |
| [util decode base64]({{ site.baseurl }}/ja/commands/util-decode-base64.html)                                                     | Base64 (RFC 4648) フォーマットからテキストをデコードします             |
| [util desktop display list]({{ site.baseurl }}/ja/commands/util-desktop-display-list.html)                                       | このマシンのディスプレイを一覧表示                                     |
| [util desktop open]({{ site.baseurl }}/ja/commands/util-desktop-open.html)                                                       | デフォルトのアプリケーションでファイルやフォルダを開く                 |
| [util desktop screenshot interval]({{ site.baseurl }}/ja/commands/util-desktop-screenshot-interval.html)                         | 定期的にスクリーンショットを撮る                                       |
| [util desktop screenshot snap]({{ site.baseurl }}/ja/commands/util-desktop-screenshot-snap.html)                                 | スクリーンショットを撮る                                               |
| [util encode base32]({{ site.baseurl }}/ja/commands/util-encode-base32.html)                                                     | テキストをBase32(RFC 4648)形式にエンコード                             |
| [util encode base64]({{ site.baseurl }}/ja/commands/util-encode-base64.html)                                                     | テキストをBase64(RFC 4648)形式にエンコード                             |
| [util file hash]({{ site.baseurl }}/ja/commands/util-file-hash.html)                                                             | ファイルダイジェストの表示                                             |
| [util git clone]({{ site.baseurl }}/ja/commands/util-git-clone.html)                                                             | git リポジトリをクローン                                               |
| [util image exif]({{ site.baseurl }}/ja/commands/util-image-exif.html)                                                           | 画像ファイルのEXIFメタデータを表示                                     |
| [util image placeholder]({{ site.baseurl }}/ja/commands/util-image-placeholder.html)                                             | プレースホルダー画像の作成                                             |
| [util net download]({{ site.baseurl }}/ja/commands/util-net-download.html)                                                       | ファイルをダウンロードする                                             |
| [util qrcode create]({{ site.baseurl }}/ja/commands/util-qrcode-create.html)                                                     | QRコード画像ファイルの作成                                             |
| [util qrcode wifi]({{ site.baseurl }}/ja/commands/util-qrcode-wifi.html)                                                         | WIFI設定用のQRコードを生成                                             |
| [util table format xlsx]({{ site.baseurl }}/ja/commands/util-table-format-xlsx.html)                                             | xlsxファイルをテキストに整形する                                       |
| [util text case down]({{ site.baseurl }}/ja/commands/util-text-case-down.html)                                                   | 小文字のテキストを表示する                                             |
| [util text case up]({{ site.baseurl }}/ja/commands/util-text-case-up.html)                                                       | 大文字のテキストを表示する                                             |
| [util text encoding from]({{ site.baseurl }}/ja/commands/util-text-encoding-from.html)                                           | 指定されたエンコーディングからUTF-8テキストファイルに変換します.       |
| [util text encoding to]({{ site.baseurl }}/ja/commands/util-text-encoding-to.html)                                               | UTF-8テキストファイルから指定されたエンコーディングに変換する.         |
| [util text nlp english entity]({{ site.baseurl }}/ja/commands/util-text-nlp-english-entity.html)                                 | 英文をエンティティに分割する                                           |
| [util text nlp english sentence]({{ site.baseurl }}/ja/commands/util-text-nlp-english-sentence.html)                             | 英文を文章に分割する                                                   |
| [util text nlp english token]({{ site.baseurl }}/ja/commands/util-text-nlp-english-token.html)                                   | 英文をトークンに分割する                                               |
| [util text nlp japanese token]({{ site.baseurl }}/ja/commands/util-text-nlp-japanese-token.html)                                 | 日本語テキストのトークン化                                             |
| [util text nlp japanese wakati]({{ site.baseurl }}/ja/commands/util-text-nlp-japanese-wakati.html)                               | 分かち書き(日本語テキストのトークン化)                                 |
| [util tidy move dispatch]({{ site.baseurl }}/ja/commands/util-tidy-move-dispatch.html)                                           | ファイルを整理                                                         |
| [util tidy move simple]({{ site.baseurl }}/ja/commands/util-tidy-move-simple.html)                                               | ローカルファイルをアーカイブします                                     |
| [util time now]({{ site.baseurl }}/ja/commands/util-time-now.html)                                                               | 現在の時刻を表示                                                       |
| [util unixtime format]({{ site.baseurl }}/ja/commands/util-unixtime-format.html)                                                 | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット |
| [util unixtime now]({{ site.baseurl }}/ja/commands/util-unixtime-now.html)                                                       | UNIX時間で現在の時刻を表示する                                         |
| [util uuid v4]({{ site.baseurl }}/ja/commands/util-uuid-v4.html)                                                                 | UUID v4（ランダムUUID）の生成                                          |
| [util video subtitles optimize]({{ site.baseurl }}/ja/commands/util-video-subtitles-optimize.html)                               | 字幕ファイルの最適化                                                   |
| [util xlsx create]({{ site.baseurl }}/ja/commands/util-xlsx-create.html)                                                         | 空のスプレッドシートを作成する                                         |
| [util xlsx sheet export]({{ site.baseurl }}/ja/commands/util-xlsx-sheet-export.html)                                             | xlsxファイルからデータをエクスポート                                   |
| [util xlsx sheet import]({{ site.baseurl }}/ja/commands/util-xlsx-sheet-import.html)                                             | データをxlsxファイルにインポート                                       |
| [util xlsx sheet list]({{ site.baseurl }}/ja/commands/util-xlsx-sheet-list.html)                                                 | xlsxファイルのシート一覧                                               |
| [version]({{ site.baseurl }}/ja/commands/version.html)                                                                           | バージョン情報                                                         |


