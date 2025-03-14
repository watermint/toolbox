---
layout: page
title: チーム向けDropboxのコマンド
lang: ja
---

# メンバー管理コマンド

## 情報コマンド

以下のコマンドは、チームメンバーの情報を取得するためのものです.

| コマンド                                                                                               | 説明                                     |
|--------------------------------------------------------------------------------------------------------|------------------------------------------|
| [dropbox team member list]({{ site.baseurl }}/ja/commands/dropbox-team-member-list.html)               | チームメンバーの一覧                     |
| [dropbox team member feature]({{ site.baseurl }}/ja/commands/dropbox-team-member-feature.html)         | メンバーの機能設定一覧                   |
| [dropbox team member folder list]({{ site.baseurl }}/ja/commands/dropbox-team-member-folder-list.html) | 各メンバーのフォルダーを一覧表示         |
| [dropbox team member quota list]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-list.html)   | メンバーの容量制限情報を一覧します       |
| [dropbox team member quota usage]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-usage.html) | チームメンバーのストレージ利用状況を取得 |
| [dropbox team activity user]({{ site.baseurl }}/ja/commands/dropbox-team-activity-user.html)           | ユーザーごとのアクティビティ             |

## 基本管理コマンド

以下のコマンドは、チームメンバーのアカウントを管理するためのものです. これらのコマンドは、CSVファイルによる一括処理を行うためのものです.

| コマンド                                                                                                                     | 説明                                                     |
|------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| [dropbox team member batch invite]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-invite.html)                     | メンバーを招待します                                     |
| [dropbox team member batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-delete.html)                     | メンバーを削除します                                     |
| [dropbox team member batch detach]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-detach.html)                     | Dropbox for teamsのアカウントをBasicアカウントに変更する |
| [dropbox team member batch reinvite]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-reinvite.html)                 | 招待済み状態メンバーをチームに再招待します               |
| [dropbox team member update batch email]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-email.html)         | メンバーのメールアドレス処理                             |
| [dropbox team member update batch profile]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-profile.html)     | メンバーのプロフィール変更                               |
| [dropbox team member update batch visible]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-visible.html)     | メンバーへのディレクトリ制限を無効にします               |
| [dropbox team member update batch invisible]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-invisible.html) | メンバーへのディレクトリ制限を有効にします               |
| [dropbox team member quota batch update]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-batch-update.html)         | チームメンバーの容量制限を変更                           |

## メンバープロファイル設定コマンド

メンバープロフィールコマンドは、メンバーのプロフィール情報を一括して更新するためのものです.
メンバーのメールアドレスを更新する必要がある場合は、`member update email`コマンドを使用します. コマンド`member update email`は、CSVファイルを受信してメールアドレスを一括更新します.
メンバーの表示名を更新する必要がある場合は、`member update profile`コマンドを使用します.

| コマンド                                                                                                                 | 説明                         |
|--------------------------------------------------------------------------------------------------------------------------|------------------------------|
| [dropbox team member update batch email]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-email.html)     | メンバーのメールアドレス処理 |
| [dropbox team member update batch profile]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-profile.html) | メンバーのプロフィール変更   |

## メンバーのストレージ クォータ制御コマンド

`dropbox team member quota list` と `dropbox team member quota usage` コマンドを使えば、既存のメンバーのストレージクォータの設定や使用状況を見ることができる。メンバーのクォータを更新する必要がある場合は、`dropbox team member quota update`コマンドを使用します。コマンド `dropbox team member quota update` はストレージのクォータ設定を一括更新するためにCSV入力を受け取ります。

| コマンド                                                                                                             | 説明                                     |
|----------------------------------------------------------------------------------------------------------------------|------------------------------------------|
| [dropbox team member quota list]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-list.html)                 | メンバーの容量制限情報を一覧します       |
| [dropbox team member quota usage]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-usage.html)               | チームメンバーのストレージ利用状況を取得 |
| [dropbox team member quota batch update]({{ site.baseurl }}/ja/commands/dropbox-team-member-quota-batch-update.html) | チームメンバーの容量制限を変更           |

## メンバーの一時停止/停止解除には、2種類のコマンドがあります. メンバーを一人ずつサスペンド/アンサスペンドしたい場合は、`dropbox team member suspend` または `dropbox team member unsuspend` を使ってください。その他、CSVファイルを通してメンバーのサスペンド/アンサスペンドを行いたい場合は、`dropbox team member batch suspend`または`dropbox member batch unsuspend`コマンドを使用してください。

