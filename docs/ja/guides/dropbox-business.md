---
layout: page
title: Dropbox Business コマンド
lang: ja
---

# メンバー管理コマンド

## 情報コマンド

以下のコマンドは、チームメンバーの情報を取得するためのものです.

| コマンド                                                                     | 説明                                     |
|------------------------------------------------------------------------------|------------------------------------------|
| [member list]({{ site.baseurl }}/ja/commands/member-list.html)               | チームメンバーの一覧                     |
| [member folder list]({{ site.baseurl }}/ja/commands/member-folder-list.html) | 各メンバーのフォルダを検索               |
| [member quota list]({{ site.baseurl }}/ja/commands/member-quota-list.html)   | メンバーの容量制限情報を一覧します       |
| [member quota usage]({{ site.baseurl }}/ja/commands/member-quota-usage.html) | チームメンバーのストレージ利用状況を取得 |
| [team activity user]({{ site.baseurl }}/ja/commands/team-activity-user.html) | ユーザーごとのアクティビティ             |

## 基本管理コマンド

以下のコマンドは、チームメンバーのアカウントを管理するためのものです. これらのコマンドは、CSVファイルによる一括処理を行うためのものです.

| コマンド                                                                               | 説明                                                |
|----------------------------------------------------------------------------------------|-----------------------------------------------------|
| [member invite]({{ site.baseurl }}/ja/commands/member-invite.html)                     | メンバーを招待します                                |
| [member delete]({{ site.baseurl }}/ja/commands/member-delete.html)                     | メンバーを削除します                                |
| [member detach]({{ site.baseurl }}/ja/commands/member-detach.html)                     | Dropbox BusinessユーザーをBasicユーザーに変更します |
| [member reinvite]({{ site.baseurl }}/ja/commands/member-reinvite.html)                 | 招待済み状態メンバーをチームに再招待します          |
| [member update email]({{ site.baseurl }}/ja/commands/member-update-email.html)         | メンバーのメールアドレス処理                        |
| [member update profile]({{ site.baseurl }}/ja/commands/member-update-profile.html)     | メンバーのプロフィール変更                          |
| [member update visible]({{ site.baseurl }}/ja/commands/member-update-visible.html)     | メンバーへのディレクトリ制限を無効にします          |
| [member update invisible]({{ site.baseurl }}/ja/commands/member-update-invisible.html) | メンバーへのディレクトリ制限を有効にします          |
| [member quota update]({{ site.baseurl }}/ja/commands/member-quota-update.html)         | チームメンバーの容量制限を変更                      |

## メンバープロファイル設定コマンド

メンバープロフィールコマンドは、メンバーのプロフィール情報を一括して更新するためのものです.
メンバーのメールアドレスを更新する必要がある場合は、`member update email`コマンドを使用します. コマンド`member update email`は、CSVファイルを受信してメールアドレスを一括更新します.
メンバーの表示名を更新する必要がある場合は、`member update profile`コマンドを使用します.

| コマンド                                                                           | 説明                         |
|------------------------------------------------------------------------------------|------------------------------|
| [member update email]({{ site.baseurl }}/ja/commands/member-update-email.html)     | メンバーのメールアドレス処理 |
| [member update profile]({{ site.baseurl }}/ja/commands/member-update-profile.html) | メンバーのプロフィール変更   |

## メンバーのストレージ クォータ制御コマンド

既存のメンバーストレージのクォータの設定や使用状況は、`member quota list`や`member quota usage`コマンドで確認できます. メンバークオータを更新する必要がある場合は、`member quota update`コマンドを使用します. コマンド `member quota update` は、ストレージのクォータ設定を一括更新するためのCSV入力を受け付けます.

