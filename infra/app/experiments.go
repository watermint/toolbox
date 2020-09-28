package app

const (
	ExperimentDbxClientConditionerNarrow20  = "dbx_client_conditioner_narrow20"  // 429 error for 20% traffic
	ExperimentDbxClientConditionerNarrow40  = "dbx_client_conditioner_narrow40"  // 429 error for 40% traffic
	ExperimentDbxClientConditionerNarrow100 = "dbx_client_conditioner_narrow100" // 429 error for 100% traffic

	ExperimentDbxClientConditionerError20  = "dbx_client_conditioner_error20"  // 500 error for 20% traffic
	ExperimentDbxClientConditionerError40  = "dbx_client_conditioner_error40"  // 500 error for 40% traffic
	ExperimentDbxClientConditionerError100 = "dbx_client_conditioner_error100" // 500 error for 100% traffic

	// Execute batch sequentially in same batchId
	ExperimentBatchSequential = "batch_sequential"
	// Execute batch with random batchId order
	ExperimentBatchRandom = "batch_random"

	ExperimentCongestionWindowNoLimit = "congestion_window_no_limit"

	ExperimentFileSyncReduceCreateFolder = "file_sync_reduce_create_folder"
)
