# watermint toolbox

[![Build](https://github.com/watermint/toolbox/actions/workflows/build.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/build.yml)
[![Test](https://github.com/watermint/toolbox/actions/workflows/test.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/test.yml)
[![CodeQL](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml)
[![Codecov](https://codecov.io/gh/watermint/toolbox/branch/main/graph/badge.svg?token=CrE8reSVvE)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

Dropbox、Dropbox for teams、Google、GitHubなどのウェブサービス用の多目的ユーティリティコマンドラインツール。

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

# 製品ライフサイクル

## メンテナンス ポリシー

この製品自体は実験的なものであり、サービスの品質を維持するためのメンテナンスの対象ではありません。プロジェクトは、重大なバグやセキュリティ上の問題を最善の努力で修正するよう努めます。しかし、それは保証されているわけではありません。

この製品は、特定のメジャーリリースのパッチリリースをリリースしません。本製品は、修正が認められた場合、次のリリースとして修正を適用します。

## 仕様変更

このプロジェクトの成果物は、スタンドアロンの実行可能プログラムです。プログラムのバージョンを明示的にアップグレードしない限り、仕様変更は適用されません。

新バージョンのリリースにおける変更は、以下の方針で行われます。

コマンドパス、引数、戻り値などは、可能な限り互換性を保つようにアップグレードされますが、廃止または変更される可能性があります。
一般的な方針は以下の通り。

* 引数の追加やメッセージの変更など、既存の動作を壊さない変更は予告なく実施されます。
* 使用頻度が低いと思われるコマンドは、予告なく廃止または移動されます。
* その他のコマンドの変更は、30～180日以上前に発表されます。

仕様の変更は[お知らせ](https://github.com/watermint/toolbox/discussions/categories/announcements)で発表されます。仕様変更予定一覧は[仕様変更](https://toolbox.watermint.org/ja/guides/spec-change.html)をご参照ください。

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

© 2016-2024 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

Dropbox用ツールとDropbox for teams

使い方:
=======

./tbx  コマンド

利用可能なコマンド:
===================

| コマンド     | 説明                         | 備考 |
|--------------|------------------------------|------|
| asana        | Asanaのコマンド              |      |
| config       | CLI設定                      |      |
| deepl        | DeepLコマンド                |      |
| dropbox      | Dropboxコマンド              |      |
| figma        | フィグマコマンド             |      |
| file         | ファイル操作                 |      |
| filerequest  | ファイルリクエストの操作     |      |
| github       | GitHubコマンド               |      |
| google       | Google コマンド              |      |
| group        | グループ管理                 |      |
| job          | ログユーティリティ（非推奨） |      |
| license      | ライセンス情報を表示します   |      |
| local        | ローカルPC用コマンド         |      |
| log          | ログユーティリティ           |      |
| member       | チームメンバーの管理         |      |
| services     | 各種サービス向けコマンド     |      |
| sharedfolder | 共有フォルダ                 |      |
| sharedlink   | 個人アカウントの共有リンク   |      |
| slack        | Slack コマンド               |      |
| team         | チーム向けDropboxのコマンド  |      |
| teamfolder   | チームフォルダの管理         |      |
| teamspace    | チームスペースコマンド       |      |
| util         | ユーティリティー             |      |
| version      | バージョン情報               |      |

```

# コマンド

## Dropbox (個人アカウント)

| コマンド                                                                                                   | 説明                                                                      |
|------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [dropbox file account feature](docs/ja/commands/dropbox-file-account-feature.md)                           | Dropboxアカウントの機能一覧                                               |
| [dropbox file account filesystem](docs/ja/commands/dropbox-file-account-filesystem.md)                     | Dropboxのファイルシステムのバージョンを表示する                           |
| [dropbox file account info](docs/ja/commands/dropbox-file-account-info.md)                                 | Dropboxアカウント情報                                                     |
| [dropbox file compare account](docs/ja/commands/dropbox-file-compare-account.md)                           | 二つのアカウントのファイルを比較します                                    |
| [dropbox file compare local](docs/ja/commands/dropbox-file-compare-local.md)                               | ローカルフォルダとDropboxフォルダの内容を比較します                       |
| [dropbox file copy](docs/ja/commands/dropbox-file-copy.md)                                                 | ファイルをコピーします                                                    |
| [dropbox file delete](docs/ja/commands/dropbox-file-delete.md)                                             | ファイルまたはフォルダは削除します.                                       |
| [dropbox file export doc](docs/ja/commands/dropbox-file-export-doc.md)                                     | ドキュメントのエクスポート                                                |
| [dropbox file export url](docs/ja/commands/dropbox-file-export-url.md)                                     | URLからドキュメントをエクスポート                                         |
| [dropbox file import batch url](docs/ja/commands/dropbox-file-import-batch-url.md)                         | URLからファイルを一括インポートします                                     |
| [dropbox file import url](docs/ja/commands/dropbox-file-import-url.md)                                     | URLからファイルをインポートします                                         |
| [dropbox file info](docs/ja/commands/dropbox-file-info.md)                                                 | パスのメタデータを解決                                                    |
| [dropbox file list](docs/ja/commands/dropbox-file-list.md)                                                 | ファイルとフォルダを一覧します                                            |
| [dropbox file lock acquire](docs/ja/commands/dropbox-file-lock-acquire.md)                                 | ファイルをロック                                                          |
| [dropbox file lock all release](docs/ja/commands/dropbox-file-lock-all-release.md)                         | 指定したパスでのすべてのロックを解除する                                  |
| [dropbox file lock batch acquire](docs/ja/commands/dropbox-file-lock-batch-acquire.md)                     | 複数のファイルをロックする                                                |
| [dropbox file lock batch release](docs/ja/commands/dropbox-file-lock-batch-release.md)                     | 複数のロックを解除                                                        |
| [dropbox file lock list](docs/ja/commands/dropbox-file-lock-list.md)                                       | 指定したパスの下にあるロックを一覧表示します                              |
| [dropbox file lock release](docs/ja/commands/dropbox-file-lock-release.md)                                 | ロックを解除します                                                        |
| [dropbox file merge](docs/ja/commands/dropbox-file-merge.md)                                               | フォルダを統合します                                                      |
| [dropbox file move](docs/ja/commands/dropbox-file-move.md)                                                 | ファイルを移動します                                                      |
| [dropbox file replication](docs/ja/commands/dropbox-file-replication.md)                                   | ファイルコンテンツを他のアカウントに複製します                            |
| [dropbox file request create](docs/ja/commands/dropbox-file-request-create.md)                             | ファイルリクエストを作成します                                            |
| [dropbox file request delete closed](docs/ja/commands/dropbox-file-request-delete-closed.md)               | このアカウントの全ての閉じられているファイルリクエストを削除します        |
| [dropbox file request delete url](docs/ja/commands/dropbox-file-request-delete-url.md)                     | ファイルリクエストのURLを指定して削除                                     |
| [dropbox file request list](docs/ja/commands/dropbox-file-request-list.md)                                 | 個人アカウントのファイルリクエストを一覧.                                 |
| [dropbox file restore all](docs/ja/commands/dropbox-file-restore-all.md)                                   | 指定されたパス以下をリストアします                                        |
| [dropbox file revision download](docs/ja/commands/dropbox-file-revision-download.md)                       | ファイルリビジョンをダウンロードする                                      |
| [dropbox file revision list](docs/ja/commands/dropbox-file-revision-list.md)                               | ファイルリビジョン一覧                                                    |
| [dropbox file revision restore](docs/ja/commands/dropbox-file-revision-restore.md)                         | ファイルリビジョンを復元する                                              |
| [dropbox file search content](docs/ja/commands/dropbox-file-search-content.md)                             | ファイルコンテンツを検索                                                  |
| [dropbox file search name](docs/ja/commands/dropbox-file-search-name.md)                                   | ファイル名を検索                                                          |
| [dropbox file share info](docs/ja/commands/dropbox-file-share-info.md)                                     | ファイルの共有情報を取得する                                              |
| [dropbox file sharedfolder leave](docs/ja/commands/dropbox-file-sharedfolder-leave.md)                     | 共有フォルダーから退出する.                                               |
| [dropbox file sharedfolder list](docs/ja/commands/dropbox-file-sharedfolder-list.md)                       | 共有フォルダの一覧                                                        |
| [dropbox file sharedfolder member add](docs/ja/commands/dropbox-file-sharedfolder-member-add.md)           | 共有フォルダへのメンバーの追加                                            |
| [dropbox file sharedfolder member delete](docs/ja/commands/dropbox-file-sharedfolder-member-delete.md)     | 共有フォルダからメンバーを削除する                                        |
| [dropbox file sharedfolder member list](docs/ja/commands/dropbox-file-sharedfolder-member-list.md)         | 共有フォルダのメンバーを一覧します                                        |
| [dropbox file sharedfolder mount add](docs/ja/commands/dropbox-file-sharedfolder-mount-add.md)             | 共有フォルダを現在のユーザーのDropboxに追加する                           |
| [dropbox file sharedfolder mount delete](docs/ja/commands/dropbox-file-sharedfolder-mount-delete.md)       | 現在のユーザーが指定されたフォルダーをアンマウントする.                   |
| [dropbox file sharedfolder mount list](docs/ja/commands/dropbox-file-sharedfolder-mount-list.md)           | 現在のユーザーがマウントしているすべての共有フォルダーを一覧表示          |
| [dropbox file sharedfolder mount mountable](docs/ja/commands/dropbox-file-sharedfolder-mount-mountable.md) | 現在のユーザーがマウントできるすべての共有フォルダーをリストアップします. |
| [dropbox file sharedfolder share](docs/ja/commands/dropbox-file-sharedfolder-share.md)                     | フォルダの共有                                                            |
| [dropbox file sharedfolder unshare](docs/ja/commands/dropbox-file-sharedfolder-unshare.md)                 | フォルダの共有解除                                                        |
| [dropbox file sharedlink create](docs/ja/commands/dropbox-file-sharedlink-create.md)                       | 共有リンクの作成                                                          |
| [dropbox file sharedlink delete](docs/ja/commands/dropbox-file-sharedlink-delete.md)                       | 共有リンクを削除します                                                    |
| [dropbox file sharedlink file list](docs/ja/commands/dropbox-file-sharedlink-file-list.md)                 | 共有リンクのファイルを一覧する                                            |
| [dropbox file sharedlink info](docs/ja/commands/dropbox-file-sharedlink-info.md)                           | 共有リンクの情報取得                                                      |
| [dropbox file sharedlink list](docs/ja/commands/dropbox-file-sharedlink-list.md)                           | 共有リンクの一覧                                                          |
| [dropbox file size](docs/ja/commands/dropbox-file-size.md)                                                 | ストレージの利用量                                                        |
| [dropbox file sync down](docs/ja/commands/dropbox-file-sync-down.md)                                       | Dropboxと下り方向で同期します                                             |
| [dropbox file sync online](docs/ja/commands/dropbox-file-sync-online.md)                                   | オンラインファイルを同期します                                            |
| [dropbox file sync up](docs/ja/commands/dropbox-file-sync-up.md)                                           | Dropboxと上り方向で同期します                                             |
| [dropbox file tag add](docs/ja/commands/dropbox-file-tag-add.md)                                           | ファイル/フォルダーにタグを追加する                                       |
| [dropbox file tag delete](docs/ja/commands/dropbox-file-tag-delete.md)                                     | ファイル/フォルダーからタグを削除する                                     |
| [dropbox file tag list](docs/ja/commands/dropbox-file-tag-list.md)                                         | パスのタグを一覧                                                          |
| [dropbox file template apply](docs/ja/commands/dropbox-file-template-apply.md)                             | Dropboxのパスにファイル/フォルダー構造のテンプレートを適用する            |
| [dropbox file template capture](docs/ja/commands/dropbox-file-template-capture.md)                         | Dropboxのパスからファイル/フォルダ構造をテンプレートとして取り込む。      |
| [dropbox file watch](docs/ja/commands/dropbox-file-watch.md)                                               | ファイルアクティビティを監視                                              |
| [dropbox paper append](docs/ja/commands/dropbox-paper-append.md)                                           | 既存のPaperドキュメントの最後にコンテンツを追加する                       |
| [dropbox paper create](docs/ja/commands/dropbox-paper-create.md)                                           | パスに新しいPaperを作成                                                   |
| [dropbox paper overwrite](docs/ja/commands/dropbox-paper-overwrite.md)                                     | 既存のPaperドキュメントを上書きする                                       |
| [dropbox paper prepend](docs/ja/commands/dropbox-paper-prepend.md)                                         | 既存のPaperドキュメントの先頭にコンテンツを追加する                       |
| [log job ship](docs/ja/commands/log-job-ship.md)                                                           | ログの転送先Dropboxパス                                                   |
| [teamspace file list](docs/ja/commands/teamspace-file-list.md)                                             | チームスペースにあるファイルやフォルダーを一覧表示                        |
| [util monitor client](docs/ja/commands/util-monitor-client.md)                                             | デバイスモニタークライアントを起動する                                    |
| [util tidy pack remote](docs/ja/commands/util-tidy-pack-remote.md)                                         | リモートフォルダをZIPファイルにパッケージする                             |

## チーム向けDropbox

| コマンド                                                                                                                       | 説明                                                                                    |
|--------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| [dropbox team activity batch user](docs/ja/commands/dropbox-team-activity-batch-user.md)                                       | 複数ユーザーのアクティビティを一括取得します                                            |
| [dropbox team activity daily event](docs/ja/commands/dropbox-team-activity-daily-event.md)                                     | アクティビティーを1日ごとに取得します                                                   |
| [dropbox team activity event](docs/ja/commands/dropbox-team-activity-event.md)                                                 | イベントログ                                                                            |
| [dropbox team activity user](docs/ja/commands/dropbox-team-activity-user.md)                                                   | ユーザーごとのアクティビティ                                                            |
| [dropbox team admin group role add](docs/ja/commands/dropbox-team-admin-group-role-add.md)                                     | グループのメンバーにロールを追加する                                                    |
| [dropbox team admin group role delete](docs/ja/commands/dropbox-team-admin-group-role-delete.md)                               | 例外グループのメンバーを除くすべてのメンバーからロールを削除する                        |
| [dropbox team admin list](docs/ja/commands/dropbox-team-admin-list.md)                                                         | メンバーの管理者権限一覧                                                                |
| [dropbox team admin role add](docs/ja/commands/dropbox-team-admin-role-add.md)                                                 | メンバーに新しいロールを追加する                                                        |
| [dropbox team admin role clear](docs/ja/commands/dropbox-team-admin-role-clear.md)                                             | メンバーからすべての管理者ロールを削除する                                              |
| [dropbox team admin role delete](docs/ja/commands/dropbox-team-admin-role-delete.md)                                           | メンバーからロールを削除する                                                            |
| [dropbox team admin role list](docs/ja/commands/dropbox-team-admin-role-list.md)                                               | チームの管理者の役割を列挙                                                              |
| [dropbox team backup device status](docs/ja/commands/dropbox-team-backup-device-status.md)                                     | Dropbox バックアップ デバイスのステータスが指定期間内に変更された場合                   |
| [dropbox team content legacypaper count](docs/ja/commands/dropbox-team-content-legacypaper-count.md)                           | メンバー1人あたりのPaper文書の枚数                                                      |
| [dropbox team content legacypaper export](docs/ja/commands/dropbox-team-content-legacypaper-export.md)                         | チームメンバー全員のPaper文書をローカルパスにエクスポート.                              |
| [dropbox team content legacypaper list](docs/ja/commands/dropbox-team-content-legacypaper-list.md)                             | チームメンバーのPaper文書リスト出力                                                     |
| [dropbox team content member list](docs/ja/commands/dropbox-team-content-member-list.md)                                       | チームフォルダや共有フォルダのメンバー一覧                                              |
| [dropbox team content member size](docs/ja/commands/dropbox-team-content-member-size.md)                                       | チームフォルダや共有フォルダのメンバー数をカウントする                                  |
| [dropbox team content mount list](docs/ja/commands/dropbox-team-content-mount-list.md)                                         | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします.  |
| [dropbox team content policy list](docs/ja/commands/dropbox-team-content-policy-list.md)                                       | チームフォルダと共有フォルダのポリシー一覧                                              |
| [dropbox team device list](docs/ja/commands/dropbox-team-device-list.md)                                                       | チーム内全てのデバイス/セッションを一覧します                                           |
| [dropbox team device unlink](docs/ja/commands/dropbox-team-device-unlink.md)                                                   | デバイスのセッションを解除します                                                        |
| [dropbox team feature](docs/ja/commands/dropbox-team-feature.md)                                                               | チームの機能を出力します                                                                |
| [dropbox team filerequest list](docs/ja/commands/dropbox-team-filerequest-list.md)                                             | チームないのファイルリクエストを一覧します                                              |
| [dropbox team filesystem](docs/ja/commands/dropbox-team-filesystem.md)                                                         | チームのファイルシステムのバージョンを特定する                                          |
| [dropbox team group add](docs/ja/commands/dropbox-team-group-add.md)                                                           | グループを作成します                                                                    |
| [dropbox team group batch add](docs/ja/commands/dropbox-team-group-batch-add.md)                                               | グループの一括追加                                                                      |
| [dropbox team group batch delete](docs/ja/commands/dropbox-team-group-batch-delete.md)                                         | グループの削除                                                                          |
| [dropbox team group clear externalid](docs/ja/commands/dropbox-team-group-clear-externalid.md)                                 | グループの外部IDをクリアする                                                            |
| [dropbox team group delete](docs/ja/commands/dropbox-team-group-delete.md)                                                     | グループを削除します                                                                    |
| [dropbox team group folder list](docs/ja/commands/dropbox-team-group-folder-list.md)                                           | 各グループのフォルダーを一覧表示                                                        |
| [dropbox team group list](docs/ja/commands/dropbox-team-group-list.md)                                                         | グループを一覧                                                                          |
| [dropbox team group member add](docs/ja/commands/dropbox-team-group-member-add.md)                                             | メンバーをグループに追加                                                                |
| [dropbox team group member batch add](docs/ja/commands/dropbox-team-group-member-batch-add.md)                                 | グループにメンバーを一括追加                                                            |
| [dropbox team group member batch delete](docs/ja/commands/dropbox-team-group-member-batch-delete.md)                           | グループからメンバーを削除                                                              |
| [dropbox team group member batch update](docs/ja/commands/dropbox-team-group-member-batch-update.md)                           | グループからメンバーを追加または削除                                                    |
| [dropbox team group member delete](docs/ja/commands/dropbox-team-group-member-delete.md)                                       | メンバーをグループから削除                                                              |
| [dropbox team group member list](docs/ja/commands/dropbox-team-group-member-list.md)                                           | グループに所属するメンバー一覧を取得します                                              |
| [dropbox team group rename](docs/ja/commands/dropbox-team-group-rename.md)                                                     | グループの改名                                                                          |
| [dropbox team group update type](docs/ja/commands/dropbox-team-group-update-type.md)                                           | グループ管理タイプの更新                                                                |
| [dropbox team info](docs/ja/commands/dropbox-team-info.md)                                                                     | チームの情報                                                                            |
| [dropbox team insight scan](docs/ja/commands/dropbox-team-insight-scan.md)                                                     | チームデータをスキャンして分析                                                          |
| [dropbox team legalhold add](docs/ja/commands/dropbox-team-legalhold-add.md)                                                   | 新しいリーガル・ホールド・ポリシーを作成する.                                           |
| [dropbox team legalhold list](docs/ja/commands/dropbox-team-legalhold-list.md)                                                 | 既存のポリシーを取得する                                                                |
| [dropbox team legalhold member batch update](docs/ja/commands/dropbox-team-legalhold-member-batch-update.md)                   | リーガル・ホールド・ポリシーのメンバーリスト更新                                        |
| [dropbox team legalhold member list](docs/ja/commands/dropbox-team-legalhold-member-list.md)                                   | リーガルホールドのメンバーをリストアップ                                                |
| [dropbox team legalhold release](docs/ja/commands/dropbox-team-legalhold-release.md)                                           | Idによるリーガルホールドを解除する                                                      |
| [dropbox team legalhold revision list](docs/ja/commands/dropbox-team-legalhold-revision-list.md)                               | リーガル・ホールド・ポリシーのリビジョンをリストアップする                              |
| [dropbox team legalhold update desc](docs/ja/commands/dropbox-team-legalhold-update-desc.md)                                   | リーガルホールド・ポリシーの説明を更新                                                  |
| [dropbox team legalhold update name](docs/ja/commands/dropbox-team-legalhold-update-name.md)                                   | リーガルホールドポリシーの名称を更新                                                    |
| [dropbox team linkedapp list](docs/ja/commands/dropbox-team-linkedapp-list.md)                                                 | リンク済みアプリを一覧                                                                  |
| [dropbox team member batch delete](docs/ja/commands/dropbox-team-member-batch-delete.md)                                       | メンバーを削除します                                                                    |
| [dropbox team member batch detach](docs/ja/commands/dropbox-team-member-batch-detach.md)                                       | Dropbox for teamsのアカウントをBasicアカウントに変更する                                |
| [dropbox team member batch invite](docs/ja/commands/dropbox-team-member-batch-invite.md)                                       | メンバーを招待します                                                                    |
| [dropbox team member batch reinvite](docs/ja/commands/dropbox-team-member-batch-reinvite.md)                                   | 招待済み状態メンバーをチームに再招待します                                              |
| [dropbox team member batch suspend](docs/ja/commands/dropbox-team-member-batch-suspend.md)                                     | メンバーの一括一時停止                                                                  |
| [dropbox team member batch unsuspend](docs/ja/commands/dropbox-team-member-batch-unsuspend.md)                                 | メンバーの一括停止解除                                                                  |
| [dropbox team member clear externalid](docs/ja/commands/dropbox-team-member-clear-externalid.md)                               | メンバーのexternal_idを初期化します                                                     |
| [dropbox team member feature](docs/ja/commands/dropbox-team-member-feature.md)                                                 | メンバーの機能設定一覧                                                                  |
| [dropbox team member file lock all release](docs/ja/commands/dropbox-team-member-file-lock-all-release.md)                     | メンバーのパスの下にあるすべてのロックを解除します                                      |
| [dropbox team member file lock list](docs/ja/commands/dropbox-team-member-file-lock-list.md)                                   | パスの下にあるメンバーのロックを一覧表示                                                |
| [dropbox team member file lock release](docs/ja/commands/dropbox-team-member-file-lock-release.md)                             | メンバーとしてパスのロックを解除します                                                  |
| [dropbox team member file permdelete](docs/ja/commands/dropbox-team-member-file-permdelete.md)                                 | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します                  |
| [dropbox team member folder list](docs/ja/commands/dropbox-team-member-folder-list.md)                                         | 各メンバーのフォルダーを一覧表示                                                        |
| [dropbox team member folder replication](docs/ja/commands/dropbox-team-member-folder-replication.md)                           | フォルダを他のメンバーの個人フォルダに複製します                                        |
| [dropbox team member list](docs/ja/commands/dropbox-team-member-list.md)                                                       | チームメンバーの一覧                                                                    |
| [dropbox team member quota batch update](docs/ja/commands/dropbox-team-member-quota-batch-update.md)                           | チームメンバーの容量制限を変更                                                          |
| [dropbox team member quota list](docs/ja/commands/dropbox-team-member-quota-list.md)                                           | メンバーの容量制限情報を一覧します                                                      |
| [dropbox team member quota usage](docs/ja/commands/dropbox-team-member-quota-usage.md)                                         | チームメンバーのストレージ利用状況を取得                                                |
| [dropbox team member replication](docs/ja/commands/dropbox-team-member-replication.md)                                         | チームメンバーのファイルを複製します                                                    |
| [dropbox team member suspend](docs/ja/commands/dropbox-team-member-suspend.md)                                                 | メンバーの一時停止処理                                                                  |
| [dropbox team member unsuspend](docs/ja/commands/dropbox-team-member-unsuspend.md)                                             | メンバーの一時停止を解除する                                                            |
| [dropbox team member update batch email](docs/ja/commands/dropbox-team-member-update-batch-email.md)                           | メンバーのメールアドレス処理                                                            |
| [dropbox team member update batch externalid](docs/ja/commands/dropbox-team-member-update-batch-externalid.md)                 | チームメンバーのExternal IDを更新します.                                                |
| [dropbox team member update batch invisible](docs/ja/commands/dropbox-team-member-update-batch-invisible.md)                   | メンバーへのディレクトリ制限を有効にします                                              |
| [dropbox team member update batch profile](docs/ja/commands/dropbox-team-member-update-batch-profile.md)                       | メンバーのプロフィール変更                                                              |
| [dropbox team member update batch visible](docs/ja/commands/dropbox-team-member-update-batch-visible.md)                       | メンバーへのディレクトリ制限を無効にします                                              |
| [dropbox team namespace file list](docs/ja/commands/dropbox-team-namespace-file-list.md)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [dropbox team namespace file size](docs/ja/commands/dropbox-team-namespace-file-size.md)                                       | チーム内全ての名前空間でのファイル・フォルダを一覧                                      |
| [dropbox team namespace list](docs/ja/commands/dropbox-team-namespace-list.md)                                                 | チーム内すべての名前空間を一覧                                                          |
| [dropbox team namespace member list](docs/ja/commands/dropbox-team-namespace-member-list.md)                                   | チームフォルダ以下のファイル・フォルダを一覧                                            |
| [dropbox team namespace summary](docs/ja/commands/dropbox-team-namespace-summary.md)                                           | チーム・ネームスペースの状態概要を報告する.                                             |
| [dropbox team runas file batch copy](docs/ja/commands/dropbox-team-runas-file-batch-copy.md)                                   | ファイル/フォルダーをメンバーとして一括コピー                                           |
| [dropbox team runas file list](docs/ja/commands/dropbox-team-runas-file-list.md)                                               | メンバーとして実行するファイルやフォルダーの一覧                                        |
| [dropbox team runas file sync batch up](docs/ja/commands/dropbox-team-runas-file-sync-batch-up.md)                             | メンバーとして動作する一括同期                                                          |
| [dropbox team runas sharedfolder batch leave](docs/ja/commands/dropbox-team-runas-sharedfolder-batch-leave.md)                 | 共有フォルダからメンバーとして一括退出                                                  |
| [dropbox team runas sharedfolder batch share](docs/ja/commands/dropbox-team-runas-sharedfolder-batch-share.md)                 | メンバーのフォルダを一括で共有                                                          |
| [dropbox team runas sharedfolder batch unshare](docs/ja/commands/dropbox-team-runas-sharedfolder-batch-unshare.md)             | メンバーのフォルダの共有を一括解除                                                      |
| [dropbox team runas sharedfolder isolate](docs/ja/commands/dropbox-team-runas-sharedfolder-isolate.md)                         | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| [dropbox team runas sharedfolder list](docs/ja/commands/dropbox-team-runas-sharedfolder-list.md)                               | 共有フォルダーの一覧をメンバーとして実行                                                |
| [dropbox team runas sharedfolder member batch add](docs/ja/commands/dropbox-team-runas-sharedfolder-member-batch-add.md)       | メンバーの共有フォルダにメンバーを一括追加                                              |
| [dropbox team runas sharedfolder member batch delete](docs/ja/commands/dropbox-team-runas-sharedfolder-member-batch-delete.md) | メンバーの共有フォルダからメンバーを一括削除                                            |
| [dropbox team runas sharedfolder mount add](docs/ja/commands/dropbox-team-runas-sharedfolder-mount-add.md)                     | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| [dropbox team runas sharedfolder mount delete](docs/ja/commands/dropbox-team-runas-sharedfolder-mount-delete.md)               | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| [dropbox team runas sharedfolder mount list](docs/ja/commands/dropbox-team-runas-sharedfolder-mount-list.md)                   | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| [dropbox team runas sharedfolder mount mountable](docs/ja/commands/dropbox-team-runas-sharedfolder-mount-mountable.md)         | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| [dropbox team sharedlink cap expiry](docs/ja/commands/dropbox-team-sharedlink-cap-expiry.md)                                   | チーム内の共有リンクに有効期限の上限を設定                                              |
| [dropbox team sharedlink cap visibility](docs/ja/commands/dropbox-team-sharedlink-cap-visibility.md)                           | チーム内の共有リンクに可視性の上限を設定                                                |
| [dropbox team sharedlink delete links](docs/ja/commands/dropbox-team-sharedlink-delete-links.md)                               | 共有リンクの一括削除                                                                    |
| [dropbox team sharedlink delete member](docs/ja/commands/dropbox-team-sharedlink-delete-member.md)                             | メンバーの共有リンクをすべて削除                                                        |
| [dropbox team sharedlink list](docs/ja/commands/dropbox-team-sharedlink-list.md)                                               | 共有リンクの一覧                                                                        |
| [dropbox team sharedlink update expiry](docs/ja/commands/dropbox-team-sharedlink-update-expiry.md)                             | チーム内の公開されている共有リンクについて有効期限を更新します                          |
| [dropbox team sharedlink update password](docs/ja/commands/dropbox-team-sharedlink-update-password.md)                         | 共有リンクのパスワードの設定・更新                                                      |
| [dropbox team sharedlink update visibility](docs/ja/commands/dropbox-team-sharedlink-update-visibility.md)                     | 共有リンクの可視性の更新                                                                |
| [dropbox team teamfolder add](docs/ja/commands/dropbox-team-teamfolder-add.md)                                                 | チームフォルダを追加します                                                              |
| [dropbox team teamfolder archive](docs/ja/commands/dropbox-team-teamfolder-archive.md)                                         | チームフォルダのアーカイブ                                                              |
| [dropbox team teamfolder batch archive](docs/ja/commands/dropbox-team-teamfolder-batch-archive.md)                             | 複数のチームフォルダをアーカイブします                                                  |
| [dropbox team teamfolder batch permdelete](docs/ja/commands/dropbox-team-teamfolder-batch-permdelete.md)                       | 複数のチームフォルダを完全に削除します                                                  |
| [dropbox team teamfolder batch replication](docs/ja/commands/dropbox-team-teamfolder-batch-replication.md)                     | チームフォルダの一括レプリケーション                                                    |
| [dropbox team teamfolder file list](docs/ja/commands/dropbox-team-teamfolder-file-list.md)                                     | チームフォルダの一覧                                                                    |
| [dropbox team teamfolder file lock all release](docs/ja/commands/dropbox-team-teamfolder-file-lock-all-release.md)             | チームフォルダのパスの下にあるすべてのロックを解除する                                  |
| [dropbox team teamfolder file lock list](docs/ja/commands/dropbox-team-teamfolder-file-lock-list.md)                           | チームフォルダ内のロックを一覧表示                                                      |
| [dropbox team teamfolder file lock release](docs/ja/commands/dropbox-team-teamfolder-file-lock-release.md)                     | チームフォルダ内のパスのロックを解除                                                    |
| [dropbox team teamfolder file size](docs/ja/commands/dropbox-team-teamfolder-file-size.md)                                     | チームフォルダのサイズを計算                                                            |
| [dropbox team teamfolder list](docs/ja/commands/dropbox-team-teamfolder-list.md)                                               | チームフォルダの一覧                                                                    |
| [dropbox team teamfolder member add](docs/ja/commands/dropbox-team-teamfolder-member-add.md)                                   | チームフォルダへのユーザー/グループの一括追加                                           |
| [dropbox team teamfolder member delete](docs/ja/commands/dropbox-team-teamfolder-member-delete.md)                             | チームフォルダからのユーザー/グループの一括削除                                         |
| [dropbox team teamfolder member list](docs/ja/commands/dropbox-team-teamfolder-member-list.md)                                 | チームフォルダのメンバー一覧                                                            |
| [dropbox team teamfolder partial replication](docs/ja/commands/dropbox-team-teamfolder-partial-replication.md)                 | 部分的なチームフォルダの他チームへのレプリケーション                                    |
| [dropbox team teamfolder permdelete](docs/ja/commands/dropbox-team-teamfolder-permdelete.md)                                   | チームフォルダを完全に削除します                                                        |
| [dropbox team teamfolder policy list](docs/ja/commands/dropbox-team-teamfolder-policy-list.md)                                 | チームフォルダのポリシー一覧                                                            |
| [dropbox team teamfolder replication](docs/ja/commands/dropbox-team-teamfolder-replication.md)                                 | チームフォルダを他のチームに複製します                                                  |
| [dropbox team teamfolder sync setting list](docs/ja/commands/dropbox-team-teamfolder-sync-setting-list.md)                     | チームフォルダーの同期設定を一覧表示                                                    |
| [dropbox team teamfolder sync setting update](docs/ja/commands/dropbox-team-teamfolder-sync-setting-update.md)                 | チームフォルダ同期設定の一括更新                                                        |
| [teamspace asadmin file list](docs/ja/commands/teamspace-asadmin-file-list.md)                                                 | チームスペース内のファイルやフォルダーを一覧表示することができます。                    |
| [teamspace asadmin folder add](docs/ja/commands/teamspace-asadmin-folder-add.md)                                               | チームスペースにトップレベルのフォルダーを作成                                          |
| [teamspace asadmin folder delete](docs/ja/commands/teamspace-asadmin-folder-delete.md)                                         | チームスペースのトップレベルフォルダーを削除する                                        |
| [teamspace asadmin folder permdelete](docs/ja/commands/teamspace-asadmin-folder-permdelete.md)                                 | チームスペースのトップレベルフォルダを完全に削除します。                                |

## DeepL

| コマンド                                                         | 説明               |
|------------------------------------------------------------------|--------------------|
| [deepl translate text](docs/ja/commands/deepl-translate-text.md) | テキストを翻訳する |

## Figma

| コマンド                                                                     | 説明                                                  |
|------------------------------------------------------------------------------|-------------------------------------------------------|
| [figma account info](docs/ja/commands/figma-account-info.md)                 | 現在のユーザー情報を取得する                          |
| [figma file export all page](docs/ja/commands/figma-file-export-all-page.md) | チーム配下のすべてのファイル/ページをエクスポートする |
| [figma file export frame](docs/ja/commands/figma-file-export-frame.md)       | Figmaファイルの全フレームを書き出す                   |
| [figma file export node](docs/ja/commands/figma-file-export-node.md)         | Figmaドキュメント・ノードの書き出し                   |
| [figma file export page](docs/ja/commands/figma-file-export-page.md)         | Figmaファイルの全ページを書き出す                     |
| [figma file info](docs/ja/commands/figma-file-info.md)                       | figmaファイルの情報を表示する                         |
| [figma file list](docs/ja/commands/figma-file-list.md)                       | Figmaプロジェクト内のファイル一覧                     |
| [figma project list](docs/ja/commands/figma-project-list.md)                 | チームのプロジェクト一覧                              |

## GitHub

| コマンド                                                                           | 説明                                                          |
|------------------------------------------------------------------------------------|---------------------------------------------------------------|
| [github content get](docs/ja/commands/github-content-get.md)                       | レポジトリのコンテンツメタデータを取得します.                 |
| [github content put](docs/ja/commands/github-content-put.md)                       | レポジトリに小さなテキストコンテンツを格納します              |
| [github issue list](docs/ja/commands/github-issue-list.md)                         | 公開・プライベートGitHubレポジトリの課題一覧                  |
| [github profile](docs/ja/commands/github-profile.md)                               | 認証したユーザーの情報を取得                                  |
| [github release asset download](docs/ja/commands/github-release-asset-download.md) | アセットをダウンロードします                                  |
| [github release asset list](docs/ja/commands/github-release-asset-list.md)         | GitHubリリースの成果物一覧                                    |
| [github release asset upload](docs/ja/commands/github-release-asset-upload.md)     | GitHub リリースへ成果物をアップロードします                   |
| [github release draft](docs/ja/commands/github-release-draft.md)                   | リリースの下書きを作成                                        |
| [github release list](docs/ja/commands/github-release-list.md)                     | リリースの一覧                                                |
| [github tag create](docs/ja/commands/github-tag-create.md)                         | レポジトリにタグを作成します                                  |
| [util release install](docs/ja/commands/util-release-install.md)                   | watermint toolboxをダウンロードし、パスにインストールします。 |

## Google Calendar

| コマンド                                                                     | 説明                                 |
|------------------------------------------------------------------------------|--------------------------------------|
| [google calendar event list](docs/ja/commands/google-calendar-event-list.md) | Googleカレンダーのイベントを一覧表示 |

## Google Gmail

| コマンド                                                                                     | 説明                                               |
|----------------------------------------------------------------------------------------------|----------------------------------------------------|
| [google mail filter add](docs/ja/commands/google-mail-filter-add.md)                         | フィルターを追加します.                            |
| [google mail filter batch add](docs/ja/commands/google-mail-filter-batch-add.md)             | クエリによるラベルの一括追加・削除                 |
| [google mail filter delete](docs/ja/commands/google-mail-filter-delete.md)                   | フィルタの削除                                     |
| [google mail filter list](docs/ja/commands/google-mail-filter-list.md)                       | フィルターの一覧                                   |
| [google mail label add](docs/ja/commands/google-mail-label-add.md)                           | ラベルの追加                                       |
| [google mail label delete](docs/ja/commands/google-mail-label-delete.md)                     | ラベルの削除.                                      |
| [google mail label list](docs/ja/commands/google-mail-label-list.md)                         | ラベルのリスト                                     |
| [google mail label rename](docs/ja/commands/google-mail-label-rename.md)                     | ラベルの名前を変更する                             |
| [google mail message label add](docs/ja/commands/google-mail-message-label-add.md)           | メッセージにラベルを追加                           |
| [google mail message label delete](docs/ja/commands/google-mail-message-label-delete.md)     | メッセージからラベルを削除する                     |
| [google mail message list](docs/ja/commands/google-mail-message-list.md)                     | メッセージの一覧                                   |
| [google mail message processed list](docs/ja/commands/google-mail-message-processed-list.md) | 処理された形式でメッセージを一覧表示します.        |
| [google mail sendas add](docs/ja/commands/google-mail-sendas-add.md)                         | カスタムの "from" send-asエイリアスの作成          |
| [google mail sendas delete](docs/ja/commands/google-mail-sendas-delete.md)                   | 指定したsend-asエイリアスを削除する                |
| [google mail sendas list](docs/ja/commands/google-mail-sendas-list.md)                       | 指定されたアカウントの送信エイリアスを一覧表示する |
| [google mail thread list](docs/ja/commands/google-mail-thread-list.md)                       | スレッド一覧                                       |

## Google Sheets

| コマンド                                                                                 | 説明                                 |
|------------------------------------------------------------------------------------------|--------------------------------------|
| [google sheets sheet append](docs/ja/commands/google-sheets-sheet-append.md)             | スプレッドシートにデータを追加する   |
| [google sheets sheet clear](docs/ja/commands/google-sheets-sheet-clear.md)               | スプレッドシートから値をクリアする   |
| [google sheets sheet create](docs/ja/commands/google-sheets-sheet-create.md)             | 新規シートの作成                     |
| [google sheets sheet delete](docs/ja/commands/google-sheets-sheet-delete.md)             | スプレッドシートからシートを削除する |
| [google sheets sheet export](docs/ja/commands/google-sheets-sheet-export.md)             | シートデータのエクスポート           |
| [google sheets sheet import](docs/ja/commands/google-sheets-sheet-import.md)             | スプレッドシートにデータをインポート |
| [google sheets sheet list](docs/ja/commands/google-sheets-sheet-list.md)                 | スプレッドシートのシート一覧         |
| [google sheets spreadsheet create](docs/ja/commands/google-sheets-spreadsheet-create.md) | 新しいスプレッドシートの作成         |

## ユーティリティー

| コマンド                                                                                                         | 説明                                                                   |
|------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [config auth delete](docs/ja/commands/config-auth-delete.md)                                                     | 既存の認証クレデンシャルの削除                                         |
| [config auth list](docs/ja/commands/config-auth-list.md)                                                         | すべての認証情報を一覧表示                                             |
| [config feature disable](docs/ja/commands/config-feature-disable.md)                                             | 機能を無効化します.                                                    |
| [config feature enable](docs/ja/commands/config-feature-enable.md)                                               | 機能を有効化します.                                                    |
| [config feature list](docs/ja/commands/config-feature-list.md)                                                   | 利用可能なオプション機能一覧.                                          |
| [dropbox team insight report teamfoldermember](docs/ja/commands/dropbox-team-insight-report-teamfoldermember.md) | チームフォルダーメンバーを報告                                         |
| [license](docs/ja/commands/license.md)                                                                           | ライセンス情報を表示します                                             |
| [local file template apply](docs/ja/commands/local-file-template-apply.md)                                       | ファイル/フォルダー構造のテンプレートをローカルパスに適用する          |
| [local file template capture](docs/ja/commands/local-file-template-capture.md)                                   | ローカルパスからファイル/フォルダ構造をテンプレートとして取り込む      |
| [log cat curl](docs/ja/commands/log-cat-curl.md)                                                                 | キャプチャログを `curl` サンプルとしてフォーマットする                 |
| [log cat job](docs/ja/commands/log-cat-job.md)                                                                   | 指定したジョブIDのログを取得する                                       |
| [log cat kind](docs/ja/commands/log-cat-kind.md)                                                                 | 指定種別のログを結合して出力します                                     |
| [log cat last](docs/ja/commands/log-cat-last.md)                                                                 | 最後のジョブのログファイルを出力.                                      |
| [log job archive](docs/ja/commands/log-job-archive.md)                                                           | ジョブのアーカイブ                                                     |
| [log job delete](docs/ja/commands/log-job-delete.md)                                                             | 古いジョブ履歴の削除                                                   |
| [log job list](docs/ja/commands/log-job-list.md)                                                                 | ジョブ履歴の表示                                                       |
| [util archive unzip](docs/ja/commands/util-archive-unzip.md)                                                     | ZIPアーカイブファイルを解凍する                                        |
| [util archive zip](docs/ja/commands/util-archive-zip.md)                                                         | 対象ファイルをZIPアーカイブに圧縮する                                  |
| [util cert selfsigned](docs/ja/commands/util-cert-selfsigned.md)                                                 | 自己署名証明書と鍵の生成                                               |
| [util database exec](docs/ja/commands/util-database-exec.md)                                                     | SQLite3データベースファイルへのクエリ実行                              |
| [util database query](docs/ja/commands/util-database-query.md)                                                   | SQLite3データベースへの問い合わせ                                      |
| [util date today](docs/ja/commands/util-date-today.md)                                                           | 現在の日付を表示                                                       |
| [util datetime now](docs/ja/commands/util-datetime-now.md)                                                       | 現在の日時を表示                                                       |
| [util decode base32](docs/ja/commands/util-decode-base32.md)                                                     | Base32 (RFC 4648) 形式からテキストをデコードします                     |
| [util decode base64](docs/ja/commands/util-decode-base64.md)                                                     | Base64 (RFC 4648) フォーマットからテキストをデコードします             |
| [util desktop display list](docs/ja/commands/util-desktop-display-list.md)                                       | このマシンのディスプレイを一覧表示                                     |
| [util desktop open](docs/ja/commands/util-desktop-open.md)                                                       | デフォルトのアプリケーションでファイルやフォルダを開く                 |
| [util desktop screenshot interval](docs/ja/commands/util-desktop-screenshot-interval.md)                         | 定期的にスクリーンショットを撮る                                       |
| [util desktop screenshot snap](docs/ja/commands/util-desktop-screenshot-snap.md)                                 | スクリーンショットを撮る                                               |
| [util encode base32](docs/ja/commands/util-encode-base32.md)                                                     | テキストをBase32(RFC 4648)形式にエンコード                             |
| [util encode base64](docs/ja/commands/util-encode-base64.md)                                                     | テキストをBase64(RFC 4648)形式にエンコード                             |
| [util file hash](docs/ja/commands/util-file-hash.md)                                                             | ファイルダイジェストの表示                                             |
| [util git clone](docs/ja/commands/util-git-clone.md)                                                             | git リポジトリをクローン                                               |
| [util image exif](docs/ja/commands/util-image-exif.md)                                                           | 画像ファイルのEXIFメタデータを表示                                     |
| [util image placeholder](docs/ja/commands/util-image-placeholder.md)                                             | プレースホルダー画像の作成                                             |
| [util net download](docs/ja/commands/util-net-download.md)                                                       | ファイルをダウンロードする                                             |
| [util qrcode create](docs/ja/commands/util-qrcode-create.md)                                                     | QRコード画像ファイルの作成                                             |
| [util qrcode wifi](docs/ja/commands/util-qrcode-wifi.md)                                                         | WIFI設定用のQRコードを生成                                             |
| [util table format xlsx](docs/ja/commands/util-table-format-xlsx.md)                                             | xlsxファイルをテキストに整形する                                       |
| [util text case down](docs/ja/commands/util-text-case-down.md)                                                   | 小文字のテキストを表示する                                             |
| [util text case up](docs/ja/commands/util-text-case-up.md)                                                       | 大文字のテキストを表示する                                             |
| [util text encoding from](docs/ja/commands/util-text-encoding-from.md)                                           | 指定されたエンコーディングからUTF-8テキストファイルに変換します.       |
| [util text encoding to](docs/ja/commands/util-text-encoding-to.md)                                               | UTF-8テキストファイルから指定されたエンコーディングに変換する.         |
| [util text nlp english entity](docs/ja/commands/util-text-nlp-english-entity.md)                                 | 英文をエンティティに分割する                                           |
| [util text nlp english sentence](docs/ja/commands/util-text-nlp-english-sentence.md)                             | 英文を文章に分割する                                                   |
| [util text nlp english token](docs/ja/commands/util-text-nlp-english-token.md)                                   | 英文をトークンに分割する                                               |
| [util text nlp japanese token](docs/ja/commands/util-text-nlp-japanese-token.md)                                 | 日本語テキストのトークン化                                             |
| [util text nlp japanese wakati](docs/ja/commands/util-text-nlp-japanese-wakati.md)                               | 分かち書き(日本語テキストのトークン化)                                 |
| [util tidy move dispatch](docs/ja/commands/util-tidy-move-dispatch.md)                                           | ファイルを整理                                                         |
| [util tidy move simple](docs/ja/commands/util-tidy-move-simple.md)                                               | ローカルファイルをアーカイブします                                     |
| [util time now](docs/ja/commands/util-time-now.md)                                                               | 現在の時刻を表示                                                       |
| [util unixtime format](docs/ja/commands/util-unixtime-format.md)                                                 | UNIX時間（1970-01-01からのエポック秒）を変換するための時間フォーマット |
| [util unixtime now](docs/ja/commands/util-unixtime-now.md)                                                       | UNIX時間で現在の時刻を表示する                                         |
| [util uuid v4](docs/ja/commands/util-uuid-v4.md)                                                                 | UUID v4（ランダムUUID）の生成                                          |
| [util video subtitles optimize](docs/ja/commands/util-video-subtitles-optimize.md)                               | 字幕ファイルの最適化                                                   |
| [util xlsx create](docs/ja/commands/util-xlsx-create.md)                                                         | 空のスプレッドシートを作成する                                         |
| [util xlsx sheet export](docs/ja/commands/util-xlsx-sheet-export.md)                                             | xlsxファイルからデータをエクスポート                                   |
| [util xlsx sheet import](docs/ja/commands/util-xlsx-sheet-import.md)                                             | データをxlsxファイルにインポート                                       |
| [util xlsx sheet list](docs/ja/commands/util-xlsx-sheet-list.md)                                                 | xlsxファイルのシート一覧                                               |
| [version](docs/ja/commands/version.md)                                                                           | バージョン情報                                                         |

