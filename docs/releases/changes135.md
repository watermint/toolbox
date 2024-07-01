---
layout: release
title: Changes of Release 134
lang: en
---

# Changes between `Release 134` to `Release 135`

# Commands added


| Command                        | Title                  |
|--------------------------------|------------------------|
| dropbox file sharedfolder info | Get shared folder info |



# Command spec changed: `google mail message send`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "send",
- 	Title:   "Send a mail",
+ 	Title:   `{"key":"citron.google.mail.message.send.title","params":{}}`,
  	Desc:    "",
  	Remarks: "(Irreversible operation)",
  	... // 20 identical fields
  }
```
# Command spec changed: `google translate text`



## Command configuration changed


```
  &dc_recipe.Recipe{
  	Name:    "text",
- 	Title:   "Translate text",
+ 	Title:   `{"key":"citron.google.translate.text.title","params":{}}`,
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
