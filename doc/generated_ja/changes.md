# `リリース 84` から `リリース 85` までの変更点

# 追加されたコマンド


| コマンド                                  | タイトル                                  |
|-------------------------------------------|-------------------------------------------|
| dev stage griddata                        | Grid data test                            |
| services google sheets sheet append       | Append data to a spreadsheet              |
| services google sheets sheet clear        | Clears values from a spreadsheet          |
| services google sheets sheet export       | Export sheet data                         |
| services google sheets sheet import       | Import data into the spreadsheet          |
| services google sheets sheet list         | List sheets of the spreadsheet            |
| services google sheets spreadsheet create | Create a new spreadsheet                  |
| util decode base_32                       | Decode text from Base32 (RFC 4648) format |
| util decode base_64                       | Decode text from Base64 (RFC 4648) format |
| util encode base_32                       | Encode text into Base32 (RFC 4648) format |
| util encode base_64                       | Encode text into Base64 (RFC 4648) format |



# コマンド仕様の変更: `config disable`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Key", Desc: "Feature key.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `config enable`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Key", Desc: "Feature key.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `config features`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `connect business_audit`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `connect business_file`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `connect business_info`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `connect business_mgmt`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `connect user_file`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev benchmark local`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to create", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "SizeMaxKb", Desc: "Maximum file size (KiB).", Default: "2048", TypeName: "int", ...}, &{Name: "SizeMinKb", Desc: "Minimum file size (KiB).", Default: "0", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev benchmark upload`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KiB", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev benchmark uploadlink`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}, &{Name: "SizeKb", Desc: "Size in KB", Default: "1024", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev build catalogue`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev build doc`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Badge", Desc: "Include badges of build status", Default: "true", TypeName: "bool", ...}, &{Name: "CommandPath", Desc: "Relative path to generate command manuals", Default: "doc/generated/", TypeName: "string", ...}, &{Name: "DocLang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Readme", Desc: "Filename of README", Default: "README.md", TypeName: "string", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev build license`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DestPath", Desc: "Dest path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "SourcePath", Desc: "Path to licenses (go-licenses output folder)", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev build preflight`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev build readme`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to create README.txt", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev ci artifact connect`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Full", Desc: "Account alias", Default: "deploy", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev ci artifact up`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path to upload", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string", ...}, &{Name: "Timeout", Desc: "Operation timeout in seconds", Default: "60", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev ci auth connect`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Audit", Desc: "Authenticate with Dropbox Business Audit scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "File", Desc: "Authenticate with Dropbox Business member file access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Full", Desc: "Authenticate with Dropbox user full access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Github", Desc: "Account alias for Github deployment", Default: "deploy", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev ci auth export`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Audit", Desc: "Authenticate with Dropbox Business Audit scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "File", Desc: "Authenticate with Dropbox Business member file access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Full", Desc: "Authenticate with Dropbox user full access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Info", Desc: "Authenticate with Dropbox Business info scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev ci auth import`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EnvName", Desc: "Environment variable name", Default: "TOOLBOX_ENDTOEND_TOKEN", TypeName: "string", ...}, &{Name: "PeerName", Desc: "Account alias", Default: "end_to_end_test", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev diag endpoint`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "JobId", Desc: "Job Id to diagnosis", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev diag throughput`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Bucket", Desc: "Bucket size in milliseconds", Default: "1000", TypeName: "int", ...}, &{Name: "EndpointName", Desc: "Filter by endpoint. Filter by exact match to the name."}, &{Name: "EndpointNamePrefix", Desc: "Filter by endpoint. Filter by name match to the prefix."}, &{Name: "EndpointNameSuffix", Desc: "Filter by endpoint. Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev kvs dump`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to KVS data", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev release candidate`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
## 追加されたレポート


| 名称   | 説明               |
|--------|--------------------|
| result | Recipe test result |


# コマンド仕様の変更: `dev release publish`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
  		"ConnGithub": "github_repo",
