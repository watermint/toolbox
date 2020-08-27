# `リリース 72` から `リリース 73` までの変更点

# 追加されたコマンド


| コマンド                              | タイトル                              |
|---------------------------------------|---------------------------------------|
| dev build catalogue                   | Generate catalogue                    |
| dev build doc                         | Document generator                    |
| dev build license                     | Generate LICENSE.txt                  |
| dev build preflight                   | Process prerequisites for the release |
| dev build readme                      | Generate README.txt                   |
| dev test async                        | Async framework test                  |
| dev test echo                         | Echo text                             |
| services asana team list              | List team                             |
| services asana team project list      | List projects of the team             |
| services asana team task list         | List task of the team                 |
| services asana workspace list         | List workspaces                       |
| services asana workspace project list | List projects of the workspace        |



# 削除されたコマンド


| コマンド      | タイトル                              |
|---------------|---------------------------------------|
| dev async     | Async framework test                  |
| dev catalogue | Generate catalogue                    |
| dev doc       | Document generator                    |
| dev echo      | Echo text                             |
| dev preflight | Process prerequisites for the release |



# コマンド仕様の変更: `file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "IncludeMediaInfo",
- 			Desc:     "Include media information",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...},
  		&{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...},
  	},
  }
```
# コマンド仕様の変更: `services google mail filter batch add`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...},
  		&{
- 			Name:     "ApplyToInboxMessages",
+ 			Name:     "ApplyToExistingMessages",
- 			Desc:     "Apply labels to messages satisfy query in INBOX.",
+ 			Desc:     "Apply labels to existing messages that satisfy the query.",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...},
  		&{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...},
  	},
  }
```
# コマンド仕様の変更: `team diag explorer`


## 追加されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# コマンド仕様の変更: `team namespace file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "IncludeDeleted", Desc: "If true, deleted file or folder will be returned", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "IncludeMediaInfo",
- 			Desc:     "If true, media info is set for photo and video in json report",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "IncludeMemberFolder", Desc: "If true, include team member folders", Default: "false", TypeName: "bool", ...},
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		... // 3 identical elements
  	},
  }
```
## 追加されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# コマンド仕様の変更: `team namespace file size`


## 追加されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# コマンド仕様の変更: `teamfolder file list`


## 追加されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


# コマンド仕様の変更: `teamfolder file size`


## 追加されたレポート


| 名称   | 説明                                      |
|--------|-------------------------------------------|
| errors | This report shows the transaction result. |


