# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![codecov](https://codecov.io/gh/watermint/toolbox/branch/main/graph/badge.svg?token=CrE8reSVvE)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

Dropbox、Dropbox Business、Google、GitHubなどのWebサービスに対応した多目的ユーティリティ・コマンドラインツール.

# ライセンスと免責条項

watermint toolboxはMITライセンスのもと配布されています.
詳細はファイル LICENSE.mdまたは LICENSE.txt ご参照ください.

以下にご留意ください:
> ソフトウェアは「現状のまま」で、明示であるか暗黙であるかを問わず、何らの保証もなく提供されます。ここでいう保証とは、商品性、特定の目的への適合性、および権利非侵害についての保証も含みますが、それに限定されるものではありません。

# ビルド済み実行ファイル

コンパイル済みバイナリは [最新のリリース](https://github.com/watermint/toolbox/releases/latest) からダウンロードいただけます. ソースコードからビルドする場合には [BUILD.md](BUILD.md) を参照してください.

## macOSでHomebrewを使いインストール

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

© 2016-2021 Takayuki Okazaki
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

| コマンド                                                                   | 説明                                                               |
|----------------------------------------------------------------------------|--------------------------------------------------------------------|
| [file compare account](docs/ja/commands/file-compare-account.md)           | 二つのアカウントのファイルを比較します                             |
| [file compare local](docs/ja/commands/file-compare-local.md)               | ローカルフォルダとDropboxフォルダの内容を比較します                |
| [file copy](docs/ja/commands/file-copy.md)                                 | ファイルをコピーします                                             |
| [file delete](docs/ja/commands/file-delete.md)                             | ファイルまたはフォルダは削除します.                                |
| [file export doc](docs/ja/commands/file-export-doc.md)                     | ドキュメントのエクスポート                                         |
| [file export url](docs/ja/commands/file-export-url.md)                     | URLからドキュメントをエクスポート                                  |
| [file import batch url](docs/ja/commands/file-import-batch-url.md)         | URLからファイルを一括インポートします                              |
| [file import url](docs/ja/commands/file-import-url.md)                     | URLからファイルをインポートします                                  |
| [file info](docs/ja/commands/file-info.md)                                 | パスのメタデータを解決                                             |
| [file list](docs/ja/commands/file-list.md)                                 | ファイルとフォルダを一覧します                                     |
| [file lock acquire](docs/ja/commands/file-lock-acquire.md)                 | ファイルをロック                                                   |
| [file lock all release](docs/ja/commands/file-lock-all-release.md)         | 指定したパスでのすべてのロックを解除する                           |
| [file lock batch acquire](docs/ja/commands/file-lock-batch-acquire.md)     | 複数のファイルをロックする                                         |
| [file lock batch release](docs/ja/commands/file-lock-batch-release.md)     | 複数のロックを解除                                                 |
| [file lock list](docs/ja/commands/file-lock-list.md)                       | 指定したパスの下にあるロックを一覧表示します                       |
| [file lock release](docs/ja/commands/file-lock-release.md)                 | ロックを解除します                                                 |
| [file merge](docs/ja/commands/file-merge.md)                               | フォルダを統合します                                               |
| [file mount list](docs/ja/commands/file-mount-list.md)                     | マウント/アンマウントされた共有フォルダの一覧                      |
| [file move](docs/ja/commands/file-move.md)                                 | ファイルを移動します                                               |
| [file paper append](docs/ja/commands/file-paper-append.md)                 | 既存のPaperドキュメントの最後にコンテンツを追加する                |
| [file paper create](docs/ja/commands/file-paper-create.md)                 | パスに新しいPaperを作成                                            |
| [file paper overwrite](docs/ja/commands/file-paper-overwrite.md)           | 既存のPaperドキュメントを上書きする                                |
| [file paper prepend](docs/ja/commands/file-paper-prepend.md)               | 既存のPaperドキュメントの先頭にコンテンツを追加する                |
| [file replication](docs/ja/commands/file-replication.md)                   | ファイルコンテンツを他のアカウントに複製します                     |
| [file restore all](docs/ja/commands/file-restore-all.md)                   | 指定されたパス以下をリストアします                                 |
| [file search content](docs/ja/commands/file-search-content.md)             | ファイルコンテンツを検索                                           |
| [file search name](docs/ja/commands/file-search-name.md)                   | ファイル名を検索                                                   |
| [file size](docs/ja/commands/file-size.md)                                 | ストレージの利用量                                                 |
| [file sync down](docs/ja/commands/file-sync-down.md)                       | Dropboxと下り方向で同期します                                      |
| [file sync online](docs/ja/commands/file-sync-online.md)                   | オンラインファイルを同期します                                     |
| [file sync up](docs/ja/commands/file-sync-up.md)                           | Dropboxと上り方向で同期します                                      |
| [file watch](docs/ja/commands/file-watch.md)                               | ファイルアクティビティを監視                                       |
| [filerequest create](docs/ja/commands/filerequest-create.md)               | ファイルリクエストを作成します                                     |
| [filerequest delete closed](docs/ja/commands/filerequest-delete-closed.md) | このアカウントの全ての閉じられているファイルリクエストを削除します |
| [filerequest delete url](docs/ja/commands/filerequest-delete-url.md)       | ファイルリクエストのURLを指定して削除                              |
| [filerequest list](docs/ja/commands/filerequest-list.md)                   | 個人アカウントのファイルリクエストを一覧.                          |
| [job history ship](docs/ja/commands/job-history-ship.md)                   | ログの転送先Dropboxパス                                            |
| [sharedfolder list](docs/ja/commands/sharedfolder-list.md)                 | 共有フォルダの一覧                                                 |
| [sharedfolder member list](docs/ja/commands/sharedfolder-member-list.md)   | 共有フォルダのメンバーを一覧します                                 |
| [sharedlink create](docs/ja/commands/sharedlink-create.md)                 | 共有リンクの作成                                                   |
| [sharedlink delete](docs/ja/commands/sharedlink-delete.md)                 | 共有リンクを削除します                                             |
| [sharedlink file list](docs/ja/commands/sharedlink-file-list.md)           | 共有リンクのファイルを一覧する                                     |
| [sharedlink info](docs/ja/commands/sharedlink-info.md)                     | 共有リンクの情報取得                                               |
| [sharedlink list](docs/ja/commands/sharedlink-list.md)                     | 共有リンクの一覧                                                   |

## Dropbox Business

| コマンド                                                                                   | 説明                                                                                   |
|--------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|
| [group add](docs/ja/commands/group-add.md)                                                 | グループを作成します                                                                   |
| [group batch delete](docs/ja/commands/group-batch-delete.md)                               | グループの削除                                                                         |
| [group delete](docs/ja/commands/group-delete.md)                                           | グループを削除します                                                                   |
| [group folder list](docs/ja/commands/group-folder-list.md)                                 | 各グループのフォルダを探す                                                             |
| [group list](docs/ja/commands/group-list.md)                                               | グループを一覧                                                                         |
| [group member add](docs/ja/commands/group-member-add.md)                                   | メンバーをグループに追加                                                               |
| [group member batch add](docs/ja/commands/group-member-batch-add.md)                       | グループにメンバーを一括追加                                                           |
| [group member batch delete](docs/ja/commands/group-member-batch-delete.md)                 | グループからメンバーを削除                                                             |
| [group member batch update](docs/ja/commands/group-member-batch-update.md)                 | グループからメンバーを追加または削除                                                   |
| [group member delete](docs/ja/commands/group-member-delete.md)                             | メンバーをグループから削除                                                             |
| [group member list](docs/ja/commands/group-member-list.md)                                 | グループに所属するメンバー一覧を取得します                                             |
| [group rename](docs/ja/commands/group-rename.md)                                           | グループの改名                                                                         |
| [member clear externalid](docs/ja/commands/member-clear-externalid.md)                     | メンバーのexternal_idを初期化します                                                    |
| [member delete](docs/ja/commands/member-delete.md)                                         | メンバーを削除します                                                                   |
| [member detach](docs/ja/commands/member-detach.md)                                         | Dropbox BusinessユーザーをBasicユーザーに変更します                                    |
| [member file lock all release](docs/ja/commands/member-file-lock-all-release.md)           | メンバーのパスの下にあるすべてのロックを解除します                                     |
| [member file lock list](docs/ja/commands/member-file-lock-list.md)                         | パスの下にあるメンバーのロックを一覧表示                                               |
| [member file lock release](docs/ja/commands/member-file-lock-release.md)                   | メンバーとしてパスのロックを解除します                                                 |
| [member file permdelete](docs/ja/commands/member-file-permdelete.md)                       | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                 |
| [member folder list](docs/ja/commands/member-folder-list.md)                               | 各メンバーのフォルダを検索                                                             |
| [member folder replication](docs/ja/commands/member-folder-replication.md)                 | フォルダを他のメンバーの個人フォルダに複製します                                       |
| [member invite](docs/ja/commands/member-invite.md)                                         | メンバーを招待します                                                                   |
| [member list](docs/ja/commands/member-list.md)                                             | チームメンバーの一覧                                                                   |
| [member quota list](docs/ja/commands/member-quota-list.md)                                 | メンバーの容量制限情報を一覧します                                                     |
| [member quota update](docs/ja/commands/member-quota-update.md)                             | チームメンバーの容量制限を変更                                                         |
| [member quota usage](docs/ja/commands/member-quota-usage.md)                               | チームメンバーのストレージ利用状況を取得                                               |
| [member reinvite](docs/ja/commands/member-reinvite.md)                                     | 招待済み状態メンバーをチームに再招待します                                             |
| [member replication](docs/ja/commands/member-replication.md)                               | チームメンバーのファイルを複製します                                                   |
| [member update email](docs/ja/commands/member-update-email.md)                             | メンバーのメールアドレス処理                                                           |
| [member update externalid](docs/ja/commands/member-update-externalid.md)                   | チームメンバーのExternal IDを更新します.                                               |
| [member update invisible](docs/ja/commands/member-update-invisible.md)                     | メンバーへのディレクトリ制限を有効にします                                             |
| [member update profile](docs/ja/commands/member-update-profile.md)                         | メンバーのプロフィール変更                                                             |
| [member update visible](docs/ja/commands/member-update-visible.md)                         | メンバーへのディレクトリ制限を無効にします                                             |
| [team activity batch user](docs/ja/commands/team-activity-batch-user.md)                   | 複数ユーザーのアクティビティを一括取得します                                           |
| [team activity daily event](docs/ja/commands/team-activity-daily-event.md)                 | アクティビティーを1日ごとに取得します                                                  |
| [team activity event](docs/ja/commands/team-activity-event.md)                             | イベントログ                                                                           |
| [team activity user](docs/ja/commands/team-activity-user.md)                               | ユーザーごとのアクティビティ                                                           |
| [team content member list](docs/ja/commands/team-content-member-list.md)                   | チームフォルダや共有フォルダのメンバー一覧                                             |
| [team content member size](docs/ja/commands/team-content-member-size.md)                   | チームフォルダや共有フォルダのメンバー数をカウントする                                 |
| [team content mount list](docs/ja/commands/team-content-mount-list.md)                     | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします. |
| [team content policy list](docs/ja/commands/team-content-policy-list.md)                   | チームフォルダと共有フォルダのポリシー一覧                                             |
| [team device list](docs/ja/commands/team-device-list.md)                                   | チーム内全てのデバイス/セッションを一覧します                                          |
| [team device unlink](docs/ja/commands/team-device-unlink.md)                               | デバイスのセッションを解除します                                                       |
| [team feature](docs/ja/commands/team-feature.md)                                           | チームの機能を出力します                                                               |
| [team filerequest list](docs/ja/commands/team-filerequest-list.md)                         | チームないのファイルリクエストを一覧します                                             |
| [team info](docs/ja/commands/team-info.md)                                                 | チームの情報                                                                           |
| [team linkedapp list](docs/ja/commands/team-linkedapp-list.md)                             | リンク済みアプリを一覧                                                                 |
| [team namespace file list](docs/ja/commands/team-namespace-file-list.md)                   | チーム内全ての名前空間でのファイル・フォルダを一覧                                     |
| [team namespace file size](docs/ja/commands/team-namespace-file-size.md)                   | チーム内全ての名前空間でのファイル・フォルダを一覧                                     |
| [team namespace list](docs/ja/commands/team-namespace-list.md)                             | チーム内すべての名前空間を一覧                                                         |
| [team namespace member list](docs/ja/commands/team-namespace-member-list.md)               | チームフォルダ以下のファイル・フォルダを一覧                                           |
| [team sharedlink cap expiry](docs/ja/commands/team-sharedlink-cap-expiry.md)               | チーム内の共有リンクに有効期限の上限を設定                                             |
| [team sharedlink cap visibility](docs/ja/commands/team-sharedlink-cap-visibility.md)       | チーム内の共有リンクに可視性の上限を設定                                               |
| [team sharedlink delete links](docs/ja/commands/team-sharedlink-delete-links.md)           | 共有リンクの一括削除                                                                   |
| [team sharedlink delete member](docs/ja/commands/team-sharedlink-delete-member.md)         | メンバーの共有リンクをすべて削除                                                       |
| [team sharedlink list](docs/ja/commands/team-sharedlink-list.md)                           | 共有リンクの一覧                                                                       |
| [team sharedlink update expiry](docs/ja/commands/team-sharedlink-update-expiry.md)         | チーム内の公開されている共有リンクについて有効期限を更新します                         |
| [team sharedlink update password](docs/ja/commands/team-sharedlink-update-password.md)     | 共有リンクのパスワードの設定・更新                                                     |
| [team sharedlink update visibility](docs/ja/commands/team-sharedlink-update-visibility.md) | 共有リンクの可視性の更新                                                               |
| [teamfolder add](docs/ja/commands/teamfolder-add.md)                                       | チームフォルダを追加します                                                             |
| [teamfolder archive](docs/ja/commands/teamfolder-archive.md)                               | チームフォルダのアーカイブ                                                             |
| [teamfolder batch archive](docs/ja/commands/teamfolder-batch-archive.md)                   | 複数のチームフォルダをアーカイブします                                                 |
| [teamfolder batch permdelete](docs/ja/commands/teamfolder-batch-permdelete.md)             | 複数のチームフォルダを完全に削除します                                                 |
| [teamfolder batch replication](docs/ja/commands/teamfolder-batch-replication.md)           | チームフォルダの一括レプリケーション                                                   |
| [teamfolder file list](docs/ja/commands/teamfolder-file-list.md)                           | チームフォルダの一覧                                                                   |
| [teamfolder file lock all release](docs/ja/commands/teamfolder-file-lock-all-release.md)   | チームフォルダのパスの下にあるすべてのロックを解除する                                 |
| [teamfolder file lock list](docs/ja/commands/teamfolder-file-lock-list.md)                 | チームフォルダ内のロックを一覧表示                                                     |
| [teamfolder file lock release](docs/ja/commands/teamfolder-file-lock-release.md)           | チームフォルダ内のパスのロックを解除                                                   |
| [teamfolder file size](docs/ja/commands/teamfolder-file-size.md)                           | チームフォルダのサイズを計算                                                           |
| [teamfolder list](docs/ja/commands/teamfolder-list.md)                                     | チームフォルダの一覧                                                                   |
| [teamfolder member add](docs/ja/commands/teamfolder-member-add.md)                         | チームフォルダへのユーザー/グループの一括追加                                          |
| [teamfolder member delete](docs/ja/commands/teamfolder-member-delete.md)                   | チームフォルダからのユーザー/グループの一括削除                                        |
| [teamfolder member list](docs/ja/commands/teamfolder-member-list.md)                       | チームフォルダのメンバー一覧                                                           |
| [teamfolder partial replication](docs/ja/commands/teamfolder-partial-replication.md)       | 部分的なチームフォルダの他チームへのレプリケーション                                   |
| [teamfolder permdelete](docs/ja/commands/teamfolder-permdelete.md)                         | チームフォルダを完全に削除します                                                       |
| [teamfolder policy list](docs/ja/commands/teamfolder-policy-list.md)                       | チームフォルダのポリシー一覧                                                           |
| [teamfolder replication](docs/ja/commands/teamfolder-replication.md)                       | チームフォルダを他のチームに複製します                                                 |

## GitHub

| コマンド                                                                                             | 説明                                             |
|------------------------------------------------------------------------------------------------------|--------------------------------------------------|
| [services github content get](docs/ja/commands/services-github-content-get.md)                       | レポジトリのコンテンツメタデータを取得します.    |
| [services github content put](docs/ja/commands/services-github-content-put.md)                       | レポジトリに小さなテキストコンテンツを格納します |
| [services github issue list](docs/ja/commands/services-github-issue-list.md)                         | 公開・プライベートGitHubレポジトリの課題一覧     |
| [services github profile](docs/ja/commands/services-github-profile.md)                               | 認証したユーザーの情報を取得                     |
| [services github release asset download](docs/ja/commands/services-github-release-asset-download.md) | アセットをダウンロードします                     |
| [services github release asset list](docs/ja/commands/services-github-release-asset-list.md)         | GitHubリリースの成果物一覧                       |
| [services github release asset upload](docs/ja/commands/services-github-release-asset-upload.md)     | GitHub リリースへ成果物をアップロードします      |
| [services github release draft](docs/ja/commands/services-github-release-draft.md)                   | リリースの下書きを作成                           |
| [services github release list](docs/ja/commands/services-github-release-list.md)                     | リリースの一覧                                   |
| [services github tag create](docs/ja/commands/services-github-tag-create.md)                         | レポジトリにタグを作成します                     |

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
| [services google sheets sheet export](docs/ja/commands/services-google-sheets-sheet-export.md)             | シートデータのエクスポート           |
| [services google sheets sheet import](docs/ja/commands/services-google-sheets-sheet-import.md)             | スプレッドシートにデータをインポート |
| [services google sheets sheet list](docs/ja/commands/services-google-sheets-sheet-list.md)                 | スプレッドシートのシート一覧         |
| [services google sheets spreadsheet create](docs/ja/commands/services-google-sheets-spreadsheet-create.md) | 新しいスプレッドシートの作成         |

## Asana

| コマンド                                                                                           | 説明                             |
|----------------------------------------------------------------------------------------------------|----------------------------------|
| [services asana team list](docs/ja/commands/services-asana-team-list.md)                           | チームのリスト                   |
| [services asana team project list](docs/ja/commands/services-asana-team-project-list.md)           | チームのプロジェクト一覧         |
| [services asana team task list](docs/ja/commands/services-asana-team-task-list.md)                 | チームのタスク一覧               |
| [services asana workspace list](docs/ja/commands/services-asana-workspace-list.md)                 | ワークスペースの一覧             |
| [services asana workspace project list](docs/ja/commands/services-asana-workspace-project-list.md) | ワークスペースのプロジェクト一覧 |

## Slack

| コマンド                                                                                 | 説明           |
|------------------------------------------------------------------------------------------|----------------|
| [services slack conversation list](docs/ja/commands/services-slack-conversation-list.md) | チャネルの一覧 |

## ユーティリティー

| コマンド                                                             | 説明                                                                   |
|----------------------------------------------------------------------|------------------------------------------------------------------------|
| [config disable](docs/ja/commands/config-disable.md)                 | 機能を無効化します.                                                    |
| [config enable](docs/ja/commands/config-enable.md)                   | 機能を有効化します.                                                    |
| [config features](docs/ja/commands/config-features.md)               | 利用可能なオプション機能一覧.                                          |
| [file archive local](docs/ja/commands/file-archive-local.md)         | ローカルファイルをアーカイブします                                     |
| [file dispatch local](docs/ja/commands/file-dispatch-local.md)       | ローカルファイルを整理します                                           |
| [job history archive](docs/ja/commands/job-history-archive.md)       | ジョブのアーカイブ                                                     |
| [job history delete](docs/ja/commands/job-history-delete.md)         | 古いジョブ履歴の削除                                                   |
| [job history list](docs/ja/commands/job-history-list.md)             | ジョブ履歴の表示                                                       |
| [job log jobid](docs/ja/commands/job-log-jobid.md)                   | 指定したジョブIDのログを取得する                                       |
| [job log kind](docs/ja/commands/job-log-kind.md)                     | 指定種別のログを結合して出力します                                     |
| [job log last](docs/ja/commands/job-log-last.md)                     | 最後のジョブのログファイルを出力.                                      |
| [license](docs/ja/commands/license.md)                               | ライセンス情報を表示します                                             |
| [util date today](docs/ja/commands/util-date-today.md)               | 現在の日付を表示                                                       |
| [util datetime now](docs/ja/commands/util-datetime-now.md)           | 現在の日時を表示                                                       |
| [util decode base_32](docs/ja/commands/util-decode-base_32.md)       | Base32 (RFC 4648) 形式からテキストをデコードします                     |
| [util decode base_64](docs/ja/commands/util-decode-base_64.md)       | Base64 (RFC 4648) フォーマットからテキストをデコードします             |
| [util encode base_32](docs/ja/commands/util-encode-base_32.md)       | テキストをBase32(RFC 4648)形式にエンコード                             |
| [util encode base_64](docs/ja/commands/util-encode-base_64.md)       | テキストをBase64(RFC 4648)形式にエンコード                             |
| [util git clone](docs/ja/commands/util-git-clone.md)                 | git リポジトリをクローン                                               |
| [util qrcode create](docs/ja/commands/util-qrcode-create.md)         | QRコード画像ファイルの作成                                             |
| [util qrcode wifi](docs/ja/commands/util-qrcode-wifi.md)             | WIFI設定用のQRコードを生成                                             |
| [util time now](docs/ja/commands/util-time-now.md)                   | 現在の時刻を表示                                                       |
| [util unixtime format](docs/ja/commands/util-unixtime-format.md)     | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット |
| [util unixtime now](docs/ja/commands/util-unixtime-now.md)           | UNIX時間で現在の時刻を表示する                                         |
| [util xlsx create](docs/ja/commands/util-xlsx-create.md)             | 空のスプレッドシートを作成する                                         |
| [util xlsx sheet export](docs/ja/commands/util-xlsx-sheet-export.md) | xlsxファイルからデータをエクスポート                                   |
| [util xlsx sheet import](docs/ja/commands/util-xlsx-sheet-import.md) | データをxlsxファイルにインポート                                       |
| [util xlsx sheet list](docs/ja/commands/util-xlsx-sheet-list.md)     | xlsxファイルのシート一覧                                               |
| [version](docs/ja/commands/version.md)                               | バージョン情報                                                         |

