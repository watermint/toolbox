---
layout: page
title: コマンド
lang: ja
---

# コマンド

## Dropbox (個人アカウント)

| コマンド                                                                                           | 説明                                                                      |
|----------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [file compare account]({{ site.baseurl }}/ja/commands/file-compare-account.html)                   | 二つのアカウントのファイルを比較します                                    |
| [file compare local]({{ site.baseurl }}/ja/commands/file-compare-local.html)                       | ローカルフォルダとDropboxフォルダの内容を比較します                       |
| [file copy]({{ site.baseurl }}/ja/commands/file-copy.html)                                         | ファイルをコピーします                                                    |
| [file delete]({{ site.baseurl }}/ja/commands/file-delete.html)                                     | ファイルまたはフォルダは削除します.                                       |
| [file export doc]({{ site.baseurl }}/ja/commands/file-export-doc.html)                             | ドキュメントのエクスポート                                                |
| [file export url]({{ site.baseurl }}/ja/commands/file-export-url.html)                             | URLからドキュメントをエクスポート                                         |
| [file import batch url]({{ site.baseurl }}/ja/commands/file-import-batch-url.html)                 | URLからファイルを一括インポートします                                     |
| [file import url]({{ site.baseurl }}/ja/commands/file-import-url.html)                             | URLからファイルをインポートします                                         |
| [file info]({{ site.baseurl }}/ja/commands/file-info.html)                                         | パスのメタデータを解決                                                    |
| [file list]({{ site.baseurl }}/ja/commands/file-list.html)                                         | ファイルとフォルダを一覧します                                            |
| [file lock acquire]({{ site.baseurl }}/ja/commands/file-lock-acquire.html)                         | ファイルをロック                                                          |
| [file lock all release]({{ site.baseurl }}/ja/commands/file-lock-all-release.html)                 | 指定したパスでのすべてのロックを解除する                                  |
| [file lock batch acquire]({{ site.baseurl }}/ja/commands/file-lock-batch-acquire.html)             | 複数のファイルをロックする                                                |
| [file lock batch release]({{ site.baseurl }}/ja/commands/file-lock-batch-release.html)             | 複数のロックを解除                                                        |
| [file lock list]({{ site.baseurl }}/ja/commands/file-lock-list.html)                               | 指定したパスの下にあるロックを一覧表示します                              |
| [file lock release]({{ site.baseurl }}/ja/commands/file-lock-release.html)                         | ロックを解除します                                                        |
| [file merge]({{ site.baseurl }}/ja/commands/file-merge.html)                                       | フォルダを統合します                                                      |
| [file move]({{ site.baseurl }}/ja/commands/file-move.html)                                         | ファイルを移動します                                                      |
| [file paper append]({{ site.baseurl }}/ja/commands/file-paper-append.html)                         | 既存のPaperドキュメントの最後にコンテンツを追加する                       |
| [file paper create]({{ site.baseurl }}/ja/commands/file-paper-create.html)                         | パスに新しいPaperを作成                                                   |
| [file paper overwrite]({{ site.baseurl }}/ja/commands/file-paper-overwrite.html)                   | 既存のPaperドキュメントを上書きする                                       |
| [file paper prepend]({{ site.baseurl }}/ja/commands/file-paper-prepend.html)                       | 既存のPaperドキュメントの先頭にコンテンツを追加する                       |
| [file replication]({{ site.baseurl }}/ja/commands/file-replication.html)                           | ファイルコンテンツを他のアカウントに複製します                            |
| [file restore all]({{ site.baseurl }}/ja/commands/file-restore-all.html)                           | 指定されたパス以下をリストアします                                        |
| [file revision download]({{ site.baseurl }}/ja/commands/file-revision-download.html)               | ファイルリビジョンをダウンロードする                                      |
| [file revision list]({{ site.baseurl }}/ja/commands/file-revision-list.html)                       | ファイルリビジョン一覧                                                    |
| [file revision restore]({{ site.baseurl }}/ja/commands/file-revision-restore.html)                 | ファイルリビジョンを復元する                                              |
| [file search content]({{ site.baseurl }}/ja/commands/file-search-content.html)                     | ファイルコンテンツを検索                                                  |
| [file search name]({{ site.baseurl }}/ja/commands/file-search-name.html)                           | ファイル名を検索                                                          |
| [file share info]({{ site.baseurl }}/ja/commands/file-share-info.html)                             | ファイルの共有情報を取得する                                              |
| [file size]({{ site.baseurl }}/ja/commands/file-size.html)                                         | ストレージの利用量                                                        |
| [file sync down]({{ site.baseurl }}/ja/commands/file-sync-down.html)                               | Dropboxと下り方向で同期します                                             |
| [file sync online]({{ site.baseurl }}/ja/commands/file-sync-online.html)                           | オンラインファイルを同期します                                            |
| [file sync up]({{ site.baseurl }}/ja/commands/file-sync-up.html)                                   | Dropboxと上り方向で同期します                                             |
| [file tag add]({{ site.baseurl }}/ja/commands/file-tag-add.html)                                   | ファイル/フォルダーにタグを追加する                                       |
| [file tag delete]({{ site.baseurl }}/ja/commands/file-tag-delete.html)                             | ファイル/フォルダーからタグを削除する                                     |
| [file tag list]({{ site.baseurl }}/ja/commands/file-tag-list.html)                                 | パスのタグを一覧                                                          |
| [file template apply remote]({{ site.baseurl }}/ja/commands/file-template-apply-remote.html)       | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する            |
| [file template capture remote]({{ site.baseurl }}/ja/commands/file-template-capture-remote.html)   | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。      |
| [file watch]({{ site.baseurl }}/ja/commands/file-watch.html)                                       | ファイルアクティビティを監視                                              |
| [filerequest create]({{ site.baseurl }}/ja/commands/filerequest-create.html)                       | ファイルリクエストを作成します                                            |
| [filerequest delete closed]({{ site.baseurl }}/ja/commands/filerequest-delete-closed.html)         | このアカウントの全ての閉じられているファイルリクエストを削除します        |
| [filerequest delete url]({{ site.baseurl }}/ja/commands/filerequest-delete-url.html)               | ファイルリクエストのURLを指定して削除                                     |
| [filerequest list]({{ site.baseurl }}/ja/commands/filerequest-list.html)                           | 個人アカウントのファイルリクエストを一覧.                                 |
| [job history ship]({{ site.baseurl }}/ja/commands/job-history-ship.html)                           | ログの転送先Dropboxパス                                                   |
| [services dropbox user feature]({{ site.baseurl }}/ja/commands/services-dropbox-user-feature.html) | 現在のユーザーの機能設定の一覧                                            |
| [services dropbox user info]({{ site.baseurl }}/ja/commands/services-dropbox-user-info.html)       | 現在のアカウント情報を取得する                                            |
| [sharedfolder leave]({{ site.baseurl }}/ja/commands/sharedfolder-leave.html)                       | 共有フォルダーから退出する.                                               |
| [sharedfolder list]({{ site.baseurl }}/ja/commands/sharedfolder-list.html)                         | 共有フォルダの一覧                                                        |
| [sharedfolder member add]({{ site.baseurl }}/ja/commands/sharedfolder-member-add.html)             | 共有フォルダへのメンバーの追加                                            |
| [sharedfolder member delete]({{ site.baseurl }}/ja/commands/sharedfolder-member-delete.html)       | 共有フォルダからメンバーを削除する                                        |
| [sharedfolder member list]({{ site.baseurl }}/ja/commands/sharedfolder-member-list.html)           | 共有フォルダのメンバーを一覧します                                        |
| [sharedfolder mount add]({{ site.baseurl }}/ja/commands/sharedfolder-mount-add.html)               | 共有フォルダを現在のユーザーのDropboxに追加する                           |
| [sharedfolder mount delete]({{ site.baseurl }}/ja/commands/sharedfolder-mount-delete.html)         | 現在のユーザーが指定されたフォルダーをアンマウントする.                   |
| [sharedfolder mount list]({{ site.baseurl }}/ja/commands/sharedfolder-mount-list.html)             | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示          |
| [sharedfolder mount mountable]({{ site.baseurl }}/ja/commands/sharedfolder-mount-mountable.html)   | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします. |
| [sharedfolder share]({{ site.baseurl }}/ja/commands/sharedfolder-share.html)                       | フォルダの共有                                                            |
| [sharedfolder unshare]({{ site.baseurl }}/ja/commands/sharedfolder-unshare.html)                   | フォルダの共有解除                                                        |
| [sharedlink create]({{ site.baseurl }}/ja/commands/sharedlink-create.html)                         | 共有リンクの作成                                                          |
| [sharedlink delete]({{ site.baseurl }}/ja/commands/sharedlink-delete.html)                         | 共有リンクを削除します                                                    |
| [sharedlink file list]({{ site.baseurl }}/ja/commands/sharedlink-file-list.html)                   | 共有リンクのファイルを一覧する                                            |
| [sharedlink info]({{ site.baseurl }}/ja/commands/sharedlink-info.html)                             | 共有リンクの情報取得                                                      |
| [sharedlink list]({{ site.baseurl }}/ja/commands/sharedlink-list.html)                             | 共有リンクの一覧                                                          |
| [teamspace file list]({{ site.baseurl }}/ja/commands/teamspace-file-list.html)                     | チームスペースにあるファイルやフォルダーを一覧表示                        |
| [util monitor client]({{ site.baseurl }}/ja/commands/util-monitor-client.html)                     | デバイスモニタークライアントを起動する                                    |
| [util tidy pack remote]({{ site.baseurl }}/ja/commands/util-tidy-pack-remote.html)                 | リモートフォルダをZIPファイルにパッケージする                             |

