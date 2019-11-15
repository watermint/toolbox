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

| コマンド                                                                           | 説明                                                           |
|------------------------------------------------------------------------------------|----------------------------------------------------------------|
| [file compare account](doc/generated_ja/file-compare-account.md)                   | 二つのアカウントのファイルを比較します                         |
| [file compare local](doc/generated_ja/file-compare-local.md)                       | {"key":"recipe.file.compare.local.title","params":{}}          |
| [file copy](doc/generated_ja/file-copy.md)                                         | ファイルをコピーします                                         |
| [file import batch url](doc/generated_ja/file-import-batch-url.md)                 | URLからファイルを一括インポートします                          |
| [file import url](doc/generated_ja/file-import-url.md)                             | URLからファイルをインポートします                              |
| [file list](doc/generated_ja/file-list.md)                                         | ファイルとフォルダを一覧します                                 |
| [file merge](doc/generated_ja/file-merge.md)                                       | フォルダを統合します                                           |
| [file move](doc/generated_ja/file-move.md)                                         | ファイルを移動します                                           |
| [file replication](doc/generated_ja/file-replication.md)                           | ファイルコンテンツを他のアカウントに複製します                 |
| [file upload](doc/generated_ja/file-upload.md)                                     | {"key":"recipe.file.upload.title","params":{}}                 |
| [group delete](doc/generated_ja/group-delete.md)                                   | グループを削除します                                           |
| [group list](doc/generated_ja/group-list.md)                                       | グループを一覧                                                 |
| [group member list](doc/generated_ja/group-member-list.md)                         | グループに所属するメンバー一覧を取得します                     |
| [license](doc/generated_ja/license.md)                                             | ライセンス情報を表示します                                     |
| [member delete](doc/generated_ja/member-delete.md)                                 | メンバーを削除します                                           |
| [member detach](doc/generated_ja/member-detach.md)                                 | Dropbox BusinessユーザーをBasicユーザーに変更します            |
| [member invite](doc/generated_ja/member-invite.md)                                 | メンバーを招待します                                           |
| [member list](doc/generated_ja/member-list.md)                                     | チームメンバーの一覧                                           |
| [member quota list](doc/generated_ja/member-quota-list.md)                         | メンバーの容量制限情報を一覧します                             |
| [member quota update](doc/generated_ja/member-quota-update.md)                     | チームメンバーの容量制限を変更                                 |
| [member quota usage](doc/generated_ja/member-quota-usage.md)                       | チームメンバーのストレージ利用状況を取得                       |
| [member update email](doc/generated_ja/member-update-email.md)                     | メンバーのメールアドレス処理                                   |
| [member update profile](doc/generated_ja/member-update-profile.md)                 | メンバーのプロフィール変更                                     |
| [sharedfolder list](doc/generated_ja/sharedfolder-list.md)                         | 共有フォルダの一覧                                             |
| [sharedfolder member list](doc/generated_ja/sharedfolder-member-list.md)           | 共有フォルダのメンバーを一覧します                             |
| [sharedlink create](doc/generated_ja/sharedlink-create.md)                         | 共有リンクの作成                                               |
| [sharedlink delete](doc/generated_ja/sharedlink-delete.md)                         | 共有リンクを削除します                                         |
| [sharedlink list](doc/generated_ja/sharedlink-list.md)                             | 共有リンクの一覧                                               |
| [team activity daily event](doc/generated_ja/team-activity-daily-event.md)         | アクティビティーを1日ごとに取得します                          |
| [team activity event](doc/generated_ja/team-activity-event.md)                     | イベントログ                                                   |
| [team device list](doc/generated_ja/team-device-list.md)                           | チーム内全てのデバイス/セッションを一覧します                  |
| [team device unlink](doc/generated_ja/team-device-unlink.md)                       | デバイスのセッションを解除します                               |
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
| [teamfolder list](doc/generated_ja/teamfolder-list.md)                             | チームフォルダの一覧                                           |
| [teamfolder permdelete](doc/generated_ja/teamfolder-permdelete.md)                 | チームフォルダを完全に削除します                               |
| [teamfolder replication](doc/generated_ja/teamfolder-replication.md)               | チームフォルダを他のチームに複製します                         |
| [web](doc/generated_ja/web.md)                                                     | Webコンソールを起動 (実験的)                                   |
