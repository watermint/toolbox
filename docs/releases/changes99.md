---
layout: release
title: Changes of Release 98
lang: en
---

# Changes between `Release 98` to `Release 99`

# Commands deleted


| Command                                     | Title                                                                     |
|---------------------------------------------|---------------------------------------------------------------------------|
| config disable                              | Disable a feature.                                                        |
| config enable                               | Enable a feature.                                                         |
| config features                             | List available optional features.                                         |
| dev benchmark local                         | Create dummy folder structure in local file system.                       |
| dev benchmark upload                        | Upload benchmark                                                          |
| dev benchmark uploadlink                    | Benchmark single file upload with upload temporary link API.              |
| dev build catalogue                         | Generate catalogue                                                        |
| dev build doc                               | Document generator                                                        |
| dev build info                              | Generate build information file                                           |
| dev build license                           | Generate LICENSE.txt                                                      |
| dev build package                           | Package a build                                                           |
| dev build preflight                         | Process prerequisites for the release                                     |
| dev build readme                            | Generate README.txt                                                       |
| dev ci artifact up                          | Upload CI artifact                                                        |
| dev ci auth connect                         | Authenticate for generating end to end testing                            |
| dev ci auth import                          | Import auth tokens of end to end test from environment variable           |
| dev diag endpoint                           | List endpoints                                                            |
| dev diag throughput                         | Evaluate throughput from capture logs                                     |
| dev kvs dump                                | Dump KVS data                                                             |
| dev release candidate                       | Validate release candidate                                                |
| dev release doc                             | Generate release documents                                                |
| dev release publish                         | Publish release                                                           |
| dev replay approve                          | Approve the replay as test bundle                                         |
| dev replay bundle                           | Run all replays                                                           |
| dev replay recipe                           | Replay recipe                                                             |
| dev replay remote                           | Run remote replay bundle                                                  |
| dev spec diff                               | Compare spec of two releases                                              |
| dev spec doc                                | Generate spec docs                                                        |
| dev stage dbxfs                             | Verify Dropbox File System impl. for cached system                        |
| dev stage gmail                             | Gmail command                                                             |
| dev stage griddata                          | Grid data test                                                            |
| dev stage gui launch                        | GUI proof of concept                                                      |
| dev stage http_range                        | HTTP Range request proof of concept                                       |
| dev stage scoped                            | Dropbox scoped OAuth app test                                             |
| dev stage teamfolder                        | Team folder operation sample                                              |
| dev stage upload_append                     | New upload API test                                                       |
| dev test auth all                           | Test for connect to Dropbox with all scopes                               |
| dev test echo                               | Echo text                                                                 |
| dev test panic                              | Panic test                                                                |
| dev test recipe                             | Test recipe                                                               |
| dev test resources                          | Binary quality test                                                       |
| dev test setup teamsharedlink               | Create demo shared links                                                  |
| dev util anonymise                          | Anonymise capture log                                                     |
| dev util curl                               | Generate cURL preview from capture log                                    |
| dev util image jpeg                         | Create dummy image files                                                  |
| dev util wait                               | Wait for specified seconds                                                |
| file archive local                          | Archive local files                                                       |
| file compare account                        | Compare files of two accounts                                             |
| file compare local                          | Compare local folders and Dropbox folders                                 |
| file copy                                   | Copy files                                                                |
| file delete                                 | Delete file or folder                                                     |
| file dispatch local                         | Dispatch local files                                                      |
| file export doc                             | Export document                                                           |
| file export url                             | Export a document from the URL                                            |
| file import batch url                       | Batch import files from URL                                               |
| file import url                             | Import file from the URL                                                  |
| file info                                   | Resolve metadata of the path                                              |
| file list                                   | List files and folders                                                    |
| file lock acquire                           | Lock a file                                                               |
| file lock all release                       | Release all locks under the specified path                                |
| file lock batch acquire                     | Lock multiple files                                                       |
| file lock batch release                     | Release multiple locks                                                    |
| file lock list                              | List locks under the specified path                                       |
| file lock release                           | Release a lock                                                            |
| file merge                                  | Merge paths                                                               |
| file mount list                             | List mounted/unmounted shared folders                                     |
| file move                                   | Move files                                                                |
| file paper append                           | Append the content to the end of the existing Paper doc                   |
| file paper create                           | Create new Paper in the path                                              |
| file paper overwrite                        | Overwrite existing Paper document                                         |
| file paper prepend                          | Append the content to the beginning of the existing Paper doc             |
| file replication                            | Replicate file content to the other account                               |
| file restore all                            | Restore files under given path                                            |
| file search content                         | Search file content                                                       |
| file search name                            | Search file name                                                          |
| file size                                   | Storage usage                                                             |
| file sync down                              | Downstream sync with Dropbox                                              |
| file sync online                            | Sync online files                                                         |
| file sync up                                | Upstream sync with Dropbox                                                |
| file watch                                  | Watch file activities                                                     |
| filerequest create                          | Create a file request                                                     |
| filerequest delete closed                   | Delete all closed file requests on this account.                          |
| filerequest delete url                      | Delete a file request by the file request URL                             |
| filerequest list                            | List file requests of the individual account                              |
| group add                                   | Create new group                                                          |
| group batch add                             | Bulk adding groups                                                        |
| group batch delete                          | Delete groups                                                             |
| group delete                                | Delete group                                                              |
| group folder list                           | Find folders of each group                                                |
| group list                                  | List group(s)                                                             |
| group member add                            | Add a member to the group                                                 |
| group member batch add                      | Bulk add members into groups                                              |
| group member batch delete                   | Delete members from groups                                                |
| group member batch update                   | Add or delete members from groups                                         |
| group member delete                         | Delete a member from the group                                            |
| group member list                           | List members of groups                                                    |
| group rename                                | Rename the group                                                          |
| image info                                  | Show EXIF information of the image file                                   |
| job history archive                         | Archive jobs                                                              |
| job history delete                          | Delete old job history                                                    |
| job history list                            | Show job history                                                          |
| job history ship                            | Ship Job logs to Dropbox path                                             |
| job log jobid                               | Retrieve logs of specified Job ID                                         |
| job log kind                                | Concatenate and print logs of specified log kind                          |
| job log last                                | Print the last job log files                                              |
| license                                     | Show license information                                                  |
| member batch suspend                        | Bulk suspend members                                                      |
| member batch unsuspend                      | Bulk unsuspend members                                                    |
| member clear externalid                     | Clear external_id of members                                              |
| member delete                               | Delete members                                                            |
| member detach                               | Convert Dropbox Business accounts to a Basic account                      |
| member feature                              | List member feature settings                                              |
| member file lock all release                | Release all locks under the path of the member                            |
| member file lock list                       | List locks of the member under the path                                   |
| member file lock release                    | Release the lock of the path as the member                                |
| member file permdelete                      | Permanently delete the file or folder at a given path of the team member. |
| member folder list                          | Find folders for each member                                              |
| member folder replication                   | Replicate a folder to another member's personal folder                    |
| member invite                               | Invite member(s)                                                          |
| member list                                 | List team member(s)                                                       |
| member quota list                           | List team member quota                                                    |
| member quota update                         | Update team member quota                                                  |
| member quota usage                          | List team member storage usage                                            |
| member reinvite                             | Reinvite invited status members to the team                               |
| member replication                          | Replicate team member files                                               |
| member suspend                              | Suspend a member                                                          |
| member unsuspend                            | Unsuspend a member                                                        |
| member update email                         | Member email operation                                                    |
| member update externalid                    | Update External ID of team members                                        |
| member update invisible                     | Enable directory restriction to members                                   |
| member update profile                       | Update member profile                                                     |
| member update visible                       | Disable directory restriction to members                                  |
| services asana team list                    | List team                                                                 |
| services asana team project list            | List projects of the team                                                 |
| services asana team task list               | List task of the team                                                     |
| services asana workspace list               | List workspaces                                                           |
| services asana workspace project list       | List projects of the workspace                                            |
| services dropbox user feature               | List feature settings for current user                                    |
| services github content get                 | Get content metadata of the repository                                    |
| services github content put                 | Put small text content into the repository                                |
| services github issue list                  | List issues of the public/private GitHub repository                       |
| services github profile                     | Get the authenticated user                                                |
| services github release asset download      | Download assets                                                           |
| services github release asset list          | List assets of GitHub Release                                             |
| services github release asset upload        | Upload assets file into the GitHub Release                                |
| services github release draft               | Create release draft                                                      |
| services github release list                | List releases                                                             |
| services github tag create                  | Create a tag on the repository                                            |
| services google mail filter add             | Add a filter.                                                             |
| services google mail filter batch add       | Batch adding/deleting labels with query                                   |
| services google mail filter delete          | Delete a filter                                                           |
| services google mail filter list            | List filters                                                              |
| services google mail label add              | Add a label                                                               |
| services google mail label delete           | Delete a label                                                            |
| services google mail label list             | List email labels                                                         |
| services google mail label rename           | Rename a label                                                            |
| services google mail message label add      | Add labels to the message                                                 |
| services google mail message label delete   | Remove labels from the message                                            |
| services google mail message list           | List messages                                                             |
| services google mail message processed list | List messages in processed format.                                        |
| services google mail message send           | Send a mail                                                               |
| services google mail sendas add             | Creates a custom "from" send-as alias                                     |
| services google mail sendas delete          | Deletes the specified send-as alias                                       |
| services google mail sendas list            | Lists the send-as aliases for the specified account                       |
| services google mail thread list            | List threads                                                              |
| services google sheets sheet append         | Append data to a spreadsheet                                              |
| services google sheets sheet clear          | Clears values from a spreadsheet                                          |
| services google sheets sheet export         | Export sheet data                                                         |
| services google sheets sheet import         | Import data into the spreadsheet                                          |
| services google sheets sheet list           | List sheets of the spreadsheet                                            |
| services google sheets spreadsheet create   | Create a new spreadsheet                                                  |
| services slack conversation list            | List channels                                                             |
| sharedfolder list                           | List shared folder(s)                                                     |
| sharedfolder member list                    | List shared folder member(s)                                              |
| sharedlink create                           | Create shared link                                                        |
| sharedlink delete                           | Remove shared links                                                       |
| sharedlink file list                        | List files for the shared link                                            |
| sharedlink info                             | Get information about the shared link                                     |
| sharedlink list                             | List of shared link(s)                                                    |
| team activity batch user                    | Scan activities for multiple users                                        |
| team activity daily event                   | Report activities by day                                                  |
| team activity event                         | Event log                                                                 |
| team activity user                          | Activities log per user                                                   |
| team admin group role add                   | Add the role to members of the group                                      |
| team admin group role delete                | Delete the role from all members except of members of the exception group |
| team admin list                             | List admin roles of members                                               |
| team admin role add                         | Add a new role to the member                                              |
| team admin role clear                       | Remove all admin roles from the member                                    |
| team admin role delete                      | Remove a role from the member                                             |
| team admin role list                        | List admin roles of the team                                              |
| team content legacypaper count              | Count number of Paper documents per member                                |
| team content legacypaper export             | Export entire team member Paper documents into local path                 |
| team content legacypaper list               | List team member Paper documents                                          |
| team content member list                    | List team folder & shared folder members                                  |
| team content member size                    | Count number of members of team folders and shared folders                |
| team content mount list                     | List all mounted/unmounted shared folders of team members.                |
| team content policy list                    | List policies of team folders and shared folders in the team              |
| team device list                            | List all devices/sessions in the team                                     |
| team device unlink                          | Unlink device sessions                                                    |
| team feature                                | Team feature                                                              |
| team filerequest clone                      | Clone file requests by given data                                         |
| team filerequest list                       | List all file requests in the team                                        |
| team info                                   | Team information                                                          |
| team linkedapp list                         | List linked applications                                                  |
| team namespace file list                    | List all files and folders of the team namespaces                         |
| team namespace file size                    | List all files and folders of the team namespaces                         |
| team namespace list                         | List all namespaces of the team                                           |
| team namespace member list                  | List members of shared folders and team folders in the team               |
| team report activity                        | Activities report                                                         |
| team report devices                         | Devices report                                                            |
| team report membership                      | Membership report                                                         |
| team report storage                         | Storage report                                                            |
| team sharedlink cap expiry                  | Set expiry cap to shared links in the team                                |
| team sharedlink cap visibility              | Set visibility cap to shared links in the team                            |
| team sharedlink delete links                | Batch delete shared links                                                 |
| team sharedlink delete member               | Delete all shared links of the member                                     |
| team sharedlink list                        | List of shared links                                                      |
| team sharedlink update expiry               | Update expiration date of public shared links within the team             |
| team sharedlink update password             | Set or update shared link passwords                                       |
| team sharedlink update visibility           | Update visibility of shared links                                         |
| teamfolder add                              | Add team folder to the team                                               |
| teamfolder archive                          | Archive team folder                                                       |
| teamfolder batch archive                    | Archiving team folders                                                    |
| teamfolder batch permdelete                 | Permanently delete team folders                                           |
| teamfolder batch replication                | Batch replication of team folders                                         |
| teamfolder file list                        | List files in team folders                                                |
| teamfolder file lock all release            | Release all locks under the path of the team folder                       |
| teamfolder file lock list                   | List locks in the team folder                                             |
| teamfolder file lock release                | Release lock of the path in the team folder                               |
| teamfolder file size                        | Calculate size of team folders                                            |
| teamfolder list                             | List team folder(s)                                                       |
| teamfolder member add                       | Batch adding users/groups to team folders                                 |
| teamfolder member delete                    | Batch removing users/groups from team folders                             |
| teamfolder member list                      | List team folder members                                                  |
| teamfolder partial replication              | Partial team folder replication to the other team                         |
| teamfolder permdelete                       | Permanently delete team folder                                            |
| teamfolder policy list                      | List policies of team folders                                             |
| teamfolder replication                      | Replicate a team folder to the other team                                 |
| util date today                             | Display current date                                                      |
| util datetime now                           | Display current date/time                                                 |
| util decode base_32                         | Decode text from Base32 (RFC 4648) format                                 |
| util decode base_64                         | Decode text from Base64 (RFC 4648) format                                 |
| util encode base_32                         | Encode text into Base32 (RFC 4648) format                                 |
| util encode base_64                         | Encode text into Base64 (RFC 4648) format                                 |
| util git clone                              | Clone git repository                                                      |
| util image placeholder                      | Create placeholder image                                                  |
| util qrcode create                          | Create a QR code image file                                               |
| util qrcode wifi                            | Generate QR code for WIFI configuration                                   |
| util time now                               | Display current time                                                      |
| util unixtime format                        | Time format to convert the unix time (epoch seconds from 1970-01-01)      |
| util unixtime now                           | Display current time in unixtime                                          |
| util xlsx create                            | Create an empty spreadsheet                                               |
| util xlsx sheet export                      | Export data from the xlsx file                                            |
| util xlsx sheet import                      | Import data into xlsx file                                                |
| util xlsx sheet list                        | List sheets of the xlsx file                                              |
| version                                     | Show version                                                              |



