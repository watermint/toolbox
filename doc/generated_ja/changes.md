# `リリース 65` から `リリース 66` までの変更点

# 追加されたコマンド

| コマンド                               | タイトル                                   |
|----------------------------------------|--------------------------------------------|
| dev catalogue                          | Generate catalogue                         |
| dev kvs dump                           | Dump KVS data                              |
| services github release asset download | Download assets                            |
| services github release asset upload   | Upload assets file into the GitHub Release |
| team filerequest clone                 | Clone file requests by given data          |


# 削除されたコマンド

| コマンド                         | タイトル                                   |
|----------------------------------|--------------------------------------------|
| dev desktop install              | Install Dropbox client                     |
| dev desktop start                | Launch Dropbox Desktop desktop app         |
| dev desktop stop                 | Try stopping Dropbox desktop app           |
| dev desktop suspendupdate        | Suspend/unsuspend Dropbox Updater          |
| dev diag procmon                 | Collect Process monitor logs               |
| services github release asset up | Upload assets file into the GitHub Release |
| web                              | Launch web console                         |


# コマンド仕様の変更: `config disable`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `config enable`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `config features`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `connect business_audit`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_audit"},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `connect business_file`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_file"},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `connect business_info`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_info"},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `connect business_mgmt`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_management"},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `connect user_file`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "user_full"},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `dev preflight`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `dev release candidate`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	Reports:        nil,
  	Feeds:          nil,
- 	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "Path to the test resource location",
- 			Default:  "test/dev/resource.json",
- 			TypeName: "string",
- 		},
- 	},
+ 	Values: []*dc_recipe.Value{},
  }
```
# コマンド仕様の変更: `dev release publish`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo"},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool"},
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "Path to test resource",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  	},
  }
```
# コマンド仕様の変更: `dev test recipe`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool"},
- 		&{
- 			Name:     "Recipe",
- 			Desc:     "Recipe name to test",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
+ 		&{
+ 			Name:     "NoTimeout",
+ 			Desc:     "Do not timeout running recipe tests",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{
- 			Name:     "Resource",
+ 			Name:     "Single",
- 			Desc:     "Test resource file path",
+ 			Desc:     "Recipe name to test",
  			Default:  "",
  			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Verbose", Desc: "Verbose output for testing", Default: "false", TypeName: "bool"},
  	},
  }
```
# コマンド仕様の変更: `dev util curl`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `dev util wait`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `file delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "Delete file or folder",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file delete",
  	CliArgs: "-path /PATH/TO/DELETE",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}},
  }
```
# コマンド仕様の変更: `file dispatch local`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "local",
  	Title:   "Dispatch local files",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file dispatch local",
  	CliArgs: "-file /PATH/TO/DATA_FILE.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Preview", Desc: "Preview mode", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `file import batch url`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Batch import files from URL",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file import batch url",
  	CliArgs: "-file /path/to/data/file -path /path/to/import",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Path", Desc: "Path to import", TypeName: "domain.common.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}},
  }
```
# コマンド仕様の変更: `file import url`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Import file from the URL",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file import url",
  	CliArgs: "-url URL -path /path/to/import",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Path", Desc: "Path to import", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "Url", Desc: "URL", TypeName: "string"}},
  }
```
# コマンド仕様の変更: `file merge`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "merge",
  	Title:   "Merge paths",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file merge",
  	CliArgs: "-from /from/path -to /path/to",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool"}, &{Name: "From", Desc: "Path for merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "KeepEmptyFolder", Desc: "Keep empty folder after merge", Default: "false", TypeName: "bool"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "To", Desc: "Path to merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "WithinSameNamespace", Desc: "Do not cross namespace. That is for preserve sharing permission including a shared link", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `file move`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "move",
  	Title:   "Move files",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file move",
  	CliArgs: "-src /SRC/PATH -dst /DST/PATH",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
  }
```
# コマンド仕様の変更: `file replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "Replicate file content to the other account",
  	Desc:    "This command will replicate files/folders. But it does not include sharing permissions. The command replicates only for folder contents of given path.",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file replication",
  	CliArgs: "-src source -src-path /path/src -dst dest -dst-path /path/dest",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Dst", Desc: "Account alias (destionation)", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "DstPath", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Src", Desc: "Account alias (source)", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "SrcPath", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
  }
