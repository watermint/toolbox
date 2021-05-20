---
layout: release
title: Changes of Release 65
lang: en
---

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
| dev desktop install              | Install Dropbox client                     |
| dev desktop start                | Launch Dropbox Desktop desktop app         |
| dev desktop stop                 | Try stopping Dropbox desktop app           |
| dev desktop suspendupdate        | Suspend/unsuspend Dropbox Updater          |
| dev diag procmon                 | Collect Process monitor logs               |
| services github release asset up | Upload assets file into the GitHub Release |
| web                              | Launch web console                         |



# Command spec changed: `config disable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `config enable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `config features`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `connect business_audit`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `connect business_file`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `connect business_info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `connect business_mgmt`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `connect user_file`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `dev preflight`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	IsTransient:    false,
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
+ 	Values:         []*dc_recipe.Value{},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release publish`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	IsTransient:    false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "ConnGithub", Desc: "Account alias", Default: "default", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "TestResource",
- 			Desc:     "Path to test resource",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev test recipe`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "All", Desc: "Test all recipes", Default: "false", TypeName: "bool", ...},
- 		&{
- 			Name:     "Recipe",
- 			Desc:     "Recipe name to test",
- 			TypeName: "domain.common.model.mo_string.opt_string",
- 		},
+ 		&{
+ 			Name:     "NoTimeout",
+ 			Desc:     "Do not timeout running recipe tests",
+ 			Default:  "false",
+ 			TypeName: "bool",
+ 		},
  		&{
- 			Name:     "Resource",
+ 			Name:     "Single",
- 			Desc:     "Test resource file path",
+ 			Desc:     "Recipe name to test",
  			Default:  "",
  			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Verbose", Desc: "Verbose output for testing", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `dev util curl`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `dev util wait`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       true,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `file delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "Delete file or folder",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file delete",
  	CliArgs: "-path /PATH/TO/DELETE",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file dispatch local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "local",
  	Title:   "Dispatch local files",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file dispatch local",
  	CliArgs: "-file /PATH/TO/DATA_FILE.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file import batch url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Batch import files from URL",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file import batch url",
  	CliArgs: "-file /path/to/data/file -path /path/to/import",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file import url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Import file from the URL",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file import url",
  	CliArgs: "-url URL -path /path/to/import",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file merge`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "merge",
  	Title:   "Merge paths",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file merge",
  	CliArgs: "-from /from/path -to /path/to",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "move",
  	Title:   "Move files",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file move",
  	CliArgs: "-src /SRC/PATH -dst /DST/PATH",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "replication",
  	Title:   "Replicate file content to the other account",
  	Desc:    "This command will replicate files/folders. But it does not inclu"...,
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file replication",
  	CliArgs: "-src source -src-path /path/src -dst dest -dst-path /path/dest",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file restore`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "restore",
  	Title:   "Restore files under given path",
  	Desc:    "",
- 	Remarks: "(Experimental)",
+ 	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "file restore",
  	CliArgs: "-path /DROPBOX/PATH/TO/RESTORE",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "up",
  	Title:   "Upstream sync with Dropbox",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file sync up",
  	CliArgs: "-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF"...,
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `file upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "upload",
  	Title:   "Upload file",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "file upload",
  	CliArgs: "-local-path /PATH/TO/UPLOAD -dropbox-path /DROPBOX/PATH",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `filerequest create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create a file request",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "filerequest create",
  	CliArgs: "-path /DROPBOX/PATH/OF/FILEREQUEST",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `filerequest delete closed`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "closed",
  	Title:   "Delete all closed file requests on this account.",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "filerequest delete closed",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `filerequest delete url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "url",
  	Title:   "Delete a file request by the file request URL",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "filerequest delete url",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Create new group",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "group add",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "add",
  	Title:   "Add a member to the group",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "group member add",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "rename",
  	Title:   "Rename the group",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "group rename",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `job history archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `job history delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 8 identical fields
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
  	CliArgs: `-runbook-path /LOCAL/PATH/TO/RUNBOOK.runbook -until "2020-04-01 `...,
  	... // 4 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	IsTransient:    false,
  	... // 7 identical fields
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
  	... // 4 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: false,
  	IsTransient:    false,
  	Reports:        nil,
  	Feeds:          nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Fork",
- 			Desc:     "Fork process on running the workflow",
- 			Default:  "false",
- 			TypeName: "bool",
- 		},
  		&{Name: "RunbookPath", Desc: "Path to the runbook", TypeName: "domain.common.model.mo_path.file_system_path_impl", TypeAttr: map[string]interface{}{"shouldExist": bool(false)}},
