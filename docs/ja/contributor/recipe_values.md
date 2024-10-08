---
layout: contributor
title: コントリビューターガイド
lang: ja
---

# レシピの値の種類

| Implementation                                                                     | Conn  | Conns | CustomValueText | ErrorHandler | Feed  | GridDataInput | GridDataOutput | JsonInput | Message | Messages | レポート | Reports | TextInput |
|------------------------------------------------------------------------------------|-------|-------|-----------------|--------------|-------|---------------|----------------|-----------|---------|----------|----------|---------|-----------|
| github.com/watermint/toolbox/infra/ui/app_msg.messageImpl                          | false | false | false           | false        | false | false         | false          | false     | true    | false    | false    | false   | false     |
| bool                                                                               | false | false | false           | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/infra/data/da_griddata.gdInput                        | false | false | true            | false        | false | true          | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/infra/data/da_griddata.gdOutput                       | false | false | true            | false        | false | false         | true           | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/infra/data/da_json.jsInput                            | false | false | true            | false        | false | false         | false          | true      | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/infra/data/da_text.txInput                            | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | true      |
| github.com/watermint/toolbox/infra/feed/fd_file_impl.RowFeed                       | false | false | true            | false        | true  | false         | false          | false     | false   | false    | false    | false   | false     |
| int64                                                                              | false | false | false           | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/essentials/kvs/kv_storage_impl.proxyImpl              | false | false | false           | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_filter.filterImpl                 | false | false | false           | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/model/mo_path.dropboxPathImpl          | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_path.fileSystemPathImpl           | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/model/mo_time.TimeImpl                 | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/model/mo_url.urlImpl                   | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_string.optString                  | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_int.rangeInt                      | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/infra/recipe/rc_value.EmptyRecipe                     | false | true  | false           | false        | false | false         | false          | false     | false   | true     | false    | true    | false     |
| github.com/watermint/toolbox/infra/report/rp_model_impl.RowReport                  | false | false | false           | false        | false | false         | false          | false     | false   | false    | true     | false   | false     |
| github.com/watermint/toolbox/infra/report/rp_model_impl.TransactionReport          | false | false | false           | false        | false | false         | false          | false     | false   | false    | true     | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_string.selectStringInternal       | false | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| string                                                                             | false | false | false           | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/asana/api/as_conn_impl.connAsanaApi            | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedIndividual | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedTeam       | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl.connDeeplApiImpl     | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/figma/api/fg_conn_impl.connFigmaApi            | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubPublic       | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubRepo         | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn_impl.connHelloSignApi  | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |
| github.com/watermint/toolbox/domain/slack/api/work_conn_impl.connSlackApi          | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false    | false   | false     |

## 接続値の種類

| Implementation                                                                     | CustomValueText | サービス名         | スコープラベル     |
|------------------------------------------------------------------------------------|-----------------|--------------------|--------------------|
| github.com/watermint/toolbox/domain/asana/api/as_conn_impl.connAsanaApi            | true            | asana              | asana              |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedIndividual | true            | dropbox_individual | dropbox_individual |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedTeam       | true            | dropbox_team       | dropbox_team       |
| github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl.connDeeplApiImpl     | true            | deepl              | deepl              |
| github.com/watermint/toolbox/domain/figma/api/fg_conn_impl.connFigmaApi            | true            | figma              | figma              |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubPublic       | true            | github_public      | github             |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubRepo         | true            | github_repo        | github             |
| github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn_impl.connHelloSignApi  | true            | dropbox_sign       | dropbox_sign       |
| github.com/watermint/toolbox/domain/slack/api/work_conn_impl.connSlackApi          | true            | slack              | slack              |


