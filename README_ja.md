# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)

![watermint toolbox](resources/watermint-toolbox-256x256.png)

DropboxおよびDropbox Business向けのツールセット

# ライセンスと免責条項

watermint toolboxはMITライセンスのもと配布されています.
詳細はファイル LICENSE.mdまたは LICENSE.txt ご参照ください.

以下にご留意ください:

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

# 利用方法

`tbx` にはたくさんの機能があります. オプションなしで実行をするとサポートされているコマンドやオプションの一覧が表示されます.
つぎのように引数なしで実行すると利用可能なコマンド・オプションがご確認いただけます.

```
% ./tbx

watermint toolbox xx.x.xxx
==========================

© 2016-2020 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

DropboxおよびDropbox Business向けのツールセット

使い方:
=======

./tbx  コマンド

利用可能なコマンド:
===================

   file          ファイル操作                
   group         グループ管理                
   license       ライセンス情報を表示します         
   member        チームメンバーの管理            
   sharedfolder  共有フォルダ                
   sharedlink    個人アカウントの共有リンク         
   team          Dropbox Business チーム  
   teamfolder    チームフォルダの管理            

```

## コマンド

| コマンド                                                                           | 説明                                                           |
|------------------------------------------------------------------------------------|----------------------------------------------------------------|
| [file compare account](doc/generated_ja/file-compare-account.md)                   | 二つのアカウントのファイルを比較します                         |
| [file compare local](doc/generated_ja/file-compare-local.md)                       | ローカルフォルダとDropboxフォルダの内容を比較します            |
| [file copy](doc/generated_ja/file-copy.md)                                         | ファイルをコピーします                                         |
| [file delete](doc/generated_ja/file-delete.md)                                     | ファイルまたはフォルダは削除します.                            |
| [file download](doc/generated_ja/file-download.md)                                 | Dropboxからファイルをダウンロードします                        |
| [file export doc](doc/generated_ja/file-export-doc.md)                             | ドキュメントのエクスポート                                     |
| [file import batch url](doc/generated_ja/file-import-batch-url.md)                 | URLからファイルを一括インポートします                          |
| [file import url](doc/generated_ja/file-import-url.md)                             | URLからファイルをインポートします                              |
| [file list](doc/generated_ja/file-list.md)                                         | ファイルとフォルダを一覧します                                 |
| [file merge](doc/generated_ja/file-merge.md)                                       | フォルダを統合します                                           |
| [file move](doc/generated_ja/file-move.md)                                         | ファイルを移動します                                           |
| [file replication](doc/generated_ja/file-replication.md)                           | ファイルコンテンツを他のアカウントに複製します                 |
| [file restore](doc/generated_ja/file-restore.md)                                   | 指定されたパス以下をリストアします                             |
| [file sync preflight up](doc/generated_ja/file-sync-preflight-up.md)               | 上り方向同期のための事前チェックを実施します                   |
| [file sync up](doc/generated_ja/file-sync-up.md)                                   | Dropboxと上り方向で同期します                                  |
| [file upload](doc/generated_ja/file-upload.md)                                     | ファイルのアップロード                                         |
| [file watch](doc/generated_ja/file-watch.md)                                       | ファイルアクティビティを監視                                   |
| [group add](doc/generated_ja/group-add.md)                                         | グループを作成します                                           |
| [group batch delete](doc/generated_ja/group-batch-delete.md)                       | グループの削除                                                 |
| [group delete](doc/generated_ja/group-delete.md)                                   | グループを削除します                                           |
| [group list](doc/generated_ja/group-list.md)                                       | グループを一覧                                                 |
| [group member add](doc/generated_ja/group-member-add.md)                           | メンバーをグループに追加                                       |
| [group member delete](doc/generated_ja/group-member-delete.md)                     | メンバーをグループから削除                                     |
| [group member list](doc/generated_ja/group-member-list.md)                         | グループに所属するメンバー一覧を取得します                     |
| [group rename](doc/generated_ja/group-rename.md)                                   | グループの改名                                                 |
| [job history archive](doc/generated_ja/job-history-archive.md)                     | ジョブのアーカイブ                                             |
| [job history delete](doc/generated_ja/job-history-delete.md)                       | 古いジョブ履歴の削除                                           |
| [job history list](doc/generated_ja/job-history-list.md)                           | ジョブ履歴の表示                                               |
| [job history ship](doc/generated_ja/job-history-ship.md)                           | ログの転送先Dropboxパス                                        |
| [license](doc/generated_ja/license.md)                                             | ライセンス情報を表示します                                     |
| [member delete](doc/generated_ja/member-delete.md)                                 | メンバーを削除します                                           |
| [member detach](doc/generated_ja/member-detach.md)                                 | Dropbox BusinessユーザーをBasicユーザーに変更します            |
| [member invite](doc/generated_ja/member-invite.md)                                 | メンバーを招待します                                           |
| [member list](doc/generated_ja/member-list.md)                                     | チームメンバーの一覧                                           |
| [member quota list](doc/generated_ja/member-quota-list.md)                         | メンバーの容量制限情報を一覧します                             |
| [member quota update](doc/generated_ja/member-quota-update.md)                     | チームメンバーの容量制限を変更                                 |
| [member quota usage](doc/generated_ja/member-quota-usage.md)                       | チームメンバーのストレージ利用状況を取得                       |
| [member reinvite](doc/generated_ja/member-reinvite.md)                             | 招待済み状態メンバーをチームに再招待します                     |
| [member replication](doc/generated_ja/member-replication.md)                       | チームメンバーのファイルを複製します                           |
| [member update email](doc/generated_ja/member-update-email.md)                     | メンバーのメールアドレス処理                                   |
| [member update externalid](doc/generated_ja/member-update-externalid.md)           | チームメンバーのExternal IDを更新します.                       |
| [member update profile](doc/generated_ja/member-update-profile.md)                 | メンバーのプロフィール変更                                     |
| [sharedfolder list](doc/generated_ja/sharedfolder-list.md)                         | 共有フォルダの一覧                                             |
| [sharedfolder member list](doc/generated_ja/sharedfolder-member-list.md)           | 共有フォルダのメンバーを一覧します                             |
| [sharedlink create](doc/generated_ja/sharedlink-create.md)                         | 共有リンクの作成                                               |
| [sharedlink delete](doc/generated_ja/sharedlink-delete.md)                         | 共有リンクを削除します                                         |
| [sharedlink list](doc/generated_ja/sharedlink-list.md)                             | 共有リンクの一覧                                               |
| [team activity daily event](doc/generated_ja/team-activity-daily-event.md)         | アクティビティーを1日ごとに取得します                          |
| [team activity event](doc/generated_ja/team-activity-event.md)                     | イベントログ                                                   |
| [team activity user](doc/generated_ja/team-activity-user.md)                       | ユーザーごとのアクティビティ                                   |
| [team device list](doc/generated_ja/team-device-list.md)                           | チーム内全てのデバイス/セッションを一覧します                  |
| [team device unlink](doc/generated_ja/team-device-unlink.md)                       | デバイスのセッションを解除します                               |
| [team diag explorer](doc/generated_ja/team-diag-explorer.md)                       | チーム全体の情報をレポートします                               |
| [team feature](doc/generated_ja/team-feature.md)                                   | チームの機能を出力します                                       |
| [team filerequest list](doc/generated_ja/team-filerequest-list.md)                 | チームないのファイルリクエストを一覧します                     |
| [team info](doc/generated_ja/team-info.md)                                         | チームの情報                                                   |
| [team linkedapp list](doc/generated_ja/team-linkedapp-list.md)                     | リンク済みアプリを一覧                                         |
| [team namespace file list](doc/generated_ja/team-namespace-file-list.md)           | チーム内全ての名前空間でのファイル・フォルダを一覧             |
| [team namespace file size](doc/generated_ja/team-namespace-file-size.md)           | チーム内全ての名前空間でのファイル・フォルダを一覧             |
| [team namespace list](doc/generated_ja/team-namespace-list.md)                     | チーム内すべての名前空間を一覧                                 |
| [team namespace member list](doc/generated_ja/team-namespace-member-list.md)       | チームフォルダ以下のファイル・フォルダを一覧                   |
| [team sharedlink list](doc/generated_ja/team-sharedlink-list.md)                   | 共有リンクの一覧                                               |
| [team sharedlink update expiry](doc/generated_ja/team-sharedlink-update-expiry.md) | チーム内の公開されている共有リンクについて有効期限を更新します |
| [teamfolder archive](doc/generated_ja/teamfolder-archive.md)                       | チームフォルダのアーカイブ                                     |
| [teamfolder batch archive](doc/generated_ja/teamfolder-batch-archive.md)           | 複数のチームフォルダをアーカイブします                         |
| [teamfolder batch permdelete](doc/generated_ja/teamfolder-batch-permdelete.md)     | 複数のチームフォルダを完全に削除します                         |
| [teamfolder batch replication](doc/generated_ja/teamfolder-batch-replication.md)   | チームフォルダの一括レプリケーション                           |
| [teamfolder file list](doc/generated_ja/teamfolder-file-list.md)                   | チームフォルダの一覧                                           |
| [teamfolder file size](doc/generated_ja/teamfolder-file-size.md)                   | チームフォルダのサイズを計算                                   |
| [teamfolder list](doc/generated_ja/teamfolder-list.md)                             | チームフォルダの一覧                                           |
| [teamfolder permdelete](doc/generated_ja/teamfolder-permdelete.md)                 | チームフォルダを完全に削除します                               |
| [teamfolder replication](doc/generated_ja/teamfolder-replication.md)               | チームフォルダを他のチームに複製します                         |

