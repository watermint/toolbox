# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=svg)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/watermint/toolbox)](https://goreportcard.com/report/github.com/watermint/toolbox)

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
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Tools for Dropbox and Dropbox Business

Usage:
./tbx  command

Available commands:
   file          File operation
   group         Group management (Dropbox Business)
   license       Show license information
   member        Team member management (Dropbox Business)
   sharedfolder  Shared folder
   sharedlink    Shared Link of Personal account
   team          Dropbox Business Team
   teamfolder    Team folder management (Dropbox Business)
   web           Launch web console (experimental)
```

## コマンド

| コマンド                                                                        | 説明                                                           |
|---------------------------------------------------------------------------------|----------------------------------------------------------------|
| [file compare account](doc/generated/file-compare-account.md)                   | 二つのアカウントのファイルを比較します                         |
| [file copy](doc/generated/file-copy.md)                                         | ファイルをコピーします                                         |
| [file import batch url](doc/generated/file-import-batch-url.md)                 | URLからファイルを一括インポートします                          |
| [file import url](doc/generated/file-import-url.md)                             | URLからファイルをインポートします                              |
| [file list](doc/generated/file-list.md)                                         | ファイルとフォルダを一覧します                                 |
| [file merge](doc/generated/file-merge.md)                                       | フォルダを統合します                                           |
| [file move](doc/generated/file-move.md)                                         | ファイルを移動します                                           |
| [file replication](doc/generated/file-replication.md)                           | ファイルコンテンツを他のアカウントに複製します                 |
| [group delete](doc/generated/group-delete.md)                                   | グループを削除します                                           |
| [group list](doc/generated/group-list.md)                                       | グループを一覧                                                 |
| [group member list](doc/generated/group-member-list.md)                         | グループに所属するメンバー一覧を取得します                     |
| [license](doc/generated/license.md)                                             | ライセンス情報を表示します                                     |
| [member delete](doc/generated/member-delete.md)                                 | メンバーを削除します                                           |
| [member detach](doc/generated/member-detach.md)                                 | Dropbox BusinessユーザーをBasicユーザーに変更します            |
| [member invite](doc/generated/member-invite.md)                                 | メンバーを招待します                                           |
| [member list](doc/generated/member-list.md)                                     | チームメンバーの一覧                                           |
| [member quota list](doc/generated/member-quota-list.md)                         | メンバーの容量制限情報を一覧します                             |
| [member quota usage](doc/generated/member-quota-usage.md)                       | チームメンバーのストレージ利用状況を取得                       |
| [member update email](doc/generated/member-update-email.md)                     | メンバーのメールアドレス処理                                   |
| [member update profile](doc/generated/member-update-profile.md)                 | メンバーのプロフィール変更                                     |
| [sharedfolder list](doc/generated/sharedfolder-list.md)                         | 共有フォルダの一覧                                             |
| [sharedfolder member list](doc/generated/sharedfolder-member-list.md)           | 共有フォルダのメンバーを一覧します                             |
| [sharedlink create](doc/generated/sharedlink-create.md)                         | 共有リンクの作成                                               |
| [sharedlink delete](doc/generated/sharedlink-delete.md)                         | 共有リンクを削除します                                         |
| [sharedlink list](doc/generated/sharedlink-list.md)                             | 共有リンクの一覧                                               |
| [team activity daily event](doc/generated/team-activity-daily-event.md)         | アクティビティーを1日ごとに取得します                          |
| [team activity event](doc/generated/team-activity-event.md)                     | イベントログ                                                   |
| [team device list](doc/generated/team-device-list.md)                           | チーム内全てのデバイス/セッションを一覧します                  |
| [team device unlink](doc/generated/team-device-unlink.md)                       | デバイスのセッションを解除します                               |
| [team feature](doc/generated/team-feature.md)                                   | チームの機能を出力します                                       |
| [team filerequest list](doc/generated/team-filerequest-list.md)                 | チームないのファイルリクエストを一覧します                     |
| [team info](doc/generated/team-info.md)                                         | チームの情報                                                   |
| [team linkedapp list](doc/generated/team-linkedapp-list.md)                     | リンク済みアプリを一覧                                         |
| [team namespace file list](doc/generated/team-namespace-file-list.md)           | チーム内全ての名前空間でのファイル・フォルダを一覧             |
| [team namespace file size](doc/generated/team-namespace-file-size.md)           | チーム内全ての名前空間でのファイル・フォルダを一覧             |
| [team namespace list](doc/generated/team-namespace-list.md)                     | チーム内すべての名前空間を一覧                                 |
| [team namespace member list](doc/generated/team-namespace-member-list.md)       | チームフォルダ以下のファイル・フォルダを一覧                   |
| [team sharedlink list](doc/generated/team-sharedlink-list.md)                   | 共有リンクの一覧                                               |
| [team sharedlink update expiry](doc/generated/team-sharedlink-update-expiry.md) | チーム内の公開されている共有リンクについて有効期限を更新します |
| [teamfolder archive](doc/generated/teamfolder-archive.md)                       | チームフォルダのアーカイブ                                     |
| [teamfolder list](doc/generated/teamfolder-list.md)                             | チームフォルダの一覧                                           |
| [teamfolder permdelete](doc/generated/teamfolder-permdelete.md)                 | チームフォルダを完全に削除します                               |
| [teamfolder replication](doc/generated/teamfolder-replication.md)               | チームフォルダを他のチームに複製します                         |
| [web](doc/generated/web.md)                                                     | Webコンソールを起動 (実験的)                                   |