```
# コマンド仕様の変更: `file restore`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "restore",
  	Title:   "Restore files under given path",
  	Desc:    "",
- 	Remarks: "(Experimental)",
+ 	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "file restore",
  	CliArgs: "-path /DROPBOX/PATH/TO/RESTORE",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}},
  }
```
# コマンド仕様の変更: `file sync up`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "up",
  	Title:   "Upstream sync with Dropbox",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file sync up",
  	CliArgs: "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KB", Default: "153600", TypeName: "domain.common.model.mo_int.range_int", TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)}}, &{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "FailOnError", Desc: "Returns error when any error happens while the operation. This command will not return any error when this flag is not enabled. All errors are written in the report.", Default: "false", TypeName: "bool"}, &{Name: "LocalPath", Desc: "Local file path", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}},
  }
```
# コマンド仕様の変更: `file upload`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "upload",
  	Title:   "Upload file",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file upload",
  	CliArgs: "-local-path /PATH/TO/UPLOAD -dropbox-path /DROPBOX/PATH",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KB", Default: "153600", TypeName: "domain.common.model.mo_int.range_int", TypeAttr: map[string]interface{}{"max": float64(153600), "min": float64(1), "value": float64(153600)}}, &{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local file path", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Overwrite", Desc: "Overwrite existing file(s)", Default: "false", TypeName: "bool"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}},
  }
```
# コマンド仕様の変更: `filerequest create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create a file request",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "filerequest create",
  	CliArgs: "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "AllowLateUploads", Desc: "If set, allow uploads after the deadline has passed (one_day/two_days/seven_days/thirty_days/always)", TypeName: "domain.common.model.mo_string.opt_string"}, &{Name: "Deadline", Desc: "The deadline for this file request.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Path", Desc: "The path for the folder in the Dropbox where uploaded files will be sent.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "Title", Desc: "The title of the file request", TypeName: "string"}},
  }
```
# コマンド仕様の変更: `filerequest delete closed`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "closed",
  	Title:   "Delete all closed file requests on this account.",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "filerequest delete closed",
  	CliArgs: "",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}},
  }
```
# コマンド仕様の変更: `filerequest delete url`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Delete a file request by the file request URL",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "filerequest delete url",
  	CliArgs: "",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Force", Desc: "Force delete the file request.", Default: "false", TypeName: "bool"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "Url", Desc: "URL of the file request.", TypeName: "string"}},
  }
```
# コマンド仕様の変更: `group add`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Create new group",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "group add",
  	CliArgs: "",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "ManagementType", Desc: "Group management type `company_managed` or `user_managed`", Default: "company_managed", TypeName: "domain.common.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string("company_managed"), string("user_managed")}}}, &{Name: "Name", Desc: "Group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}},
  }
```
# コマンド仕様の変更: `group member add`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Add a member to the group",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "group member add",
  	CliArgs: "",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "GroupName", Desc: "Group name", TypeName: "string"}, &{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}},
  }
```
# コマンド仕様の変更: `group rename`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "rename",
  	Title:   "Rename the group",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "group rename",
  	CliArgs: "",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "CurrentName", Desc: "Current group name", TypeName: "string"}, &{Name: "NewName", Desc: "New group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}},
  }
```
# コマンド仕様の変更: `job history archive`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `job history delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `job loop`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "loop",
  	Title:   "Run runbook until specified date/time",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Experimental)",
  	Path:    "job loop",
  	CliArgs: `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 17:58:38"`,
  	... // 3 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	Reports:        nil,
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `job run`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "run",
  	Title:   "Run workflow with *.runbook file",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Experimental)",
  	Path:    "job run",
  	CliArgs: "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	... // 3 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Fork",
- 			Desc:     "Fork process on running the workflow",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "RunbookPath", Desc: "Path to the runbook", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
- 		&{
- 			Name:     "TimeoutSeconds",
- 			Desc:     "Terminate process when given time passed",
- 			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(0), "value": float64(0)},
- 		},
  	},
  }
