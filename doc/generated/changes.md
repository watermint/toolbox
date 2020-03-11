# Changes between `Release 59` to `Current release`

# Commands added

| Command                   | Title                                            |
|---------------------------|--------------------------------------------------|
| dev ci artifact up        | Upload CI artifact                               |
| dev ci auth               | Authenticate for generating end to end testing   |
| dev util curl             | Generate cURL preview from capture log           |
| filerequest create        | Create a file request                            |
| filerequest delete closed | Delete all closed file requests on this account. |
| filerequest delete url    | Delete a file request by the file request URL    |
| filerequest list          | List file requests of the individual account     |



# Commands deleted

| Command       | Title                                          |
|---------------|------------------------------------------------|
| dev test auth | Authenticate for generating end to end testing |



# Command spec changed: `connect business_audit`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_audit"},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 2 identical fields
  }

```

# Command spec changed: `connect business_file`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_file"},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 2 identical fields
  }

```

# Command spec changed: `connect business_info`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_info"},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 2 identical fields
  }

```

# Command spec changed: `connect business_mgmt`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "business_management"},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 2 identical fields
  }

```

# Command spec changed: `connect user_file`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 9 identical fields
  	ConnScopes:     map[string]string{"Peer": "user_full"},
  	IsSecret:       false,
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 2 identical fields
  }

```

# Command spec changed: `group batch delete`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "group batch delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# Command spec changed: `member quota update`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member quota update",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# Command spec changed: `member replication`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "member replication",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# Command spec changed: `member update externalid`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member update externalid",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# Command spec changed: `team activity batch user`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team activity batch user",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# Command spec changed: `teamfolder batch archive`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "teamfolder batch archive",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

# Command spec changed: `teamfolder batch permdelete`



## Command configuration changed



```
  &rc_doc.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "teamfolder batch permdelete",
- 	CliArgs:         "",
+ 	CliArgs:         "-file /path/to/file.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  }

```