## Dropbox Business

| コマンド                                                                                                                       | 説明                                                                                    |
|--------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| [group add]({{ site.baseurl }}/ja/commands/group-add.html)                                                                     | グループを作成します                                                                    |
| [group batch add]({{ site.baseurl }}/ja/commands/group-batch-add.html)                                                         | グループの一括追加                                                                      |
| [group batch delete]({{ site.baseurl }}/ja/commands/group-batch-delete.html)                                                   | グループの削除                                                                          |
| [group clear externalid]({{ site.baseurl }}/ja/commands/group-clear-externalid.html)                                           | グループの外部IDをクリアする                                                            |
| [group delete]({{ site.baseurl }}/ja/commands/group-delete.html)                                                               | グループを削除します                                                                    |
| [group folder list]({{ site.baseurl }}/ja/commands/group-folder-list.html)                                                     | 各グループのフォルダーを一覧表示                                                        |
| [group list]({{ site.baseurl }}/ja/commands/group-list.html)                                                                   | グループを一覧                                                                          |
| [group member add]({{ site.baseurl }}/ja/commands/group-member-add.html)                                                       | メンバーをグループに追加                                                                |
| [group member batch add]({{ site.baseurl }}/ja/commands/group-member-batch-add.html)                                           | グループにメンバーを一括追加                                                            |
| [group member batch delete]({{ site.baseurl }}/ja/commands/group-member-batch-delete.html)                                     | グループからメンバーを削除                                                              |
| [group member batch update]({{ site.baseurl }}/ja/commands/group-member-batch-update.html)                                     | グループからメンバーを追加または削除                                                    |
| [group member delete]({{ site.baseurl }}/ja/commands/group-member-delete.html)                                                 | メンバーをグループから削除                                                              |
| [group member list]({{ site.baseurl }}/ja/commands/group-member-list.html)                                                     | グループに所属するメンバー一覧を取得します                                              |
| [group rename]({{ site.baseurl }}/ja/commands/group-rename.html)                                                               | グループの改名                                                                          |
| [member batch suspend]({{ site.baseurl }}/ja/commands/member-batch-suspend.html)                                               | メンバーの一括一時停止                                                                  |
| [member batch unsuspend]({{ site.baseurl }}/ja/commands/member-batch-unsuspend.html)                                           | メンバーの一括停止解除                                                                  |
| [member clear externalid]({{ site.baseurl }}/ja/commands/member-clear-externalid.html)                                         | メンバーのexternal_idを初期化します                                                     |
| [member delete]({{ site.baseurl }}/ja/commands/member-delete.html)                                                             | メンバーを削除します                                                                    |
| [member detach]({{ site.baseurl }}/ja/commands/member-detach.html)                                                             | Dropbox BusinessユーザーをBasicユーザーに変更します                                     |
| [member feature]({{ site.baseurl }}/ja/commands/member-feature.html)                                                           | メンバーの機能設定一覧                                                                  |
| [member file lock all release]({{ site.baseurl }}/ja/commands/member-file-lock-all-release.html)                               | メンバーのパスの下にあるすべてのロックを解除します                                      |
| [member file lock list]({{ site.baseurl }}/ja/commands/member-file-lock-list.html)                                             | パスの下にあるメンバーのロックを一覧表示                                                |
| [member file lock release]({{ site.baseurl }}/ja/commands/member-file-lock-release.html)                                       | メンバーとしてパスのロックを解除します                                                  |
| [member file permdelete]({{ site.baseurl }}/ja/commands/member-file-permdelete.html)                                           | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                  |
| [member folder list]({{ site.baseurl }}/ja/commands/member-folder-list.html)                                                   | 各メンバーのフォルダーを一覧表示                                                        |
| [member folder replication]({{ site.baseurl }}/ja/commands/member-folder-replication.html)                                     | フォルダを他のメンバーの個人フォルダに複製します                                        |
| [member invite]({{ site.baseurl }}/ja/commands/member-invite.html)                                                             | メンバーを招待します                                                                    |
| [member list]({{ site.baseurl }}/ja/commands/member-list.html)                                                                 | チームメンバーの一覧                                                                    |
| [member quota list]({{ site.baseurl }}/ja/commands/member-quota-list.html)                                                     | メンバーの容量制限情報を一覧します                                                      |
| [member quota update]({{ site.baseurl }}/ja/commands/member-quota-update.html)                                                 | チームメンバーの容量制限を変更                                                          |
| [member quota usage]({{ site.baseurl }}/ja/commands/member-quota-usage.html)                                                   | チームメンバーのストレージ利用状況を取得                                                |
| [member reinvite]({{ site.baseurl }}/ja/commands/member-reinvite.html)                                                         | 招待済み状態メンバーをチームに再招待します                                              |
| [member replication]({{ site.baseurl }}/ja/commands/member-replication.html)                                                   | チームメンバーのファイルを複製します                                                    |
| [member suspend]({{ site.baseurl }}/ja/commands/member-suspend.html)                                                           | メンバーの一時停止処理                                                                  |
| [member unsuspend]({{ site.baseurl }}/ja/commands/member-unsuspend.html)                                                       | メンバーの一時停止を解除する                                                            |
| [member update email]({{ site.baseurl }}/ja/commands/member-update-email.html)                                                 | メンバーのメールアドレス処理                                                            |
| [member update externalid]({{ site.baseurl }}/ja/commands/member-update-externalid.html)                                       | チームメンバーのExternal IDを更新します.                                                |
| [member update invisible]({{ site.baseurl }}/ja/commands/member-update-invisible.html)                                         | メンバーへのディレクトリ制限を有効にします                                              |
| [member update profile]({{ site.baseurl }}/ja/commands/member-update-profile.html)                                             | メンバーのプロフィール変更                                                              |
| [member update visible]({{ site.baseurl }}/ja/commands/member-update-visible.html)                                             | メンバーへのディレクトリ制限を無効にします                                              |
| [team activity batch user]({{ site.baseurl }}/ja/commands/team-activity-batch-user.html)                                       | 複数ユーザーのアクティビティを一括取得します                                            |
| [team activity daily event]({{ site.baseurl }}/ja/commands/team-activity-daily-event.html)                                     | アクティビティーを1日ごとに取得します                                                   |
| [team activity event]({{ site.baseurl }}/ja/commands/team-activity-event.html)                                                 | イベントログ                                                                            |
| [team activity user]({{ site.baseurl }}/ja/commands/team-activity-user.html)                                                   | ユーザーごとのアクティビティ                                                            |
| [team admin group role add]({{ site.baseurl }}/ja/commands/team-admin-group-role-add.html)                                     | グループのメンバーにロールを追加する                                                    |
| [team admin group role delete]({{ site.baseurl }}/ja/commands/team-admin-group-role-delete.html)                               | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                        |
| [team admin list]({{ site.baseurl }}/ja/commands/team-admin-list.html)                                                         | メンバーの管理者権限一覧                                                                |
| [team admin role add]({{ site.baseurl }}/ja/commands/team-admin-role-add.html)                                                 | メンバーに新しいロールを追加する                                                        |
| [team admin role clear]({{ site.baseurl }}/ja/commands/team-admin-role-clear.html)                                             | メンバーからすべての管理者ロールを削除する                                              |
| [team admin role delete]({{ site.baseurl }}/ja/commands/team-admin-role-delete.html)                                           | メンバーからロールを削除する                                                            |
| [team admin role list]({{ site.baseurl }}/ja/commands/team-admin-role-list.html)                                               | チームの管理者の役割を列挙                                                              |
| [team content legacypaper count]({{ site.baseurl }}/ja/commands/team-content-legacypaper-count.html)                           | メンバー1人あたりのPaper文書の枚数                                                      |
| [team content legacypaper export]({{ site.baseurl }}/ja/commands/team-content-legacypaper-export.html)                         | チームメンバー全員のPaper文書をローカルパスにエクスポート.                              |
| [team content legacypaper list]({{ site.baseurl }}/ja/commands/team-content-legacypaper-list.html)                             | チームメンバーのPaper文書リスト出力                                                     |
| [team content member list]({{ site.baseurl }}/ja/commands/team-content-member-list.html)                                       | チームフォルダや共有フォルダのメンバー一覧                                              |
| [team content member size]({{ site.baseurl }}/ja/commands/team-content-member-size.html)                                       | チームフォルダや共有フォルダのメンバー数をカウントする                                  |
| [team content mount list]({{ site.baseurl }}/ja/commands/team-content-mount-list.html)                                         | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.  |
| [team content policy list]({{ site.baseurl }}/ja/commands/team-content-policy-list.html)                                       | チームフォルダと共有フォルダのポリシー一覧                                              |
| [team device list]({{ site.baseurl }}/ja/commands/team-device-list.html)                                                       | チーム内全てのデバイス/セッションを一覧します                                           |
| [team device unlink]({{ site.baseurl }}/ja/commands/team-device-unlink.html)                                                   | デバイスのセッションを解除します                                                        |
| [team feature]({{ site.baseurl }}/ja/commands/team-feature.html)                                                               | チームの機能を出力します                                                                |
| [team filerequest list]({{ site.baseurl }}/ja/commands/team-filerequest-list.html)                                             | チームないのファイルリクエストを一覧します                                              |
| [team info]({{ site.baseurl }}/ja/commands/team-info.html)                                                                     | チームの情報                                                                            |
| [team legalhold add]({{ site.baseurl }}/ja/commands/team-legalhold-add.html)                                                   | 新しいリーガル・ホールド・ポリシーを作成する.                                           |
| [team legalhold list]({{ site.baseurl }}/ja/commands/team-legalhold-list.html)                                                 | 既存のポリシーを取得する                                                                |
| [team legalhold member batch update]({{ site.baseurl }}/ja/commands/team-legalhold-member-batch-update.html)                   | リーガル・ホールド・ポリシーのメンバーリスト更新                                        |
| [team legalhold member list]({{ site.baseurl }}/ja/commands/team-legalhold-member-list.html)                                   | リーガルホールドのメンバーをリストアップ                                                |
| [team legalhold release]({{ site.baseurl }}/ja/commands/team-legalhold-release.html)                                           | Idによるリーガルホールドを解除する                                                      |
| [team legalhold revision list]({{ site.baseurl }}/ja/commands/team-legalhold-revision-list.html)                               | リーガル・ホールド・ポリシーのリビジョンをリストアップする                              |
| [team legalhold update desc]({{ site.baseurl }}/ja/commands/team-legalhold-update-desc.html)                                   | リーガルホールド・ポリシーの説明を更新                                                  |
| [team legalhold update name]({{ site.baseurl }}/ja/commands/team-legalhold-update-name.html)                                   | リーガルホールドポリシーの名称を更新                                                    |
| [team linkedapp list]({{ site.baseurl }}/ja/commands/team-linkedapp-list.html)                                                 | リンク済みアプリを一覧                                                                  |
| [team namespace file list]({{ site.baseurl }}/ja/commands/team-namespace-file-list.html)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [team namespace file size]({{ site.baseurl }}/ja/commands/team-namespace-file-size.html)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [team namespace list]({{ site.baseurl }}/ja/commands/team-namespace-list.html)                                                 | チーム内すべての名前空間を一覧                                                          |
| [team namespace member list]({{ site.baseurl }}/ja/commands/team-namespace-member-list.html)                                   | チームフォルダ以下のファイル・フォルダを一覧                                            |
| [team namespace summary]({{ site.baseurl }}/ja/commands/team-namespace-summary.html)                                           | チーム・ネームスペースの状態概要を報告する.                                             |
| [team runas file batch copy]({{ site.baseurl }}/ja/commands/team-runas-file-batch-copy.html)                                   | ファイル/フォルダーをメンバーとして一括コピー                                           |
| [team runas file list]({{ site.baseurl }}/ja/commands/team-runas-file-list.html)                                               | メンバーとして実行するファイルやフォルダーの一覧                                        |
| [team runas file sync batch up]({{ site.baseurl }}/ja/commands/team-runas-file-sync-batch-up.html)                             | メンバーとして動作する一括同期                                                          |
| [team runas sharedfolder batch leave]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-batch-leave.html)                 | 共有フォルダからメンバーとして一括退出                                                  |
| [team runas sharedfolder batch share]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-batch-share.html)                 | メンバーのフォルダを一括で共有                                                          |
| [team runas sharedfolder batch unshare]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-batch-unshare.html)             | メンバーのフォルダの共有を一括解除                                                      |
| [team runas sharedfolder isolate]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-isolate.html)                         | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| [team runas sharedfolder list]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-list.html)                               | 共有フォルダーの一覧をメンバーとして実行                                                |
| [team runas sharedfolder member batch add]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-member-batch-add.html)       | メンバーの共有フォルダにメンバーを一括追加                                              |
| [team runas sharedfolder member batch delete]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-member-batch-delete.html) | メンバーの共有フォルダからメンバーを一括削除                                            |
| [team runas sharedfolder mount add]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-mount-add.html)                     | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| [team runas sharedfolder mount delete]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-mount-delete.html)               | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| [team runas sharedfolder mount list]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-mount-list.html)                   | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| [team runas sharedfolder mount mountable]({{ site.baseurl }}/ja/commands/team-runas-sharedfolder-mount-mountable.html)         | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| [team sharedlink cap expiry]({{ site.baseurl }}/ja/commands/team-sharedlink-cap-expiry.html)                                   | チーム内の共有リンクに有効期限の上限を設定                                              |
| [team sharedlink cap visibility]({{ site.baseurl }}/ja/commands/team-sharedlink-cap-visibility.html)                           | チーム内の共有リンクに可視性の上限を設定                                                |
| [team sharedlink delete links]({{ site.baseurl }}/ja/commands/team-sharedlink-delete-links.html)                               | 共有リンクの一括削除                                                                    |
| [team sharedlink delete member]({{ site.baseurl }}/ja/commands/team-sharedlink-delete-member.html)                             | メンバーの共有リンクをすべて削除                                                        |
| [team sharedlink list]({{ site.baseurl }}/ja/commands/team-sharedlink-list.html)                                               | 共有リンクの一覧                                                                        |
| [team sharedlink update expiry]({{ site.baseurl }}/ja/commands/team-sharedlink-update-expiry.html)                             | チーム内の公開されている共有リンクについて有効期限を更新します                          |
| [team sharedlink update password]({{ site.baseurl }}/ja/commands/team-sharedlink-update-password.html)                         | 共有リンクのパスワードの設定・更新                                                      |
| [team sharedlink update visibility]({{ site.baseurl }}/ja/commands/team-sharedlink-update-visibility.html)                     | 共有リンクの可視性の更新                                                                |
| [teamfolder add]({{ site.baseurl }}/ja/commands/teamfolder-add.html)                                                           | チームフォルダを追加します                                                              |
| [teamfolder archive]({{ site.baseurl }}/ja/commands/teamfolder-archive.html)                                                   | チームフォルダのアーカイブ                                                              |
| [teamfolder batch archive]({{ site.baseurl }}/ja/commands/teamfolder-batch-archive.html)                                       | 複数のチームフォルダをアーカイブします                                                  |
| [teamfolder batch permdelete]({{ site.baseurl }}/ja/commands/teamfolder-batch-permdelete.html)                                 | 複数のチームフォルダを完全に削除します                                                  |
| [teamfolder batch replication]({{ site.baseurl }}/ja/commands/teamfolder-batch-replication.html)                               | チームフォルダの一括レプリケーション                                                    |
| [teamfolder file list]({{ site.baseurl }}/ja/commands/teamfolder-file-list.html)                                               | チームフォルダの一覧                                                                    |
| [teamfolder file lock all release]({{ site.baseurl }}/ja/commands/teamfolder-file-lock-all-release.html)                       | チームフォルダのパスの下にあるすべてのロックを解除する                                  |
| [teamfolder file lock list]({{ site.baseurl }}/ja/commands/teamfolder-file-lock-list.html)                                     | チームフォルダ内のロックを一覧表示                                                      |
| [teamfolder file lock release]({{ site.baseurl }}/ja/commands/teamfolder-file-lock-release.html)                               | チームフォルダ内のパスのロックを解除                                                    |
| [teamfolder file size]({{ site.baseurl }}/ja/commands/teamfolder-file-size.html)                                               | チームフォルダのサイズを計算                                                            |
| [teamfolder list]({{ site.baseurl }}/ja/commands/teamfolder-list.html)                                                         | チームフォルダの一覧                                                                    |
| [teamfolder member add]({{ site.baseurl }}/ja/commands/teamfolder-member-add.html)                                             | チームフォルダへのユーザー/グループの一括追加                                           |
| [teamfolder member delete]({{ site.baseurl }}/ja/commands/teamfolder-member-delete.html)                                       | チームフォルダからのユーザー/グループの一括削除                                         |
| [teamfolder member list]({{ site.baseurl }}/ja/commands/teamfolder-member-list.html)                                           | チームフォルダのメンバー一覧                                                            |
| [teamfolder partial replication]({{ site.baseurl }}/ja/commands/teamfolder-partial-replication.html)                           | 部分的なチームフォルダの他チームへのレプリケーション                                    |
| [teamfolder permdelete]({{ site.baseurl }}/ja/commands/teamfolder-permdelete.html)                                             | チームフォルダを完全に削除します                                                        |
| [teamfolder policy list]({{ site.baseurl }}/ja/commands/teamfolder-policy-list.html)                                           | チームフォルダのポリシー一覧                                                            |
| [teamfolder replication]({{ site.baseurl }}/ja/commands/teamfolder-replication.html)                                           | チームフォルダを他のチームに複製します                                                  |
| [teamspace asadmin file list]({{ site.baseurl }}/ja/commands/teamspace-asadmin-file-list.html)                                 | チームスペース内のファイルやフォルダーを一覧表示することができます。                    |
| [teamspace asadmin folder add]({{ site.baseurl }}/ja/commands/teamspace-asadmin-folder-add.html)                               | チームスペースにトップレベルのフォルダーを作成                                          |
| [teamspace asadmin folder delete]({{ site.baseurl }}/ja/commands/teamspace-asadmin-folder-delete.html)                         | チームスペースのトップレベルフォルダーを削除する                                        |
| [teamspace asadmin folder permdelete]({{ site.baseurl }}/ja/commands/teamspace-asadmin-folder-permdelete.html)                 | チームスペースのトップレベルフォルダを完全に削除します。                                |

