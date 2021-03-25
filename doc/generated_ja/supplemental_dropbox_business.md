# メンバー管理コマンド

## 情報コマンド

以下のコマンドは、チームメンバーの情報を取得するためのものです.

| コマンド                                    | 説明                                     |
|---------------------------------------------|------------------------------------------|
| [member list](member-list.md)               | チームメンバーの一覧                     |
| [member folder list](member-folder-list.md) | 各メンバーのフォルダを検索               |
| [member quota list](member-quota-list.md)   | メンバーの容量制限情報を一覧します       |
| [member quota usage](member-quota-usage.md) | チームメンバーのストレージ利用状況を取得 |
| [team activity user](team-activity-user.md) | ユーザーごとのアクティビティ             |

## 基本管理コマンド

以下のコマンドは、チームメンバーのアカウントを管理するためのものです. これらのコマンドは、CSVファイルによる一括処理を行うためのものです.

| コマンド                                              | 説明                                                |
|-------------------------------------------------------|-----------------------------------------------------|
| [member invite](member-invite.md)                     | メンバーを招待します                                |
| [member delete](member-delete.md)                     | メンバーを削除します                                |
| [member detach](member-detach.md)                     | Dropbox BusinessユーザーをBasicユーザーに変更します |
| [member reinvite](member-reinvite.md)                 | 招待済み状態メンバーをチームに再招待します          |
| [member update email](member-update-email.md)         | メンバーのメールアドレス処理                        |
| [member update profile](member-update-profile.md)     | メンバーのプロフィール変更                          |
| [member update visible](member-update-visible.md)     | メンバーへのディレクトリ制限を無効にします          |
| [member update invisible](member-update-invisible.md) | メンバーへのディレクトリ制限を有効にします          |
| [member quota update](member-quota-update.md)         | チームメンバーの容量制限を変更                      |

## メンバープロファイル設定コマンド

メンバープロフィールコマンドは、メンバーのプロフィール情報を一括して更新するためのものです. メンバーのメールアドレスを更新する必要がある場合は、`member update email`コマンドを使用します.
コマンド`member update email`は、CSVファイルを受信してメールアドレスを一括更新します. メンバーの表示名を更新する必要がある場合は、`member update profile`コマンドを使用します.

| コマンド                                          | 説明                         |
|---------------------------------------------------|------------------------------|
| [member update email](member-update-email.md)     | メンバーのメールアドレス処理 |
| [member update profile](member-update-profile.md) | メンバーのプロフィール変更   |

## メンバーのストレージ クォータ制御コマンド

既存のメンバーストレージのクォータの設定や使用状況は、`member quota list`や`member quota usage`コマンドで確認できます.
メンバークオータを更新する必要がある場合は、`member quota update`コマンドを使用します. コマンド `member quota update` は、ストレージのクォータ設定を一括更新するためのCSV入力を受け付けます.

| コマンド                                      | 説明                                     |
|-----------------------------------------------|------------------------------------------|
| [member quota list](member-quota-list.md)     | メンバーの容量制限情報を一覧します       |
| [member quota usage](member-quota-usage.md)   | チームメンバーのストレージ利用状況を取得 |
| [member quota update](member-quota-update.md) | チームメンバーの容量制限を変更           |

## ディレクトリ制限コマンド

ディレクトリ制限は、Dropbox Businessの機能で、メンバーを他の人から隠すことができます. 以下のコマンドは、この設定を更新して、他の人からメンバーを隠したり、設定を解除したりします.

| コマンド                                              | 説明                                       |
|-------------------------------------------------------|--------------------------------------------|
| [member update visible](member-update-visible.md)     | メンバーへのディレクトリ制限を無効にします |
| [member update invisible](member-update-invisible.md) | メンバーへのディレクトリ制限を有効にします |

# グループのコマンド

## グループ管理コマンド

以下のコマンドはグループを管理するためのものです.

| コマンド                                    | 説明                 |
|---------------------------------------------|----------------------|
| [group add](group-add.md)                   | グループを作成します |
| [group delete](group-delete.md)             | グループを削除します |
| [group batch delete](group-batch-delete.md) | グループの削除       |
| [group list](group-list.md)                 | グループを一覧       |
| [group rename](group-rename.md)             | グループの改名       |

