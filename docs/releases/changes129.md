---
layout: release
title: Changes of Release 128
lang: en
---

# Changes between `Release 128` to `Release 129`

# Commands added


| Command                                             | Title                                                                               |
|-----------------------------------------------------|-------------------------------------------------------------------------------------|
| dropbox file compare account                        | Compare files of two accounts                                                       |
| dropbox file compare local                          | Compare local folders and Dropbox folders                                           |
| dropbox file copy                                   | Copy files                                                                          |
| dropbox file delete                                 | Delete file or folder                                                               |
| dropbox file export doc                             | Export document                                                                     |
| dropbox file export url                             | Export a document from the URL                                                      |
| dropbox file import batch url                       | Batch import files from URL                                                         |
| dropbox file import url                             | Import file from the URL                                                            |
| dropbox file info                                   | Resolve metadata of the path                                                        |
| dropbox file list                                   | List files and folders                                                              |
| dropbox file lock acquire                           | Lock a file                                                                         |
| dropbox file lock all release                       | Release all locks under the specified path                                          |
| dropbox file lock batch acquire                     | Lock multiple files                                                                 |
| dropbox file lock batch release                     | Release multiple locks                                                              |
| dropbox file lock list                              | List locks under the specified path                                                 |
| dropbox file lock release                           | Release a lock                                                                      |
| dropbox file merge                                  | Merge paths                                                                         |
| dropbox file move                                   | Move files                                                                          |
| dropbox file replication                            | Replicate file content to the other account                                         |
| dropbox file request create                         | Create a file request                                                               |
| dropbox file request delete closed                  | Delete all closed file requests on this account.                                    |
| dropbox file request delete url                     | Delete a file request by the file request URL                                       |
| dropbox file request list                           | List file requests of the individual account                                        |
| dropbox file restore all                            | Restore files under given path                                                      |
| dropbox file revision download                      | Download the file revision                                                          |
| dropbox file revision list                          | List file revisions                                                                 |
| dropbox file revision restore                       | Restore the file revision                                                           |
| dropbox file search content                         | Search file content                                                                 |
| dropbox file search name                            | Search file name                                                                    |
| dropbox file share info                             | Retrieve sharing information of the file                                            |
| dropbox file sharedfolder leave                     | Leave from the shared folder                                                        |
| dropbox file sharedfolder list                      | List shared folder(s)                                                               |
| dropbox file sharedfolder member add                | Add a member to the shared folder                                                   |
| dropbox file sharedfolder member delete             | Delete a member from the shared folder                                              |
| dropbox file sharedfolder member list               | List shared folder member(s)                                                        |
| dropbox file sharedfolder mount add                 | Add the shared folder to the current user's Dropbox                                 |
| dropbox file sharedfolder mount delete              | The current user unmounts the designated folder.                                    |
| dropbox file sharedfolder mount list                | List all shared folders the current user mounted                                    |
| dropbox file sharedfolder mount mountable           | List all shared folders the current user can mount                                  |
| dropbox file sharedfolder share                     | Share a folder                                                                      |
| dropbox file sharedfolder unshare                   | Unshare a folder                                                                    |
| dropbox file sharedlink create                      | Create shared link                                                                  |
| dropbox file sharedlink delete                      | Remove shared links                                                                 |
| dropbox file sharedlink file list                   | List files for the shared link                                                      |
| dropbox file sharedlink info                        | Get information about the shared link                                               |
| dropbox file sharedlink list                        | List of shared link(s)                                                              |
| dropbox file size                                   | Storage usage                                                                       |
| dropbox file sync down                              | Downstream sync with Dropbox                                                        |
| dropbox file sync online                            | Sync online files                                                                   |
| dropbox file sync up                                | Upstream sync with Dropbox                                                          |
| dropbox file tag add                                | Add a tag to the file/folder                                                        |
| dropbox file tag delete                             | Delete a tag from the file/folder                                                   |
| dropbox file tag list                               | List tags of the path                                                               |
| dropbox file template apply                         | Apply file/folder structure template to the Dropbox path                            |
| dropbox file template capture                       | Capture file/folder structure as template from Dropbox path                         |
| dropbox file watch                                  | Watch file activities                                                               |
| dropbox paper append                                | Append the content to the end of the existing Paper doc                             |
| dropbox paper create                                | Create new Paper in the path                                                        |
| dropbox paper overwrite                             | Overwrite existing Paper document                                                   |
| dropbox paper prepend                               | Append the content to the beginning of the existing Paper doc                       |
| dropbox team activity batch user                    | Scan activities for multiple users                                                  |
| dropbox team activity daily event                   | Report activities by day                                                            |
| dropbox team activity event                         | Event log                                                                           |
| dropbox team activity user                          | Activities log per user                                                             |
| dropbox team admin group role add                   | Add the role to members of the group                                                |
| dropbox team admin group role delete                | Delete the role from all members except of members of the exception group           |
| dropbox team admin list                             | List admin roles of members                                                         |
| dropbox team admin role add                         | Add a new role to the member                                                        |
| dropbox team admin role clear                       | Remove all admin roles from the member                                              |
| dropbox team admin role delete                      | Remove a role from the member                                                       |
| dropbox team admin role list                        | List admin roles of the team                                                        |
| dropbox team content legacypaper count              | Count number of Paper documents per member                                          |
| dropbox team content legacypaper export             | Export entire team member Paper documents into local path                           |
| dropbox team content legacypaper list               | List team member Paper documents                                                    |
| dropbox team content member list                    | List team folder & shared folder members                                            |
| dropbox team content member size                    | Count number of members of team folders and shared folders                          |
| dropbox team content mount list                     | List all mounted/unmounted shared folders of team members.                          |
| dropbox team content policy list                    | List policies of team folders and shared folders in the team                        |
| dropbox team device list                            | List all devices/sessions in the team                                               |
| dropbox team device unlink                          | Unlink device sessions                                                              |
| dropbox team feature                                | Team feature                                                                        |
| dropbox team filerequest clone                      | Clone file requests by given data                                                   |
| dropbox team filerequest list                       | List all file requests in the team                                                  |
| dropbox team filesystem                             | Identify team's file system version                                                 |
| dropbox team group add                              | Create new group                                                                    |
| dropbox team group batch add                        | Bulk adding groups                                                                  |
| dropbox team group batch delete                     | Delete groups                                                                       |
| dropbox team group clear externalid                 | Clear an external ID of a group                                                     |
| dropbox team group delete                           | Delete group                                                                        |
| dropbox team group folder list                      | List folders of each group                                                          |
| dropbox team group list                             | List group(s)                                                                       |
| dropbox team group member add                       | Add a member to the group                                                           |
| dropbox team group member batch add                 | Bulk add members into groups                                                        |
| dropbox team group member batch delete              | Delete members from groups                                                          |
| dropbox team group member batch update              | Add or delete members from groups                                                   |
| dropbox team group member delete                    | Delete a member from the group                                                      |
| dropbox team group member list                      | List members of groups                                                              |
| dropbox team group rename                           | Rename the group                                                                    |
| dropbox team group update type                      | Update group management type                                                        |
| dropbox team info                                   | Team information                                                                    |
| dropbox team insight report teamfoldermember        | Report team folder members                                                          |
| dropbox team insight scan                           | Scans team data for analysis                                                        |
| dropbox team insight scanretry                      | Retry scan for errors on the last scan                                              |
| dropbox team insight summarize                      | Summarize team data for analysis                                                    |
| dropbox team legalhold add                          | Creates new legal hold policy.                                                      |
| dropbox team legalhold list                         | Retrieve existing policies                                                          |
| dropbox team legalhold member batch update          | Update member list of legal hold policy                                             |
| dropbox team legalhold member list                  | List members of the legal hold                                                      |
| dropbox team legalhold release                      | Releases a legal hold by Id                                                         |
| dropbox team legalhold revision list                | List revisions of the legal hold policy                                             |
| dropbox team legalhold update desc                  | Update description of the legal hold policy                                         |
| dropbox team legalhold update name                  | Update name of the legal hold policy                                                |
| dropbox team linkedapp list                         | List linked applications                                                            |
| dropbox team member batch delete                    | Delete members                                                                      |
| dropbox team member batch detach                    | Convert Dropbox for teams accounts to a Basic account                               |
| dropbox team member batch invite                    | Invite member(s)                                                                    |
| dropbox team member batch reinvite                  | Reinvite invited status members to the team                                         |
| dropbox team member batch suspend                   | Bulk suspend members                                                                |
| dropbox team member batch unsuspend                 | Bulk unsuspend members                                                              |
| dropbox team member clear externalid                | Clear external_id of members                                                        |
| dropbox team member feature                         | List member feature settings                                                        |
| dropbox team member file lock all release           | Release all locks under the path of the member                                      |
| dropbox team member file lock list                  | List locks of the member under the path                                             |
| dropbox team member file lock release               | Release the lock of the path as the member                                          |
| dropbox team member file permdelete                 | Permanently delete the file or folder at a given path of the team member.           |
| dropbox team member folder list                     | List folders for each member                                                        |
| dropbox team member folder replication              | Replicate a folder to another member's personal folder                              |
| dropbox team member list                            | List team member(s)                                                                 |
| dropbox team member quota batch update              | Update team member quota                                                            |
| dropbox team member quota list                      | List team member quota                                                              |
| dropbox team member quota usage                     | List team member storage usage                                                      |
| dropbox team member replication                     | Replicate team member files                                                         |
| dropbox team member suspend                         | Suspend a member                                                                    |
| dropbox team member unsuspend                       | Unsuspend a member                                                                  |
| dropbox team member update batch email              | Member email operation                                                              |
| dropbox team member update batch externalid         | Update External ID of team members                                                  |
| dropbox team member update batch invisible          | Enable directory restriction to members                                             |
| dropbox team member update batch profile            | Update member profile                                                               |
| dropbox team member update batch visible            | Disable directory restriction to members                                            |
| dropbox team namespace file list                    | List all files and folders of the team namespaces                                   |
| dropbox team namespace file size                    | List all files and folders of the team namespaces                                   |
| dropbox team namespace list                         | List all namespaces of the team                                                     |
| dropbox team namespace member list                  | List members of shared folders and team folders in the team                         |
| dropbox team namespace summary                      | Report team namespace status summary.                                               |
| dropbox team report activity                        | Activities report                                                                   |
| dropbox team report devices                         | Devices report                                                                      |
| dropbox team report membership                      | Membership report                                                                   |
| dropbox team report storage                         | Storage report                                                                      |
| dropbox team runas file batch copy                  | Batch copy files/folders as a member                                                |
| dropbox team runas file list                        | List files and folders run as a member                                              |
| dropbox team runas file sync batch up               | Batch sync up that run as members                                                   |
| dropbox team runas sharedfolder batch leave         | Batch leave from shared folders as a member                                         |
| dropbox team runas sharedfolder batch share         | Batch share folders for members                                                     |
| dropbox team runas sharedfolder batch unshare       | Batch unshare folders for members                                                   |
| dropbox team runas sharedfolder isolate             | Unshare owned shared folders and leave from external shared folders run as a member |
| dropbox team runas sharedfolder list                | List shared folders run as the member                                               |
| dropbox team runas sharedfolder member batch add    | Batch add members to member's shared folders                                        |
| dropbox team runas sharedfolder member batch delete | Batch delete members from member's shared folders                                   |
| dropbox team runas sharedfolder mount add           | Add the shared folder to the specified member's Dropbox                             |
| dropbox team runas sharedfolder mount delete        | The specified user unmounts the designated folder.                                  |
| dropbox team runas sharedfolder mount list          | List all shared folders the specified member mounted                                |
| dropbox team runas sharedfolder mount mountable     | List all shared folders the member can mount                                        |
| dropbox team sharedlink cap expiry                  | Set expiry cap to shared links in the team                                          |
| dropbox team sharedlink cap visibility              | Set visibility cap to shared links in the team                                      |
| dropbox team sharedlink delete links                | Batch delete shared links                                                           |
| dropbox team sharedlink delete member               | Delete all shared links of the member                                               |
| dropbox team sharedlink list                        | List of shared links                                                                |
| dropbox team sharedlink update expiry               | Update expiration date of public shared links within the team                       |
| dropbox team sharedlink update password             | Set or update shared link passwords                                                 |
| dropbox team sharedlink update visibility           | Update visibility of shared links                                                   |
| dropbox team teamfolder add                         | Add team folder to the team                                                         |
| dropbox team teamfolder archive                     | Archive team folder                                                                 |
| dropbox team teamfolder batch archive               | Archiving team folders                                                              |
| dropbox team teamfolder batch permdelete            | Permanently delete team folders                                                     |
| dropbox team teamfolder batch replication           | Batch replication of team folders                                                   |
| dropbox team teamfolder file list                   | List files in team folders                                                          |
| dropbox team teamfolder file lock all release       | Release all locks under the path of the team folder                                 |
| dropbox team teamfolder file lock list              | List locks in the team folder                                                       |
| dropbox team teamfolder file lock release           | Release lock of the path in the team folder                                         |
| dropbox team teamfolder file size                   | Calculate size of team folders                                                      |
| dropbox team teamfolder list                        | List team folder(s)                                                                 |
| dropbox team teamfolder member add                  | Batch adding users/groups to team folders                                           |
| dropbox team teamfolder member delete               | Batch removing users/groups from team folders                                       |
| dropbox team teamfolder member list                 | List team folder members                                                            |
| dropbox team teamfolder partial replication         | Partial team folder replication to the other team                                   |
| dropbox team teamfolder permdelete                  | Permanently delete team folder                                                      |
| dropbox team teamfolder policy list                 | List policies of team folders                                                       |
| dropbox team teamfolder replication                 | Replicate a team folder to the other team                                           |
| dropbox team teamfolder sync setting list           | List team folder sync settings                                                      |
| dropbox team teamfolder sync setting update         | Batch update team folder sync settings                                              |
| local file template apply                           | Apply file/folder structure template to the local path                              |
| local file template capture                         | Capture file/folder structure as template from local path                           |



