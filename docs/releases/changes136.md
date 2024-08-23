---
layout: release
title: Changes of Release 135
lang: en
---

# Changes between `Release 135` to `Release 136`

# Commands added


| Command                  | Title                                                                    |
|--------------------------|--------------------------------------------------------------------------|
| config license install   | Install a license key                                                    |
| dropbox file restore ext | Restore files with a specific extension                                  |
| util feed json           | Load feed from the URL and output the content as JSON                    |
| util json query          | Query JSON data                                                          |
| util uuid timestamp      | Parse UUID timestamp                                                     |
| util uuid ulid           | Generate ULID (Universally Unique Lexicographically Sortable Identifier) |
| util uuid v7             | Generate UUID v7                                                         |
| util uuid version        | Parse version and variant of UUID                                        |



# Commands deleted


| Command                          | Title                                 |
|----------------------------------|---------------------------------------|
| util desktop display list        | List displays of the current machine  |
| util desktop screenshot interval | Take screenshots at regular intervals |
| util desktop screenshot snap     | Take a screenshot                     |



# Command spec changed: `dev benchmark upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dev benchmark uploadlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dev license issue`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 6 identical elements
  		&{Name: "LifecycleWarningAfter", Desc: "Lifecycle warning after this period from the build time (seconds)", Default: "31536000", TypeName: "int", ...},
  		&{Name: "Owner", Desc: "License repository owner", Default: "watermint", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
+ 		&{
+ 			Name:     "RecipeAllowedPrefix",
+ 			Desc:     "Recipe allowed prefix",
+ 			TypeName: "essentials.model.mo_string.opt_string",
+ 		},
  		&{Name: "RecipesAllowed", Desc: "Comma separated list of recipes allowed", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Repository", Desc: "License repository", Default: "toolbox-supplement", TypeName: "string", ...},
  		&{Name: "Scope", Desc: "License scope", TypeName: "essentials.model.mo_string.opt_string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release announcement`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "CategoryId", Desc: "Announcement category ID", Default: "DIC_kwDOBFqRm84CQesd", TypeName: "string", ...},
  		&{Name: "Owner", Desc: "Repository owner", Default: "watermint", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "Repository name", Default: "toolbox", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release asset`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Owner", Desc: "Repository owner", TypeName: "string"},
  		&{Name: "Path", Desc: "Content path", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repo", Desc: "Repository name", TypeName: "string"},
  		&{Name: "Text", Desc: "Text content", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release asseturl`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "SourceOwner", Desc: "Source repository owner", TypeName: "string"},
  		&{Name: "SourceRepo", Desc: "Source repository name", TypeName: "string"},
  		... // 3 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `dev release checkin`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
- 	IsSecret:        false,
+ 	IsSecret:        true,
  	IsConsole:       false,
  	IsExperimental:  false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Branch", Desc: "Repository branch", Default: "main", TypeName: "string", ...},
  		&{Name: "Owner", Desc: "Repository owner", Default: "watermint", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repo", Desc: "Repository name", Default: "toolbox", TypeName: "string", ...},
  		&{Name: "SupplementBranch", Desc: "Supplement repository branch name", Default: "main", TypeName: "string", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
- 			TypeAttr: string("github_public"),
+ 			TypeAttr: string("github"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev release publish`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"ConnGithub": "github_repo",
+ 		"ConnGithub": "github",
- 		"Peer":       "github_repo",
+ 		"Peer":       "github",
  	},
  	Services: {"github"},
  	IsSecret: true,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "ArtifactPath", Desc: "Path to artifacts", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Branch", Desc: "Target branch", Default: "main", TypeName: "string", ...},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "SkipTests", Desc: "Skip end to end tests.", Default: "false", TypeName: "bool", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dropbox file account feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file account filesystem`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file account info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file compare account`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Left": "dropbox_individual", "Right": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file compare local`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file export doc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file export url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file import batch url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file import url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file lock acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file lock batch acquire`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file lock batch release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file merge`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file move`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Dst": "dropbox_individual", "Src": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file request create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file request delete closed`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file request delete url`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file request list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file restore all`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file revision download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file revision list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file revision restore`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file search content`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file search name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file share info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedfolder unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedlink create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedlink delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedlink file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedlink info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sync down`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sync online`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file sync up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file tag add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file tag delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file tag list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file template apply`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file template capture`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox file watch`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox paper append`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox paper create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox paper overwrite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox paper prepend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team activity batch user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team activity daily event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team activity event`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team activity user`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin group role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin group role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin role add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin role clear`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin role delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team admin role list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team backup device status`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper count`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content legacypaper list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content member size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team content policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team device list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team device unlink`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team filerequest clone`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team filerequest list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team filesystem`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team group update type`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team info`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team insight scan`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team insight scanretry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold member batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold revision list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold update desc`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team legalhold update name`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team linkedapp list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member batch detach`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member batch invite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member batch reinvite`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member batch suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member batch unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member clear externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member feature`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member file permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member folder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member folder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member quota batch update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member quota list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member quota usage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member suspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member unsuspend`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member update batch email`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member update batch externalid`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member update batch invisible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member update batch profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team member update batch visible`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team namespace file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team namespace file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team namespace list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team namespace member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team namespace summary`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team report activity`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team report devices`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team report membership`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team report storage`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        true,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas file batch copy`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas file sync batch up`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch leave`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch share`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder batch unshare`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder isolate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder member batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder member batch delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team runas sharedfolder mount mountable`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink cap expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink cap visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink delete links`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink delete member`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink update password`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team sharedlink update visibility`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder batch archive`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder batch permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder batch replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock all release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file lock release`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder file size`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder member list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder partial replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder permdelete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder policy list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder replication`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Dst": "dropbox_team", "Src": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder sync setting list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `dropbox team teamfolder sync setting update`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: true,
  	ConnScopes:      {"Peer": "dropbox_team"},
- 	Services:        []string{"dropbox_business"},
+ 	Services:        []string{"dropbox_team"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `github content get`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to the content", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Ref", Desc: "Name of reference", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github content put`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 3 identical elements
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to the content", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github issue list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Filter", Desc: "Indicates which sorts of issues to return.", Default: "assigned", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  		&{Name: "Labels", Desc: "A list of comma separated label names.", TypeName: "essentials.model.mo_string.opt_string"},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "Repository name", TypeName: "string"},
  		&{Name: "Since", Desc: "Only show notifications updated after the given time.", TypeName: "domain.dropbox.model.mo_time.time_impl", TypeAttr: map[string]any{"optional": bool(true)}},
  		&{Name: "State", Desc: "Indicates the state of the issues to return.", Default: "open", TypeName: "essentials.model.mo_string.select_string_internal", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github profile`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github release asset download`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{Name: "Path", Desc: "Path to download", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github release asset list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github release asset upload`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Asset", Desc: "Path to assets", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "Release tag name", TypeName: "string"},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github release draft`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Name", Desc: "Name of the release", TypeName: "string"},
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  		&{Name: "Tag", Desc: "Name of the tag", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github release list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Repository owner", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "Repository name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `github tag create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_repo"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Owner", Desc: "Owner of the repository", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_repo",
- 			TypeAttr: string("github_repo"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Repository", Desc: "Name of the repository", TypeName: "string"},
  		&{Name: "Sha1", Desc: "SHA1 hash of the commit", TypeName: "string"},
  		&{Name: "Tag", Desc: "Tag name", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util monitor client`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
# Command spec changed: `util release install`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github"},
  	Services:        {"github"},
  	IsSecret:        false,
  	... // 5 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "AcceptLicenseAgreement", Desc: "Accept to the target release's license agreement", Default: "false", TypeName: "bool", ...},
  		&{Name: "Path", Desc: "Path to install", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.github.api.gh_conn_impl.conn_github_public",
- 			TypeAttr: string("github_public"),
+ 			TypeAttr: string("github"),
  		},
  		&{Name: "Release", Desc: "Release tag name", Default: "latest", TypeName: "string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util tidy pack remote`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 8 identical fields
  	ConnUseBusiness: false,
  	ConnScopes:      {"Peer": "dropbox_individual"},
- 	Services:        []string{"dropbox"},
+ 	Services:        []string{"dropbox_individual"},
  	IsSecret:        false,
  	IsConsole:       false,
  	... // 11 identical fields
  }
```
