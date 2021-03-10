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
	// Execute batch from the largest batch
	ExperimentBatchBalance = "batch_balance"

	// Do not hard limit window size
	ExperimentCongestionWindowNoLimit = "congestion_window_no_limit"
	// Aggressive initial window size
	ExperimentCongestionWindowAggressive = "congestion_window_aggressive"
	// Reduce call create_folder on sync
	ExperimentFileSyncDisableReduceCreateFolder = "file_sync_disable_reduce_create_folder"
	// Use legacy local to dropbox sync connector
	ExperimentFileSyncLegacyLocalToDropboxConnector = "legacy_local_to_dbx_connector"
	// Use non-cache dropbox file system
	ExperimentFileSyncNoCacheDropboxFileSystem = "use_no_cache_dbxfs"
)

var (
	ExperimentalFeatures = []string{
		ExperimentDbxClientConditionerNarrow20,
		ExperimentDbxClientConditionerNarrow40,
		ExperimentDbxClientConditionerNarrow100,
		ExperimentDbxClientConditionerError20,
		ExperimentDbxClientConditionerError40,
		ExperimentDbxClientConditionerError100,
		ExperimentBatchBalance,
		ExperimentBatchRandom,
		ExperimentBatchSequential,
		ExperimentCongestionWindowNoLimit,
		ExperimentCongestionWindowAggressive,
		ExperimentFileSyncDisableReduceCreateFolder,
		ExperimentFileSyncLegacyLocalToDropboxConnector,
		ExperimentFileSyncNoCacheDropboxFileSystem,
	}
)
