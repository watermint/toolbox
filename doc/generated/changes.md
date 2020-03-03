# Changes between `Release 59` to `Current release`

# Commands added

| Command            | Title                                          |
|--------------------|------------------------------------------------|
| dev ci artifact up | Upload CI artifact                             |
| dev ci auth        | Authenticate for generating end to end testing |
| filerequest create | Create a file request                          |
| filerequest list   | List file requests of the individual account   |



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
- 	IsConsole:      false,
+ 	IsConsole:      true,
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
- 	IsConsole:      false,
+ 	IsConsole:      true,
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
- 	IsConsole:      false,
+ 	IsConsole:      true,
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
- 	IsConsole:      false,
+ 	IsConsole:      true,
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
- 	IsConsole:      false,
+ 	IsConsole:      true,
  	IsExperimental: false,
  	IsIrreversible: false,
  	... // 2 identical fields
  }

```

