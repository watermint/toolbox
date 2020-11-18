# Experimental features

The experimental feature switch is for testing or accessing early access features. You can enable those features with the option `-experiment`. If you want to specify multiple features, please select those features joined with a comma. (e.g. `-experiment feature1,feature2`).

| name                                   | Description                                                                                                |
|----------------------------------------|------------------------------------------------------------------------------------------------------------|
| dbx_client_conditioner_narrow20        | Simulate rate limit errors. 20% of requests will fail with rate limitation.                                |
| dbx_client_conditioner_narrow40        | Simulate rate limit errors. 40% of requests will fail with rate limitation.                                |
| dbx_client_conditioner_narrow100       | Simulate rate limit errors. 100% of requests will fail with rate limitation.                               |
| dbx_client_conditioner_error20         | Simulate server errors. 20% of requests will fail with a server error.                                     |
| dbx_client_conditioner_error40         | Simulate server errors. 40% of requests will fail with a server error.                                     |
| dbx_client_conditioner_error100        | Simulate server errors. 100% of requests will fail with a server error.                                    |
| batch_sequential                       | Execute batch sequentially in same batchId.                                                                |
| batch_random                           | Execute batch with random batchId order.                                                                   |
| congestion_window_no_limit             | Do not limit concurrency with the congestion window.                                                       |
| congestion_window_aggressive           | Apply aggressive initial congestion window size                                                            |
| file_sync_disable_reduce_create_folder | Disable reduce create_folder on syncing file systems. That will create empty folder while syncing folders. |

