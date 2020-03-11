# `リリース 59` から `最新リリース` までの変更点

# 追加されたコマンド

| コマンド                  | タイトル                                         |
|---------------------------|--------------------------------------------------|
| dev ci artifact up        | Upload CI artifact                               |
| dev ci auth               | Authenticate for generating end to end testing   |
| dev util curl             | Generate cURL preview from capture log           |
| filerequest create        | Create a file request                            |
| filerequest delete closed | Delete all closed file requests on this account. |
| filerequest delete url    | Delete a file request by the file request URL    |
| filerequest list          | List file requests of the individual account     |



# 削除されたコマンド

| コマンド      | タイトル                                       |
|---------------|------------------------------------------------|
| dev test auth | Authenticate for generating end to end testing |



# コマンド仕様の変更: `connect business_audit`



## 設定が変更されたコマンド



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

# コマンド仕様の変更: `connect business_file`



## 設定が変更されたコマンド



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

# コマンド仕様の変更: `connect business_info`



## 設定が変更されたコマンド



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

# コマンド仕様の変更: `connect business_mgmt`



## 設定が変更されたコマンド



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

# コマンド仕様の変更: `connect user_file`



## 設定が変更されたコマンド



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

