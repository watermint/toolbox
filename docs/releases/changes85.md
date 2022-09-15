---
layout: release
title: Changes of Release 84
lang: en
---

# Changes between `Release 84` to `Release 85`

# Commands added


| Command                                   | Title                                                                |
|-------------------------------------------|----------------------------------------------------------------------|
| dev stage griddata                        | Grid data test                                                       |
| services google sheets sheet append       | Append data to a spreadsheet                                         |
| services google sheets sheet clear        | Clears values from a spreadsheet                                     |
| services google sheets sheet export       | Export sheet data                                                    |
| services google sheets sheet import       | Import data into the spreadsheet                                     |
| services google sheets sheet list         | List sheets of the spreadsheet                                       |
| services google sheets spreadsheet create | Create a new spreadsheet                                             |
| util date today                           | Display current date                                                 |
| util datetime now                         | Display current date/time                                            |
| util decode base_32                       | Decode text from Base32 (RFC 4648) format                            |
| util decode base_64                       | Decode text from Base64 (RFC 4648) format                            |
| util encode base_32                       | Encode text into Base32 (RFC 4648) format                            |
| util encode base_64                       | Encode text into Base64 (RFC 4648) format                            |
| util time now                             | Display current time                                                 |
| util unixtime format                      | Time format to convert the unix time (epoch seconds from 1970-01-01) |
| util unixtime now                         | Display current time in unixtime                                     |
| util xlsx create                          | Create an empty spreadsheet                                          |
| util xlsx sheet export                    | Export data from the xlsx file                                       |
| util xlsx sheet import                    | Import data into xlsx file                                           |
| util xlsx sheet list                      | List sheets of the xlsx file                                         |



# Command spec changed: `config disable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Key", Desc: "Feature key.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `config enable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Key", Desc: "Feature key.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `config features`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `connect business_audit`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `connect business_file`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `connect business_info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `connect business_mgmt`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `connect user_file`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev benchmark local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to create", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "SizeMaxKb", Desc: "Maximum file size (KiB).", Default: "2048", TypeName: "int", ...}, &{Name: "SizeMinKb", Desc: "Minimum file size (KiB).", Default: "0", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KiB", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "NumFiles", Desc: "Number of files.", Default: "1000", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to Dropbox", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev benchmark uploadlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}, &{Name: "SizeKb", Desc: "Size in KB", Default: "1024", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev build catalogue`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev build doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Badge", Desc: "Include badges of build status", Default: "true", TypeName: "bool", ...}, &{Name: "CommandPath", Desc: "Relative path to generate command manuals", Default: "doc/generated/", TypeName: "string", ...}, &{Name: "DocLang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Readme", Desc: "Filename of README", Default: "README.md", TypeName: "string", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev build license`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DestPath", Desc: "Dest path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "SourcePath", Desc: "Path to licenses (go-licenses output folder)", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev build preflight`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev build readme`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to create README.txt", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev ci artifact connect`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Full", Desc: "Account alias", Default: "deploy", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev ci artifact up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path to upload", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "PeerName", Desc: "Account alias", Default: "deploy", TypeName: "string", ...}, &{Name: "Timeout", Desc: "Operation timeout in seconds", Default: "60", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev ci auth connect`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Audit", Desc: "Authenticate with Dropbox Business Audit scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "File", Desc: "Authenticate with Dropbox Business member file access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Full", Desc: "Authenticate with Dropbox user full access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Github", Desc: "Account alias for Github deployment", Default: "deploy", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev ci auth export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Audit", Desc: "Authenticate with Dropbox Business Audit scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "File", Desc: "Authenticate with Dropbox Business member file access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Full", Desc: "Authenticate with Dropbox user full access scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Info", Desc: "Authenticate with Dropbox Business info scope", Default: "end_to_end_test", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev ci auth import`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EnvName", Desc: "Environment variable name", Default: "TOOLBOX_ENDTOEND_TOKEN", TypeName: "string", ...}, &{Name: "PeerName", Desc: "Account alias", Default: "end_to_end_test", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev diag endpoint`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "JobId", Desc: "Job Id to diagnosis", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev diag throughput`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Bucket", Desc: "Bucket size in milliseconds", Default: "1000", TypeName: "int", ...}, &{Name: "EndpointName", Desc: "Filter by endpoint. Filter by exact match to the name."}, &{Name: "EndpointNamePrefix", Desc: "Filter by endpoint. Filter by name match to the prefix."}, &{Name: "EndpointNameSuffix", Desc: "Filter by endpoint. Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev kvs dump`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to KVS data", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```