メンバーの一時停止/停止解除

| コマンド                                                                                                       | 説明                         |
|----------------------------------------------------------------------------------------------------------------|------------------------------|
| [dropbox team member suspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-suspend.html)                 | メンバーの一時停止処理       |
| [dropbox team member unsuspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-unsuspend.html)             | メンバーの一時停止を解除する |
| [dropbox team member batch suspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-suspend.html)     | メンバーの一括一時停止       |
| [dropbox team member batch unsuspend]({{ site.baseurl }}/ja/commands/dropbox-team-member-batch-unsuspend.html) | メンバーの一括停止解除       |

## ディレクトリ制限コマンド

ディレクトリ制限は、Dropbox for teamsの機能で、メンバーを他のメンバーから隠すことができます。以下のコマンドは、この設定を更新して、他の人からメンバーを隠したり、設定を解除したりします.

| コマンド                                                                                                                     | 説明                                       |
|------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------|
| [dropbox team member update batch visible]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-visible.html)     | メンバーへのディレクトリ制限を無効にします |
| [dropbox team member update batch invisible]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-invisible.html) | メンバーへのディレクトリ制限を有効にします |

# グループのコマンド

## グループ管理コマンド

以下のコマンドはグループを管理するためのものです.

| コマンド                                                                                               | 説明                     |
|--------------------------------------------------------------------------------------------------------|--------------------------|
| [dropbox team group add]({{ site.baseurl }}/ja/commands/dropbox-team-group-add.html)                   | グループを作成します     |
| [dropbox team group batch add]({{ site.baseurl }}/ja/commands/dropbox-team-group-batch-add.html)       | グループの一括追加       |
| [dropbox team group batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-batch-delete.html) | グループの削除           |
| [dropbox team group delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-delete.html)             | グループを削除します     |
| [dropbox team group list]({{ site.baseurl }}/ja/commands/dropbox-team-group-list.html)                 | グループを一覧           |
| [dropbox team group rename]({{ site.baseurl }}/ja/commands/dropbox-team-group-rename.html)             | グループの改名           |
| [dropbox team group update type]({{ site.baseurl }}/ja/commands/dropbox-team-group-update-type.html)   | グループ管理タイプの更新 |

## グループメンバー管理コマンド

グループメンバーの追加・削除・更新は、以下のコマンドで行うことができます. CSVファイルでグループメンバーを追加・削除・更新したい場合は、`dropbox team group member batch add`、`dropbox team group member batch delete`、`dropbox team group member batch delete`を使用します。

| コマンド                                                                                                             | 説明                                       |
|----------------------------------------------------------------------------------------------------------------------|--------------------------------------------|
| [dropbox team group member add]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-add.html)                   | メンバーをグループに追加                   |
| [dropbox team group member delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-delete.html)             | メンバーをグループから削除                 |
| [dropbox team group member list]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-list.html)                 | グループに所属するメンバー一覧を取得します |
| [dropbox team group member batch add]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-batch-add.html)       | グループにメンバーを一括追加               |
| [dropbox team group member batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-batch-delete.html) | グループからメンバーを削除                 |
| [dropbox team group member batch update]({{ site.baseurl }}/ja/commands/dropbox-team-group-member-batch-update.html) | グループからメンバーを追加または削除       |

## 未使用のグループの検索と削除

未使用のグループを探すには2つのコマンドがあります. 最初のコマンドは`dropbox team group list`である。`dropbox team group list`コマンドは各グループのメンバー数を表示します。0の場合は、フォルダに権限を追加するためのグループが現在使用されていません.
どのフォルダが各グループを使用しているかを確認したい場合は、`dropbox team group folder list`コマンドを使用します。`dropbox team group folder list`はグループとフォルダのマッピングを報告します。`group_with_no_folders`というレポートでは、フォルダがないグループが表示されます.
グループの削除は、メンバー数とフォルダ数の両方を確認すれば、安全に行うことができます. 確認後、`dropbox team group batch delete`コマンドを使ってグループを一括削除することができます。

