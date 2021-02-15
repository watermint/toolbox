# Changes between `Release 84` to `Release 85`

# Commands added

| Command                                   | Title                            |
|-------------------------------------------|----------------------------------|
| services google sheets sheet clear        | Clears values from a spreadsheet |
| services google sheets sheet export       | Export sheet data                |
| services google sheets sheet list         | List sheets of the spreadsheet   |
| services google sheets spreadsheet create | Create a new spreadsheet         |

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
+ 		"Peer":       "github_repo",
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