| コマンド                                                                       | 説明                                     |
|--------------------------------------------------------------------------------|------------------------------------------|
| [member quota list]({{ site.baseurl }}/ja/commands/member-quota-list.html)     | メンバーの容量制限情報を一覧します       |
| [member quota usage]({{ site.baseurl }}/ja/commands/member-quota-usage.html)   | チームメンバーのストレージ利用状況を取得 |
| [member quota update]({{ site.baseurl }}/ja/commands/member-quota-update.html) | チームメンバーの容量制限を変更           |

## ディレクトリ制限コマンド

ディレクトリ制限は、Dropbox Businessの機能で、メンバーを他の人から隠すことができます. 以下のコマンドは、この設定を更新して、他の人からメンバーを隠したり、設定を解除したりします.

| コマンド                                                                               | 説明                                       |
|----------------------------------------------------------------------------------------|--------------------------------------------|
| [member update visible]({{ site.baseurl }}/ja/commands/member-update-visible.html)     | メンバーへのディレクトリ制限を無効にします |
| [member update invisible]({{ site.baseurl }}/ja/commands/member-update-invisible.html) | メンバーへのディレクトリ制限を有効にします |

# グループのコマンド

## グループ管理コマンド

以下のコマンドはグループを管理するためのものです.

| コマンド                                                                     | 説明                 |
|------------------------------------------------------------------------------|----------------------|
| [group add]({{ site.baseurl }}/ja/commands/group-add.html)                   | グループを作成します |
| [group delete]({{ site.baseurl }}/ja/commands/group-delete.html)             | グループを削除します |
| [group batch delete]({{ site.baseurl }}/ja/commands/group-batch-delete.html) | グループの削除       |
| [group list]({{ site.baseurl }}/ja/commands/group-list.html)                 | グループを一覧       |
| [group rename]({{ site.baseurl }}/ja/commands/group-rename.html)             | グループの改名       |

## グループメンバー管理コマンド

グループメンバーの追加・削除・更新は、以下のコマンドで行うことができます. グループメンバーをCSVファイルで追加/削除/更新したい場合は、`group member batch add`, `group member batch delete`, `group member batch delete`を用います.

| コマンド                                                                                   | 説明                                       |
|--------------------------------------------------------------------------------------------|--------------------------------------------|
| [group member add]({{ site.baseurl }}/ja/commands/group-member-add.html)                   | メンバーをグループに追加                   |
| [group member delete]({{ site.baseurl }}/ja/commands/group-member-delete.html)             | メンバーをグループから削除                 |
| [group member list]({{ site.baseurl }}/ja/commands/group-member-list.html)                 | グループに所属するメンバー一覧を取得します |
| [group member batch add]({{ site.baseurl }}/ja/commands/group-member-batch-add.html)       | グループにメンバーを一括追加               |
| [group member batch delete]({{ site.baseurl }}/ja/commands/group-member-batch-delete.html) | グループからメンバーを削除                 |
| [group member batch update]({{ site.baseurl }}/ja/commands/group-member-batch-update.html) | グループからメンバーを追加または削除       |

## 未使用のグループの検索と削除

未使用のグループを探すには2つのコマンドがあります. 最初のコマンドは `group list` です. コマンド `group list` は、各グループのメンバー数を報告します. 0の場合は、フォルダに権限を追加するためのグループが現在使用されていません.
どのフォルダが各グループを使用しているかを確認したい場合は、`group folder list`というコマンドを使います. `group folder list`では、グループとフォルダのマッピングを報告します. `group_with_no_folders`というレポートでは、フォルダがないグループが表示されます.
グループの削除は、メンバー数とフォルダ数の両方を確認すれば、安全に行うことができます. 確認後、`group batch delete`コマンドでグループを一括削除することができます.

| コマンド                                                                     | 説明                       |
|------------------------------------------------------------------------------|----------------------------|
| [group list]({{ site.baseurl }}/ja/commands/group-list.html)                 | グループを一覧             |
| [group folder list]({{ site.baseurl }}/ja/commands/group-folder-list.html)   | 各グループのフォルダを探す |
| [group batch delete]({{ site.baseurl }}/ja/commands/group-batch-delete.html) | グループの削除             |

