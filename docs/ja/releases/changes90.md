---
layout: release
title: リリースの変更点 89
lang: ja
---

# `リリース 89` から `リリース 90` までの変更点

# コマンド仕様の変更: `dev build doc`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Badge", Desc: "ビルド状態のバッジを含める", Default: "true", TypeName: "bool", ...},
- 		&{
- 			Name:     "CommandPath",
- 			Desc:     "コマンドマニュアルを作成する相対パス",
- 			Default:  "doc/generated/",
- 			TypeName: "string",
- 		},
  		&{Name: "DocLang", Desc: "言語", TypeName: "essentials.model.mo_string.opt_string"},
- 		&{
- 			Name:     "Readme",
- 			Desc:     "README のファイル名",
- 			Default:  "README.md",
- 			TypeName: "string",
- 		},
- 		&{
- 			Name:     "Security",
- 			Desc:     "SECURITY_AND_PRIVACYのファイル名",
- 			Default:  "SECURITY_AND_PRIVACY.md",
- 			TypeName: "string",
- 		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `dev stage http_range`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	Name: "http_range",
  	Title: strings.Join({
  		"HTTP",
- 		" Range request proof of concept",
+ 		"レンジリクエストのプルーフオブコンセプト",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "DropboxPath",
- 			Desc:     "Dropbox file path to download",
+ 			Desc:     "ダウンロードするDropboxファイルのパス",
  			Default:  "",
  			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "LocalPath",
- 			Desc:     "Local path to store",
+ 			Desc:     "保存先のローカルパス",
  			Default:  "",
  			TypeName: "essentials.model.mo_path.file_system_path_impl",
  			TypeAttr: map[string]any{"shouldExist": bool(false)},
  		},
  		&{
  			Name:     "Peer",
- 			Desc:     "Account alias",
+ 			Desc:     "アカウントの別名",
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual",
  			TypeAttr: []any{string("files.content.read")},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `member file permdelete`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	Name: "permdelete",
  	Title: strings.Join({
  		"チームメンバーの指定したパスのファイルまた\xe3",
  		"\x81\xafフォルダを完全に削除します",
- 		"完全に削除については、https://www.dropbox.com/help/40",
- 		" をご覧ください.",
  	}, ""),
- 	Desc:    "",
+ 	Desc:    "完全に削除については、https://www.dropbox.com/help/40 をご覧ください.",
  	Remarks: "(試験的実装かつ非可逆な操作です)",
  	Path:    "member file permdelete",
  	... // 18 identical fields
  }
```
# コマンド仕様の変更: `services google mail sendas add`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	Name: "add",
  	Title: strings.Join({
- 		`Creates a custom "from" send-as alias`,
+ 		`カスタムの "from" send-asエイリアスの作成`,
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name: "DisplayName",
  			Desc: strings.Join({
- 				`A name that appears in the "From:" header for mail sent using th`,
- 				"is alias",
+ 				`このエイリアスを使って送信されるメールの "Fr`,
+ 				`om: "ヘッダーに表示される名前`,
  				".",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Email",
- 			Desc:     "The send-as alias email address",
+ 			Desc:     "send-asエイリアス メールアドレス",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Peer",
- 			Desc:     "Account alias",
+ 			Desc:     "アカウントの別名",
  			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_google_mail",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/gmail.settings.sharing")},
  		},
  		&{
  			Name: "ReplyTo",
  			Desc: strings.Join({
- 				`An optional email address that is included in a "Reply-To:" head`,
- 				"er for mail sent using this alias",
+ 				`このエイリアスを使って送信されたメールの "Re`,
+ 				"ply-To: \"ヘッダーに含まれる任意のメールアドレ\xe3",
+ 				"\x82\xb9",
  				".",
  			}, ""),
  			Default:  "",
  			TypeName: "essentials.model.mo_string.opt_string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "SkipVerify",
- 			Desc:     "Skip verify ",
+ 			Desc:     "検証をスキップする ",
  			Default:  "false",
  			TypeName: "bool",
  			TypeAttr: nil,
  		},
  		&{
  			Name: "UserId",
  			Desc: strings.Join({
- 				"The user's email address. The special value me can be used to in",
- 				"dicate the authenticated user",
+ 				"ユーザーのメールアドレス. 特別な値meは、認証",
+ 				"されたユーザを示すために使用することができ\xe3",
+ 				"\x81\xbeす",
  				".",
  			}, ""),
  			Default:  "me",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## 変更されたレポート: send_as

```
  &dc_recipe.Report{
  	Name: "send_as",
  	Desc: strings.Join({
- 		"Settings associated with a send-as alias, which can be either th",
- 		"e primary login address associated with the account or a custom ",
- 		`"from" address.`,
+ 		"送信先のエイリアスに関連する設定で、アカウ\xe3",
+ 		"\x83\xb3トに関連付けられたプライマリログインアド\xe3\x83",
+ 		"\xacスまたはカスタム\"from\"アドレスのいずれかを指",
+ 		"定できます",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		&{
  			Name: "send_as_email",
- 			Desc: "The send-as alias email address",
+ 			Desc: "send-asエイリアス メールアドレス",
  		},
  		&{
  			Name: "display_name",
  			Desc: strings.Join({
- 				`A name that appears in the "From:" header for mail sent using th`,
- 				"is alias",
+ 				`このエイリアスを使って送信されるメールの "Fr`,
+ 				`om: "ヘッダーに表示される名前`,
  				".",
  			}, ""),
  		},
  		&{
  			Name: "reply_to_address",
  			Desc: strings.Join({
- 				`An optional email address that is included in a "Reply-To:" head`,
- 				"er for mail sent using this alias",
+ 				`このエイリアスを使って送信されたメールの "Re`,
+ 				"ply-To: \"ヘッダーに含まれる任意のメールアドレ\xe3",
+ 				"\x82\xb9",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "signature",
  			Desc: strings.Join({
- 				"An optional HTML signature that is included in messages composed",
- 				" with this alias in the Gmail web UI",
+ 				"GmailのウェブUIでこのエイリアスを使って作成さ",
+ 				"れたメッセージに含まれるオプションのHTML署名",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "is_primary",
  			Desc: strings.Join({
- 				"Whether this address is the primary address used to login to the",
- 				" account",
+ 				"このアドレスが、アカウントへのログインに使\xe7",
+ 				"\x94\xa8されるプライマリアドレスであるかどうか",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "is_default",
  			Desc: strings.Join({
- 				`Whether this address is selected as the default "From:" address `,
- 				"in situations such as composing a new message or sending a vacat",
- 				"ion auto-reply.",
+ 				"新規メッセージの作成やバケーションの自動返\xe4",
+ 				"\xbf\xa1などの際に、このアドレスをデフォルトの\"From",
+ 				`:"アドレスとして選択するかどうか。`,
  			}, ""),
  		},
  		&{
  			Name: "treat_as_alias",
  			Desc: strings.Join({
- 				"Whether Gmail should treat this address as an alias for the user",
- 				"'s primary email address",
+ 				"Gmailがこのアドレスをユーザーのプライマリメ\xe3\x83",
+ 				"\xbcルアドレスのエイリアスとして扱うかどうかを",
+ 				"指定します",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "verification_status",
  			Desc: strings.Join({
- 				"Indicates whether this address has been verified for use as a se",
- 				"nd-as alias",
+ 				"このアドレスがsend-as aliasとして使用するために",
+ 				"検証されているかどうかを示す.",
  			}, ""),
  		},
  	},
  }
```
# コマンド仕様の変更: `services google mail sendas delete`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	Name: "delete",
  	Title: strings.Join({
- 		"Deletes the specified send-as alias",
+ 		"指定したsend-asエイリアスを削除する",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Email",
- 			Desc:     "The send-as alias email address",
+ 			Desc:     "send-asエイリアス メールアドレス",
  			Default:  "",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  		&{
  			Name:     "Peer",
- 			Desc:     "Account alias",
+ 			Desc:     "アカウントの別名",
  			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_google_mail",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/gmail.settings.sharing")},
  		},
  		&{
  			Name: "UserId",
  			Desc: strings.Join({
- 				"The user's email address. The special value me can be used to in",
- 				"dicate the authenticated user",
+ 				"ユーザーのメールアドレス. 特別な値meは、認証",
+ 				"されたユーザを示すために使用することができ\xe3",
+ 				"\x81\xbeす",
  				".",
  			}, ""),
  			Default:  "me",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```
# コマンド仕様の変更: `services google mail sendas list`



## 設定が変更されたコマンド

```
  &dc_recipe.Recipe{
  	Name: "list",
  	Title: strings.Join({
- 		"Lists the send-as aliases for the specified account",
+ 		"指定されたアカウントの送信エイリアスを一覧\xe8",
+ 		"\xa1\xa8示する",
  	}, ""),
  	Desc:    "",
  	Remarks: "",
  	... // 12 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Peer",
- 			Desc:     "Account alias",
+ 			Desc:     "アカウントの別名",
  			Default:  "default",
  			TypeName: "domain.google.api.goog_conn_impl.conn_google_mail",
  			TypeAttr: []any{string("https://www.googleapis.com/auth/gmail.readonly")},
  		},
  		&{
  			Name: "UserId",
  			Desc: strings.Join({
- 				"The user's email address. The special value me can be used to in",
- 				"dicate the authenticated user",
+ 				"ユーザーのメールアドレス. 特別な値meは、認証",
+ 				"されたユーザを示すために使用することができ\xe3",
+ 				"\x81\xbeす",
  				".",
  			}, ""),
  			Default:  "me",
  			TypeName: "string",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

## 変更されたレポート: send_as

```
  &dc_recipe.Report{
  	Name: "send_as",
  	Desc: strings.Join({
- 		"Settings associated with a send-as alias, which can be either th",
- 		"e primary login address associated with the account or a custom ",
- 		`"from" address.`,
+ 		"送信先のエイリアスに関連する設定で、アカウ\xe3",
+ 		"\x83\xb3トに関連付けられたプライマリログインアド\xe3\x83",
+ 		"\xacスまたはカスタム\"from\"アドレスのいずれかを指",
+ 		"定できます",
  	}, ""),
  	Columns: []*dc_recipe.ReportColumn{
  		&{
  			Name: "send_as_email",
- 			Desc: "The send-as alias email address",
+ 			Desc: "send-asエイリアス メールアドレス",
  		},
  		&{
  			Name: "display_name",
  			Desc: strings.Join({
- 				`A name that appears in the "From:" header for mail sent using th`,
- 				"is alias",
+ 				`このエイリアスを使って送信されるメールの "Fr`,
+ 				`om: "ヘッダーに表示される名前`,
  				".",
  			}, ""),
  		},
  		&{
  			Name: "reply_to_address",
  			Desc: strings.Join({
- 				`An optional email address that is included in a "Reply-To:" head`,
- 				"er for mail sent using this alias",
+ 				`このエイリアスを使って送信されたメールの "Re`,
+ 				"ply-To: \"ヘッダーに含まれる任意のメールアドレ\xe3",
+ 				"\x82\xb9",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "is_primary",
  			Desc: strings.Join({
- 				"Whether this address is the primary address used to login to the",
- 				" account",
+ 				"このアドレスが、アカウントへのログインに使\xe7",
+ 				"\x94\xa8されるプライマリアドレスであるかどうか",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "is_default",
  			Desc: strings.Join({
- 				`Whether this address is selected as the default "From:" address `,
- 				"in situations such as composing a new message or sending a vacat",
- 				"ion auto-reply.",
+ 				"新規メッセージの作成やバケーションの自動返\xe4",
+ 				"\xbf\xa1などの際に、このアドレスをデフォルトの\"From",
+ 				`:"アドレスとして選択するかどうか。`,
  			}, ""),
  		},
  		&{
  			Name: "treat_as_alias",
  			Desc: strings.Join({
- 				"Whether Gmail should treat this address as an alias for the user",
- 				"'s primary email address",
+ 				"Gmailがこのアドレスをユーザーのプライマリメ\xe3\x83",
+ 				"\xbcルアドレスのエイリアスとして扱うかどうかを",
+ 				"指定します",
  				".",
  			}, ""),
  		},
  		&{
  			Name: "verification_status",
  			Desc: strings.Join({
- 				"Indicates whether this address has been verified for use as a se",
- 				"nd-as alias",
+ 				"このアドレスがsend-as aliasとして使用するために",
+ 				"検証されているかどうかを示す.",
  			}, ""),
  		},
  	},
  }
```
