---
layout: release
title: Changes of Release 141
lang: en
---

# Changes between `Release 141` to `Release 142`

# Commands added


| Command                                             | Title                                                                                                                                 |
|-----------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| asana team list                                     | List team                                                                                                                             |
| asana team project list                             | List projects of the team                                                                                                             |
| asana team task list                                | List tasks of the team                                                                                                                |
| asana workspace list                                | List workspaces                                                                                                                       |
| asana workspace project list                        | List projects of the workspace                                                                                                        |
| config auth delete                                  | Delete existing auth credential                                                                                                       |
| config auth list                                    | List all auth credentials                                                                                                             |
| config feature disable                              | Disable a feature.                                                                                                                    |
| config feature enable                               | Enable a feature.                                                                                                                     |
| config feature list                                 | List available optional features.                                                                                                     |
| config license install                              | Install a license key                                                                                                                 |
| config license list                                 | List available license keys                                                                                                           |
| deepl translate text                                | Translate text                                                                                                                        |
| dev benchmark local                                 | Create dummy folder structure in local file system.                                                                                   |
| dev benchmark upload                                | Upload benchmark                                                                                                                      |
| dev benchmark uploadlink                            | Benchmark single file upload with upload temporary link API.                                                                          |
| dev build catalogue                                 | Generate catalogue                                                                                                                    |
| dev build doc                                       | Document generator                                                                                                                    |
| dev build info                                      | Generate build information file                                                                                                       |
| dev build license                                   | Generate LICENSE.txt                                                                                                                  |
| dev build package                                   | Package a build                                                                                                                       |
| dev build preflight                                 | Process prerequisites for the release                                                                                                 |
| dev build readme                                    | Generate README.txt                                                                                                                   |
| dev ci artifact up                                  | Upload CI artifact                                                                                                                    |
| dev ci auth export                                  | Export deploy token data for CI build                                                                                                 |
| dev diag endpoint                                   | List endpoints                                                                                                                        |
| dev diag throughput                                 | Evaluate throughput from capture logs                                                                                                 |
| dev doc knowledge                                   | Generate reduced knowledge base                                                                                                       |
| dev doc markdown                                    | Generate messages from markdown source                                                                                                |
| dev doc msg add                                     | Add a new message                                                                                                                     |
| dev doc msg catalogue_options                       | Generate option descriptions for all recipes in catalogue                                                                             |
| dev doc msg delete                                  | Delete a message                                                                                                                      |
| dev doc msg list                                    | List messages                                                                                                                         |
| dev doc msg options                                 | Generate option descriptions for SelectString fields                                                                                  |
| dev doc msg translate                               | Translation helper                                                                                                                    |
| dev doc msg update                                  | Update a message                                                                                                                      |
| dev doc msg verify                                  | Verify message template variables consistency                                                                                         |
| dev doc review approve                              | Mark a message as reviewed                                                                                                            |
| dev doc review batch                                | Review and approve messages in batch                                                                                                  |
| dev doc review list                                 | List unreviewed messages                                                                                                              |
| dev doc review options                              | Review missing SelectString option descriptions                                                                                       |
| dev info                                            | Dev information                                                                                                                       |
| dev kvs concurrency                                 | Concurrency test for KVS engine                                                                                                       |
| dev kvs dump                                        | Dump KVS data                                                                                                                         |
| dev license issue                                   | Issue a license                                                                                                                       |
| dev lifecycle assets                                | Remove deprecated assets                                                                                                              |
| dev lifecycle planchangepath                        | Add plan of changing path to commands                                                                                                 |
| dev lifecycle planprune                             | Add plan of the command discontinuation                                                                                               |
| dev module list                                     | Dependent module list                                                                                                                 |
| dev placeholder pathchange                          | Placeholder command for path change document generation                                                                               |
| dev placeholder prune                               | Placeholder of prune workflow messages                                                                                                |
| dev release announcement                            | Update announcements                                                                                                                  |
| dev release asset                                   | Commit a file to a repository                                                                                                         |
| dev release asseturl                                | Update asset URL of the release                                                                                                       |
| dev release candidate                               | Validate release candidate                                                                                                            |
| dev release checkin                                 | Check in the new release                                                                                                              |
| dev release doc                                     | Generate release documents                                                                                                            |
| dev release publish                                 | Publish release                                                                                                                       |
| dev replay approve                                  | Approve the replay as test bundle                                                                                                     |
| dev replay bundle                                   | Run all replays                                                                                                                       |
| dev replay recipe                                   | Replay recipe                                                                                                                         |
| dev replay remote                                   | Run remote replay bundle                                                                                                              |
| dev spec diff                                       | Compare spec of two releases                                                                                                          |
| dev spec doc                                        | Generate spec docs                                                                                                                    |
| dev test coverage list                              | Test Coverage List                                                                                                                    |
| dev test coverage missing                           | Find Missing Tests                                                                                                                    |
| dev test coverage pkg                               | Test Coverage Package                                                                                                                 |
| dev test coverage summary                           | Test Coverage Summary                                                                                                                 |
| dev test echo                                       | Echo text                                                                                                                             |
| dev test license                                    | Testing license required logic                                                                                                        |
| dev test panic                                      | Panic test                                                                                                                            |
| dev test recipe                                     | Test recipe                                                                                                                           |
| dev test resources                                  | Binary quality test                                                                                                                   |
| dev util anonymise                                  | Anonymise capture log                                                                                                                 |
| dev util image jpeg                                 | Create dummy image files                                                                                                              |
| dev util wait                                       | Wait for specified seconds                                                                                                            |
| dropbox file account feature                        | List Dropbox account features                                                                                                         |
| dropbox file account filesystem                     | Show Dropbox file system version                                                                                                      |
| dropbox file account info                           | Dropbox account info                                                                                                                  |
| dropbox file compare account                        | Compare files of two accounts                                                                                                         |
| dropbox file compare local                          | Compare local folders and Dropbox folders                                                                                             |
| dropbox file copy                                   | Copy files                                                                                                                            |
| dropbox file delete                                 | Delete file or folder                                                                                                                 |
| dropbox file export doc                             | Export document                                                                                                                       |
| dropbox file export url                             | Export a document from the URL                                                                                                        |
| dropbox file import batch url                       | Batch import files from URL                                                                                                           |
| dropbox file import url                             | Import file from the URL                                                                                                              |
| dropbox file info                                   | Resolve metadata of the path                                                                                                          |
| dropbox file list                                   | List files and folders                                                                                                                |
| dropbox file lock acquire                           | Lock a file                                                                                                                           |
| dropbox file lock all release                       | Release all locks under the specified path                                                                                            |
| dropbox file lock batch acquire                     | Lock multiple files                                                                                                                   |
| dropbox file lock batch release                     | Release multiple locks                                                                                                                |
| dropbox file lock list                              | List locks under the specified path                                                                                                   |
| dropbox file lock release                           | Release a lock                                                                                                                        |
| dropbox file merge                                  | Merge paths                                                                                                                           |
| dropbox file move                                   | Move files                                                                                                                            |
| dropbox file replication                            | Replicate file content to the other account                                                                                           |
| dropbox file request create                         | Create a file request                                                                                                                 |
| dropbox file request delete closed                  | Delete all closed file requests on this account.                                                                                      |
| dropbox file request delete url                     | Delete a file request by the file request URL                                                                                         |
| dropbox file request list                           | List file requests of the individual account                                                                                          |
| dropbox file restore all                            | Restore files under given path                                                                                                        |
| dropbox file restore ext                            | Restore files with a specific extension                                                                                               |
| dropbox file revision download                      | Download the file revision                                                                                                            |
| dropbox file revision list                          | List file revisions                                                                                                                   |
| dropbox file revision restore                       | Restore the file revision                                                                                                             |
| dropbox file search content                         | Search file content                                                                                                                   |
| dropbox file search name                            | Search file name                                                                                                                      |
| dropbox file share info                             | Retrieve sharing information of the file                                                                                              |
| dropbox file sharedfolder info                      | Get shared folder info                                                                                                                |
| dropbox file sharedfolder leave                     | Leave the shared folder                                                                                                               |
| dropbox file sharedfolder list                      | List shared folders                                                                                                                   |
| dropbox file sharedfolder member add                | Add a member to the shared folder                                                                                                     |
| dropbox file sharedfolder member delete             | Remove a member from the shared folder                                                                                                |
| dropbox file sharedfolder member list               | List shared folder members                                                                                                            |
| dropbox file sharedfolder mount add                 | Add the shared folder to the current user's Dropbox                                                                                   |
| dropbox file sharedfolder mount delete              | Unmount the shared folder                                                                                                             |
| dropbox file sharedfolder mount list                | List all shared folders the current user has mounted                                                                                  |
| dropbox file sharedfolder mount mountable           | List all shared folders the current user can mount                                                                                    |
| dropbox file sharedfolder share                     | Share a folder                                                                                                                        |
| dropbox file sharedfolder unshare                   | Unshare a folder                                                                                                                      |
| dropbox file sharedlink create                      | Create shared link                                                                                                                    |
| dropbox file sharedlink delete                      | Remove shared links                                                                                                                   |
| dropbox file sharedlink file list                   | List files for the shared link                                                                                                        |
| dropbox file sharedlink info                        | Get information about the shared link                                                                                                 |
| dropbox file sharedlink list                        | List shared links                                                                                                                     |
| dropbox file size                                   | Storage usage                                                                                                                         |
| dropbox file sync down                              | Downstream sync with Dropbox                                                                                                          |
| dropbox file sync online                            | Sync online files                                                                                                                     |
| dropbox file sync up                                | Upstream sync with Dropbox                                                                                                            |
| dropbox file tag add                                | Add tag to file or folder                                                                                                             |
| dropbox file tag delete                             | Delete a tag from the file/folder                                                                                                     |
| dropbox file tag list                               | List tags of the path                                                                                                                 |
| dropbox file template apply                         | Apply file/folder structure template to the Dropbox path                                                                              |
| dropbox file template capture                       | Capture file/folder structure as template from Dropbox path                                                                           |
| dropbox file watch                                  | Watch file activities                                                                                                                 |
| dropbox paper append                                | Append the content to the end of the existing Paper doc                                                                               |
| dropbox paper create                                | Create new Paper in the path                                                                                                          |
| dropbox paper overwrite                             | Overwrite an existing Paper document                                                                                                  |
| dropbox paper prepend                               | Append the content to the beginning of the existing Paper doc                                                                         |
| dropbox sign account info                           | Show Dropbox Sign account information                                                                                                 |
| dropbox sign request list                           | List signature requests                                                                                                               |
| dropbox sign request signature list                 | List signatures of requests                                                                                                           |
| dropbox team activity batch user                    | Scan and retrieve activity logs for multiple team members in batch, useful for compliance auditing and user behavior analysis         |
| dropbox team activity daily event                   | Generate daily activity reports showing team events grouped by date, helpful for tracking team usage patterns and security monitoring |
| dropbox team activity event                         | Retrieve detailed team activity event logs with filtering options, essential for security auditing and compliance monitoring          |
| dropbox team activity user                          | Retrieve activity logs for specific team members, showing their file operations, logins, and sharing activities                       |
| dropbox team admin group role add                   | Assign admin roles to all members of a specified group, streamlining role management for large teams                                  |
| dropbox team admin group role delete                | Remove admin roles from all team members except those in a specified exception group, useful for role cleanup and access control      |
| dropbox team admin list                             | Display all team members with their assigned admin roles, helpful for auditing administrative access and permissions                  |
| dropbox team admin role add                         | Grant a specific admin role to an individual team member, enabling granular permission management                                     |
| dropbox team admin role clear                       | Revoke all administrative privileges from a team member, useful for role transitions or security purposes                             |
| dropbox team admin role delete                      | Remove a specific admin role from a team member while preserving other roles, allowing precise permission adjustments                 |
| dropbox team admin role list                        | Display all available admin roles in the team with their descriptions and permissions                                                 |
| dropbox team backup device status                   | Track Dropbox Backup status changes for all team devices over a specified period, monitoring backup health and compliance             |
| dropbox team content legacypaper count              | Calculate the total number of legacy Paper documents owned by each team member, useful for content auditing and migration planning    |
| dropbox team content legacypaper export             | Export all legacy Paper documents from team members to local storage in HTML or Markdown format for backup or migration               |
| dropbox team content legacypaper list               | Generate a comprehensive list of all legacy Paper documents across the team with ownership and metadata information                   |
| dropbox team content member list                    | Display all members with access to team folders and shared folders, showing permission levels and folder relationships                |
| dropbox team content member size                    | Calculate member counts for each team folder and shared folder, helping identify heavily accessed content and optimize permissions    |
| dropbox team content mount list                     | Display mount status of all shared folders for team members, identifying which folders are actively synced to member devices          |
| dropbox team content policy list                    | Review all access policies and restrictions applied to team folders and shared folders for governance compliance                      |
| dropbox team device list                            | Display all devices and active sessions connected to team member accounts with device details and last activity timestamps            |
| dropbox team device unlink                          | Remotely disconnect devices from team member accounts, essential for securing lost/stolen devices or revoking access                  |
| dropbox team feature                                | Display all features and capabilities enabled for your Dropbox team account, including API limits and special features                |
| dropbox team filerequest clone                      | Duplicate existing file requests with customized settings, useful for creating similar requests across team members                   |
| dropbox team filerequest list                       | Display all active and closed file requests created by team members, helping track external file collection activities                |
| dropbox team filesystem                             | Identify whether your team uses legacy or modern file system architecture, important for feature compatibility                        |
| dropbox team group add                              | Create a new group in your team for organizing members and managing permissions collectively                                          |
| dropbox team group batch add                        | Create multiple groups at once using batch processing, efficient for large-scale team organization                                    |
| dropbox team group batch delete                     | Remove multiple groups from your team in batch, streamlining group cleanup and reorganization                                         |
| dropbox team group clear externalid                 | Remove external ID mappings from groups, useful when disconnecting from external identity providers                                   |
| dropbox team group delete                           | Remove a specific group from your team, automatically removing all member associations                                                |
| dropbox team group folder list                      | Display all folders accessible by each group, showing group-based content organization and permissions                                |
| dropbox team group list                             | Display all groups in your team with member counts and group management types                                                         |
| dropbox team group member add                       | Add individual team members to a specific group for centralized permission management                                                 |
| dropbox team group member batch add                 | Add multiple members to groups efficiently using batch processing, ideal for large team reorganizations                               |
| dropbox team group member batch delete              | Remove multiple members from groups in batch, streamlining group membership management                                                |
| dropbox team group member batch update              | Update group memberships in bulk by adding or removing members, optimizing group composition changes                                  |
| dropbox team group member delete                    | Remove a specific member from a group while preserving their other group memberships                                                  |
| dropbox team group member list                      | Display all members belonging to each group, useful for auditing group compositions and access rights                                 |
| dropbox team group rename                           | Change the name of an existing group to better reflect its purpose or organizational changes                                          |
| dropbox team group update type                      | Change how a group is managed (user-managed vs company-managed), affecting who can modify group membership                            |
| dropbox team info                                   | Display essential team account information including team ID and basic team settings                                                  |
| dropbox team insight report teamfoldermember        | Generate detailed reports on team folder membership, showing access patterns and member distribution                                  |
| dropbox team insight scan                           | Perform comprehensive data scanning across your team for analytics and insights generation                                            |
| dropbox team insight scanretry                      | Re-run failed or incomplete scans to ensure complete data collection for team insights                                                |
| dropbox team insight summarize                      | Generate summary reports from scanned team data, providing actionable insights on team usage and patterns                             |
| dropbox team legalhold add                          | Create a legal hold policy to preserve specified team content for compliance or litigation purposes                                   |
| dropbox team legalhold list                         | Display all active legal hold policies with their details, members, and preservation status                                           |
| dropbox team legalhold member batch update          | Add or remove multiple team members from legal hold policies in batch for efficient compliance management                             |
| dropbox team legalhold member list                  | Display all team members currently under legal hold policies with their preservation status                                           |
| dropbox team legalhold release                      | Release a legal hold policy and restore normal file operations for affected members and content                                       |
| dropbox team legalhold revision list                | Display all file revisions preserved under legal hold policies, ensuring comprehensive data retention                                 |
| dropbox team legalhold update desc                  | Modify the description of an existing legal hold policy to reflect changes in scope or purpose                                        |
| dropbox team legalhold update name                  | Change the name of a legal hold policy for better identification and organization                                                     |
| dropbox team linkedapp list                         | Display all third-party applications linked to team member accounts for security auditing and access control                          |
| dropbox team member batch delete                    | Remove multiple team members in batch, efficiently managing team departures and access revocation                                     |
| dropbox team member batch detach                    | Convert multiple team accounts to individual Basic accounts, preserving personal data while removing team access                      |
| dropbox team member batch invite                    | Send batch invitations to new team members, streamlining the onboarding process for multiple users                                    |
| dropbox team member batch reinvite                  | Resend invitations to pending members who haven't joined yet, ensuring all intended members receive access                            |
| dropbox team member batch suspend                   | Temporarily suspend multiple team members' access while preserving their data and settings                                            |
| dropbox team member batch unsuspend                 | Restore access for multiple suspended team members, reactivating their accounts in batch                                              |
| dropbox team member clear externalid                | Remove external ID mappings from team members, useful when disconnecting from identity management systems                             |
| dropbox team member feature                         | Display feature settings and capabilities enabled for specific team members, helping understand member permissions                    |
| dropbox team member file lock all release           | Release all file locks held by a team member under a specified path, resolving editing conflicts                                      |
| dropbox team member file lock list                  | Display all files locked by a specific team member under a given path, identifying potential collaboration blocks                     |
| dropbox team member file lock release               | Release a specific file lock held by a team member, enabling others to edit the file                                                  |
| dropbox team member file permdelete                 | Permanently delete files or folders from a team member's account, bypassing trash for immediate removal                               |
| dropbox team member folder list                     | Display all folders in each team member's account, useful for content auditing and storage analysis                                   |
| dropbox team member folder replication              | Copy folder contents from one team member to another's personal space, facilitating content transfer and backup                       |
| dropbox team member list                            | Display comprehensive list of all team members with their status, roles, and account details                                          |
| dropbox team member quota batch update              | Modify storage quotas for multiple team members in batch, managing storage allocation efficiently                                     |
| dropbox team member quota list                      | Display storage quota assignments for all team members, helping monitor and plan storage distribution                                 |
| dropbox team member quota usage                     | Show actual storage usage for each team member compared to their quotas, identifying storage needs                                    |
| dropbox team member replication                     | Replicate all files from one team member's account to another, useful for account transitions or backups                              |
| dropbox team member suspend                         | Temporarily suspend a team member's access to their account while preserving all data and settings                                    |
| dropbox team member unsuspend                       | Restore access for a suspended team member, reactivating their account and all associated permissions                                 |
| dropbox team member update batch email              | Update email addresses for multiple team members in batch, managing email changes efficiently                                         |
| dropbox team member update batch externalid         | Set or update external IDs for multiple team members, integrating with identity management systems                                    |
| dropbox team member update batch invisible          | Hide team members from the directory listing, enhancing privacy for sensitive roles or contractors                                    |
| dropbox team member update batch profile            | Update profile information for multiple team members including names and job titles in batch                                          |
| dropbox team member update batch visible            | Make hidden team members visible in the directory, restoring standard visibility settings                                             |
| dropbox team namespace file list                    | Display comprehensive file and folder listings within team namespaces for content inventory and analysis                              |
| dropbox team namespace file size                    | Calculate storage usage for files and folders in team namespaces, providing detailed size analytics                                   |
| dropbox team namespace list                         | Display all team namespaces including team folders and shared spaces with their configurations                                        |
| dropbox team namespace member list                  | Show all members with access to each namespace, detailing permissions and access levels                                               |
| dropbox team namespace summary                      | Generate comprehensive summary reports of team namespace usage, member counts, and storage statistics                                 |
| dropbox team report activity                        | Generate detailed activity reports covering all team operations, useful for compliance and usage analysis                             |
| dropbox team report devices                         | Create comprehensive device usage reports showing all connected devices, platforms, and access patterns                               |
| dropbox team report membership                      | Generate team membership reports including member status, roles, and account statistics over time                                     |
| dropbox team report storage                         | Create detailed storage usage reports showing team consumption, trends, and member distribution                                       |
| dropbox team runas file batch copy                  | Copy multiple files or folders on behalf of team members, useful for content management and organization                              |
| dropbox team runas file list                        | List files and folders in a team member's account by running operations as that member                                                |
| dropbox team runas file sync batch up               | Upload multiple local files to team members' Dropbox accounts in batch, running as those members                                      |
| dropbox team runas sharedfolder batch leave         | Remove team members from multiple shared folders in batch by running leave operations as those members                                |
| dropbox team runas sharedfolder batch share         | Share multiple folders on behalf of team members in batch, automating folder sharing processes                                        |
| dropbox team runas sharedfolder batch unshare       | Remove sharing from multiple folders on behalf of team members, managing folder access in bulk                                        |
| dropbox team runas sharedfolder isolate             | Remove all shared folder access for a team member and transfer ownership, useful for departing employees                              |
| dropbox team runas sharedfolder list                | Display all shared folders accessible by a team member, running the operation as that member                                          |
| dropbox team runas sharedfolder member batch add    | Add multiple members to shared folders in batch on behalf of folder owners, streamlining access management                            |
| dropbox team runas sharedfolder member batch delete | Remove multiple members from shared folders in batch on behalf of folder owners, managing access efficiently                          |
| dropbox team runas sharedfolder mount add           | Mount shared folders to team members' accounts on their behalf, ensuring proper folder synchronization                                |
| dropbox team runas sharedfolder mount delete        | Unmount shared folders from team members' accounts on their behalf, managing folder synchronization                                   |
| dropbox team runas sharedfolder mount list          | Display all shared folders currently mounted (synced) to a specific team member's account                                             |
| dropbox team runas sharedfolder mount mountable     | Show all available shared folders that a team member can mount but hasn't synced yet                                                  |
| dropbox team sharedlink cap expiry                  | Apply expiration date limits to all team shared links for enhanced security and compliance                                            |
| dropbox team sharedlink cap visibility              | Enforce visibility restrictions on team shared links, controlling public access levels                                                |
| dropbox team sharedlink delete links                | Delete multiple shared links in batch for security compliance or access control cleanup                                               |
| dropbox team sharedlink delete member               | Remove all shared links created by a specific team member, useful for departing employees                                             |
| dropbox team sharedlink list                        | Display comprehensive list of all shared links created by team members with visibility and expiration details                         |
| dropbox team sharedlink update expiry               | Modify expiration dates for existing shared links across the team to enforce security policies                                        |
| dropbox team sharedlink update password             | Add or change passwords on team shared links in batch for enhanced security protection                                                |
| dropbox team sharedlink update visibility           | Change access levels of existing shared links between public, team-only, and password-protected                                       |
| dropbox team teamfolder add                         | Create a new team folder for centralized team content storage and collaboration                                                       |
| dropbox team teamfolder archive                     | Archive a team folder to make it read-only while preserving all content and access history                                            |
| dropbox team teamfolder batch archive               | Archive multiple team folders in batch, efficiently managing folder lifecycle and compliance                                          |
| dropbox team teamfolder batch permdelete            | Permanently delete multiple archived team folders in batch, freeing storage space                                                     |
| dropbox team teamfolder batch replication           | Replicate multiple team folders to another team account in batch for migration or backup                                              |
| dropbox team teamfolder file list                   | Display all files and subfolders within team folders for content inventory and management                                             |
| dropbox team teamfolder file lock all release       | Release all file locks within a team folder path, resolving editing conflicts in bulk                                                 |
| dropbox team teamfolder file lock list              | Display all locked files within team folders, identifying collaboration bottlenecks                                                   |
| dropbox team teamfolder file lock release           | Release specific file locks in team folders to enable collaborative editing                                                           |
| dropbox team teamfolder file size                   | Calculate storage usage for team folders, providing detailed size analytics for capacity planning                                     |
| dropbox team teamfolder list                        | Display all team folders with their status, sync settings, and member access information                                              |
| dropbox team teamfolder member add                  | Add multiple users or groups to team folders in batch, streamlining access provisioning                                               |
| dropbox team teamfolder member delete               | Remove multiple users or groups from team folders in batch, managing access revocation efficiently                                    |
| dropbox team teamfolder member list                 | Display all members with access to each team folder, showing permission levels and access types                                       |
| dropbox team teamfolder partial replication         | Selectively replicate team folder contents to another team, enabling flexible content migration                                       |
| dropbox team teamfolder permdelete                  | Permanently delete an archived team folder and all its contents, irreversibly freeing storage                                         |
| dropbox team teamfolder policy list                 | Display all access policies and restrictions applied to team folders for governance review                                            |
| dropbox team teamfolder replication                 | Copy an entire team folder with all contents to another team account for migration or backup                                          |
| dropbox team teamfolder sync setting list           | Display sync configuration for all team folders, showing default sync behavior for members                                            |
| dropbox team teamfolder sync setting update         | Modify sync settings for multiple team folders in batch, controlling automatic synchronization behavior                               |
| figma account info                                  | Retrieve current user information                                                                                                     |
| figma file export all page                          | Export all files/pages under the team                                                                                                 |
| figma file export frame                             | Export all frames of the Figma file                                                                                                   |
| figma file export node                              | Export Figma document Node                                                                                                            |
| figma file export page                              | Export all pages of the Figma file                                                                                                    |
| figma file info                                     | Show information of the Figma file                                                                                                    |
| figma file list                                     | List files in the Figma Project                                                                                                       |
| figma project list                                  | List projects of the team                                                                                                             |
| github content get                                  | Get content metadata of the repository                                                                                                |
| github content put                                  | Put small text content into the repository                                                                                            |
| github issue list                                   | List issues of the public/private GitHub repository                                                                                   |
| github profile                                      | Get the authenticated user                                                                                                            |
| github release asset download                       | Download assets                                                                                                                       |
| github release asset list                           | List assets of GitHub Release                                                                                                         |
| github release asset upload                         | Upload assets file into the GitHub Release                                                                                            |
| github release draft                                | Create release draft                                                                                                                  |
| github release list                                 | List releases                                                                                                                         |
| github tag create                                   | Create a tag on the repository                                                                                                        |
| license                                             | Show license information                                                                                                              |
| local file template apply                           | Apply file/folder structure template to the local path                                                                                |
| local file template capture                         | Capture file/folder structure as template from local path                                                                             |
| log api job                                         | Show statistics of the API log of the job specified by the job ID                                                                     |
| log api name                                        | Show statistics of the API log of the job specified by the job name                                                                   |
| log cat curl                                        | Format capture logs as `curl` sample                                                                                                  |
| log cat job                                         | Retrieve logs of specified Job ID                                                                                                     |
| log cat kind                                        | Concatenate and print logs of specified log kind                                                                                      |
| log cat last                                        | Print the last job log files                                                                                                          |
| log job archive                                     | Archive jobs                                                                                                                          |
| log job delete                                      | Delete old job history                                                                                                                |
| log job list                                        | Show job history                                                                                                                      |
| slack conversation history                          | Conversation history                                                                                                                  |
| slack conversation list                             | List channels                                                                                                                         |
| util archive unzip                                  | Extract the zip archive file                                                                                                          |
| util archive zip                                    | Compress target files into the zip archive                                                                                            |
| util cert selfsigned                                | Generate self-signed certificate and key                                                                                              |
| util database exec                                  | Execute query on SQLite3 database file                                                                                                |
| util database query                                 | Query SQLite3 database                                                                                                                |
| util date today                                     | Display current date                                                                                                                  |
| util datetime now                                   | Display current date/time                                                                                                             |
| util decode base32                                  | Decode text from Base32 (RFC 4648) format                                                                                             |
| util decode base64                                  | Decode text from Base64 (RFC 4648) format                                                                                             |
| util desktop open                                   | Open a file or folder with the default application                                                                                    |
| util encode base32                                  | Encode text into Base32 (RFC 4648) format                                                                                             |
| util encode base64                                  | Encode text into Base64 (RFC 4648) format                                                                                             |
| util feed json                                      | Load feed from the URL and output the content as JSON                                                                                 |
| util file hash                                      | File Hash                                                                                                                             |
| util git clone                                      | Clone git repository                                                                                                                  |
| util image exif                                     | Print EXIF metadata of image file                                                                                                     |
| util image placeholder                              | Create placeholder image                                                                                                              |
| util json query                                     | Query JSON data                                                                                                                       |
| util net download                                   | Download a file                                                                                                                       |
| util qrcode create                                  | Create a QR code image file                                                                                                           |
| util qrcode wifi                                    | Generate QR code for WIFI configuration                                                                                               |
| util release install                                | Download & install watermint toolbox to the path                                                                                      |
| util table format xlsx                              | Formatting xlsx file into text                                                                                                        |
| util text case down                                 | Print lower case text                                                                                                                 |
| util text case up                                   | Print upper case text                                                                                                                 |
| util text encoding from                             | Convert text encoding to UTF-8 text file from specified encoding.                                                                     |
| util text encoding to                               | Convert text encoding to specified encoding from UTF-8 text file.                                                                     |
| util text nlp english entity                        | Split English text into entities                                                                                                      |
| util text nlp english sentence                      | Split English text into sentences                                                                                                     |
| util text nlp english token                         | Split English text into tokens                                                                                                        |
| util text nlp japanese token                        | Tokenize Japanese text                                                                                                                |
| util text nlp japanese wakati                       | Wakachigaki (tokenize Japanese text)                                                                                                  |
| util tidy move dispatch                             | Dispatch files                                                                                                                        |
| util tidy move simple                               | Archive local files                                                                                                                   |
| util tidy pack remote                               | Package remote folder into the zip file                                                                                               |
| util time now                                       | Display current time                                                                                                                  |
| util unixtime format                                | Time format to convert the unix time (epoch seconds from 1970-01-01)                                                                  |
| util unixtime now                                   | Display current time in unixtime                                                                                                      |
| util uuid timestamp                                 | UUID Timestamp                                                                                                                        |
| util uuid ulid                                      | ULID Utility                                                                                                                          |
| util uuid v4                                        | Generate UUID v4 (random UUID)                                                                                                        |
| util uuid v7                                        | Generate UUID v7                                                                                                                      |
| util uuid version                                   | Parse version and variant of UUID                                                                                                     |
| util xlsx create                                    | Create an empty spreadsheet                                                                                                           |
| util xlsx sheet export                              | Export data from the xlsx file                                                                                                        |
| util xlsx sheet import                              | Import data into xlsx file                                                                                                            |
| util xlsx sheet list                                | List sheets of the xlsx file                                                                                                          |
| version                                             | Show version                                                                                                                          |