# チームコンテンツのコマンド

管理者はDropbox Business APIを使って、チームフォルダ、共有フォルダ、メンバーのフォルダのコンテンツを扱うことができます. これらのコマンドの使用には注意が必要です.
名前空間とは、Dropbox APIの中で、フォルダの権限や設定を管理するための用語です. 共有フォルダ、チームフォルダ、チームフォルダ内のネストしたフォルダ、メンバーのルートフォルダ、メンバーのアプリフォルダなどのフォルダタイプは、すべて名前空間として管理されます.
名前空間コマンドでは、チームフォルダやメンバーズフォルダなど、あらゆる種類のフォルダを扱うことができます. しかし、特定のフォルダタイプのコマンドは、より多くの機能や詳細な情報がレポートに含まれています.

## チームフォルダ操作コマンド

以下のコマンドを使って、チームフォルダーの作成、アーカイブ、完全に削除ができます. 複数のチームフォルダを扱う必要がある場合は、`teamfolder batch`コマンドの使用をご検討ください.

| コマンド                                                                                         | 説明                                   |
|--------------------------------------------------------------------------------------------------|----------------------------------------|
| [teamfolder list]({{ site.baseurl }}/ja/commands/teamfolder-list.html)                           | チームフォルダの一覧                   |
| [teamfolder policy list]({{ site.baseurl }}/ja/commands/teamfolder-policy-list.html)             | チームフォルダのポリシー一覧           |
| [teamfolder file size]({{ site.baseurl }}/ja/commands/teamfolder-file-size.html)                 | チームフォルダのサイズを計算           |
| [teamfolder add]({{ site.baseurl }}/ja/commands/teamfolder-add.html)                             | チームフォルダを追加します             |
| [teamfolder archive]({{ site.baseurl }}/ja/commands/teamfolder-archive.html)                     | チームフォルダのアーカイブ             |
| [teamfolder permdelete]({{ site.baseurl }}/ja/commands/teamfolder-permdelete.html)               | チームフォルダを完全に削除します       |
| [teamfolder batch archive]({{ site.baseurl }}/ja/commands/teamfolder-batch-archive.html)         | 複数のチームフォルダをアーカイブします |
| [teamfolder batch permdelete]({{ site.baseurl }}/ja/commands/teamfolder-batch-permdelete.html)   | 複数のチームフォルダを完全に削除します |
| [teamfolder batch replication]({{ site.baseurl }}/ja/commands/teamfolder-batch-replication.html) | チームフォルダの一括レプリケーション   |

## チームフォルダの権限コマンド

チームフォルダーやサブフォルダーにメンバーを一括で追加・削除するには、以下のコマンドを使います.

| コマンド                                                                                 | 説明                                            |
|------------------------------------------------------------------------------------------|-------------------------------------------------|
| [teamfolder member list]({{ site.baseurl }}/ja/commands/teamfolder-member-list.html)     | チームフォルダのメンバー一覧                    |
| [teamfolder member add]({{ site.baseurl }}/ja/commands/teamfolder-member-add.html)       | チームフォルダへのユーザー/グループの一括追加   |
| [teamfolder member delete]({{ site.baseurl }}/ja/commands/teamfolder-member-delete.html) | チームフォルダからのユーザー/グループの一括削除 |

## チームフォルダと共有フォルダのコマンド

以下のコマンドは、チームフォルダとチームの共有フォルダの両方に対応しています.
特定のフォルダを実際に使用している人を知りたい場合は、`team content mount list`というコマンドの使用をご検討ください. マウントは、ユーザーが自分のDropboxアカウントに共有フォルダを追加した状態です.

