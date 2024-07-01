---
layout: release
title: リリースの変更点 134
lang: ja
---

# `リリース 134` から `リリース 135` までの変更点

# 追加されたコマンド


| コマンド                       | タイトル               |
|--------------------------------|------------------------|
| dropbox file sharedfolder info | 共有フォルダ情報の取得 |



# コマンド仕様の変更: `google mail message send`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "send",
- 	Title:   "メールの送信",
+ 	Title:   `{"key":"citron.google.mail.message.send.title","params":{}}`,
  	Desc:    "",
  	Remarks: "(非可逆な操作です)",
  	... // 20 identical fields
  }
```
# コマンド仕様の変更: `google translate text`



## 設定が変更されたコマンド


```
  &dc_recipe.Recipe{
  	Name:    "text",
- 	Title:   "テキストを翻訳する",
+ 	Title:   `{"key":"citron.google.translate.text.title","params":{}}`,
  	Desc:    "",
  	Remarks: "",
  	... // 20 identical fields
  }
```