# Commands deleted


| Command                                     | Title                                                                               |
|---------------------------------------------|-------------------------------------------------------------------------------------|
| file compare account                        | Compare files of two accounts                                                       |
| file compare local                          | Compare local folders and Dropbox folders                                           |
| file copy                                   | Copy files                                                                          |
| file delete                                 | Delete file or folder                                                               |
| file export doc                             | Export document                                                                     |
| file export url                             | Export a document from the URL                                                      |
| file import batch url                       | Batch import files from URL                                                         |
| file import url                             | Import file from the URL                                                            |
| file info                                   | Resolve metadata of the path                                                        |
| file list                                   | List files and folders                                                              |
| file lock acquire                           | Lock a file                                                                         |
| file lock all release                       | Release all locks under the specified path                                          |
| file lock batch acquire                     | Lock multiple files                                                                 |
| file lock batch release                     | Release multiple locks                                                              |
| file lock list                              | List locks under the specified path                                                 |
| file lock release                           | Release a lock                                                                      |
| file merge                                  | Merge paths                                                                         |
| file move                                   | Move files                                                                          |
| file paper append                           | Append the content to the end of the existing Paper doc                             |
| file paper create                           | Create new Paper in the path                                                        |
| file paper overwrite                        | Overwrite existing Paper document                                                   |
| file paper prepend                          | Append the content to the beginning of the existing Paper doc                       |
| file replication                            | Replicate file content to the other account                                         |
| file restore all                            | Restore files under given path                                                      |
| file revision download                      | Download the file revision                                                          |
| file revision list                          | List file revisions                                                                 |
| file revision restore                       | Restore the file revision                                                           |
| file search content                         | Search file content                                                                 |
| file search name                            | Search file name                                                                    |
| file share info                             | Retrieve sharing information of the file                                            |
| file size                                   | Storage usage                                                                       |
| file sync down                              | Downstream sync with Dropbox                                                        |
| file sync online                            | Sync online files                                                                   |
| file sync up                                | Upstream sync with Dropbox                                                          |
| file tag add                                | Add a tag to the file/folder                                                        |
| file tag delete                             | Delete a tag from the file/folder                                                   |
| file tag list                               | List tags of the path                                                               |
| file template apply local                   | Apply file/folder structure template to the local path                              |
| file template apply remote                  | Apply file/folder structure template to the Dropbox path                            |
| file template capture local                 | Capture file/folder structure as template from local path                           |
| file template capture remote                | Capture file/folder structure as template from Dropbox path                         |
| file watch                                  | Watch file activities                                                               |
| filerequest create                          | Create a file request                                                               |
| filerequest delete closed                   | Delete all closed file requests on this account.                                    |
| filerequest delete url                      | Delete a file request by the file request URL                                       |
| filerequest list                            | List file requests of the individual account                                        |
| group add                                   | Create new group                                                                    |
| group batch add                             | Bulk adding groups                                                                  |
| group batch delete                          | Delete groups                                                                       |
| group clear externalid                      | Clear an external ID of a group                                                     |
| group delete                                | Delete group                                                                        |
| group folder list                           | List folders of each group                                                          |
| group list                                  | List group(s)                                                                       |
| group member add                            | Add a member to the group                                                           |
| group member batch add                      | Bulk add members into groups                                                        |
| group member batch delete                   | Delete members from groups                                                          |
| group member batch update                   | Add or delete members from groups                                                   |
| group member delete                         | Delete a member from the group                                                      |
| group member list                           | List members of groups                                                              |
| group rename                                | Rename the group                                                                    |
| group update type                           | Update group management type                                                        |
| member batch suspend                        | Bulk suspend members                                                                |
| member batch unsuspend                      | Bulk unsuspend members                                                              |
| member clear externalid                     | Clear external_id of members                                                        |
| member delete                               | Delete members                                                                      |
| member detach                               | Convert Dropbox for teams accounts to a Basic account                               |
| member feature                              | List member feature settings                                                        |
| member file lock all release                | Release all locks under the path of the member                                      |
| member file lock list                       | List locks of the member under the path                                             |
| member file lock release                    | Release the lock of the path as the member                                          |
| member file permdelete                      | Permanently delete the file or folder at a given path of the team member.           |
| member folder list                          | List folders for each member                                                        |
| member folder replication                   | Replicate a folder to another member's personal folder                              |
| member invite                               | Invite member(s)                                                                    |
| member list                                 | List team member(s)                                                                 |
| member quota list                           | List team member quota                                                              |
| member quota update                         | Update team member quota                                                            |
| member quota usage                          | List team member storage usage                                                      |
| member reinvite                             | Reinvite invited status members to the team                                         |
| member replication                          | Replicate team member files                                                         |
| member suspend                              | Suspend a member                                                                    |
| member unsuspend                            | Unsuspend a member                                                                  |
| member update email                         | Member email operation                                                              |
| member update externalid                    | Update External ID of team members                                                  |
| member update invisible                     | Enable directory restriction to members                                             |
| member update profile                       | Update member profile                                                               |
| member update visible                       | Disable directory restriction to members                                            |
| sharedfolder leave                          | Leave from the shared folder                                                        |
| sharedfolder list                           | List shared folder(s)                                                               |
| sharedfolder member add                     | Add a member to the shared folder                                                   |
| sharedfolder member delete                  | Delete a member from the shared folder                                              |
| sharedfolder member list                    | List shared folder member(s)                                                        |
| sharedfolder mount add                      | Add the shared folder to the current user's Dropbox                                 |
| sharedfolder mount delete                   | The current user unmounts the designated folder.                                    |
| sharedfolder mount list                     | List all shared folders the current user mounted                                    |
| sharedfolder mount mountable                | List all shared folders the current user can mount                                  |
| sharedfolder share                          | Share a folder                                                                      |
| sharedfolder unshare                        | Unshare a folder                                                                    |
| sharedlink create                           | Create shared link                                                                  |
| sharedlink delete                           | Remove shared links                                                                 |
| sharedlink file list                        | List files for the shared link                                                      |
| sharedlink info                             | Get information about the shared link                                               |
| sharedlink list                             | List of shared link(s)                                                              |
| team activity batch user                    | Scan activities for multiple users                                                  |
| team activity daily event                   | Report activities by day                                                            |
| team activity event                         | Event log                                                                           |
| team activity user                          | Activities log per user                                                             |
| team admin group role add                   | Add the role to members of the group                                                |
| team admin group role delete                | Delete the role from all members except of members of the exception group           |
| team admin list                             | List admin roles of members                                                         |
| team admin role add                         | Add a new role to the member                                                        |
| team admin role clear                       | Remove all admin roles from the member                                              |
| team admin role delete                      | Remove a role from the member                                                       |
| team admin role list                        | List admin roles of the team                                                        |
| team content legacypaper count              | Count number of Paper documents per member                                          |
| team content legacypaper export             | Export entire team member Paper documents into local path                           |
| team content legacypaper list               | List team member Paper documents                                                    |
| team content member list                    | List team folder & shared folder members                                            |
| team content member size                    | Count number of members of team folders and shared folders                          |
| team content mount list                     | List all mounted/unmounted shared folders of team members.                          |
| team content policy list                    | List policies of team folders and shared folders in the team                        |
| team device list                            | List all devices/sessions in the team                                               |
| team device unlink                          | Unlink device sessions                                                              |
| team feature                                | Team feature                                                                        |
| team filerequest clone                      | Clone file requests by given data                                                   |
| team filerequest list                       | List all file requests in the team                                                  |
| team filesystem                             | Identify team's file system version                                                 |
| team info                                   | Team information                                                                    |
| team insight scan                           | Scan the team entire information                                                    |
| team insight summarize                      | Summarize scanned team information                                                  |
| team legalhold add                          | Creates new legal hold policy.                                                      |
| team legalhold list                         | Retrieve existing policies                                                          |
| team legalhold member batch update          | Update member list of legal hold policy                                             |
| team legalhold member list                  | List members of the legal hold                                                      |
| team legalhold release                      | Releases a legal hold by Id                                                         |
| team legalhold revision list                | List revisions of the legal hold policy                                             |
| team legalhold update desc                  | Update description of the legal hold policy                                         |
| team legalhold update name                  | Update name of the legal hold policy                                                |
| team linkedapp list                         | List linked applications                                                            |
| team namespace file list                    | List all files and folders of the team namespaces                                   |
| team namespace file size                    | List all files and folders of the team namespaces                                   |
| team namespace list                         | List all namespaces of the team                                                     |
| team namespace member list                  | List members of shared folders and team folders in the team                         |
| team namespace summary                      | Report team namespace status summary.                                               |
| team report activity                        | Activities report                                                                   |
| team report devices                         | Devices report                                                                      |
| team report membership                      | Membership report                                                                   |
| team report storage                         | Storage report                                                                      |
| team runas file batch copy                  | Batch copy files/folders as a member                                                |
| team runas file list                        | List files and folders run as a member                                              |
| team runas file sync batch up               | Batch sync up that run as members                                                   |
| team runas sharedfolder batch leave         | Batch leave from shared folders as a member                                         |
| team runas sharedfolder batch share         | Batch share folders for members                                                     |
| team runas sharedfolder batch unshare       | Batch unshare folders for members                                                   |
| team runas sharedfolder isolate             | Unshare owned shared folders and leave from external shared folders run as a member |
| team runas sharedfolder list                | List shared folders run as the member                                               |
| team runas sharedfolder member batch add    | Batch add members to member's shared folders                                        |
| team runas sharedfolder member batch delete | Batch delete members from member's shared folders                                   |
| team runas sharedfolder mount add           | Add the shared folder to the specified member's Dropbox                             |
| team runas sharedfolder mount delete        | The specified user unmounts the designated folder.                                  |
| team runas sharedfolder mount list          | List all shared folders the specified member mounted                                |
| team runas sharedfolder mount mountable     | List all shared folders the member can mount                                        |
| team sharedlink cap expiry                  | Set expiry cap to shared links in the team                                          |
| team sharedlink cap visibility              | Set visibility cap to shared links in the team                                      |
| team sharedlink delete links                | Batch delete shared links                                                           |
| team sharedlink delete member               | Delete all shared links of the member                                               |
| team sharedlink list                        | List of shared links                                                                |
| team sharedlink update expiry               | Update expiration date of public shared links within the team                       |
| team sharedlink update password             | Set or update shared link passwords                                                 |
| team sharedlink update visibility           | Update visibility of shared links                                                   |
| teamfolder add                              | Add team folder to the team                                                         |
| teamfolder archive                          | Archive team folder                                                                 |
| teamfolder batch archive                    | Archiving team folders                                                              |
| teamfolder batch permdelete                 | Permanently delete team folders                                                     |
| teamfolder batch replication                | Batch replication of team folders                                                   |
| teamfolder file list                        | List files in team folders                                                          |
| teamfolder file lock all release            | Release all locks under the path of the team folder                                 |
| teamfolder file lock list                   | List locks in the team folder                                                       |
| teamfolder file lock release                | Release lock of the path in the team folder                                         |
| teamfolder file size                        | Calculate size of team folders                                                      |
| teamfolder list                             | List team folder(s)                                                                 |
| teamfolder member add                       | Batch adding users/groups to team folders                                           |
| teamfolder member delete                    | Batch removing users/groups from team folders                                       |
| teamfolder member list                      | List team folder members                                                            |
| teamfolder partial replication              | Partial team folder replication to the other team                                   |
| teamfolder permdelete                       | Permanently delete team folder                                                      |
| teamfolder policy list                      | List policies of team folders                                                       |
| teamfolder replication                      | Replicate a team folder to the other team                                           |
| teamfolder sync setting list                | List team folder sync settings                                                      |
| teamfolder sync setting update              | Batch update team folder sync settings                                              |



