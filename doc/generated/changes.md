# Changes between `Release 67` to `Release 68`

# Commands added


| Command                     | Title                                      |
|-----------------------------|--------------------------------------------|
| dev util image jpeg         | Create dummy image files                   |
| services github content get | Get content metadata of the repository     |
| services github content put | Put small text content into the repository |



# Commands deleted


| Command   | Title              |
|-----------|--------------------|
| dev dummy | Create dummy files |



# Command spec changed: `member delete`


## Command configuration changed


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "File", Desc: "Data file", TypeName: "infra.feed.fd_file_impl.row_feed"},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_business_mgmt"},
+ 		&{
+ 			Name:     "TransferDestMember",
+ 			Desc:     "If provided, files from the deleted member account will be transferred to this user.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
+ 		&{
+ 			Name:     "TransferNotifyAdminEmailOnError",
+ 			Desc:     "If provided, errors during the transfer process will be sent via email to this user.",
+ 			TypeName: "domain.common.model.mo_string.opt_string",
+ 		},
  		&{Name: "WipeData", Desc: "If true, controls if the user's data will be deleted on their linked devices", Default: "true", TypeName: "bool"},
  	},
  }
```