## Figma

| コマンド                                                                                                       | 説明                                                  |
|----------------------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| [services figma account info]({{ site.baseurl }}/ja/commands/services-figma-account-info.html)                 | 現在のユーザー情報を取得する                          |
| [services figma file export all page]({{ site.baseurl }}/ja/commands/services-figma-file-export-all-page.html) | チーム配下のすべてのファイル/ページをエクスポートする |
| [services figma file export frame]({{ site.baseurl }}/ja/commands/services-figma-file-export-frame.html)       | Figmaファイルの全フレームを書き出す                   |
| [services figma file export node]({{ site.baseurl }}/ja/commands/services-figma-file-export-node.html)         | Figmaドキュメント・ノードの書き出し                   |
| [services figma file export page]({{ site.baseurl }}/ja/commands/services-figma-file-export-page.html)         | Figmaファイルの全ページを書き出す                     |
| [services figma file info]({{ site.baseurl }}/ja/commands/services-figma-file-info.html)                       | figmaファイルの情報を表示する                         |
| [services figma file list]({{ site.baseurl }}/ja/commands/services-figma-file-list.html)                       | Figmaプロジェクト内のファイル一覧                     |
| [services figma project list]({{ site.baseurl }}/ja/commands/services-figma-project-list.html)                 | チームのプロジェクト一覧                              |

