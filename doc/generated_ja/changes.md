# `リリース 76` から `リリース 77` までの変更点

# コマンド仕様の変更: `dev stage teamfolder`


## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
- 		&{
- 			Name:     "Peer",
- 			Desc:     "Account alias",
- 			Default:  "&{Peer [groups.write files.content.write] <nil>}",
- 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
- 			TypeAttr: []interface{}{string("groups.write"), string("files.content.write")},
- 		},
+ 		&{
+ 			Name:     "Peer",
+ 			Desc:     "Account alias",
+ 			Default:  "&{Peer [files.content.read files.content.write groups.write sharing.read sharing.write team_data.member team_data.team_space tea"...,
+ 			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
+ 			TypeAttr: []interface{}{
+ 				string("files.content.read"), string("files.content.write"),
+ 				string("groups.write"), string("sharing.read"), string("sharing.write"),
+ 				string("team_data.member"), string("team_data.team_space"),
+ 				string("team_info.read"),
+ 			},
+ 		},
  	},
  }
```
