---
layout: release
title: Changes of Release 140
lang: en
---

# Changes between `Release 140` to `Release 141`

# Commands deleted


| Command                                             | Title                                                                     |
|-----------------------------------------------------|---------------------------------------------------------------------------|
| asana team list                                     | List team                                                                 |
| asana team project list                             | List projects of the team                                                 |
| asana team task list                                | List task of the team                                                     |
| asana workspace list                                | List workspaces                                                           |
| asana workspace project list                        | List projects of the workspace                                            |
| config auth delete                                  | Delete existing auth credential                                           |
| config auth list                                    | List all auth credentials                                                 |
| config feature disable                              | Disable a feature.                                                        |
| config feature enable                               | Enable a feature.                                                         |
| config feature list                                 | List available optional features.                                         |
| config license install                              | Install a license key                                                     |
| config license list                                 | List available license keys                                               |
| deepl translate text                                | Translate text                                                            |
| dev benchmark local                                 | Create dummy folder structure in local file system.                       |
| dev benchmark upload                                | Upload benchmark                                                          |
| dev benchmark uploadlink                            | Benchmark single file upload with upload temporary link API.              |
| dev build catalogue                                 | Generate catalogue                                                        |
| dev build doc                                       | Document generator                                                        |
| dev build info                                      | Generate build information file                                           |
| dev build license                                   | Generate LICENSE.txt                                                      |
| dev build package                                   | Package a build                                                           |
| dev build preflight                                 | Process prerequisites for the release                                     |
| dev build readme                                    | Generate README.txt                                                       |
| dev ci artifact up                                  | Upload CI artifact                                                        |
| dev ci auth export                                  | Export deploy token data for CI build                                     |
| dev diag endpoint                                   | List endpoints                                                            |
| dev diag throughput                                 | Evaluate throughput from capture logs                                     |
| dev doc markdown                                    | Generate messages from markdown source                                    |
| dev info                                            | Dev information                                                           |
| dev kvs concurrency                                 | Concurrency test for KVS engine                                           |
| dev kvs dump                                        | Dump KVS data                                                             |
| dev license issue                                   | Issue a license                                                           |
| dev lifecycle assets                                | Remove deprecated assets                                                  |
| dev lifecycle planchangepath                        | Add plan of changing path to commands                                     |
| dev lifecycle planprune                             | Add plan of the command discontinuation                                   |
| dev module list                                     | Dependent module list                                                     |
| dev placeholder pathchange                          | Placeholder command for path change document generation                   |
| dev placeholder prune                               | Placeholder of prune workflow messages                                    |
| dev release announcement                            | Update announcements                                                      |
| dev release asset                                   | Commit a file to a repository                                             |
| dev release asseturl                                | Update asset URL of the release                                           |
| dev release candidate                               | Validate release candidate                                                |
| dev release checkin                                 | Check in the new release                                                  |
| dev release doc                                     | Generate release documents                                                |
| dev release publish                                 | Publish release                                                           |
| dev replay approve                                  | Approve the replay as test bundle                                         |
| dev replay bundle                                   | Run all replays                                                           |
| dev replay recipe                                   | Replay recipe                                                             |
| dev replay remote                                   | Run remote replay bundle                                                  |
| dev spec diff                                       | Compare spec of two releases                                              |
| dev spec doc                                        | Generate spec docs                                                        |
| dev test echo                                       | Echo text                                                                 |
| dev test license                                    | Testing license required logic                                            |
| dev test panic                                      | Panic test                                                                |
| dev test recipe                                     | Test recipe                                                               |
| dev test resources                                  | Binary quality test                                                       |
| dev util anonymise                                  | Anonymise capture log                                                     |
| dev util image jpeg                                 | Create dummy image files                                                  |
| dev util wait                                       | Wait for specified seconds                                                |
| dropbox file account feature                        | List Dropbox account features                                             |
| dropbox file account filesystem                     | Show Dropbox file system version                                          |
| dropbox file account info                           | Dropbox account info                                                      |
| dropbox file compare account                        | Compare files of two accounts                                             |
| dropbox file compare local                          | Compare local folders and Dropbox folders                                 |
| dropbox file copy                                   | Copy files                                                                |
| dropbox file delete                                 | Delete file or folder                                                     |
| dropbox file export doc                             | Export document                                                           |
| dropbox file export url                             | Export a document from the URL                                            |
| dropbox file import batch url                       | Batch import files from URL                                               |
| dropbox file import url                             | Import file from the URL                                                  |
| dropbox file info                                   | Resolve metadata of the path                                              |
| dropbox file list                                   | List files and folders                                                    |
| dropbox file lock acquire                           | Lock a file                                                               |
| dropbox file lock all release                       | Release all locks under the specified path                                |
| dropbox file lock batch acquire                     | Lock multiple files                                                       |
| dropbox file lock batch release                     | Release multiple locks                                                    |
| dropbox file lock list                              | List locks under the specified path                                       |
| dropbox file lock release                           | Release a lock                                                            |
| dropbox file merge                                  | Merge paths                                                               |
| dropbox file move                                   | Move files                                                                |
| dropbox file replication                            | Replicate file content to the other account                               |
| dropbox file request create                         | Create a file request                                                     |
| dropbox file request delete closed                  | Delete all closed file requests on this account.                          |
| dropbox file request delete url                     | Delete a file request by the file request URL                             |
| dropbox file request list                           | List file requests of the individual account                              |
| dropbox file restore all                            | Restore files under given path                                            |
| dropbox file restore ext                            | Restore files with a specific extension                                   |
| dropbox file revision download                      | Download the file revision                                                |
| dropbox file revision list                          | List file revisions                                                       |
| dropbox file revision restore                       | Restore the file revision                                                 |
| dropbox file search content                         | Search file content                                                       |
| dropbox file search name                            | Search file name                                                          |
| dropbox file share info                             | Retrieve sharing information of the file                                  |
| dropbox file sharedfolder info                      | Get shared folder info                                                    |
| dropbox file sharedfolder leave                     | Leave from the shared folder                                              |
| dropbox file sharedfolder list                      | List shared folder(s)                                                     |
| dropbox file sharedfolder member add                | Add a member to the shared folder                                         |
| dropbox file sharedfolder member delete             | Delete a member from the shared folder                                    |
| dropbox file sharedfolder member list               | List shared folder member(s)                                              |
| dropbox file sharedfolder mount add                 | Add the shared folder to the current user's Dropbox                       |
| dropbox file sharedfolder mount delete              | The current user unmounts the designated folder.                          |
| dropbox file sharedfolder mount list                | List all shared folders the current user mounted                          |
| dropbox file sharedfolder mount mountable           | List all shared folders the current user can mount                        |
| dropbox file sharedfolder share                     | Share a folder                                                            |
| dropbox file sharedfolder unshare                   | Unshare a folder                                                          |
| dropbox file sharedlink create                      | Create shared link                                                        |
| dropbox file sharedlink delete                      | Remove shared links                                                       |
| dropbox file sharedlink file list                   | List files for the shared link                                            |
| dropbox file sharedlink info                        | Get information about the shared link                                     |
| dropbox file sharedlink list                        | List of shared link(s)                                                    |
| dropbox file size                                   | Storage usage                                                             |
| dropbox file sync down                              | Downstream sync with Dropbox                                              |
| dropbox file sync online                            | Sync online files                                                         |
| dropbox file sync up                                | Upstream sync with Dropbox                                                |
| dropbox file tag add                                | Add tag to file or folder                                                 |
| dropbox file tag delete                             | Delete a tag from the file/folder                                         |
| dropbox file tag list                               | List tags of the path                                                     |
| dropbox file template apply                         | Apply file/folder structure template to the Dropbox path                  |
| dropbox file template capture                       | Capture file/folder structure as template from Dropbox path               |
| dropbox file watch                                  | Watch file activities                                                     |
| dropbox paper append                                | Append the content to the end of the existing Paper doc                   |
| dropbox paper create                                | Create new Paper in the path                                              |
| dropbox paper overwrite                             | Overwrite existing Paper document                                         |
| dropbox paper prepend                               | Append the content to the beginning of the existing Paper doc             |
| dropbox sign account info                           | Show Dropbox Sign account information                                     |
| dropbox sign request list                           | List signature requests                                                   |
| dropbox sign request signature list                 | List signatures of requests                                               |
| dropbox team activity batch user                    | Scan activities for multiple users                                        |
| dropbox team activity daily event                   | Report activities by day                                                  |
| dropbox team activity event                         | Event log                                                                 |
| dropbox team activity user                          | Activities log per user                                                   |
| dropbox team admin group role add                   | Add the role to members of the group                                      |
| dropbox team admin group role delete                | Delete the role from all members except of members of the exception group |
| dropbox team admin list                             | List admin roles of members                                               |
| dropbox team admin role add                         | Add a new role to the member                                              |
| dropbox team admin role clear                       | Remove all admin roles from the member                                    |
| dropbox team admin role delete                      | Remove a role from the member                                             |
| dropbox team admin role list                        | List admin roles of the team                                              |
| dropbox team backup device status                   | Dropbox Backup device status change in the specified period               |
| dropbox team content legacypaper count              | Count number of Paper documents per member                                |
| dropbox team content legacypaper export             | Export entire team member Paper documents into local path                 |
| dropbox team content legacypaper list               | List team member Paper documents                                          |
| dropbox team content member list                    | List team folder & shared folder members                                  |
| dropbox team content member size                    | Count number of members of team folders and shared folders                |
| dropbox team content mount list                     | List all mounted/unmounted shared folders of team members.                |
| dropbox team content policy list                    | List policies of team folders and shared folders in the team              |
| dropbox team device list                            | List all devices/sessions in the team                                     |
| dropbox team device unlink                          | Unlink device sessions                                                    |
| dropbox team feature                                | Team feature                                                              |
| dropbox team filerequest clone                      | Clone file requests by given data                                         |
| dropbox team filerequest list                       | List all file requests in the team                                        |
| dropbox team filesystem                             | Identify team's file system version                                       |
| dropbox team group add                              | Create new group                                                          |
| dropbox team group batch add                        | Bulk adding groups                                                        |
| dropbox team group batch delete                     | Delete groups                                                             |
| dropbox team group clear externalid                 | Clear an external ID of a group                                           |
| dropbox team group delete                           | Delete group                                                              |
| dropbox team group folder list                      | List folders of each group                                                |
| dropbox team group list                             | List group(s)                                                             |
| dropbox team group member add                       | Add a member to the group                                                 |
| dropbox team group member batch add                 | Bulk add members into groups                                              |
| dropbox team group member batch delete              | Delete members from groups                                                |
| dropbox team group member batch update              | Add or delete members from groups                                         |
| dropbox team group member delete                    | Delete a member from the group                                            |
| dropbox team group member list                      | List members of groups                                                    |
| dropbox team group rename                           | Rename the group                                                          |
| dropbox team group update type                      | Update group management type                                              |
| dropbox team info                                   | Team information                                                          |
| dropbox team insight report teamfoldermember        | Report team folder members                                                |
| dropbox team insight scan                           | Scans team data for analysis                                              |
| dropbox team insight scanretry                      | Retry scan for errors on the last scan                                    |
| dropbox team insight summarize                      | Summarize team data for analysis                                          |
| dropbox team legalhold add                          | Creates new legal hold policy.                                            |
| dropbox team legalhold list                         | Retrieve existing policies                                                |
| dropbox team legalhold member batch update          | Update member list of legal hold policy                                   |
| dropbox team legalhold member list                  | List members of the legal hold                                            |
| dropbox team legalhold release                      | Releases a legal hold by Id                                               |
| dropbox team legalhold revision list                | List revisions under legal hold                                           |
| dropbox team legalhold update desc                  | Update description of the legal hold policy                               |
| dropbox team legalhold update name                  | Update name of the legal hold policy                                      |
| dropbox team linkedapp list                         | List linked applications                                                  |
| dropbox team member batch delete                    | Delete members                                                            |
| dropbox team member batch detach                    | Convert Dropbox for teams accounts to a Basic account                     |
| dropbox team member batch invite                    | Invite member(s)                                                          |
| dropbox team member batch reinvite                  | Reinvite invited status members to the team                               |
| dropbox team member batch suspend                   | Bulk suspend members                                                      |
| dropbox team member batch unsuspend                 | Bulk unsuspend members                                                    |
| dropbox team member clear externalid                | Clear external_id of members                                              |
| dropbox team member feature                         | List member feature settings                                              |
| dropbox team member file lock all release           | Release all locks under the path of the member                            |
| dropbox team member file lock list                  | List locks of the member under the path                                   |
| dropbox team member file lock release               | Release the lock of the path as the member                                |
| dropbox team member file permdelete                 | Permanently delete the file or folder at a given path of the team member. |
| dropbox team member folder list                     | List folders for each member                                              |
| dropbox team member folder replication              | Replicate a folder to another member's personal folder                    |
| dropbox team member list                            | List team member(s)                                                       |
| dropbox team member quota batch update              | Update team member quota                                                  |
| dropbox team member quota list                      | List team member quota                                                    |
| dropbox team member quota usage                     | List team member storage usage                                            |
| dropbox team member replication                     | Replicate team member files                                               |
| dropbox team member suspend                         | Suspend a member                                                          |
| dropbox team member unsuspend                       | Unsuspend a member                                                        |
| dropbox team member update batch email              | Member email operation                                                    |
| dropbox team member update batch externalid         | Update External ID of team members                                        |
| dropbox team member update batch invisible          | Enable directory restriction to members                                   |
| dropbox team member update batch profile            | Batch update member profiles                                              |
| dropbox team member update batch visible            | Disable directory restriction to members                                  |
| dropbox team namespace file list                    | List all files and folders of the team namespaces                         |
| dropbox team namespace file size                    | List all files and folders of the team namespaces                         |
| dropbox team namespace list                         | List all namespaces of the team                                           |
| dropbox team namespace member list                  | List members of shared folders and team folders in the team               |
| dropbox team namespace summary                      | Report team namespace status summary.                                     |
| dropbox team report activity                        | Activities report                                                         |
| dropbox team report devices                         | Devices report                                                            |
| dropbox team report membership                      | Membership report                                                         |
| dropbox team report storage                         | Storage report                                                            |
| dropbox team runas file batch copy                  | Batch copy files/folders as a member                                      |
| dropbox team runas file list                        | List files and folders run as a member                                    |
| dropbox team runas file sync batch up               | Batch upstream sync with Dropbox                                          |
| dropbox team runas sharedfolder batch leave         | Leave shared folders in batch                                             |
| dropbox team runas sharedfolder batch share         | Share shared folders in batch                                             |
| dropbox team runas sharedfolder batch unshare       | Unshare shared folders in batch                                           |
| dropbox team runas sharedfolder isolate             | Isolate member from shared folder                                         |
| dropbox team runas sharedfolder list                | List shared folders                                                       |
| dropbox team runas sharedfolder member batch add    | Add members to shared folders in batch                                    |
| dropbox team runas sharedfolder member batch delete | Remove members from shared folders in batch                               |
| dropbox team runas sharedfolder mount add           | Mount a shared folder as another member                                   |
| dropbox team runas sharedfolder mount delete        | The specified user unmounts the designated folder.                        |
| dropbox team runas sharedfolder mount list          | List all shared folders the specified member mounted                      |
| dropbox team runas sharedfolder mount mountable     | List all shared folders the member can mount                              |
| dropbox team sharedlink cap expiry                  | Set expiry cap to shared links in the team                                |
| dropbox team sharedlink cap visibility              | Set visibility cap to shared links in the team                            |
| dropbox team sharedlink delete links                | Batch delete shared links                                                 |
| dropbox team sharedlink delete member               | Delete all shared links of the member                                     |
| dropbox team sharedlink list                        | List of shared links                                                      |
| dropbox team sharedlink update expiry               | Update expiration date of public shared links within the team             |
| dropbox team sharedlink update password             | Set or update shared link passwords                                       |
| dropbox team sharedlink update visibility           | Update visibility of shared links                                         |
| dropbox team teamfolder add                         | Add team folder to the team                                               |
| dropbox team teamfolder archive                     | Archive team folder                                                       |
| dropbox team teamfolder batch archive               | Archiving team folders                                                    |
| dropbox team teamfolder batch permdelete            | Permanently delete team folders                                           |
| dropbox team teamfolder batch replication           | Batch replication of team folders                                         |
| dropbox team teamfolder file list                   | List files in team folders                                                |
| dropbox team teamfolder file lock all release       | Release all locks under the path of the team folder                       |
| dropbox team teamfolder file lock list              | List locks in the team folder                                             |
| dropbox team teamfolder file lock release           | Release lock of the path in the team folder                               |
| dropbox team teamfolder file size                   | Calculate size of team folders                                            |
| dropbox team teamfolder list                        | List team folder(s)                                                       |
| dropbox team teamfolder member add                  | Batch adding users/groups to team folders                                 |
| dropbox team teamfolder member delete               | Batch removing users/groups from team folders                             |
| dropbox team teamfolder member list                 | List team folder members                                                  |
| dropbox team teamfolder partial replication         | Partial team folder replication to the other team                         |
| dropbox team teamfolder permdelete                  | Permanently delete team folder                                            |
| dropbox team teamfolder policy list                 | List policies of team folders                                             |
| dropbox team teamfolder replication                 | Replicate a team folder to the other team                                 |
| dropbox team teamfolder sync setting list           | List team folder sync settings                                            |
| dropbox team teamfolder sync setting update         | Batch update team folder sync settings                                    |
| figma account info                                  | Retrieve current user information                                         |
| figma file export all page                          | Export all files/pages under the team                                     |
| figma file export frame                             | Export all frames of the Figma file                                       |
| figma file export node                              | Export Figma document Node                                                |
| figma file export page                              | Export all pages of the Figma file                                        |
| figma file info                                     | Show information of the figma file                                        |
| figma file list                                     | List files in the Figma Project                                           |
| figma project list                                  | List projects of the team                                                 |
| github content get                                  | Get content metadata of the repository                                    |
| github content put                                  | Put small text content into the repository                                |
| github issue list                                   | List issues of the public/private GitHub repository                       |
| github profile                                      | Get the authenticated user                                                |
| github release asset download                       | Download assets                                                           |
| github release asset list                           | List assets of GitHub Release                                             |
| github release asset upload                         | Upload assets file into the GitHub Release                                |
| github release draft                                | Create release draft                                                      |
| github release list                                 | List releases                                                             |
| github tag create                                   | Create a tag on the repository                                            |
| license                                             | Show license information                                                  |
| local file template apply                           | Apply file/folder structure template to the local path                    |
| local file template capture                         | Capture file/folder structure as template from local path                 |
| log api job                                         | Show statistics of the API log of the job specified by the job ID         |
| log api name                                        | Show statistics of the API log of the job specified by the job name       |
| log cat curl                                        | Format capture logs as `curl` sample                                      |
| log cat job                                         | Retrieve logs of specified Job ID                                         |
| log cat kind                                        | Concatenate and print logs of specified log kind                          |
| log cat last                                        | Print the last job log files                                              |
| log job archive                                     | Archive jobs                                                              |
| log job delete                                      | Delete old job history                                                    |
| log job list                                        | Show job history                                                          |
| slack conversation history                          | Conversation history                                                      |
| slack conversation list                             | List channels                                                             |
| util archive unzip                                  | Extract the zip archive file                                              |
| util archive zip                                    | Compress target files into the zip archive                                |
| util cert selfsigned                                | Generate self-signed certificate and key                                  |
| util database exec                                  | Execute query on SQLite3 database file                                    |
| util database query                                 | Query SQLite3 database                                                    |
| util date today                                     | Display current date                                                      |
| util datetime now                                   | Display current date/time                                                 |
| util decode base32                                  | Decode text from Base32 (RFC 4648) format                                 |
| util decode base64                                  | Decode text from Base64 (RFC 4648) format                                 |
| util desktop open                                   | Open a file or folder with the default application                        |
| util encode base32                                  | Encode text into Base32 (RFC 4648) format                                 |
| util encode base64                                  | Encode text into Base64 (RFC 4648) format                                 |
| util feed json                                      | Load feed from the URL and output the content as JSON                     |
| util file hash                                      | File Hash                                                                 |
| util git clone                                      | Clone git repository                                                      |
| util image exif                                     | Print EXIF metadata of image file                                         |
| util image placeholder                              | Create placeholder image                                                  |
| util json query                                     | Query JSON data                                                           |
| util net download                                   | Download a file                                                           |
| util qrcode create                                  | Create a QR code image file                                               |
| util qrcode wifi                                    | Generate QR code for WIFI configuration                                   |
| util release install                                | Download & install watermint toolbox to the path                          |
| util table format xlsx                              | Formatting xlsx file into text                                            |
| util text case down                                 | Print lower case text                                                     |
| util text case up                                   | Print upper case text                                                     |
| util text encoding from                             | Convert text encoding to UTF-8 text file from specified encoding.         |
| util text encoding to                               | Convert text encoding to specified encoding from UTF-8 text file.         |
| util text nlp english entity                        | Split English text into entities                                          |
| util text nlp english sentence                      | Split English text into sentences                                         |
| util text nlp english token                         | Split English text into tokens                                            |
| util text nlp japanese token                        | Tokenize Japanese text                                                    |
| util text nlp japanese wakati                       | Wakati gaki (tokenize Japanese text)                                      |
| util tidy move dispatch                             | Dispatch files                                                            |
| util tidy move simple                               | Archive local files                                                       |
| util tidy pack remote                               | Package remote folder into the zip file                                   |
| util time now                                       | Display current time                                                      |
| util unixtime format                                | Time format to convert the unix time (epoch seconds from 1970-01-01)      |
| util unixtime now                                   | Display current time in unixtime                                          |
| util uuid timestamp                                 | UUID Timestamp                                                            |
| util uuid ulid                                      | ULID Utility                                                              |
| util uuid v4                                        | Generate UUID v4 (random UUID)                                            |
| util uuid v7                                        | Generate UUID v7                                                          |
| util uuid version                                   | Parse version and variant of UUID                                         |
| util xlsx create                                    | Create an empty spreadsheet                                               |
| util xlsx sheet export                              | Export data from the xlsx file                                            |
| util xlsx sheet import                              | Import data into xlsx file                                                |
| util xlsx sheet list                                | List sheets of the xlsx file                                              |
| version                                             | Show version                                                              |



