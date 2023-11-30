# watermint toolbox

[![Build](https://github.com/watermint/toolbox/actions/workflows/build.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/build.yml)
[![Test](https://github.com/watermint/toolbox/actions/workflows/test.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/test.yml)
[![CodeQL](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml)
[![Codecov](https://codecov.io/gh/watermint/toolbox/branch/main/graph/badge.svg?token=CrE8reSVvE)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

Dropbox、Dropbox Business、Google、GitHubなどのWebサービスに対応した多目的ユーティリティ・コマンドラインツール.

# ライセンスと免責条項

watermint toolboxはMITライセンスのもと配布されています.
詳細はファイル LICENSE.mdまたは LICENSE.txt ご参照ください.

以下にご留意ください:
> ソフトウェアは「現状のまま」で、明示であるか暗黙であるかを問わず、何らの保証もなく提供されます。ここでいう保証とは、商品性、特定の目的への適合性、および権利非侵害についての保証も含みますが、それに限定されるものではありません。

# ビルド済み実行ファイル

コンパイル済みバイナリは [最新のリリース](https://github.com/watermint/toolbox/releases/latest) からダウンロードいただけます. ソースコードからビルドする場合には [BUILD.md](BUILD.md) を参照してください.

## macOS/LinuxでHomebrewを使ってインストールする。

まずHomebrewをインストールします. 手順は [オフィシャルサイト](https://brew.sh/)を参照してください. 次のコマンドを実行してwatermint toolboxをインストールします.
```
brew tap watermint/toolbox
brew install toolbox
```

# セキュリティとプライバシー

## 情報は収集しません 

watermint toolboxは、第三者のサーバーに情報を収集することはありません. 

watermint toolboxは、Dropbox のようなサービスとご自身のアカウントでやりとりするためのものです. 第三者のアカウントは関与していません. コマンドは、PCのストレージにAPIトークン、ログ、ファイル、またはレポートを保存します.

## 機密データ

APIトークンなどの機密データのほとんどは、難読化されてアクセス制限された状態でPCのストレージに保存されています. しかし、それらのデータを秘密にするのはあなたの責任です.
特に、ツールボックスのワークパスの下にある`secrets`フォルダ(デフォルトでは`C:\Users\<ユーザー名>\.toolbox`、または`$HOME/.toolbox`フォルダ以下)は共有しないでください。

# 利用方法

`tbx` にはたくさんの機能があります. オプションなしで実行をするとサポートされているコマンドやオプションの一覧が表示されます.
つぎのように引数なしで実行すると利用可能なコマンド・オプションがご確認いただけます.

```
% ./tbx

watermint toolbox xx.x.xxx
==========================

© 2016-2023 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

DropboxおよびDropbox Business向けのツールセット

使い方:
=======

./tbx  コマンド

利用可能なコマンド:
===================

| コマンド     | 説明                       | 備考 |
|--------------|----------------------------|------|
| config       | watermint toolbox の設定   |      |
| file         | ファイル操作               |      |
| filerequest  | ファイルリクエストの操作   |      |
| group        | グループ管理               |      |
| license      | ライセンス情報を表示します |      |
| member       | チームメンバーの管理       |      |
| sharedfolder | 共有フォルダ               |      |
| sharedlink   | 個人アカウントの共有リンク |      |
| team         | Dropbox Business チーム    |      |
| teamfolder   | チームフォルダの管理       |      |
| version      | バージョン情報             |      |

```

# コマンド

## Dropbox (個人アカウント)

| コマンド                                                                           | 説明                                                                      |
|------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [file compare account](docs/ja/commands/file-compare-account.md)                   | 二つのアカウントのファイルを比較します                                    |
| [file compare local](docs/ja/commands/file-compare-local.md)                       | ローカルフォルダとDropboxフォルダの内容を比較します                       |
| [file copy](docs/ja/commands/file-copy.md)                                         | ファイルをコピーします                                                    |
| [file delete](docs/ja/commands/file-delete.md)                                     | ファイルまたはフォルダは削除します.                                       |
| [file export doc](docs/ja/commands/file-export-doc.md)                             | ドキュメントのエクスポート                                                |
| [file export url](docs/ja/commands/file-export-url.md)                             | URLからドキュメントをエクスポート                                         |
| [file import batch url](docs/ja/commands/file-import-batch-url.md)                 | URLからファイルを一括インポートします                                     |
| [file import url](docs/ja/commands/file-import-url.md)                             | URLからファイルをインポートします                                         |
| [file info](docs/ja/commands/file-info.md)                                         | パスのメタデータを解決                                                    |
| [file list](docs/ja/commands/file-list.md)                                         | ファイルとフォルダを一覧します                                            |
| [file lock acquire](docs/ja/commands/file-lock-acquire.md)                         | ファイルをロック                                                          |
| [file lock all release](docs/ja/commands/file-lock-all-release.md)                 | 指定したパスでのすべてのロックを解除する                                  |
| [file lock batch acquire](docs/ja/commands/file-lock-batch-acquire.md)             | 複数のファイルをロックする                                                |
| [file lock batch release](docs/ja/commands/file-lock-batch-release.md)             | 複数のロックを解除                                                        |
| [file lock list](docs/ja/commands/file-lock-list.md)                               | 指定したパスの下にあるロックを一覧表示します                              |
| [file lock release](docs/ja/commands/file-lock-release.md)                         | ロックを解除します                                                        |
| [file merge](docs/ja/commands/file-merge.md)                                       | フォルダを統合します                                                      |
| [file move](docs/ja/commands/file-move.md)                                         | ファイルを移動します                                                      |
| [file paper append](docs/ja/commands/file-paper-append.md)                         | 既存のPaperドキュメントの最後にコンテンツを追加する                       |
| [file paper create](docs/ja/commands/file-paper-create.md)                         | パスに新しいPaperを作成                                                   |
| [file paper overwrite](docs/ja/commands/file-paper-overwrite.md)                   | 既存のPaperドキュメントを上書きする                                       |
| [file paper prepend](docs/ja/commands/file-paper-prepend.md)                       | 既存のPaperドキュメントの先頭にコンテンツを追加する                       |
| [file replication](docs/ja/commands/file-replication.md)                           | ファイルコンテンツを他のアカウントに複製します                            |
| [file restore all](docs/ja/commands/file-restore-all.md)                           | 指定されたパス以下をリストアします                                        |
| [file revision download](docs/ja/commands/file-revision-download.md)               | ファイルリビジョンをダウンロードする                                      |
| [file revision list](docs/ja/commands/file-revision-list.md)                       | ファイルリビジョン一覧                                                    |
| [file revision restore](docs/ja/commands/file-revision-restore.md)                 | ファイルリビジョンを復元する                                              |
| [file search content](docs/ja/commands/file-search-content.md)                     | ファイルコンテンツを検索                                                  |
| [file search name](docs/ja/commands/file-search-name.md)                           | ファイル名を検索                                                          |
| [file share info](docs/ja/commands/file-share-info.md)                             | ファイルの共有情報を取得する                                              |
| [file size](docs/ja/commands/file-size.md)                                         | ストレージの利用量                                                        |
| [file sync down](docs/ja/commands/file-sync-down.md)                               | Dropboxと下り方向で同期します                                             |
| [file sync online](docs/ja/commands/file-sync-online.md)                           | オンラインファイルを同期します                                            |
| [file sync up](docs/ja/commands/file-sync-up.md)                                   | Dropboxと上り方向で同期します                                             |
| [file tag add](docs/ja/commands/file-tag-add.md)                                   | ファイル/フォルダーにタグを追加する                                       |
| [file tag delete](docs/ja/commands/file-tag-delete.md)                             | ファイル/フォルダーからタグを削除する                                     |
| [file tag list](docs/ja/commands/file-tag-list.md)                                 | パスのタグを一覧                                                          |
| [file template apply remote](docs/ja/commands/file-template-apply-remote.md)       | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する            |
| [file template capture remote](docs/ja/commands/file-template-capture-remote.md)   | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。      |
| [file watch](docs/ja/commands/file-watch.md)                                       | ファイルアクティビティを監視                                              |
| [filerequest create](docs/ja/commands/filerequest-create.md)                       | ファイルリクエストを作成します                                            |
| [filerequest delete closed](docs/ja/commands/filerequest-delete-closed.md)         | このアカウントの全ての閉じられているファイルリクエストを削除します        |
| [filerequest delete url](docs/ja/commands/filerequest-delete-url.md)               | ファイルリクエストのURLを指定して削除                                     |
| [filerequest list](docs/ja/commands/filerequest-list.md)                           | 個人アカウントのファイルリクエストを一覧.                                 |
| [job history ship](docs/ja/commands/job-history-ship.md)                           | ログの転送先Dropboxパス                                                   |
| [services dropbox user feature](docs/ja/commands/services-dropbox-user-feature.md) | 現在のユーザーの機能設定の一覧                                            |
| [services dropbox user info](docs/ja/commands/services-dropbox-user-info.md)       | 現在のアカウント情報を取得する                                            |
| [sharedfolder leave](docs/ja/commands/sharedfolder-leave.md)                       | 共有フォルダーから退出する.                                               |
| [sharedfolder list](docs/ja/commands/sharedfolder-list.md)                         | 共有フォルダの一覧                                                        |
| [sharedfolder member add](docs/ja/commands/sharedfolder-member-add.md)             | 共有フォルダへのメンバーの追加                                            |
| [sharedfolder member delete](docs/ja/commands/sharedfolder-member-delete.md)       | 共有フォルダからメンバーを削除する                                        |
| [sharedfolder member list](docs/ja/commands/sharedfolder-member-list.md)           | 共有フォルダのメンバーを一覧します                                        |
| [sharedfolder mount add](docs/ja/commands/sharedfolder-mount-add.md)               | 共有フォルダを現在のユーザーのDropboxに追加する                           |
| [sharedfolder mount delete](docs/ja/commands/sharedfolder-mount-delete.md)         | 現在のユーザーが指定されたフォルダーをアンマウントする.                   |
| [sharedfolder mount list](docs/ja/commands/sharedfolder-mount-list.md)             | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示          |
| [sharedfolder mount mountable](docs/ja/commands/sharedfolder-mount-mountable.md)   | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします. |
| [sharedfolder share](docs/ja/commands/sharedfolder-share.md)                       | フォルダの共有                                                            |
| [sharedfolder unshare](docs/ja/commands/sharedfolder-unshare.md)                   | フォルダの共有解除                                                        |
| [sharedlink create](docs/ja/commands/sharedlink-create.md)                         | 共有リンクの作成                                                          |
| [sharedlink delete](docs/ja/commands/sharedlink-delete.md)                         | 共有リンクを削除します                                                    |
| [sharedlink file list](docs/ja/commands/sharedlink-file-list.md)                   | 共有リンクのファイルを一覧する                                            |
| [sharedlink info](docs/ja/commands/sharedlink-info.md)                             | 共有リンクの情報取得                                                      |
| [sharedlink list](docs/ja/commands/sharedlink-list.md)                             | 共有リンクの一覧                                                          |
| [teamspace file list](docs/ja/commands/teamspace-file-list.md)                     | チームスペースにあるファイルやフォルダーを一覧表示                        |
| [util monitor client](docs/ja/commands/util-monitor-client.md)                     | デバイスモニタークライアントを起動する                                    |
| [util tidy pack remote](docs/ja/commands/util-tidy-pack-remote.md)                 | リモートフォルダをZIPファイルにパッケージする                             |

## Dropbox Business

| コマンド                                                                                                       | 説明                                                                                    |
|----------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| [group add](docs/ja/commands/group-add.md)                                                                     | グループを作成します                                                                    |
| [group batch add](docs/ja/commands/group-batch-add.md)                                                         | グループの一括追加                                                                      |
| [group batch delete](docs/ja/commands/group-batch-delete.md)                                                   | グループの削除                                                                          |
| [group clear externalid](docs/ja/commands/group-clear-externalid.md)                                           | グループの外部IDをクリアする                                                            |
| [group delete](docs/ja/commands/group-delete.md)                                                               | グループを削除します                                                                    |
| [group folder list](docs/ja/commands/group-folder-list.md)                                                     | 各グループのフォルダーを一覧表示                                                        |
| [group list](docs/ja/commands/group-list.md)                                                                   | グループを一覧                                                                          |
| [group member add](docs/ja/commands/group-member-add.md)                                                       | メンバーをグループに追加                                                                |
| [group member batch add](docs/ja/commands/group-member-batch-add.md)                                           | グループにメンバーを一括追加                                                            |
| [group member batch delete](docs/ja/commands/group-member-batch-delete.md)                                     | グループからメンバーを削除                                                              |
| [group member batch update](docs/ja/commands/group-member-batch-update.md)                                     | グループからメンバーを追加または削除                                                    |
| [group member delete](docs/ja/commands/group-member-delete.md)                                                 | メンバーをグループから削除                                                              |
| [group member list](docs/ja/commands/group-member-list.md)                                                     | グループに所属するメンバー一覧を取得します                                              |
| [group rename](docs/ja/commands/group-rename.md)                                                               | グループの改名                                                                          |
| [group update type](docs/ja/commands/group-update-type.md)                                                     | グループ管理タイプの更新                                                                |
| [member batch suspend](docs/ja/commands/member-batch-suspend.md)                                               | メンバーの一括一時停止                                                                  |
| [member batch unsuspend](docs/ja/commands/member-batch-unsuspend.md)                                           | メンバーの一括停止解除                                                                  |
| [member clear externalid](docs/ja/commands/member-clear-externalid.md)                                         | メンバーのexternal_idを初期化します                                                     |
| [member delete](docs/ja/commands/member-delete.md)                                                             | メンバーを削除します                                                                    |
| [member detach](docs/ja/commands/member-detach.md)                                                             | Dropbox BusinessユーザーをBasicユーザーに変更します                                     |
| [member feature](docs/ja/commands/member-feature.md)                                                           | メンバーの機能設定一覧                                                                  |
| [member file lock all release](docs/ja/commands/member-file-lock-all-release.md)                               | メンバーのパスの下にあるすべてのロックを解除します                                      |
| [member file lock list](docs/ja/commands/member-file-lock-list.md)                                             | パスの下にあるメンバーのロックを一覧表示                                                |
| [member file lock release](docs/ja/commands/member-file-lock-release.md)                                       | メンバーとしてパスのロックを解除します                                                  |
| [member file permdelete](docs/ja/commands/member-file-permdelete.md)                                           | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                  |
| [member folder list](docs/ja/commands/member-folder-list.md)                                                   | 各メンバーのフォルダーを一覧表示                                                        |
| [member folder replication](docs/ja/commands/member-folder-replication.md)                                     | フォルダを他のメンバーの個人フォルダに複製します                                        |
| [member invite](docs/ja/commands/member-invite.md)                                                             | メンバーを招待します                                                                    |
| [member list](docs/ja/commands/member-list.md)                                                                 | チームメンバーの一覧                                                                    |
| [member quota list](docs/ja/commands/member-quota-list.md)                                                     | メンバーの容量制限情報を一覧します                                                      |
| [member quota update](docs/ja/commands/member-quota-update.md)                                                 | チームメンバーの容量制限を変更                                                          |
| [member quota usage](docs/ja/commands/member-quota-usage.md)                                                   | チームメンバーのストレージ利用状況を取得                                                |
| [member reinvite](docs/ja/commands/member-reinvite.md)                                                         | 招待済み状態メンバーをチームに再招待します                                              |
| [member replication](docs/ja/commands/member-replication.md)                                                   | チームメンバーのファイルを複製します                                                    |
| [member suspend](docs/ja/commands/member-suspend.md)                                                           | メンバーの一時停止処理                                                                  |
| [member unsuspend](docs/ja/commands/member-unsuspend.md)                                                       | メンバーの一時停止を解除する                                                            |
| [member update email](docs/ja/commands/member-update-email.md)                                                 | メンバーのメールアドレス処理                                                            |
| [member update externalid](docs/ja/commands/member-update-externalid.md)                                       | チームメンバーのExternal IDを更新します.                                                |
| [member update invisible](docs/ja/commands/member-update-invisible.md)                                         | メンバーへのディレクトリ制限を有効にします                                              |
| [member update profile](docs/ja/commands/member-update-profile.md)                                             | メンバーのプロフィール変更                                                              |
| [member update visible](docs/ja/commands/member-update-visible.md)                                             | メンバーへのディレクトリ制限を無効にします                                              |
| [team activity batch user](docs/ja/commands/team-activity-batch-user.md)                                       | 複数ユーザーのアクティビティを一括取得します                                            |
| [team activity daily event](docs/ja/commands/team-activity-daily-event.md)                                     | アクティビティーを1日ごとに取得します                                                   |
| [team activity event](docs/ja/commands/team-activity-event.md)                                                 | イベントログ                                                                            |
| [team activity user](docs/ja/commands/team-activity-user.md)                                                   | ユーザーごとのアクティビティ                                                            |
| [team admin group role add](docs/ja/commands/team-admin-group-role-add.md)                                     | グループのメンバーにロールを追加する                                                    |
| [team admin group role delete](docs/ja/commands/team-admin-group-role-delete.md)                               | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                        |
| [team admin list](docs/ja/commands/team-admin-list.md)                                                         | メンバーの管理者権限一覧                                                                |
| [team admin role add](docs/ja/commands/team-admin-role-add.md)                                                 | メンバーに新しいロールを追加する                                                        |
| [team admin role clear](docs/ja/commands/team-admin-role-clear.md)                                             | メンバーからすべての管理者ロールを削除する                                              |
| [team admin role delete](docs/ja/commands/team-admin-role-delete.md)                                           | メンバーからロールを削除する                                                            |
| [team admin role list](docs/ja/commands/team-admin-role-list.md)                                               | チームの管理者の役割を列挙                                                              |
| [team content legacypaper count](docs/ja/commands/team-content-legacypaper-count.md)                           | メンバー1人あたりのPaper文書の枚数                                                      |
| [team content legacypaper export](docs/ja/commands/team-content-legacypaper-export.md)                         | チームメンバー全員のPaper文書をローカルパスにエクスポート.                              |
| [team content legacypaper list](docs/ja/commands/team-content-legacypaper-list.md)                             | チームメンバーのPaper文書リスト出力                                                     |
| [team content member list](docs/ja/commands/team-content-member-list.md)                                       | チームフォルダや共有フォルダのメンバー一覧                                              |
| [team content member size](docs/ja/commands/team-content-member-size.md)                                       | チームフォルダや共有フォルダのメンバー数をカウントする                                  |
| [team content mount list](docs/ja/commands/team-content-mount-list.md)                                         | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.  |
| [team content policy list](docs/ja/commands/team-content-policy-list.md)                                       | チームフォルダと共有フォルダのポリシー一覧                                              |
| [team device list](docs/ja/commands/team-device-list.md)                                                       | チーム内全てのデバイス/セッションを一覧します                                           |
| [team device unlink](docs/ja/commands/team-device-unlink.md)                                                   | デバイスのセッションを解除します                                                        |
| [team feature](docs/ja/commands/team-feature.md)                                                               | チームの機能を出力します                                                                |
| [team filerequest list](docs/ja/commands/team-filerequest-list.md)                                             | チームないのファイルリクエストを一覧します                                              |
| [team filesystem](docs/ja/commands/team-filesystem.md)                                                         | Identify team's file system version                                                     |
| [team info](docs/ja/commands/team-info.md)                                                                     | チームの情報                                                                            |
| [team legalhold add](docs/ja/commands/team-legalhold-add.md)                                                   | 新しいリーガル・ホールド・ポリシーを作成する.                                           |
| [team legalhold list](docs/ja/commands/team-legalhold-list.md)                                                 | 既存のポリシーを取得する                                                                |
| [team legalhold member batch update](docs/ja/commands/team-legalhold-member-batch-update.md)                   | リーガル・ホールド・ポリシーのメンバーリスト更新                                        |
| [team legalhold member list](docs/ja/commands/team-legalhold-member-list.md)                                   | リーガルホールドのメンバーをリストアップ                                                |
| [team legalhold release](docs/ja/commands/team-legalhold-release.md)                                           | Idによるリーガルホールドを解除する                                                      |
| [team legalhold revision list](docs/ja/commands/team-legalhold-revision-list.md)                               | リーガル・ホールド・ポリシーのリビジョンをリストアップする                              |
| [team legalhold update desc](docs/ja/commands/team-legalhold-update-desc.md)                                   | リーガルホールド・ポリシーの説明を更新                                                  |
| [team legalhold update name](docs/ja/commands/team-legalhold-update-name.md)                                   | リーガルホールドポリシーの名称を更新                                                    |
| [team linkedapp list](docs/ja/commands/team-linkedapp-list.md)                                                 | リンク済みアプリを一覧                                                                  |
| [team namespace file list](docs/ja/commands/team-namespace-file-list.md)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [team namespace file size](docs/ja/commands/team-namespace-file-size.md)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [team namespace list](docs/ja/commands/team-namespace-list.md)                                                 | チーム内すべての名前空間を一覧                                                          |
| [team namespace member list](docs/ja/commands/team-namespace-member-list.md)                                   | チームフォルダ以下のファイル・フォルダを一覧                                            |
| [team namespace summary](docs/ja/commands/team-namespace-summary.md)                                           | チーム・ネームスペースの状態概要を報告する.                                             |
| [team runas file batch copy](docs/ja/commands/team-runas-file-batch-copy.md)                                   | ファイル/フォルダーをメンバーとして一括コピー                                           |
| [team runas file list](docs/ja/commands/team-runas-file-list.md)                                               | メンバーとして実行するファイルやフォルダーの一覧                                        |
| [team runas file sync batch up](docs/ja/commands/team-runas-file-sync-batch-up.md)                             | メンバーとして動作する一括同期                                                          |
| [team runas sharedfolder batch leave](docs/ja/commands/team-runas-sharedfolder-batch-leave.md)                 | 共有フォルダからメンバーとして一括退出                                                  |
| [team runas sharedfolder batch share](docs/ja/commands/team-runas-sharedfolder-batch-share.md)                 | メンバーのフォルダを一括で共有                                                          |
| [team runas sharedfolder batch unshare](docs/ja/commands/team-runas-sharedfolder-batch-unshare.md)             | メンバーのフォルダの共有を一括解除                                                      |
| [team runas sharedfolder isolate](docs/ja/commands/team-runas-sharedfolder-isolate.md)                         | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| [team runas sharedfolder list](docs/ja/commands/team-runas-sharedfolder-list.md)                               | 共有フォルダーの一覧をメンバーとして実行                                                |
| [team runas sharedfolder member batch add](docs/ja/commands/team-runas-sharedfolder-member-batch-add.md)       | メンバーの共有フォルダにメンバーを一括追加                                              |
| [team runas sharedfolder member batch delete](docs/ja/commands/team-runas-sharedfolder-member-batch-delete.md) | メンバーの共有フォルダからメンバーを一括削除                                            |
| [team runas sharedfolder mount add](docs/ja/commands/team-runas-sharedfolder-mount-add.md)                     | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| [team runas sharedfolder mount delete](docs/ja/commands/team-runas-sharedfolder-mount-delete.md)               | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| [team runas sharedfolder mount list](docs/ja/commands/team-runas-sharedfolder-mount-list.md)                   | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| [team runas sharedfolder mount mountable](docs/ja/commands/team-runas-sharedfolder-mount-mountable.md)         | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| [team sharedlink cap expiry](docs/ja/commands/team-sharedlink-cap-expiry.md)                                   | チーム内の共有リンクに有効期限の上限を設定                                              |
| [team sharedlink cap visibility](docs/ja/commands/team-sharedlink-cap-visibility.md)                           | チーム内の共有リンクに可視性の上限を設定                                                |
| [team sharedlink delete links](docs/ja/commands/team-sharedlink-delete-links.md)                               | 共有リンクの一括削除                                                                    |
| [team sharedlink delete member](docs/ja/commands/team-sharedlink-delete-member.md)                             | メンバーの共有リンクをすべて削除                                                        |
| [team sharedlink list](docs/ja/commands/team-sharedlink-list.md)                                               | 共有リンクの一覧                                                                        |
| [team sharedlink update expiry](docs/ja/commands/team-sharedlink-update-expiry.md)                             | チーム内の公開されている共有リンクについて有効期限を更新します                          |
| [team sharedlink update password](docs/ja/commands/team-sharedlink-update-password.md)                         | 共有リンクのパスワードの設定・更新                                                      |
| [team sharedlink update visibility](docs/ja/commands/team-sharedlink-update-visibility.md)                     | 共有リンクの可視性の更新                                                                |
| [teamfolder add](docs/ja/commands/teamfolder-add.md)                                                           | チームフォルダを追加します                                                              |
| [teamfolder archive](docs/ja/commands/teamfolder-archive.md)                                                   | チームフォルダのアーカイブ                                                              |
| [teamfolder batch archive](docs/ja/commands/teamfolder-batch-archive.md)                                       | 複数のチームフォルダをアーカイブします                                                  |
| [teamfolder batch permdelete](docs/ja/commands/teamfolder-batch-permdelete.md)                                 | 複数のチームフォルダを完全に削除します                                                  |
| [teamfolder batch replication](docs/ja/commands/teamfolder-batch-replication.md)                               | チームフォルダの一括レプリケーション                                                    |
| [teamfolder file list](docs/ja/commands/teamfolder-file-list.md)                                               | チームフォルダの一覧                                                                    |
| [teamfolder file lock all release](docs/ja/commands/teamfolder-file-lock-all-release.md)                       | チームフォルダのパスの下にあるすべてのロックを解除する                                  |
| [teamfolder file lock list](docs/ja/commands/teamfolder-file-lock-list.md)                                     | チームフォルダ内のロックを一覧表示                                                      |
| [teamfolder file lock release](docs/ja/commands/teamfolder-file-lock-release.md)                               | チームフォルダ内のパスのロックを解除                                                    |
| [teamfolder file size](docs/ja/commands/teamfolder-file-size.md)                                               | チームフォルダのサイズを計算                                                            |
| [teamfolder list](docs/ja/commands/teamfolder-list.md)                                                         | チームフォルダの一覧                                                                    |
| [teamfolder member add](docs/ja/commands/teamfolder-member-add.md)                                             | チームフォルダへのユーザー/グループの一括追加                                           |
| [teamfolder member delete](docs/ja/commands/teamfolder-member-delete.md)                                       | チームフォルダからのユーザー/グループの一括削除                                         |
| [teamfolder member list](docs/ja/commands/teamfolder-member-list.md)                                           | チームフォルダのメンバー一覧                                                            |
| [teamfolder partial replication](docs/ja/commands/teamfolder-partial-replication.md)                           | 部分的なチームフォルダの他チームへのレプリケーション                                    |
| [teamfolder permdelete](docs/ja/commands/teamfolder-permdelete.md)                                             | チームフォルダを完全に削除します                                                        |
| [teamfolder policy list](docs/ja/commands/teamfolder-policy-list.md)                                           | チームフォルダのポリシー一覧                                                            |
| [teamfolder replication](docs/ja/commands/teamfolder-replication.md)                                           | チームフォルダを他のチームに複製します                                                  |
| [teamfolder sync setting list](docs/ja/commands/teamfolder-sync-setting-list.md)                               | チームフォルダーの同期設定を一覧表示                                                    |
| [teamfolder sync setting update](docs/ja/commands/teamfolder-sync-setting-update.md)                           | チームフォルダ同期設定の一括更新                                                        |
| [teamspace asadmin file list](docs/ja/commands/teamspace-asadmin-file-list.md)                                 | チームスペース内のファイルやフォルダーを一覧表示することができます。                    |
| [teamspace asadmin folder add](docs/ja/commands/teamspace-asadmin-folder-add.md)                               | チームスペースにトップレベルのフォルダーを作成                                          |
| [teamspace asadmin folder delete](docs/ja/commands/teamspace-asadmin-folder-delete.md)                         | チームスペースのトップレベルフォルダーを削除する                                        |
| [teamspace asadmin folder permdelete](docs/ja/commands/teamspace-asadmin-folder-permdelete.md)                 | チームスペースのトップレベルフォルダを完全に削除します。                                |

## DeepL

| コマンド                                                                           | 説明           |
|------------------------------------------------------------------------------------|----------------|
| [services deepl translate text](docs/ja/commands/services-deepl-translate-text.md) | Translate text |

## Figma

| コマンド                                                                                       | 説明                                                  |
|------------------------------------------------------------------------------------------------|-------------------------------------------------------|
| [services figma account info](docs/ja/commands/services-figma-account-info.md)                 | 現在のユーザー情報を取得する                          |
| [services figma file export all page](docs/ja/commands/services-figma-file-export-all-page.md) | チーム配下のすべてのファイル/ページをエクスポートする |
| [services figma file export frame](docs/ja/commands/services-figma-file-export-frame.md)       | Figmaファイルの全フレームを書き出す                   |
| [services figma file export node](docs/ja/commands/services-figma-file-export-node.md)         | Figmaドキュメント・ノードの書き出し                   |
| [services figma file export page](docs/ja/commands/services-figma-file-export-page.md)         | Figmaファイルの全ページを書き出す                     |
| [services figma file info](docs/ja/commands/services-figma-file-info.md)                       | figmaファイルの情報を表示する                         |
| [services figma file list](docs/ja/commands/services-figma-file-list.md)                       | Figmaプロジェクト内のファイル一覧                     |
| [services figma project list](docs/ja/commands/services-figma-project-list.md)                 | チームのプロジェクト一覧                              |

## GitHub

| コマンド                                                                                             | 説明                                                          |
|------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [services github content get](docs/ja/commands/services-github-content-get.md)                       | レポジトリのコンテンツメタデータを取得します.                 |
| [services github content put](docs/ja/commands/services-github-content-put.md)                       | レポジトリに小さなテキストコンテンツを格納します              |
| [services github issue list](docs/ja/commands/services-github-issue-list.md)                         | 公開・プライベートGitHubレポジトリの課題一覧                  |
| [services github profile](docs/ja/commands/services-github-profile.md)                               | 認証したユーザーの情報を取得                                  |
| [services github release asset download](docs/ja/commands/services-github-release-asset-download.md) | アセットをダウンロードします                                  |
| [services github release asset list](docs/ja/commands/services-github-release-asset-list.md)         | GitHubリリースの成果物一覧                                    |
| [services github release asset upload](docs/ja/commands/services-github-release-asset-upload.md)     | GitHub リリースへ成果物をアップロードします                   |
| [services github release draft](docs/ja/commands/services-github-release-draft.md)                   | リリースの下書きを作成                                        |
| [services github release list](docs/ja/commands/services-github-release-list.md)                     | リリースの一覧                                                |
| [services github tag create](docs/ja/commands/services-github-tag-create.md)                         | レポジトリにタグを作成します                                  |
| [util release install](docs/ja/commands/util-release-install.md)                                     | watermint toolboxをダウンロードし、パスにインストールします。 |

## Google Calendar

| コマンド                                                                                       | 説明                                 |
|------------------------------------------------------------------------------------------------|--------------------------------------|
| [services google calendar event list](docs/ja/commands/services-google-calendar-event-list.md) | Googleカレンダーのイベントを一覧表示 |

## Google Gmail

| コマンド                                                                                                       | 説明                                               |
|----------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| [services google mail filter add](docs/ja/commands/services-google-mail-filter-add.md)                         | フィルターを追加します.                            |
| [services google mail filter batch add](docs/ja/commands/services-google-mail-filter-batch-add.md)             | クエリによるラベルの一括追加・削除                 |
| [services google mail filter delete](docs/ja/commands/services-google-mail-filter-delete.md)                   | フィルタの削除                                     |
| [services google mail filter list](docs/ja/commands/services-google-mail-filter-list.md)                       | フィルターの一覧                                   |
| [services google mail label add](docs/ja/commands/services-google-mail-label-add.md)                           | ラベルの追加                                       |
| [services google mail label delete](docs/ja/commands/services-google-mail-label-delete.md)                     | ラベルの削除.                                      |
| [services google mail label list](docs/ja/commands/services-google-mail-label-list.md)                         | ラベルのリスト                                     |
| [services google mail label rename](docs/ja/commands/services-google-mail-label-rename.md)                     | ラベルの名前を変更する                             |
| [services google mail message label add](docs/ja/commands/services-google-mail-message-label-add.md)           | メッセージにラベルを追加                           |
| [services google mail message label delete](docs/ja/commands/services-google-mail-message-label-delete.md)     | メッセージからラベルを削除する                     |
| [services google mail message list](docs/ja/commands/services-google-mail-message-list.md)                     | メッセージの一覧                                   |
| [services google mail message processed list](docs/ja/commands/services-google-mail-message-processed-list.md) | 処理された形式でメッセージを一覧表示します.        |
| [services google mail sendas add](docs/ja/commands/services-google-mail-sendas-add.md)                         | カスタムの "from" send-asエイリアスの作成          |
| [services google mail sendas delete](docs/ja/commands/services-google-mail-sendas-delete.md)                   | 指定したsend-asエイリアスを削除する                |
| [services google mail sendas list](docs/ja/commands/services-google-mail-sendas-list.md)                       | 指定されたアカウントの送信エイリアスを一覧表示する |
| [services google mail thread list](docs/ja/commands/services-google-mail-thread-list.md)                       | スレッド一覧                                       |

## Google Sheets

| コマンド                                                                                                   | 説明                                 |
|------------------------------------------------------------------------------------------------------------|--------------------------------------|
| [services google sheets sheet append](docs/ja/commands/services-google-sheets-sheet-append.md)             | スプレッドシートにデータを追加する   |
| [services google sheets sheet clear](docs/ja/commands/services-google-sheets-sheet-clear.md)               | スプレッドシートから値をクリアする   |
| [services google sheets sheet create](docs/ja/commands/services-google-sheets-sheet-create.md)             | 新規シートの作成                     |
| [services google sheets sheet delete](docs/ja/commands/services-google-sheets-sheet-delete.md)             | スプレッドシートからシートを削除する |
| [services google sheets sheet export](docs/ja/commands/services-google-sheets-sheet-export.md)             | シートデータのエクスポート           |
| [services google sheets sheet import](docs/ja/commands/services-google-sheets-sheet-import.md)             | スプレッドシートにデータをインポート |
| [services google sheets sheet list](docs/ja/commands/services-google-sheets-sheet-list.md)                 | スプレッドシートのシート一覧         |
| [services google sheets spreadsheet create](docs/ja/commands/services-google-sheets-spreadsheet-create.md) | 新しいスプレッドシートの作成         |

## ユーティリティー

| コマンド                                                                                 | 説明                                                                   |
|------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [config auth delete](docs/ja/commands/config-auth-delete.md)                             | 既存の認証クレデンシャルの削除                                         |
| [config auth list](docs/ja/commands/config-auth-list.md)                                 | すべての認証情報を一覧表示                                             |
| [config disable](docs/ja/commands/config-disable.md)                                     | 機能を無効化します.                                                    |
| [config enable](docs/ja/commands/config-enable.md)                                       | 機能を有効化します.                                                    |
| [config features](docs/ja/commands/config-features.md)                                   | 利用可能なオプション機能一覧.                                          |
| [file template apply local](docs/ja/commands/file-template-apply-local.md)               | ファイル/フォルダー構造のテンプレートをローカルパスに適用する          |
| [file template capture local](docs/ja/commands/file-template-capture-local.md)           | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む      |
| [job history archive](docs/ja/commands/job-history-archive.md)                           | ジョブのアーカイブ                                                     |
| [job history delete](docs/ja/commands/job-history-delete.md)                             | 古いジョブ履歴の削除                                                   |
| [job history list](docs/ja/commands/job-history-list.md)                                 | ジョブ履歴の表示                                                       |
| [job log jobid](docs/ja/commands/job-log-jobid.md)                                       | 指定したジョブIDのログを取得する                                       |
| [job log kind](docs/ja/commands/job-log-kind.md)                                         | 指定種別のログを結合して出力します                                     |
| [job log last](docs/ja/commands/job-log-last.md)                                         | 最後のジョブのログファイルを出力.                                      |
| [license](docs/ja/commands/license.md)                                                   | ライセンス情報を表示します                                             |
| [util archive unzip](docs/ja/commands/util-archive-unzip.md)                             | ZIPアーカイブファイルを解凍する                                        |
| [util archive zip](docs/ja/commands/util-archive-zip.md)                                 | 対象ファイルをZIPアーカイブに圧縮する                                  |
| [util cert selfsigned](docs/ja/commands/util-cert-selfsigned.md)                         | 自己署名証明書と鍵の生成                                               |
| [util database exec](docs/ja/commands/util-database-exec.md)                             | SQLite3データベースファイルへのクエリ実行                              |
| [util database query](docs/ja/commands/util-database-query.md)                           | SQLite3データベースへの問い合わせ                                      |
| [util date today](docs/ja/commands/util-date-today.md)                                   | 現在の日付を表示                                                       |
| [util datetime now](docs/ja/commands/util-datetime-now.md)                               | 現在の日時を表示                                                       |
| [util decode base32](docs/ja/commands/util-decode-base32.md)                             | Base32 (RFC 4648) 形式からテキストをデコードします                     |
| [util decode base64](docs/ja/commands/util-decode-base64.md)                             | Base64 (RFC 4648) フォーマットからテキストをデコードします             |
| [util desktop display list](docs/ja/commands/util-desktop-display-list.md)               | List displays of the current machine                                   |
| [util desktop open](docs/ja/commands/util-desktop-open.md)                               | Open a file or folder with the default application                     |
| [util desktop screenshot interval](docs/ja/commands/util-desktop-screenshot-interval.md) | Take screenshots at regular intervals                                  |
| [util desktop screenshot snap](docs/ja/commands/util-desktop-screenshot-snap.md)         | Take a screenshot                                                      |
| [util encode base32](docs/ja/commands/util-encode-base32.md)                             | テキストをBase32(RFC 4648)形式にエンコード                             |
| [util encode base64](docs/ja/commands/util-encode-base64.md)                             | テキストをBase64(RFC 4648)形式にエンコード                             |
| [util file hash](docs/ja/commands/util-file-hash.md)                                     | ファイルダイジェストの表示                                             |
| [util git clone](docs/ja/commands/util-git-clone.md)                                     | git リポジトリをクローン                                               |
| [util image exif](docs/ja/commands/util-image-exif.md)                                   | 画像ファイルのEXIFメタデータを表示                                     |
| [util image placeholder](docs/ja/commands/util-image-placeholder.md)                     | プレースホルダー画像の作成                                             |
| [util net download](docs/ja/commands/util-net-download.md)                               | ファイルをダウンロードする                                             |
| [util qrcode create](docs/ja/commands/util-qrcode-create.md)                             | QRコード画像ファイルの作成                                             |
| [util qrcode wifi](docs/ja/commands/util-qrcode-wifi.md)                                 | WIFI設定用のQRコードを生成                                             |
| [util table format xlsx](docs/ja/commands/util-table-format-xlsx.md)                     | xlsxファイルをテキストに整形する                                       |
| [util text case down](docs/ja/commands/util-text-case-down.md)                           | 小文字のテキストを表示する                                             |
| [util text case up](docs/ja/commands/util-text-case-up.md)                               | 大文字のテキストを表示する                                             |
| [util text encoding from](docs/ja/commands/util-text-encoding-from.md)                   | 指定されたエンコーディングからUTF-8テキストファイルに変換します.       |
| [util text encoding to](docs/ja/commands/util-text-encoding-to.md)                       | UTF-8テキストファイルから指定されたエンコーディングに変換する.         |
| [util text nlp english entity](docs/ja/commands/util-text-nlp-english-entity.md)         | 英文をエンティティに分割する                                           |
| [util text nlp english sentence](docs/ja/commands/util-text-nlp-english-sentence.md)     | 英文を文章に分割する                                                   |
| [util text nlp english token](docs/ja/commands/util-text-nlp-english-token.md)           | 英文をトークンに分割する                                               |
| [util text nlp japanese token](docs/ja/commands/util-text-nlp-japanese-token.md)         | 日本語テキストのトークン化                                             |
| [util text nlp japanese wakati](docs/ja/commands/util-text-nlp-japanese-wakati.md)       | 分かち書き(日本語テキストのトークン化)                                 |
| [util tidy move dispatch](docs/ja/commands/util-tidy-move-dispatch.md)                   | ファイルを整理                                                         |
| [util tidy move simple](docs/ja/commands/util-tidy-move-simple.md)                       | ローカルファイルをアーカイブします                                     |
| [util time now](docs/ja/commands/util-time-now.md)                                       | 現在の時刻を表示                                                       |
| [util unixtime format](docs/ja/commands/util-unixtime-format.md)                         | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット |
| [util unixtime now](docs/ja/commands/util-unixtime-now.md)                               | UNIX時間で現在の時刻を表示する                                         |
| [util uuid v4](docs/ja/commands/util-uuid-v4.md)                                         | UUID v4（ランダムUUID）の生成                                          |
| [util video subtitles optimize](docs/ja/commands/util-video-subtitles-optimize.md)       | 字幕ファイルの最適化                                                   |
| [util xlsx create](docs/ja/commands/util-xlsx-create.md)                                 | 空のスプレッドシートを作成する                                         |
| [util xlsx sheet export](docs/ja/commands/util-xlsx-sheet-export.md)                     | xlsxファイルからデータをエクスポート                                   |
| [util xlsx sheet import](docs/ja/commands/util-xlsx-sheet-import.md)                     | データをxlsxファイルにインポート                                       |
| [util xlsx sheet list](docs/ja/commands/util-xlsx-sheet-list.md)                         | xlsxファイルのシート一覧                                               |
| [version](docs/ja/commands/version.md)                                                   | バージョン情報                                                         |