| コマンド                                                                                               | 説明                             |
|--------------------------------------------------------------------------------------------------------|----------------------------------|
| [dropbox team group list]({{ site.baseurl }}/ja/commands/dropbox-team-group-list.html)                 | グループを一覧                   |
| [dropbox team group folder list]({{ site.baseurl }}/ja/commands/dropbox-team-group-folder-list.html)   | 各グループのフォルダーを一覧表示 |
| [dropbox team group batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-group-batch-delete.html) | グループの削除                   |

# チームコンテンツのコマンド

管理者はDropbox Business APIを通じて、チームフォルダ、共有フォルダ、メンバーフォルダのコンテンツを扱うことができます。これらのコマンドの使用には注意が必要です.
名前空間とは、Dropbox APIの中で、フォルダの権限や設定を管理するための用語です. 共有フォルダ、チームフォルダ、チームフォルダ内のネストしたフォルダ、メンバーのルートフォルダ、メンバーのアプリフォルダなどのフォルダタイプは、すべて名前空間として管理されます.
名前空間コマンドでは、チームフォルダやメンバーズフォルダなど、あらゆる種類のフォルダを扱うことができます. しかし、特定のフォルダタイプのコマンドは、より多くの機能や詳細な情報がレポートに含まれています.

## チームフォルダ操作コマンド

以下のコマンドを使って、チームフォルダーの作成、アーカイブ、完全に削除ができます. 複数のチームフォルダを処理する必要がある場合は、`dropbox team teamfolder batch`コマンドの使用を検討してください。

| コマンド                                                                                                                       | 説明                                   |
|--------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| [dropbox team teamfolder add]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-add.html)                                 | チームフォルダを追加します             |
| [dropbox team teamfolder archive]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-archive.html)                         | チームフォルダのアーカイブ             |
| [dropbox team teamfolder batch archive]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-batch-archive.html)             | 複数のチームフォルダをアーカイブします |
| [dropbox team teamfolder batch permdelete]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-batch-permdelete.html)       | 複数のチームフォルダを完全に削除します |
| [dropbox team teamfolder batch replication]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-batch-replication.html)     | チームフォルダの一括レプリケーション   |
| [dropbox team teamfolder file size]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-size.html)                     | チームフォルダのサイズを計算           |
| [dropbox team teamfolder list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-list.html)                               | チームフォルダの一覧                   |
| [dropbox team teamfolder permdelete]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-permdelete.html)                   | チームフォルダを完全に削除します       |
| [dropbox team teamfolder policy list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-policy-list.html)                 | チームフォルダのポリシー一覧           |
| [dropbox team teamfolder sync setting list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-sync-setting-list.html)     | チームフォルダーの同期設定を一覧表示   |
| [dropbox team teamfolder sync setting update]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-sync-setting-update.html) | チームフォルダ同期設定の一括更新       |

## チームフォルダの権限コマンド

チームフォルダーやサブフォルダーにメンバーを一括で追加・削除するには、以下のコマンドを使います.

| コマンド                                                                                                           | 説明                                            |
|--------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| [dropbox team teamfolder member list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-member-list.html)     | チームフォルダのメンバー一覧                    |
| [dropbox team teamfolder member add]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-member-add.html)       | チームフォルダへのユーザー/グループの一括追加   |
| [dropbox team teamfolder member delete]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-member-delete.html) | チームフォルダからのユーザー/グループの一括削除 |

## チームフォルダと共有フォルダのコマンド

以下のコマンドは、チームフォルダとチームの共有フォルダの両方に対応しています.
特定のフォルダーを実際に使用している人を知りたい場合は、`dropbox team content mount list`コマンドの使用を検討してください。マウントは、ユーザーが自分のDropboxアカウントに共有フォルダを追加した状態です.

| コマンド                                                                                                 | 説明                                                                                   |
|----------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|
| [dropbox team content member list]({{ site.baseurl }}/ja/commands/dropbox-team-content-member-list.html) | チームフォルダや共有フォルダのメンバー一覧                                             |
| [dropbox team content member size]({{ site.baseurl }}/ja/commands/dropbox-team-content-member-size.html) | チームフォルダや共有フォルダのメンバー数をカウントする                                 |
| [dropbox team content mount list]({{ site.baseurl }}/ja/commands/dropbox-team-content-mount-list.html)   | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします. |
| [dropbox team content policy list]({{ site.baseurl }}/ja/commands/dropbox-team-content-policy-list.html) | チームフォルダと共有フォルダのポリシー一覧                                             |