| コマンド                                                                                 | 説明                                                                                   |
|------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|
| [team content member list]({{ site.baseurl }}/ja/commands/team-content-member-list.html) | チームフォルダや共有フォルダのメンバー一覧                                             |
| [team content member size]({{ site.baseurl }}/ja/commands/team-content-member-size.html) | チームフォルダや共有フォルダのメンバー数をカウントする                                 |
| [team content mount list]({{ site.baseurl }}/ja/commands/team-content-mount-list.html)   | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします. |
| [team content policy list]({{ site.baseurl }}/ja/commands/team-content-policy-list.html) | チームフォルダと共有フォルダのポリシー一覧                                             |

## 名前空間コマンド

| コマンド                                                                                     | 説明                                               |
|----------------------------------------------------------------------------------------------|----------------------------------------------------|
| [team namespace list]({{ site.baseurl }}/ja/commands/team-namespace-list.html)               | チーム内すべての名前空間を一覧                     |
| [team namespace file list]({{ site.baseurl }}/ja/commands/team-namespace-file-list.html)     | チーム内全ての名前空間でのファイル・フォルダを一覧 |
| [team namespace file size]({{ site.baseurl }}/ja/commands/team-namespace-file-size.html)     | チーム内全ての名前空間でのファイル・フォルダを一覧 |
| [team namespace member list]({{ site.baseurl }}/ja/commands/team-namespace-member-list.html) | チームフォルダ以下のファイル・フォルダを一覧       |

## チームのファイルリクエスト コマンド

| コマンド                                                                           | 説明                                       |
|------------------------------------------------------------------------------------|--------------------------------------------|
| [team filerequest list]({{ site.baseurl }}/ja/commands/team-filerequest-list.html) | チームないのファイルリクエストを一覧します |

## メンバーファイルのコマンド

| コマンド                                                                             | 説明                                                                   |
|--------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [member file permdelete]({{ site.baseurl }}/ja/commands/member-file-permdelete.html) | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します |

# チーム共有リンクコマンド

チーム共有リンクコマンドは、チーム内のすべての共有リンクを一覧表示したり、指定した共有リンクを更新・削除することができます.

| コマンド                                                                                                   | 説明                                                           |
|------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------|
| [team sharedlink list]({{ site.baseurl }}/ja/commands/team-sharedlink-list.html)                           | 共有リンクの一覧                                               |
| [team sharedlink cap expiry]({{ site.baseurl }}/ja/commands/team-sharedlink-cap-expiry.html)               | チーム内の共有リンクに有効期限の上限を設定                     |
| [team sharedlink cap visibility]({{ site.baseurl }}/ja/commands/team-sharedlink-cap-visibility.html)       | チーム内の共有リンクに可視性の上限を設定                       |
| [team sharedlink update expiry]({{ site.baseurl }}/ja/commands/team-sharedlink-update-expiry.html)         | チーム内の公開されている共有リンクについて有効期限を更新します |
| [team sharedlink update password]({{ site.baseurl }}/ja/commands/team-sharedlink-update-password.html)     | 共有リンクのパスワードの設定・更新                             |
| [team sharedlink update visibility]({{ site.baseurl }}/ja/commands/team-sharedlink-update-visibility.html) | 共有リンクの可視性の更新                                       |
| [team sharedlink delete links]({{ site.baseurl }}/ja/commands/team-sharedlink-delete-links.html)           | 共有リンクの一括削除                                           |
| [team sharedlink delete member]({{ site.baseurl }}/ja/commands/team-sharedlink-delete-member.html)         | メンバーの共有リンクをすべて削除                               |

## `team sharedlink cap` と `team sharedlink update` の違い

コマンド `team sharedlink update` は、共有リンクに値を設定するためのものです. コマンド `team sharedlink cap` は、共有リンクにキャップ値を設定するためのものです.
例：有効期限を2021-05-06に設定して、`team sharedlink update expiry`で設定した場合. このコマンドは、既存のリンクが2021-05-04のように短い有効期限を持っている場合でも、有効期限を2021-05-06に更新します.
一方、`team sharedlink cap expiry`は、リンクの有効期限が2021-05-07のように長い場合にリンクを更新します.