## Added report(s)


| Name   | Description        |
|--------|--------------------|
| result | Recipe test result |


# Command spec changed: `dev release publish`



## Command configuration changed


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
  	... // 6 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ArtifactPath", Desc: "Path to artifacts", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "Target branch", Default: "master", TypeName: "string", ...}, &{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```

## Added report(s)


| Name   | Description        |
|--------|--------------------|
| commit | Commit information |
| result | Recipe test result |


# Command spec changed: `dev replay approve`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Job Id.", TypeName: "string"}, &{Name: "Name", Desc: "Extra name of the approved recipe", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "ReplayPath", Desc: "Replay repository path. Fall back to the environment variable `T"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "WorkspacePath", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev replay bundle`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "PeerName", Desc: "Account alias for deployment.", Default: "deploy", TypeName: "string", ...}, &{Name: "ReplayPath", Desc: "Replay repository path. Fall back to the environment variable `T"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "ResultsPath", Desc: "Test failure result destination path.", Default: "/watermint-toolbox-logs/{% raw %}{{.{% endraw %}Date}}-{% raw %}{{.{% endraw %}Time}}/{% raw %}{{.{% endraw %}Random}}", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl", ...}, &{Name: "Timeout", Desc: "Test result upload operation timeout.", Default: "60", TypeName: "int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev replay recipe`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Job Id.", TypeName: "string"}, &{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev replay remote`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ReplayUrl", Desc: "Replay bundle shared link url", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev spec diff`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FilePath", Desc: "File path to output", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Lang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Release1", Desc: "Release name 1", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Release2", Desc: "Release name 2", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev spec doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FilePath", Desc: "File path", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Lang", Desc: "Language", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev stage gmail`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "User id. The special value 'me' can be used to indicate the auth"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev stage gui`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev stage scoped`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Individual", Desc: "Account alias for individual", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}, &{Name: "Team", Desc: "Account alias for team", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev stage teamfolder`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev test echo`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Text", Desc: "Text to echo", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev test kvsfootprint`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Count", Desc: "Test count", Default: "1", TypeName: "int", ...}, &{Name: "Duplicate", Desc: "Create duplicated records", Default: "1", TypeName: "int", ...}, &{Name: "NumEntries", Desc: "Specify number of entries to write", Default: "1", TypeName: "int", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev test monkey`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Distribution", Desc: "Number of files/folder distribution", Default: "10000", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Extension", Desc: "File extensions (comma separated)", Default: "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,"..., TypeName: "string", ...}, &{Name: "Path", Desc: "Monkey test path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev test recipe`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool", ...}, &{Name: "NoTimeout", Desc: "Do not timeout running recipe tests", Default: "false", TypeName: "bool", ...}, &{Name: "Single", Desc: "Recipe name to test", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Verbose", Desc: "Verbose output for testing", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```

## Added report(s)


| Name   | Description        |
|--------|--------------------|
| result | Recipe test result |