## 名前空間コマンド

| コマンド                                                                                                     | 説明                                               |
|--------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| [dropbox team namespace list]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-list.html)               | チーム内すべての名前空間を一覧                     |
| [dropbox team namespace summary]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-summary.html)         | チーム・ネームスペースの状態概要を報告する.        |
| [dropbox team namespace file list]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-file-list.html)     | チーム内全ての名前空間でのファイル・フォルダを一覧 |
| [dropbox team namespace file size]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-file-size.html)     | チーム内全ての名前空間でのファイル・フォルダを一覧 |
| [dropbox team namespace member list]({{ site.baseurl }}/ja/commands/dropbox-team-namespace-member-list.html) | チームフォルダ以下のファイル・フォルダを一覧       |

## チームのファイルリクエスト コマンド

| コマンド                                                                                           | 説明                                       |
|----------------------------------------------------------------------------------------------------|--------------------------------------------|
| [dropbox team filerequest list]({{ site.baseurl }}/ja/commands/dropbox-team-filerequest-list.html) | チームないのファイルリクエストを一覧します |

## メンバーファイルのコマンド

| コマンド                                                                                                       | 説明                                                                   |
|----------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| [dropbox team member file permdelete]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-permdelete.html) | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します |

## チーム インサイト

チーム インサイトは Dropbox Business の機能で、チーム フォルダのサマリーを表示します.

| コマンド                                                                                                                         | 説明                                                     |
|----------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------|
| [dropbox team insight scan]({{ site.baseurl }}/ja/commands/dropbox-team-insight-scan.html)                                       | チームデータをスキャンして分析                           |
| [dropbox team insight scanretry]({{ site.baseurl }}/ja/commands/dropbox-team-insight-scanretry.html)                             | 前回のスキャンでエラーがあった場合、スキャンを再試行する |
| [dropbox team insight summarize]({{ site.baseurl }}/ja/commands/dropbox-team-insight-summarize.html)                             | 分析のためにチームデータをまとめる                       |
| [dropbox team insight report teamfoldermember]({{ site.baseurl }}/ja/commands/dropbox-team-insight-report-teamfoldermember.html) | チームフォルダーメンバーを報告                           |

# チーム共有リンクコマンド

チーム共有リンクコマンドは、チーム内のすべての共有リンクを一覧表示したり、指定した共有リンクを更新・削除することができます.

| コマンド                                                                                                                   | 説明                                                           |
|----------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------|
| [dropbox team sharedlink list]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-list.html)                           | 共有リンクの一覧                                               |
| [dropbox team sharedlink cap expiry]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-cap-expiry.html)               | チーム内の共有リンクに有効期限の上限を設定                     |
| [dropbox team sharedlink cap visibility]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-cap-visibility.html)       | チーム内の共有リンクに可視性の上限を設定                       |
| [dropbox team sharedlink update expiry]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-update-expiry.html)         | チーム内の公開されている共有リンクについて有効期限を更新します |
| [dropbox team sharedlink update password]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-update-password.html)     | 共有リンクのパスワードの設定・更新                             |
| [dropbox team sharedlink update visibility]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-update-visibility.html) | 共有リンクの可視性の更新                                       |
| [dropbox team sharedlink delete links]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-delete-links.html)           | 共有リンクの一括削除                                           |
| [dropbox team sharedlink delete member]({{ site.baseurl }}/ja/commands/dropbox-team-sharedlink-delete-member.html)         | メンバーの共有リンクをすべて削除                               |

## `dropbox team sharelink cap`と`dropbox team sharelink update`の違い

コマンド `dropbox team sharedlink update` は共有リンクに値を設定するためのコマンドです。コマンド `dropbox team sharedlink cap` は共有リンクに上限値を設定するためのコマンドです。
例えば、`dropbox team sharedlink update expiry`で有効期限を2021-05-06に設定した場合。このコマンドは、既存のリンクが2021-05-04のように短い有効期限を持っている場合でも、有効期限を2021-05-06に更新します.
一方、`dropbox team sharedlink cap expiry`は、リンクの有効期限が2021-05-07のように長い場合、リンクを更新する。