同様に、`team sharedlink cap visibility`というコマンドは、リンクの保護された可視性が少ない場合にのみ、可視性を制限します. 例えば、パスワードのない共有リンクをteam_onlyにしたい場合などです. team sharelink cap visibility` は、リンクが公開されていてパスワードがない場合にteam_onlyへ可視性を更新します

## 例(リンクの一覧):

チーム内のすべてのパブリックリンクをリストアップ:

```
tbx team sharedlink list -visibility public
```

結果は、CSV、xlsx、JSON形式で保存されます. 共有リンクを更新するためのレポートを変更することができます.
jqというコマンドに慣れていれば、以下のように直接CSVファイルを作成することができます.

```
tbx team sharedlink list -output json | jq '.sharedlink.url' > all_links.csv
```

リンク所有者のメールアドレスでフィルタリングされたリンクを一覧表示します:

```
tbx team sharedlink list -output json | jq 'select(.member.profile.email == "username@example.com") | .sharedlink.url'
```

## 例（リンクの削除）:

CSVファイルに記載されているすべてのリンクを削除する

```
tbx team sharedlink delete links -file /PATH/TO/DATA.csv
```

jqコマンドに慣れていれば、以下のようにパイプから直接データを送ることができます(標準入力から読み込む場合は、`-file`オプションにシングルダッシュ `-`を渡します).

```
tbx team sharedlink list -visibility public -output json | tbx team sharedlink delete links -file -
```

# ファイルロック

ファイルロックコマンドは、現在のファイルロックを一覧表示したり、管理者としてファイルロックを解除することができます.

## メンバーのファイルロックコマンド

| コマンド                                                                                         | 説明                                               |
|--------------------------------------------------------------------------------------------------|----------------------------------------------------|
| [member file lock all release]({{ site.baseurl }}/ja/commands/member-file-lock-all-release.html) | メンバーのパスの下にあるすべてのロックを解除します |
| [member file lock list]({{ site.baseurl }}/ja/commands/member-file-lock-list.html)               | パスの下にあるメンバーのロックを一覧表示           |
| [member file lock release]({{ site.baseurl }}/ja/commands/member-file-lock-release.html)         | メンバーとしてパスのロックを解除します             |

## チームフォルダのファイルロックコマンド

| コマンド                                                                                                 | 説明                                                   |
|----------------------------------------------------------------------------------------------------------|--------------------------------------------------------|
| [teamfolder file list]({{ site.baseurl }}/ja/commands/teamfolder-file-list.html)                         | チームフォルダの一覧                                   |
| [teamfolder file lock all release]({{ site.baseurl }}/ja/commands/teamfolder-file-lock-all-release.html) | チームフォルダのパスの下にあるすべてのロックを解除する |
| [teamfolder file lock list]({{ site.baseurl }}/ja/commands/teamfolder-file-lock-list.html)               | チームフォルダ内のロックを一覧表示                     |
| [teamfolder file lock release]({{ site.baseurl }}/ja/commands/teamfolder-file-lock-release.html)         | チームフォルダ内のパスのロックを解除                   |

# アクティビティ ログ コマンド

チーム活動ログのコマンドでは、特定のフィルターや視点で活動ログをエクスポートすることができます.

| コマンド                                                                                   | 説明                                         |
|--------------------------------------------------------------------------------------------|----------------------------------------------|
| [team activity batch user]({{ site.baseurl }}/ja/commands/team-activity-batch-user.html)   | 複数ユーザーのアクティビティを一括取得します |
| [team activity daily event]({{ site.baseurl }}/ja/commands/team-activity-daily-event.html) | アクティビティーを1日ごとに取得します        |
| [team activity event]({{ site.baseurl }}/ja/commands/team-activity-event.html)             | イベントログ                                 |
| [team activity user]({{ site.baseurl }}/ja/commands/team-activity-user.html)               | ユーザーごとのアクティビティ                 |

# 接続されたアプリケーションとデバイスのコマンド.

以下のコマンドは、チーム内で接続されているデバイスやアプリケーションの情報を取得することができます.

| コマンド                                                                       | 説明                                          |
|--------------------------------------------------------------------------------|-----------------------------------------------|
| [team device list]({{ site.baseurl }}/ja/commands/team-device-list.html)       | チーム内全てのデバイス/セッションを一覧します |
| [team device unlink]({{ site.baseurl }}/ja/commands/team-device-unlink.html)   | デバイスのセッションを解除します              |
| [team linkedapp list]({{ site.baseurl }}/ja/commands/team-linkedapp-list.html) | リンク済みアプリを一覧                        |

# その他の使用例

## External ID

External IDは、Dropboxのどのユーザーインターフェースにも表示されない属性です. この属性は、Dropbox AD ConnectorなどのID管理ソフトウェアによって、DropboxとIDソース（Active Directoryや人事データベースなど）との関係を維持するためのものです. Dropbox AD Connectorを使用していて、新しいActive Directoryツリーを構築した場合は、以下のようになります. 古いActive Directoryツリーと新しいツリーとの関係を切断するために、既存の外部IDをクリアする必要があるかもしれません.
External IDのクリアを省略すると、Dropbox AD Connectorが新しいツリーへの構成中に意図せずアカウントを削除してしまう可能性があります.
既存の外部IDを確認したい場合は、`member list`コマンドを使います. しかし、このコマンドはデフォルトでは外部IDを含みません. [jq](https://stedolan.github.io/jq/)コマンドを使用して、以下のように実行することをご検討ください.

```
tbx member list -output json | jq -r '[.profile.email, .profile.external_id] | @csv'
```

| コマンド                                                                                 | 説明                                     |
|------------------------------------------------------------------------------------------|------------------------------------------|
| [member list]({{ site.baseurl }}/ja/commands/member-list.html)                           | チームメンバーの一覧                     |
| [member clear externalid]({{ site.baseurl }}/ja/commands/member-clear-externalid.html)   | メンバーのexternal_idを初期化します      |
| [member update externalid]({{ site.baseurl }}/ja/commands/member-update-externalid.html) | チームメンバーのExternal IDを更新します. |

## データ移行補助コマンド

データ移行補助コマンドは、メンバーフォルダやチームフォルダを別のアカウントやチームにコピーします. 実際にデータを移行する前に、これらのコマンドを使用してテストしてください.

| コマンド                                                                                             | 説明                                                 |
|------------------------------------------------------------------------------------------------------|------------------------------------------------------|
| [member folder replication]({{ site.baseurl }}/ja/commands/member-folder-replication.html)           | フォルダを他のメンバーの個人フォルダに複製します     |
| [member replication]({{ site.baseurl }}/ja/commands/member-replication.html)                         | チームメンバーのファイルを複製します                 |
| [teamfolder partial replication]({{ site.baseurl }}/ja/commands/teamfolder-partial-replication.html) | 部分的なチームフォルダの他チームへのレプリケーション |
| [teamfolder replication]({{ site.baseurl }}/ja/commands/teamfolder-replication.html)                 | チームフォルダを他のチームに複製します               |

## チーム情報コマンド

| コマンド                                                         | 説明                     |
|------------------------------------------------------------------|--------------------------|
| [team feature]({{ site.baseurl }}/ja/commands/team-feature.html) | チームの機能を出力します |
| [team info]({{ site.baseurl }}/ja/commands/team-info.html)       | チームの情報             |

# 注意事項:

Dropbox Businessのコマンドを実行するには、管理者権限が必要です. 認証トークンは、Dropboxのサポートを含め、誰とも共有してはいけません.