```
# コマンド仕様の変更: `member delete`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "Delete members",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member delete",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}, &{Name: "WipeData", Desc: "If true, controls if the user's data will be deleted on their linked devices", Default: "true", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `member detach`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "detach",
  	Title:   "Convert Dropbox Business accounts to a Basic account",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member detach",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}, &{Name: "RevokeTeamShares", Desc: "True for revoke shared folder access which owned by the team", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `member invite`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "invite",
  	Title:   "Invite member(s)",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member invite",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}, &{Name: "SilentInvite", Desc: "Do not send welcome email (requires SSO + domain verification instead)", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `member reinvite`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "reinvite",
  	Title:   "Reinvite invited status members to the team",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member reinvite",
  	CliArgs: "",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}, &{Name: "Silent", Desc: "Do not send welcome email (SSO required)", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `member update email`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "email",
  	Title:   "Member email operation",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member update email",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}, &{Name: "UpdateUnverified", Desc: "Update an account which didn't verified email. If an account email unverified, email address change may affect lose invitation to folders.", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `member update externalid`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
  	Title:   "Update External ID of team members",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member update externalid",
  	CliArgs: "-file /path/to/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}},
  }
```
# コマンド仕様の変更: `member update profile`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "profile",
  	Title:   "Update member profile",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member update profile",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"}},
  }
```
# コマンド仕様の変更: `services github issue list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `services github profile`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `services github release asset list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
## 変更されたレポート: assets

```
  &dc_recipe.Report{
  	Name: "assets",
  	Desc: "GitHub Release assets",
  	Columns: []*dc_recipe.ReportColumn{
  		... // 2 identical elements
  		&{Name: "state", Desc: "State of the asset"},
  		&{Name: "download_count", Desc: "Number of downloads"},
+ 		&{Name: "download_url", Desc: "Download URL"},
  	},
  }
```
# コマンド仕様の変更: `services github release draft`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "draft",
  	Title:   "Create release draft",
  	Desc:    "",
- 	Remarks: "(Experimental)",
+ 	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "services github release draft",
  	CliArgs: "-body-file /LOCAL/PATH/TO/body.txt",
  	... // 3 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "BodyFile", Desc: "File path to body text. THe file must encoded in UTF-8 without BOM.", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "Name of the target branch", TypeName: "string"}, &{Name: "Name", Desc: "Name of the release", TypeName: "string"}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo"}, &{Name: "Repository", Desc: "Name of the repository", TypeName: "string"}, &{Name: "Tag", Desc: "Name of the tag", TypeName: "string"}},
  }
```
# コマンド仕様の変更: `services github release list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: true,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `sharedlink create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create shared link",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "sharedlink create",
  	CliArgs: "-path /path/to/share",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Expires", Desc: "Expiration date/time of the shared link", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Password", Desc: "Password", TypeName: "domain.common.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file"}, &{Name: "TeamOnly", Desc: "Link is accessible only by team members", Default: "false", TypeName: "bool"}},
  }
```
# コマンド仕様の変更: `team diag explorer`


## 追加されたレポート

| 名称             | 説明                                                           |
|------------------|----------------------------------------------------------------|
| namespace_member | This report shows a list of members of namespaces in the team. |
| team_folder      | This report shows a list of team folders in the team.          |

# コマンド仕様の変更: `team namespace file list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool"},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool"},
  		&{
  			Name:     "Name",
  			Desc:     "List only for the folder matched to the name",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
# コマンド仕様の変更: `team namespace file size`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool"},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool"},
  		&{
  			Name:     "Name",
  			Desc:     "List only for the folder matched to the name",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
# コマンド仕様の変更: `teamfolder archive`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "Archive team folder",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "teamfolder archive",
  	CliArgs: "-name TEAMFOLDER_NAME",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"}},
  }
```
# コマンド仕様の変更: `teamfolder batch archive`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "Archiving team folders",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "teamfolder batch archive",
  	CliArgs: "-file /path/to/file.csv",
  	... // 5 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"}},
  }
```
# コマンド仕様の変更: `teamfolder replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "Replicate a team folder to the other team",
  	Desc:    "",
- 	Remarks: "(Irreversible operation)",
+ 	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "teamfolder replication",
  	CliArgs: "",
  	... // 4 identical fields
  	IsSecret:       false,
  	IsConsole:      false,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: true,
  	Reports:        nil,
  	... // 2 identical fields
  }
```
