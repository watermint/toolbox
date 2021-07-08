---
layout: release
title: Changes of Release 91
lang: en
---

# Changes between `Release 91` to `Release 92`

# Commands added


| Command                        | Title                                          |
|--------------------------------|------------------------------------------------|
| team sharedlink cap expiry     | Set expiry cap to shared links in the team     |
| team sharedlink cap visibility | Set visibility cap to shared links in the team |



# Commands deleted


| Command                | Title                                 |
|------------------------|---------------------------------------|
| connect business_audit | Connect to the team audit access      |
| connect business_file  | Connect to the team file access       |
| connect business_info  | Connect to the team info access       |
| connect business_mgmt  | Connect to the team management access |
| connect user_file      | Connect to user file access           |
| dev ci auth export     | Export auth tokens of end to end test |
| team diag explorer     | Report whole team information         |



# Command spec changed: `dev build preflight`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports:        nil,
  	Feeds:          nil,
- 	Values:         []*dc_recipe.Value{},
+ 	Values:         []*dc_recipe.Value{&{Name: "Quick", Desc: "Quick mode", Default: "false", TypeName: "bool"}},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `dev ci auth connect`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 5 identical fields
  	CliArgs:         "",
  	CliNote:         "",
- 	ConnUsePersonal: true,
+ 	ConnUsePersonal: false,
- 	ConnUseBusiness: true,
+ 	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Audit":  "business_audit",
- 		"File":   "business_file",
- 		"Full":   "user_full",
  		"Github": "github_repo",
- 		"Info":   "business_info",
- 		"Mgmt":   "business_management",
  	},
  	Services: []string{
- 		"dropbox",
- 		"dropbox_business",
  		"github",
  	},
  	IsSecret:  true,
  	IsConsole: false,
  	... // 3 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Audit",
- 			Desc:     "Authenticate with Dropbox Business Audit scope",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_audit",
- 		},
- 		&{
- 			Name:     "File",
- 			Desc:     "Authenticate with Dropbox Business member file access scope",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
- 		},
- 		&{
- 			Name:     "Full",
- 			Desc:     "Authenticate with Dropbox user full access scope",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_user_file",
- 		},
  		&{Name: "Github", Desc: "Account alias for Github deployment", Default: "deploy", TypeName: "domain.github.api.gh_conn_impl.conn_github_repo", ...},
- 		&{
- 			Name:     "Info",
- 			Desc:     "Authenticate with Dropbox Business info scope",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_info",
- 		},
- 		&{
- 			Name:     "Mgmt",
- 			Desc:     "Authenticate with Dropbox Business management scope",
- 			Default:  "end_to_end_test",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt",
- 		},
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
  	... // 5 identical fields
  	CliArgs:         "",
  	CliNote:         "",
- 	ConnUsePersonal: true,
+ 	ConnUsePersonal: false,
- 	ConnUseBusiness: true,
+ 	ConnUseBusiness: false,
  	ConnScopes: map[string]string{
- 		"Audit":  "business_audit",
- 		"File":   "business_file",
- 		"Full":   "user_full",
  		"Github": "github_repo",
- 		"Info":   "business_info",
- 		"Mgmt":   "business_management",
  		"Peer":   "github_public",
  	},
  	Services: []string{
- 		"dropbox",
- 		"dropbox_business",
  		"github",
  	},
  	IsSecret:  true,
  	IsConsole: true,
  	... // 10 identical fields
  }
```
# Command spec changed: `team sharedlink list`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 7 identical fields
  	ConnUsePersonal: false,
  	ConnUseBusiness: true,
- 	ConnScopes:      map[string]string{"Peer": "business_file"},
+ 	ConnScopes:      map[string]string{"Peer": "dropbox_scoped_team"},
  	Services:        {"dropbox_business"},
  	IsSecret:        false,
  	... // 4 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
  			Desc:     "Account alias",
  			Default:  "default",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file",
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: nil,
+ 			TypeAttr: []interface{}{string("members.read"), string("sharing.read"), string("team_data.member")},
  		},
  		&{Name: "Visibility", Desc: "Filter links by visibility (all/public/team_only/password)", Default: "all", TypeName: "essentials.model.mo_string.select_string", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# Command spec changed: `team sharedlink update expiry`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:  "expiry",
  	Title: "Update expiration date of public shared links within the team",
  	Desc: (
  		"""
  		Note: From Release 87, this command will receive a file to select shared links to update. If you wanted to update the expiry for all shared links in the team, please consider using a combination of `team sharedlink list`. For example, if you are familiar with the command [jq](https://stedolan.github.io/jq/), then you can do an equivalent operation as like below (force expiry within 28 days for every public link).
  		
  		```
- 		tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -days 28
+ 		tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -at +720h
  		```
+ 		From Release 92, the command will not receive the argument `-days`. If you want to set a relative date/time, please use `-at +HOURh` like `+720h` (720 hours = 30 days).
  		
+ 		Commands `team sharedlink update` is for setting a value to the shared links. Commands `team sharedlink cap` is for putting a cap value to the shared links. For example: if you set expiry by `team sharedlink update expiry` with the expiration date 2021-05-06. The command will update the expiry to 2021-05-06 even if the existing link has a shorter expiration date like 2021-05-04.
  		"""
  	),
  	Remarks:         "(Irreversible operation)",
  	Path:            "team sharedlink update expiry",
- 	CliArgs:         "-file /PATH/TO/DATA_FILE.csv -days 28",
+ 	CliArgs:         "-file /PATH/TO/DATA_FILE.csv -at +720h",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			... // 2 identical fields
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_time.time_impl",
- 			TypeAttr: map[string]interface{}{"optional": bool(true)},
+ 			TypeAttr: map[string]interface{}{"optional": bool(false)},
  		},
- 		&{
- 			Name:     "Days",
- 			Desc:     "Days to the new expiration date",
- 			Default:  "0",
- 			TypeName: "essentials.model.mo_int.range_int",
- 			TypeAttr: map[string]interface{}{"max": float64(2.147483647e+09), "min": float64(0), "value": float64(0)},
- 		},
  		&{Name: "File", Desc: "Path to data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
