---
layout: page
title: レポートオプション
lang: ja
---

# レポートオプション

watermint toolboxは、API経由でサービスから取得したデータからレポートを作成する。レポートはコマンドによって異なる。

* APIからのデータをレポートとして出力するコマンド。
* APIからのデータを処理後にレポートとして出力するコマンド。

コマンドの出力形式は、コアなユースケースのために設計されている。混乱を避けるため、このコマンドでは関連性のない/優先順位の低いフィールドは省略される。
以下に示す方法で省略データを出力するか、出力フィルタを使用してお好みの形式でレポートを作成することができます。

# 隠しカラム

コマンドによって作成されるCSVおよびxlsxレポートでは、一部の列が省略されることがあります。これには、たとえば内部で使用されているIDや関連性の低いデータなどが含まれる。

```
$ ./tbx dropbox file account info

watermint toolbox xxx.x.xxx
===========================

© 2016-2024 Takayuki Okazaki
Licensed under open source licenses. 詳細は`license`コマンドでご覧ください.

| email                 | email_verified | given_name | surname | display_name |
|-----------------------|----------------|------------|---------|--------------|
| xxx@xxxxxxxxxxxxx.xxx | true           | xxxx       | xxxx    | xxxxxxxx     |
```

この種のデータを出力したい場合は、`-experiment report_all_columns' オプションを追加して、定義されたすべての列を出力することができます。

```
$ ./tbx dropbox file account info -experiment report_all_columns

watermint toolbox xxx.x.xxx
===========================

© 2016-2024 Takayuki Okazaki
Licensed under open source licenses. 詳細は`license`コマンドでご覧ください.

| team_member_id                            | email                 | email_verified | status | given_name | surname | familiar_name | display_name | abbreviated_name | member_folder_id | external_id | account_id                               | persistent_id | joined_on | invited_on | role | tag |
|-------------------------------------------|-----------------------|----------------|--------|------------|---------|---------------|--------------|------------------|------------------|-------------|------------------------------------------|---------------|-----------|------------|------|-----|
| dbmid:xxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx | xxx@xxxxxxxxxxxxx.xxx | true           |        | xxxx       | xxxx    | xxxxxxxx      | xxxxxxx     | xxxx             |                  |             | dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx |               |           |            |      |     |
```

このオプションを使用しても、一部の情報が出力されない場合があります。より詳細な情報が必要な場合は、以下の出力フィルターをお試しください。

# 出力フィルター・オプション

この機能により、コマンドの出力をフィルターすることができる。
これは、出力を特定のフォーマットで処理したい場合に便利である。

> 注：このフィルタはすべての出力レポートに適用されます。複数のレポート形式を出力するコマンドでは、意図したとおりに動作しません。
> 複数のレポート形式を処理する場合は、`util json query`コマンドを使用して、出力されたJSONファイルをそれぞれ処理してください。
> 

さらに、JSON形式のデータには、より多くのデータが含まれている場合もある。
そのような隠されたデータを取得したい場合、このオプションはレポートとして抽出するのに役立ちます。

例えば、[dropbox team member list](https://toolbox.watermint.org/commands/dropbox-team-member-list.html)というコマンドは、チームメンバーのリストを返す。
JSONレポートには、Dropbox APIからの生データが含まれています。
チームメンバーのEメールアドレスと認証ステータスのみを抽出したい場合は、出力フィルターオプションを使用できます。

```
$ ./tbx dropbox team member list -output json --output-filter "[.profile.email, .profile.email_verified]"
["sugito@example.com", true]
["kajiwara@example.com", true]
["takimoto@example.com", false]
["ueno@example.com", true]
["tomioka@example.com", false]
```

次に、このデータをCSVとしてフォーマットしたい場合は、次のように`@csv`フィルターを使用します（最後に`| @csv`を追加します）：

```
$ ./tbx dropbox team member list -output json --output-filter "[.profile.email, .profile.email_verified] | @csv"
"sugito@example.com",true
"kajiwara@example.com",true
"takimoto@example.com",false
"ueno@example.com",true
"tomioka@example.com",false
```

出力フィルタをテストしたい場合は、出力フィルタオプションなしでコマンドを実行することができます。
このコマンドは生のJSON出力を生成する。
それから、[util json query](https://toolbox.watermint.org/commands/util-json-query.html)コマンドでクエリーをテストすることができる。