# Command spec changed: `dev lifecycle planchangepath`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev lifecycle planchangepath",
  	CliArgs: strings.Join({
  		"-announce-url URL -compatibility-file /LOCAL/PATH/TO/compat.json",
  		" -",
+ 		"message-file /LOCAL/PATH/TO/messages.json -",
  		`date "2020-04-01 17:58:38" -current-path RECIPE -former-path REC`,
  		"IPE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 9 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AnnounceUrl", Desc: "Announce URL", TypeName: "string"},
  		&{Name: "Compact", Desc: "Generate compact output", Default: "false", TypeName: "bool", ...},
  		&{Name: "CompatibilityFile", Desc: "Compatibility file", Default: "catalogue/catalogue_compatibility.json", TypeName: "essentials.model.mo_path.file_system_path_impl", ...},
+ 		&{
+ 			Name:     "CurrentBase",
+ 			Desc:     "Current recipe's base path",
+ 			Default:  "citron",
+ 			TypeName: "string",
+ 		},
  		&{Name: "CurrentPath", Desc: "Current CLI path", TypeName: "string"},
  		&{Name: "Date", Desc: "Effective date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(false)}},
+ 		&{
+ 			Name:     "FormerBase",
+ 			Desc:     "Former recipe's base path",
+ 			Default:  "recipe",
+ 			TypeName: "string",
+ 		},
  		&{Name: "FormerPath", Desc: "Former CLI path", TypeName: "string"},
+ 		&{
+ 			Name:     "MessageFile",
+ 			Desc:     "Message file path",
+ 			Default:  "resources/messages/en/messages.json",
+ 			TypeName: "essentials.model.mo_path.file_system_path_impl",
+ 			TypeAttr: map[string]any{"shouldExist": bool(false)},
+ 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
