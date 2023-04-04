---
layout: contributor
title: Contributor guide
lang: en
---

# Recipe value types

| Implementation                                                                     | Conn  | Conns | CustomValueText | ErrorHandler | Feed  | GridDataInput | GridDataOutput | JsonInput | Message | Messages | Report | Reports | TextInput |
|------------------------------------------------------------------------------------|-------|-------|-----------------|--------------|-------|---------------|----------------|-----------|---------|----------|--------|---------|-----------|
| github.com/watermint/toolbox/infra/ui/app_msg.messageImpl                          | false | false | false           | false        | false | false         | false          | false     | true    | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/asana/api/as_conn_impl.connAsanaApi            | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| bool                                                                               | false | false | false           | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/infra/data/da_griddata.gdInput                        | false | false | true            | false        | false | true          | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/infra/data/da_griddata.gdOutput                       | false | false | true            | false        | false | false         | true           | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/infra/data/da_json.jsInput                            | false | false | true            | false        | false | false         | false          | true      | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/infra/data/da_text.txInput                            | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | true      |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedIndividual | true  | false | true            | true         | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedTeam       | true  | false | true            | true         | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/infra/feed/fd_file_impl.RowFeed                       | false | false | true            | false        | true  | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/figma/api/fg_conn_impl.connFigmaApi            | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubPublic       | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubRepo         | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/google/api/goog_conn_impl.connGoogleCalendar   | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/google/api/goog_conn_impl.connGoogleMail       | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/google/api/goog_conn_impl.connSheets           | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn_impl.connHelloSignApi  | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| int64                                                                              | false | false | false           | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/essentials/kvs/kv_storage_impl.proxyImpl              | false | false | false           | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_filter.filterImpl                 | false | false | false           | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/model/mo_path.dropboxPathImpl          | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_path.fileSystemPathImpl           | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/model/mo_time.TimeImpl                 | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/dropbox/model/mo_url.urlImpl                   | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_string.optString                  | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_int.rangeInt                      | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/infra/recipe/rc_value.EmptyRecipe                     | false | true  | false           | false        | false | false         | false          | false     | false   | true     | false  | true    | false     |
| github.com/watermint/toolbox/infra/report/rp_model_impl.RowReport                  | false | false | false           | false        | false | false         | false          | false     | false   | false    | true   | false   | false     |
| github.com/watermint/toolbox/infra/report/rp_model_impl.TransactionReport          | false | false | false           | false        | false | false         | false          | false     | false   | false    | true   | false   | false     |
| github.com/watermint/toolbox/essentials/model/mo_string.selectString               | false | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| github.com/watermint/toolbox/domain/slack/api/work_conn_impl.connSlackApi          | true  | false | true            | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |
| string                                                                             | false | false | false           | false        | false | false         | false          | false     | false   | false    | false  | false   | false     |

## Connection value types

| Implementation                                                                     | CustomValueText | Service name     | Scope label        |
|------------------------------------------------------------------------------------|-----------------|------------------|--------------------|
| github.com/watermint/toolbox/domain/asana/api/as_conn_impl.connAsanaApi            | true            | asana            | asana              |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedIndividual | true            | dropbox          | dropbox_individual |
| github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl.connScopedTeam       | true            | dropbox_business | dropbox_team       |
| github.com/watermint/toolbox/domain/figma/api/fg_conn_impl.connFigmaApi            | true            | figma            | figma              |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubPublic       | true            | github           | github_public      |
| github.com/watermint/toolbox/domain/github/api/gh_conn_impl.ConnGithubRepo         | true            | github           | github_repo        |
| github.com/watermint/toolbox/domain/google/api/goog_conn_impl.connGoogleCalendar   | true            | google_calendar  | google_calendar    |
| github.com/watermint/toolbox/domain/google/api/goog_conn_impl.connGoogleMail       | true            | google_mail      | google_mail        |
| github.com/watermint/toolbox/domain/google/api/goog_conn_impl.connSheets           | true            | google_sheets    | google_sheets      |
| github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn_impl.connHelloSignApi  | true            | dropbox_sign     | dropbox_sign       |
| github.com/watermint/toolbox/domain/slack/api/work_conn_impl.connSlackApi          | true            | slack            | slack              |