## GitHub

| コマンド                                                                                                             | 説明                                                          |
|----------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [services github content get]({{ site.baseurl }}/ja/commands/services-github-content-get.html)                       | レポジトリのコンテンツメタデータを取得します.                 |
| [services github content put]({{ site.baseurl }}/ja/commands/services-github-content-put.html)                       | レポジトリに小さなテキストコンテンツを格納します              |
| [services github issue list]({{ site.baseurl }}/ja/commands/services-github-issue-list.html)                         | 公開・プライベートGitHubレポジトリの課題一覧                  |
| [services github profile]({{ site.baseurl }}/ja/commands/services-github-profile.html)                               | 認証したユーザーの情報を取得                                  |
| [services github release asset download]({{ site.baseurl }}/ja/commands/services-github-release-asset-download.html) | アセットをダウンロードします                                  |
| [services github release asset list]({{ site.baseurl }}/ja/commands/services-github-release-asset-list.html)         | GitHubリリースの成果物一覧                                    |
| [services github release asset upload]({{ site.baseurl }}/ja/commands/services-github-release-asset-upload.html)     | GitHub リリースへ成果物をアップロードします                   |
| [services github release draft]({{ site.baseurl }}/ja/commands/services-github-release-draft.html)                   | リリースの下書きを作成                                        |
| [services github release list]({{ site.baseurl }}/ja/commands/services-github-release-list.html)                     | リリースの一覧                                                |
| [services github tag create]({{ site.baseurl }}/ja/commands/services-github-tag-create.html)                         | レポジトリにタグを作成します                                  |
| [util release install]({{ site.baseurl }}/ja/commands/util-release-install.html)                                     | watermint toolboxをダウンロードし、パスにインストールします。 |

