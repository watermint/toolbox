# Member management commands

## Information commands

Below commands are to retrieve information about team members.

| Command                                     | Description                    |
|---------------------------------------------|--------------------------------|
| [member list](member-list.md)               | List team member(s)            |
| [member folder list](member-folder-list.md) | Find folders for each member   |
| [member quota list](member-quota-list.md)   | List team member quota         |
| [member quota usage](member-quota-usage.md) | List team member storage usage |
| [team activity user](team-activity-user.md) | Activities log per user        |

## Basic management commands

Below commands are for managing team member accounts. Those commands are for a bulk operation by a CSV file.

| Command                                               | Description                                          |
|-------------------------------------------------------|------------------------------------------------------|
| [member invite](member-invite.md)                     | Invite member(s)                                     |
| [member delete](member-delete.md)                     | Delete members                                       |
| [member detach](member-detach.md)                     | Convert Dropbox Business accounts to a Basic account |
| [member reinvite](member-reinvite.md)                 | Reinvite invited status members to the team          |
| [member update email](member-update-email.md)         | Member email operation                               |
| [member update profile](member-update-profile.md)     | Update member profile                                |
| [member update visible](member-update-visible.md)     | Disable directory restriction to members             |
| [member update invisible](member-update-invisible.md) | Enable directory restriction to members              |
| [member quota update](member-quota-update.md)         | Update team member quota                             |

# Group commands

## Group management commands

Below commands are for managing groups.

| Command                                     | Description      |
|---------------------------------------------|------------------|
| [group add](group-add.md)                   | Create new group |
| [group delete](group-delete.md)             | Delete group     |
| [group batch delete](group-batch-delete.md) | Delete groups    |
| [group list](group-list.md)                 | List group(s)    |
| [group rename](group-rename.md)             | Rename the group |

## Find and delete unused groups

There are two commands to find unused groups. The first command is `group list`. The command `group list` will report
the number of members of each group. If it's zero, a group is not currently used to adding permission to folders. If you
want to see which folder uses each group, use the command `group folder list`. `group folder list` will report the group
to folder mapping. The report `group_with_no_folders` will show groups with no folders. You can safely remove groups
once if you check both the number of members and folders. After confirmation, you can bulk delete groups by using the
command `group batch delete`.

| Command                                     | Description                |
|---------------------------------------------|----------------------------|
| [group list](group-list.md)                 | List group(s)              |
| [group folder list](group-folder-list.md)   | Find folders of each group |
| [group batch delete](group-batch-delete.md) | Delete groups              |

# Notes:

Dropbox Business commands require admin permissions to execute them. Auth tokens must not share with anyone, including
Dropbox support.

