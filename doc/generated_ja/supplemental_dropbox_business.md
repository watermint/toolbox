# メンバー管理コマンド

## 情報コマンド

以下のコマンドは、チームメンバーの情報を取得するためのものです.

| Command                                     | Description                              |
|---------------------------------------------|------------------------------------------|
| [member list](member-list.md)               | チームメンバーの一覧                     |
| [member folder list](member-folder-list.md) | 各メンバーのフォルダを検索               |
| [member quota list](member-quota-list.md)   | メンバーの容量制限情報を一覧します       |
| [member quota usage](member-quota-usage.md) | チームメンバーのストレージ利用状況を取得 |
| [team activity user](team-activity-user.md) | ユーザーごとのアクティビティ             |

## 基本管理コマンド

以下のコマンドは、チームメンバーのアカウントを管理するためのものです. これらのコマンドは、CSVファイルによる一括処理を行うためのものです.

| Command                                               | Description                                         |
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

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                     | Description          |
|---------------------------------------------|----------------------|
| [group add](group-add.md)                   | グループを作成します |
| [group delete](group-delete.md)             | グループを削除します |
| [group batch delete](group-batch-delete.md) | グループの削除       |
| [group list](group-list.md)                 | グループを一覧       |
| [group rename](group-rename.md)             | グループの改名       |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `group list`. The command `group list` will report
the number of members of each group. If it's zero, a group is not currently used to adding permission to folders. If you
want to see which folder uses each group, use the command `group folder list`. `group folder list` will report the group
to folder mapping. The report `group_with_no_folders` will show groups with no folders. You can safely remove groups
once if you check both the number of members and folders. After confirmation, you can bulk delete groups by using the
command `group batch delete`.

| Command                                     | Description                |
|---------------------------------------------|----------------------------|
| [group list](group-list.md)                 | グループを一覧             |
| [group folder list](group-folder-list.md)   | 各グループのフォルダを探す |
| [group batch delete](group-batch-delete.md) | グループの削除             |

# Notes:

Dropbox Business commands require admin permissions to execute them. Auth tokens must not share with anyone, including
Dropbox support.