## Google Calendar

| コマンド                                                                                                       | 説明                                 |
|----------------------------------------------------------------------------------------------------------------|--------------------------------------|
| [services google calendar event list]({{ site.baseurl }}/ja/commands/services-google-calendar-event-list.html) | Googleカレンダーのイベントを一覧表示 |

## Google Gmail

| コマンド                                                                                                                       | 説明                                               |
|--------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| [services google mail filter add]({{ site.baseurl }}/ja/commands/services-google-mail-filter-add.html)                         | フィルターを追加します.                            |
| [services google mail filter batch add]({{ site.baseurl }}/ja/commands/services-google-mail-filter-batch-add.html)             | クエリによるラベルの一括追加・削除                 |
| [services google mail filter delete]({{ site.baseurl }}/ja/commands/services-google-mail-filter-delete.html)                   | フィルタの削除                                     |
| [services google mail filter list]({{ site.baseurl }}/ja/commands/services-google-mail-filter-list.html)                       | フィルターの一覧                                   |
| [services google mail label add]({{ site.baseurl }}/ja/commands/services-google-mail-label-add.html)                           | ラベルの追加                                       |
| [services google mail label delete]({{ site.baseurl }}/ja/commands/services-google-mail-label-delete.html)                     | ラベルの削除.                                      |
| [services google mail label list]({{ site.baseurl }}/ja/commands/services-google-mail-label-list.html)                         | ラベルのリスト                                     |
| [services google mail label rename]({{ site.baseurl }}/ja/commands/services-google-mail-label-rename.html)                     | ラベルの名前を変更する                             |
| [services google mail message label add]({{ site.baseurl }}/ja/commands/services-google-mail-message-label-add.html)           | メッセージにラベルを追加                           |
| [services google mail message label delete]({{ site.baseurl }}/ja/commands/services-google-mail-message-label-delete.html)     | メッセージからラベルを削除する                     |
| [services google mail message list]({{ site.baseurl }}/ja/commands/services-google-mail-message-list.html)                     | メッセージの一覧                                   |
| [services google mail message processed list]({{ site.baseurl }}/ja/commands/services-google-mail-message-processed-list.html) | 処理された形式でメッセージを一覧表示します.        |
| [services google mail sendas add]({{ site.baseurl }}/ja/commands/services-google-mail-sendas-add.html)                         | カスタムの "from" send-asエイリアスの作成          |
| [services google mail sendas delete]({{ site.baseurl }}/ja/commands/services-google-mail-sendas-delete.html)                   | 指定したsend-asエイリアスを削除する                |
| [services google mail sendas list]({{ site.baseurl }}/ja/commands/services-google-mail-sendas-list.html)                       | 指定されたアカウントの送信エイリアスを一覧表示する |
| [services google mail thread list]({{ site.baseurl }}/ja/commands/services-google-mail-thread-list.html)                       | スレッド一覧                                       |

