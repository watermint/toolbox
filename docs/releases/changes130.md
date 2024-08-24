---
layout: release
title: Changes of Release 129
lang: en
---

# Changes between `Release 129` to `Release 130`

# Commands added


| Command                  | Title                          |
|--------------------------|--------------------------------|
| config license list      | List available license keys    |
| dev license issue        | Issue a license                |
| dev release announcement | Update announcements           |
| dev release checkin      | Check in the new release       |
| dev test license         | Testing license required logic |



# Commands deleted


| Command                       | Title                                       |
|-------------------------------|---------------------------------------------|
| dev test auth all             | Test for connect to Dropbox with all scopes |
| dev test setup massfiles      | Upload Wikimedia dump file as test file     |
| dev test setup teamsharedlink | Create demo shared links                    |



# Command spec changed: `dev release candidate`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "github_public"},
+ 	ConnScopes:      map[string]string{"Peer": "github_repo"},
  	Services:        {"github"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```

## Added report(s)


| Name          | Description  |
|---------------|--------------|
| announcements | Announcement |


# Command spec changed: `google calendar event list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_calendar"},
+ 	ConnScopes:      map[string]string{"Peer": "google_calendar2024"},
  	Services:        {"google_calendar"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail filter add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail filter batch add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail filter delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail filter list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail label add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail label delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail label list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail label rename`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail message label add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail message label delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail message list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail message processed list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail message send`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        true,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail sendas add`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail sendas delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail sendas list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google mail thread list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_mail"},
+ 	ConnScopes:      map[string]string{"Peer": "google_mail2024"},
  	Services:        {"google_mail"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet append`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet clear`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet delete`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet export`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet import`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets sheet list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `google sheets spreadsheet create`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: false,
- 	ConnScopes:      map[string]string{"Peer": "google_sheets"},
+ 	ConnScopes:      map[string]string{"Peer": "google_sheets2024"},
  	Services:        {"google_sheets"},
  	IsSecret:        false,
  	... // 12 identical fields
  }
```
# Command spec changed: `util desktop screenshot interval`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Count", Desc: "Number of screenshots to take. If the value is less than 1, the "..., Default: "-1", TypeName: "int", ...},
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(2),
+ 				"max":   float64(1),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Interval", Desc: "Interval seconds between screenshots.", Default: "10", TypeName: "int", ...},
  		&{Name: "NamePattern", Desc: "Name pattern of screenshot file. You can use the following place"..., Default: "{% raw %}{{.{% endraw %}Sequence}}_{% raw %}{{.{% endraw %}Timestamp}}.png", TypeName: "string", ...},
  		... // 2 identical elements
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `util desktop screenshot snap`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 17 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "0",
  			TypeName: "essentials.model.mo_int.range_int",
  			TypeAttr: map[string]any{
- 				"max":   float64(2),
+ 				"max":   float64(1),
  				"min":   float64(0),
  				"value": float64(0),
  			},
  		},
  		&{Name: "Path", Desc: "Path to save the screenshot", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