## グループメンバー管理コマンド

グループメンバーの追加・削除・更新は、以下のコマンドで行うことができます. グループメンバーをCSVファイルで追加/削除/更新したい場合は、`group member batch add`
, `group member batch delete`, `group member batch delete`を用います.

| コマンド                                                  | 説明                                       |
|-----------------------------------------------------------|--------------------------------------------|
| [group member add](group-member-add.md)                   | メンバーをグループに追加                   |
| [group member delete](group-member-delete.md)             | メンバーをグループから削除                 |
| [group member list](group-member-list.md)                 | グループに所属するメンバー一覧を取得します |
| [group member batch add](group-member-batch-add.md)       | グループにメンバーを一括追加               |
| [group member batch delete](group-member-batch-delete.md) | グループからメンバーを削除                 |
| [group member batch update](group-member-batch-update.md) | グループからメンバーを追加または削除       |

## 未使用のグループの検索と削除

未使用のグループを探すには2つのコマンドがあります. 最初のコマンドは `group list` です. コマンド `group list` は、各グループのメンバー数を報告します.
0の場合は、フォルダに権限を追加するためのグループが現在使用されていません. どのフォルダが各グループを使用しているかを確認したい場合は、`group folder list`
というコマンドを使います. `group folder list`では、グループとフォルダのマッピングを報告します. `group_with_no_folders`というレポートでは、フォルダがないグループが表示されます.
グループの削除は、メンバー数とフォルダ数の両方を確認すれば、安全に行うことができます. 確認後、`group batch delete`コマンドでグループを一括削除することができます.

| コマンド                                    | 説明                       |
|---------------------------------------------|----------------------------|
| [group list](group-list.md)                 | グループを一覧             |
| [group folder list](group-folder-list.md)   | 各グループのフォルダを探す |
| [group batch delete](group-batch-delete.md) | グループの削除             |

# チームコンテンツのコマンド

管理者はDropbox Business APIを使って、チームフォルダ、共有フォルダ、メンバーのフォルダのコンテンツを扱うことができます. これらのコマンドの使用には注意が必要です. 名前空間とは、Dropbox
APIの中で、フォルダの権限や設定を管理するための用語です. 共有フォルダ、チームフォルダ、チームフォルダ内のネストしたフォルダ、メンバーのルートフォルダ、メンバーのアプリフォルダなどのフォルダタイプは、すべて名前空間として管理されます.
名前空間コマンドでは、チームフォルダやメンバーズフォルダなど、あらゆる種類のフォルダを扱うことができます. しかし、特定のフォルダタイプのコマンドは、より多くの機能や詳細な情報がレポートに含まれています.

## チームフォルダ操作コマンド

以下のコマンドを使って、チームフォルダーの作成、アーカイブ、完全に削除ができます. 複数のチームフォルダを扱う必要がある場合は、`teamfolder batch`コマンドの使用をご検討ください.

| コマンド                                                        | 説明                                   |
|-----------------------------------------------------------------|----------------------------------------|
| [teamfolder list](teamfolder-list.md)                           | チームフォルダの一覧                   |
| [teamfolder policy list](teamfolder-policy-list.md)             | チームフォルダのポリシー一覧           |
| [teamfolder file size](teamfolder-file-size.md)                 | チームフォルダのサイズを計算           |
| [teamfolder add](teamfolder-add.md)                             | チームフォルダを追加します             |
| [teamfolder archive](teamfolder-archive.md)                     | チームフォルダのアーカイブ             |
| [teamfolder permdelete](teamfolder-permdelete.md)               | チームフォルダを完全に削除します       |
| [teamfolder batch archive](teamfolder-batch-archive.md)         | 複数のチームフォルダをアーカイブします |
| [teamfolder batch permdelete](teamfolder-batch-permdelete.md)   | 複数のチームフォルダを完全に削除します |
| [teamfolder batch replication](teamfolder-batch-replication.md) | チームフォルダの一括レプリケーション   |

## チームフォルダの権限コマンド

チームフォルダーやサブフォルダーにメンバーを一括で追加・削除するには、以下のコマンドを使います.