+ 		"Peer":       "github_repo",
  	},
  	Services: {"github"},
  	IsSecret: true,
  	... // 5 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ArtifactPath", Desc: "Path to artifacts", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "Target branch", Default: "master", TypeName: "string", ...}, &{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
## 追加されたレポート


| 名称   | 説明               |
|--------|--------------------|
| commit | Commit information |
| result | Recipe test result |


# コマンド仕様の変更: `dev replay approve`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Job Id.", TypeName: "string"}, &{Name: "Name", Desc: "Extra name of the approved recipe", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "ReplayPath", Desc: "Replay repository path. Fall back to the environment variable `T"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "WorkspacePath", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev replay bundle`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "PeerName", Desc: "Account alias for deployment.", Default: "deploy", TypeName: "string", ...}, &{Name: "ReplayPath", Desc: "Replay repository path. Fall back to the environment variable `T"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "ResultsPath", Desc: "Test failure result destination path.", Default: "/watermint-toolbox-logs/{{.Date}}-{{.Time}}/{{.Random}}", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl", ...}, &{Name: "Timeout", Desc: "Test result upload operation timeout.", Default: "60", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev replay recipe`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Job Id.", TypeName: "string"}, &{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev replay remote`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ReplayUrl", Desc: "Replay bundle shared link url", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev spec diff`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FilePath", Desc: "File path to output", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Lang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Release1", Desc: "Release name 1", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Release2", Desc: "Release name 2", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev spec doc`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FilePath", Desc: "File path", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Lang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev stage gmail`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "User id. The special value 'me' can be used to indicate the auth"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev stage gui`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev stage scoped`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Individual", Desc: "Account alias for individual", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}, &{Name: "Team", Desc: "Account alias for team", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev stage teamfolder`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev test echo`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Text", Desc: "Text to echo", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev test kvsfootprint`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Count", Desc: "Test count", Default: "1", TypeName: "int", ...}, &{Name: "Duplicate", Desc: "Create duplicated records", Default: "1", TypeName: "int", ...}, &{Name: "NumEntries", Desc: "Specify number of entries to write", Default: "1", TypeName: "int", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev test monkey`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Distribution", Desc: "Number of files/folder distribution", Default: "10000", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Extension", Desc: "File extensions (comma separated)", Default: "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"..., TypeName: "string", ...}, &{Name: "Path", Desc: "Monkey test path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev test recipe`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool", ...}, &{Name: "NoTimeout", Desc: "Do not timeout running recipe tests", Default: "false", TypeName: "bool", ...}, &{Name: "Single", Desc: "Recipe name to test", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Verbose", Desc: "Verbose output for testing", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
## 追加されたレポート


| 名称   | 説明               |
|--------|--------------------|
| result | Recipe test result |


