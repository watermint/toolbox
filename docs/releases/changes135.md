---
layout: release
title: Changes of Release 134
lang: en
---

# Changes between `Release 134` to `Release 135`

# Commands added


| Command                        | Title                  |
|--------------------------------|------------------------|
| dropbox file sharedfolder info | Get shared folder info |



# Commands deleted


| Command                             | Title                                                 |
|-------------------------------------|-------------------------------------------------------|
| google calendar event list          | List Google Calendar events                           |
| google mail filter add              | Add a filter.                                         |
| google mail filter batch add        | Batch adding/deleting labels with query               |
| google mail filter delete           | Delete a filter                                       |
| google mail filter list             | List filters                                          |
| google mail label add               | Add a label                                           |
| google mail label delete            | Delete a label                                        |
| google mail label list              | List email labels                                     |
| google mail label rename            | Rename a label                                        |
| google mail message label add       | Add labels to the message                             |
| google mail message label delete    | Remove labels from the message                        |
| google mail message list            | List messages                                         |
| google mail message processed list  | List messages in processed format.                    |
| google mail message send            | Send a mail                                           |
| google mail sendas add              | Creates a custom "from" send-as alias                 |
| google mail sendas delete           | Deletes the specified send-as alias                   |
| google mail sendas list             | Lists the send-as aliases for the specified account   |
| google mail thread list             | List threads                                          |
| google sheets sheet append          | Append data to a spreadsheet                          |
| google sheets sheet clear           | Clears values from a spreadsheet                      |
| google sheets sheet create          | Create a new sheet                                    |
| google sheets sheet delete          | Delete a sheet from the spreadsheet                   |
| google sheets sheet export          | Export sheet data                                     |
| google sheets sheet import          | Import data into the spreadsheet                      |
| google sheets sheet list            | List sheets of the spreadsheet                        |
| google sheets spreadsheet create    | Create a new spreadsheet                              |
| google translate text               | Translate text                                        |
| teamspace asadmin file list         | List files and folders in team space run as admin     |
| teamspace asadmin folder add        | Create top level folder in the team space             |
| teamspace asadmin folder delete     | Delete top level folder of the team space             |
| teamspace asadmin folder permdelete | Permanently delete top level folder of the team space |
| teamspace asadmin member list       | List top level folder members                         |
| teamspace file list                 | List files and folders in team space                  |



# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github_public"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
