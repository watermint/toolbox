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
| services github release asset up | Upload assets file into the GitHub Release |
| web                              | Launch web console                         |


# コマンド仕様の変更: `config disable`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `services github issue list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
- 	IsConsole:      true,
+ 	IsConsole:      false,
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
+ 		&{Name: "download_url", Desc: "Download URL"},
  	},
  }
```
# コマンド仕様の変更: `services github release draft`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `services github release list`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# コマンド仕様の変更: `services github tag create`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create a tag on the repository",
  	Desc:    "",
- 	Remarks: "(Experimental, and Irreversible operation)",
+ 	Remarks: "(Experimental)",
  	Path:    "services github tag create",
  	CliArgs: "",
  	... // 3 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: true,
+ 	IsIrreversible: false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values:         []*dc_recipe.Value{&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"}, &{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo"}, &{Name: "Repository", Desc: "Name of the repository", TypeName: "string"}, &{Name: "Sha1", Desc: "SHA1 hash of the commit", TypeName: "string"}, &{Name: "Tag", Desc: "Tag name", TypeName: "string"}},
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
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file"},
  	},
  }
```
# コマンド仕様の変更: `teamfolder replication`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "Replicate a team folder to the other team",
  	Desc:    "",
- 	Remarks: "(Irreversible operation)",
+ 	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "teamfolder replication",
  	CliArgs: "",
  	... // 4 identical fields
  	IsSecret:       false,
  	IsConsole:      false,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: true,
  	Reports:        nil,
  	... // 2 identical fields
  }
```