- 		&{
- 			Name:     "TimeoutSeconds",
- 			Desc:     "Terminate process when given time passed",
- 			Default:  "0",
- 			TypeName: "domain.common.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(3.1536e+07), "min": float64(0), "value": float64(0)},
- 		},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "delete",
  	Title:   "Delete members",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member delete",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `member detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "detach",
  	Title:   "Convert Dropbox Business accounts to a Basic account",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member detach",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `member invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "invite",
  	Title:   "Invite member(s)",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member invite",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `member reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "reinvite",
  	Title:   "Reinvite invited status members to the team",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member reinvite",
  	CliArgs: "",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `member update email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "email",
  	Title:   "Member email operation",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member update email",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `member update externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "externalid",
  	Title:   "Update External ID of team members",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member update externalid",
  	CliArgs: "-file /path/to/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `member update profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "profile",
  	Title:   "Update member profile",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "member update profile",
  	CliArgs: "-file /path/to/data/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `services github issue list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `services github profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `services github release asset list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
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
  	Name:    "draft",
  	Title:   "Create release draft",
  	Desc:    "",
- 	Remarks: "(Experimental)",
+ 	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "services github release draft",
  	CliArgs: "-body-file /LOCAL/PATH/TO/body.txt",
  	... // 4 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `services github release list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: false,
  	... // 8 identical fields
  }
```
# Command spec changed: `services github tag create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 10 identical fields
  	Services:       nil,
  	IsSecret:       false,
- 	IsConsole:      true,
+ 	IsConsole:      false,
  	IsExperimental: true,
  	IsIrreversible: true,
  	... // 8 identical fields
  }
```
# Command spec changed: `sharedlink create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "create",
  	Title:   "Create shared link",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "sharedlink create",
  	CliArgs: "-path /path/to/share",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
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
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "Name",
  			Desc:     "List only for the folder matched to the name",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "IncludeSharedFolder", Desc: "If true, include shared folders", Default: "true", TypeName: "bool", ...},
  		&{Name: "IncludeTeamFolder", Desc: "If true, include team folders", Default: "true", TypeName: "bool", ...},
  		&{
  			Name:     "Name",
  			Desc:     "List only for the folder matched to the name",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  	},
  	GridDataInput:  nil,
  	GridDataOutput: nil,
  	... // 2 identical fields
  }
```
# Command spec changed: `teamfolder archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "Archive team folder",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "teamfolder archive",
  	CliArgs: "-name TEAMFOLDER_NAME",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
  }
```
# Command spec changed: `teamfolder batch archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "archive",
  	Title:   "Archiving team folders",
  	Desc:    "",
- 	Remarks: "",
+ 	Remarks: "(Irreversible operation)",
  	Path:    "teamfolder batch archive",
  	CliArgs: "-file /path/to/file.csv",
  	... // 6 identical fields
  	IsConsole:      false,
  	IsExperimental: false,
- 	IsIrreversible: false,
+ 	IsIrreversible: true,
  	IsTransient:    false,
  	Reports:        nil,
  	... // 6 identical fields
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
  	... // 5 identical fields
  	IsSecret:       false,
  	IsConsole:      false,
- 	IsExperimental: false,
+ 	IsExperimental: true,
  	IsIrreversible: true,
  	IsTransient:    false,
  	... // 7 identical fields
  }
```