| コマンド                                                | 説明                                            |
|---------------------------------------------------------|-------------------------------------------------|
| [teamfolder member list](teamfolder-member-list.md)     | チームフォルダのメンバー一覧                    |
| [teamfolder member add](teamfolder-member-add.md)       | チームフォルダへのユーザー/グループの一括追加   |
| [teamfolder member delete](teamfolder-member-delete.md) | チームフォルダからのユーザー/グループの一括削除 |

## チームフォルダと共有フォルダのコマンド

以下のコマンドは、チームフォルダとチームの共有フォルダの両方に対応しています. 特定のフォルダを実際に使用している人を知りたい場合は、`team content mount list`というコマンドの使用をご検討ください.
マウントは、ユーザーが自分のDropboxアカウントに共有フォルダを追加した状態です.

| コマンド                                                | 説明                                                                                   |
|---------------------------------------------------------|----------------------------------------------------------------------------------------|
| [team content member list](team-content-member-list.md) | チームフォルダや共有フォルダのメンバー一覧                                             |
| [team content mount list](team-content-mount-list.md)   | チームメンバーのマウント済み/アンマウント済みの共有フォルダをすべてリストアップします. |
| [team content policy list](team-content-policy-list.md) | チームフォルダと共有フォルダのポリシー一覧                                             |

## 名前空間コマンド

| コマンド                                                    | 説明                                               |
|-------------------------------------------------------------|----------------------------------------------------|
| [team namespace list](team-namespace-list.md)               | チーム内すべての名前空間を一覧                     |
| [team namespace file list](team-namespace-file-list.md)     | チーム内全ての名前空間でのファイル・フォルダを一覧 |
| [team namespace file size](team-namespace-file-size.md)     | チーム内全ての名前空間でのファイル・フォルダを一覧 |
| [team namespace member list](team-namespace-member-list.md) | チームフォルダ以下のファイル・フォルダを一覧       |

## チームのファイルリクエスト コマンド

| コマンド                                          | 説明                                       |
|---------------------------------------------------|--------------------------------------------|
| [team filerequest list](team-filerequest-list.md) | チームないのファイルリクエストを一覧します |

## Member file commands

| コマンド                                            | 説明                                                                                                                                        |
|-----------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| [member file permdelete](member-file-permdelete.md) | チームメンバーの指定したパスのファイルまたはフォルダを完全に削除します完全に削除については、https://www.dropbox.com/help/40 をご覧ください. |

# チーム共有リンクコマンド

チーム共有リンクコマンドは、チーム内のすべての共有リンクを一覧表示したり、指定した共有リンクを更新・削除することができます.

| コマンド                                                                  | 説明                                                           |
|---------------------------------------------------------------------------|----------------------------------------------------------------|
| [team sharedlink list](team-sharedlink-list.md)                           | 共有リンクの一覧                                               |
| [team sharedlink update expiry](team-sharedlink-update-expiry.md)         | チーム内の公開されている共有リンクについて有効期限を更新します |
| [team sharedlink update password](team-sharedlink-update-password.md)     | 共有リンクのパスワードの設定・更新                             |
| [team sharedlink update visibility](team-sharedlink-update-visibility.md) | 共有リンクの可視性の更新                                       |
| [team sharedlink delete links](team-sharedlink-delete-links.md)           | 共有リンクの一括削除                                           |
| [team sharedlink delete member](team-sharedlink-delete-member.md)         | メンバーの共有リンクをすべて削除                               |

## 例(リンクの一覧):

チーム内のすべてのパブリックリンクをリストアップ:

```
tbx team sharedlink list -visibility public
```

結果は、CSV、xlsx、JSON形式で保存されます. 共有リンクを更新するためのレポートを変更することができます. jqというコマンドに慣れていれば、以下のように直接CSVファイルを作成することができます.

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

| コマンド                                                        | 説明                                               |
|-----------------------------------------------------------------|----------------------------------------------------|
| [member file lock all release](member-file-lock-all-release.md) | メンバーのパスの下にあるすべてのロックを解除します |
| [member file lock list](member-file-lock-list.md)               | パスの下にあるメンバーのロックを一覧表示           |
| [member file lock release](member-file-lock-release.md)         | メンバーとしてパスのロックを解除します             |

## チームフォルダのファイルロックコマンド