同様に、`dropbox team sharedlink cap visibility`コマンドは、リンクが保護されている可視性が低い場合にのみ可視性を制限します。例えば、パスワードのない共有リンクをteam_onlyにしたい場合などです. dropbox team sharedlink cap visibility` は、リンクが公開されていてパスワードがない場合にのみ、チームの可視性を更新します。

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

| コマンド                                                                                                                   | 説明                                               |
|----------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| [dropbox team member file lock all release]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-lock-all-release.html) | メンバーのパスの下にあるすべてのロックを解除します |
| [dropbox team member file lock list]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-lock-list.html)               | パスの下にあるメンバーのロックを一覧表示           |
| [dropbox team member file lock release]({{ site.baseurl }}/ja/commands/dropbox-team-member-file-lock-release.html)         | メンバーとしてパスのロックを解除します             |

## チームフォルダのファイルロックコマンド

| コマンド                                                                                                                           | 説明                                                   |
|------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------|
| [dropbox team teamfolder file list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-list.html)                         | チームフォルダの一覧                                   |
| [dropbox team teamfolder file lock all release]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-lock-all-release.html) | チームフォルダのパスの下にあるすべてのロックを解除する |
| [dropbox team teamfolder file lock list]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-lock-list.html)               | チームフォルダ内のロックを一覧表示                     |
| [dropbox team teamfolder file lock release]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-file-lock-release.html)         | チームフォルダ内のパスのロックを解除                   |

# アクティビティ ログ コマンド

チーム活動ログのコマンドでは、特定のフィルターや視点で活動ログをエクスポートすることができます.

| コマンド                                                                                                   | 説明                                         |
|------------------------------------------------------------------------------------------------------------|----------------------------------------------|
| [dropbox team activity batch user]({{ site.baseurl }}/ja/commands/dropbox-team-activity-batch-user.html)   | 複数ユーザーのアクティビティを一括取得します |
| [dropbox team activity daily event]({{ site.baseurl }}/ja/commands/dropbox-team-activity-daily-event.html) | アクティビティーを1日ごとに取得します        |
| [dropbox team activity event]({{ site.baseurl }}/ja/commands/dropbox-team-activity-event.html)             | イベントログ                                 |
| [dropbox team activity user]({{ site.baseurl }}/ja/commands/dropbox-team-activity-user.html)               | ユーザーごとのアクティビティ                 |

# 接続されたアプリケーションとデバイスのコマンド.

以下のコマンドは、チーム内で接続されているデバイスやアプリケーションの情報を取得することができます.

| コマンド                                                                                                   | 説明                                                                  |
|------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|
| [dropbox team device list]({{ site.baseurl }}/ja/commands/dropbox-team-device-list.html)                   | チーム内全てのデバイス/セッションを一覧します                         |
| [dropbox team device unlink]({{ site.baseurl }}/ja/commands/dropbox-team-device-unlink.html)               | デバイスのセッションを解除します                                      |
| [dropbox team linkedapp list]({{ site.baseurl }}/ja/commands/dropbox-team-linkedapp-list.html)             | リンク済みアプリを一覧                                                |
| [dropbox team backup device status]({{ site.baseurl }}/ja/commands/dropbox-team-backup-device-status.html) | Dropbox バックアップ デバイスのステータスが指定期間内に変更された場合 |

# その他の使用例

## External ID

External IDは、Dropboxのどのユーザーインターフェースにも表示されない属性です. この属性は、Dropbox AD ConnectorなどのID管理ソフトウェアによって、DropboxとIDソース（Active Directoryや人事データベースなど）との関係を維持するためのものです. Dropbox AD Connectorを使用していて、新しいActive Directoryツリーを構築した場合は、以下のようになります. 古いActive Directoryツリーと新しいツリーとの関係を切断するために、既存の外部IDをクリアする必要があるかもしれません.
External IDのクリアを省略すると、Dropbox AD Connectorが新しいツリーへの構成中に意図せずアカウントを削除してしまう可能性があります.
既存の外部IDを見たい場合は、`dropbox team member list`コマンドを使ってください。しかし、このコマンドはデフォルトでは外部IDを含みません. 以下のように`experiment report_all_columns`オプションを追加してください

```
tbx member list -experiment report_all_columns
```

| コマンド                                                                                                                       | 説明                                     |
|--------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|
| [dropbox team member list]({{ site.baseurl }}/ja/commands/dropbox-team-member-list.html)                                       | チームメンバーの一覧                     |
| [dropbox team member clear externalid]({{ site.baseurl }}/ja/commands/dropbox-team-member-clear-externalid.html)               | メンバーのexternal_idを初期化します      |
| [dropbox team member update batch externalid]({{ site.baseurl }}/ja/commands/dropbox-team-member-update-batch-externalid.html) | チームメンバーのExternal IDを更新します. |
| [dropbox team group list]({{ site.baseurl }}/ja/commands/dropbox-team-group-list.html)                                         | グループを一覧                           |
| [dropbox team group clear externalid]({{ site.baseurl }}/ja/commands/dropbox-team-group-clear-externalid.html)                 | グループの外部IDをクリアする             |

## データ移行補助コマンド

データ移行補助コマンドは、メンバーフォルダやチームフォルダを別のアカウントやチームにコピーします. 実際にデータを移行する前に、これらのコマンドを使用してテストしてください.

| コマンド                                                                                                                       | 説明                                                 |
|--------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------|
| [dropbox team member folder replication]({{ site.baseurl }}/ja/commands/dropbox-team-member-folder-replication.html)           | フォルダを他のメンバーの個人フォルダに複製します     |
| [dropbox team member replication]({{ site.baseurl }}/ja/commands/dropbox-team-member-replication.html)                         | チームメンバーのファイルを複製します                 |
| [dropbox team teamfolder partial replication]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-partial-replication.html) | 部分的なチームフォルダの他チームへのレプリケーション |
| [dropbox team teamfolder replication]({{ site.baseurl }}/ja/commands/dropbox-team-teamfolder-replication.html)                 | チームフォルダを他のチームに複製します               |

## チーム情報コマンド

| コマンド                                                                               | 説明                                           |
|----------------------------------------------------------------------------------------|------------------------------------------------|
| [dropbox team feature]({{ site.baseurl }}/ja/commands/dropbox-team-feature.html)       | チームの機能を出力します                       |
| [dropbox team filesystem]({{ site.baseurl }}/ja/commands/dropbox-team-filesystem.html) | チームのファイルシステムのバージョンを特定する |
| [dropbox team info]({{ site.baseurl }}/ja/commands/dropbox-team-info.html)             | チームの情報                                   |

# Paperコマンド

## レガシーPaperコマンド

チームのレガシーPaperコンテンツのコマンドです. レガシーペーパーとマイグレーションの詳細については、[公式ガイド](https://developers.dropbox.com/paper-migration-guide)をご覧ください.

| コマンド                                                                                                               | 説明                                                       |
|------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|
| [dropbox team content legacypaper count]({{ site.baseurl }}/ja/commands/dropbox-team-content-legacypaper-count.html)   | メンバー1人あたりのPaper文書の枚数                         |
| [dropbox team content legacypaper list]({{ site.baseurl }}/ja/commands/dropbox-team-content-legacypaper-list.html)     | チームメンバーのPaper文書リスト出力                        |
| [dropbox team content legacypaper export]({{ site.baseurl }}/ja/commands/dropbox-team-content-legacypaper-export.html) | チームメンバー全員のPaper文書をローカルパスにエクスポート. |

# チーム管理者用コマンド

以下のコマンドは、チーム管理者を管理するためのものです.

| コマンド                                                                                                         | 説明                                                             |
|------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------|
| [dropbox team admin list]({{ site.baseurl }}/ja/commands/dropbox-team-admin-list.html)                           | メンバーの管理者権限一覧                                         |
| [dropbox team admin role add]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-add.html)                   | メンバーに新しいロールを追加する                                 |
| [dropbox team admin role clear]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-clear.html)               | メンバーからすべての管理者ロールを削除する                       |
| [dropbox team admin role delete]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-delete.html)             | メンバーからロールを削除する                                     |
| [dropbox team admin role list]({{ site.baseurl }}/ja/commands/dropbox-team-admin-role-list.html)                 | チームの管理者の役割を列挙                                       |
| [dropbox team admin group role add]({{ site.baseurl }}/ja/commands/dropbox-team-admin-group-role-add.html)       | グループのメンバーにロールを追加する                             |
| [dropbox team admin group role delete]({{ site.baseurl }}/ja/commands/dropbox-team-admin-group-role-delete.html) | 例外グループのメンバーを除くすべてのメンバーからロールを削除する |

# チームメンバーとして実行するコマンド

チームメンバーとしてコマンドを実行することができます. 例えば、`dropbox team runas file sync batch up`を使ってメンバーのフォルダにファイルをアップロードすることができます。

| コマンド                                                                                                                                       | 説明                                                                                    |
|------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| [dropbox team runas file list]({{ site.baseurl }}/ja/commands/dropbox-team-runas-file-list.html)                                               | メンバーとして実行するファイルやフォルダーの一覧                                        |
| [dropbox team runas file batch copy]({{ site.baseurl }}/ja/commands/dropbox-team-runas-file-batch-copy.html)                                   | ファイル/フォルダーをメンバーとして一括コピー                                           |
| [dropbox team runas file sync batch up]({{ site.baseurl }}/ja/commands/dropbox-team-runas-file-sync-batch-up.html)                             | メンバーとして動作する一括同期                                                          |
| [dropbox team runas sharedfolder list]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-list.html)                               | 共有フォルダーの一覧をメンバーとして実行                                                |
| [dropbox team runas sharedfolder isolate]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-isolate.html)                         | 所有する共有フォルダの共有を解除し、メンバーとして実行する外部共有フォルダから離脱する. |
| [dropbox team runas sharedfolder mount add]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-add.html)                     | 指定したメンバーのDropboxに共有フォルダを追加する                                       |
| [dropbox team runas sharedfolder mount delete]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-delete.html)               | 指定されたユーザーが指定されたフォルダーをアンマウントする.                             |
| [dropbox team runas sharedfolder mount list]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-list.html)                   | 指定されたメンバーがマウントしているすべての共有フォルダーをリストアップします.         |
| [dropbox team runas sharedfolder mount mountable]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-mount-mountable.html)         | メンバーがマウントできるすべての共有フォルダーをリストアップ.                           |
| [dropbox team runas sharedfolder batch leave]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-batch-leave.html)                 | 共有フォルダからメンバーとして一括退出                                                  |
| [dropbox team runas sharedfolder batch share]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-batch-share.html)                 | メンバーのフォルダを一括で共有                                                          |
| [dropbox team runas sharedfolder batch unshare]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-batch-unshare.html)             | メンバーのフォルダの共有を一括解除                                                      |
| [dropbox team runas sharedfolder member batch add]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-member-batch-add.html)       | メンバーの共有フォルダにメンバーを一括追加                                              |
| [dropbox team runas sharedfolder member batch delete]({{ site.baseurl }}/ja/commands/dropbox-team-runas-sharedfolder-member-batch-delete.html) | メンバーの共有フォルダからメンバーを一括削除                                            |

# リーガルホールド

リーガルホールドでは、管理者はチームのメンバーをリーガルホールドし、そのメンバーが作成または変更したすべてのコンテンツを表示およびエクスポートすることができます.

| コマンド                                                                                                                     | 説明                                                       |
|------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|
| [dropbox team legalhold add]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-add.html)                                 | 新しいリーガル・ホールド・ポリシーを作成する.              |
| [dropbox team legalhold list]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-list.html)                               | 既存のポリシーを取得する                                   |
| [dropbox team legalhold member batch update]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-member-batch-update.html) | リーガル・ホールド・ポリシーのメンバーリスト更新           |
| [dropbox team legalhold member list]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-member-list.html)                 | リーガルホールドのメンバーをリストアップ                   |
| [dropbox team legalhold release]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-release.html)                         | Idによるリーガルホールドを解除する                         |
| [dropbox team legalhold revision list]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-revision-list.html)             | リーガル・ホールド・ポリシーのリビジョンをリストアップする |
| [dropbox team legalhold update desc]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-update-desc.html)                 | リーガルホールド・ポリシーの説明を更新                     |
| [dropbox team legalhold update name]({{ site.baseurl }}/ja/commands/dropbox-team-legalhold-update-name.html)                 | リーガルホールドポリシーの名称を更新                       |

# 注意事項:

Dropbox for teamsのコマンドを実行するには、管理者権限が必要です。認証トークンは、Dropboxのサポートを含め、誰とも共有してはいけません.


