# Changes between `Release 84` to `Release 85`

# Commands added

| Command                                   | Title                          |
|-------------------------------------------|--------------------------------|
| services google sheets sheet list         | List sheets of the spreadsheet |
| services google sheets spreadsheet create | Create a new spreadsheet       |

# Command spec changed: `dev release candidate`

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
  	... // 7 identical fields
  }
```

## Added report(s)

| Name   | Description        |
|--------|--------------------|
| commit | Commit information |
| result | Recipe test result |

# Command spec changed: `dev test recipe`

## Added report(s)

| Name   | Description        |
|--------|--------------------|
| result | Recipe test result |

# Command spec changed: `team diag explorer`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
  	ConnScopes: map[string]string{
  		"File": "business_file",
  		"Info": "business_info",
  		"Mgmt": "business_management",
- 		"Peer": "business_file",
+ 		"Peer": "business_management",
  	},
  	Services: {"dropbox_business"},
  	IsSecret: false,
  	... // 7 identical fields
  }
```
