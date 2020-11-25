# Changes between `Release 78` to `Release 79`

# Commands added


| Command       | Title                |
|---------------|----------------------|
| dev stage gui | GUI proof of concept |



# Command spec changed: `team sharedlink list`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_file", ...},
  		&{
  			Name:     "Visibility",
  			Desc:     "Filter links by visibility (public/team_only/password)",
- 			Default:  "public",
+ 			Default:  "all",
  			TypeName: "essentials.model.mo_string.select_string",
  			TypeAttr: map[string]interface{}{
  				"options": []interface{}{
+ 					string("all"),
  					string("public"),
  					string("team_only"),
  					... // 3 identical elements
  				},
  			},
  		},
  	},
  }
```