# コマンド仕様の変更: `dev test resources`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev util anonymise`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "JobIdName", Desc: "Filter by job id name Filter by exact match to the name."}, &{Name: "JobIdNamePrefix", Desc: "Filter by job id name Filter by name match to the prefix."}, &{Name: "JobIdNameSuffix", Desc: "Filter by job id name Filter by name match to the suffix."}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev util curl`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BufferSize", Desc: "Size of buffer", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Record", Desc: "Capture record(s) for the test", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev util image jpeg`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Count", Desc: "Number of files to generate", Default: "10", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Height", Desc: "Height", Default: "1080", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "NamePrefix", Desc: "Filename prefix", Default: "test_image", TypeName: "string", ...}, &{Name: "Path", Desc: "Path to generate files", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `dev util wait`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Seconds", Desc: "Wait seconds", Default: "1", TypeName: "essentials.model.mo_int.range_int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file archive local`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "The destination folder path. The command will create folders if "..., TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "ExcludeFolders", Desc: "Exclude folders", Default: "false", TypeName: "bool", ...}, &{Name: "IncludeSystemFiles", Desc: "Include system files", Default: "false", TypeName: "bool", ...}, &{Name: "Preview", Desc: "Preview mode", Default: "false", TypeName: "bool", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file compare account`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Left", Desc: "Account alias (left)", Default: "left", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "LeftPath", Desc: "The path from account root (left)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Right", Desc: "Account alias (right)", Default: "right", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "RightPath", Desc: "The path from account root (right)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file compare local`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file copy`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file delete`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file dispatch local`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Preview", Desc: "Preview mode", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file download`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "File path to download", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path to download", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file export doc`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox document path to export.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path to save", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file import batch url`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Path", Desc: "Path to import", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file import url`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to import", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "URL", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file info`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file lock acquire`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "File path to lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file lock all release`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to release locks", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file lock batch acquire`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file lock batch release`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file lock list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file lock release`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to the file", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file merge`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool", ...}, &{Name: "From", Desc: "Path for merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "KeepEmptyFolder", Desc: "Keep empty folder after merge", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file mount list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file move`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file replication`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Account alias (destionation)", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "DstPath", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Src", Desc: "Account alias (source)", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "SrcPath", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file restore`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file search content`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}}, &{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file search name`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("image"), string("document"), string("pdf"), ...}}}, &{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file size`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "Report an entry for all files and folders depth folders deep", Default: "2", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "Path to scan", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file sync down`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Delete", Desc: "Delete local file if a file removed on Dropbox", Default: "false", TypeName: "bool", ...}, &{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "NameDisableIgnore", Desc: "Filter by name. Filter system file or ignore files."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file sync online`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Delete", Desc: "Delete file if a file removed in source path", Default: "false", TypeName: "bool", ...}, &{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "NameDisableIgnore", Desc: "Filter by name. Filter system file or ignore files."}, &{Name: "NameName", Desc: "Filter by name. Filter by exact match to the name."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file sync up`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KB", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Delete", Desc: "Delete Dropbox file if a file removed locally", Default: "false", TypeName: "bool", ...}, &{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local file path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `file watch`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to watch", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "Watch path recursively", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `filerequest create`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AllowLateUploads", Desc: "If set, allow uploads after the deadline has passed (one_day/two"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Deadline", Desc: "The deadline for this file request.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Path", Desc: "The path for the folder in the Dropbox where uploaded files will"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `filerequest delete closed`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `filerequest delete url`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Force", Desc: "Force delete the file request.", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "URL of the file request.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `filerequest list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group add`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ManagementType", Desc: "Group management type `company_managed` or `user_managed`", Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Name", Desc: "Group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group batch delete`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file for group name list", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group delete`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group folder list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "GroupName", Desc: "Filter by group name. Filter by exact match to the name."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group member add`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "GroupName", Desc: "Group name", TypeName: "string"}, &{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group member batch add`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group member batch delete`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group member batch update`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group member delete`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "GroupName", Desc: "Name of the group", TypeName: "string"}, &{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group member list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `group rename`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "CurrentName", Desc: "Current group name", TypeName: "string"}, &{Name: "NewName", Desc: "New group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `image info`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to the image", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job history archive`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Days", Desc: "Target days old", Default: "7", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job history delete`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Days", Desc: "Target days old", Default: "28", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job history list`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job history ship`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job log jobid`


## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Job ID", TypeName: "string"}, &{Name: "Kind", Desc: "Kind of log", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job log kind`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Kind", Desc: "Log kind.", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `job log last`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Kind", Desc: "Log kind", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `license`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member clear externalid`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "TransferDestMember", Desc: "If provided, files from the deleted member account will be trans"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "TransferNotifyAdminEmailOnError", Desc: "If provided, errors during the transfer process will be sent via"..., TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member detach`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "RevokeTeamShares", Desc: "True for revoke shared folder access which owned by the team", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Batch operation size", Default: "100", TypeName: "int", ...}, &{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member file permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Team member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member folder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "MemberEmail", Desc: "Filter by member email address. Filter by email address."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member folder replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstMemberEmail", Desc: "Destination team member email address", TypeName: "string"}, &{Name: "DstPath", Desc: "The path for the destination team member. Note the root (/) path"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "SrcMemberEmail", Desc: "Source team member email address", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member invite`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "SilentInvite", Desc: "Do not send welcome email (requires SSO + domain verification in"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "IncludeDeleted", Desc: "Include deleted members.", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member quota list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member quota update`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "Quota", Desc: "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom q"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member quota usage`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member reinvite`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "Silent", Desc: "Do not send welcome email (SSO required)", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Destination team; team file access", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Src", Desc: "Source team; team file access", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member update email`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "UpdateUnverified", Desc: "Update an account which didn't verified email. If an account ema"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member update externalid`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member update invisible`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member update profile`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `member update visible`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services asana team list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "WorkspaceName", Desc: "Name or GID of the workspace. Filter by exact match to the name."}, &{Name: "WorkspaceNamePrefix", Desc: "Name or GID of the workspace. Filter by name match to the prefix."}, &{Name: "WorkspaceNameSuffix", Desc: "Name or GID of the workspace. Filter by name match to the suffix."}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services asana team project list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "TeamName", Desc: "Name or GID of the team Filter by exact match to the name."}, &{Name: "TeamNamePrefix", Desc: "Name or GID of the team Filter by name match to the prefix."}, &{Name: "TeamNameSuffix", Desc: "Name or GID of the team Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services asana team task list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "ProjectName", Desc: "Name or GID of the project Filter by exact match to the name."}, &{Name: "ProjectNamePrefix", Desc: "Name or GID of the project Filter by name match to the prefix."}, &{Name: "ProjectNameSuffix", Desc: "Name or GID of the project Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services asana workspace list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services asana workspace project list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "WorkspaceName", Desc: "Name or GID of the workspace. Filter by exact match to the name."}, &{Name: "WorkspaceNamePrefix", Desc: "Name or GID of the workspace. Filter by name match to the prefix."}, &{Name: "WorkspaceNameSuffix", Desc: "Name or GID of the workspace. Filter by name match to the suffix."}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github content get`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Path", Desc: "Path to the content", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Ref", Desc: "Name of reference", TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github content put`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Branch", Desc: "Name of the branch", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Content", Desc: "Path to a content file", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Message", Desc: "Commit message", TypeName: "string"}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github issue list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "Repository name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github profile`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github release asset download`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Path", Desc: "Path to download", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "Release tag name", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github release asset list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "Release tag name", TypeName: "string"}, &{Name: "Repository", Desc: "Name of the repository", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github release asset upload`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Asset", Desc: "Path to assets", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "Release tag name", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github release draft`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BodyFile", Desc: "File path to body text. THe file must encoded in UTF-8 without BOM.", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "Name of the target branch", TypeName: "string"}, &{Name: "Name", Desc: "Name of the release", TypeName: "string"}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github release list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Repository owner", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "Repository name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services github tag create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "Name of the repository", TypeName: "string"}, &{Name: "Sha1", Desc: "SHA1 hash of the commit", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail filter add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...}, &{Name: "AddLabels", Desc: "List of labels to add to the message, separated by ','.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "CriteriaExcludeChats", Desc: "Whether the response should exclude chats.", Default: "false", TypeName: "bool", ...}, &{Name: "CriteriaFrom", Desc: "The sender's display name or email address.", TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail filter batch add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...}, &{Name: "ApplyToExistingMessages", Desc: "Apply labels to existing messages that satisfy the query.", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail filter delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Filter Id", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail filter list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail label add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ColorBackground", Desc: "The background color.", TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}}}, &{Name: "ColorText", Desc: "The text color.", TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]interface{}{"options": []interface{}{string(""), string("#000000"), string("#434343"), string("#666666"), ...}}}, &{Name: "LabelListVisibility", Desc: "The visibility of the label in the label list in the Gmail web i"..., Default: "labelShow", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "MessageListVisibility", Desc: "The visibility of messages with this label in the message list i"..., Default: "show", TypeName: "essentials.model.mo_string.select_string", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail label delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Name of the label", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail label list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail label rename`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "CurrentName", Desc: "Current label name", TypeName: "string"}, &{Name: "NewName", Desc: "New label name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail message label add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...}, &{Name: "Label", Desc: "Label names to add this message.", TypeName: "string"}, &{Name: "MessageId", Desc: "The immutable ID of the message.", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail message label delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Label", Desc: "Label names to remove this message.", TypeName: "string"}, &{Name: "MessageId", Desc: "The immutable ID of the message.", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail message list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Format", Desc: "The format to return the message in. ", Default: "metadata", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...}, &{Name: "Labels", Desc: "Only return messages with labels that match all of the specified"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "MaxResults", Desc: "Maximum number of messages to return.", Default: "20", TypeName: "int", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail message processed list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Format", Desc: "The format to return the message in. ", Default: "metadata", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...}, &{Name: "Labels", Desc: "Only return messages with labels that match all of the specified"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "MaxResults", Desc: "Maximum number of messages to return.", Default: "20", TypeName: "int", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services google mail thread list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `services slack conversation list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "Peer", TypeName: "domain.slack.api.work_conn_impl.conn_slack_api", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `sharedfolder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `sharedfolder member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `sharedlink create`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Expires", Desc: "Expiration date/time of the shared link", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Password", Desc: "Password", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `sharedlink delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "File or folder path to remove shared link", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "Attempt to remove the file hierarchy", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `sharedlink file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Password", Desc: "Password for the shared link", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "Shared link URL", TypeName: "domain.dropbox.model.mo_url.url_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `sharedlink list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team activity batch user`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "File", Desc: "User email address list file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team activity daily event`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Event category", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndDate", Desc: "End date", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team activity event`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team activity user`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team content member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team content mount list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Filter members. Filter by email address."}, &{Name: "MemberName", Desc: "Filter members. Filter by exact match to the name."}, &{Name: "MemberNamePrefix", Desc: "Filter members. Filter by name match to the prefix."}, &{Name: "MemberNameSuffix", Desc: "Filter members. Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team content policy list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team device list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team device unlink`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DeleteOnUnlink", Desc: "Delete files on unlink", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team diag explorer`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "All", Desc: "Include additional reports", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Dropbox Business file access", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Info", Desc: "Dropbox Business information access", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}, &{Name: "Mgmt", Desc: "Dropbox Business management", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team feature`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team filerequest clone`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team filerequest list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team info`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team linkedapp list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team namespace file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "IncludeDeleted", Desc: "If true, deleted file or folder will be returned", Default: "false", TypeName: "bool", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team namespace file size`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "Report entry for all files and directories depth directories deep", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team namespace list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team namespace member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AllColumns", Desc: "Show all columns", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team report activity`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team report devices`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team report membership`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team report storage`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team sharedlink list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Visibility", Desc: "Filter links by visibility (public/team_only/password)", Default: "all", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `team sharedlink update expiry`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "At", Desc: "New expiration date and time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]interface{}{"optional": bool(true)}}, &{Name: "Days", Desc: "Days to the new expiration date", Default: "0", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Visibility", Desc: "Target link visibility", Default: "public", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "SyncSetting", Desc: "Sync setting for the team folder", Default: "default", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder archive`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder batch archive`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder batch permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder batch replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstPeerName", Desc: "Destination team account alias", Default: "dst", TypeName: "string", ...}, &{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "SrcPeerName", Desc: "Source team account alias", Default: "src", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder file list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock all release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder file lock release`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder file size`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "Depth", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder member add`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AdminGroupName", Desc: "Temporary group name for admin operation", Default: "watermint-toolbox-admin", TypeName: "string", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder member delete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AdminGroupName", Desc: "Temporary group name for admin operation", Default: "watermint-toolbox-admin", TypeName: "string", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder member list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder partial replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Peer name for the destination team", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "DstPath", Desc: "Relative path from the team folder (please specify '/' for the t"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "DstTeamFolderName", Desc: "Destination team folder name", TypeName: "string"}, &{Name: "Src", Desc: "Peer name for the src team", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder permdelete`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder policy list`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `teamfolder replication`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstPeerName", Desc: "Destination team account alias", Default: "dst", TypeName: "string", ...}, &{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "SrcPeerName", Desc: "Source team account alias", Default: "src", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
# コマンド仕様の変更: `version`

## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  }
```