| コマンド                                                                | 説明                                                   |
|-------------------------------------------------------------------------|--------------------------------------------------------|
| [teamfolder file list](teamfolder-file-list.md)                         | チームフォルダの一覧                                   |
| [teamfolder file lock all release](teamfolder-file-lock-all-release.md) | チームフォルダのパスの下にあるすべてのロックを解除する |
| [teamfolder file lock list](teamfolder-file-lock-list.md)               | チームフォルダ内のロックを一覧表示                     |
| [teamfolder file lock release](teamfolder-file-lock-release.md)         | チームフォルダ内のパスのロックを解除                   |

# アクティビティ ログ コマンド

チーム活動ログのコマンドでは、特定のフィルターや視点で活動ログをエクスポートすることができます.

| コマンド                                                  | 説明                                         |
|-----------------------------------------------------------|----------------------------------------------|
| [team activity batch user](team-activity-batch-user.md)   | 複数ユーザーのアクティビティを一括取得します |
| [team activity daily event](team-activity-daily-event.md) | アクティビティーを1日ごとに取得します        |
| [team activity event](team-activity-event.md)             | イベントログ                                 |
| [team activity user](team-activity-user.md)               | ユーザーごとのアクティビティ                 |

# 接続されたアプリケーションとデバイスのコマンド.

以下のコマンドは、チーム内で接続されているデバイスやアプリケーションの情報を取得することができます.

| コマンド                                      | 説明                                          |
|-----------------------------------------------|-----------------------------------------------|
| [team device list](team-device-list.md)       | チーム内全てのデバイス/セッションを一覧します |
| [team device unlink](team-device-unlink.md)   | デバイスのセッションを解除します              |
| [team linkedapp list](team-linkedapp-list.md) | リンク済みアプリを一覧                        |

# その他の使用例

## External ID

External IDは、Dropboxのどのユーザーインターフェースにも表示されない属性です. この属性は、Dropbox AD ConnectorなどのID管理ソフトウェアによって、DropboxとIDソース（Active
Directoryや人事データベースなど）との関係を維持するためのものです. Dropbox AD Connectorを使用していて、新しいActive Directoryツリーを構築した場合は、以下のようになります. 古いActive
Directoryツリーと新しいツリーとの関係を切断するために、既存の外部IDをクリアする必要があるかもしれません. External IDのクリアを省略すると、Dropbox AD
Connectorが新しいツリーへの構成中に意図せずアカウントを削除してしまう可能性があります. 既存の外部IDを確認したい場合は、`member list`コマンドを使います.
しかし、このコマンドはデフォルトでは外部IDを含みません. [jq](https://stedolan.github.io/jq/)コマンドを使用して、以下のように実行することをご検討ください.

```
tbx member list -output json | jq -r '[.profile.email, .profile.external_id] | @csv'
```

| コマンド                                                | 説明                                     |
|---------------------------------------------------------|------------------------------------------|
| [member list](member-list.md)                           | チームメンバーの一覧                     |
| [member clear externalid](member-clear-externalid.md)   | メンバーのexternal_idを初期化します      |
| [member update externalid](member-update-externalid.md) | チームメンバーのExternal IDを更新します. |

## データ移行補助コマンド

データ移行補助コマンドは、メンバーフォルダやチームフォルダを別のアカウントやチームにコピーします. 実際にデータを移行する前に、これらのコマンドを使用してテストしてください.

| コマンド                                                            | 説明                                                 |
|---------------------------------------------------------------------|------------------------------------------------------|
| [member folder replication](member-folder-replication.md)           | フォルダを他のメンバーの個人フォルダに複製します     |
| [member replication](member-replication.md)                         | チームメンバーのファイルを複製します                 |
| [teamfolder partial replication](teamfolder-partial-replication.md) | 部分的なチームフォルダの他チームへのレプリケーション |
| [teamfolder replication](teamfolder-replication.md)                 | チームフォルダを他のチームに複製します               |

## チーム情報コマンド

| コマンド                        | 説明                     |
|---------------------------------|--------------------------|
| [team feature](team-feature.md) | チームの機能を出力します |
| [team info](team-info.md)       | チームの情報             |

# 注意事項:

Dropbox Businessのコマンドを実行するには、管理者権限が必要です. 認証トークンは、Dropboxのサポートを含め、誰とも共有してはいけません.

