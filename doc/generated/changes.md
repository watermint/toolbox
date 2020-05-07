# Changes between `Release 65` to `Release 66`

# Commands added

| Command                                | Title                                      |
|----------------------------------------|--------------------------------------------|
| dev catalogue                          | Generate catalogue                         |
| dev kvs dump                           | Dump KVS data                              |
| services github release asset download | Download assets                            |
| services github release asset upload   | Upload assets file into the GitHub Release |
| team filerequest clone                 | Clone file requests by given data          |


# Commands deleted

| Command                          | Title                                      |
|----------------------------------|--------------------------------------------|
| services github release asset up | Upload assets file into the GitHub Release |
| web                              | Launch web console                         |


# Command spec changed: `config disable`


## Command configuration changed


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
# Command spec changed: `config enable`


## Command configuration changed


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
# Command spec changed: `config features`


## Command configuration changed


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
# Command spec changed: `connect business_audit`


## Command configuration changed


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
# Command spec changed: `connect business_file`


## Command configuration changed


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
# Command spec changed: `connect business_info`


## Command configuration changed


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
# Command spec changed: `connect business_mgmt`


## Command configuration changed


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
# Command spec changed: `connect user_file`


## Command configuration changed


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
# Command spec changed: `dev preflight`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# Command spec changed: `dev release candidate`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	Reports:        nil,
  	Feeds:          nil,
- 	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "Path to the test resource location",
- 			Default:  "test/dev/resource.json",
- 			TypeName: "string",
- 		},
- 	},
+ 	Values: []*dc_recipe.Value{},
  }
```
# Command spec changed: `dev release publish`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo"},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool"},
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "Path to test resource",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  	},
  }
```
# Command spec changed: `dev test recipe`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 14 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool"},
  		&{Name: "Recipe", Desc: "Recipe name to test", TypeName: "domain.common.model.mo_string.opt_string"},
- 		&{
- 			Name:     "Resource",
- 			Desc:     "Test resource file path",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  		&{Name: "Verbose", Desc: "Verbose output for testing", Default: "false", TypeName: "bool"},
  	},
  }
```
# Command spec changed: `dev util curl`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# Command spec changed: `dev util wait`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# Command spec changed: `job history archive`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# Command spec changed: `job history delete`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 3 identical fields
  }
```
# Command spec changed: `job loop`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "loop",
  	Title:   "Run runbook until specified date/time",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Experimental)",
  	Path:    "job loop",
  	CliArgs: `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 17:58:38"`,
  	... // 3 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	Reports:        nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `job run`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "run",
  	Title:   "Run workflow with *.runbook file",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Experimental)",
  	Path:    "job run",
  	CliArgs: "-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook",
  	... // 3 identical fields
  	ConnScopes:     map[string]string{},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	Reports:        nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `services github issue list`


## Command configuration changed


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
# Command spec changed: `services github profile`


## Command configuration changed


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
# Command spec changed: `services github release asset list`


## Command configuration changed


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
## Changed report: assets

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
# Command spec changed: `services github release draft`


## Command configuration changed


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
# Command spec changed: `services github release list`


## Command configuration changed


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
# Command spec changed: `services github tag create`


## Command configuration changed


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
# Command spec changed: `team diag explorer`


## Added report(s)

| Name             | Description                                                    |
|------------------|----------------------------------------------------------------|
| namespace_member | This report shows a list of members of namespaces in the team. |
| team_folder      | This report shows a list of team folders in the team.          |

# Command spec changed: `team namespace file list`


## Command configuration changed


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
# Command spec changed: `team namespace file size`


## Command configuration changed


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
# Command spec changed: `teamfolder replication`


## Command configuration changed


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
