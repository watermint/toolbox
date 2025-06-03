# Dropbox Commands Documentation Analysis

## Overview
This document provides a comprehensive analysis of all Dropbox-related commands in the watermint toolbox and their current documentation status.

## Total Commands Found: 210

### Command Categories

#### 1. File Operations (citron.dropbox.file.*)
- **Basic Operations**: copy, delete, info, list, merge, move, size, watch
- **Account**: feature, filesystem, info
- **Compare**: account, local
- **Export**: doc, url
- **Import**: url, batch/url
- **Lock**: acquire, release, list, all/release, batch/acquire, batch/release
- **Request**: create, list, delete/url, delete/closed
- **Restore**: all, ext, restore
- **Revision**: download, list, restore
- **Search**: content, name
- **Share**: info
- **Sharedfolder**: info, leave, list, share, unshare, member/add, member/delete, member/list, mount/add, mount/delete, mount/list, mount/mountable
- **Sharedlink**: create, delete, info, list, file/list
- **Sync**: down, online, up
- **Tag**: add, delete, list
- **Template**: apply, capture

#### 2. Paper Operations (citron.dropbox.paper.*)
- append, create, overwrite, prepend

#### 3. Sign Operations (citron.dropbox.sign.*)
- **Account**: info
- **Request**: list, signature/list

#### 4. Team Operations (citron.dropbox.team.*)
- **Activity**: batch/user, daily/event, event, user
- **Admin**: group/role/add, group/role/delete, list, role/add, role/clear, role/delete, role/list
- **Backup**: device/status
- **Content**: legacypaper/count, legacypaper/export, legacypaper/list, member/list, member/size, mount/list, policy/list
- **Device**: list, unlink
- **Feature**: (basic team feature info)
- **Filerequest**: clone, list
- **Filesystem**: (team filesystem info)
- **Group**: add, batch/add, batch/delete, clear/externalid, delete, folder/list, list, member/add, member/batch/add, member/batch/delete, member/batch/operation, member/batch/record, member/batch/update, member/delete, member/info, member/list, rename, update/type
- **Info**: (basic team info)
- **Insight**: report/teamfoldermember, scan, scanretry, summarize
- **Legalhold**: add, list, member/batch/update, member/list, release, revision/list, update/desc, update/name
- **Linkedapp**: list
- **Member**: batch/delete, batch/detach, batch/invite, batch/reinvite, batch/suspend, batch/unsuspend, batch/user, clear/externalid, feature, file/lock/all/release, file/lock/list, file/lock/release, file/permdelete, folder/list, folder/replication, list, quota/batch/update, quota/list, quota/usage, replication, suspend, unsuspend, update/batch/email, update/batch/externalid, update/batch/invisible, update/batch/profile, update/batch/visibility_record, update/batch/visible
- **Namespace**: file/list, file/size, list, member/list, summary
- **Report**: activity, devices, membership, storage
- **Runas**: file/batch/copy, file/batch/copymapping, file/list, file/sync/batch/up, file/sync/batch/upmapping, sharedfolder/batch/leave, sharedfolder/batch/member_folder, sharedfolder/batch/share, sharedfolder/batch/unshare, sharedfolder/isolate, sharedfolder/list, sharedfolder/member/batch/add, sharedfolder/member/batch/delete, sharedfolder/member/batch/member_folder, sharedfolder/mount/add, sharedfolder/mount/delete, sharedfolder/mount/list, sharedfolder/mount/mountable
- **Sharedlink**: cap/expiry, cap/visibility, delete/links, delete/member, list, update/expiry, update/password, update/visibility
- **Teamfolder**: add, archive, batch/archive, batch/name, batch/permdelete, batch/replication, file/list, file/lock/all/release, file/lock/list, file/lock/release, file/size, list, member/add, member/delete, member/list, partial/replication, permdelete, policy/list, replication, sync/setting/list, sync/setting/update

## Documentation Status Analysis

### Commands with Descriptions (.desc keys)
Based on the analysis of the message resources, the following commands currently have description keys:

1. `citron.dropbox.file.replication.desc` - "This command will replicate files/folders. But it does not include sharing permissions. The command replicates only for folder contents of given path."
2. `citron.dropbox.file.sharedfolder.leave.desc` - "Upon success, the current user will no longer have access to the folder. Please use `dropbox file sharedfolder list` command to find the shared_folder_id of the folder you want to leave."
3. `citron.dropbox.file.sharedfolder.mount.delete.desc` - "Upon success, the current user cannot access the folder unless adding the folder again. Please use `dropbox file sharedfolder mount list` command to find the shared_folder_id of the folder you want to leave."
4. `citron.dropbox.file.sharedlink.delete.desc` - "This command will delete shared links based on the path in the Dropbox"
5. `citron.dropbox.team.activity.event.desc` - (Extensive description with examples)
6. `citron.dropbox.team.backup.device.status.desc` - (Detailed description of backup status evaluation)
7. `citron.dropbox.team.group.delete.desc` - "This command does not confirm whether the group used in existing folders"
8. `citron.dropbox.team.insight.scan.desc` - (Description of team data collection)
9. `citron.dropbox.team.member.file.permdelete.desc` - "Please see https://www.dropbox.com/help/40 for more detail about permanent deletion."
10. `citron.dropbox.team.member.folder.replication.desc` - (Description of folder replication process)

### Commands Missing Descriptions
The vast majority of commands (approximately 200+ commands) only have `.title` keys but lack `.desc` keys for detailed descriptions.

## Specific Focus: File Tag Commands

### Current Status
The file tag commands have the following message keys:

#### citron.dropbox.file.tag.add
- ✅ `.title`: "Add tag to file or folder"
- ❌ `.desc`: Missing
- ✅ `.cli.args`: "-path /DROPBOX/PATH/TO/TARGET -tag TAG_NAME"
- ✅ Flag descriptions: base_path, path, peer, tag

#### citron.dropbox.file.tag.delete
- ✅ `.title`: "Delete a tag from the file/folder"
- ❌ `.desc`: Missing
- ✅ `.cli.args`: "-path /DROPBOX/PATH/TO/PROCESS -tag TAG_NAME"
- ✅ Flag descriptions: base_path, path, peer, tag

#### citron.dropbox.file.tag.list
- ✅ `.title`: "List tags of the path"
- ❌ `.desc`: Missing
- ✅ `.cli.args`: "-path /DROPBOX/PATH/TO/TARGET"
- ✅ Flag descriptions: base_path, path, peer

### Recommended Descriptions

#### citron.dropbox.file.tag.add.desc
"Add a custom tag to a file or folder in Dropbox. Tags help organize and categorize your content for easier searching and management. You can add multiple tags to the same file or folder."

#### citron.dropbox.file.tag.delete.desc
"Remove a specific tag from a file or folder in Dropbox. This operation only removes the tag association and does not affect the file or folder itself. Use this command to clean up outdated or incorrect tags."

#### citron.dropbox.file.tag.list.desc
"Display all tags associated with a specific file or folder in Dropbox. This command helps you see what tags have been applied to organize and categorize your content. The output shows each tag along with the file path."

## Implementation Strategy

### Phase 1: File Tag Commands (Immediate)
1. Add missing `.desc` keys for the three file tag commands
2. Test with `dev build preflight` to ensure proper integration

### Phase 2: Core File Operations (High Priority)
Focus on the most commonly used file operations:
- copy, delete, info, list, move, size
- sync/up, sync/down
- search/content, search/name

### Phase 3: Team Operations (Medium Priority)
Focus on essential team management commands:
- member/list, member/batch operations
- teamfolder operations
- activity/event

### Phase 4: Advanced Operations (Lower Priority)
- Complex batch operations
- Advanced sharing and permission management
- Specialized reporting commands

## Next Steps

1. **Immediate Action**: Add description keys for file tag commands
2. **Systematic Review**: Go through each command category and add descriptions based on code analysis
3. **Testing**: Run `dev build preflight` after each batch of changes
4. **Documentation Generation**: Verify that the new descriptions appear correctly in generated documentation
5. **Localization**: Add corresponding Japanese translations for all new description keys

## Notes

- All user-facing text must be stored in external resource files (following custom instruction #1)
- Descriptions should explain purpose, behavior, and design intent (following custom instruction #3)
- Keep descriptions concise but comprehensive
- Include usage examples where helpful
- Maintain consistency with existing description patterns 