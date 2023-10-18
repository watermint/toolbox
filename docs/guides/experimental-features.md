---
layout: page
title: Experimental features
lang: en
---

# Experimental features

The experimental feature switch is for testing or accessing early access features. You can enable those features with the option `-experiment`. If you want to specify multiple features, please select those features joined with a comma. (e.g. `-experiment feature1,feature2`).

| name                                   | Description                                                                                                                                                                                             |
|----------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| batch_balance                          | Execute batch from the largest batch                                                                                                                                                                    |
| batch_non_durable                      | Use non-durable batch framework                                                                                                                                                                         |
| batch_random                           | Execute batch with random batchId order.                                                                                                                                                                |
| batch_sequential                       | Execute batch sequentially in same batchId.                                                                                                                                                             |
| congestion_window_aggressive           | Apply aggressive initial congestion window size                                                                                                                                                         |
| congestion_window_no_limit             | Do not limit concurrency with the congestion window.                                                                                                                                                    |
| dbx_auth_course_grained_scope          | Requests all Dropbox authorization scopes instead of command-defined ones. This is used as a workaround in case the program does not work properly with the authorization scope defined in the command. |
| dbx_auth_redirect                      | Use redirect processing for authorization process to Dropbox                                                                                                                                            |
| dbx_client_conditioner_error100        | Simulate server errors. 100% of requests will fail with a server error.                                                                                                                                 |
| dbx_client_conditioner_error20         | Simulate server errors. 20% of requests will fail with a server error.                                                                                                                                  |
| dbx_client_conditioner_error40         | Simulate server errors. 40% of requests will fail with a server error.                                                                                                                                  |
| dbx_client_conditioner_narrow100       | Simulate rate limit errors. 100% of requests will fail with rate limitation.                                                                                                                            |
| dbx_client_conditioner_narrow20        | Simulate rate limit errors. 20% of requests will fail with rate limitation.                                                                                                                             |
| dbx_client_conditioner_narrow40        | Simulate rate limit errors. 40% of requests will fail with rate limitation.                                                                                                                             |
| dbx_download_block                     | Download file divide by blocks (improve concurrency)                                                                                                                                                    |
| file_sync_disable_reduce_create_folder | Disable reduce create_folder on syncing file systems. That will create empty folder while syncing folders.                                                                                              |
| legacy_local_to_dbx_connector          | Use legacy local to dropbox sync connector                                                                                                                                                              |
| use_no_cache_dbxfs                     | Use non-cache dropbox file system                                                                                                                                                                       |
| kvs_badger                             | Use Badger as KVS engine                                                                                                                                                                                |
| kvs_badger_turnstile                   | Use Badger as KVS engine with turnstile                                                                                                                                                                 |
| kvs_bitcask                            | Use Bitcask as KVS engine                                                                                                                                                                               |
| kvs_bitcask_turnstile                  | Use Bitcask as the key-value store with turnstile                                                                                                                                                       |
| kvs_sqlite                             | Use Sqlite3 as KVS engine                                                                                                                                                                               |
| kvs_sqlite_turnstile                   | Use SQLite as the key-value store with turnstile                                                                                                                                                        |
| profile_cpu                            | Enable CPU profiler                                                                                                                                                                                     |
| profile_memory                         | Enable memory profiler                                                                                                                                                                                  |
| report_all_columns                     | Show all columns defined as data structure.                                                                                                                                                             |
| suppress_progress                      | Suppress progress indicators                                                                                                                                                                            |