# Command spec changed: `dev test resources`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev util anonymise`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "JobIdName", Desc: "Filter by job id name Filter by exact match to the name."}, &{Name: "JobIdNamePrefix", Desc: "Filter by job id name Filter by name match to the prefix."}, &{Name: "JobIdNameSuffix", Desc: "Filter by job id name Filter by name match to the suffix."}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev util curl`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BufferSize", Desc: "Size of buffer", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Record", Desc: "Capture record(s) for the test", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev util image jpeg`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Count", Desc: "Number of files to generate", Default: "10", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Height", Desc: "Height", Default: "1080", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "NamePrefix", Desc: "Filename prefix", Default: "test_image", TypeName: "string", ...}, &{Name: "Path", Desc: "Path to generate files", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `dev util wait`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Seconds", Desc: "Wait seconds", Default: "1", TypeName: "essentials.model.mo_int.range_int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file archive local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "The destination folder path. The command will create folders if "..., TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "ExcludeFolders", Desc: "Exclude folders", Default: "false", TypeName: "bool", ...}, &{Name: "IncludeSystemFiles", Desc: "Include system files", Default: "false", TypeName: "bool", ...}, &{Name: "Preview", Desc: "Preview mode", Default: "false", TypeName: "bool", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file compare account`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Left", Desc: "Account alias (left)", Default: "left", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "LeftPath", Desc: "The path from account root (left)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Right", Desc: "Account alias (right)", Default: "right", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "RightPath", Desc: "The path from account root (right)", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file compare local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file dispatch local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Preview", Desc: "Preview mode", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "File path to download", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path to download", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file export doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox document path to export.", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path to save", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file import batch url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Path", Desc: "Path to import", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file import url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to import", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "URL", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "IncludeDeleted", Desc: "Include deleted files", Default: "false", TypeName: "bool", ...}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "List recursively", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file lock acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "File path to lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to release locks", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file lock batch acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file lock batch release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to the file", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file merge`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DryRun", Desc: "Dry run", Default: "true", TypeName: "bool", ...}, &{Name: "From", Desc: "Path for merge", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "KeepEmptyFolder", Desc: "Keep empty folder after merge", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Src", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Account alias (destionation)", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "DstPath", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Src", Desc: "Account alias (source)", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "SrcPath", Desc: "Source path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file restore`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file search content`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}}}, &{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file search name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Restricts search to only the file categories specified (image/do"..., TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("image"), string("document"), string("pdf"), ...}}}, &{Name: "Extension", Desc: "Restricts search to only the extensions specified.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Scopes the search to a path in the user's Dropbox.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "Report an entry for all files and folders depth folders deep", Default: "2", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "Path to scan", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file sync down`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Delete", Desc: "Delete local file if a file removed on Dropbox", Default: "false", TypeName: "bool", ...}, &{Name: "DropboxPath", Desc: "Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "NameDisableIgnore", Desc: "Filter by name. Filter system file or ignore files."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file sync online`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Delete", Desc: "Delete file if a file removed in source path", Default: "false", TypeName: "bool", ...}, &{Name: "Dst", Desc: "Destination path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "NameDisableIgnore", Desc: "Filter by name. Filter system file or ignore files."}, &{Name: "NameName", Desc: "Filter by name. Filter by exact match to the name."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ChunkSizeKb", Desc: "Upload chunk size in KB", Default: "65536", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Delete", Desc: "Delete Dropbox file if a file removed locally", Default: "false", TypeName: "bool", ...}, &{Name: "DropboxPath", Desc: "Destination Dropbox path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "LocalPath", Desc: "Local file path", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `file watch`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to watch", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "Watch path recursively", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `filerequest create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AllowLateUploads", Desc: "If set, allow uploads after the deadline has passed (one_day/two"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Deadline", Desc: "The deadline for this file request.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Path", Desc: "The path for the folder in the Dropbox where uploaded files will"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `filerequest delete closed`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `filerequest delete url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Force", Desc: "Force delete the file request.", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "URL of the file request.", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ManagementType", Desc: "Group management type `company_managed` or `user_managed`", Default: "company_managed", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Name", Desc: "Group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file for group name list", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "GroupName", Desc: "Filter by group name. Filter by exact match to the name."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "GroupName", Desc: "Group name", TypeName: "string"}, &{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "GroupName", Desc: "Name of the group", TypeName: "string"}, &{Name: "MemberEmail", Desc: "Email address of the member", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "CurrentName", Desc: "Current group name", TypeName: "string"}, &{Name: "NewName", Desc: "New group name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `image info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to the image", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job history archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Days", Desc: "Target days old", Default: "7", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job history delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Days", Desc: "Target days old", Default: "28", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job history list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job history ship`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DropboxPath", Desc: "Dropbox path to upload", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job log jobid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Job ID", TypeName: "string"}, &{Name: "Kind", Desc: "Kind of log", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "Path to the workspace", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job log kind`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Kind", Desc: "Log kind.", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `job log last`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Kind", Desc: "Log kind", Default: "toolbox", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "Path", Desc: "Path to workspace.", TypeName: "essentials.model.mo_string.opt_string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `license`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "TransferDestMember", Desc: "If provided, files from the deleted member account will be trans"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "TransferNotifyAdminEmailOnError", Desc: "If provided, errors during the transfer process will be sent via"..., TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "RevokeTeamShares", Desc: "True for revoke shared folder access which owned by the team", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Batch operation size", Default: "100", TypeName: "int", ...}, &{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member file permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Team member email address", TypeName: "string"}, &{Name: "Path", Desc: "Path to delete", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "MemberEmail", Desc: "Filter by member email address. Filter by email address."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member folder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstMemberEmail", Desc: "Destination team member email address", TypeName: "string"}, &{Name: "DstPath", Desc: "The path for the destination team member. Note the root (/) path"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "SrcMemberEmail", Desc: "Source team member email address", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "SilentInvite", Desc: "Do not send welcome email (requires SSO + domain verification in"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "IncludeDeleted", Desc: "Include deleted members.", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member quota list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member quota update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "Quota", Desc: "Custom quota in GB (1TB = 1024GB). 0 if the user has no custom q"..., Default: "0", TypeName: "essentials.model.mo_int.range_int", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member quota usage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "Silent", Desc: "Do not send welcome email (SSO required)", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Destination team; team file access", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Src", Desc: "Source team; team file access", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member update email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}, &{Name: "UpdateUnverified", Desc: "Update an account which didn't verified email. If an account ema"..., Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member update externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member update invisible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member update profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `member update visible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services asana team list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "WorkspaceName", Desc: "Name or GID of the workspace. Filter by exact match to the name."}, &{Name: "WorkspaceNamePrefix", Desc: "Name or GID of the workspace. Filter by name match to the prefix."}, &{Name: "WorkspaceNameSuffix", Desc: "Name or GID of the workspace. Filter by name match to the suffix."}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services asana team project list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "TeamName", Desc: "Name or GID of the team Filter by exact match to the name."}, &{Name: "TeamNamePrefix", Desc: "Name or GID of the team Filter by name match to the prefix."}, &{Name: "TeamNameSuffix", Desc: "Name or GID of the team Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services asana team task list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "ProjectName", Desc: "Name or GID of the project Filter by exact match to the name."}, &{Name: "ProjectNamePrefix", Desc: "Name or GID of the project Filter by name match to the prefix."}, &{Name: "ProjectNameSuffix", Desc: "Name or GID of the project Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services asana workspace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services asana workspace project list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "&{default <nil> default}", TypeName: "domain.asana.api.as_conn_impl.conn_asana_api", ...}, &{Name: "WorkspaceName", Desc: "Name or GID of the workspace. Filter by exact match to the name."}, &{Name: "WorkspaceNamePrefix", Desc: "Name or GID of the workspace. Filter by name match to the prefix."}, &{Name: "WorkspaceNameSuffix", Desc: "Name or GID of the workspace. Filter by name match to the suffix."}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github content get`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Path", Desc: "Path to the content", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Ref", Desc: "Name of reference", TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github content put`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Branch", Desc: "Name of the branch", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Content", Desc: "Path to a content file", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Message", Desc: "Commit message", TypeName: "string"}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github issue list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "Repository name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github release asset download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Path", Desc: "Path to download", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "Release tag name", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github release asset list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "Release tag name", TypeName: "string"}, &{Name: "Repository", Desc: "Name of the repository", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github release asset upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Asset", Desc: "Path to assets", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Release", Desc: "Release tag name", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github release draft`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BodyFile", Desc: "File path to body text. THe file must encoded in UTF-8 without BOM.", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}}, &{Name: "Branch", Desc: "Name of the target branch", TypeName: "string"}, &{Name: "Name", Desc: "Name of the release", TypeName: "string"}, &{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github release list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Repository owner", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "Repository name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services github tag create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...}, &{Name: "Repository", Desc: "Name of the repository", TypeName: "string"}, &{Name: "Sha1", Desc: "SHA1 hash of the commit", TypeName: "string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail filter add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...}, &{Name: "AddLabels", Desc: "List of labels to add to the message, separated by ','.", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "CriteriaExcludeChats", Desc: "Whether the response should exclude chats.", Default: "false", TypeName: "bool", ...}, &{Name: "CriteriaFrom", Desc: "The sender's display name or email address.", TypeName: "essentials.model.mo_string.opt_string"}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail filter batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...}, &{Name: "ApplyToExistingMessages", Desc: "Apply labels to existing messages that satisfy the query.", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail filter delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Id", Desc: "Filter Id", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail filter list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail label add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "ColorBackground", Desc: "The background color.", TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("#000000"), string("#434343"), string("#666666"), ...}}}, &{Name: "ColorText", Desc: "The text color.", TypeName: "essentials.model.mo_string.select_string", TypeAttr: map[string]any{"options": []any{string(""), string("#000000"), string("#434343"), string("#666666"), ...}}}, &{Name: "LabelListVisibility", Desc: "The visibility of the label in the label list in the Gmail web i"..., Default: "labelShow", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "MessageListVisibility", Desc: "The visibility of messages with this label in the message list i"..., Default: "show", TypeName: "essentials.model.mo_string.select_string", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail label delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Name of the label", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail label list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail label rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "CurrentName", Desc: "Current label name", TypeName: "string"}, &{Name: "NewName", Desc: "New label name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail message label add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AddLabelIfNotExist", Desc: "Create a label if it is not exist.", Default: "false", TypeName: "bool", ...}, &{Name: "Label", Desc: "Label names to add this message.", TypeName: "string"}, &{Name: "MessageId", Desc: "The immutable ID of the message.", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail message label delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Label", Desc: "Label names to remove this message.", TypeName: "string"}, &{Name: "MessageId", Desc: "The immutable ID of the message.", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail message list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Format", Desc: "The format to return the message in. ", Default: "metadata", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...}, &{Name: "Labels", Desc: "Only return messages with labels that match all of the specified"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "MaxResults", Desc: "Maximum number of messages to return.", Default: "20", TypeName: "int", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail message processed list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Format", Desc: "The format to return the message in. ", Default: "metadata", TypeName: "essentials.model.mo_string.select_string", ...}, &{Name: "IncludeSpamTrash", Desc: "Include messages from SPAM and TRASH in the results.", Default: "false", TypeName: "bool", ...}, &{Name: "Labels", Desc: "Only return messages with labels that match all of the specified"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "MaxResults", Desc: "Maximum number of messages to return.", Default: "20", TypeName: "int", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services google mail thread list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.google.api.goog_conn_impl.conn_google_mail", ...}, &{Name: "UserId", Desc: "The user's email address. The special value me can be used to in"..., Default: "me", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `services slack conversation list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "Peer", TypeName: "domain.slack.api.work_conn_impl.conn_slack_api", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `sharedfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `sharedlink create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Expires", Desc: "Expiration date/time of the shared link", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Password", Desc: "Password", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `sharedlink delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "File or folder path to remove shared link", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Recursive", Desc: "Attempt to remove the file hierarchy", Default: "false", TypeName: "bool", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `sharedlink file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Password", Desc: "Password for the shared link", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}, &{Name: "Url", Desc: "Shared link URL", TypeName: "domain.dropbox.model.mo_url.url_impl"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team activity batch user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "File", Desc: "User email address list file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team activity daily event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Event category", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndDate", Desc: "End date", TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team activity user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Category", Desc: "Filter the returned events to a single category. This field is o"..., TypeName: "essentials.model.mo_string.opt_string"}, &{Name: "EndTime", Desc: "Ending time (exclusive).", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit", ...}, &{Name: "StartTime", Desc: "Starting time (inclusive)", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team content member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team content mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "MemberEmail", Desc: "Filter members. Filter by email address."}, &{Name: "MemberName", Desc: "Filter members. Filter by exact match to the name."}, &{Name: "MemberNamePrefix", Desc: "Filter members. Filter by name match to the prefix."}, &{Name: "MemberNameSuffix", Desc: "Filter members. Filter by name match to the suffix."}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team content policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team device list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team device unlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DeleteOnUnlink", Desc: "Delete files on unlink", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team diag explorer`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "All", Desc: "Include additional reports", Default: "false", TypeName: "bool", ...}, &{Name: "File", Desc: "Dropbox Business file access", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Info", Desc: "Dropbox Business information access", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}, &{Name: "Mgmt", Desc: "Dropbox Business management", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team filerequest clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team linkedapp list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team namespace file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "IncludeDeleted", Desc: "If true, deleted file or folder will be returned", Default: "false", TypeName: "bool", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "Report entry for all files and directories depth directories deep", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team namespace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team namespace member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AllColumns", Desc: "Show all columns", Default: "false", TypeName: "bool", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team report activity`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team report devices`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team report membership`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team report storage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "EndDate", Desc: "End date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "StartDate", Desc: "Start date", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Visibility", Desc: "Filter links by visibility (public/team_only/password)", Default: "all", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "At", Desc: "New expiration date and time", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}}, &{Name: "Days", Desc: "Days to the new expiration date", Default: "0", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "Visibility", Desc: "Target link visibility", Default: "public", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, &{Name: "SyncSetting", Desc: "Sync setting for the team folder", Default: "default", TypeName: "essentials.model.mo_string.select_string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder batch archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder batch permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder batch replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstPeerName", Desc: "Destination team account alias", Default: "dst", TypeName: "string", ...}, &{Name: "File", Desc: "Data file for a list of team folder names", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "SrcPeerName", Desc: "Source team account alias", Default: "src", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "BatchSize", Desc: "Operation batch size", Default: "100", TypeName: "int", ...}, &{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Path", Desc: "Path to release lock", TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "TeamFolder", Desc: "Team folder name", TypeName: "string"}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Depth", Desc: "Depth", Default: "3", TypeName: "essentials.model.mo_int.range_int", ...}, &{Name: "FolderName", Desc: "List only for the folder matched to the name. Filter by exact ma"...}, &{Name: "FolderNamePrefix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, &{Name: "FolderNameSuffix", Desc: "List only for the folder matched to the name. Filter by name mat"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AdminGroupName", Desc: "Temporary group name for admin operation", Default: "watermint-toolbox-admin", TypeName: "string", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "AdminGroupName", Desc: "Temporary group name for admin operation", Default: "watermint-toolbox-admin", TypeName: "string", ...}, &{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "MemberTypeExternal", Desc: "Filter folder members. Keep only members are external (not in th"...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder partial replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Dst", Desc: "Peer name for the destination team", Default: "dst", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, &{Name: "DstPath", Desc: "Relative path from the team folder (please specify '/' for the t"..., TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl"}, &{Name: "DstTeamFolderName", Desc: "Destination team folder name", TypeName: "string"}, &{Name: "Src", Desc: "Peer name for the src team", Default: "src", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "FolderName", Desc: "Filter by folder name. Filter by exact match to the name."}, &{Name: "FolderNamePrefix", Desc: "Filter by folder name. Filter by name match to the prefix."}, &{Name: "FolderNameSuffix", Desc: "Filter by folder name. Filter by name match to the suffix."}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...}, ...},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `teamfolder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {&{Name: "DstPeerName", Desc: "Destination team account alias", Default: "dst", TypeName: "string", ...}, &{Name: "Name", Desc: "Team folder name", TypeName: "string"}, &{Name: "SrcPeerName", Desc: "Source team account alias", Default: "src", TypeName: "string", ...}},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
# Command spec changed: `version`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 18 identical fields
  	Feeds:          nil,
  	Values:         {},
- 	GridDataInput:  nil,
+ 	GridDataInput:  []*dc_recipe.DocGridDataInput{},
- 	GridDataOutput: nil,
+ 	GridDataOutput: []*dc_recipe.DocGridDataOutput{},
  	TextInput:      nil,
  	JsonInput:      nil,
  }
```
