# watermint toolbox

[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![codecov](https://codecov.io/gh/watermint/toolbox/branch/master/graph/badge.svg)](https://codecov.io/gh/watermint/toolbox)

![watermint toolbox](resources/images/watermint-toolbox-256x256.png)

DropboxおよびDropbox Business向けのツールセット

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

© 2016-2020 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

DropboxおよびDropbox Business向けのツールセット

使い方:
=======

./tbx  コマンド

利用可能なコマンド:
===================

[1mコマンド      説明                     備考[0m
config        watermint toolbox の設定  
connect       アカウントに接続する             
file          ファイル操作                 
filerequest   ファイルリクエストの操作           
group         グループ管理                 
license       ライセンス情報を表示します          
member        チームメンバーの管理             
sharedfolder  共有フォルダ                 
sharedlink    個人アカウントの共有リンク          
team          Dropbox Business チーム   
teamfolder    チームフォルダの管理             
version       バージョン情報                
```

# コマンド

## Dropbox (個人アカウント)

| コマンド                                                                   | 説明                                                               |
|----------------------------------------------------------------------------|--------------------------------------------------------------------|
| [connect user_file](doc/generated_ja/connect-user_file.md)                 | ユーザーのファイルアクセスに接続する                               |
| [file compare account](doc/generated_ja/file-compare-account.md)           | 二つのアカウントのファイルを比較します                             |
| [file compare local](doc/generated_ja/file-compare-local.md)               | ローカルフォルダとDropboxフォルダの内容を比較します                |
| [file copy](doc/generated_ja/file-copy.md)                                 | ファイルをコピーします                                             |
| [file delete](doc/generated_ja/file-delete.md)                             | ファイルまたはフォルダは削除します.                                |
| [file download](doc/generated_ja/file-download.md)                         | Dropboxからファイルをダウンロードします                            |
| [file export doc](doc/generated_ja/file-export-doc.md)                     | ドキュメントのエクスポート                                         |
| [file import batch url](doc/generated_ja/file-import-batch-url.md)         | URLからファイルを一括インポートします                              |
| [file import url](doc/generated_ja/file-import-url.md)                     | URLからファイルをインポートします                                  |
| [file list](doc/generated_ja/file-list.md)                                 | ファイルとフォルダを一覧します                                     |
| [file merge](doc/generated_ja/file-merge.md)                               | フォルダを統合します                                               |
| [file move](doc/generated_ja/file-move.md)                                 | ファイルを移動します                                               |
| [file replication](doc/generated_ja/file-replication.md)                   | ファイルコンテンツを他のアカウントに複製します                     |
| [file restore](doc/generated_ja/file-restore.md)                           | 指定されたパス以下をリストアします                                 |
| [file search content](doc/generated_ja/file-search-content.md)             | ファイルコンテンツを検索                                           |
| [file search name](doc/generated_ja/file-search-name.md)                   | ファイル名を検索                                                   |
| [file sync preflight up](doc/generated_ja/file-sync-preflight-up.md)       | 上り方向同期のための事前チェックを実施します                       |
| [file sync up](doc/generated_ja/file-sync-up.md)                           | Dropboxと上り方向で同期します                                      |
| [file upload](doc/generated_ja/file-upload.md)                             | ファイルのアップロード                                             |
| [file watch](doc/generated_ja/file-watch.md)                               | ファイルアクティビティを監視                                       |
| [filerequest create](doc/generated_ja/filerequest-create.md)               | ファイルリクエストを作成します                                     |
| [filerequest delete closed](doc/generated_ja/filerequest-delete-closed.md) | このアカウントの全ての閉じられているファイルリクエストを削除します |
| [filerequest delete url](doc/generated_ja/filerequest-delete-url.md)       | ファイルリクエストのURLを指定して削除                              |
| [filerequest list](doc/generated_ja/filerequest-list.md)                   | 個人アカウントのファイルリクエストを一覧.                          |
| [job history ship](doc/generated_ja/job-history-ship.md)                   | ログの転送先Dropboxパス                                            |
| [sharedfolder list](doc/generated_ja/sharedfolder-list.md)                 | 共有フォルダの一覧                                                 |
| [sharedfolder member list](doc/generated_ja/sharedfolder-member-list.md)   | 共有フォルダのメンバーを一覧します                                 |
| [sharedlink create](doc/generated_ja/sharedlink-create.md)                 | 共有リンクの作成                                                   |
| [sharedlink delete](doc/generated_ja/sharedlink-delete.md)                 | 共有リンクを削除します                                             |
| [sharedlink file list](doc/generated_ja/sharedlink-file-list.md)           | 共有リンクのファイルを一覧する                                     |
| [sharedlink list](doc/generated_ja/sharedlink-list.md)                     | 共有リンクの一覧                                                   |

## Dropbox Business

| コマンド                                                                           | 説明                                                           |
|------------------------------------------------------------------------------------|----------------------------------------------------------------|
| [connect business_audit](doc/generated_ja/connect-business_audit.md)               | チーム監査アクセスに接続する                                   |
| [connect business_file](doc/generated_ja/connect-business_file.md)                 | チームファイルアクセスに接続する                               |
| [connect business_info](doc/generated_ja/connect-business_info.md)                 | チームの情報アクセスに接続する                                 |
| [connect business_mgmt](doc/generated_ja/connect-business_mgmt.md)                 | チームの管理アクセスに接続する                                 |
| [group add](doc/generated_ja/group-add.md)                                         | グループを作成します                                           |
| [group batch delete](doc/generated_ja/group-batch-delete.md)                       | グループの削除                                                 |
| [group delete](doc/generated_ja/group-delete.md)                                   | グループを削除します                                           |
| [group list](doc/generated_ja/group-list.md)                                       | グループを一覧                                                 |
| [group member add](doc/generated_ja/group-member-add.md)                           | メンバーをグループに追加                                       |
| [group member delete](doc/generated_ja/group-member-delete.md)                     | メンバーをグループから削除                                     |
| [group member list](doc/generated_ja/group-member-list.md)                         | グループに所属するメンバー一覧を取得します                     |
| [group rename](doc/generated_ja/group-rename.md)                                   | グループの改名                                                 |
| [member clear externalid](doc/generated_ja/member-clear-externalid.md)             | メンバーのexternal_idを初期化します                            |
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
| [team activity batch user](doc/generated_ja/team-activity-batch-user.md)           | 複数ユーザーのアクティビティを一括取得します                   |
| [team activity daily event](doc/generated_ja/team-activity-daily-event.md)         | アクティビティーを1日ごとに取得します                          |
| [team activity event](doc/generated_ja/team-activity-event.md)                     | イベントログ                                                   |
| [team activity user](doc/generated_ja/team-activity-user.md)                       | ユーザーごとのアクティビティ                                   |
| [team content member list](doc/generated_ja/team-content-member-list.md)           | チームフォルダや共有フォルダのメンバー一覧                     |
| [team content policy list](doc/generated_ja/team-content-policy-list.md)           | チームフォルダと共有フォルダのポリシー一覧                     |
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
| [teamfolder add](doc/generated_ja/teamfolder-add.md)                               | チームフォルダを追加します                                     |
| [teamfolder archive](doc/generated_ja/teamfolder-archive.md)                       | チームフォルダのアーカイブ                                     |
| [teamfolder batch archive](doc/generated_ja/teamfolder-batch-archive.md)           | 複数のチームフォルダをアーカイブします                         |
| [teamfolder batch permdelete](doc/generated_ja/teamfolder-batch-permdelete.md)     | 複数のチームフォルダを完全に削除します                         |
| [teamfolder batch replication](doc/generated_ja/teamfolder-batch-replication.md)   | チームフォルダの一括レプリケーション                           |
| [teamfolder file list](doc/generated_ja/teamfolder-file-list.md)                   | チームフォルダの一覧                                           |
| [teamfolder file size](doc/generated_ja/teamfolder-file-size.md)                   | チームフォルダのサイズを計算                                   |
| [teamfolder list](doc/generated_ja/teamfolder-list.md)                             | チームフォルダの一覧                                           |
| [teamfolder member list](doc/generated_ja/teamfolder-member-list.md)               | チームフォルダのメンバー一覧                                   |
| [teamfolder permdelete](doc/generated_ja/teamfolder-permdelete.md)                 | チームフォルダを完全に削除します                               |
| [teamfolder policy list](doc/generated_ja/teamfolder-policy-list.md)               | チームフォルダのポリシー一覧                                   |
| [teamfolder replication](doc/generated_ja/teamfolder-replication.md)               | チームフォルダを他のチームに複製します                         |

## GitHub

| コマンド                                                                                             | 説明                                             |
|------------------------------------------------------------------------------------------------------|--------------------------------------------------|
| [services github content get](doc/generated_ja/services-github-content-get.md)                       | レポジトリのコンテンツメタデータを取得します.    |
| [services github content put](doc/generated_ja/services-github-content-put.md)                       | レポジトリに小さなテキストコンテンツを格納します |
| [services github issue list](doc/generated_ja/services-github-issue-list.md)                         | 公開・プライベートGitHubレポジトリの課題一覧     |
| [services github profile](doc/generated_ja/services-github-profile.md)                               | 認証したユーザーの情報を取得                     |
| [services github release asset download](doc/generated_ja/services-github-release-asset-download.md) | アセットをダウンロードします                     |
| [services github release asset list](doc/generated_ja/services-github-release-asset-list.md)         | GitHubリリースの成果物一覧                       |
| [services github release asset upload](doc/generated_ja/services-github-release-asset-upload.md)     | GitHub リリースへ成果物をアップロードします      |
| [services github release draft](doc/generated_ja/services-github-release-draft.md)                   | リリースの下書きを作成                           |
| [services github release list](doc/generated_ja/services-github-release-list.md)                     | リリースの一覧                                   |
| [services github tag create](doc/generated_ja/services-github-tag-create.md)                         | レポジトリにタグを作成します                     |

## Google Gmail

| コマンド                                                                                                       | 説明                                        |
|----------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| [services google mail filter add](doc/generated_ja/services-google-mail-filter-add.md)                         | フィルターを追加します.                     |
| [services google mail filter batch add](doc/generated_ja/services-google-mail-filter-batch-add.md)             | クエリによるラベルの一括追加・削除          |
| [services google mail filter delete](doc/generated_ja/services-google-mail-filter-delete.md)                   | フィルタの削除                              |
| [services google mail filter list](doc/generated_ja/services-google-mail-filter-list.md)                       | フィルターの一覧                            |
| [services google mail label add](doc/generated_ja/services-google-mail-label-add.md)                           | ラベルの追加                                |
| [services google mail label delete](doc/generated_ja/services-google-mail-label-delete.md)                     | ラベルの削除.                               |
| [services google mail label list](doc/generated_ja/services-google-mail-label-list.md)                         | ラベルのリスト                              |
| [services google mail label rename](doc/generated_ja/services-google-mail-label-rename.md)                     | ラベルの名前を変更する                      |
| [services google mail message label add](doc/generated_ja/services-google-mail-message-label-add.md)           | メッセージにラベルを追加                    |
| [services google mail message label delete](doc/generated_ja/services-google-mail-message-label-delete.md)     | メッセージからラベルを削除する              |
| [services google mail message list](doc/generated_ja/services-google-mail-message-list.md)                     | メッセージの一覧                            |
| [services google mail message processed list](doc/generated_ja/services-google-mail-message-processed-list.md) | 処理された形式でメッセージを一覧表示します. |
| [services google mail thread list](doc/generated_ja/services-google-mail-thread-list.md)                       | スレッド一覧                                |

## Asana

| コマンド                                                                                           | 説明                             |
|----------------------------------------------------------------------------------------------------|----------------------------------|
| [services asana team list](doc/generated_ja/services-asana-team-list.md)                           | チームのリスト                   |
| [services asana team project list](doc/generated_ja/services-asana-team-project-list.md)           | チームのプロジェクト一覧         |
| [services asana team task list](doc/generated_ja/services-asana-team-task-list.md)                 | チームのタスク一覧               |
| [services asana workspace list](doc/generated_ja/services-asana-workspace-list.md)                 | ワークスペースの一覧             |
| [services asana workspace project list](doc/generated_ja/services-asana-workspace-project-list.md) | ワークスペースのプロジェクト一覧 |

## ユーティリティー

| コマンド                                                       | 説明                               |
|----------------------------------------------------------------|------------------------------------|
| [config disable](doc/generated_ja/config-disable.md)           | 機能を無効化します.                |
| [config enable](doc/generated_ja/config-enable.md)             | 機能を有効化します.                |
| [config features](doc/generated_ja/config-features.md)         | 利用可能なオプション機能一覧.      |
| [file dispatch local](doc/generated_ja/file-dispatch-local.md) | ローカルファイルを整理します       |
| [job history archive](doc/generated_ja/job-history-archive.md) | ジョブのアーカイブ                 |
| [job history delete](doc/generated_ja/job-history-delete.md)   | 古いジョブ履歴の削除               |
| [job history list](doc/generated_ja/job-history-list.md)       | ジョブ履歴の表示                   |
| [job log jobid](doc/generated_ja/job-log-jobid.md)             | 指定したジョブIDのログを取得する   |
| [job log kind](doc/generated_ja/job-log-kind.md)               | 指定種別のログを結合して出力します |
| [job log last](doc/generated_ja/job-log-last.md)               | 最後のジョブのログファイルを出力.  |
| [license](doc/generated_ja/license.md)                         | ライセンス情報を表示します         |
| [version](doc/generated_ja/version.md)                         | バージョン情報                     |