## Google Sheets

| コマンド                                                                                                                   | 説明                                 |
|----------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| [services google sheets sheet append]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-append.html)             | スプレッドシートにデータを追加する   |
| [services google sheets sheet clear]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-clear.html)               | スプレッドシートから値をクリアする   |
| [services google sheets sheet create]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-create.html)             | 新規シートの作成                     |
| [services google sheets sheet delete]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-delete.html)             | スプレッドシートからシートを削除する |
| [services google sheets sheet export]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-export.html)             | シートデータのエクスポート           |
| [services google sheets sheet import]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-import.html)             | スプレッドシートにデータをインポート |
| [services google sheets sheet list]({{ site.baseurl }}/ja/commands/services-google-sheets-sheet-list.html)                 | スプレッドシートのシート一覧         |
| [services google sheets spreadsheet create]({{ site.baseurl }}/ja/commands/services-google-sheets-spreadsheet-create.html) | 新しいスプレッドシートの作成         |

## ユーティリティー

| コマンド                                                                                       | 説明                                                                   |
|------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [config auth delete]({{ site.baseurl }}/ja/commands/config-auth-delete.html)                   | 既存の認証クレデンシャルの削除                                         |
| [config auth list]({{ site.baseurl }}/ja/commands/config-auth-list.html)                       | すべての認証情報を一覧表示                                             |
| [config disable]({{ site.baseurl }}/ja/commands/config-disable.html)                           | 機能を無効化します.                                                    |
| [config enable]({{ site.baseurl }}/ja/commands/config-enable.html)                             | 機能を有効化します.                                                    |
| [config features]({{ site.baseurl }}/ja/commands/config-features.html)                         | 利用可能なオプション機能一覧.                                          |
| [file template apply local]({{ site.baseurl }}/ja/commands/file-template-apply-local.html)     | ファイル/フォルダー構造のテンプレートをローカルパスに適用する          |
| [file template capture local]({{ site.baseurl }}/ja/commands/file-template-capture-local.html) | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む      |
| [job history archive]({{ site.baseurl }}/ja/commands/job-history-archive.html)                 | ジョブのアーカイブ                                                     |
| [job history delete]({{ site.baseurl }}/ja/commands/job-history-delete.html)                   | 古いジョブ履歴の削除                                                   |
| [job history list]({{ site.baseurl }}/ja/commands/job-history-list.html)                       | ジョブ履歴の表示                                                       |
| [job log jobid]({{ site.baseurl }}/ja/commands/job-log-jobid.html)                             | 指定したジョブIDのログを取得する                                       |
| [job log kind]({{ site.baseurl }}/ja/commands/job-log-kind.html)                               | 指定種別のログを結合して出力します                                     |
| [job log last]({{ site.baseurl }}/ja/commands/job-log-last.html)                               | 最後のジョブのログファイルを出力.                                      |
| [license]({{ site.baseurl }}/ja/commands/license.html)                                         | ライセンス情報を表示します                                             |
| [util archive unzip]({{ site.baseurl }}/ja/commands/util-archive-unzip.html)                   | ZIPアーカイブファイルを解凍する                                        |
| [util archive zip]({{ site.baseurl }}/ja/commands/util-archive-zip.html)                       | 対象ファイルをZIPアーカイブに圧縮する                                  |
| [util database exec]({{ site.baseurl }}/ja/commands/util-database-exec.html)                   | SQLite3データベースファイルへのクエリ実行                              |
| [util database query]({{ site.baseurl }}/ja/commands/util-database-query.html)                 | SQLite3データベースへの問い合わせ                                      |
| [util date today]({{ site.baseurl }}/ja/commands/util-date-today.html)                         | 現在の日付を表示                                                       |
| [util datetime now]({{ site.baseurl }}/ja/commands/util-datetime-now.html)                     | 現在の日時を表示                                                       |
| [util decode base32]({{ site.baseurl }}/ja/commands/util-decode-base32.html)                   | Base32 (RFC 4648) 形式からテキストをデコードします                     |
| [util decode base64]({{ site.baseurl }}/ja/commands/util-decode-base64.html)                   | Base64 (RFC 4648) フォーマットからテキストをデコードします             |
| [util encode base32]({{ site.baseurl }}/ja/commands/util-encode-base32.html)                   | テキストをBase32(RFC 4648)形式にエンコード                             |
| [util encode base64]({{ site.baseurl }}/ja/commands/util-encode-base64.html)                   | テキストをBase64(RFC 4648)形式にエンコード                             |
| [util file hash]({{ site.baseurl }}/ja/commands/util-file-hash.html)                           | ファイルダイジェストの表示                                             |
| [util git clone]({{ site.baseurl }}/ja/commands/util-git-clone.html)                           | git リポジトリをクローン                                               |
| [util image exif]({{ site.baseurl }}/ja/commands/util-image-exif.html)                         | 画像ファイルのEXIFメタデータを表示                                     |
| [util image placeholder]({{ site.baseurl }}/ja/commands/util-image-placeholder.html)           | プレースホルダー画像の作成                                             |
| [util net download]({{ site.baseurl }}/ja/commands/util-net-download.html)                     | ファイルをダウンロードする                                             |
| [util qrcode create]({{ site.baseurl }}/ja/commands/util-qrcode-create.html)                   | QRコード画像ファイルの作成                                             |
| [util qrcode wifi]({{ site.baseurl }}/ja/commands/util-qrcode-wifi.html)                       | WIFI設定用のQRコードを生成                                             |
| [util text case down]({{ site.baseurl }}/ja/commands/util-text-case-down.html)                 | 小文字のテキストを表示する                                             |
| [util text case up]({{ site.baseurl }}/ja/commands/util-text-case-up.html)                     | 大文字のテキストを表示する                                             |
| [util text encoding from]({{ site.baseurl }}/ja/commands/util-text-encoding-from.html)         | 指定されたエンコーディングからUTF-8テキストファイルに変換します.       |
| [util text encoding to]({{ site.baseurl }}/ja/commands/util-text-encoding-to.html)             | UTF-8テキストファイルから指定されたエンコーディングに変換する.         |
| [util tidy move dispatch]({{ site.baseurl }}/ja/commands/util-tidy-move-dispatch.html)         | ファイルを整理                                                         |
| [util tidy move simple]({{ site.baseurl }}/ja/commands/util-tidy-move-simple.html)             | ローカルファイルをアーカイブします                                     |
| [util time now]({{ site.baseurl }}/ja/commands/util-time-now.html)                             | 現在の時刻を表示                                                       |
| [util unixtime format]({{ site.baseurl }}/ja/commands/util-unixtime-format.html)               | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット |
| [util unixtime now]({{ site.baseurl }}/ja/commands/util-unixtime-now.html)                     | UNIX時間で現在の時刻を表示する                                         |
| [util xlsx create]({{ site.baseurl }}/ja/commands/util-xlsx-create.html)                       | 空のスプレッドシートを作成する                                         |
| [util xlsx sheet export]({{ site.baseurl }}/ja/commands/util-xlsx-sheet-export.html)           | xlsxファイルからデータをエクスポート                                   |
| [util xlsx sheet import]({{ site.baseurl }}/ja/commands/util-xlsx-sheet-import.html)           | データをxlsxファイルにインポート                                       |
| [util xlsx sheet list]({{ site.baseurl }}/ja/commands/util-xlsx-sheet-list.html)               | xlsxファイルのシート一覧                                               |
| [version]({{ site.baseurl }}/ja/commands/version.html)                                         | バージョン情報                                                         |


