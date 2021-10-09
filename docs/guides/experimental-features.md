---
layout: page
title: Experimental features
lang: en
---

# Experimental features

The experimental feature switch is for testing or accessing early access features. You can enable those features with the option `-experiment`. If you want to specify multiple features, please select those features joined with a comma. (e.g. `-experiment feature1,feature2`).

| name                                   | Description                                                                                                |
|----------------------------------------|------------------------------------------------------------------------------------------------------------|
| batch_balance                          | Execute batch from the largest batch                                                                       |
| batch_random                           | Execute batch with random batchId order.                                                                   |
| batch_sequential                       | Execute batch sequentially in same batchId.                                                                |
| congestion_window_aggressive           | Apply aggressive initial congestion window size                                                            |
| congestion_window_no_limit             | Do not limit concurrency with the congestion window.                                                       |
| dbx_client_conditioner_error100        | Simulate server errors. 100% of requests will fail with a server error.                                    |
| dbx_client_conditioner_error20         | Simulate server errors. 20% of requests will fail with a server error.                                     |
| dbx_client_conditioner_error40         | Simulate server errors. 40% of requests will fail with a server error.                                     |
| dbx_client_conditioner_narrow100       | Simulate rate limit errors. 100% of requests will fail with rate limitation.                               |
| dbx_client_conditioner_narrow20        | Simulate rate limit errors. 20% of requests will fail with rate limitation.                                |
| dbx_client_conditioner_narrow40        | Simulate rate limit errors. 40% of requests will fail with rate limitation.                                |
| file_sync_disable_reduce_create_folder | Disable reduce create_folder on syncing file systems. That will create empty folder while syncing folders. |
| legacy_local_to_dbx_connector          | Use legacy local to dropbox sync connector                                                                 |
| use_no_cache_dbxfs                     | Use non-cache dropbox file system                                                                          |
| profile_cpu                            | Enable CPU profiler                                                                                        |
| profile_memory                         | Enable memory profiler                                                                                     |
| report_all_columns                     | Show all columns defined as data structure.                                                                |


